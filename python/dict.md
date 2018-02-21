
### 字典常用方法
```python
# 定义及初始化
dict_var = {} # 或者 dict_var = dict()

# 添加或更新字典项
dict_var["name"] = "vikey"

# 查看键是否在字典里
print("name" in dict_var)

# 查看字典里的所有键
dict_var.keys() # 返回 dict_keys(['name'])

# 查看字典里的所有值
dict_var.values() # 返回 dict_values(['vikey'])

# 查看字典里的所有项
dict_var.items() # 返回 dict_items([('name', 'vikey')])

# 字典浅拷贝
dict_var_copy = dict_var.copy()

# 清空字典
dict_var.clear()
# del dict_var["name"] 可以删除字典的某一项

# 从拷贝的字典更新字典所有项，没有键则添加
dict_var.update(dict_var_copy)

# 设置字典键的值，如果键不存在，将添加并设为默认值
dict_var.setdefault("age", 18) # 返回键的值

# 获取字典键的值，如果键不存在，返回默认值
dict_var.get("age", 28)
# dict_var["age"] 也可以，但是键不存在会抛 keyError 异常

# 删除字典指定键并返回对应的值，没有则返回默认值
dict_var.pop("age", 0) # 返回 18

# 从序列中生成字典，不改变原字典（没有默认值，则默认 None）
dict_var.fromkeys(["sex", "birthday"], "") # 返回 {'sex': '', 'birthday': ''}

# 获取字典项目总数
len(dict_var) # 返回 2
```

### 判断键的存在
```python
# 不推荐
dict_var.has_key(key)
# 推荐
key in dict_var
```

### 关于 not
```python
# 不推荐
not key in dict_var
# 推荐
key not in dict_var
```

### 关于取值
```python
# 不推荐
if key not in dict_var:
  dict_var[key] = value
dict_var[key] = dict_var[key] + 1
# 推荐
dict_var[key] = dict_var.get(key, value) + 1
```

### 关于使用默认值
```python
# 不推荐
dict_var = {}
for (key, value) in data:
    if key in dict_var:
        dict_var[key].append(value)
    else:
        dict_var[key] = [value]
# 推荐
dict_var = defaultdict(list)
for (key, value) in data:
    dict_var[key].append(value)
```
