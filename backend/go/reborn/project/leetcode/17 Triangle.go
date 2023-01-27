package main

import "fmt"

/*
给定一个数字三角形，找到从顶部到底部的最小路径和。每一步可以移动到下面一行的相邻数字上。

样例
样例 1
输入下列数字三角形：
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
输出： 11
解释： 从顶到底部的最小路径和为11 ( 2 + 3 + 5 + 1 = 11)。

样例 2
输入下列数字三角形：
[
     [2],
    [3,2],
   [6,5,7],
  [4,4,8,1]
]
输出： 12
解释： 从顶到底部的最小路径和为12 ( 2 + 2 + 7 + 1 = 12)。

挑战
额外空间复杂度O(n)，其中n是数字三角形的总行数。
*/

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			switch {
			case j == 0:
				triangle[i][0] += triangle[i-1][0]
			case j == i:
				triangle[i][i] += triangle[i-1][i-1]
			case triangle[i-1][j-1] < triangle[i-1][j]:
				triangle[i][j] += triangle[i-1][j-1]
			default:
				triangle[i][j] += triangle[i-1][j]
			}
		}
	}

	min := triangle[n-1][0]
	for j := 1; j < n; j++ {
		if min > triangle[n-1][j] {
			min = triangle[n-1][j]
		}
	}

	return min
}

func main() {
	var triangle [][]int
	triangle = append(triangle, []int{2})
	triangle = append(triangle, []int{3, 4})
	triangle = append(triangle, []int{6, 5, 7})
	triangle = append(triangle, []int{4, 1, 8, 3})

	fmt.Printf("input: %v, output: %v\n", triangle, minimumTotal(triangle))
}
