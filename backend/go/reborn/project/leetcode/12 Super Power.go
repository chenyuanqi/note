package main

import "fmt"

/*
你的任务是计算 a^b mod 1337，其中 a 是一个正整数，b 是一个超级大的正整数，以数组的形式给出。

示例 1:
输入：
a = 2
b = [2]
输出：
4

示例 2:
输入：
a = 2
b = [1,0]
输出：
1024

提示:
用数组标识指数
得到对1337求模的结果
高效地进行幂运算
*/

// (a*b) % k = ((a%k)*(b%k)) % k
func superPow(a int, b []int) int {
	base := 1337

	// return a^k mod base
	powmod := func(a, k int) int {
		a %= base
		res := 1

		for i := 0; i < k; i++ {
			res = (res * a) % base
		}
		return res
	}

	n := len(b)
	if n == 0 {
		return 1
	}

	lastB := b[n-1]
	newB := b[:n-1]

	return (powmod(superPow(a, newB), 10) * powmod(a, lastB)) % base
}

func main() {
	fmt.Printf("a=2 b=[2], return: %v\n", superPow(2, []int{2}))
}
