
### Java 泛型
早期的 Object 类型可以接收任意的对象类型，但是在实际的使用中，会有类型转换的问题。  

泛型是一种 “代码模板”，可以用一套代码套用各种类型。  
泛型的好处是使用时不必对类型进行强制转换，它通过编译器对类型进行检查。  

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
        return this.name.compareTo(other.name);
    }

    public String toString() 
    {
        return this.name + "," + this.score;
    }
}

public interface Comparable<T> 
{
    /**
     * 返回 -1: 当前实例比参数 o 小
     * 返回 0: 当前实例与参数 o 相等
     * 返回 1: 当前实例比参数 o 大
     */
    int compareTo(T o);
}
```

Java 语言的泛型实现方式是擦拭法（Type Erasure）。  
所谓擦拭法是指，虚拟机对泛型其实一无所知，所有的工作都是编译器做的。  
Java 的泛型是由编译器在编译时实行的，编译器内部永远把所有类型 T 视为 Object 处理，但是，在需要转型的时候，编译器会根据 T 的类型自动为我们实行安全地强制转型。  
所以，Java 泛型的局限由如下几点：  
1、<T\> 不能是基本类型，例如 int，因为实际类型是 Object，Object 类型无法持有基本类型  
2、无法取得带泛型的 Class  
3、无法判断带泛型的 Class  
4、不能实例化 T 类型  

泛型的应用
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
