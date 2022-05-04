package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := 10
	ip := &i

	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3
	fmt.Println(i)

	p := new(person)
	//Name 是 person 的第一个字段不用偏移，可通过指针修改
	pName := (*string)(unsafe.Pointer(p))
	*pName = "vikey"
	//Age 并不是 person 的第一个字段，所以需要进行偏移，这样才能正确定位到 Age 字段这块内存，才可以正确的修改
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.Age)))
	*pAge = 20

	fmt.Println(*p)

	fmt.Println(unsafe.Sizeof(true))
	fmt.Println(unsafe.Sizeof(int8(0)))
	fmt.Println(unsafe.Sizeof(int16(10)))
	fmt.Println(unsafe.Sizeof(int32(10000000)))
	fmt.Println(unsafe.Sizeof(int64(10000000000000)))
	fmt.Println(unsafe.Sizeof(int(10000000000000000)))
	fmt.Println(unsafe.Sizeof(string("vikey")))
	fmt.Println(unsafe.Sizeof([]string{"vikey 1", "张三"}))

}

type person struct {
	Name string
	Age  int
}
