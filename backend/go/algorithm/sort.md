
### Go 冒泡排序
冒泡排序只会操作相邻的两个数据。每次冒泡操作都会对相邻的两个元素进行比较，看是否满足大小关系要求，如果不满足就让它俩互换。一次冒泡会让至少一个元素移动到它应该在的位置，重复 n 次，就完成了 n 个数据的排序工作。
```golang
func BubbleSort(arr []int) []int {
	len := len(arr)
	if len == 0 {
		return []int{}
	}

	for i := 0; i < len - 1; i++ {
		flag := false
		for j := 0; j < len - 1; j++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
				flag = true
			}
		}

		if !flag {
			break
		}
	}

	return arr
}

// 测试
arr := []int{4, 5, 6, 7, 8, 3, 2, 1}
fmt.Println(BubbleSort(arr)) // [1 2 3 4 5 6 7 8]
```
