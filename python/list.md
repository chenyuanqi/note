

### 列表常用方法
```python
# 列表定义及初始化
list_var = [] # 或者 list_var = list()

# 切割字符串
list_var = "1,2,3".split(",") # 返回 ['1', '2', '3']

# 获取列表元素：list_var[offset]
print(list_var[1], list_var[-1]) # 返回 2 3
# 需要注意的是，offset 大于列表长度 - 1 时会抛 IndexError 异常

# 修改列表元素
list_var[2] = "4" # ['1', '2', '4']

# 列表尾部添加元素
list_var.append("4") # ['1', '2', '4', '4']
# 指定位置添加元素，若指定位置不存在则尾部添加
list_var.insert(2, "3") # ['1', '2', '3', '4', '4']

# 列表切片，格式 list_var[start:end:step]
list_var_split = list_var[-2:] # ['4', '4']
# 需要注意的是，step 默认为 1，而且当切片没有任何元素时，返回 []

# 判断指定值是否在列表中
"1" in list_var_split # 返回 False

# 获取指定值的位置
list_var.index("2") # 返回 1

# 统计指定值出现的次数
list_var.count("4") # 返回 2
# 获取列表元素个数
len(list_var) # 返回 5
# 获取列表元素最大值
max(list_var) # 返回 '4'
# 获取列表元素最小值
min(list_var) # 返回 '1'

# 列表反转
list_var.reverse() # list_var 修改为 ['4', '4', '3', '2', '1']
# 列表默认升序排序
list_var.sort() # list_var 修改为 ['1', '2', '3', '4', '4']
# 改成降序排序
list_var.sort(reverse=True) # list_var 修改为 ['4', '4', '3', '2', '1'] 
# 不修改原列表的排序处理
sorted(list_var) # 返回 ['1', '2', '3', '4', '4']

# 删除列表指定下标元素
list_var.pop(1) # 返回 '4'
# 需要注意的是，pop 默认 offset 为 -1，而且当 offset 不存在时抛 IndexError 异常
# 删除指定值的元素，多个则删除第一个出现的元素
list_var.remove("4") # 没有返回值，list_var 修改为 ['3', '2', '1'] 
# 删除列表指定下标的元素
del list_var_split[1] # list_var_split 修改为 ['4']
# 删除列表，删除后无法访问
del list_var_split

# 复制列表(浅拷贝)
another_list_var = list_var[:]
# 或者 another_list_var = list_var.copy()

# 列表连接，不改变任何原列表
list_var + another_list_var # 返回 ['3', '2', '1', '3', '2', '1']

# 合并列表
list_var += another_list_var # list_var 修改为 ['3', '2', '1', '3', '2', '1']
# 还原 list_var，使用 extend 合并
list_var = another_list_var
list_var.extend(another_list_var) # list_var 修改为 ['3', '2', '1', '3', '2', '1']
# 两种合并方式的区别：extend 把 another_list_var 放入 list_var 中，原列表各自内存地址不变；而 += 会把列表相加后重新赋值给 list_var，即内存地址改变

# 删除列表重复项
list(set(list_var)) # list_var 不改变，返回 ['1', '3', '2']

# 列表转换成字符串
",".join(list_var) # list_var 不改变，返回 '3,2,1,3,2,1'
```

### 遍历列表，同时输出键和值
```python
list_var = [1, 2, 3]
for index, value in enumerate(list_var):
    print(index, value)
```

### 列表生成式
```python
dict_var = {'x': 'A', 'y': 'B', 'z': 'C' }
# 不使用列表生成式时
for key, value in dict_var.items():
    print(key, value)
# 使用列表生成式时
[print(key, value) for key, value in dict_var.items()]
```
