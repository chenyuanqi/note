
### StringJoiner
用分隔符拼接数组的需求很常见，所以 Java 标准库还提供了一个 StringJoiner 来干这个事。  
```java
import java.util.StringJoiner;

String[] names = {"Bob", "Alice", "Grace"};
var sj = new StringJoiner(", ", "Hello ", "!"); // 第 2、3 个参数是字符串的开始和结尾定义，可以省略
for (String name : names) {
    sj.add(name);
}
System.out.println(sj.toString()); // Hello Bob, Alice, Grace!
```

StringJoiner 内部实际上就是使用了 StringBuilder，所以拼接效率和 StringBuilder 几乎是一模一样的。  

String 还提供了一个静态方法 join()，这个方法在内部使用了 StringJoiner 来拼接字符串，在不需要指定 “开头” 和 “结尾” 的时候，用 String.join() 更方便。  
```java
String[] names = {"Bob", "Alice", "Grace"};
var s = String.join(", ", names);
```
