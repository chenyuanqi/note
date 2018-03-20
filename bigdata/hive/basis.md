
### 什么是 Hive
Hive 是基于 hadoop 的数据仓库。  
Hive 的设计目的是使用 SQL 生成 MapReduce 代码对 HDFS 执行查询。   

### Hive 的系统架构
```html
                       — — Derby
       — — MetaStore -|
Hive -|                — — Mysql
       — — HDFS (/user/hive/warehouse)

```

Hive 中的元数据包括表的名字，表的列和分区及其属性，表的属性（是否为外部表等），表的数据所在目录等。  
Hive 数据存储在 HDFS 中，大部分的查询由 MapReduce 完成（SELECT * 不会生成 MapReduce 任务）。    
Hive 用户接口包括 CLI, Client, WUI，解释器、编译器、优化器完成 HQL 查询语句从词法分析、语法分析、编译、优化以及查询计划的生成。生成的查询计划存储在 HDFS 中，并在随后有 MapReduce 调用执行。  
注意：Hive 默认元数据库并不是 Mysql (存储在自带的内存数据库 Derby)，但是因为默认元数据库存在局限（Derby 数据存在内存，不稳定），所以最好使用 Mysql。

### Hive 安装步骤
1、在 [Hive Dist](http://archive.apache.org/dist/hive/) 中选择下载一个稳定版本的压缩包，然后解压到大数据目录。
```bash
# 下载，解压
wget http://archive.apache.org/dist/hive/hive-x.x.x/apache-hive.x.x.x.tar.gz
tar -xzvf apache-hive.x.x.x.tar.gz
# 添加环境变量
export HIVE_HOME=/hive_path
export PATH=$HIVE_HOME/bin: $PATH
```

2、在 Mysql 中建立 Hive 元数据库及相关用户
```bash
CREATE DATABASE hive;
GRANT ALL ON hive.* TO hive@'%' IDENTIFIED BY 'hive_password';
flush privileges;
```

3、配置 hive-site.xml  
配置文件在 $HIVE_HOME/conf/ 目录中，
```bash
cd $HIVE_HOME/conf
cp hive-default.xml.template hive-site.xml
vim hive-site.xml
```

```xml
<configuration>
	<property>
        <name>javax.jdo.option.ConnectionURL</name>
        <value>jdbc:mysql://hadoop-master:3306/hive?createDatabaseIfNotExist=true</value>
        <description>JDBC connect string for a JDBC metastore</description>    
	</property>   
	<property> 
        <name>javax.jdo.option.ConnectionDriverName</name> 
        <value>com.mysql.jdbc.Driver</value> 
        <description>Driver class name for a JDBC metastore</description>     
	</property>               
 
	<property> 
        <name>javax.jdo.option.ConnectionUserName</name>
        <value>{hive_user}<value>
        <description>username to use against metastore database</description>
	</property>
	<property>  
        <name>javax.jdo.option.ConnectionPassword</name>
        <value>{hive_password}</value>
        <description>password to use against metastore database</description>  
	</property>          
</configuration>
```

4、启动
```bash
# 启动 metastore
hive --service metastore &

# 进入 hive-cli
hive
# 进入 hive-cli, 并指定日志目录
hive -hiveconf hive.log.dir='log_dir'
# 进入 hive-cli, 并将调试信息输出控制台
hive -hiveconf hive.root.logger=DEBUG,console
> Usage: hive [-hiveconf conf_item=conf_value] [-i file_name] [-f file_name] [-e query_string] [-S]
> -i 从文件中初始化 Sql
> -f 执行文件中 Sql
> -S 静默模式
```

### Hive 内部表和外部表
内部表：加载数据到 hive 所在的 hdfs 目录，删除时，元数据和数据文件都删除  
外部表：不加载数据到 hive 所在的 hdfs 目录，删除时，只删除表结构  

建议使用外部表，因为外部表不会加载数据到 hive 所在的 hdfs 目录，减少数据传输，而且数据可以共享；  
删除表时，只删除表结构，不删除表数据；  
hive 不会修改数据，所以无需担心数据的损坏。

### Hive 虚拟列
Hive 提供了三个虚拟列 INPUT__FILE__NAME、BLOCK__OFFSET__INSIDE__FILE 和 ROW__OFFSET__INSIDE__BLOCK；  
INPUT__FILE__NAME, 用于 mapper 任务的输出文件名；  
ROW__OFFSET__INSIDE__BLOCK 用来排查有问题的输入数据，默认情况下，ROW__OFFSET__INSIDE__BLOCK 是不可用的，需要设置 hive.exec.rowoffset 值为 true 才可用；  
BLOCK__OFFSET__INSIDE__FILE, 当前全局文件的偏移量；对于块压缩文件，就是当前块的文件偏移量，即当前块的第一个字节在文件中的偏移量。

### Hive 数据结构
| 分类 | 类型| 描述 | 示例 |  
| ---- | :-----:  | :----: | ----: |
| 原始类型 | BOOLEAN | true/false | TRUE |
| | TINYINT | 1 字节的有符号整数 -128~127 | -121 |
| | SMALLINT | 2 个字节的有符号整数，-32768~32767 | 129 |
| | INT | 4 个字节的带符号整数，-2147483648~2147483647 | 1 |
| | BIGINT | 8 字节带符号整数，-9223372036854775808~9223372036854775807 | 1 |
| | FLOAT | 4字节单精度浮点数 | 1.0 |	
| | DOUBLE | 8字节双精度浮点数 | 1.0 |
| | DEICIMAL | 任意精度的带符号小数 | 1.0 |
| | STRING | 无上限可变长度字符串 | "a", 'a' |
| | VARCHAR | 可变长度字符串 | "a", 'a' |
| | CHAR | 固定长度字符串 | "a", 'a' |
| | BINARY | 字节数组 | |
| | TIMESTAMP | 时间戳，纳秒精度 | 1519714435439 |
| | DATE | 日期 | '2018-02-28' |
| 复杂类型 | ARRAY | 有序的的同类型的集合 | array(1,2) |
| | MAP | key-value,key必须为原始类型，value可以任意类型 |map('a', 1, 'b', 2) |
| | STRUCT | 字段集合,类型可以不同 | struct('1',1,1.0), named_stract('col1','1','col2',1,'clo3',1.0) |
| | UNION | 在有限取值范围内的一个值| create_union(1,'a',63) |

### Hive 基本操作
```bash
# 创建数据库
CREATE (DATABASE|SCHEMA) [IF NOT EXISTS] database_name
    [COMMENT database_comment]
    [LOCATION hdfs_path]
    [WITH DBPROPERTIES (property_name=property_value, ...)];
ep:
CREATE DATABASE IF NOT EXISTS example_database
COMMENT '创建栗子数据库'
LOCATION 'hdfs://user/hive/wherehouse/example/example_database.db/';

# 删除数据库
# 默认情况下，Hive不允许删除一个里面有表存在的数据库；
# 如果想删除数据库，要么先将数据库中的表全部删除，要么可以使用CASCADE关键字，使用该关键字后，Hive会自己将数据库下的表全部删除；
# RESTRICT关键字就是默认情况，即如果有表存在，则不允许删除数据库。
DROP (DATABASE|SCHEMA) [IF EXISTS] database_name [RESTRICT|CASCADE];
ep:
DROP DATABASE if exists example_database CASCADE;

# 查看所有数据库
SHOW databases;

# 选择数据库
USE database_name;

# 创建表
# 内部表 DROP 的时候会删除 HDFS 上的数据，而外部表不会（无特殊要求，建议使用外部表）
# CREATE EXTERNAL TABLE 表示创建外部表，如果是内部表则不指定 EXTERNAL 关键字 
CREATE EXTERNAL TABLE external_table (id INT, name STRING, day STRING)
COMMENT '这是外部表的注释'
PARTITIONED BY (day STRING) 
ROW FORMAT DELIMITED 
FIELDS TERMINATED BY ',' 
LINES TERMINATED BY '\n' 
COLLECTION ITEMS TERMINATED BY ','
MAP KEYS TERMINATED BY ':'
STORED AS textfile 
LOCATION 'hdfs://hive_table_path/'; 

# 删除表
DROP TABLE [IF EXISTS] table_name [PURGE];
ep:
DROP TABLE IF EXISTS table_name;

# 表重命名
ALTER TABLE table_name RENAME TO new_talbe_name;

# 添加字段
ALTER TABLE table_name ADD COLUMNS (column_name column_type COMMENT 'your column comment');
# 更新字段
ALTER TABLE table_name CHANGE COLUMNS column_name new_column_name column_type COMMENT 'your column comment' FIRST|AFTER another_column_name;
# 替换所有字段
ALTER TABLE table_name REPLACE COLUMNS (column_name INT/others COMMENT 'your column comment');

# 添加分区
ALTER TABLE table_name ADD PATITION (patition_field='patition_value') LOCATION 'hdfs_path';
# 删除分区
ALTER TABLE table_name DROP PATITION (patition_field='patition_value');

# 查看表分区
SHOW PATITIONS table_name;

# 查看表结构
DESC table_name; # 或 DESCRIBE table_name;

# 复制表结构
CREATE TABLE copy_table_name LIKE table_name;

# 查看所有表
SHOW TABLES;
# 正则查看部分表
SHOW TABLES '.*game';

# 加载本地文件
LOAD DATA LOCAL INPATH '/hive_data_path'
INTO TABLE example_table PARTITION (day='xxxx-xx-xx');
# 加载 HDFS 文件
LOAD DATA INPATH '/hive_data_path'
INTO TABLE example_table PARTITION (day='xxxx-xx-xx');

# 从子查询中加载数据
INSERT [OVERWRITE] TABLE example_table PARTITION (day='xxxx-xx-xx');
SELECT example_field from temp_table;
```

[更多操作](./hsql.md)
