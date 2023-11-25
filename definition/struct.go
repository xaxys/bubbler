package definition

import (
	"fmt"

	"github.com/xaxys/bubbler/util"
)

// ==================== Struct ====================

type Struct struct {
	BasePosition

	StructName    string
	StructBitSize int64
	StructFields  []Field
}

func (s Struct) ShortString() string {
	return fmt.Sprintf("%s [%s] {%d fields}", s.StructName, util.ToSizeString(s.StructBitSize), len(s.StructFields))
}

func (s Struct) String() string {
	fields := "[\n"
	for _, field := range s.StructFields {
		fields += util.IndentSpace8(field) + "\n"
	}
	fields += "    ]"

	return fmt.Sprintf("Struct {\n    StructName: %s\n    StructBitSize: %d\n    StructFields: %s\n}", s.StructName, s.StructBitSize, fields)
}

// ==================== For StructType ====================

func (s Struct) GetTypeID() TypeID {
	return TypeID_Struct
}

func (s Struct) GetTypeName() string {
	return s.StructName
}

func (s Struct) GetTypeBitSize() int64 {
	return s.StructBitSize
}

// ==================== Struct ====================

func (s *Struct) GetFieldByName(name string) Field {
	for _, field := range s.StructFields {
		switch val := field.(type) {
		case *NormalField:
			if val.FieldName == name {
				return val
			}
		}
	}
	return nil
}

func (s *Struct) GetFieldByUniqueName(name string) Field {
	for _, field := range s.StructFields {
		if field.GetFieldUniqueName() == name {
			return field
		}
	}
	return nil
}

func (s *Struct) SumFieldBitSize() (fixed_size int64, has_dynamic bool) {
	sum := int64(0)
	dynamic := false
	for _, field := range s.StructFields {
		if field.GetFieldBitSize() == -1 {
			dynamic = true
		}
		sum += field.GetFieldBitSize()
	}
	return sum, dynamic
}

func (s *Struct) GetFieldsStartBit() []int64 {
	startBits := []int64{}
	startBit := int64(0)
	for _, field := range s.StructFields {
		startBits = append(startBits, startBit)
		startBit += field.GetFieldBitSize()
	}
	return startBits
}
