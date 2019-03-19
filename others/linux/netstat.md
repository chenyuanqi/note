
### netstat 基本用法
Netstat 命令用于显示各种网络相关信息，如网络连接，路由表，接口状态 (Interface Statistics)，masquerade 连接，多播成员 (Multicast Memberships) 等等。  
```bash
$ netstat
# -a (all) 显示所有选项，默认不显示 LISTEN 相关(LISTEN 和 LISTENING 的状态只有用 - a 或者 - l 才能看到)
# -t (tcp) 仅显示 tcp 相关选项
# -u (udp) 仅显示 udp 相关选项
# -n 拒绝显示别名，能显示数字的全部转化成数字。
# -l 仅列出有在 Listen (监听) 的服務状态
# -p 显示建立相关链接的程序名
# -r 显示路由信息，路由表
# -e 显示扩展信息，例如 uid 等
# -s 按各个协议进行统计
# -c 每隔一个固定时间，执行该 netstat 命令。

Active Internet connections (w/o servers)
Proto Recv-Q Send-Q Local Address Foreign Address State
tcp 0 2 210.34.6.89:telnet 210.34.6.96:2873 ESTABLISHED
tcp 296 0 210.34.6.89:1165 210.34.6.84:netbios-ssn ESTABLISHED
tcp 0 0 localhost.localdom:9001 localhost.localdom:1162 ESTABLISHED
tcp 0 0 localhost.localdom:1162 localhost.localdom:9001 ESTABLISHED
tcp 0 80 210.34.6.89:1161 210.34.6.10:netbios-ssn CLOSE

Active UNIX domain sockets (w/o servers)
Proto RefCnt Flags Type State I-Node Path
unix 1 [ ] STREAM CONNECTED 16178 @000000dd
unix 1 [ ] STREAM CONNECTED 16176 @000000dc
unix 9 [ ] DGRAM 5292 /dev/log
unix 1 [ ] STREAM CONNECTED 16182 @000000df
```
从整体上看，netstat 的输出结果可以分为两个部分：  
一个是 Active Internet connections，称为有源 TCP 连接，其中 "Recv-Q" 和 "Send-Q" 指 %0A 的是接收队列和发送队列。这些数字一般都应该是 0。如果不是则表示软件包正在队列中堆积。这种情况只能在非常少的情况见到。  
另一个是 Active UNIX domain sockets，称为有源 Unix 域套接口 (和网络套接字一样，但是只能用于本机通信，性能可以提高一倍)。  
Proto 显示连接使用的协议，RefCnt 表示连接到本套接口上的进程号，Types 显示套接口的类型，State 显示套接口当前的状态，Path 表示连接到套接口的其它进程使用的路径名。  

```bash
# 列出所有端口 (包括监听和未监听的)
$ netstat -a

# 列出所有 tcp 端口
$ netstat -at

# 列出所有 udp 端口
$ netstat -au

# 列出所有处于监听状态的 Sockets
# 只显示监听端口 
$ netstat -l

#  只列出所有监听 tcp 端口
$ netstat -lt

# 只列出所有监听 udp 端口
$ netstat -lu

# 只列出所有监听 UNIX 端口
$ netstat -lx

# 显示每个协议的统计信息
# 显示所有端口的统计信息 
$ netstat -s

# 显示 TCP 或 UDP 端口的统计信息 
$ netstat -st 
$ netstat -su

# 在 netstat 输出中显示 PID 和进程名称 
$ netstat -p

# 在 netstat 输出中不显示主机，端口和用户名 (host, port or user)
$ netstat -n

# 持续输出 netstat 信息(每隔一秒输出网络信息)
$ netstat -c

# 显示系统不支持的地址族 
$ netstat --verbose

# 显示核心路由信息 
$ netstat -r

# 显示网络接口列表
$ netstat -i
# 显示详细信息
$ netstat -ie

# TCP 各种状态列表
$ netstat -nat |awk '{print $6}'|sort|uniq -c|sort -rn

# 查看连接某服务端口最多的的 IP 地址
$ netstat -nat | grep "192.168.1.15:22" |awk '{print $5}'|awk -F: '{print $1}'|sort|uniq -c|sort -nr|head -20

# 找出程序运行的端口
$ netstat -ap | grep ssh
$ netstat -an | grep ':80'

# 分析 access.log 获得访问前 10 位的 ip 地址
$ awk '{print $1}' access.log |sort|uniq -c|sort -nr|head -10
```
