
```bash
# 关机，如下各命令
halt, poweroff, shutdown, init 0
# 重启，如下各命令
reboot, shutdown, init 6

# 查看命令文档
man [COMMAND]
# 查看命令信息
info [COMMAND]
# 显示命令帮助信息
[COMMAND] --help

# 显示指定路径下的文件列表
ls [OPTION]... [DIR]...
# -a, --all: 显示所有文件，包括隐藏文件；
# -l: 长格式
# -h, --human-readable：单位换算
# -d: 显示目录自身的相关属性；通常要与-l一起使用；
# -r, --reverse: 逆序显示
# -R, --recursive: 递归
ls -l
# -rw-r--r-- 1 root root 44800 2月 26 14:32 LICENSE
# -rw-r--r--: 最左侧的第一位表示文件类型 -, d, l, b, c, p, s；后面的 9 位字符(rwx)表示访问权限
# 数字 1：文件被硬链接的次数；
# 左 root: 文件的 owner
# 右 root: 文件的 group
# 44800：文件的 size，单位 byte
# 2月 26 14:32 : 文件的最近一次被修改的时间
# LICENSE： 文件名

# 查看当前路径的绝对路径
pwd

# 命令的历史
history
# 显示命令历史中最近的 10 条命令
history 10
# 清空当前命令历史
history -c
# 与 history 相关的环境变量有 HISTSIZE、HISTFILE、HISTFILESIZE
# HISTSIZE：命令历史记录的条数
# HISTFILE：命令历史记录存储文件，默认 ~/.bash_history
# HISTFILESIZE：命令历史文件记录历史的条数

# 收集系统日志及架构信息并输出诊断文档
sosreport 

# 回当前用户的主目录
cd 或 cd ~
# 切换至指定用户的主目录
cd ~USERNAME
# 在上一个目录和当前目录之间来回切换
cd -

# 创建目录
mkdir [options] /path/to/somewhere
# -p: 存在于不报错，且可自动创建所需的各目录
# -v: 显示详细信息
# -m MODE: 创建目录时直接指定权限

# 创建多目录
mkdir /tmp/x/{y1,y2}/{a,b}
mkdir {x,y}_{m,n}

# 删除空目录
rmdir [OPTION]... DIRECTORY...
# -v: 显示过程

# 复制文件或目录
cp [OPTION] SRC DEST
# -i：交互式
# -r, -R: 递归复制目录及内部的所有内容；
# -a: 归档，相当于-dR --preserv=all
# -f: --force

# 远程复制文件或者文件夹
# 复制本地到远程： scp [-r] local_path user_name@ip:remote_path
# 复制远程到本地： scp [-r] user_name@ip:remote_path local_path

# 移动文件或重命名
mv [OPTION] SOURCE DEST
# -i：交互式
# --backup=<备份模式>：若需覆盖文件，则覆盖前先行备份
# -u：当源文件比目标文件新或者目标文件不存在时，才执行移动操作
# -f: --force

# 删除文件或目录
rm [OPTION] FILE
# -i: 交互式
# -f: 强制删除
# -r: 递归

# 为文件创建链接
ln aim_file_path link_file_path
# -s 参数表示对源文件建立符号链接，而非硬链接
# -f 参数表示强行删除任何已存在的目标文件

# 硬链接与软链接
# 硬链接会增加相关节点，建立目标文件的目录项
# 硬链接不能对目录、不同系统间进行连接
# 硬链接相当于目标文件的别名，意义在于防止真实文件被误操作，删除一个硬链接文件并不影响其他有相同 inode 号的文件
# 软链接也称符号链接，生成一个新文件，没有目录和系统间的限制
# 软链接相当于一个指针，删除软链接并不影响被指向的文件，删除源文件或目录只删除了数据但不会删除链接（连接变为死链接）；一旦以同样文件名创建了源文件，链接将继续指向该文件的新数据

# 打包与解包
tar [OPTION] FILE
# -c或--create：创建新的备份文件（打包）
# -x或--extract或--get：从备份文件中还原文件（解包）
# -A 将 tar 文件添加到归档文件中
# -z或--gzip或--ungzip：通过gzip指令处理备份文件
# -Z或--compress或--uncompress：通过 compress 指令处理备份文件
# -f<文件>或--file=<文件>：指定文件或目录
# -v或--verbose：显示指令执行过程

# 不同的后缀的文件
# *.tar、*.tar.Z、*.tar.bz2、*.tar.gz、*.tgz 用 tar 加、解压
# *.gz 用 gzip、gunzip 加、解压
# *.bz2 用 bzip2、bunzip2 加、解压
# *.Z 用 compress、uncompress 加、解压
# *.rar 用 rar、unrar 加、解压
# *.zip 用 zip、unzip 加、解压

# 别名的设置与删除
alias ll='ls -alhF'
unalias ll

# 显示当前在线用户
w
# 显示已经登录的用户
who
# 我是谁
whoami
# 更详细 who x x，又如 who is she\who are you
who am i
# 查看系统所有用户
cut -d: -f1 /etc/passwd
# 查看系统所有用户组
cut -d: -f1 /etc/group

# 查看用户登录日志
last

# 显示主机名称
hostname

# 显示执行命令所在的路径
which command
# 显示某个命令的所有位置
whereis command
# 查找文件位置 whereis 文件名 或 find / -name 文件名
# 查找文件夹位置 locate 文件夹名

# 显示系统内核信息
uname -a

# 显示磁盘信息
df -ah

# 显示本次开机运行的时间
uptime

# 显示某个文件或目录的磁盘使用量
du file_name
# -h 参数将返回的大小显示为人类可读的格式，即显示单位为 K、M、G 等
# -s 参数表示总结
# -x 参数表示不显示不在当前分区的目录，通常会忽略/dev、/proc、/sys 等目录
# -c 参数表示显示当前目录总共占用的空间大小
# -exclude 参数用于排除某些目录或文件，如 --exclude=*.iso
# --max-depth 参数用于设定目录大小统计到第几层，如果 -–max-depth=0，那么等同于 -s 参数

# 读取文件所有内容
cat [OPTION] [FILE]
# -n 或 --number：有1开始对所有输出的行数编号
# -A：显示不可打印字符，行尾显示 $

# 查看文件内容，并显示行号
nl [OPTION] [FILE]
# -s 参数可设定行号的后缀

# 少即是多，分屏上下翻页浏览文件内容，PageUp 键向上翻页，PageDown 键向下翻页,Q 键退出 
less [OPTION] [FILE]
# -N：每一行行首显示行号
# -s：将连续多个空行压缩成一行显示
# -S 参数表示在单行显示较长的内容，而不换行显示
# -l：搜索时忽略大小写的差异
# -g：不加亮显示搜索到的所有关键词，仅显示当前显示的关键字，以提高显示速度
# -f：强制显示文件

# 显示文件内容，每次显示一屏，Enter 键向下翻滚一行，空格向下滚动一屏，B 键显示上一屏内容，Q 键退出
more [OPTION] [FILE]
# -d: 显示翻页及退出提示
# -s：将连续多个空行压缩成一行显示

# 读取文件的头部
head [OPTION] [FILE]
# -c ?: 指定获取前 ? 字节
# -n ? 或 -?: 指定获取前 ? 行

# 读取文件的尾部
tail [OPTION] [FILE]
# -c ?: 指定获取后 ? 字节
# -n ? 或 -?: 指定获取后 ? 行
# -f: 跟踪显示文件新追加的内容

# 给其他命令传参，默认给命令 echo 传参
xargs
# 多行输入单行输出: cat test.txt | xargs
# 每 3 项一行输出多行：cat test.txt | xargs -n3
# 自定义定界符：echo "nameXnameXnameXname" | xargs -dX
# 将\0作为定界符：xargs -0
# 使用 -I 指定一个替换字符串 {}：cat arg.txt | xargs -I {} ./sk.sh -p {} -l

# 连接文件并打印到标准输出设备上
cut [OPTION] [FILE]
# -c ?：仅显示行中指定范围的字符，例如 -c-2 \ -c5- 
# -d DELIMITER: 指明分隔符
# -f FILEDS: 显示指定字段的内容
#    -f?: 第 ? 个字段
#    -f?,?[,?]：离散的多个字段，例如 -f1,3,6
#    -f?-?：连续的多个字段, 例如 -f1-6
#    -f1-3,7: 混合使用
# --output-delimiter=STRING

# 对比文件差异
diff [OPTION] [FILE...]
# --brief 显示比较后的结果
# -c      描述文件内容具体的不同

# 报告或忽略文件中的重复行 (连续且完全相同)
uniq [OPTION] [FILE]
# -c: 显示每行重复出现的次数
# -d: 仅显示重复过的行
# -u: 仅显示不曾重复的行

# 将文件进行排序并输出
sort [OPTION] [FILE]
# -f: 忽略字符大小写
# -r: 逆序
# -t DELIMITER: 字段分隔符
# -k ?：以指定字段为标准排序
# -n: 以数值大小进行排序
# -u: uniq，排序后去重

# 重定向输出到文件并显示
echo 'aaa' | tee /tmp/aaa.txt
# -a 向文件中重定向时使用追加模式
# -i 忽略中断（interrupt）信号

# 转换或删除字符
tr [OPTION] SET1 [SET2]
# -c 或 ——complerment：取代所有不属于第一字符集的字符
# -d 或 ——delete：删除所有属于第一字符集的字符
# -s 或 --squeeze-repeats：把连续重复的字符以单独一个字符表示
# -t 或 --truncate-set1：先删除第一字符集较第二字符集多出的字符

# 按照给定模式转换文本，如小写转大写
cat example.txt | tr 'a-z' 'A-Z'

# 创建新的空文件
touch [OPTION] [FILE]
# -a: 只更改存取时间
# -m: 只更改变动时间
# -t <日期时间>:  使用指定的日期时间，而非现在的时间
# -c: 如果文件不存在，则不被创建

# 查看文件类型
file file_path

# 查看文件的详细信息
stat file_path

# 修改文件权限
chmod [OPTION] [FILE]
# -R: 递归

# 修改文件的属主和属组（: 可替换为 .）
chown [OPTION] [OWNER][:[GROUP]] [FILE]
# -R: 递归

# 修改文件的属组
chgrp [OPTION] [GROUP] [FILE]
# -R: 递归

# 设置文件不允许修改、删除、移动、复制，root 用户也生效
chattr +i file_path
# 取消文件属性设置
chattr -i file_path
# 查看文件或目录属性设置
lsattr file_path

# 文件或目录创建时的遮罩码（掩码）
umask [OPTION] [FILE]
# -p：输出的权限掩码可直接作为指令来执行
# -S：以符号方式输出权限掩码

# 显示文本中的行数、字数、字节数等信息
wc -l file_path
# -l 参数表示 lines，显示行数
# -w 参数表示 words，显示字数
# -c 参数表示 characters，显示字节数

# 列出当前正在执行的进程信息
ps [OPTION]
# -u 参数列出指定用户拥有的进程
# a 参数表示显示现行终端机下的所有程序，包括其他用户的程序
# u 以用户为主的格式来显示程序状况
# x 显示所有程序，不以终端机来区分
# e 列出程序时，显示每个程序所使用的环境变量
# f 用 ASCII 字符显示树状结构，表达程序间的相互关系

# 查看进程 pid
pidof [server_name]

# 查看进程打开的文件、目录和套接字等信息（root 用户使用）
lsof [OPTION]
# -a <文件>: 列出打开文件存在的进程
# -c <进程名>：列出指定进程所打开的文件
# -p <进程号>：列出指定进程号所打开的文件
# -u 列出用户 UID 号进程详情
# -i<条件>: 列出符合条件的进程（4、6、协议、:端口、 @ip ）
# -d<文件号>：列出占用该文件号的进程
# +d<目录>：列出目录下被打开的文件；
# +D<目录>：递归列出目录下被打开的文件

# 查看所有网络接口的属性
ifconfig 
# 查看防火墙设置
iptables -L
# 查看路由表
route -n 

# 查看端口（明日黄花）
netstat [OPTION]
# -a 显示所有，默认不显示 LISTEN
# -n 不显示数字别名
# -p 显示关联的程序f

# 查询 socket 的有关统计信息
ss [OPTION]
# -s, --summary   显示套接字（socket）使用概况（网络连接统计）
# -l, --listening 显示监听状态的套接字（sockets）（所有打开的网络端口）
# -a, --all 所有 socket 连接

# 所有端口为 22（ssh）的连接
ss state all sport = :ssh

# 切换用户或以其他用户身份执行
# 非登录式切换，即不会读取目标用户的配置文件
su user_name
# 登录式切换，会读取目标用户的配置文件，完全切换
su - user_name
# 注意，root 用户 su 到其他用户无需密码，非 root 用户切换时需要密码

# 换个身份执行命令
su [- 或 -l] user_name -c 'COMMAND'

# 以其他身份来执行命令
sudo [OPTION] COMMAND  
# -u<用户> 以指定的用户作为新的身份。若不加上此参数，则预设以 root 作为新的身份

# 使用 root 超级用户重新登录一次 shell，只不过密码是使用的当前用户的密码
# 重新加载 /etc/profile 文件以及 /etc/bashrc 文件等系统配置文件，并且还会重新加载 root 用户的 $SHELL 环境变量所对应的配置文件
sudo su -

# 修改自己的登录密码
passwd
# 修改其他用户的登录密码（仅限 root）
passwd [OPTIONS] user_name
# -l: 锁定指定用户
# -u: 解锁指定用户
# -n mindays: 指定最短使用期限
# -x maxdays：最大使用期限
# -w warndays：提前多少天开始警告
# -i inactivedays：非活动期限；

# 添加用户
useradd [OPTIONS] user_name
# -u 指定UID（省略该选项, 系统会自动分配一个UID）
# -g 初始化群组, （默认创建新的群组, 群组名与帐号名称相同）
# -G 次要的群组, 可以指定该帐号的所属的其它群组,多个用,分开, 如 -G g1,g2,g3
# -c 用户的注释信息
# -d 该用户的home目录
# -s shell脚本环境, 默认为 /bin/bash

# 修改用户（大部分参数与 useradd 相同）
usermod [OPTIONS] user_name
# -l 修改帐号名称
# -U 解除密码锁定

# 删除用户
userdel [OPTIONS] user_name
# -f 强制删除用户，即使用户当前已登录
# -r 删除用户的同时，删除与用户相关的所有文件

# 添加用户组
groupadd [OPTIONS] group_name
# -g gid 指定组群 id
# -o 允许添加组 ID 号不唯一的工作组
# -r 创建系统工作组，系统工作组的组 ID 小于 500
# -K 覆盖配置文件 /ect/login.defs

# 修改用户组
groupmod [OPTIONS] group_name
# -g gid 指定组群 id
# -n group_name 修改用户组名

# 删除用户组
groupdel group_name

# 从标准输入接收用户密码
echo "PASSWORD" | passwd --stdin user_name

# 清屏 ctrl+l
clear

# 显示本月的日历
cal
# 显示当前的日期和时间
date +"%Y-%m-%d %H:%M:%S"

# 显示或管理执行中的程序 [htop](http://hisham.hm/htop/)
htop [OPTION]
# -u --user=USER_NAME 只显示一个指定用户的进程
# -d --delay=DELAY 设置更新之间的延迟，在十秒

# 同步的艺术
rsync [OPTION] local_file user@ip:/remote_file
# -t 修改时间也同步
# -I 踏实做人，逐个文件发起数据同步
# -l 保持软链接文件类型
# -P 保持权限
# -v 查看同步信息
# -r 递归同步文件夹
# -a 霸道，尽可能的保证各方面的一致性
# --delete 如果源端没有该文件，删除它
# --exclude 排除隐私文件
# --partial 断点续传
# --progress 显示传输进度

# 杀死进程
kill [-9] process_id
killall process_name

# 建立信任链接
ssh-copy-id "-p port user_name@ip"
# 提示 “No identities found”，需要给使用的登录账号创建公钥私钥（ssh-keygen -t rsa）
# 输入密码成功后，后面就可以免密登录了（ssh user_name@ip -p port ）

# 查看已存在的环境变量
env [OPTION]
# -i 开始一个新的空的环境
# -u<变量名> 从当前环境中删除指定的变量

# 任务调度
crontab [-u username] [-l|-e|-r]
# * * * * * COMMAND 分(0 - 59) 时(0 - 23) 日(1 - 31) 月(1 - 12) 周(0 - 7) (Sunday=0 or 7)
# -u 只有 root 才能进行这个任务，也就是说帮其他使用者创建/移除 crontab 工作排程
# -e 编辑 crontab 的工作内容
# -l 查阅 crontab 的工作内容
# -r 移除所有的 crontab 的工作内容，若仅要移除一项，请用 -e 去编辑

# 下载文件 
wget [OPTIONS] URL
# -O xxx 指定一个文件名
# -o 设置下载的日志文件
# -r 下载远程文件夹
# --limit-rate=300k 限速 300k
# -c 断点续传
# -b 后台下载
# -i, –-input-file=FILE 下载在 FILE 文件中出现的 URLs
# --ftp-user=USERNAME --ftp-password=PASSWORD 用户名和密码认证的 ftp 下载
# -r -A.pdf 下载指定格式（如全站 pdf 文件）
# --spider 测试下载链接
# --execute robots=off 避开网站 robots.txt 封禁
# --tries=40 失败可重试 40 次
# --user-agent="Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/534.16 (KHTML, like Gecko) Chrome/10.0.648.204 Safari/534.16" 伪代理下载

# 利用 URL 规则传输文件
curl [OPTION] URL
# --silent 不显示进度信息
# -C 偏移量 断点续传（-C - 自动判断续传位置）
# --referer "SITE" 指定参照页字符串
# --user-agent 或 -A 设置用户代理
# -H 设置头部信息
# --limit-rate 50k 限制下载速度 50k
# --max-filesize 指定可下载的最大文件大小
# -u user:pwd 完成 HTTP 或 FTP 认证
# -I 或 -head 只打印出HTTP头部信息
# -v 显示请求全过程解析
# -d 添加请求参数
# -X 以什么方式请求，如 -XPOST
# -F "key=value" 模拟表单提交数据

# 挂载磁盘设备
mount [OPTION] origin_dir aim_dir
# –t  文件系统类型
# -o ro 用唯读模式挂上

# 卸载已经挂载的文件系统
umount [OPTION] aim_dir
# -v 执行时显示详细的信息
# -a 卸除 /etc/mtab 中记录的所有文件系统
# -r 若无法成功卸除，则尝试以只读的方式重新挂入文件系统
# -n 卸除时不要将信息存入 /etc/mtab 文件中

# 并发压力测试
ab -n 100 -c 10 -l http://www.your_site.com
# -n number 总的请求数
# -c concurrency 并发数 
# -l 表示当某个请求的回复长度不与第一个请求的回复长度一致时，不把它作为失败的请求
ab -n 100 -c 10 -p post.txt http://www.your_site.com
# post.txt 的内容是参数，如 data={"name":"wgc"}，但是，需要 url 编码
```