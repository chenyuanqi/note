
### awk 的基本用法
awk 是一个强大的文本分析工具。  
awk 就是把文件逐行的读入，以空格为默认分隔符将每行切片，切开的部分再进行各种分析处理。  
awk 有 3 个不同版本: awk、nawk 和 gawk，未作特别说明，一般指 gawk，gawk 是 AWK 的 GNU 版本。   
awk 其名称得自于它的创始人 Alfred Aho 、Peter Weinberger 和 Brian Kernighan 姓氏的首个字母。实际上 AWK 的确拥有自己的语言： AWK 程序设计语言 ， 三位创建者已将它正式定义为 “样式扫描和处理语言”。它允许您创建简短的程序，这些程序读取输入文件、为数据排序、处理数据、对输入执行计算以及生成报表，还有无数其他的功能。awk 其名称得自于它的创始人 Alfred Aho 、Peter Weinberger 和 Brian Kernighan 姓氏的首个字母。实际上 AWK 的确拥有自己的语言： AWK 程序设计语言 ， 三位创建者已将它正式定义为 “样式扫描和处理语言”。它允许您创建简短的程序，这些程序读取输入文件、为数据排序、处理数据、对输入执行计算以及生成报表，还有无数其他的功能。  

使用方法如下
```bash
# 格式
$ awk '条件 动作' 文件名

# 示例，$0 代表当前行
$ awk '{print $0}' demo.txt 
$ echo 'this is a test' | awk '{print $3}' # 输出 a
```
awk 会根据空格和制表符，将每一行分成若干字段，依次用 $1、$2、$3 代表第一个字段、第二个字段、第三个字段等等。

比如，我们把 /etc/passwd 文件保存成 demo.txt
```bash
$ cat /etc/passwd > demo.txt | cat
root:x:0:0:root:/root:/usr/bin/zsh
daemon:x:1:1:daemon:/usr/sbin:/usr/sbin/nologin
bin:x:2:2:bin:/bin:/usr/sbin/nologin
sys:x:3:3:sys:/dev:/usr/sbin/nologin
sync:x:4:65534:sync:/bin:/bin/sync

$ awk -F ':' '{ print $1 }' demo.txt
root
daemon
bin
sys
sync
```

awk 还提供其他一些变量：  
NF：表示当前行有多少个字段。  
NR：表示当前处理的是第几行。  
FILENAME：当前文件名。  
FS：字段分隔符，默认是空格和制表符。  
RS：行分隔符，用于分割每一行，默认是换行符。  
OFS：输出字段的分隔符，用于打印时分隔字段，默认为空格。  
ORS：输出记录的分隔符，用于打印时分隔记录，默认为换行符。  
OFMT：数字输出的格式，默认为 ％.6g。  
```bash
$ echo 'this is a test' | awk '{print $NF}'
test

# print 命令里面的逗号，表示输出的时候，两个部分之间使用空格分隔
$ awk -F ':' '{print $1, $(NF-1)}' demo.txt
root /root
daemon /usr/sbin
bin /bin
sys /dev
sync /bin
```

awk 还提供了一些内置函数，方便对原始数据的处理：  
toupper()：字符转为大写。  
tolower()：字符转为小写。  
length()：返回字符串长度。  
substr()：返回子字符串。  
sin()：正弦。  
cos()：余弦。  
sqrt()：平方根。  
rand()：随机数。  
```bash
$ awk -F ':' '{ print toupper($1) }' demo.txt
ROOT
DAEMON
BIN
SYS
SYNC
```

awk 允许指定输出条件，只输出符合条件的行。  
```bash
# 使用正则
$ awk -F ':' '/usr/ {print $1}' demo.txt
root
daemon
bin
sys

# 输出奇数行
$ awk -F ':' 'NR % 2 == 1 {print $1}' demo.txt
root
bin
sync
```

awk 提供了 if 结构，用于编写复杂的条件。
```bash
# 输出第一个字段的第一个字符大于 m 的行
$ awk -F ':' '{if ($1 > "m") print $1}' demo.txt
root
sys
sync
```

### awk 进阶用法
awk 遵循了非常简单的工作流：读取，执行和重复。  
读取 -- （从输入流（文件，管道或者标准输入）中读取一行，然后存储到内存中）。  
执行 -- 所有的AWK命令都依次在输入上执行。默认情况下，awk 会对每一行执行命令，我们可以通过提供模式限制这种行为。  
重复 -- 处理过程不断重复，直到到达文件结尾。  

```bash
# 统计某个文件夹下的文件占用的字节数
$ ls -l |awk 'BEGIN {size=0;} {size=size+$5;} END{print "[end]size is ", size}'
[end]size is 8657198
```

awk 中的循环语句同样借鉴于 C 语言，支持 while、do/while、for、break、continue，这些关键字的语义和C语言中的语义完全相同。  
