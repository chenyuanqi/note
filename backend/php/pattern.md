
### 什么是设计模式
在面向对象中，类是用于生成对象的代码模版，设计模式是用于解决共性问题的代码模版；遵循这样的模板，我们可以快速设计出优秀的代码。
（注意：设计模式只是模板，不是具体的代码；它是为了代码复用，增加可维护性。）

设计模式的宗旨：**重用**。  
面向对象的设计模式都是类之间关系的组合，譬如依赖注入
```php
class Human{}

class Woman
{
    public function __construct(Human $human){}
}
```

### UML 类图
UML 类图（类图中的类，与面向对象语言中的类的概念是对应的）是一种结构图，用于描述一个系统的静态结构。类图以反映类结构和类之间关系为目的，用以描述软件系统的结构，是一种静态建模方法。  
UML 类图是面向对象设计的辅助工具，但并非是必须工具。  

UML 类图推荐使用免费的 [UMLet工具](http://www.umlet.com/umletino/umletino.html)。

类与类之间的关系主要有六种：继承→实现→组合→聚合→关联→依赖。

### 设计原则
设计模式有六大原则（SOLID）。  
SRP The Single Responsibility Principle: -- a class should have one, and only one, reason to change.
OCP The Open Closed Principle: -- you should be able to extend a class's behavior, without modifying it.
LSP The Liskov Substitution Principle: -- derived classes must be substitutable for their base classes.
ISP The Interface Segregation Principle: -- make fine grained interfaces that are client specific.
DIP The Dependency Inversion Principle -- depend on abstractions not on concrete implementations.