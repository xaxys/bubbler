package definition

import "strings"

var EmptyPackage = &Package{}

type Package struct {
	BasePosition
	PackageFullPaths []string // com.example.app.module.pkgname
	PackagePath      []string // com.example.app.module
	PackageName      string   // pkgname
}

func NewPackage(pos Position, fullPaths []string) *Package {
	if len(fullPaths) == 0 {
		return EmptyPackage
	}
	if pos == nil {
		pos = BasePosition{}
	}
	return &Package{
		BasePosition: BasePosition{
			File:   pos.GetFile(),
			Line:   pos.GetLine(),
			Column: pos.GetColumn(),
		},
		PackageFullPaths: fullPaths,
		PackagePath:      fullPaths[:len(fullPaths)-1],
		PackageName:      fullPaths[len(fullPaths)-1],
	}
}

func (p Package) Equals(other *Package) bool {
	if len(p.PackageFullPaths) != len(other.PackageFullPaths) {
		return false
	}
	for i, name := range p.PackageFullPaths {
		if name != other.PackageFullPaths[i] {
			return false
		}
	}
	return true
}

// Difference returns the diffence of this package and other package name
//
// e.g. com.example.app.module.pkgname and com.example.app.module.pkgname2
// returns ["pkgname2"]
//
// e.g. com.example.app.module.pkgname and com.example.app2
// returns ["..", "..", "app2"]
//
// e.g. com.example.app.module.pkgname and tt.t
// returns ["..", "..", "..", "..", "tt", "t"]
//
// e.g. com.example.app.module.pkgname and com.example.app.module.pkgname.subpkg.abc
// returns ["subpkg", "abc"]
func (p Package) Difference(other *Package) (diff []string) {
	i := 0
	for ; i < len(p.PackagePath) && i < len(other.PackagePath); i++ {
		if p.PackagePath[i] != other.PackagePath[i] {
			break
		}
	}
	// i - 1 same level
	// i current level
	// i + 1 first different level
	for j := i; j < len(p.PackagePath); j++ {
		diff = append(diff, "..")
	}
	for j := i; j < len(other.PackagePath); j++ {
		diff = append(diff, other.PackagePath[j])
	}
	diff = append(diff, other.PackageName)
	return
}

// ToPath returns the path of this package
//
// sep must be a valid path separator (e.g. "/", ".")
//
// ext must start with a dot (e.g. ".h", ".c", ".py")
func (p Package) ToPath(sep, ext string) string {
	return strings.Join(p.PackageFullPaths, sep) + ext
}

// ToFilePath returns the file path of this package
//
// ext must start with a dot (e.g. ".h", ".c", ".py")
func (p Package) ToFilePath(ext string) string {
	return strings.Join(p.PackageFullPaths, "/") + ext
}

// ToRelativePath returns the relative path of this package to other package
//
// ext must start with a dot (e.g. ".h", ".c", ".py")
func (p Package) ToRelativePath(other *Package, ext string) string {
	diff := p.Difference(other)
	return strings.Join(diff, "/") + ext
}

func (p Package) String() string {
	return strings.Join(p.PackageFullPaths, ".")
}
