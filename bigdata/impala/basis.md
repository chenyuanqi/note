
### 什么是 Impala
Impala 是性能最高的 SQL 引擎（提供类似RDBMS的体验），它提供了访问存储在 Hadoop 分布式文件系统中的数据（HDFS 或 HBase）的最快方法。  
Impala 是开源的本地分析数据库。  
Impala 不需要对存储在 Hadoop 上的数据进行数据转换和数据移动。  
与 Hive 不同，Impala 不基于 MapReduce 算法。 Impala 实现了一个基于守护进程的分布式架构，它负责在同一台机器上运行的查询执行的所有方面。因此，Impala 减少了使用MapReduce 的延迟，比 Hive 更快。  
Impala 使用 Hive 的元数据，ODBC 驱动程序和 SQL 语法。  
Impala 支持各种文件格式，如 LZO，序列文件，Avro，RCFile 以及 Parquet。  
Impala 不支持更新或删除单个记录、事务、索引、触发器等，它不提供任何对序列化和反序列化的支持；Impala 只能读取文本文件，而不能读取自定义二进制文件；每当新的记录/文件被添加到HDFS中的数据目录时，该表需要被刷新。    


### Impala 原理

### Impala 安装
1、在 [Impala Download Page](http://impala.apache.org/downloads.html) 中选择下载一个稳定版本的压缩包，然后解压到大数据目录。
```bash
# 下载，解压
wget http://archive.apache.org/dist/impala/x.x.x/apache-impala.x.x.x.tar.gz
tar -xzvf apache-impala.x.x.x.tar.gz
# 添加环境变量
export IMPALA_HOME=/impala_path
export PATH=$HIVE_HOME/bin: $PATH
```

### Impala shell
1、执行 impala-shell 进入命令行  
```bash
# 启动
impala-shell

# 获取可用的命令列表
help;

# 查看版本
version;

# 查看执行命令历史(默认 10 条)
history;

# 退出
exit; # 或者 quit;

# 连接 impala 实例
connect;

# 查看查询的执行计划
explain select_sentence;

# 查看最近查询的低级信息(用于查询的诊断和性能的调整)
profile;

# 执行系统命令
shell command; 
# 如 shell date;
# 又如操作 hdfs：shell hdfs fs -mkdir /tmp/impala
```

2、使用 impala-shell 执行 sql  
```bash
# 执行包含 SQL 语句的文件
impala-shell -f sql_file

# 直接 sql 查询
impala-shell -q "select_sentance"

# 查询结果输出到指定文件
impala-shell -f sql_file \
    -o output_file \
    --delimited \
    --output_delimeter 'symbol'  
```

### Impala 数据结构
| 分类 | 类型| 描述 | 示例 |  
| ---- | :-----:  | :----: | ----: |
| 原始类型 | BOOLEAN | true/false | TRUE |  
| | TINYINT | 1 字节的有符号整数 -128~127 | -121 |
| | SMALLINT | 2 个字节的有符号整数，-32768~32767 | 129 |
| | INT | 4 个字节的带符号整数，-2147483648~2147483647 | 1 |
| | BIGINT | 8 字节带符号整数，-9223372036854775808~9223372036854775807 | 1 |
| | FLOAT | 4字节单精度浮点数 | 1.0 |	
| | DOUBLE | 8字节双精度浮点数 | 1.0 |
| | DEICIMAL | 任意精度的带符号小数，用于存储十进制值 | 1.0 |
| | COMPLEX | 复数 | 1 + 2j |
| | STRING | 无上限可变长度字符串 | "a", 'a' |
| | VARCHAR | 可变长度字符串，最大长度为 65535 | "a", 'a' |
| | CHAR | 固定长度字符串，用空格填充，可存储最大长度为 255 | "a", 'a' |
| | TIMESTAMP | 时间戳，纳秒精度 | 1519714435439 |
| 复杂类型 | ARRAY | 有序的的同类型的集合 | array(1,2) |
| | MAP | key-value,key必须为原始类型，value可以任意类型 |map('a', 1, 'b', 2) |
| | STRUCT | 字段集合,类型可以不同 | struct('1',1,1.0), named_stract('col1','1','col2',1,'clo3',1.0) |

### Impala 注释
```bash
-- 这是单行注释

/*
这是多行注释_1.
这是多行注释_2.
这是多行注释_3.
*/
```
