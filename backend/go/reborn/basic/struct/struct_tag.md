

1. `json` 标签：这个标签常用于定义 JSON 序列化和反序列化时字段的名称。它可以让我们在序列化和反序列化时自定义字段名，实现更加优雅的数据交互。

```go
type User struct {
    ID       int64  `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email,omitempty"`
}
```

2. `xml` 标签：与 `json` 标签类似，`xml` 标签用于定义 XML 序列化和反序列化时的字段名称。

```go
type Person struct {
    Name string `xml:"name"`
    Age  int    `xml:"age"`
}
```

3. `bson` 标签：`bson` 标签用于处理 MongoDB 中的 BSON 数据格式，类似于 JSON，它也可以定义序列化和反序列化时的字段名称。

```go
type Product struct {
    Name  string `bson:"name"`
    Price int    `bson:"price"`
}
```

4. `gorm` 标签：`gorm` 是一款流行的 Go 语言 ORM 库，它的标签用于定义结构体与数据库表之间的映射关系，如表名、字段名、索引等。

```go
type Employee struct {
    ID     uint   `gorm:"primary_key"`
    Name   string `gorm:"column:name;type:varchar(100);not null"`
    Salary int    `gorm:"column:salary;default:0"`
}
```

5. `validate` 标签：`validate` 标签用于数据验证，我们可以在此标签中定义各种验证规则，例如非空、最小值、最大值等。

```go
type RegisterForm struct {
    Email    string `validate:"required,email"`
    Password string `validate:"required,min=6,max=20"`
}
```

6. `yaml` 标签：`yaml` 标签用于处理 YAML 数据格式，类似于 JSON 和 XML，我们可以使用它来定义序列化和反序列化时的字段名称。

```go
type Configuration struct {
    Address string `yaml:"address"`
    Port    int    `yaml:"port"`
}
```

7. `msgpack` 标签：`msgpack` 标签用于处理 MessagePack 数据格式，它是一种二进制数据交换格式。我们可以使用此标签定义序列化和反序列化时的字段名称。

```go
type Book struct {
    Title  string `msgpack:"title"`
    Author string `msgpack:"author"`
}
```

8. `protobuf` 标签：`protobuf` 标签用于处理 Protocol Buffers 数据格式，它是 Google 开发的一种数据序列化协议。通过这个标签，我们可以定义字段在序列化和反序列化时的编号。

```protobuf
message Person {
    string name = 1;
    int32 age = 2;
}
```

9. `csv` 标签：`csv` 标签用于处理 CSV 数据格式，我们可以使用它来定义 CSV 文件中的字段名称。

```go
type Record struct {
    Date string `csv:"date"`
    Open string `csv:"open"`
    High string `csv:"high"`
}
```

10. `form` 标签：`form` 标签用于处理表单数据，我们可以通过这个标签定义表单字段名称，方便在解析表单时自动映射到结构体字段。

```go
type LoginForm struct {
    Username string `form:"username"`
    Password string `form:"password"`
}
```

11. 自定义标签：在一些特定的场景下，我们可能需要自定义一些标签来满足特殊需求。自定义标签的用法与内置标签类似，只需在结构体字段上添加相应的标签即可。

```go
type UserInfo struct {
    Name     string `mytag:"name"`
    Birthdate string `mytag:"birthdate"`
}

func main() {
    u := UserInfo{
        Name: "张三",
        Birthdate: "1990-01-01",
    }

    t := reflect.TypeOf(u)
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        fmt.Printf("Field: %s, Tag: %v\n", field.Name, field.Tag.Get("mytag"))
    }
}
```

12. 标签的组合使用：在实际项目中，我们可能需要在一个结构体字段上应用多个标签。此时，可以通过在标签字符串中用空格隔开多个标签来实现。

```go
type User struct {
    ID       int64  `json:"id" bson:"_id"`
    Username string `json:"username" bson:"username"`
}
```

13. 利用反射获取标签：在某些情况下，我们需要在运行时动态获取结构体标签的信息。此时，可以使用 Go 语言的反射（reflect）包来实现。

```go
func PrintTags(obj interface{}) {
    t := reflect.TypeOf(obj)
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        fmt.Printf("Field: %s, JSON Tag: %v\n", field.Name, field.Tag.Get("json"))
    }
}
```

14. 在 Gorm 中自定义全局的标签
在 Gorm 中，我们可以通过实现 `gorm.SchemaNaming` 接口来自定义全局的标签。该接口包含一个 `NameSchema` 方法，用于处理结构体和字段的名称。

```go
type CustomSchemaNaming struct{}

func (c CustomSchemaNaming) NameSchema(schema gorm.SchemaNaming, value interface{}) string {
    // 自定义逻辑，例如将字段名称转为下划线形式
    return toSnakeCase(schema.DefaultName(value))
}

func toSnakeCase(str string) string {
    var result []rune
    for i, r := range str {
        if unicode.IsUpper(r) {
            if i > 0 {
                result = append(result, '_')
            }
            result = append(result, unicode.ToLower(r))
        } else {
            result = append(result, r)
        }
    }
    return string(result)
}
```

接下来，在创建 Gorm 实例时，将自定义的命名策略注册到 Gorm 中。

```go
func main() {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
        NamingStrategy: CustomSchemaNaming{},
    })
    if err != nil {
        panic("Failed to connect database")
    }

    // Your code...
}
```

以上代码示例展示了如何在 Gorm 中自定义全局的标签。通过实现 `gorm.SchemaNaming` 接口并在创建 Gorm 实例时注册，我们可以轻松地实现自定义的结构体和字段命名策略。
