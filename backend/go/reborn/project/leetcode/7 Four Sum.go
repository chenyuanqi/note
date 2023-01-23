package main

/*
给一个包含n个数的整数数组S，在S中找到所有使得和为给定整数target的四元组(a, b, c, d)。

四元组(a, b, c, d)中，需要满足a <= b <= c <= d

答案中不可以包含重复的四元组。
例1:
输入:[2,7,11,15],3
输出:[]

例2:
输入:[1,0,-1,0,-2,2],0
输出:
[[-1, 0, 0, 1]
,[-2, -1, 1, 2]
,[-2, 0, 0, 2]]
*/

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			l, r := j+1, len(nums)-1
			for l < r {
				s := nums[i] + nums[j] + nums[l] + nums[r]
				switch {
				case s < target:
					l++
				case s > target:
					r--
				default:
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l, r = next(nums, l, r)
				}
			}
		}

	}
	return res
}

func next(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}

	return l, r
}

func main() {
	number_list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 10
	fmt.Printf("number list: [1,2,3,4,5,6,7,8,9], target: 10, return: %d", fourSum(number_list, target))
}
