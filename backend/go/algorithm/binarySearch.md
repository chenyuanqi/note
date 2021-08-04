
### Go 二分查找
对于基于数字索引的数组 / 切片元素查找，我们可能第一反应都是遍历这个数组 / 切片，直到给定元素值和待查找的值相等时，返回索引值并退出，否则一直遍历到最后一个元素，如果还是没有找到则返回 -1。  

所谓二分查找，针对的是一个有序的数据集合（这点很重要），查找思想有点类似分治思想 —— 每次都通过跟区间的中间元素对比，将待查找的区间缩小为之前的一半，直到找到要查找的元素，或者区间被缩小为 0。  
```golang
func BinarySearch(nums []int, num int) int {
	len := len(nums)
	if len == 0 {
		return -1
	}

	return binarySearchInternal(nums, num, 0, len - 1)
}

func binarySearchInternal(nums []int, num int, low int, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2
	if num > nums[mid] {
		return binarySearchInternal(nums, num, mid + 1, high)
	} else if num < nums[mid] {
		return binarySearchInternal(nums, num, low, mid - 1)
	}
	
	return mid
}

// 测试
arr := []int{1,5,7,9,10}
fmt.Println(BinarySearch(arr, 5))
```
