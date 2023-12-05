package helper

import (
	"fmt"
	"math"
)

// 四舍五入保留小数
func RoundFloat(val float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}