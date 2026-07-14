package main

import (
    "bytes"
    "errors"
    "fmt"
    "io/fs"
    "os"
    "os/exec"
    "path/filepath"
    goruntime "runtime"
    "sort"
    "strconv"
    "strings"
)

// MatrixConfig controls the unified codec test matrix. The matrix deliberately
// lives beside the language-neutral scenario spec so test data, driver
// generation, generator options, compilation and execution stay in one place.
type MatrixConfig struct {
    Bubbler string
    Root    string
    Work    string
    Langs   string
    Profile string
    Keep    bool
}

type matrixVariant struct {
    Name  string
    Flags []string
}

type matrixTarget struct {
    Name       string
    BubblerID  string
    SkipEnv    string
    Tools      []string
    LoopMarker string
    Options    []matrixVariant
}

type matrixRuntime struct {
    root     string
    e2e      string
    bubbler  string
    work     string
    profile  string
    keep     bool
    selected map[string]bool
    spec     *Spec
}

var unrollThresholds = []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 31, 32}

var matrixTargets = []matrixTarget{
    {
        Name:       "c",
        BubblerID:  "c",
        SkipEnv:    "SKIP_C",
        Tools:      []string{"gcc"},
        LoopMarker: "for (int64_t _i",
        Options: []matrixVariant{
            {Name: "minimal", Flags: []string{"-minimal"}},
            {Name: "memcpy", Flags: []string{"-memcpy"}},
            {Name: "decnum", Flags: []string{"-decnum"}},
            {Name: "signext_arith", Flags: []string{"-signext=arith"}},
        },
    },
    {
        Name:       "cpp",
        BubblerID:  "cpp",
        SkipEnv:    "SKIP_CPP",
        Tools:      []string{"g++"},
        LoopMarker: "for (int64_t _i",
        Options: []matrixVariant{
            {Name: "minimal", Flags: []string{"-minimal"}},
            {Name: "memcpy", Flags: []string{"-memcpy"}},
            {Name: "decnum", Flags: []string{"-decnum"}},
            {Name: "signext_arith", Flags: []string{"-signext=arith"}},
            {Name: "compat", Flags: []string{"-compat"}},
            {Name: "single", Flags: []string{"-single"}},
            {Name: "single_compat", Flags: []string{"-single", "-compat"}},
        },
    },
    {
        Name:       "go",
        BubblerID:  "go",
        SkipEnv:    "SKIP_GO",
        Tools:      []string{"go"},
        LoopMarker: "for _i :=",
        Options: []matrixVariant{
            {Name: "minimal", Flags: []string{"-minimal"}},
            {Name: "memcpy", Flags: []string{"-memcpy"}},
            {Name: "decnum", Flags: []string{"-decnum"}},
            {Name: "signext_arith", Flags: []string{"-signext=arith"}},
        },
    },
    {
        Name:       "java",
        BubblerID:  "java",
        SkipEnv:    "SKIP_JAVA",
        Tools:      []string{"javac", "java"},
        LoopMarker: "for (int _i",
        Options: []matrixVariant{
            {Name: "decnum", Flags: []string{"-decnum"}},
        },
    },
    {
        Name:       "python",
        BubblerID:  "py",
        SkipEnv:    "SKIP_PY",
        Tools:      nil,
        LoopMarker: "for _i in range",
        Options: []matrixVariant{
            {Name: "decnum", Flags: []string{"-decnum"}},
            {Name: "signext_arith", Flags: []string{"-signext=arith"}},
        },
    },
    {
        Name:       "csharp",
        BubblerID:  "cs",
        SkipEnv:    "SKIP_CS",
        Tools:      []string{"dotnet"},
        LoopMarker: "for (int _i",
        Options: []matrixVariant{
            {Name: "memcpy", Flags: []string{"-memcpy"}},
            {Name: "decnum", Flags: []string{"-decnum"}},
        },
    },
    {
        Name:       "cjs",
        BubblerID:  "cjs",
        SkipEnv:    "SKIP_CJS",
        Tools:      []string{"node"},
        LoopMarker: "for (let _i",
        Options: []matrixVariant{
            {Name: "compat", Flags: []string{"-compat"}},
            {Name: "decnum", Flags: []string{"-decnum"}},
        },
    },
    {
        Name:       "esm",
        BubblerID:  "mjs",
        SkipEnv:    "SKIP_ESM",
        Tools:      []string{"node"},
        LoopMarker: "for (let _i",
        Options: []matrixVariant{
            {Name: "compat", Flags: []string{"-compat"}},
            {Name: "decnum", Flags: []string{"-decnum"}},
        },
    },
}

func RunMatrix(config MatrixConfig) error {
    runtime, err := newMatrixRuntime(config)
    if err != nil {
        return err
    }
    if !runtime.keep {
        defer os.RemoveAll(runtime.work)
    }
    if err := os.RemoveAll(runtime.work); err != nil {
        return fmt.Errorf("clean matrix work directory: %w", err)
    }
    if err := os.MkdirAll(runtime.work, 0o755); err != nil {
        return fmt.Errorf("create matrix work directory: %w", err)
    }

    fmt.Printf("[matrix] root:     %s\n", runtime.root)
    fmt.Printf("[matrix] bubbler: %s\n", runtime.bubbler)
    fmt.Printf("[matrix] profile: %s\n", runtime.profile)
    fmt.Printf("[matrix] work:     %s\n", runtime.work)

    var failures []string
    ran := 0
    skipped := 0
    for _, target := range matrixTargets {
        if !runtime.isSelected(target.Name) {
            continue
        }
        reason := targetSkipReason(target)
        if reason != "" {
            fmt.Printf("[matrix] SKIP %-8s %s\n", target.Name, reason)
            skipped++
            continue
        }
        fmt.Printf("\n[matrix] === %s ===\n", target.Name)
        if err := runtime.runTarget(target); err != nil {
            failures = append(failures, err.Error())
        }
        ran++
    }

    if ran == 0 && skipped == 0 {
        return errors.New("matrix selected no known languages")
    }
    if len(failures) > 0 {
        return fmt.Errorf("matrix failed (%d target(s)):\n  %s", len(failures), strings.Join(failures, "\n  "))
    }
    fmt.Printf("\n[matrix] PASS: %d target(s), %d skipped\n", ran, skipped)
    return nil
}

func newMatrixRuntime(config MatrixConfig) (*matrixRuntime, error) {
    root, err := findProjectRoot(config.Root)
    if err != nil {
        return nil, err
    }
    bubbler := config.Bubbler
    if bubbler == "" {
        bubbler = os.Getenv("BUBBLER")
    }
    if bubbler == "" {
        name := "bubbler"
        if filepath.Separator == '\\' {
            name = "bubbler.exe"
        }
        bubbler = filepath.Join(root, name)
    } else if !filepath.IsAbs(bubbler) && strings.ContainsRune(bubbler, filepath.Separator) {
        bubbler = filepath.Join(root, bubbler)
    }
    if _, err := os.Stat(bubbler); err != nil {
        if resolved, lookErr := exec.LookPath(bubbler); lookErr == nil {
            bubbler = resolved
        } else {
            return nil, fmt.Errorf("bubbler executable %q: %w", bubbler, err)
        }
    }

    work := config.Work
    if work == "" {
        work = filepath.Join(os.TempDir(), fmt.Sprintf("bubbler-e2e-matrix-%d", os.Getpid()))
    } else if !filepath.IsAbs(work) {
        work = filepath.Join(root, work)
    }

    profile := strings.ToLower(strings.TrimSpace(config.Profile))
    switch profile {
    case "smoke", "unroll", "options", "full":
    default:
        return nil, fmt.Errorf("unknown matrix profile %q (want smoke, unroll, options, full)", config.Profile)
    }

    selected := make(map[string]bool)
    for _, item := range strings.Split(config.Langs, ",") {
        item = strings.ToLower(strings.TrimSpace(item))
        if item != "" {
            selected[item] = true
        }
    }

    return &matrixRuntime{
        root:     root,
        e2e:      filepath.Join(root, "e2e"),
        bubbler:  bubbler,
        work:     work,
        profile:  profile,
        keep:     config.Keep,
        selected: selected,
        spec:     AllScenarios(),
    }, nil
}

func findProjectRoot(explicit string) (string, error) {
    if explicit != "" {
        root, err := filepath.Abs(explicit)
        if err != nil {
            return "", err
        }
        if isProjectRoot(root) {
            return root, nil
        }
        return "", fmt.Errorf("%s is not the bubbler project root", root)
    }
    dir, err := os.Getwd()
    if err != nil {
        return "", err
    }
    for {
        if isProjectRoot(dir) {
            return dir, nil
        }
        parent := filepath.Dir(dir)
        if parent == dir {
            break
        }
        dir = parent
    }
    return "", errors.New("cannot locate project root (go.mod and e2e/testcase.bb not found)")
}

func isProjectRoot(dir string) bool {
    if _, err := os.Stat(filepath.Join(dir, "go.mod")); err != nil {
        return false
    }
    if _, err := os.Stat(filepath.Join(dir, "e2e", "testcase.bb")); err != nil {
        return false
    }
    return true
}

func (runtime *matrixRuntime) isSelected(name string) bool {
    if len(runtime.selected) == 0 {
        return true
    }
    return runtime.selected[name]
}

func targetSkipReason(target matrixTarget) string {
    if target.Name == "cjs" && os.Getenv("SKIP_JS") == "1" && os.Getenv("SKIP_CJS") == "" {
        return "SKIP_JS=1"
    }
    if os.Getenv(target.SkipEnv) == "1" {
        return target.SkipEnv + "=1"
    }
    if target.Name == "python" && pythonExecutable() == "" {
        return "python interpreter not found"
    }
    if target.Name == "csharp" {
        output, err := exec.Command("dotnet", "--version").Output()
        if err == nil {
            majorText := strings.SplitN(strings.TrimSpace(string(output)), ".", 2)[0]
            major, parseErr := strconv.Atoi(majorText)
            if parseErr == nil && major < 8 {
                return "dotnet SDK 8+ not found"
            }
        }
    }
    for _, tool := range target.Tools {
        if _, err := exec.LookPath(tool); err != nil {
            return tool + " not found"
        }
    }
    return ""
}

func (runtime *matrixRuntime) runTarget(target matrixTarget) error {
    variants := runtime.variants(target)
    loopCounts := make(map[int]int)
    var failures []string
    for _, variant := range variants {
        fmt.Printf("[matrix] RUN  %-8s %s\n", target.Name, variant.Name)
        dir := filepath.Join(runtime.work, target.Name, variant.Name)
        if err := os.RemoveAll(dir); err != nil {
            failures = append(failures, fmt.Sprintf("%s: clean: %v", variant.Name, err))
            continue
        }
        if err := os.MkdirAll(dir, 0o755); err != nil {
            failures = append(failures, fmt.Sprintf("%s: mkdir: %v", variant.Name, err))
            continue
        }
        if err := runtime.runVariant(target, variant, dir); err != nil {
            fmt.Printf("[matrix] FAIL %-8s %s\n", target.Name, variant.Name)
            failures = append(failures, fmt.Sprintf("%s/%s: %v", target.Name, variant.Name, err))
            continue
        }
        if err := validateGeneratedIndent(dir); err != nil {
            failures = append(failures, fmt.Sprintf("%s/%s: %v", target.Name, variant.Name, err))
            continue
        }
        if threshold, ok := variantUnroll(variant); ok {
            count, err := countMarker(dir, target.LoopMarker)
            if err != nil {
                failures = append(failures, fmt.Sprintf("%s/%s: count loops: %v", target.Name, variant.Name, err))
                continue
            }
            loopCounts[threshold] = count
            fmt.Printf("[matrix]      %-8s %s loop markers=%d\n", target.Name, variant.Name, count)
        }
        fmt.Printf("[matrix] PASS %-8s %s\n", target.Name, variant.Name)
    }
    if runtime.profile == "unroll" || runtime.profile == "full" {
        if err := validateUnrollShape(target.Name, loopCounts); err != nil {
            failures = append(failures, err.Error())
        }
    }
    if len(failures) > 0 {
        return errors.New(strings.Join(failures, "; "))
    }
    return nil
}

func (runtime *matrixRuntime) variants(target matrixTarget) []matrixVariant {
    variants := make([]matrixVariant, 0, 1+len(target.Options)+len(unrollThresholds))
    if runtime.profile == "smoke" || runtime.profile == "options" || runtime.profile == "full" {
        variants = append(variants, matrixVariant{Name: "default"})
    }
    if runtime.profile == "options" || runtime.profile == "full" {
        variants = append(variants, target.Options...)
    }
    if runtime.profile == "unroll" || runtime.profile == "full" {
        for _, threshold := range unrollThresholds {
            name := "unroll_" + strconv.Itoa(threshold)
            if threshold < 0 {
                name = "unroll_neg" + strconv.Itoa(-threshold)
            }
            variants = append(variants, matrixVariant{
                Name:  name,
                Flags: []string{"-unroll=" + strconv.Itoa(threshold)},
            })
        }
    }
    return variants
}

func variantUnroll(variant matrixVariant) (int, bool) {
    for _, flag := range variant.Flags {
        if strings.HasPrefix(flag, "-unroll=") {
            value, err := strconv.Atoi(strings.TrimPrefix(flag, "-unroll="))
            return value, err == nil
        }
    }
    return 0, false
}

func (runtime *matrixRuntime) runVariant(target matrixTarget, variant matrixVariant, dir string) error {
    switch target.Name {
    case "c":
        return runtime.runC(target, variant, dir)
    case "cpp":
        return runtime.runCpp(target, variant, dir)
    case "go":
        return runtime.runGo(target, variant, dir)
    case "java":
        return runtime.runJava(target, variant, dir)
    case "python":
        return runtime.runPython(target, variant, dir)
    case "csharp":
        return runtime.runCsharp(target, variant, dir)
    case "cjs":
        return runtime.runJS(target, variant, dir, false)
    case "esm":
        return runtime.runJS(target, variant, dir, true)
    default:
        return fmt.Errorf("unsupported matrix target %s", target.Name)
    }
}

func (runtime *matrixRuntime) generate(target, output, schema string, flags ...string) error {
    args := []string{"-t", target}
    args = append(args, flags...)
    args = append(args, "-o", output, schema)
    return runCommand(runtime.root, runtime.bubbler, args...)
}

func (runtime *matrixRuntime) schemas() []string {
    return []string{
        filepath.Join(runtime.e2e, "testcase.bb"),
        filepath.Join(runtime.e2e, "features", "bitwid.bb"),
    }
}

func (runtime *matrixRuntime) runC(target matrixTarget, variant matrixVariant, dir string) error {
    genDir := filepath.Join(dir, "gen")
    if err := os.MkdirAll(genDir, 0o755); err != nil {
        return err
    }
    for _, schema := range runtime.schemas() {
        if err := runtime.generate(target.BubblerID, genDir+string(filepath.Separator), schema, variant.Flags...); err != nil {
            return err
        }
    }
    if err := writeMatrixFile(filepath.Join(dir, "main.c"), EmitC(runtime.spec)); err != nil {
        return err
    }
    sources, err := filesWithSuffix(genDir, ".c")
    if err != nil {
        return err
    }
    args := []string{"-std=c11", "-I" + genDir, "-o", filepath.Join(dir, "run_test"), filepath.Join(dir, "main.c")}
    args = append(args, sources...)
    args = append(args, "-lm")
    if err := runCommand(dir, "gcc", args...); err != nil {
        return err
    }
    return runCommand(dir, filepath.Join(dir, "run_test"))
}

func (runtime *matrixRuntime) runCpp(target matrixTarget, variant matrixVariant, dir string) error {
    genDir := filepath.Join(dir, "gen")
    if err := os.MkdirAll(genDir, 0o755); err != nil {
        return err
    }
    single := hasFlag(variant.Flags, "-single")
    schemas := runtime.schemas()
    if single {
        outputs := []string{filepath.Join(genDir, "testpkg.bb.cpp"), filepath.Join(genDir, "bitwid.bb.cpp")}
        for i, schema := range schemas {
            if err := runtime.generate(target.BubblerID, outputs[i], schema, variant.Flags...); err != nil {
                return err
            }
        }
    } else {
        for _, schema := range schemas {
            if err := runtime.generate(target.BubblerID, genDir+string(filepath.Separator), schema, variant.Flags...); err != nil {
                return err
            }
        }
    }
    if err := writeMatrixFile(filepath.Join(dir, "main.cpp"), EmitCpp(runtime.spec)); err != nil {
        return err
    }
    args := []string{"-std=c++20", "-I" + genDir}
    if single {
        args = append(args, "-DBUBBLER_CPP_SINGLE")
    }
    if hasFlag(variant.Flags, "-compat") {
        args = append(args, "-DBUBBLER_CPP_COMPAT")
    }
    args = append(args, "-o", filepath.Join(dir, "run_test"), filepath.Join(dir, "main.cpp"))
    if !single {
        sources, err := filesWithSuffix(genDir, ".cpp")
        if err != nil {
            return err
        }
        args = append(args, sources...)
    }
    args = append(args, "-lm")
    if err := runCommand(dir, "g++", args...); err != nil {
        return err
    }
    return runCommand(dir, filepath.Join(dir, "run_test"))
}

func (runtime *matrixRuntime) runGo(target matrixTarget, variant matrixVariant, dir string) error {
    if err := copyFile(filepath.Join(runtime.e2e, "tests", "go", "go.mod"), filepath.Join(dir, "go.mod")); err != nil {
        return err
    }
    for _, schema := range runtime.schemas() {
        if err := runtime.generate(target.BubblerID, dir+string(filepath.Separator), schema, variant.Flags...); err != nil {
            return err
        }
    }
    if err := EmitGo(runtime.spec, dir); err != nil {
        return err
    }
    return runCommand(dir, "go", "test", "./...", "-count=1")
}

func (runtime *matrixRuntime) runJava(target matrixTarget, variant matrixVariant, dir string) error {
    genDir := filepath.Join(dir, "gen")
    outDir := filepath.Join(dir, "out")
    if err := os.MkdirAll(genDir, 0o755); err != nil {
        return err
    }
    if err := os.MkdirAll(outDir, 0o755); err != nil {
        return err
    }
    for _, schema := range runtime.schemas() {
        if err := runtime.generate(target.BubblerID, genDir+string(filepath.Separator), schema, variant.Flags...); err != nil {
            return err
        }
    }
    if err := writeMatrixFile(filepath.Join(dir, "Main.java"), EmitJava(runtime.spec)); err != nil {
        return err
    }
    sources, err := filesWithSuffix(genDir, ".java")
    if err != nil {
        return err
    }
    args := []string{"-encoding", "UTF-8", "-d", outDir}
    args = append(args, sources...)
    if err := runCommand(dir, "javac", args...); err != nil {
        return err
    }
    if err := runCommand(dir, "javac", "-encoding", "UTF-8", "-cp", outDir, "-d", outDir, filepath.Join(dir, "Main.java")); err != nil {
        return err
    }
    return runCommand(dir, "java", "-cp", outDir, "Main")
}

func (runtime *matrixRuntime) runPython(target matrixTarget, variant matrixVariant, dir string) error {
    genDir := filepath.Join(dir, "gen")
    if err := os.MkdirAll(genDir, 0o755); err != nil {
        return err
    }
    flags := appendUniqueFlag(variant.Flags, "-single")
    outputs := []string{filepath.Join(genDir, "testcase_bb.py"), filepath.Join(genDir, "bitwid_bb.py")}
    for i, schema := range runtime.schemas() {
        if err := runtime.generate(target.BubblerID, outputs[i], schema, flags...); err != nil {
            return err
        }
    }
    if err := writeMatrixFile(filepath.Join(dir, "test_main.py"), EmitPython(runtime.spec)); err != nil {
        return err
    }
    return runCommand(dir, pythonExecutable(), "test_main.py")
}

func pythonExecutable() string {
    candidates := []string{"python3", "python"}
    if goruntime.GOOS == "windows" {
        candidates = []string{"python", "python3"}
    }
    for _, candidate := range candidates {
        if _, err := exec.LookPath(candidate); err == nil {
            return candidate
        }
    }
    return ""
}

func (runtime *matrixRuntime) runCsharp(target matrixTarget, variant matrixVariant, dir string) error {
    genDir := filepath.Join(dir, "gen")
    if err := os.MkdirAll(genDir, 0o755); err != nil {
        return err
    }
    if err := copyFile(filepath.Join(runtime.e2e, "tests", "csharp", "test.csproj"), filepath.Join(dir, "test.csproj")); err != nil {
        return err
    }
    flags := appendUniqueFlag(variant.Flags, "-single")
    outputs := []string{filepath.Join(genDir, "testcase.bb.cs"), filepath.Join(genDir, "bitwid.bb.cs")}
    for i, schema := range runtime.schemas() {
        if err := runtime.generate(target.BubblerID, outputs[i], schema, flags...); err != nil {
            return err
        }
    }
    if err := writeMatrixFile(filepath.Join(dir, "Program.cs"), EmitCsharp(runtime.spec)); err != nil {
        return err
    }
    return runCommand(dir, "dotnet", "run", "-f", "net8.0", "--project", "test.csproj", "--nologo")
}

func (runtime *matrixRuntime) runJS(target matrixTarget, variant matrixVariant, dir string, esm bool) error {
    genDir := filepath.Join(dir, "gen")
    if err := os.MkdirAll(genDir, 0o755); err != nil {
        return err
    }
    if esm {
        if err := copyFile(filepath.Join(runtime.e2e, "tests", "esm", "package.json"), filepath.Join(dir, "package.json")); err != nil {
            return err
        }
    }
    flags := appendUniqueFlag(variant.Flags, "-single")
    outputs := []string{filepath.Join(genDir, "testcase.bb.js"), filepath.Join(genDir, "bitwid.bb.js")}
    for i, schema := range runtime.schemas() {
        if err := runtime.generate(target.BubblerID, outputs[i], schema, flags...); err != nil {
            return err
        }
    }
    var driver string
    if esm {
        driver = EmitEsm(runtime.spec)
    } else {
        driver = EmitCjs(runtime.spec)
    }
    if err := writeMatrixFile(filepath.Join(dir, "test.mjs"), driver); err != nil {
        return err
    }
    return runCommand(dir, "node", "test.mjs")
}

func appendUniqueFlag(flags []string, required string) []string {
    out := append([]string(nil), flags...)
    if !hasFlag(out, required) {
        out = append(out, required)
    }
    return out
}

func hasFlag(flags []string, wanted string) bool {
    for _, flag := range flags {
        if flag == wanted {
            return true
        }
    }
    return false
}

func validateUnrollShape(target string, counts map[int]int) error {
    if len(counts) != len(unrollThresholds) {
        return fmt.Errorf("%s unroll shape: got %d threshold results, want %d", target, len(counts), len(unrollThresholds))
    }
    if counts[-1] != 0 {
        return fmt.Errorf("%s unroll=-1 emitted %d loop marker(s), want 0", target, counts[-1])
    }
    if counts[0] <= 0 {
        return fmt.Errorf("%s unroll=0 emitted no loop markers", target)
    }
    previous := counts[0]
    changed := false
    for _, threshold := range unrollThresholds {
        if threshold <= 0 {
            continue
        }
        current := counts[threshold]
        if current > previous {
            return fmt.Errorf("%s loop count increased at threshold %d: %d > %d", target, threshold, current, previous)
        }
        if current < previous {
            changed = true
        }
        previous = current
    }
    if !changed {
        return fmt.Errorf("%s loop count never changed across positive thresholds: %v", target, counts)
    }
    return nil
}

func countMarker(root, marker string) (int, error) {
    total := 0
    err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if entry.IsDir() || !isGeneratedSource(path) {
            return nil
        }
        data, err := os.ReadFile(path)
        if err != nil {
            return err
        }
        total += bytes.Count(data, []byte(marker))
        return nil
    })
    return total, err
}

func validateGeneratedIndent(root string) error {
    return filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if entry.IsDir() || !isGeneratedSource(path) {
            return nil
        }
        data, err := os.ReadFile(path)
        if err != nil {
            return err
        }
        if bytes.ContainsRune(data, '\t') {
            return fmt.Errorf("generated source contains a tab: %s", path)
        }
        return nil
    })
}

func isGeneratedSource(path string) bool {
    ext := strings.ToLower(filepath.Ext(path))
    switch ext {
    case ".c", ".cpp", ".go", ".java", ".cs", ".js", ".py":
        return strings.Contains(path, string(filepath.Separator)+"gen"+string(filepath.Separator)) || strings.Contains(filepath.Base(path), ".bb.")
    default:
        return false
    }
}

func filesWithSuffix(root, suffix string) ([]string, error) {
    var files []string
    err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if !entry.IsDir() && strings.HasSuffix(strings.ToLower(path), strings.ToLower(suffix)) {
            files = append(files, path)
        }
        return nil
    })
    sort.Strings(files)
    return files, err
}

func writeMatrixFile(path, content string) error {
    if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
        return err
    }
    content = strings.ReplaceAll(content, "\t", "    ")
    if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
        return fmt.Errorf("write %s: %w", path, err)
    }
    return nil
}

func copyFile(source, destination string) error {
    data, err := os.ReadFile(source)
    if err != nil {
        return err
    }
    if err := os.MkdirAll(filepath.Dir(destination), 0o755); err != nil {
        return err
    }
    return os.WriteFile(destination, data, 0o644)
}

func runCommand(dir, name string, args ...string) error {
    command := exec.Command(name, args...)
    command.Dir = dir
    command.Env = append(os.Environ(), "GOWORK=off")
    var output bytes.Buffer
    command.Stdout = &output
    command.Stderr = &output
    if err := command.Run(); err != nil {
        text := strings.TrimSpace(output.String())
        if text == "" {
            return fmt.Errorf("%s %s: %w", name, strings.Join(args, " "), err)
        }
        return fmt.Errorf("%s %s: %w\n%s", name, strings.Join(args, " "), err, text)
    }
    return nil
}
