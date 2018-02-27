
### 什么是 Hive
hive 是基于 hadoop 的数据仓库

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
注意：Hive 默认元数据库并不是 Mysql，但是因为默认元数据库存在局限，所以最好使用 Mysql。

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
GRANT ALL ON hive.* TO hive@’%’ IDENTIFIED BY 'hive_password';
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
        <value>hive<value>
        <description>username to use against metastore database</description>
	</property>
	<property>  
        <name>javax.jdo.option.ConnectionPassword</name>
        <value>hive</value>
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
```

### Hive 基本操作
```bash
# 创建数据库
CREATE (DATABASE|SCHEMA) [IF NOT EXISTS] database_name
    [COMMENT database_comment]
    [LOCATION hdfs_path]
    [WITH DBPROPERTIES (property_name=property_value, ...)];
ep:
CREATE DATABASE IF NOT EXISTS example_database
COMMENT '这是一个例子'
LOCATION 'hdfs://user/hive/wherehouse/example/example_database.db/';

# 删除数据库
# 默认情况下，Hive不允许删除一个里面有表存在的数据库；
# 如果想删除数据库，要么先将数据库中的表全部删除，要么可以使用CASCADE关键字，使用该关键字后，Hive会自己将数据库下的表全部删除；
# RESTRICT关键字就是默认情况，即如果有表存在，则不允许删除数据库。
DROP (DATABASE|SCHEMA) [IF EXISTS] database_name [RESTRICT|CASCADE];
ep:
DROP DATABASE if exists example_database CASCADE;

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

# 加载本地文件
LOAD DATA LOCAL INPATH '/hive_data_path'
INTO TABLE example_table PARTITION (day='xxxx-xx-xx');
# 加载 HDFS 文件
LOAD DATA INPATH '/hive_data_path'
INTO TABLE example_table PARTITION (day='xxxx-xx-xx');

# 从子查询中加载数据
INSERT overwrite TABLE example_table PARTITION (day='xxxx-xx-xx');
SELECT example_field from temp_table;
```

[更多操作](./hsql.md)
