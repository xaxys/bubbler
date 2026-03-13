// Bubbler E2E Test — Go target
//
// The generated testpkg.bb.go is placed alongside this file (in e2e/tests/go/testpkg/).
// Run:  go test ./testpkg/  (from e2e/tests/go/)
package testpkg

import (
	"bytes"
	"encoding/binary"
	"math"
	"testing"
)

/* ------------------------------------------------------------------ */
/* Helper                                                               */
/* ------------------------------------------------------------------ */

func nearF64(a, b, eps float64) bool {
	return math.Abs(a-b) < eps
}

/* ------------------------------------------------------------------ */
/* Test 1: SimpleScalars                                               */
/* ------------------------------------------------------------------ */
func TestSimpleScalars(t *testing.T) {
	s := SimpleScalars{
		U8Zero: 0,
		U8Max:  255,
		I8Min:  -128,
		I8Max:  127,
		U16Val: 0xBEEF,
		I16Neg: -1234,
		U32Val: 0xDEADBEEF,
		I32Neg: -100000,
		U64Val: 0xCAFEBABEDEADBEEF,
		I64Neg: -1,
		F32Val: 3.14159,
		F64Val: 2.718281828459045,
		FlagT:  true,
		FlagF:  false,
		Color:  BLUE,
	}

	buf := s.Encode()
	if len(buf) != int(SimpleScalars{}.Size()) {
		t.Fatalf("encode length: got %d, want %d", len(buf), SimpleScalars{}.Size())
	}

	var d SimpleScalars
	n := d.Decode(buf)
	if n < 0 {
		t.Fatalf("Decode returned %d", n)
	}

	if d.U8Zero != 0 {
		t.Errorf("U8Zero: got %d", d.U8Zero)
	}
	if d.U8Max != 255 {
		t.Errorf("U8Max: got %d", d.U8Max)
	}
	if d.I8Min != -128 {
		t.Errorf("I8Min: got %d", d.I8Min)
	}
	if d.I8Max != 127 {
		t.Errorf("I8Max: got %d", d.I8Max)
	}
	if d.U16Val != 0xBEEF {
		t.Errorf("U16Val: got 0x%X", d.U16Val)
	}
	if d.I16Neg != -1234 {
		t.Errorf("I16Neg: got %d", d.I16Neg)
	}
	if d.U32Val != 0xDEADBEEF {
		t.Errorf("U32Val: got 0x%X", d.U32Val)
	}
	if d.I32Neg != -100000 {
		t.Errorf("I32Neg: got %d", d.I32Neg)
	}
	if d.U64Val != 0xCAFEBABEDEADBEEF {
		t.Errorf("U64Val: got 0x%X", d.U64Val)
	}
	if d.I64Neg != -1 {
		t.Errorf("I64Neg: got %d", d.I64Neg)
	}
	if !nearF64(float64(d.F32Val), 3.14159, 1e-3) {
		t.Errorf("F32Val: got %v", d.F32Val)
	}
	if !nearF64(d.F64Val, 2.718281828459045, 1e-12) {
		t.Errorf("F64Val: got %v", d.F64Val)
	}
	if d.FlagT != true {
		t.Errorf("FlagT: got false")
	}
	if d.FlagF != false {
		t.Errorf("FlagF: got true")
	}
	if d.Color != BLUE {
		t.Errorf("Color: got %v", d.Color)
	}
}

/* ------------------------------------------------------------------ */
/* Test 2: BitPacked — sign extension, bit boundaries                  */
/* ------------------------------------------------------------------ */
func TestBitPacked(t *testing.T) {
	cases := []struct {
		name   string
		b0     bool
		nibble uint8
		i2     int8
		u4     uint8
		i20    int32
		i48    int64
		status Status
	}{
		{"typical", true, 0xA, -1, 0xF, 524287, 0x7FFFFFFFFFFF, BROKEN},
		{"min_signed", false, 0, -2, 0, -524288, -1, IDLE},
		{"zeros", false, 0, 0, 0, 0, 0, IDLE},
		{"max_i48", true, 0xF, 1, 0xF, 524287, 0x7FFFFFFFFFFF, ACTIVE},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := BitPacked{
				B0:     c.b0,
				B1:     false,
				Nibble: c.nibble,
				I2:     c.i2,
				U4:     c.u4,
				I20:    c.i20,
				I48:    c.i48,
				Status: c.status,
			}
			buf := s.Encode()
			var d BitPacked
			n := d.Decode(buf)
			if n < 0 {
				t.Fatalf("Decode returned %d", n)
			}
			if d.I2 != c.i2 {
				t.Errorf("I2: got %d, want %d", d.I2, c.i2)
			}
			if d.I20 != c.i20 {
				t.Errorf("I20: got %d, want %d", d.I20, c.i20)
			}
			if d.I48 != c.i48 {
				t.Errorf("I48: got %d, want %d", d.I48, c.i48)
			}
			if d.Status != c.status {
				t.Errorf("Status: got %v, want %v", d.Status, c.status)
			}
			if d.Nibble != c.nibble {
				t.Errorf("Nibble: got %d, want %d", d.Nibble, c.nibble)
			}
		})
	}
}

/* ------------------------------------------------------------------ */
/* Test 3: BigEndianFields — byte-order verification                   */
/* ------------------------------------------------------------------ */
func TestBigEndianFields(t *testing.T) {
	s := BigEndianFields{
		U16: 0x1234,
		I32: -1,
		F32: 3.14,
		Arr: [4]int16{100, -200, 300, -400},
	}
	buf := s.Encode()

	// Constant 0xBE at byte offset 0
	if buf[0] != 0xBE {
		t.Errorf("buf[0]: got 0x%X, want 0xBE", buf[0])
	}
	// u16=0x1234 big-endian
	if buf[1] != 0x12 || buf[2] != 0x34 {
		t.Errorf("u16 big-endian wrong: 0x%02X 0x%02X", buf[1], buf[2])
	}

	var d BigEndianFields
	if n := d.Decode(buf); n < 0 {
		t.Fatalf("Decode returned %d", n)
	}
	if d.U16 != 0x1234 {
		t.Errorf("U16: got 0x%X", d.U16)
	}
	if d.I32 != -1 {
		t.Errorf("I32: got %d", d.I32)
	}
	if !nearF64(float64(d.F32), 3.14, 1e-3) {
		t.Errorf("F32: got %v", d.F32)
	}
	expected := [4]int16{100, -200, 300, -400}
	if d.Arr != expected {
		t.Errorf("Arr: got %v, want %v", d.Arr, expected)
	}
}

/* ------------------------------------------------------------------ */
/* Test 4: ArrayFields                                                 */
/* ------------------------------------------------------------------ */
func TestArrayFields(t *testing.T) {
	s := ArrayFields{
		U8Arr:    [4]uint8{0, 127, 128, 255},
		I16Arr:   [3]int16{-32768, 0, 32767},
		ColorArr: [2]Color{RED, BLUE},
		PointArr: [2]Point{
			{X: -100, Y: 200},
			{X: 30000, Y: -30000},
		},
	}
	buf := s.Encode()
	var d ArrayFields
	if n := d.Decode(buf); n < 0 {
		t.Fatalf("Decode returned %d", n)
	}
	if d.U8Arr != s.U8Arr {
		t.Errorf("U8Arr: got %v", d.U8Arr)
	}
	if d.I16Arr != s.I16Arr {
		t.Errorf("I16Arr: got %v", d.I16Arr)
	}
	if d.ColorArr != s.ColorArr {
		t.Errorf("ColorArr: got %v", d.ColorArr)
	}
	if d.PointArr != s.PointArr {
		t.Errorf("PointArr: got %v", d.PointArr)
	}
}

/* ------------------------------------------------------------------ */
/* Test 5: EmbedStructs                                               */
/* ------------------------------------------------------------------ */
func TestEmbedStructs(t *testing.T) {
	s := EmbedStructs{
		Id:    42,
		Ax:    10,
		Ay:    20,
		Pt:    Point{X: -500, Y: 500},
		Flags: 0xFF,
	}
	buf := s.Encode()
	var d EmbedStructs
	if n := d.Decode(buf); n < 0 {
		t.Fatalf("Decode returned %d", n)
	}
	if d.Id != 42 || d.Ax != 10 || d.Ay != 20 {
		t.Errorf("scalar fields mismatch: %+v", d)
	}
	if d.Pt.X != -500 || d.Pt.Y != 500 {
		t.Errorf("Pt mismatch: %+v", d.Pt)
	}
	if d.Flags != 0xFF {
		t.Errorf("Flags: got 0x%X", d.Flags)
	}
}

/* ------------------------------------------------------------------ */
/* Test 6: ConstantFields — encode constants; decode failure on mismatch */
/* ------------------------------------------------------------------ */
func TestConstantFields(t *testing.T) {
	s := ConstantFields{Length: 1024}
	buf := s.Encode()

	if buf[0] != 0xAA {
		t.Errorf("header: got 0x%X, want 0xAA", buf[0])
	}
	if buf[1] != 0x02 {
		t.Errorf("version: got 0x%X, want 0x02", buf[1])
	}
	if buf[4] != 0x02 {
		t.Errorf("magic_color: got 0x%X, want 0x02", buf[4])
	}

	var d ConstantFields
	if n := d.Decode(buf); n < 0 {
		t.Fatalf("happy-path Decode failed: %d", n)
	}
	if d.Length != 1024 {
		t.Errorf("Length: got %d", d.Length)
	}

	// Corrupt header
	bad := append([]byte{}, buf...)
	bad[0] = 0xFF
	if n := d.Decode(bad); n >= 0 {
		t.Errorf("Decode should fail on bad header, got %d", n)
	}

	// Corrupt version
	bad2 := append([]byte{}, buf...)
	bad2[1] = 0x00
	if n := d.Decode(bad2); n >= 0 {
		t.Errorf("Decode should fail on bad version, got %d", n)
	}
}

/* ------------------------------------------------------------------ */
/* Test 7: GetterSetter — custom getter/setter round-trip              */
/* ------------------------------------------------------------------ */
func TestGetterSetter(t *testing.T) {
	var s GetterSetter

	s.SetVoltage(3.3)
	if !nearF64(s.GetVoltage(), 3.3, 5e-3) {
		t.Errorf("GetVoltage: got %v", s.GetVoltage())
	}

	s.SetVoltage(0.0)
	if !nearF64(s.GetVoltage(), 0.0, 1e-6) {
		t.Errorf("GetVoltage 0: got %v", s.GetVoltage())
	}

	s.SetCelsius(36.5)
	if !nearF64(s.GetCelsius(), 36.5, 1e-2) {
		t.Errorf("GetCelsius: got %v", s.GetCelsius())
	}

	s.SetCelsius(-40.0)
	if !nearF64(s.GetCelsius(), -40.0, 1e-2) {
		t.Errorf("GetCelsius -40: got %v", s.GetCelsius())
	}

	// Encode/decode preserves raw values
	var s2 GetterSetter
	s2.SetVoltage(1.65)
	s2.SetCelsius(25.0)
	buf := s2.Encode()
	var d GetterSetter
	if n := d.Decode(buf); n < 0 {
		t.Fatalf("Decode failed: %d", n)
	}
	if !nearF64(d.GetVoltage(), 1.65, 5e-3) {
		t.Errorf("round-trip voltage: got %v", d.GetVoltage())
	}
	if !nearF64(d.GetCelsius(), 25.0, 0.1) {
		t.Errorf("round-trip celsius: got %v", d.GetCelsius())
	}
}

/* ------------------------------------------------------------------ */
/* Test 8: DynamicFields — variable-length string + bytes             */
/* ------------------------------------------------------------------ */
func TestDynamicFields(t *testing.T) {
	t.Run("non_empty", func(t *testing.T) {
		blob := []byte{0xDE, 0xAD, 0xBE, 0xEF, 0x00, 0x01, 0x02}
		s := DynamicFields{
			Id:    0xDEADBEEF,
			Label: "hello, bubbler!",
			Data:  blob,
		}
		buf := s.Encode()
		var d DynamicFields
		if n := d.Decode(buf); n < 0 {
			t.Fatalf("Decode returned %d", n)
		}
		if d.Id != 0xDEADBEEF {
			t.Errorf("Id: got 0x%X", d.Id)
		}
		if d.Label != "hello, bubbler!" {
			t.Errorf("Label: got %q", d.Label)
		}
		if !bytes.Equal(d.Data, blob) {
			t.Errorf("Data: got %v", d.Data)
		}
	})

	t.Run("empty", func(t *testing.T) {
		s := DynamicFields{Id: 0, Label: "", Data: nil}
		buf := s.Encode()
		var d DynamicFields
		if n := d.Decode(buf); n < 0 {
			t.Fatalf("Decode returned %d", n)
		}
		if d.Label != "" {
			t.Errorf("Label should be empty, got %q", d.Label)
		}
		if len(d.Data) != 0 {
			t.Errorf("Data should be empty, got %v", d.Data)
		}
	})

	t.Run("utf8", func(t *testing.T) {
		utf8str := "\xe4\xb8\xad\xe6\x96\x87" // 中文
		s := DynamicFields{Id: 1, Label: utf8str}
		buf := s.Encode()
		var d DynamicFields
		d.Decode(buf)
		if d.Label != utf8str {
			t.Errorf("UTF-8 label mismatch: got %q", d.Label)
		}
	})
}

func TestDynamicFieldsDecodeSize(t *testing.T) {
	buildPacket := func(label string, payload []byte) []byte {
		s := DynamicFields{Id: 1, Label: label, Data: payload}
		return s.Encode()
	}

	t.Run("complete packet", func(t *testing.T) {
		buf := buildPacket("hello", []byte{0xAA, 0xBB, 0xCC})
		var d DynamicFields
		if got := d.DecodeSize(buf); got != len(buf) {
			t.Fatalf("DecodeSize complete: got %d, want %d", got, len(buf))
		}
	})

	t.Run("truncate one byte", func(t *testing.T) {
		buf := buildPacket("hello", []byte{0xAA, 0xBB, 0xCC})
		var d DynamicFields
		if got := d.DecodeSize(buf[:len(buf)-1]); got != -len(buf) {
			t.Fatalf("DecodeSize truncated: got %d, want %d", got, -len(buf))
		}
	})

	t.Run("string without terminator", func(t *testing.T) {
		malformed := []byte{1, 0, 0, 0, 'A'}
		var d DynamicFields
		if got := d.DecodeSize(malformed); got != -6 {
			t.Fatalf("DecodeSize missing string terminator: got %d, want -6", got)
		}
	})

	t.Run("bytes varint length truncated", func(t *testing.T) {
		malformed := []byte{1, 0, 0, 0, 'A', 0, 0x80}
		var d DynamicFields
		if got := d.DecodeSize(malformed); got != -8 {
			t.Fatalf("DecodeSize truncated varint: got %d, want -8", got)
		}
	})

	t.Run("bytes payload truncated", func(t *testing.T) {
		malformed := []byte{1, 0, 0, 0, 'A', 0, 0x03, 0xAA, 0xBB}
		var d DynamicFields
		if got := d.DecodeSize(malformed); got != -10 {
			t.Fatalf("DecodeSize truncated payload: got %d, want -10", got)
		}
	})

	t.Run("empty dynamic fields", func(t *testing.T) {
		buf := buildPacket("", nil)
		var d DynamicFields
		if got := d.DecodeSize(buf); got != len(buf) {
			t.Fatalf("DecodeSize empty dynamic: got %d, want %d", got, len(buf))
		}
	})

	t.Run("multi-byte varint payload", func(t *testing.T) {
		payload := make([]byte, 130)
		for i := range payload {
			payload[i] = byte(i)
		}
		buf := buildPacket("a", payload)
		var d DynamicFields
		if got := d.DecodeSize(buf); got != len(buf) {
			t.Fatalf("DecodeSize multi-byte varint: got %d, want %d", got, len(buf))
		}

		// Build a deterministic malformed frame: id + "a\0" + varint(130) + only 129 payload bytes.
		malformed := make([]byte, 0, 4+2+2+129)
		id := make([]byte, 4)
		binary.LittleEndian.PutUint32(id, 1)
		malformed = append(malformed, id...)
		malformed = append(malformed, 'a', 0)
		malformed = append(malformed, 0x82, 0x01)
		malformed = append(malformed, payload[:129]...)
		if got := d.DecodeSize(malformed); got != -138 {
			t.Fatalf("DecodeSize multi-byte varint truncated payload: got %d, want -138", got)
		}
	})
}
