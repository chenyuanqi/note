package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

func conversion(number int64, binary int64) int64 {
	negative := number < 0
	number = int64(math.Abs(float64(number)))

	var str strings.Builder
	for number != 0 {
		str.WriteString(strconv.Itoa(int(number % binary)))
		number /= binary
	}

	if negative {
		str.WriteString("-")
	}

	// fmt.Printf("%#v", str)
	result, _ := strconv.ParseInt(reverse(str.String()), 10, 64)
	return int64(result)
}

func main() {
	fmt.Println(conversion(-100, 7))
	fmt.Println(conversion(50, 2))
}
