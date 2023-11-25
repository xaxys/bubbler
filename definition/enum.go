package definition

import (
	"fmt"

	"github.com/xaxys/bubbler/util"
)

type EnumValue struct {
	BasePosition

	EnumValueName string
	EnumValue     int64
	EnumBelongs   *Enum
}

func (e EnumValue) String() string {
	return fmt.Sprintf("EnumValue {\n    EnumValueName: %s\n    EnumValue: %d\n}", e.EnumValueName, e.EnumValue)
}

type Enum struct {
	BasePosition

	EnumName    string
	EnumBitSize int64
	EnumValues  []*EnumValue
}

func (e Enum) ShortString() string {
	return fmt.Sprintf("%s [%s] {%d values}", e.EnumName, util.ToSizeString(e.EnumBitSize), len(e.EnumValues))
}

func (e Enum) String() string {
	values := "[\n"
	for _, value := range e.EnumValues {
		values += util.IndentSpace8(value) + "\n"
	}
	values += "    ]"

	return fmt.Sprintf("Enum {\n    EnumName: %s\n    EnumBitSize: %d\n    EnumValues: %s\n}", e.EnumName, e.EnumBitSize, values)
}

// ==================== EnumType ====================

func (e *Enum) GetTypeID() TypeID {
	return TypeID_Enum
}

func (e *Enum) GetTypeName() string {
	return e.EnumName
}

func (e *Enum) GetTypeBitSize() int64 {
	return e.EnumBitSize
}

func (e *Enum) GetEnumValueByName(name string) *EnumValue {
	for _, enumValue := range e.EnumValues {
		if enumValue.EnumValueName == name {
			return enumValue
		}
	}
	return nil
}

func (e *Enum) GetEnumValueByValue(value int64) *EnumValue {
	for _, enumValue := range e.EnumValues {
		if enumValue.EnumValue == value {
			return enumValue
		}
	}
	return nil
}
