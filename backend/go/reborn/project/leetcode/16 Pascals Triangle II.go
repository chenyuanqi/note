package main

import "fmt"

/*
给定非负索引k，其中k≤33，返回杨辉三角形的第k个索引行。

在杨辉三角中，每个数是它左上方和右上方的数的和。

样例
样例1
输入: 3
输出: [1,3,3,1]

样例2
输入: 4
输出: [1,4,6,4,1]
挑战
你可以优化你的算法到空间复杂度为O(k)吗？
*/
func PascalsTriangle(rowIndex int) []int {
	res := make([]int, 1, rowIndex+1)
	res[0] = 1
	if rowIndex == 0 {
		return res
	}

	for i := 0; i < rowIndex; i++ {
		res = append(res, 1)
		for j := len(res) - 2; j > 0; j-- {
			res[j] += res[j-1]
		}
	}

	return res
}

func main() {
	fmt.Printf("pascal triangle 5: %+v \n", PascalsTriangle(5))
	fmt.Printf("pascal triangle 3: %+v \n", PascalsTriangle(3))
}
