package main

/*
将两个整数相除，要求不使用乘法、除法和 mod 运算符。

如果溢出(超出32位有符号整型表示范围)，返回 2147483647 。

整数除法应截断为零。

样例 1:
输入: 被除数 = 0, 除数 = 1
输出: 0

样例 2:
输入: 被除数 = 100, 除数 = 9
输出: 11

提示:
被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31, 2^31 − 1]。本题中，如果除法结果溢出，则返回 2^31 − 1。
*/
import (
	"fmt"
	"math"
)

func divide(m, n int) int {
	if n == 0 {
		return math.MaxInt32
	}

	signM, absM := analysis(m)
	signN, absN := analysis(n)

	res, _ := d(absM, absN, 1)

	if signM != signN {
		res = res - res - res
	}

	if res < math.MinInt32 || res > math.MaxInt32 {
		return math.MaxInt32
	}

	return res
}

func analysis(num int) (sign, abs int) {
	sign = 1
	abs = num
	if num < 0 {
		sign = -1
		abs = num - num - num
	}

	return
}

func d(m, n, count int) (res, remainder int) {
	switch {
	case m < n:
		return 0, m
	case n <= m && m < n+n:
		return count, m - n
	default:
		res, remainder = d(m, n+n, count+count)
		if remainder >= n {
			return res + count, remainder - n
		}

		return
	}
}

func main() {
	fmt.Printf("divide two int: 100, 9, result: %d\n", divide(100, 9))
}
