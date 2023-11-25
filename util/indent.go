package util

import (
	"fmt"
	"strings"
)

func IndentTab(v any, n int) string {
	tab := strings.Repeat("\t", n)
	s := fmt.Sprintf("%s%s", tab, v)
	s = strings.ReplaceAll(s, "\n", "\n"+tab)
	s = strings.TrimRight(s, "\t")
	return s
}

func IndentTab1(v any) string {
	return IndentTab(v, 1)
}

func IndentSpace(v any, n int) string {
	space := strings.Repeat(" ", n)
	s := fmt.Sprintf("%s%s", space, v)
	s = strings.ReplaceAll(s, "\n", "\n"+space)
	s = strings.TrimRight(s, " ")
	return s
}

func IndentSpace4(v any) string {
	return IndentSpace(v, 4)
}

func IndentSpace8(v any) string {
	return IndentSpace(v, 8)
}

func IndentSpaceNoFirst(v any, n int) string {
	space := strings.Repeat(" ", n)
	s := fmt.Sprintf("%s", v)
	s = strings.ReplaceAll(s, "\n", "\n"+space)
	s = strings.TrimRight(s, " ")
	return s
}

func IndentSpace4NoFirst(v any) string {
	return IndentSpaceNoFirst(v, 4)
}

func IndentSpace8NoFirst(v any) string {
	return IndentSpaceNoFirst(v, 8)
}
