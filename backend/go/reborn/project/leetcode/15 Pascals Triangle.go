package main

import "fmt"

/*
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:
样例 1:
输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]

样例 2:
输入: 3
输出:
[
     [1],
    [1,1],
   [1,2,1]
]
*/
func PascalsTriangle(n int) [][]int {
	var result [][]int
	var cur_list, last_list []int
	for i := 0; i < n; i++ {
		cur_list = []int{}
		if i > 0 {
			cur_list = append(cur_list, 1)
			for j := 1; j < i; j++ {
				cur_list = append(cur_list, last_list[j-1]+last_list[j])
			}
		}

		cur_list = append(cur_list, 1)
		result = append(result, cur_list)
		last_list = cur_list
	}

	return result
}

func main() {
	fmt.Printf("pascal triangle 5: %+v \n", PascalsTriangle(5))
	fmt.Printf("pascal triangle 3: %+v \n", PascalsTriangle(3))
}
