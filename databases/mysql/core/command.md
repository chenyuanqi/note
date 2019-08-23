
### 常用命令
```bash
# 查看 mysql 进程列表
show processlist;

# Mysql 5.7 以上，查看持有 MDL 读锁的线程 id
select blocking_pid from sys.schema_table_lock_waits;

```
