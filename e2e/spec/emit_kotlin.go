package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func EmitKotlin(spec *Spec) string {
	var b strings.Builder
	b.WriteString("@file:OptIn(ExperimentalUnsignedTypes::class)\n\n")
	b.WriteString("import com.example.testpkg.*\nimport com.example.bitwid.*\nimport kotlin.math.abs\n\n")
	b.WriteString(`private var current = ""

private fun check(condition: Boolean, message: String) {
    if (!condition) error("$current: $message")
}

private fun checkNear(actual: Double, expected: Double, tolerance: Double, message: String) {
    if (actual.isNaN() && expected.isNaN()) return
    check(abs(actual - expected) <= tolerance, "$message: got=$actual expected=$expected")
}

`)
	for _, sc := range spec.Scenarios {
		if sc.StructName == "GetterSetter" {
			b.WriteString(emitKotlinGetterSetter(sc))
		} else {
			b.WriteString(emitKotlinScenario(sc))
		}
		b.WriteByte('\n')
	}
	b.WriteString("fun main() {\n")
	for _, sc := range spec.Scenarios {
		fmt.Fprintf(&b, "    test%s()\n", sc.StructName)
	}
	b.WriteString("    println(\"Kotlin e2e OK\")\n}\n")
	return b.String()
}

func emitKotlinScenario(sc Scenario) string {
	var b strings.Builder
	fmt.Fprintf(&b, "private fun test%s() {\n", sc.StructName)
	fmt.Fprintf(&b, "    current = %s\n", strconv.Quote(sc.StructName))
	for _, c := range sc.Cases {
		fmt.Fprintf(&b, "    run { // %s\n", c.Name)
		fmt.Fprintf(&b, "        val source = %s()\n", sc.StructName)
		for _, f := range c.Setup {
			b.WriteString(emitKotlinSetNamed("source", f.Name, f.V, "        "))
		}
		b.WriteString("        val encoded = source.encode()\n")
		b.WriteString("        check(encoded.size == source.encodeSize(), \"encodeSize\")\n")
		if len(c.Wire) > 0 {
			fmt.Fprintf(&b, "        check(encoded.contentEquals(%s), \"golden wire\")\n", kotlinBytes(c.Wire))
		}
		fmt.Fprintf(&b, "        val decoded = %s()\n", sc.StructName)
		b.WriteString("        check(decoded.decode(encoded) == encoded.size, \"decode return\")\n")
		for _, a := range c.resolveAssert() {
			b.WriteString(emitKotlinAssertNamed("decoded", a.Name, a.V, a.Tol, "        "))
		}
		for _, failure := range c.Errors {
			b.WriteString("        run {\n")
			b.WriteString("            val bad = encoded.copyOf()\n")
			for _, patch := range failure.Patches {
				fmt.Fprintf(&b, "            bad[%d] = 0x%02X.toByte()\n", patch.Offset, patch.Value)
			}
			fmt.Fprintf(&b, "            val probe = %s()\n", sc.StructName)
			if failure.ExpectedRet == -1 {
				fmt.Fprintf(&b, "            check(probe.decode(bad) < 0, %s)\n", strconv.Quote(failure.Name))
			} else {
				fmt.Fprintf(&b, "            check(probe.decode(bad) == %d, %s)\n", failure.ExpectedRet, strconv.Quote(failure.Name))
			}
			b.WriteString("        }\n")
		}
		b.WriteString("    }\n")
	}
	if sc.StructName == "SimpleScalars" && len(sc.Cases) > 0 {
		b.WriteString("    run { // unknown enum values decode to null\n")
		fmt.Fprintf(&b, "        val source = %s()\n", sc.StructName)
		for _, f := range sc.Cases[0].Setup {
			b.WriteString(emitKotlinSetNamed("source", f.Name, f.V, "        "))
		}
		b.WriteString("        val encoded = source.encode()\n")
		b.WriteString("        encoded[encoded.lastIndex] = 0x7F.toByte()\n")
		fmt.Fprintf(&b, "        val decoded = %s()\n", sc.StructName)
		b.WriteString("        check(decoded.decode(encoded) == encoded.size, \"unknown enum decode succeeds\")\n")
		b.WriteString("        check(decoded.color == null, \"unknown enum becomes null\")\n")
		b.WriteString("    }\n")
	}
	b.WriteString(emitKotlinDecodeSize(sc))
	b.WriteString("}\n")
	return b.String()
}

func emitKotlinDecodeSize(sc Scenario) string {
	if len(sc.DecodeSizeChecks) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString("    run { // decodeSize\n")
	sourceCases := map[string]bool{}
	for _, checkDef := range sc.DecodeSizeChecks {
		if checkDef.SourceCase != "" {
			sourceCases[checkDef.SourceCase] = true
		}
	}
	for name := range sourceCases {
		for _, c := range sc.Cases {
			if c.Name != name {
				continue
			}
			id := kotlinIdentifier("full_" + name)
			fmt.Fprintf(&b, "        val source_%s = %s()\n", id, sc.StructName)
			for _, f := range c.Setup {
				b.WriteString(emitKotlinSetNamed("source_"+id, f.Name, f.V, "        "))
			}
			fmt.Fprintf(&b, "        val %s = source_%s.encode()\n", id, id)
		}
	}
	fmt.Fprintf(&b, "        val probe = %s()\n", sc.StructName)
	for _, d := range sc.DecodeSizeChecks {
		if d.SourceCase != "" {
			id := kotlinIdentifier("full_" + d.SourceCase)
			if d.Truncate == 0 {
				fmt.Fprintf(&b, "        check(probe.decodeSize(%s) == %s.size, %s)\n", id, id, strconv.Quote(d.Name))
			} else {
				fmt.Fprintf(&b, "        check(probe.decodeSize(%s.copyOf(%s.size - %d)) == -%s.size, %s)\n", id, id, d.Truncate, id, strconv.Quote(d.Name))
			}
		} else {
			fmt.Fprintf(&b, "        check(probe.decodeSize(%s) == %d, %s)\n", kotlinBytes(d.Bytes), d.Expected, strconv.Quote(d.Name))
		}
	}
	b.WriteString("    }\n")
	return b.String()
}

func emitKotlinGetterSetter(sc Scenario) string {
	var b strings.Builder
	fmt.Fprintf(&b, "private fun test%s() {\n", sc.StructName)
	fmt.Fprintf(&b, "    current = %s\n", strconv.Quote(sc.StructName))
	for _, c := range sc.Cases {
		b.WriteString("    run {\n")
		fmt.Fprintf(&b, "        val source = %s()\n", sc.StructName)
		for _, a := range c.Accessors {
			fmt.Fprintf(&b, "        source.%s = Double.fromBits(%s.toLong())\n", camelCase(a.SetterName), kotlinULong(a.SetArg.U))
		}
		roundTrip := false
		for _, a := range c.Accessors {
			if !a.NoRoundTrip {
				roundTrip = true
			}
		}
		target := "source"
		if roundTrip {
			b.WriteString("        val encoded = source.encode()\n")
			fmt.Fprintf(&b, "        val decoded = %s()\n", sc.StructName)
			b.WriteString("        check(decoded.decode(encoded) == encoded.size, \"decode return\")\n")
			target = "decoded"
		}
		for _, a := range c.Accessors {
			tol := a.Tol
			if tol <= 0 {
				tol = 1e-9
			}
			fmt.Fprintf(&b, "        checkNear(%s.%s, Double.fromBits(%s.toLong()), %s, %s)\n", target, camelCase(a.GetterName), kotlinULong(a.Expected.U), strconv.FormatFloat(tol, 'g', -1, 64), strconv.Quote(a.GetterName))
		}
		b.WriteString("    }\n")
	}
	b.WriteString("}\n")
	return b.String()
}

func emitKotlinSetNamed(target, name string, v Val, indent string) string {
	property := target + "." + camelCase(name)
	if v.Kind == VKStruct {
		var b strings.Builder
		tmp := kotlinIdentifier(target + "_" + name)
		fmt.Fprintf(&b, "%sval %s = %s()\n", indent, tmp, v.Struct.Type)
		for _, f := range v.Struct.Fields {
			b.WriteString(emitKotlinSetNamed(tmp, f.Name, f.V, indent))
		}
		fmt.Fprintf(&b, "%s%s = %s\n", indent, property, tmp)
		return b.String()
	}
	if v.Kind == VKArray {
		var b strings.Builder
		for i, item := range v.Array.Items {
			b.WriteString(emitKotlinSetAt(property, i, item, indent))
		}
		return b.String()
	}
	return fmt.Sprintf("%s%s = %s\n", indent, property, kotlinValue(v))
}

func emitKotlinSetAt(target string, index int, v Val, indent string) string {
	if v.Kind == VKStruct {
		var b strings.Builder
		tmp := kotlinIdentifier(fmt.Sprintf("%s_%d", target, index))
		fmt.Fprintf(&b, "%sval %s = %s()\n", indent, tmp, v.Struct.Type)
		for _, f := range v.Struct.Fields {
			b.WriteString(emitKotlinSetNamed(tmp, f.Name, f.V, indent))
		}
		fmt.Fprintf(&b, "%s%s[%d] = %s\n", indent, target, index, tmp)
		return b.String()
	}
	return fmt.Sprintf("%s%s[%d] = %s\n", indent, target, index, kotlinValue(v))
}

func emitKotlinAssertNamed(target, name string, v Val, tolerance float64, indent string) string {
	return emitKotlinAssertValue(target+"."+camelCase(name), name, v, tolerance, indent)
}

func emitKotlinAssertValue(actual, label string, v Val, tolerance float64, indent string) string {
	switch v.Kind {
	case VKF32:
		if tolerance > 0 {
			return fmt.Sprintf("%scheckNear(%s.toDouble(), Float.fromBits(0x%08X.toInt()).toDouble(), %s, %s)\n", indent, actual, uint32(v.U), strconv.FormatFloat(tolerance, 'g', -1, 64), strconv.Quote(label))
		}
		return fmt.Sprintf("%scheck(%s.toRawBits().toUInt() == 0x%08Xu, %s)\n", indent, actual, uint32(v.U), strconv.Quote(label+" bits"))
	case VKF64:
		if tolerance > 0 {
			return fmt.Sprintf("%scheckNear(%s, Double.fromBits(%s.toLong()), %s, %s)\n", indent, actual, kotlinULong(v.U), strconv.FormatFloat(tolerance, 'g', -1, 64), strconv.Quote(label))
		}
		return fmt.Sprintf("%scheck(%s.toRawBits().toULong() == %s, %s)\n", indent, actual, kotlinULong(v.U), strconv.Quote(label+" bits"))
	case VKBytes:
		return fmt.Sprintf("%scheck(%s.contentEquals(%s), %s)\n", indent, actual, kotlinBytes(v.Bytes), strconv.Quote(label))
	case VKStruct:
		var b strings.Builder
		for _, f := range v.Struct.Fields {
			b.WriteString(emitKotlinAssertNamed(actual, f.Name, f.V, tolerance, indent))
		}
		return b.String()
	case VKArray:
		var b strings.Builder
		for i, item := range v.Array.Items {
			b.WriteString(emitKotlinAssertValue(fmt.Sprintf("%s[%d]", actual, i), fmt.Sprintf("%s[%d]", label, i), item, tolerance, indent))
		}
		return b.String()
	default:
		return fmt.Sprintf("%scheck(%s == %s, %s)\n", indent, actual, kotlinValue(v), strconv.Quote(label))
	}
}

func kotlinValue(v Val) string {
	switch v.Kind {
	case VKBool:
		if v.U != 0 {
			return "true"
		}
		return "false"
	case VKU8:
		return fmt.Sprintf("%du.toUByte()", uint8(v.U))
	case VKU16:
		return fmt.Sprintf("%du.toUShort()", uint16(v.U))
	case VKU32:
		return fmt.Sprintf("%du", uint32(v.U))
	case VKU64:
		return kotlinULong(v.U)
	case VKI8:
		return fmt.Sprintf("(%d).toByte()", int8(v.I))
	case VKI16:
		return fmt.Sprintf("(%d).toShort()", int16(v.I))
	case VKI32:
		return fmt.Sprintf("%d", int32(v.I))
	case VKI64:
		if v.I == math.MinInt64 {
			return "Long.MIN_VALUE"
		}
		return fmt.Sprintf("%dL", v.I)
	case VKF32:
		return fmt.Sprintf("Float.fromBits(0x%08X.toInt())", uint32(v.U))
	case VKF64:
		return "Double.fromBits(" + kotlinULong(v.U) + ".toLong())"
	case VKString:
		return strconv.Quote(v.Str)
	case VKBytes:
		return kotlinBytes(v.Bytes)
	case VKEnum:
		return v.EnumType + "." + v.EnumName
	}
	return "error(\"unsupported test value\")"
}

func kotlinULong(v uint64) string { return fmt.Sprintf("0x%016XuL", v) }

func kotlinBytes(v []byte) string {
	if len(v) == 0 {
		return "byteArrayOf()"
	}
	parts := make([]string, len(v))
	for i, item := range v {
		parts[i] = fmt.Sprintf("0x%02X.toByte()", item)
	}
	return "byteArrayOf(" + strings.Join(parts, ", ") + ")"
}

func kotlinIdentifier(value string) string {
	var b strings.Builder
	b.WriteByte('_')
	for _, r := range value {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
			b.WriteRune(r)
		} else {
			b.WriteByte('_')
		}
	}
	return b.String()
}
