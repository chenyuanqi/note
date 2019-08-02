
### ifconfig
在 CentOS 7 中 ifconfig 命令已经不灵了，换成了 ip 命令工具；可以安装工具 net-tools 将 ifconfig 等网络命令找回来。
```bash
yum install net-tools
ifconfig
```

网络相关目录
```
/etc/sysconfig/network # 包括主机基本网络信息，用于系统启动
/etc/sysconfig/network-script/ # 此目录下是系统启动最初始化网络的信息
/etc/sysconfig/network-script/ifcfg-em1 # 网络配置信息，每个人的配置名字不一样通过命令查看
/etc/xinetd.conf 定义了由超级进程 XINETD 启动的网络服务
/etc/protocols # 设定了主机使用的协议以及各个协议的协议号
/etc/services  # 设定了主机的不同端口的网络服务
```

### ip
ip [选项] 操作对象 {link|addr|route...}
```bash
ip link show                    # 显示网络接口信息
ip link set eth0 upi            # 开启网卡
ip link set eth0 down           # 关闭网卡
ip link set eth0 promisc on     # 开启网卡的混合模式
ip link set eth0 promisc offi   # 关闭网卡的混个模式
ip link set eth0 txqueuelen 1200 # 设置网卡队列长度
ip link set eth0 mtu 1400        # 设置网卡最大传输单元
ip addr show     # 显示网卡IP信息
ip addr add 192.168.0.1/24 dev eth0 # 设置eth0网卡IP地址192.168.0.1
ip addr del 192.168.0.1/24 dev eth0 # 删除eth0网卡IP地址

ip route list                 # 查看路由信息
ip route add 192.168.4.0/24  via  192.168.0.254 dev eth0 # 设置192.168.4.0网段的网关为192.168.0.254,数据走eth0接口
ip route add default via  192.168.0.254  dev eth0        # 设置默认网关为192.168.0.254
ip route del 192.168.4.0/24   # 删除192.168.4.0网段的网关
ip route del default          # 删除默认路由
```
通过 ip addr 命令 查看 ip。看到两个配置，lo 和 em1 ，lo 代表 127.0.0.1，即 localhost。eth0 这个是你的网卡，如果上面没有 inet 字段后面跟着 IP 的话，你需要去配置文件中修改配置。

默认配置文件在这里 /etc/sysconfig/network-scripts/ifcfg-eth0 ，一般情况这个配置文件是 ifcfg-<网卡名字> 加上你网卡名字。
```
vim /etc/sysconfig/network-scripts/ifcfg-eth0 
```
主要更改这三项：TYPE=Ethernet、BOOTPROTO=dhcp、ONBOOT=yes
```
# TYPE：配置文件接口类型。在 /etc/sysconfig/network-scripts/ 目录有多种网络配置文件，有 Ethernet 、IPsec 等类型，网络接口类型为 Ethernet
TYPE=Ethernet    

BOOTPROTO=dhcp     # 自动获取 IP
# BOOTPROTO=none   # 静态配置，若该值为 “dhcp” 则为动态获得，另外 static 也是表示静态ip地址
DEFROUTE=yes
PEERDNS=yes
PEERROUTES=yes
IPV4_FAILURE_FATAL=no
IPV6INIT=yes        # 是否执行 IPv6。yes：支持IPv6、no：不支持IPv6。
IPV6_AUTOCONF=yes
IPV6_DEFROUTE=yes
IPV6_PEERDNS=yes
IPV6_PEERROUTES=yes
IPV6_FAILURE_FATAL=no
NAME=em1                # 网卡名称
UUID=99-6217--a******   # UUID 号，没事不要动它，否则你会后悔的
DEVICE=em1              # 接口名称
ONBOOT=yes              # 设置网络开机自动启动

IPADDR=<这里固定IP配置的地方>  # 设置 IP 地址
PREFIX=24                   # 设置子网掩码
GATEWAY=<这里设置网关>        # 设置网关
DNS1=<这里设置DNS>           # DNS

DNS1=8.8.8.8            # 设置主 DNS
DNS2=8.8.4.4            # 设置备 DNS
```

网卡启动、关闭
```bash
ifup <设备名/网卡名字>    # 激活网卡
ifdown <设备名/网卡名字>  # 关闭网卡
```

### 网络服务
```bash
service network start    # 启动网络服务
service network stop     # 停止网络服务
service network restart  # 重启网络服务
service network status   # 查看网络服务状态
nmcli dev status         # 检查受网络管理器管理的网络接口
systemctl status NetworkManager.service # 验证网络管理器服务的状态
```

名词解析  
> wlan0 表示第一块无线以太网卡  
> Link encap 表示该网卡位于 OSI 物理层 (Physical Layer）的名称  
> HWaddr 表示网卡的 MAC 地址（Hardware Address）  
> inet addr 表示该网卡在 TCP/IP 网络中的 IP 地址  
> Bcast 表示广播地址（Broad Address）  
> Mask 表示子网掩码（Subnet Mask）  
> MTU 表示最大传送单元，不同局域网 MTU 值不一定相同，对以太网来说，MTU 的默认设置是 1500 个字节  
> Metric 表示度量值，通常用于计算路由成本  
> RX 表示接收的数据包  
> TX 表示发送的数据包  
> collisions 表示数据包冲突的次数  
> txqueuelen 表示传送列队（Transfer Queue）长度  
> interrupt 表示该网卡的 IRQ 中断号  
> Base address 表示 I/O 地址  