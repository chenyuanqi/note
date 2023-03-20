
### Gorm 框架简要使用
golang orm 库旨在友好数据库开发。

**特性**  
- 全功能 ORM
- 关联 (Has One，Has Many，Belongs To，Many To Many，多态，单表继承)
- Create，Save，Update，Delete，Find 中钩子方法
- 支持 Preload、Joins 的预加载
- 事务，嵌套事务，Save Point，Rollback To Saved Point
- Context、预编译模式、DryRun 模式
- 批量插入，FindInBatches，Find/Create with Map，使用 SQL 表达式、Context Valuer 进行 CRUD
- SQL 构建器，Upsert，数据库锁，Optimizer/Index/Comment Hint，命名参数，子查询
- 复合主键，索引，约束
- Auto Migration
- 自定义 Logger
- 灵活的可扩展插件 API：Database Resolver（多数据库，读写分离）、Prometheus…
- 每个特性都经过了测试的重重考验
- 开发者友好

**安装**  
```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```

**快速使用**  
```go
package main

import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

func main() {
  db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // 迁移 schema
  db.AutoMigrate(&Product{})

  // Create
  db.Create(&Product{Code: "D42", Price: 100})

  // Read
  var product Product
  db.First(&product, 1) // 根据整型主键查找
  db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

  // Update - 将 product 的 price 更新为 200
  db.Model(&product).Update("Price", 200)
  // Update - 更新多个字段
  db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
  db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

  // Delete - 删除 product
  db.Delete(&product, 1)
}
```

### 快速增删改查
1. 数据库增加操作：

```go
// 创建一个新的 User 对象
user := &User{Name: "Tom", Age: 20}

// 将该对象插入到数据库中
if err := db.Create(user).Error; err != nil {
    // 处理错误
}
```

2. 数据库查询操作：

```go
// 查询单个对象
var user User
if err := db.First(&user, "name = ?", "Tom").Error; err != nil {
    // 处理错误
}

// 查询多个对象
var users []User
if err := db.Where("age > ?", 18).Find(&users).Error; err != nil {
    // 处理错误
}
```

3. 数据库修改操作：

```go
// 更新单个对象
if err := db.Model(&user).Update("name", "Jerry").Error; err != nil {
    // 处理错误
}

// 更新多个对象
if err := db.Model(&User{}).Where("age < ?", 18).Update("name", "Child").Error; err != nil {
    // 处理错误
}
```

4. 数据库删除操作：

```go
// 删除单个对象
if err := db.Delete(&user).Error; err != nil {
    // 处理错误
}

// 删除多个对象
if err := db.Where("age < ?", 18).Delete(&User{}).Error; err != nil {
    // 处理错误
}
```

5. 批量添加：

```go
// 创建一个 User 对象的切片
users := []User{
    {Name: "Tom", Age: 20},
    {Name: "Jerry", Age: 22},
    {Name: "Alice", Age: 24},
}

// 批量插入到数据库中
if err := db.Create(&users).Error; err != nil {
    // 处理错误
}
```

6. 批量删除：

```go
// 将所有年龄小于 18 岁的用户删除
if err := db.Where("age < ?", 18).Delete(&User{}).Error; err != nil {
    // 处理错误
}
```

7. 批量修改：

```go
// 将所有年龄小于 18 岁的用户的名字修改为 Child
if err := db.Model(&User{}).Where("age < ?", 18).Update("name", "Child").Error; err != nil {
    // 处理错误
}
```

以上是 GORM 对数据库增删查改和批量操作的代码说明。

8. 在 GORM 中，可以使用 FirstOrCreate 方法来实现记录不存在就创建，存在就更新的功能

```go
user := User{Name: "Alice", Age: 20}
db.FirstOrCreate(&user, User{Name: "Alice"})
```

9. 多条记录同时判断不存在就创建，存在就更新
```go
// 定义要创建或更新的记录结构体
type User struct {
    ID    int
    Name  string
    Email string
}

// 定义要创建或更新的记录列表
users := []User{
    {Name: "user1", Email: "user1@example.com"},
    {Name: "user2", Email: "user2@example.com"},
    {Name: "user3", Email: "user3@example.com"},
}

// 批量创建或更新记录
for _, user := range users {
    // 使用FirstOrCreate方法查询是否存在该记录，如果不存在则创建，存在则更新
    db.FirstOrCreate(&user, User{Name: user.Name})
    // 使用Updates方法更新记录，如果存在则更新，不存在则忽略
    db.Model(&user).Updates(User{Name: user.Name, Email: user.Email})
}
```

