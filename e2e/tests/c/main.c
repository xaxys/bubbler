/*
 * Bubbler E2E Test — C target
 *
 * Build (from this directory, after code-gen step):
 *   gcc -std=c11 -Igen -o run_test main.c gen/testpkg.bb.c gen/bitwid.bb.c -lm
 *
 * The run_tests.sh script generates gen/ then invokes this.
 */
#include "gen/testpkg.bb.h"
#include "gen/bitwid.bb.h"
#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/* ------------------------------------------------------------------ */
/* Minimal test framework                                               */
/* ------------------------------------------------------------------ */
static int g_pass = 0;
static int g_fail = 0;
static const char *g_current = "";

static void _check(int cond, const char *expr, const char *file, int line) {
    if (cond) {
        g_pass++;
    } else {
        fprintf(stderr, "  FAIL [%s] %s:%d  %s\n", g_current, file, line, expr);
        g_fail++;
    }
}

#define CHECK(cond)           _check(!!(cond), #cond, __FILE__, __LINE__)
#define CHECK_EQ(a,b)         _check((a)==(b), #a " == " #b, __FILE__, __LINE__)
#define CHECK_NEAR(a,b,eps)   _check(fabs((double)(a)-(double)(b))<(eps), \
                                  #a " ~= " #b, __FILE__, __LINE__)

/* ------------------------------------------------------------------ */
/* Test 1: SimpleScalars                                               */
/* Covers all primitive types; min/max/negative edge values            */
/* ------------------------------------------------------------------ */
static void test_simple_scalars(void) {
    g_current = "SimpleScalars";
    struct SimpleScalars s = {0};
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
    s.color   = BLUE;

    uint8_t buf[SimpleScalars_size];
    memset(buf, 0, sizeof(buf));
    uint64_t enc = SimpleScalars_encode(&s, buf);
    CHECK_EQ((int)enc, (int)SimpleScalars_size);

    struct SimpleScalars d = {0};
    int64_t dec = SimpleScalars_decode(buf, &d);
    CHECK_EQ((int)dec, (int)SimpleScalars_size);

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
    CHECK_EQ((int)d.flag_t,    1);
    CHECK_EQ((int)d.flag_f,    0);
    CHECK_EQ((int)d.color,     (int)BLUE);
}

/* ------------------------------------------------------------------ */
/* Test 2: BitPacked                                                   */
/* Covers sub-byte fields, cross-byte signed fields, sign extension    */
/* ------------------------------------------------------------------ */
static void test_bit_packed(void) {
    g_current = "BitPacked";
    struct BitPacked s = {0};

    /* Case 1: typical values */
    s.b0     = true;
    s.b1     = false;
    s.nibble = 0xA;
    s.i2     = -1;      /* 2-bit signed: 0b11 → sign-extend → -1 */
    s.u4     = 0xF;
    s.i20    = 524287;  /* max 20-bit positive = 0x7FFFF */
    s.i48    = 0x7FFFFFFFFFFFLL; /* max 48-bit positive */
    s.status = BROKEN;

    uint8_t buf[BitPacked_size];
    memset(buf, 0, sizeof(buf));
    CHECK_EQ((int)BitPacked_encode(&s, buf), (int)BitPacked_size);

    struct BitPacked d = {0};
    CHECK_EQ((int)BitPacked_decode(buf, &d), (int)BitPacked_size);

    CHECK_EQ((int)d.b0,     1);
    CHECK_EQ((int)d.b1,     0);
    CHECK_EQ((int)d.nibble, 0xA);
    CHECK_EQ((int)d.i2,     -1);
    CHECK_EQ((int)d.u4,     0xF);
    CHECK_EQ((long long)d.i20, 524287LL);
    CHECK_EQ((long long)d.i48, 0x7FFFFFFFFFFFLL);
    CHECK_EQ((int)d.status, (int)BROKEN);

    /* Case 2: sign-extension boundary values */
    s.i2  = -2;        /* min 2-bit signed */
    s.i20 = -524288;   /* min 20-bit signed = 0x80000 */
    s.i48 = -1LL;      /* 48-bit all-ones */
    memset(buf, 0, sizeof(buf));
    BitPacked_encode(&s, buf);
    memset(&d, 0, sizeof(d));
    BitPacked_decode(buf, &d);
    CHECK_EQ((int)d.i2,           -2);
    CHECK_EQ((long long)d.i20,    -524288LL);
    CHECK_EQ((long long)d.i48,    -1LL);

    /* Case 3: zero / min positive */
    s.i2  = 0;
    s.i20 = 0;
    s.i48 = 0;
    s.b0  = false;
    s.nibble = 0;
    s.u4  = 0;
    s.status = IDLE;
    memset(buf, 0, sizeof(buf));
    BitPacked_encode(&s, buf);
    memset(&d, 0, sizeof(d));
    BitPacked_decode(buf, &d);
    CHECK_EQ((int)d.i2,   0);
    CHECK_EQ((int)d.i20,  0);
    CHECK_EQ((long long)d.i48, 0LL);
    CHECK_EQ((int)d.status, (int)IDLE);
}

/* ------------------------------------------------------------------ */
/* Test 3: BigEndianFields                                             */
/* Covers big-endian byte order; verifies raw byte layout             */
/* ------------------------------------------------------------------ */
static void test_big_endian_fields(void) {
    g_current = "BigEndianFields";
    struct BigEndianFields s = {0};
    s.u16    = 0x1234;
    s.i32    = -1;
    s.f32    = 3.14f;
    s.arr[0] = 100; s.arr[1] = -200;
    s.arr[2] = 300; s.arr[3] = -400;

    uint8_t buf[BigEndianFields_size];
    memset(buf, 0, sizeof(buf));
    CHECK_EQ((int)BigEndianFields_encode(&s, buf), (int)BigEndianFields_size);

    /* Constant 0xBE written at byte 0 */
    CHECK_EQ(buf[0], 0xBE);
    /* u16=0x1234 big-endian: MSB first */
    CHECK_EQ(buf[1], 0x12);
    CHECK_EQ(buf[2], 0x34);
    /* i32=-1 big-endian: 0xFF FF FF FF */
    CHECK_EQ(buf[3], 0xFF);
    CHECK_EQ(buf[6], 0xFF);

    struct BigEndianFields d = {0};
    CHECK_EQ((int)BigEndianFields_decode(buf, &d), (int)BigEndianFields_size);
    CHECK_EQ((int)d.u16,    0x1234);
    CHECK_EQ((long long)d.i32, -1LL);
    CHECK_NEAR(d.f32, 3.14f, 1e-3f);
    CHECK_EQ((int)d.arr[0],  100);
    CHECK_EQ((int)d.arr[1], -200);
    CHECK_EQ((int)d.arr[2],  300);
    CHECK_EQ((int)d.arr[3], -400);
}

/* ------------------------------------------------------------------ */
/* Test 4: ArrayFields                                                 */
/* Covers primitive, enum, struct arrays; boundary indices            */
/* ------------------------------------------------------------------ */
static void test_array_fields(void) {
    g_current = "ArrayFields";
    struct ArrayFields s = {0};
    s.u8_arr[0] = 0;   s.u8_arr[1] = 127;
    s.u8_arr[2] = 128; s.u8_arr[3] = 255;
    s.i16_arr[0] = -32768; s.i16_arr[1] = 0; s.i16_arr[2] = 32767;
    s.color_arr[0] = RED;  s.color_arr[1] = BLUE;
    s.point_arr[0].x = -100; s.point_arr[0].y =  200;
    s.point_arr[1].x =  30000; s.point_arr[1].y = -30000;

    uint8_t buf[ArrayFields_size];
    memset(buf, 0, sizeof(buf));
    CHECK_EQ((int)ArrayFields_encode(&s, buf), (int)ArrayFields_size);

    struct ArrayFields d = {0};
    CHECK_EQ((int)ArrayFields_decode(buf, &d), (int)ArrayFields_size);

    CHECK_EQ((int)d.u8_arr[0],   0);
    CHECK_EQ((int)d.u8_arr[1], 127);
    CHECK_EQ((int)d.u8_arr[2], 128);
    CHECK_EQ((int)d.u8_arr[3], 255);
    CHECK_EQ((int)d.i16_arr[0], -32768);
    CHECK_EQ((int)d.i16_arr[1],      0);
    CHECK_EQ((int)d.i16_arr[2],  32767);
    CHECK_EQ((int)d.color_arr[0], (int)RED);
    CHECK_EQ((int)d.color_arr[1], (int)BLUE);
    CHECK_EQ((int)d.point_arr[0].x,   -100);
    CHECK_EQ((int)d.point_arr[0].y,    200);
    CHECK_EQ((int)d.point_arr[1].x,  30000);
    CHECK_EQ((int)d.point_arr[1].y, -30000);
}

/* ------------------------------------------------------------------ */
/* Test 5: EmbedStructs                                               */
/* Covers anonymous embed + named struct embed                        */
/* ------------------------------------------------------------------ */
static void test_embed_structs(void) {
    g_current = "EmbedStructs";
    struct EmbedStructs s = {0};
    s.id     = 42;
    s.ax     = 10;
    s.ay     = 20;
    s.pt.x   = -500;
    s.pt.y   =  500;
    s.flags  = 0xFF;

    uint8_t buf[EmbedStructs_size];
    memset(buf, 0, sizeof(buf));
    CHECK_EQ((int)EmbedStructs_encode(&s, buf), (int)EmbedStructs_size);

    struct EmbedStructs d = {0};
    CHECK_EQ((int)EmbedStructs_decode(buf, &d), (int)EmbedStructs_size);

    CHECK_EQ((int)d.id,   42);
    CHECK_EQ((int)d.ax,   10);
    CHECK_EQ((int)d.ay,   20);
    CHECK_EQ((int)d.pt.x, -500);
    CHECK_EQ((int)d.pt.y,  500);
    CHECK_EQ((int)d.flags, 0xFF);
}

/* ------------------------------------------------------------------ */
/* Test 6: ConstantFields                                             */
/* Covers constant field encoding; decode failure on mismatch         */
/* ------------------------------------------------------------------ */
static void test_constant_fields(void) {
    g_current = "ConstantFields";
    struct ConstantFields s = {0};
    s.length = 1024;

    uint8_t buf[ConstantFields_size];
    memset(buf, 0, sizeof(buf));
    CHECK_EQ((int)ConstantFields_encode(&s, buf), (int)ConstantFields_size);

    /* Constants must be written verbatim */
    CHECK_EQ(buf[0], 0xAA);          /* header  */
    CHECK_EQ(buf[1], 0x02);          /* version */
    CHECK_EQ(buf[4], 0x02);          /* magic_color = GREEN */

    /* Happy-path decode */
    struct ConstantFields d = {0};
    CHECK_EQ((int)ConstantFields_decode(buf, &d), (int)ConstantFields_size);
    CHECK_EQ((int)d.length, 1024);

    /* Decode must fail when a constant byte is corrupted */
    buf[0] = 0xFF;
    CHECK_EQ((int)ConstantFields_decode(buf, &d), -1);
    buf[0] = 0xAA;  /* restore */
    buf[1] = 0x00;  /* corrupt version */
    CHECK_EQ((int)ConstantFields_decode(buf, &d), -1);
}

/* ------------------------------------------------------------------ */
/* Test 7: GetterSetter                                               */
/* Covers custom getter/setter; set-then-get round-trip               */
/* ------------------------------------------------------------------ */
static void test_getter_setter(void) {
    g_current = "GetterSetter";
    struct GetterSetter s = {0};

    /* voltage = raw_adc * 3.3 / 4096 */
    GetterSetter_set_voltage(&s, 3.3);
    CHECK_NEAR(GetterSetter_get_voltage(&s), 3.3, 5e-3);

    GetterSetter_set_voltage(&s, 0.0);
    CHECK_NEAR(GetterSetter_get_voltage(&s), 0.0, 1e-6);

    /* celsius = raw_temp / 100.0 */
    GetterSetter_set_celsius(&s, 36.5);
    CHECK_NEAR(GetterSetter_get_celsius(&s), 36.5, 1e-2);

    GetterSetter_set_celsius(&s, -40.0);
    CHECK_NEAR(GetterSetter_get_celsius(&s), -40.0, 1e-2);

    /* Encode/decode preserves raw values → getters still correct */
    struct GetterSetter s2 = {0};
    GetterSetter_set_voltage(&s2, 1.65);
    GetterSetter_set_celsius(&s2, 25.0);

    uint8_t buf[GetterSetter_size];
    memset(buf, 0, sizeof(buf));
    GetterSetter_encode(&s2, buf);

    struct GetterSetter d = {0};
    CHECK_EQ((int)GetterSetter_decode(buf, &d), (int)GetterSetter_size);
    CHECK_NEAR(GetterSetter_get_voltage(&d), 1.65, 5e-3);
    CHECK_NEAR(GetterSetter_get_celsius(&d), 25.0,  0.1);
}

/* ------------------------------------------------------------------ */
/* Test 8: DynamicFields                                              */
/* Covers variable-length string + bytes; zero-copy decode semantics  */
/* ------------------------------------------------------------------ */
static void test_dynamic_fields(void) {
    g_current = "DynamicFields";

    /* --- non-empty string and bytes --- */
    const char *label    = "hello, bubbler!";
    uint8_t     blob[]   = {0xDE, 0xAD, 0xBE, 0xEF, 0x00, 0x01, 0x02};
    struct DynamicFields s = {0};
    s.id           = 0xDEADBEEFU;
    s.label        = (char *)label;
    s.data.data    = blob;
    s.data.length  = sizeof(blob);

    uint64_t sz  = DynamicFields_encode_size(&s);
    uint8_t *buf = (uint8_t *)malloc(sz);
    CHECK_EQ((long long)DynamicFields_encode(&s, buf), (long long)sz);

    struct DynamicFields d = {0};
    int64_t dec = DynamicFields_decode(buf, &d);
    CHECK((int)dec > 0);
    CHECK_EQ(d.id, 0xDEADBEEFU);
    CHECK(strcmp(d.label, label) == 0);
    CHECK_EQ((int)d.data.length, (int)sizeof(blob));
    CHECK(memcmp(d.data.data, blob, sizeof(blob)) == 0);
    free(buf);

    /* --- empty string and zero-length bytes --- */
    struct DynamicFields s2 = {0};
    s2.id          = 0;
    s2.label       = "";
    s2.data.data   = NULL;
    s2.data.length = 0;

    uint64_t sz2  = DynamicFields_encode_size(&s2);
    uint8_t *buf2 = (uint8_t *)malloc(sz2 + 1); /* +1 safety */
    DynamicFields_encode(&s2, buf2);

    struct DynamicFields d2 = {0};
    int64_t dec2 = DynamicFields_decode(buf2, &d2);
    CHECK((int)dec2 > 0);
    CHECK_EQ((int)d2.id, 0);
    CHECK(strcmp(d2.label, "") == 0);
    CHECK_EQ((int)d2.data.length, 0);
    free(buf2);

    /* --- multi-byte UTF-8 string --- */
    const char *utf8 = "\xe4\xb8\xad\xe6\x96\x87";  /* 中文 */
    struct DynamicFields s3 = {0};
    s3.id    = 1;
    s3.label = (char *)utf8;
    uint64_t sz3  = DynamicFields_encode_size(&s3);
    uint8_t *buf3 = (uint8_t *)malloc(sz3);
    DynamicFields_encode(&s3, buf3);
    struct DynamicFields d3 = {0};
    DynamicFields_decode(buf3, &d3);
    CHECK(strcmp(d3.label, utf8) == 0);
    free(buf3);
}

/* ------------------------------------------------------------------ */
/* Test 9: NarrowBWTest (bitwid)                                       */
/* Covers narrow bit-width arrays: element bits < type bits           */
/* ------------------------------------------------------------------ */
static void test_narrow_bw(void) {
    g_current = "NarrowBWTest";

    struct NarrowBWTest s = {0};
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

    uint8_t buf[NarrowBWTest_size];
    memset(buf, 0, sizeof(buf));
    CHECK_EQ((int)NarrowBWTest_size, 21);
    CHECK_EQ((int)NarrowBWTest_encode(&s, buf), 21);

    struct NarrowBWTest d = {0};
    CHECK_EQ((int)NarrowBWTest_decode(buf, &d), 21);

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

    /* Verify zero-round-trip */
    struct NarrowBWTest z = {0};
    uint8_t zbuf[NarrowBWTest_size];
    memset(zbuf, 0, sizeof(zbuf));
    NarrowBWTest_encode(&z, zbuf);
    struct NarrowBWTest dz = {0};
    NarrowBWTest_decode(zbuf, &dz);
    CHECK_EQ((int)dz.narrow12[0], 0);
    CHECK_EQ((int)dz.narrow6[3],  0);
}

/* ------------------------------------------------------------------ */
/* Entry point                                                         */
/* ------------------------------------------------------------------ */
int main(void) {
    printf("=== Bubbler E2E Test — C ===\n\n");
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
