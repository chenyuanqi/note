
### 常用命令
```bash
# 连接数据库
mysql [-h 127.0.0.1] [-P 3306] -uroot -p

# 修改 mysql 密码
mysqladmin -u 用户名 -p 旧密码 password 新密码

# 查看 mysql 线程列表
SHOW PROCESSLIST;

# 查看 mysql 变量列表
SHOW VARIABLES;

# Mysql 5.7 以上，查看持有 MDL 读锁的线程 id
select blocking_pid from sys.schema_table_lock_waits;

# 备份与还原
# 备份，将数据的结构与表内数据保存起来,利用 mysqldump 指令完成
# 导出一张表
mysqldump -u用户名 -p密码 库名 表名 > 文件名(/tmp/a.sql)
# 导出多张表
mysqldump -u用户名 -p密码 库名 表1 表2 表3 > 文件名(/tmp/a.sql)
# 导出所有表
mysqldump -u用户名 -p密码 库名 > 文件名(/tmp/a.sql)
# 导出一个库 
mysqldump -u用户名 -p密码 -B 库名 > 文件名(/tmp/a.sql)

# 导入
# 在登录 mysql 的情况下
source  备份文件
# 在不登录的情况下
mysql -u用户名 -p密码 库名 < 备份文件

# 创建用户并授权
# -- 创建用户 laowang
create user 'laowang'@'localhost' identified by '123456';
# -- 授权 test 数据库给 laowang
grant all on test.* to 'laowang'@'localhost';

# 创建数据库并指定编码格式
create database db_name default charset utf8 collate utf8_general_ci;

# 查看 mysql 最大连接数
show variables like 'max_connections%'; # 配置文件 my.ini 中可以设置
# 查询线程连接数
show global status like 'threads_%';
```

### 常用函数
- sum(field) – 求某个字段的和值；
- count(\*) – 查询总条数；
- min(field) – 某列中最小的值；
- max(field) – 某列中最大的值；
- avg(field) – 求平均数；
- current_date() – 获取当前日期；
- now() – 获取当前日期和时间；
- concat(a, b) – 连接两个字符串值以创建单个字符串输出；
- datediff(a, b) – 确定两个日期之间的差异，通常用于计算年龄。  
