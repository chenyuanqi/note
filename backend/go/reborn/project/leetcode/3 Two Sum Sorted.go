package main

import "fmt"

/*
给定一个已经 按升序排列 的数组，找到两个数使他们加起来的和等于特定数。
函数应该返回这两个数的下标，index1必须小于index2。

你可以假设每个输入刚好只有一个答案
返回的下标值（index1 和 index2）不是从零开始的。
例1:
输入: nums = [2, 7, 11, 15]
target = 9
输出: [1, 2]
例2:

输入: nums = [2,3]
target = 5
输出: [1, 2]
提示:
输入数据已经排序，故使用双指针能大幅提高遍历效率（左指针，右指针向不同方向移动）
*/

func sum(numbers []int, target int) []int {
	hash := make(map[int]int, len(numbers))
	for i, n := range numbers {
		if j, ok := hash[target-n]; ok {
			return []int{j, i + 1}
		}

		hash[n] = i + 1
	}

	return nil
}

func main() {
	number_list := []int{1, 2, 3, 4, 5}
	target := 6
	fmt.Printf("numbers = []int{1, 2, 3, 4, 5}, target = 6, return: %d\n", sum(number_list, target))
}
