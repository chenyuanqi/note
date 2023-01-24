package main

import "fmt"

/*
实现 pow(x, n). (n是一个整数)
不用担心精度，当答案和标准输出差绝对值小于 1e-3 时都算正确

样例 1:
输入: x = 9.88023, n = 3
输出: 964.498

样例 2:
输入: x = 2.1, n = 3
输出: 9.261

样例 3:
输入: x = 1, n = 0
输出: 1

挑战
时间复杂度O(logn)
*/
func PowerOfN(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / pow(x, -n)
	}

	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	res := pow(x, n>>1)
	if n&1 == 0 {
		return res * res
	}

	return res * res * x
}

func main() {
	fmt.Printf("x = 1, n = 0: %v\n", PowerOfN(1, 0))
}
