// Bubbler E2E Feature Test — Narrow Bit-Width Arrays
//
// Tests the corner case where the declared [size] in bytes is LESS than
// the natural width of the element type × element count, causing each
// element to occupy fewer bits than its natural type width.
//
// Formula:  bits_per_element = (total_bytes × 8) / element_count
//
// When bits_per_element < type_bits, sign extension is used for signed types.

package bitwid;

option go_package       = "bitwid";
option java_package     = "com.example.bitwid";
option csharp_namespace = "Bitwid";

// ============================================================
// NarrowBWTest — four sub-cases of narrow element width
//
// | Field     | Type   | Count | Bytes | Bits/elem | Range             |
// |-----------|--------|-------|-------|-----------|-------------------|
// | narrow12  | int16  |   4   |   6   |    12     | [-2048, 2047]     |
// | narrow16  | int32  |   3   |   6   |    16     | [-32768, 32767]   |
// | narrow24  | int64  |   2   |   6   |    24     | [-8388608, 8388607] |
// | narrow6   | uint8  |   4   |   3   |     6     | [0, 63]           |
//
// Total struct size: 6+6+6+3 = 21 bytes
// ============================================================
struct NarrowBWTest {
    int16<4> narrow12 [6];
    int32<3> narrow16 [6];
    int64<2> narrow24 [6];
    uint8<4> narrow6  [3];
}
