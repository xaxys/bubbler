package kotlin

import (
	"fmt"
	"strings"

	"github.com/xaxys/bubbler/definition"
	"github.com/xaxys/bubbler/generator/gen"
	"github.com/xaxys/bubbler/util"
)

func (g *KotlinGenerator) generateDecoder(s *definition.Struct) (string, error) {
	var b strings.Builder
	b.WriteString("    fun decodeSize(data: ByteArray, start: Int = 0): Int {\n")
	if !s.StructDynamic {
		fmt.Fprintf(&b, "        return if (data.size - start < %d) -%d else %d\n", s.StructBitSize/8, s.StructBitSize/8, s.StructBitSize/8)
	} else {
		b.WriteString("        var offset = 0\n")
		if err := s.ForEachField(func(field definition.Field, _ int, startBits int64, _ bool) error {
			f, ok := field.(*definition.NormalField)
			if !ok || !f.FieldType.GetTypeDynamic() {
				return nil
			}
			return g.emitDecodeSizeField(&b, f, startBits, "        ")
		}); err != nil {
			return "", err
		}
		fixed := s.StructBitSize / 8
		fmt.Fprintf(&b, "        if (data.size - start < offset + %d) return -(offset + %d)\n", fixed, fixed)
		fmt.Fprintf(&b, "        return offset + %d\n", fixed)
	}
	b.WriteString("    }\n\n")
	b.WriteString("    fun decode(data: ByteArray, start: Int = 0): Int {\n")
	b.WriteString("        if (decodeSize(data, start) < 0) return -1\n")
	if s.StructDynamic {
		b.WriteString("        var offset = 0\n")
	}
	if err := s.ForEachFieldWithPos(func(field definition.Field, _ int, startBits int64, dynamic bool, pos string) error {
		stmts, err := g.decodeField(field, startBits, dynamic, "        ")
		if err != nil {
			return err
		}
		if stmts != "" {
			fmt.Fprintf(&b, "        // %s %s: %s\n", pos, field.GetFieldKind(), field)
			b.WriteString(stmts)
		}
		return nil
	}); err != nil {
		return "", err
	}
	if s.StructDynamic {
		fmt.Fprintf(&b, "        return offset + %d\n", s.StructBitSize/8)
	} else {
		fmt.Fprintf(&b, "        return %d\n", s.StructBitSize/8)
	}
	b.WriteString("    }\n")
	return b.String(), nil
}

func (g *KotlinGenerator) emitDecodeSizeField(b *strings.Builder, f *definition.NormalField, startBits int64, indent string) error {
	name := "this." + util.TocamelCase(f.FieldName)
	if a, ok := f.FieldType.(*definition.Array); ok {
		return g.emitArray(indent, a.Length, func(index, inner string) error {
			return g.emitDecodeSizeValue(b, a.ElementType, name+"["+index+"]", startBits, inner)
		}, b)
	}
	return g.emitDecodeSizeValue(b, f.FieldType, name, startBits, indent)
}

func (g *KotlinGenerator) emitDecodeSizeValue(b *strings.Builder, ty definition.Type, expr string, startBits int64, indent string) error {
	fromByte := startBits / 8
	pos := fmt.Sprintf("start + offset + %d", fromByte)
	switch t := ty.(type) {
	case *definition.String:
		fmt.Fprintf(b, "%srun {\n", indent)
		fmt.Fprintf(b, "%s    if (data.size <= %s) return -(%s + 1 - start)\n", indent, pos, pos)
		fmt.Fprintf(b, "%s    var length = 0\n", indent)
		fmt.Fprintf(b, "%s    while (data[%s + length].toInt() != 0) {\n", indent, pos)
		fmt.Fprintf(b, "%s        length++\n", indent)
		fmt.Fprintf(b, "%s        if (data.size <= %s + length) return -(%s + length + 1 - start)\n", indent, pos, pos)
		fmt.Fprintf(b, "%s    }\n", indent)
		fmt.Fprintf(b, "%s    offset += length + 1\n", indent)
		fmt.Fprintf(b, "%s}\n", indent)
	case *definition.Bytes:
		fmt.Fprintf(b, "%srun {\n", indent)
		fmt.Fprintf(b, "%s    if (data.size <= %s) return -(%s + 1 - start)\n", indent, pos, pos)
		fmt.Fprintf(b, "%s    var length = 0\n", indent)
		fmt.Fprintf(b, "%s    var shift = 0\n", indent)
		fmt.Fprintf(b, "%s    while (true) {\n", indent)
		fmt.Fprintf(b, "%s        if (data.size <= %s) return -(%s + 1 - start)\n", indent, pos, pos)
		fmt.Fprintf(b, "%s        val part = data[%s].toInt() and 0xFF\n", indent, pos)
		fmt.Fprintf(b, "%s        length = length or ((part and 0x7F) shl shift)\n", indent)
		fmt.Fprintf(b, "%s        offset++\n", indent)
		fmt.Fprintf(b, "%s        if ((part and 0x80) == 0) break\n", indent)
		fmt.Fprintf(b, "%s        shift += 7\n", indent)
		fmt.Fprintf(b, "%s        if (shift > 28) return -1\n", indent)
		fmt.Fprintf(b, "%s    }\n", indent)
		fmt.Fprintf(b, "%s    if (data.size - start < offset + %d + length) return -(offset + %d + length)\n", indent, fromByte, fromByte)
		fmt.Fprintf(b, "%s    offset += length\n", indent)
		fmt.Fprintf(b, "%s}\n", indent)
	case *definition.Struct:
		if t.StructDynamic {
			fmt.Fprintf(b, "%srun {\n", indent)
			fmt.Fprintf(b, "%s    val subSize = %s.decodeSize(data, %s)\n", indent, expr, pos)
			fmt.Fprintf(b, "%s    if (subSize < 0) return -(offset + %d) + subSize\n", indent, fromByte)
			fmt.Fprintf(b, "%s    offset += subSize\n", indent)
			fmt.Fprintf(b, "%s}\n", indent)
		}
	default:
		return fmt.Errorf("unsupported Kotlin decodeSize type %T", ty)
	}
	return nil
}

func (g *KotlinGenerator) decodeField(field definition.Field, startBits int64, dynamic bool, indent string) (string, error) {
	switch f := field.(type) {
	case *definition.VoidField, *definition.EmbeddedField:
		return "", nil
	case *definition.ConstantField:
		lit, err := g.literalForType(f.FieldConstant, f.FieldType)
		if err != nil {
			return "", err
		}
		expected := g.basicToULong(f.FieldType, lit)
		return fmt.Sprintf("%sif (bbReadFieldBits(data, %s, %dL, %t) != (%s and bbMask(%d))) return -1\n", indent, bitBase(startBits, dynamic), f.FieldBitSize, gen.MatchOption(f.FieldOptions, "order", "big"), expected, f.FieldBitSize), nil
	case *definition.NormalField:
		return g.decodeNormalField(f, startBits, dynamic, indent)
	default:
		return "", fmt.Errorf("unsupported Kotlin decode field %T", field)
	}
}

func (g *KotlinGenerator) decodeNormalField(f *definition.NormalField, startBits int64, dynamic bool, indent string) (string, error) {
	name := "this." + util.TocamelCase(f.FieldName)
	big := gen.MatchOption(f.FieldOptions, "order", "big")
	read := fmt.Sprintf("bbReadFieldBits(data, %s, %dL, %t)", bitBase(startBits, dynamic), f.FieldBitSize, big)
	var b strings.Builder
	switch ty := f.FieldType.(type) {
	case *definition.BasicType:
		fmt.Fprintf(&b, "%s%s = %s\n", indent, name, g.decodeBasic(ty, read, f.FieldBitSize))
	case *definition.Enum:
		fmt.Fprintf(&b, "%s%s = %s.fromNumber(%s.toLong())\n", indent, name, ty.EnumName, read)
	case *definition.Struct:
		fmt.Fprintf(&b, "%sif (%s.decode(data, %s) < 0) return -1\n", indent, name, byteBase(startBits, dynamic))
		if ty.StructDynamic {
			fmt.Fprintf(&b, "%soffset += %s.decodeSize(data, %s)\n", indent, name, byteBase(startBits, dynamic))
		}
	case *definition.String:
		g.emitDecodeString(&b, name, startBits, indent)
	case *definition.Bytes:
		g.emitDecodeBytes(&b, name, startBits, indent)
	case *definition.Array:
		elemBits := f.FieldBitSize
		if elemBits >= 0 {
			elemBits /= ty.Length
		}
		err := g.emitArray(indent, ty.Length, func(index, inner string) error {
			expr := name + "[" + index + "]"
			if ty.ElementType.GetTypeDynamic() {
				return g.emitDecodeDynamicValue(&b, ty.ElementType, expr, startBits, inner)
			}
			subBase := fmt.Sprintf("%s + %s.toLong() * %dL", bitBase(startBits, dynamic), index, elemBits)
			subRead := fmt.Sprintf("bbReadFieldBits(data, %s, %dL, %t)", subBase, elemBits, big)
			switch elem := ty.ElementType.(type) {
			case *definition.BasicType:
				fmt.Fprintf(&b, "%s%s = %s\n", inner, expr, g.decodeBasic(elem, subRead, elemBits))
			case *definition.Enum:
				fmt.Fprintf(&b, "%s%s = %s.fromNumber(%s.toLong())\n", inner, expr, elem.EnumName, subRead)
			case *definition.Struct:
				fmt.Fprintf(&b, "%sif (%s.decode(data, %s) < 0) return -1\n", inner, expr, arrayStructByteBase(startBits, dynamic, index, elemBits))
			default:
				return fmt.Errorf("unsupported fixed Kotlin decode array element %T", ty.ElementType)
			}
			return nil
		}, &b)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("unsupported Kotlin decode type %T", f.FieldType)
	}
	return b.String(), nil
}

func (g *KotlinGenerator) emitDecodeDynamicValue(b *strings.Builder, ty definition.Type, expr string, startBits int64, indent string) error {
	switch t := ty.(type) {
	case *definition.String:
		g.emitDecodeString(b, expr, startBits, indent)
	case *definition.Bytes:
		g.emitDecodeBytes(b, expr, startBits, indent)
	case *definition.Struct:
		fmt.Fprintf(b, "%sif (%s.decode(data, %s) < 0) return -1\n", indent, expr, byteBase(startBits, true))
		if t.StructDynamic {
			fmt.Fprintf(b, "%soffset += %s.decodeSize(data, %s)\n", indent, expr, byteBase(startBits, true))
		}
	default:
		return fmt.Errorf("unsupported dynamic Kotlin decode array element %T", ty)
	}
	return nil
}

func (g *KotlinGenerator) emitDecodeString(b *strings.Builder, target string, startBits int64, indent string) {
	fmt.Fprintf(b, "%srun {\n", indent)
	fmt.Fprintf(b, "%s    var length = 0\n", indent)
	fmt.Fprintf(b, "%s    while (data[%s + length].toInt() != 0) length++\n", indent, byteBase(startBits, true))
	fmt.Fprintf(b, "%s    %s = data.copyOfRange(%s, %s + length).decodeToString()\n", indent, target, byteBase(startBits, true), byteBase(startBits, true))
	fmt.Fprintf(b, "%s    offset += length + 1\n", indent)
	fmt.Fprintf(b, "%s}\n", indent)
}

func (g *KotlinGenerator) emitDecodeBytes(b *strings.Builder, target string, startBits int64, indent string) {
	fmt.Fprintf(b, "%srun {\n", indent)
	fmt.Fprintf(b, "%s    var length = 0\n", indent)
	fmt.Fprintf(b, "%s    var shift = 0\n", indent)
	fmt.Fprintf(b, "%s    while (true) {\n", indent)
	fmt.Fprintf(b, "%s        val part = data[%s].toInt() and 0xFF\n", indent, byteBase(startBits, true))
	fmt.Fprintf(b, "%s        length = length or ((part and 0x7F) shl shift)\n", indent)
	fmt.Fprintf(b, "%s        offset++\n", indent)
	fmt.Fprintf(b, "%s        if ((part and 0x80) == 0) break\n", indent)
	fmt.Fprintf(b, "%s        shift += 7\n", indent)
	fmt.Fprintf(b, "%s    }\n", indent)
	fmt.Fprintf(b, "%s    %s = data.copyOfRange(%s, %s + length)\n", indent, target, byteBase(startBits, true), byteBase(startBits, true))
	fmt.Fprintf(b, "%s    offset += length\n", indent)
	fmt.Fprintf(b, "%s}\n", indent)
}

func (g *KotlinGenerator) decodeBasic(ty *definition.BasicType, read string, bits int64) string {
	switch ty.GetTypeID() {
	case definition.TypeID_Bool:
		return read + " != 0uL"
	case definition.TypeID_Uint8:
		return read + ".toUByte()"
	case definition.TypeID_Uint16:
		return read + ".toUShort()"
	case definition.TypeID_Uint32:
		return read + ".toUInt()"
	case definition.TypeID_Uint64:
		return read
	case definition.TypeID_Int8:
		return g.signExtend(read, bits) + ".toByte()"
	case definition.TypeID_Int16:
		return g.signExtend(read, bits) + ".toShort()"
	case definition.TypeID_Int32:
		return g.signExtend(read, bits) + ".toInt()"
	case definition.TypeID_Int64:
		return g.signExtend(read, bits)
	case definition.TypeID_Float32:
		return "Float.fromBits(" + read + ".toInt())"
	case definition.TypeID_Float64:
		return "Double.fromBits(" + read + ".toLong())"
	default:
		panic("unreachable Kotlin basic decode")
	}
}

func (g *KotlinGenerator) signExtend(read string, bits int64) string {
	if g.GenCtx.GenOptions.SignExtMethod == gen.SignExtMethodArith && bits < 64 {
		mask := fmt.Sprintf("(1uL shl %d)", bits-1)
		return "((" + read + " xor " + mask + ") - " + mask + ").toLong()"
	}
	return fmt.Sprintf("bbSignExtend(%s, %d)", read, bits)
}
