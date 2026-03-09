/*
 * Bubbler E2E Test — C++ target
 *
 * Build (from this directory, after code-gen step):
 *   g++ -std=c++17 -Igen -o run_test main.cpp -lm
 *
 * The run_tests.sh script generates gen/ then invokes this.
 */
#include "gen/testcase.bb.hpp"
#include "gen/bitwid.bb.hpp"
#include <cmath>
#include <cstdio>
#include <cstring>
#include <memory>
#include <vector>

/* ------------------------------------------------------------------ */
/* Minimal test framework                                               */
/* ------------------------------------------------------------------ */
static int g_pass = 0;
static int g_fail = 0;
static const char *g_current = "";

static void _check(bool cond, const char *expr, const char *file, int line) {
    if (cond) {
        ++g_pass;
    } else {
        fprintf(stderr, "  FAIL [%s] %s:%d  %s\n", g_current, file, line, expr);
        ++g_fail;
    }
}

#define CHECK(cond)           _check(!!(cond), #cond, __FILE__, __LINE__)
#define CHECK_EQ(a,b)         _check((a)==(b), #a " == " #b, __FILE__, __LINE__)
#define CHECK_NEAR(a,b,eps)   _check(std::abs((double)(a)-(double)(b))<(eps), \
                                  #a " ~= " #b, __FILE__, __LINE__)

using namespace testpkg;

/* ------------------------------------------------------------------ */
/* Helper: null-deleter shared_ptr (no ownership)                      */
/* ------------------------------------------------------------------ */
static struct bytes make_bytes_ref(uint8_t *ptr, uint64_t len) {
    struct bytes b;
    b.data   = std::shared_ptr<uint8_t[]>(ptr, [](uint8_t*){});
    b.length = len;
    return b;
}

/* ------------------------------------------------------------------ */
/* Test 1: SimpleScalars                                               */
/* ------------------------------------------------------------------ */
static void test_simple_scalars() {
    g_current = "SimpleScalars";
    SimpleScalars s{};
    s.u8_zero = 0;
    s.u8_max  = 255;
    s.i8_min  = -128;
    s.i8_max  = 127;
    s.u16_val = 0xBEEF;
    s.i16_neg = -1234;
    s.u32_val = 0xDEADBEEFU;
    s.i32_neg = -100000;
    s.u64_val = 0xCAFEBABEDEADBEEFULL;
    s.i64_neg = -1LL;
    s.f32_val = 3.14159f;
    s.f64_val = 2.718281828459045;
    s.flag_t  = true;
    s.flag_f  = false;
    s.color   = Color::BLUE;

    std::vector<uint8_t> buf(SimpleScalars::size);
    CHECK_EQ((int)s.encode(buf.data()), (int)SimpleScalars::size);

    SimpleScalars d{};
    CHECK_EQ((int)d.decode(buf.data()), (int)SimpleScalars::size);

    CHECK_EQ((int)d.u8_zero,   0);
    CHECK_EQ((int)d.u8_max,    255);
    CHECK_EQ((int)d.i8_min,    -128);
    CHECK_EQ((int)d.i8_max,    127);
    CHECK_EQ((unsigned)d.u16_val, 0xBEEFu);
    CHECK_EQ((int)d.i16_neg,   -1234);
    CHECK_EQ(d.u32_val,        0xDEADBEEFU);
    CHECK_EQ((long long)d.i32_neg, -100000LL);
    CHECK_EQ(d.u64_val,        0xCAFEBABEDEADBEEFULL);
    CHECK_EQ((long long)d.i64_neg, -1LL);
    CHECK_NEAR(d.f32_val, 3.14159f, 1e-3);
    CHECK_NEAR(d.f64_val, 2.718281828459045, 1e-12);
    CHECK_EQ(d.flag_t, true);
    CHECK_EQ(d.flag_f, false);
    CHECK_EQ(d.color, Color::BLUE);
}

/* ------------------------------------------------------------------ */
/* Test 2: BitPacked                                                   */
/* ------------------------------------------------------------------ */
static void test_bit_packed() {
    g_current = "BitPacked";
    BitPacked s{};

    /* Case 1: typical values */
    s.b0     = true;
    s.b1     = false;
    s.nibble = 0xA;
    s.i2     = -1;
    s.u4     = 0xF;
    s.i20    = 524287;
    s.i48    = 0x7FFFFFFFFFFFLL;
    s.status = Status::BROKEN;

    std::vector<uint8_t> buf(BitPacked::size);
    CHECK_EQ((int)s.encode(buf.data()), (int)BitPacked::size);

    BitPacked d{};
    CHECK_EQ((int)d.decode(buf.data()), (int)BitPacked::size);

    CHECK_EQ(d.b0, true);
    CHECK_EQ(d.b1, false);
    CHECK_EQ((int)d.nibble, 0xA);
    CHECK_EQ((int)d.i2,    -1);
    CHECK_EQ((int)d.u4,    0xF);
    CHECK_EQ((long long)d.i20, 524287LL);
    CHECK_EQ((long long)d.i48, 0x7FFFFFFFFFFFLL);
    CHECK_EQ(d.status, Status::BROKEN);

    /* Case 2: sign-extension boundaries */
    s.i2  = -2;
    s.i20 = -524288;
    s.i48 = -1LL;
    buf.assign(BitPacked::size, 0);
    s.encode(buf.data());
    d = BitPacked{};
    d.decode(buf.data());
    CHECK_EQ((int)d.i2,        -2);
    CHECK_EQ((long long)d.i20, -524288LL);
    CHECK_EQ((long long)d.i48, -1LL);

    /* Case 3: zeros */
    s = BitPacked{};
    buf.assign(BitPacked::size, 0);
    s.encode(buf.data());
    d = BitPacked{};
    d.decode(buf.data());
    CHECK_EQ((int)d.i2,  0);
    CHECK_EQ(d.status, Status::IDLE);
}

/* ------------------------------------------------------------------ */
/* Test 3: BigEndianFields                                             */
/* ------------------------------------------------------------------ */
static void test_big_endian_fields() {
    g_current = "BigEndianFields";
    BigEndianFields s{};
    s.u16    = 0x1234;
    s.i32    = -1;
    s.f32    = 3.14f;
    s.arr[0] = 100;  s.arr[1] = -200;
    s.arr[2] = 300;  s.arr[3] = -400;

    std::vector<uint8_t> buf(BigEndianFields::size);
    CHECK_EQ((int)s.encode(buf.data()), (int)BigEndianFields::size);

    /* Constant byte at position 0 */
    CHECK_EQ(buf[0], 0xBE);
    /* u16=0x1234 big-endian */
    CHECK_EQ(buf[1], 0x12);
    CHECK_EQ(buf[2], 0x34);

    BigEndianFields d{};
    CHECK_EQ((int)d.decode(buf.data()), (int)BigEndianFields::size);
    CHECK_EQ((int)d.u16,   0x1234);
    CHECK_EQ((long long)d.i32, -1LL);
    CHECK_NEAR(d.f32, 3.14f, 1e-3f);
    CHECK_EQ((int)d.arr[0],  100);
    CHECK_EQ((int)d.arr[1], -200);
    CHECK_EQ((int)d.arr[2],  300);
    CHECK_EQ((int)d.arr[3], -400);
}

/* ------------------------------------------------------------------ */
/* Test 4: ArrayFields                                                 */
/* ------------------------------------------------------------------ */
static void test_array_fields() {
    g_current = "ArrayFields";
    ArrayFields s{};
    s.u8_arr[0] = 0;   s.u8_arr[1] = 127;
    s.u8_arr[2] = 128; s.u8_arr[3] = 255;
    s.i16_arr[0] = -32768; s.i16_arr[1] = 0;  s.i16_arr[2] = 32767;
    s.color_arr[0] = Color::RED;  s.color_arr[1] = Color::BLUE;
    s.point_arr[0].x = -100; s.point_arr[0].y =  200;
    s.point_arr[1].x =  30000; s.point_arr[1].y = -30000;

    std::vector<uint8_t> buf(ArrayFields::size);
    CHECK_EQ((int)s.encode(buf.data()), (int)ArrayFields::size);

    ArrayFields d{};
    CHECK_EQ((int)d.decode(buf.data()), (int)ArrayFields::size);

    CHECK_EQ((int)d.u8_arr[0],   0);
    CHECK_EQ((int)d.u8_arr[3], 255);
    CHECK_EQ((int)d.i16_arr[0], -32768);
    CHECK_EQ((int)d.i16_arr[2],  32767);
    CHECK_EQ(d.color_arr[0], Color::RED);
    CHECK_EQ(d.color_arr[1], Color::BLUE);
    CHECK_EQ((int)d.point_arr[0].x,  -100);
    CHECK_EQ((int)d.point_arr[0].y,   200);
    CHECK_EQ((int)d.point_arr[1].x, 30000);
    CHECK_EQ((int)d.point_arr[1].y,-30000);
}

/* ------------------------------------------------------------------ */
/* Test 5: EmbedStructs                                               */
/* ------------------------------------------------------------------ */
static void test_embed_structs() {
    g_current = "EmbedStructs";
    EmbedStructs s{};
    s.id    = 42;
    s.ax    = 10;
    s.ay    = 20;
    s.pt.x  = -500;
    s.pt.y  =  500;
    s.flags = 0xFF;

    std::vector<uint8_t> buf(EmbedStructs::size);
    CHECK_EQ((int)s.encode(buf.data()), (int)EmbedStructs::size);

    EmbedStructs d{};
    CHECK_EQ((int)d.decode(buf.data()), (int)EmbedStructs::size);
    CHECK_EQ((int)d.id,   42);
    CHECK_EQ((int)d.ax,   10);
    CHECK_EQ((int)d.ay,   20);
    CHECK_EQ((int)d.pt.x, -500);
    CHECK_EQ((int)d.pt.y,  500);
    CHECK_EQ(d.flags, (uint8_t)0xFF);
}

/* ------------------------------------------------------------------ */
/* Test 6: ConstantFields                                             */
/* ------------------------------------------------------------------ */
static void test_constant_fields() {
    g_current = "ConstantFields";
    ConstantFields s{};
    s.length = 1024;

    std::vector<uint8_t> buf(ConstantFields::size);
    CHECK_EQ((int)s.encode(buf.data()), (int)ConstantFields::size);

    CHECK_EQ(buf[0], 0xAA);
    CHECK_EQ(buf[1], 0x02);
    CHECK_EQ(buf[4], 0x02);

    ConstantFields d{};
    CHECK_EQ((int)d.decode(buf.data()), (int)ConstantFields::size);
    CHECK_EQ((int)d.length, 1024);

    /* Corrupt and verify decode failure */
    buf[0] = 0xFF;
    CHECK_EQ((int)d.decode(buf.data()), -1);
    buf[0] = 0xAA;
    buf[1] = 0x00;
    CHECK_EQ((int)d.decode(buf.data()), -1);
}

/* ------------------------------------------------------------------ */
/* Test 7: GetterSetter                                               */
/* ------------------------------------------------------------------ */
static void test_getter_setter() {
    g_current = "GetterSetter";
    GetterSetter s{};

    s.set_voltage(3.3);
    CHECK_NEAR(s.get_voltage(), 3.3, 5e-3);

    s.set_voltage(0.0);
    CHECK_NEAR(s.get_voltage(), 0.0, 1e-6);

    s.set_celsius(36.5);
    CHECK_NEAR(s.get_celsius(), 36.5, 1e-2);

    s.set_celsius(-40.0);
    CHECK_NEAR(s.get_celsius(), -40.0, 1e-2);

    /* Round-trip through encode/decode */
    GetterSetter s2{};
    s2.set_voltage(1.65);
    s2.set_celsius(25.0);

    std::vector<uint8_t> buf(GetterSetter::size);
    s2.encode(buf.data());

    GetterSetter d{};
    CHECK_EQ((int)d.decode(buf.data()), (int)GetterSetter::size);
    CHECK_NEAR(d.get_voltage(), 1.65, 5e-3);
    CHECK_NEAR(d.get_celsius(), 25.0, 0.1);
}

/* ------------------------------------------------------------------ */
/* Test 8: DynamicFields                                              */
/* ------------------------------------------------------------------ */
static void test_dynamic_fields() {
    g_current = "DynamicFields";

    /* --- non-empty label + data --- */
    uint8_t blob[] = {0xDE, 0xAD, 0xBE, 0xEF, 0x00, 0x01, 0x02};
    DynamicFields s{};
    s.id    = 0xDEADBEEFU;
    s.label = "hello, bubbler!";
    s.data  = make_bytes_ref(blob, sizeof(blob));

    uint64_t sz = s.encode_size();
    std::vector<uint8_t> buf(sz);
    CHECK_EQ((long long)s.encode(buf.data()), (long long)sz);

    DynamicFields d{};
    int64_t dec = d.decode(buf.data());
    CHECK(dec > 0);
    CHECK_EQ(d.id, 0xDEADBEEFU);
    CHECK(d.label == "hello, bubbler!");
    CHECK_EQ((int)d.data.length, (int)sizeof(blob));
    CHECK(memcmp(d.data.data.get(), blob, sizeof(blob)) == 0);

    /* --- empty string, zero-length bytes --- */
    DynamicFields s2{};
    s2.id    = 0;
    s2.label = "";
    s2.data  = make_bytes_ref(nullptr, 0);

    std::vector<uint8_t> buf2(s2.encode_size());
    s2.encode(buf2.data());

    DynamicFields d2{};
    int64_t dec2 = d2.decode(buf2.data());
    CHECK(dec2 > 0);
    CHECK_EQ((int)d2.id, 0);
    CHECK(d2.label.empty());
    CHECK_EQ((int)d2.data.length, 0);

    /* --- UTF-8 string round-trip --- */
    const char *utf8 = "\xe4\xb8\xad\xe6\x96\x87"; /* 中文 */
    DynamicFields s3{};
    s3.id    = 1;
    s3.label = std::string(utf8);
    std::vector<uint8_t> buf3(s3.encode_size());
    s3.encode(buf3.data());
    DynamicFields d3{};
    d3.decode(buf3.data());
    CHECK(d3.label == utf8);
}

/* ------------------------------------------------------------------ */
/* Test 9: NarrowBWTest (bitwid)                                       */
/* Covers narrow bit-width arrays: element bits < type bits           */
/* ------------------------------------------------------------------ */
static void test_narrow_bw() {
    g_current = "NarrowBWTest";

    bitwid::NarrowBWTest s{};
    /* narrow12: int16[4] in 6 bytes = 12 bits/elem, range [-2048, 2047] */
    s.narrow12[0] =  2047;
    s.narrow12[1] = -2048;
    s.narrow12[2] =     0;
    s.narrow12[3] =    -1;
    /* narrow16: int32[3] in 6 bytes = 16 bits/elem, range [-32768, 32767] */
    s.narrow16[0] =  32767;
    s.narrow16[1] = -32768;
    s.narrow16[2] =      0;
    /* narrow24: int64[2] in 6 bytes = 24 bits/elem, range [-8388608, 8388607] */
    s.narrow24[0] =  8388607LL;
    s.narrow24[1] = -8388608LL;
    /* narrow6: uint8[4] in 3 bytes = 6 bits/elem, range [0, 63] */
    s.narrow6[0] = 63;
    s.narrow6[1] =  0;
    s.narrow6[2] = 32;
    s.narrow6[3] =  1;

    CHECK_EQ((int)bitwid::NarrowBWTest::size, 21);
    std::vector<uint8_t> buf(bitwid::NarrowBWTest::size);
    CHECK_EQ((int)s.encode(buf.data()), 21);

    bitwid::NarrowBWTest d{};
    CHECK_EQ((int)d.decode(buf.data()), 21);

    CHECK_EQ((int)d.narrow12[0],  2047);
    CHECK_EQ((int)d.narrow12[1], -2048);
    CHECK_EQ((int)d.narrow12[2],     0);
    CHECK_EQ((int)d.narrow12[3],    -1);

    CHECK_EQ((int)d.narrow16[0],  32767);
    CHECK_EQ((int)d.narrow16[1], -32768);
    CHECK_EQ((int)d.narrow16[2],      0);

    CHECK_EQ((long long)d.narrow24[0],  8388607LL);
    CHECK_EQ((long long)d.narrow24[1], -8388608LL);

    CHECK_EQ((int)d.narrow6[0], 63);
    CHECK_EQ((int)d.narrow6[1],  0);
    CHECK_EQ((int)d.narrow6[2], 32);
    CHECK_EQ((int)d.narrow6[3],  1);

    /* Zero round-trip */
    bitwid::NarrowBWTest z{};
    std::vector<uint8_t> zbuf(bitwid::NarrowBWTest::size, 0);
    z.encode(zbuf.data());
    bitwid::NarrowBWTest dz{};
    dz.decode(zbuf.data());
    CHECK_EQ((int)dz.narrow12[0], 0);
    CHECK_EQ((int)dz.narrow6[3],  0);
}

/* ------------------------------------------------------------------ */
/* Entry point                                                         */
/* ------------------------------------------------------------------ */
int main() {
    printf("=== Bubbler E2E Test — C++ ===\n\n");
    test_simple_scalars();
    test_bit_packed();
    test_big_endian_fields();
    test_array_fields();
    test_embed_structs();
    test_constant_fields();
    test_getter_setter();
    test_dynamic_fields();
    test_narrow_bw();
    printf("\n=== Results: %d passed, %d failed ===\n", g_pass, g_fail);
    return g_fail > 0 ? 1 : 0;
}
