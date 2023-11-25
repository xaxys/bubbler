package dump

import (
	"fmt"

	"github.com/xaxys/bubbler/definition"
)

type DumpGenerator struct{}

func NewDumpGenerator() *DumpGenerator {
	return &DumpGenerator{}
}

func (g *DumpGenerator) Generate(unit *definition.CompilationUnit) (string, error) {
	str := fmt.Sprintf("%v", unit)
	return str, nil
}
