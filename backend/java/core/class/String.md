
### 字符串 String
在 Java 中，String 是一个引用类型，它本身也是一个 class，CharSequence 是 String 的父类。  
但是，Java 编译器对 String 有特殊处理，即可以直接用 "..." 来表示一个字符串。  
```java
String s1 = "Hello";
// 实际上,字符串在 String 内部是通过一个 char[] 数组表示的
String s2 = new String(new char[] {'H', 'e', 'l', 'l', 'o'});
```

Java 字符串的一个重要特点就是字符串内容不可变，但是引用可以变。这种不可变性是通过内部的 private final char[] 字段，以及没有任何修改 char[] 的方法实现的。  
```java
// 早期 JDK 版本的 String 总是以 char[] 存储
public final class String 
{
    private final char[] value;
    private final int offset;
    private final int count;
}

// 较新的 JDK 版本的 String 则以 byte[] 存储
// 如果 String 仅包含 ASCII 字符，则每个 byte 存储一个字符
// 否则每两个 byte 存储一个字符
public final class String 
{
    private final byte[] value;
    private final byte coder; // 0 = LATIN1, 1 = UTF16
}
```

### 字符串 String 常用方法
**字符串对比**  
两个字符串比较，必须总是使用 equals() 方法；要忽略大小写比较，使用 equalsIgnoreCase() 方法。
```java
String s1 = "hello";
String s2 = "HELLO".toLowerCase();
System.out.println(s1 == s2); // 字符串也可以用 == 比较，但是引用不一致会导致不准确
System.out.println(s1.equals(s2));
```

**字符串检索**  
```java
// 是否包含子串:
// contains() 方法的参数是 CharSequence
"Hello".contains("ll"); // true

// 搜索子串的位置
// 索引从 0 开始
"Hello".indexOf("l"); // 2
"Hello".lastIndexOf("l"); // 3
"Hello".startsWith("He"); // true
"Hello".endsWith("lo"); // true

// 提取子串
"Hello".substring(2); // "llo"
"Hello".substring(2, 4); // "ll"
```

**去除首尾空白字符**
```java
"  \tHello\r\n ".trim(); // "Hello"

// 中文的空格字符也会被去除
"\u3000Hello\u3000".strip(); // "Hello"
" Hello ".stripLeading(); // "Hello "
" Hello ".stripTrailing(); // " Hello"
```

**是否为空和空白字符串**  
```java
"".isEmpty(); // true，因为字符串长度为 0
"  ".isEmpty(); // false，因为字符串长度不为 0
"  \n".isBlank(); // true，因为只包含空白字符
" Hello ".isBlank(); // false，因为包含非空白字符
```

**替换子串**
```java
String s = "hello";
s.replace('l', 'w'); // "hewwo"，所有字符'l'被替换为'w'
s.replace("ll", "~~"); // "he~~o"，所有子串"ll"被替换为"~~"

// 正则替换
String s = "A,,B;C ,D";
s.replaceAll("[\\,\\;\\s]+", ","); // "A,B,C,D"
```

**分割字符串**  
```java
String s = "A,B,C,D";
String[] ss = s.split("\\,"); // {"A", "B", "C", "D"}
```

**拼接字符串**
```java
String[] arr = {"A", "B", "C"};
String s = String.join("***", arr); // "A***B***C"
```

**类型转换**
```java
// 任意基本类型或引用类型转换为字符串
// valueOf() 是一个重载方法
String.valueOf(123); // "123"
String.valueOf(45.67); // "45.67"
String.valueOf(true); // "true"
String.valueOf(new Object()); // 类似java.lang.Object@636be97c

// 把字符串转成其他类型
int n1 = Integer.parseInt("123"); // 123
int n2 = Integer.parseInt("ff", 16); // 按十六进制转换，255
boolean b1 = Boolean.parseBoolean("true"); // true
boolean b2 = Boolean.parseBoolean("FALSE"); // false
// 把字符串对应的系统变量转换为 Integer
Integer.getInteger("java.version"); // 版本号，12

// 转换为 char[]
char[] cs = "Hello".toCharArray(); // String -> char[]
String s = new String(cs); // char[] -> String
cs[0] = 'X'; // 修改不会影响 s，因为 new String() 只是复制了 cs
// 如果在 class 中传递 char[] 则会对 s 产生影响
```

**字符串编码**  
Java 的 String 和 char 在内存中总是以 Unicode 编码表示。
```java
import java.nio.charset.StandardCharsets;

byte[] b1 = "Hello".getBytes(); // 按 ISO8859-1 编码转换，不推荐
byte[] b2 = "Hello".getBytes("UTF-8"); // 按 UTF-8 编码转换
byte[] b2 = "Hello".getBytes("GBK"); // 按 GBK 编码转换
byte[] b3 = "Hello".getBytes(StandardCharsets.UTF_8); // 按 UTF-8 编码转换

byte[] b = "Hello".getBytes("UTF-8");
String s1 = new String(b, "GBK"); // 按GBK转换
String s2 = new String(b, StandardCharsets.UTF_8); // 按UTF-8转换
```

