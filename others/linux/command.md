
```bash
# 关机，如下各命令
halt, poweroff, shutdown, init 0
# 重启，如下各命令
reboot, shutdown, init 6

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

# 回当前用户的主目录
cd 或 cd ~
# 切换至指定用户的主目录
cd ~USERNAME
# 在上一个目录和当前目录之间来回切换
cd -

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

# 显示执行命令所在的路径
which command
# 显示某个命令的所有位置
whereis command

# 显示内核信息
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

# 查看文件内容 
cat file_path
# 查看文件内容，并显示行号
nl file_path
# -s 参数可设定行号的后缀

# 显示文件内容，每次显示一屏，Enter 键向下翻滚一行，空格向下滚动一屏，B 键显示上一屏内容，Q 键退出
more file_path
# 分屏上下翻页浏览文件内容，PageUp 键向上翻页，PageDown 键向下翻页,Q 键退出
less file_path
# -l 参数表示搜索时忽略大小写的差异
# -N 参数表示每一行行首显示行号
# -s 参数表示将连续多个空行压缩成一行显示
# -S 参数表示在单行显示较长的内容，而不换行显示

# 查看文件类型
file file_path

# 按照给定模式转换文本，如小写转大写
cat example.txt | tr 'a-z' 'A-Z'

# 显示文本中的行数、字数、字节数等信息
wc -l file_path
# -l 参数表示 lines，显示行数
# -w 参数表示 words，显示字数
# -c 参数表示 characters，显示字节数

# 列出当前正在执行的进程信息
ps
# -u 参数列出指定用户拥有的进程
# a 参数表示显示现行终端机下的所有程序，包括其他用户的程序
# u 以用户为主的格式来显示程序状况
# x 显示所有程序，不以终端机来区分
# e 列出程序时，显示每个程序所使用的环境变量
# f 用 ASCII 字符显示树状结构，表达程序间的相互关系

# 清屏 ctrl+l
clear

# 显示本月的日历
cal
# 显示当前的日期和时间
date +"%Y-%m-%d %H:%M:%S"

# 杀死进程
kill [-9] process_id
killall process_name
```