// Bubbler E2E Test Case
// This file covers a wide range of corner cases for all code generators.

package testpkg;

option java_package     = "com.example.testpkg";
option go_package       = "testpkg";
option csharp_namespace = "Testpkg";

// ============================================================
// Enums
// ============================================================

// 1-byte enum
enum Color[1] {
    RED   = 0x01,
    GREEN = 0x02,
    BLUE  = 0x03,
}

// 12-bit enum (1 byte + 4 bits)
enum Status[1#4] {
    IDLE   = 0x000,
    ACTIVE = 0x001,
    BROKEN = 0xFFF,
}

// ============================================================
// Helper structs
// ============================================================

struct Point {
    int16 x;
    int16 y;
}

// ============================================================
// Test 1: SimpleScalars
// All primitive types, little-endian encoding
// Corner cases: min/max values, negative, zero, +/-inf NA
// ============================================================
struct SimpleScalars {
    uint8   u8_zero;
    uint8   u8_max;
    int8    i8_min;
    int8    i8_max;
    uint16  u16_val;
    int16   i16_neg;
    uint32  u32_val;
    int32   i32_neg;
    uint64  u64_val;
    int64   i64_neg;
    float32 f32_val;
    float64 f64_val;
    bool    flag_t;
    bool    flag_f;
    Color   color;
}

// ============================================================
// Test 2: BitPacked
// Sub-byte fields, bit-packing, cross-byte fields
// Struct total: 13 bytes (byte-aligned)
// Layout:
//   Byte  0: b0[#1]+b1[#1]+nibble[#4]+void[#2]      = 8 bits
//   Byte  1: i2[#2]+u4[#4]+void[#2]                 = 8 bits
//   Bytes 2-4: i20[2#4]+void[#4]                    = 24 bits
//   Bytes 5-10: i48[6]                              = 48 bits
//   Bytes 11-12: status[1#4]+void[#4]               = 16 bits
// ============================================================
struct BitPacked {
    struct Byte0[1] {
        bool  b0[#1];
        bool  b1[#1];
        uint8 nibble[#4];
        void  [#2];
    };
    int8   i2[#2];
    uint8  u4[#4];
    void   [#2];
    int32  i20[2#4];
    void   [#4];
    int64  i48[6];
    Status status;
    void   [#4];
}

// ============================================================
// Test 3: BigEndianFields
// Big-endian byte order for all field types
// ============================================================
struct BigEndianFields {
    uint8    magic = 0xBE [order = "big"];
    uint16   u16            [order = "big"];
    int32    i32            [order = "big"];
    float32  f32            [order = "big"];
    int16<4> arr            [order = "big"];
}

// ============================================================
// Test 4: ArrayFields
// Arrays of primitives, enums, and structs
// ============================================================
struct ArrayFields {
    uint8<4>  u8_arr;
    int16<3>  i16_arr;
    Color<2>  color_arr;
    Point<2>  point_arr;
}

// ============================================================
// Test 5: EmbedStructs
// Anonymous inline struct embed and named struct embed
// ============================================================
struct EmbedStructs {
    uint8 id;
    struct AnonInner[2] {
        uint8 ax;
        uint8 ay;
    };
    Point pt;
    uint8 flags;
}

// ============================================================
// Test 6: ConstantFields
// Constant field encoding/decoding (with decode validation)
// ============================================================
struct ConstantFields {
    uint8  header      = 0xAA;
    uint8  version     = 0x02;
    uint16 length;
    uint8  magic_color = 0x02;
}

// ============================================================
// Test 7: GetterSetter
// Custom getter/setter accessors
// ============================================================
struct GetterSetter {
    uint16 raw_adc {
        get voltage(float64): (float64)(value) * 3.3 / 4096.0;
        set voltage(float64): (uint16)(value * 4096.0 / 3.3);
    };
    int16 raw_temp {
        get celsius(float64): (float64)(value) / 100.0;
        set celsius(float64): (int16)(value * 100.0);
    };
}

// ============================================================
// Test 8: DynamicFields
// Variable-length string and bytes fields
// ============================================================
struct DynamicFields {
    uint32 id;
    string label;
    bytes  data;
}
