package main

import "fmt"

/*
给一个整数数组，找到两个数使得他们的和等于一个给定的数 target。

你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。
示例:
numbers = [1, 2, 3, 4, 5]
target = 6

返回 [1, 3]
提示:
哈希查找的时间复杂度为O(1) ，或许你该考虑使用哈希表

挑战:
O(n) 空间复杂度，O(nlogn) 时间复杂度，
O(n) 空间复杂度，O(n) 时间复杂度，
*/

func sum(numbers []int, target int) []int {
	/*
		number_length := len(numbers)
		for i := 0; i < number_length - 1; i++ {
			for j := 1; j < number_length; j++ {
				if  numbers[i] + numbers[j] == target {
					return []int{numbers[i], numbers[j]}
				}
			}
		}
	*/
	hash := make(map[int]int, len(numbers))
	for i, n := range numbers {
		if j, ok := hash[target-n]; ok {
			return []int{numbers[j], n}
		}

		hash[n] = i
	}

	return nil
}

func main() {
	number_list := []int{1, 2, 3, 4, 5}
	target := 6
	fmt.Printf("numbers = []int{1, 2, 3, 4, 5}, target = 6, return: %d\n", sum(number_list, target))
}
