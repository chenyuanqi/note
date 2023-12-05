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

// 向上取整保留小数
func CeilFloat(val float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Ceil(val*ratio) / ratio
}

// 向下取整保留小数
func FloorFloat(val float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Floor(val*ratio) / ratio
}
