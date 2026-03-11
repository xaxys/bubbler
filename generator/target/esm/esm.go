package esm

import (
	"fmt"
	"text/template"

	"github.com/xaxys/bubbler/definition"
	"github.com/xaxys/bubbler/generator/gen"
	"github.com/xaxys/bubbler/util"
)

type GeneratedUnit struct {
	SourceUnit     *definition.CompilationUnit
	GeneratedTypes *util.OrderedMap[string, *GeneratedType]
}

type GeneratedType struct {
	GeneratedDef string
	JsComments   string
	IsClass      bool
}

// ==================== ESModule Generator ====================

type ESModuleGeneratorState struct {
	UseFloat32     bool
	UseFloat64     bool
	UseString      bool
	UseStructArray bool
}

func NewESModuleGeneratorState() *ESModuleGeneratorState {
	g := &ESModuleGeneratorState{}
	g.Reset()
	return g
}

func (g *ESModuleGeneratorState) Reset() {
	g.UseFloat32 = false
	g.UseFloat64 = false
	g.UseString = false
	g.UseStructArray = false
}

type ESModuleGenerator struct {
	*gen.GenDispatcher
	GenCtx   *gen.GenCtx
	GenUnits *util.OrderedMap[string, *GeneratedUnit]
	GenTypes *util.OrderedMap[string, *GeneratedType]
	GenStack *util.OrderedMap[string, any]
	GenState *ESModuleGeneratorState
	Warning  definition.TopLevelWarning
}

func NewESModuleGenerator() *ESModuleGenerator {
	generator := &ESModuleGenerator{
		GenDispatcher: nil,
		GenUnits:      util.NewOrderedMap[string, *GeneratedUnit](),
		GenTypes:      util.NewOrderedMap[string, *GeneratedType](),
		GenStack:      util.NewOrderedMap[string, any](),
		GenState:      NewESModuleGeneratorState(),
		Warning:       nil,
	}
	generator.GenDispatcher = gen.NewGenDispatcher(generator)
	return generator
}

// ==================== Util ====================

func generateDec(value any) string {
	return fmt.Sprintf("%d", value)
}

func generateDecBigInt(value any) string {
	return fmt.Sprintf("%sn", generateDec(value))
}

func (g *ESModuleGenerator) generateDec(value any) string {
	return generateDec(value)
}

func (g *ESModuleGenerator) generateHex(value any) string {
	if g.GenCtx.GenOptions.DecimalNumber {
		return fmt.Sprintf("%d", value)
	}
	return fmt.Sprintf("0x%X", value)
}

func (g *ESModuleGenerator) generateBin(value any) string {
	if g.GenCtx.GenOptions.DecimalNumber {
		return fmt.Sprintf("%d", value)
	}
	return fmt.Sprintf("0b%b", value)
}

func (g *ESModuleGenerator) generateDecBigInt(value any) string {
	return generateDecBigInt(value)
}

func (g *ESModuleGenerator) generateHexBigInt(value any) string {
	return fmt.Sprintf("%sn", g.generateHex(value))
}

func (g *ESModuleGenerator) generateBinBigInt(value any) string {
	return fmt.Sprintf("%sn", g.generateBin(value))
}

var structPackagePrefixTemplate = `
{{- define "structPackagePrefix" -}}
{{- $structDef := .StructDef -}}
{{- $inPackage := .GenUnit.LocalNames.Has $structDef.StructName -}}
{{- if .SingleFile -}}
    {{ $structDef.StructName }}
{{- else -}}
    {{- if $inPackage -}}
    {{ $structDef.StructName }}
    {{- else -}}
    ${{ $structDef.StructBelongs.Package.PackageName }}.{{ $structDef.StructName }}
    {{- end -}}
{{- end -}}
{{- end -}}
`

func (g *ESModuleGenerator) generateStructPackagePrefix(structDef *definition.Struct) string {
	data := map[string]any{
		"StructDef":  structDef,
		"GenUnit":    g.GenUnits.Last().Value.SourceUnit,
		"SingleFile": g.GenCtx.GenOptions.SingleFile,
	}
	return util.ExecuteTemplate(structPackagePrefixTemplate, "structPackagePrefix", nil, data)
}

// ==================== Generate ====================

func (g *ESModuleGenerator) Generate(ctx *gen.GenCtx) (retErr error, retWarnings error) {
	g.GenCtx = ctx
	if !ctx.GenOptions.RelativePath {
		warn := &definition.GenerateWarning{
			Warning: &definition.OptionNotSetWarning{
				OptionName: "relpath",
				Reason:     "ESModule target generates relative path by default, this option will be forced to be enabled",
			},
		}
		g.Warning = definition.TopLevelWarningsJoin(g.Warning, warn)
		ctx.GenOptions.RelativePath = true
	}
	if ctx.GenOptions.SignExtMethod == gen.SignExtMethodShift {
		warn := &definition.GenerateWarning{
			Warning: &definition.OptionNotAvailableWarning{
				OptionName: "signext",
				Reason:     "ESModule target does not support sign extension method 'shift', default method will be used",
			},
		}
		g.Warning = definition.TopLevelWarningsJoin(g.Warning, warn)
	}
	if ctx.GenOptions.InnerClass {
		warn := &definition.GenerateWarning{
			Warning: &definition.OptionNotAvailableWarning{
				OptionName: "inner",
				Reason:     "ESModule target does not support inner class yet, this option will be ignored",
			},
		}
		g.Warning = definition.TopLevelWarningsJoin(g.Warning, warn)
	}
	if ctx.GenOptions.MinimalCode {
		warn := &definition.GenerateWarning{
			Warning: &definition.OptionNotAvailableWarning{
				OptionName: "minimal",
				Reason:     "ESModule target does not support minimal code generation yet, this option will be ignored",
			},
		}
		g.Warning = definition.TopLevelWarningsJoin(g.Warning, warn)
	}
	if !ctx.GenOptions.MemoryCopy {
		warn := &definition.GenerateWarning{
			Warning: &definition.OptionNotSetWarning{
				OptionName: "memcpy",
				Reason:     "ESModule target does not support zero-copy, this option will be forced to be enabled",
			},
		}
		g.Warning = definition.TopLevelWarningsJoin(g.Warning, warn)
		ctx.GenOptions.MemoryCopy = true
	}
	var topErr definition.TopLevelError
	genErr := g.AcceptGenCtx(ctx)
	if genErr != nil {
		if _, ok := genErr.(definition.TopLevelError); !ok {
			genErr = &definition.GenerateError{
				Err: genErr,
			}
		}
		topErr = definition.TopLevelErrorsJoin(topErr, genErr.(definition.TopLevelError))
	}

	// generate single file
	if g.GenCtx.GenOptions.SingleFile {
		singleData := map[string]any{
			"GenUnits":   g.GenUnits,
			"GenState":   g.GenState,
			"GenOptions": g.GenCtx.GenOptions,
		}

		singleStr := util.ExecuteTemplate(fileTemplate, "singleFile", nil, singleData)
		err := g.GenCtx.WriteFile("", singleStr)
		if err != nil {
			topErr = definition.TopLevelErrorsJoin(topErr, &definition.GenerateError{
				Err: err,
			})
		}
	}

	return topErr, g.Warning
}

// ==================== GenerateUnit ====================

var fileTemplate = `
{{- define "file" -}}
{{- $curUnit := .Unit -}}
// Target: ESModule
// Generated by bubbler
// DO NOT EDIT

/* eslint-disable */
{{- range $unit := .Unit.LocalImports.Values }}
import ${{ $unit.Package.PackageName }} from "{{ $curUnit.Package.ToRelativePathStrict $unit.Package ".bb.js" }}";
{{- end }}

{{- if .GenState.UseStructArray }}

function createArray(length, init) {
    const arr = new Array(length);
    for (let i = 0; i < length; i++) {
        arr[i] = init();
    }
    return arr;
}
{{- end }}

{{- if .GenState.UseFloat32 }}

function floatToUint32Bits(value) {
    const buffer = new ArrayBuffer(4);
    const view = new DataView(buffer);
    view.setFloat32(0, value, true);
    return view.getUint32(0, true);
}

function uint32BitsToFloat(value) {
    const buffer = new ArrayBuffer(4);
    const view = new DataView(buffer);
    view.setUint32(0, value, true);
    return view.getFloat32(0, true);
}
{{- end }}

{{- if .GenState.UseFloat64 }}

function doubleToUint64Bits(value) {
    const buffer = new ArrayBuffer(8);
    const view = new DataView(buffer);
    view.setFloat64(0, value, true);
    return view.getBigUint64(0, true);
}

function uint64BitsToDouble(value) {
    const buffer = new ArrayBuffer(8);
    const view = new DataView(buffer);
    view.setBigUint64(0, value, true);
    return view.getFloat64(0, true);
}
{{- end }}

{{- if .GenState.UseString }}

function stringToUTF8BytesCount(str) {
    let count = 0;
    for (let i = 0; i < str.length; i++) {
        let code = str.charCodeAt(i);
        if (code < 0x80) count += 1;
        else if (code < 0x800) count += 2;
        else if (code < 0xD800 || code >= 0xE000) count += 3;
        else { count += 4; i++; }
    }
    return count;
}

function stringToUTF8Bytes(str, data, start) {
    let offset = start;
    for (let i = 0; i < str.length; i++) {
        let code = str.charCodeAt(i);
        if (code < 0x80) {
            data[offset++] = code;
        } else if (code < 0x800) {
            data[offset++] = 0xC0 | (code >> 6);
            data[offset++] = 0x80 | (code & 0x3F);
        } else if (code < 0xD800 || code >= 0xE000) {
            data[offset++] = 0xE0 | (code >> 12);
            data[offset++] = 0x80 | ((code >> 6) & 0x3F);
            data[offset++] = 0x80 | (code & 0x3F);
        } else {
            i++;
            code = 0x10000 + (((code & 0x3FF) << 10) | (str.charCodeAt(i) & 0x3FF));
            data[offset++] = 0xF0 | (code >> 18);
            data[offset++] = 0x80 | ((code >> 12) & 0x3F);
            data[offset++] = 0x80 | ((code >> 6) & 0x3F);
            data[offset++] = 0x80 | (code & 0x3F);
        }
    }
    return offset - start;
}

function stringFromUTF8Bytes(data, start) {
    let str = "";
    let offset = start;
    while (offset < data.length && data[offset] && data[offset] !== 0) {
        let code = data[offset++];
        if (code < 0x80) str += String.fromCharCode(code);
        else if (code < 0xE0) str += String.fromCharCode(((code & 0x1F) << 6) | (data[offset++] & 0x3F));
        else if (code < 0xF0) str += String.fromCharCode(((code & 0xF) << 12) | ((data[offset++] & 0x3F) << 6) | (data[offset++] & 0x3F));
        else {
            code = ((code & 0x7) << 18) | ((data[offset++] & 0x3F) << 12) | ((data[offset++] & 0x3F) << 6) | (data[offset++] & 0x3F);
            if (code < 0x10000) str += String.fromCharCode(code);
            else {
                code -= 0x10000;
                str += String.fromCharCode(0xD800 + (code >> 10), 0xDC00 + (code & 0x3FF));
            }
        }
    }
    return [str, offset - start];
}
{{- end }}

{{ range $entry := .GenTypes.Entries -}}
{{- $type := $entry.Value -}}
// ====================== {{ $entry.Key }} ======================
{{ if $type.GeneratedDef }}
{{- if $type.JsComments }}
{{ $type.JsComments }}
{{- end }}
{{- if and $type.IsClass (not $.GenOptions.CompatibleMode) }}
{{ $type.GeneratedDef }}
{{- else }}
const {{ $entry.Key }} = (function() {
{{ indent 4 $type.GeneratedDef }}
{{- if $type.IsClass }}
    return {{ $entry.Key }};
{{- end }}
})();
{{- end }}
{{ end }}
// ==================== End {{ $entry.Key }} ====================

{{ end -}}
export default {
{{- range $entry := .GenTypes.Entries }}
{{- if $entry.Value.GeneratedDef }}
	{{ $entry.Key }},
{{- end }}
{{- end }}
};
{{- end -}}

{{- define "singleFile" -}}
// Target: ESModule (single file)
// Generated by bubbler
// DO NOT EDIT

/* eslint-disable */

{{- if .GenState.UseStructArray }}

function createArray(length, init) {
    const arr = new Array(length);
    for (let i = 0; i < length; i++) {
        arr[i] = init();
    }
    return arr;
}
{{- end }}

{{- if .GenState.UseFloat32 }}

function floatToUint32Bits(value) {
    const buffer = new ArrayBuffer(4);
    const view = new DataView(buffer);
    view.setFloat32(0, value, true);
    return view.getUint32(0, true);
}

function uint32BitsToFloat(value) {
    const buffer = new ArrayBuffer(4);
    const view = new DataView(buffer);
    view.setUint32(0, value, true);
    return view.getFloat32(0, true);
}
{{- end }}

{{- if .GenState.UseFloat64 }}

function doubleToUint64Bits(value) {
    const buffer = new ArrayBuffer(8);
    const view = new DataView(buffer);
    view.setFloat64(0, value, true);
    return view.getBigUint64(0, true);
}

function uint64BitsToDouble(value) {
    const buffer = new ArrayBuffer(8);
    const view = new DataView(buffer);
    view.setBigUint64(0, value, true);
    return view.getFloat64(0, true);
}
{{- end }}

{{- if .GenState.UseString }}

function stringToUTF8BytesCount(str) {
    let count = 0;
    for (let i = 0; i < str.length; i++) {
        let code = str.charCodeAt(i);
        if (code < 0x80) count += 1;
        else if (code < 0x800) count += 2;
        else if (code < 0xD800 || code >= 0xE000) count += 3;
        else { count += 4; i++; }
    }
    return count;
}

function stringToUTF8Bytes(str, data, start) {
    let offset = start;
    for (let i = 0; i < str.length; i++) {
        let code = str.charCodeAt(i);
        if (code < 0x80) {
            data[offset++] = code;
        } else if (code < 0x800) {
            data[offset++] = 0xC0 | (code >> 6);
            data[offset++] = 0x80 | (code & 0x3F);
        } else if (code < 0xD800 || code >= 0xE000) {
            data[offset++] = 0xE0 | (code >> 12);
            data[offset++] = 0x80 | ((code >> 6) & 0x3F);
            data[offset++] = 0x80 | (code & 0x3F);
        } else {
            i++;
            code = 0x10000 + (((code & 0x3FF) << 10) | (str.charCodeAt(i) & 0x3FF));
            data[offset++] = 0xF0 | (code >> 18);
            data[offset++] = 0x80 | ((code >> 12) & 0x3F);
            data[offset++] = 0x80 | ((code >> 6) & 0x3F);
            data[offset++] = 0x80 | (code & 0x3F);
        }
    }
    return offset - start;
}

function stringFromUTF8Bytes(data, start) {
    let str = "";
    let offset = start;
    while (offset < data.length && data[offset] && data[offset] !== 0) {
        let code = data[offset++];
        if (code < 0x80) str += String.fromCharCode(code);
        else if (code < 0xE0) str += String.fromCharCode(((code & 0x1F) << 6) | (data[offset++] & 0x3F));
        else if (code < 0xF0) str += String.fromCharCode(((code & 0xF) << 12) | ((data[offset++] & 0x3F) << 6) | (data[offset++] & 0x3F));
        else {
            code = ((code & 0x7) << 18) | ((data[offset++] & 0x3F) << 12) | ((data[offset++] & 0x3F) << 6) | (data[offset++] & 0x3F);
            if (code < 0x10000) str += String.fromCharCode(code);
            else {
                code -= 0x10000;
                str += String.fromCharCode(0xD800 + (code >> 10), 0xDC00 + (code & 0x3FF));
            }
        }
    }
    return [str, offset - start];
}
{{- end }}

{{ range $genUnit := .GenUnits.Values -}}
{{- $curUnit := $genUnit.SourceUnit -}}
{{- $genTypes := $genUnit.GeneratedTypes -}}
// ====================== {{ $curUnit.Package }} ======================

{{ range $entry := $genTypes.Entries -}}
{{- $type := $entry.Value -}}
// ====================== {{ $entry.Key }} ======================
{{ if $type.GeneratedDef }}
{{- if $type.JsComments }}
{{ $type.JsComments }}
{{- end }}
{{- if and $type.IsClass (not $.GenOptions.CompatibleMode) }}
{{ $type.GeneratedDef }}
{{- else }}
const {{ $entry.Key }} = (function() {
{{ indent 4 $type.GeneratedDef }}
{{- if $type.IsClass }}
    return {{ $entry.Key }};
{{- end }}
})();
{{- end }}
{{ end }}
// ==================== End {{ $entry.Key }} ====================

{{ end -}}
// ==================== End {{ $curUnit.Package }} ====================


{{ end -}}
export default {
{{ range $genUnit := .GenUnits.Values }}
	{{ $genUnit.SourceUnit.Package.PackageName }}: {
{{- range $entry := $genUnit.GeneratedTypes.Entries }}
{{- if $entry.Value.GeneratedDef }}
        {{ $entry.Key }},
{{- end }}
{{- end }}
    },
{{ end -}}
};
{{- end -}}
`

func (g ESModuleGenerator) GenerateUnit(unit *definition.CompilationUnit) error {
	if unit.LocalTypes.Len() == 0 && gen.MatchOption(unit.Options, "omit_empty", true) {
		return nil
	}

	genUnit := &GeneratedUnit{
		SourceUnit:     unit,
		GeneratedTypes: util.NewOrderedMap[string, *GeneratedType](),
	}

	g.GenUnits.Put(unit.Package.String(), genUnit)

	start := g.GenTypes.Len()

	for _, type_ := range unit.LocalTypes.Values() {
		_, err := g.GenerateType(type_)
		if err != nil {
			return err
		}
	}

	end := g.GenTypes.Len()

	genTypes := g.GenTypes.Sub(start, end)
	genUnit.GeneratedTypes = genTypes

	// do not generate file if single file
	if g.GenCtx.GenOptions.SingleFile {
		return nil
	}

	fileData := map[string]any{
		"Unit":       unit,
		"GenTypes":   genTypes,
		"GenState":   g.GenState,
		"GenOptions": g.GenCtx.GenOptions,
	}

	fileStr := util.ExecuteTemplate(fileTemplate, "file", nil, fileData)
	err := g.GenCtx.WritePackage(unit.Package, ".bb.js", fileStr)
	if err != nil {
		return err
	}

	// clear state for next unit
	g.GenState.Reset()

	return nil
}

func (g ESModuleGenerator) GenerateType(type_ definition.Type) (string, error) {
	return g.AcceptType(type_)
}

func (g ESModuleGenerator) GenerateTypeDefaultValue(type_ definition.Type) (string, error) {
	return g.AcceptTypeDefaultValue(type_)
}

var typeMap = map[definition.TypeID]string{
	definition.TypeID_Bool:    "Boolean",
	definition.TypeID_Uint8:   "Number",
	definition.TypeID_Uint16:  "Number",
	definition.TypeID_Uint32:  "Number",
	definition.TypeID_Uint64:  "BigInt",
	definition.TypeID_Int8:    "Number",
	definition.TypeID_Int16:   "Number",
	definition.TypeID_Int32:   "Number",
	definition.TypeID_Int64:   "BigInt",
	definition.TypeID_Float32: "Number",
	definition.TypeID_Float64: "Number",
	definition.TypeID_String:  "String",
	definition.TypeID_Bytes:   "Array",
}

func (g ESModuleGenerator) GenerateBasicType(type_ *definition.BasicType) (string, error) {
	switch type_.TypeTypeID {
	case definition.TypeID_Float32:
		g.GenState.UseFloat32 = true
	case definition.TypeID_Float64:
		g.GenState.UseFloat64 = true
	}
	if str, ok := typeMap[type_.TypeTypeID]; ok {
		return str, nil
	}
	return "", fmt.Errorf("unknown basic type: %s", type_.String())
}

var typeDefValueMap = map[definition.TypeID]string{
	definition.TypeID_Bool:    "false",
	definition.TypeID_Uint8:   "0",
	definition.TypeID_Uint16:  "0",
	definition.TypeID_Uint32:  "0",
	definition.TypeID_Uint64:  "0n",
	definition.TypeID_Int8:    "0",
	definition.TypeID_Int16:   "0",
	definition.TypeID_Int32:   "0",
	definition.TypeID_Int64:   "0n",
	definition.TypeID_Float32: "0",
	definition.TypeID_Float64: "0",
}

func (g ESModuleGenerator) GenerateBasicTypeDefaultValue(type_ *definition.BasicType) (string, error) {
	switch type_.TypeTypeID {
	case definition.TypeID_Float32:
		g.GenState.UseFloat32 = true
	case definition.TypeID_Float64:
		g.GenState.UseFloat64 = true
	}
	if str, ok := typeDefValueMap[type_.TypeTypeID]; ok {
		return str, nil
	}
	return "", fmt.Errorf("unknown basic type: %s", type_.String())
}

func (g ESModuleGenerator) GenerateString(string_ *definition.String) (string, error) {
	g.GenState.UseString = true
	return "String", nil
}

func (g ESModuleGenerator) GenerateStringDefaultValue(string_ *definition.String) (string, error) {
	g.GenState.UseString = true
	return `""`, nil
}

func (g ESModuleGenerator) GenerateBytes(bytes *definition.Bytes) (string, error) {
	if g.GenCtx.GenOptions.CompatibleMode {
		return "Array", nil
	}
	return "Uint8Array", nil
}

func (g ESModuleGenerator) GenerateBytesDefaultValue(bytes *definition.Bytes) (string, error) {
	if g.GenCtx.GenOptions.CompatibleMode {
		return "[]", nil
	}
	return "new Uint8Array(0)", nil
}

func (g ESModuleGenerator) GenerateArray(array *definition.Array) (string, error) {
	_, err := g.GenerateType(array.ElementType)
	if err != nil {
		return "", err
	}
	return "Array", nil
}

func (g ESModuleGenerator) GenerateArrayDefaultValue(array *definition.Array) (string, error) {
	if array.ElementType.GetTypeID().IsStruct() {
		g.GenState.UseStructArray = true
		elemDefValue, err := g.GenerateTypeDefaultValue(array.ElementType)
		if err != nil {
			return "", err
		}
		length := g.generateDec(array.Length)
		return fmt.Sprintf("createArray(%s, function() { return %s; })", length, elemDefValue), nil
	}
	return "[]", nil
}

var structTemplate = `
{{- define "field" -}}
        // {{ .Pos }} {{ .Field.GetFieldKind }}: {{ .Field }}
        {{- $fieldStrLen := len .FieldStr -}}
        {{- if gt $fieldStrLen 0 }}
        {{ .FieldStr }}
        {{- end }}
{{- end -}}

{{- define "methodGroupDef" -}}
{{- range $method := .MethodGroupDef.Values }}
    {{ GenerateMethod $method }}
{{- end }}
{{- end }}

{{- define "structDef" -}}
{{- $structDef := .StructDef -}}
{{- $fieldStrs := .FieldStrs -}}
/**
 * @memberof {{ $structDef.StructBelongs.Package }}
 * @constructor
 */
class {{ $structDef.StructName }} {
    constructor(properties) {
        {{- range $fieldStr := $fieldStrs }}
        {{ $fieldStr }}
        {{- end }}
        if (properties) Object.assign(this, properties);
    }

    static create(properties) {
        return new {{ $structDef.StructName }}(properties);
    }

    static dynamic() {
        return {{ if $structDef.StructDynamic }}true{{ else }}false{{ end }};
    }

    static size() {
        return {{ calc $structDef.StructBitSize "/" 8 }};
    }
    {{- range $methodStr := .MethodStrs }}
	{{ $methodStr }}
    {{- end }}

    {{ indentspace 4 .EncoderStr }}

    {{ indentspace 4 .DecoderStr }}
}
{{- end -}}
`

func (g ESModuleGenerator) GenerateStruct(structDef *definition.Struct) (string, error) {
	name := structDef.StructName
	if g.GenTypes.Has(structDef.StructName) {
		return name, nil
	}
	if g.GenStack.Has(structDef.StructName) {
		return name, nil
	}
	g.GenStack.Put(structDef.StructName, nil)
	defer g.GenStack.Remove(structDef.StructName)

	genTy, err := g.generateStruct(structDef)
	if err != nil {
		return "", err
	}
	g.GenTypes.Put(structDef.StructName, genTy)
	return name, nil
}

func (g ESModuleGenerator) generateStruct(structDef *definition.Struct) (*GeneratedType, error) {
	funcMap := template.FuncMap{
		"GenerateMethod": g.GenerateMethod,
		"GenerateType":   g.GenerateType,
	}

	fieldInitStrs := make([]string, structDef.StructFields.Len())
	if err := structDef.ForEachFieldWithPos(func(field definition.Field, index int, start int64, dynamic bool, pos string) error {
		fieldStr, err := g.GenerateField(field)
		if err != nil {
			return err
		}
		fieldData := map[string]any{
			"Pos":      pos,
			"Field":    field,
			"FieldStr": fieldStr,
		}
		str := util.ExecuteTemplate(structTemplate, "field", funcMap, fieldData)
		fieldInitStrs[index] = str
		return nil
	}); err != nil {
		return nil, err
	}

	methodStrs := []string{}
	if err := structDef.ForEachField(func(field definition.Field, index int, start int64, dynamic bool) error {
		if !field.GetFieldKind().IsNormal() {
			return nil
		}
		normalField := field.(*definition.NormalField)
		for _, group := range normalField.FieldMethods.Values() {
			methodGroupDefData := map[string]any{
				"MethodGroupDef": group,
			}
			methodGroupStr := util.ExecuteTemplate(structTemplate, "methodGroupDef", funcMap, methodGroupDefData)
			methodStrs = append(methodStrs, methodGroupStr)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	encoderStr, err := g.GenerateEncoder(structDef)
	if err != nil {
		return nil, err
	}
	decoderStr, err := g.GenerateDecoder(structDef)
	if err != nil {
		return nil, err
	}

	defData := map[string]any{
		"StructDef":  structDef,
		"FieldStrs":  fieldInitStrs,
		"MethodStrs": methodStrs,
		"EncoderStr": encoderStr,
		"DecoderStr": decoderStr,
	}
	defStr := util.ExecuteTemplate(structTemplate, "structDef", funcMap, defData)
	code := &GeneratedType{
		GeneratedDef: defStr,
		IsClass:      true,
	}
	return code, nil
}

func (g *ESModuleGenerator) GenerateStructDefaultValue(structDef *definition.Struct) (string, error) {
	prefix := g.generateStructPackagePrefix(structDef)
	return fmt.Sprintf("%s.create()", prefix), nil
}

func (g ESModuleGenerator) GenerateField(field definition.Field) (string, error) {
	return g.AcceptField(field)
}

var normalFieldTemplate = `
{{- define "normalField" -}}
    {{- $field := .Field -}}
    {{- $fieldName := TocamelCase .Field.FieldName -}}
    {{- $tyDefValue := GenerateTypeDefaultValue $field.FieldType -}}
    this.{{ $fieldName }} = {{ $tyDefValue }};
{{- end -}}
`

func (g ESModuleGenerator) GenerateNormalField(field *definition.NormalField) (string, error) {
	funcMap := template.FuncMap{
		"GenerateType":             g.GenerateType,
		"GenerateTypeDefaultValue": g.GenerateTypeDefaultValue,
	}
	fieldData := map[string]any{"Field": field}
	return util.ExecuteTemplate(normalFieldTemplate, "normalField", funcMap, fieldData), nil
}

func (g ESModuleGenerator) GenerateVoidField(field *definition.VoidField) (string, error) {
	return "", nil
}

func (g ESModuleGenerator) GenerateEmbeddedField(field *definition.EmbeddedField) (string, error) {
	return "", nil
}

var constantFieldTemplate = `
{{- define "constantField" -}}
    {{- $field := .Field -}}
    {{- $fieldName := TocamelCase .Field.FieldName -}}
    {{- $constValue := GenerateLiteral $field.FieldConstant -}}
    this.{{ $fieldName }} = {{ $constValue }};
{{- end -}}
`

func (g ESModuleGenerator) GenerateConstantField(field *definition.ConstantField) (string, error) {
	literalGentor := NewESModuleLiteralGenerator()
	funcMap := template.FuncMap{
		"GenerateType":    g.GenerateType,
		"GenerateLiteral": literalGentor.GenerateLiteral,
	}
	fieldData := map[string]any{"Field": field}
	return util.ExecuteTemplate(constantFieldTemplate, "constantField", funcMap, fieldData), nil
}

func (g ESModuleGenerator) GenerateMethod(method definition.Method) (string, error) {
	return g.AcceptMethod(method)
}

func (g ESModuleGenerator) GenerateMethodDecl(method definition.Method) (string, error) {
	return g.AcceptMethodDecl(method)
}

func (g ESModuleGenerator) GenerateDefaultGetterDecl(method *definition.GetMethod) (string, error) {
	panic("not implemented")
}
func (g ESModuleGenerator) GenerateDefaultSetterDecl(method *definition.SetMethod) (string, error) {
	panic("not implemented")
}
func (g ESModuleGenerator) GenerateCustomGetterDecl(method *definition.GetMethod) (string, error) {
	panic("not implemented")
}
func (g ESModuleGenerator) GenerateCustomSetterDecl(method *definition.SetMethod) (string, error) {
	panic("not implemented")
}
func (g ESModuleGenerator) GenerateRawGetterDecl(field definition.Field) (string, error) {
	panic("not implemented")
}
func (g ESModuleGenerator) GenerateRawSetterDecl(field definition.Field) (string, error) {
	panic("not implemented")
}
func (g ESModuleGenerator) GenerateDefaultGetter(method *definition.GetMethod) (string, error) {
	panic("not implemented")
}
func (g ESModuleGenerator) GenerateDefaultSetter(method *definition.SetMethod) (string, error) {
	panic("not implemented")
}

var customGetterTemplate = `
{{- define "customGetter" -}}
{{- $field := .MethodDef.MethodBelongs -}}
{{- $fieldName := TocamelCase $field.FieldName -}}
{{- $methodName := TocamelCase .MethodDef.MethodName -}}
{{- $valueStr := printf "this.%s" $fieldName -}}
{{- $exprStr := GenerateExpr .MethodDef.MethodExpr $valueStr -}}
        get {{ $methodName }}() { return {{ $exprStr }}; }
{{- end -}}
`

func (g ESModuleGenerator) GenerateCustomGetter(method *definition.GetMethod) (string, error) {
	funcMap := template.FuncMap{
		"GenerateType": g.GenerateType,
		"GenerateExpr": g.GenerateExpr,
	}
	fieldData := map[string]any{
		"MethodDef":  method,
		"GenOptions": g.GenCtx.GenOptions,
	}
	return util.ExecuteTemplate(customGetterTemplate, "customGetter", funcMap, fieldData), nil
}

var customSetterTemplate = `
{{- define "customSetter" -}}
{{- $field := .MethodDef.MethodBelongs -}}
{{- $fieldName := TocamelCase $field.FieldName -}}
{{- $methodName := TocamelCase .MethodDef.MethodName -}}
{{- $exprStr := GenerateExpr .MethodDef.MethodExpr "value" -}}
        set {{ $methodName }}(value) { this.{{ $fieldName }} = {{ $exprStr }}; }
{{- end -}}
`

func (g ESModuleGenerator) GenerateCustomSetter(method *definition.SetMethod) (string, error) {
	funcMap := template.FuncMap{
		"GenerateType": g.GenerateType,
		"GenerateExpr": g.GenerateExpr,
	}
	fieldData := map[string]any{
		"MethodDef":  method,
		"GenOptions": g.GenCtx.GenOptions,
	}
	return util.ExecuteTemplate(customSetterTemplate, "customSetter", funcMap, fieldData), nil
}

func (g ESModuleGenerator) GenerateRawGetter(field definition.Field) (string, error) {
	panic("not implemented")
}
func (g ESModuleGenerator) GenerateRawSetter(field definition.Field) (string, error) {
	panic("not implemented")
}

func (g ESModuleGenerator) GenerateEnum(enumDef *definition.Enum) (string, error) {
	name := enumDef.EnumName
	if g.GenTypes.Has(enumDef.EnumName) {
		return name, nil
	}
	if g.GenStack.Has(enumDef.EnumName) {
		return name, nil
	}
	g.GenStack.Put(enumDef.EnumName, nil)
	defer g.GenStack.Remove(enumDef.EnumName)
	genTy, err := g.generateEnum(enumDef)
	if err != nil {
		return "", err
	}
	g.GenTypes.Put(enumDef.EnumName, genTy)
	return name, nil
}

var enumTemplate = `
{{- define "enumDef" -}}
const valuesById = {}, values = Object.create(valuesById);
{{- range .EnumDef.EnumValues.Values }}
values[valuesById[{{ .EnumValue }}] = "{{ .EnumValueName }}"] = {{ .EnumValue }};
{{- end }}
return values;
{{- end -}}
`

func (g ESModuleGenerator) generateEnum(enumDef *definition.Enum) (*GeneratedType, error) {
	enumDefData := map[string]any{"EnumDef": enumDef}
	enumDefStr := util.ExecuteTemplate(enumTemplate, "enumDef", nil, enumDefData)
	code := &GeneratedType{GeneratedDef: enumDefStr}
	return code, nil
}

func (g ESModuleGenerator) GenerateEnumDefaultValue(enumDef *definition.Enum) (string, error) {
	return "null", nil
}
