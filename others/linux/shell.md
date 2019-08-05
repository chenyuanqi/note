
### shell link
[shell check](https://github.com/koalaman/shellcheck)  
[shell awesome](https://github.com/alebcay/awesome-shell)  

### shell 脚本
shell 脚本以.sh 结尾的或者不需要后缀都可以。
```sh
#!/bin/bash
# 文件名 hello.sh
NUM=100
# 输出
echo "学习 shell $NUM 分"  
# 格式化输出
printf "%d" $NUM
```
运行 shell 脚本需要注意分配执行权限，注意头部的 bom 头和首行行尾多余的看不见的空格  
```bash
sh hello.sh
```

### shell 注释与非注释
```sh
# 这是一个注释

echo "这里的 # 不是注释"
echo ${PATH#*:} # PATH 后的 # 是参数替换，不是注释
echo $(( 2#101011 ))  # 2 后面的 # 是基本转换，也不是注释
```

### shell 单引号与双引号
引号一个很重要的作用是保护命令行上的一个参数不被 shell 解释，而把此参数传递给要执行的程序来处理它；引号 echo 输出换行。  
单引号：任何字符都会原样输出，变量无效，不能再次出现单引号即使转义也不行  
双引号：可以有变量，可以出现转义字符

### shell 分隔符
分割符允许在同一行里有两个或更多的命令
```sh
echo here; echo there
filePath='/tmp/shell/hello.sh'
echo ${filePath} # 加花括号是为了帮助解释器识别变量的边界
if [ -x "$filePath" ]; then # 注意："if" and "then"需要分隔符，否则 then 需要换行
  echo "$filePath exists."
else
  echo "$filePath not found."
fi
```

### shell 操作符
**计算操作符**  

| 操作符 | 描述 |  
| ---- | ---- |  
| `+` | 加 |  
| `-` | 减 |  
| `*` | 乘 |  
| `/` | 除 |  
| `%` | 取模 |  
| `**` | 求幂 |  
| `+=` | 加等(plus-equal) 把原变量值增加一个常量并重新赋值给变量 |  
| `-=` | 减等(minus-equal) 把原变量值减少一个常量并重新赋值给变量 |  
| `*=` | 乘等(times-equal) 把原变量值乘上一个常量并重新赋值给变量 |  
| `/=` | 除等(slash-equal) 把原变量值除以一个常量并重新赋值给变量 |  
| `%=` | 模等(mod-equal) 把原变量值除以一个常量整除（即取模）并重新赋余数的值给变量 |  

**位操作符**

| 操作符 | 描述 | 操作符 | 描述 |  
| ---- | ---- | ---- | ---- |  
| `<<` | 位左移（每移一位相当乘以2） | `|` | 位或 |  
| `<<=` | 位左移赋值 | `|=` | 位或赋值 |  
| `>>` | 位右移（每移一位相当除以2） | `~` | 位反 |  
| `>>=` | "位右移赋值"（和<<=相反） | `!` | 位非 |  
| `&` | 位与 | `^` | 位或 |  
| `&=` | 位于赋值 | `^=` | 位或赋值 |  

**文件操作符**  

| 操作符 | 描述 |   
| ---- | ---- |  
| -e | 文件存在 |  
| -a | 文件存在，这个和-e的作用一样. 它是不赞成使用的，所以它的用处不大。 |  
| -f | 文件是一个普通文件(不是一个目录或是一个设备文件) |  
| -s | 文件大小不为零 |  
| -d | 文件是一个目录 |  
| -b | 文件是一个块设备(软盘，光驱，等等。) |  
| -c | 文件是一个字符设备(键盘，调制解调器，声卡，等等。) |  
| -p | 文件是一个管道 |  
| -h | 文件是一个符号链接 |  
| -L | 文件是一个符号链接 |  
| -S | 文件是一个socket |  
| -t | 文件(描述符)与一个终端设备相关。|  
| -r | 文件是否可读 (指运行这个测试命令的用户的读权限) |  
| -w | 文件是否可写 (指运行这个测试命令的用户的读权限) |  
| -x | 文件是否可执行 (指运行这个测试命令的用户的读权限) |  
| -g | 文件或目录的设置-组-ID(sgid)标记被设置。 |  
| -u | 文件的设置-用户-ID(suid)标志被设置 |  
| -k | 粘住位设置 |  
| -N | 文件最后一次读后被修改 |  
| f1 -nt f2 | 文件f1比f2新 |  
| f1 -ot f2 | 文件f1比f2旧 |  
| f1 -ef f2 | 文件f1和f2 是相同文件的硬链接 |  
| ! | "非" -- 反转上面所有测试的结果(如果没有给出条件则返回真)。|  

⚠️  

1. `-t` 这个测试选项可以用于检查脚本中是否标准输入 ([ -t 0 ])或标准输出([ -t 1 ])是一个终端。  
1. `-g` 如果一个目录的 sgid 标志被设置，在这个目录下创建的文件都属于拥有此目录的用户组，而不必是创建文件的用户所属的组。  

**整数比较操作符**  

| 比较操作符 | 描述 | 例子 |  
| ---- | ---- | ---- |  
| `-eq` | 等于 | `if [ "$a" -eq "$b" ]` |  
| `-ne` | 不等于 | `if [ "$a" -ne "$b" ]` |   
| `-gt` | 大于 | `if [ "$a" -gt "$b" ]` |  
| `-ge` | 大于等于 | `if [ "$a" -ge "$b" ]` |  
| `-lt` | 小于 | `if [ "$a" -lt "$b" ]` |  
| `-le` | 小于等于 | `if [ "$a" -le "$b" ]` |  
| `<` | 小于(在双括号里使用) | `(("$a" < "$b"))` |  
| `<=` | 小于等于 (在双括号里使用) | `(("$a" <= "$b"))` |  
| `>` | 大于 (在双括号里使用) | `(("$a" > "$b"))` |  
| `>=` | 大于等于(在双括号里使用) | `(("$a" >= "$b"))` |  

`注意：> =号 < 号在双括号外使用必须转义，不然 shell 当作重定向符号`

**字符串比较操作符**  

| 比较操作符 | 描述 | 例子 |  
| ---- | ---- | ---- |  
| = | 等于 | `if [ "$a" = "$b" ]` |  
| == | 等于，它和=是同义词。 | `if [ "$a" == "$b" ]` |  
| != | 不相等，操作符在[[ ... ]]结构里使用模式匹配. | `if [ "$a" != "$b" ]` |  
| < | 小于，依照ASCII字符排列顺序，注意"<"字符在[ ] 结构里需要转义 | `if [[ "$a" < "$b" ]]` `if [ "$a" \< "$b" ]` |  
| > | 大于，依照ASCII字符排列顺序，注意">"字符在[ ] 结构里需要转义. | `if [[ "$a" > "$b" ]]` `if [ "$a" \> "$b" ]`|   
| -z | 字符串为"null"，即是指字符串长度为零。 | - |  
| -n | 字符串不为"null"，即长度不为零。 | - |  

**混合比较操作符**  

| 比较操作符 | 描述 | 例子 |  
| ---- | ---- | ---- |  
| -a / && | 逻辑与，如果 exp1 和 exp2 都为真，则 exp1 -a exp2 返回真 | `if [ "$exp1" -a "$exp2" ]` |  
| -o / || | 逻辑或，只要 exp1 和 exp2 任何一个为真，则 exp1 -o exp2 返回真 | `if [ "$exp1" -o "$exp2" ]` |  

`字符串变量附加一个外部的字符串（"x$string"）给可能有空字符串变量比较的所有变量`  

**其他操作符**
```sh
# 逗号操作符
let "t1 = ((5 + 3, 7 - 1, 15 - 4))" # 11
```

### shell 变量
shell 的变量分为局部变量、环境变量和 shell 变量。  
> 局部变量  
> 局部变量在脚本或命令中定义，仅在当前 shell 实例中有效，其他 shell 启动的程序不能访问局部变量  
> 环境变量  
> 所有的程序，包括 shell 启动的程序，都能访问环境变量，有些程序需要环境变量来保证其正常运行，必要的时候 shell 脚本也可以定义环境变量  
> shell 变量  
> shell 变量是由 shell 程序设置的特殊变量。  
> shell 变量中有一部分是环境变量，有一部分是局部变量，这些变量保证了 shell 的正常运行  

shell 变量的一些说明  
> 变量的定义，等号两边不能有空格  
> 声明变量为局部变量，local filePath  
> 只读的变量（可以理解为常量），readonly filePath  
> 取消变量使用 unset filePath  
> 如果变量没有被设置，使用默认值 ${varName:=defaultValue}  
> declare 或 typeset 内建命令 (它们是完全相同的) 可以用来限定变量的属性    

**declare 选项说明**

| 参数 | 说明 | 例子 |  
| ---- | ---- | ---- |  
| `-r` | 只读 | `declare -r var1` |  
| `-i` | 整数 | `declare -i number;number=3;` |  
| `-a` | 数组 | `declare -a indices` |  
| `-f` | 函数 | `declare -f` 会列出所有在此脚本前面已定义的函数|  
| `-x export` | 函这样将声明一个变量作为脚本的环境变量而被导出。 | `declare -x var3` |  
| `-x var=$value` | declare命令允许在声明变量类型的时候同时给变量赋值。| `declare -x var3=373` |  

**内部变量说明**

| 内部变量 | 说明 |  
| ---- | ---- |  
| $BASH | Bash二进制程序文件的路径 |  
| $BASH_ENV |   该环境变量保存一个Bash启动文件路径，当启动一个脚本程序时会去读该环境变量指定的文件。 |  
| $BASH_SUBSHELL | 一个指示子shell(subshell)等级的变量。它是Bash版本3新加入的。 |  
| $BASH_VERSINFO[n] | 这个数组含有6个元素，指示了安装的Bash版本的信息。它和$BASH_VERSION相似，但它们还是有一些小小的不同。|  
| $BASH_VERSION | 安装在系统里的Bash版本。|  
| $DIRSTACK | 在目录堆栈里面最顶端的值(它受pushd和popd的控制) |  
| $EDITOR | 由脚本调用的默认的编辑器，一般是vi或是emacs。 |    
| $EUID | 有效用户ID |   
| $FUNCNAME | 当前函数的名字 |  
| $GLOBIGNORE | 由通配符(globbing)扩展的一列文件名模式。|  
| $GROUPS | 目前用户所属的组|  
| $HOME | 用户的家目录，通常是/home/username |  
| $HOSTNAME | 在系统启动时由一个初始化脚本中用hostname命令给系统指派一个名字。然而，gethostname()函数能设置Bash内部变量E$HOSTNAME。|  
| $HOSTTYPE | 机器类型，像$MACHTYPE一样标识系统硬件。|  
| $IFS | 内部字段分隔符 |  
| $IGNOREEOF | 忽略EOF：在退出控制台前有多少文件结尾标识（end-of-files,control-D）会被shell忽略。|  
| $LC_COLLATE | 它常常在.bashrc或/etc/profile文件里被设置，它控制文件名扩展和模式匹配的展开顺序。|  
| $LINENO | 这个变量表示在本shell脚本中该变量出现时所在的行数。它只在脚本中它出现时有意义，它一般可用于调试。|  
| $MACHTYPE | 机器类型，识别系统的硬件类型。|  
| $OLDPWD | 上一次工作的目录("OLD-print-working-directory",你上一次进入工作的目录)|
| $TZ | 时区 |  
| $MAILCHECK | 每隔多少秒检查是否有新的信件 |  
| $OSTYPE | 操作系统类型 |  
| $MANPATH man | 指令的搜寻路径 |  
| $PATH | 可执行程序文件的搜索路径。一般有/usr/bin/, /usr/X11R6/bin/, /usr/local/bin,等等。|  
| $PIPESTATUS | 此数组变量保存了最后执行的前台管道的退出状态。相当有趣的是，它不一定和最后执行的命令的退出状态一样。|  
| $PPID | 一个进程的$PPID变量保存它的父进程的进程ID(pid)。用这个变量和pidof命令比较。|  
| $PROMPT_COMMAND | 这个变量在主提示符前($PS1显示之前)执行它的值里保存的命令。|  
| $PS1 | 这是主提示符（第一提示符），它能在命令行上看见。|  
| $PS2 | 副提示符（第二提示符），它在期望有附加的输入时能看见。它显示像">"的提示。|  
| $PS3 | 第三提示符。它在一个select循环里显示 (参考例子 10-29)。|  
| $PS4 | 第四提示符，它在用-x选项调用一个脚本时的输出的每一行开头显示。它通常显示像"+"的提示。|  
| $PWD | 工作目录(即你现在所处的目录) ，它类似于内建命令pwd。|  
| $REPLY | 没有变量提供给read命令时的默认变量．这也适用于select命令的目录，但只是提供被选择的变量项目编号而不是变量本身的值。 |  
| $SECONDS | 脚本已运行的秒数。|  
| $SHELLOPTS | 已经激活的shell选项列表，它是一个只读变量。|  
| $SHLVL | SHELL的嵌套级别．指示了Bash被嵌套了多深．在命令行里，$SHLVL是1，因此在一个脚本里，它是2 |  
| $TMOUT | 如果$TMOUT环境变量被设为非零值时间值time，那么经过time这么长的时间后，shell提示符会超时．这将使此shell退出登录 |  
| $USER | 用户名，这是当前用户的用户名称 |  
| $UID | 用户ID号，这是当前用户的用户标识号，它在 /etc/passwd 文件中记录 |  

**特殊变量说明**  

| 参数处理 | 说明 |  
| ---- | ----  |  
| $#	| 传递到脚本的参数个数 |  
| $*	| 以一个单字符串显示所有向脚本传递的参数，* 改成 0[参数个数-1] 则是获取脚本相应参数 |  
| $$	| 脚本运行的当前进程 ID 号 |  
| $!	| 后台运行的最后一个进程的 ID 号 |  
| $@	| 与∗相同，但是使用时加引号，并在引号中返回每个参数。如 "@" 用「"」括起来的情况、以"1""2" … "$n" 的形式输出所有参数 |  
| $-	| 显示 Shell 使用的当前选项，与 set 命令功能相同 |  
| $?	| 显示最后命令的退出状态。0 表示没有错误，其他任何值表明有错误 |    

### shell 数组
shell 支持一维数组（不支持多维数组），并且没有限定数组的大小
```sh
# 数组的定义和初始化
fooArray=(
	"key=0"
	"key=1"
)
# 也可以这样 
# fooArray[0]="key=0"
# fooArray[1]="key=1"

echo "数组第一个元素: ${fooArray[0]}"
echo "数组第二个元素: ${fooArray[1]}"
# 获取数组所有元素 ${fooArray[@]} 或 ${fooArray[*]}
echo "数组的长度是：${#fooArray[@]}"

# 遍历数组
for item in  ${fooArray[*]}
do
    echo $item;
done
```

### shell 字符串
```sh
# 字符串的拼接
name="vikey"
echo "hello, "$name" !"
echo "hello, ${name} !"

# 获取字符串的长度 ${#string} 
echo ${#name}

# 索引
echo `expr index 'vikey' 'e'` # 4

# 提取子字符串 ${string:position:length}
echo ${name:1:3} # ike

# 移动
# ${string#substring} 从 $string 左边开始，剥去最短匹配 $substring 子串
# ${string##substring} 从 $string 左边开始，剥去最长匹配 $substring 子串
# ${string%substring} 从 $string 结尾开始，剥去最短匹配 $substring 子串
# ${string%%substring} 从 $string 结尾开始，剥去最长匹配 $substring 子串
echo ${name##v*i} # key

# 字符串比对
string1='compare1'
string2='compare2'
if [ "$string1" = "$string2" ]
# if [ "X$string1" = "X$string2" ] 会更安全，它为了防止其中有一个字符串为空时产生错误信息

then
   # ...
fi

# 遍历字符串列表
list="Alabama Alaska Arizona"
list=$list" Connecticut"
for state in $list
do
  echo "Have you ever visited $state"
done

```

### shell 执行命令
```sh
users=$(who | wc -l)
echo "Logged in user are $users"

scriptName=`basename $0`
echo "The name of this script is $script_name."

rm `cat filePath` # "filePath" 包含了需要被删除的文件列表
```

### shell 运算
```sh
# 自增运算 (( index++ ))
# 自减运算 (( index-- ))

index=1
# 运算的各种姿势
index=`expr $index + 1`   
index=$(($index+1))
index=$[$index+1]
let "index+=1"
index=`echo $index+1 | bc`


```

### shell if 语句
```sh
if [ condition ]; then
  # condition is true.
else
  # condition is false.
fi   
```

### shell 选择
case 语句类似于 switch
```sh
#!/bin/bash
read Keypress
case "$Keypress" in
  [[:lower:]]   ) echo "Lowercase letter";;
  [[:upper:]]   ) echo "Uppercase letter";;
  [0-9]         ) echo "Digit";;
  *             ) echo "Punctuation, whitespace, or other";;
esac
exit 0
```

select 结构是建立菜单的另一种工具
```sh
#!/bin/bash
echo '选择你喜欢的蔬菜:  '
choiceFood(){
  select vegetable
  # [in list] 被忽略, 所以'select'用传递给函数的参数.
  do
    echo "你最喜欢的蔬菜是  $vegetable"
    break
  done
}

choiceFood "胡萝卜" "土豆" "洋葱"
```

### shell 循环
```sh
# for arg in [list]
# do
#     ...
# done
for planet in Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune Pluto
do
  echo $planet
done

# while [condition]
# do
#     ...
# done
index=0
limit=10
while [ "$index" -lt "$limit" ]
do
  echo -n "${index} "     # -n 会阻止产生新行
  index=`expr $index + 1`   
done 

# 知道条件为真才停止
# until [condition]
# do
#     ...
# done
index=0
limit=10
until [ "$index" -gt "$limit" ]
do
  echo -n "${index} "
  index=`expr $index + 1`   
done 

# 循环内的控制
# continue 跳出当前循环，continue 2 跳出两层循环
# break 跳出所有循环
```

### shell 函数
一个函数是一个子程序，用于实现一串操作的代码块，它是完成特定任务的 "黑盒子"
```sh
# function 可以不写
function functionName { 
  # ... 
  # 可以选择定义 return 
}

functionName()
# 获取返回值 $?
```

### shell 正则
**正则表达式的基本组成部分**

| 正则表达式 | 描述 | 示例 | Basic RegEx | Extended RegEx | Python RegEx | Perl regEx | 
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | 
| \ | 转义符，将特殊字符进行转义，忽略其特殊意义 | a\.b匹配a.b，但不能匹配ajb，.被转义为特殊意义 | \ | \ | \ | \ | 
| ^ | 匹配行首，awk中，^则是匹配字符串的开始 | ^tux匹配以tux开头的行 | ^ | ^ | ^ | ^ |
| $ | 匹配行尾，awk中，$则是匹配字符串的结尾 | tux$匹配以tux结尾的行 | $ | $ | $ | $ |
| . | 匹配除换行符\n之外的任意单个字符，awk则中可以 | ab.匹配abc或bad，不可匹配abcd或abde，只能匹配单字符 | . | . | . | . |
| [] | 匹配包含在[字符]之中的任意一个字符 | coo[kl]可以匹配cook或cool | [] | [] | [] | [] |
| [^] | 匹配[^字符]之外的任意一个字符 | 123[^45]不可以匹配1234或1235，1236、1237都可以 | [^] | [^] | [^] | [^] |
| [-] | 匹配[]中指定范围内的任意一个字符，要写成递增 | [0-9]可以匹配1、2或3等其中任意一个数字 | [-] | [-] | [-] | [-] |
| ? | 匹配之前的项1次或者0次 | colou?r可以匹配color或者colour，不能匹配colouur | 不支持 | ? | ? | ? |
| + | 匹配之前的项1次或者多次 | sa-6+匹配sa-6、sa-666，不能匹配sa- | 不支持 | + | + | + |
| \* | 匹配之前的项0次或者多次 | co\*l匹配cl、col、cool、coool等 | \* | \* | \* | \* |
| () | 匹配表达式，创建一个用于匹配的子串 | ma(tri)?匹配max或maxtrix | 不支持 | () | () | () |
| {n} | 匹配之前的项n次，n是可以为0的正整数 | [0-9]{3}匹配任意一个三位数，可以扩展为[0-9][0-9][0-9] | 不支持 | {n} | {n} | {n} |
| {n,} | 之前的项至少需要匹配n次 | [0-9]{2,}匹配任意一个两位数或更多位数 | 不支持 | {n,} | {n,} | {n,} |
| {n,m} | 指定之前的项至少匹配n次，最多匹配m次，n&lt;=m | [0-9]{2,5}匹配从两位数到五位数之间的任意一个数字 | 不支持 | {n,m} | {n,m} | {n,m} |
| `|` | 交替匹配|两边的任意一项 | `ab(c|d)`匹配abc或abd | 不支持 | `|` | `|` | `|` |

POSIX 字符类是一个形如[:...:]的特殊元序列（meta sequence），他可以用于匹配特定的字符范围  
**POSIX 字符类**

| 正则表达式 | 描述 | 示例 | Basic RegEx | Extended RegEx | Python RegEx | Perl regEx |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | 
| [:alnum:] | 匹配任意一个字母或数字字符 | [[:alnum:]]+ | [:alnum:] | [:alnum:] | [:alnum:] | [:alnum:] |
| [:alpha:] | 匹配任意一个字母字符（包括大小写字母） | [[:alpha:]]{4} | [:alpha:] | [:alpha:] | [:alpha:] | [:alpha:] |
| [:blank:] | 空格与制表符（横向和纵向） | [[:blank:]]* | [:blank:] | [:blank:] | [:blank:] | [:blank:] |
| [:digit:] | 匹配任意一个数字字符 | [[:digit:]]? | [:digit:] | [:digit:] | [:digit:] | [:digit:] |
| [:lower:] | 匹配小写字母 | [[:lower:]]{5,} | [:lower:] | [:lower:] | [:lower:] | [:lower:] |
| [:upper:] | 匹配大写字母 | ([[:upper:]]+)? | [:upper:] | [:upper:] | [:upper:] | [:upper:] |
| [:punct:] | 匹配标点符号 | [[:punct:]] | [:punct:] | [:punct:] | [:punct:] | [:punct:] |
| [:space:] | 匹配一个包括换行符、回车等在内的所有空白符 | [[:space:]]+ | [:space:] | [:space:] | [:space:] | [:space:] |
| [:graph:] | 匹配任何一个可以看得见的且可以打印的字符 | [[:graph:]] | [:graph:] | [:graph:] | [:graph:] | [:graph:] |
| [:xdigit:] | 任何一个十六进制数（即：0-9，a-f，A-F） | [[:xdigit:]]+ | [:xdigit:] | [:xdigit:] | [:xdigit:] | [:xdigit:] |
| [:cntrl:] | 任何一个控制字符（[ASCII](http://zh.wikipedia.org/zh/ASCII)字符集中的前32个字符) | [[:cntrl:]] | [:cntrl:] | [:cntrl:] | [:cntrl:] | [:cntrl:] |
| [:print:] | 任何一个可以打印的字符 | [[:print:]] | [:print:] | [:print:] | [:print:] | [:print:] |

元字符（meta character）是一种Perl风格的正则表达式，只有一部分文本处理工具支持它，并不是所有的文本处理工具都支持  
**元字符**

| 正则表达式 | 描述 | 示例 | Basic RegEx | Extended RegEx | Python RegEx | Perl regEx |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | 
| \b | 单词边界 | \bcool\b 匹配cool，不匹配coolant | \b | \b | \b | \b |
| \B | 非单词边界 | cool\B 匹配coolant，不匹配cool | \B | \B | \B | \B |
| \d | 单个数字字符 | b\db 匹配b2b，不匹配bcb | 不支持 | 不支持 | \d | \d |
| \D | 单个非数字字符 | b\Db 匹配bcb，不匹配b2b | 不支持 | 不支持 | \D | \D |
| \w | 单个单词字符（字母、数字与_） | \w 匹配1或a，不匹配& | \w | \w | \w | \w |
| \W | 单个非单词字符 | \W 匹配&，不匹配1或a | \W | \W | \W | \W |
| \n | 换行符 | \n 匹配一个新行 | 不支持 | 不支持 | \n | \n |
| \s | 单个空白字符 | x\sx 匹配x x，不匹配xx | 不支持 | 不支持 | \s | \s |
| \S | 单个非空白字符 | x\S\x 匹配xkx，不匹配xx | 不支持 | 不支持 | \S | \S |
| \r | 回车 | \r 匹配回车 | 不支持 | 不支持 | \r | \r |
| \t | 横向制表符 | \t 匹配一个横向制表符 | 不支持 | 不支持 | \t | \t |
| \v | 垂直制表符 | \v 匹配一个垂直制表符 | 不支持 | 不支持 | \v | \v |
| \f | 换页符 | \f 匹配一个换页符 | 不支持 | 不支持 | \f | \f |

### shell 其他
```sh
# 获取用户输入
echo -n "Enter your name:"
read name
echo "Hello $name, welcome~"
# 如果不指定变量，read 命令就会把它收到的任何数据都放到特殊环境变量 REPLY 中
read -p "Enter a number:"
echo $REPLY
# 限时 5s，输入放到变量 number
read -t 5 -p "Please enter the number again:" number
# 隐形输入
read -s -p "Please enter your password: " password

# 定时执行脚本
at -f at.sh 23:59

# msgbox 部件 dialog 的使用
dialog --title text --msgbox "This is a test" 10 20

# 常用系统脚本
# 
# 查看僵尸进程
ps -al | gawk '{print $2,$4}' | grep Z
# 查看在线用户数
uptime | sed 's/user.*$//' | gawk '{print $NF}'
# 查看内存使用百分比
free | sed -n '2p' | gawk 'x = int(( $3 / $2 ) * 100) {print x}' | sed 's/$/%/'
# 查看磁盘使用百分比
df -h /dev/sda1 | sed -n '/% \//p' | gawk '{ print $5 }'
# 邮箱验证
gawk --re-interval '/^([a-zA-Z0-9_\-\.\+]+)@([a-zA-Z0-9_\-\+]+)\.([a-zA-Z]{2,5})/{print $0}'
```

