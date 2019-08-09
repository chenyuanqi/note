
### JavaBean
在 Java 中，有很多 class 的定义都符合这样的规范：  
若干 private 实例字段；  
通过 public 方法来读写实例字段  
```java
// 读方法:
public Type getXyz()
// 写方法:
public void setXyz(Type value)

// boolean比较特殊
// 读方法:
public boolean isChild()
// 写方法:
public void setChild(boolean value)
```
如果读写方法符合以上的命名规范，那么这种 class 被称为 JavaBean。  

通常把一组对应的读方法（getter）和写方法（setter）称为属性（property），只有 getter 的属性称为只读属性（read-only），只有 setter 的属性称为只写属性（write-only）。属性只需要定义 getter 和 setter 方法，不一定需要对应的字段（比如 Boolean 属性，可以借助其他字段判断）。  

JavaBean 主要用来传递数据，即把一组数据组合成一个 JavaBean 便于传输。而且，JavaBean 可以方便地被 IDE 工具分析（通过定义属性即可生成相应的 getter 和 setter），生成读写属性的代码，主要用在图形界面的可视化设计中。  

```java
import java.beans.*;

// 枚举一个 JavaBean 的所有属性
BeanInfo info = Introspector.getBeanInfo(Person.class);
for (PropertyDescriptor pd : info.getPropertyDescriptors()) {
    System.out.println(pd.getName());
    System.out.println("  " + pd.getReadMethod());
    System.out.println("  " + pd.getWriteMethod());
}
```
