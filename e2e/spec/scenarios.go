package main

import (
    "math"
    "strings"
)

// AllScenarios is the single source of truth for every codec test case
// emitted into per-language driver source. Adding a new struct or case is
// a one-place edit here — every language driver picks it up on next run.
func AllScenarios() *Spec {
    spec := &Spec{
        Scenarios: []Scenario{
            simpleScalars(),
            bitPacked(),
            bigEndianFields(),
            arrayFields(),
            embedStructs(),
            constantFields(),
            getterSetter(),
            dynamicFields(),
            floatSpecials(),
            smallArrays(),
            mediumArrays(),
            largeArrays(),
            veryLargeArrays(),
            mixedArrays(),
            codecArrays(),
            narrowBW(),
        },
    }
    for i := range spec.Scenarios {
        scenario := &spec.Scenarios[i]
        if len(scenario.Cases) == 0 || len(scenario.Cases[0].Setup) == 0 {
            continue
        }
        hasEncodedSourceCheck := false
        for _, check := range scenario.DecodeSizeChecks {
            if check.SourceCase != "" {
                hasEncodedSourceCheck = true
                break
            }
        }
        if hasEncodedSourceCheck {
            continue
        }
        source := scenario.Cases[0].Name
        scenario.DecodeSizeChecks = append(scenario.DecodeSizeChecks,
            DecodeSizeCheck{Name: "complete_packet", SourceCase: source},
            DecodeSizeCheck{Name: "truncate_one_byte", SourceCase: source, Truncate: 1},
        )
    }
    return spec
}

// ──────────────────────────────────────────────────────────────────────────
// testcase.bb / package "testpkg"
// ──────────────────────────────────────────────────────────────────────────

func simpleScalars() Scenario {
    return Scenario{
        BBFile:     "testcase",
        Package:    "testpkg",
        StructName: "SimpleScalars",
        Notes:      "Every primitive type, including signed/unsigned extremes and bool/enum",
        Cases: []Case{{
            Name: "round_trip",
            Setup: []FieldVal{
                F("u8_zero", U8(0)),
                F("u8_max", U8(255)),
                F("i8_min", I8(-128)),
                F("i8_max", I8(127)),
                F("u16_val", U16(0xBEEF)),
                F("i16_neg", I16(-1234)),
                F("u32_val", U32(0xDEADBEEF)),
                F("i32_neg", I32(-100000)),
                F("u64_val", U64(0xCAFEBABEDEADBEEF)),
                F("i64_neg", I64(-1)),
                F("f32_val", F32(3.14159)),
                F("f64_val", F64(2.718281828459045)),
                F("flag_t", B(true)),
                F("flag_f", B(false)),
                F("color", E("Color", "BLUE")),
            },
            Assert: []AssertField{
                {Name: "u8_zero", V: U8(0)},
                {Name: "u8_max", V: U8(255)},
                {Name: "i8_min", V: I8(-128)},
                {Name: "i8_max", V: I8(127)},
                {Name: "u16_val", V: U16(0xBEEF)},
                {Name: "i16_neg", V: I16(-1234)},
                {Name: "u32_val", V: U32(0xDEADBEEF)},
                {Name: "i32_neg", V: I32(-100000)},
                {Name: "u64_val", V: U64(0xCAFEBABEDEADBEEF)},
                {Name: "i64_neg", V: I64(-1)},
                // Floats use a tolerance — the input was a literal, not bits;
                // bit-exact float coverage lives in FloatSpecials.
                {Name: "f32_val", V: F32(3.14159), Tol: 1e-3},
                {Name: "f64_val", V: F64(2.718281828459045), Tol: 1e-12},
                {Name: "flag_t", V: B(true)},
                {Name: "flag_f", V: B(false)},
                {Name: "color", V: E("Color", "BLUE")},
            },
            Wire: []byte{
                0x00, 0xFF, 0x80, 0x7F, 0xEF, 0xBE, 0x2E, 0xFB,
                0xEF, 0xBE, 0xAD, 0xDE, 0x60, 0x79, 0xFE, 0xFF,
                0xEF, 0xBE, 0xAD, 0xDE, 0xBE, 0xBA, 0xFE, 0xCA,
                0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
                0xD0, 0x0F, 0x49, 0x40,
                0x69, 0x57, 0x14, 0x8B, 0x0A, 0xBF, 0x05, 0x40,
                0x01, 0x00, 0x03,
            },
        }},
    }
}

func bitPacked() Scenario {
    return Scenario{
        BBFile:     "testcase",
        Package:    "testpkg",
        StructName: "BitPacked",
        Notes:      "Sub-byte fields, sign extension across non-byte-aligned widths",
        Cases: []Case{
            {
                Name: "typical",
                Setup: []FieldVal{
                    F("b0", B(true)),
                    F("b1", B(false)),
                    F("nibble", U8(0xA)),
                    F("i2", I8(-1)),
                    F("u4", U8(0xF)),
                    F("i20", I32(524287)),
                    F("i48", I64(0x7FFFFFFFFFFF)),
                    F("status", E("Status", "BROKEN")),
                },
            },
            {
                Name: "min_signed",
                Setup: []FieldVal{
                    F("b0", B(false)),
                    F("b1", B(false)),
                    F("nibble", U8(0)),
                    F("i2", I8(-2)),
                    F("u4", U8(0)),
                    F("i20", I32(-524288)),
                    F("i48", I64(-1)),
                    F("status", E("Status", "IDLE")),
                },
            },
            {
                Name: "zeros",
                Setup: []FieldVal{
                    F("b0", B(false)),
                    F("b1", B(false)),
                    F("nibble", U8(0)),
                    F("i2", I8(0)),
                    F("u4", U8(0)),
                    F("i20", I32(0)),
                    F("i48", I64(0)),
                    F("status", E("Status", "IDLE")),
                },
            },
        },
    }
}

func bigEndianFields() Scenario {
    arr := A(VKI16, "", I16(100), I16(-200), I16(300), I16(-400))
    return Scenario{
        BBFile:     "testcase",
        Package:    "testpkg",
        StructName: "BigEndianFields",
        Notes:      "Big-endian byte order, including byte-layout assertions",
        Cases: []Case{{
            Name: "round_trip",
            Setup: []FieldVal{
                F("u16", U16(0x1234)),
                F("i32", I32(-1)),
                F("f32", F32(3.14)),
                F("arr", arr),
            },
            Assert: []AssertField{
                {Name: "u16", V: U16(0x1234)},
                {Name: "i32", V: I32(-1)},
                {Name: "f32", V: F32(3.14), Tol: 1e-3},
                {Name: "arr", V: arr},
            },
            Wire: []byte{
                0xBE, 0x12, 0x34, 0xFF, 0xFF, 0xFF, 0xFF,
                0x40, 0x48, 0xF5, 0xC3,
                0x00, 0x64, 0xFF, 0x38, 0x01, 0x2C, 0xFE, 0x70,
            },
        }},
    }
}

func arrayFields() Scenario {
    return Scenario{
        BBFile:     "testcase",
        Package:    "testpkg",
        StructName: "ArrayFields",
        Notes:      "Arrays of primitives, enums, and structs across element-count boundaries",
        Cases: []Case{{
            Name: "round_trip",
            Setup: []FieldVal{
                F("u8_arr", A(VKU8, "", U8(0), U8(127), U8(128), U8(255))),
                F("i16_arr", A(VKI16, "", I16(-32768), I16(0), I16(32767))),
                F("color_arr", A(VKEnum, "Color", E("Color", "RED"), E("Color", "BLUE"))),
                F("point_arr", A(VKStruct, "Point",
                    Struct("Point", F("x", I16(-100)), F("y", I16(200))),
                    Struct("Point", F("x", I16(30000)), F("y", I16(-30000))),
                )),
                F("i64_arr", A(VKI64, "",
                    I64(-1), I64(0), I64(1), I64(-0x123456789A), I64(0x123456789A),
                )),
                F("u64_arr", A(VKU64, "",
                    U64(0), U64(1), U64(0xFFFFFFFF),
                    U64(0x123456789ABCDEF0), U64(0x8000000000000000), U64(0xFFFFFFFFFFFFFFFF),
                )),
                F("large_enum_arr", A(VKEnum, "Color",
                    E("Color", "RED"), E("Color", "GREEN"), E("Color", "BLUE"),
                    E("Color", "RED"), E("Color", "GREEN"), E("Color", "BLUE"),
                    E("Color", "RED"),
                )),
                F("large_struct_arr", A(VKStruct, "Point",
                    Struct("Point", F("x", I16(-1000)), F("y", I16(1000))),
                    Struct("Point", F("x", I16(-2000)), F("y", I16(2000))),
                    Struct("Point", F("x", I16(-3000)), F("y", I16(3000))),
                    Struct("Point", F("x", I16(-4000)), F("y", I16(4000))),
                    Struct("Point", F("x", I16(-5000)), F("y", I16(5000))),
                    Struct("Point", F("x", I16(-6000)), F("y", I16(6000))),
                    Struct("Point", F("x", I16(-7000)), F("y", I16(7000))),
                    Struct("Point", F("x", I16(-8000)), F("y", I16(8000))),
                )),
            },
        }},
    }
}

func embedStructs() Scenario {
    return Scenario{
        BBFile:     "testcase",
        Package:    "testpkg",
        StructName: "EmbedStructs",
        Notes:      "Anonymous embed (AnonInner promoted as ax/ay) plus named struct embed",
        Cases: []Case{{
            Name: "round_trip",
            Setup: []FieldVal{
                F("id", U8(42)),
                F("ax", U8(10)),
                F("ay", U8(20)),
                F("pt", Struct("Point", F("x", I16(-500)), F("y", I16(500)))),
                F("flags", U8(0xFF)),
            },
        }},
    }
}

func constantFields() Scenario {
    return Scenario{
        BBFile:     "testcase",
        Package:    "testpkg",
        StructName: "ConstantFields",
        Notes:      "Constant-field encode (verbatim) and decode (validation)",
        Cases: []Case{{
            Name:  "round_trip",
            Setup: []FieldVal{F("length", U16(1024))},
            Wire:  []byte{0xAA, 0x02, 0x00, 0x04, 0x02},
            Errors: []DecodeError{
                {Name: "bad_header", Patches: []BytePatch{{Offset: 0, Value: 0xFF}}, ExpectedRet: -1},
                {Name: "bad_version", Patches: []BytePatch{{Offset: 1, Value: 0x00}}, ExpectedRet: -1},
            },
        }},
    }
}

func getterSetter() Scenario {
    // Special-cased struct: scenario carries Accessors instead of Setup/Assert.
    return Scenario{
        BBFile:     "testcase",
        Package:    "testpkg",
        StructName: "GetterSetter",
        Notes:      "Custom getter/setter accessors and round-trip preservation",
        Cases: []Case{
            {
                Name: "voltage_3_3",
                Accessors: []AccessorRoundTrip{
                    {SetterName: "voltage", SetArg: F64(3.3), GetterName: "voltage", Expected: F64(3.3), Tol: 5e-3, NoRoundTrip: true},
                },
            },
            {
                Name: "voltage_zero",
                Accessors: []AccessorRoundTrip{
                    {SetterName: "voltage", SetArg: F64(0.0), GetterName: "voltage", Expected: F64(0.0), Tol: 1e-6, NoRoundTrip: true},
                },
            },
            {
                Name: "celsius_36_5",
                Accessors: []AccessorRoundTrip{
                    {SetterName: "celsius", SetArg: F64(36.5), GetterName: "celsius", Expected: F64(36.5), Tol: 1e-2, NoRoundTrip: true},
                },
            },
            {
                Name: "celsius_neg40",
                Accessors: []AccessorRoundTrip{
                    {SetterName: "celsius", SetArg: F64(-40.0), GetterName: "celsius", Expected: F64(-40.0), Tol: 1e-2, NoRoundTrip: true},
                },
            },
            {
                Name: "round_trip",
                Accessors: []AccessorRoundTrip{
                    {SetterName: "voltage", SetArg: F64(1.65), GetterName: "voltage", Expected: F64(1.65), Tol: 5e-3},
                    {SetterName: "celsius", SetArg: F64(25.0), GetterName: "celsius", Expected: F64(25.0), Tol: 0.1},
                },
            },
        },
    }
}

func dynamicFields() Scenario {
    blob := []byte{0xDE, 0xAD, 0xBE, 0xEF, 0x00, 0x01, 0x02}
    return Scenario{
        BBFile:     "testcase",
        Package:    "testpkg",
        StructName: "DynamicFields",
        IsDynamic:  true,
        Notes:      "Variable-length string and bytes; UTF-8; decode_size truncation",
        Cases: []Case{
            {
                Name: "non_empty",
                Setup: []FieldVal{
                    F("id", U32(0xDEADBEEF)),
                    F("label", S("hello, bubbler!")),
                    F("data", Bs(blob)),
                },
            },
            {
                Name: "empty",
                Setup: []FieldVal{
                    F("id", U32(0)),
                    F("label", S("")),
                    F("data", Bs(nil)),
                },
            },
            {
                Name: "utf8",
                Setup: []FieldVal{
                    F("id", U32(1)),
                    F("label", S("\xe4\xb8\xad\xe6\x96\x87")), // 中文
                    F("data", Bs(nil)),
                },
            },
            {
                Name: "large_payload",
                Setup: []FieldVal{
                    F("id", U32(999)),
                    F("label", S("large")),
                    // 200 bytes triggers multi-byte length varint
                    F("data", Bs(seq(200))),
                },
            },
        },
        DecodeSizeChecks: []DecodeSizeCheck{
            {Name: "complete_packet", SourceCase: "non_empty"},
            {Name: "truncate_one_byte", SourceCase: "non_empty", Truncate: 1},
            {Name: "missing_string_terminator", Bytes: []byte{1, 0, 0, 0, 'A'}, Expected: -6},
            {Name: "truncated_bytes_varint", Bytes: []byte{1, 0, 0, 0, 'A', 0, 0x80}, Expected: -8},
            {Name: "truncated_bytes_payload", Bytes: []byte{1, 0, 0, 0, 'A', 0, 0x03, 0xAA, 0xBB}, Expected: -10},
            {Name: "only_fixed_header", Bytes: []byte{1, 0, 0, 0}, Expected: -5},
        },
    }
}

func floatSpecials() Scenario {
    // Bit-exact assertions for IEEE-754 corner cases.
    f32qnanPos := uint32(0x7FC00000) // canonical quiet NaN
    f32qnanNeg := uint32(0xFFC00000) // negative quiet NaN
    f64qnanPos := uint64(0x7FF8000000000000)
    f64qnanNeg := uint64(0xFFF8000000000000)

    f32MinNormal := math.Float32bits(math.Float32frombits(0x00800000))     // 1.17549435e-38
    f32Subnormal := math.Float32bits(math.Float32frombits(0x00000001))     // smallest positive denormal
    f64MinNormal := math.Float64bits(math.Float64frombits(0x0010000000000000))
    f64Subnormal := math.Float64bits(math.Float64frombits(0x0000000000000001))

    setup := []FieldVal{
        F("f32_pos_inf", F32Bits(0x7F800000)),
        F("f32_neg_inf", F32Bits(0xFF800000)),
        F("f32_qnan_pos", F32Bits(f32qnanPos)),
        F("f32_qnan_neg", F32Bits(f32qnanNeg)),
        F("f32_neg_zero", F32Bits(0x80000000)),
        F("f32_pos_zero", F32Bits(0x00000000)),
        F("f32_max", F32Bits(0x7F7FFFFF)), // FLT_MAX
        F("f32_min_normal", F32Bits(f32MinNormal)),
        F("f32_subnormal", F32Bits(f32Subnormal)),

        F("f64_pos_inf", F64Bits(0x7FF0000000000000)),
        F("f64_neg_inf", F64Bits(0xFFF0000000000000)),
        F("f64_qnan_pos", F64Bits(f64qnanPos)),
        F("f64_qnan_neg", F64Bits(f64qnanNeg)),
        F("f64_neg_zero", F64Bits(0x8000000000000000)),
        F("f64_pos_zero", F64Bits(0x0000000000000000)),
        F("f64_max", F64Bits(0x7FEFFFFFFFFFFFFF)), // DBL_MAX
        F("f64_min_normal", F64Bits(f64MinNormal)),
        F("f64_subnormal", F64Bits(f64Subnormal)),
    }
    // Default assertions = setup with Tol == 0 (bit-exact for floats).
    return Scenario{
        BBFile:     "testcase",
        Package:    "testpkg",
        StructName: "FloatSpecials",
        Notes:      "IEEE-754 corner cases. Round-trip must preserve all bits, NaN sign/payload included.",
        Cases:      []Case{{Name: "round_trip", Setup: setup}},
    }
}

// ──────────────────────────────────────────────────────────────────────────
// Unroll-coverage scenarios. test_options.sh #11 generates each language at
// every -unroll value and compiles+runs the driver, so these cases get
// re-run under unroll = -1, 0, 1, 2, 4, 6, 8, 16, 32.
// ──────────────────────────────────────────────────────────────────────────

func smallArrays() Scenario {
    return Scenario{
        BBFile:     "testcase",
        Package:    "testpkg",
        StructName: "SmallArrays",
        Notes:      "Length-2 arrays — typically inlined under default unroll",
        Cases: []Case{{
            Name: "round_trip",
            Setup: []FieldVal{
                F("small_uint8", A(VKU8, "", U8(0), U8(255))),
                F("small_int16", A(VKI16, "", I16(-32768), I16(32767))),
                F("small_enum", A(VKEnum, "Color", E("Color", "RED"), E("Color", "BLUE"))),
            },
        }},
    }
}

func mediumArrays() Scenario {
    return Scenario{
        BBFile:     "testcase",
        Package:    "testpkg",
        StructName: "MediumArrays",
        Notes:      "Length-4 arrays — sit on the default unroll threshold",
        Cases: []Case{{
            Name: "round_trip",
            Setup: []FieldVal{
                F("medium_uint8", A(VKU8, "", U8(1), U8(2), U8(3), U8(4))),
                F("medium_int16", A(VKI16, "", I16(-1), I16(0), I16(1), I16(0x7FFF))),
                F("medium_enum", A(VKEnum, "Color",
                    E("Color", "RED"), E("Color", "GREEN"), E("Color", "BLUE"), E("Color", "RED"))),
                F("medium_float32", A(VKF32, "", F32(1.0), F32(-1.0), F32(0.5), F32(-0.5))),
            },
            Assert: []AssertField{
                {Name: "medium_uint8", V: A(VKU8, "", U8(1), U8(2), U8(3), U8(4))},
                {Name: "medium_int16", V: A(VKI16, "", I16(-1), I16(0), I16(1), I16(0x7FFF))},
                {Name: "medium_enum", V: A(VKEnum, "Color",
                    E("Color", "RED"), E("Color", "GREEN"), E("Color", "BLUE"), E("Color", "RED"))},
                {Name: "medium_float32", V: A(VKF32, "",
                    F32(1.0), F32(-1.0), F32(0.5), F32(-0.5)), Tol: 1e-6},
            },
        }},
    }
}

func largeArrays() Scenario {
    setF64 := A(VKF64, "",
        F64(1.5), F64(-1.5), F64(0.0), F64(-0.0),
        F64(math.Pi), F64(math.E), F64(1e-10), F64(1e10),
    )
    return Scenario{
        BBFile:     "testcase",
        Package:    "testpkg",
        StructName: "LargeArrays",
        Notes:      "Length-8 arrays — looped under default unroll",
        Cases: []Case{{
            Name: "round_trip",
            Setup: []FieldVal{
                F("large_uint8", A(VKU8, "",
                    U8(0), U8(1), U8(2), U8(3), U8(4), U8(5), U8(6), U8(7))),
                F("large_int16", A(VKI16, "",
                    I16(-100), I16(-50), I16(-1), I16(0),
                    I16(1), I16(50), I16(100), I16(0x7FFF))),
                F("large_enum", A(VKEnum, "Color",
                    E("Color", "RED"), E("Color", "GREEN"), E("Color", "BLUE"),
                    E("Color", "RED"), E("Color", "GREEN"), E("Color", "BLUE"),
                    E("Color", "RED"), E("Color", "GREEN"))),
                F("large_float64", setF64),
            },
            Assert: []AssertField{
                {Name: "large_uint8", V: A(VKU8, "",
                    U8(0), U8(1), U8(2), U8(3), U8(4), U8(5), U8(6), U8(7))},
                {Name: "large_int16", V: A(VKI16, "",
                    I16(-100), I16(-50), I16(-1), I16(0),
                    I16(1), I16(50), I16(100), I16(0x7FFF))},
                {Name: "large_enum", V: A(VKEnum, "Color",
                    E("Color", "RED"), E("Color", "GREEN"), E("Color", "BLUE"),
                    E("Color", "RED"), E("Color", "GREEN"), E("Color", "BLUE"),
                    E("Color", "RED"), E("Color", "GREEN"))},
                {Name: "large_float64", V: setF64, Tol: 1e-12},
            },
        }},
    }
}

func veryLargeArrays() Scenario {
    u8 := make([]Val, 32)
    for i := 0; i < 32; i++ {
        u8[i] = U8(uint8(i * 7))
    }
    i16 := make([]Val, 32)
    for i := 0; i < 32; i++ {
        i16[i] = I16(int16(i*1000 - 16000))
    }
    enum := make([]Val, 32)
    names := []string{"RED", "GREEN", "BLUE"}
    for i := 0; i < 32; i++ {
        enum[i] = E("Color", names[i%3])
    }
    return Scenario{
        BBFile:     "testcase",
        Package:    "testpkg",
        StructName: "VeryLargeArrays",
        Notes:      "Length-32 arrays — always looped at any reasonable unroll threshold",
        Cases: []Case{{
            Name: "round_trip",
            Setup: []FieldVal{
                F("very_large_uint8", A(VKU8, "", u8...)),
                F("very_large_int16", A(VKI16, "", i16...)),
                F("very_large_enum", A(VKEnum, "Color", enum...)),
            },
        }},
    }
}

func mixedArrays() Scenario {
    return Scenario{
        BBFile:     "testcase",
        Package:    "testpkg",
        StructName: "MixedArrays",
        Notes:      "Lengths 1/2/4/8 — verifies unroll boundary behaviour element-by-element",
        Cases: []Case{{
            Name: "round_trip",
            Setup: []FieldVal{
                F("single_elem", A(VKU8, "", U8(0xAB))),
                F("pair", A(VKU16, "", U16(0x1234), U16(0xBEEF))),
                F("quad", A(VKU32, "",
                    U32(0xDEADBEEF), U32(0xCAFEBABE), U32(0), U32(0xFFFFFFFF))),
                F("octet", A(VKU64, "",
                    U64(0), U64(1), U64(0xFFFFFFFFFFFFFFFF),
                    U64(0x123456789ABCDEF0), U64(0x8000000000000000),
                    U64(0xFEDCBA9876543210), U64(0xAAAAAAAAAAAAAAAA),
                    U64(0x5555555555555555))),
            },
        }},
    }
}

func codecArrays() Scenario {
    stringsValue := A(VKString, "",
        S(""),
        S("ascii"),
        S("中文"),
        S("quote=\" slash=\\"),
        S(strings.Repeat("x", 130)),
    )
    blobsValue := A(VKBytes, "",
        Bs(nil),
        Bs([]byte{0}),
        Bs([]byte{0x00, 0x7F, 0x80, 0xFF}),
        Bs(seq(127)),
        Bs(seq(128)),
    )
    recordsValue := A(VKStruct, "BlobRecord",
        Struct("BlobRecord", F("code", U16(0)), F("label", S("")), F("payload", Bs(nil))),
        Struct("BlobRecord", F("code", U16(1)), F("label", S("one")), F("payload", Bs([]byte{1}))),
        Struct("BlobRecord", F("code", U16(0x7FFF)), F("label", S("中文")), F("payload", Bs(seq(3)))),
        Struct("BlobRecord", F("code", U16(0x8000)), F("label", S("boundary127")), F("payload", Bs(seq(127)))),
        Struct("BlobRecord", F("code", U16(0xFFFF)), F("label", S("boundary128")), F("payload", Bs(seq(128)))),
    )
    return Scenario{
        BBFile:     "testcase",
        Package:    "testpkg",
        StructName: "CodecArrays",
        IsDynamic:  true,
        Notes:      "Every array element category, including strings, bytes and dynamic structs",
        Cases: []Case{{
            Name: "round_trip",
            Setup: []FieldVal{
                F("bools", A(VKBool, "", B(false), B(true), B(true), B(false), B(true))),
                F("i8s", A(VKI8, "", I8(-128), I8(-1), I8(0), I8(1), I8(127))),
                F("u16s", A(VKU16, "", U16(0), U16(1), U16(0x7FFF), U16(0x8000), U16(0xFFFF))),
                F("i32s", A(VKI32, "", I32(-2147483648), I32(-1), I32(0), I32(1), I32(2147483647))),
                F("u32s", A(VKU32, "", U32(0), U32(1), U32(0x7FFFFFFF), U32(0x80000000), U32(0xFFFFFFFF))),
                F("f32s", A(VKF32, "", F32Bits(0x80000000), F32(1.5), F32(-2.25), F32Bits(0x7F800000), F32Bits(0x7FC00000))),
                F("f64s", A(VKF64, "", F64Bits(0x8000000000000000), F64(math.Pi), F64(-math.E), F64Bits(0x7FF0000000000000), F64Bits(0x7FF8000000000000))),
                F("strings", stringsValue),
                F("blobs", blobsValue),
                F("records", recordsValue),
            },
        }},
        DecodeSizeChecks: []DecodeSizeCheck{
            {Name: "complete_packet", SourceCase: "round_trip"},
            {Name: "truncate_one_byte", SourceCase: "round_trip", Truncate: 1},
        },
    }
}

// ──────────────────────────────────────────────────────────────────────────
// features/bitwid.bb / package "bitwid"
// ──────────────────────────────────────────────────────────────────────────

func narrowBW() Scenario {
    return Scenario{
        BBFile:     "bitwid",
        Package:    "bitwid",
        StructName: "NarrowBWTest",
        Notes:      "Narrow bit-width arrays: element bits < type bits, with sign extension",
        Cases: []Case{
            {
                Name: "typical_values",
                Setup: []FieldVal{
                    F("narrow12", A(VKI16, "", I16(2047), I16(-2048), I16(0), I16(-1))),
                    F("narrow16", A(VKI32, "", I32(32767), I32(-32768), I32(0))),
                    F("narrow24", A(VKI64, "", I64(8388607), I64(-8388608))),
                    F("narrow6", A(VKU8, "", U8(63), U8(0), U8(32), U8(1))),
                },
            },
            {
                Name: "zero_round_trip",
                Setup: []FieldVal{
                    F("narrow12", A(VKI16, "", I16(0), I16(0), I16(0), I16(0))),
                    F("narrow16", A(VKI32, "", I32(0), I32(0), I32(0))),
                    F("narrow24", A(VKI64, "", I64(0), I64(0))),
                    F("narrow6", A(VKU8, "", U8(0), U8(0), U8(0), U8(0))),
                },
            },
        },
    }
}

func seq(n int) []byte {
    out := make([]byte, n)
    for i := range out {
        out[i] = byte(i & 0xFF)
    }
    return out
}
