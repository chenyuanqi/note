package main

import "fmt"

// 快速排序的基本思想是通过选定一个基准点，将待排序的数组分为两个子数组，左边的子数组中的元素都小于等于基准点，右边的子数组中的元素都大于基准点
// 然后对左右两个子数组分别递归进行快速排序，最终完成整个排序过程。在上面的代码中，quickSort 函数就是实现了这个过程，其中的 partition 函数用于实现数组分割的过程
func main() {
	// 待排序的数组
	arr := []int{5, 3, 7, 8, 1, 4, 6, 2}

	// 调用快速排序函数进行排序
	quickSort(arr, 0, len(arr)-1)

	// 输出排序后的结果
	fmt.Println(arr)
}

// 快速排序函数
func quickSort(arr []int, left, right int) {
	if left < right {
		// 将数组分为两个子数组，并返回分割位置
		pivot := partition(arr, left, right)

		// 递归对左子数组进行排序
		quickSort(arr, left, pivot-1)

		// 递归对右子数组进行排序
		quickSort(arr, pivot+1, right)
	}
}

// 分割函数，返回分割位置
func partition(arr []int, left, right int) int {
	// 取最右边的元素作为基准点
	pivot := arr[right]

	// i 指向左边第一个大于基准点的元素
	i := left - 1

	// j 从左到右遍历数组，将小于等于基准点的元素移到左边，大于基准点的元素移到右边
	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// 将基准点移到正确的位置
	arr[i+1], arr[right] = arr[right], arr[i+1]

	// 返回基准点位置
	return i + 1
}
