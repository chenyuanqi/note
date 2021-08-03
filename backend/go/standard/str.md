
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

// 多行字符串
results := `Search results for "Golang":
- Go
- Golang
Golang Programming
`
fmt.Printf("%s", results)
// 或是这样
results := "Search results for \"Golang\":\n" +
"- Go\n" +
"- Golang\n" +
"- Golang Programming\n"
fmt.Printf("%s", results)

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

// 字符串的类型转换
v1 := "100"
v2, _ := strconv.Atoi(v1)  // 将字符串转化为整型，v2 = 100

v3 := 100
v4 := strconv.Itoa(v3)   // 将整型转化为字符串, v4 = "100"

v5 := "true"
v6, _ := strconv.ParseBool(v5)  // 将字符串转化为布尔型
v5 = strconv.FormatBool(v6)  // 将布尔值转化为字符串

v7 := "100"
v8, _ := strconv.ParseInt(v7, 10, 64)   // 将字符串转化为整型，第二个参数表示进制，第三个参数表示最大位数
v7 = strconv.FormatInt(v8, 10)   // 将整型转化为字符串，第二个参数表示进制

v9, _ := strconv.ParseUint(v7, 10, 64)   // 将字符串转化为无符号整型，参数含义同 ParseInt
v7 = strconv.FormatUint(v9, 10)  // 将字符串转化为无符号整型，参数含义同 FormatInt

v10 := "99.99"
v11, _ := strconv.ParseFloat(v10, 64)   // 将字符串转化为浮点型，第二个参数表示精度
v10 = strconv.FormatFloat(v11, 'E', -1, 64)

q := strconv.Quote("Hello, 世界")    // 为字符串加引号
q = strconv.QuoteToASCII("Hello, 世界")  // 将字符串转化为 ASCII 编码
```

### Go 字符串格式化输出
fmt 包实现了类似 C 语言 printf 和 scanf 的格式化 I/O。格式化 verb（'verb'）源自 C 语言但更简单。  

**通用**  
%v    值的默认格式表示。当输出结构体时，扩展标志（%+v）会添加字段名  
%+v   类似%v，但输出结构体时会添加字段名  
%#v   值的 Go 语法表示  
%T    值的类型的 Go 语法表示  
%%    百分号  

**布尔值**  
%t    单词true或false  

**整数**  
%b    表示为二进制  
%c    该值对应的 unicode 码值  
%d    表示为十进制   
%o    表示为八进制  
%q    该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示  
%x    表示为十六进制，使用 a-f  
%X    表示为十六进制，使用 A-F  
%U    表示为 Unicode 格式：U+1234，等价于"U+%04X"  

**浮点数/复数**  
%f:    默认宽度，默认精度  
%9f    宽度9，默认精度  
%.2f   默认宽度，精度2 %9.2f  宽度9，精度2 %9.f 宽度9，精度0   
%b    无小数部分、二进制指数的科学计数法，如-123456p-78  
	  参见strconv.FormatFloat %e    科学计数法，如-1234.456e+78 %E    
	  科学计数法，如-1234.456E+78 %f    
	  有小数部分但无指数部分，如123.456 %F    等价于%f %g     
	  根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）  
	  
%e     科学计数法，例如 -1234.456e+78   
%E     科学计数法，例如 -1234.456E+78   
%f     有小数点而无指数，例如 123.456   
%g     根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出   
%G     根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出  

**字符串和 []byte**  
%s    直接输出字符串或者[]byte %q 该值对应的双引号括起来的 go 语法字符串字面值，必要时会采用安全的转义表示  
%x    每个字节用两字符十六进制数表示（使用a-f）  
%X    每个字节用两字符十六进制数表示（使用A-F）  

```golang
fmt.Sprintf(格式化样式, 参数列表…)

user := User{"xiaoming", 13}
//Go默认形式
fmt.Printf("%v",user)
fmt.Println()
//类型+值对象
fmt.Printf("%#v",user)
fmt.Println()
//输出字段名和字段值形式
fmt.Printf("%+v",user)
fmt.Println()
//值类型的Go语法表示形式
fmt.Printf("%T",user)
fmt.Println()
fmt.Printf("%%")

fmt.Println(&k)  // 获取变量在计算机内存中的地址，可在变量名前面加上&字符
```


