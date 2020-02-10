
### 字符常量
```python
# 包 0 到 9 的数字
string.digits '0123456789'

# 包含 26 字母的小写
string.ascii_lowercase # 'abcdefghijklmnopqrstuvwxyz'
# 包含 26 字母的大写
string.ascii_uppercase # 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
# 包含 26 个字母的大小写
string.ascii_letters # 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'

# 包含八进制的所有数字
string.octdigits # '01234567'
# 包含十六进制所有数字和字母
string.hexdigits # '0123456789abcdefABCDEF'

# 包含所有空白符
string.whitespace # ' \t\n\r\x0b\x0c'

# 包含所有标点符号
string.punctuation # '!"#$%&\'()*+,-./:;<=>?@[\\]^_`{|}~'

# 可打印的所有 ascii 字符
string.printable # '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!"#$%&\'()*+,-./:;<=>?@[\\]^_`{|}~ \t\n\r\x0b\x0c'
```

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
# 编码和解码
string = 'a string'
bstring = string.encode('ascii') # b'a string'
bstring.decode('ascii') # a string

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
那么，在 2.x 中，为避免乱码，这句头部代码就显得尤为重要了 # -\*- coding:utf-8 -\*-

以 u 开头的字符串称为 Unicode 字符串，r 开头的是原始字符串（定义原始字符串可以用小写 r 或者大写 R 开头，比如 r"\b" 或者  R"\b" 都是允许的）。  
```python
# 它们都是由普通文本字符组成的串，在这里没什么区别
foo = "hello"
bar = r"hello"
foo is bar # True

foo = "\n"
bar = r"\n"
len(foo) # 1
len(bar) # 2
foo == bar # False
# "\n" 是一个转义字符，它在 ASCII 中表示换行符
# r"\n" 是一个原始字符串，原始字符串不对特殊字符进行转义，它就是你看到的字面意思，由 “\” 和 “n” 两个字符组成的字符串
```


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

### pickle 模块
```python
import pickle

dict_var = {
    "name": "Python"
}

# dumps: 字典 => 字节 （序列化）
pickle_var = pickle.dumps(dict_var) # b'\x80\x03}q\x00X\x04\x00\x00\x00nameq\x01X\x06\x00\x00\x00Pythonq\x02s.'
# loads: 字节 => 字典 （反序列化）
pickle.loads(pickle_var) # {'name': 'Python'}

# 处理 pickle 文件
with open('static/pickle_file.txt', 'wb') as f:
    # 字典 => 字节，写入 pickle_file.txt 文件 （序列化）
    pickle.dump(dict_var, f)

with open('static/pickle_file.txt', 'rb') as f:
    # 字节 => 字典，并赋值给 dict_var （反序列化）
    dict_var = pickle.load(f) # {'name': 'Python'}
```

### 字符串使用示例
#### 1 反转字符串

```python
st="python"
#方法1
''.join(reversed(st))
#方法2
st[::-1]
```

#### 2 字符串切片操作

```python
字符串切片操作——查找替换3或5的倍数
In [1]:[str("java"[i%3*4:]+"python"[i%5*6:] or i) for i in range(1,15)]
OUT[1]:['1',
 '2',
 'java',
 '4',
 'python',
 'java',
 '7',
 '8',
 'java',
 'python',
 '11',
 'java',
 '13',
 '14']
```
#### 3 join串联字符串
```python
In [4]: mystr = ['1',
   ...:  '2',
   ...:  'java',
   ...:  '4',
   ...:  'python',
   ...:  'java',
   ...:  '7',
   ...:  '8',
   ...:  'java',
   ...:  'python',
   ...:  '11',
   ...:  'java',
   ...:  '13',
   ...:  '14']

In [5]: ','.join(mystr) #用逗号连接字符串
Out[5]: '1,2,java,4,python,java,7,8,java,python,11,java,13,14'
```

#### 4 字符串的字节长度

```python
def str_byte_len(mystr):
    return (len(mystr.encode('utf-8')))


str_byte_len('i love python')  # 13(个字节)
str_byte_len('字符')  # 6(个字节)
```

### 正则使用示例
```python
import re
```

#### 1 查找第一个匹配串

```python
s = 'i love python very much'
pat = 'python' 
r = re.search(pat,s)
print(r.span()) #(7,13)
```

#### 2 查找所有1的索引

```python
s = '山东省潍坊市青州第1中学高三1班'
pat = '1'
r = re.finditer(pat,s)
for i in r:
    print(i)

# <re.Match object; span=(9, 10), match='1'>
# <re.Match object; span=(14, 15), match='1'>
```

#### 3 \d 匹配数字[0-9]
findall找出全部位置的所有匹配
```python
s = '一共20行代码运行时间13.59s'
pat = r'\d+' # +表示匹配数字(\d表示数字的通用字符)1次或多次
r = re.findall(pat,s)
print(r)
# ['20', '13', '59']
```

#### 4 匹配浮点数和整数

?表示前一个字符匹配0或1次
```python
s = '一共20行代码运行时间13.59s'
pat = r'\d+\.?\d+' # ?表示匹配小数点(\.)0次或1次，这种写法有个小bug，不能匹配到个位数的整数
r = re.findall(pat,s)
print(r)
# ['20', '13.59']

# 更好的写法：
pat = r'\d+\.\d+|\d+' # A|B，匹配A失败才匹配B
```
#### 5 ^匹配字符串的开头

```python
s = 'This module provides regular expression matching operations similar to those found in Perl'
pat = r'^[emrt]' # 查找以字符e,m,r或t开始的字符串
r = re.findall(pat,s)
print(r)
# [],因为字符串的开头是字符`T`，不在emrt匹配范围内，所以返回为空
IN [11]: s2 = 'email for me is guozhennianhua@163.com'
re.findall('^[emrt].*',s2)# 匹配以e,m,r,t开始的字符串，后面是多个任意字符
Out[11]: ['email for me is guozhennianhua@163.com']

```
#### 6 re.I 忽略大小写

```python
s = 'That'
pat = r't' 
r = re.findall(pat,s,re.I)
In [22]: r
Out[22]: ['T', 't']
```
#### 7 理解compile的作用
如果要做很多次匹配，可以先编译匹配串：
```python
import re
pat = re.compile('\W+') # \W 匹配不是数字和字母的字符
has_special_chars = pat.search('ed#2@edc') 
if has_special_chars:
    print(f'str contains special characters:{has_special_chars.group(0)}')

###输出结果: 
 # str contains special characters:#   

### 再次使用pat正则编译对象 做匹配
again_pattern = pat.findall('guozhennianhua@163.com')
if '@' in again_pattern:
    print('possibly it is an email')

```

#### 8 使用()捕获单词，不想带空格
使用`()`捕获
```python
s = 'This module provides regular expression matching operations similar to those found in Perl'
pat = r'\s([a-zA-Z]+)'  
r = re.findall(pat,s)
print(r) #['module', 'provides', 'regular', 'expression', 'matching', 'operations', 'similar', 'to', 'those', 'found', 'in', 'Perl']
```
看到提取单词中未包括第一个单词，使用`?`表示前面字符出现0次或1次，但是此字符还有表示贪心或非贪心匹配含义，使用时要谨慎。
```python
s = 'This module provides regular expression matching operations similar to those found in Perl'
pat = r'\s?([a-zA-Z]+)'  
r = re.findall(pat,s)
print(r) #['This', 'module', 'provides', 'regular', 'expression', 'matching', 'operations', 'similar', 'to', 'those', 'found', 'in', 'Perl']
```

#### 9 split分割单词
使用以上方法分割单词不是简洁的，仅仅是为了演示。分割单词最简单还是使用`split`函数。
```python
s = 'This module provides regular expression matching operations similar to those found in Perl'
pat = r'\s+'  
r = re.split(pat,s)
print(r) # ['This', 'module', 'provides', 'regular', 'expression', 'matching', 'operations', 'similar', 'to', 'those', 'found', 'in', 'Perl']

### 上面这句话也可直接使用str自带的split函数：
s.split(' ') #使用空格分隔

### 但是，对于风格符更加复杂的情况，split无能为力，只能使用正则

s = 'This,,,   module ; \t   provides|| regular ; '
words = re.split('[,\s;|]+',s)  #这样分隔出来，最后会有一个空字符串
words = [i for i in words if len(i)>0]
```

#### 10 match从字符串开始位置匹配
注意`match`,`search`等的不同：
1) match函数
```python
import re
### match
mystr = 'This'
pat = re.compile('hi')
pat.match(mystr) # None
pat.match(mystr,1) # 从位置1处开始匹配
Out[90]: <re.Match object; span=(1, 3), match='hi'>
```
2) search函数
search是从字符串的任意位置开始匹配
```python
In [91]: mystr = 'This'
    ...: pat = re.compile('hi')
    ...: pat.search(mystr)
Out[91]: <re.Match object; span=(1, 3), match='hi'>
```

#### 11 替换匹配的子串
`sub`函数实现对匹配子串的替换
```python
content="hello 12345, hello 456321"    
pat=re.compile(r'\d+') #要替换的部分
m=pat.sub("666",content)
print(m) # hello 666, hello 666
```

#### 12 贪心捕获
(.*)表示捕获任意多个字符，尽可能多的匹配字符
```python
content='<h>ddedadsad</h><div>graph</div>bb<div>math</div>cc'
pat=re.compile(r"<div>(.*)</div>")  #贪婪模式
m=pat.findall(content)
print(m) #匹配结果为： ['graph</div>bb<div>math']
```
#### 13 非贪心捕获
仅添加一个问号(`?`)，得到结果完全不同，这是非贪心匹配，通过这个例子体会贪心和非贪心的匹配的不同。
```python
content='<h>ddedadsad</h><div>graph</div>bb<div>math</div>cc'
pat=re.compile(r"<div>(.*?)</div>")
m=pat.findall(content)
print(m) # ['graph', 'math']
```
非贪心捕获，见好就收。

#### 14 常用元字符总结

    . 匹配任意字符  
    ^ 匹配字符串开始位置 
    $ 匹配字符串中结束的位置 
    * 前面的原子重复0次、1次、多次 
    ? 前面的原子重复0次或者1次 
    + 前面的原子重复1次或多次
    {n} 前面的原子出现了 n 次
    {n,} 前面的原子至少出现 n 次
    {n,m} 前面的原子出现次数介于 n-m 之间
    ( ) 分组,需要输出的部分

#### 15 常用通用字符总结

    \s  匹配空白字符 
    \w  匹配任意字母/数字/下划线 
    \W  和小写 w 相反，匹配任意字母/数字/下划线以外的字符
    \d  匹配十进制数字
    \D  匹配除了十进制数以外的值 
    [0-9]  匹配一个0-9之间的数字
    [a-z]  匹配小写英文字母
    [A-Z]  匹配大写英文字母

#### 14 密码安全检查

密码安全要求：1)要求密码为6到20位; 2)密码只包含英文字母和数字

```python
pat = re.compile(r'\w{6,20}') # 这是错误的，因为\w通配符匹配的是字母，数字和下划线，题目要求不能含有下划线
# 使用最稳的方法：\da-zA-Z满足`密码只包含英文字母和数字`
pat = re.compile(r'[\da-zA-Z]{6,20}')
```
选用最保险的`fullmatch`方法，查看是否整个字符串都匹配：
```python
pat.fullmatch('qaz12') # 返回 None, 长度小于6
pat.fullmatch('qaz12wsxedcrfvtgb67890942234343434') # None 长度大于22
pat.fullmatch('qaz_231') # None 含有下划线
pat.fullmatch('n0passw0Rd')
Out[4]: <re.Match object; span=(0, 10), match='n0passw0Rd'>
```

#### 15 爬取百度首页标题

```python
import re
from urllib import request

#爬虫爬取百度首页内容
data=request.urlopen("http://www.baidu.com/").read().decode()

#分析网页,确定正则表达式
pat=r'<title>(.*?)</title>'

result=re.search(pat,data)
print(result) <re.Match object; span=(1358, 1382), match='<title>百度一下，你就知道</title>'>

result.group() # 百度一下，你就知道
```

#### 16 批量转化为驼峰格式(Camel)

数据库字段名批量转化为驼峰格式

分析过程

```python
# 用到的正则串讲解
# \s 指匹配： [ \t\n\r\f\v]
# A|B：表示匹配A串或B串
# re.sub(pattern, newchar, string): 
# substitue代替，用newchar字符替代与pattern匹配的字符所有.
```



```python
# title(): 转化为大写，例子：
# 'Hello world'.title() # 'Hello World'
```



```python
# print(re.sub(r"\s|_|", "", "He llo_worl\td"))
s = re.sub(r"(\s|_|-)+", " ",
           'some_database_field_name').title().replace(" ", "")  
#结果： SomeDatabaseFieldName
```



```python
# 可以看到此时的第一个字符为大写，需要转化为小写
s = s[0].lower()+s[1:]  # 最终结果
```

 

整理以上分析得到如下代码：

```python
import re
def camel(s):
    s = re.sub(r"(\s|_|-)+", " ", s).title().replace(" ", "")
    return s[0].lower() + s[1:]

# 批量转化
def batch_camel(slist):
    return [camel(s) for s in slist]
```

测试结果：

```python
s = batch_camel(['student_id', 'student\tname', 'student-add'])
print(s)
# 结果
['studentId', 'studentName', 'studentAdd']
```
