package main

import (
	"fmt"
	"math"
	"sort"
)

/*
给一个包含 n 个整数的数组 S, 找到和与给定整数 target 最接近的三元组，返回这三个数的和。

例1:
输入:[1,2,3,4,5],10
输出:10
解释:
2+3+5=10

例2:
输入:[0,-2,1,-3],2
输出:1
解释:
-2+0+1=1
提示:
双指针

挑战
O(n^2) 时间, O(1) 额外空间。
*/

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res, delta := 0, math.MaxInt64

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s < target:
				l++
				if delta > target-s {
					delta = target - s
					res = s
				}
			case s > target:
				r--
				if delta > s-target {
					delta = s - target
					res = s
				}
			default:
				return s
			}
		}
	}

	return res
}

func main() {
	number_list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 10
	fmt.Printf("number list: [1,2,3,4,5,6,7,8,9], target: 10, return: %d", threeSumClosest(number_list, target))
}
