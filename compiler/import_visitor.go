package compiler

import (
	"strconv"

	"github.com/xaxys/bubbler/definition"
	"github.com/xaxys/bubbler/fileio"
	"github.com/xaxys/bubbler/parser"
)

type ImportVisitor struct {
	parser.BasebubblerVisitor
	File    *definition.FileIdentifer
	Imports []*definition.FileIdentifer
	PosList []*definition.BasePosition
}

func NewImportVisitor(File *definition.FileIdentifer) *ImportVisitor {
	return &ImportVisitor{
		File:    File,
		Imports: make([]*definition.FileIdentifer, 0),
		PosList: make([]*definition.BasePosition, 0),
	}
}

func (v *ImportVisitor) VisitProto(ctx *parser.ProtoContext) any {
	importStmts := ctx.AllImportStatement()
	for _, importStmt := range importStmts {
		err := importStmt.Accept(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v *ImportVisitor) VisitImportStatement(ctx *parser.ImportStatementContext) any {
	pos := &definition.BasePosition{
		File:   v.File.Path,
		Line:   ctx.StrLit().STR_LIT().GetSymbol().GetLine(),
		Column: ctx.StrLit().STR_LIT().GetSymbol().GetColumn(),
	}
	topLevelErrorWrapper := func(err error) definition.TopLevelError {
		return &definition.CompileError{
			Position: pos,
			Err:      err,
		}
	}

	fileLit := ctx.StrLit().STR_LIT().GetText()
	file, err := strconv.Unquote(fileLit)
	if err != nil {
		return topLevelErrorWrapper(
			&definition.InvalidLiteralError{
				Literal: fileLit,
				Err:     err,
			},
		)
	}

	path, err := fileio.GetAbsolutelyPathWithBasePath(file, v.File.Path)
	if err != nil {
		return topLevelErrorWrapper(err)
	}

	ident := &definition.FileIdentifer{
		Name: file,
		Path: path,
	}

	v.Imports = append(v.Imports, ident)
	v.PosList = append(v.PosList, pos)
	return nil
}
