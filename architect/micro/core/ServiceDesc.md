
### 服务发布和引用
服务发布和引用常见的三种方式：Restful API、XML 配置以及 IDL 文件。  

### 服务描述之 XML 配置
**XML 配置方式的服务发布和引用流程**  
服务提供者定义好接口，并且在服务发布配置文件中配置要发布的接口名，在进程启动时加载服务发布配置文件就可以对外提供服务。   
服务消费者通过在服务引用配置文件中定义相同的接口名，并且在服务引用配置文件中配置要引用的接口名，在进程启动时加载服务引用配置文件就可以引用服务。  

1、服务提供者定义接口  
服务提供者发布服务之前，首先要定义接口，声明接口名、传递参数以及返回值类型，然后把接口打包成 JAR 包发布出去。  
```java
package com.api.common.status.service;
 
public interface UserLastStatusService {
    /*
     * @param uids
     * @return
     */
    public long getLastStatusId(long uid);
 
    /**
     *
     * @param uids
     * @return
     */
    public Map<Long, Long> getLastStatusIds(long[] uids);
}
```
2、服务提供者发布接口
服务提供者发布的接口是通过在服务发布配置文件中定义接口来实现的。在配置文件中定义要发布的接口、对外暴露的协议（如 Motan 协议）、端口、方法的超时及重试次数等；然后，服务发布者在进程启动的时候，会加载配置文件，把接口对外暴露出去。  
```xml
<?xml version="1.0" encoding="UTF-8"?>
<beans xmlns="http://www.springframework.org/schema/beans"
      xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
      xmlns:context="http://www.springframework.org/schema/context"
      xmlns:aop="http://www.springframework.org/schema/aop" 
     xsi:schemaLocation="http://www.springframework.org/schema/context
            http://www.springframework.org/schema/context/spring-context-2.5.xsd
http://www.springframework.org/schema/beans http://www.springframework.org/schema/beans/spring-beans-2.5.xsd
http://www.springframework.org/schema/aop http://www.springframework.org/schema/aop/spring-aop-2.5.xsd
">
   <motan:service ref="userLastStatusLocalService"
            requestTimeout="50" retries="2"    interface="com.api.common.status.service.UserLastStatusService"
            basicService="serviceBasicConfig" export="motan:8882">
       <motan:method name="getLastStatusId" requestTimeout="300"
                  retries="0" />
       <motan:method name="getLastStatusIds" requestTimeout="300"
                  retries="0" />
    </motan:service>
</beans>
```
3、服务消费者引用接口  
服务消费者引用接口是通过在服务引用配置文件中定义要引用的接口，并把包含接口定义的 JAR 包引入到代码依赖中；然后服务消费者在进程启动时，会加载配置文件来完成服务引用。  
```xml
<?xml version="1.0" encoding="UTF-8"?>
<beans xmlns="http://www.springframework.org/schema/beans"
      xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
      xmlns:context="http://www.springframework.org/schema/context"
      xmlns:aop="http://www.springframework.org/schema/aop" 
     xsi:schemaLocation="http://www.springframework.org/schema/context
            http://www.springframework.org/schema/context/spring-context-2.5.xsd
http://www.springframework.org/schema/beans http://www.springframework.org/schema/beans/spring-beans-2.5.xsd
http://www.springframework.org/schema/aop http://www.springframework.org/schema/aop/spring-aop-2.5.xsd
">
<motan:protocol name="motan" default="true" loadbalance="${service.loadbalance.name}" />
<motan:basicReferer id="userLastStatusServiceClientBasicConfig"
               protocol="motan" />
<!-- 导出接口 -->
<motan:referer id="commonUserLastStatusService" interface="com.api.common.status.service.UserLastStatusService"
            basicReferer="userLastStatusServiceClientBasicConfig" />
</beans>
```

### 服务发布和引用的常见问题
在服务发布的时候，最好预定义好接口的各种配置。  
在一个服务被多个服务消费者引用的情况下，由于业务经验的参差不齐，可能不同的服务消费者对服务的认知水平不一，比如某个服务可能调用超时了，最好可以重试来提供调用成功率。  
可能有的服务消费者会忽视这一点，并没有在服务引用配置文件中配置接口调用超时重试的次数。所以，最好是可以在服务发布的配置文件中预定义好类似超时重试次数，即使服务消费者没有在服务引用配置文件中定义，也能继承服务提供者的定义。  

在服务规模不大，业务比较简单的时候，这样做比较合适。但是对于复杂业务，虽然服务发布时预定义好接口的各种配置，但在引用的服务消费者众多且同时访问的时候，可能会引起网络风暴。这种情况下，比较保险的方式是，把接口的各种配置放在服务引用配置文件里。  
`在进行服务配置升级过程时，要考虑好步骤，在所有服务消费者完成升级之前，服务提供者还不能把服务的详细信息去掉，否则可能会导致没有升级的服务消费者引用异常` 

1、服务发布预定义配置  
比如服务发布配置文件 server.xml 提供一个服务 contentSliceRPCService，并且明确了其中三个方法的调用超时时间为 500ms 以及超时重试次数为 3。  
```xml
<motan:service ref="contentSliceRPCService" interface="cn.sina.api.data.service.ContentSliceRPCService"
            basicService="serviceBasicConfig" export="motan:8882" >
   <motan:method name="saveContent" requestTimeout="500"
              retries="3" />
   <motan:method name="deleteContent" requestTimeout="500"
              retries="3" />
   <motan:method name="updateContent" requestTimeout="500"
              retries="3" />
</motan:service>
```
服务引用的配置文件 client.xml，服务消费者会默认继承服务发布配置文件中设置的方法调用的超时时间以及超时重试次数。  
```xml
<motan:referer id="contentSliceRPCService" interface="cn.sina.api.data.service.ContentSliceRPCService" basicReferer="contentSliceClientBasicConfig" >
</motan:referer>
```
通过服务发布预定义配置可以解决多个服务消费者引用服务可能带来的配置复杂的问题。  
但是，一个服务提供者发布的服务有上百个方法，并且每个方法都有各自的超时时间、重试次数等信息。服务消费者引用服务时，完全继承了服务发布预定义的各项配置。这种情况下，服务提供者所发布服务的详细配置信息都需要存储在注册中心中，这样服务消费者才能在实际引用时从服务发布预定义配置中继承各种配置。这里就存在一种风险，当服务提供者发生节点变更，尤其是在网络频繁抖动的情况下，所有的服务消费者都会从注册中心拉取最新的服务节点信息，就包括了服务发布配置中预定的各项接口信息，这个信息不加限制的话可能达到 1M 以上，如果同时有上百个服务消费者从注册中心拉取服务节点信息，在注册中心机器部署为百兆带宽的情况下，很有可能会导致网络带宽打满的情况发生。面对这种情况，最好的办法是把服务发布端的详细服务配置信息转移到服务引用端。  

2、服务引用定义配置  
服务发布配置文件详细定义了服务 userInfoService 的各个方法的配置信息，比如超时时间和重试次数等。  
```xml
<motan:service ref="userInfoService" requestTimeout="50" retries="2"                   interface="cn.api.user.service.UserInfoService" basicService="serviceBasicConfig">
    <motan:method name="addUserInfo" requestTimeout="300" retries="0"/>
    <motan:method name="updateUserPortrait" requestTimeout="300" retries="0"/>
    <motan:method name="modifyUserInfo" requestTimeout="300" retries="0"/>
    <motan:method name="addUserTags" requestTimeout="300" retries="0"/>
    <motan:method name="delUserTags" requestTimeout="300" retries="0"/>
    <motan:method name="processUserCacheByNewMyTriggerQ" requestTimeout="300" retries="0"/>
    <motan:method name="modifyObjectUserInfo" requestTimeout="300" retries="0"/>
    <motan:method name="addObjectUserInfo" requestTimeout="300" retries="0"/>
    <motan:method name="updateObjectUserPortrait" requestTimeout="300" retries="0"/>
    <motan:method name="updateObjectManager" requestTimeout="300" retries="0"/>
    <motan:method name="add" requestTimeout="300" retries="0"/>
    <motan:method name="deleteObjectManager" requestTimeout="300" retries="0"/>
    <motan:method name="getUserAttr" requestTimeout="300" retries="1" />
    <motan:method name="getUserAttrList" requestTimeout="300" retries="1" />
    <motan:method name="getAllUserAttr" requestTimeout="300" retries="1" />
    <motan:method name="getUserAttr2" requestTimeout="300" retries="1" />
</motan:service>
```
可以像下面一样，把服务 userInfoService 的详细配置信息转移到服务引用配置文件中。  
```xml
<motan:referer id="userInfoService" interface="cn.api.user.service.UserInfoService" basicReferer="userClientBasicConfig">
    <motan:method name="addUserInfo" requestTimeout="300" retries="0"/>
    <motan:method name="updateUserPortrait" requestTimeout="300" retries="0"/>
    <motan:method name="modifyUserInfo" requestTimeout="300" retries="0"/>
    <motan:method name="addUserTags" requestTimeout="300" retries="0"/>
    <motan:method name="delUserTags" requestTimeout="300" retries="0"/>
    <motan:method name="processUserCacheByNewMyTriggerQ" requestTimeout="300" retries="0"/>
    <motan:method name="modifyObjectUserInfo" requestTimeout="300" retries="0"/>
    <motan:method name="addObjectUserInfo" requestTimeout="300" retries="0"/>
    <motan:method name="updateObjectUserPortrait" requestTimeout="300" retries="0"/>
    <motan:method name="updateObjectManager" requestTimeout="300" retries="0"/>
    <motan:method name="add" requestTimeout="300" retries="0"/>
    <motan:method name="deleteObjectManager" requestTimeout="300" retries="0"/>
    <motan:method name="getUserAttr" requestTimeout="300" retries="1" />
    <motan:method name="getUserAttrList" requestTimeout="300" retries="1" />
    <motan:method name="getAllUserAttr" requestTimeout="300" retries="1" />
    <motan:method name="getUserAttr2" requestTimeout="300" retries="1" />
</motan:referer>
```
这样的话，服务发布配置文件可以简化为下面这段代码
```xml
<motan:service ref="userInfoService" requestTimeout="50" retries="2"                   interface="cn.api.user.service.UserInfoService" basicService="serviceBasicConfig">
</motan:service>
```
在进行类似的服务详细信息配置，由服务发布配置文件迁移到服务引用配置文件的过程时，尤其要注意迁移步骤问题。  

3、服务配置升级  
由于引用服务的服务消费者众多，并且涉及多个部门，升级步骤就显得异常重要，通常可以按照下面步骤操作。  
> 1、各个服务消费者在服务引用配置文件中添加服务详细信息。  
> 2、服务提供者升级两台服务器，在服务发布配置文件中删除服务详细信息，并观察是否所有的服务消费者引用时都包含服务详细信息。  
> 3、如果都包含，说明所有服务消费者均完成升级，那么服务提供者就可以删除服务发布配置中的服务详细信息。  
> 4、如果有不包含服务详细信息的服务消费者，排查出相应的业务方进行升级，直至所有业务方完成升级。


