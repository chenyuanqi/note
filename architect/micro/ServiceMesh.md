
### 下一代微服务框架 —— ServiceMesh
Service Mesh 是一种新型的用于处理服务与服务之间通信的技术，尤其适用以云原生应用形式部署的服务，能够保证服务与服务之间调用的可靠性。在实际部署时，Service Mesh 通常以轻量级的网络代理的方式跟应用的代码部署在一起，从而以应用无感知的方式实现服务治理。  
> Service Mesh 以轻量级的网络代理的方式与应用的代码部署在一起，用于保证服务与服务之间调用的可靠性，与传统的微服务架构有着本质的区别：  
> 1、跨语言服务调用的需要  
> 2、云原生应用服务治理的需要

Service Mesh 思想的孕育而生，一方面出于各大公司微服务技术的普及，增加了对跨语言服务调用的需求；另一方面得益于微服务容器化后，采用 Kubernetes 等云平台部署的云原生应用越来越多，服务治理的需求也越来越强烈。

### ServiceMesh 实现原理
Service Mesh 实现的关键就在于两点：一个是轻量级的网络代理也叫 SideCar，它的作用就是转发服务之间的调用；一个是基于 SideCar 的服务治理也被叫作 Control Plane，它的作用是向 SideCar 发送各种指令，以完成各种服务治理功能。  

- SideCar
> 服务框架的功能都集中实现在 SideCar 里，并在每一个服务消费者和服务提供者的本地都部署一个 SideCar，服务消费者和服务提供者只管自己的业务实现，服务消费者向本地的 SideCar 发起请求，本地的 SideCar 根据请求的路径向注册中心查询，得到服务提供者的可用节点列表后，再根据负载均衡策略选择一个服务提供者节点，并向这个节点上的 SideCar 转发请求，服务提供者节点上的 SideCar 完成流量统计、限流等功能后，再把请求转发给本地部署的服务提供者进程，从而完成一次服务请求。  
> 把服务消费者节点上的 SideCar 叫作正向代理，服务提供者节点上的 SideCar 叫作反向代理，那么 Service Mesh 架构的关键点就在于服务消费者发出的请求如何通过正向代理转发以及服务提供者收到的请求如何通过反向代理转发：  
> 1、基于 iptables 的网络拦截  
> 2、采用协议转换的方式

![SideCar](./images/service-mesh-sidecar.png)

- Control Plane  
> 通过 Control Plane 与各个 SideCar 交互，能控制网格中流量的运转。Control Plane 的主要作用：  
> 1、服务发现。服务提供者会通过 SideCar 注册到 Control Plane 的注册中心，这样的话服务消费者把请求发送给 SideCar 后，SideCar 就会查询 Control Plane 的注册中心来获取服务提供者节点列表。  
> 2、负载均衡。SideCar 从 Control Plane 获取到服务提供者节点列表信息后，就需要按照一定的负载均衡算法从可用的节点列表中选取一个节点发起调用，可以通过 Control Plane 动态修改 SideCar 中的负载均衡配置。  
> 3、请求路由。SideCar 从 Control Plane 获取的服务提供者节点列表，也可以通过 Control Plane 来动态改变，比如需要进行 A/B 测试、灰度发布或者流量切换时，就可以动态地改变请求路由。  
> 4、故障处理。服务之间的调用如果出现故障，就需要加以控制，通常的手段有超时重试、熔断等，这些都可以在 SideCar 转发请求时，通过 Control Plane 动态配置。  
> 5、安全认证。可以通过 Control Plane 控制一个服务可以被谁访问，以及访问哪些信息。  
> 6、监控上报。所有 SideCar 转发的请求信息，都会发送到 Control Plane，再由 Control Plane 发送给监控系统，比如 Prometheus 等。  
> 7、日志记录。所有 SideCar 转发的日志信息，也会发送到 Control Plane，再由 Control Plane 发送给日志系统，比如 Stackdriver 等。  
> 8、配额控制。可以在 Control Plane 里给服务的每个调用方配置最大调用次数，在 SideCar 转发请求给某个服务时，会审计调用是否超出服务对应的次数限制。  

### Istio：Service Mesh 产品代表 
Istio 的架构可以说由两部分组成，分别是 Proxy 和 Control Plane。  
> Proxy，就是 SideCar，与应用程序部署在同一个主机上，应用程序之间的调用都通过 Proxy 来转发，目前支持 HTTP/1.1、HTTP/2、gRPC 以及 TCP 请求。  
> Control Plane，与 Proxy 通信，来实现各种服务治理功能，包括三个基本组件：Pilot、Mixer 以及 Citadel。  

![Istio 架构](./images/istio-framework.png)  

相比 Linkerd，Istio 引入了 Control Plane 的理念，通过 Control Plane 能带来强大的服务治理能力，可以称得上是 Linkerd 的进化，算是第二代的 Service Mesh 产品。  
Istio 默认的 SideCar 采用了Envoy，它是用 C++ 语言实现的，在性能和资源消耗上要比采用 Scala 语言实现的 Linkerd 小，这一点对于延迟敏感型和资源敏感型的服务来说，尤其重要。  
有 Google 和 IBM 的背书，尤其是在微服务容器化的大趋势下，云原生应用越来越受欢迎，而 Google 开源的 Kubernetes 可以说已经成为云原生应用默认采用的容器平台，基于此 Google 可以将 Kubernetes 与 Istio 很自然的整合，打造成云原生应用默认的服务治理方案。  





