package compiler

import (
	"github.com/xaxys/bubbler/definition"
	"github.com/xaxys/bubbler/util"
)

var FileOptionValidator *OptionValidator
var FieldOptionValidator *OptionValidator
var EnumValueOptionValidator *OptionValidator

func init() {
	FileOptionValidator = NewOptionValidator()
	FileOptionValidator.AddOption("csharp_namespace", definition.LiteralKindID_String)
	FileOptionValidator.AddOption("cpp_namespace", definition.LiteralKindID_String)
	FileOptionValidator.AddOption("go_package", definition.LiteralKindID_String)
	FileOptionValidator.AddOption("java_package", definition.LiteralKindID_String)

	FieldOptionValidator = NewOptionValidator()
	FieldOptionValidator.AddOption("order", definition.LiteralKindID_String, "big", "little")

	EnumValueOptionValidator = NewOptionValidator()
}

// ==================== OptionValidator ====================

type OptionValidator struct {
	OptionType   *util.OrderedMap[string, definition.LiteralKindID]
	OptionValues *util.OrderedMap[string, *util.OrderedMap[any, any]]
}

func NewOptionValidator() *OptionValidator {
	return &OptionValidator{
		OptionType:   util.NewOrderedMap[string, definition.LiteralKindID](),
		OptionValues: util.NewOrderedMap[string, *util.OrderedMap[any, any]](),
	}
}

func (v *OptionValidator) AddOption(name string, ty definition.LiteralKindID, values ...any) {
	v.OptionType.Put(name, ty)
	if len(values) > 0 {
		valuesMap := util.NewOrderedMap[any, any]()
		for _, value := range values {
			valuesMap.Put(value, nil)
		}
		v.OptionValues.Put(name, valuesMap)
	}
}

func (v *OptionValidator) ValidateOption(option *definition.Option) error {
	checked := false

	if optionType, ok := v.OptionType.Get(option.OptionName); ok {
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
	if optionValues, ok := v.OptionValues.Get(option.OptionName); ok {
		checked = true
		if _, ok := optionValues.Get(value); !ok {
			return &definition.CompileError{
				Position: option.BasePosition,
				Err: &definition.OptionValueError{
					OptionName: option.OptionName,
					Expect:     optionValues.Keys(),
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
