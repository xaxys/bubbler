#!/usr/bin/env bash
##############################################################################
# Bubbler E2E — Compiler Error Tests
#
# Negative tests: each .bb file is expected to FAIL compilation.
# Verifies the compiler exits with a non-zero code and optionally that the
# error message matches the expected pattern.
#
# Run from the e2e/ directory:   bash compiler/test_errors.sh
#
# Environment:
#   BUBBLER  — Path to the bubbler binary (default: ../bubbler)
##############################################################################
set -euo pipefail
cd "$(dirname "$0")/.."   # ensure we run from e2e/

if [[ -z "${BUBBLER:-}" ]]; then
    if [[ -x "../bubbler.exe" ]]; then
        BUBBLER="../bubbler.exe"
    else
        BUBBLER="../bubbler"
    fi
fi
TMPDIR_BASE="tests/.tmp_bubbler_errs_$$"
mkdir -p "$TMPDIR_BASE"
trap 'rm -rf "$TMPDIR_BASE"' EXIT

pass=0
fail=0

##############################################################################
# Helper: run compiler, expect failure (exit != 0), check optional message
# Usage: expect_fail  <label>  <optional_pattern>  [-- bubbler args...]
##############################################################################
expect_fail() {
    local label="$1"; shift
    local pattern="$1"; shift

    local out code
    out=$("$BUBBLER" "$@" 2>&1) && code=$? || code=$?

    if [[ $code -eq 0 ]]; then
        echo "  FAIL: ${label}: expected compile error but compiler succeeded" >&2
        fail=$((fail + 1))
        return
    fi

    if [[ -n "$pattern" ]] && ! echo "$out" | grep -qiF "$pattern"; then
        echo "  FAIL: ${label}: exit=${code} but error message did not contain '${pattern}'" >&2
        echo "        actual output: $out" >&2
        fail=$((fail + 1))
    else
        echo "  PASS: ${label}"
        pass=$((pass + 1))
    fi
}

##############################################################################
# 1. Duplicate struct name in same file
##############################################################################
echo
echo "=== #1: duplicate struct name ==="
expect_fail \
    "dup_struct: duplicate definition" \
    "duplicate definition" \
    -t c -o "$TMPDIR_BASE/dup_struct/" errors/dup_struct.bb

##############################################################################
# 2. Duplicate option in same file
##############################################################################
echo
echo "=== #2: duplicate option ==="
expect_fail \
    "dup_option: option already set" \
    "already been set" \
    -t c -o "$TMPDIR_BASE/dup_option/" errors/dup_option.bb

##############################################################################
# 3. Circular import  (a.bb → b.bb → a.bb)
##############################################################################
echo
echo "=== #3: circular import ==="
expect_fail \
    "cycle: import cycle detected" \
    "import cycle" \
    -t c -o "$TMPDIR_BASE/cycle/" errors/cycle/a.bb

##############################################################################
# 4. Duplicate package name via imports
##############################################################################
echo
echo "=== #4: duplicate package name ==="
expect_fail \
    "dup_pkg: duplicate package name" \
    "duplicate package name" \
    -t c -o "$TMPDIR_BASE/dup_pkg/" errors/dup_pkg/entry.bb

##############################################################################
# Summary
##############################################################################
echo
echo "=== Compiler Error Tests: ${pass} passed, ${fail} failed ==="
exit "$fail"
