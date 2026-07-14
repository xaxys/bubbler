package gen

import "testing"

func TestNewGenOptionsDefaults(t *testing.T) {
    options := NewGenOptions()
    if options.LoopUnroll != 4 {
        t.Fatalf("LoopUnroll default = %d, want 4", options.LoopUnroll)
    }
}

func TestLoopUnrollBoundaries(t *testing.T) {
    for _, value := range []int{-1, 0, 1, 4, 32} {
        setter, err := LoopUnroll(value)
        if err != nil {
            t.Fatalf("LoopUnroll(%d): %v", value, err)
        }
        options := NewGenOptions(setter)
        if options.LoopUnroll != value {
            t.Fatalf("LoopUnroll(%d) stored %d", value, options.LoopUnroll)
        }
    }
}

func TestLoopUnrollRejectsLessThanNegativeOne(t *testing.T) {
    if _, err := LoopUnroll(-2); err == nil {
        t.Fatal("LoopUnroll(-2) succeeded, want error")
    }
}

func TestRemovePathTrimsEveryTrailingSlash(t *testing.T) {
    options := NewGenOptions(RemovePath("example.com/a///,example.com/b/"))
    want := []string{"example.com/a", "example.com/b"}
    if len(options.RemovePath) != len(want) {
        t.Fatalf("RemovePath length = %d, want %d", len(options.RemovePath), len(want))
    }
    for i := range want {
        if options.RemovePath[i] != want[i] {
            t.Fatalf("RemovePath[%d] = %q, want %q", i, options.RemovePath[i], want[i])
        }
    }
}
