package main

import "fmt"

// BubbleSort 冒泡排序
// 参数 arr 是一个整数切片，返回值是一个排好序的整数切片
func BubbleSort(arr []int) []int {
	// swap 是一个布尔值，用于记录在本轮排序中是否发生了交换
	var swap bool
	// length 记录了切片 arr 的长度
	length := len(arr)

	// i 从 0 开始遍历，到 length-1 结束
	// 这个循环用于控制排序的轮数
	for i := 0; i < length-1; i++ {
		// 初始化 swap 为 false
		swap = false
		// j 从 0 开始遍历，到 length-1 结束
		// 这个循环用于比较相邻的两个数并交换
		for j := 0; j < length-i-1; j++ {
			// 如果 arr[j] 大于 arr[j+1]，则交换两个数的位置
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				// 设置 swap 为 true，表示发生了交换
				swap = true
			}
		}

		// 如果在本轮排序中没有发生交换，则说明已经排好序了，直接 break
		if !swap {
			break
		}
	}

	// 返回排好序的整数切片
	return arr
}

func main() {
	// 定义一个整数切片 arr
	arr := []int{3, 5, 1, 8, 4, 9, 6, 2, 7}
	// 调用 BubbleSort 函数，并打印结果
	fmt.Println(BubbleSort(arr))
}
