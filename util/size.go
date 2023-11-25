package util

import "fmt"

func ToSizeString(bitSize int64) string {
	bytes := bitSize / 8
	bits := bitSize % 8
	if bytes == 0 && bits == 0 {
		return "0"
	}
	if bytes == 0 {
		return fmt.Sprintf("#%d", bits)
	}
	if bits == 0 {
		return fmt.Sprintf("%d", bytes)
	}
	return fmt.Sprintf("%d#%d", bytes, bits)
}
