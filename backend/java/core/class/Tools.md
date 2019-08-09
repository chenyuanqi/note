
### 工具类
Java 的核心库提供了大量的现成的类供我们使用。  

#### Math
Math 类就是用来进行数学计算的，它提供了大量的静态方法来便于我们实现数学计算。  
```java
// 数学常量
double pi = Math.PI; // 3.14159...
double e = Math.E; // 2.7182818...

// 求绝对值
Math.abs(-100); // 100
Math.abs(-7.8); // 7.8

// 取最大或最小值
Math.max(100, 99); // 100
Math.min(1.2, 2.3); // 1.2

// 计算 power
Math.pow(2, 10); // 2 的 10 次方 = 1024
// 计算√x
Math.sqrt(2); // 1.414...
// 计算 e 的 x 次方
Math.exp(2); // 7.389...
// 计算以 e 为底的对数
Math.log(4); // 1.386...
// 计算以 10 为底的对数
Math.log10(100); // 2

// 三角函数
Math.sin(3.14); // 0.00159...
Math.sin(Math.PI / 6); // sin(π/6) = 0.5
Math.cos(3.14); // -0.9999...
Math.tan(3.14); // -0.0015...
Math.asin(1.0); // 1.57079...
Math.acos(1.0); // 0.0

// 生成一个随机数 x，x 的范围是 0 <= x < 1
Math.random(); // 0.53907... 每次都不一样
```
StrictMath 保证所有平台计算结果都是完全相同的，而 Math 会尽量针对平台优化计算速度。

#### Random 和 SecureRandom
Random 用来创建伪随机数。  
所谓伪随机数，是指只要给定一个初始的种子，产生的随机数序列是完全一样的。  
```java
// 如果不给定种子，就使用系统当前时间戳作为种子
Random r = new Random();
r.nextInt(); // 2071575453,每次都不一样
r.nextInt(10); // 5,生成一个 [0,10) 之间的 int
r.nextLong(); // 8811649292570369305,每次都不一样
r.nextFloat(); // 0.54335...生成一个[0,1)之间的float
r.nextDouble(); // 0.3716...生成一个[0,1)之间的double

// 指定种子，就会得到完全确定的随机数序列
Random r = new Random(12345);
for (int i = 0; i < 10; i++) {
    System.out.println(r.nextInt(100));
}
```

有伪随机数，就有真随机数。  
实际上真正的真随机数只能通过量子力学原理来获取，而我们想要的是一个不可预测的安全的随机数，SecureRandom 就是用来创建安全的随机数的。  
SecureRandom 无法指定种子，它使用 RNG（random number generator）算法。
```java
SecureRandom sr = new SecureRandom();
System.out.println(sr.nextInt(100));

// JDK 的 SecureRandom 实际上有多种不同的底层实现
// 有的使用安全随机种子加上伪随机数算法来产生安全的随机数，有的使用真正的随机数生成器
// 实际使用的时候，可以优先获取高强度的安全随机数生成器，如果没有提供，再使用普通等级的安全随机数生成器
import java.util.Arrays;
import java.security.SecureRandom;
import java.security.NoSuchAlgorithmException;
SecureRandom sr = null;
try {
    sr = SecureRandom.getInstanceStrong(); // 获取高强度安全随机数生成器
} catch (NoSuchAlgorithmException e) {
    sr = new SecureRandom(); // 获取普通的安全随机数生成器
}
byte[] buffer = new byte[16];
sr.nextBytes(buffer); // 用安全随机数填充 buffer
System.out.println(Arrays.toString(buffer));
```
SecureRandom 的安全性是通过操作系统提供的安全的随机种子来生成随机数。这个种子是通过 CPU 的热噪声、读写磁盘的字节、网络流量等各种随机事件产生的 “熵”。  
在密码学中，安全的随机数非常重要。如果使用不安全的伪随机数，所有加密体系都将被攻破。  
