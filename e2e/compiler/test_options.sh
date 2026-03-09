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

BUBBLER="${BUBBLER:-../bubbler}"
TMPDIR_BASE="${TMPDIR:-/tmp}/bubbler_opts_$$"
mkdir -p "$TMPDIR_BASE"
trap 'rm -rf "$TMPDIR_BASE"' EXIT

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
# Summary
##############################################################################
echo
echo "=== Compiler Options Tests: ${pass} passed, ${fail} failed ==="
exit "$fail"
