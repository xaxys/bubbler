package definition

import "fmt"

// ==================== Literal ====================

type LiteralKindID int

const (
	LiteralKindID_Bool LiteralKindID = iota
	LiteralKindID_Int
	LiteralKindID_Float
	LiteralKindID_String
)

func (l LiteralKindID) String() string {
	switch l {
	case LiteralKindID_Bool:
		return "bool"
	case LiteralKindID_Int:
		return "int"
	case LiteralKindID_Float:
		return "float"
	case LiteralKindID_String:
		return "string"
	}
	return "unknown"
}

func (l LiteralKindID) ToTypeID() TypeID {
	switch l {
	case LiteralKindID_Bool:
		return TypeID_Bool
	case LiteralKindID_Int:
		return TypeID_Int64
	case LiteralKindID_Float:
		return TypeID_Float64
	case LiteralKindID_String:
		return TypeID_String
	}
	return TypeID_Int32
}

type Literal interface {
	GetLiteralKind() LiteralKindID
	GetLiteralValue() interface{}
	GetLiteralValueIn(ty *BasicType) interface{}
}

type BoolLiteral struct {
	BasePosition
	BoolValue bool
}

func (l BoolLiteral) String() string {
	return fmt.Sprint(l.BoolValue)
}

func (l BoolLiteral) GetLiteralKind() LiteralKindID {
	return LiteralKindID_Bool
}

func (l BoolLiteral) GetLiteralValue() interface{} {
	return l.BoolValue
}

func (l BoolLiteral) GetLiteralValueIn(ty *BasicType) interface{} {
	if ty.TypeTypeID == TypeID_Bool {
		return l.BoolValue
	}
	return nil
}

type IntLiteral struct {
	BasePosition
	IntValue int64
}

func (l IntLiteral) String() string {
	return fmt.Sprint(l.IntValue)
}

func (l IntLiteral) GetLiteralKind() LiteralKindID {
	return LiteralKindID_Int
}

func (l IntLiteral) GetLiteralValue() interface{} {
	return l.IntValue
}

func (l IntLiteral) GetLiteralValueIn(ty *BasicType) interface{} {
	switch ty.TypeTypeID {
	case TypeID_Uint8:
		return uint8(l.IntValue)
	case TypeID_Int8:
		return int8(l.IntValue)
	case TypeID_Uint16:
		return uint16(l.IntValue)
	case TypeID_Int16:
		return int16(l.IntValue)
	case TypeID_Uint32:
		return uint32(l.IntValue)
	case TypeID_Int32:
		return int32(l.IntValue)
	case TypeID_Uint64:
		return uint64(l.IntValue)
	case TypeID_Int64:
		return l.IntValue
	}
	return nil
}

type FloatLiteral struct {
	BasePosition
	FloatValue float64
}

func (l FloatLiteral) String() string {
	return fmt.Sprint(l.FloatValue)
}

func (l FloatLiteral) GetLiteralKind() LiteralKindID {
	return LiteralKindID_Float
}

func (l FloatLiteral) GetLiteralValue() interface{} {
	return l.FloatValue
}

func (l FloatLiteral) GetLiteralValueIn(ty *BasicType) interface{} {
	switch ty.TypeTypeID {
	case TypeID_Float32:
		return float32(l.FloatValue)
	case TypeID_Float64:
		return l.FloatValue
	}
	return nil
}

type StringLiteral struct {
	BasePosition
	StringValue string
}

func (l StringLiteral) String() string {
	return fmt.Sprintf(`"%s"`, l.StringValue)
}

func (l StringLiteral) GetLiteralKind() LiteralKindID {
	return LiteralKindID_String
}

func (l StringLiteral) GetLiteralValue() interface{} {
	return l.StringValue
}

func (l StringLiteral) GetLiteralValueIn(ty *BasicType) interface{} {
	if ty.TypeTypeID == TypeID_String {
		return l.StringValue
	}
	return nil
}
