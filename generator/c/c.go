package c

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"text/template"

	"github.com/xaxys/bubbler/definition"
	"github.com/xaxys/bubbler/util"
)

type GeneratedType struct {
	GeneratedConst string
	GeneratedDef   string
	GeneratedFunc  string
}

type CGenerator struct {
	GenTypes *util.OrderedMap[string, *GeneratedType]
	GenStack *util.OrderedMap[string, any]
}

func NewCGenerator() *CGenerator {
	return &CGenerator{}
}

var fileTemplate = `
{{- define "file" -}}
// Target: C
// Generated by bubbler
// DO NOT EDIT
#include <stdint.h>
#include <stdbool.h>
#include <math.h>

{{ range $type := .GenTypes.Values -}}
{{- if $type.GeneratedConst }}
{{ $type.GeneratedConst }}
{{ end }}

{{- if $type.GeneratedDef }}
{{ $type.GeneratedDef }};
{{ end }}

{{- if $type.GeneratedFunc }}
{{ $type.GeneratedFunc }}
{{ end }}

{{- end -}}
{{- end -}}
`

func (g *CGenerator) Generate(unit *definition.CompilationUnit) (string, error) {
	g.GenTypes = util.NewOrderedMap[string, *GeneratedType]()
	g.GenStack = util.NewOrderedMap[string, any]()

	for _, type_ := range unit.Types.Values() {
		g.generateType(type_)
	}

	data := map[string]any{
		"GenTypes": g.GenTypes,
	}

	str := util.ExecuteTemplate(fileTemplate, "file", nil, data)

	return str, nil
}

var typeMap = map[definition.TypeID]string{
	definition.TypeID_Bool:    "bool",
	definition.TypeID_Uint8:   "uint8_t",
	definition.TypeID_Uint16:  "uint16_t",
	definition.TypeID_Uint32:  "uint32_t",
	definition.TypeID_Uint64:  "uint64_t",
	definition.TypeID_Int8:    "int8_t",
	definition.TypeID_Int16:   "int16_t",
	definition.TypeID_Int32:   "int32_t",
	definition.TypeID_Int64:   "int64_t",
	definition.TypeID_Float32: "float",
	definition.TypeID_Float64: "double",
	definition.TypeID_String:  "char*",
	definition.TypeID_Bytes:   "uint8_t*",
}

var typeSizeMapInt = map[int64]string{
	8:  "int8_t",
	16: "int16_t",
	32: "int32_t",
	64: "int64_t",
}

var typeSizeMapUint = map[int64]string{
	8:  "uint8_t",
	16: "uint16_t",
	32: "uint32_t",
	64: "uint64_t",
}

func (g *CGenerator) generateType(type_ definition.Type) (def string) {
	switch val := type_.(type) {
	case *definition.Struct:
		name := fmt.Sprintf("struct %s", val.StructName)
		if g.GenTypes.Has(val.StructName) || g.GenStack.Has(val.StructName) {
			return name
		}
		g.GenStack.Put(val.StructName, nil)
		genTy := g.generateStruct(val)
		g.GenStack.Remove(val.StructName)
		g.GenTypes.Put(val.StructName, genTy)
		return name

	case *definition.Enum:
		name := fmt.Sprintf("enum %s", val.EnumName)
		if g.GenTypes.Has(val.EnumName) || g.GenStack.Has(val.EnumName) {
			return name
		}
		g.GenStack.Put(val.EnumName, nil)
		genTy := g.generateEnum(val)
		g.GenStack.Remove(val.EnumName)
		g.GenTypes.Put(val.EnumName, genTy)
		return name

	case *definition.BasicType:
		return typeMap[val.TypeTypeID]

	case *definition.Array:
		panic("unreachable")

	case *definition.String:
		panic("unimplemented")

	case *definition.Bytes:
		panic("unimplemented")

	default:
		panic("unreachable")
	}
}

var structTemplate = `
{{- define "constantField" -}}
	{{- $pos := .Pos -}}
	{{- $field := .Field -}}
	{{- $fieldName := Tosnake_case .Field.FieldName -}}
	// {{ $pos }} ConstantField: {{ $field.ShortString }}
	{{ if $fieldName }}
		{{- $tyStr := generateType $field.FieldType -}}
		{{ $tyStr }} {{ $fieldName }};
	{{- end }}
{{- end -}}

{{- define "voidField" -}}
	{{- $pos := .Pos -}}
	{{- $field := .Field -}}
	// {{ $pos }} VoidField: {{ $field.ShortString }}
{{- end -}}

{{- define "embeddedField" -}}
	{{- $pos := .Pos -}}
	{{- $field := .Field -}}
	// {{ $pos }} EmbeddedField: {{ $field.ShortString }}
{{- end -}}

{{- define "normalField" -}}
	{{- $pos := .Pos -}}
	{{- $field := .Field -}}
	{{- $fieldName := Tosnake_case .Field.FieldName -}}
	// {{ $pos }} NormalField: {{ $field.ShortString }}
	{{ if $field.FieldType.GetTypeID.IsArray }}
		{{- $arrayType := $field.FieldType.ElementType -}}
		{{- $tyStr := generateType $arrayType -}}
		{{- $length := $field.FieldType.Length -}}
		{{ $tyStr }} {{ $fieldName }}[{{ $length }}];
	{{- else }}
		{{- $tyStr := generateType $field.FieldType -}}
		{{ $tyStr }} {{ $fieldName }};
	{{- end -}}
{{- end -}}

{{- define "field" -}}
	{{- if eq .Field.GetFieldKind .FieldKindID_Constant }}
		{{- template "constantField" . }}
	{{- else if eq .Field.GetFieldKind .FieldKindID_Void }}
		{{- template "voidField" . }}
	{{- else if eq .Field.GetFieldKind .FieldKindID_Embedded }}
		{{- template "embeddedField" . }}
	{{- else if eq .Field.GetFieldKind .FieldKindID_Normal }}
		{{- template "normalField" . }}
	{{- else }}
		{{- panic "unreachable" }}
	{{- end -}}
{{- end -}}

{{- define "structConst" -}}
static const uint64_t {{ .StructName }}_Size = {{ .StructSize }};
{{- end -}}

{{- define "structDef" -}}
{{- $structDef := .StructDef -}}
{{- $fieldStrs := .FieldStrs -}}
// Struct: {{ $structDef.ShortString }}
struct {{ $structDef.StructName }} {
{{- range $fieldStr := $fieldStrs }}
	{{ $fieldStr }}
{{- end }}
}
{{- end -}}
`

func (g *CGenerator) generateStruct(structDef *definition.Struct) *GeneratedType {
	funcMap := template.FuncMap{
		"generateType": g.generateType,
	}

	fields := []string{}
	startBits := int64(0)
	for _, field := range structDef.StructFields {
		from := startBits
		to := startBits + field.GetFieldBitSize()
		startBits += field.GetFieldBitSize()

		pos := fmt.Sprintf("[%s, %s)", util.ToSizeString(from), util.ToSizeString(to))
		if from == to {
			pos = "[virtual]"
		}

		fieldData := map[string]any{
			"Pos":                  pos,
			"Field":                field,
			"FieldKindID_Constant": definition.FieldKindID_Constant,
			"FieldKindID_Void":     definition.FieldKindID_Void,
			"FieldKindID_Embedded": definition.FieldKindID_Embedded,
			"FieldKindID_Normal":   definition.FieldKindID_Normal,
		}

		fieldStr := util.ExecuteTemplate(structTemplate, "field", funcMap, fieldData)
		fields = append(fields, fieldStr)
	}

	constData := map[string]any{
		"StructName": structDef.StructName,
		"StructSize": structDef.StructBitSize / 8,
	}
	constStr := util.ExecuteTemplate(structTemplate, "structConst", funcMap, constData)

	defData := map[string]any{
		"StructDef": structDef,
		"FieldStrs": fields,
	}

	defStr := util.ExecuteTemplate(structTemplate, "structDef", funcMap, defData)

	funcStr := g.generateEncoder(structDef) + "\n" + g.generateDecoder(structDef) + "\n"

	funcStr += g.generateStructMethods(structDef)

	code := &GeneratedType{
		GeneratedConst: constStr,
		GeneratedDef:   defStr,
		GeneratedFunc:  funcStr,
	}
	return code
}

var methodsTemplate = `
{{- define "defaultGetter" -}}
{{- $returnType := generateType .ReturnType -}}
{{- $fieldName := Tosnake_case .FieldDef.FieldName -}}
{{- $valueReplacement := printf "structPtr->%s" $fieldName -}}
{{- $expr := generateExpr .Expr $valueReplacement -}}
// DefaultGetter: {{ .StructName }}_get_{{ $fieldName }}
static {{ $returnType }} {{ .StructName }}_get_{{ $fieldName }}(struct {{ .StructName }}* structPtr) {
	return {{ $expr }};
}

{{ end -}}

{{- define "defaultSetter" -}}
{{- $paramType := generateType .ParamType -}}
{{- $fieldName := Tosnake_case .FieldDef.FieldName -}}
{{- $expr := generateExpr .Expr "value" -}}
// DefaultSetter: {{ .StructName }}_set_{{ $fieldName }}
static void {{ .StructName }}_set_{{ $fieldName }}(struct {{ .StructName }}* structPtr, {{ $paramType }} value) {
	structPtr->{{ $fieldName }} = {{ $expr }};
}

{{ end -}}

{{- define "customGetter" -}}
{{- $returnType := generateType .ReturnType -}}
{{- $fieldName := Tosnake_case .FieldDef.FieldName -}}
{{- $methodName := Tosnake_case .Name -}}
{{- $valueReplacement := printf "structPtr->%s" $fieldName -}}
{{- $expr := generateExpr .Expr $valueReplacement -}}
// CustomGetter: {{ .StructName }}_get_{{ $fieldName }}_{{ $methodName }}
static {{ $returnType }} {{ .StructName }}_get_{{ $fieldName }}_{{ $methodName }}(struct {{ .StructName }}* structPtr) {
	return {{ $expr }};
}

{{ end -}}

{{- define "customSetter" -}}
{{- $paramType := generateType .ParamType -}}
{{- $fieldName := Tosnake_case .FieldDef.FieldName -}}
{{- $methodName := Tosnake_case .Name -}}
{{- $expr := generateExpr .Expr "value" -}}
// CustomSetter: {{ .StructName }}_set_{{ $fieldName }}_{{ $methodName }}
static void {{ .StructName }}_set_{{ $fieldName }}_{{ $methodName }}(struct {{ .StructName }}* structPtr, {{ $paramType }} value) {
	structPtr->{{ $fieldName }} = {{ $expr }};
}

{{ end -}}

{{- define "rawGetter" -}}
{{- $returnType := generateType .FieldDef.FieldType -}}
{{- $fieldName := Tosnake_case .FieldDef.FieldName -}}
// RawGetter: {{ .StructName }}_get_{{ $fieldName }}
static {{ $returnType }} {{ .StructName }}_getraw_{{ $fieldName }}(struct {{ .StructName }}* structPtr) {
	return structPtr->{{ $fieldName }};
}

{{ end -}}

{{- define "rawSetter" -}}
{{- $paramType := generateType .FieldDef.FieldType -}}
{{- $fieldName := Tosnake_case .FieldDef.FieldName -}}
// RawSetter: {{ .StructName }}_set_{{ $fieldName }}
static void {{ .StructName }}_setraw_{{ $fieldName }}(struct {{ .StructName }}* structPtr, {{ $paramType }} value) {
	structPtr->{{ $fieldName }} = value;
}

{{ end -}}

{{- define "rawGetterArray" -}}
{{- $elemType := generateType .FieldDef.FieldType.ElementType -}}
{{- $length := .FieldDef.FieldType.Length -}}
{{- $fieldName := Tosnake_case .FieldDef.FieldName -}}
// RawGetterArray: {{ .StructName }}_get_{{ $fieldName }}
static {{ $elemType }}* {{ .StructName }}_getraw_{{ $fieldName }}(struct {{ .StructName }}* structPtr) {
	return structPtr->{{ $fieldName }};
}

// RawGetterArrayItem: {{ .StructName }}_getraw_{{ $fieldName }}_item
static {{ $elemType }} {{ .StructName }}_getraw_{{ $fieldName }}_item(struct {{ .StructName }}* structPtr, int index) {
	return structPtr->{{ $fieldName }}[index];
}

// RawGetterArrayLength: {{ .StructName }}_getraw_{{ $fieldName }}_length
static uint64_t {{ .StructName }}_getraw_{{ $fieldName }}_length(struct {{ .StructName }}* structPtr) {
	return {{ $length }};
}

{{ end -}}

{{- define "rawSetterArray" -}}
{{- $elemType := generateType .FieldDef.FieldType.ElementType -}}
{{- $length := .FieldDef.FieldType.Length -}}
{{- $fieldName := Tosnake_case .FieldDef.FieldName -}}
// RawSetterArray: {{ .StructName }}_set_{{ $fieldName }}
static void {{ .StructName }}_setraw_{{ $fieldName }}(struct {{ .StructName }}* structPtr, {{ $elemType }}* value) {
{{- range $i := iterate 0 $length }}
	structPtr->{{ $fieldName }}[{{ $i }}] = value[{{ $i }}];
{{- end }}
}

// RawSetterArrayItem: {{ .StructName }}_setraw_{{ $fieldName }}_item
static void {{ .StructName }}_setraw_{{ $fieldName }}_item(struct {{ .StructName }}* structPtr, int index, {{ $elemType }} value) {
	structPtr->{{ $fieldName }}[index] = value;
}

{{ end -}}

{{- define "fieldMethods" -}}
{{- $structName := .StructName -}}
{{- $fieldDef := .FieldDef -}}
{{- $methodKindID_Get := .MethodKindID_Get -}}
{{- $methodKindID_Set := .MethodKindID_Set -}}
// FieldMethods: {{ .StructName }}.{{ Tosnake_case .FieldDef.FieldName }}
{{ if $fieldDef.FieldType.GetTypeID.IsArray }}
	{{- template "rawGetterArray" (dict "StructName" $structName "FieldDef" $fieldDef) }}
	{{- template "rawSetterArray" (dict "StructName" $structName "FieldDef" $fieldDef) }}
{{- else -}}
	{{- template "rawGetter" (dict "StructName" $structName "FieldDef" $fieldDef) }}
	{{- template "rawSetter" (dict "StructName" $structName "FieldDef" $fieldDef) }}
{{- end -}}
{{ range $method := .FieldDef.FieldMethods }}
	{{- if eq $method.MethodKind $methodKindID_Get }}
		{{- if $method.MethodName }}
			{{- $getterData := dict "StructName" $structName "FieldDef" $fieldDef "ReturnType" $method.MethodParamType "Expr" $method.MethodExpr "Name" $method.MethodName }}
			{{- template "customGetter" $getterData }}
		{{- else }}
			{{- $getterData := dict "StructName" $structName "FieldDef" $fieldDef "ReturnType" $method.MethodParamType "Expr" $method.MethodExpr }}
			{{- template "defaultGetter" $getterData }}
		{{- end }}
	{{- else if eq $method.MethodKind $methodKindID_Set }}
		{{- if $method.MethodName }}
			{{- $setterData := dict "StructName" $structName "FieldDef" $fieldDef "ParamType" $method.MethodParamType "Expr" $method.MethodExpr "Name" $method.MethodName }}
			{{- template "customSetter" $setterData }}
		{{- else }}
			{{- $setterData := dict "StructName" $structName "FieldDef" $fieldDef "ParamType" $method.MethodParamType "Expr" $method.MethodExpr }}
			{{- template "defaultSetter" $setterData }}
		{{- end }}
	{{- else }}
		{{- panic "unreachable" }}
	{{- end -}}
{{- end -}}
{{- end -}}

{{- define "structMethods" -}}
{{- $structName := .StructDef.StructName -}}
{{- $fieldKindID_Constant := .FieldKindID_Constant -}}
{{- $fieldKindID_Void := .FieldKindID_Void -}}
{{- $fieldKindID_Embedded := .FieldKindID_Embedded -}}
{{- $fieldKindID_Normal := .FieldKindID_Normal -}}
{{- $methodKindID_Get := .MethodKindID_Get -}}
{{- $methodKindID_Set := .MethodKindID_Set -}}
{{- range $field := .StructDef.StructFields }}
	{{- if eq $field.GetFieldKind $fieldKindID_Normal }}
		{{- template "fieldMethods" dict "StructName" $structName "FieldDef" $field "MethodKindID_Get" $methodKindID_Get "MethodKindID_Set" $methodKindID_Set }}
	{{- end -}}
{{- end -}}
{{- end -}}
`

func (g *CGenerator) generateStructMethods(structDef *definition.Struct) string {
	funcMap := template.FuncMap{
		"generateExpr": g.generateExpr,
		"generateType": g.generateType,
	}

	data := map[string]interface{}{
		"StructDef":            structDef,
		"FieldKindID_Constant": definition.FieldKindID_Constant,
		"FieldKindID_Void":     definition.FieldKindID_Void,
		"FieldKindID_Embedded": definition.FieldKindID_Embedded,
		"FieldKindID_Normal":   definition.FieldKindID_Normal,
		"MethodKindID_Get":     definition.MethodKindID_Get,
		"MethodKindID_Set":     definition.MethodKindID_Set,
	}

	methodsStr := util.ExecuteTemplate(methodsTemplate, "structMethods", funcMap, data)
	return methodsStr
}

var exprTemplate = `
{{- define "unopExpr" -}}
    ({{.Expr.Op}}{{generateExpr .Expr.Expr1 .ValueReplacement}})
{{- end -}}

{{- define "powExpr" -}}
    pow({{generateExpr .Expr.Expr1 .ValueReplacement}}, {{generateExpr .Expr.Expr2 .ValueReplacement}})
{{- end -}}

{{- define "binopExpr" -}}
    ({{generateExpr .Expr.Expr1 .ValueReplacement}} {{.Expr.Op}} {{generateExpr .Expr.Expr2 .ValueReplacement}})
{{- end -}}

{{- define "castExpr" -}}
    ({{generateType .Expr.ToType}})({{generateExpr .Expr.Expr1 .ValueReplacement}})
{{- end -}}

{{- define "tenaryExpr" -}}
    ({{generateExpr .Expr.Cond .ValueReplacement}} ? {{generateExpr .Expr.Expr1 .ValueReplacement}} : {{generateExpr .Expr.Expr2 .ValueReplacement}})
{{- end -}}

{{- define "constantExpr" -}}
    {{.Expr.ConstantValue}}
{{- end -}}

{{- define "valueExpr" -}}
    {{.ValueReplacement}}
{{- end -}}

{{- define "expr" -}}
    {{- if eq .Expr.GetExprKind .ExprKindID_UnopExpr -}}
        {{- template "unopExpr" . -}}
    {{- else if eq .Expr.GetExprKind .ExprKindID_BinopExpr -}}
        {{- if eq .Expr.Op .ExprOp_POW -}}
            {{- template "powExpr" . -}}
        {{- else -}}
            {{- template "binopExpr" . -}}
        {{- end -}}
    {{- else if eq .Expr.GetExprKind .ExprKindID_CastExpr -}}
        {{- template "castExpr" . -}}
    {{- else if eq .Expr.GetExprKind .ExprKindID_TenaryExpr -}}
        {{- template "tenaryExpr" . -}}
    {{- else if eq .Expr.GetExprKind .ExprKindID_ConstantExpr -}}
        {{- template "constantExpr" . -}}
    {{- else if eq .Expr.GetExprKind .ExprKindID_ValueExpr -}}
        {{- template "valueExpr" . -}}
    {{- else -}}
        {{- panic .Expr -}}
    {{- end -}}
{{- end -}}
`

func (g *CGenerator) generateExpr(expr definition.Expr, valueReplacement string) string {
	funcMap := template.FuncMap{
		"generateExpr": g.generateExpr,
		"generateType": g.generateType,
	}

	data := map[string]interface{}{
		"Expr":                    expr,
		"ValueReplacement":        valueReplacement,
		"ExprKindID_UnopExpr":     definition.ExprKindID_UnopExpr,
		"ExprKindID_BinopExpr":    definition.ExprKindID_BinopExpr,
		"ExprKindID_CastExpr":     definition.ExprKindID_CastExpr,
		"ExprKindID_TenaryExpr":   definition.ExprKindID_TenaryExpr,
		"ExprKindID_ConstantExpr": definition.ExprKindID_ConstantExpr,
		"ExprKindID_ValueExpr":    definition.ExprKindID_ValueExpr,
		"ExprOp_ADD":              definition.ExprOp_ADD,
		"ExprOp_SUB":              definition.ExprOp_SUB,
		"ExprOp_MUL":              definition.ExprOp_MUL,
		"ExprOp_DIV":              definition.ExprOp_DIV,
		"ExprOp_MOD":              definition.ExprOp_MOD,
		"ExprOp_POW":              definition.ExprOp_POW,
		"ExprOp_SHL":              definition.ExprOp_SHL,
		"ExprOp_SHR":              definition.ExprOp_SHR,
		"ExprOp_LT":               definition.ExprOp_LT,
		"ExprOp_LE":               definition.ExprOp_LE,
		"ExprOp_GT":               definition.ExprOp_GT,
		"ExprOp_GE":               definition.ExprOp_GE,
		"ExprOp_EQ":               definition.ExprOp_EQ,
		"ExprOp_NE":               definition.ExprOp_NE,
		"ExprOp_BAND":             definition.ExprOp_BAND,
		"ExprOp_BXOR":             definition.ExprOp_BXOR,
		"ExprOp_BOR":              definition.ExprOp_BOR,
		"ExprOp_AND":              definition.ExprOp_AND,
		"ExprOp_OR":               definition.ExprOp_OR,
		"ExprOp_NOT":              definition.ExprOp_NOT,
		"ExprOp_BNOT":             definition.ExprOp_BNOT,
	}

	exprStr := util.ExecuteTemplate(exprTemplate, "expr", funcMap, data)
	return exprStr
}

var enumTemplate = `
{{- define "enumDef" -}}
// Enum: {{ .ShortString }}
enum {{ .EnumName }} {
{{- range .EnumValues }}
    {{ .EnumValueName }} = {{ .EnumValue }},
{{- end }}
}
{{- end -}}
`

func (g *CGenerator) generateEnum(enumDef *definition.Enum) *GeneratedType {
	enumDefStr := util.ExecuteTemplate(enumTemplate, "enumDef", nil, enumDef)

	code := &GeneratedType{
		GeneratedConst: "",
		GeneratedDef:   enumDefStr,
		GeneratedFunc:  "",
	}

	return code
}

//
// ==================== Encoder & Decoder ====================
// TODO: Templateify the encoder & decoder
//

func (g *CGenerator) generateEncoder(structDef *definition.Struct) string {
	genEncode := func(from, to int64, fieldData func(int64) string) string {
		encodeStr := ""
		for i := from; i < to; i = (i + 8) & (^7) {
			nextI := min(to, (i+8)&(^7))
			dataMask := ((1 << (((nextI - 1) & 7) + 1)) - 1) & (^((1 << (i & 7)) - 1))
			operator := "="
			if i%8 != 0 {
				operator = "|="
			}

			begin := i - from
			end := nextI - from
			fieldStr := ""

			j := begin
			if j < end {
				nextJ := min(end, (j+8)&(^7))
				fieldMask := ((1 << (((nextJ - 1) & 7) + 1)) - 1) & (^((1 << (j & 7)) - 1))
				shiftRight := j % 8
				fieldStr += fmt.Sprintf("(((%s) & 0b%b) >> %d)", fieldData(j/8), fieldMask, shiftRight)
				j = nextJ
			}
			if j < end {
				nextJ := min(end, (j+8)&(^7))
				fieldMask := ((1 << (((nextJ - 1) & 7) + 1)) - 1) & (^((1 << (j & 7)) - 1))
				shiftLeft := 8 - nextJ%8
				fieldStr += fmt.Sprintf(" | (((%s) & 0b%b) << %d)", fieldData(j/8), fieldMask, shiftLeft)
				j = nextJ
			}

			shiftLeft := i % 8
			encodeStr += fmt.Sprintf("((uint8_t*)data)[%d] %s ((%s << %d) & 0b%b);\n", i/8, operator, fieldStr, shiftLeft, dataMask)
		}
		return encodeStr
	}

	str := fmt.Sprintf("// Encoder: %s\n", structDef.StructName)
	str += fmt.Sprintf("static void %s_encode(struct %s* structPtr, void* data) {\n", structDef.StructName, structDef.StructName)
	encodeStr := ""
	startBits := int64(0)
	for fieldIndex, field := range structDef.StructFields {
		from := startBits
		to := startBits + field.GetFieldBitSize()
		startBits += field.GetFieldBitSize()
		pos := fmt.Sprintf("[%s, %s)", util.ToSizeString(from), util.ToSizeString(to))

		switch val := field.(type) {
		case *definition.ConstantField:
			encodeStr += fmt.Sprintf("// %s %s\n", pos, val.ShortString())

			var byteOrder binary.ByteOrder = binary.LittleEndian
			if option, ok := val.FieldOptions.Get("order"); ok {
				if option.OptionValue.GetLiteralValue() == "big" {
					byteOrder = binary.BigEndian
				}
			}

			buffer := &bytes.Buffer{}
			value := val.FieldConstant.GetLiteralValueIn(val.FieldType)
			err := binary.Write(buffer, byteOrder, value)
			if err != nil {
				panic(fmt.Errorf("internal error: %s", err))
			}
			data := buffer.Bytes()

			fieldData := func(i int64) string {
				return fmt.Sprintf("0x%X", data[i])
			}
			encodeStr += genEncode(from, to, fieldData)

		case *definition.VoidField:
			encodeStr += fmt.Sprintf("// %s %s\n", pos, val.ShortString())
			encodeStr += genEncode(from, to, func(i int64) string { return "0" })

		case *definition.EmbeddedField:
			continue

		case *definition.NormalField:
			encodeStr += fmt.Sprintf("// %s %s\n", pos, val.ShortString())
			name := fmt.Sprintf("structPtr->%s", val.FieldName)

			switch ty := val.FieldType.(type) {
			case *definition.Struct:
				encodeStr += fmt.Sprintf("%s_encode(&(%s), ((uint8_t*)data) + %d);\n", ty.StructName, name, from/8)

			case *definition.Enum:
				fieldSize := (val.FieldBitSize + 7) / 8
				tySize := util.HighBit(fieldSize)
				tyUint := typeSizeMapUint[tySize*8]
				tempName := fmt.Sprintf("temp_field_%d", fieldIndex)
				encodeStr += fmt.Sprintf("%s %s = (%s)%s;\n", tyUint, tempName, tyUint, name)

				var fieldData func(int64) string
				// 默认编码为小端序
				if option, ok := val.FieldOptions.Get("order"); ok {
					if option.OptionValue.GetLiteralValue() == "big" {
						// 大端序
						fieldData = func(i int64) string {
							return fmt.Sprintf("(%s >> %d)", tempName, (fieldSize-i-1)*8)
						}
					}
				}
				if fieldData == nil {
					// 小端序
					fieldData = func(i int64) string {
						return fmt.Sprintf("(%s >> %d)", tempName, i*8)
					}
				}
				encodeStr += genEncode(from, to, fieldData)

			case *definition.BasicType:
				tySize := (ty.TypeBitSize + 7) / 8
				tyUint := typeSizeMapUint[tySize*8]
				fieldSize := (val.FieldBitSize + 7) / 8
				var fieldData func(int64) string
				// 默认编码为小端序
				if option, ok := val.FieldOptions.Get("order"); ok {
					if option.OptionValue.GetLiteralValue() == "big" {
						// 大端序
						fieldData = func(i int64) string {
							return fmt.Sprintf("((*(%s*)(&(%s))) >> %d)", tyUint, name, (fieldSize-i-1)*8)
						}
					}
				}
				if fieldData == nil {
					// 小端序
					fieldData = func(i int64) string {
						return fmt.Sprintf("((*(%s*)(&(%s))) >> %d)", tyUint, name, i*8)
					}
				}
				encodeStr += genEncode(from, to, fieldData)

			case *definition.Array:
				elemTySize := (ty.ElementType.GetTypeBitSize() + 7) / 8
				elemBitSize := val.FieldBitSize / ty.Length
				elemSize := (elemBitSize + 7) / 8
				tyUint := typeSizeMapUint[elemTySize*8]
				var fieldDataIndex func(int64) func(int64) string
				// 默认编码为小端序
				if option, ok := val.FieldOptions.Get("order"); ok {
					if option.OptionValue.GetLiteralValue() == "big" {
						// 大端序
						fieldDataIndex = func(index int64) func(int64) string {
							return func(i int64) string {
								return fmt.Sprintf("((*(%s*)(&((%s)[%d]))) >> %d)", tyUint, name, index, (elemSize-i-1)*8)
							}
						}
					}
				}
				if fieldDataIndex == nil {
					// 小端序
					fieldDataIndex = func(index int64) func(int64) string {
						return func(i int64) string {
							return fmt.Sprintf("((*(%s*)(&((%s)[%d]))) >> %d)", tyUint, name, index, i*8)
						}
					}
				}
				for i := int64(0); i < ty.Length; i++ {
					subFrom := from + i*elemBitSize
					subTo := from + (i+1)*elemBitSize
					fieldData := fieldDataIndex(i)
					encodeStr += genEncode(subFrom, subTo, fieldData)
				}

			default:
				fieldData := func(i int64) string {
					return fmt.Sprintf("((uint8_t*)&(%s))[%d]", name, i)
				}
				encodeStr += genEncode(from, to, fieldData)
			}
		default:
			panic("unreachable")
		}
	}
	encodeStr = util.IndentSpace4(encodeStr)
	str += encodeStr
	str += fmt.Sprintf("}\n")
	return str
}

func (g *CGenerator) generateDecoder(structDef *definition.Struct) string {
	genDecode := func(from, to int64, fieldProcessor func(string, int64) string) string {
		decodeStr := ""
		for i := int64(0); i < to-from; i += 8 {
			nextI := min(to-from, (i+8)&(^7))
			// dataMask := ((1 << (((nextI - 1) & 7) + 1)) - 1) & (^((1 << (i & 7)) - 1))

			begin := from + i
			end := from + nextI
			fieldStr := ""

			j := begin
			if j < end {
				nextJ := min(end, (j+8)&(^7))
				fieldMask := ((1 << (((nextJ - 1) & 7) + 1)) - 1) & (^((1 << (j & 7)) - 1))
				shiftLeft := j % 8
				fieldStr += fmt.Sprintf("((((uint8_t*)data)[%d] & 0b%b) >> %d)", j/8, fieldMask, shiftLeft)
				j = nextJ
			}
			if j < end {
				nextJ := min(end, (j+8)&(^7))
				fieldMask := ((1 << (((nextJ - 1) & 7) + 1)) - 1) & (^((1 << (j & 7)) - 1))
				shiftRight := 8 - nextJ%8
				fieldStr += fmt.Sprintf(" | ((((uint8_t*)data)[%d] & 0b%b) << %d)", j/8, fieldMask, shiftRight)
				j = nextJ
			}

			decodeStr += fmt.Sprintf("%s;\n", fieldProcessor(fieldStr, i/8))
		}
		return decodeStr
	}

	signExtend := func(from, to int64, fieldStr string) string {
		if from >= to {
			return fieldStr
		}
		return fmt.Sprintf("((int%d_t)((%s) << %d) >> %d)", to, fieldStr, to-from, to-from)
	}

	// another sign extend implementation
	// signExtend2 := func(from, to int64, fieldStr string) string {
	//     if from >= to {
	//         return fieldStr
	//     }
	//     signMask := int64(1) << (from - 1)
	//     return fmt.Sprintf("((%s ^ 0x%X) - 0x%X)", fieldStr, signMask, signMask)
	// }

	str := fmt.Sprintf("// Decoder: %s\n", structDef.StructName)
	str += fmt.Sprintf("static bool %s_decode(void* data, struct %s* structPtr) {\n", structDef.StructName, structDef.StructName)
	decodeStr := ""
	startBits := int64(0)
	for fieldIndex, field := range structDef.StructFields {
		from := startBits
		to := startBits + field.GetFieldBitSize()
		startBits += field.GetFieldBitSize()
		pos := fmt.Sprintf("[%s, %s)", util.ToSizeString(from), util.ToSizeString(to))

		switch val := field.(type) {

		case *definition.ConstantField:
			decodeStr += fmt.Sprintf("// %s %s\n", pos, val.ShortString())

			name := fmt.Sprintf("structPtr->%s", val.FieldName)
			if val.FieldName == "" {
				tyStr := g.generateType(val.FieldType)
				name = fmt.Sprintf("temp_field_%d", fieldIndex)
				decodeStr += fmt.Sprintf("%s %s;\n", tyStr, name)
			}

			tySize := (val.FieldType.GetTypeBitSize() + 7) / 8
			tyUint := typeSizeMapUint[tySize*8]
			fieldSize := (val.FieldBitSize + 7) / 8
			var fieldProcessor func(string, int64) string

			// 默认解码为小端序
			if option, ok := val.FieldOptions.Get("order"); ok {
				if option.OptionValue.GetLiteralValue() == "big" {
					// 大端序
					fieldProcessor = func(fieldStr string, i int64) string {
						operator := "="
						if i != 0 {
							operator = "|="
						}
						return fmt.Sprintf("(*(%s*)(&(%s))) %s ((%s)(%s) << %d)", tyUint, name, operator, tyUint, fieldStr, (fieldSize-i-1)*8)
					}
				}
			}

			if fieldProcessor == nil {
				// 小端序
				fieldProcessor = func(fieldStr string, i int64) string {
					operator := "="
					if i != 0 {
						operator = "|="
					}
					return fmt.Sprintf("(*(%s*)(&(%s))) %s ((%s)(%s) << %d)", tyUint, name, operator, tyUint, fieldStr, i*8)
				}
			}

			decodeStr += genDecode(from, to, fieldProcessor)

			if val.FieldType.TypeTypeID.IsInt() && val.FieldType.TypeBitSize > val.FieldBitSize {
				decodeStr += fmt.Sprintf("%s = %s;\n", name, signExtend(val.FieldBitSize, val.FieldType.TypeBitSize, name))
			}

			decodeStr += fmt.Sprintf("if (%s != %s) return false;\n", name, val.FieldConstant)

		case *definition.VoidField:
			decodeStr += fmt.Sprintf("// %s %s\n", pos, val.ShortString())
			continue

		case *definition.EmbeddedField:
			continue

		case *definition.NormalField:
			decodeStr += fmt.Sprintf("// %s %s\n", pos, val.ShortString())
			name := fmt.Sprintf("structPtr->%s", val.FieldName)

			switch ty := val.FieldType.(type) {
			case *definition.Struct:
				decodeStr += fmt.Sprintf("if (!%s_decode((void*)((uint8_t*)data + %d), &(%s))) return false;\n", ty.StructName, from/8, name)

			case *definition.Enum:
				fieldSize := (val.FieldBitSize + 7) / 8
				tySize := util.HighBit(fieldSize)
				tyUint := typeSizeMapUint[tySize*8]
				tempName := fmt.Sprintf("temp_field_%d", fieldIndex)
				decodeStr += fmt.Sprintf("%s %s;\n", tyUint, tempName)

				var fieldProcessor func(string, int64) string
				// 默认解码为小端序
				if option, ok := val.FieldOptions.Get("order"); ok {
					if option.OptionValue.GetLiteralValue() == "big" {
						// 大端序
						fieldProcessor = func(fieldStr string, i int64) string {
							operator := "="
							if i != 0 {
								operator = "|="
							}
							return fmt.Sprintf("%s %s ((%s)(%s) << %d)", tempName, operator, tyUint, fieldStr, (fieldSize-i-1)*8)
						}
					}
				}
				if fieldProcessor == nil {
					// 小端序
					fieldProcessor = func(fieldStr string, i int64) string {
						operator := "="
						if i != 0 {
							operator = "|="
						}
						return fmt.Sprintf("%s %s ((%s)(%s) << %d)", tempName, operator, tyUint, fieldStr, i*8)
					}
				}
				decodeStr += genDecode(from, to, fieldProcessor)
				decodeStr += fmt.Sprintf("%s = (enum %s)%s;\n", name, ty.EnumName, tempName)

				// TODO: sign extend of enum. DO THIS AFTER SUPPORTING NEGATIVE ENUM VALUES
				// if basicTy, ok := ty.(*definition.BasicType); ok {
				// 	if basicTy.TypeTypeID.IsInt() && basicTy.TypeBitSize > val.FieldBitSize {
				// 		decodeStr += fmt.Sprintf("%s = %s;\n", name, signExtend(val.FieldBitSize, basicTy.TypeBitSize, name))
				// 	}
				// }

			case *definition.BasicType:
				tySize := (ty.TypeBitSize + 7) / 8
				tyUint := typeSizeMapUint[tySize*8]
				fieldSize := (val.FieldBitSize + 7) / 8
				var fieldProcessor func(string, int64) string
				// 默认解码为小端序
				if option, ok := val.FieldOptions.Get("order"); ok {
					if option.OptionValue.GetLiteralValue() == "big" {
						// 大端序
						fieldProcessor = func(fieldStr string, i int64) string {
							operator := "="
							if i != 0 {
								operator = "|="
							}
							return fmt.Sprintf("(*(%s*)(&(%s))) %s ((%s)(%s) << %d)", tyUint, name, operator, tyUint, fieldStr, (fieldSize-i-1)*8)
						}
					}
				}
				if fieldProcessor == nil {
					// 小端序
					fieldProcessor = func(fieldStr string, i int64) string {
						operator := "="
						if i != 0 {
							operator = "|="
						}
						return fmt.Sprintf("(*(%s*)(&(%s))) %s ((%s)(%s) << %d)", tyUint, name, operator, tyUint, fieldStr, i*8)
					}
				}
				decodeStr += genDecode(from, to, fieldProcessor)

				if ty.TypeTypeID.IsInt() && ty.TypeBitSize > val.FieldBitSize {
					decodeStr += fmt.Sprintf("%s = %s;\n", name, signExtend(val.FieldBitSize, ty.TypeBitSize, name))
				}

			case *definition.Array:
				elemTySize := (ty.ElementType.GetTypeBitSize() + 7) / 8
				elemBitSize := val.FieldBitSize / ty.Length
				elemSize := (elemBitSize + 7) / 8
				tyUint := typeSizeMapUint[elemTySize*8]
				var fieldProcessorIndex func(int64) func(string, int64) string
				// 默认解码为小端序
				if option, ok := val.FieldOptions.Get("order"); ok {
					if option.OptionValue.GetLiteralValue() == "big" {
						// 大端序
						fieldProcessorIndex = func(index int64) func(string, int64) string {
							return func(fieldStr string, i int64) string {
								operator := "="
								if i != 0 {
									operator = "|="
								}
								return fmt.Sprintf("(*(%s*)(&((%s)[%d]))) %s ((%s)(%s) << %d)", tyUint, name, index, operator, tyUint, fieldStr, (elemSize-i-1)*8)
							}
						}
					}
				}
				if fieldProcessorIndex == nil {
					// 小端序
					fieldProcessorIndex = func(index int64) func(string, int64) string {
						return func(fieldStr string, i int64) string {
							operator := "="
							if i != 0 {
								operator = "|="
							}
							return fmt.Sprintf("(*(%s*)(&((%s)[%d]))) %s ((%s)(%s) << %d)", tyUint, name, index, operator, tyUint, fieldStr, i*8)
						}
					}
				}
				for i := int64(0); i < ty.Length; i++ {
					subFrom := from + i*elemBitSize
					subTo := from + (i+1)*elemBitSize
					fieldProcessor := fieldProcessorIndex(i)
					decodeStr += genDecode(subFrom, subTo, fieldProcessor)
				}

				if basicTy, ok := ty.ElementType.(*definition.BasicType); ok {
					if basicTy.TypeTypeID.IsInt() && basicTy.TypeBitSize > val.FieldBitSize {
						for i := int64(0); i < ty.Length; i++ {
							elemName := fmt.Sprintf("(%s)[%d]", name, i)
							decodeStr += fmt.Sprintf("%s = %s;\n", elemName, signExtend(val.FieldBitSize, basicTy.TypeBitSize, elemName))
						}
					}
				}

			default:
				fieldProcessor := func(fieldStr string, i int64) string {
					operator := "="
					return fmt.Sprintf("((uint8_t*)&(structPtr->%s))[%d] %s %s", val.FieldName, i, operator, fieldStr)
				}
				decodeStr += genDecode(from, to, fieldProcessor)
			}

		default:
			panic("unreachable")
		}
	}
	decodeStr += "return true;\n"
	decodeStr = util.IndentSpace4(decodeStr)
	str += decodeStr
	str += "}\n"
	return str
}
