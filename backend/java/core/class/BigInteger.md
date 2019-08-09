
### BigInteger
在 Java 中，由 CPU 原生提供的整型最大范围是 64 位 long 型整数。使用 long 型整数可以直接通过 CPU 指令进行计算，速度非常快。  
但是，整数范围如果超过了 long 型，就只能使用 BigInteger 了。  

java.math.BigInteger 就是用来表示任意大小的整数，但是速度比较慢。  
BigInteger 内部用一个 int[] 数组来模拟一个非常大的整数。BigInteger 和 Integer、Long 一样，也是不可变类，并且也继承自 Number 类。  
```java
import java.math.BigInteger;

BigInteger bi = new BigInteger("1234567890");
System.out.println(bi.pow(5)); // 2867971860299718107233761438093672048294900000

// 加法运算
BigInteger i1 = new BigInteger("1234567890");
BigInteger i2 = new BigInteger("12345678901234567890");
BigInteger sum = i1.add(i2); // 12345678902469135780

// 把 BigInteger 转换成 long 型
BigInteger i = new BigInteger("123456789000");
System.out.println(i.longValue()); // 123456789000
// 如果需要准确地转换成基本类型，可以使用 intValueExact()、longValueExact() 等方法
// 在转换时如果超出范围，将直接抛出 ArithmeticException 异常
System.out.println(i.multiply(i).longValueExact()); // java.lang.ArithmeticException: BigInteger out of long range 超出范围异常

// Number 定义了转换为基本类型的方法，BigInteger 也可以使用
// 如果 BigInteger 表示的范围超过了基本类型的范围，转换时将丢失高位信息，即结果不一定是准确的
BigInteger i = new BigInteger("123");
System.out.println(i.byteValue()); // 123
System.out.println(i.shortValue()); // 123
System.out.println(i.intValue()); // 123
System.out.println(i.longValue()); // 123
System.out.println(i.floatValue()); // 123.0
System.out.println(i.doubleValue()); // 123.0
```
