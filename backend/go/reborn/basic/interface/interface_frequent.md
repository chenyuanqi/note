
## 常用 interface
```go
// 获取 tempItem 所属结构体
v := reflect.ValueOf(tempItem)
// 字段数量 NumField()
for i := 0; i < v.NumField(); i++ {
    // 字段
    field := v.Field(i)
    // 字段值
    value := field.Interface()
    if value == nil || value == "" || value == 0 {
        break
    }
}
```
