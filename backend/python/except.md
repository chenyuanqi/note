
### Python 异常
良好的异常处理可以让你的程序更加健壮，清晰的错误信息更能帮助你快速修复问题。  
Python 使用 try/except/finally 语句块来处理异常。  

```python
def div(a, b):
    try:
        print(a / b)
    except ZeroDivisionError:
        print("Error: b should not be 0 !!")
    except Exception as e:
        print("Unexpected Error: {}".format(e))
    else:
        print('Run into else only when everything goes well')
    finally:
        print('Always run into finally block.')

# tests
div(2, 0)
div(2, 'bad type')
div(1, 2)

# Mutiple exception in one line
try:
    print(a / b)
except (ZeroDivisionError, TypeError) as e:
    print(e)

# Except block is optional when there is finally
try:
    open(database)
finally:
    close(database)

# catch all errors and log it
try:
    do_work()
except:    
    # get detail from logging module
    logging.exception('Exception caught!')
    
    # get detail from sys.exc_info() method
    error_type, error_value, trace_back = sys.exc_info()
    print(error_value)
    raise
```

使用 Python 异常处理，  
1、except 语句不是必须的，finally 语句也不是必须的，但是二者必须要有一个，否则就没有 try 的意义了  
2、except 语句可以有多个，Python 会按 except 语句的顺序依次匹配你指定的异常，如果异常已经处理就不会再进入后面的 except 语句  
3、except 语句可以以元组形式同时指定多个异常  
4、except 语句后面如果不指定异常类型，则默认捕获所有异常，你可以通过 logging 或者 sys 模块获取当前异常  
5、如果要捕获异常后要重复抛出，请使用 raise，后面不要带任何参数或信息  
6、不建议捕获并抛出同一个异常，请考虑重构你的代码  
7、不建议在不清楚逻辑的情况下捕获所有异常，有可能你隐藏了很严重的问题  
8、尽量使用内置的异常处理语句来替换 try/except 语句，比如 with 语句，getattr() 方法  

**Python 异常处理格式**  
```python
try:
    ......
except XXXError as XXX:
    ......
else:
    ......
finally:
    ......
```

**抛出异常 raise**  
如果你需要自主抛出异常一个异常，可以使用 raise 关键字。  
raise 关键字后面可以指定你要抛出的异常实例，一般来说抛出的异常越详细越好。  
```python
raise NameError("bad name!")

# 查看 exceptions 中的异常类型
import exceptions
print dir(exceptions)
# ['ArithmeticError', 'AssertionError'...]
```

**自定义异常类型**  
Python 中自定义自己的异常类型非常简单，只需要要从 Exception 类继承即可 (直接或间接)：
```python
class XXXError(Exception):
    def __init__(self, *args, **kwargs):
        ...
    def logger(self):
        ...

# 手动触发异常
raise XXXError()

# 断言
assert py_express

class SomeCustomException(Exception):
    pass

class AnotherException(SomeCustomException):
    pass
```
一般在自定义异常类型时，需要考虑的问题应该是这个异常所应用的场景。  
如果内置异常已经包括了你需要的异常，建议考虑使用内置的异常类型。比如你希望在函数参数错误时抛出一个异常，你可能并不需要定义一个 InvalidArgumentError，使用内置的 ValueError 即可。  

**Exception 和 BaseException**  
当我们要捕获一个通用异常时，应该用 Exception 还是 BaseException？我建议你还是看一下 官方文档说明，这两个异常到底有啥区别呢？ 请看它们之间的继承关系。Python 异常层次：  

BaseException  
&nbsp;&nbsp;&nbsp;&nbsp;+-- SystemExit  
&nbsp;&nbsp;&nbsp;&nbsp;+-- KeyboardInterrupt  
&nbsp;&nbsp;&nbsp;&nbsp;+-- GeneratorExit  
&nbsp;&nbsp;&nbsp;&nbsp;+-- Exception  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;+-- StopIteration  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;+-- ArithmeticError  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;+-- AssertionError  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;+-- AttributeError  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;+-- BufferError  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;+-- EOFError  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;+-- ImportError  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;+-- LookupError  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;+-- MemoryError  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;+-- NameError  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;+-- OSError  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;+-- ReferenceError  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;+-- RuntimeError  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;+-- SyntaxError  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;+-- SystemError  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;+-- TypeError  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;+-- ValueError  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;+-- Warning  

从 Exception 的层级结构来看，BaseException 是最基础的异常类，Exception 继承了它。BaseException 除了包含所有的 Exception 外还包含了 SystemExit，KeyboardInterrupt 和 GeneratorExit 三个异常。  
在捕获所有异常时更应该使用 Exception 而不是 BaseException，因为被排除的三个异常属于更高级别的异常，合理的做法应该是交给 Python 的解释器处理。  

**except Exception as e 和 except Exception, e**  
```python
try:
    do_something()
except NameError as e:  # should
    pass
except KeyError, e:  # should not
    pass
```
在 Python2 的时代，你可以使用以上两种写法中的任意一种。在 Python3 中你只能使用第一种写法，第二种写法已经不再支持。第一个种写法可读性更好，而且为了程序的兼容性和后期移植的成本，请你果断抛弃第二种写法。

**raise "Exception string"**  
把字符串当成异常抛出看上去是一个非常简洁的办法，但其实是一个非常不好的习惯。  
```python
if is_work_done():
    pass
else:
    raise "Work is not done!" # not cool

# 抛出异常会是这样的
# Traceback (most recent call last):
#   File "/demo/exception_hanlding.py", line 48, in <module>
#     raise "Work is not done!"
# TypeError: exceptions must be old-style classes or derived from BaseException, not str
```
这在 Python2.4 以前是可以接受的做法，但是没有指定异常类型有可能会让下游没办法正确捕获并处理这个异常，从而导致你的程序难以维护。简单说，这种写法是是封建时代的陋习，应该扔了。  

**使用内置的语法范式代替 try/except**  
Python 本身提供了很多的语法范式简化了异常的处理，比如 for 语句就处理了的 StopIteration 异常，让你很流畅地写出一个循环。  

with 语句在打开文件后会自动调用 finally 并关闭文件。我们在写 Python 代码时应该尽量避免在遇到这种情况时还使用 try/except/finally 的思维来处理。  
```python
# should not
try:
    f = open(a_file)
    do_something(f)
finally:
    f.close()

# should 
with open(a_file) as f:
    do_something(f)
```

当我们需要访问一个不确定的属性时，
```python
try:
    test = Test()
    name = test.name  # not sure if we can get its name
except AttributeError:
    name = 'default'

# 可以使用更简单的 getattr() 来达到你的目的
name = getattr(test, 'name', 'default')
```

### Python 异常最佳实践
1、只处理你知道的异常，避免捕获所有异常然后吞掉它们。  
2、抛出的异常应该说明原因，有时候你知道异常类型也猜不出所以然。  
3、避免在 catch 语句块中干一些没意义的事情，捕获异常也是需要成本的。  
4、不要使用异常来控制流程，那样你的程序会无比难懂和难维护。  
5、如果有需要，切记使用 finally 来释放资源。  
6、如果有需要，请不要忘记在处理异常后做清理工作或者回滚操作。  

### Python 错误及解决办法
1、SyntaxError 语法错误  
（1）引号没有成对出现：SyntaxError:EOL while scanning string literal  
```python
string = 'hello world
```
字符串切记要放在引号中，单引号双引号无所谓。  
当一个字符串中包含单引号或双引号时，很容易出现引号不配对的情况。  

（2）圆括号没有成对出现：SyntaxError:unexpected EOF while parsing  
```python
result = (1024+(512*2)/128
# 或者
print('hello world'
```
使圆括号成对出现。  
在书写复杂的表达式或调用函数时会经常遇到这个问题。  

（3）错误使用自操作运算符 ++ 或 -- 等：SyntaxError:invalid syntax  
```python
v = 64
v++ # 正确示例：v += 1
```
在 Python 语言中，没有类似 C 语言的 ++ 或 -- 等自操作运算符。  
与之类似功能的用法是 += 或 -= 运算符。  

（4）试图使用等号（=）判断两个运算量是否相等：SyntaxError:invalid syntax  
```python
if v=64:
    print('hello world')
```
在 Python 语言中,使用两个等号（==）作为判断两个运算量是否相等的关系运算符，而等号（=）是赋值运算符。  

（5）错误使用 Python 语言关键字作为变量名：SyntaxError: can't assign to keyword  
```python
False = 1
```
不要使用 Python 语言关键字作为变量名、函数名或类名等。  
在 Python Shell 窗口中，使用 help('keywords') 指令可以查看 Python 语言的关键字列表。  

（6）忘记在 if/elif/else/while/for/def/class 等语句末尾添加冒号（:）：SyntaxError:invalid syntax  
```python
a = '12345'
for i  in a
    print(i)
```
在 if/elif/else/while/for/def/class 等语句末尾添加冒号（:）即可。  
牢记语法规则，多多练习多多敲代码。  

（7）错误地使用了中文标点符号：SyntaxError: invalid character in identifier  
```python
print('hello'，'world')
# 错误原因：逗号是中文标点符号

for i in range(10)：
# 错误原因：冒号是中文标点符号
```
除了字符串中可以有中文外，其它任何情况均使用英文状态进行编辑。  

2、IndentationError 缩进错误  
报错信息：IndentationError：unindent does not match any outer indentation level  
或者：IndentationError：expected an indented block  
```python
a = 2
while a < 0:
        print('hello')
    a -= 1
else:
    print('0.0')
```
上述代码中 while 语句体内的代码缩进没有对齐。  
正确使用缩进排版代码。当代码是从其它地方复制并粘贴过来的时候，这个错误较多见。  

3、NameError 名字错误  
当变量名、函数名或类名等书写错误，或者函数在定义之前就被调用等情况下，就会导致名字错误。  
报错信息：NameError: name 'pirnt' is not defined  
```python
pirnt('hello world')
# 错误原因：print 拼写错误

sayhi()
def sayhi():
    pass
# 错误原因：在函数定义之前对函数进行调用

pd.read_excel(r'file.xlsx')
# 错误原因：在调用 pandas 方法前并未导入 pandas 库或者并未起别名为 pd
```
正确书写变量名、函数名或类名等，在使用变量前先进行赋值，将函数的定义放在函数调用之前，在使用第三方库前先进行导入、调包等等。即保证某个名字（标识符）先存在，才能被使用。  

4、TypeError 类型错误  
（1）整数和字符串不能进行连接操作：TypeError: Can't convert 'int' object to str implicitly 或 TypeError: unsupported operand type(s) for + : 'float' and 'str'  
```python
print('score:'+100)
print(9.8 + 'seconds')
```
在整数、浮点数或布尔值与字符串进行连接操作之前，先使用str()函数将其转换为字符串类型。  

（2）调用函数时参数的个数不正确，或者未传递参数：TypeError: input expected at most 1 arguments,got 2 或 TypeError: say() missing 1 required positional argument:'words'  
```python
input('输入姓名','年龄')
# 错误原因：试图给 input() 函数提供第 2 个参数

def say(words):
    print(words)
say()
# 错误原因：调用函数时未传递参数
```
记住函数用法，了解函数的参数定义，使用正确的方法调用函数即可。  

5、KeyError 键错误  
使用不存在的键名访问字典中的元素，就会发生这个错误。  
报错信息：KeyError: 'c'  
```python
d = {'a':1,'b':2}
print(d['c'])
```
在访问字典中的元素时，先用 in 关键字检测要访问的键名是否存在，或者是使用字典和 get() 方法安全地访问字典元素。  

6、 IndexError 索引错误  
当访问列表的索引超出列表范围时，就会出现索引错误。  
报错信息：IndexError: list index out of range  
```python
a = [1,2,3]
print(a[3])
# 错误原因：列表 a 中不存在第 4 个索引。列表的索引从 0 开始编号。
```
通过len()函数获取列表的长度，然后判断要访问的索引是否超出列表范围。  

7、UNboundLocalError 未初始化本地变量错误  
在函数中，如果对未声明的全局变量进行修改操作，将会遇到这个错误。  
报错信息：UnboundLocalError: local variable 's' referenced before assignment  
```python
s = 1

def test():
    s += 1
    print(s)

test()
# 错误原因：在函数内对未声明的全局变量 s 进行了自增操作
# Python 将变量 s 视为一个本地的局部变量，但该变量未初始化
```
在函数内使用全局变量时，使用 global 关键字对其进行声明即可。  

8、AttributeError 属性错误  
报错信息：AttributeError: 'tuple' object has no attribute 'append' 或 AttributeError: 'DataFrame' object has no attribute 'col'  
```python
t = (1,2,3)
t.append(4)
# 错误原因：元祖不可变

df = pd.read_excel(r'data.xlsx')
df.col
# 错误原因：DataFrame 没有 col 属性，应该为 columns
```
正确书写类的属性名，不要发生书写错误。  
深刻理解元祖，列表的区别，可将元祖转换为列表添加元素。  

9、ModuleNotFoundError 模块不存在  
报错信息：ModuleNotFoundError: No module named 'pandas'  
```python
import pandas as pd
# 没有导入成功，报上面错误
```
这种报错常见于两种场景中，  
第一、未下载、安装该模块；  
第二、将调用的模块路径与被调用的模块路径不一致等。  
第一种情况直接下载安装即可，在cmd中，pip install xxx；第二种情况电脑中可能存在多个版本的 Python，建议保留一个常用的即可。  

10、FileNotFoundError 文件不存在  
报错信息：FileNotFoundError: File b'E:\test\test_data.csv' does not exist  
```python
pd.read_csv('E:\test\test_data.csv')
# 错误原因：路径中包含'\t'，系统错误地认为是制表符
```
在确保该路径下确实存在所写文件后，在读取文件路径前面加'r'，表示只读，作为文件路径读取；或者使用双斜杠'\ \'来进行转义，形如：'E:\ \test\ \test_data.csv'。偶尔也会发生文件名、路径的确写错，犯一些低级错误。  
