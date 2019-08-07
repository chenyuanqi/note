
### Java Object
所有的 class 最终都继承自 Object，Object 定义了几个重要的方法：  
> toString() 把 instance 输出为 String  
> equals() 判断两个 instance 是否逻辑相等  
> hashCode() 计算一个 instance 的哈希值  

### Java 面向对象
世间万物都可以看成一个对象，每个物体包括动态的行为和静态的属性，这些就构成了一个对象。类是对象的抽象，对象是类的具体，类是对象的模板，对象是类的实例。  

面向对象是基于面向过程的编程思想。  
面向对象更符合我们的思维习惯，可以把复杂的事情简单化，角色从执行者变成指挥者。  
在 Java 中，创建一个类 Person，包含多个字段（field），初步认识 Java 的面向对象。  

Java 内建的访问权限包括 public、protected、private 和 package 权限；  
在方法内部定义的变量是局部变量，局部变量的作用域从变量声明开始，到一个块结束；  
final 修饰符不是访问权限，它可以修饰 class、field 和 method。  
> public: 字段属性和方法可以被其他包、实例和子类访问，以及子类的子类  
> protected: 字段属性和方法可以被实例和子类访问，以及子类的子类，不能被其他包访问  
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

Java 抽象的特点：  
抽象类和抽象方法必须用 abstract 关键字修饰；  
抽象类不一定有抽象方法，有抽象方法的类一定是抽象类；   
抽象类不能实例化（按照多态的方式，由具体的子类实例化，简称抽象类多态）；  
抽象类的子类要么是抽象类，要么重写抽象类中的所有抽象方法；    
抽象类可以有构造方法，用于子类访问父类数据的初始化；  
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

### Java 面向对象的一些关键字
super 关键字：子类调用父类的属性或方法。

final 关键字：一个父类不允许子类对它的某个方法进行覆写，如果类声明 final 则表示断子绝孙。  
final 的特点是修饰的类不能被继承，修饰的方法不能被 Override，修饰的字段属性在初始化后不能被修改。

static 关键字：静态修饰，可以修饰成员变量和成员方法，经常用于工具类和辅助方法。  
static 的特点是随着类的加载而加载，优先于对象存在，被类的所有对象共享，可以通过类名调用也可以通过对象调用。  
`注意：在静态方法中是没有 this 关键字的；静态方法只能访问静态的成员变量和静态的成员方法；不推荐使用实例访问类的静态属性或方法，即便可以这么做`

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
// 方式三：import static（很少使用）
```

在编写 class 的时候，编译器会自动帮我们做两个 import 动作：  
1、默认自动 import 当前 package 的其他 class  
2、默认自动 import java.lang.\*   

所以，要注意不要和 java.lang 包的类重名（如 String、System、Runtime...)，也不要和 JDK 常用类重名（如 java.util.List, java.text.Format, java.math.BigInteger...）。  

