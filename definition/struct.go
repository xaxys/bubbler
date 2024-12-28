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
	StructDynamic bool
	StructFields  *util.OrderedMap[string, Field]
	StructBelongs *CompilationUnit
}

func (s Struct) String() string {
	return fmt.Sprintf("%s [%s] {%d fields}", s.StructName, util.ToSizeStringDynamic(s.StructBitSize, s.StructDynamic, 0), s.StructFields.Len())
}

func (s Struct) DumpString() string {
	fields := "[\n"
	for _, field := range s.StructFields.Values() {
		fields += util.IndentSpace8(field) + "\n"
	}
	fields += "    ]"

	return fmt.Sprintf("Struct {\n    StructName: %s\n    StructDynamic: %t\n    StructBitSize: %s\n    StructFields: %s\n}", s.StructName, s.StructDynamic, util.ToSizeString(s.StructBitSize), fields)
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

func (s Struct) GetTypeDynamic() bool {
	return s.StructDynamic
}

func (s Struct) GetBelongs() *CompilationUnit {
	return s.StructBelongs
}

func (s *Struct) SetBelongs(c *CompilationUnit) {
	s.StructBelongs = c
	for _, field := range s.StructFields.Values() {
		if normalField, ok := field.(*NormalField); ok {
			switch ty := normalField.FieldType.(type) {
			case CustomType:
				if ty.GetBelongs() == nil {
					ty.SetBelongs(c)
				}
			case *Array:
				if customType, ok := ty.ElementType.(CustomType); ok {
					if customType.GetBelongs() == nil {
						customType.SetBelongs(c)
					}
				}
			}
		}
	}
}

// ==================== Struct ====================

func (s *Struct) SumFieldBitSize() (fixedSize int64, hasDynamic bool) {
	sum := int64(0)
	dynamic := false
	for _, field := range s.StructFields.Values() {
		switch ty := field.(type) {
		case *ConstantField:
		case *NormalField:
			if ty.FieldType.GetTypeDynamic() {
				dynamic = true
			}
		}
		if field.GetFieldBitSize() > 0 {
			sum += field.GetFieldBitSize()
		}
	}
	return sum, dynamic
}

func (s *Struct) ForEachField(f func(field Field, index int, fixedStart int64, isDynamicStart bool) error) error {
	startBit := int64(0)
	dynamicStart := false
	for i, field := range s.StructFields.Values() {
		if err := f(field, i, startBit, dynamicStart); err != nil {
			return err
		}
		if field.GetFieldBitSize() != -1 {
			startBit += field.GetFieldBitSize()
		} else {
			dynamicStart = true
		}
	}
	return nil
}

func (s *Struct) ForEachFieldWithPos(f func(field Field, index int, fixedStart int64, isDynamicStart bool, pos string) error) error {
	ToPosStr := func(start int64, size int64, isDynamicStart bool) string {
		if size == 0 {
			return "[virtual]"
		}
		startStr := util.ToSizeStringDynamic(start, isDynamicStart, 0)
		endStr := util.ToSizeStringDynamic(start, isDynamicStart, size)
		return fmt.Sprintf("[%s:%s)", startStr, endStr)
	}

	return s.ForEachField(func(field Field, index int, fixedStart int64, isDynamicStart bool) error {
		return f(field, index, fixedStart, isDynamicStart, ToPosStr(fixedStart, field.GetFieldBitSize(), isDynamicStart))
	})
}
