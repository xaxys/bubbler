/**
 * Bubbler E2E Test — CommonJS target
 *
 * Run (from this directory, after code-gen step):
 *   node test.mjs
 *
 * The run_tests.sh script generates gen/testcase.bb.js then invokes node.
 */
import { createRequire } from "module";
const require = createRequire(import.meta.url);
const pkg = require("./gen/testcase.bb.js").testpkg;
const bw  = require("./gen/bitwid.bb.js").bitwid;

/* ------------------------------------------------------------------ */
/* Minimal test framework                                               */
/* ------------------------------------------------------------------ */
let g_pass = 0, g_fail = 0, g_current = "";
function check(cond, msg) {
    if (cond) { g_pass++; }
    else { process.stderr.write(`  FAIL [${g_current}] ${msg}\n`); g_fail++; }
}
function checkEq(a, b, msg) {
    check(a === b, `${msg}: got ${a}, want ${b}`);
}
function checkNear(a, b, eps, msg) {
    check(Math.abs(Number(a) - Number(b)) < eps, `${msg}: got ${a}, want ~${b}`);
}

/* ------------------------------------------------------------------ */
/* Test 1: SimpleScalars                                               */
/* ------------------------------------------------------------------ */
g_current = "SimpleScalars";
{
    const s = new pkg.SimpleScalars();
    // Use values that stay positive as JS int32 for uint32 (avoid sign flip)
    s.u8Zero = 0;
    s.u8Max  = 255;
    s.i8Min  = -128;
    s.i8Max  = 127;
    s.u16Val = 0xBEEF;
    s.i16Neg = -1234;
    s.u32Val = 0x12345678;   // stays positive as int32
    s.i32Neg = -100000;
    s.u64Val = BigInt("0xCAFEBABEDEADBEEF");
    s.i64Neg = -1n;
    s.f32Val = 3.14159;
    s.f64Val = 2.718281828459045;
    s.flagT  = true;
    s.flagF  = false;
    s.color  = pkg.Color.BLUE;

    const buf = pkg.SimpleScalars.encode(s);           // returns new Array
    check(buf.length === 47, "encode length 47");

    const d = new pkg.SimpleScalars();
    const n = pkg.SimpleScalars.decode(d, buf);
    check(n === 47, `decode length: got ${n}`);

    checkEq(d.u8Zero, 0,     "u8Zero");
    checkEq(d.u8Max,  255,   "u8Max");
    checkEq(d.i8Min,  -128,  "i8Min");
    checkEq(d.i8Max,  127,   "i8Max");
    checkEq(d.u16Val, 0xBEEF, "u16Val");
    checkEq(d.i16Neg, -1234, "i16Neg");
    checkEq(d.u32Val, 0x12345678, "u32Val");
    checkEq(d.i32Neg, -100000, "i32Neg");
    checkEq(d.u64Val, BigInt("0xCAFEBABEDEADBEEF"), "u64Val");
    checkEq(d.i64Neg, -1n, "i64Neg");
    checkNear(d.f32Val, 3.14159, 1e-3, "f32Val");
    checkNear(d.f64Val, 2.718281828459045, 1e-12, "f64Val");
    checkEq(d.flagT, true, "flagT");
    checkEq(d.flagF, false, "flagF");
    checkEq(d.color, "BLUE", "color");
}

/* ------------------------------------------------------------------ */
/* Test 2: BitPacked — sign extension, bit boundaries                  */
/* ------------------------------------------------------------------ */
g_current = "BitPacked";
{
    // Case 1: typical values
    const s = new pkg.BitPacked();
    s.b0     = true;
    s.b1     = false;
    s.nibble = 0xA;
    s.i2     = -1;
    s.u4     = 0xF;
    s.i20    = 524287;
    s.i48    = 0x7FFFFFFFFFFFn;
    s.status = pkg.Status.BROKEN;

    const buf = pkg.BitPacked.encode(s);
    const d = new pkg.BitPacked();
    check(pkg.BitPacked.decode(d, buf) > 0, "decode > 0");
    checkEq(d.i2,  -1, "i2 -1");
    checkEq(d.i20, 524287, "i20 max");
    checkEq(d.i48, 0x7FFFFFFFFFFFn, "i48 max");
    checkEq(d.status, "BROKEN", "status BROKEN");
    checkEq(d.nibble, 0xA, "nibble");

    // Case 2: min signed values
    s.i2  = -2;
    s.i20 = -524288;
    s.i48 = -1n;
    const buf2 = pkg.BitPacked.encode(s);
    const d2 = new pkg.BitPacked();
    pkg.BitPacked.decode(d2, buf2);
    checkEq(d2.i2,  -2, "i2 -2");
    checkEq(d2.i20, -524288, "i20 min");
    checkEq(d2.i48, -1n, "i48 -1");

    // Case 3: zeros
    const s3 = new pkg.BitPacked();
    const buf3 = pkg.BitPacked.encode(s3);
    const d3 = new pkg.BitPacked();
    pkg.BitPacked.decode(d3, buf3);
    checkEq(d3.i2,  0, "i2 zero");
    checkEq(d3.i20, 0, "i20 zero");
    checkEq(d3.i48, 0n, "i48 zero");
}

/* ------------------------------------------------------------------ */
/* Test 3: BigEndianFields — byte-order verification                   */
/* ------------------------------------------------------------------ */
g_current = "BigEndianFields";
{
    const s = new pkg.BigEndianFields();
    s.u16 = 0x1234;
    s.i32 = -1;
    s.f32 = 3.14;
    s.arr = [100, -200, 300, -400];

    const buf = pkg.BigEndianFields.encode(s);
    checkEq(buf[0], 0xBE, "buf[0]=0xBE constant");
    checkEq(buf[1], 0x12, "buf[1] BE high");
    checkEq(buf[2], 0x34, "buf[2] BE low");

    const d = new pkg.BigEndianFields();
    check(pkg.BigEndianFields.decode(d, buf) > 0, "decode > 0");
    checkEq(d.u16, 0x1234, "u16");
    checkEq(d.i32, -1, "i32 -1");
    checkNear(d.f32, 3.14, 1e-3, "f32");
    checkEq(d.arr[0],  100, "arr[0]");
    checkEq(d.arr[1], -200, "arr[1]");
    checkEq(d.arr[2],  300, "arr[2]");
    checkEq(d.arr[3], -400, "arr[3]");
}

/* ------------------------------------------------------------------ */
/* Test 4: ArrayFields                                                 */
/* ------------------------------------------------------------------ */
g_current = "ArrayFields";
{
    const s = new pkg.ArrayFields();
    s.u8Arr    = [0, 127, 128, 255];
    s.i16Arr   = [-32768, 0, 32767];
    s.colorArr = [pkg.Color.RED, pkg.Color.BLUE];

    // Point objects have x/y properties
    s.pointArr = [];
    const p0 = new pkg.Point(); p0.x = -100; p0.y = 200;
    const p1 = new pkg.Point(); p1.x = 30000; p1.y = -30000;
    s.pointArr.push(p0); s.pointArr.push(p1);

    const buf = pkg.ArrayFields.encode(s);
    const d = new pkg.ArrayFields();
    check(pkg.ArrayFields.decode(d, buf) > 0, "decode > 0");

    checkEq(d.u8Arr[0],   0, "u8Arr[0]");
    checkEq(d.u8Arr[1], 127, "u8Arr[1]");
    checkEq(d.u8Arr[2], 128, "u8Arr[2]");
    checkEq(d.u8Arr[3], 255, "u8Arr[3]");
    checkEq(d.i16Arr[0], -32768, "i16Arr[0]");
    checkEq(d.i16Arr[1],      0, "i16Arr[1]");
    checkEq(d.i16Arr[2],  32767, "i16Arr[2]");
    checkEq(d.colorArr[0], "RED",  "colorArr[0]");
    checkEq(d.colorArr[1], "BLUE", "colorArr[1]");
    checkEq(d.pointArr[0].x,    -100, "pt[0].x");
    checkEq(d.pointArr[0].y,     200, "pt[0].y");
    checkEq(d.pointArr[1].x,  30000, "pt[1].x");
    checkEq(d.pointArr[1].y, -30000, "pt[1].y");
}

/* ------------------------------------------------------------------ */
/* Test 5: EmbedStructs                                               */
/* ------------------------------------------------------------------ */
g_current = "EmbedStructs";
{
    const s = new pkg.EmbedStructs();
    s.id    = 42;
    s.ax    = 10;
    s.ay    = 20;
    const pt = new pkg.Point(); pt.x = -500; pt.y = 500;
    s.pt    = pt;
    s.flags = 0xFF;

    const buf = pkg.EmbedStructs.encode(s);
    const d = new pkg.EmbedStructs();
    check(pkg.EmbedStructs.decode(d, buf) > 0, "decode > 0");
    checkEq(d.id,   42, "id");
    checkEq(d.ax,   10, "ax");
    checkEq(d.ay,   20, "ay");
    checkEq(d.pt.x, -500, "pt.x");
    checkEq(d.pt.y,  500, "pt.y");
    checkEq(d.flags, 0xFF, "flags");
}

/* ------------------------------------------------------------------ */
/* Test 6: ConstantFields                                             */
/* ------------------------------------------------------------------ */
g_current = "ConstantFields";
{
    const s = new pkg.ConstantFields();
    s.length = 1024;

    const buf = pkg.ConstantFields.encode(s);
    checkEq(buf[0], 0xAA, "header=0xAA");
    checkEq(buf[1], 0x02, "version=0x02");
    checkEq(buf[4], 0x02, "magic_color=0x02");

    const d = new pkg.ConstantFields();
    check(pkg.ConstantFields.decode(d, buf) > 0, "happy decode > 0");
    checkEq(d.length, 1024, "length");

    // Corrupt header → decode must fail
    const bad = [...buf];
    bad[0] = 0xFF;
    check(pkg.ConstantFields.decode(new pkg.ConstantFields(), bad) === -1, "bad header → -1");

    // Corrupt version → decode must fail
    const bad2 = [...buf];
    bad2[1] = 0x00;
    check(pkg.ConstantFields.decode(new pkg.ConstantFields(), bad2) === -1, "bad version → -1");
}

/* ------------------------------------------------------------------ */
/* Test 7: GetterSetter                                               */
/* ------------------------------------------------------------------ */
g_current = "GetterSetter";
{
    const s = new pkg.GetterSetter();

    s.voltage = 3.3;
    checkNear(s.voltage, 3.3, 5e-3, "voltage 3.3");

    s.voltage = 0.0;
    checkNear(s.voltage, 0.0, 1e-6, "voltage 0");

    s.celsius = 36.5;
    checkNear(s.celsius, 36.5, 1e-2, "celsius 36.5");

    s.celsius = -40.0;
    checkNear(s.celsius, -40.0, 1e-2, "celsius -40");

    // Round-trip via encode/decode
    const s2 = new pkg.GetterSetter();
    s2.voltage = 1.65;
    s2.celsius = 25.0;
    const buf = pkg.GetterSetter.encode(s2);

    const d = new pkg.GetterSetter();
    check(pkg.GetterSetter.decode(d, buf) > 0, "decode > 0");
    checkNear(d.voltage, 1.65, 5e-3, "rt voltage");
    checkNear(d.celsius, 25.0, 0.1,  "rt celsius");
}

/* ------------------------------------------------------------------ */
/* Test 8: DynamicFields                                              */
/* ------------------------------------------------------------------ */
g_current = "DynamicFields";
{
    // Non-empty label and data
    const blob = [0xDE, 0xAD, 0xBE, 0xEF, 0x00, 0x01, 0x02];
    const s = new pkg.DynamicFields();
    s.id    = 42;         // use small positive value (no int32 sign issue)
    s.label = "hello, bubbler!";
    s.data  = [...blob];  // Array of numbers

    const buf = pkg.DynamicFields.encode(s);
    check(buf instanceof Uint8Array, "encode returns Uint8Array");
    const d = new pkg.DynamicFields();
    const n = pkg.DynamicFields.decode(d, buf);
    check(n > 0, `decode > 0: got ${n}`);
    checkEq(d.id, 42, "id");
    checkEq(d.label, "hello, bubbler!", "label");
    checkEq(d.data.length, blob.length, "data.length");
    for (let i = 0; i < blob.length; i++) {
        checkEq(d.data[i], blob[i], `data[${i}]`);
    }

    // Empty label and data
    const s2 = new pkg.DynamicFields();
    s2.id    = 0;
    s2.label = "";
    s2.data  = [];
    const buf2 = pkg.DynamicFields.encode(s2);
    const d2 = new pkg.DynamicFields();
    check(pkg.DynamicFields.decode(d2, buf2) > 0, "empty decode > 0");
    checkEq(d2.label, "", "empty label");
    checkEq(d2.data.length, 0, "empty data length");

    // UTF-8 string round-trip
    const utf8 = "\u4e2d\u6587"; // 中文
    const s3 = new pkg.DynamicFields();
    s3.id    = 1;
    s3.label = utf8;
    s3.data  = [];
    const buf3 = pkg.DynamicFields.encode(s3);
    const d3 = new pkg.DynamicFields();
    pkg.DynamicFields.decode(d3, buf3);
    checkEq(d3.label, utf8, "utf8 label");

    // Large data block (tests multi-byte length varint for bytes field)
    const largeData = new Array(200).fill(0).map((_, i) => i & 0xFF);
    const s4 = new pkg.DynamicFields();
    s4.id    = 999;
    s4.label = "large";
    s4.data  = largeData;
    const buf4 = pkg.DynamicFields.encode(s4);
    const d4 = new pkg.DynamicFields();
    pkg.DynamicFields.decode(d4, buf4);
    checkEq(d4.data.length, 200, "large data length");
    checkEq(d4.data[199], 199 & 0xFF, "large data last byte");
}

/* ------------------------------------------------------------------ */
/* Test 9: NarrowBWTest (bitwid)                                       */
/* Narrow bit-width arrays: element bits < type bits                  */
/* ------------------------------------------------------------------ */
g_current = "NarrowBWTest";
{
    const s = new bw.NarrowBWTest();
    // narrow12: int16[4] in 6 bytes = 12 bits/elem [-2048, 2047]
    s.narrow12 = [2047, -2048, 0, -1];
    // narrow16: int32[3] in 6 bytes = 16 bits/elem [-32768, 32767]
    s.narrow16 = [32767, -32768, 0];
    // narrow24: int64[2] in 6 bytes = 24 bits/elem [-8388608, 8388607]
    // Note: int64 uses BigInt in JS
    s.narrow24 = [BigInt(8388607), BigInt(-8388608)];
    // narrow6: uint8[4] in 3 bytes = 6 bits/elem [0, 63]
    s.narrow6 = [63, 0, 32, 1];

    const buf = bw.NarrowBWTest.encode(s);
    check(buf.length === 21, `encode length 21, got ${buf.length}`);

    const d = new bw.NarrowBWTest();
    const n = bw.NarrowBWTest.decode(d, buf);
    check(n === 21, `decode length 21, got ${n}`);

    checkEq(d.narrow12[0],  2047, "narrow12[0]");
    checkEq(d.narrow12[1], -2048, "narrow12[1]");
    checkEq(d.narrow12[2],     0, "narrow12[2]");
    checkEq(d.narrow12[3],    -1, "narrow12[3]");

    checkEq(d.narrow16[0],  32767, "narrow16[0]");
    checkEq(d.narrow16[1], -32768, "narrow16[1]");
    checkEq(d.narrow16[2],      0, "narrow16[2]");

    checkEq(d.narrow24[0], BigInt( 8388607), "narrow24[0]");
    checkEq(d.narrow24[1], BigInt(-8388608), "narrow24[1]");

    checkEq(d.narrow6[0], 63, "narrow6[0]");
    checkEq(d.narrow6[1],  0, "narrow6[1]");
    checkEq(d.narrow6[2], 32, "narrow6[2]");
    checkEq(d.narrow6[3],  1, "narrow6[3]");

    // Zero round-trip
    const z = new bw.NarrowBWTest();
    z.narrow12 = [0, 0, 0, 0];
    z.narrow16 = [0, 0, 0];
    z.narrow24 = [0n, 0n];
    z.narrow6  = [0, 0, 0, 0];
    const zbuf = bw.NarrowBWTest.encode(z);
    const dz = new bw.NarrowBWTest();
    bw.NarrowBWTest.decode(dz, zbuf);
    checkEq(dz.narrow12[0], 0, "zero narrow12[0]");
    checkEq(dz.narrow6[3],  0, "zero narrow6[3]");
}

/* ------------------------------------------------------------------ */
/* Summary                                                             */
/* ------------------------------------------------------------------ */
console.log(`\n=== Bubbler E2E Test — CommonJS ===`);
console.log(`=== Results: ${g_pass} passed, ${g_fail} failed ===`);
process.exit(g_fail > 0 ? 1 : 0);
