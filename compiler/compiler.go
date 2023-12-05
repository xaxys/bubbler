package compiler

import (
	"github.com/xaxys/bubbler/definition"
)

func Compile(file string) (retUnit []*definition.CompilationUnit, retErr error, retWarning error) {
	rt := NewCompilationRuntime()
	_, err, warning := rt.CompileFile(file)
	return rt.Units.Values(), err, warning
}
