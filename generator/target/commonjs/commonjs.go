package commonjs

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
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
}

// ==================== CommonJS Generator ====================

type CommonJSGeneratorState struct {
	UseFloat32     bool
	UseFloat64     bool
	UseString      bool
	UseStructArray bool
}

func NewCommonJSGeneratorState() *CommonJSGeneratorState {
	return &CommonJSGeneratorState{
		UseFloat32:     false,
		UseFloat64:     false,
		UseString:      false,
		UseStructArray: false,
	}
}

type CommonJSGenerator struct {
	*gen.GenDispatcher
	GenCtx   *gen.GenCtx
	GenUnits *util.OrderedMap[string, *GeneratedUnit]
	GenTypes *util.OrderedMap[string, *GeneratedType]
	GenStack *util.OrderedMap[string, any]
	GenState *CommonJSGeneratorState
	Warning  definition.TopLevelWarning
}

func NewCommonJSGenerator() *CommonJSGenerator {
	generator := &CommonJSGenerator{
		GenDispatcher: nil,
		GenUnits:      util.NewOrderedMap[string, *GeneratedUnit](),
		GenTypes:      util.NewOrderedMap[string, *GeneratedType](),
		GenStack:      util.NewOrderedMap[string, any](),
		GenState:      NewCommonJSGeneratorState(),
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

func (g *CommonJSGenerator) generateDec(value any) string {
	return generateDec(value)
}

func (g *CommonJSGenerator) generateHex(value any) string {
	if g.GenCtx.GenOptions.DecimalNumber {
		return fmt.Sprintf("%d", value)
	}
	return fmt.Sprintf("0x%X", value)
}

func (g *CommonJSGenerator) generateBin(value any) string {
	if g.GenCtx.GenOptions.DecimalNumber {
		return fmt.Sprintf("%d", value)
	}
	return fmt.Sprintf("0b%b", value)
}

func (g *CommonJSGenerator) generateDecBigInt(value any) string {
	return generateDecBigInt(value)
}

func (g *CommonJSGenerator) generateHexBigInt(value any) string {
	return fmt.Sprintf("%sn", g.generateHex(value))
}

func (g *CommonJSGenerator) generateBinBigInt(value any) string {
	return fmt.Sprintf("%sn", g.generateBin(value))
}

var structPackagePrefixTemplate = `
{{- define "structPackagePrefix" -}}
{{- $structDef := .StructDef -}}
{{- $inPackage := .GenUnit.LocalNames.Has $structDef.StructName -}}
{{- if $inPackage -}}
    $package.{{ $structDef.StructName }}
{{- else -}}
    $root.{{ $structDef.StructBelongs.Package }}.{{ $structDef.StructName }}
{{- end -}}
{{- end -}}
`

func (g *CommonJSGenerator) generateStructPackagePrefix(structDef *definition.Struct) string {
	data := map[string]any{
		"StructDef": structDef,
		"GenUnit":   g.GenUnits.Last().Value.SourceUnit,
	}
	return util.ExecuteTemplate(structPackagePrefixTemplate, "structPackagePrefix", nil, data)
}

// ==================== Generate ====================

func (g *CommonJSGenerator) Generate(ctx *gen.GenCtx) (retErr error, retWarnings error) {
	g.GenCtx = ctx
	if ctx.GenOptions.SignExtMethod == gen.SignExtMethodShift {
		warn := &definition.GenerateWarning{
			Warning: &definition.OptionNotAvailableWarning{
				OptionName: "signext",
				Reason:     "CommonJS target does not support sign extension method 'shift', default method will be used",
			},
		}
		g.Warning = definition.TopLevelWarningsJoin(g.Warning, warn)
	}
	if ctx.GenOptions.InnerClass {
		warn := &definition.GenerateWarning{
			Warning: &definition.OptionNotAvailableWarning{
				OptionName: "inner",
				Reason:     "CommonJS target does not support inner class yet, this option will be ignored",
			},
		}
		g.Warning = definition.TopLevelWarningsJoin(g.Warning, warn)
	}
	if ctx.GenOptions.MinimalCode {
		warn := &definition.GenerateWarning{
			Warning: &definition.OptionNotAvailableWarning{
				OptionName: "minimal",
				Reason:     "CommonJS target does not support minimal code generation yet, this option will be ignored",
			},
		}
		g.Warning = definition.TopLevelWarningsJoin(g.Warning, warn)
	}
	if !ctx.GenOptions.MemoryCopy {
		warn := &definition.GenerateWarning{
			Warning: &definition.OptionNotSetWarning{
				OptionName: "memcpy",
				Reason:     "CommonJS target does not support zero-copy, this option will be forced to be enabled",
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
			"GenUnits": g.GenUnits,
			"GenState": g.GenState,
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
// Target: CommonJS
// Generated by bubbler
// DO NOT EDIT

/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
(function (global, factory) {

    var isObj = (function (item) {
        return (item && typeof item === "object" && !Array.isArray(item));
    });

    var mergeDeep = (function (target, ...sources) {
        if (!sources.length) return target;
        var source = sources.shift();
        if (isObj(target) && isObj(source)) {
            for (var key in source) {
                if (isObj(source[key])) {
                    if (!target[key]) Object.assign(target, { [key]: {} });
                    mergeDeep(target[key], source[key]);
                } else {
                    Object.assign(target, { [key]: source[key] });
                }
            }
        }
        return mergeDeep(target, ...sources);
    });

    $root = mergeDeep({},
        {{- range $unit := .Unit.LocalImports.Values }}
        require("{{ $curUnit.Package.ToRelativePathStrict $unit.Package ".bb" }}"),
        {{- end }}
    );

    module.exports = factory($root);

})(this, function ($root) {
    "use strict";

    function isObj(item) {
        return (item && typeof item === "object" && !Array.isArray(item));
    }

    function mergeDeep(target, ...sources) {
        if (!sources.length) return target;
        var source = sources.shift();
        if (isObj(target) && isObj(source)) {
            for (var key in source) {
                if (isObj(source[key])) {
                    if (!target[key]) Object.assign(target, { [key]: {} });
                    mergeDeep(target[key], source[key]);
                } else {
                    Object.assign(target, { [key]: source[key] });
                }
            }
        }
        return mergeDeep(target, ...sources);
    }

    {{- if .GenState.UseStructArray }}

    function createArray(length, init) {
        var arr = new Array(length);
        for (var i = 0; i < length; i++) {
            arr[i] = init();
        }
        return arr;
    }
    {{- end }}

    {{- if .GenState.UseFloat32 }}

    function floatToUint32Bits(value) {
        var buffer = new ArrayBuffer(4);
        var view = new DataView(buffer);
        view.setFloat32(0, value, true); // true for little-endian
        return view.getUint32(0, true);
    }

    function uint32BitsToFloat(value) {
        var buffer = new ArrayBuffer(4);
        var view = new DataView(buffer);
        view.setUint32(0, value, true); // true for little-endian
        return view.getFloat32(0, true);
    }
    {{- end }}

    {{- if .GenState.UseFloat64 }}

    function doubleToUint64Bits(value) {
        var buffer = new ArrayBuffer(8);
        var view = new DataView(buffer);
        view.setFloat64(0, value, true); // true for little-endian
        return view.getBigUint64(0, true);
    }

    function uint64BitsToDouble(value) {
        var buffer = new ArrayBuffer(8);
        var view = new DataView(buffer);
        view.setBigUint64(0, value, true); // true for little-endian
        return view.getFloat64(0, true);
    }
    {{- end }}

    {{- if .GenState.UseString }}

    function stringToUTF8BytesCount(str) {
        var count = 0;
        for (var i = 0; i < str.length; i++) {
            var code = str.charCodeAt(i);
            if (code < 0x80) count += 1;
            else if (code < 0x800) count += 2;
            else if (code < 0xD800 || code >= 0xE000) count += 3;
            else { count += 4; i++; }
        }
        return count;
    }

    function stringToUTF8Bytes(str, data, start) {
        var offset = start;
        for (var i = 0; i < str.length; i++) {
            var code = str.charCodeAt(i);
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
        var str = "";
        var offset = start;
        while (offset < data.length && data[offset] && data[offset] !== 0) {
            var code = data[offset++];
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
    
    // ====================== {{ $curUnit.Package }} ======================

    (function () {

        var root = {};
        var namespace = root;
        {{- range $pkgPart := $curUnit.Package.PackagePath }}
        namespace.{{ $pkgPart }} = {};
        namespace = namespace.{{ $pkgPart }};
        {{- end }}

        namespace.{{ $curUnit.Package.PackageName }} = (function () {
            var $package = {};

            {{ range $entry := .GenTypes.Entries -}}
            {{- $type := $entry.Value -}}
            // ====================== {{ $entry.Key }} ======================
            {{ if $type.GeneratedDef }}
            $package.{{ $entry.Key }} = (function() {
                {{ $type.GeneratedDef }}
            })();
            {{ end }}
            // ==================== End {{ $entry.Key }} ====================
    
            {{ end -}}
            return $package;
        })();

        $root = mergeDeep($root, root);
    })();
    // ==================== End {{ $curUnit.Package }} ====================

    return $root;
});
{{- end -}}

{{- define "singleFile" -}}
// Target: CommonJS (single file)
// Generated by bubbler
// DO NOT EDIT

/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
(function (global, factory) {

    $root = {};
    module.exports = factory($root);

})(this, function ($root) {
    "use strict";

    function isObj(item) {
        return (item && typeof item === "object" && !Array.isArray(item));
    }

    function mergeDeep(target, ...sources) {
        if (!sources.length) return target;
        var source = sources.shift();
        if (isObj(target) && isObj(source)) {
            for (var key in source) {
                if (isObj(source[key])) {
                    if (!target[key]) Object.assign(target, { [key]: {} });
                    mergeDeep(target[key], source[key]);
                } else {
                    Object.assign(target, { [key]: source[key] });
                }
            }
        }
        return mergeDeep(target, ...sources);
    }

    {{- if .GenState.UseStructArray }}

    function createArray(length, init) {
        var arr = new Array(length);
        for (var i = 0; i < length; i++) {
            arr[i] = init();
        }
        return arr;
    }
    {{- end }}

    {{- if .GenState.UseFloat32 }}

    function floatToUint32Bits(value) {
        var buffer = new ArrayBuffer(4);
        var view = new DataView(buffer);
        view.setFloat32(0, value, true); // true for little-endian
        return view.getUint32(0, true);
    }

    function uint32BitsToFloat(value) {
        var buffer = new ArrayBuffer(4);
        var view = new DataView(buffer);
        view.setUint32(0, value, true); // true for little-endian
        return view.getFloat32(0, true);
    }
    {{- end }}

    {{- if .GenState.UseFloat64 }}

    function doubleToUint64Bits(value) {
        var buffer = new ArrayBuffer(8);
        var view = new DataView(buffer);
        view.setFloat64(0, value, true); // true for little-endian
        return view.getBigUint64(0, true);
    }

    function uint64BitsToDouble(value) {
        var buffer = new ArrayBuffer(8);
        var view = new DataView(buffer);
        view.setBigUint64(0, value, true); // true for little-endian
        return view.getFloat64(0, true);
    }
    {{- end }}

    {{- if .GenState.UseString }}

    function stringToUTF8BytesCount(str) {
        var count = 0;
        for (var i = 0; i < str.length; i++) {
            var code = str.charCodeAt(i);
            if (code < 0x80) count += 1;
            else if (code < 0x800) count += 2;
            else if (code < 0xD800 || code >= 0xE000) count += 3;
            else { count += 4; i++; }
        }
        return count;
    }

    function stringToUTF8Bytes(str, data, start) {
        var offset = start;
        for (var i = 0; i < str.length; i++) {
            var code = str.charCodeAt(i);
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
        var str = "";
        var offset = start;
        while (offset < data.length && data[offset] && data[offset] !== 0) {
            var code = data[offset++];
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

    (function () {

        var root = {};
        var namespace = root;
        {{- range $pkgPart := $curUnit.Package.PackagePath }}
        namespace.{{ $pkgPart }} = {};
        namespace = namespace.{{ $pkgPart }};
        {{- end }}

        namespace.{{ $curUnit.Package.PackageName }} = (function () {
            var $package = {};

            {{ range $entry := $genTypes.Entries -}}
            {{- $type := $entry.Value -}}
            // ====================== {{ $entry.Key }} ======================
            {{ if $type.GeneratedDef }}
            {{- if $type.JsComments }}
            {{ $type.JsComments }}
            {{- end }}
            $package.{{ $entry.Key }} = (function() {
                {{ $type.GeneratedDef }}
            })();
            {{ end }}
            // ==================== End {{ $entry.Key }} ====================

            {{ end -}}
            return $package;
        })();

        $root = mergeDeep($root, root);
    })();

    // ==================== End {{ $curUnit.Package }} ====================

    {{ end -}}
    return $root;
});
{{- end -}}
`

func (g *CommonJSGenerator) GenerateUnit(unit *definition.CompilationUnit) error {
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
		"Unit":     unit,
		"GenTypes": genTypes,
		"GenState": g.GenState,
	}

	fileStr := util.ExecuteTemplate(fileTemplate, "file", nil, fileData)
	err := g.GenCtx.WritePackage(unit.Package, ".bb.js", fileStr)
	if err != nil {
		return err
	}

	// clear state for next unit
	g.GenState = NewCommonJSGeneratorState()

	return nil
}

// ==================== GenerateType ====================

func (g CommonJSGenerator) GenerateType(type_ definition.Type) (string, error) {
	return g.AcceptType(type_)
}

// ==================== GenerateTypeDefaultValue ====================

func (g CommonJSGenerator) GenerateTypeDefaultValue(type_ definition.Type) (string, error) {
	return g.AcceptTypeDefaultValue(type_)
}

// ==================== GenerateBasicType ====================

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

func (g CommonJSGenerator) GenerateBasicType(type_ *definition.BasicType) (string, error) {
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

// ==================== GenerateBasicTypeDefaultValue ====================

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

func (g CommonJSGenerator) GenerateBasicTypeDefaultValue(type_ *definition.BasicType) (string, error) {
	if str, ok := typeDefValueMap[type_.TypeTypeID]; ok {
		return str, nil
	}
	return "", fmt.Errorf("unknown basic type: %s", type_.String())
}

// ==================== GenerateString ====================

func (g CommonJSGenerator) GenerateString(string_ *definition.String) (string, error) {
	g.GenState.UseString = true
	return "String", nil
}

// ==================== GenerateStringDefaultValue ====================

func (g CommonJSGenerator) GenerateStringDefaultValue(string_ *definition.String) (string, error) {
	return `""`, nil
}

// ==================== GenerateBytes ====================

func (g CommonJSGenerator) GenerateBytes(bytes *definition.Bytes) (string, error) {
	return "Array", nil
}

// ==================== GenerateBytesDefaultValue ====================

func (g CommonJSGenerator) GenerateBytesDefaultValue(bytes *definition.Bytes) (string, error) {
	return `[]`, nil
}

// ==================== GenerateArray ====================

func (g CommonJSGenerator) GenerateArray(array *definition.Array) (string, error) {
	_, err := g.GenerateType(array.ElementType)
	if err != nil {
		return "", err
	}
	return "Array", nil
}

// ==================== GenerateArrayDefaultValue ====================

func (g CommonJSGenerator) GenerateArrayDefaultValue(array *definition.Array) (string, error) {
	// special case for struct array
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

// ==================== GenerateStruct ====================

var structTemplate = `
{{- define "field" -}}
                    // {{ .Pos }} {{ .Field.GetFieldKind }}: {{ .Field }}
                    {{- $fieldStrLen := len .FieldStr -}}
                    {{- if gt $fieldStrLen 0 }}
                    {{ .FieldStr }}
                    {{- end }}
{{- end -}}

{{- define "structConst" -}}
                /**
                 * Check if this struct has dynamic size
                 * @function dynamic
                 * @memberof {{ .StructDef.StructBelongs.Package }}.{{ .StructDef.StructName }}
                 * @static
                 * @returns {Boolean} If this struct has dynamic size
                 */
                {{ .StructDef.StructName }}.dynamic = function() {
                    return {{ if .StructDef.StructDynamic }}true{{ else }}false{{ end }};
                };

                /**
                 * Get the size of this struct
                 * @function size
                 * @memberof {{ .StructDef.StructBelongs.Package }}.{{ .StructDef.StructName }}
                 * @static
                 * @returns {Number} Size of this struct
                 */
                {{ .StructDef.StructName }}.size = function() {
                    return {{ calc .StructDef.StructBitSize "/" 8 }};
                };
{{ end }}

{{- define "methodGroupDef" -}}
{{- $methods := .MethodGroupDef.Values -}}
{{- $firstMethod := index $methods 0 -}}
{{- $methodName := TocamelCase $firstMethod.MethodName -}}
{{- $fieldDef := $firstMethod.MethodBelongs -}}
{{- $fieldName := TocamelCase $fieldDef.FieldName -}}
{{- $structDef := $fieldDef.FieldBelongs -}}
{{- $structName := $structDef.StructName -}}
                /**
                 * Method group: {{ $methodName }}
                 * @member {"{{ $fieldName }}"} {{ $methodName }}
                 * @memberof {{ $structDef.StructBelongs.Package }}.{{ $structName }}
                 * @instance
                 */
                Object.defineProperty({{ $structName }}.prototype, "{{ $methodName }}", {
                    {{- range $method := .MethodGroupDef.Values }}
                    {{ GenerateMethod $method }}
                    {{- end }}
                });
{{ end }}

{{- define "structDef" -}}
{{- $structDef := .StructDef -}}
{{- $fieldStrs := .FieldStrs -}}

                /**
                 * Struct: {{ $structDef }}
                 * @memberof {{ $structDef.StructBelongs.Package }}
                 * @interface I{{ $structDef.StructName }}
                {{- range $field := .StructDef.StructFields.Values }}
                {{- if $field.GetFieldKind.IsNormal }}
                {{- $fieldName := TocamelCase $field.FieldName }}
                {{- if $field.FieldType.GetTypeID.IsStruct }}
                 * @property {{"{"}}{{ $field.FieldType.StructBelongs.Package }}.I{{ $field.FieldType.StructName }}|null} {{ $fieldName }} {{ $structDef.StructName }} {{ $field.FieldName }}
                {{- else }}
                 * @property {{"{"}}{{ GenerateType $field.FieldType }}|null} {{ $fieldName }} {{ $structDef.StructName }} {{ $field.FieldName }}
                {{- end }}
                {{- end }}
                {{- end }}
                 */

                /**
                 * Constructor for {{ $structDef.StructName }}
                 * @memberof {{ $structDef.StructBelongs.Package }}
                 * @classdesc Structure for {{ $structDef.StructName }}
                 * @implements I{{ $structDef.StructName }}
                 * @constructor
                 * @param {{"{"}}{{ $structDef.StructBelongs.Package }}.I{{ $structDef.StructName }}=} properties Properties to set
                 */
                function {{ $structDef.StructName }}(properties) {
                    {{- range $fieldStr := $fieldStrs }}
                    {{ $fieldStr }}
                    {{- end }}
                    if (properties) {
                        for (var key = Object.keys(properties), i = 0; i < key.length; i++) {
                            if (properties[key[i]] != null) {
                                this[key[i]] = properties[key[i]];
                            }
                        }
                    }
                }

                /**
                 * Create a new {{ $structDef.StructName }} instance
                 * @function create
                 * @memberof {{ $structDef.StructBelongs.Package }}.{{ $structDef.StructName }}
                 * @static
                 * @param {{"{"}}{{ $structDef.StructBelongs.Package }}.I{{ $structDef.StructName }}=} properties Properties to set
                 * @returns {{"{"}}{{ $structDef.StructBelongs.Package }}.I{{ $structDef.StructName }}} {{ $structDef.StructName }} instance
                 */
                {{ $structDef.StructName }}.create = function(properties) {
                    return new {{ $structDef.StructName }}(properties);
                };

                {{ .StructConstStr }}
                {{- range $methodStr := .MethodStrs }}
                {{ $methodStr }}
                {{- end }}
                {{ .EncoderStr }}

                {{ .DecoderStr }}

                return {{ $structDef.StructName }};
{{- end -}}
`

func (g CommonJSGenerator) GenerateStruct(structDef *definition.Struct) (string, error) {
	name := structDef.StructName
	// check if this struct is already generated
	if g.GenTypes.Has(structDef.StructName) {
		return name, nil
	}
	// check if this struct is in generating
	if g.GenStack.Has(structDef.StructName) {
		return name, nil
	}
	// push to stack
	g.GenStack.Put(structDef.StructName, nil)
	defer g.GenStack.Remove(structDef.StructName)

	genTy, err := g.generateStruct(structDef)
	if err != nil {
		return "", err
	}
	g.GenTypes.Put(structDef.StructName, genTy)

	return name, nil
}

func (g CommonJSGenerator) generateStruct(structDef *definition.Struct) (*GeneratedType, error) {
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

	constData := map[string]any{
		"StructDef": structDef,
	}
	constStr := util.ExecuteTemplate(structTemplate, "structConst", funcMap, constData)

	encoderStr, err := g.GenerateEncoder(structDef)
	if err != nil {
		return nil, err
	}

	decoderStr, err := g.GenerateDecoder(structDef)
	if err != nil {
		return nil, err
	}

	defData := map[string]any{
		"StructDef":      structDef,
		"FieldStrs":      fieldInitStrs,
		"StructConstStr": constStr,
		"MethodStrs":     methodStrs,
		"EncoderStr":     encoderStr,
		"DecoderStr":     decoderStr,
	}

	defStr := util.ExecuteTemplate(structTemplate, "structDef", funcMap, defData)

	code := &GeneratedType{
		GeneratedDef: defStr,
	}
	return code, nil
}

// ==================== GenerateStructDefaultValue ====================

func (g *CommonJSGenerator) GenerateStructDefaultValue(structDef *definition.Struct) (string, error) {
	prefix := g.generateStructPackagePrefix(structDef)
	return fmt.Sprintf("%s.create()", prefix), nil
}

// ==================== GenerateField ====================

func (g CommonJSGenerator) GenerateField(field definition.Field) (string, error) {
	return g.AcceptField(field)
}

// ==================== GenerateNormalField ====================

var normalFieldTemplate = `
{{- define "normalField" -}}
    {{- $field := .Field -}}
    {{- $fieldName := TocamelCase .Field.FieldName -}}
    {{- $tyStr := GenerateType $field.FieldType -}}
    {{- $tyDefValue := GenerateTypeDefaultValue $field.FieldType -}}
    this.{{ $fieldName }} = {{ $tyDefValue }};
{{- end -}}
`

func (g CommonJSGenerator) GenerateNormalField(field *definition.NormalField) (string, error) {
	funcMap := template.FuncMap{
		"GenerateType":             g.GenerateType,
		"GenerateTypeDefaultValue": g.GenerateTypeDefaultValue,
	}

	fieldData := map[string]any{
		"Field": field,
	}

	fieldStr := util.ExecuteTemplate(normalFieldTemplate, "normalField", funcMap, fieldData)
	return fieldStr, nil
}

// ==================== GenerateVoidField ====================

func (g CommonJSGenerator) GenerateVoidField(field *definition.VoidField) (string, error) {
	return "", nil
}

// ==================== GenerateEmbeddedField ====================

func (g CommonJSGenerator) GenerateEmbeddedField(field *definition.EmbeddedField) (string, error) {
	return "", nil
}

// ==================== GenerateConstantField ====================

var constantFieldTemplate = `
{{- define "constantField" -}}
    {{- $field := .Field -}}
    {{- $fieldName := TocamelCase .Field.FieldName -}}
    {{- $tyStr := GenerateType $field.FieldType -}}
    {{- $constValue := GenerateLiteral $field.FieldConstant -}}
    this.{{ $fieldName }} = {{ $constValue }};
{{- end -}}
`

func (g CommonJSGenerator) GenerateConstantField(field *definition.ConstantField) (string, error) {
	literalGentor := NewCommonJSLiteralGenerator()

	funcMap := template.FuncMap{
		"GenerateType":    g.GenerateType,
		"GenerateLiteral": literalGentor.GenerateLiteral,
	}

	fieldData := map[string]any{
		"Field": field,
	}

	fieldStr := util.ExecuteTemplate(constantFieldTemplate, "constantField", funcMap, fieldData)
	return fieldStr, nil
}

// ==================== GenerateMethod ====================

func (g CommonJSGenerator) GenerateMethod(method definition.Method) (string, error) {
	return g.AcceptMethod(method)
}

// ==================== GenerateMethodDecl ====================

func (g CommonJSGenerator) GenerateMethodDecl(method definition.Method) (string, error) {
	return g.AcceptMethodDecl(method)
}

// ==================== GenerateDefaultGetterDecl ====================

func (g CommonJSGenerator) GenerateDefaultGetterDecl(method *definition.GetMethod) (string, error) {
	panic("not implemented")
}

// ==================== GenerateDefaultSetterDecl ====================

func (g CommonJSGenerator) GenerateDefaultSetterDecl(method *definition.SetMethod) (string, error) {
	panic("not implemented")
}

// ==================== GenerateCustomGetterDecl ====================

func (g CommonJSGenerator) GenerateCustomGetterDecl(method *definition.GetMethod) (string, error) {
	panic("not implemented")
}

// ==================== GenerateCustomSetterDecl ====================

func (g CommonJSGenerator) GenerateCustomSetterDecl(method *definition.SetMethod) (string, error) {
	panic("not implemented")
}

// ==================== GenerateRawGetterDecl ====================

func (g CommonJSGenerator) GenerateRawGetterDecl(field definition.Field) (string, error) {
	panic("not implemented")
}

// ==================== GenerateRawSetterDecl ====================

func (g CommonJSGenerator) GenerateRawSetterDecl(field definition.Field) (string, error) {
	panic("not implemented")
}

// ==================== GenerateDefaultGetter ====================

// TODO: support default getter

func (g CommonJSGenerator) GenerateDefaultGetter(method *definition.GetMethod) (string, error) {
	panic("not implemented")
}

// ==================== GenerateDefaultSetter ====================

// TODO: support default setter

func (g CommonJSGenerator) GenerateDefaultSetter(method *definition.SetMethod) (string, error) {
	panic("not implemented")
}

// ==================== GenerateCustomGetter ====================

var customGetterTemplate = `
{{- define "customGetter" -}}
{{- $retTyStr := GenerateType .MethodDef.MethodRetType -}}
{{- $field := .MethodDef.MethodBelongs -}}
{{- $fieldName := TocamelCase $field.FieldName -}}
{{- $valueStr := printf "this.%s" $fieldName -}}
{{- $exprStr := GenerateExpr .MethodDef.MethodExpr $valueStr -}}
                    // CustomGetter: {{ $fieldName }}
                    get: function () { return {{ $exprStr }}; },
{{- end -}}
`

func (g CommonJSGenerator) GenerateCustomGetter(method *definition.GetMethod) (string, error) {
	funcMap := template.FuncMap{
		"GenerateType": g.GenerateType,
		"GenerateExpr": g.GenerateExpr,
	}

	fieldData := map[string]any{
		"MethodDef": method,
		"GenOption": g.GenCtx.GenOptions,
	}
	fieldStr := util.ExecuteTemplate(customGetterTemplate, "customGetter", funcMap, fieldData)
	return fieldStr, nil
}

// ==================== GenerateCustomSetter ====================

var customSetterTemplate = `
{{- define "customSetter" -}}
{{- $paramTyStr := GenerateType .MethodDef.MethodParamType -}}
{{- $field := .MethodDef.MethodBelongs -}}
{{- $fieldName := TocamelCase $field.FieldName -}}
{{- $exprStr := GenerateExpr .MethodDef.MethodExpr "value" -}}
                    // CustomSetter: {{ $fieldName }}
                    set: function (value) { this.{{ $fieldName }} = {{ $exprStr }}; },
{{- end -}}
`

func (g CommonJSGenerator) GenerateCustomSetter(method *definition.SetMethod) (string, error) {
	funcMap := template.FuncMap{
		"GenerateType": g.GenerateType,
		"GenerateExpr": g.GenerateExpr,
	}

	fieldData := map[string]any{
		"MethodDef": method,
		"GenOption": g.GenCtx.GenOptions,
	}
	fieldStr := util.ExecuteTemplate(customSetterTemplate, "customSetter", funcMap, fieldData)
	return fieldStr, nil
}

// ==================== GenerateRawGetter ====================

func (g CommonJSGenerator) GenerateRawGetter(field definition.Field) (string, error) {
	panic("not implemented")
}

// ==================== GenerateRawSetter ====================

func (g CommonJSGenerator) GenerateRawSetter(field definition.Field) (string, error) {
	panic("not implemented")
}

// ==================== GenerateEnum ====================

func (g CommonJSGenerator) GenerateEnum(enumDef *definition.Enum) (string, error) {
	name := enumDef.EnumName
	// check if this enum is already generated
	if g.GenTypes.Has(enumDef.EnumName) {
		return name, nil
	}
	// check if this enum is in generating
	if g.GenStack.Has(enumDef.EnumName) {
		return name, nil
	}
	// push to stack
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
                var valuesById = {}, values = Object.create(valuesById);
                {{- range .EnumDef.EnumValues.Values }}
                values[valuesById[{{ .EnumValue }}] = "{{ .EnumValueName }}"] = {{ .EnumValue }};
                {{- end }}
                return values;
{{- end -}}

{{- define "enumComments" -}}
            /**
             * Enum: {{ .EnumDef.ShortString }}
             * @name {{ .EnumDef.EnumName }}
             * @enum {number}
            {{- range .EnumDef.EnumValues.Values }}
             * @property {number} {{ .EnumValueName }}={{ .EnumValue }} {{ .EnumValueName }} value
            {{- end }}
             */
{{- end -}}
`

func (g CommonJSGenerator) generateEnum(enumDef *definition.Enum) (*GeneratedType, error) {
	enumDefData := map[string]any{
		"EnumDef": enumDef,
	}

	enumDefStr := util.ExecuteTemplate(enumTemplate, "enumDef", nil, enumDefData)

	enumCommentsData := map[string]any{
		"EnumDef": enumDef,
	}

	enumCommentsStr := util.ExecuteTemplate(enumTemplate, "enumComments", nil, enumCommentsData)

	code := &GeneratedType{
		GeneratedDef: enumDefStr,
		JsComments:   enumCommentsStr,
	}

	return code, nil
}

// ==================== GenerateEnumDefaultValue ====================

func (g CommonJSGenerator) GenerateEnumDefaultValue(enumDef *definition.Enum) (string, error) {
	return "null", nil
}

// ==================== GenerateEncoder ====================

var encoderTemplate = `
{{- define "encodeField" -}}
                // {{ .Pos }} {{ .Field.GetFieldKind }}: {{ .Field }}
                {{- range $encodeStmt := .EncodeStmts }}
                    {{ $encodeStmt }}
                {{- end -}}
{{- end -}}

{{- define "encoder" -}}
{{- $structName := .StructDef.StructName -}}
                /**
                 * Calculate encoded size of {{ $structName }} object
                 * @function encode_size
                 * @memberof {{ .StructDef.StructBelongs.Package }}.{{ $structName }}
                 * @static
                 * @param {{"{"}}{{ .StructDef.StructBelongs.Package }}.I{{ $structName }}} obj {{ $structName }} object
                 * @returns {Number} Encoded size of {{ $structName }} object
                 */
                {{ $structName }}.encode_size = function(obj) {
                    {{- if .StructDef.StructDynamic }}
                    var size = {{ calc .StructDef.StructBitSize "/" 8 }}
                    {{- range $field := .StructDef.StructFields.Values }}
                        {{- if $field.GetFieldKind.IsNormal }}
                        {{- $fieldName := TocamelCase $field.FieldName }}
                            {{- if $field.FieldType.GetTypeID.IsArray }}
                                {{- if $field.FieldType.ElementType.GetTypeID.IsStruct }}
                                    {{- if $field.FieldType.ElementType.GetTypeDynamic }}
                                        {{- range $i := iterate 0 $field.FieldType.Length }}
                    size += obj.{{ $fieldName }}[{{ $i }}].encode_size();
                                        {{- end }}
                                    {{- end }}
                                {{- else if $field.FieldType.ElementType.GetTypeID.IsString }}
                                    {{- range $i := iterate 0 $field.FieldType.Length }}
                    size += stringToUTF8BytesCount(obj.{{ $fieldName }}[{{ $i }}]) + 1;
                                    {{- end }}
                                {{- else if $field.FieldType.ElementType.GetTypeID.IsBytes }}
                                    {{- range $i := iterate 0 $field.FieldType.Length }}
                    size += obj.{{ $fieldName }}[{{ $i }}].length;
                    size += 1{{ range $j := iterate 1 5 }} + Boolean(obj.{{ $fieldName }}[{{ $i }}].length >> {{ calc $j "*" 7 }}){{ end }};
                                    {{- end }}
                                {{- end }}
                            {{- else if $field.FieldType.GetTypeID.IsStruct }}
                                {{- if $field.FieldType.GetTypeDynamic }}
                    size += obj.{{ $fieldName }}.encode_size();
                                {{- end }}
                            {{- else if $field.FieldType.GetTypeID.IsString }}
                    size += stringToUTF8BytesCount(obj.{{ $fieldName }}) + 1;
                            {{- else if $field.FieldType.GetTypeID.IsBytes }}
                    size += obj.{{ $fieldName }}.length;
                    size += 1{{ range $j := iterate 1 5 }} + Boolean(obj.{{ $fieldName }}.length >> {{ calc $j "*" 7 }}){{ end }};
                            {{- end }}
                        {{- end }}
                    {{- end }}
                    return size;
                    {{- else }}
                    return {{ calc .StructDef.StructBitSize "/" 8 }};
                    {{- end }}
                };

                /**
                 * Calculate encoded size of {{ $structName }} object
                 * @function encode_size
                 * @memberof {{ .StructDef.StructBelongs.Package }}.{{ $structName }}
                 * @instance
                 * @returns {Number} Encoded size of {{ $structName }} object
                 */
                {{ $structName }}.prototype.encode_size = function() {
                    return {{ $structName }}.encode_size(this);
                };

                /**
                 * Encode {{ $structName }} object to buffer
                 * @function encode
                 * @memberof {{ .StructDef.StructBelongs.Package }}.{{ $structName }}
                 * @static
                 * @param {{"{"}}{{ .StructDef.StructBelongs.Package }}.I{{ $structName }}} obj {{ $structName }} object
                 * @param {Array | Uint8Array} [buffer] The buffer to encode data
                 * @param {Number} [start] The start position to store the encoded data
                 * @returns {Array | Number} Array with encoded data if data is not provided, otherwise number of bytes encoded
                 */
                {{ $structName }}.encode = function(obj, buffer, start) {
                    if (obj === undefined) return buffer === undefined ? -1 : undefined;
                    var data = buffer;
                    if (data === undefined) data = new Array({{ if .Dynamic }}obj.encode_size(){{ else }}{{ calc .StructDef.StructBitSize "/" 8 }}{{ end }});
                    if (start === undefined) start = 0;
                    {{- if .Dynamic }}
                    var offset = 0;
                    {{- end }}
                    {{- range $encodeStr := .EncodeStrs }}
                    {{ $encodeStr }}
                    {{- end }}
                    return buffer === undefined ? data : {{ if .Dynamic }}offset + {{ end }}{{ calc .StructDef.StructBitSize "/" 8 }}
                };

                /**
                 * Encode {{ $structName }} object to buffer
                 * @function encode
                 * @memberof {{ .StructDef.StructBelongs.Package }}.{{ $structName }}
                 * @instance
                 * @param {Array | Uint8Array} [data] The buffer to encode data
                 * @param {Number} [start] The start position to store the encoded data
                 * @returns {Array | Number} Array with encoded data if data is not provided, otherwise number of bytes encoded
                 */
                {{ $structName }}.prototype.encode = function(data, start) {
                    return {{ $structName }}.encode(this, data, start);
                };
{{- end -}}
`

func (g CommonJSGenerator) GenerateEncoder(structDef *definition.Struct) (string, error) {
	encodeStrs := []string{}
	if err := structDef.ForEachFieldWithPos(func(field definition.Field, fieldIndex int, startBits int64, dynamic bool, pos string) error {
		encodeStmts, err := g.generateEncodeField(field, startBits)
		if err != nil {
			return err
		}
		// skip
		if len(encodeStmts) == 0 {
			return nil
		}
		// remove blank lines
		filteredEncodeStmts := []string{}
		for _, stmt := range encodeStmts {
			if stmt != "" {
				filteredEncodeStmts = append(filteredEncodeStmts, stmt)
			}
		}
		encodeStmts = filteredEncodeStmts

		encodeFieldData := map[string]any{
			"Pos":         pos,
			"Field":       field,
			"EncodeStmts": encodeStmts,
		}

		str := util.ExecuteTemplate(encoderTemplate, "encodeField", nil, encodeFieldData)
		encodeStrs = append(encodeStrs, str)
		return nil
	}); err != nil {
		return "", err
	}

	fieldData := map[string]any{
		"StructDef":  structDef,
		"EncodeStrs": encodeStrs,
		"GenOption":  g.GenCtx.GenOptions,
		"Dynamic":    structDef.GetTypeDynamic(),
	}

	encoderStr := util.ExecuteTemplate(encoderTemplate, "encoder", nil, fieldData)
	return encoderStr, nil
}

// return multiple statements (lines) to encode a field
func (g CommonJSGenerator) generateEncodeField(field definition.Field, startBits int64) ([]string, error) {
	switch val := field.(type) {
	case *definition.ConstantField:
		return g.generateEncodeConstantField(val, startBits)
	case *definition.VoidField:
		return g.generateEncodeVoidField(val, startBits)
	case *definition.EmbeddedField:
		return g.generateEncodeEmbeddedField(val, startBits)
	case *definition.NormalField:
		return g.generateEncodeNormalField(val, startBits)
	default:
		return nil, fmt.Errorf("internal error: unknown field kind %T", val)
	}
}

var fieldEncoderTemplate = `
{{- define "encodeTempVarName" -}}
    tempFieldAtPos{{ .StartBits }}
{{- end -}}

{{- define "encodeStructFieldName" -}}
{{- $fieldName := TocamelCase .FieldName -}}
    obj.{{ $fieldName }}
{{- end -}}

{{- define "encodeNormalFieldStruct" -}}
{{- $packagePrefix := call .GenerateStructPackagePrefix .FieldStruct -}}
    {{ if .FieldStruct.GetTypeDynamic }}offset += {{ end }}{{ $packagePrefix }}.encode({{ .FieldName }}, data, {{ if .Dynamic }}offset + {{ end }}start + {{ .FromByte }});
{{- end -}}

{{- define "encodeNormalFieldTempVarDecl" -}}
    var {{ .TempName }} = {{ .FieldName }};
{{- end -}}

{{- define "encodeNormalFieldTempVarDeclEnum" -}}
{{- $inPackage := .FieldDef.FieldBelongs.StructBelongs.LocalNames.Has .EnumDef.EnumName -}}
{{- if not $inPackage -}}
    var {{ .TempName }} = (typeof {{ .FieldName }} === "number") ? {{ .FieldName }} : $root.{{ .EnumDef.EnumBelongs.Package }}.{{ .EnumDef.EnumName }}[{{ .FieldName }}];
{{- else -}}
    var {{ .TempName }} = (typeof {{ .FieldName }} === "number") ? {{ .FieldName }} : $package.{{ .EnumDef.EnumName }}[{{ .FieldName }}];
{{- end -}}
{{- end -}}

{{- define "encodeNormalFieldTempVarDeclFloatCast" -}}
{{- if .IsFloat32 -}}
    var {{ .TempName }} = floatToUint32Bits({{ .FieldName }});
{{- else -}}
    var {{ .TempName }} = doubleToUint64Bits({{ .FieldName }});
{{- end -}}
{{- end -}}

{{- define "encodeNormalFieldTempVarDeclOnly" -}}
    var {{ .TempName }};
{{- end -}}

{{- define "encodeNormalFieldTempVarAssignEnum" -}}
{{- $inPackage := .FieldDef.FieldBelongs.StructBelongs.LocalNames.Has .EnumDef.EnumName -}}
{{- if not $inPackage -}}
    {{ .TempName }} = (typeof {{ .FieldName }} === "number") ? {{ .FieldName }} : $root.{{ .EnumDef.EnumBelongs.Package }}.{{ .EnumDef.EnumName }}[{{ .FieldName }}];
{{- else -}}
    {{ .TempName }} = (typeof {{ .FieldName }} === "number") ? {{ .FieldName }} : $package.{{ .EnumDef.EnumName }}[{{ .FieldName }}];
{{- end -}}
{{- end -}}

{{- define "encodeNormalFieldTempVarAssignFloatCast" -}}
{{- if .IsFloat32 -}}
    {{ .TempName }} = floatToUint32Bits({{ .FieldName }});
{{- else -}}
    {{ .TempName }} = doubleToUint64Bits({{ .FieldName }});
{{- end -}}
{{- end -}}

{{- define "encodeNormalFieldString" -}}
                    (function() {
                        var {{ .TempName }} = stringToUTF8Bytes({{ .FieldName }}, data, offset + start + {{ .FromByte }});
                        data[offset + start + {{ .FromByte }} + {{ .TempName }}] = 0;
                        offset += {{ .TempName }} + 1;
                    })();
{{- end -}}

{{- define "encodeNormalFieldBytes" -}}
                    (function() {
                        var {{ .TempName }} = {{ .FieldName }}.length;
                        do { data[offset + start + {{ .FromByte }}] = {{ .TempName }} & {{ .GetMask }} | {{ .SetMask }}; offset++; {{ .TempName }} >>= {{ .Shift }}; } while ({{ .TempName }} > 0);
                        data[offset - 1 + start + {{ .FromByte }}] &= ~{{ .SetMask }};
                        for (var i = 0; i < {{ .FieldName }}.length; i++) data[offset + start + {{ .FromByte }} + i] = {{ .FieldName }}[i];
                        offset += {{ .FieldName }}.length;
                    })();
{{- end -}}

{{- define "encodeImpl" -}}
    data[{{ if .Dynamic }}offset + {{ end }}start + {{ .BytePos }}] {{ .Operator }} {{ .FieldData }};
{{- end -}}
`

func (g CommonJSGenerator) generateEncodeTempVarName(startBits int64) string {
	encodeTempVarNameData := map[string]any{
		"StartBits": startBits,
	}

	return util.ExecuteTemplate(fieldEncoderTemplate, "encodeTempVarName", nil, encodeTempVarNameData)
}

func (g CommonJSGenerator) generateEncodeStructFieldName(name string) string {
	encodeStructFieldNameData := map[string]any{
		"FieldName": name,
	}

	return util.ExecuteTemplate(fieldEncoderTemplate, "encodeStructFieldName", nil, encodeStructFieldNameData)
}

func (g CommonJSGenerator) generateEncodeConstantField(field *definition.ConstantField, startBits int64) ([]string, error) {
	structDynamic := field.FieldBelongs.GetTypeDynamic()
	var byteOrder binary.ByteOrder = binary.LittleEndian
	if gen.MatchOption(field.FieldOptions, "order", "big") {
		byteOrder = binary.BigEndian
	}

	buffer := &bytes.Buffer{}
	value := field.FieldConstant.GetLiteralValueIn(field.FieldType.TypeTypeID)
	err := binary.Write(buffer, byteOrder, value)
	if err != nil {
		return nil, fmt.Errorf("internal error: %s", err)
	}

	data := buffer.Bytes()
	fieldData := func(i int64) string {
		return g.generateHex(data[i])
	}

	from := startBits
	to := startBits + field.GetFieldBitSize()

	encodeStmts := g.generateEncodeImpl(from, to, fieldData, g.generateBin, g.generateDec, false, structDynamic)
	return encodeStmts, nil
}

func (g CommonJSGenerator) generateEncodeVoidField(field *definition.VoidField, startBits int64) ([]string, error) {
	return []string{""}, nil
}

func (g CommonJSGenerator) generateEncodeEmbeddedField(field *definition.EmbeddedField, startBits int64) ([]string, error) {
	return nil, nil
}

func (g CommonJSGenerator) generateEncodeNormalField(field *definition.NormalField, startBits int64) ([]string, error) {
	from := startBits
	to := startBits + field.GetFieldBitSize()
	structDynamic := field.FieldBelongs.GetTypeDynamic()
	encodeStmts := []string{}

	switch ty := field.FieldType.(type) {
	case *definition.Struct, *definition.String, *definition.Bytes:
		name := g.generateEncodeStructFieldName(field.FieldName)
		stmts, err := g.generateEncodeNormalFieldImpl(name, ty, field.FieldOptions, structDynamic, from, to)
		if err != nil {
			return nil, err
		}
		encodeStmts = append(encodeStmts, stmts...)

	case *definition.BasicType:
		name := g.generateEncodeStructFieldName(field.FieldName)

		if ty.GetTypeID().IsFloat() {
			// decl and assign temp variable to cast float to uint
			// optimization for commonjs only
			tempName := g.generateEncodeTempVarName(startBits)
			encodeNormalFieldTempVarDeclFloatCastData := map[string]any{
				"TempName":  tempName,
				"FieldName": name,
				"IsFloat32": ty.GetTypeID() == definition.TypeID_Float32,
			}
			stmt := util.ExecuteTemplate(fieldEncoderTemplate, "encodeNormalFieldTempVarDeclFloatCast", nil, encodeNormalFieldTempVarDeclFloatCastData)
			encodeStmts = append(encodeStmts, stmt)

			name = tempName
		}

		stmts, err := g.generateEncodeNormalFieldImpl(name, ty, field.FieldOptions, structDynamic, from, to)
		if err != nil {
			return nil, err
		}
		encodeStmts = append(encodeStmts, stmts...)

	case *definition.Enum:
		tempName := g.generateEncodeTempVarName(startBits)
		// any integer type is ok (except 64-bit)
		tempTy := &definition.Uint32

		encodeNormalFieldTempVarDeclData := map[string]any{
			"EnumDef":   ty,
			"TempName":  tempName,
			"FieldName": g.generateEncodeStructFieldName(field.FieldName),
			"FieldDef":  field,
		}
		declStr := util.ExecuteTemplate(fieldEncoderTemplate, "encodeNormalFieldTempVarDeclEnum", nil, encodeNormalFieldTempVarDeclData)
		encodeStmts = append(encodeStmts, declStr)

		stmts, err := g.generateEncodeNormalFieldImpl(tempName, tempTy, field.FieldOptions, structDynamic, from, to)
		if err != nil {
			return nil, err
		}
		encodeStmts = append(encodeStmts, stmts...)

	case *definition.Array:
		elemTy := ty.ElementType
		elemBitSize := field.FieldBitSize / ty.Length

		name := g.generateEncodeStructFieldName(field.FieldName)

		// temp variable declaration
		var nameIndex func(int64) string
		switch ty.ElementType.(type) {
		case *definition.Struct, *definition.String, *definition.Bytes:
			nameIndex = func(index int64) string {
				return fmt.Sprintf("%s[%d]", name, index)
			}

		case *definition.BasicType:
			nameIndex = func(index int64) string {
				return fmt.Sprintf("%s[%d]", name, index)
			}
			// same as enum
			if ty.ElementType.GetTypeID().IsFloat() {
				tempName := g.generateEncodeTempVarName(startBits)

				encodeNormalFieldTempVarDeclOnlyData := map[string]any{
					"TempName": tempName,
				}
				declStr := util.ExecuteTemplate(fieldEncoderTemplate, "encodeNormalFieldTempVarDeclOnly", nil, encodeNormalFieldTempVarDeclOnlyData)
				encodeStmts = append(encodeStmts, declStr)

				nameIndex = func(_ int64) string {
					return tempName
				}

				// change elemTy to 32 or 64 bit integer type
				switch ty.ElementType.GetTypeID() {
				case definition.TypeID_Float32:
					elemTy = &definition.Uint32
				case definition.TypeID_Float64:
					elemTy = &definition.Uint64
				}
			}

		case *definition.Enum:
			tempName := g.generateEncodeTempVarName(startBits)

			encodeNormalFieldTempVarDeclOnlyData := map[string]any{
				"TempName": tempName,
			}
			declStr := util.ExecuteTemplate(fieldEncoderTemplate, "encodeNormalFieldTempVarDeclOnly", nil, encodeNormalFieldTempVarDeclOnlyData)
			encodeStmts = append(encodeStmts, declStr)

			nameIndex = func(_ int64) string {
				return tempName
			}

			// change elemTy to any integer type (except 64-bit)
			elemTy = &definition.Uint32
		default:
			return nil, fmt.Errorf("internal error: unsupported array element type %T", ty.ElementType)
		}

		for i := int64(0); i < ty.Length; i++ {
			subFrom := from + i*elemBitSize
			subTo := from + (i+1)*elemBitSize

			subName := nameIndex(i)

			switch ty.ElementType.(type) {
			case *definition.BasicType:
				if ty.ElementType.GetTypeID().IsFloat() {
					encodeNormalFieldTempVarAssignFloatCastData := map[string]any{
						"TempName":  subName,
						"FieldName": fmt.Sprintf("%s[%d]", name, i),
						"IsFloat32": ty.ElementType.GetTypeID() == definition.TypeID_Float32,
					}
					assignStr := util.ExecuteTemplate(fieldEncoderTemplate, "encodeNormalFieldTempVarAssignFloatCast", nil, encodeNormalFieldTempVarAssignFloatCastData)
					encodeStmts = append(encodeStmts, assignStr)
				}
			case *definition.Enum:
				encodeNormalFieldTempVarAssignData := map[string]any{
					"EnumDef":   ty.ElementType,
					"TempName":  subName,
					"FieldName": fmt.Sprintf("%s[%d]", name, i),
					"FieldDef":  field,
				}
				assignStr := util.ExecuteTemplate(fieldEncoderTemplate, "encodeNormalFieldTempVarAssignEnum", nil, encodeNormalFieldTempVarAssignData)
				encodeStmts = append(encodeStmts, assignStr)
			default:
			}

			stmts, err := g.generateEncodeNormalFieldImpl(subName, elemTy, field.FieldOptions, structDynamic, subFrom, subTo)
			if err != nil {
				return nil, err
			}

			encodeStmts = append(encodeStmts, stmts...)
		}

	default:
		return nil, fmt.Errorf("internal error: unknown field kind %T", ty)
	}

	return encodeStmts, nil
}

// generateEncodeNormalFieldImpl does not handle array field or generate temp variable declaration
func (g CommonJSGenerator) generateEncodeNormalFieldImpl(fieldNameStr string, fieldType definition.Type, fieldOptions *util.OrderedMap[string, *definition.Option], structDynamic bool, from, to int64) ([]string, error) {
	encodeStmts := []string{}
	fieldBitSize := to - from

	switch ty := fieldType.(type) {
	case *definition.Struct:
		encodeNormalFieldStructData := map[string]any{
			"FieldStruct":                 ty,
			"FieldName":                   fieldNameStr,
			"FromByte":                    from / 8,
			"Dynamic":                     structDynamic,
			"GenerateStructPackagePrefix": g.generateStructPackagePrefix,
		}

		stmt := util.ExecuteTemplate(fieldEncoderTemplate, "encodeNormalFieldStruct", nil, encodeNormalFieldStructData)
		encodeStmts = append(encodeStmts, stmt)

	case *definition.Enum:
		panic("unreachable, enum field should be handled in generateEncodeNormalField")

	case *definition.String:
		encodeNormalFieldStringData := map[string]any{
			"FieldName": fieldNameStr,
			"FromByte":  from / 8,
			"TempName":  g.generateEncodeTempVarName(from),
		}

		stmt := util.ExecuteTemplate(fieldEncoderTemplate, "encodeNormalFieldString", nil, encodeNormalFieldStringData)
		encodeStmts = append(encodeStmts, stmt)

	case *definition.Bytes:
		encodeNormalFieldBytesData := map[string]any{
			"FieldName": fieldNameStr,
			"FromByte":  from / 8,
			"GetMask":   g.generateHex((1 << 7) - 1),
			"SetMask":   g.generateHex(1 << 7),
			"Shift":     7,
			"TempName":  g.generateEncodeTempVarName(from),
		}

		stmt := util.ExecuteTemplate(fieldEncoderTemplate, "encodeNormalFieldBytes", nil, encodeNormalFieldBytesData)
		encodeStmts = append(encodeStmts, stmt)

	case *definition.BasicType:
		generateDec := g.generateDec
		generateBin := g.generateBin
		cast := false
		if ty.GetTypeBitSize() > 32 {
			generateDec = g.generateDecBigInt
			generateBin = g.generateBinBigInt
			cast = true
		}
		// little endian as default
		fieldData := func(i int64) string {
			// expr = fieldNameStr >> 8*i
			expr := &definition.BinopExpr{
				Op: definition.ExprOp_SHR,
				Expr1: &definition.RawExpr{
					Expr: fieldNameStr,
				},
				Expr2: &definition.RawExpr{
					Expr: generateDec(8 * i),
				},
			}
			exprStr, err := g.GenerateExpr(expr, "")
			if err != nil {
				panic(fmt.Errorf("internal error: %s", err))
			}
			return exprStr
		}
		// big endian
		if gen.MatchOption(fieldOptions, "order", "big") {
			fieldData = func(i int64) string {
				// expr = fieldNameStr >> max(0, fieldBitSize-8*(i+1))
				expr := &definition.BinopExpr{
					Op: definition.ExprOp_SHR,
					Expr1: &definition.RawExpr{
						Expr: fieldNameStr,
					},
					Expr2: &definition.RawExpr{
						Expr: generateDec(max(0, fieldBitSize-8*(i+1))),
					},
				}
				exprStr, err := g.GenerateExpr(expr, "")
				if err != nil {
					panic(fmt.Errorf("internal error: %s", err))
				}
				return exprStr
			}
		}
		encodeStmts = append(encodeStmts, g.generateEncodeImpl(from, to, fieldData, generateBin, generateDec, cast, structDynamic)...)

	case *definition.Array:
		panic("unreachable, array field should be handled in generateEncodeNormalField")

	default:
		return nil, fmt.Errorf("internal error: unknown field kind %T", ty)
	}

	return encodeStmts, nil
}

// generateEncodeImpl generates encode implementation from 'from' bit to 'to' bit and align to 8 bits
// from: start bit position of encoded data
// to: end bit position of encoded data
// from is inclusive, to is exclusive, i.e. [from, to)
//
// e.g. from = 3, to = 11
//
//	data[0] = (data[0] & 0b00000111) | ((fieldData(0) << 3) & 0b11111000)
//	data[1] = (data[1] & 0b11111000) | ((fieldData(0) >> 5) & 0b00000111)
//
// fieldData: a function callback to get expression of x-th byte data
//
// e.g. big endian
//
//	fieldData(0) -> ((structPtr->intField >> 24) & 0xff)
//	fieldData(1) -> ((structPtr->intField >> 16) & 0xff)
//	fieldData(2) -> ((structPtr->intField >> 8) & 0xff)
//	fieldData(3) -> ((structPtr->intField >> 0) & 0xff)
func (g CommonJSGenerator) generateEncodeImpl(from, to int64, fieldData func(int64) string, generateBin func(any) string, generateDec func(any) string, cast bool, structDynamic bool) []string {
	encodeStmts := []string{}
	// generate encode implentation from 'from' bit to 'to' bit and align to 8 bits
	// e.g. from = 3, to = 11 -> loop 2 times: 3-7, 8-11
	for i := from; i < to; i = (i + 8) & (^7) {
		// nextI is the right bound of current encode expression
		// use nextI to calculate the mask of current encode expression
		nextI := min(to, (i+8)&(^7))
		dataMask := ((1 << (((nextI - 1) & 7) + 1)) - 1) & (^((1 << (i & 7)) - 1))

		// operator is '=' if is filling the whole byte, otherwise is '|='
		operator := ""
		if i%8 == 0 {
			operator = exprOpToString[definition.ExprOp_ASSIGN] // "="
		} else {
			operator = exprOpToString[definition.ExprOp_BOR] + exprOpToString[definition.ExprOp_ASSIGN] // "|="
		}

		// we use 'from' and 'to' to denote the bit position in encoded data
		// we use 'begin' and 'end' to denote the bit position in raw field data
		begin := i - from
		end := nextI - from

		var expr definition.Expr

		// 'i' is the start bit position in encoded data
		// 'j' is the start bit position in raw field data
		j := begin
		// first half
		// e.g. j = 3, end = 11
		//      j = 3, nextJ = 8, fieldMask = 0b11111000, shiftRight = 3
		if j < end {
			// nextJ is the right bound of current field data (aligned to 8 bits)
			nextJ := min(end, (j+8)&(^7))
			fieldMask := ((1 << (((nextJ - 1) & 7) + 1)) - 1) & (^((1 << (j & 7)) - 1))
			shiftRight := j % 8
			// expr = (fieldData(j/8) & fieldMask) >> j%8
			expr = &definition.BinopExpr{
				Op: definition.ExprOp_SHR,
				Expr1: &definition.BinopExpr{
					Op: definition.ExprOp_BAND,
					Expr1: &definition.RawExpr{
						Expr: fieldData(j / 8),
					},
					Expr2: &definition.RawExpr{
						Expr: generateBin(fieldMask),
					},
				},
				Expr2: &definition.RawExpr{
					Expr: generateDec(shiftRight),
				},
			}

			// jump to second half
			j = nextJ
		}
		// second half (if exists)
		// e.g. j = 8, end = 11
		//      j = 8, nextJ = 11, fieldMask = 0b00000111, shiftLeft = 5
		if j < end {
			nextJ := min(end, (j+8)&(^7))
			fieldMask := ((1 << (((nextJ - 1) & 7) + 1)) - 1) & (^((1 << (j & 7)) - 1))
			shiftLeft := 8 - nextJ%8
			// expr = expr | (fieldData(j/8) & fieldMask) << (8 - nextJ%8)
			expr = &definition.BinopExpr{
				Op:    definition.ExprOp_BOR,
				Expr1: expr,
				Expr2: &definition.BinopExpr{
					Op: definition.ExprOp_SHL,
					Expr1: &definition.BinopExpr{
						Op: definition.ExprOp_BAND,
						Expr1: &definition.RawExpr{
							Expr: fieldData(j / 8),
						},
						Expr2: &definition.RawExpr{
							Expr: generateBin(fieldMask),
						},
					},
					Expr2: &definition.RawExpr{
						Expr: generateDec(shiftLeft),
					},
				},
			}

			j = nextJ
		}

		// shift expr to match the bit position in encoded data (concerning 'i')
		shiftLeft := i % 8
		// expr = (expr << i%8) & dataMask
		expr = &definition.BinopExpr{
			Op: definition.ExprOp_BAND,
			Expr1: &definition.BinopExpr{
				Op:    definition.ExprOp_SHL,
				Expr1: expr,
				Expr2: &definition.RawExpr{
					Expr: generateDec(shiftLeft),
				},
			},
			Expr2: &definition.RawExpr{
				Expr: generateBin(dataMask),
			},
		}

		if cast {
			expr = &definition.CastExpr{
				ToType: &definition.BasicType{
					TypeTypeID: definition.TypeID_Uint8,
				},
				Expr1: expr,
			}
		}

		// generate encode expression
		exprStr, err := g.GenerateExpr(expr, "")
		if err != nil {
			panic(fmt.Errorf("internal error: %s", err))
		}

		// generate encode statement
		encodeImplData := map[string]any{
			"TyUint8":   typeMap[definition.TypeID_Uint8],
			"BytePos":   i / 8,
			"Operator":  operator,
			"FieldData": exprStr,
			"Dynamic":   structDynamic,
		}

		encodeStmt := util.ExecuteTemplate(fieldEncoderTemplate, "encodeImpl", nil, encodeImplData)
		encodeStmts = append(encodeStmts, encodeStmt)
	}
	return encodeStmts
}

// ==================== GenerateDecoder ====================

var decoderTemplate = `
{{- define "decodeField" -}}
                // {{ .Pos }} {{ .Field.GetFieldKind }}: {{ .Field }}
                {{- range $decodeStmt := .DecodeStmts }}
                    {{ $decodeStmt }}
                {{- end -}}
{{- end -}}

{{- define "decoder" -}}
{{- $structName := .StructDef.StructName -}}
                // Decoder: {{ $structName }}
                /**
                 * Decode buffer to {{ $structName }} object
                 * @function decode
                 * @memberof {{ .StructDef.StructBelongs.Package }}.{{ $structName }}
                 * @static
                 * @param {{"{"}}{{ .StructDef.StructBelongs.Package }}.I{{ $structName }}} obj {{ $structName }} object
                 * @param {Array | Uint8Array} data The buffer to decode from
                 * @param {Number} [start] Start position of buffer to decode
                 * @returns {Number} -1 if decode failed, otherwise number of bytes decoded
                 */
                {{ $structName }}.decode = function(obj, data, start) {
                    if (obj === undefined) return -1;
                    if (data === undefined) return -1;
                    if (start === undefined) start = 0;
                    {{- if .Dynamic }}
                    var offset = 0;
                    {{- end }}
                    {{- range $decodeStr := .DecodeStrs }}
                    {{ $decodeStr }}
                    {{- end }}
                    return {{ if .Dynamic }}offset + {{ end }}{{ calc .StructDef.StructBitSize "/" 8 }};
                };

                /**
                 * Decode buffer to {{ $structName }} object
                 * @function decode
                 * @memberof {{ .StructDef.StructBelongs.Package }}.{{ $structName }}
                 * @instance
                 * @param {Array | Uint8Array} data The buffer to decode from
                 * @param {Number} [start] Start position of buffer to decode
                 * @returns {Number} -1 if decode failed, otherwise number of bytes decoded
                 */
                {{ $structName }}.prototype.decode = function(data, start) {
                    return {{ $structName }}.decode(this, data, start);
                };
{{- end -}}
`

func (g CommonJSGenerator) GenerateDecoder(structDef *definition.Struct) (string, error) {
	decodeStrs := []string{}
	if err := structDef.ForEachFieldWithPos(func(field definition.Field, fieldIndex int, startBits int64, dynamic bool, pos string) error {
		decodeStmts, err := g.generateDecodeField(field, startBits)
		if err != nil {
			return err
		}
		// skip
		if len(decodeStmts) == 0 {
			return nil
		}
		// remove blank lines
		filteredDecodeStmts := []string{}
		for _, stmt := range decodeStmts {
			if stmt != "" {
				filteredDecodeStmts = append(filteredDecodeStmts, stmt)
			}
		}
		decodeStmts = filteredDecodeStmts

		decodeFieldData := map[string]any{
			"Pos":         pos,
			"Field":       field,
			"DecodeStmts": decodeStmts,
		}

		str := util.ExecuteTemplate(decoderTemplate, "decodeField", nil, decodeFieldData)
		decodeStrs = append(decodeStrs, str)
		return nil
	}); err != nil {
		return "", err
	}

	fieldData := map[string]any{
		"StructDef":  structDef,
		"DecodeStrs": decodeStrs,
		"GenOption":  g.GenCtx.GenOptions,
		"Dynamic":    structDef.GetTypeDynamic(),
	}

	decoderStr := util.ExecuteTemplate(decoderTemplate, "decoder", nil, fieldData)
	return decoderStr, nil
}

func (g CommonJSGenerator) generateDecodeField(field definition.Field, startBits int64) ([]string, error) {
	switch val := field.(type) {
	case *definition.ConstantField:
		return g.generateDecodeConstantField(val, startBits)
	case *definition.VoidField:
		return g.generateDecodeVoidField(val, startBits)
	case *definition.EmbeddedField:
		return g.generateDecodeEmbeddedField(val, startBits)
	case *definition.NormalField:
		return g.generateDecodeNormalField(val, startBits)
	default:
		return nil, fmt.Errorf("internal error: unknown field kind %T", val)
	}
}

var fieldDecoderTemplate = `
{{- define "decodeTempVarName" -}}
    tempFieldAtPos{{ .StartBits }}
{{- end -}}

{{- define "decodeStructFieldName" -}}
{{- $fieldName := TocamelCase .FieldName -}}
    obj.{{ $fieldName }}
{{- end -}}

{{- define "decodeConstantField" -}}
    if ({{ .TempName }} !== {{ .ConstantValue }}) return -1;
{{- end -}}

{{- define "decodeNormalFieldStruct" -}}
{{- $packagePrefix := call .GenerateStructPackagePrefix .FieldStruct -}}
{{- if .FieldStruct.GetTypeDynamic -}}
                    (function() {
                        var {{ .TempName }} = {{ $packagePrefix }}.decode({{ .FieldName }}, data, offset + start + {{ .FromByte }});
                        if ({{ .TempName }} < 0) return -1;
                        offset += {{ .TempName }};
                    })();
{{- else -}}
    if ({{ $packagePrefix }}.decode({{ .FieldName }}, data, {{ if .Dynamic }}offset + {{ end }}start + {{ .FromByte }}) < 0) return -1;
{{- end -}}
{{- end -}}

{{- define "decodeNormalFieldTempVarDecl" -}}
    var {{ .TempName }};
{{- end -}}

{{- define "decodeNormalFieldTempVarAssignEnum" -}}
{{- $inPackage := .FieldDef.FieldBelongs.StructBelongs.LocalNames.Has .EnumDef.EnumName -}}
{{- if not $inPackage -}}
    {{ .FieldName }} = $root.{{ .EnumDef.EnumBelongs.Package }}.{{ .EnumDef.EnumName }}[{{ .TempName }}];
{{- else -}}
    {{ .FieldName }} = $package.{{ .EnumDef.EnumName }}[{{ .TempName }}];
{{- end -}}
{{- end -}}

{{- define "decodeNormalFieldTempVarAssignFloatCast" -}}
{{- if .IsFloat32 -}}
    {{ .FieldName }} = uint32BitsToFloat({{ .TempName }});
{{- else -}}
    {{ .FieldName }} = uint64BitsToDouble({{ .TempName }});
{{- end -}}
{{- end -}}

{{- define "decodeNormalFieldString" -}}
                    (function() {
                        var result = stringFromUTF8Bytes(data, offset + start + {{ .FromByte }});
                        {{ .FieldName }} = result[0];
                        offset += result[1] + 1;
                    })();
{{- end -}}

{{- define "decodeNormalFieldBytes" -}}
                    (function() {
                        var {{ .TempName }} = 0;
                        var shift = 0;
                        while ((data[offset + start + {{ .FromByte }}] & {{ .SetMask }}) !== 0) { {{ .TempName }} |= (data[offset + start + {{ .FromByte }}] & {{ .GetMask }}) << shift; shift += {{ .Shift }}; offset++; }
                        {{ .TempName }} |= (data[offset + start + {{ .FromByte }}] & {{ .GetMask }}) << shift; offset++;
                        {{ .FieldName }} = new Array({{ .TempName }});
                        for (var i = 0; i < {{ .TempName }}; i++) {{ .FieldName }}[i] = data[offset + start + {{ .FromByte }} + i];
                        offset += {{ .TempName }};
                    })();
{{- end -}}

{{- define "decodeData" -}}
    data[{{ if .Dynamic }}offset + {{ end }}start + {{ .BytePos }}]
{{- end -}}

{{- define "decodeImpl" -}}
    {{ .FieldName }} {{ .Operator }} {{ .Expr }};
{{- end -}}

{{- define "signExtendArith" -}}
    {{ .FieldName }} = (({{ .FieldName }}) ^ {{ .SignMask }}) - {{ .SignMask }};
{{- end -}}

{{- define "signExtendLogic" -}}
    {{ .FieldName }} = ({{ .FieldName }} >>> {{ .SignShift }}) ? -(~{{ .FieldName }} + 1) : {{ .FieldName }};
{{- end -}}
`

func (g CommonJSGenerator) generateDecodeTempVarName(startBits int64) string {
	decoderTempVarNameData := map[string]any{
		"StartBits": startBits,
	}

	return util.ExecuteTemplate(fieldDecoderTemplate, "decodeTempVarName", nil, decoderTempVarNameData)
}

func (g CommonJSGenerator) generateDecodeStructFieldName(name string) string {
	decodeStructFieldNameData := map[string]any{
		"FieldName": name,
	}

	return util.ExecuteTemplate(fieldDecoderTemplate, "decodeStructFieldName", nil, decodeStructFieldNameData)
}

func (g CommonJSGenerator) generateDecodeConstantField(field *definition.ConstantField, startBits int64) ([]string, error) {
	structDynamic := field.FieldBelongs.GetTypeDynamic()
	decodeStmts := []string{}

	tempName := g.generateDecodeTempVarName(startBits)
	decodeNormalFieldTempVarDeclData := map[string]any{
		"TempName": tempName,
	}
	declStr := util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldTempVarDecl", nil, decodeNormalFieldTempVarDeclData)
	decodeStmts = append(decodeStmts, declStr)

	from := startBits
	to := startBits + field.GetFieldBitSize()

	stmts, err := g.generateDecodeNormalFieldImpl(tempName, field.FieldType, field.FieldOptions, structDynamic, from, to)
	if err != nil {
		return nil, err
	}

	decodeStmts = append(decodeStmts, stmts...)

	literalValue, err := NewCommonJSLiteralGenerator().GenerateLiteral(field.FieldConstant)
	if err != nil {
		return nil, err
	}

	decodeConstantFieldData := map[string]any{
		"TempName":      tempName,
		"ConstantValue": literalValue,
	}

	checkStr := util.ExecuteTemplate(fieldDecoderTemplate, "decodeConstantField", nil, decodeConstantFieldData)
	decodeStmts = append(decodeStmts, checkStr)

	return decodeStmts, nil
}

func (g CommonJSGenerator) generateDecodeVoidField(field *definition.VoidField, startBits int64) ([]string, error) {
	return []string{""}, nil
}

func (g CommonJSGenerator) generateDecodeEmbeddedField(field *definition.EmbeddedField, startBits int64) ([]string, error) {
	return nil, nil
}

func (g CommonJSGenerator) generateDecodeNormalField(field *definition.NormalField, startBits int64) ([]string, error) {
	from := startBits
	to := startBits + field.GetFieldBitSize()
	structDynamic := field.FieldBelongs.GetTypeDynamic()
	decodeStmts := []string{}

	switch ty := field.FieldType.(type) {
	case *definition.Struct, *definition.String, *definition.Bytes:
		name := g.generateDecodeStructFieldName(field.FieldName)
		stmts, err := g.generateDecodeNormalFieldImpl(name, ty, field.FieldOptions, structDynamic, from, to)
		if err != nil {
			return nil, err
		}
		decodeStmts = append(decodeStmts, stmts...)

	case *definition.BasicType:
		if ty.GetTypeID().IsFloat() {
			// any integer type is ok
			tempTy := &definition.Uint32
			if ty.GetTypeID() == definition.TypeID_Float64 {
				tempTy = &definition.Uint64
			}
			tempName := g.generateDecodeTempVarName(startBits)

			decodeNormalFieldTempVarDeclData := map[string]any{
				"TempName":  tempName,
				"FieldName": g.generateDecodeStructFieldName(field.FieldName),
			}
			declStr := util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldTempVarDecl", nil, decodeNormalFieldTempVarDeclData)
			decodeStmts = append(decodeStmts, declStr)

			stmts, err := g.generateDecodeNormalFieldImpl(tempName, tempTy, field.FieldOptions, structDynamic, from, to)
			if err != nil {
				return nil, err
			}
			decodeStmts = append(decodeStmts, stmts...)

			decodeNormalFieldTempVarAssignFloatCastData := map[string]any{
				"FieldName": g.generateDecodeStructFieldName(field.FieldName),
				"TempName":  tempName,
				"IsFloat32": ty.GetTypeID() == definition.TypeID_Float32,
			}

			assignStr := util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldTempVarAssignFloatCast", nil, decodeNormalFieldTempVarAssignFloatCastData)
			decodeStmts = append(decodeStmts, assignStr)
		} else {
			// same as struct
			name := g.generateDecodeStructFieldName(field.FieldName)
			stmts, err := g.generateDecodeNormalFieldImpl(name, ty, field.FieldOptions, structDynamic, from, to)
			if err != nil {
				return nil, err
			}
			decodeStmts = append(decodeStmts, stmts...)
		}

	case *definition.Enum:
		// any integer type is ok (except 64-bit)
		tempTy := &definition.Uint32
		tempName := g.generateDecodeTempVarName(startBits)

		decodeNormalFieldTempVarDeclOnlyData := map[string]any{
			"TempName": tempName,
		}
		declStr := util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldTempVarDecl", nil, decodeNormalFieldTempVarDeclOnlyData)
		decodeStmts = append(decodeStmts, declStr)

		stmts, err := g.generateDecodeNormalFieldImpl(tempName, tempTy, field.FieldOptions, structDynamic, from, to)
		if err != nil {
			return nil, err
		}
		decodeStmts = append(decodeStmts, stmts...)

		decodeNormalFieldTempVarAssignEnumData := map[string]any{
			"FieldName": g.generateDecodeStructFieldName(field.FieldName),
			"FieldDef":  field,
			"EnumDef":   ty,
			"TempName":  tempName,
		}

		assignStr := util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldTempVarAssignEnum", nil, decodeNormalFieldTempVarAssignEnumData)
		decodeStmts = append(decodeStmts, assignStr)

	case *definition.Array:
		elemTy := ty.ElementType
		elemBitSize := field.FieldBitSize / ty.Length

		name := g.generateDecodeStructFieldName(field.FieldName)

		// temp variable declaration
		var nameIndex func(int64) string
		switch ty.ElementType.(type) {
		case *definition.Struct, *definition.String, *definition.Bytes:
			nameIndex = func(index int64) string {
				return fmt.Sprintf("%s[%d]", name, index)
			}
		case *definition.BasicType:
			nameIndex = func(index int64) string {
				return fmt.Sprintf("%s[%d]", name, index)
			}
			// same as enum
			if ty.ElementType.GetTypeID().IsFloat() {
				tempName := g.generateDecodeTempVarName(startBits)

				decodeNormalFieldTempVarDeclOnlyData := map[string]any{
					"TempName": tempName,
				}
				declStr := util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldTempVarDecl", nil, decodeNormalFieldTempVarDeclOnlyData)
				decodeStmts = append(decodeStmts, declStr)

				nameIndex = func(_ int64) string {
					return tempName
				}

				// change elemTy to 32 or 64 bit integer type
				switch ty.ElementType.GetTypeID() {
				case definition.TypeID_Float32:
					elemTy = &definition.Uint32
				case definition.TypeID_Float64:
					elemTy = &definition.Uint64
				}
			}
		case *definition.Enum:
			tempName := g.generateDecodeTempVarName(startBits)

			decodeNormalFieldTempVarDeclOnly := map[string]any{
				"TempName": tempName,
			}
			declStr := util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldTempVarDecl", nil, decodeNormalFieldTempVarDeclOnly)
			decodeStmts = append(decodeStmts, declStr)

			nameIndex = func(_ int64) string {
				return tempName
			}

			// change elemTy to any integer type (except 64-bit)
			elemTy = &definition.Uint32
		default:
			return nil, fmt.Errorf("internal error: unsupported array element type %T", ty.ElementType)
		}

		for i := int64(0); i < ty.Length; i++ {
			subFrom := from + i*elemBitSize
			subTo := from + (i+1)*elemBitSize
			subName := nameIndex(i)

			stmts, err := g.generateDecodeNormalFieldImpl(subName, elemTy, field.FieldOptions, structDynamic, subFrom, subTo)
			if err != nil {
				return nil, err
			}
			decodeStmts = append(decodeStmts, stmts...)

			switch elemTy := ty.ElementType.(type) {
			case *definition.BasicType:
				if elemTy.GetTypeID().IsFloat() {
					decodeNormalFieldTempVarAssignFloatCastData := map[string]any{
						"TempName":  subName,
						"FieldName": fmt.Sprintf("%s[%d]", name, i),
						"IsFloat32": elemTy.GetTypeID() == definition.TypeID_Float32,
					}
					assignStr := util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldTempVarAssignFloatCast", nil, decodeNormalFieldTempVarAssignFloatCastData)
					decodeStmts = append(decodeStmts, assignStr)
				}
			case *definition.Enum:
				decodeNormalFieldTempVarAssignEnumData := map[string]any{
					"TempName":  subName,
					"EnumDef":   elemTy,
					"FieldName": fmt.Sprintf("%s[%d]", name, i),
					"FieldDef":  field,
				}
				assignStr := util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldTempVarAssignEnum", nil, decodeNormalFieldTempVarAssignEnumData)
				decodeStmts = append(decodeStmts, assignStr)
			default:
			}
		}

	default:
		return nil, fmt.Errorf("internal error: unknown field kind %T", ty)
	}

	return decodeStmts, nil
}

func (g CommonJSGenerator) generateDecodeNormalFieldImpl(fieldNameStr string, fieldType definition.Type, fieldOptions *util.OrderedMap[string, *definition.Option], structDynamic bool, from, to int64) ([]string, error) {
	decodeStmts := []string{}
	fieldBitSize := to - from

	dataDataFunc := func(i int64) string {
		decodeDataData := map[string]any{
			"Dynamic": structDynamic,
			"BytePos": i,
		}
		return util.ExecuteTemplate(fieldDecoderTemplate, "decodeData", nil, decodeDataData)
	}

	switch ty := fieldType.(type) {
	case *definition.Struct:
		decodeNormalFieldStructData := map[string]any{
			"FieldStruct":                 ty,
			"FieldName":                   fieldNameStr,
			"FromByte":                    from / 8,
			"ToByte":                      (to + 7) / 8,
			"Dynamic":                     structDynamic,
			"TempName":                    g.generateDecodeTempVarName(from),
			"GenerateStructPackagePrefix": g.generateStructPackagePrefix,
		}

		stmt := util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldStruct", nil, decodeNormalFieldStructData)
		decodeStmts = append(decodeStmts, stmt)

	case *definition.Enum:
		panic("unreachable, enum field should be handled in generateDecodeNormalField")

	case *definition.String:
		decodeNormalFieldStringData := map[string]any{
			"FieldName": fieldNameStr,
			"FromByte":  from / 8,
		}

		stmt := util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldString", nil, decodeNormalFieldStringData)
		decodeStmts = append(decodeStmts, stmt)

	case *definition.Bytes:
		decodeNormalFieldBytesData := map[string]any{
			"FieldName": fieldNameStr,
			"FromByte":  from / 8,
			"GetMask":   g.generateHex((1 << 7) - 1),
			"SetMask":   g.generateHex(1 << 7),
			"Shift":     7,
			"TempName":  g.generateDecodeTempVarName(from),
		}

		stmt := util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldBytes", nil, decodeNormalFieldBytesData)
		decodeStmts = append(decodeStmts, stmt)

	case *definition.BasicType:
		generateBin := g.generateBin
		generateDec := g.generateDec
		generateHex := g.generateHex
		if ty.GetTypeBitSize() > 32 {
			generateBin = g.generateBinBigInt
			generateDec = g.generateDecBigInt
			generateHex = g.generateHexBigInt
			oldDataDataFunc := dataDataFunc
			dataDataFunc = func(i int64) string {
				expr := &definition.CastExpr{
					ToType: &definition.Uint64,
					Expr1: &definition.RawExpr{
						Expr: oldDataDataFunc(i),
					},
				}
				exprStr, err := g.GenerateExpr(expr, "")
				if err != nil {
					panic(fmt.Errorf("internal error: %s", err))
				}
				return exprStr
			}
		}
		// little endian as default
		shiftCalc := func(i int64) string {
			return generateDec(i * 8)
		}
		if gen.MatchOption(fieldOptions, "order", "big") {
			// big endian
			shiftCalc = func(i int64) string {
				return generateDec(max(0, fieldBitSize-8*(i+1)))
			}
		}

		fieldProcessor := func(expr string, i int64) string {
			// operator is '=' if is filling the first bit, otherwise is '|='
			operator := ""
			if i == 0 {
				operator = exprOpToString[definition.ExprOp_ASSIGN] // "="
			} else {
				operator = exprOpToString[definition.ExprOp_BOR] + exprOpToString[definition.ExprOp_ASSIGN] // "|="
			}

			// exprExp = expr << 8*i or expr << max(0, fieldBitSize-8*(i+1))
			exprExp := &definition.BinopExpr{
				Op: definition.ExprOp_SHL,
				Expr1: &definition.RawExpr{
					Expr: expr,
				},
				Expr2: &definition.RawExpr{
					Expr: shiftCalc(i),
				},
			}

			exprExpStr, err := g.GenerateExpr(exprExp, "")
			if err != nil {
				panic(fmt.Errorf("internal error: %s", err))
			}

			decodeImplData := map[string]any{
				"FieldName": fieldNameStr,
				"Operator":  operator,
				"Expr":      exprExpStr,
			}

			return util.ExecuteTemplate(fieldDecoderTemplate, "decodeImpl", nil, decodeImplData)
		}

		decodeStmts = append(decodeStmts, g.generateDecodeImpl(from, to, fieldProcessor, dataDataFunc, generateBin, generateDec)...)

		// set sign bit
		if ty.GetTypeID().IsInt() {
			originFromBitSize := fieldBitSize

			switch g.GenCtx.GenOptions.SignExtMethod {
			case gen.SignExtMethodDefault, gen.SignExtMethodShift, gen.SignExtMethodArith:
				// special case: sign extension for 32-bit integer
				if originFromBitSize == 32 {
					signShift := originFromBitSize - 1
					signShiftStr := generateDec(signShift)
					signExtendData := map[string]any{
						"FieldName": fieldNameStr,
						"SignShift": signShiftStr,
					}

					stmt := util.ExecuteTemplate(fieldDecoderTemplate, "signExtendLogic", nil, signExtendData)
					decodeStmts = append(decodeStmts, stmt)
				} else {
					signMask := int64(1) << (originFromBitSize - 1)
					signMaskStr := generateHex(signMask)
					signExtendData := map[string]any{
						"FieldName": fieldNameStr,
						"SignMask":  signMaskStr,
					}

					stmt := util.ExecuteTemplate(fieldDecoderTemplate, "signExtendArith", nil, signExtendData)
					decodeStmts = append(decodeStmts, stmt)
				}

			default:
				panic("unreachable, unknown sign extension method")
			}
		}

	case *definition.Array:
		panic("unreachable, array field should be handled in generateDecodeNormalField")

	default:
		return nil, fmt.Errorf("internal error: unknown field kind %T", ty)
	}

	return decodeStmts, nil
}

// generateDecodeImpl generates decode implementation from 'from' bit to 'to' bit and align to 8 bits
// from: start bit position of encoded data
// to: end bit position of encoded data
// from is inclusive, to is exclusive, i.e. [from, to)
//
// e.g. from = 3, to = 19
//
//	exprOfExtract1stByteFromEncodedData = (((((uint8_t*)data)[0] & 0b11111000) >> 3) | ((((uint8_t*)data)[1] & 0b00000111) << 5))
//	exprOfExtract2ndByteFromEncodedData = (((((uint8_t*)data)[1] & 0b11111000) >> 3) | ((((uint8_t*)data)[2] & 0b00000111) << 5))
//
// fieldData: a function callback to generate expression of decoding x-th byte data (x is the byte index)
//
// e.g. little endian
//
//	fieldProcessor(exprOfExtract1stByteFromEncodedData, 0) -> (*(uint32_t*)(&(structPtr->intField))) = (exprOfExtract1stByteFromEncodedData << 0)
//	fieldProcessor(exprOfExtract2ndByteFromEncodedData, 1) -> (*(uint32_t*)(&(structPtr->intField))) |= (exprOfExtract2ndByteFromEncodedData << 8)
//	fieldProcessor(exprOfExtract3rdByteFromEncodedData, 2) -> (*(uint32_t*)(&(structPtr->intField))) |= (exprOfExtract3rdByteFromEncodedData << 16)
//	fieldProcessor(exprOfExtract4thByteFromEncodedData, 3) -> (*(uint32_t*)(&(structPtr->intField))) |= (exprOfExtract4thByteFromEncodedData << 24)
func (g CommonJSGenerator) generateDecodeImpl(from, to int64, fieldProcessor func(string, int64) string, dataData func(int64) string, generateBin func(any) string, generateDec func(any) string) []string {
	decodeStmts := []string{}
	// generate decode implentation from 'from' bit to 'to' bit per 8 bits
	// e.g. from = 3, to = 19 -> loop 2 times: 3-10, 11-19 (not aligned to 8 bits!!!)
	for i := from; i < to; i += 8 {

		// we use 'from' and 'to' to denote the bit position in encoded data
		begin := i
		end := min(to, i+8)

		var expr definition.Expr

		// separator to check if is aligned to 8 bits
		sep := min(end, (begin+8)&(^7))
		// first half
		// e.g. begin = 3, end = 10
		//      sep = 8, fieldMask = 0b11111000, shiftRight = 3
		if begin < sep { // always true, just for beauty
			fieldMask := ((1 << (((sep - 1) & 7) + 1)) - 1) & (^((1 << (begin & 7)) - 1))
			shiftRight := begin % 8
			// expr = (((data[begin/8] & fieldMask) >> shiftRight)
			expr = &definition.BinopExpr{
				Op: definition.ExprOp_SHR,
				Expr1: &definition.BinopExpr{
					Op: definition.ExprOp_BAND,
					Expr1: &definition.RawExpr{
						Expr: dataData(begin / 8),
					},
					Expr2: &definition.RawExpr{
						Expr: generateBin(fieldMask),
					},
				},
				Expr2: &definition.RawExpr{
					Expr: generateDec(shiftRight),
				},
			}
		}
		// second half
		// e.g. begin = 8, end = 10
		//      sep = 8, fieldMask = 0b00000111, shiftLeft = 5
		if sep < end {
			fieldMask := ((1 << (((end - 1) & 7) + 1)) - 1) & (^((1 << (sep & 7)) - 1))
			shiftLeft := 8 - end%8
			// expr = expr | (((data[sep/8] & fieldMask) << shiftLeft)
			expr = &definition.BinopExpr{
				Op:    definition.ExprOp_BOR,
				Expr1: expr,
				Expr2: &definition.BinopExpr{
					Op: definition.ExprOp_SHL,
					Expr1: &definition.BinopExpr{
						Op: definition.ExprOp_BAND,
						Expr1: &definition.RawExpr{
							Expr: dataData(sep / 8),
						},
						Expr2: &definition.RawExpr{
							Expr: generateBin(fieldMask),
						},
					},
					Expr2: &definition.RawExpr{
						Expr: generateDec(shiftLeft),
					},
				},
			}
		}

		// generate decode expression
		exprStr, err := g.GenerateExpr(expr, "")
		if err != nil {
			panic(fmt.Errorf("internal error: %s", err))
		}

		// generate decode statement
		decodeStmt := fieldProcessor(exprStr, (i-from)/8)
		decodeStmts = append(decodeStmts, decodeStmt)
	}
	return decodeStmts
}

// ==================== GenerateExpr ====================

func (g CommonJSGenerator) GenerateExpr(expr definition.Expr, valueStr string) (string, error) {
	generator := NewCommonJSExprGenerator(g.GenerateType, valueStr)
	return g.AcceptExpr(expr, generator)
}

// ==================== Expr Generator ====================

type CommonJSExprGenerator struct {
	*gen.GenExprDispatcher
	GenType          func(definition.Type) (string, error)
	ValueStr         string
	LiteralGenerator gen.LiteralGeneratorImpl // optional
}

func NewCommonJSExprGenerator(genType func(definition.Type) (string, error), valueStr string) *CommonJSExprGenerator {
	generator := &CommonJSExprGenerator{
		GenExprDispatcher: nil,
		GenType:           genType,
		ValueStr:          valueStr,
	}
	generator.GenExprDispatcher = gen.NewGenExprDispatcher(generator)
	return generator
}

func (g CommonJSExprGenerator) GenerateExpr(expr definition.Expr) (string, error) {
	return g.AcceptExpr(expr)
}

var exprOpToString = map[definition.ExprOp]string{
	definition.ExprOp_ADD:    "+",
	definition.ExprOp_SUB:    "-",
	definition.ExprOp_MUL:    "*",
	definition.ExprOp_DIV:    "/",
	definition.ExprOp_MOD:    "%",
	definition.ExprOp_POW:    "**",
	definition.ExprOp_SHL:    "<<",
	definition.ExprOp_SHR:    ">>",
	definition.ExprOp_LT:     "<",
	definition.ExprOp_LE:     "<=",
	definition.ExprOp_GT:     ">",
	definition.ExprOp_GE:     ">=",
	definition.ExprOp_EQ:     "==",
	definition.ExprOp_NE:     "!=",
	definition.ExprOp_BAND:   "&",
	definition.ExprOp_BXOR:   "^",
	definition.ExprOp_BOR:    "|",
	definition.ExprOp_AND:    "&&",
	definition.ExprOp_OR:     "||",
	definition.ExprOp_NOT:    "!",
	definition.ExprOp_BNOT:   "~",
	definition.ExprOp_ASSIGN: "=",
}

func (g CommonJSExprGenerator) GenerateUnopExpr(expr *definition.UnopExpr) (string, error) {
	opStr, ok := exprOpToString[expr.Op]
	if !ok {
		return "", fmt.Errorf("unknown unop expr op: %s", expr.Op.String())
	}
	expr1, err := g.GenerateExpr(expr.Expr1)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s%s)", opStr, expr1), nil
}

func (g CommonJSExprGenerator) GenerateBinopExpr(expr *definition.BinopExpr) (string, error) {
	opStr, ok := exprOpToString[expr.Op]
	if !ok {
		return "", fmt.Errorf("unknown binop expr op: %s", expr.Op.String())
	}
	expr1, err := g.GenerateExpr(expr.Expr1)
	if err != nil {
		return "", err
	}
	expr2, err := g.GenerateExpr(expr.Expr2)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s %s %s)", expr1, opStr, expr2), nil
}

func (g CommonJSExprGenerator) GenerateCastExpr(expr *definition.CastExpr) (string, error) {
	expr1, err := g.GenerateExpr(expr.Expr1)
	if err != nil {
		return "", err
	}
	ty, err := g.GenType(expr.ToType)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s(%s)", ty, expr1), nil
}

func (g CommonJSExprGenerator) GenerateConstantExpr(expr *definition.ConstantExpr) (string, error) {
	generator := g.LiteralGenerator
	if generator == nil {
		generator = NewCommonJSLiteralGenerator()
	}
	return g.AcceptLiteral(expr.ConstantValue, generator)
}

func (g CommonJSExprGenerator) GenerateTenaryExpr(expr *definition.TenaryExpr) (string, error) {
	cond, err := g.GenerateExpr(expr.Cond)
	if err != nil {
		return "", err
	}
	expr1, err := g.GenerateExpr(expr.Expr1)
	if err != nil {
		return "", err
	}
	expr2, err := g.GenerateExpr(expr.Expr2)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s ? %s : %s)", cond, expr1, expr2), nil
}

func (g CommonJSExprGenerator) GenerateValueExpr(expr *definition.ValueExpr) (string, error) {
	return g.ValueStr, nil
}

func (g CommonJSExprGenerator) GenerateRawExpr(expr *definition.RawExpr) (string, error) {
	return expr.Expr, nil
}

// ==================== Literal Generator ====================

type CommonJSLiteralGenerator struct {
	*gen.GenLiteralDispatcher
}

func NewCommonJSLiteralGenerator() *CommonJSLiteralGenerator {
	generator := &CommonJSLiteralGenerator{
		GenLiteralDispatcher: nil,
	}
	generator.GenLiteralDispatcher = gen.NewGenLiteralDispatcher(generator)
	return generator
}

func (g CommonJSLiteralGenerator) GenerateLiteral(literal definition.Literal) (string, error) {
	return g.AcceptLiteral(literal)
}

func (g CommonJSLiteralGenerator) GenerateBoolLiteral(literal *definition.BoolLiteral) (string, error) {
	return fmt.Sprintf("%t", literal.BoolValue), nil
}

func (g CommonJSLiteralGenerator) GenerateIntLiteral(literal *definition.IntLiteral) (string, error) {
	if literal.IntValue > math.MaxInt32 || literal.IntValue < math.MinInt32 {
		return generateDecBigInt(literal.IntValue), nil
	}
	return generateDec(literal.IntValue), nil
}

func (g CommonJSLiteralGenerator) GenerateFloatLiteral(literal *definition.FloatLiteral) (string, error) {
	return fmt.Sprintf("%f", literal.FloatValue), nil
}

func (g CommonJSLiteralGenerator) GenerateStringLiteral(literal *definition.StringLiteral) (string, error) {
	return fmt.Sprintf(`"%s"`, literal.StringValue), nil
}
