/*
 * Bubbler E2E Test - Java target
 *
 * Compile (from this directory, after code-gen step):
 *   javac -sourcepath gen -cp gen -d out gen/com/example/testpkg/*.java gen/com/example/bitwid/*.java Main.java
 * Run:
 *   java -cp out:gen Main
 *
 * The run_tests.sh script generates gen/ then invokes javac + java.
 */
import com.example.testpkg.*;
import com.example.bitwid.*;

public class Main {
    private static int pass = 0;
    private static int fail = 0;
    private static String current = "";

    /* ------------------------------------------------------------------ */
    /* Minimal test framework                                               */
    /* ------------------------------------------------------------------ */
    static void check(boolean cond, String msg) {
        if (cond) {
            pass++;
        } else {
            System.err.println("  FAIL [" + current + "] " + msg);
            fail++;
        }
    }

    static void checkEq(long a, long b, String msg) {
        check(a == b, msg + " (got " + a + ", want " + b + ")");
    }

    static void checkNear(double a, double b, double eps, String msg) {
        check(Math.abs(a - b) < eps, msg + " (got " + a + ", want ~" + b + ")");
    }

    /* ------------------------------------------------------------------ */
    /* Test 1: SimpleScalars                                               */
    /* ------------------------------------------------------------------ */
    static void testSimpleScalars() {
        current = "SimpleScalars";
        SimpleScalars s = new SimpleScalars();
        s.setU8Zero((byte) 0);
        s.setU8Max((byte) 255);
        s.setI8Min((byte) -128);
        s.setI8Max((byte) 127);
        s.setU16Val((short) 0xBEEF);
        s.setI16Neg((short) -1234);
        s.setU32Val(0xDEADBEEF);
        s.setI32Neg(-100000);
        s.setU64Val(0xCAFEBABEDEADBEEFL);
        s.setI64Neg(-1L);
        s.setF32Val(3.14159f);
        s.setF64Val(2.718281828459045);
        s.setFlagT(true);
        s.setFlagF(false);
        s.setColor(Color.BLUE);

        byte[] buf = s.encode();
        check(buf.length == SimpleScalars.size(), "encode length");

        SimpleScalars d = new SimpleScalars();
        check(d.decode(buf) == SimpleScalars.size(), "decode length");

        checkEq(Byte.toUnsignedInt(d.getU8Zero()), 0,   "u8_zero");
        checkEq(Byte.toUnsignedInt(d.getU8Max()),  255, "u8_max");
        checkEq((int) d.getI8Min(), -128, "i8_min");
        checkEq((int) d.getI8Max(),  127, "i8_max");
        checkEq(Short.toUnsignedInt(d.getU16Val()), 0xBEEF, "u16_val");
        checkEq(d.getI16Neg(), -1234, "i16_neg");
        checkEq(Integer.toUnsignedLong(d.getU32Val()), 0xDEADBEEFL, "u32_val");
        checkEq(d.getI32Neg(), -100000, "i32_neg");
        check(Long.compareUnsigned(d.getU64Val(), 0xCAFEBABEDEADBEEFL) == 0, "u64_val");
        checkEq(d.getI64Neg(), -1L, "i64_neg");
        checkNear(d.getF32Val(), 3.14159f, 1e-3, "f32_val");
        checkNear(d.getF64Val(), 2.718281828459045, 1e-12, "f64_val");
        check(d.getFlagT() == true,  "flag_t");
        check(d.getFlagF() == false, "flag_f");
        check(d.getColor() == Color.BLUE, "color");
    }

    /* ------------------------------------------------------------------ */
    /* Test 2: BitPacked                                                   */
    /* ------------------------------------------------------------------ */
    static void testBitPacked() {
        current = "BitPacked";

        // Case 1: typical values
        BitPacked s = new BitPacked();
        s.setB0(true);
        s.setB1(false);
        s.setNibble((byte) 0xA);
        s.setI2((byte) -1);
        s.setU4((byte) 0xF);
        s.setI20(524287);
        s.setI48(0x7FFFFFFFFFFFL);
        s.setStatus(Status.BROKEN);

        byte[] buf = s.encode();
        BitPacked d = new BitPacked();
        check(d.decode(buf) == BitPacked.size(), "decode size");

        checkEq(d.getI2(), -1, "i2 -1");
        checkEq(d.getI20(), 524287, "i20 max");
        checkEq(d.getI48(), 0x7FFFFFFFFFFFL, "i48 max");
        check(d.getStatus() == Status.BROKEN, "status BROKEN");
        checkEq(Byte.toUnsignedInt(d.getNibble()), 0xA, "nibble");

        // Case 2: minimum signed values
        s.setI2((byte) -2);
        s.setI20(-524288);
        s.setI48(-1L);
        buf = s.encode();
        d = new BitPacked();
        d.decode(buf);
        checkEq(d.getI2(), -2, "i2 -2");
        checkEq(d.getI20(), -524288, "i20 min");
        checkEq(d.getI48(), -1L, "i48 -1");

        // Case 3: zeros
        s = new BitPacked();
        s.setStatus(Status.IDLE);
        buf = s.encode();
        d = new BitPacked();
        d.decode(buf);
        checkEq(d.getI2(), 0, "i2 zero");
        checkEq(d.getI20(), 0, "i20 zero");
        checkEq(d.getI48(), 0, "i48 zero");
    }

    /* ------------------------------------------------------------------ */
    /* Test 3: BigEndianFields                                             */
    /* ------------------------------------------------------------------ */
    static void testBigEndianFields() {
        current = "BigEndianFields";
        BigEndianFields s = new BigEndianFields();
        s.setU16((short) 0x1234);
        s.setI32(-1);
        s.setF32(3.14f);
        s.setArrAt(0, (short)  100);
        s.setArrAt(1, (short) -200);
        s.setArrAt(2, (short)  300);
        s.setArrAt(3, (short) -400);

        byte[] buf = s.encode();

        checkEq(Byte.toUnsignedInt(buf[0]), 0xBE, "buf[0]=0xBE constant");
        checkEq(Byte.toUnsignedInt(buf[1]), 0x12, "buf[1]=0x12 BE high");
        checkEq(Byte.toUnsignedInt(buf[2]), 0x34, "buf[2]=0x34 BE low");

        BigEndianFields d = new BigEndianFields();
        check(d.decode(buf) == BigEndianFields.size(), "decode size");
        checkEq(Short.toUnsignedInt(d.getU16()), 0x1234, "u16");
        checkEq(d.getI32(), -1, "i32 -1");
        checkNear(d.getF32(), 3.14f, 1e-3, "f32");
        checkEq(d.getArrAt(0),  100, "arr[0]");
        checkEq(d.getArrAt(1), -200, "arr[1]");
        checkEq(d.getArrAt(2),  300, "arr[2]");
        checkEq(d.getArrAt(3), -400, "arr[3]");
    }

    /* ------------------------------------------------------------------ */
    /* Test 4: ArrayFields                                                 */
    /* ------------------------------------------------------------------ */
    static void testArrayFields() {
        current = "ArrayFields";
        ArrayFields s = new ArrayFields();
        s.setU8ArrAt(0, (byte) 0);
        s.setU8ArrAt(1, (byte) 127);
        s.setU8ArrAt(2, (byte) 128);
        s.setU8ArrAt(3, (byte) 255);
        s.setI16ArrAt(0, (short) -32768);
        s.setI16ArrAt(1, (short) 0);
        s.setI16ArrAt(2, (short) 32767);
        s.setColorArrAt(0, Color.RED);
        s.setColorArrAt(1, Color.BLUE);
        Point p0 = new Point();
        p0.setX((short) -100); p0.setY((short) 200);
        Point p1 = new Point();
        p1.setX((short) 30000); p1.setY((short) -30000);
        s.setPointArrAt(0, p0);
        s.setPointArrAt(1, p1);

        byte[] buf = s.encode();
        ArrayFields d = new ArrayFields();
        check(d.decode(buf) == ArrayFields.size(), "decode size");

        checkEq(Byte.toUnsignedInt(d.getU8ArrAt(0)),   0, "u8[0]");
        checkEq(Byte.toUnsignedInt(d.getU8ArrAt(1)), 127, "u8[1]");
        checkEq(Byte.toUnsignedInt(d.getU8ArrAt(2)), 128, "u8[2]");
        checkEq(Byte.toUnsignedInt(d.getU8ArrAt(3)), 255, "u8[3]");
        checkEq(d.getI16ArrAt(0), -32768, "i16[0]");
        checkEq(d.getI16ArrAt(1),      0, "i16[1]");
        checkEq(d.getI16ArrAt(2),  32767, "i16[2]");
        check(d.getColorArrAt(0) == Color.RED,  "color[0]");
        check(d.getColorArrAt(1) == Color.BLUE, "color[1]");
        checkEq(d.getPointArrAt(0).getX(),    -100, "pt[0].x");
        checkEq(d.getPointArrAt(0).getY(),     200, "pt[0].y");
        checkEq(d.getPointArrAt(1).getX(),   30000, "pt[1].x");
        checkEq(d.getPointArrAt(1).getY(),  -30000, "pt[1].y");
    }

    /* ------------------------------------------------------------------ */
    /* Test 5: EmbedStructs                                               */
    /* ------------------------------------------------------------------ */
    static void testEmbedStructs() {
        current = "EmbedStructs";
        EmbedStructs s = new EmbedStructs();
        s.setId((byte) 42);
        s.setAx((byte) 10);
        s.setAy((byte) 20);
        Point pt = new Point();
        pt.setX((short) -500); pt.setY((short) 500);
        s.setPt(pt);
        s.setFlags((byte) 0xFF);

        byte[] buf = s.encode();
        EmbedStructs d = new EmbedStructs();
        check(d.decode(buf) == EmbedStructs.size(), "decode size");

        checkEq(Byte.toUnsignedInt(d.getId()), 42, "id");
        checkEq(Byte.toUnsignedInt(d.getAx()), 10, "ax");
        checkEq(Byte.toUnsignedInt(d.getAy()), 20, "ay");
        checkEq(d.getPt().getX(), -500, "pt.x");
        checkEq(d.getPt().getY(),  500, "pt.y");
        checkEq(Byte.toUnsignedInt(d.getFlags()), 0xFF, "flags");
    }

    /* ------------------------------------------------------------------ */
    /* Test 6: ConstantFields                                             */
    /* ------------------------------------------------------------------ */
    static void testConstantFields() {
        current = "ConstantFields";
        ConstantFields s = new ConstantFields();
        s.setLength((short) 1024);

        byte[] buf = s.encode();
        checkEq(Byte.toUnsignedInt(buf[0]), 0xAA, "header=0xAA");
        checkEq(Byte.toUnsignedInt(buf[1]), 0x02, "version=0x02");
        checkEq(Byte.toUnsignedInt(buf[4]), 0x02, "magic_color=0x02");

        ConstantFields d = new ConstantFields();
        check(d.decode(buf) == ConstantFields.size(), "happy-path decode");
        checkEq(Short.toUnsignedInt(d.getLength()), 1024, "length");

        // Corrupt header -> decode must fail
        byte[] bad = buf.clone();
        bad[0] = (byte) 0xFF;
        check(d.decode(bad) == -1, "bad header -> -1");

        // Corrupt version -> decode must fail
        byte[] bad2 = buf.clone();
        bad2[1] = 0x00;
        check(d.decode(bad2) == -1, "bad version -> -1");
    }

    /* ------------------------------------------------------------------ */
    /* Test 7: GetterSetter                                               */
    /* ------------------------------------------------------------------ */
    static void testGetterSetter() {
        current = "GetterSetter";
        GetterSetter s = new GetterSetter();

        s.setVoltage(3.3);
        checkNear(s.getVoltage(), 3.3, 5e-3, "voltage 3.3");

        s.setVoltage(0.0);
        checkNear(s.getVoltage(), 0.0, 1e-6, "voltage 0");

        s.setCelsius(36.5);
        checkNear(s.getCelsius(), 36.5, 1e-2, "celsius 36.5");

        s.setCelsius(-40.0);
        checkNear(s.getCelsius(), -40.0, 1e-2, "celsius -40");

        // Round-trip via encode/decode
        GetterSetter s2 = new GetterSetter();
        s2.setVoltage(1.65);
        s2.setCelsius(25.0);
        byte[] buf = s2.encode();

        GetterSetter d = new GetterSetter();
        check(d.decode(buf) == GetterSetter.size(), "decode size");
        checkNear(d.getVoltage(), 1.65, 5e-3, "rt voltage");
        checkNear(d.getCelsius(), 25.0, 0.1,  "rt celsius");
    }

    /* ------------------------------------------------------------------ */
    /* Test 8: DynamicFields                                              */
    /* ------------------------------------------------------------------ */
    static void testDynamicFields() {
        current = "DynamicFields";

        // Non-empty label and data
        byte[] blob = {(byte) 0xDE, (byte) 0xAD, (byte) 0xBE, (byte) 0xEF,
                       0x00, 0x01, 0x02};
        DynamicFields s = new DynamicFields();
        s.setId(0xDEADBEEF);
        s.setLabel("hello, bubbler!");
        s.setData(blob);

        byte[] buf = s.encode();
        DynamicFields d = new DynamicFields();
        check(d.decode(buf) > 0, "non-empty decode > 0");
        check(Integer.toUnsignedLong(d.getId()) == 0xDEADBEEFL, "id");
        check("hello, bubbler!".equals(d.getLabel()), "label");
        check(java.util.Arrays.equals(d.getData(), blob), "data bytes");

        // Empty label and data
        DynamicFields s2 = new DynamicFields();
        s2.setId(0);
        s2.setLabel("");
        s2.setData(new byte[0]);
        byte[] buf2 = s2.encode();
        DynamicFields d2 = new DynamicFields();
        check(d2.decode(buf2) > 0, "empty decode > 0");
        check("".equals(d2.getLabel()), "empty label");
        check(d2.getData().length == 0, "empty data");

        // UTF-8 string
        String utf8 = "\u4e2d\u6587"; // 中文
        DynamicFields s3 = new DynamicFields();
        s3.setId(1);
        s3.setLabel(utf8);
        s3.setData(new byte[0]);
        byte[] buf3 = s3.encode();
        DynamicFields d3 = new DynamicFields();
        d3.decode(buf3);
        check(utf8.equals(d3.getLabel()), "utf8 label");
    }

    /* ------------------------------------------------------------------ */
    /* Test 9: NarrowBWTest (bitwid)                                       */
    /* ------------------------------------------------------------------ */
    static void testNarrowBW() {
        current = "NarrowBWTest";

        NarrowBWTest s = new NarrowBWTest();
        s.setNarrow12(new short[]{ 2047, -2048, 0, -1 });
        s.setNarrow16(new int[]  { 32767, -32768, 0 });
        s.setNarrow24(new long[] { 8388607L, -8388608L });
        s.setNarrow6 (new byte[] { 63, 0, 32, 1 });

        byte[] buf = s.encode();
        checkEq(buf.length, 21, "encode length");

        NarrowBWTest d = new NarrowBWTest();
        int n = d.decode(buf);
        checkEq(n, 21, "decode length");

        checkEq((int)d.getNarrow12At(0),  2047, "narrow12[0]");
        checkEq((int)d.getNarrow12At(1), -2048, "narrow12[1]");
        checkEq((int)d.getNarrow12At(2),     0, "narrow12[2]");
        checkEq((int)d.getNarrow12At(3),    -1, "narrow12[3]");

        checkEq(d.getNarrow16At(0),  32767L, "narrow16[0]");
        checkEq(d.getNarrow16At(1), -32768L, "narrow16[1]");
        checkEq(d.getNarrow16At(2),      0L, "narrow16[2]");

        checkEq(d.getNarrow24At(0),  8388607L, "narrow24[0]");
        checkEq(d.getNarrow24At(1), -8388608L, "narrow24[1]");

        checkEq(Byte.toUnsignedInt(d.getNarrow6At(0)), 63, "narrow6[0]");
        checkEq(Byte.toUnsignedInt(d.getNarrow6At(1)),  0, "narrow6[1]");
        checkEq(Byte.toUnsignedInt(d.getNarrow6At(2)), 32, "narrow6[2]");
        checkEq(Byte.toUnsignedInt(d.getNarrow6At(3)),  1, "narrow6[3]");

        // Zero round-trip
        NarrowBWTest z = new NarrowBWTest();
        NarrowBWTest dz = new NarrowBWTest();
        dz.decode(z.encode());
        checkEq((int)dz.getNarrow12At(0), 0, "zero narrow12[0]");
        checkEq(Byte.toUnsignedInt(dz.getNarrow6At(3)),  0, "zero narrow6[3]");
    }

    /* ------------------------------------------------------------------ */
    /* Entry point                                                         */
    /* ------------------------------------------------------------------ */
    public static void main(String[] args) {
        System.out.println("=== Bubbler E2E Test - Java ===\n");
        testSimpleScalars();
        testBitPacked();
        testBigEndianFields();
        testArrayFields();
        testEmbedStructs();
        testConstantFields();
        testGetterSetter();
        testDynamicFields();
        testNarrowBW();
        System.out.printf("%n=== Results: %d passed, %d failed ===%n", pass, fail);
        System.exit(fail > 0 ? 1 : 0);
    }
}
