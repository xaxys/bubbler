#!/usr/bin/env bash
##############################################################################
# Bubbler E2E — Compiler Options Test
#
# Tests various bubbler compiler flags and .bb file options.
# Run from the e2e/ directory:   bash compiler/test_options.sh
# or let run_tests.sh invoke it.
#
# Environment:
#   BUBBLER  — Path to the bubbler binary (default: ../bubbler)
##############################################################################
set -euo pipefail
cd "$(dirname "$0")/.."   # ensure we run from e2e/
E2E_DIR="$(pwd)"
PROJECT_ROOT="$(cd .. && pwd)"

if [[ -z "${BUBBLER:-}" ]]; then
    if [[ -x "../bubbler.exe" ]]; then
        BUBBLER="../bubbler.exe"
    else
        BUBBLER="../bubbler"
    fi
fi
TMPDIR_BASE="tests/.tmp_bubbler_opts_$$"
mkdir -p "$TMPDIR_BASE"
trap 'rm -rf "$TMPDIR_BASE"' EXIT

# Build the test-driver generator once per variant invocation. Test drivers
# are identical across compiler-flag variants — only the *.bb codegen
# changes — so we render once into the variant directory.
gen_driver() {
    local lang="$1" out="$2"
    (cd "$PROJECT_ROOT" && go run ./e2e/spec -lang="$lang" -out "$out")
}

pass=0
fail=0

# ─── helpers ────────────────────────────────────────────────────────────────

ok()   { echo "  PASS: $*"; pass=$((pass + 1)); }
fail() { echo "  FAIL: $*" >&2; fail=$((fail + 1)); }

check_exists()     { [[ -f "$1" ]] && ok "$2: file exists"      || fail "$2: missing $1"; }
check_not_exists() { [[ ! -f "$1" ]] && ok "$2: file absent"    || fail "$2: should not exist $1"; }
check_dir_exists() { [[ -d "$1" ]] && ok "$2: dir exists"       || fail "$2: missing dir $1"; }

check_contains() {
    local file="$1" pattern="$2" label="$3"
    grep -q "$pattern" "$file" && ok "$label" || fail "$label: pattern '$pattern' not found in $file"
}

check_not_contains() {
    local file="$1" pattern="$2" label="$3"
    grep -q "$pattern" "$file" && fail "$label: pattern '$pattern' should NOT be in $file" || ok "$label"
}

run_variant() {
    local label="$1"
    shift
    if "$@"; then
        ok "$label"
    else
        fail "$label"
    fi
}

variant_c() {
    local name="$1"
    shift
    local out="$TMPDIR_BASE/variants/c/$name"
    mkdir -p "$out/gen"
    "$BUBBLER" -t c "$@" -o "$out/gen/" testcase.bb
    "$BUBBLER" -t c "$@" -o "$out/gen/" features/bitwid.bb
    gen_driver c "$E2E_DIR/$out/main.c"
    gcc -std=c11 -I"$out" -I"$out/gen" -o "$out/run_test" "$out/main.c" "$out/gen/testpkg.bb.c" "$out/gen/bitwid.bb.c" -lm
    "$out/run_test" > /dev/null
}

variant_cpp() {
    local name="$1"
    shift
    local out="$TMPDIR_BASE/variants/cpp/$name"
    local -a bubbler_flags=("$@")
    local -a cxx_flags=()
    local is_single=0
    local is_compat=0

    for arg in "${bubbler_flags[@]}"; do
        if [[ "$arg" == "-single" ]]; then
            is_single=1
        fi
        if [[ "$arg" == "-compat" ]]; then
            is_compat=1
        fi
    done

    if [[ "$is_single" -eq 1 ]]; then
        cxx_flags+=("-DBUBBLER_CPP_SINGLE")
    fi
    if [[ "$is_compat" -eq 1 ]]; then
        cxx_flags+=("-DBUBBLER_CPP_COMPAT")
    fi

    mkdir -p "$out/gen"
    gen_driver cpp "$E2E_DIR/$out/main.cpp"
    if [[ "$is_single" -eq 1 ]]; then
        "$BUBBLER" -t cpp "${bubbler_flags[@]}" -o "$out/gen/testpkg.bb.cpp" testcase.bb
        "$BUBBLER" -t cpp "${bubbler_flags[@]}" -o "$out/gen/bitwid.bb.cpp" features/bitwid.bb
        g++ -std=c++20 "${cxx_flags[@]}" -I"$out" -I"$out/gen" -o "$out/run_test" "$out/main.cpp" -lm
    else
        "$BUBBLER" -t cpp "${bubbler_flags[@]}" -o "$out/gen/" testcase.bb
        "$BUBBLER" -t cpp "${bubbler_flags[@]}" -o "$out/gen/" features/bitwid.bb
        g++ -std=c++20 "${cxx_flags[@]}" -I"$out" -I"$out/gen" -o "$out/run_test" "$out/main.cpp" "$out/gen/testpkg.bb.cpp" "$out/gen/bitwid.bb.cpp" -lm
    fi
    "$out/run_test" > /dev/null
}

variant_go() {
    local name="$1"
    shift
    local out="$TMPDIR_BASE/variants/go/$name"
    mkdir -p "$out/bitwid" "$out/testpkg"
    cp tests/go/go.mod "$out/go.mod"
    "$BUBBLER" -t go "$@" -o "$out/" testcase.bb
    "$BUBBLER" -t go "$@" -o "$out/" features/bitwid.bb
    gen_driver go "$E2E_DIR/$out/"
    (cd "$out" && go test ./... > /dev/null)
}

variant_java() {
    local name="$1"
    shift
    local out="$TMPDIR_BASE/variants/java/$name"
    mkdir -p "$out/gen"
    gen_driver java "$E2E_DIR/$out/Main.java"
    "$BUBBLER" -t java "$@" -o "$out/gen/" testcase.bb
    "$BUBBLER" -t java "$@" -o "$out/gen/" features/bitwid.bb
    (cd "$out" && rm -rf out && mkdir out && find gen -name "*.java" -exec javac -encoding UTF-8 -d out {} + && javac -encoding UTF-8 -cp out -d out Main.java && java -cp out Main > /dev/null)
}

variant_python() {
    local name="$1"
    shift
    local out="$TMPDIR_BASE/variants/python/$name"
    mkdir -p "$out/gen"
    gen_driver python "$E2E_DIR/$out/test_main.py"
    "$BUBBLER" -t py "$@" -single -o "$out/gen/testcase_bb.py" testcase.bb
    "$BUBBLER" -t py "$@" -single -o "$out/gen/bitwid_bb.py" features/bitwid.bb
    (cd "$out" && python3 test_main.py > /dev/null)
}

variant_csharp() {
    local name="$1"
    shift
    local out="$TMPDIR_BASE/variants/csharp/$name"
    mkdir -p "$out/gen"
    gen_driver csharp "$E2E_DIR/$out/Program.cs"
    cp tests/csharp/test.csproj "$out/test.csproj"
    "$BUBBLER" -t cs "$@" -single -o "$out/gen/testcase.bb.cs" testcase.bb
    "$BUBBLER" -t cs "$@" -single -o "$out/gen/bitwid.bb.cs" features/bitwid.bb
    (
        cd "$out"
        dotnet run -f net8.0 --project test.csproj > dotnet.log 2>&1 || {
            echo "[csharp:$name] dotnet run failed" >&2
            cat dotnet.log >&2
            return 1
        }
    )
}

variant_cjs() {
    local name="$1"
    shift
    local out="$TMPDIR_BASE/variants/cjs/$name"
    mkdir -p "$out/gen"
    gen_driver cjs "$E2E_DIR/$out/test.mjs"
    "$BUBBLER" -t cjs "$@" -single -o "$out/gen/testcase.bb.js" testcase.bb
    "$BUBBLER" -t cjs "$@" -single -o "$out/gen/bitwid.bb.js" features/bitwid.bb
    (cd "$out" && node test.mjs > /dev/null)
}

variant_esm() {
    local name="$1"
    shift
    local out="$TMPDIR_BASE/variants/esm/$name"
    mkdir -p "$out/gen"
    gen_driver esm "$E2E_DIR/$out/test.mjs"
    cp tests/esm/package.json "$out/package.json"
    "$BUBBLER" -t mjs "$@" -single -o "$out/gen/testcase.bb.js" testcase.bb
    "$BUBBLER" -t mjs "$@" -single -o "$out/gen/bitwid.bb.js" features/bitwid.bb
    (cd "$out" && node test.mjs > /dev/null)
}

##############################################################################
# 1. go_package — output directory reflects option value
##############################################################################
echo
echo "=== #1: go_package output directory ==="
out="$TMPDIR_BASE/gopkg"
"$BUBBLER" -t go -o "$out/" features/imports/protocol.bb
check_dir_exists "$out/github.com/example/protocol" "go_package deep dir"
check_exists     "$out/github.com/example/protocol/protocol.bb.go" "protocol.bb.go"

##############################################################################
# 2. -rmpath — strip go_package prefix from output path
##############################################################################
echo
echo "=== #2: -rmpath strip prefix ==="
out="$TMPDIR_BASE/rmpath"
"$BUBBLER" -t go -rmpath github.com/example -o "$out/" features/imports/protocol.bb
check_dir_exists "$out/protocol" "protocol dir (after rmpath)"
check_exists     "$out/protocol/protocol.bb.go" "protocol.bb.go (after rmpath)"
check_not_exists "$out/github.com/example/protocol/protocol.bb.go" "deep dir absent (after rmpath)"

##############################################################################
# 3. java_package — output directory reflects com.example.bitwid → com/example/bitwid
##############################################################################
echo
echo "=== #3: java_package output directory ==="
out="$TMPDIR_BASE/javapkg"
"$BUBBLER" -t java -o "$out/" features/bitwid.bb
check_dir_exists "$out/com/example/bitwid" "java package dir"
check_exists     "$out/com/example/bitwid/NarrowBWTest.java" "NarrowBWTest.java"

##############################################################################
# 4. -single — single output file vs multiple files
##############################################################################
echo
echo "=== #4: -single produces one file ==="
# Without -single: Go target produces directory structure
out_multi="$TMPDIR_BASE/multi"
"$BUBBLER" -t go -o "$out_multi/" testcase.bb
check_dir_exists "$out_multi/testpkg" "-single absent: testpkg dir"
check_exists     "$out_multi/testpkg/testpkg.bb.go" "testpkg.bb.go without -single"

# With -single: Go target produces single file at -o path
single_file="$TMPDIR_BASE/testcase_single.go"
"$BUBBLER" -t go -single -o "$single_file" testcase.bb
check_exists     "$single_file" "-single: single file created"
check_not_exists "$TMPDIR_BASE/testpkg/testpkg.bb.go" "-single: no subdirectory"

##############################################################################
# 5. -minimal — removes raw accessor functions
##############################################################################
echo
echo "=== #5: -minimal removes raw getters/setters ==="
out_min="$TMPDIR_BASE/minimal"
out_full="$TMPDIR_BASE/full"
"$BUBBLER" -t c -minimal -o "$out_min/" testcase.bb
"$BUBBLER" -t c          -o "$out_full/" testcase.bb

check_not_contains "$out_min/testpkg.bb.h"  "RawGetter"   "-minimal: no RawGetter"
check_not_contains "$out_min/testpkg.bb.h"  "RawSetter"   "-minimal: no RawSetter"
# Custom getters/setters are preserved even with -minimal
check_contains     "$out_min/testpkg.bb.h"  "get_voltage" "-minimal: custom getter kept"
# Non-minimal has raw getters
check_contains     "$out_full/testpkg.bb.h" "RawGetter"   "no -minimal: RawGetter present"

##############################################################################
# 6. -decnum — decimal constants instead of hex/binary
##############################################################################
echo
echo "=== #6: -decnum uses decimal constants ==="
out_dec="$TMPDIR_BASE/decnum"
out_hex="$TMPDIR_BASE/hexnum"
"$BUBBLER" -t c -decnum -o "$out_dec/" testcase.bb
"$BUBBLER" -t c         -o "$out_hex/" testcase.bb

# Without -decnum: uses 0xAA hex literal for header constant
check_contains     "$out_hex/testpkg.bb.c" "0xAA"  "no -decnum: hex literal 0xAA"
# With -decnum: uses 170 decimal literal
check_contains     "$out_dec/testpkg.bb.c" "170"   "-decnum: decimal literal 170"
check_not_contains "$out_dec/testpkg.bb.c" "0xAA"  "-decnum: no hex literal"
# With -decnum: uses decimal mask (255) instead of binary (0b11111111)
check_contains     "$out_dec/testpkg.bb.c" "255"   "-decnum: decimal mask 255"
check_not_contains "$out_dec/testpkg.bb.c" "0b11111111" "-decnum: no binary mask"

##############################################################################
# 7. omit_empty — empty package generates no output file
##############################################################################
echo
echo "=== #7: omit_empty skips generation for empty packages ==="
out="$TMPDIR_BASE/omit"
"$BUBBLER" -t c -o "$out/" features/omit_e.bb
# types.bb (imported) should be generated
check_exists     "$out/types.bb.h"      "omit_empty: types.bb.h generated"
# importonly package has no local structs + omit_empty = true → NOT generated
check_not_exists "$out/importonly.bb.h" "omit_empty: importonly.bb.h NOT generated"

##############################################################################
# 8. Multi-file import — protocol.bb imports types.bb
##############################################################################
echo
echo "=== #8: multi-file import generates both packages ==="
out="$TMPDIR_BASE/import"
"$BUBBLER" -t c -o "$out/" features/imports/protocol.bb
check_exists "$out/types.bb.h"    "import: types.bb.h"
check_exists "$out/types.bb.c"    "import: types.bb.c"
check_exists "$out/protocol.bb.h" "import: protocol.bb.h"
check_exists "$out/protocol.bb.c" "import: protocol.bb.c"
# Verify that Vec2 and Direction (from types.bb) are referenced in protocol output
check_contains "$out/protocol.bb.h" "Vec2"      "import: Vec2 referenced"
check_contains "$out/protocol.bb.h" "Direction" "import: Direction referenced"

##############################################################################
# 9. -signext arith — different codegen but same encode/decode result
#    (just verify the code compiles; correctness is verified by language tests)
##############################################################################
echo
echo "=== #9: -signext variants produce valid code ==="
out_shift="$TMPDIR_BASE/signext_shift"
out_arith="$TMPDIR_BASE/signext_arith"
"$BUBBLER" -t c -signext shift -o "$out_shift/" testcase.bb
check_exists "$out_shift/testpkg.bb.c" "-signext shift: code generated"
"$BUBBLER" -t c -signext arith -o "$out_arith/" testcase.bb
check_exists "$out_arith/testpkg.bb.c" "-signext arith: code generated"

##############################################################################
# 10. -compat — CommonJS uses Array instead of Uint8Array
##############################################################################
echo
echo "=== #10: -compat switches CommonJS from Uint8Array to Array ==="
out_cjs_default="$TMPDIR_BASE/cjs_default"
out_cjs_compat="$TMPDIR_BASE/cjs_compat"
"$BUBBLER" -t cjs -single -o "$out_cjs_default/testcase.bb.js" testcase.bb
"$BUBBLER" -t cjs -compat -single -o "$out_cjs_compat/testcase.bb.js" testcase.bb

# Without -compat: allocates Uint8Array for encode buffers
check_contains     "$out_cjs_default/testcase.bb.js" "new Uint8Array(" "-compat absent: Uint8Array used"
check_not_contains "$out_cjs_default/testcase.bb.js" "if (data === undefined) data = new Array(" "-compat absent: no Array alloc"
# With -compat: allocates Array for encode buffers
check_contains     "$out_cjs_compat/testcase.bb.js" "if (data === undefined) data = new Array(" "-compat: Array alloc used"
check_not_contains "$out_cjs_compat/testcase.bb.js" "if (data === undefined) data = new Uint8Array(" "-compat: no Uint8Array alloc"

##############################################################################
# 11. -unroll with different thresholds — full encode/decode tests
#     Tests array handling with different unroll thresholds:
#     - unroll=-1: always unroll (no loops)
#     - unroll=0: use loops even for fixed-width arrays
#     - unroll=1: use loops for arrays > 1 element
#     - unroll=4: default (use loops for arrays > 4 elements)
#     - unroll=8: use loops only for large arrays
##############################################################################
echo
echo "=== #11: -unroll with various thresholds across all languages ==="
echo "Testing Color<2>, uint8<4> and other arrays with different unroll values"

for unroll_val in -1 0 1 2 4 6 8 16 32; do
    echo ""
    echo "  Testing -unroll=$unroll_val..."

    if command -v gcc >/dev/null 2>&1; then
        run_variant "C -unroll=$unroll_val: full encode/decode test" \
            variant_c unroll_${unroll_val}_c -unroll "$unroll_val"
    fi

    if command -v g++ >/dev/null 2>&1; then
        run_variant "C++ -unroll=$unroll_val: full encode/decode test" \
            variant_cpp unroll_${unroll_val}_cpp -unroll "$unroll_val"
    fi

    if command -v go >/dev/null 2>&1; then
        run_variant "Go -unroll=$unroll_val: full encode/decode test" \
            variant_go unroll_${unroll_val}_go -unroll "$unroll_val"
    fi

    if command -v python3 >/dev/null 2>&1; then
        run_variant "Python -unroll=$unroll_val: full encode/decode test" \
            variant_python unroll_${unroll_val}_py -unroll "$unroll_val"
    fi

    if command -v javac >/dev/null 2>&1 && command -v java >/dev/null 2>&1; then
        run_variant "Java -unroll=$unroll_val: full encode/decode test" \
            variant_java unroll_${unroll_val}_java -unroll "$unroll_val"
    fi

    if command -v dotnet >/dev/null 2>&1; then
        run_variant "C# -unroll=$unroll_val: full encode/decode test" \
            variant_csharp unroll_${unroll_val}_csharp -unroll "$unroll_val"
    fi

    if command -v node >/dev/null 2>&1; then
        run_variant "CommonJS -unroll=$unroll_val: full encode/decode test" \
            variant_cjs unroll_${unroll_val}_cjs -unroll "$unroll_val"
        run_variant "ESModule -unroll=$unroll_val: full encode/decode test" \
            variant_esm unroll_${unroll_val}_esm -unroll "$unroll_val"
    fi
done

##############################################################################
# 11b. -unroll static checks — generated code actually uses loops vs. unrolls
#      depending on threshold. Uses testcase.bb (arrays of size 2/4/8/32).
##############################################################################
echo
echo "=== #11b: -unroll static checks against testcase.bb ==="

# Helper: count occurrences of a loop marker in a generated file.
count_loops() {
    local file="$1" pattern="$2"
    if [[ ! -f "$file" ]]; then echo 0; return; fi
    local cnt
    cnt=$(grep -c -- "$pattern" "$file" 2>/dev/null) || cnt=0
    echo "$cnt"
}

unroll_static_dir="$TMPDIR_BASE/unroll_static"
mkdir -p "$unroll_static_dir"

# Generate Go output for testcase.bb at -unroll=-1 (never loop) and -unroll=0
# (always loop) and verify the loop count meaningfully differs.
if command -v go >/dev/null 2>&1; then
    "$BUBBLER" -t go -unroll=-1 -single -o "$unroll_static_dir/go_neg1.go" testcase.bb
    "$BUBBLER" -t go -unroll=0  -single -o "$unroll_static_dir/go_0.go"     testcase.bb

    loops_neg1=$(count_loops "$unroll_static_dir/go_neg1.go" 'for _i := int64')
    loops_0=$(count_loops "$unroll_static_dir/go_0.go"      'for _i := int64')

    if [[ "$loops_neg1" -eq 0 && "$loops_0" -gt 0 ]]; then
        ok "Go static: unroll=-1 yields no loops; unroll=0 yields $loops_0 loops"
    else
        fail "Go static: expected 0/positive loops, got neg1=$loops_neg1 / 0=$loops_0"
    fi
fi

# Same for C++ (file-private bb_write/read_field_bits and `for (int64_t _i` loop).
if command -v g++ >/dev/null 2>&1; then
    "$BUBBLER" -t cpp -unroll=-1 -single -o "$unroll_static_dir/cpp_neg1.cpp" testcase.bb
    "$BUBBLER" -t cpp -unroll=0  -single -o "$unroll_static_dir/cpp_0.cpp"     testcase.bb

    loops_neg1=$(count_loops "$unroll_static_dir/cpp_neg1.cpp" 'for (int64_t _i')
    loops_0=$(count_loops "$unroll_static_dir/cpp_0.cpp"      'for (int64_t _i')

    if [[ "$loops_neg1" -eq 0 && "$loops_0" -gt 0 ]]; then
        ok "C++ static: unroll=-1 yields no loops; unroll=0 yields $loops_0 loops"
    else
        fail "C++ static: expected 0/positive loops, got neg1=$loops_neg1 / 0=$loops_0"
    fi
fi

# Same for Python.
if command -v python3 >/dev/null 2>&1; then
    "$BUBBLER" -t py -unroll=-1 -single -o "$unroll_static_dir/py_neg1.py" testcase.bb
    "$BUBBLER" -t py -unroll=0  -single -o "$unroll_static_dir/py_0.py"     testcase.bb

    loops_neg1=$(count_loops "$unroll_static_dir/py_neg1.py" 'for _i in range')
    loops_0=$(count_loops "$unroll_static_dir/py_0.py"      'for _i in range')

    if [[ "$loops_neg1" -eq 0 && "$loops_0" -gt 0 ]]; then
        ok "Python static: unroll=-1 yields no loops; unroll=0 yields $loops_0 loops"
    else
        fail "Python static: expected 0/positive loops, got neg1=$loops_neg1 / 0=$loops_0"
    fi
fi

# Threshold boundary check (Go): -unroll=4 must NOT loop arrays of length 4
# (MediumArrays) but MUST loop arrays of length 8 (LargeArrays). Spot-check by
# confirming loop count at unroll=4 is between unroll=-1 and unroll=0.
if command -v go >/dev/null 2>&1; then
    "$BUBBLER" -t go -unroll=4 -single -o "$unroll_static_dir/go_4.go" testcase.bb
    loops_4=$(count_loops "$unroll_static_dir/go_4.go" 'for _i := int64')
    loops_neg1=$(count_loops "$unroll_static_dir/go_neg1.go" 'for _i := int64')
    loops_0=$(count_loops "$unroll_static_dir/go_0.go" 'for _i := int64')
    if [[ "$loops_4" -gt "$loops_neg1" && "$loops_4" -lt "$loops_0" ]]; then
        ok "Go threshold: unroll=4 loops ($loops_4) sit between unroll=-1 ($loops_neg1) and unroll=0 ($loops_0)"
    else
        fail "Go threshold: expected -1 < 4 < 0, got -1=$loops_neg1 / 4=$loops_4 / 0=$loops_0"
    fi
fi

# C: bb_write_field_bits / bb_read_field_bits + `for (int64_t _i` loop.
if command -v gcc >/dev/null 2>&1; then
    "$BUBBLER" -t c -unroll=-1 -single -o "$unroll_static_dir/c_neg1.c" testcase.bb
    "$BUBBLER" -t c -unroll=0  -single -o "$unroll_static_dir/c_0.c"     testcase.bb
    loops_neg1=$(count_loops "$unroll_static_dir/c_neg1.c" 'for (int64_t _i')
    loops_0=$(count_loops "$unroll_static_dir/c_0.c"      'for (int64_t _i')
    if [[ "$loops_neg1" -eq 0 && "$loops_0" -gt 0 ]]; then
        ok "C static: unroll=-1 yields no loops; unroll=0 yields $loops_0 loops"
    else
        fail "C static: expected 0/positive loops, got neg1=$loops_neg1 / 0=$loops_0"
    fi
fi

# Java: bbWriteFieldBits / bbReadFieldBits + `for (int _i` loop.
if command -v javac >/dev/null 2>&1; then
    mkdir -p "$unroll_static_dir/java_neg1" "$unroll_static_dir/java_0"
    "$BUBBLER" -t java -unroll=-1 -o "$unroll_static_dir/java_neg1" testcase.bb
    "$BUBBLER" -t java -unroll=0  -o "$unroll_static_dir/java_0"     testcase.bb
    loops_neg1=0
    while IFS= read -r f; do
        c=$(count_loops "$f" 'for (int _i')
        loops_neg1=$((loops_neg1 + c))
    done < <(find "$unroll_static_dir/java_neg1" -name "*.java")
    loops_0=0
    while IFS= read -r f; do
        c=$(count_loops "$f" 'for (int _i')
        loops_0=$((loops_0 + c))
    done < <(find "$unroll_static_dir/java_0" -name "*.java")
    if [[ "$loops_neg1" -eq 0 && "$loops_0" -gt 0 ]]; then
        ok "Java static: unroll=-1 yields no loops; unroll=0 yields $loops_0 loops"
    else
        fail "Java static: expected 0/positive loops, got neg1=$loops_neg1 / 0=$loops_0"
    fi
fi

# C#: BbWriteFieldBits / BbReadFieldBits + `for (int _i` loop.
if command -v dotnet >/dev/null 2>&1; then
    "$BUBBLER" -t cs -unroll=-1 -single -o "$unroll_static_dir/cs_neg1.cs" testcase.bb
    "$BUBBLER" -t cs -unroll=0  -single -o "$unroll_static_dir/cs_0.cs"     testcase.bb
    loops_neg1=$(count_loops "$unroll_static_dir/cs_neg1.cs" 'for (int _i')
    loops_0=$(count_loops "$unroll_static_dir/cs_0.cs"      'for (int _i')
    if [[ "$loops_neg1" -eq 0 && "$loops_0" -gt 0 ]]; then
        ok "C# static: unroll=-1 yields no loops; unroll=0 yields $loops_0 loops"
    else
        fail "C# static: expected 0/positive loops, got neg1=$loops_neg1 / 0=$loops_0"
    fi
fi

# CommonJS / ESM: bbWriteFieldBits / bbReadFieldBits + `for (let _i` loop.
if command -v node >/dev/null 2>&1; then
    "$BUBBLER" -t cjs -unroll=-1 -single -o "$unroll_static_dir/cjs_neg1.js" testcase.bb
    "$BUBBLER" -t cjs -unroll=0  -single -o "$unroll_static_dir/cjs_0.js"     testcase.bb
    loops_neg1=$(count_loops "$unroll_static_dir/cjs_neg1.js" 'for (let _i')
    loops_0=$(count_loops "$unroll_static_dir/cjs_0.js"      'for (let _i')
    if [[ "$loops_neg1" -eq 0 && "$loops_0" -gt 0 ]]; then
        ok "CommonJS static: unroll=-1 yields no loops; unroll=0 yields $loops_0 loops"
    else
        fail "CommonJS static: expected 0/positive loops, got neg1=$loops_neg1 / 0=$loops_0"
    fi

    "$BUBBLER" -t mjs -unroll=-1 -single -o "$unroll_static_dir/esm_neg1.js" testcase.bb
    "$BUBBLER" -t mjs -unroll=0  -single -o "$unroll_static_dir/esm_0.js"     testcase.bb
    loops_neg1=$(count_loops "$unroll_static_dir/esm_neg1.js" 'for (let _i')
    loops_0=$(count_loops "$unroll_static_dir/esm_0.js"      'for (let _i')
    if [[ "$loops_neg1" -eq 0 && "$loops_0" -gt 0 ]]; then
        ok "ESModule static: unroll=-1 yields no loops; unroll=0 yields $loops_0 loops"
    else
        fail "ESModule static: expected 0/positive loops, got neg1=$loops_neg1 / 0=$loops_0"
    fi
fi

##############################################################################
# 12. Runtime matrix — per-target CLI option variants must pass codec tests
##############################################################################
echo
echo "=== #12: per-target option variants runtime matrix ==="

if command -v gcc >/dev/null 2>&1; then
    run_variant "C runtime: default"            variant_c default
    run_variant "C runtime: -minimal"           variant_c minimal -minimal
    run_variant "C runtime: -decnum"            variant_c decnum -decnum
    run_variant "C runtime: -signext arith"     variant_c signext_arith -signext arith
    run_variant "C runtime: -unroll=0"          variant_c unroll_0 -unroll 0
else
    ok "C runtime matrix skipped (gcc not found)"
fi

if command -v g++ >/dev/null 2>&1; then
    run_variant "C++ runtime: default"          variant_cpp default
    run_variant "C++ runtime: -minimal"         variant_cpp minimal -minimal
    run_variant "C++ runtime: -decnum"          variant_cpp decnum -decnum
    run_variant "C++ runtime: -signext arith"   variant_cpp signext_arith -signext arith
    run_variant "C++ runtime: -unroll=0"        variant_cpp unroll_0 -unroll 0
    run_variant "C++ runtime combo: multi non-compat"   variant_cpp combo_multi_nocompat
    run_variant "C++ runtime combo: multi compat"       variant_cpp combo_multi_compat -compat
    run_variant "C++ runtime combo: single non-compat"  variant_cpp combo_single_nocompat -single
    run_variant "C++ runtime combo: single compat"      variant_cpp combo_single_compat -single -compat
else
    ok "C++ runtime matrix skipped (g++ not found)"
fi

if command -v go >/dev/null 2>&1; then
    run_variant "Go runtime: default"           variant_go default
    run_variant "Go runtime: -decnum"           variant_go decnum -decnum
    run_variant "Go runtime: -signext arith"    variant_go signext_arith -signext arith
else
    ok "Go runtime matrix skipped (go not found)"
fi

if command -v python3 >/dev/null 2>&1; then
    run_variant "Python runtime: default"       variant_python default
    run_variant "Python runtime: -decnum"       variant_python decnum -decnum
    run_variant "Python runtime: -signext arith" variant_python signext_arith -signext arith
else
    ok "Python runtime matrix skipped (python3 not found)"
fi

if command -v javac >/dev/null 2>&1 && command -v java >/dev/null 2>&1; then
    run_variant "Java runtime: default"         variant_java default
    run_variant "Java runtime: -decnum"         variant_java decnum -decnum
else
    ok "Java runtime matrix skipped (javac/java not found)"
fi

if command -v dotnet >/dev/null 2>&1; then
    run_variant "C# runtime: default"           variant_csharp default
    run_variant "C# runtime: -memcpy"           variant_csharp memcpy -memcpy
else
    ok "C# runtime matrix skipped (dotnet not found)"
fi

if command -v node >/dev/null 2>&1; then
    run_variant "CommonJS runtime: default"     variant_cjs default
    run_variant "CommonJS runtime: -compat"     variant_cjs compat -compat
    run_variant "ESModule runtime: default"     variant_esm default
    run_variant "ESModule runtime: -compat"     variant_esm compat -compat
else
    ok "JS runtime matrices skipped (node not found)"
fi

##############################################################################
# Summary
##############################################################################
echo
echo "=== Compiler Options Tests: ${pass} passed, ${fail} failed ==="
exit "$fail"
