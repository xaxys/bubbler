package definition

import (
	"fmt"
)

// ==================== Literal ====================

type LiteralKindID int

const (
	LiteralKindID_Bool LiteralKindID = iota
	LiteralKindID_Int
	LiteralKindID_Float
	LiteralKindID_String
)

func (l LiteralKindID) IsBool() bool {
	return l == LiteralKindID_Bool
}

func (l LiteralKindID) IsInt() bool {
	return l == LiteralKindID_Int
}

func (l LiteralKindID) IsFloat() bool {
	return l == LiteralKindID_Float
}

func (l LiteralKindID) IsString() bool {
	return l == LiteralKindID_String
}

var literalKindToString = map[LiteralKindID]string{
	LiteralKindID_Bool:   "bool",
	LiteralKindID_Int:    "int",
	LiteralKindID_Float:  "float",
	LiteralKindID_String: "string",
}

func (l LiteralKindID) String() string {
	if str, ok := literalKindToString[l]; ok {
		return str
	}
	return "unknown"
}

// var literalKindToTypeID = map[LiteralKindID]TypeID{
// 	LiteralKindID_Bool:   TypeID_Bool,
// 	LiteralKindID_Int:    TypeID_Int64,
// 	LiteralKindID_Float:  TypeID_Float64,
// 	LiteralKindID_String: TypeID_String,
// }

// func (l LiteralKindID) ToTypeID() TypeID {
// 	if typeID, ok := literalKindToTypeID[l]; ok {
// 		return typeID
// 	}
// 	return TypeID_Int32
// }

// ==================== Literal ====================

type Literal interface {
	GetLiteralKind() LiteralKindID
	GetLiteralValue() interface{}
	GetLiteralValueIn(ty TypeID) interface{}
	GetMinimalTypeID() TypeID
}

// ==================== BoolLiteral ====================

// ensure that BoolLiteral implements Literal
var _ Literal = (*BoolLiteral)(nil)

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

func (l BoolLiteral) GetLiteralValueIn(ty TypeID) interface{} {
	switch ty {
	case TypeID_Bool:
		return l.BoolValue
	case TypeID_String:
		return fmt.Sprint(l.BoolValue)
	case TypeID_Bytes:
		if l.BoolValue {
			return []byte{1}
		} else {
			return []byte{0}
		}
	}
	return nil
}

func (l BoolLiteral) GetMinimalTypeID() TypeID {
	return TypeID_Bool
}

// ==================== IntLiteral ====================

// ensure that IntLiteral implements Literal
var _ Literal = (*IntLiteral)(nil)

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

func (l IntLiteral) GetLiteralValueIn(ty TypeID) interface{} {
	switch ty {
	case TypeID_Bool:
		return l.IntValue != 0
	case TypeID_Uint8:
		return uint8(l.IntValue)
	case TypeID_Uint16:
		return uint16(l.IntValue)
	case TypeID_Uint32:
		return uint32(l.IntValue)
	case TypeID_Uint64:
		return uint64(l.IntValue)
	case TypeID_Int8:
		return int8(l.IntValue)
	case TypeID_Int16:
		return int16(l.IntValue)
	case TypeID_Int32:
		return int32(l.IntValue)
	case TypeID_Int64:
		return l.IntValue
	case TypeID_Float32:
		return float32(l.IntValue)
	case TypeID_Float64:
		return float64(l.IntValue)
	case TypeID_String:
		return fmt.Sprint(l.IntValue)
	case TypeID_Bytes:
		return []byte(fmt.Sprint(l.IntValue))
	}
	return nil
}

func (l IntLiteral) GetMinimalTypeID() TypeID {
	if l.IntValue >= 0 {
		if l.IntValue <= int64(^uint8(0)) {
			return TypeID_Uint8
		}
		if l.IntValue <= int64(^uint16(0)) {
			return TypeID_Uint16
		}
		if l.IntValue <= int64(^uint32(0)) {
			return TypeID_Uint32
		}
		return TypeID_Uint64
	} else {
		if l.IntValue >= int64(^int8(0)) {
			return TypeID_Int8
		}
		if l.IntValue >= int64(^int16(0)) {
			return TypeID_Int16
		}
		if l.IntValue >= int64(^int32(0)) {
			return TypeID_Int32
		}
		return TypeID_Int64
	}
}

// ==================== FloatLiteral ====================

// ensure that FloatLiteral implements Literal
var _ Literal = (*FloatLiteral)(nil)

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

func (l FloatLiteral) GetLiteralValueIn(ty TypeID) interface{} {
	switch ty {
	case TypeID_Bool:
		return l.FloatValue != 0
	case TypeID_Uint8:
		return uint8(l.FloatValue)
	case TypeID_Uint16:
		return uint16(l.FloatValue)
	case TypeID_Uint32:
		return uint32(l.FloatValue)
	case TypeID_Uint64:
		return uint64(l.FloatValue)
	case TypeID_Int8:
		return int8(l.FloatValue)
	case TypeID_Int16:
		return int16(l.FloatValue)
	case TypeID_Int32:
		return int32(l.FloatValue)
	case TypeID_Int64:
		return int64(l.FloatValue)
	case TypeID_Float32:
		return float32(l.FloatValue)
	case TypeID_Float64:
		return l.FloatValue
	case TypeID_String:
		return fmt.Sprint(l.FloatValue)
	case TypeID_Bytes:
		return []byte(fmt.Sprint(l.FloatValue))
	}
	return nil
}

func (l FloatLiteral) GetMinimalTypeID() TypeID {
	return TypeID_Float64
}

// ==================== StringLiteral ====================

// ensure that StringLiteral implements Literal
var _ Literal = (*StringLiteral)(nil)

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

func (l StringLiteral) GetLiteralValueIn(ty TypeID) interface{} {
	switch ty {
	case TypeID_Bool:
		return l.StringValue != ""
	case TypeID_String:
		return l.StringValue
	case TypeID_Bytes:
		return []byte(l.StringValue)
	}
	return nil
}

func (l StringLiteral) GetMinimalTypeID() TypeID {
	return TypeID_String
}
