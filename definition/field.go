package definition

import (
	"fmt"

	"github.com/xaxys/bubbler/util"
)

// ==================== FieldKindID ====================

type FieldKindID int

const (
	FieldKindID_Normal FieldKindID = iota
	FieldKindID_Void
	FieldKindID_Embedded
	FieldKindID_Constant
)

func (f FieldKindID) IsNormal() bool {
	return f == FieldKindID_Normal
}

func (f FieldKindID) IsVoid() bool {
	return f == FieldKindID_Void
}

func (f FieldKindID) IsEmbedded() bool {
	return f == FieldKindID_Embedded
}

func (f FieldKindID) IsConstant() bool {
	return f == FieldKindID_Constant
}

var FieldKindIDStringsMap = map[FieldKindID]string{
	FieldKindID_Normal:   "NormalField",
	FieldKindID_Void:     "VoidField",
	FieldKindID_Embedded: "EmbeddedField",
	FieldKindID_Constant: "ConstantField",
}

func (f FieldKindID) String() string {
	if str, ok := FieldKindIDStringsMap[f]; ok {
		return str
	}
	return fmt.Sprintf("FieldKindID(%d)", f)
}

// ==================== Field ====================

type Field interface {
	Position
	GetFieldKind() FieldKindID
	GetFieldUniqueName() string
	// FieldBitSize only indicates the size defined in the field,
	// which may be different from the actual field type size.
	// For example, a field of type "int32" may have a FieldBitSize
	// of 16, although its type size is 32.
	// Or a embedded field may have GetFieldBitSize return 0. But
	// stored actual size in FieldBitSize.
	// Or a variable length field may have GetFieldBitSize return -1.
	GetFieldBitSize() int64
	GetFieldBelongs() *Struct
	SetFieldBelongs(*Struct)
	Copy() Field
}

type NormalField struct {
	BasePosition

	FieldName    string
	FieldType    Type
	FieldBitSize int64
	FieldMethods *util.OrderedMap[string, *util.OrderedMap[MethodKindID, Method]]
	FieldOptions *util.OrderedMap[string, *Option]
	FieldBelongs *Struct
}

func (f NormalField) GetFieldKind() FieldKindID {
	return FieldKindID_Normal
}

func (f NormalField) String() string {
	return fmt.Sprintf("%s %s[%s]", f.FieldType.GetTypeName(), f.FieldName, util.ToSizeString(f.FieldBitSize))
}

func (f NormalField) DumpString() string {
	methods := "[\n"
	for _, method := range f.FieldMethods.Values() {
		methods += util.IndentSpace8(method) + "\n"
	}
	methods += "    ]"

	options := "[\n"
	for _, option := range f.FieldOptions.Values() {
		options += util.IndentSpace8(option) + "\n"
	}
	options += "    ]"

	typeStr := util.IndentSpace4NoFirst(f.FieldType)

	belongsStr := "<nil>"
	if f.FieldBelongs != nil {
		belongsStr = f.FieldBelongs.StructName
	}

	return fmt.Sprintf("NormalField {\n    FieldName: %s\n    FieldType: %s\n    FieldBitSize: %d\n    FieldMethods: %s\n    FieldOptions: %s\n    FieldBelongs: %s\n}", f.FieldName, typeStr, f.FieldBitSize, methods, options, belongsStr)
}

func (f NormalField) GetFieldUniqueName() string {
	return f.FieldName
}

func (f NormalField) GetFieldBitSize() int64 {
	return f.FieldBitSize
}

func (f NormalField) GetFieldBelongs() *Struct {
	return f.FieldBelongs
}

func (f *NormalField) SetFieldBelongs(s *Struct) {
	f.FieldBelongs = s
}

func (f NormalField) Copy() Field {
	newField := NormalField{
		BasePosition: f.BasePosition,
		FieldName:    f.FieldName,
		FieldType:    f.FieldType,
		FieldBitSize: f.FieldBitSize,
		FieldMethods: f.FieldMethods,
		FieldOptions: f.FieldOptions,
		FieldBelongs: f.FieldBelongs,
	}
	return &newField
}

type VoidField struct {
	BasePosition

	FieldBitSize int64
	FieldOptions *util.OrderedMap[string, *Option]
	FieldBelongs *Struct
}

func (f VoidField) GetFieldKind() FieldKindID {
	return FieldKindID_Void
}

func (f VoidField) String() string {
	return fmt.Sprintf("void [%s]", util.ToSizeString(f.FieldBitSize))
}

func (f VoidField) DumpString() string {
	options := "[\n"
	for _, option := range f.FieldOptions.Values() {
		options += util.IndentSpace8(option) + "\n"
	}
	options += "    ]"

	belongsStr := "<nil>"
	if f.FieldBelongs != nil {
		belongsStr = f.FieldBelongs.StructName
	}

	return fmt.Sprintf("VoidField {\n    FieldBitSize: %d\n    FieldOptions: %s\n    FieldBelongs: %s\n}", f.FieldBitSize, options, belongsStr)
}

func (f VoidField) GetFieldUniqueName() string {
	return ""
}

func (f VoidField) GetFieldBitSize() int64 {
	return f.FieldBitSize
}

func (f VoidField) GetFieldBelongs() *Struct {
	return f.FieldBelongs
}

func (f *VoidField) SetFieldBelongs(s *Struct) {
	f.FieldBelongs = s
}

func (f VoidField) Copy() Field {
	newField := VoidField{
		BasePosition: f.BasePosition,
		FieldBitSize: f.FieldBitSize,
		FieldOptions: f.FieldOptions,
		FieldBelongs: f.FieldBelongs,
	}
	return &newField
}

type EmbeddedField struct {
	BasePosition

	FieldType    *Struct
	FieldBitSize int64
	FieldOptions *util.OrderedMap[string, *Option]
	FieldBelongs *Struct
}

func (f EmbeddedField) GetFieldKind() FieldKindID {
	return FieldKindID_Embedded
}

func (f EmbeddedField) String() string {
	return fmt.Sprintf("%s [%s]", f.FieldType.StructName, util.ToSizeString(f.FieldBitSize))
}

func (f EmbeddedField) DumpString() string {
	typeStr := util.IndentSpace4NoFirst(f.FieldType)

	options := "[\n"
	for _, option := range f.FieldOptions.Values() {
		options += util.IndentSpace8(option) + "\n"
	}
	options += "    ]"

	belongsStr := "<nil>"
	if f.FieldBelongs != nil {
		belongsStr = f.FieldBelongs.StructName
	}

	return fmt.Sprintf("EmbeddedField {\n    FieldType: %s\n    FieldBitSize: %d\n    FieldOptions: %s\n    FieldBelongs: %s\n}", typeStr, f.FieldBitSize, options, belongsStr)
}

func (f EmbeddedField) GetFieldUniqueName() string {
	return f.FieldType.StructName
}

func (f EmbeddedField) GetFieldBitSize() int64 {
	// EmbeddedField is not a real field, so it has no bit size
	// It is only used to flatten embedded struct fields to toplevel
	// You can get bit size directly from EmbeddedField.FieldBitSize
	return 0
}

func (f EmbeddedField) GetFieldBelongs() *Struct {
	return f.FieldBelongs
}

func (f *EmbeddedField) SetFieldBelongs(s *Struct) {
	f.FieldBelongs = s
}

func (f EmbeddedField) Copy() Field {
	newField := EmbeddedField{
		BasePosition: f.BasePosition,
		FieldType:    f.FieldType,
		FieldBitSize: f.FieldBitSize,
		FieldOptions: f.FieldOptions,
		FieldBelongs: f.FieldBelongs,
	}
	return &newField
}

type ConstantField struct {
	BasePosition

	FieldName     string // may be empty
	FieldType     *BasicType
	FieldBitSize  int64
	FieldConstant Literal
	FieldOptions  *util.OrderedMap[string, *Option]
	FieldBelongs  *Struct
}

func (f ConstantField) GetFieldKind() FieldKindID {
	return FieldKindID_Constant
}

func (f ConstantField) String() string {
	return fmt.Sprintf("%s %s[%s] = %v", f.FieldType.GetTypeName(), f.FieldName, util.ToSizeString(f.FieldBitSize), f.FieldConstant)
}

func (f ConstantField) DumpString() string {
	typeStr := util.IndentSpace4NoFirst(f.FieldType)

	options := "[\n"
	for _, option := range f.FieldOptions.Values() {
		options += util.IndentSpace8(option) + "\n"
	}
	options += "    ]"

	belongsStr := "<nil>"
	if f.FieldBelongs != nil {
		belongsStr = f.FieldBelongs.StructName
	}

	return fmt.Sprintf("ConstantField {\n    FieldName: %s\n    FieldType: %s\n    FieldBitSize: %d\n    FieldConstant: %s\n    FieldOptions: %s\n    FieldBelongs: %s\n}", f.FieldName, typeStr, f.FieldBitSize, f.FieldConstant, options, belongsStr)
}

func (f ConstantField) GetFieldUniqueName() string {
	return f.FieldName // may be empty
}

func (f ConstantField) GetFieldBitSize() int64 {
	return f.FieldBitSize
}

func (f ConstantField) GetFieldBelongs() *Struct {
	return f.FieldBelongs
}

func (f *ConstantField) SetFieldBelongs(s *Struct) {
	f.FieldBelongs = s
}

func (f ConstantField) Copy() Field {
	newField := ConstantField{
		BasePosition:  f.BasePosition,
		FieldName:     f.FieldName,
		FieldType:     f.FieldType,
		FieldBitSize:  f.FieldBitSize,
		FieldConstant: f.FieldConstant,
		FieldOptions:  f.FieldOptions,
		FieldBelongs:  f.FieldBelongs,
	}
	return &newField
}
