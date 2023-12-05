package dump

import (
	"fmt"

	"github.com/xaxys/bubbler/generator/genctx"
)

type DumpGenerator struct{}

func NewDumpGenerator() *DumpGenerator {
	return &DumpGenerator{}
}

func (g *DumpGenerator) Generate(ctx *genctx.GenCtx) error {
	unit := ctx.Units[0]

	str := fmt.Sprintf("%v", unit)

	err := ctx.WriteFileAbs(ctx.OutputPath, str)
	if err != nil {
		return err
	}

	return nil
}
