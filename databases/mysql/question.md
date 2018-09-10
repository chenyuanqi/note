
### 开发问题汇集
- MySQL 的复制原理以及流程
> 主：binlog 线程 —— 记录下所有改变了数据库数据的语句，放进 master 上的 binlog 中；  
> 从：io 线程 —— 在使用 start slave 之后，负责从 master 上拉取 binlog 内容，放进 自己的 relay log 中；  
> 从：sql 执行线程 —— 执行 relay log 中的语句  

- MySQL 中存储引擎 MyISAM 与 InnoDB 的区别
> InnoDB 支持事务，而 MyISAM 不支持事务
> InnoDB 支持行级锁，而 MyISAM 支持表级锁
> InnoDB 支持 MVCC, 而 MyISAM 不支持
> InnoDB 支持外键，而 MyISAM 不支持
> InnoDB 不支持全文索引，而 MyISAM 支持

- MySQL 中 varchar 与 char 的区别以及 varchar(50) 中的 50 代表的含义
> varchar 与 char 的区别在于 char 是一种固定长度的类型，varchar 则是一种可变长度的类型  
> varchar(50) 中 50 的含义是最多存放 50 个字符，varchar(50) 和 varchar(200) 存储 hello 所占空间一样，但后者在排序时会消耗更多内存，因为 order by column 采用 fixed_length 计算 column 长度(memory 存储引擎也一样)  
> 
> 扩展 int(20) 中 20 的含义  
> int(20）中 20 的含义是指显示字符的长度（默认不补全 0），最大为 255，比如它是记录行数的 id,插入 10 笔资料，它就显示 00000000001 ~~~ 00000000010
> 当字符的位数超过 11,它也只显示 11 位，如果你没有加那个让它未满 11 位就前面加 0 的参数，20 不会显示为 020，但仍占 4 字节存储，存储范围不变；  
> mysql 这么设计对大多数应用是没有意义的，只是规定一些工具用来显示字符的个数；int(1) 和 int(20) 存储和计算均一样  

- 如何从 MySQL 全库备份中恢复某个库和某张表
> 这里主要用到的参数是 --one-database（简写 -o），极大方便了我们的恢复灵活性。
```bash
# 全库备份
mysqldump -uroot -p --single-transaction -A --master-data=2 >dump.sql

# 只还原erp库的内容
mysql -uroot -pMANAGER erp --one-database <dump.sql

# 从全库备份中抽取出 test 表的表结构
sed -e'/./{H;$!d;}' -e 'x;/CREATE TABLE `test`/!d;q' dump.sql

# 从全库备份中抽取出 test 表的内容
grep'INSERT INTO `t`' dump.sql
```

- MySQL 如何支持 emoji 表情
> 使用 utf8_mb4 字符集

- 500 台 db，在最快时间之内重启
> puppet，dsh

- 监控数据库的工具
> 如 zabbix，lepus

- 主从一致性校验的工具
> 如 checksum、mysqldiff、pt-table-checksum

- 如何维护数据库的数据字典
> 我们一般是直接在生产库进行注释，利用工具导出成 excel 方便沟通