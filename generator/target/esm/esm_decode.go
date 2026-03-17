package esm

import (
	"fmt"
	"strings"

	"github.com/xaxys/bubbler/definition"
	"github.com/xaxys/bubbler/generator/gen"
	"github.com/xaxys/bubbler/util"
)

// ==================== GenerateDecoder ====================

var decoderTemplate = `
{{- define "decodeField" -}}
// {{ .Pos }} {{ .Field.GetFieldKind }}: {{ .Field }}
{{- range $decodeStmt := .DecodeStmts }}
    {{ $decodeStmt }}
{{- end -}}
{{- end -}}

{{- define "decoderFunc" -}}
{{- $structName := .StructDef.StructName -}}
/**
 * Decode buffer to {{ $structName }} object.
 * @function decode
 * @memberof {{ .StructDef.StructBelongs.Package }}.{{ $structName }}
 * @static
 * @param {Object} obj {{ $structName }} object.
 * @param {Array|Uint8Array} data The buffer to decode from.
 * @param {Number} [start] Start position of buffer to decode.
 * @returns {Number} -1 if decode failed, otherwise number of bytes decoded.
 */
static decode(obj, data, start) {
    if (obj === undefined) return -1;
    if (data === undefined) return -1;
    if (start === undefined) start = 0;
    {{- if .Dynamic }}
    let offset = 0;
    {{- end }}
    {{- range $decodeStr := .DecodeStrs }}
    {{ $decodeStr }}
    {{- end }}
    return {{ if .Dynamic }}offset + {{ end }}{{ calc .StructDef.StructBitSize "/" 8 }};
}

/**
 * Decode buffer to {{ $structName }} object.
 * @function decode
 * @memberof {{ .StructDef.StructBelongs.Package }}.{{ $structName }}
 * @instance
 * @param {Array|Uint8Array} data The buffer to decode from.
 * @param {Number} [start] Start position of buffer to decode.
 * @returns {Number} -1 if decode failed, otherwise number of bytes decoded.
 */
decode(data, start) {
    return {{ $structName }}.decode(this, data, start);
}
{{- end -}}

{{- define "decoder" -}}
{{ template "decoderSize" . }}

{{ template "decoderFunc" . }}
{{- end -}}

{{- define "decoderSizeField" -}}
{{- $fromByte := .FromByte -}}
{{- $f := .Field -}}
{{- if $f.FieldType.GetTypeID.IsArray -}}
    {{- if $f.FieldType.ElementType.GetTypeID.IsString -}}
        {{- range $i := iterate 0 $f.FieldType.Length }}
    {   // {{ $f }}: [{{ $i }}]
        if (data.length - start <= offset + {{ $fromByte }}) return -(offset + {{ $fromByte }} + 1);
        let _length = 0;
        while (data[offset + start + {{ $fromByte }} + _length] !== 0) {
            _length++;
            if (data.length - start <= offset + {{ $fromByte }} + _length) { return -(offset + {{ $fromByte }} + _length + 1); }
        }
        offset += _length + 1;
    }
        {{- end -}}
    {{- else if $f.FieldType.ElementType.GetTypeID.IsBytes -}}
        {{- range $i := iterate 0 $f.FieldType.Length }}
    {   // {{ $f }}: [{{ $i }}]
        if (data.length - start <= offset + {{ $fromByte }}) { return -(offset + {{ $fromByte }} + 1); }
        let _length = 0;
        let _shift = 0;
        while ((data[offset + start + {{ $fromByte }}] & 0x80) !== 0) {
            _length |= (data[offset + start + {{ $fromByte }}] & 0x7F) << _shift;
            _shift += 7;
            offset++;
            if (data.length - start <= offset + {{ $fromByte }}) { return -(offset + {{ $fromByte }} + 1); }
        }
        _length |= (data[offset + start + {{ $fromByte }}] & 0x7F) << _shift;
        offset++;
        if (data.length - start < offset + {{ $fromByte }} + _length) return -(offset + {{ $fromByte }} + _length);
        offset += _length;
    }
        {{- end -}}
    {{- else if $f.FieldType.ElementType.GetTypeID.IsStruct -}}
        {{- if $f.FieldType.ElementType.GetTypeDynamic -}}
            {{- $pkgPrefix := FieldStructPkgPrefix $f -}}
            {{- range $i := iterate 0 $f.FieldType.Length }}
    {   // {{ $f }}: [{{ $i }}]
        const _subSize = {{ $pkgPrefix }}.decode_size(data, offset + start + {{ $fromByte }});
        if (_subSize < 0) { return -(offset + {{ $fromByte }}) + _subSize; }
        offset += _subSize;
    }
            {{- end -}}
        {{- end -}}
    {{- end -}}
{{- else if $f.FieldType.GetTypeID.IsString }}
    {   // {{ $f }}
        if (data.length - start <= offset + {{ $fromByte }}) return -(offset + {{ $fromByte }} + 1);
        let _length = 0;
        while (data[offset + start + {{ $fromByte }} + _length] !== 0) {
            _length++;
            if (data.length - start <= offset + {{ $fromByte }} + _length) { return -(offset + {{ $fromByte }} + _length + 1); }
        }
        offset += _length + 1;
    }
{{- else if $f.FieldType.GetTypeID.IsBytes }}
    {   // {{ $f }}
        if (data.length - start <= offset + {{ $fromByte }}) { return -(offset + {{ $fromByte }} + 1); }
        let _length = 0;
        let _shift = 0;
        while ((data[offset + start + {{ $fromByte }}] & 0x80) !== 0) {
            _length |= (data[offset + start + {{ $fromByte }}] & 0x7F) << _shift;
            _shift += 7;
            offset++;
            if (data.length - start <= offset + {{ $fromByte }}) { return -(offset + {{ $fromByte }} + 1); }
        }
        _length |= (data[offset + start + {{ $fromByte }}] & 0x7F) << _shift;
        offset++;
        if (data.length - start < offset + {{ $fromByte }} + _length) return -(offset + {{ $fromByte }} + _length);
        offset += _length;
    }
{{- else if $f.FieldType.GetTypeID.IsStruct -}}
    {{- if $f.FieldType.GetTypeDynamic -}}
    {{- $pkgPrefix := FieldStructPkgPrefix $f -}}
    {   // {{ $f }}
        const _subSize = {{ $pkgPrefix }}.decode_size(data, offset + start + {{ $fromByte }});
        if (_subSize < 0) { return -(offset + {{ $fromByte }}) + _subSize; }
        offset += _subSize;
    }
    {{- end -}}
{{- end -}}
{{- end -}}

{{- define "decoderSize" -}}
{{- $structName := .StructDef.StructName -}}
{{- $structBytes := calc .StructDef.StructBitSize "/" 8 -}}
/**
 * Calculate size of {{ $structName }} from buffer.
 * @function decode_size
 * @description Returns the encoded size (> 0) if successful, or the negative minimum required size (< 0) if data is insufficient. The required size may change as more data is provided.
 * @memberof {{ .StructDef.StructBelongs.Package }}.{{ $structName }}
 * @static
 * @param {Array|Uint8Array} data The buffer to probe.
 * @param {Number} [start] Start position of buffer.
 * @returns {Number} Positive frame size if complete, negative minimum needed size if incomplete.
 */
static decode_size(data, start) {
    if (data === undefined) return -1;
    if (start === undefined) start = 0;
    {{- if .StructDef.StructDynamic }}
    let offset = 0;
    {{- $fixedStart := 0 }}
    {{- range $field := .StructDef.StructFields.Values }}
        {{- if and $field.GetFieldKind.IsNormal (lt $field.GetFieldBitSize 0) }}
    {{- template "decoderSizeField" (dict "Field" $field "FromByte" (calc $fixedStart "/" 8)) }}
        {{- end }}
        {{- if ne $field.GetFieldBitSize -1 }}
    {{- $fixedStart = calc $fixedStart "+" $field.GetFieldBitSize }}
        {{- end }}
    {{- end }}
    if (data.length - start < offset + {{ $structBytes }}) return -(offset + {{ $structBytes }});
    return offset + {{ $structBytes }};
    {{- else }}
    if (data.length - start < {{ $structBytes }}) return -({{ $structBytes }});
    return {{ $structBytes }};
    {{- end }}
}

/**
 * Calculate size of {{ $structName }} from buffer.
 * @function decode_size
 * @memberof {{ .StructDef.StructBelongs.Package }}.{{ $structName }}
 * @instance
 * @param {Array|Uint8Array} data The buffer to probe.
 * @param {Number} [start] Start position of buffer.
 * @returns {Number} Positive frame size if complete, negative minimum needed size if incomplete.
 */
decode_size(data, start) {
    return {{ $structName }}.decode_size(data, start);
}
{{- end -}}
`

func (g ESModuleGenerator) GenerateDecoder(structDef *definition.Struct) (string, error) {
	funcMap := map[string]any{
		"FieldStructPkgPrefix": func(f interface{}) string {
			if nf, ok := f.(*definition.NormalField); ok {
				if s, ok := nf.FieldType.(*definition.Struct); ok {
					return strings.TrimSpace(g.generateStructPackagePrefix(s))
				}
				if a, ok := nf.FieldType.(*definition.Array); ok {
					if s, ok := a.ElementType.(*definition.Struct); ok {
						return strings.TrimSpace(g.generateStructPackagePrefix(s))
					}
				}
			}
			return ""
		},
	}

	decodeStrs := []string{}
	if err := structDef.ForEachFieldWithPos(func(field definition.Field, fieldIndex int, startBits int64, dynamic bool, pos string) error {
		decodeStmts, err := g.generateDecodeField(field, startBits)
		if err != nil {
			return err
		}
		if len(decodeStmts) == 0 {
			return nil
		}
		filtered := []string{}
		for _, s := range decodeStmts {
			if s != "" {
				filtered = append(filtered, s)
			}
		}
		decodeFieldData := map[string]any{
			"Pos":         pos,
			"Field":       field,
			"DecodeStmts": filtered,
		}
		str := util.ExecuteTemplate(decoderTemplate, "decodeField", funcMap, decodeFieldData)
		decodeStrs = append(decodeStrs, str)
		return nil
	}); err != nil {
		return "", err
	}

	fieldData := map[string]any{
		"StructDef":  structDef,
		"DecodeStrs": decodeStrs,
		"GenOptions": g.GenCtx.GenOptions,
		"Dynamic":    structDef.GetTypeDynamic(),
	}
	return util.ExecuteTemplate(decoderTemplate, "decoder", funcMap, fieldData), nil
}

func (g ESModuleGenerator) generateDecodeField(field definition.Field, startBits int64) ([]string, error) {
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
        const {{ .TempName }} = {{ $packagePrefix }}.decode({{ .FieldName }}, data, offset + start + {{ .FromByte }});
        if ({{ .TempName }} < 0) return -1;
        offset += {{ .TempName }};
    })();
{{- else -}}
    if ({{ $packagePrefix }}.decode({{ .FieldName }}, data, {{ if .Dynamic }}offset + {{ end }}start + {{ .FromByte }}) < 0) return -1;
{{- end -}}
{{- end -}}

{{- define "decodeNormalFieldTempVarDecl" -}}
    let {{ .TempName }};
{{- end -}}

{{- define "decodeNormalFieldTempVarAssignEnum" -}}
{{- $inPackage := .FieldDef.FieldBelongs.StructBelongs.LocalNames.Has .EnumDef.EnumName -}}
{{- if or .SingleFile $inPackage -}}
    {{ .FieldName }} = {{ .EnumDef.EnumName }}[{{ .TempName }}];
{{- else -}}
    {{ .FieldName }} = ${{ .EnumDef.EnumBelongs.Package.PackageName }}.{{ .EnumDef.EnumName }}[{{ .TempName }}];
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
        const result = stringFromUTF8Bytes(data, offset + start + {{ .FromByte }});
        {{ .FieldName }} = result[0];
        offset += result[1] + 1;
    })();
{{- end -}}

{{- define "decodeNormalFieldBytes" -}}
    (function() {
        let {{ .TempName }} = 0;
        let shift = 0;
        while ((data[offset + start + {{ .FromByte }}] & {{ .SetMask }}) !== 0) { {{ .TempName }} |= (data[offset + start + {{ .FromByte }}] & {{ .GetMask }}) << shift; shift += {{ .Shift }}; offset++; }
        {{ .TempName }} |= (data[offset + start + {{ .FromByte }}] & {{ .GetMask }}) << shift; offset++;
        {{ .FieldName }} = new {{ if .GenOptions.CompatibleMode }}Array{{ else }}Uint8Array{{ end }}({{ .TempName }});
        for (let i = 0; i < {{ .TempName }}; i++) {{ .FieldName }}[i] = data[offset + start + {{ .FromByte }} + i];
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

func (g ESModuleGenerator) generateDecodeTempVarName(startBits int64) string {
	return util.ExecuteTemplate(fieldDecoderTemplate, "decodeTempVarName", nil, map[string]any{"StartBits": startBits})
}

func (g ESModuleGenerator) generateDecodeStructFieldName(name string) string {
	return util.ExecuteTemplate(fieldDecoderTemplate, "decodeStructFieldName", nil, map[string]any{"FieldName": name})
}

func (g ESModuleGenerator) generateDecodeConstantField(field *definition.ConstantField, startBits int64) ([]string, error) {
	structDynamic := field.FieldBelongs.GetTypeDynamic()
	decodeStmts := []string{}
	tempName := g.generateDecodeTempVarName(startBits)
	decodeStmts = append(decodeStmts, util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldTempVarDecl", nil, map[string]any{"TempName": tempName}))
	stmts, err := g.generateDecodeNormalFieldImpl(tempName, field.FieldType, field.FieldOptions, structDynamic, startBits, startBits+field.GetFieldBitSize())
	if err != nil {
		return nil, err
	}
	decodeStmts = append(decodeStmts, stmts...)
	literalValue, err := NewESModuleLiteralGenerator().GenerateLiteral(field.FieldConstant)
	if err != nil {
		return nil, err
	}
	checkStr := util.ExecuteTemplate(fieldDecoderTemplate, "decodeConstantField", nil, map[string]any{"TempName": tempName, "ConstantValue": literalValue})
	decodeStmts = append(decodeStmts, checkStr)
	return decodeStmts, nil
}

func (g ESModuleGenerator) generateDecodeVoidField(field *definition.VoidField, startBits int64) ([]string, error) {
	return []string{""}, nil
}

func (g ESModuleGenerator) generateDecodeEmbeddedField(field *definition.EmbeddedField, startBits int64) ([]string, error) {
	return nil, nil
}

func (g ESModuleGenerator) generateDecodeNormalField(field *definition.NormalField, startBits int64) ([]string, error) {
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
			tempTy := &definition.Uint32
			if ty.GetTypeID() == definition.TypeID_Float64 {
				tempTy = &definition.Uint64
			}
			tempName := g.generateDecodeTempVarName(startBits)
			decodeStmts = append(decodeStmts, util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldTempVarDecl", nil, map[string]any{"TempName": tempName}))
			stmts, err := g.generateDecodeNormalFieldImpl(tempName, tempTy, field.FieldOptions, structDynamic, from, to)
			if err != nil {
				return nil, err
			}
			decodeStmts = append(decodeStmts, stmts...)
			decodeStmts = append(decodeStmts, util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldTempVarAssignFloatCast", nil, map[string]any{
				"FieldName": g.generateDecodeStructFieldName(field.FieldName), "TempName": tempName, "IsFloat32": ty.GetTypeID() == definition.TypeID_Float32,
			}))
		} else {
			name := g.generateDecodeStructFieldName(field.FieldName)
			stmts, err := g.generateDecodeNormalFieldImpl(name, ty, field.FieldOptions, structDynamic, from, to)
			if err != nil {
				return nil, err
			}
			decodeStmts = append(decodeStmts, stmts...)
		}

	case *definition.Enum:
		tempTy := &definition.Uint32
		tempName := g.generateDecodeTempVarName(startBits)
		decodeStmts = append(decodeStmts, util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldTempVarDecl", nil, map[string]any{"TempName": tempName}))
		stmts, err := g.generateDecodeNormalFieldImpl(tempName, tempTy, field.FieldOptions, structDynamic, from, to)
		if err != nil {
			return nil, err
		}
		decodeStmts = append(decodeStmts, stmts...)
		decodeStmts = append(decodeStmts, util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldTempVarAssignEnum", nil, map[string]any{
			"FieldName": g.generateDecodeStructFieldName(field.FieldName), "FieldDef": field, "EnumDef": ty, "TempName": tempName,
			"SingleFile": g.GenCtx.GenOptions.SingleFile,
		}))

	case *definition.Array:
		elemTy := ty.ElementType
		elemBitSize := field.FieldBitSize / ty.Length
		name := g.generateDecodeStructFieldName(field.FieldName)

		var nameIndex func(int64) string
		switch ty.ElementType.(type) {
		case *definition.Struct, *definition.String, *definition.Bytes:
			nameIndex = func(index int64) string { return fmt.Sprintf("%s[%d]", name, index) }
		case *definition.BasicType:
			nameIndex = func(index int64) string { return fmt.Sprintf("%s[%d]", name, index) }
			if ty.ElementType.GetTypeID().IsFloat() {
				tempName := g.generateDecodeTempVarName(startBits)
				decodeStmts = append(decodeStmts, util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldTempVarDecl", nil, map[string]any{"TempName": tempName}))
				nameIndex = func(_ int64) string { return tempName }
				switch ty.ElementType.GetTypeID() {
				case definition.TypeID_Float32:
					elemTy = &definition.Uint32
				case definition.TypeID_Float64:
					elemTy = &definition.Uint64
				}
			}
		case *definition.Enum:
			tempName := g.generateDecodeTempVarName(startBits)
			decodeStmts = append(decodeStmts, util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldTempVarDecl", nil, map[string]any{"TempName": tempName}))
			nameIndex = func(_ int64) string { return tempName }
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
			switch elemTy2 := ty.ElementType.(type) {
			case *definition.BasicType:
				if elemTy2.GetTypeID().IsFloat() {
					decodeStmts = append(decodeStmts, util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldTempVarAssignFloatCast", nil, map[string]any{
						"TempName": subName, "FieldName": fmt.Sprintf("%s[%d]", name, i), "IsFloat32": elemTy2.GetTypeID() == definition.TypeID_Float32,
					}))
				}
			case *definition.Enum:
				decodeStmts = append(decodeStmts, util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldTempVarAssignEnum", nil, map[string]any{
					"TempName": subName, "EnumDef": elemTy2, "FieldName": fmt.Sprintf("%s[%d]", name, i), "FieldDef": field,
					"SingleFile": g.GenCtx.GenOptions.SingleFile,
				}))
			}
		}

	default:
		return nil, fmt.Errorf("internal error: unknown field kind %T", ty)
	}

	return decodeStmts, nil
}

func (g ESModuleGenerator) generateDecodeNormalFieldImpl(fieldNameStr string, fieldType definition.Type, fieldOptions *util.OrderedMap[string, *definition.Option], structDynamic bool, from, to int64) ([]string, error) {
	decodeStmts := []string{}
	fieldBitSize := to - from

	dataDataFunc := func(i int64) string {
		return util.ExecuteTemplate(fieldDecoderTemplate, "decodeData", nil, map[string]any{"Dynamic": structDynamic, "BytePos": i})
	}

	switch ty := fieldType.(type) {
	case *definition.Struct:
		stmt := util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldStruct", nil, map[string]any{
			"FieldStruct": ty, "FieldName": fieldNameStr, "FromByte": from / 8, "ToByte": (to + 7) / 8,
			"Dynamic": structDynamic, "TempName": g.generateDecodeTempVarName(from), "GenerateStructPackagePrefix": g.generateStructPackagePrefix,
		})
		decodeStmts = append(decodeStmts, stmt)

	case *definition.Enum:
		panic("unreachable")

	case *definition.String:
		stmt := util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldString", nil, map[string]any{
			"FieldName": fieldNameStr, "FromByte": from / 8,
		})
		decodeStmts = append(decodeStmts, stmt)

	case *definition.Bytes:
		stmt := util.ExecuteTemplate(fieldDecoderTemplate, "decodeNormalFieldBytes", nil, map[string]any{
			"FieldName": fieldNameStr, "FromByte": from / 8,
			"GetMask": g.generateHex((1 << 7) - 1), "SetMask": g.generateHex(1 << 7), "Shift": 7,
			"TempName": g.generateDecodeTempVarName(from), "GenOptions": g.GenCtx.GenOptions,
		})
		decodeStmts = append(decodeStmts, stmt)

	case *definition.BasicType:
		generateBin := g.generateBin
		generateDec := g.generateDec
		generateHex := g.generateHex
		if ty.GetTypeBitSize() > 32 {
			generateBin = g.generateBinBigInt
			generateDec = g.generateDecBigInt
			generateHex = g.generateHexBigInt
			oldDataFunc := dataDataFunc
			dataDataFunc = func(i int64) string {
				expr := &definition.CastExpr{ToType: &definition.Uint64, Expr1: &definition.RawExpr{Expr: oldDataFunc(i)}}
				s, err := g.GenerateExpr(expr, "")
				if err != nil {
					panic(fmt.Errorf("internal error: %s", err))
				}
				return s
			}
		}
		shiftCalc := func(i int64) string { return generateDec(i * 8) }
		if gen.MatchOption(fieldOptions, "order", "big") {
			shiftCalc = func(i int64) string { return generateDec(max(0, fieldBitSize-8*(i+1))) }
		}

		fieldProcessor := func(expr string, i int64) string {
			operator := ""
			if i == 0 {
				operator = exprOpToString[definition.ExprOp_ASSIGN]
			} else {
				operator = exprOpToString[definition.ExprOp_BOR] + exprOpToString[definition.ExprOp_ASSIGN]
			}
			exprExp := &definition.BinopExpr{
				Op:    definition.ExprOp_SHL,
				Expr1: &definition.RawExpr{Expr: expr},
				Expr2: &definition.RawExpr{Expr: shiftCalc(i)},
			}
			exprExpStr, err := g.GenerateExpr(exprExp, "")
			if err != nil {
				panic(fmt.Errorf("internal error: %s", err))
			}
			return util.ExecuteTemplate(fieldDecoderTemplate, "decodeImpl", nil, map[string]any{
				"FieldName": fieldNameStr, "Operator": operator, "Expr": exprExpStr,
			})
		}

		decodeStmts = append(decodeStmts, g.generateDecodeImpl(from, to, fieldProcessor, dataDataFunc, generateBin, generateDec)...)

		if ty.GetTypeID().IsBool() {
			castExpr := &definition.CastExpr{ToType: &definition.Bool, Expr1: &definition.RawExpr{Expr: fieldNameStr}}
			castExprStr, err := g.GenerateExpr(castExpr, "")
			if err != nil {
				panic(fmt.Errorf("internal error: %s", err))
			}
			stmt := util.ExecuteTemplate(fieldDecoderTemplate, "decodeImpl", nil, map[string]any{
				"FieldName": fieldNameStr, "Operator": exprOpToString[definition.ExprOp_ASSIGN], "Expr": castExprStr,
			})
			decodeStmts = append(decodeStmts, stmt)
		}

		if ty.GetTypeID().IsInt() {
			originFromBitSize := fieldBitSize
			switch g.GenCtx.GenOptions.SignExtMethod {
			case gen.SignExtMethodDefault, gen.SignExtMethodShift, gen.SignExtMethodArith:
				if originFromBitSize == 32 {
					stmt := util.ExecuteTemplate(fieldDecoderTemplate, "signExtendLogic", nil, map[string]any{
						"FieldName": fieldNameStr, "SignShift": generateDec(originFromBitSize - 1),
					})
					decodeStmts = append(decodeStmts, stmt)
				} else {
					signMask := uint64(1) << (originFromBitSize - 1)
					stmt := util.ExecuteTemplate(fieldDecoderTemplate, "signExtendArith", nil, map[string]any{
						"FieldName": fieldNameStr, "SignMask": generateHex(signMask),
					})
					decodeStmts = append(decodeStmts, stmt)
				}
			default:
				panic("unreachable")
			}
		}

	case *definition.Array:
		panic("unreachable")

	default:
		return nil, fmt.Errorf("internal error: unknown field kind %T", ty)
	}

	return decodeStmts, nil
}

func (g ESModuleGenerator) generateDecodeImpl(from, to int64, fieldProcessor func(string, int64) string, dataData func(int64) string, generateBin func(any) string, generateDec func(any) string) []string {
	decodeStmts := []string{}
	for i := from; i < to; i += 8 {
		begin := i
		end := min(to, i+8)
		width := end - begin
		var expr definition.Expr
		sep := min(end, (begin+8)&(^7))
		if begin < sep {
			fieldMask := ((1 << (((sep - 1) & 7) + 1)) - 1) & (^((1 << (begin & 7)) - 1))
			shiftRight := begin % 8
			expr = &definition.BinopExpr{
				Op:    definition.ExprOp_SHR,
				Expr1: &definition.BinopExpr{Op: definition.ExprOp_BAND, Expr1: &definition.RawExpr{Expr: dataData(begin / 8)}, Expr2: &definition.RawExpr{Expr: generateBin(fieldMask)}},
				Expr2: &definition.RawExpr{Expr: generateDec(shiftRight)},
			}
		}
		if sep < end {
			fieldMask := ((1 << (((end - 1) & 7) + 1)) - 1) & (^((1 << (sep & 7)) - 1))
			shiftLeft := width - end%8
			expr = &definition.BinopExpr{
				Op:    definition.ExprOp_BOR,
				Expr1: expr,
				Expr2: &definition.BinopExpr{
					Op:    definition.ExprOp_SHL,
					Expr1: &definition.BinopExpr{Op: definition.ExprOp_BAND, Expr1: &definition.RawExpr{Expr: dataData(sep / 8)}, Expr2: &definition.RawExpr{Expr: generateBin(fieldMask)}},
					Expr2: &definition.RawExpr{Expr: generateDec(shiftLeft)},
				},
			}
		}
		exprStr, err := g.GenerateExpr(expr, "")
		if err != nil {
			panic(fmt.Errorf("internal error: %s", err))
		}
		decodeStmts = append(decodeStmts, fieldProcessor(exprStr, (i-from)/8))
	}
	return decodeStmts
}

// ==================== GenerateExpr ====================

func (g ESModuleGenerator) GenerateExpr(expr definition.Expr, valueStr string) (string, error) {
	generator := NewESModuleExprGenerator(g.GenerateType, valueStr)
	return g.AcceptExpr(expr, generator)
}

// ==================== Expr Generator ====================

type ESModuleExprGenerator struct {
	*gen.GenExprDispatcher
	GenType          func(definition.Type) (string, error)
	ValueStr         string
	LiteralGenerator gen.LiteralGeneratorImpl
}

func NewESModuleExprGenerator(genType func(definition.Type) (string, error), valueStr string) *ESModuleExprGenerator {
	g := &ESModuleExprGenerator{GenType: genType, ValueStr: valueStr}
	g.GenExprDispatcher = gen.NewGenExprDispatcher(g)
	return g
}

func (g ESModuleExprGenerator) GenerateExpr(expr definition.Expr) (string, error) {
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

func (g ESModuleExprGenerator) GenerateUnopExpr(expr *definition.UnopExpr) (string, error) {
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

func (g ESModuleExprGenerator) GenerateBinopExpr(expr *definition.BinopExpr) (string, error) {
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

func (g ESModuleExprGenerator) GenerateCastExpr(expr *definition.CastExpr) (string, error) {
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

func (g ESModuleExprGenerator) GenerateConstantExpr(expr *definition.ConstantExpr) (string, error) {
	generator := g.LiteralGenerator
	if generator == nil {
		generator = NewESModuleLiteralGenerator()
	}
	return g.AcceptLiteral(expr.ConstantValue, generator)
}

func (g ESModuleExprGenerator) GenerateTenaryExpr(expr *definition.TenaryExpr) (string, error) {
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

func (g ESModuleExprGenerator) GenerateValueExpr(expr *definition.ValueExpr) (string, error) {
	return g.ValueStr, nil
}

func (g ESModuleExprGenerator) GenerateRawExpr(expr *definition.RawExpr) (string, error) {
	return expr.Expr, nil
}

// ==================== Literal Generator ====================

type ESModuleLiteralGenerator struct {
	*gen.GenLiteralDispatcher
}

func NewESModuleLiteralGenerator() *ESModuleLiteralGenerator {
	g := &ESModuleLiteralGenerator{}
	g.GenLiteralDispatcher = gen.NewGenLiteralDispatcher(g)
	return g
}

func (g ESModuleLiteralGenerator) GenerateLiteral(literal definition.Literal) (string, error) {
	return g.AcceptLiteral(literal)
}

func (g ESModuleLiteralGenerator) GenerateBoolLiteral(literal *definition.BoolLiteral) (string, error) {
	return fmt.Sprintf("%t", literal.BoolValue), nil
}

func (g ESModuleLiteralGenerator) GenerateIntLiteral(literal *definition.IntLiteral) (string, error) {
	if literal.IntValue > 2147483647 || literal.IntValue < -2147483648 {
		return generateDecBigInt(literal.IntValue), nil
	}
	return generateDec(literal.IntValue), nil
}

func (g ESModuleLiteralGenerator) GenerateFloatLiteral(literal *definition.FloatLiteral) (string, error) {
	return fmt.Sprintf("%g", literal.FloatValue), nil
}

func (g ESModuleLiteralGenerator) GenerateStringLiteral(literal *definition.StringLiteral) (string, error) {
	return fmt.Sprintf("%q", literal.StringValue), nil
}
