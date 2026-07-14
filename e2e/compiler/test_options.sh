#!/usr/bin/env bash
##############################################################################
# Bubbler E2E — compiler option shape checks
#
# Codec correctness for option combinations lives in the declarative matrix
# at e2e/spec/matrix.go. This script only checks filesystem layout and
# generated API/text properties that cannot be asserted by runtime round trips.
##############################################################################
set -euo pipefail

cd "$(dirname "$0")/.."

if [[ -z "${BUBBLER:-}" ]]; then
    if [[ -x "../bubbler.exe" ]]; then
        BUBBLER="../bubbler.exe"
    else
        BUBBLER="../bubbler"
    fi
fi

work="tests/.tmp_bubbler_opts_$$"
mkdir -p "$work"
cleanup() {
    local status=$?
    if [[ "${KEEP_E2E_WORK:-0}" == "1" || "$status" -ne 0 ]]; then
        echo "[options] keeping work directory: $work"
    else
        rm -rf "$work"
    fi
}
trap cleanup EXIT

pass=0
fail=0

ok() {
    echo "  PASS: $*"
    pass=$((pass + 1))
}

fail_case() {
    echo "  FAIL: $*" >&2
    fail=$((fail + 1))
}

check_exists() {
    [[ -f "$1" ]] && ok "$2" || fail_case "$2: missing $1"
}

check_not_exists() {
    [[ ! -f "$1" ]] && ok "$2" || fail_case "$2: should not exist: $1"
}

check_dir_exists() {
    [[ -d "$1" ]] && ok "$2" || fail_case "$2: missing directory $1"
}

check_contains() {
    local file="$1"
    local pattern="$2"
    local label="$3"
    grep -q -- "$pattern" "$file" && ok "$label" || fail_case "$label: '$pattern' not found in $file"
}

check_not_contains() {
    local file="$1"
    local pattern="$2"
    local label="$3"
    if grep -q -- "$pattern" "$file"; then
        fail_case "$label: '$pattern' should not occur in $file"
    else
        ok "$label"
    fi
}

echo
echo "=== package output paths ==="
out="$work/gopkg"
"$BUBBLER" -t go -o "$out/" features/imports/protocol.bb
check_dir_exists "$out/github.com/example/protocol" "go_package creates nested output path"
check_exists "$out/github.com/example/protocol/protocol.bb.go" "go_package output file"

out="$work/rmpath"
"$BUBBLER" -t go -rmpath github.com/example -o "$out/" features/imports/protocol.bb
check_dir_exists "$out/protocol" "-rmpath strips configured prefix"
check_exists "$out/protocol/protocol.bb.go" "-rmpath output file"
check_not_exists "$out/github.com/example/protocol/protocol.bb.go" "-rmpath removes original nested file"

out="$work/javapkg"
"$BUBBLER" -t java -o "$out/" features/bitwid.bb
check_dir_exists "$out/com/example/bitwid" "java_package creates nested output path"
check_exists "$out/com/example/bitwid/NarrowBWTest.java" "java_package output file"

out="$work/kotlinpkg"
"$BUBBLER" -t kotlin -o "$out/" features/imports/protocol.bb
check_dir_exists "$out/com/example/protocol" "kotlin_package creates nested output path"
check_exists "$out/com/example/protocol/Packet.kt" "kotlin_package struct output file"
check_exists "$out/com/example/types/Direction.kt" "Kotlin imported enum output file"
check_contains "$out/com/example/protocol/Packet.kt" "import com.example.types.*" "Kotlin import uses kotlin_package"

out="$work/kotlinfallback"
"$BUBBLER" -t kotlin -o "$out/" features/kotlin_java_package_independent.bb
check_exists "$out/kotlinfallback/PackageFallback.kt" "Kotlin falls back to .bb package"
check_not_exists "$out/wrong/java/package/PackageFallback.kt" "java_package does not affect Kotlin output"

echo
echo "=== -single ==="
multi="$work/multi"
single="$work/testcase_single.go"
"$BUBBLER" -t go -o "$multi/" testcase.bb
"$BUBBLER" -t go -single -o "$single" testcase.bb
check_exists "$multi/testpkg/testpkg.bb.go" "multi-file Go output"
check_exists "$single" "single-file Go output"
check_not_exists "$work/testpkg/testpkg.bb.go" "single-file mode does not create package directory"

echo
echo "=== -minimal ==="
minimal="$work/minimal"
full="$work/full"
"$BUBBLER" -t c -minimal -o "$minimal/" testcase.bb
"$BUBBLER" -t c -o "$full/" testcase.bb
check_not_contains "$minimal/testpkg.bb.h" "RawGetter" "-minimal removes raw getters"
check_not_contains "$minimal/testpkg.bb.h" "RawSetter" "-minimal removes raw setters"
check_contains "$minimal/testpkg.bb.h" "get_voltage" "-minimal preserves custom getter"
check_contains "$full/testpkg.bb.h" "RawGetter" "default mode emits raw getters"

java_minimal="$work/java_minimal"
python_minimal="$work/python_minimal.py"
cs_minimal="$work/csharp_minimal"
kotlin_minimal="$work/kotlin_minimal"
"$BUBBLER" -t java -minimal -o "$java_minimal/" testcase.bb
"$BUBBLER" -t kotlin -minimal -o "$kotlin_minimal/" testcase.bb
"$BUBBLER" -t py -minimal -single -o "$python_minimal" testcase.bb
mkdir -p "$cs_minimal"
"$BUBBLER" -t cs -minimal -single -o "$cs_minimal/testcase.bb.cs" testcase.bb
if command -v javac >/dev/null 2>&1; then
    mapfile -t java_sources < <(find "$java_minimal" -name '*.java' -print)
    javac -encoding UTF-8 -d "$work/java_minimal_out" "${java_sources[@]}"
    ok "Java -minimal generated sources compile"
else
    ok "Java -minimal compile skipped (javac not found)"
fi
check_contains "$kotlin_minimal/com/example/testpkg/GetterSetter.kt" "private var rawAdc" "Kotlin -minimal hides raw property"
check_contains "$kotlin_minimal/com/example/testpkg/GetterSetter.kt" "var voltage: Double" "Kotlin -minimal preserves custom property"
if command -v kotlinc >/dev/null 2>&1; then
    mapfile -t kotlin_sources < <(find "$kotlin_minimal" -name '*.kt' -print)
    kotlinc "${kotlin_sources[@]}" -jvm-target 17 -d "$work/kotlin_minimal.jar"
    ok "Kotlin -minimal generated sources compile"
else
    fail_case "Kotlin compiler not found; Kotlin option tests must run with Kotlin 1.9.24"
fi

kotlin_getter_only="$work/kotlin_getter_only"
"$BUBBLER" -t kotlin -o "$kotlin_getter_only/" features/kotlin_getter_only.bb
check_contains "$kotlin_getter_only/kotlinaccessor/GetterOnly.kt" "val doubled: UInt" "Kotlin getter-only accessor generates val"
check_not_contains "$kotlin_getter_only/kotlinaccessor/GetterOnly.kt" "var doubled: UInt" "Kotlin getter-only accessor is not mutable"
if command -v python3 >/dev/null 2>&1; then
    python3 -m py_compile "$python_minimal"
    ok "Python -minimal generated source compiles"
elif command -v python >/dev/null 2>&1; then
    python -m py_compile "$python_minimal"
    ok "Python -minimal generated source compiles"
else
    ok "Python -minimal compile skipped (python not found)"
fi
if command -v dotnet >/dev/null 2>&1 && [[ "$(dotnet --version | cut -d. -f1)" -ge 8 ]]; then
    cp framework/csharp_library.csproj "$cs_minimal/csharp_library.csproj"
    dotnet build "$cs_minimal/csharp_library.csproj" --nologo --verbosity quiet
    ok "C# -minimal generated source compiles on .NET 8"
else
    ok "C# -minimal compile skipped (.NET 8 SDK not found)"
fi

echo
echo "=== -decnum ==="
decimal="$work/decnum"
hex="$work/hexnum"
"$BUBBLER" -t c -decnum -o "$decimal/" testcase.bb
"$BUBBLER" -t c -o "$hex/" testcase.bb
check_contains "$hex/testpkg.bb.c" "0xAA" "default constants preserve hexadecimal form"
check_contains "$decimal/testpkg.bb.c" "170" "-decnum emits decimal constant"
check_not_contains "$decimal/testpkg.bb.c" "0xAA" "-decnum removes hexadecimal constant"
check_not_contains "$decimal/testpkg.bb.c" "0b11111111" "-decnum removes binary masks"

kotlin_decimal="$work/kotlin_decnum"
"$BUBBLER" -t kotlin -decnum -o "$kotlin_decimal/" testcase.bb
check_not_contains "$kotlin_decimal/com/example/testpkg/ConstantFields.kt" "0xAA" "Kotlin -decnum removes hexadecimal constant"

echo
echo "=== omit_empty and imports ==="
omit="$work/omit"
"$BUBBLER" -t c -o "$omit/" features/omit_e.bb
check_exists "$omit/types.bb.h" "omit_empty still generates imported definitions"
check_not_exists "$omit/importonly.bb.h" "omit_empty suppresses empty package"

imports="$work/imports"
"$BUBBLER" -t c -o "$imports/" features/imports/protocol.bb
check_exists "$imports/types.bb.h" "imported header generated"
check_exists "$imports/types.bb.c" "imported source generated"
check_exists "$imports/protocol.bb.h" "entry header generated"
check_exists "$imports/protocol.bb.c" "entry source generated"
check_contains "$imports/protocol.bb.h" "Vec2" "imported struct is referenced"
check_contains "$imports/protocol.bb.h" "Direction" "imported enum is referenced"

echo
echo "=== Kotlin setter-only accessor rejection ==="
if "$BUBBLER" -t kotlin -o "$work/kotlin_setter_only/" features/kotlin_setter_only.bb >"$work/kotlin_setter_only.out" 2>&1; then
    fail_case "Kotlin setter-only accessor should fail generation"
else
    check_contains "$work/kotlin_setter_only.out" "Kotlin target cannot generate setter-only custom property \"calibrated\" on field \"raw\"" "Kotlin setter-only error identifies accessor and field"
    check_contains "$work/kotlin_setter_only.out" "does not support write-only properties" "Kotlin setter-only error explains language limitation"
fi

echo
echo "=== -signext and -compat ==="
shift="$work/signext_shift"
arith="$work/signext_arith"
"$BUBBLER" -t c -signext shift -o "$shift/" testcase.bb
"$BUBBLER" -t c -signext arith -o "$arith/" testcase.bb
check_exists "$shift/testpkg.bb.c" "-signext=shift generates code"
check_exists "$arith/testpkg.bb.c" "-signext=arith generates code"

cjs_default="$work/cjs_default"
cjs_compat="$work/cjs_compat"
"$BUBBLER" -t cjs -single -o "$cjs_default/testcase.bb.js" testcase.bb
"$BUBBLER" -t cjs -compat -single -o "$cjs_compat/testcase.bb.js" testcase.bb
check_contains "$cjs_default/testcase.bb.js" "new Uint8Array(" "CommonJS default uses Uint8Array"
check_not_contains "$cjs_default/testcase.bb.js" "if (data === undefined) data = new Array(" "CommonJS default does not allocate Array"
check_contains "$cjs_compat/testcase.bb.js" "if (data === undefined) data = new Array(" "CommonJS -compat uses Array"
check_not_contains "$cjs_compat/testcase.bb.js" "if (data === undefined) data = new Uint8Array(" "CommonJS -compat does not allocate Uint8Array"

echo
echo "=== compiler option shape checks: ${pass} passed, ${fail} failed ==="
exit "$fail"
