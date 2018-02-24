
### 字符串常用方法
```python
# 定义及初始化字符串
str_var = 'I am a string' # 或者 str_var = "I am a string"
# 注意，这里的单引号和双引号没有 PHP 效率上的问题，只是为了避免字符串内的引号转义，如 "I'm a string"

# 多行字符串的定义
multiple_str_var = """
I
am
a
string
"""
# 实际输出：'\nI\nam\na\nstring\n'
# 多行字符串也可以是三个单引号，通常用于多行的注释

# 字符串的换行
multiple_str_var = ("I"
"am"
"a"
"string")
# 或者这样（方法的链式操作也是如下写法）
multiple_str_var = "I" \
"am" \
"a" \
"string"

# 字符串替换
str_var.replace("string", "str") # 返回 'I am a str', 原字符串不改变

# 大小写转换
str_var.lower() # 返回 'i am a string', 原字符串不改变
str_var.upper() # 返回 'I AM A STRING', 原字符串不改变
# 首字母大写
str_var.capitalize() # 返回 'I am a string', 原字符串不改变
# 每个单词首字母大写
str_var.title() # 返回  'I Am A String', 原字符串不改变
# 大小写互换
str_var.swapcase() # 返回 'i AM A STRING'，原字符串不变

# 字符串填充 
# 中部扩充 center(str_len, char)
str_var.center(30, '$') # 返回 '$$$$$$$$I am a string$$$$$$$$$'
# 零填充 zfill(str_len)
str_var.zfill(30) # 返回 '00000000000000000I am a string'
# 左右填充 ljust(str_len, char) / rjust(str_len, char)
str_var.ljust(30, 'X') # 返回 'I am a stringXXXXXXXXXXXXXXXXX'
str_var.rjust(30, 'X') # 返回 'XXXXXXXXXXXXXXXXXI am a string'
# 制表符转换指定空格，默认转为 8 个空格
str_var = "I\tam\ta\tstring"
str_var.expandtabs(1) # 返回 'I am a string'

# 字符串查找
str_var.index('a') # 返回 2，sub 也可以是字符串
str_var.find('a') # 返回 2
# index 和 find 功能相同，区别在于 find 查找失败会返回 -1，不会影响程序运行，而 index 会抛 ValueError 异常

# 包含
"am" in str_var # 返回 True
"am" not in str_var # 返回 False

# 字符统计 count(char, start_offset, end_offset)
str_var.count('a') # 返回 2

# 字符串的切片
str_var[:4] # 返回 'I am'
# 与列表类似，str_var[start:end:step]

# 去除多余空格,（可指定去除的字符，默认空格）
str_var = '  I am a string  '
str_var.lstrip() # 返回 'I am a string  ', 原字符串不改变
str_var.rstrip() # 返回 '  I am a string', 原字符串不改变
str_var.strip() # 返回 'I am a string', 原字符串不改变

# 字符串分割与连接
str_var = '1, 2, 3, 4, 5'
list_var = str_var.split(',') # 返回 ['1', ' 2', ' 3', ' 4', ' 5']
' '.join(list_var) # 返回 '1  2  3  4  5'
# 需要注意的是，list_var 中的项必须是字符串，否则抛 TypeError 异常（除了列表，元组也可以 join，也必须遵循字符串规则）

# 字符串的拼接
str_var = "Hello " + "Python" # 返回 'Hello Python'
# 同样的，这样也可以：str_var = "Hello " "Python"
str_var * 2 # 返回 'Hello PythonHello Python'

# 字符串的判断
# startswith(prefix[,start[,end]])
str_var.startswith('H') # 返回 True
# endswith(suffix[,start[,end]])
str_var.endswith('H') # 返回 False
# 是否全是字母和数字，并至少有一个字符
str_var.isalnum() # 返回 False
# 是否全是字母，并至少有一个字符
str_var.isalpha() # 返回 False
# 是否全是数字，并至少有一个字符 
str_var.isdigit() # 返回 False
str_var.isnumeric() # 返回 False
# isdigit 判断的是 Unicode数字，byte数字（单字节），全角数字（双字节），罗马数字，无异常处理
# isnumeric 判断的是 Unicode数字，全角数字（双字节），罗马数字，汉字数字，有异常处理
# isdecimal 判断的是 Unicode数字，，全角数字（双字节），有异常处理
# 是否全是空白字符，并至少有一个字符
str_var.isspace() # 返回 False
# 是否全是小写
str_var.islower() # 返回 False
# 是否全是大写
str_var.isupper() # 返回 False
# 是否是首字母大写
str_var.istitle() # 返回 True


# 格式化字符串
'{} {} {}'.format('a', 'b', 'c') # 返回 'a b c'
'{2} {1} {0}'.format('a', 'b', 'c') # 返回 'c b a'
'{color} {i} {f}'.format(i=1, f=1.5, color='red') # 返回 'red 1 1.5'
'{color} {0} {x} {1}'.format(0, '1', x = 1.5, color='green') # 返回 'green 0 1.5 1'
from math import pi
'{0:3} {1:3d} {2:3.2f}'.format('foo', 5, 2 * pi) # 返回 'foo   5 6.28'
```

### 中文编码问题
计算机存储的一切数据都是由一串 0 或 1 的字节序列构成，1 byte = 8 bit；
字符是一个符号，比如一个汉字；
字节方便存储和网络传输，而字符用于显示，方便阅读。

编码 (encode)：字符 => 字节；
解码 (decode)：字节 => 字符

在 2.x，sys.getdefaultencoding() 的值为 'ascii'，而 ascii 不能处理中文；
在 3.x, sys.getdefaultencoding() 的值 'utf-8'

在 2.x 中，字符串被分为 unicode 和 str 两种类型，本质上 str 是一串字节序列，unicode 是字符（符号）；
```python
# UnicodeEncodeError
def main():
    name = u'大蟒蛇'
    with open("output.txt", "w") as f:
        f.write(name)
# 抛 UnicodeEncodeError 异常的原因：
# 调用 write 方法时，Python 会先判断字符串是什么类型，
# 如果是 str，就直接写入文件；
# 如果字符串是 unicode 类型，那么它会先调用 encode 方法把 unicode 转换成 str 字节，才保存到文件，而 encode 方法会使用 python 默认的 ascii 码来编码（不识别中文）
# 所以，正确的解决之道是 n = u'大蟒蛇'.encode('utf-8')

# UnicodeDecodeError
python_en = u'Python'
python_cn = '大蟒蛇'
python_en + python_cn
# 抛 UnicodeDecodeError 异常的原因
# str 与 unicode 执行 + 操作时，
# Python 会把 str 隐式地转换（解码）成和 python_en 一样的 unicode 类型，
# 但是，Python 使用默认的是 ascii 编码来转换的
# 正确的解法是 u'大蟒蛇'
```

在 3.x 中，字符串只保留 str 类型。
那么，在 2.x 中，为避免乱码，这句头部代码就显得尤为重要了 # -*- coding:utf-8 -*-

### json 模块
```python
import json

dict_var = {
    "name": "Python"
}

# dumps: 字典 => json （卸货）
json_var = json.dumps(dict_var) # '{"name": "Python"}'
# loads: json => 字典 （装货）
json.loads(json_var) # {'name': 'Python'}

# 处理 json 文件
with open('static/json_file.json', 'w') as f:
    # 字典 => json 数据, 并写入 json_file.json 中
    json.dump(dict_var, f) # {"name": "Python"}

with open('static/json_file.json', 'r') as f:
    # json 数据 => 字典，并赋值给 dict_var
    dict_var = json.load(f) # {"name": "Python"}
```
