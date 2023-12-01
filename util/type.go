package util

func ToInt64(v interface{}) int64 {
	switch v := v.(type) {
	case int64:
		return v
	case int32:
		return int64(v)
	case int16:
		return int64(v)
	case int8:
		return int64(v)
	case int:
		return int64(v)
	case uint64:
		return int64(v)
	case uint32:
		return int64(v)
	case uint16:
		return int64(v)
	case uint8:
		return int64(v)
	case uint:
		return int64(v)
	default:
		panic("invalid type")
	}
}

func ToInt32(v interface{}) int32 {
	switch v := v.(type) {
	case int64:
		return int32(v)
	case int32:
		return v
	case int16:
		return int32(v)
	case int8:
		return int32(v)
	case int:
		return int32(v)
	case uint64:
		return int32(v)
	case uint32:
		return int32(v)
	case uint16:
		return int32(v)
	case uint8:
		return int32(v)
	case uint:
		return int32(v)
	default:
		panic("invalid type")
	}
}

func ToInt16(v interface{}) int16 {
	switch v := v.(type) {
	case int64:
		return int16(v)
	case int32:
		return int16(v)
	case int16:
		return v
	case int8:
		return int16(v)
	case int:
		return int16(v)
	case uint64:
		return int16(v)
	case uint32:
		return int16(v)
	case uint16:
		return int16(v)
	case uint8:
		return int16(v)
	case uint:
		return int16(v)
	default:
		panic("invalid type")
	}
}

func ToInt8(v interface{}) int8 {
	switch v := v.(type) {
	case int64:
		return int8(v)
	case int32:
		return int8(v)
	case int16:
		return int8(v)
	case int8:
		return v
	case int:
		return int8(v)
	case uint64:
		return int8(v)
	case uint32:
		return int8(v)
	case uint16:
		return int8(v)
	case uint8:
		return int8(v)
	case uint:
		return int8(v)
	default:
		panic("invalid type")
	}
}

func ToInt(v interface{}) int {
	switch v := v.(type) {
	case int64:
		return int(v)
	case int32:
		return int(v)
	case int16:
		return int(v)
	case int8:
		return int(v)
	case int:
		return v
	case uint64:
		return int(v)
	case uint32:
		return int(v)
	case uint16:
		return int(v)
	case uint8:
		return int(v)
	case uint:
		return int(v)
	default:
		panic("invalid type")
	}
}

func ToUint64(v interface{}) uint64 {
	switch v := v.(type) {
	case int64:
		return uint64(v)
	case int32:
		return uint64(v)
	case int16:
		return uint64(v)
	case int8:
		return uint64(v)
	case int:
		return uint64(v)
	case uint64:
		return v
	case uint32:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint:
		return uint64(v)
	default:
		panic("invalid type")
	}
}

func ToUint32(v interface{}) uint32 {
	switch v := v.(type) {
	case int64:
		return uint32(v)
	case int32:
		return uint32(v)
	case int16:
		return uint32(v)
	case int8:
		return uint32(v)
	case int:
		return uint32(v)
	case uint64:
		return uint32(v)
	case uint32:
		return v
	case uint16:
		return uint32(v)
	case uint8:
		return uint32(v)
	case uint:
		return uint32(v)
	default:
		panic("invalid type")
	}
}

func ToUint16(v interface{}) uint16 {
	switch v := v.(type) {
	case int64:
		return uint16(v)
	case int32:
		return uint16(v)
	case int16:
		return uint16(v)
	case int8:
		return uint16(v)
	case int:
		return uint16(v)
	case uint64:
		return uint16(v)
	case uint32:
		return uint16(v)
	case uint16:
		return v
	case uint8:
		return uint16(v)
	case uint:
		return uint16(v)
	default:
		panic("invalid type")
	}
}

func ToUint8(v interface{}) uint8 {
	switch v := v.(type) {
	case int64:
		return uint8(v)
	case int32:
		return uint8(v)
	case int16:
		return uint8(v)
	case int8:
		return uint8(v)
	case int:
		return uint8(v)
	case uint64:
		return uint8(v)
	case uint32:
		return uint8(v)
	case uint16:
		return uint8(v)
	case uint8:
		return v
	case uint:
		return uint8(v)
	default:
		panic("invalid type")
	}
}

func ToUint(v interface{}) uint {
	switch v := v.(type) {
	case int64:
		return uint(v)
	case int32:
		return uint(v)
	case int16:
		return uint(v)
	case int8:
		return uint(v)
	case int:
		return uint(v)
	case uint64:
		return uint(v)
	case uint32:
		return uint(v)
	case uint16:
		return uint(v)
	case uint8:
		return uint(v)
	case uint:
		return v
	default:
		panic("invalid type")
	}
}

func ToFloat64(v interface{}) float64 {
	switch v := v.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	case int64:
		return float64(v)
	case int32:
		return float64(v)
	case int16:
		return float64(v)
	case int8:
		return float64(v)
	case int:
		return float64(v)
	case uint64:
		return float64(v)
	case uint32:
		return float64(v)
	case uint16:
		return float64(v)
	case uint8:
		return float64(v)
	case uint:
		return float64(v)
	default:
		panic("invalid type")
	}
}

func ToFloat32(v interface{}) float32 {
	switch v := v.(type) {
	case float64:
		return float32(v)
	case float32:
		return v
	case int64:
		return float32(v)
	case int32:
		return float32(v)
	case int16:
		return float32(v)
	case int8:
		return float32(v)
	case int:
		return float32(v)
	case uint64:
		return float32(v)
	case uint32:
		return float32(v)
	case uint16:
		return float32(v)
	case uint8:
		return float32(v)
	case uint:
		return float32(v)
	default:
		panic("invalid type")
	}
}

func ToBool(v interface{}) bool {
	switch v := v.(type) {
	case bool:
		return v
	case int64:
		return v != 0
	case int32:
		return v != 0
	case int16:
		return v != 0
	case int8:
		return v != 0
	case int:
		return v != 0
	case uint64:
		return v != 0
	case uint32:
		return v != 0
	case uint16:
		return v != 0
	case uint8:
		return v != 0
	case uint:
		return v != 0
	default:
		panic("invalid type")
	}
}

type Bool interface {
	~bool
}

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32
}

type Float interface {
	~float32 | ~float64
}

type IntNumber interface {
	Int | Uint
}

type Number interface {
	IntNumber | Float
}

// func To[T, U Number](v U) T {
//     return T(v)
// }

// func ToInt64[T Number](v T) int64 {
//     return int64(v)
// }

// func ToInt32[T Number](v T) int32 {
//     return int32(v)
// }

// func ToInt16[T Number](v T) int16 {
//     return int16(v)
// }

// func ToInt8[T Number](v T) int8 {
//     return int8(v)
// }

// func ToInt[T Number](v T) int {
//     return int(v)
// }

// func ToUint64[T Number](v T) uint64 {
//     return uint64(v)
// }

// func ToUint32[T Number](v T) uint32 {
//     return uint32(v)
// }

// func ToUint16[T Number](v T) uint16 {
//     return uint16(v)
// }

// func ToUint8[T Number](v T) uint8 {
//     return uint8(v)
// }

// func ToUint[T Number](v T) uint {
//     return uint(v)
// }

// func ToFloat64[T Number](v T) float64 {
//     return float64(v)
// }

// func ToFloat32[T Number](v T) float32 {
//     return float32(v)
// }

// func ToBool[T Number](v T) bool {
//     return v != 0
// }
