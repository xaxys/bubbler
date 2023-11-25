// Code generated from D:/GoProject/bubbler/tools/../bubbler.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // bubbler
import "github.com/antlr4-go/antlr/v4"

type BasebubblerVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasebubblerVisitor) VisitProto(ctx *ProtoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitImportStatement(ctx *ImportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitTopLevelDef(ctx *TopLevelDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitSize_(ctx *Size_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitByteSize(ctx *ByteSizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitBitSize(ctx *BitSizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitOptionName(ctx *OptionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitFieldVoid(ctx *FieldVoidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitFieldConstant(ctx *FieldConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitFieldEmbedded(ctx *FieldEmbeddedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitFieldNormal(ctx *FieldNormalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitFieldOptions(ctx *FieldOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitFieldOption(ctx *FieldOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitFieldMethods(ctx *FieldMethodsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitFieldMethod(ctx *FieldMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitType_(ctx *Type_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitBasicType(ctx *BasicTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitEnumDef(ctx *EnumDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitEnumBody(ctx *EnumBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitEnumElement(ctx *EnumElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitEnumField(ctx *EnumFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitEnumValueOptions(ctx *EnumValueOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitEnumValueOption(ctx *EnumValueOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitStructDef(ctx *StructDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitStructBody(ctx *StructBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitStructElement(ctx *StructElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitExprMulDivMod(ctx *ExprMulDivModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitExprTernary(ctx *ExprTernaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitExprBitXor(ctx *ExprBitXorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitExprParens(ctx *ExprParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitExprConstant(ctx *ExprConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitExprPower(ctx *ExprPowerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitExprLogicalOr(ctx *ExprLogicalOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitExprRelational(ctx *ExprRelationalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitExprShift(ctx *ExprShiftContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitExprBitOr(ctx *ExprBitOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitExprAddSub(ctx *ExprAddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitExprValue(ctx *ExprValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitExprCast(ctx *ExprCastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitExprUnary(ctx *ExprUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitExprBitAnd(ctx *ExprBitAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitExprEquality(ctx *ExprEqualityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitExprLogicalAnd(ctx *ExprLogicalAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitEmptyStatement_(ctx *EmptyStatement_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitIdent(ctx *IdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitStructName(ctx *StructNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitEnumName(ctx *EnumNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitFieldName(ctx *FieldNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitMethodName(ctx *MethodNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitStructType(ctx *StructTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitEnumType(ctx *EnumTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitIntLit(ctx *IntLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitStrLit(ctx *StrLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitBoolLit(ctx *BoolLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebubblerVisitor) VisitFloatLit(ctx *FloatLitContext) interface{} {
	return v.VisitChildren(ctx)
}
