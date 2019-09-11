
### Java Object
所有的 class 最终都继承自 Object，Object 定义了几个重要的方法：  
> toString() 把 instance 输出为 String  
> equals() 判断两个 instance 是否逻辑相等  
> hashCode() 计算一个 instance 的哈希值  
> finalize() GC 在回收对象之前不确定性的调用  

### Java 面向对象
世间万物都可以看成一个对象，每个物体包括动态的行为和静态的属性，这些就构成了一个对象。类是对象的抽象，对象是类的具体，类是对象的模板，对象是类的实例。  

面向对象是基于面向过程的编程思想。  
面向对象更符合我们的思维习惯，可以把复杂的事情简单化，角色从执行者变成指挥者。  
在 Java 中，创建一个类 Person，包含多个字段（field），初步认识 Java 的面向对象。  

Java 内建的访问权限包括 public、protected、private 和 package 权限；  
在方法内部定义的变量是局部变量，局部变量的作用域从变量声明开始，到一个块结束；  
final 修饰符不是访问权限，它可以修饰 class、field 和 method。  
> 默认（friendly）： 字段属性和方法只能被当前类和同一个包下的方法  
> public: 字段属性和方法可以被当前类和同一包、其他包、实例和子类访问，以及子类的子类  
> protected: 字段属性和方法可以被当前类和同一包、实例和子类访问，以及子类的子类，不能被其他包访问  
> private: 字段属性和方法不可以被实例或子类访问，只能被类自己访问，可以用在 static 方法上  
> package: 一个类允许访问同一个 package 的没有 public、private 修饰的 class，以及没有 public、protected、private 修饰的字段和方法；在同一个包，就可以访问 package 权限的 class、field 和 method；包没有父子关系；使用它有助于测试

一个 .java 文件只能包含一个 public 类，但可以包含多个非 public 类；如果有 public 类，文件名必须和 public 类的名字相同。  
```java
import java.io.*;
import java.util.*;

// 没有 extends，编译器会自动加上 extends Object（Object 是顶级基类）
class Person 
{
    // private 拒绝外部访问，保证类的封装型
    // 创建实例：先初始化字段（无默认值分配默认值），再执行构造方法
    private String name;
    /* private int age; */
    private int currentYear;
    private int birthYear;
    private String[] hobbies;
    private boolean isGoodman = true;
    protected String sex;

    // 构造方法的方法名与类名相同，它没有返回值类型（连 void 都没有）
    // 默认构造方法 public Person(){}，如果类内已经自定义则不默认创建
    // 可以并存多个构造方法，不同参数的构造方法对应不同的实例化，编译器会自动区分
    // 构造方法内部还可以调用其他的构造方法，便于代码复用
    public Person(int currentYear)
    {
        this.currentYear = currentYear;
    }

    public String getName()
    {
        // this 始终指向当前实例
        // 即方法被哪个对象调用，this 就代表那个对象
        return this.name;
    }

    public void setName(String name)
    {
        // 不允许传入 null 和空字符串
        if (name == null || name.isBlank()) {
            throw new IllegalArgumentException("invalid name value.");
        }

        this.name = name.strip();
    }

    // 方法重载（Overload）
    // 方法重载是指多个方法的方法名相同，但各自的参数不同
    // 重载方法应该完成类似的功能，返回值类型也应该相同
    public void setName(String familyName, String fullName)
    {
        // 不允许传入 null 和空字符串
        if (familyName == null || familyName.isBlank()) {
            throw new IllegalArgumentException("invalid family name value.");
        } else if (fullName == null || fullName.isBlank()) {
            throw new IllegalArgumentException("invalid full name value.");
        }

        this.name = familyName.strip() + fullName.strip();
    }

    public int getAge()
    {
        // return this.age;
        return calcAge();
    }

    /*public void setAge(int age)
    {
        if (age < 0 || age > 150) {
            throw new IllegalArgumentException("invalid age value.");
        }

        this.age = age;
    }*/

    public void setBirthYear(int year)
    {
        this.birthYear = year;
    }

    private int calcAge()
    {
        return this.currentYear - this.birthYear;
    }

    public String[] getHobbies()
    {
        return this.hobbies;
    }

    // 可变参数无法保证传入的是 null 或者 0 个参数
    public void setHobbies(String... hobbies)
    {
        this.hobbies = hobbies;
    }
}

public class Main {
    public static void main(String[] args) {
        String name = "xiaoqi";
        int birthYear = 2001;
        String[] hobbies = {"football", "basketball"};
        // 创建 Perspon 的实例
        // Person qi 是声明变量 qi 的类型（引用变量），new Person 则是创建实例 instance
        Person qi = new Person(2019);
        qi.setName(name); // 不影响 name 变量，因为基本类型参数的传递只是值的复制
        qi.setBirthYear(birthYear);
        qi.setHobbies(hobbies);
        System.out.println(qi.getName() + "年年" + qi.getAge() + "岁");
        // 引用类型参数的传递，变量和参数变量指向同一个对象，修改会互相影响
        System.out.println("hobbies:" + Arrays.toString(qi.getHobbies())); // hobbies:[football, basketball]
        hobbies[0] = "handball";
        System.out.println("hobbies:" + Arrays.toString(qi.getHobbies())); // hobbies:[handball, basketball]
    }
}
```

面向对象的特征：封装、继承、多态。  

封装是指隐藏对象的属性和实现细节，仅对外提供公共访问方式。封装的好处是隐藏实现细节，提供公共的访问方式；提高了代码的复用性和安全性。   
封装的原则是将不需要对外提供的内容都隐藏起来；把属性隐藏，提供公共方法对其访问。  

继承的意义是多个类中存在相同属性和行为时，将这些内容抽取到单独一个类中，那么多个类无需再定义这些属性和行为，只要继承那个类即可。单独的这个类称为父类，基类或者超类；这多个类可以称为子类或者派生类。继承的好处是提高了代码的复用性、维护性，让类与类之间产生了关系，是多态的前提。  
Java 只支持单继承，不支持多继承（即一个类只能有一个父类），可以多层继承。子类只能继承父类所有非私有的成员(成员方法和成员变量)，子类不能继承父类的构造方法，但是可以通过 super 关键字去访问父类构造方法。  
继承的原则是类与类之间存在父子、从属的关系（is）。
```java
// 子类不能访问父类的 private 属性或者 private 方法（只能访问 protected 以下的）
class Man extends Person {
    // 添加 Person 的额外属性
    private int strength;

    public Man()
    {
        // 调用父类的构造方法
        // 任何 class 的构造方法，第一行语句必须是调用父类的构造方法
        // 如果没有明确调用父类的构造方法，编译器会自动补上 super()  
        // 如果父类的构造方法没有无参数的声明将编译失败
        super(2019);
    }

    public int getStrength()
    {
        return this.strength;
    }

    public void setStrength(int strength)
    {
        this.strength = strength;
    }

    // 多态的体现，不同继承子类展示不同处理
    public void setSex()
    {
        // 可以访问父类的 protected 以下的属性
        // super 关键字表示父类，这里也可以直接用 this.sex 或 sex
        super.sex = "male";
    }
}
```

多态是指某一个事物，在不同时刻表现出来的不同状态，比如水在不同时刻的状态。多态的前提是有继承关系，有方法重写，有父类引用指向子类对象。多态的好处是提高程序的维护性、扩展性（like），但是也有弊端，比如不能访问子类特有的功能（需要使用多态的转型）。  
继承体系中，向上转型是从子到父，父类引用指向子类对象；而向下转型是从父到子，父类引用转为子类对象。转型的意义在于多态中访问不同层级的属性和方法。  
```java
class Woman extends Person
{
    public Woman()
    {
        super(2019);
    }

    // 多态的体现，不同继承子类展示不同处理
    public void setSex()
    {
        this.sex = "female";
    }

    // 方法覆写
    // 方法覆写的特点是返回值、参数保持一致，不然会编译报错
    @Override
    public int getAge()
    {
        return 18;
    }
}
// 向上转型（upcasting）
// 继承树：Woman > Person > Object
Woman wm1 = new Woman();
Person ps1 = wm1;

// 向下转型（downcasting）
// 向下转型会出错，报 ClassCastException
Person ps2 = new Woman();
Woman wm2 = (Woman)ps2;
// 判断是否为某类的实例
System.out.println(wm2 instanceof Woman); // true
```

### Java 抽象类
通过 abstract 定义的方法是抽象方法，它只有定义，没有实现。抽象方法定义了子类必须实现的接口规范；定义了抽象方法的 class 必须被定义为抽象类，从抽象类继承的子类必须实现抽象方法（如果不实现抽象方法，则该子类仍是一个抽象类）。  
面向抽象编程使得调用者只关心抽象方法的定义，不关心子类的具体实现。  

使用 abstract 关键字修饰的方法叫做抽象方法，抽象方法仅有声明没有方法体。  

Java 抽象的特点：  
抽象类和抽象方法必须用 abstract 关键字修饰；  
抽象类不一定有抽象方法，有抽象方法的类一定是抽象类；   
抽象类不能实例化（按照多态的方式，由具体的子类实例化，简称抽象类多态）；  
抽象类的子类要么是抽象类，要么重写抽象类中的所有抽象方法；    
抽象类可以有构造方法，用于子类访问父类数据的初始化；  
抽象方法不能为 private、static、final 等关键字修饰；  
抽象方法用于限定子类必须完成某些动作，非抽象方法用于提高代码复用性（共性）。  
```java
abstract class Animal 
{
    public abstract void run();
}

class Person extends Animal 
{
    public void run()
    {
        System.out.println("I am a person who can run.");
    }
}
```

### Java 接口
所谓接口（interface），就是比抽象类还要抽象的纯抽象接口，因为它连字段都不能有。  
在 Java 语言设计中，接口不是类，而是对类的一组需求描述，这些类必须要遵循接口描述的统一格式进行定义。  

Java 接口的特点：  
接口用关键字 interface 表示，类实现接口用 implements 表示；  
接口不能实例化（按照多态的方式，由具体的子类实例化，简称接口多态）；  
接口的成员变量只能是常量，默认修饰符 public static final；  
接口的成员方法只能是抽象方法，默认修饰符 public abstract；  
接口没有构造方法，因为接口主要是扩展功能的，而没有具体存在；  
类和接口可以单实现，也可以多实现，还可以在继承一个类的同时实现多个接口（扩展）；  
接口和接口可以存在继承关系，既可以单继承也可以多继承；  
接口可以定义 default 方法，目的是子类覆写即可全部子类使用。  
```java
interface God
{
    String getName();
    default void run() {
        System.out.println(getName() + " run");
    }
}

class Person implements God 
{
    private String name;

    public Person(String name)
    {
        this.name = name;
    }

    public String getName()
    {
        return this.name;
    }
}

God qi = new Person("Xiaoqi");
qi.run(); // Xiaoqi run
```

Java 8 中接口的改动：  
1、JDK 8 之前接口不能有方法体，JDK 8 之后接口中增加了 default 方法和 static 方法，可以有方法体  
```java
interface IAnimal {
    static void printSex() {
        System.out.println("Male Dog");
    }
    default void printAge() {
        System.out.println("18");
    }
}
```
2、接口中的静态变量会被继承  
3、新增函数式接口  
函数式接口（Function Interface）是一个特殊的接口，使用 \@FunctionInterface 注解声明，定义这种接口可以使用 Lambda 表达式直接调用。  
```java
@FunctionalInterface
interface IAnimal {
    static String animalName = "Animal Name";

    static void printSex() {
        System.out.println("Male Dog");
    }

    default void printAge() {
        System.out.println("18");
    }

    void sayHi(String name);
}

class FunctionInterfaceTest {
    public static void main(String[] args) {
        IAnimal animal = name -> System.out.println(name);
        animal.sayHi("WangWang");
    }
}
```

`注意：使用 @FunctionInterface 声明的函数式接口，抽象方法必须有且仅有一个，但可以包含其他非抽象方法。`

### Java 内部类
在一个类中定义了另一个类，则将定义在类中的那个类称之为成员内部类，成员内部类也是最普通的内部类。  
```java
class InnerTest {
    public static void main(String[] args) {
        Outer out = new Outer();
        // 创建成员内部类：Outer.Inner inner = new Outer().new Inner();
        Outer.Inner inner = out.new Inner();
        inner.sayHi();
    }
}

class Outer {
    private String name = "OuterClass";
    
    public Outer() {
        System.out.println("Outer Class.");
        // 外部类访问内部类
        System.out.println(new Inner().name);
    }

    public void sayHi() {
        System.out.println("Hi, Outer.");
    }

    class Inner {
        String name = "InnerClass";

        public void sayHi() {
            System.out.println("Hi, Inner.");
            // 内部类访问外部类
            Outer.this.sayHi();
        }
    }
}
```

在一个类中定义了另一个 static 类，则将定义在类中的那个 static 类称之为静态成员内部类。  
静态成员内部类也就是给内部成员类加上 static 修饰符。  
使用静态内部类的好处是作用域不会扩散到包外（可以通过 “外部类.内部类” 的方式直接访问，内部类可以访问外部类中的所有静态属性和方法）。  
```java
class OuterClass {
    public OuterClass() {
        System.out.println("OuterClass Init.");
    }
    protected static class InnerClass {
        public void sayHi() {
            System.out.println("Hi, InnerClass.");
        }
    }
}
class InnerClassTest {
    public static void main(String[] args) {
        // 与内部成员类的创建方式 new Outer().new Inner() 不同，静态成员内部类使用 new OuterClass.InnerClass() 
        // 注意：不能从静态成员内部类中访问非静态外部类对象
        OuterClass.InnerClass innerClass = new OuterClass.InnerClass();
        innerClass.sayHi();
    }
}
```

一个类定义在另一个类的局部（方法或者任意作用域），这个类就称之为局部内部类。  
局部内部类特点：  

- 局部内部类不能使用任何访问修饰符；
- 局部类如果在方法中，可以直接使用方法中的变量，不需要通过 OutClass.this.xxx 的方式获得。

```java
class OutClass {
    public void sayHi() {
        class InnerClass {
            InnerClass(String name) {
                System.out.println("InnerClass:" + name);
            }
        }
        System.out.println(new InnerClass("Three"));
        System.out.println("Hi, OutClass");
    }
}
class OutTest {
    public static void main(String[] args) {
        new OutClass().sayHi();
    }
}
```

没有名字的内部类就叫做匿名内部类。  
匿名内部类特点：  

- 匿名内部类必须继承一个父类或者实现一个接口
- 匿名内部类不能定义任何静态成员和方法
- 匿名内部类中的方法不能是抽象的

```java
interface AnonymityOuter {
    void hi();
}
class AnonymityTest {
    public static void main(String[] args) {
        AnonymityOuter anonymityOuter = new AnonymityOuter() {
            @Override
            public void hi() {
                System.out.println("Hi, AnonymityOuter.");
            }
        };
        anonymityOuter.hi();
    }
}
```

枚举类是 JDK 1.5 引入的新特性，使用关键字 “enum” 声明。枚举功能虽小，却非常实用，大大方便了程序的开发者。  
```java
enum ColorEnum {
    RED,
    BLUE,
    YELLOW,
    GREEN
}
class EnumTest {
    public static void main(String[] args) {
        ColorEnum color = ColorEnum.GREEN;
        switch (color) {
            case RED:
                System.out.println("Red");
                break;
            case BLUE:
                System.out.println("Blue");
                break;
            case YELLOW:
                System.out.println("Yellow");
                break;
            case GREEN:
                System.out.println("Green");
                break;
            default:
                break;
        }
    }
}
```

### Java 匿名类
Java 匿名类相当于在定义类的同时再新建这个类的实例，既是匿名，自然无法在别的地方使用这个类。  
Java 匿名类放在 class 中使用叫做匿名内部类，它能访问外层 Class 里面的字段，但是不能访问外层方法中的本地变量（除非该变量是 final 声明），内部类的名称和外面能访问的名称相同，则会把外部名称覆盖。  
Java 匿名类不能定义静态初始化代码块，也不能定义接口甚至定义构造方法。  
```java
// Runnable 是一个接口，没有构造函数
Runnable demo = new Runnable() 
{  
    @Override
    public void run() 
    {  
        System.out.println("this is a anonymous class");  
    }  
}; 
```

内部类其实是一种语法糖，可以让你更优雅地设计你的程序结构。
```java
public class Outer 
{  
  public class Inner 
  {

  }  
} 

// 编译之后
public class Outer 
{  
  public static class Inner 
  {  
    private final Outer parent; 

    public Inner(Outer parent) 
    {  
      this.parent = parent;  
    }  
  }  
} 
```

### Java 面向对象的一些关键字及方法
super 关键字：子类调用父类的属性或方法。

final 关键字：一个父类不允许子类对它的某个方法进行覆写，如果类声明 final 则表示断子绝孙。  
final 的特点是修饰的类不能被继承，修饰的方法不能被 Override，修饰的字段属性在初始化后不能被修改。  
对于一个 final 变量，如果是基本数据类型的变量，则其数值一旦在初始化之后便不能更改；如果是引用类型的变量，则在对其初始化之后便不能再让其指向另一个对象。  
使用 final 方法的原因有两个。第一个原因是把方法锁定，以防任何继承类修改它的含义；第二个原因是效率。在早期的 Java 实现版本中，会将 final 方法转为内嵌调用。但是如果方法过于庞大，可能看不到内嵌调用带来的任何性能提升（现在的 Java 版本已经不需要使用 final 方法进行这些优化了），类中所有的 private 方法都隐式地指定为 final。

static 关键字：静态修饰，可以修饰成员变量和成员方法，经常用于工具类和辅助方法。  
static 的特点是随着类的加载而加载，优先于对象存在，被类的所有对象共享，可以通过类名调用也可以通过对象调用。  
`注意：在静态方法中是没有 this 关键字的；静态方法只能访问静态的成员变量和静态的成员方法(其他属性和方法可以通过实例获得)；不推荐使用实例访问类的静态属性或方法，即便可以这么做`

finalize() 是 Object 的 protected 方法，子类可以覆盖该方法以实现资源清理工作，GC 在回收对象之前调用该方法。不建议用 finalize 方法完成 “非内存资源” 的清理工作，但建议用于清理本地对象（通过 JNI 创建的对象）、作为确保某些非内存资源（如 Socket、文件等）释放的一个补充。  
finalize() 的执行过程：当对象变成 (GC Roots) 不可达时，GC 会判断该对象是否覆盖了 finalize 方法，若未覆盖则直接将其回收。否则，若对象未执行过 finalize 方法，将其放入 F-Queue 队列，由一个低优先级线程执行该队列中对象的 finalize 方法。执行 finalize 方法完毕后，GC 会再次判断该对象是否可达，若不可达，则进行回收，否则对象 “复活”。

### Java 克隆
使用等号复制时，对于值类型来说，彼此之间的修改操作是相对独立的；对于引用类型来说，因为复制的是引用对象的内存地址，所以修改其中一个值，另一个值也会跟着变化。  
为了防止这种问题的发生，就要使用对象克隆来解决引用类型复制的问题。  

克隆的好处：  
1、使用方便：假如要复制一个对象，但这个对象中的部分属性已经被修改过了，如果不使用克隆的话，需要给属性手动赋值，相比克隆而已麻烦很多；  
2、性能高：查看 clone 方法可以知道，它是 native 方法，native 方法是原生函数，使用操作系统底层的语言实现的，因此执行效率更高；  
3、隔离性：克隆可以确保对象操作时相互隔离。  

克隆分浅克隆和深克隆。  
> 浅克隆：只会复制对象的值类型，而不会复制对象的引用类型；  
> 深克隆：复制整个对象，包含值类型和引用类型。  

实现浅克隆：克隆的对象实现 Cloneable 接口，并重写 clone() 方法（虽然所有类都是 Object 的子类，但因为 Object 中的 clone() 方法被声明为 protected 访问级别，所以非 java.lang 包下的其他类是不能直接使用的）。  
```java
class CloneTest 
{
    public static void main(String[] args) throws CloneNotSupportedException 
    {
        Dog dog = new Dog();
        dog.name = "旺财";
        dog.age = 5;
        // 克隆
        Dog dog3 = (Dog) dog.clone();
        dog3.name = "小白";
        dog3.age = 2;
        System.out.println(dog.name + "，" + dog.age + "岁");
        System.out.println(dog3.name + "，" + dog3.age + "岁");
    }
}

class Dog implements Cloneable 
{
    public String name;
    public int age;

    @Override
    protected Object clone() throws CloneNotSupportedException 
    {
        return super.clone();
    }
}
```

实现深克隆一般有两种方式：通过序列化实现深克隆（Java 原生序列化、JSON 序列化、Hessian 序列化）；所有引用类型都实现克隆，从而实现深克隆。  
1、序列化实现深克隆的原理：先将原对象序列化到内存的字节流中，再从字节流中反序列化出刚刚存储的对象，这个新对象和原对象就不存在任何地址上的共享，从而实现深克隆。  
```java
class CloneTest 
{
    public static void main(String[] args) throws CloneNotSupportedException 
    {
        BirdChild birdChild = new BirdChild();
        birdChild.name = "小小鸟";
        Bird bird = new Bird();
        bird.name = "小鸟";
        bird.birdChild = birdChild;
        // 使用序列化克隆对象
        Bird bird2 = CloneUtils.clone(bird);
        bird2.name = "黄雀";
        bird2.birdChild.name = "小黄雀";
        System.out.println("bird name:" + bird.name);
        System.out.println("bird child name:" + bird.birdChild.name);
        System.out.println("bird name 2:" + bird2.name);
        System.out.println("bird child name 2:" + bird2.birdChild.name);
    }
}

class CloneUtils 
{
    public static <T extends Serializable> T clone(T obj) 
    {
        T cloneObj = null;
        try {
            //写入字节流
            ByteArrayOutputStream bo = new ByteArrayOutputStream();
            ObjectOutputStream oos = new ObjectOutputStream(bo);
            oos.writeObject(obj);
            oos.close();
            //分配内存,写入原始对象,生成新对象
            ByteArrayInputStream bi = new ByteArrayInputStream(bo.toByteArray());//获取上面的输出字节流
            ObjectInputStream oi = new ObjectInputStream(bi);
            //返回生成的新对象
            cloneObj = (T) oi.readObject();
            oi.close();
        } catch (Exception e) {
            e.printStackTrace();
        }
        return cloneObj;
    }
}
```

2、所有引用类型都实现克隆  
```java
class SerializableTest 
{
    public static void main(String[] args) throws IOException, ClassNotFoundException 
    {
    ParrotChild parrotChild = new ParrotChild();
        parrotChild.name = "小鹦鹉";
        Parrot parrot = new Parrot();
        parrot.name = "大鹦鹉";
        parrot.parrotChild = parrotChild;
        // 克隆
        Parrot parrot2 = (Parrot) parrot.clone();
        parrot2.name = "老鹦鹉";
        parrot2.parrotChild.name = "少鹦鹉";
        System.out.println("parrot name:" + parrot.name);
        System.out.println("parrot child name:" + parrot.parrotChild.name);
        System.out.println("parrot name 2:" + parrot2.name);
        System.out.println("parrot child name 2:" + parrot2.parrotChild.name);
    }
}

class Parrot implements Cloneable 
{
    public String name;
    public ParrotChild parrotChild;

    @Override
    protected Object clone() throws CloneNotSupportedException 
    {
        Parrot bird = (Parrot) super.clone();
        bird.parrotChild = (ParrotChild) parrotChild.clone();
        return bird;
    }
}

class ParrotChild implements Cloneable 
{
    public String name;

    @Override
    protected Object clone() throws CloneNotSupportedException 
    {
        return super.clone();
    }
}
```

### Java 序列化与反序列化
序列化是一种对象持久化的手段，普遍应用在网络传输、RMI 等场景中。  
Java 对象序列化，在保存对象时会把其状态保存为一组字节，反序列化时再将这些字节组装成对象。  
```java
import java.io.Serializable;
import java.io.ObjectInputStream;
import java.io.IOException;
import org.apache.commons.io.FileUtils;
import org.apache.commons.io.IOUtils;

// Person 类实现 java.io.Serializable 接口，它可以被序列化
class Person implements Serializable
{
    private String name;
    private transient String gender; // transient 阻止该变量被序列化

    public String getName() 
    {
        return name;
    }

    public void setName(String name) 
    {
        this.name = name;
    }

    @Override
    public String toString()
    {
        return "Person name：" + this.name;
    }

    // 自定义序列化策略
    // @Override
    // private void readObject(java.io.ObjectInputStream s)
    // @Override
    // private void writeObject(java.io.ObjectOutputStream s)
}

Person p = new Person();
p.setName("vikey");
System.out.println(p);

// 通过 ObjectOutputStream  进行序列化
ObjectOutputStream oos = null;
try {
    oos = new ObjectOutputStream(new FileOutputStream("tempFile"));
    oos.writeObject(p);
} catch (IOException e) {
    e.printStackTrace();
} finally {
    IOUtils.closeQuietly(oos);
}

// 通过 ObjectInputStream 进行反序列化
File file = new File("tempFile");
ObjectInputStream ois = null;
try {
    ois = new ObjectInputStream(new FileInputStream(file));
    Person newPerson = (Person)ois.readObject();
    System.out.println(newPerson);
} catch (IOException e) {
    e.printStackTrace();
} catch (ClassNotFoundException e) {
    e.printStackTrace();
} finally {
    IOUtils.closeQuietly(ois);
    try {
        FileUtils.forceDelete(file);
    } catch (IOException e) {
        e.printStackTrace();
    }
}
```

### Java 包
在 Java 中，使用 package 来解决名字冲突。  
Java 定义了一种名字空间，称之为包（package）。一个类总是属于某个包，类名（比如 Person）只是一个简写，真正的完整类名是"包名.类名"。没有定义包名的 class 使用默认包，不推荐这样做！推荐所有 Java 文件对应的目录层次要和包的层次一致。    

包作用域：位于同一个包的类，可以访问包作用域的字段和方法，不用 public、protected、private 修饰的字段和方法就是包作用域。
```java
package qi;

public class Person
{
    void sayHello() 
    {
        System.out.println("Hello");
    }
}

public class Main 
{
    public static void main(String[] args) 
    {
        Person p = new Person();
        p.hello(); // Main 和 Person 在同一个包，可以直接调用
    }
}
```

包的导入（import）  
在 class 中引用其他 class，需要使用包的导入。
```java
// 方式一：完整类名
qi.Person p = new qi.Person();
// 方式二：import（推荐使用）
import qi.Person; // 如果是导入完整的包 qi.Person.* 但是不建议使用
Person p = new Person();
// 方式三：import static（包的静态导入，很少使用）
import static java.lang.Math.abs;
System.out.println(abs(-12));
```

在编写 class 的时候，编译器会自动帮我们做两个 import 动作：  
1、默认自动 import 当前 package 的其他 class  
2、默认自动 import java.lang.\*   

所以，要注意不要和 java.lang 包的类重名（如 String、System、Runtime...)，也不要和 JDK 常用类重名（如 java.util.List, java.text.Format, java.math.BigInteger...）。  

### jar 包
jar 包能把目录打一个 jar 包，变成一个文件以便管理。  
jar 包实际上就是一个 zip 格式的压缩文件，而 jar 包相当于目录。  
```bash
# 执行一个 jar 包的 class，把 jar 包放到 classpath 即可
# JVM 会自动在 hello.jar 文件里去搜索 abc.xyz.Hello
java -cp ./hello.jar abc.xyz.Hello
```

jar 包的创建，可以使用 zip 的压缩方式，将 src 目录（不包含 bin 目录）打包即可。  
jar 包可以包含一个特殊的纯文本文件 /META-INF/MANIFEST.MF，可以指定 Main-Class 和其它信息。  
```bash
# JVM 会自动读取 MANIFEST.MF 文件
# 如果存在 Main-Class，我们就不必在命令行指定启动的类名
java -jar hello.jar
```

jar 包还可以包含其它 jar 包，需要在 MANIFEST.MF 文件里配置 classpath。  
在大型项目中，不可能手动编写 MANIFEST.MF 文件再手动创建 zip 包，Java 社区提供了大量的开源构建工具，比如 Maven 就可以非常方便地创建 jar 包。  

### Java 模块
jar 包只是用于存放 class 的容器，它并不关心 class 之间的依赖。  
为了解决依赖问题，自带 “依赖关系” 的 class 容器 ———— 模块诞生了。  

为了表明 Java 模块化的决心，从 Java 9 开始，原有的 Java 标准库已经由一个单一巨大的 rt.jar 分拆成了几十个模块，这些模块以 .jmod 扩展名标识，可以在 $JAVA_HOME/jmods 目录下找到它们（java.base.jmod java.compiler.jmod ...）。  

模块之间的依赖关系被写入到模块内的 module-info.class 文件。所有的模块都直接或间接地依赖 java.base 模块，只有 java.base 模块不依赖任何模块，它可以被看作是 “根模块”，好比所有的类都是从 Object 直接或间接继承而来。  
把一堆 class 封装为 jar 仅仅是一个打包的过程，而把一堆 class 封装为模块则不但需要打包，还需要写入依赖关系，并且还可以包含二进制代码（通常是 JNI 扩展）。而且，模块还支持多版本，即在同一个模块中可以为不同的 JVM 提供不同的版本。  
```
# bin 目录存放编译后的 class 文件
# src 目录存放源码
# module-info.java 模块的描述文件
demo-module
|- bin
|- build.sh
|- src
----|- com
----|---|- demo
----|---|----|- Main.java
----|- module-info.java

# module-info.java
module demo.person {
    requires java.base; // 可以省略，任何模块都会自动引入
    requires java.xml; // 引入 xml 模块
}

# Main.java
package com.demo;

import javax.xml.XMLConstants;

public class Main 
{
    public static void main(String[] args) 
    {
        System.out.println(XMLConstants.XML_NS_PREFIX);
    }
}

# 在 demo-module 目录下编译所有的 .java 文件并存放到 bin 目录
javac -d bin src/module-info.java src/com/demo/*.java
# 把 bin 目录下的所有 class 文件先打包成 jar
jar --create --file demo.jar --main-class com.demo.Main -C bin .
# 继续使用 JDK 自带的 jmod 命令把一个 jar 包转换成模块
jmod create --class-path demo.jar demo.jmod

# 上面，我们已经成功得到一个模块
# 运行模块，使用 jar 包
java --module-path demo.jar --module demo.person
# jmod 可以用来打包 JRE
```

打包 JRE  
过去发布一个 Java 应用程序，要运行它必须下载一个完整的 JRE，再运行 jar 包。而完整的 JRE 块头很大，大约有 100 多 M。  
现在，JRE 自身的标准库已经分拆成了模块，只需要带上程序用到的模块（“复制” 一份 JRE，只带上用到的模块），其他的模块就可以被裁剪掉。  
```bash
jlink --module-path demo.jmod --add-modules java.base,java.xml,demo.person --output jre/
# 运行 jre
jre/bin/java --module demo.person
```

其他说明  
模块进一步隔离了代码的访问权限。  
比如模块 java.xml 的一个类 javax.xml.XMLConstants，我们之所以能直接使用这个类，是因为模块 java.xml 的 module-info.java 中声明了若干导出。  
比如外部代码需要访问 demo.person 中 com.demo.Main，我们就需要把它在自己的 module-info.java 文件中声明 `exports com.demo`。
```java
module java.xml {
    exports java.xml;
    exports javax.xml.catalog;
    exports javax.xml.datatype;
    ...
}
```
