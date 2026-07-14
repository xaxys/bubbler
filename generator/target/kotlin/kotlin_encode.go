package kotlin

import (
	"fmt"
	"strings"

	"github.com/xaxys/bubbler/definition"
	"github.com/xaxys/bubbler/generator/gen"
	"github.com/xaxys/bubbler/util"
)

func (g *KotlinGenerator) generateEncoder(s *definition.Struct) (string, error) {
	var b strings.Builder
	b.WriteString("    fun encodeSize(): Int {\n")
	if !s.StructDynamic {
		fmt.Fprintf(&b, "        return %d\n", s.StructBitSize/8)
	} else {
		fmt.Fprintf(&b, "        var result = %d\n", s.StructBitSize/8)
		if err := s.ForEachField(func(field definition.Field, _ int, _ int64, _ bool) error {
			f, ok := field.(*definition.NormalField)
			if !ok || !f.FieldType.GetTypeDynamic() {
				return nil
			}
			return g.emitEncodeSizeField(&b, f, "        ")
		}); err != nil {
			return "", err
		}
		b.WriteString("        return result\n")
	}
	b.WriteString("    }\n\n")
	b.WriteString("    fun encode(): ByteArray {\n")
	b.WriteString("        val data = ByteArray(encodeSize())\n")
	b.WriteString("        encode(data, 0)\n")
	b.WriteString("        return data\n")
	b.WriteString("    }\n\n")
	b.WriteString("    fun encode(data: ByteArray, start: Int = 0): Int {\n")
	if s.StructDynamic {
		b.WriteString("        var offset = 0\n")
	}
	if err := s.ForEachFieldWithPos(func(field definition.Field, _ int, startBits int64, dynamic bool, pos string) error {
		stmts, err := g.encodeField(field, startBits, dynamic, "        ")
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

func (g *KotlinGenerator) emitEncodeSizeField(b *strings.Builder, f *definition.NormalField, indent string) error {
	name := "this." + util.TocamelCase(f.FieldName)
	if a, ok := f.FieldType.(*definition.Array); ok {
		return g.emitArray(indent, a.Length, func(index string, inner string) error {
			return g.emitEncodeSizeValue(b, a.ElementType, name+"["+index+"]", inner)
		}, b)
	}
	return g.emitEncodeSizeValue(b, f.FieldType, name, indent)
}

func (g *KotlinGenerator) emitEncodeSizeValue(b *strings.Builder, ty definition.Type, expr, indent string) error {
	switch t := ty.(type) {
	case *definition.String:
		fmt.Fprintf(b, "%sresult += %s.encodeToByteArray().size + 1\n", indent, expr)
	case *definition.Bytes:
		fmt.Fprintf(b, "%sresult += %s.size + bbVarUIntSize(%s.size)\n", indent, expr, expr)
	case *definition.Struct:
		if t.StructDynamic {
			fmt.Fprintf(b, "%sresult += %s.encodeSize()\n", indent, expr)
		}
	default:
		return fmt.Errorf("unsupported dynamic Kotlin size type %T", ty)
	}
	return nil
}

func (g *KotlinGenerator) encodeField(field definition.Field, startBits int64, dynamic bool, indent string) (string, error) {
	switch f := field.(type) {
	case *definition.VoidField, *definition.EmbeddedField:
		return "", nil
	case *definition.ConstantField:
		lit, err := g.literalForType(f.FieldConstant, f.FieldType)
		if err != nil {
			return "", err
		}
		value := g.basicToULong(f.FieldType, lit)
		return fmt.Sprintf("%sbbWriteFieldBits(data, %s, %dL, %s, %t)\n", indent, bitBase(startBits, dynamic), f.FieldBitSize, value, gen.MatchOption(f.FieldOptions, "order", "big")), nil
	case *definition.NormalField:
		return g.encodeNormalField(f, startBits, dynamic, indent)
	default:
		return "", fmt.Errorf("unsupported Kotlin encode field %T", field)
	}
}

func (g *KotlinGenerator) encodeNormalField(f *definition.NormalField, startBits int64, dynamic bool, indent string) (string, error) {
	name := "this." + util.TocamelCase(f.FieldName)
	big := gen.MatchOption(f.FieldOptions, "order", "big")
	var b strings.Builder
	switch ty := f.FieldType.(type) {
	case *definition.BasicType:
		fmt.Fprintf(&b, "%sbbWriteFieldBits(data, %s, %dL, %s, %t)\n", indent, bitBase(startBits, dynamic), f.FieldBitSize, g.basicToULong(ty, name), big)
	case *definition.Enum:
		fmt.Fprintf(&b, "%sbbWriteFieldBits(data, %s, %dL, %s!!.number.toULong(), %t)\n", indent, bitBase(startBits, dynamic), f.FieldBitSize, name, big)
	case *definition.Struct:
		fmt.Fprintf(&b, "%s%s.encode(data, %s)\n", indent, name, byteBase(startBits, dynamic))
		if ty.StructDynamic {
			fmt.Fprintf(&b, "%soffset += %s.encodeSize()\n", indent, name)
		}
	case *definition.String:
		fmt.Fprintf(&b, "%srun {\n", indent)
		fmt.Fprintf(&b, "%s    val bytes = %s.encodeToByteArray()\n", indent, name)
		fmt.Fprintf(&b, "%s    bytes.copyInto(data, %s)\n", indent, byteBase(startBits, true))
		fmt.Fprintf(&b, "%s    data[%s + bytes.size] = 0\n", indent, byteBase(startBits, true))
		fmt.Fprintf(&b, "%s    offset += bytes.size + 1\n", indent)
		fmt.Fprintf(&b, "%s}\n", indent)
	case *definition.Bytes:
		g.emitEncodeBytes(&b, name, startBits, indent)
	case *definition.Array:
		elemBits := f.FieldBitSize
		if elemBits >= 0 {
			elemBits /= ty.Length
		}
		err := g.emitArray(indent, ty.Length, func(index string, inner string) error {
			expr := name + "[" + index + "]"
			if ty.ElementType.GetTypeDynamic() {
				return g.emitEncodeDynamicValue(&b, ty.ElementType, expr, startBits, inner)
			}
			subStart := fmt.Sprintf("%s + %s.toLong() * %dL", bitBase(startBits, dynamic), index, elemBits)
			switch elem := ty.ElementType.(type) {
			case *definition.BasicType:
				fmt.Fprintf(&b, "%sbbWriteFieldBits(data, %s, %dL, %s, %t)\n", inner, subStart, elemBits, g.basicToULong(elem, expr), big)
			case *definition.Enum:
				fmt.Fprintf(&b, "%sbbWriteFieldBits(data, %s, %dL, %s!!.number.toULong(), %t)\n", inner, subStart, elemBits, expr, big)
			case *definition.Struct:
				fmt.Fprintf(&b, "%s%s.encode(data, %s)\n", inner, expr, arrayStructByteBase(startBits, dynamic, index, elemBits))
			default:
				return fmt.Errorf("unsupported fixed Kotlin array element %T", ty.ElementType)
			}
			return nil
		}, &b)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("unsupported Kotlin encode type %T", f.FieldType)
	}
	return b.String(), nil
}

func (g *KotlinGenerator) emitEncodeDynamicValue(b *strings.Builder, ty definition.Type, expr string, startBits int64, indent string) error {
	switch t := ty.(type) {
	case *definition.String:
		fmt.Fprintf(b, "%srun {\n", indent)
		fmt.Fprintf(b, "%s    val bytes = %s.encodeToByteArray()\n", indent, expr)
		fmt.Fprintf(b, "%s    bytes.copyInto(data, %s)\n", indent, byteBase(startBits, true))
		fmt.Fprintf(b, "%s    data[%s + bytes.size] = 0\n", indent, byteBase(startBits, true))
		fmt.Fprintf(b, "%s    offset += bytes.size + 1\n", indent)
		fmt.Fprintf(b, "%s}\n", indent)
	case *definition.Bytes:
		g.emitEncodeBytes(b, expr, startBits, indent)
	case *definition.Struct:
		fmt.Fprintf(b, "%s%s.encode(data, %s)\n", indent, expr, byteBase(startBits, true))
		if t.StructDynamic {
			fmt.Fprintf(b, "%soffset += %s.encodeSize()\n", indent, expr)
		}
	default:
		return fmt.Errorf("unsupported dynamic Kotlin array element %T", ty)
	}
	return nil
}

func (g *KotlinGenerator) emitEncodeBytes(b *strings.Builder, expr string, startBits int64, indent string) {
	fmt.Fprintf(b, "%srun {\n", indent)
	fmt.Fprintf(b, "%s    var length = %s.size\n", indent, expr)
	fmt.Fprintf(b, "%s    do {\n", indent)
	fmt.Fprintf(b, "%s        var part = length and 0x7F\n", indent)
	fmt.Fprintf(b, "%s        length = length ushr 7\n", indent)
	fmt.Fprintf(b, "%s        if (length != 0) part = part or 0x80\n", indent)
	fmt.Fprintf(b, "%s        data[%s] = part.toByte()\n", indent, byteBase(startBits, true))
	fmt.Fprintf(b, "%s        offset++\n", indent)
	fmt.Fprintf(b, "%s    } while (length != 0)\n", indent)
	fmt.Fprintf(b, "%s    %s.copyInto(data, %s)\n", indent, expr, byteBase(startBits, true))
	fmt.Fprintf(b, "%s    offset += %s.size\n", indent, expr)
	fmt.Fprintf(b, "%s}\n", indent)
}

func (g *KotlinGenerator) emitArray(indent string, length int64, body func(index, indent string) error, b *strings.Builder) error {
	useLoop := g.GenCtx.GenOptions.LoopUnroll >= 0 && length > int64(g.GenCtx.GenOptions.LoopUnroll)
	if useLoop {
		fmt.Fprintf(b, "%sfor (_i in 0 until %d) {\n", indent, length)
		if err := body("_i", indent+"    "); err != nil {
			return err
		}
		fmt.Fprintf(b, "%s}\n", indent)
		return nil
	}
	for i := int64(0); i < length; i++ {
		if err := body(fmt.Sprintf("%d", i), indent); err != nil {
			return err
		}
	}
	return nil
}

func (g *KotlinGenerator) basicToULong(ty *definition.BasicType, expr string) string {
	switch ty.GetTypeID() {
	case definition.TypeID_Bool:
		return "(if (" + expr + ") 1uL else 0uL)"
	case definition.TypeID_Float32:
		return "" + expr + ".toRawBits().toUInt().toULong()"
	case definition.TypeID_Float64:
		return "" + expr + ".toRawBits().toULong()"
	default:
		return "(" + expr + ").toULong()"
	}
}

func bitBase(startBits int64, dynamic bool) string {
	bytePos := startBits / 8
	bitPos := startBits % 8
	offset := "start"
	if dynamic {
		offset = "start + offset"
	}
	return fmt.Sprintf("((%s + %d).toLong() * 8L + %dL)", offset, bytePos, bitPos)
}

func byteBase(startBits int64, dynamic bool) string {
	base := fmt.Sprintf("start + %d", startBits/8)
	if dynamic {
		base = fmt.Sprintf("start + offset + %d", startBits/8)
	}
	return "(" + base + ")"
}

func arrayStructByteBase(startBits int64, dynamic bool, index string, elemBits int64) string {
	base := fmt.Sprintf("start + %d + %s * %d", startBits/8, index, elemBits/8)
	if dynamic {
		base = fmt.Sprintf("start + offset + %d + %s * %d", startBits/8, index, elemBits/8)
	}
	return "(" + base + ")"
}
