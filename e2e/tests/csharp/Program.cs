/*
 * Bubbler E2E Test — C# target
 *
 * Build (from this directory, after code-gen step):
 *   dotnet build
 * Run:
 *   dotnet run
 *
 * The run_tests.sh script generates gen/ then invokes dotnet run.
 */
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using Testpkg;
using Bitwid;

int pass = 0, fail = 0;
string current = "";

void Check(bool cond, string msg) {
    if (cond) { pass++; }
    else {
        Console.Error.WriteLine($"  FAIL [{current}] {msg}");
        fail++;
    }
}

void CheckEq<T>(T a, T b, string msg)
    => Check(EqualityComparer<T>.Default.Equals(a, b), $"{msg}: got {a}, want {b}");

void CheckNear(double a, double b, double eps, string msg)
    => Check(Math.Abs(a - b) < eps, $"{msg}: got {a}, want ~{b}");

void SetDynamicData(DynamicFields msg, byte[] data) {
    var prop = typeof(DynamicFields).GetProperty("Data");
    if (prop == null) throw new InvalidOperationException("DynamicFields.Data not found");

    if (prop.PropertyType == typeof(byte[])) {
        prop.SetValue(msg, data);
        return;
    }
    if (prop.PropertyType == typeof(Memory<byte>)) {
        prop.SetValue(msg, new Memory<byte>(data));
        return;
    }
    if (prop.PropertyType == typeof(ReadOnlyMemory<byte>)) {
        prop.SetValue(msg, new ReadOnlyMemory<byte>(data));
        return;
    }

    throw new InvalidOperationException($"Unsupported DynamicFields.Data type: {prop.PropertyType}");
}

int DecodeSizeDynamic(DynamicFields msg, byte[] data, int start = 0) {
    var t = typeof(DynamicFields);

    var m2 = t.GetMethod("DecodeSize", new[] { typeof(byte[]), typeof(int) });
    if (m2 != null) {
        return (int)m2.Invoke(msg, new object[] { data, start })!;
    }

    var m1 = t.GetMethod("DecodeSize", new[] { typeof(byte[]) });
    if (m1 != null) {
        if (start == 0) {
            return (int)m1.Invoke(msg, new object[] { data })!;
        }
        byte[] sliced = data.Skip(start).ToArray();
        return (int)m1.Invoke(msg, new object[] { sliced })!;
    }

    throw new InvalidOperationException("DynamicFields.DecodeSize overloads are not available");
}

/* ------------------------------------------------------------------ */
/* Test 1: SimpleScalars                                               */
/* ------------------------------------------------------------------ */
current = "SimpleScalars";
{
    var s = new SimpleScalars {
        U8Zero = 0,
        U8Max  = 255,
        I8Min  = -128,
        I8Max  = 127,
        U16Val = 0xBEEF,
        I16Neg = -1234,
        U32Val = 0xDEADBEEF,
        I32Neg = -100000,
        U64Val = 0xCAFEBABEDEADBEEF,
        I64Neg = -1,
        F32Val = 3.14159f,
        F64Val = 2.718281828459045,
        FlagT  = true,
        FlagF  = false,
        Color  = Color.BLUE,
    };
    byte[] buf = s.Encode();
    Check(buf.Length == SimpleScalars.Size, "encode length");

    var d = new SimpleScalars();
    Check(d.Decode(buf) == SimpleScalars.Size, "decode length");

    CheckEq(d.U8Zero, (byte)0,   "u8_zero");
    CheckEq(d.U8Max,  (byte)255, "u8_max");
    CheckEq(d.I8Min,  (sbyte)-128, "i8_min");
    CheckEq(d.I8Max,  (sbyte)127,  "i8_max");
    CheckEq((ushort)d.U16Val, (ushort)0xBEEF, "u16_val");
    CheckEq(d.I16Neg, (short)-1234, "i16_neg");
    CheckEq(d.U32Val, 0xDEADBEEFu, "u32_val");
    CheckEq(d.I32Neg, -100000, "i32_neg");
    CheckEq(d.U64Val, 0xCAFEBABEDEADBEEFul, "u64_val");
    CheckEq(d.I64Neg, -1L, "i64_neg");
    CheckNear(d.F32Val, 3.14159f, 1e-3, "f32_val");
    CheckNear(d.F64Val, 2.718281828459045, 1e-12, "f64_val");
    CheckEq(d.FlagT, true,  "flag_t");
    CheckEq(d.FlagF, false, "flag_f");
    CheckEq(d.Color, Color.BLUE, "color");
}

/* ------------------------------------------------------------------ */
/* Test 2: BitPacked                                                   */
/* ------------------------------------------------------------------ */
current = "BitPacked";
{
    // Case 1: typical values
    var s = new BitPacked {
        B0 = true, B1 = false, Nibble = 0xA,
        I2 = -1, U4 = 0xF, I20 = 524287, I48 = 0x7FFFFFFFFFFFL,
        Status = Status.BROKEN,
    };
    byte[] buf = s.Encode();
    var d = new BitPacked();
    Check(d.Decode(buf) == BitPacked.Size, "decode size");
    CheckEq(d.I2,  (sbyte)-1,   "i2 -1");
    CheckEq(d.I20, 524287,      "i20 max");
    CheckEq(d.I48, 0x7FFFFFFFFFFFL, "i48 max");
    CheckEq(d.Status, Status.BROKEN, "status");
    CheckEq(d.Nibble, (byte)0xA, "nibble");

    // Case 2: min signed
    s.I2  = -2;
    s.I20 = -524288;
    s.I48 = -1L;
    buf = s.Encode();
    d = new BitPacked();
    d.Decode(buf);
    CheckEq(d.I2,  (sbyte)-2,   "i2 -2");
    CheckEq(d.I20, -524288,     "i20 min");
    CheckEq(d.I48, -1L,         "i48 -1");

    // Case 3: zeros
    s = new BitPacked();
    buf = s.Encode();
    d = new BitPacked();
    d.Decode(buf);
    CheckEq(d.I2,  (sbyte)0, "i2 zero");
    CheckEq(d.I20, 0,        "i20 zero");
    CheckEq(d.I48, 0L,       "i48 zero");
}

/* ------------------------------------------------------------------ */
/* Test 3: BigEndianFields                                             */
/* ------------------------------------------------------------------ */
current = "BigEndianFields";
{
    var s = new BigEndianFields {
        U16 = 0x1234,
        I32 = -1,
        F32 = 3.14f,
        Arr = new short[] { 100, -200, 300, -400 },
    };
    byte[] buf = s.Encode();
    CheckEq(buf[0], (byte)0xBE, "buf[0]=0xBE");
    CheckEq(buf[1], (byte)0x12, "buf[1]=0x12 BE hi");
    CheckEq(buf[2], (byte)0x34, "buf[2]=0x34 BE lo");

    var d = new BigEndianFields();
    Check(d.Decode(buf) == BigEndianFields.Size, "decode size");
    CheckEq((ushort)d.U16, (ushort)0x1234, "u16");
    CheckEq(d.I32, -1, "i32");
    CheckNear(d.F32, 3.14f, 1e-3, "f32");
    CheckEq(d.Arr[0],  (short)100,  "arr[0]");
    CheckEq(d.Arr[1], (short)-200,  "arr[1]");
    CheckEq(d.Arr[2],  (short)300,  "arr[2]");
    CheckEq(d.Arr[3], (short)-400,  "arr[3]");
}

/* ------------------------------------------------------------------ */
/* Test 4: ArrayFields                                                 */
/* ------------------------------------------------------------------ */
current = "ArrayFields";
{
    var s = new ArrayFields {
        U8Arr    = new byte[]  { 0, 127, 128, 255 },
        I16Arr   = new short[] { -32768, 0, 32767 },
        ColorArr = new Color[] { Color.RED, Color.BLUE },
        PointArr = new Point[] {
            new Point { X = -100, Y = 200 },
            new Point { X = 30000, Y = -30000 },
        },
    };
    byte[] buf = s.Encode();
    var d = new ArrayFields();
    Check(d.Decode(buf) == ArrayFields.Size, "decode size");

    CheckEq(d.U8Arr[0], (byte)0,   "u8[0]");
    CheckEq(d.U8Arr[1], (byte)127, "u8[1]");
    CheckEq(d.U8Arr[2], (byte)128, "u8[2]");
    CheckEq(d.U8Arr[3], (byte)255, "u8[3]");
    CheckEq(d.I16Arr[0], (short)-32768, "i16[0]");
    CheckEq(d.I16Arr[1], (short)0,      "i16[1]");
    CheckEq(d.I16Arr[2], (short)32767,  "i16[2]");
    CheckEq(d.ColorArr[0], Color.RED,  "color[0]");
    CheckEq(d.ColorArr[1], Color.BLUE, "color[1]");
    CheckEq(d.PointArr[0].X, (short)-100,  "pt[0].x");
    CheckEq(d.PointArr[0].Y, (short)200,   "pt[0].y");
    CheckEq(d.PointArr[1].X, (short)30000, "pt[1].x");
    CheckEq(d.PointArr[1].Y, (short)-30000,"pt[1].y");
}

/* ------------------------------------------------------------------ */
/* Test 5: EmbedStructs                                               */
/* ------------------------------------------------------------------ */
current = "EmbedStructs";
{
    var s = new EmbedStructs {
        Id = 42, Ax = 10, Ay = 20,
        Pt = new Point { X = -500, Y = 500 },
        Flags = 0xFF,
    };
    byte[] buf = s.Encode();
    var d = new EmbedStructs();
    Check(d.Decode(buf) == EmbedStructs.Size, "decode size");
    CheckEq(d.Id, (byte)42, "id");
    CheckEq(d.Ax, (byte)10, "ax");
    CheckEq(d.Ay, (byte)20, "ay");
    CheckEq(d.Pt.X, (short)-500, "pt.x");
    CheckEq(d.Pt.Y, (short)500,  "pt.y");
    CheckEq(d.Flags, (byte)0xFF, "flags");
}

/* ------------------------------------------------------------------ */
/* Test 6: ConstantFields                                             */
/* ------------------------------------------------------------------ */
current = "ConstantFields";
{
    var s = new ConstantFields { Length = 1024 };
    byte[] buf = s.Encode();
    CheckEq(buf[0], (byte)0xAA, "header=0xAA");
    CheckEq(buf[1], (byte)0x02, "version=0x02");
    CheckEq(buf[4], (byte)0x02, "magic_color=0x02");

    var d = new ConstantFields();
    Check(d.Decode(buf) == ConstantFields.Size, "happy decode");
    CheckEq(d.Length, (ushort)1024, "length");

    byte[] bad = (byte[])buf.Clone();
    bad[0] = 0xFF;
    Check(d.Decode(bad) == -1, "bad header → -1");

    byte[] bad2 = (byte[])buf.Clone();
    bad2[1] = 0x00;
    Check(d.Decode(bad2) == -1, "bad version → -1");
}

/* ------------------------------------------------------------------ */
/* Test 7: GetterSetter                                               */
/* ------------------------------------------------------------------ */
current = "GetterSetter";
{
    var s = new GetterSetter();

    s.Voltage = 3.3;
    CheckNear(s.Voltage, 3.3, 5e-3, "voltage 3.3");

    s.Voltage = 0.0;
    CheckNear(s.Voltage, 0.0, 1e-6, "voltage 0");

    s.Celsius = 36.5;
    CheckNear(s.Celsius, 36.5, 1e-2, "celsius 36.5");

    s.Celsius = -40.0;
    CheckNear(s.Celsius, -40.0, 1e-2, "celsius -40");

    var s2 = new GetterSetter();
    s2.Voltage = 1.65;
    s2.Celsius = 25.0;
    byte[] buf = s2.Encode();
    var d = new GetterSetter();
    Check(d.Decode(buf) == GetterSetter.Size, "decode size");
    CheckNear(d.Voltage, 1.65, 5e-3, "rt voltage");
    CheckNear(d.Celsius, 25.0, 0.1,  "rt celsius");
}

/* ------------------------------------------------------------------ */
/* Test 8: DynamicFields                                              */
/* ------------------------------------------------------------------ */
current = "DynamicFields";
{
    byte[] blob = { 0xDE, 0xAD, 0xBE, 0xEF, 0x00, 0x01, 0x02 };
    var s = new DynamicFields {
        Id    = 0xDEADBEEF,
        Label = "hello, bubbler!",
    };
    SetDynamicData(s, blob);
    var buf = s.Encode();
    var d = new DynamicFields();
    int dec = d.Decode(buf.ToArray());
    Check(dec > 0, "non-empty decode > 0");
    CheckEq(d.Id, 0xDEADBEEFu, "id");
    CheckEq(d.Label, "hello, bubbler!", "label");
    Check(d.Data.ToArray().SequenceEqual(blob), "data bytes");

    // Empty
    var s2 = new DynamicFields { Id = 0, Label = "" };
    SetDynamicData(s2, Array.Empty<byte>());
    var buf2 = s2.Encode();
    var d2 = new DynamicFields();
    Check(d2.Decode(buf2.ToArray()) > 0, "empty decode > 0");
    CheckEq(d2.Label, "", "empty label");
    Check(d2.Data.Length == 0, "empty data");

    // UTF-8
    string utf8 = "\u4e2d\u6587"; // 中文
    var s3 = new DynamicFields { Id = 1, Label = utf8 };
    SetDynamicData(s3, Array.Empty<byte>());
    var buf3 = s3.Encode();
    var d3 = new DynamicFields();
    d3.Decode(buf3.ToArray());
    CheckEq(d3.Label, utf8, "utf8 label");

    // DecodeSize boundary checks
    var s4 = new DynamicFields { Id = 7, Label = "probe" };
    SetDynamicData(s4, blob);
    byte[] full = s4.Encode().ToArray();
    var probe = new DynamicFields();
    CheckEq(DecodeSizeDynamic(probe, full), full.Length, "decodeSize complete");
    CheckEq(DecodeSizeDynamic(probe, full.Take(full.Length - 1).ToArray()), -full.Length, "decodeSize truncated 1 byte");

    CheckEq(DecodeSizeDynamic(probe, new byte[] { 1, 0, 0, 0, (byte)'A' }), -6, "decodeSize missing string terminator");
    CheckEq(DecodeSizeDynamic(probe, new byte[] { 1, 0, 0, 0, (byte)'A', 0, 0x80 }), -8, "decodeSize truncated bytes varint");
    CheckEq(DecodeSizeDynamic(probe, new byte[] { 1, 0, 0, 0, (byte)'A', 0, 0x03, 0xAA, 0xBB }), -10, "decodeSize truncated bytes payload");
    CheckEq(DecodeSizeDynamic(probe, new byte[] { 1, 0, 0, 0 }), -5, "decodeSize only fixed header");
}

/* ------------------------------------------------------------------ */
/* Test 9: NarrowBWTest (bitwid)                                       */
/* Covers narrow bit-width arrays: element bits < type bits           */
/* ------------------------------------------------------------------ */
current = "NarrowBWTest";
{
    var s = new NarrowBWTest {
        Narrow12 = new short[] { 2047, -2048, 0, -1 },
        Narrow16 = new int[]   { 32767, -32768, 0 },
        Narrow24 = new long[]  { 8388607L, -8388608L },
        Narrow6  = new byte[]  { 63, 0, 32, 1 },
    };

    var buf = new byte[21];
    CheckEq(s.Encode(buf, 0), 21, "encode length");

    var d = new NarrowBWTest();
    CheckEq(d.Decode(buf, 0), 21, "decode length");

    CheckEq((int)d.Narrow12[0],  2047, "narrow12[0]");
    CheckEq((int)d.Narrow12[1], -2048, "narrow12[1]");
    CheckEq((int)d.Narrow12[2],     0, "narrow12[2]");
    CheckEq((int)d.Narrow12[3],    -1, "narrow12[3]");

    CheckEq(d.Narrow16[0],  32767, "narrow16[0]");
    CheckEq(d.Narrow16[1], -32768, "narrow16[1]");
    CheckEq(d.Narrow16[2],      0, "narrow16[2]");

    CheckEq(d.Narrow24[0],  8388607L, "narrow24[0]");
    CheckEq(d.Narrow24[1], -8388608L, "narrow24[1]");

    CheckEq(d.Narrow6[0], (byte)63, "narrow6[0]");
    CheckEq(d.Narrow6[1], (byte) 0, "narrow6[1]");
    CheckEq(d.Narrow6[2], (byte)32, "narrow6[2]");
    CheckEq(d.Narrow6[3], (byte) 1, "narrow6[3]");

    // Zero round-trip
    var z = new NarrowBWTest();
    var zbuf = new byte[21];
    z.Encode(zbuf, 0);
    var dz = new NarrowBWTest();
    dz.Decode(zbuf, 0);
    CheckEq((int)dz.Narrow12[0], 0, "zero narrow12[0]");
    CheckEq(dz.Narrow6[3], (byte)0, "zero narrow6[3]");
}

/* ------------------------------------------------------------------ */
/* Summary                                                             */
/* ------------------------------------------------------------------ */
Console.WriteLine($"\n=== Bubbler E2E Test — C# ===");
Console.WriteLine($"=== Results: {pass} passed, {fail} failed ===");
Environment.Exit(fail > 0 ? 1 : 0);
