
### Java 反射
Java 反射（Reflection）是指程序在运行期可以拿到一个对象的所有信息。反射是为了解决在运行期，对某个实例一无所知的情况下，如何调用其方法。  

除了 int 等基本类型外，Java 的其他类型全部都是 class。class 是由 JVM 在执行过程中动态加载的，JVM 在第一次读取到一种 class 类型时，将其加载进内存；每加载一种 class，JVM 就为其创建一个 Class 类型的实例（一个名叫 Class 的 class），并关联起来，该 Class 实例包含了该 class 的所有完整信息。  

JVM 动态加载 class 的特性：JVM 在执行 Java 程序的时候，并不是一次性把所有用到的 class 全部加载到内存，而是第一次需要用到 class 时才加载。  
利用 JVM 动态加载 class 的特性，我们才能在运行期根据条件加载不同的实现类。  
```java
import java.lang.reflect.Field;
import java.lang.reflect.Modifier;
import java.lang.reflect.Method;
import java.lang.reflect.Constructor;

// 获取一个 class 的 Class 实例
// 1、直接通过一个 class 的静态变量 class 获取
Class cls = Person.class;
// 2、通过该实例变量提供的 getClass() 方法获取
Person p = new Person();
Class cls = p.getClass();
// 3、通过静态方法 Class.forName(class 的完整类名) 获取 
Class cls = Class.forName("qi.Person");

//  Class 实例比较和 instanceof 的差别
// == 判断 class 实例可以精确地判断数据类型，但不能作子类型比较
Integer n = new Integer(123);
boolean b1 = n.getClass() == Integer.class; // true
boolean b2 = n.getClass() == Number.class; // false
// instanceof 不但匹配当前类型，还匹配当前类型的子类
boolean b3 = n instanceof Integer; // true
boolean b4 = n instanceof Number; // true


// 通过反射获取实例的 class 信息
System.out.println("Class name: " + cls.getName());
System.out.println("Simple name: " + cls.getSimpleName());
if (cls.getPackage() != null) {
    System.out.println("Package name: " + cls.getPackage().getName());
}
System.out.println("is interface: " + cls.isInterface());
System.out.println("is enum: " + cls.isEnum());
System.out.println("is array: " + cls.isArray());
System.out.println("is primitive: " + cls.isPrimitive());


// 获取字段属性信息，返回 Field 或 Field[]，不存在会抛异常 NoSuchFieldException
Field f = cls.getField("public_field"); // public java.lang.String Person.name
System.out.println("Field name：" + f.getName()); // name
System.out.println("Field type" + f.getType());  // typeclass java.lang.String
int m = f.getModifiers();
System.out.println("Field modifier：" + m); // 1
Modifier.isFinal(m); 
Modifier.isPublic(m); 
Modifier.isProtected(m); 
Modifier.isPrivate(m); 
Modifier.isStatic(m); 
// 获取所有 public field 的方法：cls.getFields();
System.out.println(f); 
System.out.println(cls.getDeclaredField("private_field")); 
// 获取所有 private field 的方法：cls.getDeclaredFields();

// 获取字段属性值
Field f = cls.getField("public_field");
// 如果是 private 字段，需要加这句：f.setAccessible(true);
// setAccessible 可能会失败，如果 JVM 运行期存在 SecurityManager 阻止
Object v = f.get(p);

// 设置字段属性值
Field f = cls.getField("public_field");
// 同样，如果是 private 字段，需要加这句：f.setAccessible(true);
f.set(p, "public_field_value");


// 获取 public 方法，返回 Method 或 Method[]，不存在会抛异常 NoSuchMethodException
Method m = cls.getMethod("method_name", String.class); // 参数为 String 时
System.out.println("Method name：" + m.getName());
System.out.println("Method return type" + m.getReturnType()); // Class 实例
Class[] c = m.getParameterTypes(); // Class 数组
System.out.println("Method parameter types" + c); 
int mo = m.getModifiers(); // 方法的修饰符
// 获取所有 public 方法：cls.getMethods();
Method m = cls.getDeclaredMethod("method_name", int.class); // 参数为 int 时
// 获取所有 private 方法：cls.getDeclaredMethods()

// 调用 public 方法
String r = (String) m.invoke(p, "qi"); // 调用静态方法时，第一个参数为 null
System.out.println(r); // 方法返回的结果
// 调用非 private 方法
Method m = cls.getDeclaredMethod("setName", String.class);
m.setAccessible(true); // 设置允许调用
m.invoke(p, "Bob");
System.out.println(p.name);

// 构造方法总是当前类定义的构造方法，与父类无关，不存在多态问题
// 调用构造方法（调用该类 public 无参数的构造方法）
Person p = Person.class.newInstance(); 
// 调用任意构造方法（该类 public 带参数方法）
Constructor cons1 = cls.getConstructor(int.class);
Integer n1 = (Integer) cons1.newInstance(123);
System.out.println(n1);
Constructor cons2 = cls.getConstructor(String.class);
Integer n2 = (Integer) cons2.newInstance("456");
System.out.println(n2);
// 获取所有 public 的构造方法：cls.getConstructors();
// 调用任意构造方法（该类非 public 带参数方法）
Constructor cons3 = cls.getDeclaredConstructor(int.class);
cons3.setAccessible(true);
Integer n3 = (Integer) cons3.newInstance(789);
System.out.println(n3);
// 获取所有非 public 的构造方法：cls.getDeclaredConstructors();
```

`注意：当类的方法被子类 @override 之后，反射仍然遵循多态原则：即总是调用实际类型的覆写方法（如果存在）。`

除 Object 外，其他任何非 interface 的 Class 都必定存在一个父类类型。  
```java
// 获取类的父级
Class parent = cls.getSuperclass();
System.out.println("Person 的父类：" + parent); // class java.lang.Object

// 获取类的实现接口（只返回当前类直接实现的接口类型，并不包括其父类实现的接口类型）
Class[] is = cls.getInterfaces(); // 没有时，返回 null
for (Class i : is) {
    System.out.println(i);
}

// 判断一个实例是否属于某个类，使用 instanceof 即可
System.out.println(p instanceof Person); // true
// 如果时两个 class 的实例，判断向上转型是否成立，用 isAssignableFrom()
Number.class.isAssignableFrom(Integer.class); // true，因为 Integer 可以赋值给 Number
Integer.class.isAssignableFrom(Number.class); // false，因为 Number 不能赋值给 Integer
```

### Java 动态代理
我们知道，非抽象类可以被实例化，接口不能被实例化。  
Java 标准库提供了一种动态代理（Dynamic Proxy）的机制：可以在运行期动态创建某个 interface 的实例。  
```java
public interface Animal 
{
    void say(String name);
}

// 正常情况下，我们是这样实现接口的
public class Person implements Animal 
{
    public void say(String name) 
    {
        System.out.println("hello, " + name);
    }
}
Person p = new Person();
p.say("Baby");

// 动态方式（JDK 在运行期动态创建 class 字节码并加载的过程）
import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;
// 定义 InvocationHandler 实例，负责实现接口的方法调用
InvocationHandler handler = new InvocationHandler() {
    @Override
    public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
        if (method.getName().equals("say")) {
            System.out.println("hello, " + args[0]);
        }

        return null;
    }
};

Animal p = (Animal) Proxy.newProxyInstance(
    Animal.class.getClassLoader(), // 接口类的 ClassLoader
    new Class[] { Animal.class },  // 需要实现的接口数组，至少需要传入一个接口进去
    handler
);
p.say("Baby");
```
