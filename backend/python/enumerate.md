
### 枚举 Enumerate
枚举(enumerate)是 Python 内置函数，它允许我们遍历数据并⾃动计数。  
```python
for counter, value in enumerate(some_list):
    print(counter, value)

# enumerate 也接受⼀些可选参数，如第二个蚕食是定制从哪个数字开始枚举
my_list = ['apple', 'banana', 'grapes', 'pear']
for c, value in enumerate(my_list, 1):
    print(c, value)
```
