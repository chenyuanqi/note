
### 什么是 Kubernetes
使用 Web 服务，用户希望应用程序能够 7\*24 小时全天运行，开发人员希望每天多次部署新的应用版本。  
通过应用容器化可以实现这些目标，使应用简单、快捷的方式更新和发布，也能实现热更新、迁移等操作。使用 Kubernetes 能确保程序在任何时间、任何地方运行，还能扩展更多有需求的工具 / 资源。Kubernetes 积累了 Google 在容器化应用业务方面的经验，以及社区成员的实践，是能在生产环境使用的开源平台。  

一个“容器”，实际上是一个由 Linux Namespace、Linux Cgroups 和 rootfs 三种技术构建出来的进程的隔离环境。  
Kubernetes 是容器集群管理系统，是一个开源的平台，可以实现容器集群的自动化部署、自动扩缩容、维护等功能。Kubernetes 项目的本质，是为用户提供一个具有普遍意义的容器编排工具。  

Kubernetes 的名字来自希腊语，意思是“舵手” 或 “领航员”。K8s 是将中间的 8 个字母 “ubernete” 替换为 “8” 的缩写。

**Kubernetes 特点**

- 可移植: 支持公有云，私有云，混合云，多重云（multi-cloud）
- 可扩展: 模块化, 插件化, 可挂载, 可组合
- 自动化: 自动部署，自动重启，自动复制，自动伸缩/扩展

**Kubernets 技能图谱**   
![Kubernets 技能图谱](../../../others/static/images/kubernets-tech-graph.png)  

### 为什么 Kubernets


### Kubernets 基础概念


### Kubernets 部署
Kubernetes 作为一个 Golang 项目，已经免去了很多类似于 Python 项目要安装语言级别依赖的麻烦。但是，除了将各个组件编译成二进制文件外，用户还要负责为这些二进制文件编写对应 的配置文件、配置自启动脚本，以及为 kube-apiserver 配置授权文件等等诸多运维工作。  
目前，各大云厂商常用的部署的方法，是使用 SaltStack、Ansible 等运维工具自动化地执行这些步骤。但是，这个部署过程依然非常繁琐。


