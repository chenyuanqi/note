
### AOP 是什么
AOP（Aspect Oriented Programming），即面向切面编程，可以说是 OOP（Object Oriented Programming，面向对象编程）的补充和完善。OOP 允许开发者定义纵向的关系，但并不适合定义横向的关系（比如日志代码往往横向地散布在所有对象层次中，却与它对应的对象的核心功能毫无关系）。  

AOP 的本质是在一系列纵向的控制流程中，把那些相同的子流程提取成一个横向的面。AOP 利用一种称为 "横切" 的技术，剖解开封装的对象内部，并将那些影响了多个类的公共行为封装到一个可重用模块，并将其命名为 "Aspect"，即切面。  
所谓 "切面"，简单说就是那些与业务无关，却为业务模块所共同调用的逻辑或责任封装起来，便于减少系统的重复代码，降低模块之间的耦合度，并有利于未来的可操作性和可维护性。  

### AOP 核心概念
**横切关注点**  
对哪些方法进行拦截，拦截后怎么处理，这些关注点称之为横切关注点。  

**切面（aspect）**
类是对物体特征的抽象，切面就是对横切关注点的抽象。  

**连接点（joinpoint）**  
被拦截到的点，因为 Spring 只支持方法类型的连接点，所以在 Spring 中连接点指的就是被拦截到的方法，实际上连接点还可以是字段或者构造器。  

**切入点（pointcut）**  
对连接点进行拦截的定义。  

**通知（advice）**  
所谓通知指的就是指拦截到连接点之后要执行的代码，通知分为前置、后置、异常、最终、环绕通知五类。  

**目标对象**  
代理的目标对象。  

**织入（weave）**  
将切面应用到目标对象并导致代理对象创建的过程。  

**引入（introduction）**  
在不修改代码的前提下，引入可以在运行期为类动态地添加一些方法或字段。  

### Spring AOP 
使用 Spring AOP，需要额外下载以下两个 jar 包：  
> aopalliance.jar  
> aspectjweaver.jar  

```java
public interface HelloWorld
{
    void printHelloWorld();
    void doPrint();
}

public class HelloWorldImpl1 implements HelloWorld
{
    public void printHelloWorld()
    {
        System.out.println("Enter HelloWorldImpl1.printHelloWorld()");
    }
    
    public void doPrint()
    {
        System.out.println("Enter HelloWorldImpl1.doPrint()");
        return ;
    }
}

public class HelloWorldImpl2 implements HelloWorld
{
    public void printHelloWorld()
    {
        System.out.println("Enter HelloWorldImpl2.printHelloWorld()");
    }
    
    public void doPrint()
    {
        System.out.println("Enter HelloWorldImpl2.doPrint()");
        return ;
    }
}

public class TimeHandler
{
    public void printTime()
    {
        System.out.println("CurrentTime = " + System.currentTimeMillis());
    }
}
```

添加  aop.xml 配置：  
```xml
<?xml version="1.0" encoding="UTF-8"?>
<beans xmlns="http://www.springframework.org/schema/beans"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xmlns:aop="http://www.springframework.org/schema/aop"
    xmlns:tx="http://www.springframework.org/schema/tx"
    xsi:schemaLocation="http://www.springframework.org/schema/beans
        http://www.springframework.org/schema/beans/spring-beans-4.2.xsd
        http://www.springframework.org/schema/aop
        http://www.springframework.org/schema/aop/spring-aop-4.2.xsd">
        
        <bean id="helloWorldImpl1" class="com.xrq.aop.HelloWorldImpl1" />
        <bean id="helloWorldImpl2" class="com.xrq.aop.HelloWorldImpl2" />
        <bean id="timeHandler" class="com.xrq.aop.TimeHandler" />
        
        <aop:config>
            <aop:aspect id="time" ref="timeHandler">
                <aop:pointcut id="addAllMethod" expression="execution(* com.xrq.aop.HelloWorld.*(..))" />
                <aop:before method="printTime" pointcut-ref="addAllMethod" />
                <aop:after method="printTime" pointcut-ref="addAllMethod" />
            </aop:aspect>
        </aop:config>
</beans>
```

使用 aop：  
```java
public static void main(String[] args)
{
    ApplicationContext ctx = new ClassPathXmlApplicationContext("aop.xml");
        
    HelloWorld hw1 = (HelloWorld)ctx.getBean("helloWorldImpl1");
    HelloWorld hw2 = (HelloWorld)ctx.getBean("helloWorldImpl2");
    hw1.printHelloWorld();
    System.out.println();
    hw1.doPrint();
    
    System.out.println();
    hw2.printHelloWorld();
    System.out.println();
    hw2.doPrint();
}
```
