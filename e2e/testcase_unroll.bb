// Bubbler E2E Test Case for Loop Unroll Feature
// This file tests the -unroll flag with arrays of various sizes

package testpkg;

option java_package     = "com.example.testpkg";
option go_package       = "testpkg";
option csharp_namespace = "Testpkg";

// ============================================================
// Helper Types
// ============================================================

enum Color[1] {
    RED   = 0x01,
    GREEN = 0x02,
    BLUE  = 0x03,
}

struct Point {
    int16 x;
    int16 y;
}

// ============================================================
// Test Case: Arrays with various sizes
// For testing -unroll flag behavior
// ============================================================

// Small array (smaller than typical unroll threshold)
struct SmallArrays {
    uint8<2>  small_uint8;
    int16<2>  small_int16;
    Color<2>  small_enum;
}

// Medium array (around typical unroll threshold of 4)
struct MediumArrays {
    uint8<4>   medium_uint8;
    int16<4>   medium_int16;
    Color<4>   medium_enum;
    float32<4> medium_float32;
}

// Large array (larger than typical unroll threshold)
struct LargeArrays {
    uint8<8>   large_uint8;
    int16<8>   large_int16;
    Color<8>   large_enum;
    float64<8> large_float64;
}

// Very large array (forces loop generation)
struct VeryLargeArrays {
    uint8<32>  very_large_uint8;
    int16<32>  very_large_int16;
    Color<32>  very_large_enum;
}

// Mixed test case with different array sizes
struct MixedArrays {
    uint8<1>   single_elem;
    uint16<2>  pair;
    uint32<4>  quad;
    uint64<8>  octet;
}
