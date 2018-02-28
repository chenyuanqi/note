
### Hive Sql 与 传统 Sql 的对比
Hive 对分号没有那么智能，需要转为八进制的 Ascii 码  
Hive 不支持 INSERT INTO table_name Values(), UPDATE, DELETE 操作  
Hive 不支持将数据插入现有的表或分区中，仅支持覆盖重写整个表（INSERT OVERWRITE TABLE）  
HiveQL 的空字符串 IS NULL 的判断结果是 False。
Hive 中，SELECT table_name.* FROM table_name JOIN another_table_name ON (table_name.id = another_table_name.id)；  
传统 SQL 中，SELECT table_name.* FROM table_name, another_table_name WHERE table_name.id = another_table_name.id。  


### 数据查询
```bash
# 基本查询格式
SELECT [ALL | DISTINCT] select_expr, select_expr, ...
FROM table_reference
[WHERE where_condition]
[GROUP BY col_list [HAVING condition]]
[CLUSTER BY col_list] | [DISTRIBUTE BY col_list]  
[SORT BY| ORDER BY col_list]
[LIMIT number]

# join 连接查询
# 仅支持等值 join
SELECT table_name.* FROM table_name JOIN another_table_name ON (table_name.id = another_table_name.id)

# UNION ALL 联合查询
# 需保证 SELECT 字段一致
select_statement UNION ALL select_statement UNION ALL select_statement...
```

### 使用 Py 脚本
```python
import sys
import datetime

for line in sys.stdin:
    line = line.strip()
    user_id, movie_id, rating, unix_time = line.split('\t')
    week_day = datetime.datetime.fromtimestamp(float(unix_time)).isoweekday()
    print('\t'.join([user_id, movie_id, rating, str(week_day)]))
```

```bash
# 添加脚本
ADD FILE python_file_path;
# 写入数据
INSERT OVERWRITE TABLE table_name
SELECT
  TRANSFORM (field_name)
  USING 'python_file_path'
  AS (python_file_field)
FROM another_table_name;
```

### 内置函数及 UDF
```bash
# 查看所有内置函数
show functions;
# 查看内置函数的用法
describe function function_name;
```
