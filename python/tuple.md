
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

# 获取元组元素：tuple_var[offset]
print(tuple_var[0], tuple_var[-1]) # 返回 1 'tuple'
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
