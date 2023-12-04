package compiler

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/xaxys/bubbler/definition"
	"github.com/xaxys/bubbler/fileio"
	"github.com/xaxys/bubbler/parser"
	"github.com/xaxys/bubbler/util"
)

type CompilationRuntime struct {
	Units      *util.OrderedMap[string, *definition.CompilationUnit] // Units contains all the units compiled, full path as key
	Packages   *util.OrderedMap[string, *definition.CompilationUnit] // another view of Units, for faster package lookup, package name as key
	ParseStack *util.OrderedMap[string, *definition.CompilationUnit] // ParseStack contains all the units being parsed, full path as key
}

func NewCompilationRuntime() *CompilationRuntime {
	return &CompilationRuntime{
		Units:      util.NewOrderedMap[string, *definition.CompilationUnit](),
		Packages:   util.NewOrderedMap[string, *definition.CompilationUnit](),
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

	skip, err := r.precheckFilePath(ident)
	if err != nil {
		return nil, err, nil
	}
	if skip {
		return r.Units.MustGet(ident.Path), nil, nil
	}

	ast, err := getProtoAST(content, ident)
	if err != nil {
		return nil, err, nil
	}

	var warnings definition.TopLevelWarning

	importVisitor := NewImportVisitor(ident)
	importErr := ast.Accept(importVisitor)
	if err, ok := importErr.(error); ok {
		return nil, err, nil
	}

	unit := definition.NewCompilationUnit(ident)

	infoVisitor := NewInfoVisitor(unit)
	infoErr := ast.Accept(infoVisitor)
	warnings = definition.TopLevelWarningsJoin(warnings, infoVisitor.Warning)
	if err, ok := infoErr.(error); ok {
		return nil, err, warnings
	}

	err = r.precheckPackageName(unit.Package)
	if err != nil {
		return nil, err, warnings
	}

	r.Units.Put(ident.Path, unit)
	r.Packages.Put(unit.Package.String(), unit)
	r.ParseStack.Put(ident.Path, unit)

	// compile imports
	var errs definition.TopLevelError
	for i, other := range importVisitor.Imports {
		otherUnit, err, warning := r.Compile(other)
		if warning != nil {
			var w definition.TopLevelWarning
			topLevelWarning, ok := warning.(definition.TopLevelWarning)
			if !ok {
				w = &definition.CompileWarning{
					Position: importVisitor.PosList[i],
					Warning:  warning,
				}
			} else {
				w = &definition.CompileWarning{
					Position: importVisitor.PosList[i],
					Warning: &definition.ImportingWarning{
						File:    other,
						Warning: topLevelWarning,
					},
				}
			}
			warnings = definition.TopLevelWarningsJoin(warnings, w)
		}
		if err != nil {
			topLevelErr, ok := err.(definition.TopLevelError)
			if !ok {
				ex := &definition.CompileError{
					Position: importVisitor.PosList[i],
					Err:      err,
				}
				errs = definition.TopLevelErrorsJoin(errs, ex)
				continue
			}
			ex := &definition.CompileError{
				Position: importVisitor.PosList[i],
				Err: &definition.ImportingError{
					File: other,
					Err:  topLevelErr,
				},
			}
			errs = definition.TopLevelErrorsJoin(errs, ex)
			continue
		}
		if otherUnit == nil {
			// no error, and no unit, skip
			continue
		}

		err = unit.AddImport(otherUnit)
		if err != nil {
			ex := &definition.CompileError{
				Position: importVisitor.PosList[i],
				Err:      err,
			}
			errs = definition.TopLevelErrorsJoin(errs, ex)
		}
	}

	if errs != nil {
		return nil, errs, warnings
	}

	r.ParseStack.Remove(ident.Path)

	protoVisitor := NewParseVisitor(unit)
	protoErr := ast.Accept(protoVisitor)
	warnings = definition.TopLevelWarningsJoin(warnings, protoVisitor.Warning)
	if err, ok := protoErr.(error); ok {
		return nil, err, warnings
	}

	return unit, nil, warnings
}

func (r *CompilationRuntime) precheckFilePath(file *definition.FileIdentifer) (skip bool, err error) {
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

func (r *CompilationRuntime) precheckPackageName(pkg *definition.Package) error {
	prev, ok := r.Packages.Get(pkg.String())
	if ok {
		return &definition.CompileError{
			Position: pkg,
			Err: &definition.PackageDuplicateError{
				PrevDef: prev.Package.BasePosition,
				Package: pkg,
			},
		}
	}
	return nil
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
