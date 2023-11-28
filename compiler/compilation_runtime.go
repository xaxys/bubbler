package compiler

import (
	"errors"

	"github.com/antlr4-go/antlr/v4"
	"github.com/xaxys/bubbler/definition"
	"github.com/xaxys/bubbler/fileio"
	"github.com/xaxys/bubbler/parser"
	"github.com/xaxys/bubbler/util"
)

type CompilationRuntime struct {
	Units      *util.OrderedMap[string, *definition.CompilationUnit]
	ParseStack *util.OrderedMap[string, *definition.CompilationUnit]
}

func NewCompilationRuntime() *CompilationRuntime {
	return &CompilationRuntime{
		Units:      util.NewOrderedMap[string, *definition.CompilationUnit](),
		ParseStack: util.NewOrderedMap[string, *definition.CompilationUnit](),
	}
}

func (r *CompilationRuntime) CompileFile(file string) (retUnit *definition.CompilationUnit, retErr error, retWarning error) {
	ident := &definition.FileIdentifer{
		Name: file,
		Path: "",
	}
	return r.Compile(ident)
}

func (r *CompilationRuntime) Compile(file *definition.FileIdentifer) (retUnit *definition.CompilationUnit, retErr error, retWarning error) {
	if file == nil {
		return nil, nil, nil
	}
	if file.Name == "" && file.Path == "" {
		return nil, nil, nil
	}

	var ident *definition.FileIdentifer
	if file.Name == "" {
		ident = &definition.FileIdentifer{
			Name: file.Path,
			Path: file.Path,
		}
	} else if file.Path == "" {
		id, err := fileio.GetFileIdentifer(file.Name)
		if err != nil {
			return nil, err, nil
		}
		ident = id
	} else {
		ident = &definition.FileIdentifer{
			Name: file.Name,
			Path: file.Path,
		}
	}

	err := fileio.GetFileExistingStatus(ident)
	if err != nil {
		return nil, err, nil
	}

	content, err := fileio.GetFileContent(ident)
	if err != nil {
		return nil, err, nil
	}

	return r.compile(ident, content)
}

func (r *CompilationRuntime) compile(ident *definition.FileIdentifer, content string) (retUnit *definition.CompilationUnit, retErr error, retWarning error) {
	if ident == nil {
		return nil, nil, nil
	}

	skip, err := r.CheckImport(ident)
	if err != nil {
		return nil, err, nil
	}
	if skip {
		return nil, nil, nil
	}

	ast, err := getProtoAST(content, ident)
	if err != nil {
		return nil, err, nil
	}

	importor := NewImportVisitor(ident)
	importErr := ast.Accept(importor)
	if err, ok := importErr.(error); ok {
		return nil, err, nil
	}

	unit := definition.NewCompilationUnit(ident)
	r.Units.Put(ident.Path, unit)
	r.ParseStack.Put(ident.Path, unit)

	var warnings error
	// compile imports
	for i, other := range importor.Imports {
		otherUnit, err, warning := r.Compile(other)
		if otherUnit == nil && err == nil && warning == nil {
			continue
		}
		if warning != nil {
			w := &definition.CompileWarning{
				Position: importor.PosList[i],
				Warning: &definition.ImportingWarning{
					File:    other,
					Warning: warning,
				},
			}
			warnings = errors.Join(warnings, w)
		}
		if err != nil {
			topLevelErr, ok := err.(definition.TopLevelError)
			if !ok {
				return nil, &definition.CompileError{
					Position: importor.PosList[i],
					Err:      err,
				}, warnings
			}
			return nil, &definition.CompileError{
				Position: importor.PosList[i],
				Err: &definition.ImportingError{
					File: other,
					Err:  topLevelErr,
				},
			}, warnings
		}
		err = unit.AddImport(otherUnit)
		if err != nil {
			return nil, &definition.CompileError{
				Position: importor.PosList[i],
				Err:      err,
			}, warnings
		}
	}

	r.ParseStack.Remove(ident.Path)

	protoVisitor := NewParseVisitor(unit)
	protoErr := ast.Accept(protoVisitor)
	warnings = errors.Join(warnings, protoVisitor.Warning)
	if err, ok := protoErr.(error); ok {
		return nil, err, warnings
	}

	return unit, nil, warnings
}

func (r *CompilationRuntime) CheckImport(file *definition.FileIdentifer) (skip bool, err error) {
	if r.ParseStack.Has(file.Path) {
		if r.ParseStack.Last().Value.UnitName.Path == file.Path {
			return true, &definition.ImportSelfError{}
		}
		return true, &definition.ImportCycleError{File: file}
	}
	if r.Units.Has(file.Path) {
		// already imported
		// skip by diamond import
		return true, nil
	}
	return false, nil
}

func getProtoAST(content string, file *definition.FileIdentifer) (antlr.ParseTree, error) {
	listener := &errorListener{File: file.Path}

	inputStream := antlr.NewInputStream(content)
	lexer := parser.NewbubblerLexer(inputStream)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(listener)

	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	parsar := parser.NewbubblerParser(tokenStream)
	parsar.RemoveErrorListeners()
	parsar.AddErrorListener(listener)

	ast := parsar.Proto()
	if listener.Err != nil {
		return nil, listener.Err
	}
	return ast, nil
}
