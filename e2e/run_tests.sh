#!/usr/bin/env bash
##############################################################################
# Bubbler E2E — unified entry point
#
# Codec scenarios, language adapters, option variants and unroll thresholds
# are declared in e2e/spec. Shell remains only as a portable suite launcher.
##############################################################################
set -uo pipefail

script_dir="$(cd "$(dirname "$0")" && pwd)"
project_root="$(cd "$script_dir/.." && pwd)"

if [[ -z "${BUBBLER:-}" ]]; then
    if [[ -x "$project_root/bubbler.exe" ]]; then
        BUBBLER="$project_root/bubbler.exe"
    elif [[ -x "$project_root/bubbler" ]]; then
        BUBBLER="$project_root/bubbler"
    elif command -v bubbler >/dev/null 2>&1; then
        BUBBLER="bubbler"
    else
        echo "[setup] Building bubbler from source..."
        (cd "$project_root" && go build -o bubbler .) || exit 1
        BUBBLER="$project_root/bubbler"
    fi
fi

echo "[setup] Using bubbler: $BUBBLER"
echo "[setup] Project root:  $project_root"

pass=0
fail=0

run_suite() {
    local name="$1"
    shift
    echo
    echo "════════════════════════════════════════════"
    echo " $name"
    echo "════════════════════════════════════════════"
    if "$@"; then
        echo "  [OK] $name"
        pass=$((pass + 1))
    else
        local code=$?
        echo "  [FAIL] $name (exit=$code)" >&2
        fail=$((fail + 1))
    fi
}

run_suite \
    "Unified Codec / Options / Unroll Matrix" \
    bash -c 'cd "$1" && exec go run ./e2e/spec -matrix -profile full -bubbler "$2"' \
    _ "$project_root" "$BUBBLER"

run_suite \
    "Empty Schema Codegen Matrix" \
    bash -c 'cd "$1/e2e" && BUBBLER="$2" exec bash framework/test_empty_codegen.sh' \
    _ "$project_root" "$BUBBLER"

run_suite \
    "Compiler Option Shape Checks" \
    bash -c 'cd "$1/e2e" && BUBBLER="$2" exec bash compiler/test_options.sh' \
    _ "$project_root" "$BUBBLER"

run_suite \
    "Compiler Negative Tests" \
    bash -c 'cd "$1/e2e" && BUBBLER="$2" exec bash compiler/test_errors.sh' \
    _ "$project_root" "$BUBBLER"

echo
echo "════════════════════════════════════════════"
if [[ "$fail" -eq 0 ]]; then
    echo " ALL PASSED: $pass suites"
else
    echo " RESULT: $pass passed, $fail failed"
fi
echo "════════════════════════════════════════════"

exit "$fail"
