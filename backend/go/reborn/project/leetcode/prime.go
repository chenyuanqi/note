package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	sqrtNum := int(math.Sqrt(float64(num)))
	for i := 2; i < sqrtNum; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	num := 7
	fmt.Printf("%d is prime: %t\n", num, isPrime(num))
}
