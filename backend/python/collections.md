
### 容器(Collections)
Python 附带⼀个模块，它包含许多容器数据类型，名字叫作 collections。  
collections 主要包含 defaultdict、counter、deque、namedtuple、enum.Enum。  

**defaultdict**    
与 dict 类型不同，你不需要检查 key 是否存在。  
```python
from collections import defaultdict

colours = (
	('Yasoob', 'Yellow'),
	('Ali', 'Blue'),
	('Arham', 'Green'),
	('Ali', 'Black'),
	('Yasoob', 'Red'),
	('Ahmed', 'Silver'),
)
favourite_colours = defaultdict(list)

for name, colour in colours:
	favourite_colours[name].append(colour)

print(favourite_colours)
# defaultdict(<type 'list'>,
# {'Arham': ['Green'],
# 'Yasoob': ['Yellow', 'Red'],
# 'Ahmed': ['Silver'],
# 'Ali': ['Blue', 'Black']
# })
```

当你在⼀个字典中对⼀个键进⾏嵌套赋值时，如果这个键不存在，会触发 keyError 异常。 defaultdict 允许我们⽤⼀个聪明的⽅式绕过这个问题。
```python
some_dict = {}
some_dict['colours']['favourite'] = "yellow"
## 异常输出：KeyError: 'colours'

import collections
import json

tree = lambda: collections.defaultdict(tree)
some_dict = tree()
some_dict['colours']['favourite'] = "yellow"

print(json.dumps(some_dict))
# {"colours": {"favourite": "yellow"}}
```

**counter**  
Counter是⼀个计数器，它可以帮助我们针对某项数据进⾏计数。  
```python
from collections import Counter

# 计算每个⼈喜欢多少种颜⾊
colours = (
	('Yasoob', 'Yellow'),
	('Ali', 'Blue'),
	('Arham', 'Green'),
	('Ali', 'Black'),
	('Yasoob', 'Red'),
	('Ahmed', 'Silver'),
)
favs = Counter(name for name, colour in colours)
print(favs)
# Counter({
# 'Yasoob': 2,
# 'Ali': 2,
# 'Arham': 1,
# 'Ahmed': 1
# })

# 统计⽂件
with open('filename', 'rb') as f:
    line_count = Counter(f)

print(line_count)
```

**deque**  
deque 提供了⼀个双端队列，你可以从头/尾两端添加或删除元素。  
```python
from collections import deque

d = deque()
d.append('1')
d.append('2')
d.append('3')
print(len(d)) # 3
print(d[0]) # 1
print(d[-1]) # 3

d = deque(range(5))
print(len(d)) # 5
d.popleft() # 0
d.pop() # 4
print(d) # deque([1, 2, 3])

# 也可以限制这个列表的⼤⼩，当超出你设定的限制时，数据会从对队列另⼀端被挤出去(pop)
d = deque(maxlen=30) # 当你插⼊ 30 条数据时，最左边⼀端的数据将从队列中删除
d = deque([1,2,3,4,5])
d.extendleft([0])
d.extend([6,7,8])
print(d) # deque([0, 1, 2, 3, 4, 5, 6, 7, 8])
```

**namedtuple**  
⼀个元组是⼀个不可变的列表，你可以存储⼀个数据的序列，它和命名元组(namedtuples)⾮常像，但有⼏个关键的不同。  

主要相似点是都不像列表，你不能修改元组中的数据。  
namedtuples 把元组变成⼀个针对简单任务的容器。你不必使⽤整数索引来访问⼀个 namedtuples 的数据，你可以像字典(dict)⼀样访问 namedtuples，但 namedtuples 是不可变的。  
namedtuple 让你的元组变得⾃⽂档，不必使⽤整数索引来访问⼀个命名元组，这让你的代码更易于维护；namedtuple 的每个实例没有对象字典，所以它们很轻量，与普通的元组⽐，并不需要更多的内存，这使得它们⽐字典更快。  
`注意：要记住它是⼀个元组，属性值在 namedtuple 中是不可变的`

```python
man = ('Ali', 30)
print(man[0]) # Ali

from collections import namedtuple

Animal = namedtuple('Animal', 'name age type') # 元组名称，字段名称（元组名称是 Animal，字段名称是 'name'，'age' 和 'type'）
perry = Animal(name="perry", age=31, type="cat")
print(perry) # Animal(name='perry', age=31, type='cat')
print(perry.name) # perry

perry.age = 42
# Traceback (most recent call last):
# File "", line 1, in
# AttributeError: can't set attribute

# 以将⼀个命名元组转换为字典
print(perry._asdict()) #  OrderedDict([('name', 'Perry'), ('age', 31), ...
```

**enum.Enum**  
枚举对象属于 enum 模块，存在于 Python 3.4 以上版本中（同时作为⼀个独⽴的PyPI包enum34供⽼版本使⽤）。  
Enums(枚举类型)基本上是⼀种组织各种东西的⽅式。
```python
from collections import namedtuple
from enum import Enum

class Species(Enum):
	cat = 1
	dog = 2
	horse = 3
	aardvark = 4
	butterfly = 5
	owl = 6
	platypus = 7
	dragon = 8
	unicorn = 9
	# 依次类推

	# 但我们并不想关⼼同⼀物种的年龄，所以我们可以使⽤⼀个别名
	kitten = 1 # (译者注：幼⼩的猫咪)
	puppy = 2 # (译者注：幼⼩的狗狗)

Animal = namedtuple('Animal', 'name age type')
perry = Animal(name="Perry", age=31, type=Species.cat)
drogon = Animal(name="Drogon", age=4, type=Species.dragon)
tom = Animal(name="Tom", age=75, type=Species.cat)
charlie = Animal(name="Charlie", age=2, type=Species.kitten)

print(charlie.type == tom.type) # True
charlie.type # <Species.cat: 1>

# 以下⽅法都可以获取到'cat'的值：
Species(1)
Species['cat']
Species.cat
```

