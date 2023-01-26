package main

import "fmt"

/*
给定一个包含红，白，蓝且长度为 n 的数组，将数组元素进行分类使相同颜色的元素相邻，并按照红、白、蓝的顺序进行排序。

使用整数 0，1 和 2 分别代表 红，白，蓝 。

不能使用代码库中的排序函数来解决这个问题。
排序需要在原数组中进行。
样例
输入 : [1, 0, 1, 2]
输出 : [0, 1, 1, 2]

挑战
一个相当直接的解决方案是使用计数排序扫描2遍的算法。首先，迭代数组计算 0,1,2 出现的次数，然后依次用 0,1,2 出现的次数去覆盖数组。
你否能想出一个仅使用常数级额外空间复杂度且只扫描遍历一遍数组的算法？
*/

func SortColors(a []int) []int {
	i, j, k := 0, 0, len(a)-1

	for j <= k {
		switch a[j] {
		case 0:
			a[i], a[j] = a[j], a[i]
			i++
			j++
		case 1:
			j++
		case 2:
			a[j], a[k] = a[k], a[j]
			k--
		}
	}

	return a
}

func main() {
	number_list := []int{1, 0, 1, 2}
	fmt.Printf("[1, 0, 1, 2], result: %+v\n", SortColors(number_list))
}
