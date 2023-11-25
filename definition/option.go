package definition

import "fmt"

type Option struct {
	BasePosition

	OptionName  string
	OptionValue Literal
}

func (o Option) String() string {
	return fmt.Sprintf("%s = %s", o.OptionName, o.OptionValue)
}
