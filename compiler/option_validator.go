package compiler

import (
	"github.com/xaxys/bubbler/definition"
)

var optionTypeMap = map[string]definition.LiteralKindID{
	"order": definition.LiteralKindID_String,
}

var optionValuesMap = map[string]map[any]any{
	"order": {
		"big":    nil,
		"little": nil,
	},
}

func MapKeys(m map[any]any) []any {
	keys := make([]any, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func ValidateOption(option *definition.Option) error {
	checked := false

	if optionType, ok := optionTypeMap[option.OptionName]; ok {
		checked = true
		ty := option.OptionValue.GetLiteralKind()
		if ty != optionType {
			return &definition.CompileError{
				Position: option.BasePosition,
				Err: &definition.OptionTypeError{
					OptionName: option.OptionName,
					Expect:     optionType.String(),
					Got:        ty.String(),
				},
			}
		}
	}

	value := option.OptionValue.GetLiteralValue()
	if optionValues, ok := optionValuesMap[option.OptionName]; ok {
		checked = true
		if _, ok := optionValues[value]; !ok {
			return &definition.CompileError{
				Position: option.BasePosition,
				Err: &definition.OptionValueError{
					OptionName: option.OptionName,
					Expect:     MapKeys(optionValues),
					Got:        value,
				},
			}
		}
	}

	if !checked {
		return &definition.CompileWarning{
			Position: option.BasePosition,
			Warning: &definition.OptionUnknownWarning{
				OptionName: option.OptionName,
			},
		}
	}

	return nil
}
