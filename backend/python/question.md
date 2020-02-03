
### 常见问题及 Python 进阶
1、int() 强制转换浮点数     
在 int（）的强制转换浮点数时候，不管是正数还是负数，只取整数部分。  
`注意：这里不是向上或者向下取整，也不是四舍五入。`

2、字符串是不可变的  
```python
string = 'test'
string[0] = 'a' # TypeError

# 如果非要修改，联合使用字符串的切片和拼接完成
string = 'test'
string = 'a' + string[1:]
```

3、注意操作的返回值  
print 是没有返回值的，同样的，在使用一些操作时也需要注意这样的情况。  
```python
ls1 = ['a', 'c', 'b']
ls2 = ls1.sort()
print(ls2) # None
```

4、字符串快速转换为列表  
```python
ls1 = [i for i in 'Python']
ls2 = list('Python')

string = 'I love Python'
ls3 = string.split(' ')
```

5、字典中按照键的顺序输出  
字典有三大特性其中有一个就是无序性，如果要按照键的排序，可以给字典套上个 sorted 的外衣  
```python
dict1 = {'b': 111, 'a': 212, 'c': 333}
for k in sorted(dict1):
	print(k, dict1[k])
```

6、可变长元组参数  
函数的可变长参数 * 开头的收集到的的是元组参数，既然有收集，那就可能会有分散。  
```python
def div_mod(num1, num2):
	return num1 // num2, num1 % num2

tuple1 = (3, 2)
# 传入的元组前加个 * 号，即可将元组拆开
div_mod(*tuple1)
```

7、文件名和路径  
```python
import os


# 获取当前目录路径
cwd = os.getcwd()
print(cwd)
# 检查该目录下是否存在某个文件
print(os.path.exists('xxx'))
```

8、序列合体   
将序列（字符串，列表，元组）相互合并的操作。
```python
ls = [1,2,3,4,5]
tp = (1,2,3,4,5)
new_ls = [(x,y) for x in ls for y in tp if x == y]
# 更简便的办法是 zip(ls, tp)
```

9、\_\_init\_\_ 和 \_\_new\_\_，及 \_\_call\_\_  
\_\_init\_\_ 方法负责对象的初始化，系统执行该方法前，其实该对象已经存在了，要不然初始化什么东西呢？  
事实上，对象实例化在初始化之前，即 \_\_new\_\_ 方法先被调用，返回一个实例对象，接着 \_\_init\_\_ 被调用。  
```python
# class A(object): python2 必须显示地继承object
class A:
    def __init__(self):
	    print("__init__ ")
	    print(self)
	    super(A, self).__init__()
	    # Python 规定，__init__只能返回 None 值，否则报错

	def __new__(cls):
	    print("__new__ ")
	    self = super(A, cls).__new__(cls)
	    print(self)
	    # __new__ 方法的返回值就是类的实例对象，这个实例对象会传递给 __init__ 方法中定义的 self 参数，以便实例对象可以被正确地初始化
	    # 如果 __new__ 方法不返回值（或者说返回 None）那么 __init__ 将不会得到调用（因为实例对象都没创建出来）
	    return self

    def __call__(self):  # 可以定义任意参数
        print('__call__ ')

A() 
# 依次输出
# __new__ 
# <__main__.A object at 0x1007a95f8>
# __init__ 
# <__main__.A object at 0x1007a95f8>
```
一般地，我们不会去重写 \_\_new\_\_ 方法，除非你确切知道怎么做，什么时候你会去关心它呢，它作为构造函数用于创建对象，是一个工厂函数，专用于生产实例对象。著名的设计模式之一 —— 单例模式，就可以通过此方法来实现。  
```python
class BaseController(object):
    _singleton = None
    def __new__(cls, *a, **k):
        if not cls._singleton:
            cls._singleton = object.__new__(cls, *a, **k)
        return cls._singleton
```
自定义的函数、内置函数和类都属于可调用对象（callable）。  
凡是可以把一对括号 () 应用到某个对象身上都可称之为可调用对象，判断对象是否为可调用对象可以用函数 callable。  
\_\_call\_\_ 方法就是在调用可调用对象时触发的方法。  
```python
a = A()
print(callable(a))  # True
a()  # __call__
```

10、\*args 和 \*\*kwargs
\*args 和 \*\*kwargs 主要⽤于函数定义，可以将不定数量的参数（预先并不知道, 函数使⽤者会传递多少个参数）传递给⼀个函数。  
\*args 是⽤来发送⼀个⾮键值对的可变数量的参数列表给⼀个函数。  
\*\*kwargs 允许你将不定长度的键值对, 作为参数传递给⼀个函数。   
标准参数与 \*args、\*\*kwargs 在使⽤时的顺序 func(fargs, \*args, \*\*kwargs)。  
```python
def test_var_args(f_arg, *argv): 
    print("first normal arg:", f_arg)
    for arg in argv:
        print("another arg through *argv:", arg)


test_var_args('yasoob', 'python', 'eggs', 'test')
# first normal arg: yasoob
# another arg through *argv: python 
# another arg through *argv: eggs 
# another arg through *argv: test


def greet_me(**kwargs):
    for key, value in kwargs.items(): 
        print("{0} == {1}".format(key, value))


greet_me(name="yasoob") 
# name == yasoob


def test_args_kwargs(arg1, arg2, arg3): 
    print("arg1:", arg1) 
    print("arg2:", arg2) 
    print("arg3:", arg3)


# ⾸先使⽤ *args
args = ("two", 3, 5)
test_args_kwargs(*args)
# 现在使⽤ **kwargs:
kwargs = {"arg3": 3, "arg2": "two", "arg1": 5} 
test_args_kwargs(**kwargs)
```

11、调试（Debugging）  
利⽤好调试，能⼤⼤提⾼你捕捉代码 Bug，[参考连接](https://docs.python.org/3/library/pdb.html)。  

- 从命令⾏运⾏
在命令⾏使⽤ Python debugger 运⾏⼀个脚本，它会触发 debugger 在脚本的第⼀⾏指令处停⽌执⾏。这在脚本很短时会很有帮助，你可以通过(Pdb)模式接着查看变量信息，并且逐⾏调试。  
```bash
python -m pdb my_script.py
```

- 从脚本内部运⾏
在脚本内部设置断点，这样就可以在某些特定点查看变量信息和各种执⾏时信息了，这⾥将使⽤ pdb.set_trace() ⽅法来实现。  
```python
import pdb


def make_bread(): 
    pdb.set_trace()
    return "I don't have time"


print(make_bread())
```

12、对象变动(Mutation)  
Python中可变(mutable)与不可变(immutable)的数据类型让新⼿很是头痛。简单的说，可变(mutable)意味着"可以被改动"，⽽不可变(immutable)的意思是“常量(constant)”。  
```python
foo = ['hi']
print(foo)
# Output: ['hi']

# 每当你将⼀个变量赋值为另⼀个可变类型的变量时，对这个数据的任意改动会同时反映到这两个变量上去
# 新变量只不过是⽼变量的⼀个别名⽽已
bar = foo
bar += ['bye']
print(foo)
# Output: ['hi', 'bye']

# 在Python中当函数被定义时，默认参数只会运算⼀次，⽽不是每次被调⽤时都会重新运算
def add_to(num, target=[]):
	target.append(num)
	return target
add_to(1)
# Output: [1]
add_to(2)
# Output: [2]
add_to(3)
# Output: [3]

# 永远不要定义可变类型的默认参数，应该这样做
def add_to(element, target=None):
	if target is None:
		target = []
		target.append(element)
	
	return target
# 每当你在调⽤这个函数不传⼊ target 参数的时候，⼀个新的列表会被创建
add_to(42)
# Output: [42]
add_to(42)
# Output: [42]
add_to(42)
# Output: [42]
```


999、Python 三大利器：迭代器，生成器，装饰器  
#### 1 寻找第n次出现位置

```python
def search_n(s, c, n):
    size = 0
    for i, x in enumerate(s):
        if x == c:
            size += 1
        if size == n:
            return i
    return -1



print(search_n("fdasadfadf", "a", 3))# 结果为7，正确
print(search_n("fdasadfadf", "a", 30))# 结果为-1，正确
```


#### 2 斐波那契数列前n项

```python
def fibonacci(n):
    a, b = 1, 1
    for _ in range(n):
        yield a
        a, b = b, a + b


list(fibonacci(5))  # [1, 1, 2, 3, 5]
```

#### 3 找出所有重复元素

```python
from collections import Counter


def find_all_duplicates(lst):
    c = Counter(lst)
    return list(filter(lambda k: c[k] > 1, c))


find_all_duplicates([1, 2, 2, 3, 3, 3])  # [2,3]
```

#### 4 联合统计次数
Counter对象间可以做数学运算

```python
from collections import Counter
a = ['apple', 'orange', 'computer', 'orange']
b = ['computer', 'orange']

ca = Counter(a)
cb = Counter(b)
#Counter对象间可以做数学运算
ca + cb  # Counter({'orange': 3, 'computer': 2, 'apple': 1})


# 进一步抽象，实现多个列表内元素的个数统计


def sumc(*c):
    if (len(c) < 1):
        return
    mapc = map(Counter, c)
    s = Counter([])
    for ic in mapc: # ic 是一个Counter对象
        s += ic
    return s


#Counter({'orange': 3, 'computer': 3, 'apple': 1, 'abc': 1, 'face': 1})
sumc(a, b, ['abc'], ['face', 'computer'])

```

#### 5 groupby单字段分组

天气记录：

```python
a = [{'date': '2019-12-15', 'weather': 'cloud'},
 {'date': '2019-12-13', 'weather': 'sunny'},
 {'date': '2019-12-14', 'weather': 'cloud'}]
```

按照天气字段`weather`分组汇总：

```python
from itertools import groupby
for k, items in  groupby(a,key=lambda x:x['weather']):
     print(k)
```

输出结果看出，分组失败！原因：分组前必须按照分组字段`排序`，这个很坑~

```python
cloud
sunny
cloud
```

修改代码：

```python
a.sort(key=lambda x: x['weather'])
for k, items in  groupby(a,key=lambda x:x['weather']):
     print(k)
     for i in items:
         print(i)
```

输出结果：

```python
cloud
{'date': '2019-12-15', 'weather': 'cloud'}
{'date': '2019-12-14', 'weather': 'cloud'}
sunny
{'date': '2019-12-13', 'weather': 'sunny'}
```

#### 6 itemgetter和key函数

注意到`sort`和`groupby`所用的`key`函数，除了`lambda`写法外，还有一种简写，就是使用`itemgetter`：

```python
a = [{'date': '2019-12-15', 'weather': 'cloud'},
 {'date': '2019-12-13', 'weather': 'sunny'},
 {'date': '2019-12-14', 'weather': 'cloud'}]
from operator import itemgetter
from itertools import groupby

a.sort(key=itemgetter('weather'))
for k, items in groupby(a, key=itemgetter('weather')):
     print(k)
     for i in items:
         print(i)
```

结果：

```python
cloud
{'date': '2019-12-15', 'weather': 'cloud'}
{'date': '2019-12-14', 'weather': 'cloud'}
sunny
{'date': '2019-12-13', 'weather': 'sunny'}
```

#### 7 groupby多字段分组

`itemgetter`是一个类，`itemgetter('weather')`返回一个可调用的对象，它的参数可有多个：

```python
from operator import itemgetter
from itertools import groupby

a.sort(key=itemgetter('weather', 'date'))
for k, items in groupby(a, key=itemgetter('weather')):
     print(k)
     for i in items:
         print(i)
```

结果如下，使用`weather`和`date`两个字段排序`a`，

```python
cloud
{'date': '2019-12-14', 'weather': 'cloud'}
{'date': '2019-12-15', 'weather': 'cloud'}
sunny
{'date': '2019-12-13', 'weather': 'sunny'}
```

注意这个结果与上面结果有些微妙不同，这个更多是我们想看到和使用更多的。

#### 8 sum函数计算和聚合同时做

Python中的聚合类函数`sum`,`min`,`max`第一个参数是`iterable`类型，一般使用方法如下：

```python
a = [4,2,5,1]
sum([i+1 for i in a]) # 16
```

使用列表生成式`[i+1 for i in a]`创建一个长度与`a`一行的临时列表，这步完成后，再做`sum`聚合。

试想如果你的数组`a`长度十百万级，再创建一个这样的临时列表就很不划算，最好是一边算一边聚合，稍改动为如下：

```python
a = [4,2,5,1]
sum(i+1 for i in a) # 16
```

此时`i+1 for i in a`是`(i+1 for i in a)`的简写，得到一个生成器(`generator`)对象，如下所示：

```python
In [8]:(i+1 for i in a)
OUT [8]:<generator object <genexpr> at 0x000002AC7FFA8CF0>
```

生成器每迭代一步吐出(`yield`)一个元素并计算和聚合后，进入下一次迭代，直到终点。

#### 9 list分组(生成器版)

```python
from math import ceil

def divide_iter(lst, n):
    if n <= 0:
        yield lst
        return
    i, div = 0, ceil(len(lst) / n)
    while i < n:
        yield lst[i * div: (i + 1) * div]
        i += 1

list(divide_iter([1, 2, 3, 4, 5], 0))  # [[1, 2, 3, 4, 5]]
list(divide_iter([1, 2, 3, 4, 5], 2))  # [[1, 2, 3], [4, 5]]
```

#### 10 列表全展开（生成器版）
```python
#多层列表展开成单层列表
a=[1,2,[3,4,[5,6],7],8,["python",6],9]
def function(lst):
    for i in lst:
        if type(i)==list:
            yield from function(i)
        else:
            yield i
print(list(function(a))) # [1, 2, 3, 4, 5, 6, 7, 8, 'python', 6, 9]
```

#### 11 测试函数运行时间的装饰器
```python
#测试函数执行时间的装饰器示例
import time
def timing_func(fn):
    def wrapper():
        start=time.time()
        fn()   #执行传入的fn参数
        stop=time.time()
        return (stop-start)
    return wrapper
@timing_func
def test_list_append():
    lst=[]
    for i in range(0,100000):
        lst.append(i)  
@timing_func
def test_list_compre():
    [i for i in range(0,100000)]  #列表生成式
a=test_list_append()
c=test_list_compre()
print("test list append time:",a)
print("test list comprehension time:",c)
print("append/compre:",round(a/c,3))

test list append time: 0.0219423770904541
test list comprehension time: 0.007980823516845703
append/compre: 2.749
```

#### 12 统计异常出现次数和时间的装饰器


写一个装饰器，统计某个异常重复出现指定次数时，经历的时长。
```python
import time
import math


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

关键词`nonlocal`常用于函数嵌套中，声明变量i为非局部变量；

如果不声明，`i+=1`表明`i`为函数`wrapper`内的局部变量，因为在`i+=1`引用(reference)时,`i`未被声明，所以会报`unreferenced variable`的错误。

使用创建的装饰函数`excepter`, `n`是异常出现的次数。

共测试了两类常见的异常：`被零除`和`数组越界`。

```python
n = 10 # except count

@excepter
def divide_zero_except():
    time.sleep(0.1)
    j = 1/(40-20*2)

# test zero divived except
for _ in range(n):
    divide_zero_except()


@excepter
def outof_range_except():
    a = [1,3,5]
    time.sleep(0.1)
    print(a[3])
# test out of range except
for _ in range(n):
    outof_range_except()

```

打印出来的结果如下：
```python
division by zero: 1
division by zero: 2
division by zero: 3
division by zero: 4
division by zero: 5
division by zero: 6
division by zero: 7
division by zero: 8
division by zero: 9
division by zero: 10
spending time:1.01
list index out of range: 1
list index out of range: 2
list index out of range: 3
list index out of range: 4
list index out of range: 5
list index out of range: 6
list index out of range: 7
list index out of range: 8
list index out of range: 9
list index out of range: 10
spending time:1.01
```


#### 13 测试运行时长的装饰器


```python
#测试函数执行时间的装饰器示例
import time
def timing(fn):
    def wrapper():
        start=time.time()
        fn()   #执行传入的fn参数
        stop=time.time()
        return (stop-start)
    return wrapper

@timing
def test_list_append():
    lst=[]
    for i in range(0,100000):
        lst.append(i)  

@timing
def test_list_compre():
    [i for i in range(0,100000)]  #列表生成式

a=test_list_append()
c=test_list_compre()
print("test list append time:",a)
print("test list comprehension time:",c)
print("append/compre:",round(a/c,3))

# test list append time: 0.0219
# test list comprehension time: 0.00798
# append/compre: 2.749
```



#### 14 装饰器通俗理解

再看一个装饰器：

```python
def call_print(f):
  def g():
    print('you\'re calling %s function'%(f.__name__,))
  return g
```

使用`call_print`装饰器：

```python
@call_print
def myfun():
  pass
 
@call_print
def myfun2():
  pass
```

myfun()后返回：

```python
In [27]: myfun()
you're calling myfun function

In [28]: myfun2()
you're calling myfun2 function
```

**使用call_print**

你看，`@call_print`放置在任何一个新定义的函数上面，都会默认输出一行，你正在调用这个函数的名。

这是为什么呢？注意观察新定义的`call_print`函数(加上@后便是装饰器):

```python
def call_print(f):
  def g():
    print('you\'re calling %s function'%(f.__name__,))
  return g
```

它必须接受一个函数`f`，然后返回另外一个函数`g`.

**装饰器本质**

本质上，它与下面的调用方式效果是等效的：

```
def myfun():
  pass

def myfun2():
  pass
  
def call_print(f):
  def g():
    print('you\'re calling %s function'%(f.__name__,))
  return g
```

下面是最重要的代码：

```
myfun = call_print(myfun)
myfun2 = call_print(myfun2)
```

大家看明白吗？也就是call_print(myfun)后不是返回一个函数吗，然后再赋值给myfun.

再次调用myfun, myfun2时，效果是这样的：

```python
In [32]: myfun()
you're calling myfun function

In [33]: myfun2()
you're calling myfun2 function
```

你看，这与装饰器的实现效果是一模一样的。装饰器的写法可能更加直观些，所以不用显示的这样赋值：`myfun = call_print(myfun)`，`myfun2 = call_print(myfun2)`，但是装饰器的这种封装，猛一看，有些不好理解。

#### 15 定制递减迭代器

```python
#编写一个迭代器，通过循环语句，实现对某个正整数的依次递减1，直到0.
class Descend(Iterator):
    def __init__(self,N):
        self.N=N
        self.a=0
    def __iter__(self):
        return self 
    def __next__(self):
        while self.a<self.N:
            self.N-=1
            return self.N
        raise StopIteration
    
descend_iter=Descend(10)
print(list(descend_iter))
[9, 8, 7, 6, 5, 4, 3, 2, 1, 0]
```

核心要点：

1 `__nex__ `名字不能变，实现定制的迭代逻辑

2 `raise StopIteration`：通过 raise 中断程序，必须这样写


### Python 之坑
#### 1 含单个元素的元组

Python中有些函数的参数类型为元组，其内有1个元素，这样创建是错误的：

```python
c = (5) # NO!
```

它实际创建一个整型元素5，必须要在元素后加一个`逗号`:

```python
c = (5,) # YES!
```

#### 2 默认参数设为空

含有默认参数的函数，如果类型为容器，且设置为空：

```python
def f(a,b=[]):  # NO!
    print(b)
    return b

ret = f(1)
ret.append(1)
ret.append(2)
# 当再调用f(1)时，预计打印为 []
f(1)
# 但是却为 [1,2]
```

这是可变类型的默认参数之坑，请务必设置此类默认参数为None：

```python
def f(a,b=None): # YES!
    pass
```

#### 3 共享变量未绑定之坑

有时想要多个函数共享一个全局变量，但却在某个函数内试图修改它为局部变量：

```python
i = 1
def f():
    i+=1 #NO!
    
def g():
    print(i)
```

应该在f函数内显示声明`i`为global变量：

```python
i = 1
def f():
    global i # YES!
    i+=1
```

#### 4 lambda自由参数之坑

排序和分组的key函数常使用lambda，表达更加简洁，但是有个坑新手容易掉进去：

```python
a = [lambda x: x+i for i in range(3)] # NO!
for f in a:
    print(f(1))
# 你可能期望输出：1,2,3
```

但是实际却输出: 3,3,3. 定义lambda使用的`i`被称为自由参数，它只在调用lambda函数时，值才被真正确定下来，这就犹如下面打印出2，你肯定确信无疑吧。

```python
a = 0
a = 1
a = 2
def f(a):
    print(a)
```

正确做法是转化`自由参数`为lambda函数的`默认参数`：

```python
a = [lambda x,i=i: x+i for i in range(3)] # YES!
```

#### 5 各种参数使用之坑

Python强大多变，原因之一在于函数参数类型的多样化。方便的同时，也为使用者带来更多的约束规则。如果不了解这些规则，调用函数时，可能会出现如下一些语法异常：

*(1) SyntaxError: positional argument follows keyword argument*


*(2) TypeError: f() missing 1 required keyword-only argument: 'b'*


*(3) SyntaxError: keyword argument repeated*

*(4) TypeError: f() missing 1 required positional argument: 'b'*

*(5) TypeError: f() got an unexpected keyword argument 'a'*

*(6) TypeError: f() takes 0 positional arguments but 1 was given*


总结主要的参数使用规则

位置参数

`位置参数`的定义：`函数调用`时根据函数定义的参数位（形参）置来传递参数，是最常见的参数类型。

```python
def f(a):
  return a

f(1) # 位置参数 
```
位置参数不能缺少：
```python
def f(a,b):
  pass

f(1) # TypeError: f() missing 1 required positional argument: 'b'
```

**规则1：位置参数必须一一对应，缺一不可**

关键字参数

在函数调用时，通过‘键--值’方式为函数形参传值，不用按照位置为函数形参传值。

```python
def f(a):
  print(f'a:{a}')
```
这么调用，`a`就是关键字参数：
```python
f(a=1)
```
但是下面调用就不OK:
```python
f(a=1,20.) # SyntaxError: positional argument follows keyword argument
```

**规则2：关键字参数必须在位置参数右边**


下面调用也不OK:
```python
f(1,width=20.,width=30.) #SyntaxError: keyword argument repeated

```

**规则3：对同一个形参不能重复传值**


默认参数

在定义函数时，可以为形参提供默认值。对于有默认值的形参，调用函数时如果为该参数传值，则使用传入的值，否则使用默认值。如下`b`是默认参数：
```python
def f(a,b=1):
  print(f'a:{a}, b:{b}')

```


**规则4：无论是函数的定义还是调用，默认参数的定义应该在位置形参右面**

只在定义时赋值一次；默认参数通常应该定义成不可变类型


可变位置参数

如下定义的参数a为可变位置参数：
```python
def f(*a):
  print(a)
```
调用方法：
```python
f(1) #打印结果为元组： (1,)
f(1,2,3) #打印结果：(1, 2, 3)
```

但是，不能这么调用：
```python
f(a=1) # TypeError: f() got an unexpected keyword argument 'a'
```


可变关键字参数

如下`a`是可变关键字参数：
```python
def f(**a):
  print(a)
```
调用方法：
```python
f(a=1) #打印结果为字典：{'a': 1}
f(a=1,b=2,width=3) #打印结果：{'a': 1, 'b': 2, 'width': 3}
```

但是，不能这么调用：
```python
f(1) TypeError: f() takes 0 positional arguments but 1 was given
```

接下来，单独推送分析一个小例子，综合以上各种参数类型的函数及调用方法，敬请关注。

#### 6 列表删除之坑

删除一个列表中的元素，此元素可能在列表中重复多次：

```python
def del_item(lst,e):
    return [lst.remove(i) for i in e if i==e] # NO!
```

考虑删除这个序列[1,3,3,3,5]中的元素3，结果发现只删除其中两个：

```python
del_item([1,3,3,3,5],3) # 结果：[1,3,5]
```

正确做法：

```python
def del_item(lst,e):
    d = dict(zip(range(len(lst)),lst)) # YES! 构造字典
    return [v for k,v in d.items() if v!=e]

```

#### 7 列表快速复制之坑

在python中`*`与列表操作，实现快速元素复制：

```python
a = [1,3,5] * 3 # [1,3,5,1,3,5,1,3,5]
a[0] = 10 # [10, 2, 3, 1, 2, 3, 1, 2, 3]
```

如果列表元素为列表或字典等复合类型：

```python
a = [[1,3,5],[2,4]] * 3 # [[1, 3, 5], [2, 4], [1, 3, 5], [2, 4], [1, 3, 5], [2, 4]]

a[0][0] = 10 #  
```

结果可能出乎你的意料，其他`a[1[0]`等也被修改为10

```python
[[10, 3, 5], [2, 4], [10, 3, 5], [2, 4], [10, 3, 5], [2, 4]]
```

这是因为*复制的复合对象都是浅引用，也就是说id(a[0])与id(a[2])门牌号是相等的。如果想要实现深复制效果，这么做：

```python
a = [[] for _ in range(3)]
```

#### 8 字符串驻留
```python
In [1]: a = 'something'
    ...: b = 'some'+'thing'
    ...: id(a)==id(b)
Out[1]: True
```
如果上面例子返回`True`，但是下面例子为什么是`False`:
```python
In [1]: a = '@zglg.com'

In [2]: b = '@zglg'+'.com'

In [3]: id(a)==id(b)
Out[3]: False
```
这与Cpython 编译优化相关，行为称为`字符串驻留`，但驻留的字符串中只包含字母，数字或下划线。

#### 9 相同值的不可变对象
```python
In [5]: d = {}
    ...: d[1] = 'java'
    ...: d[1.0] = 'python'

In [6]: d
Out[6]: {1: 'python'}

### key=1,value=java的键值对神器消失了
In [7]: d[1]
Out[7]: 'python'
In [8]: d[1.0]
Out[8]: 'python'
```
这是因为具有相同值的不可变对象在Python中始终具有`相同的哈希值`

由于存在`哈希冲突`，不同值的对象也可能具有相同的哈希值。

#### 10 对象销毁顺序
创建一个类`SE`:
```python
class SE(object):
  def __init__(self):
    print('init')
  def __del__(self):
    print('del')
```
创建两个SE实例，使用`is`判断：
```python
In [63]: SE() is SE()
init
init
del
del
Out[63]: False

```
创建两个SE实例，使用`id`判断：
```python
In [64]: id(SE()) == id(SE())
init
del
init
del
Out[64]: True
```

调用`id`函数, Python 创建一个 SE 类的实例，并使用`id`函数获得内存地址后，销毁内存丢弃这个对象。

当连续两次进行此操作, Python会将相同的内存地址分配给第二个对象，所以两个对象的id值是相同的.


但是is行为却与之不同，通过打印顺序就可以看到。

#### 11 充分认识for
```python
In [65]: for i in range(5):
    ...:   print(i)
    ...:   i = 10
0
1
2
3
4
```
为什么不是执行一次就退出？

按照for在Python中的工作方式, i = 10 并不会影响循环。range(5)生成的下一个元素就被解包，并赋值给目标列表的变量`i`.

#### 12 认识执行时机

```python
array = [1, 3, 5]
g = (x for x in array if array.count(x) > 0)
```
`g`为生成器，list(g)后返回`[1,3,5]`，因为每个元素肯定至少都出现一次。所以这个结果这不足为奇。但是，请看下例：
```python
array = [1, 3, 5]
g = (x for x in array if array.count(x) > 0)
array = [5, 7, 9]
```
请问,list(g)等于多少？这不是和上面那个例子结果一样吗，结果也是`[1,3,5]`，但是：
```python
In [74]: list(g)
Out[74]: [5]
```

这有些不可思议~~ 原因在于：

生成器表达式中, in 子句在声明时执行, 而条件子句则是在运行时执行。


所以代码：
```python
array = [1, 3, 5]
g = (x for x in array if array.count(x) > 0)
array = [5, 7, 9]
```

等价于：
```python
g = (x for x in [1,3,5] if [5,7,9].count(x) > 0)
```



#### 13 创建空集合错误

这是Python的一个集合：`{1,3,5}`，它里面没有重复元素，在去重等场景有重要应用。下面这样创建空集合是错误的：

```python
empty = {} #NO!
```

cpython会解释它为字典

使用内置函数`set()`创建空集合：

```python
empty = set() #YES!
```

