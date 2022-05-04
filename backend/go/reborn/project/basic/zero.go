package main

import "fmt"

func main() {
	var s string
	var i int
	var b bool
	var f float64
	var st struct{}
	var mi map[string]int
	var sl []string
	var ia interface{}
	var fn func()
	var ch chan string
	fmt.Println("string 的零值为", s)
	fmt.Println("int 的零值为", i)
	fmt.Println("bool 的零值为", b)
	fmt.Println("float64 的零值为", f)
	fmt.Println("struct 的零值为", st)
	fmt.Println("map 的零值为", mi)
	fmt.Println("slice 的零值为", sl)
	fmt.Println("interface 的零值为", ia)
	fmt.Println("func 的零值为", fn)
	fmt.Println("chan 的零值为", ch)
}
