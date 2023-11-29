package util

func LowBit[T IntNumber](v T) T {
	return v & -v
}

func HighBit[T IntNumber](v T) T {
	v = v - 1
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v |= v >> 32
	return v + 1
}
