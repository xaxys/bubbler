#!/usr/bin/env bash
##############################################################################
# Bubbler E2E — Master Test Runner
#
# Generates code for all language targets, runs all language tests, and runs
# the compiler option + error tests.
#
# Usage:
#   bash e2e/run_tests.sh           # from project root
#   bash run_tests.sh               # from e2e/ directory
#
# Environment:
#   BUBBLER   — Path to the bubbler binary (default: auto-detected)
#   SKIP_C    — Set to 1 to skip C tests
#   SKIP_CPP  — Set to 1 to skip C++ tests
#   SKIP_GO   — Set to 1 to skip Go tests
#   SKIP_PY   — Set to 1 to skip Python tests
#   SKIP_JAVA — Set to 1 to skip Java tests
#   SKIP_CS   — Set to 1 to skip C# tests
#   SKIP_JS   — Set to 1 to skip JS tests
##############################################################################
set -euo pipefail

# ─── locate directories ─────────────────────────────────────────────────────
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
E2E_DIR="$SCRIPT_DIR"

# Support running from project root ( bash e2e/run_tests.sh )
if [[ "$(basename "$SCRIPT_DIR")" != "e2e" ]]; then
    E2E_DIR="$SCRIPT_DIR/e2e"
fi

PROJECT_ROOT="$(cd "$E2E_DIR/.." && pwd)"
cd "$E2E_DIR"

# ─── locate bubbler binary ───────────────────────────────────────────────────
if [[ -z "${BUBBLER:-}" ]]; then
    if [[ -x "$PROJECT_ROOT/bubbler" ]]; then
        BUBBLER="$PROJECT_ROOT/bubbler"
    elif command -v bubbler &>/dev/null; then
        BUBBLER="bubbler"
    else
        # Try building from source
        echo "[setup] Building bubbler from source..."
        (cd "$PROJECT_ROOT" && go build -o bubbler .)
        BUBBLER="$PROJECT_ROOT/bubbler"
    fi
fi
echo "[setup] Using bubbler: $BUBBLER"
echo "[setup] E2E root:      $E2E_DIR"

# ─── test tracking ───────────────────────────────────────────────────────────
total_pass=0
total_fail=0

section_pass() {
    local name="$1" exit_code="$2"
    if [[ "$exit_code" -eq 0 ]]; then
        echo "  [OK]  $name"
        total_pass=$((total_pass + 1))
    else
        echo "  [FAIL] $name (exit=$exit_code)" >&2
        total_fail=$((total_fail + 1))
    fi
}

# ─── directories ─────────────────────────────────────────────────────────────
C_DIR="$E2E_DIR/tests/c"
CPP_DIR="$E2E_DIR/tests/cpp"
GO_DIR="$E2E_DIR/tests/go"
PY_DIR="$E2E_DIR/tests/python"
JAVA_DIR="$E2E_DIR/tests/java"
CS_DIR="$E2E_DIR/tests/csharp"
JS_DIR="$E2E_DIR/tests/js"

mkdir -p "$C_DIR/gen" "$CPP_DIR/gen" "$JAVA_DIR/gen" "$PY_DIR/gen" "$CS_DIR/gen" "$JS_DIR/gen"

##############################################################################
# CODE GENERATION
##############################################################################
echo
echo "════════════════════════════════════════════"
echo " Code Generation"
echo "════════════════════════════════════════════"

# — testcase.bb for all targets —
echo "[gen] testcase.bb"
"$BUBBLER" -t c   -o "$C_DIR/gen/"                              testcase.bb
"$BUBBLER" -t cpp -single -o "$CPP_DIR/gen/testcase.bb.hpp"    testcase.bb
"$BUBBLER" -t go  -o "$GO_DIR/"                                 testcase.bb
"$BUBBLER" -t java        -o "$JAVA_DIR/gen/"                   testcase.bb
"$BUBBLER" -t py  -single -o "$PY_DIR/gen/testcase_bb.py"      testcase.bb
"$BUBBLER" -t cs  -single -o "$CS_DIR/gen/testcase.bb.cs"      testcase.bb
"$BUBBLER" -t cjs -single -o "$JS_DIR/gen/testcase.bb.js"      testcase.bb

# — bitwid.bb (narrow array feature) for all targets —
echo "[gen] bitwid.bb"
"$BUBBLER" -t c   -o "$C_DIR/gen/"                              features/bitwid.bb
"$BUBBLER" -t cpp -single -o "$CPP_DIR/gen/bitwid.bb.hpp"      features/bitwid.bb
"$BUBBLER" -t go  -o "$GO_DIR/"                                 features/bitwid.bb
"$BUBBLER" -t java        -o "$JAVA_DIR/gen/"                   features/bitwid.bb
"$BUBBLER" -t py  -single -o "$PY_DIR/gen/bitwid_bb.py"        features/bitwid.bb
"$BUBBLER" -t cs  -single -o "$CS_DIR/gen/bitwid.bb.cs"        features/bitwid.bb
"$BUBBLER" -t cjs -single -o "$JS_DIR/gen/bitwid.bb.js"        features/bitwid.bb

echo "[gen] done"

##############################################################################
# LANGUAGE TESTS
##############################################################################
echo
echo "════════════════════════════════════════════"
echo " Language Codec Tests"
echo "════════════════════════════════════════════"

# ── C ────────────────────────────────────────────────────────────────────────
if [[ "${SKIP_C:-0}" != "1" ]]; then
    echo
    echo "--- C ---"
    (
        cd "$C_DIR"
        gcc -std=c11 -Igen -o run_test main.c gen/testpkg.bb.c gen/bitwid.bb.c -lm
        ./run_test
    )
    section_pass "C" "$?"
else
    echo "  [SKIP] C"
fi

# ── C++ ──────────────────────────────────────────────────────────────────────
if [[ "${SKIP_CPP:-0}" != "1" ]]; then
    echo
    echo "--- C++ ---"
    (
        cd "$CPP_DIR"
        g++ -std=c++17 -Igen -o run_test main.cpp -lm
        ./run_test
    )
    section_pass "C++" "$?"
else
    echo "  [SKIP] C++"
fi

# ── Go ───────────────────────────────────────────────────────────────────────
if [[ "${SKIP_GO:-0}" != "1" ]]; then
    echo
    echo "--- Go ---"
    (
        cd "$GO_DIR"
        go test ./... -v
    )
    section_pass "Go" "$?"
else
    echo "  [SKIP] Go"
fi

# ── Python ───────────────────────────────────────────────────────────────────
if [[ "${SKIP_PY:-0}" != "1" ]]; then
    echo
    echo "--- Python ---"
    (
        cd "$PY_DIR"
        python3 test_main.py
    )
    section_pass "Python" "$?"
else
    echo "  [SKIP] Python"
fi

# ── Java ─────────────────────────────────────────────────────────────────────
if [[ "${SKIP_JAVA:-0}" != "1" ]]; then
    echo
    echo "--- Java ---"
    (
        cd "$JAVA_DIR"
        rm -rf out && mkdir out
        find gen -name "*.java" -exec javac -encoding UTF-8 -d out {} +
        javac -encoding UTF-8 -cp out -d out Main.java
        java -cp out Main
    )
    section_pass "Java" "$?"
else
    echo "  [SKIP] Java"
fi

# ── C# (net8.0) ─────────────────────────────────────────────────────────────
if [[ "${SKIP_CS:-0}" != "1" ]]; then
    echo
    echo "--- C# (net8.0) ---"
    (
        cd "$CS_DIR"
        dotnet run -f net8.0 --project test.csproj
    )
    section_pass "C# net8.0" "$?"

    echo
    echo "--- C# (net472 / Mono) ---"
    if [[ "$(uname)" == "Linux" ]]; then
        echo "  [SKIP] C# net472 (Linux: net472 requires Windows; tested in CI Windows runner)"
    else
        (
            cd "$CS_DIR"
            dotnet build test.csproj -f net472 -c Release --nologo -q
            mono bin/Release/net472/test.exe
        )
        section_pass "C# net472" "$?"
    fi
else
    echo "  [SKIP] C#"
fi

# ── JavaScript (CommonJS) ────────────────────────────────────────────────────
if [[ "${SKIP_JS:-0}" != "1" ]]; then
    echo
    echo "--- JavaScript ---"
    (
        cd "$JS_DIR"
        node test.mjs
    )
    section_pass "JavaScript" "$?"
else
    echo "  [SKIP] JavaScript"
fi

##############################################################################
# COMPILER OPTION TESTS
##############################################################################
echo
echo "════════════════════════════════════════════"
echo " Compiler Option Tests"
echo "════════════════════════════════════════════"
(
    export BUBBLER
    bash compiler/test_options.sh
)
section_pass "Compiler Options" "$?"

##############################################################################
# COMPILER ERROR TESTS
##############################################################################
echo
echo "════════════════════════════════════════════"
echo " Compiler Error Tests (negative)"
echo "════════════════════════════════════════════"
(
    export BUBBLER
    bash compiler/test_errors.sh
)
section_pass "Compiler Errors" "$?"

##############################################################################
# SUMMARY
##############################################################################
echo
echo "════════════════════════════════════════════"
if [[ "$total_fail" -eq 0 ]]; then
    echo " ALL PASSED: ${total_pass} test suites"
else
    echo " RESULT: ${total_pass} passed, ${total_fail} FAILED"
fi
echo "════════════════════════════════════════════"

exit "$total_fail"
