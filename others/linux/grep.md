
### grep 基本用法
grep (global search regular expression (RE) and print out the line, 全面搜索正则表达式并把行打印出来) 是一种强大的文本搜索工具，它能使用正则表达式搜索文本，并把匹配的行打印出来。  

Unix 的 grep 家族包括 grep、egrep 和 fgrep。egrep 和 fgrep 的命令只跟 grep 有很小不同。egrep 是 grep 的扩展，支持更多的 re 元字符， fgrep 就是 fixed grep 或 fast grep，它们把所有的字母都看作单词，也就是说，正则表达式中的元字符表示回其自身的字面意义，不再特殊。linux 使用 GNU 版本的 grep。它功能更强，可以通过 - G、-E、-F 命令行选项来使用 egrep 和 fgrep 的功能。  
```bash
# grep 用于在文本中执行关键词搜索，并显示匹配的结果
$ grep [OPTION] file_path
# -a 将 binary 文件以 text 文件的方式搜寻数据
# -b 将可执行文件当作文本文件搜索
# -c 仅显示找到的行数
# -i 忽略大小写
# -n 显示行号
# -v 反向选择 —— 仅列出没有“关键词”的行
# --color=auto 可以将找到的关键词部分加上颜色的显示

# 譬如，查找当前系统种不允许登录的用户信息
$ grep /sbin/nologin /etc/passwd
# 将 /etc/passwd，将没有出现 root 的行取出来
$ grep -v root /etc/passwd
# 将 /etc/passwd，将没有出现 root 和 nologin 的行取出来
$ grep -v root /etc/passwd | grep -v nologin

# 用 dmesg 列出核心信息，再以 grep 找出内含 eth 那行，要将捉到的关键字显色，且加上行号来表示
$ dmesg | grep -n --color=auto 'eth'
# 用 dmesg 列出核心信息，再以 grep 找出内含 eth 那行，在关键字所在行的前两行与后三行也一起捉出来显示
$ dmesg | grep -n -A3 -B2 --color=auto 'eth'

# 在当前目录搜索带'energywise'行的文件
$ grep 'energywise' *
# 在当前目录及其子目录下搜索'energywise'行的文件
$ grep -r 'energywise' *
# 在当前目录及其子目录下搜索'energywise'行的文件，但是不显示匹配的行，只显示匹配的文件
$ grep -l -r 'energywise' *
```

### grep 与正规表达式
字符类的搜索：如果我想要搜寻 test 或 taste 这两个单字时，可以发现到，其实她们有共通的 't?st' 存在～这个时候，我可以这样来搜寻
```bash
# [] 里面不论有几个字节，他都谨代表某『一个』字节
$ grep -n 't[ae]st' regular_express.txt
8:I can't finish the test.
9:Oh! The soup taste good.
```
字符类的反向选择 [^] ：如果想要搜索到有 oo 的行，但不想要 oo 前面有 g
```bash
$ grep -n '[^g]oo' regular_express.txt
2:apple is my favorite food.
3:Football game is not use feet only.
18:google is the best tools for search keyword.
19:goooooogle yes!

# 假设我 oo 前面不想要有小写字节，所以，我可以这样写
# 例如大写英文 / 小写英文 / 数字等等， 就可以使用 [a-z],[A-Z],[0-9] 等方式来书写
$ grep -n '[^a-z]oo' regular_express.txt
```
- 注意，行首与行尾字节符号：^ 和 $。另外，因为小数点具有其他意义，所以必须要使用转义字符 (\.) 来加以解除其特殊意义！  
> . (小数点)：代表『一定有一个任意字节』的意思；  
> * (星号)：代表『重复前一个字符， 0 到无穷多次』的意思，为组合形态；  
> {} 限定连续 RE 字符范围  

```bash
$ grep -n '\.$' regular_express.txt

# 找出空白行
$ grep -n '^$' regular_express.txt

# 找出『任意数字』的行
$ grep -n '[0-9][0-9]*' regular_express.txt

# 找出 g 后面接 2 到 5 个 o ，然后再接一个 g 的字串
$ grep -n 'go\{2,5\}g' regular_express.txt
#  2 个 o 以上的 goooo....g 
$ grep -n 'go\{2,\}g' regular_express.txt
```

### 扩展 grep (grep -E 或者 egrep)
使用扩展 grep 的主要好处是增加了额外的正则表达式元字符集。

比如，打印所有包含 NW 或 EA 的行。如果不是使用 egrep，而是 grep，将不会有结果查出。
```bash
$ egrep 'NW|EA' testfile
# 对于标准 grep，如果在扩展元字符前面加 \，grep 会自动启用扩展选项 - E
$ grep 'NW\|EA' testfile

# 搜索所有包含一个或多个 3 的行
$ egrep '3+' testfile  
$ grep -E '3+' testfile  
$ grep '3\+' testfile  
```

不使用正则表达式。  
fgrep 查询速度比 grep 命令快，但是不够灵活：它只能找固定的文本，而不是规则表达式。  
```bash
# 在一个文件或者输出中找到包含星号字符的行
$ fgrep  '*' /etc/profile for i in /etc/profile.d/*.sh ; do
# 或
$ grep -F '*' /etc/profile for i in /etc/profile.d/*.sh ; do
```
