
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

Python 也有自己的缺点，比如：  
速度慢，Python 的运行速度相比 C 语言确实慢很多，跟 JAVA 相比也要慢一些  
代码不能加密  
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
# mac upgrade
brew update
brew upgrade Python3

# pip 升级
python -m pip install --upgrade pip

# 进入交互模式
python -i

# 交互模式中，最近一个表达式的值赋给变量 _
price = 100
price + 100
_ - 50 # 150
```

[python环境管理](https://github.com/pyenv/pyenv)  

### Python 工具
**Jupyter Notebook**  
[Jupyter Notebook](https://jupyter.org/index.html) 是现代 Python 的必学技术。   
Jupyter Notebook 是什么呢？安装 Jupyter 的创始人 Fernando Perez 的说法，他最初的梦想是做一个综合 Ju（Julia）、Py（Python）和 R 三种科学运算语言的计算工具平台，所以将其命名为 Ju-Py-te-R。  
发展到现状，Jupyter 已经成为一个几乎支持所有语言，能把软件代码、计算输出、解析文档、多媒体资源整合在一起的多功能科学运算平台。

Jupyter 的优点如下：

- 整合所有资源  
- 交互性编程体验
- 零成本重现结果

```bash
# Jupyter Notebook 安装
pip install jupyterlab notebook

# 在相关文件夹下运行起来
jupyter notebook
# 浏览器打开 http://localhost:8888/ 即可查看

# 安装主题（https://github.com/dunovank/jupyter-themes）
pip install jupyterthemes
# 查看主题列表
jt -l
# 设置主题
jt -t oceans16
# 设置主题，带参数（-t是设置主题，-f设置代码的字体，-nf设置notebook的字体）
jt -t monokai -f roboto -nf robotosans -tf robotosans -N -T -cellw 70% -dfs 10 -ofs 10
# 恢复默认主题
jt -r

# Markdown 文件自动生成目录&自动补全代码
python -m pip install jupyter_contrib_nbextensions
jupyter contrib nbextension install --user --skip-running-check
# 安装完成后，在 Nbextensions 选项卡里勾选 Table of Contents 以及 Hinterland
```

**Virtualenv**  
virtualenv 是用来为一个应用创建一套 “隔离” 的 Python 运行环境，解决了不同应用间多版本的冲突问题。  
[参考文档](https://pythonguidecn.readthedocs.io/zh/latest/dev/virtualenvs.html#virtualenv)   

```bash
# centos\linux can install virtualenvwrapper
pip install virtualenv
# 查看版本
virtualenv --version

# 创建项目
mkdir myproject_path
cd myproject_path
# 创建一个独立的 Python 运行环境，命名为 venv（--no-site-packages 是系统 Python 环境中的所有第三方包都不会复制过来）
# 参数 -p 指定 python 版本，值有 python2、python3
virtualenv --no-site-packages venv
# 用 source 进入该环境（win: source venv/Scripts/activate）
source venv/bin/activate
# 退出当前的 venv 环境
deactivate 

# 为了保持您的环境的一致性，“冷冻住（freeze）” 环境包当前的状态是个好主意
pip freeze > requirements.txt # 创建一个 requirements.txt 文件，其中包含了当前环境中所有包及 各自的版本的简单列表
# 可以使用 pip list 在不产生 requirements 文件的情况下， 查看已安装包的列表
pip list
# 以后安装相同版本的相同包变得容易
pip install -r requirements.txt
```

### Python3 基础
Python 之父两年前就已宣布 Python 2.7 将于 2020 年 1 月 1 日终止支持。
[Python2.7 与 Python3.x 的主要差异](http://chenqx.github.io/2014/11/10/Key-differences-between-Python-2-7-x-and-Python-3-x/)  

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
# 取十进制小数点后的精度为 3 ，得到的浮点数为 '0.333'
print('{0:.3f}'.format(1.0/3)) # 0.333
# 填充下划线 (_) ，文本居中，将 '___hello___' 的宽度扩充为 11 
print('{0:_^11}'.format('hello')) # ___hello___

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
# 注意，Python 中的字符串是不可变的，不能随意更改字符串中字符的值（+= 拼接操作例外）
```

**浅拷贝和深拷贝**  
赋值（=），就是创建了对象的一个新的引用，修改其中任意一个变量都会影响到另一个。  
浅拷贝 copy.copy：创建一个新的对象，但它包含的是对原始对象中包含项的引用（如果用引用的方式修改其中一个对象，另外一个也会修改改变）。  
深拷贝：创建一个新的对象，并且递归的复制它所包含的对象（修改其中一个，另外一个不会改变）{copy 模块的 deep.deepcopy () 函数}


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
集合的本质是哈希表。  
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

filled_dict = {"one": 1, "two": 2, "three": 3}
filled_dict.keys() # 获取所有键，dict_keys(['one', 'two', 'three'])
filled_dict.values() # 获取所有值，dict_values([3, 2, 1])
filled_dict.get("one") # 用 get 来避免 KeyError，1
filled_dict.setdefault("four", 3) # setdefault 方法只有当键不存在的时候插入新值
filled_dict.update({"four": 4}) # 字典赋值，或filled_dict["four"] = 4

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
else:
	pass # for 结束继续其他操作
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
else：
	pass # while 结束后继续进行其它你想做的操作
```

**输入和输出**  
程序的输出可以有多种形式：可以将数据以人类可读的形式打印到屏幕上，或者将其写入到文件中以供后续使用。  

在 Python 中存在两种输出值的方法：表达式语句以及 print() 函数。（第三种方法是使用文件对象的 write() 方法；标准文件输出可以参考 sys.stdout 方法）  
如果仅仅想要在调试时打印某些变量，而不进行格式化输出，那么可以使用 repr() 函数或者 str() 函数将任意值转化成字符串。str() 函数能够将值以人类可读的形式呈现，而 repr() 函数则是将值以解释器可读的形式呈现（如果没有与之对应的转换语法，则会产生 SyntaxError 异常）。若某个对象没有适用于人类可读的形式，则 str() 函数的返回值与 repr() 函数相同。诸如数值、或者是链表、字典这样的结构，上述两种函数各自有统一的呈现方式。但是对于字符串，上述两种函数各自有独特的呈现方式。  

Python 提供了被称为 JSON (JavaScript Object Notation) 的流行数据交换格式。  
序列化（dumps）：json 采用 Python 式的数据层级，并且转换成字符串的形式；  
反序列化：从字符串转化成数据类型。  

```python
# f-strings 格式化字符串
year = 2020
event = 'Referendum'
f'Results of the {year} {event}' # Results of the 2016 Referendum
# {} 中 ： 后 .xf 表示四舍五入到 x 位，如果 : 后是整数则表示最小字符数
# {} 中 !a 表示 ascii()，!s 表示 str()，!r 表示 repr()
print(f'The value of pi is approximately {math.pi:.3f}.') # The value of pi is approximately 3.142

# str.format() 格式化字符串
yes_votes = 42_572_654 
no_votes = 43_132_495
percentage = yes_votes/(yes_votes+no_votes)
'{:-9} YES votes  {:2.2%}'.format(yes_votes, percentage) #  42572654 YES votes  49.67%

# 字符串转化
s = 'Hello, world.'
str(s) # Hello, world.
repr(s) # 'Hello, world.'
str(1/7) # 0.14285714285714285
x = 10 * 3.25
y = 200 * 200
s = 'The value of x is ' + repr(x) + ', and y is ' + repr(y) # The value of x is 32.5, and y is 40000
repr((x, y, ('spam', 'eggs'))) # (32.5, 40000, ('spam', 'eggs'))

# 读写文件
# mode 参数可以使用 'r' 表示只读模式；'w' 表示只写入模式；'a' 表示在文件末尾追加写入；'r+' 表示读写操作
# mode 参数中 'b' 表示以 二进制模式 打开文件并且追加数据， 数据以字节的形式读写到文件中，但是这种模式应该被用来打开不包含文本的文件中
f = open('workfile', 'w')

# with 是非常适合用于处理文件对象
# with 的优势在于，即使发生了 exception，文件操作序列结束后也可以自动关闭
with open('workfile') as f:
	read_data = f.read()

# 使用 json 保存数据
import json
x = json.dumps([1, 'simple', 'list']) # '[1, "simple", "list"]'
# f 是文本文件已打开的可写入文件对象
json.dump(x, f)
x = json.load(f)
```

**错误和异常（error）**  
错误信息（至少）可以分为两类：语法错误（也称解析错误） 和 异常（语法正确却在运行的时候报错）。  

Python 允许编程处理特定的异常。

```python
while True print('Hello world') # SyntaxError: invalid syntax
10 * (1/0) # ZeroDivisionError: division by zero

# 用 try/except 块处理异常状况
try:
    # 用 raise 抛出异常
    raise IndexError("This is an index error")
except IndexError as e:
    pass    # pass是无操作，但是应该在这里处理错误
except (TypeError, NameError):
    pass    # 可以同时处理不同类的错误
else:   # else 语句是可选的，必须在所有的 except 之后
    print("All good!")   # 只有当 try 运行完没有错误的时候这句才会运行
finally: # finally 语句是可选的，用于定义所有情况下都之行的操作
    # 当异常产生在 try 子句中并未被 except 子句捕获（或异常在 except 或 else 子句中产生）时，异常将在 finally 子句被执行后再引发
    # 当其他子句通过 break, continue or return 等语句离开 try 语句时，finally 子句也会被执行
    pass
```

用户还可以自定义异常。  
不要滥用异常。通常情况下，使用系统自带的异常就可以了。异常处理，通常用在你不确定某段代码能否成功执行，也无法轻易判断的情况下，比如数 据库的连接、读取等等。正常的 flow-control 逻辑，不要使用异常处理，直接用条件语句解决就可以了。  
```python
class MyInputError(Exception):
    """Exception raised when there're errors in input"""
    def __init__(self, value): # 自定义异常类型的初始化 
        self.value = value
    def __str__(self): # 自定义异常类型的 string 表达形式 
        return ("{} is invalid input".format(repr(self.value)))

try:
    raise MyInputError(1) # 抛出 MyInputError 这个异常 
except MyInputError as err:
    print('error: {}'.format(err)) 

# 输出 error: 1 is invalid input
```

**函数（def）**  
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

**类（class）**  
类就是一组数据和函数的集合。  
创建了一个类意味着创建了一个对象类型，同时允许创建此类型的实例；每个类实例都带有属性以维护其状态，同样也有方法（在它自己的类中定义）来修改这些状态。  

namespace（命名空间） 是一个从名字到对象的映射。  
大部分命名空间当前都由 Python 字典实现，但一般情况下基本不会去关注它们（除了要面对性能问题时）。  
命名空间的案例：存放内置函数的集合（里面含有 abs() 这样的函数，和其他的内置名称）；模块中的全局名称；函数调用中的本地名称。  

scope (作用域) 是一段 Python 程序的文本区域，处于其中的命名空间是可直接访问的。  
如果某名称是在全局进行的声明，那么所有的引用和分配都会直接导向中间的这层包含模块的全局名称的作用域中。  
要想让最内层的作用域重新绑定一个在外层出现过的变量，我们可以用 nonlocal 声明来完成；如果不声明 nonlocal (非本地)，这些变量则都是只读的（任何尝试写入这种变量的行为都将会创建一个 全新 的本地变量，不会对最外层的那个有丝毫影响。）  
一般地，本地作用域引用着当前函数的本地名称；外层的函数引用的是和全局作用域一样的命名空间：模块的命名空间。类定义放置在本地作用域的另一个命名空间中。  

```python
def scope_test():
    def do_local():
        spam = "local spam"

    def do_nonlocal():
        nonlocal spam
        spam = "nonlocal spam"

    def do_global():
        global spam
        spam = "global spam"

    spam = "test spam"
    do_local()
    print("After local assignment:", spam)
    do_nonlocal()
    print("After nonlocal assignment:", spam)
    do_global()
    print("After global assignment:", spam)

scope_test()
# After local assignment: test spam
# After nonlocal assignment: nonlocal spam
# After global assignment: nonlocal spam
print("In global scope:", spam) # In global scope: global spam
# 注意：本地的分配并未改变 scope_test 中绑定的 spam，而 nonlocal 标明过的分配则改变了 scope_test 绑定的 spam，global 则更改的是模块层面的绑定（global 之前没有绑定 spam）  
```

类的定义与函数定义（def statements）差不多，在它们生效前需要预先执行这些定义（也可以在  if 分支或函数内部声明类）。  
类定义内的声明通常是函数定义，也有其他声明。在类中定义的函数通常有一个特有的参数列表，指代是作为方法调用的。类定义后，会创建一个新的命名空间作为本地作用域，从而所有的本地变量的指定都会进到这个新的作用域里。类定义正常结束时，一个新的类对象就被创建出来了，这是类定义在命名空间中最基本的一层包装。  

类对象支持两种操作：属性引用和实例化。  
属性引用（Attribute references ）使用的是 Python 中标准的属性引用语法： obj.name，每个类都有 \_\_doc\_\_ 属性（用于返回类中文档字符串）、\_\_self\_\_ 属性（返回类的实例对象）、\_\_func\_\_ 属性（和方法相对应的函数对象）等。  
实例化类似函数的形式。

类的继承。  
当一个类对象被创建，它会记录它的基类。这将被用于解析对象的属性：如果一个需要的属性不存在于当前类中，紧接着就会去基类中寻找；如果该基类也是从其他类派生出来的，那么相同的过程也会递归地被应用到这些类中。  
重写类的方法，一般用于扩展同名的基类方法，而非简单的替换。调用基类方法的方案：调用 BaseClassName.methodname(self, arguments) 即可。  
类的多重继承中，继承的属性是深度优先，从左到右，而不是在继承结构中重叠的同一个类中搜索两次。

Python 提供了两个判断继承关系的内建函数：

- isinstance() 检查一个实例的类型：当且仅当 obj.\_\_class\_\_ 是 int 或其它从 int 派生的类时， isinstance(obj, int) 才会返回 True 。
- issubclass() 检查类之间的继承关系：因为 bool 是 int 的一个子类，所以 issubclass(bool, int) 返回 True；然而，因为 float 不是 int 的派生类，所以 issubclass(float, int) 返回 False。

```python
class MyClass:
    """我的类"""
    i = 12345

    """类被实例化时自动调用 __init__() 方法"""
    def __init__(self):
        print(self.i)

    def f(self):
        return 'hello, class'

    # 类方法，被所有此类的实例共用，第一个参数是这个类对象
    @classmethod
    def class_func(cls):
    	pass

    # 静态方法，调用时没有实例或类的绑定
    @staticmethod
    def static_func():
    	pass

class MySon(MyClass):
	def f(self):
		# 这里调用父类，self 不可省略
		tmp = MyClass.f(self)

		return 'dad:' + tmp + ' son: hello, I am son.'

x = MySon()
print(x.f())
```

只能从对象内部访问的『私有』实例变量，在 Python 中不存在。  
然而，在大多数 Python 代码中存在一个这样的约定：以一个下划线开头的命名（例如 \_spam ）会被处理为 API 的非公开部分（无论它是一个函数、方法或数据成员）。  

**迭代器（iter）**  
大部分容器对象都能被 for 所循环，包括 list/set/tuple/str/dict 等数据结构以及生成器。  
for 声明会调用容器对象的 iter() 函数，这个函数则返回一个迭代器对象；迭代器对象有 \_\_next\_\_() 方法，它会让容器中的元素一次返回一个。 \_\_next\_\_()  会抛出 StopIteration 异常来让 for 结束，也可以用 next() 函数来调用它的 \_\_next\_\_() 方法。  
在类中，我们可以相应的迭代器行为。  

```python
class Reverse:
    def __init__(self, data):
        self.data = data
        self.index = len(data)

    def __iter__(self):
        return self

    def __next__(self):
        if self.index == 0:
            raise StopIteration
        self.index = self.index - 1

        return self.data[self.index]

rev = Reverse('spam')
for char in rev:
	print(char)
```

**生成器（Generator）**  
生成器是一个简单又强大的创建迭代器的工具。  
生成器的 \_\_iter\_\_() 和 \_\_next\_\_() 方法都被自动隐式的创建了。每次 next() 调用生成器时，生成器就会从它断开的地方恢复（本地变量和执行条件都会被自动保存）。  

生成器表达式用来一般用在在函数内需要写即用即删的数据的时候。  
生成器表达式比起完整的生成器要更加紧凑但并不如它功能强大，不过比起列表表达式来说内存占用更少。  

生成器的好处是不用占用很多内存，只需要在用的时候计算元素的值就行了。

```python
def reverse(data):
    for index in range(len(data)-1, -1, -1):
        yield data[index]

for char in reverse('golf'):
	print(char)

# 生成器表达式
sum(i*i for i in range(10)) # 285
xvec = [10, 20, 30]
yvec = [7, 5, 3]
sum(x*y for x,y in zip(xvec, yvec)) # 260
from math import pi, sin
sin_table = {x: sin(x*pi/180) for x in range(0, 91)}
```

**装饰器（Decorator）**  
装饰器就是在代码运行期间动态增加功能的方式。  

装饰器本质上是一个 Python 函数或类，它可以让其他函数或类在不需要做任何代码修改的前提下增加额外功能，装饰器的返回值也是一个函数 / 类对象。  
装饰器经常用于有切面需求的场景，比如：插入日志、性能测试、事务处理、缓存、权限校验等场景，装饰器是解决这类问题的绝佳设计。有了装饰器，我们就可以抽离出大量与函数功能本身无关的雷同代码到装饰器中并继续重用。概括的讲，装饰器的作用就是为已经存在的对象添加额外的功能。  

```python
from functools import wraps

def beg(target_function):
	# decorator 就是一个返回函数的高阶函数
    @wraps(target_function)
    def wrapper(*args, **kwargs):
        msg, say_please = target_function(*args, **kwargs)
        if say_please:
            return "{} {}".format(msg, "Please! I am poor :(")
        return msg

    return wrapper


@beg
def say(say_please=False):
    msg = "Can you buy me a beer?"
    return msg, say_please


print(say())  # Can you buy me a beer?
print(say(say_please=True))  # Can you buy me a beer? Please! I am poor :(
```

**模块（module）**  
Python 可以把定义放入一个文件中然后在一个脚本或交互式解释器实例中使用它，这个文件被叫做模块 （module）。  

模块中的定义可以通过导入进入到其他模块或者主模块（在顶层和计算器模式下执行的脚本中可以访问的变量集合）。  
一个模块是一个包含 Python 定义和声明的文件，文件是模块名加上 .py 后缀；在一个模块中，模块名（字符串类型）可以通过全局变量 \_\_name\_\_ 获取。  

```python
# fibo.py
# 斐波那契数模块

def fib(n):    # 打印斐波那契数直到 n
    a, b = 0, 1
    while a < n:
        print(a, end=' ')
        a, b = b, a+b
    print()

def fib2(n):   # 返回到 n 的斐波那契数
    result = []
    a, b = 0, 1
    while a < n:
        result.append(a)
        a, b = b, a+b
    return result

# 同目录下另一个文件 temp.py
import fibo

fibo.fib(1000) # 0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987
fibo.fib2(100) # [0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89]
print(fibo.__name__) # fibo
```

一个模块可以包含可执行声明包括函数定义，这些声明被用于初始化模块，它们只在模块被第一次导入时执行（作为脚本运行也会执行）。  

每个模块都有其私有的符号表，模块中定义的所有函数将这个符号表作为全局符号表。因此，一个模块的作者可以在模块中使用全局变量而无需担心与其他模块的全部变量冲突。另一方面，如果你知道你在干什么，你同样可以使用 模块.变量 的方式来获取一个模块的全局变量。  

模块可以导入其他模块。  
将所有 import 语句放在模块（或者脚本，如果这个问题重要的话）的开头不是必须的，但习惯如此，被导入的模块名被放置于当前模块的全局符号表中。

`注意：由于性能原因，每个模块在每个解释器会话中只会被导入一次。因此，如果你改变了你的模块，你必须重启解释器；或者你只想交互式地测试一个模块，你可以使用 importlib.reload()，如 import importlib; importlib.reload(modulename)。`

```python
# import 声明的一种变体可以把一个模块中的变量直接导入当前模块的符号表中
# 这样做不会把模块名引入本地符号表中
from fibo import fib, fib2
# 还有一种导入声明的变体可以导入一个模块中定义的所有变量： from fibo import *
# 但是 * 过于粗暴，它会为解释器引入一系列位置未知变量，从而有可能覆盖你已经定义的某些变量，通常不建议使用

fib(500) # 0 1 1 2 3 5 8 13 21 34 55 89 144 233 377

# 导入的别名问题
# 如果模块名后紧跟 as, 那么 as 后的变量名会与被导入的模块名绑定
import fibo as fib

fib.fib(500)
# 使用 from 时可以使用这个机制达到相同的效果
from fibo import fib as fibonacci

fibonacci(500)
```

当可之行模块被当作脚本时。  
```python
# 使用 import 语句时，__name__ 就会被赋值为该模块的名字；所以 __name__ == "__main__" 是避开 import 的执行
if __name__ == "__main__":
    import sys
    fib(int(sys.argv[1]))

# 在命令行里之行即可： python fibo.py [number]
```

当一个模块被导入时，解释器首先寻找同名的内建模块。如果没有发现同名内建模块，解释器会根据 sys.path 提供的一系列路径下寻找文件。sys.path 根据下面这些位置进行初始化：  

- 包含输入脚本的目录（如果没有指明文件则为当前目录）  
- PYTHONPATH 一个目录的列表，语法与 shell 的 PATH 变量相同
- 安装依赖默认路径

为了加快模块载入速度，Python 将每个模块的编译版本以 module.\*version\*.pyc 的名称缓存在 \_\_pycache\_\_ 目录下，"version" 编码编译文件的格式，它通常包含 Python 版本号（如在 CPython 3.3 编译的 spam.py 文件缓存在 \_\_pycache\_\_/spam.cpython-33.pyc 中，这种命名方式允许不同发行版本和不同版本的 Python 编译文件共存）。  
Python 检查源文件修改日期并与编译的文件进行比较以确认编译文件是否过时，需要重新编译，这是一个全自动过程。同样的，编译的模块不依赖于操作系统，所以相同的库可以在不同架构的系统之间分享。  
Python 在两种情况下不检查缓存：首先，Python 总会重新编译且不会缓存从命令行直接导入模块；其次，如果没有源模块，Python 也不会检查缓存。为了支持无源文件（只有编译文件）发布，编译的模块必须位于源目录，且不能有一个源模块。  

内置函数 dir() 用于按模块名搜索模块定义，它返回一个字符串类型的存储列表。dir() 的参数为空时，默认返回当前定义的命名。  
`注意：dir() 不会列出内置函数和变量名，如果你想列出这些内容，它们在标准模块 builtins 中定义。`

```python
import fibo

dir(fibo) # ['__name__', 'fib', 'fib2']
dir() # ['__builtins__', '__name__', 'a', 'fib', 'fibo', 'sys']
```

**包**  
包通常是使用『圆点模块名』的结构化模块命名空间，如名为 A.B 的模块表示了名为 A 的包中名为 B 的子模块。  

假设你现在想要设计一个模块集（一个 "包"）来统一处理声音文件和声音数据。存在几种不同的声音格式（通常由它们的扩展名来标识，例如：.wav， .aiff，.au ），于是，为了在不同类型的文件格式之间转换，你需要维护一个不断增长的包集合。可能你还想要对声音数据做很多不同的操作（例如混音，添加回声，应用平衡 功能，创建一个人造效果），所以你要加入一个无限流模块来执行这些操作。  
你的包可能会是这个样子（通过分级的文件体系来进行分组）：  
```
sound/                          Top-level package
      __init__.py               Initialize the sound package
      formats/                  Subpackage for file format conversions
              __init__.py
              wavread.py
              wavwrite.py
              aiffread.py
              aiffwrite.py
              auread.py
              auwrite.py
              ...
      effects/                  Subpackage for sound effects
              __init__.py
              echo.py
              surround.py
              reverse.py
              ...
      filters/                  Subpackage for filters
              __init__.py
              equalizer.py
              vocoder.py
              karaoke.py
              ...
```

为了让 Python 将目录当做内容包，目录中必须包含 \_\_init\_\_.py 文件（为了避免一个含有烂俗名字的目录无意中隐藏了稍后在模块搜索路径中出现的有效模块）。最简单的情况下，只需要一个空的 \_\_init\_\_.py 文件即可（事实上，Python3 并非强制要求）。  

`注意：使用 from package import item 方式导入包时，这个子项（item）既可以是包中的一个子模块（或一个子包），也可以是包中定义的其它命名，像函数、类或变量。import 语句首先核对是否包中有这个子项，如果没有，它假定这是一个模块，并尝试加载它。如果没有找到它，会引发一个 ImportError 异常。类似 import item.subitem.subsubitem 这样的语法时，这些子项必须是包，最后的子项可以是包或模块，但不能是前面子项中定义的类、函数或变量。`

```python
# 用户可以每次只导入包里的特定模块，但是引用的时候必需通过完整的名称
import sound.effects.echo

sound.effects.echo.echofilter(input, output, delay=0.7, atten=4)

# 另一种方式
from sound.effects import echo

echo.echofilter(input, output, delay=0.7, atten=4)

# 另一种变体用于直接导入函数或变量
from sound.effects.echo import echofilter

chofilter(input, output, delay=0.7, atten=4)
```

使用 * 导入包时，可能会消耗很长时间或边界效应。对于包的作者来说唯一的解决方案就是给提供一个明确的包索引，执行 from package import * 时，如果包中的 \_\_init\_\_.py 代码定义了一个名为 \_\_all\_\_ 的列表，就会按照列表中给出的模块名进行导入。  
```python
# sound/effects/__init__.py

__all__ = ["echo", "surround", "reverse"] # 意味着 from sound.effects import * 语句会从 sound 包中导入以上三个已命名的子模块
# 如果没有定义 __all__ ，from sound.effects import * 语句不会从 sound.effects 包中导入所有的子模块
# 无论包中定义多少命名，只能确定的是导入了 sound.effects 包（可能会运行 __init__.py 中的初始化代码）以及包中定义的所有命名会随之导入
```

包内引用。  
如果包中使用了子包结构，可以按绝对位置从相邻的包中引入子模块。比如 sound.filters.vocoder 包需要使用 sound.effects 包中的 echo 模块，它可以 from sound.Effects import echo。也可以用这样的形式 from module import name 来写显式的相对位置导入，那些显式相对导入用点号标明关联导入当前和上级包。  
`注意：无论显式或隐式相对位置导入都基于当前模块的命名，因为主模块的名字总是 "__main__"，Python 应用程序的主模块应该总是用绝对导入`。

```python
# surround 模块
from . import echo
from .. import formats
from ..filters import equalizer # 从本模块所在的层级的上一层级引用(相对)，从本模块所在的层级的上上层级引用 from ...xxx import xxxx 以此类推
```

通常，项目内的模块/包引用，可以直接使用 . 号引用项目的相对路径。  
这时候，就需要系统的环境变量里添加项目路径或者修改系统环境变量的第一项为项目根目录的绝对路径。  
```python
import sys

print(sys.path) # ["", "xxx"]
sys.path[0] = "project root path"
```
但是，显然这不是最佳的解决方案。  

比较好的方案是修改 PYTHONHOME。  
> Python 的 Virtual Environment（虚拟运行环境）。  
> Python 可以通过 Virtualenv 工具，非常方便地创建一 个全新的 Python 运行环境。  
> 建议对于每一个项目来说，最好要有一个独立的运行环境来保持包和模块的纯净性。  
> 
> 在一个 Virtual Environment 里，你能找到一个文件叫 activate，在这个文件的末尾，填上下面的内容：  
> export PYTHONPATH="project root path"  
> 这样，每次你通过 activate 激活这个运行时环境的时候，它就会自动将项目的根目录添加 到搜索路径中去。

### Python 垃圾回收机制
python 采用的是引用计数机制为主，标记-清除和分代收集两种机制为辅的策略。  

[参考：Python 中的垃圾回收机制](https://foofish.net/python-gc.html)  

### Python 代码风格
[PEP 8](https://www.python.org/dev/peps/pep-0008/) 是大多数 Python 项目使用的代码风格指南，它提供了高可读性和养眼的代码风格，以下划重点。  

- 缩减使用四个空格而不是制表符  
- 每行不要超过 79 个字符  
- 使用空行分隔函数、类或者函数内较大的代码段  
- 尽量将注释和代码放在一起  
- 用 docstrings  
- 用在操作符前后和逗号之后加空格，但是括号之内不需要： a= f(1, 2) + g(3, 4)   
- 一致性的命名类与函数：惯例是用 CamelCase 命名类，用 lower_case_with_underscores 命名函数和方法  
- 变量命名使用小写（多单词以 \_ 连接，不超过 3 个单词），保护型变量命名前加 \_，私有变量命名前加 \_\_  
- Python 默认使用 UTF-8，甚至纯 ASCII 在任何情况下都能最好地工作  

更多参考 [Google 开源项目风格指南]](https://zh-google-styleguide.readthedocs.io/en/latest/google-python-styleguide/python_style_rules/)


### 基础使用示例
#### 1 求绝对值

绝对值或复数的模

```python
In [1]: abs(-6)
Out[1]: 6
```

#### 2 元素都为真

接受一个迭代器，如果迭代器的`所有元素`都为真，那么返回`True`，否则返回`False`

```python
In [2]: all([1,0,3,6])
Out[2]: False

In [3]: all([1,2,3])
Out[3]: True
```

#### 3 元素至少一个为真　

接受一个迭代器，如果迭代器里`至少有一个`元素为真，那么返回`True`，否则返回`False`

```python
In [4]: any([0,0,0,[]])
Out[4]: False

In [5]: any([0,0,1])
Out[5]: True
```

#### 4 ascii展示对象　　

调用对象的__repr__() 方法，获得该方法的返回值，如下例子返回值为字符串

```python
In [1]: class Student():
   ...:     def __init__(self,id,name):
   ...:         self.id = id
   ...:         self.name = name
   ...:     def __repr__(self):
   ...:         return 'id = '+self.id +', name = '+self.name
   ...: 
   ...: 

In [2]: xiaoming = Student(id='001',name='xiaoming')

In [3]: print(xiaoming)
id = 001, name = xiaoming

In [4]: ascii(xiaoming)
Out[4]: 'id = 001, name = xiaoming'
```

#### 5  十转二

将`十进制`转换为`二进制`

```python
In [1]: bin(10)
Out[1]: '0b1010'
```

#### 6 十转八

将`十进制`转换为`八进制`

```python
In [1]: oct(9)
Out[1]: '0o11'
```

#### 7 十转十六

将`十进制`转换为`十六进制`

```python
In [1]: hex(15)
Out[1]: '0xf'
```

#### 8 判断是真是假　　

测试一个对象是True, 还是False.

```python
In [1]: bool([0,0,0])
Out[1]: True

In [2]: bool([])
Out[2]: False

In [3]: bool([1,0,1])
Out[3]: True
```

#### 9  字符串转字节　　

将一个`字符串`转换成`字节`类型

```python
In [1]: s = "apple"

In [2]: bytes(s,encoding='utf-8')
Out[2]: b'apple'
```

#### 10 转为字符串　　

将`字符类型`、`数值类型`等转换为`字符串`类型

```python
In [1]: i = 100

In [2]: str(i)
Out[2]: '100'
```

#### 11 是否可调用　　

判断对象是否可被调用，能被调用的对象就是一个`callable` 对象，比如函数 `str`, `int` 等都是可被调用的，但是例子**4** 中`xiaoming`实例是不可被调用的：

```python
In [1]: callable(str)
Out[1]: True

In [2]: callable(int)
Out[2]: True

In [3]: xiaoming
Out[3]: id = 001, name = xiaoming

In [4]: callable(xiaoming)
Out[4]: False
```
如果想让`xiaoming`能被调用 xiaoming(), 需要重写`Student`类的`__call__`方法：

```python
In [1]: class Student():
    ...:     def __init__(self,id,name):
    ...:         self.id = id
    ...:         self.name = name
    ...:     def __repr__(self):
    ...:         return 'id = '+self.id +', name = '+self.name
    ...:     def __call__(self):
    ...:         print('I can be called')
    ...:         print(f'my name is {self.name}')
    ...: 
    ...: 

In [2]: t = Student('001','xiaoming')

In [3]: t()
I can be called
my name is xiaoming

```



#### 12 十转ASCII

查看十进制整数对应的`ASCII字符`

```python
In [1]: chr(65)
Out[1]: 'A'
```

#### 13 ASCII转十

查看某个`ASCII字符`对应的十进制数

```python
In [1]: ord('A')
Out[1]: 65
```

#### 14 静态方法　

`classmethod` 装饰器对应的函数不需要实例化，不需要 `self `参数，但第一个参数需要是表示自身类的 cls 参数，可以来调用类的属性，类的方法，实例化对象等。

```python
In [1]: class Student():
    ...:     def __init__(self,id,name):
    ...:         self.id = id
    ...:         self.name = name
    ...:     def __repr__(self):
    ...:         return 'id = '+self.id +', name = '+self.name
    ...:     @classmethod
    ...:     def f(cls):
    ...:         print(cls)
```

#### 15 执行字符串表示的代码

将字符串编译成python能识别或可执行的代码，也可以将文字读成字符串再编译。

```python
In [1]: s  = "print('helloworld')"
    
In [2]: r = compile(s,"<string>", "exec")
    
In [3]: r
Out[3]: <code object <module> at 0x0000000005DE75D0, file "<string>", line 1>
    
In [4]: exec(r)
helloworld
```

#### 16  创建复数

创建一个复数

```python
In [1]: complex(1,2)
Out[1]: (1+2j)
```

#### 17 动态删除属性　　

删除对象的属性

```python
In [1]: delattr(xiaoming,'id')

In [2]: hasattr(xiaoming,'id')
Out[2]: False
```

#### 18 转为字典　　

创建数据字典

```python
In [1]: dict()
Out[1]: {}

In [2]: dict(a='a',b='b')
Out[2]: {'a': 'a', 'b': 'b'}

In [3]: dict(zip(['a','b'],[1,2]))
Out[3]: {'a': 1, 'b': 2}

In [4]: dict([('a',1),('b',2)])
Out[4]: {'a': 1, 'b': 2}
```

#### 19 一键查看对象所有方法　

不带参数时返回`当前范围`内的变量、方法和定义的类型列表；带参数时返回`参数`的属性，方法列表。

```python
In [96]: dir(xiaoming)
Out[96]:
['__class__',
 '__delattr__',
 '__dict__',
 '__dir__',
 '__doc__',
 '__eq__',
 '__format__',
 '__ge__',
 '__getattribute__',
 '__gt__',
 '__hash__',
 '__init__',
 '__init_subclass__',
 '__le__',
 '__lt__',
 '__module__',
 '__ne__',
 '__new__',
 '__reduce__',
 '__reduce_ex__',
 '__repr__',
 '__setattr__',
 '__sizeof__',
 '__str__',
 '__subclasshook__',
 '__weakref__',
 
 'name']
```

#### 20 取商和余数　　

分别取商和余数

```python
In [1]: divmod(10,3)
Out[1]: (3, 1)
```

#### 21 枚举对象　　

返回一个可以枚举的对象，该对象的next()方法将返回一个元组。

```python
In [1]: s = ["a","b","c"]
    ...: for i ,v in enumerate(s,1):
    ...:     print(i,v)
    ...:
1 a
2 b
3 c
```

#### 22 计算表达式

将字符串str 当成有效的表达式来求值并返回计算结果取出字符串中内容

```python
In [1]: s = "1 + 3 +5"
    ...: eval(s)
    ...:
Out[1]: 9
```

#### 23 查看变量所占字节数

```python
In [1]: import sys

In [2]: a = {'a':1,'b':2.0}

In [3]: sys.getsizeof(a) # 占用240个字节
Out[3]: 240
```

#### 24 过滤器　　

在函数中设定过滤条件，迭代元素，保留返回值为`True`的元素：

```python
In [1]: fil = filter(lambda x: x>10,[1,11,2,45,7,6,13])

In [2]: list(fil)
Out[2]: [11, 45, 13]
```

#### 25 转为浮点类型　

将一个整数或数值型字符串转换为浮点数

```python
In [1]: float(3)
Out[1]: 3.0
```
如果不能转化为浮点数，则会报`ValueError`:
```python
In [2]: float('a')
ValueError                                Traceback (most recent call last)
<ipython-input-11-99859da4e72c> in <module>()
----> 1 float('a')

ValueError: could not convert string to float: 'a'

```

#### 26 字符串格式化　

格式化输出字符串，format(value, format_spec)实质上是调用了value的__format__(format_spec)方法。

```
In [104]: print("i am {0},age{1}".format("tom",18))
i am tom,age18
```

| 3.1415926  | {:.2f}  | 3.14      | 保留小数点后两位             |
| ---------- | ------- | --------- | ---------------------------- |
| 3.1415926  | {:+.2f} | +3.14     | 带符号保留小数点后两位       |
| -1         | {:+.2f} | -1.00     | 带符号保留小数点后两位       |
| 2.71828    | {:.0f}  | 3         | 不带小数                     |
| 5          | {:0>2d} | 05        | 数字补零 (填充左边, 宽度为2) |
| 5          | {:x<4d} | 5xxx      | 数字补x (填充右边, 宽度为4)  |
| 10         | {:x<4d} | 10xx      | 数字补x (填充右边, 宽度为4)  |
| 1000000    | {:,}    | 1,000,000 | 以逗号分隔的数字格式         |
| 0.25       | {:.2%}  | 25.00%    | 百分比格式                   |
| 1000000000 | {:.2e}  | 1.00e+09  | 指数记法                     |
| 18         | {:>10d} | ' 18'     | 右对齐 (默认, 宽度为10)      |
| 18         | {:<10d} | '18 '     | 左对齐 (宽度为10)            |
| 18         | {:^10d} | ' 18 '    | 中间对齐 (宽度为10)          |

#### 27 冻结集合　　

创建一个不可修改的集合。

```python
In [1]: frozenset([1,1,3,2,3])
Out[1]: frozenset({1, 2, 3})
```
因为不可修改，所以没有像`set`那样的`add`和`pop`方法

#### 28 动态获取对象属性　

获取对象的属性

```python
In [1]: class Student():
   ...:     def __init__(self,id,name):
   ...:         self.id = id
   ...:         self.name = name
   ...:     def __repr__(self):
   ...:         return 'id = '+self.id +', name = '+self.name

In [2]: xiaoming = Student(id='001',name='xiaoming')
In [3]: getattr(xiaoming,'name') # 获取xiaoming这个实例的name属性值
Out[3]: 'xiaoming'
```

#### 29 对象是否有这个属性

```python
In [1]: class Student():
   ...:     def __init__(self,id,name):
   ...:         self.id = id
   ...:         self.name = name
   ...:     def __repr__(self):
   ...:         return 'id = '+self.id +', name = '+self.name

In [2]: xiaoming = Student(id='001',name='xiaoming')
In [3]: hasattr(xiaoming,'name')
Out[3]: True

In [4]: hasattr(xiaoming,'address')
Out[4]: False
```

#### 30 返回对象的哈希值　　

返回对象的哈希值，值得注意的是自定义的实例都是可哈希的，`list`, `dict`, `set`等可变对象都是不可哈希的(unhashable)

  ```python
In [1]: hash(xiaoming)
Out[1]: 6139638

In [2]: hash([1,2,3])
TypeError                                 Traceback (most recent call last)
<ipython-input-32-fb5b1b1d9906> in <module>()
----> 1 hash([1,2,3])

TypeError: unhashable type: 'list'
  ```

#### 31  一键帮助　

返回对象的帮助文档

```python
In [1]: help(xiaoming)
Help on Student in module __main__ object:

class Student(builtins.object)
 |  Methods defined here:
 |
 |  __init__(self, id, name)
 |
 |  __repr__(self)
 |
 |  Data descriptors defined here:
 |
 |  __dict__
 |      dictionary for instance variables (if defined)
 |
 |  __weakref__
 |      list of weak references to the object (if defined)
```

#### 32 对象门牌号　

返回对象的内存地址

```python
In [1]: id(xiaoming)
Out[1]: 98234208
```

#### 33 获取用户输入　

获取用户输入内容

```python
In [1]: input()
aa
Out[1]: 'aa'
```

#### 34  转为整型　　

int(x, base =10) , x可能为字符串或数值，将x 转换为一个普通整数。如果参数是字符串，那么它可能包含符号和小数点。如果超出了普通整数的表示范围，一个长整数被返回。  

```python
In [1]: int('12',16)
Out[1]: 18
```

#### 35 isinstance

判断*object*是否为类*classinfo*的实例，是返回true

```python
In [1]: class Student():
   ...:     def __init__(self,id,name):
   ...:         self.id = id
   ...:         self.name = name
   ...:     def __repr__(self):
   ...:         return 'id = '+self.id +', name = '+self.name

In [2]: xiaoming = Student(id='001',name='xiaoming')

In [3]: isinstance(xiaoming,Student)
Out[3]: True
```

#### 36 父子关系鉴定

```python
In [1]: class undergraduate(Student):
    ...:     def studyClass(self):
    ...:         pass
    ...:     def attendActivity(self):
    ...:         pass

In [2]: issubclass(undergraduate,Student)
Out[2]: True

In [3]: issubclass(object,Student)
Out[3]: False

In [4]: issubclass(Student,object)
Out[4]: True
```

如果class是classinfo元组中某个元素的子类，也会返回True

```python
In [1]: issubclass(int,(int,float))
Out[1]: True
```

#### 37 创建迭代器类型

使用`iter(obj, sentinel)`, 返回一个可迭代对象, sentinel可省略(一旦迭代到此元素，立即终止)

```python
In [1]: lst = [1,3,5]

In [2]: for i in iter(lst):
    ...:     print(i)
    ...:
1
3
5
```

```python
In [1]: class TestIter(object):
    ...:     def __init__(self):
    ...:         self.l=[1,3,2,3,4,5]
    ...:         self.i=iter(self.l)
    ...:     def __call__(self):  #定义了__call__方法的类的实例是可调用的
    ...:         item = next(self.i)
    ...:         print ("__call__ is called,fowhich would return",item)
    ...:         return item
    ...:     def __iter__(self): #支持迭代协议(即定义有__iter__()函数)
    ...:         print ("__iter__ is called!!")
    ...:         return iter(self.l)
In [2]: t = TestIter()
In [3]: t() # 因为实现了__call__，所以t实例能被调用
__call__ is called,which would return 1
Out[3]: 1

In [4]: for e in TestIter(): # 因为实现了__iter__方法，所以t能被迭代
    ...:     print(e)
    ...: 
__iter__ is called!!
1
3
2
3
4
5
```

#### 38 所有对象之根

object 是所有类的基类

```python
In [1]: o = object()

In [2]: type(o)
Out[2]: object
```

#### 39 打开文件

返回文件对象

```python
In [1]: fo = open('D:/a.txt',mode='r', encoding='utf-8')

In [2]: fo.read()
Out[2]: '\ufefflife is not so long,\nI use Python to play.'
```

mode取值表：

| 字符  | 意义                             |
| :---- | :------------------------------- |
| `'r'` | 读取（默认）                     |
| `'w'` | 写入，并先截断文件               |
| `'x'` | 排它性创建，如果文件已存在则失败 |
| `'a'` | 写入，如果文件存在则在末尾追加   |
| `'b'` | 二进制模式                       |
| `'t'` | 文本模式（默认）                 |
| `'+'` | 打开用于更新（读取与写入）       |

#### 40 次幂

base为底的exp次幂，如果mod给出，取余

```python
In [1]: pow(3, 2, 4)
Out[1]: 1
```

#### 41 打印

```python
In [5]: lst = [1,3,5]

In [6]: print(lst)
[1, 3, 5]

In [7]: print(f'lst: {lst}')
lst: [1, 3, 5]

In [8]: print('lst:{}'.format(lst))
lst:[1, 3, 5]

In [9]: print('lst:',lst)
lst: [1, 3, 5]
```



#### 42  创建属性的两种方式

返回 property 属性，典型的用法：

```python
class C:
    def __init__(self):
        self._x = None

    def getx(self):
        return self._x

    def setx(self, value):
        self._x = value

    def delx(self):
        del self._x
    # 使用property类创建 property 属性
    x = property(getx, setx, delx, "I'm the 'x' property.")
```

使用python装饰器，实现与上完全一样的效果代码：

```python
class C:
    def __init__(self):
        self._x = None

    @property
    def x(self):
        return self._x

    @x.setter
    def x(self, value):
        self._x = value

    @x.deleter
    def x(self):
        del self._x
```

#### 43 创建range序列

1) range(stop)
2) range(start, stop[,step])

生成一个不可变序列：

```python
In [1]: range(11)
Out[1]: range(0, 11)

In [2]: range(0,11,1)
Out[2]: range(0, 11)
```

#### 44 反向迭代器

```python
In [1]: rev = reversed([1,4,2,3,1])

In [2]: for i in rev:
     ...:     print(i)
     ...:
1
3
2
4
1
```

#### 45 四舍五入

四舍五入，`ndigits`代表小数点后保留几位：

```python
In [11]: round(10.0222222, 3)
Out[11]: 10.022

In [12]: round(10.05,1)
Out[12]: 10.1
```

#### 46 转为集合类型

返回一个set对象，集合内不允许有重复元素：

```python
In [159]: a = [1,4,2,3,1]

In [160]: set(a)
Out[160]: {1, 2, 3, 4}
```

#### 47 转为切片对象

*class* slice(*start*, *stop*[, *step*])

返回一个表示由 range(start, stop, step) 所指定索引集的 slice对象，它让代码可读性、可维护性变好。

```python
In [1]: a = [1,4,2,3,1]

In [2]: my_slice_meaning = slice(0,5,2)

In [3]: a[my_slice_meaning]
Out[3]: [1, 2, 1]
```

#### 48 拿来就用的排序函数

排序：

```python
In [1]: a = [1,4,2,3,1]

In [2]: sorted(a,reverse=True)
Out[2]: [4, 3, 2, 1, 1]

In [3]: a = [{'name':'xiaoming','age':18,'gender':'male'},{'name':'
     ...: xiaohong','age':20,'gender':'female'}]
In [4]: sorted(a,key=lambda x: x['age'],reverse=False)
Out[4]:
[{'name': 'xiaoming', 'age': 18, 'gender': 'male'},
 {'name': 'xiaohong', 'age': 20, 'gender': 'female'}]
```

####49 求和函数

求和：

```python
In [181]: a = [1,4,2,3,1]

In [182]: sum(a)
Out[182]: 11

In [185]: sum(a,10) #求和的初始值为10
Out[185]: 21
```

#### 50 转元组

 `tuple()` 将对象转为一个不可变的序列类型

 ```python
 In [16]: i_am_list = [1,3,5]
 In [17]: i_am_tuple = tuple(i_am_list)
 In [18]: i_am_tuple
 Out[18]: (1, 3, 5)
 ```

#### 51 查看对象类型

*class* `type`(*name*, *bases*, *dict*)

传入一个参数时，返回 *object* 的类型：

```python
In [1]: class Student():
   ...:     def __init__(self,id,name):
   ...:         self.id = id
   ...:         self.name = name
   ...:     def __repr__(self):
   ...:         return 'id = '+self.id +', name = '+self.name
   ...: 
   ...: 

In [2]: xiaoming = Student(id='001',name='xiaoming')
In [3]: type(xiaoming)
Out[3]: __main__.Student

In [4]: type(tuple())
Out[4]: tuple
```

#### 52 聚合迭代器

创建一个聚合了来自每个可迭代对象中的元素的迭代器：

```python
In [1]: x = [3,2,1]
In [2]: y = [4,5,6]
In [3]: list(zip(y,x))
Out[3]: [(4, 3), (5, 2), (6, 1)]

In [4]: a = range(5)
In [5]: b = list('abcde')
In [6]: b
Out[6]: ['a', 'b', 'c', 'd', 'e']
In [7]: [str(y) + str(x) for x,y in zip(a,b)]
Out[7]: ['a0', 'b1', 'c2', 'd3', 'e4']
```

#### 53 nonlocal用于内嵌函数中

关键词`nonlocal`常用于函数嵌套中，声明变量`i`为非局部变量；
如果不声明，`i+=1`表明`i`为函数`wrapper`内的局部变量，因为在`i+=1`引用(reference)时,i未被声明，所以会报`unreferenced variable`的错误。
```python
def excepter(f):
    i = 0
    t1 = time.time()
    def wrapper(): 
        try:
            f()
        except Exception as e:
            nonlocal i
            i += 1
            print(f'{e.args[0]}: {i}')
            t2 = time.time()
            if i == n:
                print(f'spending time:{round(t2-t1,2)}')
    return wrapper
```

#### 54 global 声明全局变量
先回答为什么要有`global`，一个变量被多个函数引用，想让全局变量被所有函数共享。有的伙伴可能会想这还不简单，这样写：
```python
i = 5
def f():
    print(i)

def g():
    print(i)
    pass

f()
g()

```
f和g两个函数都能共享变量`i`，程序没有报错，所以他们依然不明白为什么要用`global`.

但是，如果我想要有个函数对`i`递增，这样：
```python
def h():
    i += 1

h()
```
此时执行程序，bang, 出错了！ 抛出异常：`UnboundLocalError`，原来编译器在解释`i+=1`时会把`i`解析为函数`h()`内的局部变量，很显然在此函数内，编译器找不到对变量`i`的定义，所以会报错。

`global`就是为解决此问题而被提出，在函数h内，显示地告诉编译器`i`为全局变量，然后编译器会在函数外面寻找`i`的定义，执行完`i+=1`后，`i`还为全局变量，值加1：
```python
i = 0
def h():
    global i
    i += 1

h()
print(i)
```

#### 55 链式比较

```python
i = 3
print(1 < i < 3)  # False
print(1 < i <= 3)  # True
```


#### 56 不用else和if实现计算器

```python
from operator import *


def calculator(a, b, k):
    return {
        '+': add,
        '-': sub,
        '*': mul,
        '/': truediv,
        '**': pow
    }[k](a, b)


calculator(1, 2, '+')  # 3
calculator(3, 4, '**')  # 81
```

#### 57 链式操作

```python
from operator import (add, sub)


def add_or_sub(a, b, oper):
    return (add if oper == '+' else sub)(a, b)


add_or_sub(1, 2, '-')  # -1
```

#### 58 交换两元素

```python
def swap(a, b):
    return b, a


print(swap(1, 0))  # (0,1)
```

#### 59 去最求平均

```python
def score_mean(lst):
    lst.sort()
    lst2=lst[1:(len(lst)-1)]
    return round((sum(lst2)/len(lst2)),1)

lst=[9.1, 9.0,8.1, 9.7, 19,8.2, 8.6,9.8]
score_mean(lst) # 9.1
```

#### 60 打印99乘法表

打印出如下格式的乘法表

```python
1*1=1
1*2=2   2*2=4
1*3=3   2*3=6   3*3=9
1*4=4   2*4=8   3*4=12  4*4=16
1*5=5   2*5=10  3*5=15  4*5=20  5*5=25
1*6=6   2*6=12  3*6=18  4*6=24  5*6=30  6*6=36
1*7=7   2*7=14  3*7=21  4*7=28  5*7=35  6*7=42  7*7=49
1*8=8   2*8=16  3*8=24  4*8=32  5*8=40  6*8=48  7*8=56  8*8=64
1*9=9   2*9=18  3*9=27  4*9=36  5*9=45  6*9=54  7*9=63  8*9=72  9*9=81
```

一共有10 行，第`i`行的第`j`列等于：`j*i`，

其中,

 `i`取值范围：`1<=i<=9`

 `j`取值范围：`1<=j<=i`

根据`例子分析`的语言描述，转化为如下代码：

```python
for i in range(1,10):
    ...:     for j in range(1,i+1):
    ...:         print('%d*%d=%d'%(j,i,j*i),end="\t")
    ...:     print()
```

#### 61 全展开

对于如下数组：

```
[[[1,2,3],[4,5]]]
```

如何完全展开成一维的。这个小例子实现的`flatten`是递归版，两个参数分别表示带展开的数组，输出数组。

```python
from collections.abc import *

def flatten(lst, out_lst=None):
    if out_lst is None:
        out_lst = []
    for i in lst:
        if isinstance(i, Iterable): # 判断i是否可迭代
            flatten(i, out_lst)  # 尾数递归
        else:
            out_lst.append(i)    # 产生结果
    return out_lst
```

调用`flatten`:

```python
print(flatten([[1,2,3],[4,5]]))
print(flatten([[1,2,3],[4,5]], [6,7]))
print(flatten([[[1,2,3],[4,5,6]]]))
# 结果：
[1, 2, 3, 4, 5]
[6, 7, 1, 2, 3, 4, 5]
[1, 2, 3, 4, 5, 6]
```

numpy里的`flatten`与上面的函数实现有些微妙的不同：

```python
import numpy
b = numpy.array([[1,2,3],[4,5]])
b.flatten()
array([list([1, 2, 3]), list([4, 5])], dtype=object)
```

#### 62 列表等分

```python
from math import ceil

def divide(lst, size):
    if size <= 0:
        return [lst]
    return [lst[i * size:(i+1)*size] for i in range(0, ceil(len(lst) / size))]


r = divide([1, 3, 5, 7, 9], 2)
print(r)  # [[1, 3], [5, 7], [9]]

r = divide([1, 3, 5, 7, 9], 0)
print(r)  # [[1, 3, 5, 7, 9]]

r = divide([1, 3, 5, 7, 9], -3)
print(r)  # [[1, 3, 5, 7, 9]]

```

#### 63 列表压缩

```python
def filter_false(lst):
    return list(filter(bool, lst))


r = filter_false([None, 0, False, '', [], 'ok', [1, 2]])
print(r)  # ['ok', [1, 2]]

```

#### 64 更长列表

```python
def max_length(*lst):
    return max(*lst, key=lambda v: len(v))


r = max_length([1, 2, 3], [4, 5, 6, 7], [8])
print(f'更长的列表是{r}')  # [4, 5, 6, 7]

r = max_length([1, 2, 3], [4, 5, 6, 7], [8, 9])
print(f'更长的列表是{r}')  # [4, 5, 6, 7]
```

#### 65 求众数

```python
def top1(lst):
    return max(lst, default='列表为空', key=lambda v: lst.count(v))

lst = [1, 3, 3, 2, 1, 1, 2]
r = top1(lst)
print(f'{lst}中出现次数最多的元素为:{r}')  # [1, 3, 3, 2, 1, 1, 2]中出现次数最多的元素为:1
```

#### 66 多表之最
```python 
def max_lists(*lst):
    return max(max(*lst, key=lambda v: max(v)))


r = max_lists([1, 2, 3], [6, 7, 8], [4, 5])
print(r)  # 8
```

#### 67 列表查重

```python
def has_duplicates(lst):
    return len(lst) == len(set(lst))


x = [1, 1, 2, 2, 3, 2, 3, 4, 5, 6]
y = [1, 2, 3, 4, 5]
has_duplicates(x)  # False
has_duplicates(y)  # True
```




#### 68 列表反转

```python
def reverse(lst):
    return lst[::-1]


r = reverse([1, -2, 3, 4, 1, 2])
print(r)  # [2, 1, 4, 3, -2, 1]
```

#### 69 浮点数等差数列

```python
def rang(start, stop, n):
    start,stop,n = float('%.2f' % start), float('%.2f' % stop),int('%.d' % n)
    step = (stop-start)/n
    lst = [start]
    while n > 0:
        start,n = start+step,n-1
        lst.append(round((start), 2))
    return lst

rang(1, 8, 10) # [1.0, 1.7, 2.4, 3.1, 3.8, 4.5, 5.2, 5.9, 6.6, 7.3, 8.0]
```

#### 70 按条件分组

```python
def bif_by(lst, f):
    return [ [x for x in lst if f(x)],[x for x in lst if not f(x)]]

records = [25,89,31,34] 
bif_by(records, lambda x: x<80) # [[25, 31, 34], [89]]
```


#### 71 map实现向量运算

```python
#多序列运算函数—map(function,iterabel,iterable2)
lst1=[1,2,3,4,5,6]
lst2=[3,4,5,6,3,2]
list(map(lambda x,y:x*y+1,lst1,lst2))
### [4, 9, 16, 25, 16, 13]
```

#### 72 值最大的字典

```python
def max_pairs(dic):
    if len(dic) == 0:
        return dic
    max_val = max(map(lambda v: v[1], dic.items()))
    return [item for item in dic.items() if item[1] == max_val]


r = max_pairs({'a': -10, 'b': 5, 'c': 3, 'd': 5})
print(r)  # [('b', 5), ('d', 5)]
```

#### 73 合并两个字典

```python
def merge_dict(dic1, dic2):
    return {**dic1, **dic2}  # python3.5后支持的一行代码实现合并字典

merge_dict({'a': 1, 'b': 2}, {'c': 3})  # {'a': 1, 'b': 2, 'c': 3}
```

#### 74 topn字典

```python
from heapq import nlargest

# 返回字典d前n个最大值对应的键

def topn_dict(d, n):
    return nlargest(n, d, key=lambda k: d[k])

topn_dict({'a': 10, 'b': 8, 'c': 9, 'd': 10}, 3)  # ['a', 'd', 'c']
```


#### 75 异位词

```python
from collections import Counter

# 检查两个字符串是否 相同字母异序词，简称：互为变位词

def anagram(str1, str2):
    return Counter(str1) == Counter(str2)

anagram('eleven+two', 'twelve+one')  # True 这是一对神器的变位词
anagram('eleven', 'twelve')  # False
```


#### 76 逻辑上合并字典
(1) 两种合并字典方法
这是一般的字典合并写法

```python
dic1 = {'x': 1, 'y': 2 }
dic2 = {'y': 3, 'z': 4 }
merged1 = {**dic1, **dic2} # {'x': 1, 'y': 3, 'z': 4}
```

修改merged['x']=10，dic1中的x值`不变`，`merged`是重新生成的一个`新字典`。

但是，`ChainMap`却不同，它在内部创建了一个容纳这些字典的列表。因此使用ChainMap合并字典，修改merged['x']=10后，dic1中的x值`改变`，如下所示：

```python
from collections import ChainMap
merged2 = ChainMap(dic1,dic2)
print(merged2) # ChainMap({'x': 1, 'y': 2}, {'y': 3, 'z': 4})
```

#### 77 命名元组提高可读性

```python
from collections import namedtuple
Point = namedtuple('Point', ['x', 'y', 'z'])  # 定义名字为Point的元祖，字段属性有x,y,z
lst = [Point(1.5, 2, 3.0), Point(-0.3, -1.0, 2.1), Point(1.3, 2.8, -2.5)]
print(lst[0].y - lst[1].y)
```

使用命名元组写出来的代码可读性更好，尤其处理上百上千个属性时作用更加凸显。

#### 78 样本抽样

使用`sample`抽样，如下例子从100个样本中随机抽样10个。

```python
from random import randint,sample
lst = [randint(0,50) for _ in range(100)]
print(lst[:5])# [38, 19, 11, 3, 6]
lst_sample = sample(lst,10)
print(lst_sample) # [33, 40, 35, 49, 24, 15, 48, 29, 37, 24]
```

#### 79 重洗数据集

使用`shuffle`用来重洗数据集，**值得注意`shuffle`是对lst就地(in place)洗牌，节省存储空间**

```python
from random import shuffle
lst = [randint(0,50) for _ in range(100)]
shuffle(lst)
print(lst[:5]) # [50, 3, 48, 1, 26]
```

#### 80 10个均匀分布的坐标点

random模块中的`uniform(a,b)`生成[a,b)内的一个随机数，如下生成10个均匀分布的二维坐标点

```python
from random import uniform
In [1]: [(uniform(0,10),uniform(0,10)) for _ in range(10)]
Out[1]: 
[(9.244361194237328, 7.684326645514235),
 (8.129267671737324, 9.988395854203773),
 (9.505278771040661, 2.8650440524834107),
 (3.84320100484284, 1.7687190176304601),
 (6.095385729409376, 2.377133802224657),
 (8.522913365698605, 3.2395995841267844),
 (8.827829601859406, 3.9298809217233766),
 (1.4749644859469302, 8.038753079253127),
 (9.005430657826324, 7.58011186920019),
 (8.700789540392917, 1.2217577293254112)]
```

#### 81 10个高斯分布的坐标点

random模块中的`gauss(u,sigma)`生成均值为u, 标准差为sigma的满足高斯分布的值，如下生成10个二维坐标点，样本误差(y-2*x-1)满足均值为0，标准差为1的高斯分布：

```python
from random import gauss
x = range(10)
y = [2*xi+1+gauss(0,1) for xi in x]
points = list(zip(x,y))
### 10个二维点：
[(0, -0.86789025305992),
 (1, 4.738439437453464),
 (2, 5.190278040856102),
 (3, 8.05270893133576),
 (4, 9.979481700775292),
 (5, 11.960781766216384),
 (6, 13.025427054303737),
 (7, 14.02384035204836),
 (8, 15.33755823101161),
 (9, 17.565074449028497)]
```

#### 82 chain高效串联多个容器对象

`chain`函数串联a和b，兼顾内存效率同时写法更加优雅。

```python
from itertools import chain
a = [1,3,5,0]
b = (2,4,6)

for i in chain(a,b):
  print(i)
### 结果
1
3
5
0
2
4
6
```

#### 83 操作函数对象

```python
In [31]: def f():
    ...:     print('i\'m f')
    ...:

In [32]: def g():
    ...:     print('i\'m g')
    ...:

In [33]: [f,g][1]()
i'm g
```

创建函数对象的list，根据想要调用的index，方便统一调用。

#### 84 生成逆序序列

```python
list(range(10,-1,-1)) # [10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0]
```

第三个参数为负时，表示从第一个参数开始递减，终止到第二个参数(不包括此边界)

#### 85 函数的五类参数使用例子

python五类参数：位置参数，关键字参数，默认参数，可变位置或关键字参数的使用。

```python
def f(a,*b,c=10,**d):
  print(f'a:{a},b:{b},c:{c},d:{d}')
```
*默认参数`c`不能位于可变关键字参数`d`后.*

调用f:
```python
In [10]: f(1,2,5,width=10,height=20)
a:1,b:(2, 5),c:10,d:{'width': 10, 'height': 20}
```
可变位置参数`b`实参后被解析为元组`(2,5)`;而c取得默认值10; d被解析为字典.

再次调用f:
```python
In [11]: f(a=1,c=12)
a:1,b:(),c:12,d:{}
```
a=1传入时a就是关键字参数，b,d都未传值，c被传入12，而非默认值。

注意观察参数`a`, 既可以`f(1)`,也可以`f(a=1)` 其可读性比第一种更好，建议使用f(a=1)。如果要强制使用`f(a=1)`，需要在前面添加一个**星号**:
```python
def f(*,a,*b):
  print(f'a:{a},b:{b}')
```
此时f(1)调用，将会报错：`TypeError: f() takes 0 positional arguments but 1 was given`

只能`f(a=1)`才能OK.

说明前面的`*`发挥作用，它变为只能传入关键字参数，那么如何查看这个参数的类型呢？借助python的`inspect`模块：

```python
In [22]: for name,val in signature(f).parameters.items():
    ...:     print(name,val.kind)
    ...:
a KEYWORD_ONLY
b VAR_KEYWORD
```

可看到参数`a`的类型为`KEYWORD_ONLY`，也就是仅仅为关键字参数。

但是，如果f定义为：
```python
def f(a,*b):
  print(f'a:{a},b:{b}')
```
查看参数类型：
```python
In [24]: for name,val in signature(f).parameters.items():
    ...:     print(name,val.kind)
    ...:
a POSITIONAL_OR_KEYWORD
b VAR_POSITIONAL
```
可以看到参数`a`既可以是位置参数也可是关键字参数。

#### 86  使用slice对象

生成关于蛋糕的序列cake1：

```
In [1]: cake1 = list(range(5,0,-1))

In [2]: b = cake1[1:10:2]

In [3]: b
Out[3]: [4, 2]

In [4]: cake1
Out[4]: [5, 4, 3, 2, 1]
```

再生成一个序列：

```
In [5]: from random import randint
   ...: cake2 = [randint(1,100) for _ in range(100)]
   ...: # 同样以间隔为2切前10个元素，得到切片d
   ...: d = cake2[1:10:2]
In [6]: d
Out[6]: [75, 33, 63, 93, 15]
```

你看，我们使用同一种切法，分别切开两个蛋糕cake1,cake2. 后来发现这种切法`极为经典`，又拿它去切更多的容器对象。

那么，为什么不把这种切法封装为一个对象呢？于是就有了slice对象。

定义slice对象极为简单，如把上面的切法定义成slice对象：

```
perfect_cake_slice_way = slice(1,10,2)
#去切cake1
cake1_slice = cake1[perfect_cake_slice_way] 
cake2_slice = cake2[perfect_cake_slice_way]

In [11]: cake1_slice
Out[11]: [4, 2]

In [12]: cake2_slice
Out[12]: [75, 33, 63, 93, 15]
```

与上面的结果一致。

对于逆向序列切片，`slice`对象一样可行：

```
a = [1,3,5,7,9,0,3,5,7]
a_ = a[5:1:-1]

named_slice = slice(5,1,-1)
a_slice = a[named_slice] 

In [14]: a_
Out[14]: [0, 9, 7, 5]

In [15]: a_slice
Out[15]: [0, 9, 7, 5]
```

频繁使用同一切片的操作可使用slice对象抽出来，复用的同时还能提高代码可读性。

#### 87 str1是否为str2的permutation

排序词(permutation)：两个字符串含有相同字符，但字符顺序不同。

```python
from collections import defaultdict


def is_permutation(str1, str2):
    if str1 is None or str2 is None:
        return False
    if len(str1) != len(str2):
        return False
    unq_s1 = defaultdict(int)
    unq_s2 = defaultdict(int)
    for c1 in unq_s1:
        unq_s1[c1] += 1
    for c2 in unq_s2:
        unq_s2[c2] += 1

    return unq_s1 == unq_s2
```

这个小例子，使用python内置的`defaultdict`，默认类型初始化为`int`，计数默次数都为0. 这个解法本质是 `hash map lookup`

统计出的两个defaultdict：unq_s1，unq_s2，如果相等，就表明str1、 str2互为排序词。

下面测试：
```python
r = is_permutation('nice', 'cine')
print(r)  # True

r = is_permutation('', '')
print(r)  # True

r = is_permutation('', None)
print(r)  # False

r = is_permutation('work', 'woo')
print(r)  # False

```
以上就是使用defaultdict的小例子，希望对读者朋友理解此类型有帮助。

#### 88 str1是否由str2旋转而来

`stringbook`旋转后得到`bookstring`,写一段代码验证`str1`是否为`str2`旋转得到。

**思路**

转化为判断：`str1`是否为`str2+str2`的子串

```python
def is_rotation(s1: str, s2: str) -> bool:
    if s1 is None or s2 is None:
        return False
    if len(s1) != len(s2):
        return False

    def is_substring(s1: str, s2: str) -> bool:
        return s1 in s2
    return is_substring(s1, s2 + s2)
```

**测试**
```python
r = is_rotation('stringbook', 'bookstring')
print(r)  # True

r = is_rotation('greatman', 'maneatgr')
print(r)  # False
```
