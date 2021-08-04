
### json 包
Go 语言标准库 encoding/json 提供了操作 JSON 的方法，一般可以使用 json.Marshal 和 json.Unmarshal 来序列化和解析 JSON 字符串。  
```golang
// 定义结构体
type User struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

// 序列化
buf, err := json.Marshal(User{
    Email:    "xxxx@example.com",
    Password: "123456",
})

// 解析
user := User{}
err := json.Unmarshal([]byte(`{
  "email": "xxx@example.com",
  "password": "123456"
}`), &user)
```

**更加灵活和更好性能的 jsoniter 模块**  
标准库 encoding/json 在使用时需要预先定义结构体，使用时显得不够灵活。这时候可以尝试使用 github.com/json-iterator/go 模块，其除了提供与标准库一样的接口之外，还提供了一系列更加灵活的操作方法。
```golang
val := []byte(`{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`)

// 仅解析 Colors 字段，并直接得到 string 类型
str := jsoniter.Get(val, "Colors", 0).ToString()
```

**另辟蹊径提高性能的 easyjson 模块**  
标准库 encoding/json 需要依赖反射来实现，因此性能上会比较差。 github.com/mailru/easyjson 则是利用 go generate 机制自动为结构体生成实现了 MarshalJSON 和 UnmarshalJSON 方法的代码，在序列化和解析时可以直接生成对应字段的 JSON 数据，而不需要运行时反射。据官方的介绍，其性能是标准库的 4～5 倍，是其他 json 模块的 2～3 倍。  

要使用 easyjson 模块，首先执行以下命令安装 easyjson 命令。
```bash
go get -u github.com/mailru/easyjson/
go install  github.com/mailru/easyjson/easyjson
# or
go build -o easyjson github.com/mailru/easyjson/easyjson
```

```golang
import (
	"fmt"
	"studygo/easyjson"
	"time"
)

// easyjson:json
type School struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
}

func main() {
	s := easyjson.Student{
		Id:   11,
		Name: "qq",
		School: easyjson.School{
			Name: "CUMT",
			Addr: "xz",
		},
		Birthday: time.Now(),
	}
	bt, err := s.MarshalJSON()
	fmt.Println(string(bt), err)
	json := `{"id":11,"s_name":"qq","s_chool":{"name":"CUMT","addr":"xz"},"birthday":"2017-08-04T20:58:07.9894603+08:00"}`
	ss := easyjson.Student{}
	ss.UnmarshalJSON([]byte(json))
	fmt.Println(ss)
}
```
