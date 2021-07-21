
### Go 字符串
在 go 语言中，字符串实际上是一个只读的字节切片，其数据结构定义如下：
```golang
// runtime/string.go
type stringStruct struct {
	str unsafe.Pointer	// 指向底层字节数组的指针
	len int				// 字节数组的长度 
}
```
`注意：byte 其实是 uint8 的类型别名`


### Go 字符串使用
```golang
// 使用字符串字面量初始化
var a = "hi,狗"
fmt.Println(a) // hi,狗

// 可以使用下标访问，但不可修改
fmt.Printf("a[0] is %d\n", a[0]) // a[0] is 104
fmt.Printf("a[0:2] is %s\n", a[0:2]) // a[0:2] is hi
// a[0] = 'a' 编译报错，Cannot assign to a[0]

// 字符串拼接
var b = a + "狗"
fmt.Printf("b is %s\n", b) // b is hi,狗狗
// 将 string 类型的 slice 中的元素用指定字符拼接起来
strings.Join(sli, " ") // 仅支持 [] string，不支持 array，也不支持 [] int

// 使用内置 len() 函数获取其长度
fmt.Printf("a's length is: %d\n", len(a)) // a's length is: 6

// 使用 for;len 遍历
for i := 0; i < len(a); i++ {
	fmt.Println(i, a[i])
}
// 0 104
// 1 105
// 2 44
// 3 231
// 4 139
// 5 151

// 使用 for;range 遍历
for i, v := range a {
	fmt.Println(i, v)
}
// 0 104
// 1 105
// 2 44
// 3 29399

// 字符串常量会在编译期分配到只读段，对应数据地址不可写入，相同的字符串常量不会重复存储
var a = "hello"
fmt.Println(a, &a, (*reflect.StringHeader)(unsafe.Pointer(&a))) // hello 0xc0000381f0 &{5033779 5}
a = "world"
fmt.Println(a, &a, (*reflect.StringHeader)(unsafe.Pointer(&a))) // world 0xc0000381f0 &{5033844 5}
var b = "hello"
fmt.Println(b, &b, (*reflect.StringHeader)(unsafe.Pointer(&b))) // hello 0xc000038220 &{5033779 5}

// 将字符串转化为字节切片
var a = "hi,狗"
b := []byte(a)
fmt.Println(b)	// [104 105 44 231 139 151]

// 也可以将字符串转化为 rune 切片（rune 其实是 int32 的类型别名）
var a = "hi,狗"
r := []rune(a)
fmt.Println(r) // [104 105 44 29399]

// 也可以使用 "unicode/utf8" 标准库，手动实现 for;range 语法糖相同的效果
var a = "hi,狗,lang"
for i, w := 0, 0; i < len(a); i += w {
	runeValue, width := utf8.DecodeRuneInString(a[i:])
	fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
	w = width
}
/* 
U+0068 'h' starts at byte position 0
U+0069 'i' starts at byte position 1
U+002C ',' starts at byte position 2
U+72D7 '狗' starts at byte position 3
U+002C ',' starts at byte position 6  # 此处可以看到狗占了3个字节
U+006C 'l' starts at byte position 7
U+0061 'a' starts at byte position 8
U+006E 'n' starts at byte position 9
U+0067 'g' starts at byte position 10
*/
```
