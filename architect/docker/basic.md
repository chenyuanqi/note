
### Docker 是什么
Docker 是容器技术的一种，业界公认的容器标准。  

容器翻译自英文的 Container 一词，而 Container 又可以翻译成集装箱。在软件的世界里，容器封装的是软件的运行环境。容器的本质就是 Linux 操作系统里的进程，但与操作系统中运行的一般进程不同的是，容器通过 Namespace 和 Cgroups 这两种机制，可以拥有自己的 root 文件系统、自己的网络配置、自己的进程空间，甚至是自己的用户 ID 空间，这样的话容器里的进程就像是运行在宿主机上的另外一个单独的操作系统内，从而实现与宿主机操作系统里运行的其他进程隔离。  
一句话概括容器：容器就是将软件打包成标准化单元，以用于开发、交付和部署。  

- 容器镜像是轻量的、可执行的独立软件包，包含软件运行所需的所有内容：代码、运行时环境、系统工具、系统库和设置。  
- 容器化软件适用于基于 Linux 和 Windows 的应用，在任何环境中都能够始终如一地运行。  
- 容器赋予了软件独立性，使其免受外在环境差异（例如，开发和预演环境的差异）的影响，从而有助于减少团队间在相同基础设施上运行不同软件时的冲突。  
- 容器虚拟化的是操作系统而不是硬件，容器之间是共享同一套操作系统资源的。虚拟机技术是虚拟出一套硬件后，在其上运行一个完整操作系统。因此容器的隔离级别会稍低一些。  

虽然容器解决了应用程序运行时隔离的问题，但是要想实现应用能够从一台机器迁移到另外一台机器上还能正常运行，就必须保证另外一台机器上的操作系统是一致的，而且应用程序依赖的各种环境也必须是一致的。  

Docker —— 代码集装箱装卸工。  
Docker 容器的特点：轻量、标准、安全。  

- Docker 是世界领先的软件容器平台。  
- Docker 使用 Google 公司推出的 Go 语言 进行开发实现，基于 Linux 内核 的 cgroup，namespace，以及 AUFS 类的 UnionFS 等技术， 对进程进行封装隔离，属于操作系统层面的虚拟化技术。 由于隔离的进程独立于宿主和其它的隔离的进程，因此也称其为容器。Docker 最初实现是基于 LXC。  
- Docker 能够自动执行重复性任务，例如搭建和配置开发环境，从而解放了开发人员以便他们专注在真正重要的事情上：构建杰出的软件。
- 用户可以方便地创建和使用容器，把自己的应用放入容器。容器还可以进行版本管理、复制、分享、修改，就像管理普通的代码一样。

Docker 镜像不光可以打包应用程序本身，而且还可以打包应用程序的所有依赖，甚至可以包含整个操作系统。  
Docker 镜像解决了 DevOps 中微服务运行的环境难以在本地环境、测试环境以及线上环境保持一致的难题。开发可以把在本地环境中运行测试通过的代码，以及依赖的软件和操作系统本身打包成一个镜像，然后自动部署在测试环境中进行测试，测试通过后再自动发布到线上环境上去，整个开发、测试和发布的流程就打通了。  
无论是使用内部物理机还是公有云的机器部署服务，都可以利用 Docker 镜像把微服务运行环境封装起来，从而屏蔽机器内部物理机和公有云机器运行环境的差异，实现同等对待，降低了运维的复杂度。  

[Docker 镜像仓库](https://hub.docker.com/)  

### 为什么 Docker
每当说起容器，我们不得不将其与虚拟机做一个比较。实际上，对于两者无所谓谁会取代谁，而是两者可以和谐共存。  

- Docker 的镜像提供了除内核外完整的运行时环境，确保了应用运行环境一致性，从而不会再出现 “这段代码在我机器上没问题啊” 这类问题；—— 一致的运行环境
- 可以做到秒级、甚至毫秒级的启动时间。大大的节约了开发、测试、部署的时间。—— 更快速的启动时间
- 避免公用的服务器，资源会容易受到其他用户的影响。—— 隔离性
- 善于处理集中爆发的服务器使用压力；—— 弹性伸缩，快速扩展
- 可以很轻易的将在一个平台上运行的应用，迁移到另一个平台上，而不用担心运行环境的变化导致应用无法正常运行的情况。—— 迁移方便
- 使用 Docker 可以通过定制应用镜像来实现持续集成、持续交付、部署。—— 持续交付和部署

### Docker 基本概念
Docker 包括三个基本概念：镜像（Image）、容器（Container）、仓库（Repository）。理解了这三个概念，就理解 Docker 的整个生命周期。  

**镜像 (Image): 一个特殊的文件系统**  
操作系统分为内核和用户空间。对于 Linux 而言，内核启动后，会挂载 root 文件系统为其提供用户空间支持。而 Docker 镜像（Image），就相当于是一个 root 文件系统。  
Docker 镜像是一个特殊的文件系统，除了提供容器运行时所需的程序、库、资源、配置等文件外，还包含了一些为运行时准备的一些配置参数（如匿名卷、环境变量、用户等）。 镜像不包含任何动态数据，其内容在构建之后也不会被改变。  
Docker 设计时，就充分利用 Union FS 的技术，将其设计为分层存储的架构。镜像实际是由多层文件系统联合组成。  

镜像构建时，会一层层构建，前一层是后一层的基础。每一层构建完就不会再发生改变，后一层上的任何改变只发生在自己这一层。　比如，删除前一层文件的操作，实际不是真的删除前一层的文件，而是仅在当前层标记为该文件已删除。在最终容器运行的时候，虽然不会看到这个文件，但是实际上该文件会一直跟随镜像。因此，在构建镜像的时候，需要额外小心，每一层尽量只包含该层需要添加的东西，任何额外的东西应该在该层构建结束前清理掉。  
分层存储的特征还使得镜像的复用、定制变的更为容易。甚至可以用之前构建好的镜像作为基础层，然后进一步添加新的层，以定制自己所需的内容，构建新的镜像。  

**容器 (Container): 镜像运行时的实体**  
镜像（Image）和容器（Container）的关系，就像是面向对象程序设计中的 类和实例一样，镜像是静态的定义，容器是镜像运行时的实体。容器可以被创建、启动、停止、删除、暂停等。  

容器的实质是进程，但与直接在宿主执行的进程不同，容器进程运行于属于自己的独立的命名空间。镜像使用的是分层存储，容器也是如此。  
容器存储层的生存周期和容器一样，容器消亡时，容器存储层也随之消亡。因此，任何保存于容器存储层的信息都会随容器删除而丢失。  

按照 Docker 最佳实践的要求，容器不应该向其存储层内写入任何数据 ，容器存储层要保持无状态化。所有的文件写入操作，都应该使用数据卷（Volume）、或者绑定宿主目录，在这些位置的读写会跳过容器存储层，直接对宿主 (或网络存储) 发生读写，其性能和稳定性更高。数据卷的生存周期独立于容器，容器消亡，数据卷不会消亡。因此，使用数据卷后，容器可以随意删除、重新 run ，数据却不会丢失。  

**仓库 (Repository): 集中存放镜像文件的地方**  
镜像构建完成后，可以很容易的在当前宿主上运行，但是，如果需要在其它服务器上使用这个镜像，我们就需要一个集中的存储、分发镜像的服务，Docker Registry 就是这样的服务。  
一个 Docker Registry 中可以包含多个仓库（Repository）；每个仓库可以包含多个标签（Tag）；每个标签对应一个镜像。所以说：镜像仓库是 Docker 用来集中存放镜像文件的地方类似于我们之前常用的代码仓库。  

通常，一个仓库会包含同一个软件不同版本的镜像，而标签就常用于对应该软件的各个版本。我们可以通过 <仓库名>:<标签> 的格式来指定具体是这个软件哪个版本的镜像。如果不给出标签，将以 latest 作为默认标签。

Docker Registry 公开服务是开放给用户使用、允许用户管理镜像的 Registry 服务。一般这类公开服务允许用户免费上传、下载公开的镜像，并可能提供收费服务供用户管理私有镜像。  
> 最常使用的 Registry 公开服务是官方的 [Docker Hub](https://hub.docker.com/)，也是默认的 Registry，并拥有大量的高质量的官方镜像。  
> 在国内访问 Docker Hub 可能会比较慢，国内也有一些云服务商提供类似于 Docker Hub 的公开服务，比如[时速云镜像库](https://hub.tenxcloud.com/)、[网易云镜像服务](https://www.163yun.com/product/repo)、 [DaoCloud 镜像市场](https://www.daocloud.io/)、[阿里云镜像库](https://www.aliyun.com/product/containerservice?utm_content=se_1292836)等。  

除了使用公开服务外，用户还可以在本地搭建私有 Docker Registry。Docker 官方提供了 Docker Registry 镜像，可以直接使用做为私有 Registry 服务。开源的 Docker Registry 镜像只提供了 Docker Registry API 的服务端实现，足以支持 docker 命令，不影响使用。但不包含图形界面，以及镜像维护、用户管理、访问控制等高级功能。


Docker 运行过程也就是去仓库把镜像拉到本地，然后用一条命令把镜像运行起来变成容器。所以，我们也常常将 Docker 称为码头工人或码头装卸工，这和 Docker 的中文翻译搬运工人如出一辙。  

- Build（构建镜像）： 镜像就像是集装箱包括文件以及运行环境等等资源。
- Ship（运输镜像）：主机和仓库间运输，这里的仓库就像是超级码头一样。
- Run （运行镜像）：运行的镜像就是一个容器，容器就是运行程序的地方。

### Docker 安装
[windows](https://docs.docker.com/docker-for-windows/install/) 和 [macos](https://docs.docker.com/docker-for-mac/install/) 的安装比较简单，只要下载对应的软件，一直 Next 即可。  
```bash
# ubuntu
# 
# 卸载已经安装的 docker
sudo apt-get remove docker docker-engine docker.io
# 设置镜像仓库
sudo apt-get update
sudo apt-get install apt-transport-https ca-certificates curl software-properties-common
curl -fsSL https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu/gpg | sudo apt-key add - 
sudo add-apt-repository 
"deb [arch=amd64] https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu $(lsb_release -cs) stable"
sudo apt update 
# 安装 docker-ce【社区版】
sudo apt install docker-ce
# 验证
docker run helloworld
docker --version
```

**设置国内镜像**  
docker 默认的配置文件在 /etc/docker/daemon.json，选择网易、科大的镜像。   
```json
{
  "registry-mirrors": ["https://docker.mirrors.ustc.edu.cn","http://hub-mirror.c.163.com"]
}
```

**设置免 sudo**  
默认安装完 docker 后，每次执行 docker 都需要运行 sudo 命令。
```bash
sudo groupadd docker
sudo gpasswd -a ${USER} docker
sudo service docker restart
newgrp - docker
```

**docker 端口和本地宿主机器端口映射**  
假设容器的名字叫 ubuntu。  
1、对容器暴露所有的端口，随机映射宿主机端口
```bash
docker run -P -it ubuntu /bin/bash
```
2、映射宿主机随机端口到容器指定的端口
```bash
docker run -p 80 -it ubuntu /bin/bash
```
3、映射宿主机的指定端口到容器指定端口 1 对 1
```bash
docker run -p 8080:8080 -it ubuntu /bin/bash  
```
4、指定容器 ip 和容器端口，宿主机端口随机映射
```bash
docker run -p 127.0.0.1::80 -it ubuntu /bin/bash  
```

### Docker 基本操作
```bash
# 查看运行的容器
docker ps
# 停止容器
docker stop container-id
# 进入容器
docker exec -it [CONTAINER-ID] /bin/sh
# 退出
exit
```






