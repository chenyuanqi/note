
### 设计模式之单例模式
单例模式，是设计模式中最简单的一种。  
单例对象的类必须保证只有一个实例存在，并提供一个访问它的全局访问点。  

在我们的系统中，有一些对象其实我们只需要一个，比如线程池、缓存、对话框、注册表、日志对象、充当打印机、显卡等设备驱动程序的对象。事实上，这一类对象只能有一个实例，如果制造出多个实例就可能会导致一些问题的产生，比如：程序的行为异常、资源使用过量、或者不一致性的结果。  
单例模式对于频繁使用的对象，可以省略创建对象所花费的时间，这对于那些重量级对象而言，是非常可观的一笔系统开销；由于 new 操作的次数减少，因而对系统内存的使用频率也会降低，这将减轻 GC 压力，缩短 GC 停顿时间。

单例模式的特点：  
1、单例类只能有一个实例  
2、单例类必须自己自己创建自己的唯一实例  
3、单例类必须给所有其他对象提供这一实例  

单例模式的要点：  
1、私有的构造方法  
2、指向自己实例的私有静态引用  
3、以自己实例为返回值的静态的公有的方法  

单例模式有很多需要注意的东西，比如性能、线程安全等。  

```java
// 懒汉式（线程不安全）
// 懒汉式在调用取得实例方法的时候才会实例化对象
public class Singleton 
{  
    private static Singleton instance;  
    private Singleton (){}  

    public static Singleton getInstance() 
    {  
        if (instance == null) {  
            instance = new Singleton();  
        }  

        return instance;  
    }  
}  

// 懒汉式（线程安全）
public class Singleton 
{  
    private static Singleton instance;  
    private Singleton (){}  

    // synchronized 同步锁
    public static synchronized Singleton getInstance() 
    {  
        if (instance == null) {  
            instance = new Singleton();  
        }  

        return instance;  
    }  
} 

// 饿汉式（基于 classloder 机制）
// 饿汉式单例在单例类被加载时候，就实例化一个对象交给自己的引用
public class Singleton 
{  
    private static Singleton instance = new Singleton();  
    private Singleton (){}  

    public static Singleton getInstance() 
    {  
        return instance;  
    }  
}  

// 饿汉式变种
public class Singleton 
{  
    private Singleton instance = null;  
    static {  
        instance = new Singleton();  
    } 
    private Singleton (){}  

    public static Singleton getInstance() 
    {  
        return this.instance;  
    }  
}  

// 静态内部类
public class Singleton 
{  
    private static class SingletonHolder 
    {  
        // 利用 classloder 的机制来保证初始化 instance 时只有一个线程
        private static final Singleton INSTANCE = new Singleton();  
    }  
    private Singleton (){} 

    public static final Singleton getInstance() 
    {  
        return SingletonHolder.INSTANCE;  
    }  
}  

// 双重校验锁（大部分代码保证线程安全）
public class Singleton 
{  
    private volatile static Singleton singleton;  
    private Singleton (){}  

    public static Singleton getSingleton() 
    {  
        if (singleton == null) {  
            synchronized (Singleton.class) {  
                if (singleton == null) {  
                    singleton = new Singleton();  
                }  
            }  
        }  
        return singleton;  
    }  
}  

// 枚举
public enum Singleton {  
    INSTANCE;  
    public void whateverMethod() {  
    }  
} 

```

