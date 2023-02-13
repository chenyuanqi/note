package main

import "fmt"

func BubbleSort(arr []int) []int {
	var swap bool
	length := len(arr)

	for i := 0; i < length-1; i++ {
		swap = false
		for j := 0; j < length-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swap = true
			}
		}

		if !swap {
			break
		}
	}

	return arr
}

func main() {
	arr := []int{3, 5, 1, 8, 4, 9, 6, 2, 7}
	fmt.Println(BubbleSort(arr))
}
