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
	StructFields  *util.OrderedMap[string, Field]
}

func (s Struct) String() string {
	return fmt.Sprintf("%s [%s] {%d fields}", s.StructName, util.ToSizeString(s.StructBitSize), s.StructFields.Len())
}

func (s Struct) DumpString() string {
	fields := "[\n"
	for _, field := range s.StructFields.Values() {
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

func (s *Struct) SumFieldBitSize() (fixed_size int64, has_dynamic bool) {
	sum := int64(0)
	dynamic := false
	for _, field := range s.StructFields.Values() {
		if field.GetFieldBitSize() == -1 {
			dynamic = true
		} else {
			sum += field.GetFieldBitSize()
		}
	}
	return sum, dynamic
}

func (s *Struct) ForEachField(f func(field Field, index int, start int64) error) error {
	startBit := int64(0)
	for i, field := range s.StructFields.Values() {
		if err := f(field, i, startBit); err != nil {
			return err
		}
		if startBit != -1 && field.GetFieldBitSize() != -1 {
			startBit += field.GetFieldBitSize()
		} else {
			startBit = -1
		}
	}
	return nil
}

func (s *Struct) ForEachFieldWithPos(f func(field Field, index int, start int64, pos string) error) error {
	ToPosStr := func(start int64, size int64) string {
		if size == 0 {
			return "[virtual]"
		}
		if start == -1 {
			return fmt.Sprintf("[dynamic:dynamic):%s", util.ToSizeString(size))
		}
		if size == -1 {
			return fmt.Sprintf("[%s:dynamic)", util.ToSizeString(start))
		}
		return fmt.Sprintf("[%s:%s)", util.ToSizeString(start), util.ToSizeString(start+size))
	}

	return s.ForEachField(func(field Field, index int, start int64) error {
		return f(field, index, start, ToPosStr(start, field.GetFieldBitSize()))
	})
}
