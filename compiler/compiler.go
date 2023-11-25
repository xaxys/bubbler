package compiler

import (
	"github.com/xaxys/bubbler/definition"
)

func Compile(file string) (retUnit *definition.CompilationUnit, retErr error, retWarning error) {
	rt := NewCompilationRuntime()
	return rt.CompileFile(file)
}
