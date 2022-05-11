package types

import (
	"blog/pkg/logger"

	"strconv"
)

// Int64ToString 将 int64 转换为 string
func Uint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}

// StringToUint64 将字符串转换为 uint64
func StringToUint64(str string) uint64 {
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		logger.LogError(err)
	}
	return i
}
