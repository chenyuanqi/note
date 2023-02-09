package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	// 简单的想要获取命令行参数
	// os.Args 是一个 []string，0 => 执行文件路径
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}

	// flag 使用
	// 定义命令行参数方式
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "delay", 0, "延迟的时间间隔")

	//解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())

	// 查看帮助：go run flag.go --help
	// 示例：go run flag.go --age=18 --delay=1h30m --name=李四 --married=false
}
