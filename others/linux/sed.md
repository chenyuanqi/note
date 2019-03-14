
### sed 的基本用法
sed 是一种在线编辑器，它一次处理一行内容。  
处理时，把当前处理的行存储在临时缓冲区中，称为 “模式空间”（pattern space），接着用 sed 命令处理缓冲区中的内容，处理完成后，把缓冲区的内容送往屏幕。接着处理下一行，这样不断重复，直到文件末尾。文件内容并没有 改变，除非你使用重定向存储输出。  
sed 主要用来自动编辑一个或多个文件；简化对文件的反复操作；编写转换程序等。  
```bash
$ sed [-nefr] [动作]
# 选项与参数：
# -n ：使用安静(silent)模式。在一般 sed 的用法中，所有来自 STDIN 的数据一般都会被列出到终端上。但如果加上 -n 参数后，则只有经过sed 特殊处理的那一行(或者动作)才会被列出来。
# -e ：直接在命令列模式上进行 sed 的动作编辑；
# -f ：直接将 sed 的动作写在一个文件内， -f filename 则可以运行 filename 内的 sed 动作；
# -r ：sed 的动作支持的是延伸型正规表示法的语法。(默认是基础正规表示法语法)
# -i ：直接修改读取的文件内容，而不是输出到终端。

# 动作说明： [n1[,n2]]function
# n1, n2 ：不见得会存在，一般代表『选择进行动作的行数』，举例来说，如果我的动作是需要在 10 到 20 行之间进行的，则『 10,20[动作行为] 』

# function：
# a ：新增， a 的后面可以接字串，而这些字串会在新的一行出现(目前的下一行)～
# c ：取代， c 的后面可以接字串，这些字串可以取代 n1,n2 之间的行！
# d ：删除，因为是删除啊，所以 d 后面通常不接任何咚咚；
# i ：插入， i 的后面可以接字串，而这些字串会在新的一行出现(目前的上一行)；
# p ：列印，亦即将某个选择的数据印出。通常 p 会与参数 sed -n 一起运行～
# s ：取代，可以直接进行取代的工作哩！通常这个 s 的动作可以搭配正规表示法！例如 1,20s/old/new/g 就是啦！

# 将 /etc/passwd 的内容列出并且列印行号，同时，请将第 2~5 行删除
$ nl /etc/passwd | sed '2,5d'
1 root:x:0:0:root:/root:/bin/bash
6 sync:x:5:0:sync:/sbin:/bin/sync
7 shutdown:x:6:0:shutdown:/sbin:/sbin/shutdown
# 删除第 3 到最后一行
$ nl /etc/passwd | sed '3,$d' 
# 在第 2 行后 (亦即是加在第三行) 加上『drink tea』字样
$ nl /etc/passwd | sed '2a drink tea'
# 在第 2 行前 (亦即是加在第三行) 加上『drink tea』字样
$ nl /etc/passwd | sed '2i drink tea' 
# 如果是加上多行，每一行之间都必须要以反斜杠『 \ 』来进行新行的添加
$ nl /etc/passwd | sed '2a Drink tea or ......\
> drink beer ?'

# 将第 2-5 行的内容取代成为『No 2-5 number』
$ nl /etc/passwd | sed '2,5c No 2-5 number'

# 仅列出 /etc/passwd 文件内的第 5-7 行
$ nl /etc/passwd | sed -n '5,7p'

# 搜索 /etc/passwd 有 root 关键字的行（如果 root 找到，除了输出所有行，还会输出匹配行）
$ nl /etc/passwd | sed '/root/p'
1	root:x:0:0:root:/root:/bin/bash
1	root:x:0:0:root:/root:/bin/bash
2	bin:x:1:1:bin:/bin:/sbin/nologin
3	daemon:x:2:2:daemon:/sbin:/sbin/nologin
# 使用 -n 的时候，将只打印包含模板的行
$ nl /etc/passwd | sed -n '/root/p'
1  root:x:0:0:root:/root:/bin/bash
# 删除 /etc/passwd 所有包含 root 的行，其他行输出
$ nl /etc/passwd | sed  '/root/d'
2  daemon:x:1:1:daemon:/usr/sbin:/bin/sh
3  bin:x:2:2:bin:/bin:/bin/sh

# 搜索 /etc/passwd, 找到 root 对应的行，执行后面花括号中的一组命令，每个命令之间用分号分隔，这里把 bash 替换为 blueshell，再输出这行
$ nl /etc/passwd | sed -n '/root/{s/bash/blueshell/;p}'
1  root:x:0:0:root:/root:/bin/blueshell
# 如果只替换 /etc/passwd 的第一个 bash 关键字为 blueshell，就退出（q 是退出）
$ nl /etc/passwd | sed -n '/bash/{s/bash/blueshell/;p;q}'    
1  root:x:0:0:root:/root:/bin/blueshell
```

除了整行的处理模式之外， sed 还可以用行为单位进行部分数据的搜寻并取代。基本上 sed 的搜寻与替代的与 vi 相当的类似。
```bash
# 格式
$ sed 's/要被取代的字串/新的字串/g'

# 利用 /sbin/ifconfig 查询 IP，将 IP 前面的部分予以删除
$ /sbin/ifconfig eth0 | grep 'inet addr' | sed 's/^.*addr://g'
192.168.1.100 Bcast:192.168.1.255 Mask:255.255.255.0
```

sed 实现多点编辑。
```bash
# 一条 sed 命令，删除 /etc/passwd 第三行到末尾的数据，并把 bash 替换为 blueshell
# -e 表示多点编辑，第一个编辑命令删除 /etc/passwd 第三行到末尾的数据，第二条命令搜索 bash 替换为 blueshell
$ nl /etc/passwd | sed -e '3,$d' -e 's/bash/blueshell/'
1  root:x:0:0:root:/root:/bin/blueshell
2  daemon:x:1:1:daemon:/usr/sbin:/bin/sh
```

sed 可以直接修改文件的内容，不必使用管道命令或数据流重导向（慎操作！）。  
sed 的『 -i 』选项可以直接修改文件内容，这功能非常有帮助！举例来说，如果你有一个 100 万行的文件，你要在第 100 行加某些文字，此时使用 vim 可能会疯掉！  
```bash
# 利用 sed 将 regular_express.txt 内每一行结尾若为.则换成！
$ sed -i 's/\.$/\!/g' regular_express.txt

# 利用 sed 直接在 regular_express.txt 最后一行加入『# This is a test』
#（$ 代表的是最后一行，而 a 的动作是新增）
$ sed -i '$a # This is a test' regular_express.txt
```