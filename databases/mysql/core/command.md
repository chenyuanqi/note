
### 常用命令
```bash
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
```
