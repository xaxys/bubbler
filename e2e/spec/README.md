# Bubbler unified E2E spec

`e2e/spec` is the single source of truth for codec tests. It owns:

- language-neutral values and scenarios (`spec.go`, `scenarios.go`);
- generated test drivers for C, C++, Go, Java, Python, C#, CommonJS and ESM;
- the language/toolchain adapters used to generate, compile and run code;
- generator option variants and the `-unroll` threshold matrix;
- structural checks proving loop counts follow the configured threshold.

The shell scripts under `e2e/` are only suite launchers and compiler
shape/error checks. They do not contain per-language codec cases.

## Run

Generate one compatible driver (the original interface is preserved):

```sh
go run ./e2e/spec -lang=c -out=e2e/tests/c/main.c
go run ./e2e/spec -lang=go -out=e2e/tests/go
```

Run the unified matrix:

```sh
go run ./e2e/spec -matrix -bubbler=./bubbler -profile=full
```

Profiles:

- `smoke`: default generator options once per language;
- `unroll`: all declared unroll boundary values;
- `options`: default plus supported per-language option variants;
- `full`: smoke, options and unroll coverage (used by `make e2e`).

Use `-langs=c,go,python` to select languages and `-keep-work` to retain the
generated work directory. Existing `SKIP_C`, `SKIP_CPP`, `SKIP_GO`,
`SKIP_PY`, `SKIP_JAVA`, `SKIP_CS`, `SKIP_CJS`/`SKIP_JS` and `SKIP_ESM`
environment variables are supported.

Linux runs C# on .NET 8 only. .NET Framework 4.7.2 remains a Windows CI
job because it cannot be executed natively in the Linux matrix.

## Add a codec case

1. Add or reuse a struct in `e2e/testcase.bb` or `e2e/features/bitwid.bb`.
2. Add one `Scenario` in `scenarios.go` and register it in `AllScenarios()`.
3. Describe input fields with `Setup`; omit `Assert` to assert the same values
   after decode, or provide explicit assertions/tolerances.
4. Add `Wire` when the canonical byte layout is known. Golden wire assertions
   catch symmetric encoder/decoder bugs that a round trip alone cannot find.
5. Add `DecodeSizeChecks` and `DecodeError` probes for malformed/truncated
   input where relevant.
6. Run the smoke profile while iterating, then `make e2e` before merging.

The `CodecArrays` scenario is the reference for array coverage. It includes
bool, signed/unsigned integers, floats, enum/struct arrays, strings, bytes,
dynamic structs, UTF-8, empty values and the 127/128 byte length boundary.

## Unroll coverage

The matrix runs thresholds `-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 31, 32`.
Together with array lengths `1, 2, 3, 4, 5, 7, 8, 32`, this covers both sides
of every important threshold and the exact equality boundary.

Every threshold performs full encode/decode/decode-size tests. The runner also
counts language-specific generated loop markers and asserts:

- `-unroll=-1` emits no array loops;
- `-unroll=0` emits loops;
- loop counts never increase as the threshold increases;
- at least one boundary changes the generated loop count.
