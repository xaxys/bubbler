package gen

import (
	"github.com/xaxys/bubbler/definition"
	"github.com/xaxys/bubbler/util"
)

func MatchOption(options *util.OrderedMap[string, *definition.Option], name string, value any) bool {
	if options == nil {
		return false
	}
	option, ok := options.Get(name)
	if !ok {
		return false
	}
	return option.OptionValue.GetLiteralValue() == value
}
