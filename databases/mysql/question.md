
### 开发问题汇集
- MySQL 的复制原理以及流程
> mysql 要做到主从复制，其实依靠的是二进制日志。  
> 假设主服务器叫 A，从服务器叫 B；主从复制就是 B 跟着 A 学，A 做什么，B 就做什么。那么 B 怎么同步 A 的动作呢？现在 A 有一个日志功能，把自己所做的增删改查的动作全都记录在日志中，B 只需要拿到这份日志，照着日志上面的动作施加到自己身上就可以了。这样就实现了主从复制。  
> 
> 主：binlog 线程 —— 记录下所有改变了数据库数据的语句，放进 master 上的 binlog 中；  
> 从：io 线程 —— 在使用 start slave 之后，负责从 master 上拉取 binlog 内容，放进 自己的 relay log 中；  
> 从：sql 执行线程 —— 执行 relay log 中的语句  

- MySQL 中存储引擎 MyISAM 与 InnoDB 的区别
> InnoDB 支持事务，而 MyISAM 不支持事务
> InnoDB 支持行级锁，而 MyISAM 支持表级锁
> InnoDB 支持 MVCC, 而 MyISAM 不支持
> InnoDB 支持外键，而 MyISAM 不支持
> InnoDB 不支持全文索引，而 MyISAM 支持

- MySQL 中索引方法 btree 和 hash 的区别
> 1.在精确查找的情况下：hash 索引要高于 btree 索引，因为 hash 索引查找数据基本上能一次定位数据（大量 hash 值相等的情况下性能会有所降低，也可能低于 btree）,而 btree 索引基于节点上查找，所以在精确查找方面 hash 索引一般会高于 btree 索引。  
> 2.在范围性查找情况下：比如 'like'等范围性查找 hash 索引无效，因为 hash 算法是基于等值计算的。  
> 3.btree 支持的联合索引的最优前缀；hash 是无法支持的，hash 联合索引要么全用，要么全不用。  
> 4.hash 是不支持索引排序的，索引值和 hash 计算出来的 hash 值大小并不一定一致  
> 
> 因为 hash 结构每个键只对应一个值, 而且是散列的方式分布. 所以他并不支持范围查找和排序等功能；B + 树在查找单条记录的速度虽然比不上 hash 索引, 但是因为更适合排序等操作

- MySQL 中 varchar 与 char 的区别以及 varchar(50) 中的 50 代表的含义
> varchar 与 char 的区别在于 char 是一种固定长度的类型，varchar 则是一种可变长度的类型  
> varchar(50) 中 50 的含义是最多存放 50 个字符，varchar(50) 和 varchar(200) 存储 hello 所占空间一样，但后者在排序时会消耗更多内存，因为 order by column 采用 fixed_length 计算 column 长度(memory 存储引擎也一样)  
> 
> 扩展 int(20) 中 20 的含义  
> int(20）中 20 的含义是指显示字符的长度（默认不补全 0），最大为 255，比如它是记录行数的 id,插入 10 笔资料，它就显示 00000000001 ~~~ 00000000010
> 当字符的位数超过 11,它也只显示 11 位，如果你没有加那个让它未满 11 位就前面加 0 的参数，20 不会显示为 020，但仍占 4 字节存储，存储范围不变；  
> mysql 这么设计对大多数应用是没有意义的，只是规定一些工具用来显示字符的个数；int(1) 和 int(20) 存储和计算均一样  

- MySQL 统计 sum 返回 null，那么如何变成 0？
> 使用 COALESCE 即可：SELECT COALESCE(SUM(total),0)  FROM table_name 

- MySQL 统计最近 7 天数据
```mysql
select a.click_date,ifnull(b.count,0) as count_reg  
from (  
    SELECT curdate() as click_date  
    union all  
    SELECT date_sub(curdate(), interval 1 day) as click_date  
    union all  
    SELECT date_sub(curdate(), interval 2 day) as click_date  
    union all  
    SELECT date_sub(curdate(), interval 3 day) as click_date  
    union all  
    SELECT date_sub(curdate(), interval 4 day) as click_date  
    union all  
    SELECT date_sub(curdate(), interval 5 day) as click_date  
    union all  
    SELECT date_sub(curdate(), interval 6 day) as click_date  
) a left join (  
  select date(FROM_UNIXTIME(reg_time)) as datetime, count(*) as count  
  from member  
  group by date(FROM_UNIXTIME(reg_time))  
) b on a.click_date = b.datetime;
```

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

- Explain 列的解释
> EXPLAIN 显示了 MySQL 如何使用索引来处理 SELECT 语句以及连接表，可以帮助选择更好的索引和写出更优化的查询语句。  
```mysql
EXPLAIN 列的解释：  
| 列 | 描述 |  
| :-----:  | :----:  |  
| table	| 显示这一行的数据是关于哪张表的。 |  
| type | 这是重要的列，显示连接使用了何种类型。从最好到最差的连接类型为 const、eq_ref、ref、range、index和ALL。|  
| possible_keys | 显示可能应用在这张表中的索引。如果为空，没有可能的索引。可以为相关的域从WHERE语句中选择一个合适的语句。|  
| key |	实际使用的索引。如果为NULL，则没有使用索引。很少的情况下，MySQL 会选择优化不足的索引。这种情况下，可以在SELECT语句中使用USE INDEX（indexname） 来强制使用一个索引或者用IGNORE INDEX（indexname）来强制 MySQL 忽略索引。 |  
| key_len |	使用的索引的长度。在不损失精确性的情况下，长度越短越好。|  
| ref |	显示索引的哪一列被使用了，如果可能的话，是一个常数。|  
| rows | MySQL 认为必须检查的用来返回请求数据的行数。|  
| extra | 关于 MySQL 如何解析查询的额外信息。将在表 4.3 中讨论，但这里可以看到的坏的例子是Using temporary和Using filesort，意思 MySQL 根本不能使用索引，结果是检索会很慢。|  

> Extra 列返回的描述的意义：  
| 值 | 意义 |  
| :-----:  | :----:  |
| Distinct | 一旦 MySQL 找到了与行相联合匹配的行，就不再搜索了。|  
| Not exists | MySQL 优化了LEFT JOIN，一旦它找到了匹配LEFT JOIN标准的行，就不再搜索了。|  
| Range checked for each Record（index map:#）| 没有找到理想的索引，因此对于从前面表中来的每一个行组合，MySQL 检查使用哪个索引，并用它来从表中返回行。这是使用索引的最慢的连接之一。|  
| Using filesort | 看到这个的时候，查询就需要优化了。MySQL 需要进行额外的步骤来发现如何对返回的行排序。它根据连接类型以及存储排序键值和匹配条件的全部行的行指针来排序全部行。|  
| Using index | 列数据是从仅仅使用了索引中的信息而没有读取实际的行动的表返回的，这发生在对表的全部的请求列都是同一个索引的部分的时候。|  
| Using temporary | 看到这个的时候，查询需要优化了。这里，MySQL 需要创建一个临时表来存储结果，这通常发生在对不同的列集进行ORDER BY上，而不是GROUP BY上。|  
| Where used | 	使用了WHERE从句来限制哪些行将与下一张表匹配或者是返回给用户。如果不想返回表中的全部行，并且连接类型ALL或index，这就会发生，或者是查询有问题不同连接类型的解释（按照效率高低的顺序排序）。|  
| system | 表只有一行 system 表。这是 const 连接类型的特殊情况 。|  
| const | 表中的一个记录的最大值能够匹配这个查询（索引可以是主键或惟一索引）。因为只有一行，这个值实际就是常数，因为 MySQL 先读这个值然后把它当做常数来对待。|  
| eq_ref | 在连接中，MySQL 在查询时，从前面的表中，对每一个记录的联合都从表中读取一个记录，它在查询使用了索引为主键或惟一键的全部时使用。|  
| ref	| 这个连接类型只有在查询使用了不是惟一或主键的键或者是这些类型的部分（比如，利用最左边前缀）时发生。对于之前的表的每一个行联合，全部记录都将从表中读出。这个类型严重依赖于根据索引匹配的记录多少—越少越好。|  
| range	| 这个连接类型使用索引返回一个范围中的行，比如使用 > 或 < 查找东西时发生的情况。|  
| index	| 这个连接类型对前面的表中的每一个记录联合进行完全扫描（比ALL更好，因为索引一般小于表数据）。|  
| ALL	| 这个连接类型对于前面的每一个记录联合进行完全扫描，这一般比较糟糕，应该尽量避免。|  

EXPLAIN SELECT `surname`,`first_name` FORM `a`,`b` WHERE `a`.`id`=`b`.`id`
```

- 数据库设计和查询优化
> Schema 设计时主要考虑: 标准化,数据类型,索引  
> 一个数据库设计可以混合使用,一部分表格标准化,一部分表格非标准化(非标准化表格适当冗余)  
> 最优的数据类型,使表在磁盘上占据的空间尽可能小,读写快,占用内存少(索引也尽量建立在较小的列上)  
> 正确索引,提高 Select,Update,Delete 性能  
> 
> 不同的Sql不同的优化方案  
> Explain Sql 查看结果,分析查询  
> 查询使用匹配的类型  
> 使用 long-slow-queries 记录较慢查询,分析优化  
>
> 服务器端优化  
> 安装适当的 MySql 版本  
> 如果服务器使用 Intel 处理器, 使用 Intel C++ 版本可提高 30 % 效率  
> 
> 配置优化，常见优化项:  
> charset  
> max_allowed_packet  
> max_connections  
> table_cache_size  
> query_cache_size  
> 
> 存储引擎优化  
> MyISAM  
> MyISAM 引擎特点  
> 不支持事务, 提供高速存储, 检索以及全文搜索能力.  
> 宕机会破坏表.  
> 使用的磁盘和内存空间小.  
> 基于表的锁, 并发更新数据会出现严重性能问题.  
> MySql 只缓存索引, 数据由 OS 缓存.  
> MyISAM 适用情况
> 日志系统.  
> 只读操作或者大部分读操作.  
> 全表扫描.  
> 批量导入数据.  
> 没有事务的低并发读写.  
> MyISAM 优化策略  
> NOT NULL, 可以减少磁盘存储.  
> Optimize Table, 碎片整理, 回收空闲空间.  
> Deleting/updating/adding 大量数据的时候禁止使用 index.  
> 参数优化, key_buffer_size_variable 索引缓存设置.  
> 避免并发 Inset Update.  
>
> InnoDB  
> InnoDB 引擎特点  
> 具有提交, 回滚和崩溃恢复能力的事务安全存储引擎.  
> 处理巨大数据量性能卓越, 它的 CPU 使用效率非常高.  
> 需要更多的内存和磁盘存储空间.  
> 数据和索引都缓存在内存中.  
> InnoDB 适用情况  
> 需要事务的应用.  
> 高并发的应用.  
> 自动恢复.  
> 较快速的基于主键的操作.  
> InnoDB 优化策略  
> 尽量使用 short,integer 的主键.  
> 使用 prefix keys, 因为 InnoDB 没有 key 压缩功能.  
> 参数优化, innodb_buffer_pool_size,innodb_data_home_dir 等等  

- mysql 中 in 和 exists 区别
> mysql 中的 in 语句是把外表和内表作 hash 连接，而 exists 语句是对外表作 loop 循环，每次 loop 循环再对内表进行查询。一直大家都认为 exists 比 in 语句的效率要高，这种说法其实是不准确的。这个是要区分环境的。  
> 如果查询的两个表大小相当，那么用 in 和 exists 差别不大。
> 如果两个表中一个较小，一个是大表，则子查询表大的用 exists，子查询表小的用 in。
> not in 和 not exists 如果查询语句使用了 not in 那么内外表都进行全表扫描，没有用到索引；而 not extsts 的子查询依然能用到表上的索引。所以无论那个表大，用 not exists 都比 not in 要快。

- 索引设计有哪些原则
> 适合索引的列是出现在 where 子句中的列，或者连接子句中指定的列  
> 基数较小的类，索引效果较差，没有必要在此列建立索引  
> 使用短索引，如果对长字符串列进行索引，应该指定一个前缀长度，这样能够节省大量索引空间  
> 不要过度索引。索引需要额外的磁盘空间，并降低写操作的性能。在修改表内容的时候，索引会进行更新甚至重构，索引列越多，这个时间就会越长。所以只保持需要的索引有利于查询即可。

- mysql 主从同步如何实现
> 


- 悲观锁和乐观锁
> 悲观锁——总是假设最坏的情况，每次去拿数据的时候都认为别人会修改，所以每次在拿数据的时候都会上锁，这样别人想拿这个数据就会阻塞直到它拿到锁（共享资源每次只给一个线程使用，其它线程阻塞，用完后再把资源转让给其它线程）。传统的关系型数据库里边就用到了很多这种锁机制，比如行锁，表锁等，读锁，写锁等，都是在做操作之前先上锁。  
> 乐观锁——总是假设最好的情况，每次去拿数据的时候都认为别人不会修改，所以不会上锁，但是在更新的时候会判断一下在此期间别人有没有去更新这个数据，可以使用版本号机制和 CAS 算法实现。乐观锁适用于多读的应用类型，这样可以提高吞吐量，像数据库提供的类似于 write_condition 机制，其实都是提供的乐观锁。  
> 
> 乐观锁适用于写比较少的情况下（多读场景），即冲突真的很少发生的时候，这样可以省去了锁的开销，加大了系统的整个吞吐量。但如果是多写的情况，一般会经常产生冲突，这就会导致上层应用会不断的进行 retry，这样反倒是降低了性能，所以一般多写的场景下用悲观锁就比较合适。  
> 
> 乐观锁一般会使用版本号机制或 CAS 算法实现。  
> 版本号机制——一般是在数据表中加上一个数据版本号 version 字段，表示数据被修改的次数，当数据被修改时，version 值会加一。当线程 A 要更新数据值时，在读取数据的同时也会读取 version 值，在提交更新时，若刚才读取到的 version 值为当前数据库中的 version 值相等时才更新，否则重试更新操作，直到更新成功。  
> CAS 算法——即 compare and swap（比较与交换），是一种有名的无锁算法。无锁编程，即不使用锁的情况下实现多线程之间的变量同步，也就是在没有线程被阻塞的情况下实现变量的同步，所以也叫非阻塞同步（Non-blocking Synchronization）。CAS 算法涉及到三个操作数：需要读写的内存值 V、进行比较的值 A、拟写入的新值 B，当且仅当 V 的值等于 A 时，CAS 通过原子方式用新值 B 来更新 V 的值，否则不会执行任何操作（比较和替换是一个原子操作）。一般情况下是一个自旋操作，即不断的重试。  
> 
> 乐观锁的缺点  
> 如果一个变量 V 初次读取的时候是 A 值，并且在准备赋值的时候检查到它仍然是 A 值，那我们就能说明它的值没有被其他线程修改过了吗？很明显是不能的，因为在这段时间它的值可能被改为其他值，然后又改回 A，那 CAS 操作就会误认为它从来没有被修改过。这个问题被称为 CAS 操作的 "ABA" 问题。  
> 自旋 CAS（也就是不成功就一直循环执行直到成功）如果长时间不成功，会给 CPU 带来非常大的执行开销。  
> CAS 只对单个共享变量有效，当操作涉及跨多个共享变量时 CAS 无效。  






