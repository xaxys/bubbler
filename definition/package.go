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

func (p Package) Equals(other Package) bool {
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

// ToFilePath returns the file path of this package
//
// ext must start with a dot (e.g. ".h", ".c", ".py")
func (p Package) ToFilePath(ext string) string {
	return strings.Join(p.PackageFullPaths, "/") + ext
}

func (p Package) String() string {
	return strings.Join(p.PackageFullPaths, ".")
}
