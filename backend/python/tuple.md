
### 元组常用方法
```python
# 元组，初看时就像是不可变的列表
# 元组，再看时 ...

# 元组定义及初始化
tuple_var = () # 或者 tuple_var = tuple()
# 需要注意的是，元组只有一个元素时逗号不可省略，tuple_var = (1,)

# 元组的连接
tuple_var = (1,) + ("tuple",) # 返回 (1, 'tuple')

# 元组的复制
tuple_var_copy = tuple_var[:]
# 元组的所有元素复制
tuple_var * 2 # 返回 (1, 'tuple', 1, 'tuple')

# 获取元组元素：tuple_var[offset]
tuple_var[0], tuple_var[-1] # 返回 1 'tuple'
# 需要注意的是，offset 大于列表长度 - 1 时会抛 IndexError 异常

# 元组切片，格式 tuple_var[start:end:step]
tuple_var_split = tuple_var[-1:] # ('tuple',)
# 需要注意的是，step 默认为 1，而且当切片没有任何元素时，返回 ()

# 判断指定值是否在元组中
1 in tuple_var # 返回 True

# 获取指定值的位置
tuple_var.index(1) # 返回 0

# 统计指定值出现的次数
tuple_var.count(1) # 返回 1
# 获取元组的元素个数
len(tuple_var) # 返回 2
# 获取元组的元素最大值
another_tuple_var = (1, 3, 5)
max(another_tuple_var) # 返回 5
# 获取列表元素最小值
min(another_tuple_var) # 返回 1

# 删除元组，删除后无法访问
del another_tuple_var

# 打包多个序列到元组
list_var_1 = ['name', 'age']
list_var_2 = ['python', '28']
tuple_var_zip = zip(list_var_1, list_var_2) # [('name', 'python'), ('age', 28)]
# 解包
zip(*tuple_var_zip) # [('name', 'age'), ('python', '28')]

# 再来看看元组连接，不改变任何原元组
another_tuple_var = (1, 3, 5)
tuple_var + another_tuple_var # 返回 (1, 'tuple', 1, 3, 5)

# 元组中含可变元素时，如包含列表
tuple_include_list_var = (1, [1, 3], 5)
tuple_include_list_var[1][0] = 0 # tuple_include_list_var 修改为 (1, [0, 3], 5)
```

### 为什么需要元组？
> 在 python 2.x 或以下版本中，字符串格式化中参数要用元组；  
> 在字典中，元组可以当作键使用而列表不可以；  
> 元组可以作为函数的多个返回值 ...  
> 元组的不可变性，提供了一种完整的约束。

### 元组也可变？
```python
tuple_var = (1, 2, 3, 4, 5)
tuple_var = tuple_var[:1] + (666,) + tuple_var[1:]
# 利用切片修改元组，原元组内存地址也会被修改
# 这相当于重新赋值给了元组而已
```

### 说说元组和列表的区别
相同点：  
list 与 tuple 都是序列类型的容器对象，可以存放任何类型的数据、支持切片、迭代等操作。  

不同点：  
两种类型除了字面上的区别(括号与方括号)之外，最重要的一点是tuple是不可变类型，大小固定，而 list 是可变类型、数据可以动态变化，这种差异使得两者提供的方法、应用场景、性能上都有很大的区别。  
同样大小的数据，初始化和迭代 tuple 都要快于 list；  
同样大小的数据，tuple 占用的内存空间更少；  
原子性的 tuple 对象还可作为字典的键。  

tuple 用于存储异构(heterogeneous)数据，当做没有字段名的记录来用，比如用 tuple 来记录一个人的身高、体重、年龄。  
而列表一般用于存储同构数据(homogenous)，同构数据就是具有相同意义的数据。  
因为 tuple 作为没有名字的记录来使用在某些场景有一定的局限性，所以又有了一个 namedtuple 类型的存在，namedtuple 可以指定字段名，用来当做一种轻量级的类来使用。  
