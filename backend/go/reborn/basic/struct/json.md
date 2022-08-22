
#### 内置库 encoding/json
内置库中 encoding/json，Go 代码中主要调用 Marshal 从数据结构生成 JSON 字符串的过程，并调用 Unmarshal 解析 JSON 的行为解组到数据结构。
```go
// 函数原型

// 接收任意类型的k,v，返回JSON encoding 数据
func Marshal(v interface{}) ([]byte, error)
// 将json对象转成对应的数据结构
func Unmarshal(data []byte, v interface{}) error
```


**json 读取**  
struct 结构体数据，转 json 格式。
```go
func structData() {
    type UserInfo struct {
        Name string `json:"name"` // 指定格式
        Age int `json:"age"`
    }
    user1 := &UserInfo{
        Name: "changhao",
        Age: 25,
    }
    res, err := json.Marshal(user1)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("%s\n", res) // {"age":25,"name":"changhao"}
}
```

map 数据，转 json格式。
```go
func mapData() {
    user := make(map[string]interface{})
    user["name"] = "changhao"
    user["age"] = 25
    res, err := json.Marshal(user)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("%s\n", res) // {"age":25,"name":"changhao"}
}
```


使用 MarshalIndent 函数输出格式化 json。  
MarshalIndent 类似于 Marshal，但应用 Indent 来格式化输出。根据缩进嵌套，输出中的每个 JSON 元素都将在以前缀开头的新行开始，后跟一个或多个缩进副本。
```go
// 函数原型

// prefix 定义前缀字符， indent 设置缩进
func MarshalIndent(v interface{}, prefix, indent string) ([] byte , error)

func indentData(){
    user := make(map[string]interface{})
    user["name"] = "changhao"
    user["age"] = 25
    res, err := json.MarshalIndent(user, "", "   ")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("%s\n", res)
}
// {
// "age": 25,
//  "name": "changhao"
//}
```

**json 写入**  
将数据格式通过 json.MarshalIndent() 函数转成 json 对象， 再借助 ioutil.WriteFile() 函数写入文件，后缀务必为 .json。
```go
func writeJson() {
    type UserInfo struct {
        Name string `json:"name"`
        Age int `json:"age"`
        Hobby []string `json:"hobby"`
    }
    user := &UserInfo{
        Name: "changhao",
        Age: 25,
        Hobby: []string{"看书","学习新技术","编程"},
    }
    data, err := json.MarshalIndent(user, "", "    ")
    if err != nil {
        panic(err)
    }
    err = ioutil.WriteFile("data.json", data, 0755)
    if err != nil {
        panic(err)
    }
}
```

**json 修改**  
将通过 Unmarshal 函数读取 json 数据， 然后转成 map 数据类型。 删除后，在通过写入流程写入到 json 文件中。
```go
func checkErr(err error) {
    if err != nil {
        panic(err)
    } 
}

func updateJson() {
    // 读取json数据
    data := make(map[string]interface{})
    file, err := ioutil.ReadFile("data.json")
    checkErr(err)
    err = json.Unmarshal(file, &data)
    checkErr(err)
    
    // 读取爱好列表, 新增新的爱好
    hobbyList := make([]string,0)
    for _, hobby := range data["hobby"].([]interface{}) {
        hobbyList = append(hobbyList, hobby.(string))
    }
    hobbyList = append(hobbyList, "玩游戏")
    data["hobby"] = hobbyList
    
    // 写入 data.json
    dataJson, err := json.MarshalIndent(data, "", "   ")
    checkErr(err)
    err = ioutil.WriteFile("data.json", dataJson, 0755)
    checkErr(err)
}
// {
//   "age": 25,
//   "hobby": [
//     "看书",
//     "学习新技术",
//     "编程",
//     "玩游戏"
//   ],
//   "name": "changhao"
// }
```

**json 删除**    
```go
func checkErr(err error) {
    if err != nil {
        panic(err)
    } 
}

func removeJson() {
    // 读取json数据
    data := make(map[string]interface{})
    file, err := ioutil.ReadFile("data.json")
    checkErr(err)
    err = json.Unmarshal(file, &data)
    checkErr(err)
    
    // 删除 key
    delete(data, "age")

    // 写入 data.json
    dataJson, err := json.MarshalIndent(data, "", "   ")
    checkErr(err)
    err = ioutil.WriteFile("data.json", dataJson, 0755)
    checkErr(err)
}

// {
//   "list": [
//      {"name": "小明", "age": 18},
//      {"name": "小红", "age": 19},
//   ],
// }
```

### json 使用场景

**场景一：map 数据结构嵌套**  
写入 -- 首先，定义数据结构，map 写入数据。 调用 json.MarshalIndent 函数设置 json 格式，最后使用 ioutil.WriteFile 函数写入 json 文件。
```go
func checkErr(err error) {
    if err != nil {
        panic(err)
    } 
}

func write() {
    // 创建json结构
    data := make(map[string]interface{})
    // 定义list 列表
    list := make([]map[string]interface{},0)
    // 定义类型
    type infoMap = map[string]interface{}
    info1 := make(infoMap)
    info1["name"] = "小明"
    info1["age"] = 18
    list = append(list, info1)
    info2 := make(infoMap)
    info2["name"] = "小红"
    info2["age"] = 19
    list = append(list, info2)
    data["list"] = list

    // 转换json数据格式
    dataJson, err := json.MarshalIndent(data, "", "   ")
    checkErr(err)
    // 写入json文件
    err = ioutil.WriteFile("data.json", dataJson, 0755)
    checkErr(err)
}
// {
//    "list": [
//      {
//        "age": 18,
//        "name": "小明"
//      },
//      {
//        "age": 19,
//        "name": "小红"
//      }
//    ]
// }
```

修改 -- 首先，使用 ioutil.ReadFile函数 读取json文件，定义map结构体，用于接收Unmarshal 函数返回的 json对象，然后将列表中的小红、小明年龄统一修改为 20 岁。最后在调用 ioutil.WriteFile 函数将修改后的数据格式， 写入json文件。

```go
func checkErr(err error) {
    if err != nil {
        panic(err)
    } 
}

func update() {
    // 接收json结构
    data := make(map[string]interface{})
    file, err := ioutil.ReadFile("data.json")
    checkErr(err)
    err = json.Unmarshal(file, &data)
    checkErr(err)
    // 定义list 列表
    list := make([]map[string]interface{},0)
    for _, info := range data["list"].([]interface{}) {
        v := info.(map[string]interface{})
        v["age"] = 20
        list = append(list, v)
    }
    data["list"] = list

    // 转换json数据格式
    dataJson, err := json.MarshalIndent(data, "", "   ")
    checkErr(err)
    // 写入json文件
    err = ioutil.WriteFile("data.json", dataJson, 0755)
    checkErr(err)
}
// {
//     "list": [
//         {
//             "age": 20,
//             "name": "小明"
//         },
//         {
//             "age": 20,
//             "name": "小红"
//         }
//     ]
// }
```

**场景二：结构体数据嵌套**  
写入 -- 和场景一相同， 主要将数据格式以结构体封装实现。
```go
func checkErr(err error) {
    if err != nil {
        panic(err)
    } 
}

func write() {
    // 定义结构体
    type ListItem struct {
        Name string `json:"name"`
        Age int `json:"age"`
    }
    type Info struct {
        List []ListItem `json:"list"`
    }

    // 写入数据
    var info Info
    item := make([]ListItem, 0)
    listItem1 := ListItem{
        Name: "小明",
        Age: 18,
    }
    item = append(item, listItem1)

    listItem2 := ListItem{
        Name: "小红",
        Age: 19,
    }
    item = append(item, listItem2)
    info.List = item

    data, err := json.MarshalIndent(info, "", "    ")
    checkErr(err)
    err = ioutil.WriteFile("data.json", data, 0755)
    checkErr(err)
}
```

修改。
```go
func checkErr(err error) {
    if err != nil {
        panic(err)
    } 
}

func update() {
    // 定义结构体
    type ListItem struct {
        Name string `json:"name"`
        Age int `json:"age"`
    }
    type Info struct {
        List []ListItem `json:"list"`
    }

    // 接收json结构
    var info Info
    file, err := ioutil.ReadFile("data.json")
    checkErr(err)
    err = json.Unmarshal(file, &info)
    checkErr(err)
    // 修改年龄
    itemList := make([]ListItem, 0)
    for _, item := range info.List {
        item.Age = 20
        itemList = append(itemList, item)

    }
    info.List = itemList

    // 写入json
    data, err := json.MarshalIndent(info, "", "    ")
    checkErr(err)
    err = ioutil.WriteFile("data.json", data, 0755)
    checkErr(err)
}
```


