package definition

import "fmt"

type TypeID int

const (
	TypeID_Bool TypeID = iota
	TypeID_Uint8
	TypeID_Uint16
	TypeID_Uint32
	TypeID_Uint64
	TypeID_Int8
	TypeID_Int16
	TypeID_Int32
	TypeID_Int64
	TypeID_Float32
	TypeID_Float64
	TypeID_String
	TypeID_Bytes
	TypeID_Array
	TypeID_Struct
	TypeID_Enum
)

func (t TypeID) IsBasic() bool {
	return t >= TypeID_Bool && t <= TypeID_Float64
}

func (t TypeID) IsBool() bool {
	return t == TypeID_Bool
}

func (t TypeID) IsUint() bool {
	return t >= TypeID_Uint8 && t <= TypeID_Uint64
}

func (t TypeID) IsInt() bool {
	return t >= TypeID_Int8 && t <= TypeID_Int64
}

func (t TypeID) IsFloat() bool {
	return t >= TypeID_Float32 && t <= TypeID_Float64
}

func (t TypeID) IsIntOrUint() bool {
	return t.IsInt() || t.IsUint()
}

func (t TypeID) IsNumber() bool {
	return t.IsInt() || t.IsUint() || t.IsFloat()
}

func (t TypeID) IsArray() bool {
	return t == TypeID_Array
}

func (t TypeID) IsBytes() bool {
	return t == TypeID_Bytes
}

func (t TypeID) IsString() bool {
	return t == TypeID_String
}

func (t TypeID) IsStruct() bool {
	return t == TypeID_Struct
}

func (t TypeID) IsEnum() bool {
	return t == TypeID_Enum
}

func (t TypeID) String() string {
	switch t {
	case TypeID_Bool:
		return "bool"
	case TypeID_Uint8:
		return "uint8"
	case TypeID_Uint16:
		return "uint16"
	case TypeID_Uint32:
		return "uint32"
	case TypeID_Uint64:
		return "uint64"
	case TypeID_Int8:
		return "int8"
	case TypeID_Int16:
		return "int16"
	case TypeID_Int32:
		return "int32"
	case TypeID_Int64:
		return "int64"
	case TypeID_Float32:
		return "float32"
	case TypeID_Float64:
		return "float64"
	case TypeID_String:
		return "string"
	case TypeID_Bytes:
		return "bytes"
	case TypeID_Array:
		return "array"
	case TypeID_Struct:
		return "struct"
	case TypeID_Enum:
		return "enum"
	default:
		panic(fmt.Sprintf("unknown type id: %d", t))
	}
}

type CustomType interface {
	Position
	Type
	GetBelongs() *CompilationUnit
	SetBelongs(*CompilationUnit)
}

type Type interface {
	GetTypeID() TypeID
	GetTypeName() string
	GetTypeBitSize() int64
}

// ==================== BasicType ====================

type BasicType struct {
	TypeTypeID  TypeID
	TypeName    string
	TypeBitSize int64
}

func (t BasicType) String() string {
	return t.TypeName
}

func (t BasicType) GetTypeID() TypeID {
	return t.TypeTypeID
}

func (t BasicType) GetTypeName() string {
	return t.TypeName
}

func (t BasicType) GetTypeBitSize() int64 {
	return t.TypeBitSize
}

var (
	Bool    = BasicType{TypeTypeID: TypeID_Bool, TypeName: "bool", TypeBitSize: 8}
	Uint8   = BasicType{TypeTypeID: TypeID_Uint8, TypeName: "uint8", TypeBitSize: 8}
	Uint16  = BasicType{TypeTypeID: TypeID_Uint16, TypeName: "uint16", TypeBitSize: 16}
	Uint32  = BasicType{TypeTypeID: TypeID_Uint32, TypeName: "uint32", TypeBitSize: 32}
	Uint64  = BasicType{TypeTypeID: TypeID_Uint64, TypeName: "uint64", TypeBitSize: 64}
	Int8    = BasicType{TypeTypeID: TypeID_Int8, TypeName: "int8", TypeBitSize: 8}
	Int16   = BasicType{TypeTypeID: TypeID_Int16, TypeName: "int16", TypeBitSize: 16}
	Int32   = BasicType{TypeTypeID: TypeID_Int32, TypeName: "int32", TypeBitSize: 32}
	Int64   = BasicType{TypeTypeID: TypeID_Int64, TypeName: "int64", TypeBitSize: 64}
	Float32 = BasicType{TypeTypeID: TypeID_Float32, TypeName: "float32", TypeBitSize: 32}
	Float64 = BasicType{TypeTypeID: TypeID_Float64, TypeName: "float64", TypeBitSize: 64}
)

var basicTypeMap = map[TypeID]BasicType{
	TypeID_Bool:    Bool,
	TypeID_Uint8:   Uint8,
	TypeID_Uint16:  Uint16,
	TypeID_Uint32:  Uint32,
	TypeID_Uint64:  Uint64,
	TypeID_Int8:    Int8,
	TypeID_Int16:   Int16,
	TypeID_Int32:   Int32,
	TypeID_Int64:   Int64,
	TypeID_Float32: Float32,
	TypeID_Float64: Float64,
}

func GetBasicType(typeID TypeID) *BasicType {
	if t, ok := basicTypeMap[typeID]; ok {
		return &t
	}
	return nil
}

// ==================== StringType ====================

type String struct{}

func (t String) String() string {
	return t.GetTypeName()
}

func (t String) GetTypeID() TypeID {
	return TypeID_String
}

func (t String) GetTypeName() string {
	return "string"
}

func (t String) GetTypeBitSize() int64 {
	return -1
}

// ==================== BytesType ====================

type Bytes struct{}

func (t Bytes) String() string {
	return t.GetTypeName()
}

func (t Bytes) GetTypeID() TypeID {
	return TypeID_Bytes
}

func (t Bytes) GetTypeName() string {
	return "bytes"
}

func (t Bytes) GetTypeBitSize() int64 {
	return -1
}

// ==================== ArrayType ====================

type Array struct {
	ElementType Type
	Length      int64
}

func (t Array) String() string {
	return t.GetTypeName()
}

func (t Array) GetTypeID() TypeID {
	return TypeID_Array
}

func (t Array) GetTypeName() string {
	// Example: int32<10>
	return fmt.Sprintf("%s<%d>", t.ElementType.GetTypeName(), t.Length)
}

func (t Array) GetTypeBitSize() int64 {
	return t.ElementType.GetTypeBitSize() * t.Length
}
