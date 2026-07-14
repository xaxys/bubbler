// Package main is the bubbler e2e test driver generator.
//
// Single source of truth for all 7 language test drivers: scenarios.go
// declares test cases as language-neutral data, and the emit_<lang>.go
// files render that data into per-language source files.
//
// Run: go run ./e2e/spec -lang=<lang> -out=<path>
package main

import (
    "fmt"
    "math"
    "strings"
)

// ValKind tags every primitive carried by Val. Arrays carry an inner
// element kind in Val.ElemKind.
type ValKind int

const (
    VKUnknown ValKind = iota
    VKBool
    VKU8
    VKU16
    VKU32
    VKU64
    VKI8
    VKI16
    VKI32
    VKI64
    VKF32
    VKF64
    VKString
    VKBytes
    VKEnum
    VKStruct
    VKArray
)

// Val is a language-neutral literal carried through the spec. Each emitter
// renders it to a target-language expression. Floats store raw IEEE bits in
// U so NaN payloads and signed zeros are preserved exactly.
type Val struct {
    Kind     ValKind
    U        uint64
    I        int64
    Str      string
    Bytes    []byte
    EnumType string
    EnumName string
    Struct   *StructLit
    Array    *ArrayLit
}

type StructLit struct {
    Type   string
    Fields []FieldVal
}

type ArrayLit struct {
    ElemKind ValKind
    ElemType string
    Items    []Val
}

type FieldVal struct {
    Name string
    V    Val
}

// AccessorRoundTrip drives the GetterSetter struct: call setter, then either
// just call getter (NoRoundTrip) or encode+decode then call getter on the
// decoded struct.
type AccessorRoundTrip struct {
    SetterName string
    SetArg     Val
    GetterName string
    Expected   Val
    Tol        float64
    NoRoundTrip bool
}

type Case struct {
    Name        string
    Setup       []FieldVal
    Assert      []AssertField
    Wire        []byte
    Errors      []DecodeError
    Accessors   []AccessorRoundTrip
}

type AssertField struct {
    Name string
    V    Val
    // Tol is interpreted in language-specific terms: 0 means exact
    // (CHECK_EQ for non-floats); positive means CHECK_NEAR with eps.
    // For float fields, leaving Tol == 0 selects bit-exact comparison.
    Tol float64
}

type BytePatch struct {
    Offset int
    Value  uint8
}

type DecodeError struct {
    Name        string
    Patches     []BytePatch
    ExpectedRet int64 // -1 (any failure) or specific negative value
}

type DecodeSizeCheck struct {
    Name     string
    Bytes    []byte
    Expected int64
    // SourceCase: if non-empty, the bytes are produced by encoding the case
    // with the given name (instead of the literal Bytes). Useful for
    // "complete packet" / "truncate one byte" checks.
    SourceCase string
    Truncate   int // optional: if > 0 and SourceCase set, drop this many bytes from the end
}

type Scenario struct {
    BBFile           string // logical group key: "testcase" or "bitwid"
    Package          string // bb package: "testpkg" / "bitwid"
    StructName       string
    IsDynamic        bool   // true = string/bytes-bearing; uses encode_size at runtime
    Notes            string
    Cases            []Case
    DecodeSizeChecks []DecodeSizeCheck
}

// Spec is the complete scenario set.
type Spec struct {
    Scenarios []Scenario
}

// ──────────────────────────────────────────────────────────────────────────
// Constructors
// ──────────────────────────────────────────────────────────────────────────

func B(b bool) Val {
    var u uint64
    if b {
        u = 1
    }
    return Val{Kind: VKBool, U: u}
}

func U8(x uint8) Val   { return Val{Kind: VKU8, U: uint64(x)} }
func U16(x uint16) Val { return Val{Kind: VKU16, U: uint64(x)} }
func U32(x uint32) Val { return Val{Kind: VKU32, U: uint64(x)} }
func U64(x uint64) Val { return Val{Kind: VKU64, U: x} }

func I8(x int8) Val   { return Val{Kind: VKI8, I: int64(x)} }
func I16(x int16) Val { return Val{Kind: VKI16, I: int64(x)} }
func I32(x int32) Val { return Val{Kind: VKI32, I: int64(x)} }
func I64(x int64) Val { return Val{Kind: VKI64, I: x} }

// F32 stores the raw IEEE-754 bit pattern so NaN payloads / signed zeros
// survive emission. Use F32Bits when you want to spell the bits explicitly.
func F32(x float32) Val   { return F32Bits(math.Float32bits(x)) }
func F32Bits(b uint32) Val { return Val{Kind: VKF32, U: uint64(b)} }
func F64(x float64) Val   { return F64Bits(math.Float64bits(x)) }
func F64Bits(b uint64) Val { return Val{Kind: VKF64, U: b} }

func S(s string) Val      { return Val{Kind: VKString, Str: s} }
func Bs(b []byte) Val     { return Val{Kind: VKBytes, Bytes: append([]byte(nil), b...)} }
func E(typ, name string) Val {
    return Val{Kind: VKEnum, EnumType: typ, EnumName: name}
}

func Struct(typ string, fs ...FieldVal) Val {
    return Val{Kind: VKStruct, Struct: &StructLit{Type: typ, Fields: fs}}
}

func F(name string, v Val) FieldVal { return FieldVal{Name: name, V: v} }

func A(elemKind ValKind, elemType string, items ...Val) Val {
    return Val{Kind: VKArray, Array: &ArrayLit{ElemKind: elemKind, ElemType: elemType, Items: items}}
}

// AssertEq promotes a Setup field-value pair into an exact assertion.
func AssertEq(name string, v Val) AssertField { return AssertField{Name: name, V: v} }

func AssertNear(name string, v Val, tol float64) AssertField {
    return AssertField{Name: name, V: v, Tol: tol}
}

// resolveAssert returns the assertions for a case. If Assert is empty the
// Setup values double as the expected output (with default tolerance 0).
func (c Case) resolveAssert() []AssertField {
    if len(c.Assert) > 0 {
        return c.Assert
    }
    out := make([]AssertField, 0, len(c.Setup))
    for _, f := range c.Setup {
        out = append(out, AssertField{Name: f.Name, V: f.V})
    }
    return out
}

// ──────────────────────────────────────────────────────────────────────────
// Helpers shared by emitters
// ──────────────────────────────────────────────────────────────────────────

// snake → camelCase (lowerCamelCase). Matches the JS / Go / Java field
// naming conventions used by the bubbler generators.
func camelCase(s string) string {
    parts := strings.Split(s, "_")
    var sb strings.Builder
    for i, p := range parts {
        if p == "" {
            continue
        }
        if i == 0 {
            sb.WriteString(p)
        } else {
            sb.WriteString(strings.ToUpper(p[:1]))
            sb.WriteString(p[1:])
        }
    }
    return sb.String()
}

// snake → PascalCase (used by Go field names).
func pascalCase(s string) string {
    parts := strings.Split(s, "_")
    var sb strings.Builder
    for _, p := range parts {
        if p == "" {
            continue
        }
        sb.WriteString(strings.ToUpper(p[:1]))
        sb.WriteString(p[1:])
    }
    return sb.String()
}

// indent prefixes every line of s with the given indent string.
func indent(s, prefix string) string {
    if s == "" {
        return s
    }
    var sb strings.Builder
    lines := strings.Split(s, "\n")
    for i, ln := range lines {
        if ln != "" {
            sb.WriteString(prefix)
        }
        sb.WriteString(ln)
        if i < len(lines)-1 {
            sb.WriteByte('\n')
        }
    }
    return sb.String()
}

// hexByte renders a uint8 as "0xAA".
func hexByte(b uint8) string { return fmt.Sprintf("0x%02X", b) }
