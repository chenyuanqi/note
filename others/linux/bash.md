
```bash
#!/bin/bash

# 常用环境变量
# HOME  用户的主目录
# LANG  字符集以及语言编码，比如zh_CN.UTF-8
# PATH  由冒号分开的目录列表，当输入可执行程序名后，会搜索这个目录列表
# PS1   Shell提示符
# PWD   当前工作目录
# SHELL Shell的名字
# TERM  终端类型名，即终端仿真器所用的协议
# USER  用户名

# 常用内置变量
# $? 代表最近一次程式返回值
# $$ 代表脚本的 PID（进程 ID）
# $# 代表脚本参数的数量
# $@ 代表脚本所有参数
# $0 代表脚本本身的路径名
# $1 代表脚本的第一个参数

# 条件的与或非
# -a   与
# -o   或
# -not 非（! 也可以）

# 文件的状态判断
# file1 -ef file2	file1 和 file2 拥有相同的索引号（通过硬链接两个文件名指向相同的文件）
# file1 -nt file2	file1 新于 file2
# file1 -ot file2	file1 早于 file2
# -b file	file 存在并且是一个块（设备）文件
# -c file	file 存在并且是一个字符（设备）文件
# -d file	file 存在并且是一个目录
# -e file	file 存在
# -f file	file 存在并且是一个普通文件
# -g file	file 存在并且设置了组 ID
# -G file	file 存在并且由有效组 ID 拥有
# -k file	file 存在并且设置了它的 "sticky bit"
# -L file	file 存在并且是一个符号链接。
# -O file	file 存在并且由有效用户 ID 拥有
# -p file	file 存在并且是一个命名管道
# -r file	file 存在并且可读（有效用户有可读权限）
# -s file	file 存在且其长度大于零
# -S file	file 存在且是一个网络 socket
# -t fd	    fd 是一个定向到终端／从终端定向的文件描述符, 这可以被用来决定是否重定向了标准输入／输出错误
# -u file	file 存在并且设置了 setuid 位
# -w file	file 存在并且可写（有效用户拥有可写权限）
# -x file	file 存在并且可执行（有效用户有执行／搜索权限）

# 整数表达式的判断
# integer1 -eq integer2	integer1 等于 integer2
# integer1 -ne integer2	integer1 不等于 integer2
# integer1 -le integer2	integer1 小于或等于 integer2
# integer1 -lt integer2	integer1 小于 integer2
# integer1 -ge integer2	integer1 大于或等于 integer2
# integer1 -gt integer2	integer1 大于 integer2

# 字符串的判断
# string	字符串 string 不为 null
# -n string	字符串 string 的长度大于零
# -z string	字符串 string 的长度为零
# string1 = string2  字符串 string1 和字符串 string2 相同
# string1 == string2 字符串 string1 和字符串 string2 相同
# string1 != string2 字符串 string1 和字符串 string2 不相同
# string1 > string2	 字符串 sting1 排列在字符串 string2 之后
# string1 < string2	 字符串 string1 排列在字符串 string2 之前

# 执行命令的方式
command
`command`
$(command)

# 读取输入
echo "What's your name?"
read user_name

str="xxx"
multiple_str=$(cat <<EOF
xxx
xxx
xxx
EOF
)
# 变量的默认值设置
echo ${foo:-"default_value"}
# 变量存在且不为 null，返回它的值，否则将它设为 default_value 且返回 default_value
echo ${foo:="default_value"}
# 变量名存在且不为null，返回 default_value，否则返回 null
echo ${foo:+"default_value"}

# 注意单双引号的区别，单引号不解析变量不解析转义符号
echo "${str}"
# 字符变量内部进行字符串代换，x => X
echo ${str/x/X}
# 字符变量截取 ${str:start:length}
echo ${str:0:1}

# 数学表达式的计算，支持 +、-、*、/、%、**
$((number1 + number2))
# 数字自增，相当于 number++
number= `expr ${number} + 1`

# ====================
# 复杂类型：数组
arr_var[0]=0
arr_var[1]=1
arr_var[2]=2
# 或者这样初始化数组：week_day=('Sun' 'Mon' 'Tue' 'Wed' 'Thu' 'Fri' 'Sat')
# 关联数组
color_arr["red"]="#ff0000"
color_arr["green"]="#00ff00"
color_arr["blue"]="#0000ff"
echo ${color_arr['red']}
# 获取指定位置数组成员
echo ${arr_var[2]} 
# 数组长度，@ 换成数组下标则是返回指定数组成员的长度
echo ${#arr_var[@]}
# 遍历数组值，arr_var[*] 或 arr_var[@]
for index in ${arr_var[*]}
do 
    echo ${index} 
done
# 遍历数组键，!arr_var[*] 或 !arr_var[@]
for index in ${!arr_var[@]}
do 
    echo ${index} 
done
# 新数组成员拼接在数组后
arr_var+=(3 4 5)
# 不指定数组下标，默认修改数组第一个成员
arr_var=100
# 数组的复制
arr_copy=("${arr_var[@]}")
# 删除数组某个成员，单双引号可有可无
unset 'arr_var[2]'
# 模式删除数组元素
declare -a patter=( ${arr_var[@]/2*/} )
# 清空数组
unset arr_var
# ====================

if [ 判断条件 ]; 
then
    # if 条件为真的分支代码
elif [ 判断条件 ]; 
then
    # elif 条件为真的分支代码    
else
    # if 条件为假的分支代码
fi

# 根据上一个指令执行结果（或条件语句）决定是否执行下一个指令
command_1 || command_2
command_1 && command_2

# case 语句（switch）
case "${number}" in
    # 列出需要匹配的字符串
    0) echo "There is a zero.";;
    1) echo "There is a one.";;
    *) echo "It is not null.";;
esac

# 循环遍历给定的参数序列，变量 $number 的值会被打印 3 次
for ((number=1; number<=3; number++))
do
    echo $number
done
# 或者这样
for number in {1..3}
do
    echo "${number}"
done

# for 用于输出多个文件的内容，多个文件也可使用正则的方式
for content in file1 file2
do
    cat "${content}"
done

# while 循环（条件满足则循环）
while [ true ]
do
    # continue
    # break
    statements
done

# until 循环（条件满足则终止循环）
until [ condition ]; 
do
  statements
done

# 函数的定义，关键字 function 可省略（return 后也可以没有值）
function foo(number)
{
    # 可以使用局部变量
    local temp
    # ...
    return return_value
}
# 函数的调用
foo 99

# 休息 1 秒
sleep 1
# 休息 1 分钟
sleep 1m

# 退出
exit exit_code
```