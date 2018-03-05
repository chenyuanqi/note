
### Python 异常层次
BaseException  
    +-- SystemExit  
    +-- KeyboardInterrupt  
    +-- GeneratorExit  
    +-- Exception  
        +-- StopIteration  
        +-- ArithmeticError  
        +-- AssertionError  
        +-- AttributeError  
        +-- BufferError  
        +-- EOFError  
        +-- ImportError  
        +-- LookupError  
        +-- MemoryError  
        +-- NameError  
        +-- OSError  
        +-- ReferenceError  
        +-- RuntimeError  
        +-- SyntaxError  
        +-- SystemError  
        +-- TypeError  
        +-- ValueError  
        +-- Warning  

### Python 异常处理格式
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

### Python 自定义异常
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
```
