#!/usr/bin/env bash
set -euo pipefail

# Empty schema codegen matrix test:
# 1) code generation succeeds for all targets
# 2) generated code has sane formatting shape (especially ESM export block)
# 3) generated code passes compile/syntax smoke checks

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
E2E_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
PROJECT_ROOT="$(cd "$E2E_DIR/.." && pwd)"
cd "$E2E_DIR"

if [[ -z "${BUBBLER:-}" ]]; then
  if [[ -f "$PROJECT_ROOT/bubbler.exe" ]]; then
    BUBBLER="$PROJECT_ROOT/bubbler.exe"
  elif [[ -x "$PROJECT_ROOT/bubbler" ]]; then
        BUBBLER="$PROJECT_ROOT/bubbler"
    else
        BUBBLER="bubbler"
    fi
fi

WORK_DIR="tests/empty_matrix"
SRC="features/empty.bb"

rm -rf "$WORK_DIR"
mkdir -p "$WORK_DIR"/{c,cpp,go,java,py,cs,cjs,mjs}

echo "[empty] generate all targets"
"$BUBBLER" -t c   -o "$WORK_DIR/c"                  "$SRC"
"$BUBBLER" -t cpp -single -o "$WORK_DIR/cpp/empty.bb.hpp" "$SRC"
"$BUBBLER" -t go  -o "$WORK_DIR/go"                 "$SRC"
"$BUBBLER" -t java -o "$WORK_DIR/java"              "$SRC"
"$BUBBLER" -t py  -single -o "$WORK_DIR/py/empty_bb.py"    "$SRC"
"$BUBBLER" -t cs  -single -o "$WORK_DIR/cs/empty.bb.cs"    "$SRC"
"$BUBBLER" -t cjs -single -o "$WORK_DIR/cjs/empty.bb.js"   "$SRC"
"$BUBBLER" -t mjs -single -o "$WORK_DIR/mjs/empty.bb.js"   "$SRC"

echo "[empty] C compile smoke"
if [[ "${SKIP_C:-0}" == "1" ]]; then
  echo "  [SKIP] C"
elif ! command -v gcc >/dev/null 2>&1; then
  echo "  [SKIP] C (gcc not found)"
else
  gcc -std=c11 -I"$WORK_DIR/c" -c "$WORK_DIR/c/emptypkg.bb.c" -o "$WORK_DIR/c/emptypkg.o"
fi

echo "[empty] C++ compile smoke"
cat > "$WORK_DIR/cpp/main.cpp" <<'CPP'
#include "empty.bb.hpp"
int main() { return 0; }
CPP
if [[ "${SKIP_CPP:-0}" == "1" ]]; then
  echo "  [SKIP] C++"
elif ! command -v g++ >/dev/null 2>&1; then
  echo "  [SKIP] C++ (g++ not found)"
else
  g++ -std=c++17 -I"$WORK_DIR/cpp" -o "$WORK_DIR/cpp/run_test" "$WORK_DIR/cpp/main.cpp"
  "$WORK_DIR/cpp/run_test"
fi

echo "[empty] Go compile smoke"
cat > "$WORK_DIR/go/go.mod" <<'GOMOD'
module empty_codegen_smoke

go 1.20
GOMOD
if [[ "${SKIP_GO:-0}" == "1" ]]; then
  echo "  [SKIP] Go"
elif ! command -v go >/dev/null 2>&1; then
  echo "  [SKIP] Go (go not found)"
else
  (
    cd "$WORK_DIR/go"
    go test ./... -v
  )
fi

echo "[empty] Java empty-output policy check"
# Current behavior: empty schema with no types generates no Java file.
# We assert the behavior so regressions are visible.
if find "$WORK_DIR/java" -type f | grep -q .; then
    echo "[FAIL] Java empty schema should not emit files under current policy" >&2
    exit 1
fi

echo "[empty] Python syntax smoke"
if [[ "${SKIP_PY:-0}" == "1" ]]; then
  echo "  [SKIP] Python"
elif command -v python3 >/dev/null 2>&1; then
  python3 -m py_compile "$WORK_DIR/py/empty_bb.py"
elif command -v python >/dev/null 2>&1; then
  python -m py_compile "$WORK_DIR/py/empty_bb.py"
else
  echo "  [SKIP] Python (python not found)"
fi

echo "[empty] C# compile smoke"
cat > "$WORK_DIR/cs/empty.csproj" <<'CSPROJ'
<Project Sdk="Microsoft.NET.Sdk">
  <PropertyGroup>
    <TargetFramework>net8.0</TargetFramework>
    <ImplicitUsings>disable</ImplicitUsings>
    <Nullable>disable</Nullable>
    <EnableDefaultCompileItems>false</EnableDefaultCompileItems>
  </PropertyGroup>
  <ItemGroup>
    <Compile Include="empty.bb.cs" />
  </ItemGroup>
</Project>
CSPROJ
if [[ "${SKIP_CS:-0}" == "1" ]]; then
  echo "  [SKIP] C#"
elif ! command -v dotnet >/dev/null 2>&1; then
  echo "  [SKIP] C# (dotnet not found)"
else
  dotnet build "$WORK_DIR/cs/empty.csproj" -nologo -v minimal > /dev/null
fi

echo "[empty] CommonJS syntax/runtime smoke"
if [[ -n "${SKIP_JS:-}" && -z "${SKIP_CJS:-}" ]]; then
  SKIP_CJS="$SKIP_JS"
fi
if [[ "${SKIP_CJS:-0}" == "1" ]]; then
  echo "  [SKIP] CommonJS (cjs)"
elif ! command -v node >/dev/null 2>&1; then
  echo "  [SKIP] CommonJS (cjs, node not found)"
else
  node --check "$WORK_DIR/cjs/empty.bb.js"
  node -e "const m=require('./tests/empty_matrix/cjs/empty.bb.js'); if(!m.emptypkg) process.exit(2); if(Object.keys(m.emptypkg).length!==0) process.exit(3);"
fi

echo "[empty] ESModule syntax/runtime + formatting smoke"
if [[ "${SKIP_ESM:-0}" == "1" ]]; then
  echo "  [SKIP] ESModule"
elif ! command -v node >/dev/null 2>&1; then
  echo "  [SKIP] ESModule (node not found)"
else
  cp "$WORK_DIR/mjs/empty.bb.js" "$WORK_DIR/mjs/empty.bb.mjs"
  node --check "$WORK_DIR/mjs/empty.bb.mjs"
  node --input-type=module -e "import('./tests/empty_matrix/mjs/empty.bb.mjs').then((m)=>{const r=m.default; if(!r.emptypkg) process.exit(2); if(Object.keys(r.emptypkg).length!==0) process.exit(3);}).catch(()=>process.exit(4));"
  grep -Eq '^export default \{$' "$WORK_DIR/mjs/empty.bb.js"
  grep -Eq '^[[:space:]]*emptypkg: \{$' "$WORK_DIR/mjs/empty.bb.js"
fi

echo "[empty] PASS"
