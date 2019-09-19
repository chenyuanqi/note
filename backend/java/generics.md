
### Java 泛型
早期的 Object 类型可以接收任意的对象类型，但是在实际的使用中，会有类型转换的问题。  

泛型本质上是类型参数化，解决了不确定对象的类型问题。  

泛型是一种 “代码模板”，可以用一套代码套用各种类型。  
泛型的好处是使用时不必对类型进行强制转换，它通过编译器对类型进行检查。所有容器的内部存放的都是 Object 对象，泛型机制只是简化了编程，由编译器自动帮我们完成了强制类型转换而已。JDK 1.4 以及之前版本不支持泛型，类型转换需要显式完成。  

泛型的优点：  

- 安全：不用担心程序运行过程中出现类型转换的错误
- 避免了类型转换：如果是非泛型，获取到的元素是 Object 类型的，需要强制类型转换
- 可读性高：编码阶段就明确的知道集合中元素的类型

虚拟机中没有泛型，只有普通类和普通方法，所有泛型类的类型参数在编译时都会被擦除，泛型类并没有自己独有的 Class 类对象。比如并不存在 List<String>.class 或是 List<Integer>.class，而只有 List.class。

`注意泛型的继承关系：可以把 ArrayList<Integer> 向上转型为 List<Integer>（T 不能变！），但不能把 ArrayList<Integer> 向上转型为 ArrayList<Number>（T 不能变成父类）`

```java
public class ArrayList<T> {
    private T[] array;
    private int size;
    public void add(T e) {...}
    public void remove(int index) {...}
    public T get(int index) {...}
}

// 不指定泛型参数类型时，编译器会给出警告，且只能将 <T> 视为 Object 类型；
// List list = new ArrayList();
// 
// 创建可以存储 String 的 ArrayList:
ArrayList<String> strList = new ArrayList<String>();
// 创建可以存储 Float 的 ArrayList:
ArrayList<Float> floatList = new ArrayList<Float>();
// 创建可以存储 Person 的 ArrayList:
ArrayList<Person> personList = new ArrayList<Person>();

// 创建可以存储 Number 的 ArrayList（可以省略后面的 Number，编译器可以自动推断泛型类型）:
List<Number> list = new ArrayList<>();
// 在 Java 标准库中的 ArrayList<T> 实现了 List<T> 接口，它可以向上转型为 List<T>

// 可以在接口中定义泛型类型
// 实现此接口的类必须实现正确的泛型类型
// 
// Person 排序的例子
import java.util.Arrays;

public class Main 
{
    public static void main(String[] args) 
    {
        Person[] ps = new Person[] {
            new Person("Bob", 61),
            new Person("Alice", 88),
            new Person("Lily", 75),
        };
        Arrays.sort(ps);
        System.out.println(Arrays.toString(ps));
    }
}

class Person implements Comparable<Person> 
{
    String name;
    int score;
    Person(String name, int score) 
    {
        this.name = name;
        this.score = score;
    }

    public int compareTo(Person other) 
    {
        return this.score.compareTo(other.score);
    }

    public String toString() 
    {
        return this.name + "," + this.score;
    }
}

public interface Comparable<T> 
{
    /**
     * 返回 -1: 当前实例比参数 t 小
     * 返回 0: 当前实例与参数 t 相等
     * 返回 1: 当前实例比参数 t 大
     */
    int compareTo(T t);
}
```

### 泛型的工作原理
泛型是通过类型擦除来实现的，类型擦除指的是编译器在编译时，会擦除了所有类型相关的信息，比如 List<String> 在编译后就会变成 List 类型，这样做的目的就是确保能和 Java 5 之前的版本（二进制类库）进行兼容。  

### 编写泛型
泛型类一般用在集合类中，我们几乎不需要自己编写泛型，但是有必要学习一下。  
1、按照某种类型来编写类，标记所有的特定类型（比如String）  
```java
public class Pair 
{
    private String first;
    private String last;
    public Pair(String first, String last) 
    {
        this.first = first;
        this.last = last;
    }

    public String getFirst() 
    {
        return first;
    }

    public String getLast() 
    {
        return last;
    }
}
```
2、把特定类型替换为 T，并声明 <T\>  
```java
public class Pair<T> 
{
    private T first;
    private T last;
    public Pair(T first, T last) 
    {
        this.first = first;
        this.last = last;
    }

    public T getFirst() 
    {
        return first;
    }

    public T getLast() 
    {
        return last;
    }
}
```

`注意：泛型类型 <T> 不能用于静态方法`  
对于静态方法，我们可以单独改写为 “泛型” 方法，只需要使用另一个类型即可。这样，静态方法的泛型类型和实例类型的泛型类型就可以区分开。   
```java
public static <K> Pair<K> create(K first, K last) {
    return new Pair<K>(first, last);
}
```

泛型还可以定义多种类型。Java 标准库的 Map<K, V> 就是使用两种泛型类型的例子。它对 Key 使用一种类型，对 Value 使用另一种类型。  
```java
public class Pair<T, K> 
{
    private T first;
    private K last;
    public Pair(T first, K last) 
    {
        this.first = first;
        this.last = last;
    }
}

Pair<String, Integer> p = new Pair<>("test", 123);
```

### 擦拭法
Java 语言的泛型实现方式是擦拭法（Type Erasure）。  
所谓擦拭法是指，虚拟机对泛型其实一无所知，所有的工作都是编译器做的。  
Java 的泛型是由编译器在编译时实行的，编译器内部永远把所有类型 T 视为 Object 处理，但是，在需要转型的时候，编译器会根据 T 的类型自动为我们实行安全地强制转型。  
所以，Java 泛型的局限由如下几点：  
1、<T\> 不能是基本类型，例如 int，因为实际类型是 Object，Object 类型无法持有基本类型  
2、无法取得带泛型的 Class  
3、无法判断带泛型的 Class  
4、不能直接实例化 T 类型，需要借助额外的 Class<T> 参数  

**不恰当的覆写方法**  
定义的 equals(T t) 方法实际上会被擦拭成 equals(Object t)，而这个方法是继承自 Object 的，编译器会阻止一个实际上会变成覆写的泛型方法定义。我们需要换个方法名就可以通过编译。   
```java
public class Pair<T> {
    public boolean equals(T t) {
        return this == t;
    }
}
```

**泛型的继承**  
一个类可以继承自一个泛型类。  
我们无法获取 Pair<T> 的 T 类型，即给定一个变量 Pair<Integer> p，无法从 p 中获取到 Integer 类型；但是，在继承了泛型类型的情况下，子类可以获取父类的泛型类型。  
```java
public class IntPair extends Pair<Integer> 
{
}

IntPair ip = new IntPair(1, 2);
```

**泛型的应用**  
```java
// 实现最小值函数
private static <T extends Number & Comparable<? super T>> T min(T[] values) {
    if (values == null || values.length == 0) return null;
    T min = values[0];
    for (int i = 1; i < values.length; i++) {
        if (min.compareTo(values[i]) > 0) min = values[i];
    }
    return min;
}

int minInteger = min(new Integer[]{1, 2, 3}); // 1
double minDouble = min(new Double[]{1.2, 2.2, -1d}); // -1d
```

### 泛型通配符
**通配符 extends**  
使用 extends 通配符表示可以读，不能写。  
给方法传入 Pair<Integer\> 类型时，它符合参数 Pair<? extends Number\> 类型。这种使用 <? extends Number\> 的泛型定义称之为上界通配符（Upper Bounds Wildcards），即把泛型类型 T 的上界限定在 Number 了。  
除了可以传入 Pair<Integer\> 类型，我们还可以传入 Pair<Double\> 类型，Pair<BigDecimal\> 类型等等，因为 Double 和 BigDecimal 都是 Number 的子类。
方法内部无法调用传入 Number 引用的方法（null 除外），例如：obj.setFirst(Number n);。  

**通配符 super**  
使用 super 通配符表示只能写不能读。  
使用类似 <? super Integer\> 通配符作为方法参数时，泛型类型限定为 Integer 或 Integer 的超类，表示：  
方法内部可以调用传入 Integer 引用的方法，例如：obj.setFirst(Integer n);；  
方法内部无法调用获取 Integer 引用的方法（Object 除外），例如：Integer n = obj.getFirst();。  

通配符 extends 和 super 通配符要遵循 PECS 原则（Producer Extends Consumer Super）。即如果需要返回 T，它是生产者（Producer），要使用 extends 通配符；如果需要写入 T，它是消费者（Consumer），要使用 super 通配符。  

**无限定通配符**  
Java 的泛型还允许使用无限定通配符（Unbounded Wildcard Type），即只定义一个 <?\>。  
无限定通配符不允许调用 set(T) 方法并传入引用（null 除外）；不允许调用 T get() 方法并获取 T 引用（只能获取 Object 引用）。换句话说，既不能读，也不能写，那只能做一些 null 判断。  
```java
static boolean isNull(Pair<?> p) 
{
    return p.getFirst() == null || p.getLast() == null;
}
```
<?> 通配符有一个独特的特点，就是：Pair<?> 是所有 Pair<T> 的超类。  

### 泛型与反射
Java 的部分反射 API 也是泛型，比如 Class<T\> 就是泛型。  
```java
// compile warning:
Class clazz = String.class;
String str = (String) clazz.newInstance();

// no warning:
Class<String> clazz = String.class;
String str = clazz.newInstance();

// 返回的 Class 类型是 Class<? super T>
Class<? super String> sup = String.class.getSuperclass();

// 构造方法 Constructor<T> 也是泛型
Class<Integer> clazz = Integer.class;
Constructor<Integer> cons = clazz.getConstructor(int.class);
Integer i = cons.newInstance(123);

// 可以声明带泛型的数组，但不能用 new 操作符创建带泛型的数组
Pair<String>[] ps = null; // ok
Pair<String>[] ps = new Pair<String>[2]; // compile error!
// 必须通过强制转型实现带泛型的数组
Pair<String>[] ps = (Pair<String>[]) new Pair[2];
```

使用泛型的可变参数时需要特别小心  
```java
static <K> K[] pickTwo(K k1, K k2, K k3) {
    return asArray(k1, k2);
}

static <T> T[] asArray(T... objs) {
    return objs;
}
```
直接调用 asArray(T...) 似乎没有问题，但是在另一个方法中，我们返回一个泛型数组就会产生 ClassCastException，原因还是因为擦拭法，在 pickTwo() 方法内部，编译器无法检测 K[] 的正确类型，因此返回了 Object[]。  
编译器对所有可变泛型参数都会发出警告，除非确认完全没有问题，才可以用 \@SafeVarargs 消除警告。  
`如果在方法内部创建了泛型数组，最好不要将它返回给外部使用。`  
