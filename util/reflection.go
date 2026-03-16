package util

import (
	"fmt"
	"reflect"
	"strings"
)

func ToString(v any) string {
	return formatValue(reflect.ValueOf(v), 0)
}

func formatValue(val reflect.Value, depth int) string {
	ident := strings.Repeat("  ", depth)
	nextIndent := strings.Repeat("  ", depth+1)

	if !val.IsValid() {
		return "nil"
	}

	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return "nil"
		}
		return fmt.Sprintf("(ptr)->%s", formatValue(val.Elem(), depth))
	}

	if val.Kind() == reflect.Interface {
		if val.IsNil() {
			return "nil"
		}
		return fmt.Sprintf("(iface)->%s", formatValue(val.Elem(), depth))
	}

	switch val.Kind() {
	case reflect.Struct:
		var sb strings.Builder
		typ := val.Type()

		sb.WriteString(fmt.Sprintf("%s:\n", typ.Name()))
		for i := 0; i < val.NumField(); i++ {
			fieldTyp := typ.Field(i)
			fieldVal := val.Field(i)

			sb.WriteString(fmt.Sprintf("%s%s: ", nextIndent, fieldTyp.Name))

			if fieldVal.CanInterface() {
				sb.WriteString(formatValue(fieldVal, depth+1))
			}
			sb.WriteString("\n")
		}
		return sb.String()

	case reflect.Slice, reflect.Array:
		if val.Len() == 0 {
			return "[]"
		}
		var sb strings.Builder
		sb.WriteString("[\n")
		for i := 0; i < val.Len(); i++ {
			sb.WriteString(fmt.Sprintf("%s%s,\n", nextIndent, formatValue(val.Index(i), depth+1)))
		}
		sb.WriteString(fmt.Sprintf("%s]", ident))
		return sb.String()

	case reflect.String:
		return fmt.Sprintf(`"%v"`, val.Interface())

	default:
		if val.CanInterface() {
			if stringer, ok := val.Interface().(fmt.Stringer); ok {
				return stringer.String()
			}
		}

		return fmt.Sprintf("%v", val.Interface())
	}
}
