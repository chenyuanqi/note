

### 什么是 API 网关
基于微服务架构的后端服务通常是动态的，为了简化前端的调用逻辑，通常会引入 API Gateway 作为轻量级网关，同时 API Gateway 中也会实现相关的认证逻辑从而简化内部服务之间相互调用的复杂度。   

对于大多数基于微服务的应用程序而言，API 网关都应该是系统的入口，它会负责服务请求路由、组合及协议转换。

### 为什么 API 网关

- 简化客户端调用复杂度
- 数据裁剪以及聚合
- 多渠道支持
- 遗留系统的微服务化改造

### 实现 API 网关
**Spring Cloud 的 Zuul 组件**  
Spring Cloud 的 Zuul 组件提供了轻量级网关的功能支持，通过定义路由规则可以快速实现一个轻量级的 API 网关。同时，除了通过 serviceId 关联已经注册到 Consul 的服务实例以外，我们也可以通过 zuul 直接定义实现对已有服务的直接集成。  
```
zuul:
  ignoredPatterns: /api/auth
  sensitive-headers: "*"
  ignoreLocalService: true
  retryable: false
  host:
    max-total-connections: 500
  routes:
    service01:
      pateh: /service01/**
      serviceId: service01
      stripPrefix: true
    thirdpart:
      pateh: /thirdpart/**
      url: http://thirdpart.api.com
```
[zuul 参考](https://blog.csdn.net/zhanglh046/article/details/78651993)  

但是，直接使用 Zuul 会存在诸多问题：  
> 性能问题：当存在大量请求超时后会造成 Zuul 阻塞，目前只能通过横向扩展 Zuul 实例实现对高并发的支持  
> WebSocket 的支持问题： Zuul 中并不直接提供对 WebSocket 的支持，需要添加额外的过滤器实现对 WebSocket 的支持

**Nginx 动态代理**  
为了解决以上 Zuul 的问题，可以通过在 Zuul 前端部署 Nginx 实现对 Zuul 实例的反向代理，同时适当的通过添加 Cache 以及请求压缩减少对后端 Zuul 实例的压力。  
通过 Nginx 我们可以实现对多实例 Zuul 的请求代理，同时通过添加适当的缓存以及请求压缩配置可以提升前端 UI 的请求响应时间。这里需要解决的问题是 Nginx 如何动态发现 Zuul 实例信息并且将请求转发到 Zuul 当中。  

consul-template 是一个命令行工具，结合 consul 实现配置文件的动态生成并且支持在配置文件发生变化后触发用户自定义命令。

**另外的选择：Kong**  
Kong 是 Mashape 开源的高性能高可用 API 网关和 API 服务管理层。它基于 OpenResty，进行 API 管理，并提供了插件实现 API 的 AOP。  

- Kong 核心基于 OpenResty 构建，实现了请求 / 响应的 Lua 处理化；
- Kong 插件拦截请求 / 响应，如果接触过 Java Servlet，等价于拦截器，实现请求 / 响应的 AOP 处理；
- Kong Restful 管理 API 提供了 API/API 消费者 / 插件的管理；
- 数据中心用于存储 Kong 集群节点信息、API、消费者、插件等信息，目前提供了 PostgreSQL 和 Cassandra 支持，如果需要高可用建议使用 Cassandra；
- dnsmasq 用于提供给 Nginx DNS 解析功能；
- Kong 集群中的节点通过 gossip 协议自动发现其他节点，当通过一个 Kong 节点的管理 API 进行一些变更时也会通知其他节点。每个 Kong 节点的配置信息是会缓存的，如插件，那么当在某一个 Kong 节点修改了插件配置时，需要通知其他节点配置的变更。

有一些特性 Kong 默认是缺失的，如 API 级别的超时、重试、fallback 策略、缓存、API 聚合、ABTest 等，这些需要开发者自己定制和扩展。  
