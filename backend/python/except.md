
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
