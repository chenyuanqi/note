
### Docker 镜像
镜像作为 Docker 三大核心概念中，最重要的一个关键词，它有很多操作，是学习容器技术不得不掌握的关键。  

**下载镜像**  
镜像是运行容器的前提，我们可以使用 docker pull[IMAGE_NAME]:[TAG] 命令来下载镜像（不显式指定 TAG, 会默认下载 latest 标签）。  
`注意：不推荐下载 latest 标签，因为该镜像的内容会跟踪镜像的最新版本，并随之变化，所以它是不稳定的。在生产环境中，可能会出现莫名其妙的 bug, 推荐最好还是显示的指定具体的 TAG。`  

在使用 docker pull 命令时，还需要在镜像前面指定仓库地址 ( Registry), 如果不指定，则 Docker 会使用您默认配置的仓库地址。  
严格意义上，我们在使用 docker pull 命令时，还需要在镜像前面指定仓库地址 ( Registry), 如果不指定，则 Docker 会使用您默认配置的仓库地址。Docker 通过前缀地址的不同，来保证不同仓库中，重名镜像的唯一性。  
```bash
# 下载 mysql 5.7 镜像，如下命令相当于 docker pull Registry/mysql:5.7
# -a,--all-tags=true|false: 是否获取仓库中所有镜像，默认为否；
# --disable-content-trust: 跳过镜像内容的校验，默认为 true;
docker pull mysql:5.7

# 查看本地已有镜像，显示的字段参数如下：
# REPOSITORY: 来自于哪个仓库；
# TAG: 镜像的标签信息，比如 5.7、latest 表示不同的版本信息；
# IMAGE ID: 镜像的 ID, 如果您看到两个 ID 完全相同，那么实际上，它们指向的是同一个镜像，只是标签名称不同罢了；
# CREATED: 镜像最后的更新时间；
# SIZE: 镜像的大小，优秀的镜像一般体积都比较小，这也是我更倾向于使用轻量级的 alpine 版本的原因；
# 镜像大小信息只是逻辑上的大小信息，因为一个镜像是由多个镜像层（ layer）组成的，而相同的镜像层本地只会存储一份。
# 所以，真实情况下，占用的物理存储空间大小，可能会小于逻辑大小
docker images # 或 docker image ls
```

一个镜像一般是由多个层（ layer） 组成，如果多个不同的镜像中，同时包含了同一个层（ layer）,Docker 在下载之前，会去检测本地是否会有同样 ID 的层，如果本地已经存在了，就直接使用本地的就好了。  

**为镜像添加标签**  
通常情况下，为了方便在后续工作中，快速地找到某个镜像，我们可以使用 docker tag 命令，为本地镜像添加一个新的标签。docker tag 命令功能更像是，为指定镜像添加快捷方式一样。
```bash
docker tag mysql:5.7 vikey_mysql:5.7
```

**查看镜像详细信息**  
通过 docker inspect 命令，我们可以获取镜像的详细信息，其中，包括创建者，各层的数字摘要等。  
docker inspect 返回的是 JSON 格式的信息，如果您想获取其中指定的一项内容，可以通过 -f 来指定。
```bash
docker inspect mysql:5.7
# 获取镜像大小
docker inspect -f {{".Size"}} mysql:5.7
```

**查看镜像历史**  
通过 docker history 命令，可以列出各个层（layer）的创建信息。  
```bash
docker history mysql:5.7
# 查看具体信息
docker history --no-trunc mysql:5.7
```

### Docker 镜像查找
可以通过 docker search [option] keyword 命令进行搜索。
```bash
# -f,--filter filter: 过滤输出的内容；
# --limitint：指定搜索内容展示个数；
# --no-index: 不截断输出内容；
# --no-trunc：不截断输出内容；
docker search mysql
# 搜索官方提供的 mysql 镜像
docker search --filter=is-offical=true mysql
# 搜索 Stars 数超过 100 的 mysql 镜像
docker search --filter=stars=100 mysql
```

### Docker 镜像删除
通过 docker rmi [image] 或 docker image rm [image] 命令都可以删除镜像。  
```bash
# -f,-force: 强制删除镜像，即便有容器引用该镜像；
# -no-prune: 不要删除未带标签的父镜像；
docker rmi vikey_mysql:5.7
```
`实际上，当同一个镜像拥有多个标签时，执行 docker rmi 命令，只是会删除了该镜像众多标签中，您指定的标签而已，并不会影响原始的那个镜像文件。`

**通过镜像 ID 删除镜像**  
除了通过标签名称来删除镜像，我们还可以通过制定镜像 ID 来删除镜像。  
```bash
docker rmi ee7cbd482336
```
`一旦指定了通过 ID 来删除镜像，它会先尝试删除所有指向该镜像的标签，然后在删除镜像本身。`  

**删除镜像的限制**  
删除镜像很简单，但也不是我们何时何地都能删除的，它存在一些限制条件：  
当通过该镜像创建的容器未被销毁时，镜像是无法被删除的（除非通过添加 -f 子命令）；这时候，正确的做法是先删除引用这个镜像的容器，再删除这个镜像。  

**清理镜像**  
在使用 Docker 一段时间后，系统一般都会残存一些临时的、没有被使用的镜像文件，可以通过以下命令进行清理。
```bash
# -a,--all: 删除所有没有用的镜像，而不仅仅是临时文件；
# -f,--force：强制删除镜像文件，无需弹出提示确认；
docker image prune
```

### Docker 镜像创建
Docker 创建镜像主要有三种：  

- 基于已有的镜像创建；
```bash
# 运行 ubuntu 镜像
docker run -it docker.io/ubuntu:latest /bin/bash
# 创建一个测试文件
touch test.txt
# 创建镜像
# -a,--author="": 作者信息；
# -c,--change=[]: 可以在提交的时候执行 Dockerfile 指令，如 CMD、ENTRYPOINT、ENV、EXPOSE、LABEL、ONBUILD、USER、VOLUME、WORIR 等；
# -m,--message="": 提交信息；
# -p,--pause=true: 提交时，暂停容器运行。
docker container commit -m "Added test.txt file" -a "Allen" a0a0c8cfec3a test:0.1
```
- 基于 Dockerfile 来创建；
> 通过 Dockerfile 的方式来创建镜像，是最常见的一种方式了，也是比较推荐的方式。  
> Dockerfile 是一个文本指令文件，它描述了是如何基于一个父镜像，来创建一个新镜像的过程。  

```bash
# Dockerfile 内容
FROM docker.io/ubuntu:latest
LABEL version="1.0" maintainer="Allen <weiwosuo@github>"
RUN apt-get update && \    apt-get install -y python3 && \    apt-get clean && \    rm -rf /var/lib/apt/lists/*

# 构建镜像
docker image build -t python:3 .
```
- 基于本地模板来导入；

### Docker 镜像导出与加载
通常需要将镜像分享给别人，可以将镜像导出成 tar 包，别人直接通过加载这个 tar 包，快速地将镜像引入到本地镜像库。  
```bash
docker save -o python_3.tar python:3
```

别人拿到了这个 tar 包后，再导入到本地的镜像库。  
```bash
docker load -i python_3.tar # 或 docker load < python_3.tar
```

### Docker 镜像上传
以上传到 Docker Hub 上为示例。  
想要上传镜像到 Docker Hub 上，首先，我们需要注册 Docker Hub 账号。打开 [Docker Hub](https://hub.docker.com)，填写 Docker ID (即账号)、密码、Email 完成注册，然后登陆到主界面。  
选择 Create a Respository 填写相关信息，创建仓库。  
```bash
# 在命令行登陆 Docker Hub
docker login
# 打 tag
docker tag python:3 vikey/python:3
# 查看本地镜像是否存在 tag
docker images
# 上传
docker push vikey/python:3
```
