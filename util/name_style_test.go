package util

import (
	"math/rand"
	"testing"
)

func TestIsCapitalized(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"HelloWorld", true},
		{"helloWorld", false},
		{"_helloWorld", false},
		{"_HelloWorld", true},
	}

	for _, test := range tests {
		result := IsCapitalized(test.input)
		if result != test.expected {
			t.Errorf("IsCapitalized(%s) = %t, expected %t", test.input, result, test.expected)
		}
	}
}

func TestIsUncapitalized(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"helloWorld", true},
		{"HelloWorld", false},
		{"_helloWorld", true},
		{"_HelloWorld", false},
	}

	for _, test := range tests {
		result := IsUncapitalized(test.input)
		if result != test.expected {
			t.Errorf("IsUncapitalized(%s) = %t, expected %t", test.input, result, test.expected)
		}
	}
}

func TestToPascalCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"PropellerAData", "PropellerAData"},
		{"propellerAData", "PropellerAData"},
		{"propeller_a_data", "PropellerAData"},
		{"PROPELLER_A_DATA", "PropellerAData"},
		{"hello_world", "HelloWorld"},
		{"foo_bar_baz", "FooBarBaz"},
		{"abc_def_ghi", "AbcDefGhi"},
		{"____He___ll__o__", "____HeLlO__"},
		{"", ""},
		{"helloWorld", "HelloWorld"},
		{"fooBarBaz", "FooBarBaz"},
		{"abcDefGhi", "AbcDefGhi"},
		{"HelloWorld", "HelloWorld"},
		{"FooBarBaz", "FooBarBaz"},
		{"AbcDefGhi", "AbcDefGhi"},
		{"_hello_world", "_HelloWorld"},
		{"_foo_bar_baz", "_FooBarBaz"},
		{"_abc_def_ghi", "_AbcDefGhi"},
		{"CAPITALS_WITH_UNDERSCORES", "CapitalsWithUnderscores"},
		{"_CAPITALS_WITH_UNDERSCORES", "_CapitalsWithUnderscores"},
		{"hello_world_", "HelloWorld_"},
		{"hello__world", "HelloWorld"},
		{"_hello__world", "_HelloWorld"},
		{"__hello__world", "__HelloWorld"},
		{"hello_world123", "HelloWorld123"},
		{"hello_world_123", "HelloWorld123"},
	}

	for _, test := range tests {
		result := ToPascalCase(test.input)
		if result != test.expected {
			t.Errorf("ToPascalCase(%s) = %s, expected %s", test.input, result, test.expected)
		}
	}
}

func TestTocamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"PropellerAData", "propellerAData"},
		{"propellerAData", "propellerAData"},
		{"propeller_a_data", "propellerAData"},
		{"PROPELLER_A_DATA", "propellerAData"},
		{"hello_world", "helloWorld"},
		{"foo_bar_baz", "fooBarBaz"},
		{"abc_def_ghi", "abcDefGhi"},
		{"____He___ll__o__", "____heLlO__"},
		{"", ""},
		{"helloWorld", "helloWorld"},
		{"fooBarBaz", "fooBarBaz"},
		{"abcDefGhi", "abcDefGhi"},
		{"HelloWorld", "helloWorld"},
		{"FooBarBaz", "fooBarBaz"},
		{"AbcDefGhi", "abcDefGhi"},
		{"_hello_world", "_helloWorld"},
		{"_foo_bar_baz", "_fooBarBaz"},
		{"_abc_def_ghi", "_abcDefGhi"},
		{"CAPITALS_WITH_UNDERSCORES", "capitalsWithUnderscores"},
		{"_CAPITALS_WITH_UNDERSCORES", "_capitalsWithUnderscores"},
		{"hello_world_", "helloWorld_"},
		{"hello__world", "helloWorld"},
		{"_hello__world", "_helloWorld"},
		{"__hello__world", "__helloWorld"},
		{"hello_world123", "helloWorld123"},
		{"hello_world_123", "helloWorld123"},
	}

	for _, test := range tests {
		result := TocamelCase(test.input)
		if result != test.expected {
			t.Errorf("TocamelCase(%s) = %s, expected %s", test.input, result, test.expected)
		}
	}
}

func TestTosnake_case(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"PropellerAData", "propeller_a_data"},
		{"propellerAData", "propeller_a_data"},
		{"propeller_a_data", "propeller_a_data"},
		{"PROPELLER_A_DATA", "propeller_a_data"},
		{"HelloWorld", "hello_world"},
		{"FooBarBaz", "foo_bar_baz"},
		{"AbcDefGhi", "abc_def_ghi"},
		{"hello_world", "hello_world"},
		{"foo_bar_baz", "foo_bar_baz"},
		{"abc_def_ghi", "abc_def_ghi"},
		{"____He___ll__o__", "____he_ll_o__"},
		{"", ""},
		{"helloWorld", "hello_world"},
		{"fooBarBaz", "foo_bar_baz"},
		{"abcDefGhi", "abc_def_ghi"},
		{"_helloWorld", "_hello_world"},
		{"_fooBarBaz", "_foo_bar_baz"},
		{"_abcDefGhi", "_abc_def_ghi"},
		{"CAPITALS_WITH_UNDERSCORES", "capitals_with_underscores"},
		{"_CAPITALS_WITH_UNDERSCORES", "_capitals_with_underscores"},
		{"HelloWorld_", "hello_world_"},
		{"Hello__World", "hello_world"},
		{"_Hello__World", "_hello_world"},
		{"__Hello__World", "__hello_world"},
		{"HelloWorld123", "hello_world123"},
		{"HelloWorld_123", "hello_world_123"},
	}

	for _, test := range tests {
		result := Tosnake_case(test.input)
		if result != test.expected {
			t.Errorf("Tosnake_case(%s) = %s, expected %s", test.input, result, test.expected)
		}
	}
}

func TestToALLCAP_CASE(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"PropellerAData", "PROPELLER_A_DATA"},
		{"propellerAData", "PROPELLER_A_DATA"},
		{"propeller_a_data", "PROPELLER_A_DATA"},
		{"PROPELLER_A_DATA", "PROPELLER_A_DATA"},
		{"HelloWorld", "HELLO_WORLD"},
		{"FooBarBaz", "FOO_BAR_BAZ"},
		{"AbcDefGhi", "ABC_DEF_GHI"},
		{"hello_world", "HELLO_WORLD"},
		{"foo_bar_baz", "FOO_BAR_BAZ"},
		{"abc_def_ghi", "ABC_DEF_GHI"},
		{"____He___ll__o__", "____HE_LL_O__"},
		{"", ""},
		{"helloWorld", "HELLO_WORLD"},
		{"fooBarBaz", "FOO_BAR_BAZ"},
		{"abcDefGhi", "ABC_DEF_GHI"},
		{"_helloWorld", "_HELLO_WORLD"},
		{"_fooBarBaz", "_FOO_BAR_BAZ"},
		{"_abcDefGhi", "_ABC_DEF_GHI"},
		{"CAPITALS_WITH_UNDERSCORES", "CAPITALS_WITH_UNDERSCORES"},
		{"_CAPITALS_WITH_UNDERSCORES", "_CAPITALS_WITH_UNDERSCORES"},
		{"HelloWorld_", "HELLO_WORLD_"},
		{"Hello__World", "HELLO_WORLD"},
		{"_Hello__World", "_HELLO_WORLD"},
		{"__Hello__World", "__HELLO_WORLD"},
		{"HelloWorld123", "HELLO_WORLD123"},
		{"HelloWorld_123", "HELLO_WORLD_123"},
	}

	for _, test := range tests {
		result := ToALLCAP_CASE(test.input)
		if result != test.expected {
			t.Errorf("ToALLCAP_CASE(%s) = %s, expected %s", test.input, result, test.expected)
		}
	}
}

func TestMuxNameStyle(t *testing.T) {
	generator := func(length int) string {
		str := ""
		for i := 0; i < length; i++ {
			ty := rand.Intn(4)
			switch ty {
			case 0:
				str += "_"
			case 1:
				str += string(rune('A' + rand.Intn(26)))
			case 2:
				str += string(rune('a' + rand.Intn(26)))
			case 3:
				str += string(rune('0' + rand.Intn(10)))
			}
		}
		return str
	}

	funcList := map[string]func(string) string{
		"ToPascalCase":  ToPascalCase,
		"TocamelCase":   TocamelCase,
		"Tosnake_case":  Tosnake_case,
		"ToALLCAP_CASE": ToALLCAP_CASE,
	}

	caseNum := 2000
	caseLen := 2000
	for i := 0; i < caseNum; i++ {
		caseStr := generator(caseLen)

		for baseName, baseFunc := range funcList {
			baseStr := baseFunc(baseFunc(caseStr))
			// for otherName, otherFunc := range funcList {
			//     result := baseFunc(otherFunc(caseStr))
			//     if result != baseStr {
			//         t.Errorf("MuxNameStyle(%s, %s, %s) = %s, expected %s", baseName, otherName, caseStr, result, baseStr)
			//     }
			// }
			result := baseFunc(baseStr)
			if result != baseStr {
				t.Errorf("MuxNameStyle(%s, %s) = %s, expected %s", baseName, caseStr, result, baseStr)
			}
		}
	}
}
