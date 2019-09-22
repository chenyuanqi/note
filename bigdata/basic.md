
### 理解 Map/Reduce
Map/Reduce 是一种编程模型，是一种编程方法，抽象理论。  
Map 的作用是过滤一些原始数据，Reduce 则是处理这些数据，得到我们想要的结果。  
```java
public class Test 
{
  public static void main(String[] args) 
  {
    Map<Long, Integer> map = Maps.newHashMapWithExpectedSize(5);
    map.put(1l, Objects.hashCode(Lists.newArrayList(1, 2, 3)));
    map.put(2l, Objects.hashCode(Lists.newArrayList(2, 3, 4)));
    map.put(3l, Objects.hashCode(Lists.newArrayList(1, 2, 3)));
    map.put(4l, Objects.hashCode(Lists.newArrayList(1, 2, 5)));
    map.forEach((k, v) -> System.out.println(k + " - " + v));
    System.out.println("---------------------------------");
    
    System.out.println(
        map.entrySet().stream().collect(groupingBy(Map.Entry::getValue))
            .values().stream()
            .map(v -> v.get(0).getKey()).collect(Collectors.toList())
    );
    System.out.println("---------------------------------");

    System.out.println(
        map.entrySet().stream().collect(groupingBy(Map.Entry::getValue)).values().stream()
            .map(e -> e.stream().reduce((e1, e2) -> e1).get().getKey()).collect(Collectors.toList()));
  }
}
```
