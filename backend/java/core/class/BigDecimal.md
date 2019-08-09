
### BigDecimal 
BigDecimal 可以表示一个任意大小且精度完全准确的浮点数，常用于财务计算。  
BigDecimal 也是从 Number 继承的，也是不可变对象。实际上一个 BigDecimal 是通过一个 BigInteger 和一个 scale 来表示的，即 BigInteger 表示一个完整的整数，而 scale 表示小数位数。  
```java
public class BigDecimal extends Number implements Comparable<BigDecimal> {
    private final BigInteger intVal;
    private final int scale;
}
```

`比较 BigDecimal 的值是否相等，必须使用 compareTo() 而不能使用 equals()。`
  
```java
import java.math.BigDecimal;
import java.math.RoundingMode;

BigDecimal bd = new BigDecimal("123.4567");
System.out.println(bd.multiply(bd)); // 15241.55677489

// BigDecimal 用 scale() 表示小数位数
BigDecimal d1 = new BigDecimal("123.45");
BigDecimal d2 = new BigDecimal("123.4500");
BigDecimal d3 = d2.stripTrailingZeros(); // 去掉了小数末尾 0
BigDecimal d4 = new BigDecimal("1234500");
BigDecimal d5 = d4.stripTrailingZeros(); // 经过处理的整数，小数位变成负数表示整数末尾的 0 个数
System.out.println(d1.scale()); // 2, 两位小数
System.out.println(d2.scale()); // 4
System.out.println(d3.scale()); // 2
System.out.println(d4.scale()); // 0
System.out.println(d5.scale()); // -2

BigDecimal d1 = new BigDecimal("123.456789");
BigDecimal d2 = d1.setScale(4, RoundingMode.HALF_UP); // 四舍五入，123.4568
BigDecimal d3 = d1.setScale(4, RoundingMode.DOWN); // 直接截断，123.4567
System.out.println(d2);
System.out.println(d3);

// 对 BigDecimal 做加、减、乘时，精度不会丢失
// 做除法时，存在无法除尽的情况时，必须指定精度以及如何进行截断
BigDecimal d1 = new BigDecimal("123.456");
BigDecimal d2 = new BigDecimal("23.456789");
BigDecimal d3 = d1.divide(d2, 10, RoundingMode.HALF_UP); // 保留 10 位小数并四舍五入

// 比较两个 BigDecimal 的值是否相等
// 使用 equals() 方法不但要求两个 BigDecimal 的值相等，还要求它们的 scale() 相等
// 使用 compareTo() 方法来比较，它根据两个值的大小分别返回负数、正数和 0，分别表示小于、大于和等于
BigDecimal d1 = new BigDecimal("123.456");
BigDecimal d2 = new BigDecimal("123.45600");
System.out.println(d1.equals(d2)); // false, 因为scale不同
System.out.println(d1.equals(d2.stripTrailingZeros())); // true, 因为 d2 去除尾部 0 后 scal e变为 2
System.out.println(d1.compareTo(d2)); // 0
```
