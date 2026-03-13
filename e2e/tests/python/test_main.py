"""
Bubbler E2E Test — Python target

Run (from this directory, after code-gen step):
  python -m pytest test_main.py -v
or simply:
  python test_main.py

The run_tests.sh script generates gen/testcase_bb.py + gen/bitwid_bb.py then invokes pytest.
"""
import sys
import math
import unittest
sys.path.insert(0, "gen")

from testcase_bb import (
    Color, Status,
    Point,
    SimpleScalars,
    BitPacked,
    BigEndianFields,
    ArrayFields,
    EmbedStructs,
    ConstantFields,
    GetterSetter,
    DynamicFields,
)
from bitwid_bb import NarrowBWTest


class TestSimpleScalars(unittest.TestCase):
    def test_round_trip(self):
        s = SimpleScalars()
        s.u8_zero = 0
        s.u8_max  = 255
        s.i8_min  = -128
        s.i8_max  = 127
        s.u16_val = 0xBEEF
        s.i16_neg = -1234
        s.u32_val = 0xDEADBEEF
        s.i32_neg = -100000
        s.u64_val = 0xCAFEBABEDEADBEEF
        s.i64_neg = -1
        s.f32_val = 3.14159
        s.f64_val = 2.718281828459045
        s.flag_t  = True
        s.flag_f  = False
        s.color   = Color.BLUE

        buf = s.encode()
        self.assertIsInstance(buf, (bytes, bytearray))

        d = SimpleScalars()
        ok, n = d.decode(buf)
        self.assertTrue(ok, f"decode returned ok={ok}")
        self.assertGreater(n, 0)

        self.assertEqual(d.u8_zero, 0)
        self.assertEqual(d.u8_max,  255)
        self.assertEqual(d.i8_min,  -128)
        self.assertEqual(d.i8_max,  127)
        self.assertEqual(d.u16_val, 0xBEEF)
        self.assertEqual(d.i16_neg, -1234)
        self.assertEqual(d.u32_val, 0xDEADBEEF)
        self.assertEqual(d.i32_neg, -100000)
        self.assertEqual(d.u64_val, 0xCAFEBABEDEADBEEF)
        self.assertEqual(d.i64_neg, -1)
        self.assertAlmostEqual(d.f32_val, 3.14159, delta=1e-3)
        self.assertAlmostEqual(d.f64_val, 2.718281828459045, delta=1e-12)
        self.assertTrue(d.flag_t)
        self.assertFalse(d.flag_f)
        self.assertEqual(d.color, Color.BLUE)


class TestBitPacked(unittest.TestCase):
    def _roundtrip(self, **kwargs):
        s = BitPacked()
        for k, v in kwargs.items():
            setattr(s, k, v)
        buf = s.encode()
        d = BitPacked()
        ok, _ = d.decode(buf)
        self.assertTrue(ok)
        return d

    def test_typical(self):
        d = self._roundtrip(b0=True, b1=False, nibble=0xA, i2=-1,
                            u4=0xF, i20=524287, i48=0x7FFFFFFFFFFF,
                            status=Status.BROKEN)
        self.assertEqual(d.i2, -1)
        self.assertEqual(d.i20, 524287)
        self.assertEqual(d.i48, 0x7FFFFFFFFFFF)
        self.assertEqual(d.status, Status.BROKEN)
        self.assertEqual(d.nibble, 0xA)

    def test_min_signed(self):
        d = self._roundtrip(i2=-2, i20=-524288, i48=-1, status=Status.IDLE)
        self.assertEqual(d.i2, -2)
        self.assertEqual(d.i20, -524288)
        self.assertEqual(d.i48, -1)

    def test_zeros(self):
        d = self._roundtrip(i2=0, i20=0, i48=0, status=Status.IDLE)
        self.assertEqual(d.i2, 0)
        self.assertEqual(d.i20, 0)
        self.assertEqual(d.i48, 0)


class TestBigEndianFields(unittest.TestCase):
    def test_round_trip_and_byte_order(self):
        s = BigEndianFields()
        s.u16   = 0x1234
        s.i32   = -1
        s.f32   = 3.14
        s.arr   = [100, -200, 300, -400]

        buf = s.encode()

        # Constant 0xBE at offset 0
        self.assertEqual(buf[0], 0xBE)
        # u16 big-endian MSB first
        self.assertEqual(buf[1], 0x12)
        self.assertEqual(buf[2], 0x34)

        d = BigEndianFields()
        ok, _ = d.decode(buf)
        self.assertTrue(ok)
        self.assertEqual(d.u16, 0x1234)
        self.assertEqual(d.i32, -1)
        self.assertAlmostEqual(d.f32, 3.14, delta=1e-3)
        self.assertEqual(list(d.arr), [100, -200, 300, -400])


class TestArrayFields(unittest.TestCase):
    def test_round_trip(self):
        s = ArrayFields()
        s.u8_arr    = [0, 127, 128, 255]
        s.i16_arr   = [-32768, 0, 32767]
        s.color_arr = [Color.RED, Color.BLUE]
        s.point_arr = [None, None]
        p0, p1 = Point(), Point()
        p0.x, p0.y = -100, 200
        p1.x, p1.y = 30000, -30000
        s.point_arr = [p0, p1]

        buf = s.encode()
        d = ArrayFields()
        ok, _ = d.decode(buf)
        self.assertTrue(ok)

        self.assertEqual(list(d.u8_arr), [0, 127, 128, 255])
        self.assertEqual(list(d.i16_arr), [-32768, 0, 32767])
        self.assertEqual(list(d.color_arr), [Color.RED, Color.BLUE])
        self.assertEqual(d.point_arr[0].x, -100)
        self.assertEqual(d.point_arr[0].y,  200)
        self.assertEqual(d.point_arr[1].x, 30000)
        self.assertEqual(d.point_arr[1].y, -30000)


class TestEmbedStructs(unittest.TestCase):
    def test_round_trip(self):
        s = EmbedStructs()
        s.id    = 42
        s.ax    = 10
        s.ay    = 20
        pt = Point()
        pt.x, pt.y = -500, 500
        s.pt    = pt
        s.flags = 0xFF

        buf = s.encode()
        d = EmbedStructs()
        ok, _ = d.decode(buf)
        self.assertTrue(ok)

        self.assertEqual(d.id,    42)
        self.assertEqual(d.ax,    10)
        self.assertEqual(d.ay,    20)
        self.assertEqual(d.pt.x, -500)
        self.assertEqual(d.pt.y,  500)
        self.assertEqual(d.flags, 0xFF)


class TestConstantFields(unittest.TestCase):
    def test_encode_constants(self):
        s = ConstantFields()
        s.length = 1024
        buf = s.encode()
        self.assertEqual(buf[0], 0xAA)
        self.assertEqual(buf[1], 0x02)
        self.assertEqual(buf[4], 0x02)

    def test_decode_success(self):
        s = ConstantFields()
        s.length = 1024
        buf = s.encode()
        d = ConstantFields()
        ok, _ = d.decode(buf)
        self.assertTrue(ok)
        self.assertEqual(d.length, 1024)

    def test_decode_fail_bad_header(self):
        s = ConstantFields()
        s.length = 0
        buf = bytearray(s.encode())
        buf[0] = 0xFF
        d = ConstantFields()
        ok, _ = d.decode(bytes(buf))
        self.assertFalse(ok)

    def test_decode_fail_bad_version(self):
        s = ConstantFields()
        s.length = 0
        buf = bytearray(s.encode())
        buf[1] = 0x00
        d = ConstantFields()
        ok, _ = d.decode(bytes(buf))
        self.assertFalse(ok)


class TestGetterSetter(unittest.TestCase):
    def test_voltage(self):
        s = GetterSetter()
        s.voltage = 3.3
        self.assertAlmostEqual(s.voltage, 3.3, delta=5e-3)

        s.voltage = 0.0
        self.assertAlmostEqual(s.voltage, 0.0, delta=1e-6)

    def test_celsius(self):
        s = GetterSetter()
        s.celsius = 36.5
        self.assertAlmostEqual(s.celsius, 36.5, delta=1e-2)

        s.celsius = -40.0
        self.assertAlmostEqual(s.celsius, -40.0, delta=1e-2)

    def test_round_trip(self):
        s = GetterSetter()
        s.voltage = 1.65
        s.celsius = 25.0

        buf = s.encode()
        d = GetterSetter()
        ok, _ = d.decode(buf)
        self.assertTrue(ok)
        self.assertAlmostEqual(d.voltage, 1.65, delta=5e-3)
        self.assertAlmostEqual(d.celsius, 25.0, delta=0.1)


class TestDynamicFields(unittest.TestCase):
    def test_non_empty(self):
        blob = bytes([0xDE, 0xAD, 0xBE, 0xEF, 0x00, 0x01, 0x02])
        s = DynamicFields()
        s.id    = 0xDEADBEEF
        s.label = "hello, bubbler!"
        s.data  = blob

        buf = s.encode()
        d = DynamicFields()
        ok, n = d.decode(buf)
        self.assertTrue(ok)
        self.assertGreater(n, 0)
        self.assertEqual(d.id, 0xDEADBEEF)
        self.assertEqual(d.label, "hello, bubbler!")
        self.assertEqual(bytes(d.data), blob)

    def test_empty(self):
        s = DynamicFields()
        s.id    = 0
        s.label = ""
        s.data  = b""

        buf = s.encode()
        d = DynamicFields()
        ok, _ = d.decode(buf)
        self.assertTrue(ok)
        self.assertEqual(d.id, 0)
        self.assertEqual(d.label, "")
        self.assertEqual(len(d.data), 0)

    def test_utf8(self):
        utf8str = "\u4e2d\u6587"  # 中文
        s = DynamicFields()
        s.id    = 1
        s.label = utf8str
        s.data  = b""

        buf = s.encode()
        d = DynamicFields()
        ok, _ = d.decode(buf)
        self.assertTrue(ok)
        self.assertEqual(d.label, utf8str)

    def test_decode_size_boundaries(self):
        blob = bytes([0xDE, 0xAD, 0xBE, 0xEF, 0x00, 0x01, 0x02])
        s = DynamicFields()
        s.id = 7
        s.label = "probe"
        s.data = blob
        full = s.encode()

        d = DynamicFields()
        self.assertEqual(d.decode_size(full), len(full), "decode_size complete")
        self.assertEqual(d.decode_size(full[:-1]), -len(full), "decode_size truncated 1 byte")

        self.assertEqual(d.decode_size(bytes([1, 0, 0, 0, ord('A')])), -6, "missing string terminator")
        self.assertEqual(d.decode_size(bytes([1, 0, 0, 0, ord('A'), 0, 0x80])), -8, "truncated bytes varint")
        self.assertEqual(d.decode_size(bytes([1, 0, 0, 0, ord('A'), 0, 0x03, 0xAA, 0xBB])), -10, "truncated bytes payload")
        self.assertEqual(d.decode_size(bytes([1, 0, 0, 0])), -5, "only fixed header")


class TestNarrowBWTest(unittest.TestCase):
    """Narrow bit-width array encoding/decoding (bitwid.bb)

    Field layouts:
      narrow12 : int16[4]  in 6 bytes = 12 bits/elem  [-2048, 2047]
      narrow16 : int32[3]  in 6 bytes = 16 bits/elem  [-32768, 32767]
      narrow24 : int64[2]  in 6 bytes = 24 bits/elem  [-8388608, 8388607]
      narrow6  : uint8[4]  in 3 bytes =  6 bits/elem  [0, 63]
    """

    def _roundtrip(self, s):
        buf = s.encode()
        d = NarrowBWTest()
        ok, _ = d.decode(buf)
        self.assertTrue(ok, "decode returned False")
        return d

    def test_typical_values(self):
        s = NarrowBWTest()
        s.narrow12 = [2047, -2048, 0, -1]
        s.narrow16 = [32767, -32768, 0]
        s.narrow24 = [8388607, -8388608]
        s.narrow6  = [63, 0, 32, 1]

        buf = s.encode()
        self.assertEqual(len(buf), 21, "encode length should be 21 bytes")

        d = self._roundtrip(s)

        self.assertEqual(d.narrow12[0],  2047)
        self.assertEqual(d.narrow12[1], -2048)
        self.assertEqual(d.narrow12[2],     0)
        self.assertEqual(d.narrow12[3],    -1)

        self.assertEqual(d.narrow16[0],  32767)
        self.assertEqual(d.narrow16[1], -32768)
        self.assertEqual(d.narrow16[2],      0)

        self.assertEqual(d.narrow24[0],  8388607)
        self.assertEqual(d.narrow24[1], -8388608)

        self.assertEqual(d.narrow6[0], 63)
        self.assertEqual(d.narrow6[1],  0)
        self.assertEqual(d.narrow6[2], 32)
        self.assertEqual(d.narrow6[3],  1)

    def test_zero_roundtrip(self):
        s = NarrowBWTest()
        d = self._roundtrip(s)
        self.assertEqual(list(d.narrow12), [0, 0, 0, 0])
        self.assertEqual(list(d.narrow16), [0, 0, 0])
        self.assertEqual(list(d.narrow24), [0, 0])
        self.assertEqual(list(d.narrow6),  [0, 0, 0, 0])


if __name__ == "__main__":
    result = unittest.main(verbosity=2, exit=False)
    sys.exit(0 if result.result.wasSuccessful() else 1)
