package main

import "fmt"

func main() {
	fmt.Println(FibRecursion(29))
}

func FibRecursion(num int) int {
	if num <= 0 {
		return 0
	} else if num == 1 || num == 2 {
		return num
	}

	return FibRecursion(num-1) + FibRecursion(num-2)
}
