package gen

import (
	"strings"

	"github.com/xaxys/bubbler/definition"
)

type GenOptionSetter func(*GenOptions)

type SignExtMethodID int

const (
	SignExtMethodDefault SignExtMethodID = iota
	SignExtMethodShift
	SignExtMethodArith
)

type GenOptions struct {
	RemovePath    []string
	InnerClass    bool
	SingleFile    bool
	MinimalCode   bool
	DecimalNumber bool
	MemoryCopy    bool
	SignExtMethod SignExtMethodID
}

func NewGenOptions(setter ...GenOptionSetter) *GenOptions {
	options := &GenOptions{
		RemovePath:    nil,
		SingleFile:    false,
		InnerClass:    false,
		MinimalCode:   false,
		DecimalNumber: false,
		MemoryCopy:    false,
		SignExtMethod: SignExtMethodDefault,
	}
	for _, s := range setter {
		s(options)
	}
	return options
}

func RemovePath(path string) GenOptionSetter {
	paths := strings.Split(path, ",")
	for _, p := range paths {
		for strings.HasSuffix(p, "/") {
			p = p[:len(p)-1]
		}
	}
	return func(options *GenOptions) {
		options.RemovePath = paths
	}
}

func SingleFile(single bool) GenOptionSetter {
	return func(options *GenOptions) {
		options.SingleFile = single
	}
}

func InnerClass(inner bool) GenOptionSetter {
	return func(options *GenOptions) {
		options.InnerClass = inner
	}
}

func MinimalCode(minimal bool) GenOptionSetter {
	return func(options *GenOptions) {
		options.MinimalCode = minimal
	}
}

func DecimalNumber(decnum bool) GenOptionSetter {
	return func(options *GenOptions) {
		options.DecimalNumber = decnum
	}
}

func MemoryCopy(memcopy bool) GenOptionSetter {
	return func(options *GenOptions) {
		options.MemoryCopy = memcopy
	}
}

var SignExtMethodMap = map[string]SignExtMethodID{
	"":        SignExtMethodDefault,
	"default": SignExtMethodDefault,
	"shift":   SignExtMethodShift,
	"arith":   SignExtMethodArith,
}

func SignExtMethod(signext string) (GenOptionSetter, error) {
	signext = strings.ToLower(signext)
	signextID, ok := SignExtMethodMap[signext]
	if !ok {
		return nil, &definition.GeneralError{
			Err: &definition.OptionValueError{
				OptionName: "signext",
				Expect:     []any{"shift", "arith"},
				Got:        signext,
			},
		}
	}
	return func(options *GenOptions) {
		options.SignExtMethod = signextID
	}, nil
}
