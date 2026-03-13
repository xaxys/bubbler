package esm

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/xaxys/bubbler/definition"
	"github.com/xaxys/bubbler/generator/gen"
	"github.com/xaxys/bubbler/util"
)

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
 * Calculate encoded size of {{ $structName }} object.
 * @function encode_size
 * @description Returns an estimation of the encoded size in bytes.
 * @memberof {{ .StructDef.StructBelongs.Package }}.{{ $structName }}
 * @static
 * @param {Object} obj {{ $structName }} object.
 * @returns {Number} Encoded size of {{ $structName }} object.
 */
static encode_size(obj) {
{{- if .StructDef.StructDynamic }}
    let size = {{ calc .StructDef.StructBitSize "/" 8 }}
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
}

/**
 * Calculate encoded size of {{ $structName }} object.
 * @function encode_size
 * @memberof {{ .StructDef.StructBelongs.Package }}.{{ $structName }}
 * @instance
 * @returns {Number} Encoded size of {{ $structName }} object.
 */
encode_size() {
    return {{ $structName }}.encode_size(this);
}

/**
 * Encode {{ $structName }} object to buffer.
 * @function encode
 * @memberof {{ .StructDef.StructBelongs.Package }}.{{ $structName }}
 * @static
 * @param {Object} obj {{ $structName }} object.
 * @param {Array|Uint8Array} [buffer] The buffer to encode data.
 * @param {Number} [start] The start position to store the encoded data.
 * @returns {Array|Uint8Array|Number} Encoded data buffer when buffer is omitted, otherwise encoded size.
 */
static encode(obj, buffer, start) {
    if (obj === undefined) return buffer === undefined ? -1 : undefined;
    let data = buffer;
    if (data === undefined) data = new {{ if .GenOptions.CompatibleMode }}Array{{ else }}Uint8Array{{ end }}({{ if .Dynamic }}obj.encode_size(){{ else }}{{ calc .StructDef.StructBitSize "/" 8 }}{{ end }});
    if (start === undefined) start = 0;
    {{- if .Dynamic }}
    let offset = 0;
    {{- end }}
    {{- range $encodeStr := .EncodeStrs }}
    {{ $encodeStr }}
    {{- end }}
    return buffer === undefined ? data : {{ if .Dynamic }}offset + {{ end }}{{ calc .StructDef.StructBitSize "/" 8 }}
}

/**
 * Encode {{ $structName }} object to buffer.
 * @function encode
 * @memberof {{ .StructDef.StructBelongs.Package }}.{{ $structName }}
 * @instance
 * @param {Array|Uint8Array} [data] The buffer to encode data.
 * @param {Number} [start] The start position to store the encoded data.
 * @returns {Array|Uint8Array|Number} Encoded data buffer when data is omitted, otherwise encoded size.
 */
encode(data, start) {
    return {{ $structName }}.encode(this, data, start);
}
{{- end -}}
`

// ==================== GenerateEncoder ====================

func (g ESModuleGenerator) GenerateEncoder(structDef *definition.Struct) (string, error) {
	encodeStrs := []string{}
	if err := structDef.ForEachFieldWithPos(func(field definition.Field, fieldIndex int, startBits int64, dynamic bool, pos string) error {
		encodeStmts, err := g.generateEncodeField(field, startBits)
		if err != nil {
			return err
		}
		if len(encodeStmts) == 0 {
			return nil
		}
		filtered := []string{}
		for _, s := range encodeStmts {
			if s != "" {
				filtered = append(filtered, s)
			}
		}
		encodeFieldData := map[string]any{
			"Pos":         pos,
			"Field":       field,
			"EncodeStmts": filtered,
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
		"GenOptions": g.GenCtx.GenOptions,
		"Dynamic":    structDef.GetTypeDynamic(),
	}
	return util.ExecuteTemplate(encoderTemplate, "encoder", nil, fieldData), nil
}

func (g ESModuleGenerator) generateEncodeField(field definition.Field, startBits int64) ([]string, error) {
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
    const {{ .TempName }} = {{ .FieldName }};
{{- end -}}

{{- define "encodeNormalFieldTempVarDeclEnum" -}}
{{- $inPackage := .FieldDef.FieldBelongs.StructBelongs.LocalNames.Has .EnumDef.EnumName -}}
{{- if or .SingleFile $inPackage -}}
    const {{ .TempName }} = (typeof {{ .FieldName }} === "number") ? {{ .FieldName }} : {{ .EnumDef.EnumName }}[{{ .FieldName }}];
{{- else -}}
    const {{ .TempName }} = (typeof {{ .FieldName }} === "number") ? {{ .FieldName }} : ${{ .EnumDef.EnumBelongs.Package.PackageName }}.{{ .EnumDef.EnumName }}[{{ .FieldName }}];
{{- end -}}
{{- end -}}

{{- define "encodeNormalFieldTempVarDeclFloatCast" -}}
{{- if .IsFloat32 -}}
    const {{ .TempName }} = floatToUint32Bits({{ .FieldName }});
{{- else -}}
    const {{ .TempName }} = doubleToUint64Bits({{ .FieldName }});
{{- end -}}
{{- end -}}

{{- define "encodeNormalFieldTempVarDeclOnly" -}}
    let {{ .TempName }};
{{- end -}}

{{- define "encodeNormalFieldTempVarAssignEnum" -}}
{{- $inPackage := .FieldDef.FieldBelongs.StructBelongs.LocalNames.Has .EnumDef.EnumName -}}
{{- if or .SingleFile $inPackage -}}
    {{ .TempName }} = (typeof {{ .FieldName }} === "number") ? {{ .FieldName }} : {{ .EnumDef.EnumName }}[{{ .FieldName }}];
{{- else -}}
    {{ .TempName }} = (typeof {{ .FieldName }} === "number") ? {{ .FieldName }} : ${{ .EnumDef.EnumBelongs.Package.PackageName }}.{{ .EnumDef.EnumName }}[{{ .FieldName }}];
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
        const {{ .TempName }} = stringToUTF8Bytes({{ .FieldName }}, data, offset + start + {{ .FromByte }});
        data[offset + start + {{ .FromByte }} + {{ .TempName }}] = 0;
        offset += {{ .TempName }} + 1;
    })();
{{- end -}}

{{- define "encodeNormalFieldBytes" -}}
    (function() {
        let {{ .TempName }} = {{ .FieldName }}.length;
        do { data[offset + start + {{ .FromByte }}] = {{ .TempName }} & {{ .GetMask }} | {{ .SetMask }}; offset++; {{ .TempName }} >>= {{ .Shift }}; } while ({{ .TempName }} > 0);
        data[offset - 1 + start + {{ .FromByte }}] &= ~{{ .SetMask }};
        for (let i = 0; i < {{ .FieldName }}.length; i++) data[offset + start + {{ .FromByte }} + i] = {{ .FieldName }}[i];
        offset += {{ .FieldName }}.length;
    })();
{{- end -}}

{{- define "encodeImpl" -}}
    data[{{ if .Dynamic }}offset + {{ end }}start + {{ .BytePos }}] {{ .Operator }} {{ .FieldData }};
{{- end -}}
`

func (g ESModuleGenerator) generateEncodeTempVarName(startBits int64) string {
	return util.ExecuteTemplate(fieldEncoderTemplate, "encodeTempVarName", nil, map[string]any{"StartBits": startBits})
}

func (g ESModuleGenerator) generateEncodeStructFieldName(name string) string {
	return util.ExecuteTemplate(fieldEncoderTemplate, "encodeStructFieldName", nil, map[string]any{"FieldName": name})
}

func (g ESModuleGenerator) generateEncodeConstantField(field *definition.ConstantField, startBits int64) ([]string, error) {
	structDynamic := field.FieldBelongs.GetTypeDynamic()
	var byteOrder binary.ByteOrder = binary.LittleEndian
	if gen.MatchOption(field.FieldOptions, "order", "big") {
		byteOrder = binary.BigEndian
	}
	buffer := &bytes.Buffer{}
	value := field.FieldConstant.GetLiteralValueIn(field.FieldType.TypeTypeID)
	if err := binary.Write(buffer, byteOrder, value); err != nil {
		return nil, fmt.Errorf("internal error: %s", err)
	}
	data := buffer.Bytes()
	fieldData := func(i int64) string { return g.generateHex(data[i]) }
	return g.generateEncodeImpl(startBits, startBits+field.GetFieldBitSize(), fieldData, g.generateBin, g.generateDec, false, structDynamic), nil
}

func (g ESModuleGenerator) generateEncodeVoidField(field *definition.VoidField, startBits int64) ([]string, error) {
	return []string{""}, nil
}

func (g ESModuleGenerator) generateEncodeEmbeddedField(field *definition.EmbeddedField, startBits int64) ([]string, error) {
	return nil, nil
}

func (g ESModuleGenerator) generateEncodeNormalField(field *definition.NormalField, startBits int64) ([]string, error) {
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
			tempName := g.generateEncodeTempVarName(startBits)
			stmt := util.ExecuteTemplate(fieldEncoderTemplate, "encodeNormalFieldTempVarDeclFloatCast", nil, map[string]any{
				"TempName": tempName, "FieldName": name, "IsFloat32": ty.GetTypeID() == definition.TypeID_Float32,
			})
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
		tempTy := &definition.Uint32
		declStr := util.ExecuteTemplate(fieldEncoderTemplate, "encodeNormalFieldTempVarDeclEnum", nil, map[string]any{
			"EnumDef": ty, "TempName": tempName, "FieldName": g.generateEncodeStructFieldName(field.FieldName), "FieldDef": field,
			"SingleFile": g.GenCtx.GenOptions.SingleFile,
		})
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

		var nameIndex func(int64) string
		switch ty.ElementType.(type) {
		case *definition.Struct, *definition.String, *definition.Bytes:
			nameIndex = func(index int64) string { return fmt.Sprintf("%s[%d]", name, index) }
		case *definition.BasicType:
			nameIndex = func(index int64) string { return fmt.Sprintf("%s[%d]", name, index) }
			if ty.ElementType.GetTypeID().IsFloat() {
				tempName := g.generateEncodeTempVarName(startBits)
				encodeStmts = append(encodeStmts, util.ExecuteTemplate(fieldEncoderTemplate, "encodeNormalFieldTempVarDeclOnly", nil, map[string]any{"TempName": tempName}))
				nameIndex = func(_ int64) string { return tempName }
				switch ty.ElementType.GetTypeID() {
				case definition.TypeID_Float32:
					elemTy = &definition.Uint32
				case definition.TypeID_Float64:
					elemTy = &definition.Uint64
				}
			}
		case *definition.Enum:
			tempName := g.generateEncodeTempVarName(startBits)
			encodeStmts = append(encodeStmts, util.ExecuteTemplate(fieldEncoderTemplate, "encodeNormalFieldTempVarDeclOnly", nil, map[string]any{"TempName": tempName}))
			nameIndex = func(_ int64) string { return tempName }
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
					encodeStmts = append(encodeStmts, util.ExecuteTemplate(fieldEncoderTemplate, "encodeNormalFieldTempVarAssignFloatCast", nil, map[string]any{
						"TempName": subName, "FieldName": fmt.Sprintf("%s[%d]", name, i), "IsFloat32": ty.ElementType.GetTypeID() == definition.TypeID_Float32,
					}))
				}
			case *definition.Enum:
				encodeStmts = append(encodeStmts, util.ExecuteTemplate(fieldEncoderTemplate, "encodeNormalFieldTempVarAssignEnum", nil, map[string]any{
					"EnumDef": ty.ElementType, "TempName": subName, "FieldName": fmt.Sprintf("%s[%d]", name, i), "FieldDef": field,
				}))
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

func (g ESModuleGenerator) generateEncodeNormalFieldImpl(fieldNameStr string, fieldType definition.Type, fieldOptions *util.OrderedMap[string, *definition.Option], structDynamic bool, from, to int64) ([]string, error) {
	encodeStmts := []string{}
	fieldBitSize := to - from

	switch ty := fieldType.(type) {
	case *definition.Struct:
		stmt := util.ExecuteTemplate(fieldEncoderTemplate, "encodeNormalFieldStruct", nil, map[string]any{
			"FieldStruct": ty, "FieldName": fieldNameStr, "FromByte": from / 8, "Dynamic": structDynamic, "GenerateStructPackagePrefix": g.generateStructPackagePrefix,
		})
		encodeStmts = append(encodeStmts, stmt)

	case *definition.Enum:
		panic("unreachable")

	case *definition.String:
		stmt := util.ExecuteTemplate(fieldEncoderTemplate, "encodeNormalFieldString", nil, map[string]any{
			"FieldName": fieldNameStr, "FromByte": from / 8, "TempName": g.generateEncodeTempVarName(from),
		})
		encodeStmts = append(encodeStmts, stmt)

	case *definition.Bytes:
		stmt := util.ExecuteTemplate(fieldEncoderTemplate, "encodeNormalFieldBytes", nil, map[string]any{
			"FieldName": fieldNameStr, "FromByte": from / 8,
			"GetMask": g.generateHex((1 << 7) - 1), "SetMask": g.generateHex(1 << 7), "Shift": 7,
			"TempName": g.generateEncodeTempVarName(from),
		})
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
		fieldData := func(i int64) string {
			expr := &definition.BinopExpr{
				Op:    definition.ExprOp_SHR,
				Expr1: &definition.RawExpr{Expr: fieldNameStr},
				Expr2: &definition.RawExpr{Expr: generateDec(8 * i)},
			}
			s, err := g.GenerateExpr(expr, "")
			if err != nil {
				panic(fmt.Errorf("internal error: %s", err))
			}
			return s
		}
		if gen.MatchOption(fieldOptions, "order", "big") {
			fieldData = func(i int64) string {
				expr := &definition.BinopExpr{
					Op:    definition.ExprOp_SHR,
					Expr1: &definition.RawExpr{Expr: fieldNameStr},
					Expr2: &definition.RawExpr{Expr: generateDec(max(0, fieldBitSize-8*(i+1)))},
				}
				s, err := g.GenerateExpr(expr, "")
				if err != nil {
					panic(fmt.Errorf("internal error: %s", err))
				}
				return s
			}
		}
		encodeStmts = append(encodeStmts, g.generateEncodeImpl(from, to, fieldData, generateBin, generateDec, cast, structDynamic)...)

	case *definition.Array:
		panic("unreachable")

	default:
		return nil, fmt.Errorf("internal error: unknown field kind %T", ty)
	}

	return encodeStmts, nil
}

func (g ESModuleGenerator) generateEncodeImpl(from, to int64, fieldData func(int64) string, generateBin func(any) string, generateDec func(any) string, cast bool, structDynamic bool) []string {
	encodeStmts := []string{}
	for i := from; i < to; i = (i + 8) & (^7) {
		nextI := min(to, (i+8)&(^7))
		dataMask := ((1 << (((nextI - 1) & 7) + 1)) - 1) & (^((1 << (i & 7)) - 1))
		operator := ""
		if i%8 == 0 {
			operator = exprOpToString[definition.ExprOp_ASSIGN]
		} else {
			operator = exprOpToString[definition.ExprOp_BOR] + exprOpToString[definition.ExprOp_ASSIGN]
		}
		begin := i - from
		end := nextI - from
		var expr definition.Expr
		j := begin
		if j < end {
			nextJ := min(end, (j+8)&(^7))
			fieldMask := ((1 << (((nextJ - 1) & 7) + 1)) - 1) & (^((1 << (j & 7)) - 1))
			shiftRight := j % 8
			expr = &definition.BinopExpr{
				Op:    definition.ExprOp_SHR,
				Expr1: &definition.BinopExpr{Op: definition.ExprOp_BAND, Expr1: &definition.RawExpr{Expr: fieldData(j / 8)}, Expr2: &definition.RawExpr{Expr: generateBin(fieldMask)}},
				Expr2: &definition.RawExpr{Expr: generateDec(shiftRight)},
			}
			j = nextJ
		}
		if j < end {
			nextJ := min(end, (j+8)&(^7))
			fieldMask := ((1 << (((nextJ - 1) & 7) + 1)) - 1) & (^((1 << (j & 7)) - 1))
			shiftLeft := 8 - nextJ%8
			expr = &definition.BinopExpr{
				Op:    definition.ExprOp_BOR,
				Expr1: expr,
				Expr2: &definition.BinopExpr{
					Op:    definition.ExprOp_SHL,
					Expr1: &definition.BinopExpr{Op: definition.ExprOp_BAND, Expr1: &definition.RawExpr{Expr: fieldData(j / 8)}, Expr2: &definition.RawExpr{Expr: generateBin(fieldMask)}},
					Expr2: &definition.RawExpr{Expr: generateDec(shiftLeft)},
				},
			}
			j = nextJ
		}
		shiftLeft := i % 8
		expr = &definition.BinopExpr{
			Op:    definition.ExprOp_BAND,
			Expr1: &definition.BinopExpr{Op: definition.ExprOp_SHL, Expr1: expr, Expr2: &definition.RawExpr{Expr: generateDec(shiftLeft)}},
			Expr2: &definition.RawExpr{Expr: generateBin(dataMask)},
		}
		if cast {
			expr = &definition.CastExpr{ToType: &definition.BasicType{TypeTypeID: definition.TypeID_Uint8}, Expr1: expr}
		}
		exprStr, err := g.GenerateExpr(expr, "")
		if err != nil {
			panic(fmt.Errorf("internal error: %s", err))
		}
		encodeStmt := util.ExecuteTemplate(fieldEncoderTemplate, "encodeImpl", nil, map[string]any{
			"BytePos": i / 8, "Operator": operator, "FieldData": exprStr, "Dynamic": structDynamic,
		})
		encodeStmts = append(encodeStmts, encodeStmt)
	}
	return encodeStmts
}
