package main

import "fmt"

func noRepeatLongestStr(str string) int {
	var start, maxLength int
	var lastOccurred = make(map[rune]int)
	for i, c := range []rune(str) {
		if lastI, ok := lastOccurred[c]; ok && lastI >= start {
			start = i
		}
		if (i + 1 - start) > maxLength {
			maxLength = i + 1 - start
		}

		lastOccurred[c] = i
	}

	return maxLength
}

func main() {
	fmt.Println(noRepeatLongestStr("我是最长子串串"))
	fmt.Println(noRepeatLongestStr("test"))
}
