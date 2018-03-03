
### Impala 基本 crud
```bash
# 创建数据库
CREATE (DATABASE|SCHEMA) [IF NOT EXISTS]database_name[COMMENT 'database_comment']
  [LOCATION hdfs_path];
# 使用数据库
USE database_name;
# 删除数据库
DROP (DATABASE|SCHEMA) [IF EXISTS] database_name;

# 创建表
CREATE [EXTERNAL] TABLE [IF NOT EXISTS] [db_name.]table_name[(col_name data_type[COMMENT 'col_comment'], ...)]
  [COMMENT 'table_comment']
  [PARTITIONED BY (col_name data_type[COMMENT 'col_comment'], ...)]
  [
   [ROW FORMAT row_format] [STORED AS file_format]
  ]
  [LOCATION 'hdfs_path']
  [WITH SERDEPROPERTIES ('key1'='value1', 'key2'='value2', ...)]
  [TBLPROPERTIES ('key1'='value1', 'key2'='value2', ...)];
# row_format：DELIMITED [FIELDS TERMINATED BY 'char' [ESCAPED BY 'char']] [LINES TERMINATED BY 'char']
# file_format: PARQUET | PARQUETFILE | TEXTFILE | SEQUENCEFILE | RCFILE

# 克隆表（指定 hdfs）
CREATE [EXTERNAL] TABLE [IF NOT EXISTS] [db_name.]table_name 
  LIKE [db_name.]table_name
  [COMMENT 'table_comment']
  [STORED ASfile_format]
  [LOCATION 'hdfs_path'];
# 克隆表（查询其他表）
CREATE [EXTERNAL] TABLE [IF NOT EXISTS]db_name.]table_name
  [COMMENT 'table_comment']
  [STORED ASfile_format]
  [LOCATION 'hdfs_path']
AS 
  select_statement;

# 重命名表
ALTER TABLE table_name RENAME TO new_table_name;

# 删除表
DROP TABLE [IF EXISTS] [db_name.]table_name

# 显示表的元数据
DESCRIBE [FORMATTED] table_name;

# 修改表的数据物理位置
ALTER TABLE table_name SET LOCATION 'hdfs_path_of_directory';

# 修改表的文件格式
ALTER TABLE table_name SET FILEFORMAT { PARQUET | PARQUETFILE | TEXTFILE | RCFILE | SEQUENCEFILE };

# 添加字段
ALTER TABLE table_name ADD COLUMNS (column_defs);
# 重新定义字段
ALTER TABLE table_name REPLACE COLUMNS (column_defs);
# 更新字段
ALTER TABLE table_name CHANGE column_name new_name new_spec;
# 删除字段
ALTER TABLE table_name DROP column_name;

# 分析查询语句
EXPLAIN select_sentence;
```

### Impala 导入导出
```bash
# impala 仅支持导入 hdfs 文件
LOAD DATA INPATH 'hdfs_file_or_directory_path' [OVERWRITE] INTO TABLE table_name
[PARTITION (partcol1=val1,partcol2=val2...)]
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

### 自定义函数
```bash
# 执行 CREATE FUNCTION 语句创建标量函数
CREATE FUNCTION [IF NOT EXISTS] [db_name.]function_name([arg_type[,arg_type...])
  RETURNS return_type LOCATION 'hdfs_path'
  SYMBOL='symbol_or_class'

# 执行 CREATE AGGREGATE FUNCTION 语句创建 UDA
CREATE [AGGREGATE] FUNCTION [IF NOT EXISTS] [db_name.]function_name([arg_type[,arg_type...])
  RETURNSreturn_typeLOCATION 'hdfs_path'
  [INIT_FN='function]
  UPDATE_FN='function MERGE_FN='function [FINALIZE_FN='function]
```
