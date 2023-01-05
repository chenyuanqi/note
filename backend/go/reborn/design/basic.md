
### Go 设计模式   
四人组（Gang of Four），简称 GoF。  
GoF给软件设计模式提供了定义：软件设计模式(Design Pattern)是一套被反复使用、多数人知晓的、经过分类编目的、代码设计经验的总结，使用设计模式是为了可重用代码、让代码更容易被他人理解并且保证代码可靠性。”  
一句大白话可以总结为：“在一定环境下，用固定套路解决问题。”  

GoF 提出的设计模式有 23 个，包括：  
（1）创建型(Creational)模式：如何创建对象；  
（2）结构型(Structural )模式：如何实现类或对象的组合；  
（3）行为型(Behavioral)模式：类或对象怎样交互以及怎样分配职责。  

设计模式的基础是：多态。  
初学者：积累案例，不要盲目的背类图。  
初级开发人员：多思考，多梳理，归纳总结，尊重事物的认知规律，注意临界点的突破，不要浮躁。  
中级开发人员：合适的开发环境，寻找合适的设计模式来解决问题。  
多应用，对经典则组合设计模式的大量，自由的运用。要不断的追求。  

### 面向对象设计原则
对于面向对象软件系统的设计而言，在支持可维护性的同时，提高系统的可复用性是一个至关重要的问题，*如何同时提高一个软件系统的可维护性和可复用性是面向对象设计需要解决的核心问题之一*。在面向对象设计中，可维护性的复用是以设计原则为基础的。每一个原则都蕴含一些面向对象设计的思想，可以从不同的角度提升一个软件结构的设计水平。  
*面向对象设计原则为支持可维护性复用而诞生，这些原则蕴含在很多设计模式中，它们是从许多设计方案中总结出的指导性原则*。面向对象设计原则也是我们用于评价一个设计模式的使用效果的重要指标之一。
> 原则的目的： 高内聚，低耦合

1、单一职责原则 (Single Responsibility Principle, SRP)  
类的职责单一，对外只提供一种功能，而引起类变化的原因都应该只有一个。  
```go
package main

import "fmt"

type ClothesShop struct {}

func (cs *ClothesShop) OnShop() {
	fmt.Println("休闲的装扮")
}

type ClothesWork struct {}

func (cw *ClothesWork) OnWork() {
	fmt.Println("工作的装扮")
}

func main() {
	//工作的时候
	cw := new(ClothesWork)
	cw.OnWork()

	//shopping的时候
	cs := new(ClothesShop)
	cs.OnShop()
}
```

2、开闭原则 (Open-Closed Principle, OCP)  
类的改动是通过增加代码进行的，而不是修改源代码。  
```go
package main

import "fmt"

//抽象的银行业务员
type AbstractBanker interface{
	DoBusi()	//抽象的处理业务接口
}

//存款的业务员
type SaveBanker struct {
	//AbstractBanker
}

func (sb *SaveBanker) DoBusi() {
	fmt.Println("进行了存款")
}

//转账的业务员
type TransferBanker struct {
	//AbstractBanker
}

func (tb *TransferBanker) DoBusi() {
	fmt.Println("进行了转账")
}

//支付的业务员
type PayBanker struct {
	//AbstractBanker
}

func (pb *PayBanker) DoBusi() {
	fmt.Println("进行了支付")
}


func main() {
	//进行存款
	sb := &SaveBanker{}
	sb.DoBusi()

	//进行转账
	tb := &TransferBanker{}
	tb.DoBusi()
	
	//进行支付
	pb := &PayBanker{}
	pb.DoBusi()

}
```

3、里氏代换原则 (Liskov Substitution Principle, LSP）  
任何抽象类（interface接口）出现的地方都可以用他的实现类进行替换，实际就是虚拟机制，语言级别实现面向对象功能。  

4、依赖倒转原则 (Dependence  Inversion Principle, DIP)  
依赖于抽象(接口)，不要依赖具体的实现(类)，也就是针对接口编程。  

5、接口隔离原则 (Interface Segregation Principle, ISP）   
不应该强迫用户的程序依赖他们不需要的接口方法。一个接口应该只提供一种对外功能，不应该把所有操作都封装到一个接口中去。  

6、合成复用原则 (Composite Reuse Principle, CRP)  
如果使用继承，会导致父类的任何变换都可能影响到子类的行为。如果使用对象组合，就降低了这种依赖关系。对于继承和组合，优先使用组合。  

7、迪米特法则 (Law of Demeter, LoD）  
一个对象应当对其他对象尽可能少的了解，从而降低各个对象之间的耦合，提高系统的可维护性。例如在一个程序中，各个模块之间相互调用时，通常会提供一个统一的接口来实现。这样其他模块不需要了解另外一个模块的内部实现细节，这样当一个模块内部的实现发生改变时，不会影响其他模块的使用。（黑盒原理）  

