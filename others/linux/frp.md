
### frp 内网穿透
frp 是一个可用于内网穿透的高性能的反向代理应用，支持 tcp, udp 协议，为 http 和 https 应用协议提供了额外的能力，且尝试性支持了点对点穿透。

为什么要内网穿透呢？  
> 利用处于内网或防火墙后的机器，对外网环境提供 http 或 https 服务。  
> 利用处于内网或防火墙后的机器，对外网环境提供 tcp 和 udp 服务，例如在家里通过 ssh 访问处于公司内网环境内的主机。  
> 对于 http, https 服务支持基于域名的虚拟主机，支持自定义域名绑定，使多个域名可以共用一个 80 端口。  

[frp github](https://github.com/fatedier/frp)  
[frp 文档](https://github.com/fatedier/frp/blob/master/README_zh.md)  

### frp 服务端配置
具有公网 IP 的机器上，执行如下操作
```bash
wget https://github.com/fatedier/frp/releases/download/v0.28.0/frp_0.28.0_linux_amd64.tar.gz
tar -xzvf frp_0.28.0_linux_amd64.tar.gz
```
解压之后有如下 7 个文件
```
frpc # 客户端 linux 程序
frpc_full.ini # 客户端 所有配置
frpc.ini
frps # 服务端 linux 程序
frps_full.ini # 服务端 所有配置
frps.ini
LICENSE
```

修改服务端配置信息 frps.ini
```
[common]
bind_port = 7000
# auth token
token = frps_2019
```
服务端配置需要有一个外网可以访问的 ip，假设我的 ip 是 111.111.111.111, 当前服务对外端口为 7000
```bash
# 先尝试执行
./frps -c frps.ini
# 后台运行，退出则找出进程好 ps aux | grep frps 再 kill 掉
nohup ./frps -c frps.ini >> frps.log 2>&1 &
```

### frp 客户端配置
内网环境的机器上，比如 windows 系统就下载 frp 的 windows 版本。然后修改客户端配置信息 frpc.ini
```
[common]
server_addr = 111.111.111.111 # 服务端对外 IP
server_port = 7000 # 服务端口号
token = frps_2019 # 保持跟服务端 token 一致

[ssh]
type = tcp
local_ip = 127.0.0.1
local_port = 22
remote_port = 6000 # 服务器端对外提供本机服务的端口号
```
客户端服务在后台运行
```bash
# 启动后，服务端产生日志 [ssh] tcp proxy listen port [6000]
./frpc -c frpc.ini >> frpc.log 2>&1 &
```

### 外网测试
```bash
ssh root@111.111.111.111 -p 6000
```
