
### Map
Map 接口是将键映射到值的对象，一个映射不能包含重复的键，并且每个键最多只能映射到一个值。  

和 Collection 不同的是，Map 是双列的而 Collection 是单列的；Map 集合的数据结构值针对键有效，跟值无关，Collection 集合的数据结构是针对元素有效；Map 的键是唯一的而 Collection 的子体系 Set 才是唯一的。  

HashMap 键是哈希表结构，可以保证键的唯一性。  
LinkedHashMap 接口的哈希表和链接列表实现，具有可预知的迭代顺序。  
TreeMap 键是红黑树结构，可以保证键的排序和唯一性。  

Map 接口的成员方法
```java
V put(K key,V value)
V remove(Object key)
void clear()
boolean containsKey(Object key)
boolean containsValue(Object value)
boolean isEmpty()
int size()
V get(Object key)
Set<K> keySet()
Collection<V> values()
Set<Map.Entry<K,V>> entrySet()
```
