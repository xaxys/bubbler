// Code generated from D:/GoProject/bubbler/tools/../bubbler.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // bubbler
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by bubblerParser.
type bubblerVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by bubblerParser#proto.
	VisitProto(ctx *ProtoContext) interface{}

	// Visit a parse tree produced by bubblerParser#importStatement.
	VisitImportStatement(ctx *ImportStatementContext) interface{}

	// Visit a parse tree produced by bubblerParser#topLevelDef.
	VisitTopLevelDef(ctx *TopLevelDefContext) interface{}

	// Visit a parse tree produced by bubblerParser#size_.
	VisitSize_(ctx *Size_Context) interface{}

	// Visit a parse tree produced by bubblerParser#byteSize.
	VisitByteSize(ctx *ByteSizeContext) interface{}

	// Visit a parse tree produced by bubblerParser#bitSize.
	VisitBitSize(ctx *BitSizeContext) interface{}

	// Visit a parse tree produced by bubblerParser#optionName.
	VisitOptionName(ctx *OptionNameContext) interface{}

	// Visit a parse tree produced by bubblerParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by bubblerParser#fieldVoid.
	VisitFieldVoid(ctx *FieldVoidContext) interface{}

	// Visit a parse tree produced by bubblerParser#fieldConstant.
	VisitFieldConstant(ctx *FieldConstantContext) interface{}

	// Visit a parse tree produced by bubblerParser#fieldEmbedded.
	VisitFieldEmbedded(ctx *FieldEmbeddedContext) interface{}

	// Visit a parse tree produced by bubblerParser#fieldNormal.
	VisitFieldNormal(ctx *FieldNormalContext) interface{}

	// Visit a parse tree produced by bubblerParser#fieldOptions.
	VisitFieldOptions(ctx *FieldOptionsContext) interface{}

	// Visit a parse tree produced by bubblerParser#fieldOption.
	VisitFieldOption(ctx *FieldOptionContext) interface{}

	// Visit a parse tree produced by bubblerParser#fieldMethods.
	VisitFieldMethods(ctx *FieldMethodsContext) interface{}

	// Visit a parse tree produced by bubblerParser#fieldMethod.
	VisitFieldMethod(ctx *FieldMethodContext) interface{}

	// Visit a parse tree produced by bubblerParser#type_.
	VisitType_(ctx *Type_Context) interface{}

	// Visit a parse tree produced by bubblerParser#basicType.
	VisitBasicType(ctx *BasicTypeContext) interface{}

	// Visit a parse tree produced by bubblerParser#enumDef.
	VisitEnumDef(ctx *EnumDefContext) interface{}

	// Visit a parse tree produced by bubblerParser#enumBody.
	VisitEnumBody(ctx *EnumBodyContext) interface{}

	// Visit a parse tree produced by bubblerParser#enumElement.
	VisitEnumElement(ctx *EnumElementContext) interface{}

	// Visit a parse tree produced by bubblerParser#enumField.
	VisitEnumField(ctx *EnumFieldContext) interface{}

	// Visit a parse tree produced by bubblerParser#enumValueOptions.
	VisitEnumValueOptions(ctx *EnumValueOptionsContext) interface{}

	// Visit a parse tree produced by bubblerParser#enumValueOption.
	VisitEnumValueOption(ctx *EnumValueOptionContext) interface{}

	// Visit a parse tree produced by bubblerParser#structDef.
	VisitStructDef(ctx *StructDefContext) interface{}

	// Visit a parse tree produced by bubblerParser#structBody.
	VisitStructBody(ctx *StructBodyContext) interface{}

	// Visit a parse tree produced by bubblerParser#structElement.
	VisitStructElement(ctx *StructElementContext) interface{}

	// Visit a parse tree produced by bubblerParser#ExprMulDivMod.
	VisitExprMulDivMod(ctx *ExprMulDivModContext) interface{}

	// Visit a parse tree produced by bubblerParser#ExprTernary.
	VisitExprTernary(ctx *ExprTernaryContext) interface{}

	// Visit a parse tree produced by bubblerParser#ExprBitXor.
	VisitExprBitXor(ctx *ExprBitXorContext) interface{}

	// Visit a parse tree produced by bubblerParser#ExprParens.
	VisitExprParens(ctx *ExprParensContext) interface{}

	// Visit a parse tree produced by bubblerParser#ExprConstant.
	VisitExprConstant(ctx *ExprConstantContext) interface{}

	// Visit a parse tree produced by bubblerParser#ExprPower.
	VisitExprPower(ctx *ExprPowerContext) interface{}

	// Visit a parse tree produced by bubblerParser#ExprLogicalOr.
	VisitExprLogicalOr(ctx *ExprLogicalOrContext) interface{}

	// Visit a parse tree produced by bubblerParser#ExprRelational.
	VisitExprRelational(ctx *ExprRelationalContext) interface{}

	// Visit a parse tree produced by bubblerParser#ExprShift.
	VisitExprShift(ctx *ExprShiftContext) interface{}

	// Visit a parse tree produced by bubblerParser#ExprBitOr.
	VisitExprBitOr(ctx *ExprBitOrContext) interface{}

	// Visit a parse tree produced by bubblerParser#ExprAddSub.
	VisitExprAddSub(ctx *ExprAddSubContext) interface{}

	// Visit a parse tree produced by bubblerParser#ExprValue.
	VisitExprValue(ctx *ExprValueContext) interface{}

	// Visit a parse tree produced by bubblerParser#ExprCast.
	VisitExprCast(ctx *ExprCastContext) interface{}

	// Visit a parse tree produced by bubblerParser#ExprUnary.
	VisitExprUnary(ctx *ExprUnaryContext) interface{}

	// Visit a parse tree produced by bubblerParser#ExprBitAnd.
	VisitExprBitAnd(ctx *ExprBitAndContext) interface{}

	// Visit a parse tree produced by bubblerParser#ExprEquality.
	VisitExprEquality(ctx *ExprEqualityContext) interface{}

	// Visit a parse tree produced by bubblerParser#ExprLogicalAnd.
	VisitExprLogicalAnd(ctx *ExprLogicalAndContext) interface{}

	// Visit a parse tree produced by bubblerParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by bubblerParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by bubblerParser#emptyStatement_.
	VisitEmptyStatement_(ctx *EmptyStatement_Context) interface{}

	// Visit a parse tree produced by bubblerParser#ident.
	VisitIdent(ctx *IdentContext) interface{}

	// Visit a parse tree produced by bubblerParser#structName.
	VisitStructName(ctx *StructNameContext) interface{}

	// Visit a parse tree produced by bubblerParser#enumName.
	VisitEnumName(ctx *EnumNameContext) interface{}

	// Visit a parse tree produced by bubblerParser#fieldName.
	VisitFieldName(ctx *FieldNameContext) interface{}

	// Visit a parse tree produced by bubblerParser#methodName.
	VisitMethodName(ctx *MethodNameContext) interface{}

	// Visit a parse tree produced by bubblerParser#structType.
	VisitStructType(ctx *StructTypeContext) interface{}

	// Visit a parse tree produced by bubblerParser#enumType.
	VisitEnumType(ctx *EnumTypeContext) interface{}

	// Visit a parse tree produced by bubblerParser#arrayType.
	VisitArrayType(ctx *ArrayTypeContext) interface{}

	// Visit a parse tree produced by bubblerParser#intLit.
	VisitIntLit(ctx *IntLitContext) interface{}

	// Visit a parse tree produced by bubblerParser#strLit.
	VisitStrLit(ctx *StrLitContext) interface{}

	// Visit a parse tree produced by bubblerParser#boolLit.
	VisitBoolLit(ctx *BoolLitContext) interface{}

	// Visit a parse tree produced by bubblerParser#floatLit.
	VisitFloatLit(ctx *FloatLitContext) interface{}
}
