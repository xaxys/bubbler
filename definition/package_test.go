package definition

import "testing"

func TestPackage_ToFilePath(t *testing.T) {
	p := NewPackage(nil, []string{"com", "example", "app", "module", "pkgname"})
	ext := ".go"
	expected := "com/example/app/module/pkgname.go"

	filePath := p.ToFilePath(ext)

	if filePath != expected {
		t.Errorf("File path is %s, expected %s", filePath, expected)
	}
}

func TestPackage_ToFilePath_WithDotExt(t *testing.T) {
	p := NewPackage(nil, []string{"com", "example", "app", "module", "pkgname"})
	ext := ".py"
	expected := "com/example/app/module/pkgname.py"

	filePath := p.ToFilePath(ext)

	if filePath != expected {
		t.Errorf("File path is %s, expected %s", filePath, expected)
	}
}

func TestPackage_Equals(t *testing.T) {
	p1 := NewPackage(nil, []string{"a", "b", "c", "d"})
	p2 := NewPackage(nil, []string{"a", "b", "c", "d"})
	p3 := NewPackage(nil, []string{"a", "b", "c", "e"})

	if !p1.PackageEquals(p2) {
		t.Errorf("Expected p1 to be equal to p2")
	}

	if p1.PackageEquals(p3) {
		t.Errorf("Expected p1 to be not equal to p3")
	}
}

func TestPackage_String(t *testing.T) {
	p := NewPackage(nil, []string{"com", "example", "app", "module", "pkgname"})
	expected := "com.example.app.module.pkgname"

	str := p.String()

	if str != expected {
		t.Errorf("String representation is %s, expected %s", str, expected)
	}
}

func TestPackage_ToRelativePath(t *testing.T) {
	p := NewPackage(nil, []string{"com", "example", "app", "module", "pkgname"})
	other := NewPackage(nil, []string{"com", "example", "app", "module"})
	ext := ".go"
	expected := "../module.go"

	relativePath := p.ToRelativePath(other, ext)

	if relativePath != expected {
		t.Errorf("Relative path is %s, expected %s", relativePath, expected)
	}
}

func TestPackage_ToRelativePath_WithDifferentExtension(t *testing.T) {
	p := NewPackage(nil, []string{"com", "example", "app", "module", "pkgname"})
	other := NewPackage(nil, []string{"com", "example", "app", "module"})
	ext := ".py"
	expected := "../module.py"

	relativePath := p.ToRelativePath(other, ext)

	if relativePath != expected {
		t.Errorf("Relative path is %s, expected %s", relativePath, expected)
	}
}

func TestPackage_Difference(t *testing.T) {
	p := NewPackage(nil, []string{"a", "b", "c", "d"})
	tests := []struct {
		other    *Package
		expected []string
	}{
		{
			other:    NewPackage(nil, []string{"a", "b", "c", "d"}),
			expected: []string{"d"},
		},
		{
			other:    NewPackage(nil, []string{"a", "b", "c", "d", "e"}),
			expected: []string{"d", "e"},
		},
		{
			other:    NewPackage(nil, []string{"a", "b", "c"}),
			expected: []string{"..", "c"},
		},
		{
			other:    NewPackage(nil, []string{"a", "b", "e", "f", "g"}),
			expected: []string{"..", "e", "f", "g"},
		},
		{
			other:    NewPackage(nil, []string{"a", "b", "c", "d", "e", "f"}),
			expected: []string{"d", "e", "f"},
		},
		{
			other:    NewPackage(nil, []string{"e", "f", "g", "h"}),
			expected: []string{"..", "..", "..", "e", "f", "g", "h"},
		},
	}

	for _, test := range tests {
		diff := p.Difference(test.other)

		if len(diff) != len(test.expected) {
			t.Errorf("diff is %v, expected %v", diff, test.expected)
			continue
		}

		for i := 0; i < len(diff); i++ {
			if diff[i] != test.expected[i] {
				t.Errorf("diff is %v, expected %v", diff, test.expected)
				break
			}
		}
	}
}
