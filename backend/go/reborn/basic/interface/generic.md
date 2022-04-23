
### Go 泛型
所谓泛型，可以简单地理解为数据类型中的“万能牌”，它能存放任何类型的数据。Go 语言中的空接口正是为了实现泛型所采用的手段。  
泛型是类型中的“万能牌”，使用泛型作为函数参数，实际上就相当于告诉调用者：“我能兼容任何类型的参数，尽管将数据传给我就是了。”泛型以超级宽广的胸怀接纳所有类型的数据。在 Go 语言中的泛型，则使用空接口来实现。  
而所谓的“空接口”，使用代码表示非常简单，就是：interface{}。和普通接口的定义格式不同，空接口内部无需填写任何方法。空接口能接纳所有类型的数据，因此可以将任何类型的数据赋值给它的变量。  

什么时候该使用泛型呢？举个例子，如果我们想要封装一个函数，该函数的作用便是实现传入参数数据的原样输出，该如何做呢？
```go
func main() {
   dataOutput("Hello")
   dataOutput(123)
   dataOutput(true)
}

// 在函数参数中使用空接口，可以使其能接受所有类型的数据传入
func dataOutput(data interface{}) {
   fmt.Println(data)
}
```

**泛型使用示例：货车容量计算器**  
在实际开发中，非常巧妙的使用空接口的方式，它可以规避数据类型的不同，将不同类型的数据存放于同一个切片/数组中，对于组织大量具有不同类型的数据是非常有效的做法。

我们计划进行一次搬家，正在预估需要多大容量的货车来存放全部家当。  
为了方便，简化各种家具家电的体积计算方式。把它们简单粗暴地分为正方体、长方体和圆柱体三种体积形式，这三种形状的物品分别对应代码中的三种结构体类型。  
此外，还需实现为这三种形状的物品编写体积计算的方法。如此一来，我们便可通过调用这个体积计算的方法，将其计算结果累加在一起，便可得知需要至少多大容量的货车了。  
```go
package main

import (
   "fmt"
   "math"
)

func main() {
   truckSize := 0.0
   // 声明空接口类型变量materials，存放各种不同体积的家具
   var materials []interface{}
   materials = append(materials, cube{12.5})
   materials = append(materials, cuboid{25, 13, 60})
   materials = append(materials, cylinder{5, 25.3})
   // 遍历materials切片，依次计算每个家具的体积，并相加求和
   for _, singleMaterial := range materials {
      truckSize += calcSize(singleMaterial)
   }
   fmt.Println(truckSize)
}

// 计算某个物体的体积
func calcSize(material interface{}) float64 {
   cubeMaterial, cubeOk := material.(cube)
   cuboidMaterial, cuboidOk := material.(cuboid)
   // 判断某个数据是否属于某种类型的方法被称为“类型断言”
   cylinderMaterial, cylinderOk := material.(cylinder)
   if cubeOk {
      return cubeMaterial.cubeVolume()
   } else if cuboidOk {
      return cuboidMaterial.cuboidVolume()
   } else if cylinderOk {
      return cylinderMaterial.cylinderVolume()
   } else {
      return 0
   }
}

// 正方体
type cube struct {
   // 边长
   length float64
}

// 正方体的体积计算
func (c *cube) cubeVolume() float64 {
   return c.length * c.length * c.length
}

// 长方体
type cuboid struct {
   // 长
   length float64
   // 宽
   width float64
   // 高
   height float64
}

// 长方体的体积计算
func (c *cuboid) cuboidVolume() float64 {
   return c.length * c.width * c.height
}

// 圆柱体
type cylinder struct {
   // 直径
   diameter float64
   // 高度
   height float64
}

// 圆柱体的体积计算
func (c *cylinder) cylinderVolume() float64 {
   return math.Pi * (c.diameter / 2) * (c.diameter / 2) * c.height
}
```
