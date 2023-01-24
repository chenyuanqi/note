package main

import "fmt"

/*
给定一个整数（32位有符号整数），写一个方法判断这个数字是否为4的乘方。

样例 1:
输入：num = 16
输出：True

样例 2:
输入：num = 5
输出：False

挑战：
不使用循环/递归，怎么解决？
*/

func PowerOf4(number int) bool {
	if number < 1 {
		return false
	}

	for number%4 == 0 {
		number /= 4
	}

	return number == 1
}

func main() {
	fmt.Printf("13 is power of 4: %v\n", PowerOf4(13))
	fmt.Printf("16 is power of 4: %v\n", PowerOf4(16))
}
