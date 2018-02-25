
### 集合常用方法
```python
# 集合是无序不重复的数据集合
# 集合访问速度快，天然解决重复问题

# 定义及初始化
set_var = set() # 注意，不可以使用 {} 来初始化集合，因为这个已经被 dict 占有辣

# 添加元素
set_var.add('set_001') # set_var 修改为 {'set_001'}
# 批量添加元素
set_var.update([1, 2, 3]) # set_var 修改为 {1, 2, 3, 'set_001'}

# 集合的长度
len(set_var) # 返回 4

# 集合的更新
# 由于 set 存储的是一组不重复的无序元素，更新操作实际分两步
# ① 删除旧元素
set_var.discard(3)
# ② 新元素添加到集合中
set_var.add(4)


# 集合的元素删除
set_var.remove('set_001') # 元素不存在，抛 KeyError 异常
set_var.discard('set_001') # 元素不存在，不报错
# 移除最后一个元素（由于集合无序，移除的元素不确定）
set_var.pop() # 集合为空，抛 KeyError 异常

# 集合复制
set_var_copy = set_var.copy() 

# 清空集合，原集合变成 set()
set_var.clear()
```

### 集合与集合间的关系
```python
set_var_1 = {1, 2, 3, 'set'}
set_var_2 = {2, 3, 'set'}
# 集合的并集或合集
set_var_union = set_var_1.union(set_var_2) # {1, 2, 3, 'set'}
# 等价于 set_var_union = set_var_1 | set_var_2

# 判断集合是否没有交值
set_var_1.isdisjoint(set_var_2) # 有交值返回 False, 没有交值返回 True

# 集合的交值
set_var_intersection = set_var_1.intersection(set_var_2) # {'set', 2, 3}
# 等价于 set_var_intersection = set_var_1 & set_var_2
# 取集合的交值，并更新到集合里
set_var_1.intersection_update(set_var_2) # set_var_1 被修改为 set_var_intersection 的值

# 判断集合是否为另一集合的子集
set_var_1.issubset(set_var_2) # set_var_1 不是 set_var_2 的子集，所以返回 False
# 判断集合是否为另一集合的父集
set_var_1.issuperset(set_var_2) # set_var_1 是 set_var_2 的父集，所以返回 True

# 集合的差集或相对补值
set_var_difference = set_var_1.difference(set_var_2) # {1}
# 等价于 set_var_difference = set_var_1 - set_var_2
# 取集合的差值，并更新到集合里
set_var_1.difference_update(set_var_2) # set_var_1 被修改为 set_var_difference 的值

# 集合的对称差值，即不相交的部分
set_var_symmetric_difference = set_var_1.symmetric_difference(set_var_2) # {1}
# 等价于 set_var_symmetric_difference = set_var_1 ^ set_var_2
# 取集合的对称差值，并更新到集合里
set_var_1.symmetric_difference_update(set_var_2) # set_var_1 被修改为 set_var_symmetric_difference 的值

```

### frozenset 是什么
frozenset 即不可变集合，存在 hash 值，可作为 dict 的键，也可作为其他集合的键；  
frozenset 没有 add、update、remove、discard、pop、clear、intersection_update、difference_update、symmetric_difference_update 等方法；  
frozenset 的定义：frozenset_var = frozenset({1, 2, 3})  

set 或者 frozenset 都不支持创建整数的集合，如 set_var = set(1)。