
### shell 脚本
shell 脚本以.sh 结尾的或者不需要后缀都可以。
```sh
#!/bin/bash
# 文件名 hello.sh
NUM=100
echo "学习 shell $NUM 分"
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
index=1
index=`expr $index + 1`   
index=$(($index+1))
let "index+=1"

# 自减运算 (( index-- ))
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
# continue 跳出当前循环
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
