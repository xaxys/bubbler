package util

import "fmt"

func ToSizeStringDynamic(bitSizeStart int64, dynamicStart bool, bitSize int64) string {
	bitSizeStr := ""
	if bitSize == -1 {
		bitSizeStr = fmt.Sprintf("%s+dynamic", ToSizeString(bitSizeStart))
	} else {
		bitSizeStr = ToSizeString(bitSize + bitSizeStart)
	}
	if dynamicStart {
		return fmt.Sprintf("dynamic+%s", bitSizeStr)
	}
	return bitSizeStr
}

func ToSizeString(bitSize int64) string {
	if bitSize == -1 {
		return "dynamic"
	}
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
