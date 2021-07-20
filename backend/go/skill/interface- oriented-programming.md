
### 面向接口编程
接口，在表面上是由几个没有主体代码的方法定义组成的集合体，有唯一的名称，可以被类或其他接口所实现（或者也可以说继承）。  
接口是一组规则的集合，它规定了实现本接口的类或接口必须拥有的一组规则。体现了自然界 “如果你是…… 则必须能……” 的理念。接口是在一定粒度视图上同类事物的抽象表示。注意这里我强调了在一定粒度视图上，因为 “同类事物” 这个概念是相对的，它因为粒度视图不同而不同。  

在系统分析和架构中，分清层次和依赖关系，每个层次不是直接向其上层提供服务（即不是直接实例化在上层中），而是通过定义一组接口，仅向上层暴露其接口功能，上层对于下层仅仅是接口依赖，而不依赖具体类。这样做的好处是显而易见的，首先对系统灵活性大有好处；另一个好处就是不同部件或层次的开发人员可以并行开工，从而提高效率。  

### Go 面向接口编程
```golang
package main

import "fmt"

// 想学功夫的对象
type KongFuer interface {
	// 必须具备的条件
	Leg()  // 腿
	Arm()  // 胳膊
	Fist() // 拳头
}

// 教跆拳道
func TeachTaekwondo(k KongFuer) {
	// 练习5遍腿法
	for i := 0; i < 5; i++ {
		k.Leg()
	}
	// 练习1遍拳头
	k.Fist()
	fmt.Println("跆拳道修炼完毕")
}

// 教泰拳
func TeachThaiBoxing(k KongFuer) {
	// 练习3遍腿法
	for i := 0; i < 3; i++ {
		k.Leg()
	}
	// 练习3遍胳膊
	for i := 0; i < 3; i++ {
		k.Arm()
	}
	// 练习1遍拳头
	k.Fist()
	fmt.Println("泰拳修炼完毕")
}


// 小明同学
type xiaoMing struct {
	Age int
	Sex string
}

// 实现KongFuer接口所要求的三个方法
func (x xiaoMing) Leg() {
	fmt.Println("小明1.8m的大长腿，开始练习")
}

func (x xiaoMing) Arm() {
	fmt.Println("小明的麒麟臂，开始练习")
}

func (x xiaoMing) Fist() {
	fmt.Println("小明沙包大的拳头，开始练习")
}

func main() {
	// 小明同学登场
	xiaoming := xiaoMing{
		Age: 18,
		Sex: "男",
	}

	fmt.Println("小明学习跆拳道")
	TeachTaekwondo(xiaoming)

	fmt.Println("-------------------")

	fmt.Println("小明学习泰拳")
	TeachThaiBoxing(xiaoming)
}
```

