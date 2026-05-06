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
    int64<5>  i64_arr;
    uint64<6> u64_arr;
    Color<7>  large_enum_arr;
    Point<8>  large_struct_arr;
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

// ============================================================
// Test 9: FloatSpecials
// IEEE-754 corner cases: +/-Inf, NaN (positive/negative), -0.0, +0.0,
// FLT_MAX/DBL_MAX, smallest normal positive, smallest subnormal positive.
// Round-trip must preserve every bit, including NaN sign / payload.
// ============================================================
struct FloatSpecials {
    float32 f32_pos_inf;
    float32 f32_neg_inf;
    float32 f32_qnan_pos;
    float32 f32_qnan_neg;
    float32 f32_neg_zero;
    float32 f32_pos_zero;
    float32 f32_max;
    float32 f32_min_normal;
    float32 f32_subnormal;

    float64 f64_pos_inf;
    float64 f64_neg_inf;
    float64 f64_qnan_pos;
    float64 f64_qnan_neg;
    float64 f64_neg_zero;
    float64 f64_pos_zero;
    float64 f64_max;
    float64 f64_min_normal;
    float64 f64_subnormal;
}

// ============================================================
// Test 10-14: Array sizing matrix for -unroll feature
// SmallArrays / MediumArrays / LargeArrays / VeryLargeArrays / MixedArrays.
// Each language driver round-trips these so test_options.sh's runtime
// matrix exercises every unroll threshold against arrays of every size class.
// ============================================================
struct SmallArrays {
    uint8<2>  small_uint8;
    int16<2>  small_int16;
    Color<2>  small_enum;
}

struct MediumArrays {
    uint8<4>   medium_uint8;
    int16<4>   medium_int16;
    Color<4>   medium_enum;
    float32<4> medium_float32;
}

struct LargeArrays {
    uint8<8>   large_uint8;
    int16<8>   large_int16;
    Color<8>   large_enum;
    float64<8> large_float64;
}

struct VeryLargeArrays {
    uint8<32>  very_large_uint8;
    int16<32>  very_large_int16;
    Color<32>  very_large_enum;
}

struct MixedArrays {
    uint8<1>   single_elem;
    uint16<2>  pair;
    uint32<4>  quad;
    uint64<8>  octet;
}
