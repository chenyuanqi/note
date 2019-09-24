
### 什么是堡垒机
堡垒机也叫跳板机，是指在一个特定的网络环境下，为了保障网络和数据不受来自外部和内部用户的操控和破坏，而运用各种技术手段实时收集和监控网络环境中每一个组成部分的系统状态、安全事件、网络活动，以便集中报警、及时处理及审计定责。  
从功能上，堡垒机综合了核心系统运维和安全审计管控两大主干功能。  
从技术实现上，通过切断终端计算机对网络和服务器资源的直接访问，而采用协议代理的方式，接管了终端计算机对网络和服务器的访问。  

简易的堡垒机功能简单，主要核心功能是远程登录服务器和日志审计。堡垒机就是一台服务器，维护人员在维护过程中，首先要统一登录到这台服务器上，然后从这台服务器再登录到目标设备进行维护。  

比较优秀的用于搭建堡垒机的开源软件 jumpserver，主要功能有：认证、授权、审计、自动化、资产管理等。  
商业堡垒机的功能比开源的要强大，比较出名的有：齐治，Citrix XenApp 等。  

### 为什么需要堡垒机
近年来数据安全事故频发，包括斯诺登事件、希拉里邮件丑闻以及携程宕机事件等，数据安全与防止泄露成为政府和企业都非常关心的议题，因此云堡垒机也应运而生。

### 搭建堡垒机
堡垒机需要具有公网 IP 以及内网 IP。  
其中，内网 IP 用于和机房其他机器通信；公网 IP 是用于在外部登录，通过公网 IP 登录到堡垒机后，才能访问内网的机器。  

搭建堡垒机，首先需要限制端口，留出可以远程登录的端口，其他的端口都封闭掉。  
然后，还需要配置白名单 IP，规定只有哪些 IP 可以登录，以及禁止密码登录，只允许密钥登录等，做这些事情的目的是为了增加堡垒机的安全性。  
除此之外，还需要限制登录的用户，限制为普通用户登录，和限制用户可以执行的命令等。同时，还需要在客户机器上做日志审计。  

**安装 jailkit 实现 chroot**
安装 jailkit 实现 chroot 的目的是为了限制登录的用户能够执行的命令（因为要防止登录的用户对堡垒机进行其他的操作）。  
jailkit 可以把用户限制在一个虚拟的系统中，这个虚拟系统的环境是 chroot 的，让用户无法直接操作真实系统。  
```bash
# 安装 jailkit
cd /usr/local/src/
# 下载 url：https://olivier.sessink.nl/jailkit/index.html#download
wget https://olivier.sessink.nl/jailkit/jailkit-2.20.tar.bz2
tar jxvf jailkit-2.20.tar.bz2
cd jailkit-2.20
./configure && make && make install

# 创建一个目录作为虚拟系统的根目录
mkdir /home/jail
# 给虚拟系统初始化一些命令，让这个系统具有基本的文件结构、网络相关的以及常用命令等
jk_init -v -j /home/jail/ basicshell
jk_init -v -j /home/jail/ editors
jk_init -v -j /home/jail/ netutils
jk_init -v -j /home/jail/ ssh
ls /home/jail

# 创建真实系统的用户
useradd jailUser
passwd jailUser

# 创建虚拟系统的 sbin 目录，并拷贝虚拟系统的 shell 文件
mkdir /home/jail/usr/sbin
cp /usr/sbin/jk_lsh /home/jail/usr/sbin/jk_lsh
# 创建虚拟系统的用户
jk_jailuser -m -j /home/jail jailUser
# 编辑虚拟系统用户的密码文件内容
vim /home/jail/etc/passwd
# root:x:0:0:root:/root:/bin/bash
# jailUser:x:1011:1011::/home/jailUser:/bin/bash  # 改成 /bin/bash 后才能被远程登录
```

这时，已经可以使用 jailUser 登陆系统了。  
能够成功登录后就可以生成密钥，添加密钥认证。配置完密钥后，还需要在真实系统上编辑 sshd_config 文件把 PasswordAuthentication 的值改为 no，这样就能禁止远程密码登录了。  
真实系统上，还需要开启防火墙规则，除了远程登录端口外其他端口都需要封闭掉。除此之外，还需要配置白名单 IP，这个在 hosts.allow 文件里配置，hosts.deny 里也需要进行相应的配置。  

**日志审计**  
在每个客户机上做一个简单的日志审计，记录在该机器执行过的命令。  
```bash
# 创建日志文件存放的目录
mkdir /usr/local/records 
chmod 777 !$
chmod +t !$
vim /etc/profile  
source /etc/profile

# /etc/profile 文件末尾增加以下内容
if [ ! -d  /usr/local/records/${LOGNAME} ]
then
    mkdir -p /usr/local/records/${LOGNAME}
    chmod 300 /usr/local/records/${LOGNAME}
fi
export HISTORY_FILE="/usr/local/records/${LOGNAME}/bash_history"
export PROMPT_COMMAND='{ date "+%Y-%m-%d %T ##### $(who am i |awk "{print \$1\" \"\$2\" \"\$5}") #### $(history 1 | { read x cmd; echo "$cmd"; })"; } >>$HISTORY_FILE'
```

配置后，就可以查看生成的日志了。
```bash
ls /usr/local/records/
ls /usr/local/records/root/
cat /usr/local/records/root/bash_history 
```
