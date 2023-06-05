package helpers

// 三元运算函数
func ternary(condition bool, valueIfTrue, valueIfFalse interface{}) interface{} {
	if condition {
		return valueIfTrue
	}

	return valueIfFalse
}
