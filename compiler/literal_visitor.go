package compiler

import (
	"fmt"
	"strconv"

	"github.com/xaxys/bubbler/definition"
	"github.com/xaxys/bubbler/parser"
)

type LiteralVisitor struct {
	parser.BasebubblerVisitor
	Unit *definition.CompilationUnit
}

func NewLiteralVisitor(unit *definition.CompilationUnit) *LiteralVisitor {
	return &LiteralVisitor{
		Unit: unit,
	}
}

// ==================== Literal ====================

// VisitConstant returns Literal or error
func (v *LiteralVisitor) VisitConstant(ctx *parser.ConstantContext) any {
	if ctx.IntLit() != nil {
		constant := int64(0)
		ret := ctx.IntLit().Accept(v)
		switch val := ret.(type) {
		case error:
			return val
		case int64:
			constant = val
		default:
			panic("unreachable")
		}

		if ctx.SUB() != nil {
			constant = -constant
		}
		return &definition.IntLiteral{
			BasePosition: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.IntLit().GetStart().GetLine(),
				Column: ctx.IntLit().GetStart().GetColumn(),
			},
			IntValue: constant,
		}
	}
	if ctx.FloatLit() != nil {
		constant := float64(0)
		ret := ctx.FloatLit().Accept(v)
		switch val := ret.(type) {
		case error:
			return val
		case float64:
			constant = val
		default:
			panic("unreachable")
		}

		if ctx.SUB() != nil {
			constant = -constant
		}
		return &definition.FloatLiteral{
			BasePosition: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.FloatLit().GetStart().GetLine(),
				Column: ctx.FloatLit().GetStart().GetColumn(),
			},
			FloatValue: constant,
		}
	}
	if ctx.BoolLit() != nil {
		constant := ctx.BoolLit().Accept(v).(bool)
		return &definition.BoolLiteral{
			BasePosition: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.BoolLit().GetStart().GetLine(),
				Column: ctx.BoolLit().GetStart().GetColumn(),
			},
			BoolValue: constant,
		}
	}
	if ctx.StrLit() != nil {
		constant := ""
		ret := ctx.StrLit().Accept(v)
		switch val := ret.(type) {
		case error:
			return val
		case string:
			constant = val
		default:
			panic("unreachable")
		}
		return &definition.StringLiteral{
			BasePosition: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.StrLit().GetStart().GetLine(),
				Column: ctx.StrLit().GetStart().GetColumn(),
			},
			StringValue: constant,
		}
	}
	panic("unreachable")
}

// VisitIdent noexcept returns string
func (v *LiteralVisitor) VisitIdent(ctx *parser.IdentContext) any {
	ident := ctx.IDENTIFIER().GetText()
	return ident
}

// VisitIntLit returns int64 or error
func (v *LiteralVisitor) VisitIntLit(ctx *parser.IntLitContext) any {
	lit := ctx.INT_LIT().GetText()
	val, err := strconv.ParseInt(lit, 0, 64)
	if err != nil {
		return &definition.CompileError{
			Position: &definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.INT_LIT().GetSymbol().GetLine(),
				Column: ctx.INT_LIT().GetSymbol().GetColumn(),
			},
			Err: &definition.InvalidLiteralError{
				Literal: lit,
				Err:     err,
			},
		}
	}
	return val
}

// VisitFloatLit returns float64 or error
func (v *LiteralVisitor) VisitFloatLit(ctx *parser.FloatLitContext) any {
	lit := ctx.FLOAT_LIT().GetText()
	val, err := strconv.ParseFloat(lit, 64)
	if err != nil {
		return &definition.CompileError{
			Position: &definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.FLOAT_LIT().GetSymbol().GetLine(),
				Column: ctx.FLOAT_LIT().GetSymbol().GetColumn(),
			},
			Err: &definition.InvalidLiteralError{
				Literal: lit,
				Err:     err,
			},
		}
	}
	return val
}

// VisitBoolLit noexcept returns bool
func (v *LiteralVisitor) VisitBoolLit(ctx *parser.BoolLitContext) any {
	lit := ctx.BOOL_LIT().GetText()
	val, err := strconv.ParseBool(lit)
	if err != nil {
		// according to the grammar, this never returns error
		panic("unreachable")
	}
	return val
}

// VisitStrLit returns string or error
func (v *LiteralVisitor) VisitStrLit(ctx *parser.StrLitContext) any {
	lit := ctx.STR_LIT().GetText()
	// 'abc' -> "abc"
	if lit[0] == '\'' && lit[len(lit)-1] == '\'' {
		lit = fmt.Sprintf(`"%s"`, lit[1:len(lit)-1])
	}
	lit, err := strconv.Unquote(lit)
	if err != nil {
		return &definition.CompileError{
			Position: &definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.STR_LIT().GetSymbol().GetLine(),
				Column: ctx.STR_LIT().GetSymbol().GetColumn(),
			},
			Err: &definition.InvalidLiteralError{
				Literal: lit,
				Err:     err,
			},
		}
	}
	return lit
}

// VisitValue returns ValueExpr
func (v *LiteralVisitor) VisitValue(ctx *parser.ValueContext) any {
	return &definition.ValueExpr{}
}
