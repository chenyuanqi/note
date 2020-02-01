
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
