package definition

import (
	"fmt"

	"github.com/xaxys/bubbler/util"
)

// ==================== Field ====================

type FieldKindID int

const (
	FieldKindID_Normal FieldKindID = iota
	FieldKindID_Void
	FieldKindID_Embedded
	FieldKindID_Constant
)

type Field interface {
	Position
	GetFieldKind() FieldKindID
	GetFieldUniqueName() string
	GetFieldBitSize() int64
	GetFieldBelongs() *Struct
	// GetFieldFromEmbedded() *EmbeddedField
	SetFieldBelongs(*Struct)
	// SetFieldFromEmbedded(*EmbeddedField)
	Copy() Field
}

type NormalField struct {
	BasePosition
	FieldName    string
	FieldType    Type
	FieldBitSize int64
	FieldBelongs *Struct
	FieldMethods []*Method
	FieldOptions *util.OrderedMap[string, *Option]
	FromEmbedded *EmbeddedField
}

func (f NormalField) GetFieldKind() FieldKindID {
	return FieldKindID_Normal
}

func (f NormalField) ShortString() string {
	return fmt.Sprintf("%s %s[%s]", f.FieldType.GetTypeName(), f.FieldName, util.ToSizeString(f.FieldBitSize))
}

func (f NormalField) String() string {
	methods := "[\n"
	for _, method := range f.FieldMethods {
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

	fromEmbeddedStr := "<nil>"
	if f.FromEmbedded != nil {
		fromEmbeddedStr = f.FromEmbedded.FieldType.StructName
	}

	return fmt.Sprintf("NormalField {\n    FieldName: %s\n    FieldType: %s\n    FieldBitSize: %d\n    FieldMethods: %s\n    FieldOptions: %s\n    FieldBelongs: %s\n    FromEmbedded: %s\n}", f.FieldName, typeStr, f.FieldBitSize, methods, options, belongsStr, fromEmbeddedStr)
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

func (f NormalField) GetFieldFromEmbedded() *EmbeddedField {
	return f.FromEmbedded
}

func (f *NormalField) SetFieldBelongs(s *Struct) {
	f.FieldBelongs = s
}

func (f *NormalField) SetFieldFromEmbedded(e *EmbeddedField) {
	f.FromEmbedded = e
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
		FromEmbedded: f.FromEmbedded,
	}
	return &newField
}

type VoidField struct {
	BasePosition
	FieldBitSize int64
	FieldBelongs *Struct
	FieldOptions *util.OrderedMap[string, *Option]
	FromEmbedded *EmbeddedField
}

func (f VoidField) GetFieldKind() FieldKindID {
	return FieldKindID_Void
}

func (f VoidField) ShortString() string {
	return fmt.Sprintf("void [%s]", util.ToSizeString(f.FieldBitSize))
}

func (f VoidField) String() string {
	options := "[\n"
	for _, option := range f.FieldOptions.Values() {
		options += util.IndentSpace8(option) + "\n"
	}
	options += "    ]"

	belongsStr := "<nil>"
	if f.FieldBelongs != nil {
		belongsStr = f.FieldBelongs.StructName
	}

	fromEmbeddedStr := "<nil>"
	if f.FromEmbedded != nil {
		fromEmbeddedStr = f.FromEmbedded.FieldType.StructName
	}

	return fmt.Sprintf("VoidField {\n    FieldBitSize: %d\n    FieldOptions: %s\n    FieldBelongs: %s\n    FromEmbedded: %s\n}", f.FieldBitSize, options, belongsStr, fromEmbeddedStr)
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

func (f VoidField) GetFieldFromEmbedded() *EmbeddedField {
	return f.FromEmbedded
}

func (f *VoidField) SetFieldBelongs(s *Struct) {
	f.FieldBelongs = s
}

func (f *VoidField) SetFieldFromEmbedded(e *EmbeddedField) {
	f.FromEmbedded = e
}

func (f VoidField) Copy() Field {
	newField := VoidField{
		BasePosition: f.BasePosition,
		FieldBitSize: f.FieldBitSize,
		FieldBelongs: f.FieldBelongs,
		FieldOptions: f.FieldOptions,
		FromEmbedded: f.FromEmbedded,
	}
	return &newField
}

type EmbeddedField struct {
	BasePosition
	FieldType    *Struct
	FieldBitSize int64
	FieldBelongs *Struct
	FieldOptions *util.OrderedMap[string, *Option]
	FromEmbedded *EmbeddedField
}

func (f EmbeddedField) GetFieldKind() FieldKindID {
	return FieldKindID_Embedded
}

func (f EmbeddedField) ShortString() string {
	return fmt.Sprintf("%s [%s]", f.FieldType.StructName, util.ToSizeString(f.FieldBitSize))
}

func (f EmbeddedField) String() string {
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

	fromEmbeddedStr := "<nil>"
	if f.FromEmbedded != nil {
		fromEmbeddedStr = f.FromEmbedded.FieldType.StructName
	}

	return fmt.Sprintf("EmbeddedField {\n    FieldType: %s\n    FieldBitSize: %d\n    FieldOptions: %s\n    FieldBelongs: %s\n    FromEmbedded: %s\n}", typeStr, f.FieldBitSize, options, belongsStr, fromEmbeddedStr)
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

func (f EmbeddedField) GetFieldFromEmbedded() *EmbeddedField {
	return f.FromEmbedded
}

func (f *EmbeddedField) SetFieldBelongs(s *Struct) {
	f.FieldBelongs = s
}

func (f *EmbeddedField) SetFieldFromEmbedded(e *EmbeddedField) {
	f.FromEmbedded = e
}

func (f EmbeddedField) Copy() Field {
	newField := EmbeddedField{
		BasePosition: f.BasePosition,
		FieldType:    f.FieldType,
		FieldBitSize: f.FieldBitSize,
		FieldBelongs: f.FieldBelongs,
		FieldOptions: f.FieldOptions,
		FromEmbedded: f.FromEmbedded,
	}
	return &newField
}

type ConstantField struct {
	BasePosition
	FieldName     string // may be empty
	FieldType     *BasicType
	FieldBitSize  int64
	FieldConstant Literal
	FieldBelongs  *Struct
	FieldOptions  *util.OrderedMap[string, *Option]
	FromEmbedded  *EmbeddedField
}

func (f ConstantField) GetFieldKind() FieldKindID {
	return FieldKindID_Constant
}

func (f ConstantField) ShortString() string {
	return fmt.Sprintf("%s %s[%s] = %v", f.FieldType.GetTypeName(), f.FieldName, util.ToSizeString(f.FieldBitSize), f.FieldConstant)
}

func (f ConstantField) String() string {
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

	fromEmbeddedStr := "<nil>"
	if f.FromEmbedded != nil {
		fromEmbeddedStr = f.FromEmbedded.FieldType.StructName
	}

	return fmt.Sprintf("ConstantField {\n    FieldName: %s\n    FieldType: %s\n    FieldBitSize: %d\n    FieldConstant: %s\n    FieldOptions: %s\n    FieldBelongs: %s\n    FromEmbedded: %s\n}", f.FieldName, typeStr, f.FieldBitSize, f.FieldConstant, options, belongsStr, fromEmbeddedStr)
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

func (f ConstantField) GetFieldFromEmbedded() *EmbeddedField {
	return f.FromEmbedded
}

func (f *ConstantField) SetFieldBelongs(s *Struct) {
	f.FieldBelongs = s
}

func (f *ConstantField) SetFieldFromEmbedded(e *EmbeddedField) {
	f.FromEmbedded = e
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
		FromEmbedded:  f.FromEmbedded,
	}
	return &newField
}
