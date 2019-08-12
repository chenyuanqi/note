
### Java 集合
面向对象语言常见的操作就是操作对象。  
为了对多个对象进行操作，可以使用对象数组；但是，对象数组的长度是固定的，不适应变化的需求，所以集合来了。  

Java 集合的特点是长度可变、可以存储多种引用类型并且只能存储引用类型。  
不同集合类的数据结构不同，但是功能相似。  
```
# 集合的继承体系结构

Collection
	|--List
		|--ArrayList
		|--Vector
		|--LinkedList
	|--Set
		|--HashSet
		|--TreeSet
		|--EnumSet
```

集合的通用 API
```java
// 添加项 
add(Object obj)
// 添加集合所有项
addAll(Collection c)
// 移除项
remove(Object obj)
// 移除集合所有项
removeAll(Collection c)
// 清空
clear()
// 判断是否包含项
contains(Object obj)
// 判断是否包含整个集合
containsAll(Collection c)
// 判断是否为空
isEmpty()
// 获取迭代器
iterator()
// 获取长度
size()
// 交集
retainAll(Collection c)
// 转数组
toArray()
```

Collection 是最基本的集合接口，一个 Collection 代表一组 Object 的集合，这些 Object 被称作 Collection 的元素。Collection 是一个接口，用以提供规范定义，不能被实例化使用。  
Collections 则是包含各种有关集合操作的多态化静态方法，是工具类。

### 迭代器
迭代器是集合的特有遍历方式，依赖于集合而存在的。  
迭代器被定义为接口，由具体的集合类通过内部类的方式提供实现。由于数据结构不同，每种集合的存储元素和逼历元素的方式会不一样；为了统一标准，把判断（hasNext）和获取功能（next）提取出来定义成接口，使用迭代器的集合必须重写它们。  

```java
Collection c = new ArrayList();
c.add("php");
c.add("java");

Iterator it = c.iterator();
// hasNext() 判断是否有下一个元素
while(it.hasNext()) {
	// next() 获取下一个元素，并自动移动位置
	String s = (String)it.next();
	System.out.println(s);
```

### List
List 是有序的 collection（也称为序列）。此接口可以对列表中每个元素的插入位置进行精确地控制，可以根据元素的整数索引（在列表中的位置）访问元素，并搜索列表中的元素。与 set 不同的是，列表允许有重复的元素。  

ArrayList 底层数据结构是数组，查询快，增删慢；线程不安全，但是效率高。  
Vector 底层数据结构是数组，查询快，增删慢；线程安全，效率低。  
LinkedList 底层数据结构是链表，查询慢，增删快；线程不安全，效率高。  
```java
// ArrayList 的使用
// 斗地主
// 造一个牌盒(集合)
ArrayList<String> array = new ArrayList<String>();
// 定义花色数组
String[] colors = { "♠", "♥", "♣", "♦" };
// 定义点数数组
String[] numbers = { "A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K" };
// 纸牌构建
for (String color : colors) {
	for (String number : numbers) {
		array.add(color.concat(number));
	}
}
array.add("小王");
array.add("大王");
// 洗牌
Collections.shuffle(array);
// 发牌（三个选手）
ArrayList<String> linQingXia = new ArrayList<String>();
ArrayList<String> fengQingYang = new ArrayList<String>();
ArrayList<String> liuYi = new ArrayList<String>();
for (int x = 0; x < array.size(); x++) {
	if (x % 3 == 0) {
		linQingXia.add(array.get(x));
	} else if (x % 3 == 1) {
		fengQingYang.add(array.get(x));
	} else if (x % 3 == 2) {
		liuYi.add(array.get(x));
	}
}
// 看牌
lookPoker("林青霞", linQingXia);
lookPoker("风清扬", fengQingYang);
lookPoker("刘意", liuYi);
// 遍历
public static void lookPoker(String name, ArrayList<String> array) {
	System.out.print(name + "的牌是：");
	for (String s : array) {
		System.out.print(s + " ");
	}
	System.out.println();
}

// Vector 的使用
Vector v = new Vector();
v.addElement("php");
v.addElement("java");

Enumeration en = v.elements();
while(en.hasMoreElements()) {
	String s = (String) en.nextElement();
	System.out.println(s);

// LinkedList 的使用
LinkedList link = new LinkedList();
link.add("php");
link.add("java");
System.out.println("getFirst:" + link.getFirst()); // link.get(0)
System.out.println("getLast:" + link.getLast());
System.out.println("link:" + link);
```

### Set
Set 是一个不包含重复元素的 collection。  

HashSet 底层数据结构是哈希表(元素是链表的数组)，哈希表依赖于哈希值存储；不保证 set 的迭代顺序，特别是它不保证该顺序恒久不变。当向 HashSet 集合中存入一个元素时，HashSet 会调用该对象的 hashCode() 方法来得到该对象的 hashCode 值，然后根据该 HashCode 值决定该对象在 HashSet 中的存储位置。   
LinkedHashSet 元素有序唯一（由链表保证元素有序，哈希表保证元素唯一），LinkedHashSet 集合也是根据元素的 hashCode 值来决定元素的存储位置，但和 HashSet 不同的是，它同时使用链表维护元素的次序，这样使得元素看起来是以插入的顺序保存的。  
TreeSet 底层数据结构是红黑树(红黑树是一种自平衡的二叉树)，使用元素的自然顺序对元素进行排序或者根据创建 set 时提供的 Comparator 进行排序（具体取决于使用的构造方法）。  
EnumSet 是一个专门为枚举类设计的集合类，所有元素都必须是指定枚举类型的枚举值，该枚举类型在创建 EnumSet 时显式、或隐式地指定；EnumSet 的集合元素也是有序的，它们以枚举值在 Enum 类内的定义顺序来决定集合元素的顺序。  

```java
import java.util.Set;
import java.util.HashSet;
import java.util.Iterator;

Set<String> set = new HashSet<String>();

set.add("hello");
set.add("world");

// 迭代器的遍历方式
Iterator<String> it = set.iterator();
while(it.hasNext()) {
	String s = it.next();
	System.out.println(s);
}

// foreach 的遍历方式
for(String s : set) {
	System.out.println(s);
}
System.out.println("set:" + set);
```

