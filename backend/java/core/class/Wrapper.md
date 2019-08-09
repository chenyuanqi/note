
### Java 装箱拆箱 
我们知道，引用类型可以赋值为 null 表示空，但基本类型不能赋值为 null。  
如果要把基本类型转成引用类型，比如定义 Integer 为 int 的包装类。  
```java
public class Integer 
{
    private final int value;

    public Integer(int value) 
    {
        this.value = value;
    }

    public int intValue() 
    {
        return this.value;
    }
}

// int 和 Integer 互相转换
Integer n = null;
Integer n2 = new Integer(99);
int n3 = n2.intValue();
```

Java 核心库为每种基本类型都提供了对应的包装类型。  

| 基本类型 | 对应的引用类型 |  
| ----: | :---- |  
| boolean | java.lang.Boolean |  
| byte | java.lang.Byte |  
| short | java.lang.Short |  
| int | java.lang.Integer |  
| long| java.lang.Long |  
| float | java.lang.Float |  
| double | java.lang.Double |  
| char | java.lang.Character |  

所有的包装类型都是不变类，所以一旦创建了 Integer 对象，该对象就是不变的。
```java
int i = 100;
// 通过 new 操作符创建 Integer 实例(不推荐使用,会有编译警告):
Integer n1 = new Integer(i);
// 通过静态方法 valueOf(int) 创建 Integer 实例:
Integer n2 = Integer.valueOf(i);
// 通过静态方法 valueOf(String) 创建 Integer 实例:
Integer n3 = Integer.valueOf("100");
System.out.println(n3.intValue());

// 两个 Integer 实例进行比较，必须使用 equals()
System.out.println("n3.equals(n2): " + n3.equals(n2));

// 进制转换
int x1 = Integer.parseInt("100"); // 100
int x2 = Integer.parseInt("100", 16); // 256,因为按 16 进制解析
// 把整数格式化为指定进制的字符串
System.out.println(Integer.toString(100)); // "100",表示为 10 进制
System.out.println(Integer.toString(100, 36)); // "2s",表示为 36 进制
System.out.println(Integer.toHexString(100)); // "64",表示为 16 进制
System.out.println(Integer.toOctalString(100)); // "144",表示为 8 进制
System.out.println(Integer.toBinaryString(100)); // "1100100",表示为 2 进制

// 一些有用的静态变量
// boolean只有两个值true/false，其包装类型只需要引用Boolean提供的静态字段:
Boolean t = Boolean.TRUE;
Boolean f = Boolean.FALSE;
// int可表示的最大/最小值:
int max = Integer.MAX_VALUE; // 2147483647
int min = Integer.MIN_VALUE; // -2147483648
// long类型占用的bit和byte数量:
int sizeOfLong = Long.SIZE; // 64 (bits)
int bytesOfLong = Long.BYTES; // 8 (bytes)

// 通过包装类型获取各种基本类型
// Integer 向上转型为 Number:
Number num = new Integer(999);
// 获取 byte, int, long, float, double:
byte b = num.byteValue();
int n = num.intValue();
long ln = num.longValue();
float f = num.floatValue();
double d = num.doubleValue();

// 无符号整型和有符号整型的转换在 Java 中就需要借助包装类型的静态方法完成
// byte 是有符号整型，范围是 -128~+127，但如果把 byte 看作无符号整型，它的范围就是 0~255 
// 把一个负的 byte 按无符号整型转换为 int：
byte x = -1;
byte y = 127;
System.out.println(Byte.toUnsignedInt(x)); // 255
System.out.println(Byte.toUnsignedInt(y)); // 127
```
Java 编译器可以帮助我们自动在 int 和 Integer 之间转型。  
这种直接把 int 变为 Integer 的赋值写法，称为自动装箱（Auto Boxing）；把 Integer 变为 int 的赋值写法，称为自动拆箱（Auto Unboxing）。  
`注意：自动装箱和自动拆箱只发生在编译阶段，目的是为了少写代码。`  

装箱和拆箱会影响代码的执行效率，因为编译后的 class 代码是严格区分基本类型和引用类型的。并且，自动拆箱执行时可能会报 NullPointerException（如 Integer num = null;）。  
