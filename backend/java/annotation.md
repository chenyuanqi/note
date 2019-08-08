
### Java 注解
Java 程序的一种特殊 “注释”—— 注解（Annotation），注解是放在 Java 源码的类、方法、字段、参数前的一种特殊 “注释”。   

注释会被编译器直接忽略，注解则可以被编译器打包进入 class 文件，因此，注解是一种用作标注的 “元数据”。  
从 JVM 的角度看，注解本身对代码逻辑没有任何影响，如何使用注解完全由工具决定。  

Java 的注解可以分为三类：  
1、由编译器使用的注解，如 \@Override 让编译器检查该方法是否正确地实现了覆写，\@SuppressWarnings 告诉编译器忽略此处代码产生的警告。这类注解不会被编译进入 .class 文件，它们在编译后就被编译器扔掉了  
2、由工具处理.class 文件使用的注解，如有些工具会在加载 class 的时候，对 class 做动态修改，实现一些特殊的功能。这类注解会被编译进入 .class 文件，但加载结束后并不会存在于内存中。这类注解只被一些底层库使用，一般我们不必自己处理。  
3、在程序运行期能够读取的注解，它们在加载后一直存在于 JVM 中，这也是最常用的注解。如一个配置了 \@PostConstruct 的方法会在调用构造方法后自动被调用（这是 Java 代码读取该注解实现的功能，JVM 并不会识别该注解）。  

定义一个注解时，还可以定义配置参数：所有基本类型、String、枚举、基本类型或 String 或枚举的数组。  
配置参数必须是常量。   
注解的配置参数可以有默认值，缺少某个配置参数时将使用默认值。  
大部分注解会有一个名为 value 的配置参数，对此参数赋值可以只写常量，相当于省略了 value 参数。  
如果只写注解，相当于全部使用默认值。  
```java
public class Demo {
    @Check(min=0, max=100, value=55)
    public int m;

    @Check(value=99)
    public int n;

    @Check(99) // @Check(value=99)
    public int x;

    @Check
    public int y;
}
```

Java 使用 \@interface 语法来定义注解。  

元注解（meta annotation）:可以修饰其他注解的注解。  
Java 标准库已经定义了一些元注解，我们只需要使用元注解，通常不需要自己去编写元注解。  
> \@Target 定义 Annotation 能够被应用于源码的哪些位置 【required】:  
> 类或接口：ElementType.TYPE  
> 字段：ElementType.FIELD  
> 方法：ElementType.METHOD  
> 构造方法：ElementType.CONSTRUCTOR  
> 方法参数：ElementType.PARAMETER  
> 
> \@Retention 定义了 Annotation 的生命周期，默认 CLASS 【required】:  
> 仅编译期：RetentionPolicy.SOURCE（编译器使用）  
> 仅 class 文件：RetentionPolicy.CLASS（底层工具库使用）  
> 运行期：RetentionPolicy.RUNTIME  
> 
> \@Repeatable 定义 Annotation 可重复  
> 
> \@Inherited 定义子类可继承父类定义的 Annotation  
> \@Inherited 仅针对 \@Target(ElementType.TYPE) 类型的 annotation 有效  
> \@Inherited 仅针对 class 的继承，对 interface 的继承无效  
> 

```java
/*
应用多个注解参数时
@Target({
    ElementType.METHOD,
    ElementType.FIELD
})
 */
@Target(ElementType.METHOD)
// 自定义的 Annotation 都是 RUNTIME
@Retention(RetentionPolicy.RUNTIME)
@Repeatable
public @interface Log 
{
    // 建议所有参数都设置默认值
    int type() default 0;
    String level() default "info";
    String value() default ""; // 最常用参数
}

@Log(type=1, level="debug")
@Log(type=2, level="warning")
public class Demo 
{}
```

注解定义后也是一种 class。  
所有的注解都继承自 java.lang.annotation.Annotation。

Java 提供的使用反射 API 读取 Annotation 的方法：  
> 判断某个注解是否存在于 Class、Field、Method 或 Constructor：  
> Class.isAnnotationPresent(Class)  
> Field.isAnnotationPresent(Class)  
> Method.isAnnotationPresent(Class)  
> Constructor.isAnnotationPresent(Class)  
> 
> 使用反射 API 读取 Annotation：  
> Class.getAnnotation(Class)  
> Field.getAnnotation(Class)  
> Method.getAnnotation(Class)  
> Constructor.getAnnotation(Class)  

```java
// 判断 Log 是否存在于 Demo 类
Demo.class.isAnnotationPresent(Log.class);

// 获取 Demo 定义的 @Log 注解:  
Class cls = Demo.class;
if (cls.isAnnotationPresent(Log.class)) {
    Log log = Demo.class.getAnnotation(Log.class); // 如果不存在会返回 null
    int type = log.type();
    String level = log.level();
}

// 读取方法参数的 Annotation
public void hello(@NotNull @Range(max=5) String name, @NotNull String prefix){}
// 需要先通过反射获取 Method 实例: Method m = ...  
// 获取所有参数的 Annotation:
Annotation[][] annos = m.getParameterAnnotations();
// 第一个参数（索引为 0）的所有 Annotation:
Annotation[] annosOfName = annos[0];
for (Annotation anno : annosOfName) {
    if (anno instanceof Range) { // @Range注解
        Range r = (Range) anno;
    }
    if (anno instanceof NotNull) { // @NotNull注解
        NotNull n = (NotNull) anno;
    }
}
```

由于 JVM 不会自动给注解添加任何额外的逻辑，所以我们自定义的注解需要编写相关逻辑才能达到想要的效果。  
```java
import java.lang.annotation.ElementType;
import java.lang.annotation.RetentionPolicy;

import java.lang.annotation.Retention;
import java.lang.annotation.Target;

import java.lang.reflect.Field;

@Target(ElementType.FIELD)
@Retention(RetentionPolicy.RUNTIME)
public @interface Range 
{
    int min() default 0;
    int max() default 255;
}

public class Person 
{
    @Range(min=1, max=20)
    public String name;

    public Person(String name) 
    {
        this.name = name;
    }

    @Override
    public String toString() 
    {
        return String.format("{Person: name=%s}", name);
    }
}

static void check(Person person) throws IllegalArgumentException, ReflectiveOperationException {
    // 遍历所有 Field
    for (Field field : person.getClass().getFields()) {
        // 获取 Field 定义的 @Range
        Range range = field.getAnnotation(Range.class);
        // 如果 @Range存在
        if (range != null) {
            // 获取 Field 的值
            Object value = field.get(person);
            // 如果值是 String
            if (value instanceof String) {
                String s = (String) value;
                // 判断值是否满足@Range的min/max:
                if (s.length() < range.min() || s.length() > range.max()) {
                    throw new IllegalArgumentException("Invalid field: " + field.getName());
                }
            }
        }
    }
}

Person p1 = new Person("Bob");
Person p2 = new Person("");
for (Person p : new Person[] { p1, p2 }) {
    try {
        check(p);
        System.out.println("Person " + p + " checked ok.");
    } catch (IllegalArgumentException e) {
        System.out.println("Person " + p + " checked failed: " + e);
    }
}
```

