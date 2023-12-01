package compiler

import (
	"github.com/xaxys/bubbler/definition"
	"github.com/xaxys/bubbler/parser"
	"github.com/xaxys/bubbler/util"
)

type InfoVisitor struct {
	parser.BasebubblerVisitor
	Unit    *definition.CompilationUnit
	Warning definition.TopLevelWarning
}

func NewInfoVisitor(unit *definition.CompilationUnit) *InfoVisitor {
	return &InfoVisitor{
		Unit: unit,
	}
}

func (v *InfoVisitor) VisitProto(ctx *parser.ProtoContext) any {
	var err definition.TopLevelError
	packageStmts := ctx.AllPackageStatement()
	if len(packageStmts) == 0 {
		warn := &definition.CompileWarning{
			Position: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
			},
			Warning: &definition.PackageNameNotSetWarning{},
		}
		v.Warning = definition.TopLevelWarningsJoin(v.Warning, warn)
	} else {
		var pkgs []*definition.Package
		for _, pkg := range packageStmts {
			pkgRet := pkg.Accept(v)
			switch val := pkgRet.(type) {
			case definition.TopLevelError:
				err = definition.TopLevelErrorsJoin(err, val)
			case error:
				return val
			case *definition.Package:
				pkgs = append(pkgs, val)
			default:
				panic("unreachable")
			}
		}

		if err != nil {
			return err
		}

		v.Unit.Package = pkgs[0]
		for _, pkg := range pkgs[1:] {
			ex := &definition.CompileError{
				Position: pkg,
				Err: &definition.PackageDuplicateError{
					PrevDef:     v.Unit.Package,
					PackageName: v.Unit.Package.String(),
				},
			}
			err = definition.TopLevelErrorsJoin(err, ex)
		}
	}

	options := util.NewOrderedMap[string, *definition.Option]()
	addOption := func(o *definition.Option) definition.TopLevelError {
		prev, ok := options.Get(o.OptionName)
		if ok {
			return &definition.CompileError{
				Position: o,
				Err: &definition.OptionDuplicateError{
					PrevDef:    prev,
					OptionName: o.OptionName,
				},
			}
		}
		options.Put(o.OptionName, o)
		return nil
	}

	elems := ctx.AllOptionStatement()
	for _, elem := range elems {
		ret := elem.Accept(v)
		switch val := ret.(type) {
		case definition.TopLevelError:
			err = definition.TopLevelErrorsJoin(err, val)
		case error:
			return val
		case *definition.Option:
			ex := FileOptionValidator.ValidateOption(val)
			if ex != nil {
				if warn, ok := ex.(definition.TopLevelWarning); ok {
					v.Warning = definition.TopLevelWarningsJoin(v.Warning, warn)
				} else if topErr, ok := ex.(definition.TopLevelError); ok {
					err = definition.TopLevelErrorsJoin(err, topErr)
					continue
				} else {
					return ex
				}
			}
			topEx := addOption(val)
			if topEx != nil {
				err = definition.TopLevelErrorsJoin(err, topEx)
			}

		default:
			panic("unreachable")
		}
	}

	if err != nil {
		return err
	}

	return nil
}

// ==================== Package ====================

func (v *InfoVisitor) VisitPackageStatement(ctx *parser.PackageStatementContext) any {
	var paths []string
	pathsRet := ctx.FullIdent().Accept(v)
	switch val := pathsRet.(type) {
	case error:
		return val
	case []string:
		paths = val
	default:
		panic("unreachable")
	}

	pkg := &definition.Package{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		PackagePaths: paths,
	}
	return pkg
}

// ==================== Option ====================

func (v *InfoVisitor) VisitOptionStatement(ctx *parser.OptionStatementContext) any {
	name := ctx.OptionName().GetText()

	var value definition.Literal
	lv := NewLiteralVisitor(v.Unit)
	valueRet := ctx.Constant().Accept(lv)
	switch val := valueRet.(type) {
	case error:
		return val
	case definition.Literal:
		value = val
	default:
		panic("unreachable")
	}

	option := &definition.Option{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		OptionName:  name,
		OptionValue: value,
	}

	return option
}

// ==================== FullIdent ====================

func (v *InfoVisitor) VisitFullIdent(ctx *parser.FullIdentContext) any {
	var idents []string
	for _, child := range ctx.AllIdent() {
		part := child.GetText()
		idents = append(idents, part)
	}
	return idents
}
