
### 字典常用方法

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
