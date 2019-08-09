
### 枚举
枚举是一种特殊的数据类型，可以把相关的常量分组到一个枚举类型里，并提供比常量更多的方法。    
在 Java 中，定义枚举类是通过关键字 enum，是一个 class。  
> 定义的 enum 类型总是继承自 java.lang.Enum，且无法被继承；  
> 只能定义出 enum 的实例，而无法通过 new 操作符创建 enum 的实例；  
> 定义的每个实例都是引用类型的唯一实例；  
> 可以将 enum 类型用于 switch 语句。  


```java
enum Weekday {
    SUN, MON, TUE, WED, THU, FRI, SAT;
}

// 编译出的 class 大概就像这样
// 继承自 Enum，标记为 final class
public final class Weekday extends Enum { 
    // 每个实例均为全局唯一
    public static final Weekday SUN = new Weekday();
    public static final Weekday MON = new Weekday();
    // private构造方法，确保外部无法调用 new 操作符
    private Weekday() {}
}

// enum 是一个 class，每个枚举的值都是 class 实例
// 获取常量名
Weekday.SUN.name(); // "SUN"
// 获取常量的顺序，从 0 开始计数
Weekday.SUN.ordinal(); // 0

```

使用 enum 枚举的好处：  
1、实现的 enum 常量本身带有类型信息，编译器会自动检查出类型错误；  
2、不可能引用到非枚举的值，因为无法通过编译；  
3、不同类型的枚举不能互相比较或者赋值，因为类型不符  

使用 enum 定义的枚举类是一种引用类型。  
我们知道，引用类型比较，要使用 equals() 方法；但是 enum 类型可以例外。因为 enum 类型的每个常量在 JVM 中只有一个唯一实例，所以可以直接用 == 比较。

```java
// 这样定义枚举更健壮：有序号、说明
enum Weekday {
    MON(1, "星期一"), TUE(2, "星期二"), WED(3, "星期三"), THU(4, "星期四"), FRI(5, "星期五"), SAT(6, "星期六"), SUN(0, "星期日");

    public final int dayValue;
    public final String chinese;

    private Weekday(int dayValue, String chinese) {
        this.dayValue = dayValue;
        this.chinese = chinese;
    }

    @Override
    public String toString() {
        // 默认情况下，对枚举常量调用 toString() 会返回和 name() 一样的字符串
        // 覆写 toString() 的目的是在输出时更有可读性
        return this.chinese;
    }
}

Weekday day = Weekday.SUN;
if (day.dayValue == 6 || day.dayValue == 0) {
    System.out.println("Today is " + day + ". Work at home!");
} else {
    System.out.println("Today is " + day + ". Work at office!");
}

// 枚举类天生具有类型信息和有限个枚举常量，更适合用在 switch 语句中
switch(day) {
    case MON:
    case TUE:
    case WED:
    case THU:
    case FRI:
        System.out.println("Today is " + day + ". Work at office!");
        break;
    case SAT:
    case SUN:
        System.out.println("Today is " + day + ". Work at home!");
        break;
    default:
        throw new RuntimeException("cannot process " + day);
}
```