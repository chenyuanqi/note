package main

import (
	"fmt"
	"strconv"
)

func dump(params ...interface{}) {
	fmt.Println(params...)
}

func main() {
	x1 := 5.1
	n := int(x1) // convert float to int
	dump(n)

	a := uint16(0x10fe) // 0001 0000 1111 1110 (4350)
	b := int8(a)        //           1111 1110 (truncated to -2)
	c := uint16(b)      // 1111 1111 1111 1110 (sign extended to 0xfffe, 65534)
	dump(a, b, c)

	var x2 float64 = 1.9
	m := int64(x2) // 1
	dump(m)
	m = int64(-x2) // -1
	dump(m)

	m = 1234567890
	y := float32(m) // 1.234568e+09
	dump(y)

	// Integer to string
	dump(string(97)) // "a"
	dump(string(-1)) // "\ufffd" == "\xef\xbf\xbd"
	// get the decimal string representation of an integer
	dump(strconv.Itoa(97))
}
