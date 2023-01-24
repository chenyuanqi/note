package main

import "fmt"

/*
给定一个整数，判断它是否为3的幂。

样例1

输入: n = 0
输出: False
样例2

输入: n = 9
输出: True
示例 3:

输入: 100
输出: false
挑战：
能否不用循环/递归实现？
*/
func PowerOf3(number int) bool {
	if number < 1 {
		return false
	}

	for number%3 == 0 {
		number /= 3
	}

	return number == 1
}

func main() {
	fmt.Printf("13 is power of 3: %v\n", PowerOf3(13))
	fmt.Printf("81 is power of 3: %v\n", PowerOf3(81))
}
