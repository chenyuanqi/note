
### 关系型数据库 Mysql
关系型数据库，是指采用了关系模型来组织数据的数据库。  

- 关系型数据库的优点  
> 容易理解：二维表结构是非常贴近逻辑世界的一个概念，关系模型相对网状、层次等其他模型来说更容易理解  
> 使用方便：通用的SQL语言使得操作关系型数据库非常方便  
> 易于维护：丰富的完整性(实体完整性、参照完整性和用户定义的完整性)大大减低了数据冗余和数据不一致的概率

- 关系型数据库瓶颈  
> 1).高并发读写需求  
> 网站的用户并发性非常高，往往达到每秒上万次读写请求，对于传统关系型数据库来说，硬盘 I/O 是一个很大的瓶颈
> 2).海量数据的高效率读写  
> 网站每天产生的数据量是巨大的，对于关系型数据库来说，在一张包含海量数据的表中查询，效率是非常低的  
> 3).高扩展性和可用性  
> 在基于 web 的结构当中，数据库是最难进行横向扩展的，当一个应用系统的用户量和访问量与日俱增的时候，数据库却没有办法像 web server 和 app server 那样简单的通过添加更多的硬件和服务节点来扩展性能和负载能力。对于很多需要提供 24 小时不间断服务的网站来说，对数据库系统进行升级和扩展是非常痛苦的事情，往往需要停机维护和数据迁移。  

### Mysql 锁和事务
- 锁
> 因为数据库要解决并发控制问题。在同一时刻，可能会有多个客户端对同一张表进行操作，比如有的在读取该行数据，其他的尝试去删除它。为了保证数据的一致性，数据库就要对这种并发操作进行控制，因此就有了锁的概念。  

- 锁的分类
> 从对数据库操作的类型分  
> 读锁（共享锁）：针对同一块数据，多个读操作可以同时进行而不会互相影响。由读表操作加上的锁，加锁后其他用户只能获取该表或行的共享锁，不能获取排它 锁，也就是说只能读不能写。  
> 写锁（排它锁）：当当前写操作没有完成之前，它会阻断其他写锁和读锁。由写表操作加上的锁，加锁后其他用户不能获取该表或行的任何锁。  
> 
> 从锁定的数据范围分  
> 表锁：锁定某个表。  
> 行锁 ：锁定某行。  
> 

- 锁粒度
> 为了尽可能提高数据库的并发度，每次锁定的数据范围越小越好。理论上每次只锁定当前操作的数据的方案会得到最大的并发度，但是管理锁是很耗费资源的事情。因此数据库系统需要在高并发响应和系统性能两方面进行平衡，这样就产生了“锁粒度”的概念。
>
> 表锁：管理锁的开销最小，同时允许的并发量也最小的锁机制。MyIsam 存储引擎使用的锁机制。当要写入数据时，把整个表都锁上，此时其他读、写动作一律等待。在 MySql 中，除了 MyIsam 存储引擎使用这种锁策略外，MySql 本身也使用表锁来执行某些特定动作，比如 alter table.  
> 行锁：可以支持最大并发的锁策略。InnoDB和Falcon两张存储引擎都采用这种策略。  
> MySql 是一种开放的架构，你可以实现自己的存储引擎，并实现自己的锁粒度策略，不像 Oracle，你没有机会改变锁策略，Oracle 采用的是行锁。从大到小，mysql 服务器仅支持表级锁，行锁需要存储引擎完成。粒度越精细，并发性越好。即行锁的并发性最好，但需要存储引擎的支持。

- 事务  
mysql 事务主要用于处理操作量大，复杂度高的数据。比如说，在人员管理系统中，你删除一个人员，你既要删除人员的基本资料，也要删除和该人员相关的信息，如信箱，文章等。这样，这些数据库操作语句就构成一个事务。注意以下几点：  
* 在 mysql 中只有使用了 Innodb 数据库引擎的数据库或表才支持事务   
* 事务处理可以用来维护数据库的完整性，保证成批的 sql 语句要么全部执行，要么全部不执行  
* 事务用来管理 insert，update，delete 语句  

事务是 DBMS 的执行单位，它由有限的数据库操作序列组成，但不是任意的数据库操作序列都能成为事务。  
一般来说，事务必须满足4个条件（ACID）：  
原子性：组成事务处理的语句形成了一个逻辑单元，不能只执行其中一部分。换句话说，事务是不可分割的最小单元。比如银行转账过程中，必须同时从一个账户减去转账金额，并加到另一个账户中，只改变一个账户是不合理的。  
一致性：在事务处理执行前后，mysql数据库是一致的。也就是说，事务应该正确的转换系统状态。比如银行转账过程中，要么转账金额从一个账户到另一个账户，要么两个账户都不变，没有其他情况。  
隔离性：一个事务处理对另外一个事务处理没有影响。比如说银行转账过程中，在转账事务没有提交之前，另一个转账事务只能处于等待状态 。  
可靠性：事务处理的效果能够被永久保存下来。反过来说，事务能够承受所有的失败。包括服务器，进程，通信以及媒体失败等等。比如银行转账过程中，转账后账户的状态要能被保存下来。  

- 事务隔离级别
> SQL 标准定义了 4 种隔离级别，包括了一些具体规则，用来限定事务内外的哪些改变是可见的，哪些是不可见的。低级别的隔离级一般支持更高的并发处理，并拥有更低的系统开销。  
> 1、Read Uncommitted（读取未提交内容）  
> 在该隔离级别，所有事务都可以看到其他未提交事务的执行结果。本隔离级别很少用于实际应用，因为它的性能也不比其他级别好多少。读取未提交的数据，也被称之为脏读（Dirty Read）。    
> 2、Read Committed（读取提交内容）  
> 这是大多数数据库系统的默认隔离级别（但不是 MySQL 默认的）。它满足了隔离的简单定义：一个事务只能看见已经提交事务所做的改变。这种隔离级别也支持所谓的不可重复读（Nonrepeatable Read），因为同一事务的其他实例在该实例处理其间可能会有新的 commit，所以同一 select 可能返回不同结果。   
> 3、Repeatable Read（可重读）  
> 这是 MySQL 的默认事务隔离级别，它确保同一事务的多个实例在并发读取数据时，会看到同样的数据行。 不过理论上，这会导致另一个棘手的问题：幻读（Phantom Read）。简单的说，幻读指当用户读取某一范围的数据行时， 另一个事务又在该范围内插入了新行，当用户再读取该范围的数据行时，会发现有新的“幻影” 行。 InnoDB 和 Falcon 存储引擎通过多版本并发控制（MVCC，Multiversion Concurrency Control）机制解决了该问题。   
> 4、Serializable（可串行化）  
> 这是最高的隔离级别，它通过强制事务排序，使之不可能相互冲突，从而解决幻读问题。简言之，它是在每个读的数据行上加上共享锁。在这个级别，可能导致大量的超时现象和锁竞争。
>  
> 这四种隔离级别采取不同的锁类型来实现，若读取的是同一个数据的话，就容易发生问题。例如：  
> 脏读(Drity Read)：某个事务已更新一份数据，另一个事务在此时读取了同一份数据，由于某些原因，前一个RollBack了操作，则后一个事务所读取的数据就会是不正确的。  
> 幻读(Phantom Read):在一个事务的两次查询中数据笔数不一致，例如有一个事务查询了几列(Row)数据，而另一个事务却在此时插入了新的几列数据，先前的事务在接下来的查询中，就会发现有几列数据是它先前所没有的。  
> 不可重复读(Non-repeatable read):在一个事务的两次查询之中数据不一致，这可能是两次查询过程中间插入了一个事务更新的原有的数据。

### MyISAM 和 InnoDB 的区别
> 1、MyISAM 查询效率更高，但是不支持事物  
> 2、InnoDB 插入、更新较高，支持事物处理  
> 3、MyISAM 支持表锁， InnoDb 支持行锁  
> 4、MyISAM 是默认引擎，InnoDB 需要指定  
> 5、InnoDB 不支持 FULLTEXT 类型的索引  

### Mysql 增删改查
```bash
# 连接 mysql
mysql [-h host] -u user -p [database]

# 分配用户权限
GRANT ALL ON menagerie.* TO 'your_mysql_name'@'your_client_host';

# 创建数据库
create database test charset=utf8;

# 删除数据库
DROP DATABASE test;

# 创建数据表
CREATE TABLE event (name VARCHAR(20), date DATE, type VARCHAR(15), remark VARCHAR(255));

# 修改表名
ALTER TABLE table_name RENAME TO wanted_table_name

# 删除数据表
DROP TABLE table_name;

# 添加列
alter table table_name add column column_name varchar(30);
# 删除列
alter table table_name drop column column_name;
# 修改列名
alter table table_name change column_name new_column_name new_column_type;
# 修改列属性，修改字段数据类型
alter table table_name modify column_name column_type;
# 建表时设置主键
create table table_name( id int primary key);
# 或者
create table table_name(id int, primary key(id));
# 非建表时设置主键
alter table table_name add primary key;
# 删除主键
alter table table_name drop primary key;
# 建表时添加唯一键
create table table_name(columnName int unique);
# 或者
create table table_name(columnName int, unique key(column_name));
# 非建表时添加唯一键
alter table table_name add unique key(column_name);
# 删除唯一键
alter table table_name drop index unique_index_name;
# 建表时添加索引
create table table_name(column_name int key);
# 或者
create table table_name(column_name int, key/index index_name(column_name));
# 建表时添加多列索引
create table table_name(column_name1 int, column_name2 int, key/index index_name(column_name1, column_name2));
# 非建表时添加索引
alter table table_name add key/index index_name(column_name1, column_name2);
# 删除索引
alter table table_name drop key/index column_name;

# 写入数据
# INSERT INTO table_name ( field1, field2,...fieldN ) VALUES (value1, value2,...valueN );
INSERT INTO pet VALUES ('Puffball','Diane','hamster','f','1999-03-30',NULL);
INSERT INTO table_name (column1, column2) VALUES (value1, value2);

# 更新数据
# UPDATE table_name SET field1=new-value1, field2=new-value2 [WHERE Clause]
UPDATE pet SET birth = '1989-08-31' WHERE name = 'Bowser';

# 删除数据
# DELETE FROM table_name [WHERE Clause]
DELETE FROM table_name WHERE condition;
DELETE * FROM table_name;

# 基本条件查询
SELECT name, email FROM pet WHERE birth >= '1998-1-1';
SELECT column1, column2 FROM table_name WHERE condition;
SELECT * FROM table_name WHERE condition1 AND condition2;
SELECT * FROM table_name WHERE condition1 OR condition2;
SELECT * FROM table_name WHERE NOT condition;
SELECT * FROM table_name WHERE condition1 AND (condition2 OR condition3);
SELECT * FROM table_name WHERE EXISTS (SELECT column_name FROM table_name WHERE condition);

# AS 用于给表或者列分配别名
SELECT column_name AS alias_name FROM table_name;
SELECT column_name FROM table_name AS alias_name;
SELECT column_name AS alias_name1, column_name2 AS alias_name2;
SELECT column_name1, column_name2 + ‘, ‘ + column_name3 AS alias_name;

# IN 运算符是多个 OR 条件的简写
SELECT column_names FROM table_name WHERE column_name IN (value1, value2, …);
SELECT column_names FROM table_name WHERE column_name IN (SELECT STATEMENT);

# BETWEEN 用于过滤给定范围的值的运算符
SELECT column_names FROM table_name WHERE column_name BETWEEN value1 AND value2;
SELECT * FROM Products WHERE (column_name BETWEEN value1 AND value2) AND NOT column_name2 IN (value3, value4);
SELECT * FROM Products WHERE column_name BETWEEN \#01/07/1999\# AND \#03/12/1999\#;

# 去重查询
SELECT DISTINCT owner FROM pet;

# is null 查询，代表一个字段没有值
SELECT name, birth FROM pet WHERE death IS NOT NULL;
SELECT * FROM table_name WHERE column_name IS NULL;

# 模糊查询(_ 匹配一个，% 匹配零个或多个)
SELECT * FROM pet WHERE name LIKE '_b%';
# LIKE ‘a_%_%’ （查找任何以 “a” 开头且长度至少为 3 的值）
# LIKE ‘[a-c]%’（查找任何以 “a” 或“b”或 “c” 开头的值）

# 正则查询(^ 开头定位符，$ 结尾定位符，. 任意字符，{number} 出现次数...)
SELECT * FROM pet WHERE name REGEXP '^b.{5}$';

# 统计
# COUNT 返回出现次数
SELECT COUNT(*) FROM pet;
SELECT COUNT (DISTINCT column_name);

# MIN() and MAX() 返回所选列的最小 / 最大值
SELECT MIN (column_names) FROM table_name WHERE condition;
SELECT MAX (column_names) FROM table_name WHERE condition;

# AVG() 返回数字列的平均值
SELECT AVG (column_name) FROM table_name WHERE condition;

# SUM() 返回数值列的总和
SELECT SUM (column_name) FROM table_name WHERE condition;

# 排序(默认升序)
SELECT name, birth FROM pet ORDER BY birth [DESC];

# 分组，通常与聚合函数（COUNT，MAX，MIN，SUM，AVG）一起使用，用于将结果集分组为一列或多列
SELECT sex, COUNT(*) FROM pet GROUP BY sex;
SELECT column_name1, COUNT(column_name2) FROM table_name WHERE condition GROUP BY column_name1 ORDER BY COUNT(column_name2) DESC;
# HAVING 子句指定 SELECT 语句应仅返回聚合值满足指定条件的行。它被添加到 SQL 语言中，因为 WHERE 关键字不能与聚合函数一起使用
SELECT COUNT(column_name1), column_name2 FROM table GROUP BY column_name2 HAVING COUNT(column_name1) > 5;

# UNION 用于组合两个或者多个 SELECT 语句的结果集的运算符
SELECT columns_names FROM table1 UNION SELECT column_name FROM table2;
# 每个 SELECT 语句必须拥有相同的列数  
# 列必须拥有相似的数据类型  
# 每个 SELECT 语句中的列也必须具有相同的顺序
# UNION 仅允许选择不同的值, UNION ALL 允许重复

# ANY|ALL 用于检查 WHERE 或 HAVING 子句中使用的子查询条件的运算符
SELECT columns_names FROM table1 WHERE column_name operator (ANY|ALL) (SELECT column_name FROM table_name WHERE condition);
# ANY 如果任何子查询值满足条件，则返回 true
# ALL 如果所有子查询值都满足条件，则返回 true

# 导入数据
LOAD DATA LOCAL INFILE '路径/pet.txt' INTO TABLE event;

# 导出数据(或可使用 mysqldump)
SELECT * FROM pet INTO OUTFILE '/tmp/pet.txt';

# 查看表结构
DESC pet;

# 查看版本号和当前日期
SELECT VERSION(), CURRENT_DATE;

# 包含查询
# 6 在  '[1,2,3,6]' 内，示例
SELECT * FROM `articles_onlines` WHERE FIND_IN_SET('6', TRIM(TRAILING ']' FROM TRIM(LEADING '[' FROM `article_tags`)));
# 6 在 '1,2,3,6' 内，示例
SELECT * FROM `articles_onlines` WHERE FIND_IN_SET(6, `article_tags`);
# 某字段是否存在字符串中，譬如 tags_id in ('1,2,3')
SELECT * FROM `article_tabs` WHERE `tags_id` IN '1,2,3';

# 创建视图
CREATE VIEW view_name AS SELECT column1, column2 FROM table_name WHERE condition;
# 检索视图
SELECT * FROM view_name;
# 删除视图
DROP VIEW view_name;

# 连表查询
SELECT pet.name,TIMESTAMPDIFF(YEAR,birth,date) AS age,remark FROM pet INNER JOIN event ON pet.name = event.name WHERE event.type = 'litter';
# INNER JOIN 内连接，返回在两张表中具有匹配值的记录
SELECT column_names FROM table1 INNER JOIN table2 ON table1.column_name=table2.column_name;
SELECT table1.column_name1, table2.column_name2, table3.column_name3 FROM ((table1 INNER JOIN table2 ON relationship) INNER JOIN table3 ON relationship);

# LEFT (OUTER) JOIN 左外连接，返回左表（table1）中的所有记录，以及右表中的匹配记录（table2）
SELECT column_names FROM table1 LEFT JOIN table2 ON table1.column_name=table2.column_name;
# RIGHT (OUTER) JOIN 右外连接，返回右表（table2）中的所有记录，以及左表（table1）中匹配的记录
SELECT column_names FROM table1 RIGHT JOIN table2 ON table1.column_name=table2.column_name;
# FULL (OUTER) JOIN 全外连接，全连接是左右外连接的并集. 连接表包含被连接的表的所有记录, 如果缺少匹配的记录, 以 NULL 填充。
SELECT column_names FROM table1 FULL OUTER JOIN table2 ON table1.column_name=table2.column_name;
# Self JOIN 自连接，表自身连接
SELECT column_names FROM table1 T1, table1 T2 WHERE condition;
```

