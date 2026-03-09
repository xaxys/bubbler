// Bubbler E2E Test — Go target (bitwid package)
// Tests narrow bit-width array encoding/decoding.
//
// Run (from e2e/tests/go/):  go test ./bitwid/

package bitwid

import (
	"testing"
)

// TestNarrowBW verifies element values are encoded/decoded correctly when
// each element occupies fewer bits than the natural type width.
//
// Field layouts (from bitwid.bb):
//
//	narrow12 : int16[4]  in 6 bytes = 12 bits/elem  [-2048, 2047]
//	narrow16 : int32[3]  in 6 bytes = 16 bits/elem  [-32768, 32767]
//	narrow24 : int64[2]  in 6 bytes = 24 bits/elem  [-8388608, 8388607]
//	narrow6  : uint8[4]  in 3 bytes =  6 bits/elem  [0, 63]
func TestNarrowBW(t *testing.T) {
	s := NarrowBWTest{
		Narrow12: [4]int16{2047, -2048, 0, -1},
		Narrow16: [3]int32{32767, -32768, 0},
		Narrow24: [2]int64{8388607, -8388608},
		Narrow6:  [4]uint8{63, 0, 32, 1},
	}

	if NarrowBWTest(s).Size() != 21 {
		t.Fatalf("Size() = %d, want 21", NarrowBWTest(s).Size())
	}

	buf := s.Encode()
	if len(buf) != 21 {
		t.Fatalf("Encode len = %d, want 21", len(buf))
	}

	var d NarrowBWTest
	n := d.Decode(buf)
	if n != 21 {
		t.Fatalf("Decode n = %d, want 21", n)
	}

	// narrow12 checks
	if d.Narrow12[0] != 2047 {
		t.Errorf("Narrow12[0] = %d, want 2047", d.Narrow12[0])
	}
	if d.Narrow12[1] != -2048 {
		t.Errorf("Narrow12[1] = %d, want -2048", d.Narrow12[1])
	}
	if d.Narrow12[2] != 0 {
		t.Errorf("Narrow12[2] = %d, want 0", d.Narrow12[2])
	}
	if d.Narrow12[3] != -1 {
		t.Errorf("Narrow12[3] = %d, want -1", d.Narrow12[3])
	}

	// narrow16 checks
	if d.Narrow16[0] != 32767 {
		t.Errorf("Narrow16[0] = %d, want 32767", d.Narrow16[0])
	}
	if d.Narrow16[1] != -32768 {
		t.Errorf("Narrow16[1] = %d, want -32768", d.Narrow16[1])
	}
	if d.Narrow16[2] != 0 {
		t.Errorf("Narrow16[2] = %d, want 0", d.Narrow16[2])
	}

	// narrow24 checks
	if d.Narrow24[0] != 8388607 {
		t.Errorf("Narrow24[0] = %d, want 8388607", d.Narrow24[0])
	}
	if d.Narrow24[1] != -8388608 {
		t.Errorf("Narrow24[1] = %d, want -8388608", d.Narrow24[1])
	}

	// narrow6 checks
	if d.Narrow6[0] != 63 {
		t.Errorf("Narrow6[0] = %d, want 63", d.Narrow6[0])
	}
	if d.Narrow6[1] != 0 {
		t.Errorf("Narrow6[1] = %d, want 0", d.Narrow6[1])
	}
	if d.Narrow6[2] != 32 {
		t.Errorf("Narrow6[2] = %d, want 32", d.Narrow6[2])
	}
	if d.Narrow6[3] != 1 {
		t.Errorf("Narrow6[3] = %d, want 1", d.Narrow6[3])
	}
}

// TestNarrowBWZero verifies all-zero round-trip.
func TestNarrowBWZero(t *testing.T) {
	var s NarrowBWTest
	buf := s.Encode()
	var d NarrowBWTest
	if d.Decode(buf) < 0 {
		t.Fatal("Decode zero failed")
	}
	for i, v := range d.Narrow12 {
		if v != 0 {
			t.Errorf("Narrow12[%d] = %d, want 0", i, v)
		}
	}
	for i, v := range d.Narrow6 {
		if v != 0 {
			t.Errorf("Narrow6[%d] = %d, want 0", i, v)
		}
	}
}
