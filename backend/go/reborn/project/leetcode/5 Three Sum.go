package main

import (
	"fmt"
	"sort"
)

/*
给出一个有n个整数的数组S，在S中找到三个整数a, b, c，找到所有使得a + b + c = 0的三元组。

在三元组(a, b, c)，要求 a <= b <= c 。结果不能包含重复的三元组。
例1:

输入:
[1,2,3,4,5]
输出:
[]
例2:

输入:
[-1,0,1,2,-1,-2]
输出:
[[-1, 0, 1],[-1, -1, 2],[-2, 0, 2]]
提示:
转换成两数之和后使用双指针
*/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s < 0:
				l++
			case s > 0:
				r--
			default:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l, r = next(nums, l, r)
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
	number_list1 := []int{1, 2, 3, 4, 5}
	fmt.Println("a+b+c=0, result: ", threeSum(number_list1))
	number_list2 := []int{-1, 0, 1, 2, -1, -2}
	fmt.Println("a+b+c=0, result: ", threeSum(number_list2))
}
