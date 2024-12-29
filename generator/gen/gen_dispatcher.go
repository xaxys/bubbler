package gen

import (
	"github.com/xaxys/bubbler/definition"
)

// ==================== Literal Generator ====================

type LiteralGeneratorImpl interface {
	GenerateBoolLiteral(literal *definition.BoolLiteral) (string, error)
	GenerateIntLiteral(literal *definition.IntLiteral) (string, error)
	GenerateFloatLiteral(literal *definition.FloatLiteral) (string, error)
	GenerateStringLiteral(literal *definition.StringLiteral) (string, error)
}

type GenLiteralDispatcher struct {
	LiteralGeneratorImpl
}

func NewGenLiteralDispatcher(literalGeneratorImpl LiteralGeneratorImpl) *GenLiteralDispatcher {
	return &GenLiteralDispatcher{
		LiteralGeneratorImpl: literalGeneratorImpl,
	}
}

func (d *GenLiteralDispatcher) AcceptLiteral(literal definition.Literal) (string, error) {
	switch val := literal.(type) {
	case *definition.BoolLiteral:
		return d.LiteralGeneratorImpl.GenerateBoolLiteral(val)
	case *definition.IntLiteral:
		return d.LiteralGeneratorImpl.GenerateIntLiteral(val)
	case *definition.FloatLiteral:
		return d.LiteralGeneratorImpl.GenerateFloatLiteral(val)
	case *definition.StringLiteral:
		return d.LiteralGeneratorImpl.GenerateStringLiteral(val)
	default:
		panic("unreachable")
	}
}

// ==================== Expr Generator ====================

type ExprGeneratorImpl interface {
	GenerateUnopExpr(expr *definition.UnopExpr) (string, error)
	GenerateBinopExpr(expr *definition.BinopExpr) (string, error)
	GenerateCastExpr(expr *definition.CastExpr) (string, error)
	GenerateConstantExpr(expr *definition.ConstantExpr) (string, error)
	GenerateTenaryExpr(expr *definition.TenaryExpr) (string, error)
	GenerateValueExpr(expr *definition.ValueExpr) (string, error)
	GenerateRawExpr(expr *definition.RawExpr) (string, error)
}

type GenExprDispatcher struct {
	ExprGeneratorImpl
}

func NewGenExprDispatcher(exprGeneratorImpl ExprGeneratorImpl) *GenExprDispatcher {
	return &GenExprDispatcher{
		ExprGeneratorImpl: exprGeneratorImpl,
	}
}

func (d *GenExprDispatcher) AcceptExpr(expr definition.Expr) (string, error) {
	switch val := expr.(type) {
	case *definition.UnopExpr:
		return d.ExprGeneratorImpl.GenerateUnopExpr(val)
	case *definition.BinopExpr:
		return d.ExprGeneratorImpl.GenerateBinopExpr(val)
	case *definition.CastExpr:
		return d.ExprGeneratorImpl.GenerateCastExpr(val)
	case *definition.ConstantExpr:
		return d.ExprGeneratorImpl.GenerateConstantExpr(val)
	case *definition.TenaryExpr:
		return d.ExprGeneratorImpl.GenerateTenaryExpr(val)
	case *definition.ValueExpr:
		return d.ExprGeneratorImpl.GenerateValueExpr(val)
	case *definition.RawExpr:
		return d.ExprGeneratorImpl.GenerateRawExpr(val)
	default:
		panic("unreachable")
	}
}

func (d *GenExprDispatcher) AcceptLiteral(literal definition.Literal, impl ...LiteralGeneratorImpl) (string, error) {
	var literalImpl LiteralGeneratorImpl
	if len(impl) > 0 {
		literalImpl = impl[0]
	} else if self, ok := d.ExprGeneratorImpl.(LiteralGeneratorImpl); ok {
		literalImpl = self
	} else {
		panic("no LiteralGeneratorImpl provided")
	}

	if dispatcher, ok := literalImpl.(*GenLiteralDispatcher); ok {
		return dispatcher.AcceptLiteral(literal)
	}
	return NewGenLiteralDispatcher(literalImpl).AcceptLiteral(literal)
}

// ==================== Generator ====================

type GeneratorImpl interface {
	GenerateUnit(unit *definition.CompilationUnit) error

	// GenerateType(type_ definition.Type) (string, error)
	// GenerateTypeDefaultValue(type_ definition.Type) (string, error)
	GenerateBasicType(type_ *definition.BasicType) (string, error)
	GenerateBasicTypeDefaultValue(type_ *definition.BasicType) (string, error)
	GenerateString(string_ *definition.String) (string, error)
	GenerateStringDefaultValue(string_ *definition.String) (string, error)
	GenerateBytes(bytes *definition.Bytes) (string, error)
	GenerateBytesDefaultValue(bytes *definition.Bytes) (string, error)
	GenerateArray(array *definition.Array) (string, error)
	GenerateArrayDefaultValue(array *definition.Array) (string, error)

	GenerateStruct(structDef *definition.Struct) (string, error)
	GenerateStructDefaultValue(structDef *definition.Struct) (string, error)

	// GenerateField(field definition.Field) (string, error)
	GenerateNormalField(field *definition.NormalField) (string, error)
	GenerateVoidField(field *definition.VoidField) (string, error)
	GenerateEmbeddedField(field *definition.EmbeddedField) (string, error)
	GenerateConstantField(field *definition.ConstantField) (string, error)

	// GenerateMethod(method *definition.Method) (string, error)
	// GenerateMethodDecl(method *definition.Method) (string, error)
	GenerateDefaultGetterDecl(method *definition.GetMethod) (string, error) // unsupported for now
	GenerateDefaultSetterDecl(method *definition.SetMethod) (string, error) // unsupported for now
	GenerateCustomGetterDecl(method *definition.GetMethod) (string, error)
	GenerateCustomSetterDecl(method *definition.SetMethod) (string, error)
	// GenerateRawGetterDecl(method *definition.GetMethod) (string, error)
	// GenerateRawSetterDecl(method *definition.SetMethod) (string, error)
	GenerateDefaultGetter(method *definition.GetMethod) (string, error) // unsupported for now
	GenerateDefaultSetter(method *definition.SetMethod) (string, error) // unsupported for now
	GenerateCustomGetter(method *definition.GetMethod) (string, error)
	GenerateCustomSetter(method *definition.SetMethod) (string, error)
	// GenerateRawGetter(field definition.Field) (string, error)
	// GenerateRawSetter(field definition.Field) (string, error)

	// This part has been moved to GenExprDispatcher
	// ExprGeneratorImpl
	// GenerateExpr(expr *definition.Expr) (string, error)

	// GenerateEncoderDecl(structDef *definition.Struct) (string, error)
	// GenerateDecoderDecl(structDef *definition.Struct) (string, error)
	// GenerateEncoder(structDef *definition.Struct) (string, error)
	// GenerateDecoder(structDef *definition.Struct) (string, error)

	GenerateEnum(enumDef *definition.Enum) (string, error)
	GenerateEnumDefaultValue(enumDef *definition.Enum) (string, error)
	// GenerateEnumValue(enumValue *definition.EnumValue) (string, error)
}

type GenDispatcher struct {
	GeneratorImpl
}

func NewGenDispatcher(generatorImpl GeneratorImpl) *GenDispatcher {
	return &GenDispatcher{
		GeneratorImpl: generatorImpl,
	}
}

func (d *GenDispatcher) AcceptGenCtx(ctx *GenCtx) error {
	for _, unit := range ctx.Units {
		err := d.GeneratorImpl.GenerateUnit(unit)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *GenDispatcher) AcceptType(type_ definition.Type) (string, error) {
	switch val := type_.(type) {
	case *definition.BasicType:
		return d.GeneratorImpl.GenerateBasicType(val)
	case *definition.String:
		return d.GeneratorImpl.GenerateString(val)
	case *definition.Bytes:
		return d.GeneratorImpl.GenerateBytes(val)
	case *definition.Array:
		return d.GeneratorImpl.GenerateArray(val)
	case *definition.Struct:
		return d.GeneratorImpl.GenerateStruct(val)
	case *definition.Enum:
		return d.GeneratorImpl.GenerateEnum(val)
	default:
		panic("unreachable")
	}
}

func (d *GenDispatcher) AcceptTypeDefaultValue(type_ definition.Type) (string, error) {
	switch val := type_.(type) {
	case *definition.BasicType:
		return d.GeneratorImpl.GenerateBasicTypeDefaultValue(val)
	case *definition.String:
		return d.GeneratorImpl.GenerateStringDefaultValue(val)
	case *definition.Bytes:
		return d.GeneratorImpl.GenerateBytesDefaultValue(val)
	case *definition.Array:
		return d.GeneratorImpl.GenerateArrayDefaultValue(val)
	case *definition.Struct:
		return d.GeneratorImpl.GenerateStructDefaultValue(val)
	case *definition.Enum:
		return d.GeneratorImpl.GenerateEnumDefaultValue(val)
	default:
		panic("unreachable")
	}
}

func (d *GenDispatcher) AcceptField(field definition.Field) (string, error) {
	switch val := field.(type) {
	case *definition.NormalField:
		return d.GeneratorImpl.GenerateNormalField(val)
	case *definition.VoidField:
		return d.GeneratorImpl.GenerateVoidField(val)
	case *definition.EmbeddedField:
		return d.GeneratorImpl.GenerateEmbeddedField(val)
	case *definition.ConstantField:
		return d.GeneratorImpl.GenerateConstantField(val)
	default:
		panic("unreachable")
	}
}

func (d *GenDispatcher) AcceptMethod(method definition.Method) (string, error) {
	switch val := method.(type) {
	case *definition.GetMethod:
		if val.MethodName == "" {
			return d.GeneratorImpl.GenerateDefaultGetter(val)
		}
		return d.GeneratorImpl.GenerateCustomGetter(val)

	case *definition.SetMethod:
		if val.MethodName == "" {
			return d.GeneratorImpl.GenerateDefaultSetter(val)
		}
		return d.GeneratorImpl.GenerateCustomSetter(val)

	default:
		panic("unreachable")
	}
}

func (d *GenDispatcher) AcceptMethodDecl(method definition.Method) (string, error) {
	switch val := method.(type) {
	case *definition.GetMethod:
		if val.MethodName == "" {
			return d.GeneratorImpl.GenerateDefaultGetterDecl(val)
		}
		return d.GeneratorImpl.GenerateCustomGetterDecl(val)

	case *definition.SetMethod:
		if val.MethodName == "" {
			return d.GeneratorImpl.GenerateDefaultSetterDecl(val)
		}
		return d.GeneratorImpl.GenerateCustomSetterDecl(val)

	default:
		panic("unreachable")
	}
}

// AcceptExpr(expr definition.Expr) (string, error)
// AcceptExpr(expr definition.Expr, impl ExprGeneratorImpl) (string, error)
func (d *GenDispatcher) AcceptExpr(expr definition.Expr, impl ...ExprGeneratorImpl) (string, error) {
	var exprImpl ExprGeneratorImpl
	if len(impl) > 0 {
		exprImpl = impl[0]
	} else if self, ok := d.GeneratorImpl.(ExprGeneratorImpl); ok {
		exprImpl = self
	} else {
		panic("no ExprGeneratorImpl provided")
	}

	if dispatcher, ok := exprImpl.(*GenExprDispatcher); ok {
		return dispatcher.AcceptExpr(expr)
	}
	return NewGenExprDispatcher(exprImpl).AcceptExpr(expr)
}
