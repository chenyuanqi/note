
### Map
Map 接口是将键映射到值的对象，一个映射不能包含重复的键，并且每个键最多只能映射到一个值。  

和 Collection 不同的是，Map 是双列的而 Collection 是单列的；Map 集合的数据结构值针对键有效，跟值无关，Collection 集合的数据结构是针对元素有效；Map 的键是唯一的而 Collection 的子体系 Set 才是唯一的。  

HashMap 键是哈希表结构，可以保证键的唯一性。HashMap 和 Hashtable 都实现了 Map 接口，都是键值对保存数据的方式；HashMap 可以存放 null 但是 Hashtable 不能；HashMap 不是线程安全的类而 Hashtable 是。    
LinkedHashMap 接口的哈希表和链接列表实现，具有可预知的迭代顺序。  
TreeMap 键是红黑树结构，可以保证键的排序和唯一性。  

```java
// Map 接口的成员方法
// 
// V put(K key,V value)
// V remove(Object key)
// void clear()
// boolean containsKey(Object key)
// boolean containsValue(Object value)
// boolean isEmpty()
// int size()
// V get(Object key)
// Set<K> keySet()
// Collection<V> values()
// Set<Map.Entry<K,V>> entrySet()

// HashMap 的使用
// 斗地主
// 造一个牌盒(集合)
HashMap<Integer, String> hm = new HashMap<Integer, String>();
// 创建一个 ArrayList 集合
ArrayList<Integer> array = new ArrayList<Integer>();
// 定义花色数组
String[] colors = { "♠", "♥", "♣", "♦" };
// 定义点数数组
String[] numbers = { "A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K" };
// 定义一个索引
int key = 0;
// 纸牌构建
for (String color : colors) {
    for (String number : numbers) {
        String value = color.concat(number);
        hm.put(key, value);
        array.add(key);
        key++;
    }
}
hm.put(key, "小王");
array.add(key);
key++;
hm.put(key, "大王");
array.add(key);
// 洗牌
Collections.shuffle(array);
// 发牌（三个选手）
TreeSet<Integer> linQingXia = new TreeSet<Integer>();
TreeSet<Integer> fengQingYang = new TreeSet<Integer>();
TreeSet<Integer> liuYi = new TreeSet<Integer>();
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
lookPoker("林青霞", linQingXia, hm);
lookPoker("风清扬", fengQingYang, hm);
lookPoker("刘意", liuYi, hm);
// 遍历
public static void lookPoker(String name, TreeSet<Integer> array, HashMap<Integer, String> hm) {
    System.out.print(name + "的牌是：");
    for (Integer key : ts) {
        String value = hm.get(key);
        System.out.print(value + " ");
    }
    System.out.println();
}
```
