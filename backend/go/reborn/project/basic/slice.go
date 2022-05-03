package main

import (
	"fmt"
	"reflect"
)

func main() {
	// create
	var s []int // a nil slice
	fmt.Println(s)
	s1 := []string{"foo", "bar"}
	fmt.Println(s1)
	s2 := make([]int, 2)                            // same as []int{0, 0}
	s3 := make([]int, 2, 4)                         // same as new([4]int)[:2]
	fmt.Println(len(s2), cap(s2), len(s3), cap(s3)) // 2 4

	// slicing
	a := [...]int{0, 1, 2, 3} // an array
	s4 := a[1:3]              // s == []int{1, 2}        cap(s) == 3
	s4 = a[:2]                // s == []int{0, 1}        cap(s) == 4
	s4 = a[2:]                // s == []int{2, 3}        cap(s) == 2
	s4 = a[:]                 // s == []int{0, 1, 2, 3}  cap(s) == 4
	fmt.Println(s4)

	// iteration
	s5 := []string{"Foo", "Bar"}
	for i, v := range s5 {
		fmt.Println(i, v)
	}

	// compare - array
	s6 := [2]int{1, 2}
	s7 := [2]int{1, 3}
	fmt.Println(s6 == s7) // false
	// compare - slice
	var s8 []int = nil
	var s9 []int = make([]int, 0)
	fmt.Println(reflect.DeepEqual(s8, s9)) // false
}
