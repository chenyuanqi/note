
### 结构体标签
Go 语言允许我们通过结构体字段标签给一个字段附加可以被反射获取的”元信息“。  
通常情况下，结构体标签被用于提供结构体字段如何被编码为或者解码自另外一种格式的转换信息（或者是以何种形式被保存至/获取自数据库）。不过，你也可以用它存储任何你想要设置的”元信息“，供其他包或者自己使用。  

结构体标签在使用上通常是遵守下面三个规范。
- 结构体标签字符串的值是一个由空格分隔的 key:"value" 对列表，如 Name string `json:"name" xml:"name"`  
- 键，通常表示后面跟的“值”是被哪个包使用的，例如json这个键会被encoding/json包处理使用。如果要键对应的“值”中传递多个信息，通常通过用逗号（'，'）分隔来指定，如 Name string `json:"name,omitempty"`  
- 按照管理，如果一个字段的结构体标签里某个键的“值”被设置成了的破折号 ('-')，那么就意味着告诉处理该结构体标签键值的进程排除该字段，如 Name string `json:"-"` 进行JSON编码/解码时忽略 Name 这个字段

怎么获取到结构体标签  
从一开始我们就说结构体标签是给反射准备的，那么怎么在Go程序里用反射获得到字段的结构体标签呢？结构体字段类型相关的信息，在反射的世界里使用reflect.StructFiled 这个类型表示的。
```go
type StructField struct {
	Name string
	Type      Type      // field type
	Tag       StructTag // field tag string，字段声明中的结构体标签信息
  ......
}
```
用反射获取到自定义的结构体标签。  
使用反射 reflect 包访问结构体字段的标签值，我们需要先获取到结构体的类型信息 Type，然后使用 Type.Field(i int) 或 Type.FieldByName(name string)，方法查询字段信息，这两个方法都会返回一个 StructField 类型的值，上面我们也说了它在反射的世界里用于描述一个结构体字段；而 StructField.Tag 是一个 StructTag 类型的值，它描述了字段的标签值。  
```go
// 解析标签的值并返回你指定的键的“值”
func (tag StructTag) Get(key string) string
// 判断一个给定的key是否存在与标签中，返回的ok值告知给定key是否存在与标签中
func (tag StructTag) Lookup(key string) (value string, ok bool)

package main
import (
	"fmt"
	"reflect"
)
type User struct {
	Name  string `mytag:"MyName"`
	Email string `mytag:"MyEmail"`
}
func main() {
	u := User{"Bob", "bob@mycompany.com"}
	t := reflect.TypeOf(u)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field: User.%s\n", field.Name)
		fmt.Printf("\tWhole tag value : %s\n", field.Tag)
		fmt.Printf("\tValue of 'mytag': %s\n", field.Tag.Get("mytag"))
	}
    // Field: User.Name
    //     Whole tag value : mytag:"MyName"
    //     Value of 'mytag': MyName
    // Field: User.Email
    //     Whole tag value : mytag:"MyEmail"
    //     Value of 'mytag': MyEmail
}
```

常用的结构体标签键  
常用的结构体标签 Key，指的是那些被一些常用的开源包声明使用的结构体标签键。一些我们平时会用到的包，它们是：
- json: 由encoding/json 包使用，详见json.Marshal()的使用方法和实现逻辑。
- xml : 由encoding/xml包使用，详见xml.Marshal()。
- bson: 由gobson包，和mongo-go包使用。
- protobuf: 由github.com/golang/protobuf/proto 使用，在包文档中有详细说明。
- yaml: 由gopkg.in/yaml.v2 包使用，详见yaml.Marshal()。
- gorm: 由gorm.io/gorm包使用，示例可以在GORM的文档中找到。


