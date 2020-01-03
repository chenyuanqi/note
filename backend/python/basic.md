
### 什么是 Python
Python 发音 /ˈpaɪθən/ 拍森。  
Python 是解释型的编程语言。  
Python 是由吉多・范罗苏姆 (Guido Van Rossum) 在 90 年代早期设计。  

Python 命名的由来颇具感性色彩，1989 年圣诞节期间，在阿姆斯特丹，Guido 为了打发圣诞节的无趣，决心开发一个新的脚本解释程序，作为 ABC 语言的一种继承。之所以选中 Python（意为大蟒蛇）作为该编程语言的名字，是因为他是一个叫 Monty Python 的喜剧团体的爱好者。  
Python 诞生至今已经成为最受欢迎的程序设计语言之一。自 2004 年以后，Python 的使用率一直呈线性增长。

以下是 Python 的版本发布时间轴。  

- Python 1.0 - 1994 年 1 月  
Python 1.2 - 1995 年 4 月 10 号  
Python 1.3 - 1995 年 10 月 12 号  
Python 1.4 - 1996 年 10 月 25 号  
Python 1.5 - 1997 年 12 月 31 号  
Python 1.6 - 2000 年 09 月 05 号  
- Python 2.0 - 2000 年 10 月 16 号  
Python 2.1 - 2001 年 4 月 17 号  
Python 2.2 - 2001 年 12 月 21 号  
Python 2.3 - 2003 年 7 月 29 号  
Python 2.4 - 2004 年 11 月 30 号  
Python 2.5 - 2006 年 12 月 19 号  
Python 2.6 - 2008 年 10 月 1 号  
Python 2.7 - 2010 年 7 月 3 号  
- Python 3.0 - 2008 年 12 月 3 号  
Python 3.1 - 2009 年 6 月 27 号  
Python 3.2 - 2011 年 2 月 20 号  
Python 3.3 - 2012 年 9 月 29 号  
Python 3.4 - 2014 年 5 月 16 号  
Python 3.5 - 2015 年 9 月 13 号  
Python 3.6 - 2016 年 12 月 23 号  
Python 3.7 - 2018 年 6 月 28 日  
Python 3.8 - 2019 年 10 月 14 日  
Python 3.9 - ...  

### 为什么 Python
Python 易于使用，却也是一门完整的编程语言；  
Python 语法简洁且优美，几乎就是可执行的伪代码；  
Python 程序可以是独立的多个模块；  
Python 是可扩展的；  
Python 是可嵌入的；  
...

### Python 安装使用
Windows 和 Linux 系统可以在[这里](https://www.python.org/ftp/python/)下载相关的安装包。  
`注意：Windows 配置环境变量：【右键计算机】--》【属性】--》【高级系统设置】--》【高级】--》【环境变量】--》【在内容框中找到 Path 行，双击】 --> 【Python 安装目录追加到变值值中，用 ; 分割】（默认安装时会自动添加环境变量）`

```bash
# centos install
yum install -y make gcc gcc-c++
wget https://www.python.org/ftp/python/3.9.0/Python-3.9.0a2.tgz Python3
tar -xzvf Python3
cd Python3
./configure --prefix=/usr/local/python3 --enable-optimizations
make all
make install
# 如果需要修改默认 Python 版本
# mv /usr/bin/python /usr/bin/python2.7
# ln -s /usr/local/bin/python3 /usr/bin/python
# 修改配置文件 /usr/bin/yum 和 /usr/libexec/urlgrabber-ext-down 的 #!/usr/bin/python 改为 #!/usr/bin/python2.7

# mac install
brew install python3

# pip 升级
python -m pip install --upgrade pip

# 进入交互模式
python -i

# 交互模式中，最近一个表达式的值赋给变量 _
price = 100
price + 100
_ - 50 # 150
```

### Python3 基础
Python 之父两年前就已宣布 Python 2.7 将于 2020 年 1 月 1 日终止支持。

```python
# 这是一个注释

"""
    这是多行注释之一行
    这是多行注释之一行
"""

# print 是内置的打印函数
print("I'm Python,Nice to meet you!")

17 / 3 # 这是除法运算，5.666666666666667
17 // 3 # 这是板整除（即丢弃分数部分），5
17 % 3 # 这是取余，2
5 ** 2 # 这是计算幂乘方，5 的 2 次方， 25
4 * 3.75 - 1 # 这是组合运算，并对浮点数处理，14.0
(1 + 3) * 2  # 用括号决定优先级，8

# 布尔值 True 和 False
not True # 用 not 取非
True and False # 逻辑运算符，注意 and 和 or 都是小写
0 == False # True
2 != 1 # True
1 > 10 # False
2 <= 2 # True
1 < 2 < 3 # True
# 位运算
a = 0001 0101
b = 0000 0110
a&b = 0000 0100
a|b = 0001 0111
a^b = 0001 0011
~a = 1110 1010
a<<2 = 01010100
a>>2 = 0000 0101
# 成员运算
temp = "ABCDEFG"
a = "CDE"
a in temp # True
# 身份运算
a = 1
a is True # False
a is not True # True

# None 是一个对象
None
"etc" is None # False

# None，0，空字符串，空列表，空元组，空字典都算是 False
# 其他都为 True
bool(0) # False
bool("") # False
bool([]) # False
bool(()) # False
bool({}) # False

# 变量，在使用前必须先进行定义 (赋值)
width = 29
# 注意，赋值不能在表达式内部进行，也避免了这种错误：把 == 不小心写成了 =

'eggs'  # 这是使用单引号的字符串
"doesn't"  # 这是使用双引号的字符串，也可以使用转义符表示 'doesn\'t'
'First line.\nSecond line.'  # \n 意味着换行
r'/some/thing' # 这是 raw 原字符输出

# 多行字符串的定义，行尾使用 \ 表示避免行尾换行符被自动包含到字符串中（多行也可以使用 """...""" 表示）
multiple = '''\
	Usage: thing [OPTIONS]
	     -h                        Display this usage message
	     -H hostname               Hostname to connect to
'''

# 用 .format 来格式化字符串
"{} can be {}".format("strings", "interpolated")
# 重复参数
"{0} be nimble, {0} be quick, {0} jump over the {1}".format("Jack", "candle stick") # Jack be nimble, Jack be quick, Jack jump over the candle stick
# 关键字参数
"{name} wants to eat {food}".format(name="Bob", food="lasagna")  # Bob wants to eat lasagna

# 这里的 * 表示重复
3 * 'un' + 'ium' # unununium

'Py' 'thon' # 相邻的两个及其以上的字符串文本会自动合并在一起，Python
# 注意：变量和字符串文本无法合并，比如 prefix = 'Py';prefix 'thon' 会报错，此时应该使用 + 拼接字符串 prefix + 'thon'
# 利用相邻合并的特性可以切分大块的文本
text = ('Put several strings within parentheses '
        'to have them joined together.')

word = 'Python'
word[0]  # 位置 0 的字符，P
word[-1]  # 最后一个字符，n
# 切片
word[0:2]  # 从 0 （包含） 到 2（不包含），Py
word[:3]   # 从开始到 2 （不包含）的字符串，Pyt
# 获取字符串长度
len(word) # 6
```

Python 有几个复合数据类型，用于将一些值组合在一起。  

**列表（List）**  
列表可以写为中括号之间的一列逗号分隔的值，而且列表的元素不必是同一类型。  
列表和字符串有许多共同点，如可以用索引来访问、切割的操作等，它们都属于序列类型的公共特性。
```python
squares = [1, 4, 9, 16, 25]
squares[0]  # 通过索引返回元素， 1
squares[-1] # 25
squares[-3:]  # 切片返回新列表，[9, 16, 25]
squares[:] # 浅拷贝副本
squares + [36, 49, 64, 81, 100] # 列表的连接（列表是可变的），[1, 4, 9, 16, 25, 36, 49, 64, 81, 100]
# 列表修改元素
cubes = [1, 8, 27, 65, 125]
cubes[3] = 64 # [1, 8, 27, 64, 125]
cubes.append(216)  # 添加 6 的 立方
cubes.append(7 ** 3)  # 添加 7 的立方
# 使用切片批量修改列表元素
letters = ['a', 'b', 'c', 'd', 'e', 'f', 'g']
letters[2:5] = ['C', 'D', 'E'] # ['a', 'b', 'C', 'D', 'E', 'f', 'g']
letters[2:5] = [] # 移除元素，['a', 'b', 'f', 'g']
letters[:] = [] # 清空整个列表

letters = ['a', 'b', 'c', 'd']
len(letters) # 取列表长度，4

# 列表的嵌套
a = ['a', 'b', 'c']
n = [1, 2, 3]
x = [a, n] # [['a', 'b', 'c'], [1, 2, 3]]
x[0][1] # b
# 嵌套列表的矩阵处理
matrix = [
    [1, 2, 3, 4],
    [5, 6, 7, 8],
    [9, 10, 11, 12],
]
# 转置矩阵的行和列
[[row[i] for row in matrix] for i in range(4)] # [[1, 5, 9], [2, 6, 10], [3, 7, 11], [4, 8, 12]]
# 用 zip 处理复杂问题
list(zip(*matrix)) # [(1, 5, 9), (2, 6, 10), (3, 7, 11), (4, 8, 12)]

fruits = ['orange', 'apple', 'pear', 'banana', 'kiwi', 'apple', 'banana']
fruits.count('apple') # 统计 apple 出现的次数，2
fruits.index('banana') # 元素索引位置，格式是 list.index(*x*[, *start*[, *end*]])，3
fruits.index('banana', 4)  # 从索引 4 开始找 banana，6
fruits.reverse() # 逆序排列，['banana', 'apple', 'kiwi', 'banana', 'pear', 'apple', 'orange']
fruits.append('grape') # 追加元素，['banana', 'apple', 'kiwi', 'banana', 'pear', 'apple', 'orange', 'grape']
fruits.sort() # 排序，格式是 list.sort(*key=None*, *reverse=False*)，['apple', 'apple', 'banana', 'banana', 'grape', 'kiwi', 'orange', 'pear']
fruits.pop() # 移除并返回列表中给定位置的元素（默认移除最后一个）
fruits_copy = fruits.copy() # 浅拷贝，相当于 fruits[:]
fruits.clear() # 移除所有元素，相当于 del fruits[:]
fruits_copy.insert(1, 'pear') # 插入元素，格式 list.insert(*i*, *x*)，a.insert(0, x) 将元素插入列表最前面，a.insert(len(a), x) 相当于 a.append(x)，['apple', 'pear', 'apple', 'banana', 'banana', 'grape', 'kiwi', 'orange']
fruits_copy.extend(['x', 'xx']) # 将一个 iterable 的对象中的所有元素添加到列表末端来拓展这个列表，相当于 a[len(a):] =iterable，['apple', 'pear', 'apple', 'banana', 'banana', 'grape', 'kiwi', 'orange', 'x', 'xx']
fruits_copy.remove('apple') # 移除列表中第一个值为 apple 的元素（如不存在抛出 ValueError），['pear', 'apple', 'banana', 'banana', 'grape', 'kiwi', 'orange', 'x', 'xx']

vec = [-4, -2, 0, 2, 4]
# 创建一个新列表，将原列表中的每个元素乘以 2
[x*2 for x in vec] # [-8, -4, 0, 4, 8]
# 去除原列表中的负数
[x for x in vec if x >= 0] # [0, 2, 4]
# 对原列表中的每个元素调用函数
[abs(x) for x in vec] # [4, 2, 0, 2, 4]
freshfruit = ['  banana', '  loganberry ', 'passion fruit  ']
[weapon.strip() for weapon in freshfruit] # ['banana', 'loganberry', 'passion fruit']
# 创建一个由二元组构成的列表，元素形如 (number, square)
[(x, x**2) for x in range(6)] # [(0, 0), (1, 1), (2, 4), (3, 9), (4, 16), (5, 25)]
# 用一个含有两个 `for` 的列表初始化表达式将一个多维列表降维
vec = [[1,2,3], [4,5,6], [7,8,9]]
[num for elem in vec for num in elem] # [1, 2, 3, 4, 5, 6, 7, 8, 9]
# 列表初始化表达式可以包含复杂语句和嵌套的函数调用
from math import pi
[str(round(pi, i)) for i in range(1, 6)] # ['3.1', '3.14', '3.142', '3.1416', '3.14159']

a = [-1, 1, 66.25, 333, 333, 1234.5]
del a[0] # [1, 66.25, 333, 333, 1234.5]
del a[2:4] # [1, 66.25, 1234.5]
del a[:] # 清空但保留 a 的定义
del a # 清空及去除 a 的定义
```

**元组（Tuple）**  
元组由一系列被逗号分隔开的值组成。  
元组在输入时并不一定要带上两侧的括号（尽管有时带上括号非常有必要），元组在输出时总是带有两侧的括号。
```python
t = 12345, 54321, 'hello!'
t[0] # 12345
u = t, (1, 2, 3, 4, 5) # 嵌套元组，((12345, 54321, 'hello!'), (1, 2, 3, 4, 5))
# 不能对元组中的项进行赋值，但在创建元组时，可以传入可修改的对象，如列表
t[0] = 444 # 元组不可被修改，抛出 TypeError 异常
v = ([1, 2, 3], [3, 2, 1]) # 却可以包含可以被修改的对象

# 创建一个空的或只有一个元素的元组
empty = ()
singleton = 'hello',    # 注意后面的逗号不能省略
len(empty) # 0
len(singleton) # 1

# 元组解包
# 也称序列解包（所有右值的序列都可以使用这种语法），解包要求在等号的左边有着和序列内元素数量相同的变量
t = 12345, 54321, 'hello!'
x, y, z = t
print(x, y, z) # 12345 54321 'hello!'
```

列表和元组的区别：  

- 元组和列表都是序列类型，可以存放任何类型的数据、支持切片、迭代等操作。  
- 元组是不可修改的，并且经常包含着不同类型的元素（异构数据 (heterogeneous)数据），总是通过解包（本章的后面会介绍）或索引（命名元组可以通过属性索引的方式）来访问；  
- 列表是可修改的，并且往往包含着同样类型的元素（同构数据 (homogenous)），通过遍历的方式来进行访问。  
- Python 的内部实现对元组做了大量优化，访问速度比列表更快；元组在内部实现上不允许修改其元素值，从而使得代码更加安全，例如调用函数时使用元组传递参数可以防止在函数中修改元组，而使用列表则很难做到这一点。  
- 作为不可变序列，与整数、字符串一样，元组可用作字典的键，也可以作为集合的元素，而列表则永远都不能当做字典键使用，也不能作为集合中的元素，因为列表不是不可变的，或者说不可哈希。  

因为 tuple 作为没有名字的记录来使用在某些场景有一定的局限性，所以又有了一个 namedtuple 类型的存在，namedtuple 可以指定字段名，用来当做一种轻量级的类来使用。

**集合（Set）**  
集合是由多个无重复元素构成的无序整体，支持的基本功能包括成员检查以及重复元素的去除，同时支持求并集、交集、差集以及对称差集等操作。  

```python
# 集合可以通过 {} 或者调用 set() 函数创建
# 注意：如果需要创建一个空的集合实例，需使用 set() 而非 {} ，因为后者会创建一个空的字典实例
basket = set()
basket = {'apple', 'orange', 'apple', 'pear', 'orange', 'banana'}
'orange' in basket # 快速成员检查，True

a = set('abracadabra') # {'a', 'r', 'b', 'c', 'd'}
b = set('alacazam') # {'z', 'a', 'l', 'm', 'c'}
a - b # 在 a 中但是不在 b 中的字母，{'r', 'd', 'b'}
a | b # 在 a 中或在 b 中的字母，{'a', 'c', 'r', 'd', 'b', 'm', 'z', 'l'}
a & b # a 和 b 共有的字母，{'a', 'c'}
a ^ b # 在 a 中或在 b 中但两者不共有的字母，{'r', 'd', 'b', 'm', 'z', 'l'}

# 递推式构造集合
{x for x in 'abracadabra' if x not in 'abc'} # {'r', 'd'}
```

**字典（Dict）**  
字典这种数据类型与序列类数据结构不同的是通过 键 来进行索引， 键 可以是任何不可变类型，而序列类的数据类型通常使用数字进行索引。  
`注意：如果元组内只含有字符串，数字或者元组，那这个元组也可以作为键，但是如果元组内直接或间接的含有任何的可变类型的数据，则不可以作为键。一般不能使用列表来作为键。`  
可以把字典当成 键：值 对的集合类理解，字典要求一个字典内的键是不能重复的。  

字典主要的操作符就是通过键来存储对应的数据，以及根据键来取出对应的数据，也可以通过 del 来删除一个键值对。如果在存储数据的时候使用了字典中已有的键，则该键对应的值会被更新为当前新赋给值；如果使用字典中不存在的键来获取值，则会产生 error ，提示不存在这样的键。  

```python
# 创建一个空字典
dict_var = {}

# 通过一系列的 键值对 产生一个字典
dict([('sape', 4139), ('guido', 4127), ('jack', 4098)]) # {'sape': 4139, 'guido': 4127, 'jack': 4098}
# 从任意的键和值的表达式来创建字典
{x: x**2 for x in (2, 4, 6)} # {2: 4, 4: 16, 6: 36}
# 当键是字符串的时候，使用参数赋值的方式来指定键值对更方便
dict(sape=4139, guido=4127, jack=4098) # {'sape': 4139, 'guido': 4127, 'jack': 4098}

tel = {'jack': 4098, 'sape': 4139}
tel['guido'] = 4127 # {'jack': 4098, 'sape': 4139, 'guido': 4127}
del tel['sape'] # {'jack': 4098, 'guido': 4127}
tel['irv'] = 4127 # {'jack': 4098, 'guido': 4127, 'irv': 4127}
list(tel) # 字典中所有键组成的列表，['jack', 'guido', 'irv']
sorted(tel) # 经过排序的键的列表，['guido', 'irv', 'jack']
'guido' in tel # True
'jack' not in tel # False

# 遍历字典
knights = {'gallahad': 'the pure', 'robin': 'the brave'}
for k, v in knights.items():
	print(k, v)
# 遍历一个序列时，位置索引和对应的值可以用 enumerate() 方法一次性全部得到
for i, v in enumerate(['tic', 'tac', 'toe']):
	print(i, v)
# 当需要同时遍历两个或多个序列时，可以使用 zip() 方法将他们合并在一起
questions = ['name', 'quest', 'favorite color']
answers = ['lancelot', 'the holy grail', 'blue']
for q, a in zip(questions, answers):
	print('What is your {0}?  It is {1}.'.format(q, a))
# 当需要反过来遍历一个序列的时候，使用 reversed() 方法来将一个正的序列倒序
for i in reversed(range(1, 10, 2)):
	print(i)
# 需要按顺序遍历一个序列，可以把未排序的序列传到 sorted() 方法中来获得一个新的排好序的列表
basket = ['apple', 'orange', 'apple', 'pear', 'orange', 'banana']
for f in sorted(set(basket)):
	print(f)
# 有时需要在遍历的过程中修改列表，但这种时候创建一个新的列表会更方便也更安全
import math
raw_data = [56.2, float('NaN'), 51.7, 55.3, 52.5, float('NaN'), 47.8]
filtered_data = []
for value in raw_data:
	if not math.isnan(value):
		filtered_data.append(value)
```

**序列及其他类型的比较**  
拥有相同序列类型的序列对象之间可以进行比较。  
序列间的比较基于字典排序：首先比较两序列的首项，如果它们不同，那么比较就有了结果；如果它们相同，接下来的两项将继续进行比较，以此类推，直到两者中任何一个序列被遍历完毕。如果比较的项所在的序列是同样的类型，那么可以按照字典排序的方法递归进行下去。如果两序列所有的项比较过后都是相同的，则认为这两个序列相等。如果其中一个序列是另一个序列从头开始的一个子序列，那么更短的一个被认为更小。字符串的字典排序对于单个字符按照 Unicode 的编码大小进行排序。   
需要注意的是，如果有适当的比较方法，对于不同类型对象间的比较使用 < 或者 > 也是合法的。例如，混合数字类型可以根据它们的数值大小进行比较，如 0 等于 0.0 ，以此类推。否则，Python 解释器会抛出一个 TypeError  的异常，而非给出一个随机的排序。

```python
(1, 2, 3)              < (1, 2, 4) # True
[1, 2, 3]              < [1, 2, 4] # True
'ABC' < 'C' < 'Pascal' < 'Python' # True
(1, 2, 3, 4)           < (1, 2, 4) # True
(1, 2)                 < (1, 2, -1) # True
(1, 2, 3)             == (1.0, 2.0, 3.0) # True
(1, 2, ('aa', 'ab'))   < (1, 2, ('abc', 'a'), 4) # True
```

### Python 编程

**流程控制**  
注意，这里的 elseif 是 elif。
```python
if x < 0:
    x = 0
    print('Negative changed to zero')
elif x == 0:
    print('Zero')
elif x == 1:
    print('Single')
else:
    print('More')
```

注意：
被用在 while 和 if 语句中的判断条件不仅仅可以包含比较运算，还可以包含任何的运算符。  
比较运算符 in 和 not in 能够检查某个值是否在一个序列里出现（或不出现）；比较运算符 is 和 is not 比较两个对象是否是同一个对象；这只会影响如列表之类可修改的对象。所有的比较运算符的优先级都相同，比所有的算术运算法的优先级都要低。  
比较运算符可以采用连写的方式，例如 a < b == c 用来检查是否 a 小于 b 并且 b 等于 c 。  
比较运算符可以用布尔运算符 and 和 or 进行组合，然后他们的结果（或者任何其他的布尔表达式）可以被 not 否定。布尔运算符的优先级又比比较运算符更低；在他们之间， not 的优先级最高，而 or 的优先级最低，因此 A and not B or C 就等价于 (A and (not B)) or C。当然，括号可以用来提升优先级。  
布尔运算符 and 和 or 往往被称为“短路运算符”：它们的参数从左往右一个个被计算，而当最终结果已经能够确定时，就不再计算剩余的参数了。举个例子，如果 A 和 C 是真的，而 B 是假的，那么 A and B and C 不会计算表达式 C 的值。`值得注意的是，当“短路运算符”不作为布尔值使用，而是作为普通的值来使用时，短路运算符的返回值将会是最后一个被计算的参数`。  
```python
string1, string2, string3 = '', 'Trondheim', 'Hammer Dance'
non_null = string1 or string2 or string3 # Trondheim
```

**循环**  
```python
words = ['cat', 'window', 'defenestrate']
for w in words:
	print(w, len(w))

# 迭代一系列的数字
for i in range(5):
	print(i)
# 扩展
range(5, 10) # 5, 6, 7, 8, 9
range(0, 10, 3) # 0, 3, 6, 9
range(-10, -100, -30) # -10, -40, -70

# 使用 break 跳出循环
for n in range(2, 10):
	for x in range(2, n):
		if n % x == 0:
			print(n, 'equals', x, '*', n//x)
			break
		else:
			# 没有找到一个因数导致的循环失败
			print(n, 'is a prime number')

# 使用 continue 继续循环的下一次迭代
for num in range(2, 10):
	if num % 2 == 0:
		print("Found an even number", num)
		continue

	print("Found a number", num)

while True:
	pass # 什么也不做，可用于语法上需要但程序不需要做什么的时候
```

**函数**  
```python
# 打印斐波那契数列
def fib(n):
	a, b = 0, 1

	while a < n:
		print(a, end=' ')
		a, b = b, a+b

	print()

# 调用函数
fib(2000) # 0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597

# 函数定义之默认参数
def ask(prompt, retries=4, reminder='Please try again!'):
	pass
# 函数定义之关键字参数
def parrot(voltage, state='a stiff', action='voom', type='Norwegian Blue'):
	pass
parrot(1000)                                          # 一个位置参数
parrot(voltage=1000)                                  # 一个关键字参数
parrot(voltage=1000000, action='VOOOOOM')             # 2个关键字参数
parrot(action='VOOOOOM', voltage=1000000)             # 2个关键字参数
parrot('a million', 'bereft of life', 'jump')         # 3个位置参数
parrot('a thousand', state='pushing up the daisies')  # 一个位置参数，一个关键字参数
# 函数定义之参数混合
# 注意 *arguments 必须在 **keywords 前面
def cheeseshop(kind, *arguments, **keywords):
	pass
cheeseshop("Limburger", "It's very runny, sir.",
           "It's really very, VERY runny, sir.",
           shopkeeper="Michael Palin",
           client="John Cleese",
           sketch="Cheese Shop Sketch")

# 函数定义之可变参数
def concat(*args, sep="/"):
	return sep.join(args)
concat("earth", "mars", "venus") # earth/mars/venus
concat("earth", "mars", "venus", sep=".") # earth.mars.venus

# 分离参数列表
args = [3, 6]
list(range(*args)) # * 符号是指从列表（或元组）中提取参数来调用函数，[3, 4, 5]
def parrot(voltage, state='a stiff', action='voom'):
	pass
d = {"voltage": "four million", "state": "bleedin' demised", "action": "VOOM"}
parrot(**d) # ** 符号是指将关键字参数从字典中提取出来

# 使用 lambda关键字来创建匿名函数
def make_incrementor(n):
	return lambda x: x + n
f = make_incrementor(42)
f(0) # 42
f(3) # 45
# lambda 用于传递参数
pairs = [(1, 'one'), (2, 'two'), (3, 'three'), (4, 'four')]
pairs.sort(key=lambda pair: pair[1]) # [(4, 'four'), (1, 'one'), (3, 'three'), (2, 'two')]

# 文档字符串
# 第一行应该始终是一个对对象目的的精简的总结
# 如果文档字符串不止一行，则第二行应为空白，从而能在视觉上将总结与其余的描述分开
# Python 解析器并不会删除多行字符串文字的缩进，因此处理文档的工具必须在有必要之时删除缩进
def my_function():
	"""只要写文档，其他啥都别做

	...
	确实，它也啥都不做
	"""
	pass
print(my_function.__doc__)

# 函数注解 (Function annotations) 应用于用户自定义的函数，可使用的类型是完全可选的元数据
# 注解（Annotations）是以字典的形式存放在函数的  __annotations__  属性中，并且不对函数有任何影响
# 参数注解 (Parameter annotations) 是由参数名称后面加上冒号来定义的，后面紧跟一个描述，来表示注解的值
# 返回注解 (Return annotations) 的定义方式是：由 -> 符号开始，在参数列表和表示函数 def结束的冒号之间，加上你的描述
def f(ham: str, eggs: str = 'eggs') -> str:
	print("Annotations:", f.__annotations__)
	print("Arguments:", ham, eggs)

	return ham + ' and ' + eggs

f('spam')
# Annotations: {'ham': <class 'str'>, 'return': <class 'str'>, 'eggs': <class 'str'>}
# Arguments: spam eggs
# 'spam and eggs'
```

### Python 代码风格
[PEP 8](https://www.python.org/dev/peps/pep-0008/) 是大多数 Python 项目使用的代码风格指南，它提供了高可读性和养眼的代码风格，以下划重点。  

- 缩减使用四个空格而不是制表符  
- 每行不要超过 80 个字符  
- 使用空行分隔函数、类或者函数内较大的代码段  
- 尽量将注释和代码放在一起  
- 用 docstrings  
- 用在操作符前后和逗号之后加空格，但是括号之内不需要： a= f(1, 2) + g(3, 4)   
- 一致性的命名类与函数：惯例是用 CamelCase 命名类，用 lower_case_with_underscores 命名函数和方法  
-  Python 默认使用 UTF-8，甚至纯 ASCII 在任何情况下都能最好地工作  
