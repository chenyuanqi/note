package main

import "fmt"

/*
给定一个整数，写一个函数来确定它是否是2的幂。

示例:
输入:
2
输出:
true

挑战:
使用位运算，O(1) 时间和空间复杂度
*/
func PowerOf2(number int) bool {
	if number < 1 {
		return false
	}

	for number > 1 {
		if number%2 == 1 {
			return false
		}

		number /= 2
	}

	return true
}

func main() {
	fmt.Printf("13 is power of 2: %v\n", PowerOf2(13))
	fmt.Printf("1024 is power of 2: %v\n", PowerOf2(1024))
}
