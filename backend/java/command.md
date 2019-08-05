
### Java 命令行
Java 程序的入口是 main 方法，而 main 方法可以接受一个命令行参数，它是一个 String[] 数组。  

```java
// 当命令行运行 java Main --version 时就会输出 v1.0
public class Main {
    public static void main(String[] args) {
        for (String arg : args) {
            if ("--version".equals(arg)) {
                System.out.println("v1.0");
                break;
            }
        }
    }
}
```



