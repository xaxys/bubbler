package util

import (
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
		{"hello_world", "HelloWorld"},
		{"foo_bar_baz", "FooBarBaz"},
		{"abc_def_ghi", "AbcDefGhi"},
		{"____He___ll__o_", "____HeLlO"},
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
		{"hello_world_", "HelloWorld"},
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
		{"hello_world", "helloWorld"},
		{"foo_bar_baz", "fooBarBaz"},
		{"abc_def_ghi", "abcDefGhi"},
		{"____He___ll__o_", "____heLlO"},
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
		{"hello_world_", "helloWorld"},
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
		{"HelloWorld", "hello_world"},
		{"FooBarBaz", "foo_bar_baz"},
		{"AbcDefGhi", "abc_def_ghi"},
		{"hello_world", "hello_world"},
		{"foo_bar_baz", "foo_bar_baz"},
		{"abc_def_ghi", "abc_def_ghi"},
		{"____He___ll__o_", "____he_ll_o_"},
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
