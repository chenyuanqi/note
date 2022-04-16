
### goto 语句
Go 语言中 goto 语句通过标签进行代码间的无条件跳转，同时 goto 语句在快速跳出循环、避免重复退出上也有一定的帮助，使用 goto 语句能简化一些代码的实现过程。  
goto 和” 条件语句 “、” 循环语句 “都没有关系，意思是直接跳到指定位置，继续执行逻辑，位置的命名随意。  
```go
    for x := 0; x < 10; x++ {
        for y := 0; y < 10; y++ {
            if y == 2 {
                // 跳转到标签
                goto breakHere
            }
        }
    }
    // 手动返回, 避免执行进入标签
    return
    // 标签
breakHere:
    fmt.Println("done")

// 使用 goto 集中处理错误
    err := firstCheckError()
    if err != nil {
        goto onExit
    }
    err = secondCheckError()
    if err != nil {
        goto onExit
    }
    fmt.Println("done")
    return
onExit:
    fmt.Println(err)
    exitProcess()
```
