

### 官方网站
[Python](https://www.python.org/)  

### IDE
[Pycharm](https://www.jetbrains.com/pycharm/)  

### 文档
[官方文档](https://docs.python.org/3/)  
[中文文档](http://www.pythondoc.com/)  

### 代码规范
[官方风格](https://www.python.org/dev/peps/pep-0008/)  
[Google 风格](http://zh-google-styleguide.readthedocs.io/en/latest/google-python-styleguide/) 

### 代码检测工具
[语法检查 pyflakes](https://pypi.python.org/pypi/pyflakes)  
[代码规范 pylint](https://www.pylint.org/)  
> 工具同样可以聚合在 pycharm  

### 框架
[Django](https://www.djangoproject.com/)  
[Django 最佳实践](https://github.com/yangyubo/zh-django-best-practices/blob/master/readme.rst)   

[Flask](http://flask.pocoo.org/)  
[Flask 之旅](http://spacewander.github.io/explore-flask-zh/) 

### 标准库
[2.x 标准库](https://docs.python.org/2/library/)  
[2.x how to](https://docs.python.org/2/howto/)  
[3.x 标准库](https://docs.python.org/3/library/)  
[3.x how to](https://docs.python.org/3/howto/)  

### 第三方库
[兼容性工具 Six](http://pythonhosted.org/six/)  
[HTTP 请求操作库 requests](https://github.com/kennethreitz/requests)  
[URL 处理器 purl](https://github.com/codeinthehole/purl)  
[清晰友好的 HTTP 库 urllib3](https://github.com/shazow/urllib3)  
[文件系统处理库 path.py](https://github.com/jaraco/path.py)  
[数据格式检查库 schema](https://github.com/halst/schema)  
[函数式增强库 fn.py](https://github.com/kachayev/fn.py)  
[时间日期库 when.py](https://github.com/dirn/When.py)  
[更好的时间日期库 arrow](https://github.com/crsmithdev/arrow)  
[命令行参数解析库 docopt](https://github.com/docopt/docopt)  
[代码风格处理库 autopep8](https://github.com/hhatto/autopep8)  

### 爬虫
[HTML\XML 搜索修改库 bs4](https://www.crummy.com/software/BeautifulSoup/bs4/doc/)  
[快速高级的屏幕爬取及网页采集框架 scrapy](https://github.com/scrapy/scrapy)  

### 数据分析
[科学计算基础包 Numpy](http://www.numpy.org/)  
[2D 绘图库 matplotlib](https://matplotlib.org/)  
[数据分析工具 Pandas](http://pandas.pydata.org/)  

### 深入理解  
[Python 格式化](https://pyformat.info/)  
[Python 挑战](http://www.pythonchallenge.com/)  
[Python 之旅](http://funhacks.net/explore-python/)  
[Python 进阶](http://interpy.eastlakeside.com/)  
[Python 最佳实践](http://pythonguidecn.readthedocs.io/zh/latest/)  
[Python 装饰器](https://zhuanlan.zhihu.com/p/65968462)  
[Python 设计模式](https://github.com/faif/python-patterns)  
[Python 代码分析](http://www.pythontutor.com/)  
[Python 数据分析](https://github.com/BrambleXu/pydata-notebook)  
[stackoverflow 中文问题集](https://taizilongxu.gitbooks.io/stackoverflow-about-python/content/index.html)  
[Python 罕见问题](http://norvig.com/python-iaq.html)  

### 其他
[Python 之禅](https://foofish.net/)  
[Python 茶馆](https://pythoncaff.com/)  
[爱湃森 2017 榜单](https://annual2017.pycourses.com/)  
[爱湃森 2018 榜单](https://annual2018.pycourses.com/)  
[爱湃森 2019 榜单](https://annual2019.pycourses.com/)  

------
[廖雪峰 Python](http://www.liaoxuefeng.com/wiki/0014316089557264a6b348958f449949df42a6d3a2e542c000)  
[莫烦 Python](https://morvanzhou.github.io/)  

------
[更多 awesome](https://awesome-python.com/)  

### python 之禅
-- The Zen of Python, by Tim Peters  

Beautiful is better than ugly.  
Explicit is better than implicit.  
Simple is better than complex.  
Complex is better than complicated.  
Flat is better than nested.  
Sparse is better than dense.  
Readability counts.  
Special cases aren't special enough to break the rules.  
Although practicality beats purity.  
Errors should never pass silently.  
Unless explicitly silenced.  
In the face of ambiguity, refuse the temptation to guess.  
There should be one-- and preferably only one --obvious way to do it.  
Although that way may not be obvious at first unless you're Dutch.  
Now is better than never.  
Although never is often better than *right* now.  
If the implementation is hard to explain, it's a bad idea.  
If the implementation is easy to explain, it may be a good idea.  
Namespaces are one honking great idea -- let's do more of those!  

-- 中文版  
优美胜于丑陋（Python 以编写优美的代码为目标）  
明了胜于晦涩（优美的代码应当是明了的，命名规范，风格相似）  
简洁胜于复杂（优美的代码应当是简洁的，不要有复杂的内部实现）  
复杂胜于凌乱（如果复杂不可避免，那代码间也不能有难懂的关系，要保持接口简洁）  
扁平胜于嵌套（优美的代码应当是扁平的，不能有太多的嵌套）  
间隔胜于紧凑（优美的代码有适当的间隔，不要奢望一行代码解决问题）  
可读性很重要（优美的代码是可读的）  
即便假借特例的实用性之名，也不可违背这些规则（这些规则至高无上）  
不要包容所有错误，除非你确定需要这样做（精准地捕获异常，不写 except:pass 风格的代码）  
当存在多种可能，不要尝试去猜测  
而是尽量找一种，最好是唯一一种明显的解决方案（如果不确定，就用穷举法）  
虽然这并不容易，因为你不是 Python 之父（这里的 Dutch 是指 Guido ）  
做也许好过不做，但不假思索就动手还不如不做（动手之前要细思量）  
如果你无法向人描述你的方案，那肯定不是一个好方案；反之亦然（方案测评标准）  
命名空间是一种绝妙的理念，我们应当多加利用（倡导与号召）  

### 这样的 Python 很优雅
1、变量交换  
```python
a, b = b, a
```

2、循环遍历区间元素  
```python
# xrange 返回的是生成器对象，生成器比列表更加节省内存
# 不过需要注意的是 xrange 是 python2 中的写法，python3 只有 range 方法，特点和 xrange 是一样的
for i in range(6):
    print(i)
```

3、带有索引位置的集合遍历  
遍历集合时如果需要使用到集合的索引位置时，直接对集合迭代是没有索引信息的。  
```python
colors = ['red', 'green', 'blue', 'yellow']
for i, color in enumerate(colors):
    print(i, '--->', color)
```

4、字符串连接  
字符串连接时，普通的方式可以用 + 操作  
```python
names = ['raymond', 'rachel', 'matthew']
# 使用 + 操作时，每执行一次 + 操作就会导致在内存中生成一个新的字符串对象
# 使用 join 方法整个过程只会产生一个字符串对象
print(', '.join(names))
```

5、打开 / 关闭文件  
执行文件操作时，最后一定不能忘记的操作是关闭文件，即使报错了也要 close。  
```python
with open('test.txt') as f:
    data = f.read()
```

6、列表推导式  
能够用一行代码简明扼要地解决问题时，绝不要用两行。   
```python
[i for i in range(10)]
```

7、善用装饰器  
装饰器可以把与业务逻辑无关的代码抽离出来，让代码保持干净清爽，而且装饰器还能被多个地方重复利用。  
```python
import urllib.request as urllib

def cache(func):
    saved = {}

    def wrapper(url):
        if url in saved:
            return saved[url]
        else:
            page = func(url)
            saved[url] = page
            return page

    return wrapper

@cache
def web_lookup(url):
	# 一个爬虫网页的函数，如果该 URL 曾经被爬过就直接从缓存中获取，否则爬下来之后加入到缓存，防止后续重复爬取
    return urllib.urlopen(url).read()
```

8、合理使用列表  
列表对象（list）是一个查询效率高于更新操作的数据结构，比如删除一个元素和插入一个元素时执行效率就非常低，因为还要对剩下的元素进行移动。  
```python
from collections import deque

# deque 是一个双向队列的数据结构，删除元素和插入元素会很快
names = deque(['raymond', 'rachel', 'matthew', 'roger'])
names.popleft()
names.appendleft('mark')
```

9、序列解包  
```python
p = 'vttalk', 'female', 30, 'python@qq.com'
name, gender, age, email = p
```

10、遍历字典的 key 和 value  
```python
# 方法一
# 速度没那么快，因为每次迭代的时候还要重新进行 hash 查找 key 对应的 value
for k in d:
    print k, '--->', d[k]

# 方法二
# Python2 遇到字典非常大的时候，会导致内存的消耗增加一倍以上；iteritems 返回迭代器对象，可节省更多的内存
# Python3 的 items 等值于 iteritems
for k, v in d.items():
    print k, '--->', v
```

11、链式比较操作  
```python
age = 18
if 18 < age < 60:
    print("yong man")
```

12、if/else 三目运算  
```python
# 能够用 if/else 清晰表达逻辑时，就没必要再额外新增一种方式来实现
text = '男' if gender == 'male' else '女'
```

13、真值判断  
检查某个对象是否为真值时，还显示地与 True 和 False 做比较就显得多此一举，不专业。  
```python
if attr: # attr == True
    do_something()

if values: # len(values) != 0
    do_something()
```

14、for/else 语句  
for else 是 Python 中特有的语法格式，else 中的代码在 for 循环遍历完所有元素之后执行。  
```python
for i in mylist:
    if i == theflag:
        break
    process(i)
else:
    raise ValueError("List argument missing terminal flag.")
```

15、字符串格式化  
很难说用 format 比用 % s 的代码量少，但是 format 更易于理解。  
```python
s1 = "foofish.net"
s2 = "vttalk"
# s3 = "welcome to %s and following %s" % (s1, s2)
s3 = "welcome to {blog} and following {wechat}".format(blog="foofish.net", wechat="vttalk")
```

16、列表切片  
获取列表中的部分元素最先想到的就是用 for 循环根据条件提取元素，这也是其它语言中惯用的手段，而在 Python 中还有强大的切片功能。  
`注意：列表元素的下标不仅可以用正数表示，还是用负数表示，最后一个元素的位置是 -1，从右往左，依次递减。`
```python
items = range(10)
# 第1到第4个元素的范围区间
sub_items = items[1:4]
# 奇数
odd_items = items[1::2]
# 拷贝
copy_items = items[::] # 或者 items[:]
```

17、善用生成器  
生成器的好处就是无需一次性把所有元素加载到内存，只有迭代获取元素时才返回该元素，而列表是预先一次性把全部元素加载到了内存。此外用 yield 代码看起来更清晰。  
```python
# 用生成器生成费波那契数列
def fib(n):
    a, b = 0, 1
    while a < n:
        yield a
        a, b = b, a + b

for i in fib(9):
    print(i)
```

18、获取字典元素  
```python
d = {'name': 'foo'}
# 字典 key 不存在，设置默认值
d.get("name", "unknow")
```

19、预设字典默认值  
通过 key 分组的时候，不得不每次检查 key 是否已经存在于字典中。  
```python
data = [('foo', 10), ('bar', 20), ('foo', 39), ('bar', 49)]
groups = {}
for (key, value) in data:
    if key in groups:
        groups[key].append(value)
    else:
        groups[key] = [value]

#　第一种方式
groups = {}
for (key, value) in data:
    groups.setdefault(key, []).append(value) 

# 第二种方式
from collections import defaultdict
groups = defaultdict(list)
for (key, value) in data:
    groups[key].append(value)
```

20、字典推导式  
字典推导式是 python2.7 新增的特性，可读性增强了很多，类似的还是列表推导式和集合推导式。  
```python
numbers = [1, 2, 3]
my_dict = {number: number * 2 for number in numbers}
print(my_dict)  # {1: 2, 2: 4, 3: 6}

# 还可以指定过滤条件
my_dict = {number: number * 2 for number in numbers if number > 1}
print(my_dict)  # {2: 4, 3: 6}
```
