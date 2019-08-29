
### Docker 是什么
Docker 是容器技术的一种，业界公认的容器标准。  

容器翻译自英文的 Container 一词，而 Container 又可以翻译成集装箱。在软件的世界里，容器封装的是软件的运行环境。容器的本质就是 Linux 操作系统里的进程，但与操作系统中运行的一般进程不同的是，容器通过 Namespace 和 Cgroups 这两种机制，可以拥有自己的 root 文件系统、自己的网络配置、自己的进程空间，甚至是自己的用户 ID 空间，这样的话容器里的进程就像是运行在宿主机上的另外一个单独的操作系统内，从而实现与宿主机操作系统里运行的其他进程隔离。  

虽然容器解决了应用程序运行时隔离的问题，但是要想实现应用能够从一台机器迁移到另外一台机器上还能正常运行，就必须保证另外一台机器上的操作系统是一致的，而且应用程序依赖的各种环境也必须是一致的。  
Docker 镜像不光可以打包应用程序本身，而且还可以打包应用程序的所有依赖，甚至可以包含整个操作系统。  
Docker 镜像解决了 DevOps 中微服务运行的环境难以在本地环境、测试环境以及线上环境保持一致的难题。开发可以把在本地环境中运行测试通过的代码，以及依赖的软件和操作系统本身打包成一个镜像，然后自动部署在测试环境中进行测试，测试通过后再自动发布到线上环境上去，整个开发、测试和发布的流程就打通了。  
无论是使用内部物理机还是公有云的机器部署服务，都可以利用 Docker 镜像把微服务运行环境封装起来，从而屏蔽机器内部物理机和公有云机器运行环境的差异，实现同等对待，降低了运维的复杂度。  

[Docker 镜像仓库](https://hub.docker.com/)  

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






