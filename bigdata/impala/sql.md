
### Impala 基本 crud
```bash

```

### Impala 导入导出
```bash

```

### Impala 常用内置函数
```bash
# 统计
max(field_name)
min(field_name)
avg(field_name)

# 取绝对值
abs(field_name/numeric_var);

# 四舍五入
round(field_name/numeric_var)

# 幂运算
power(numeric_var, power_number)

# 类型转换
# 整形转为字符型
cast(10 as string)

# 字符串拼接
concat(string a, string b...) 

# 集合查找
find_in_set(string str, string str_list) 

# 当前时间戳
now()

# 日期增加和减少
date_add(string start_date, int days)   
date_sub(string start_date, int days)   

# 日期比较
datediff(string end_date, string start_date) 
```
