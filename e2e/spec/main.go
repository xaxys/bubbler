package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var (
		lang     = flag.String("lang", "", "Target language: c, cpp, go, java, kotlin, python, csharp, cjs, esm")
		out      = flag.String("out", "", "Output path. For most languages this is a single file; for Go it is a directory containing per-package subdirs")
		matrix   = flag.Bool("matrix", false, "Run the declarative codec/option/unroll test matrix")
		bubbler  = flag.String("bubbler", "", "Path to the bubbler executable used by -matrix")
		root     = flag.String("root", "", "Project root used by -matrix (default: auto-detect)")
		work     = flag.String("work", "", "Matrix work directory (default: a temporary directory)")
		langs    = flag.String("langs", "", "Comma-separated matrix languages (default: all available)")
		profile  = flag.String("profile", "full", "Matrix profile: smoke, unroll, options, full")
		keepWork = flag.Bool("keep-work", false, "Keep matrix work files after the run")
	)
	flag.Parse()

	if *matrix {
		err := RunMatrix(MatrixConfig{
			Bubbler: *bubbler,
			Root:    *root,
			Work:    *work,
			Langs:   *langs,
			Profile: *profile,
			Keep:    *keepWork,
		})
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}

	if *lang == "" || *out == "" {
		fmt.Fprintln(os.Stderr, "usage: go run ./e2e/spec -lang=<lang> -out=<path>")
		fmt.Fprintln(os.Stderr, "   or: go run ./e2e/spec -matrix -bubbler=<path> [-profile=full]")
		os.Exit(2)
	}

	spec := AllScenarios()

	if err := os.MkdirAll(filepath.Dir(*out), 0o755); err != nil {
		fmt.Fprintf(os.Stderr, "mkdir parent: %v\n", err)
		os.Exit(1)
	}

	switch *lang {
	case "c":
		write(*out, EmitC(spec))
	case "cpp":
		write(*out, EmitCpp(spec))
	case "go":
		// Go is multi-file: write *_test.go into per-package subdirs of out.
		if err := EmitGo(spec, *out); err != nil {
			fmt.Fprintf(os.Stderr, "emit go: %v\n", err)
			os.Exit(1)
		}
	case "java":
		write(*out, EmitJava(spec))
	case "kotlin":
		write(*out, EmitKotlin(spec))
	case "python":
		write(*out, EmitPython(spec))
	case "csharp":
		write(*out, EmitCsharp(spec))
	case "cjs":
		write(*out, EmitCjs(spec))
	case "esm":
		write(*out, EmitEsm(spec))
	default:
		fmt.Fprintf(os.Stderr, "unknown -lang=%s\n", *lang)
		os.Exit(2)
	}
}

func write(path, content string) {
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "write %s: %v\n", path, err)
		os.Exit(1)
	}
}
