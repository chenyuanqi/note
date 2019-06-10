
### Mysql 数据库设计
规则 1：  
一般情况可以选择 MyISAM 存储引擎，如果需要事务支持必须使用 InnoDB 存储引擎。  
注意：  
MyISAM 存储引擎 B-tree 索引有一个很大的限制：参与一个索引的所有字段的长度之和不能超过 1000 字节。另外，MyISAM 数据和索引是分开的，而InnoDB 的数据存储是按聚簇(cluster)索引有序排列的，主键是默认的聚簇(cluster)索引，因此 MyISAM 虽然在一般情况下查询性能比 InnoDB 高，但 InnoDB 的以主键为条件的查询性能是非常高的。  
 
规则 2：  
命名规则。  
> 1. 数据库和表名应尽可能和所服务的业务模块名一致  
> 2. 服务与同一个子模块的一类表应尽量以子模块名(或部分单词)为前缀或后缀  
> 3. 表名应尽量包含与所存放数据对应的单词  
> 4. 字段名称也应尽量保持和实际数据相对应  
> 5. 联合索引名称应尽量包含所有索引键字段名或缩写，且各字段名在索引名中的顺序应与索引键在索引中的索引顺序一致，并尽量包含一个类似idx的前缀或后缀，以表明期对象类型是索引  
> 6. 约束等其他对象也应该尽可能包含所属表或其他对象的名称，以表明各自的关系  
 
规则 3：  
数据库字段类型定义。  
> 1. 经常需要计算和排序等消耗 CPU 的字段,应该尽量选择更为迅速的字段，如用 TIMESTAMP(4 个字节，最小值 1970-01-01 00:00:00)代替Datetime（8 个字节，最小值 1001-01-01 00:00:00）,通过整型替代浮点型和字符型  
> 2. 变长字段使用 varchar，不要使用 char  
> 3. 对于二进制多媒体数据，流水队列数据(如日志)，超大文本数据不要放在数据库字段中  
 
规则 4：  
业务逻辑执行过程必须读到的表中必须要有初始的值，避免业务读出为负或无穷大的值导致程序失败。  
 
规则 5：  
并不需要一定遵守范式理论，适度的冗余，让 Query 尽量减少 Join。  
 
规则 6：
访问频率较低的大字段拆分出数据表。  
有些大字段占用空间多，访问频率较其他字段明显要少很多，这种情况进行拆分，频繁的查询中就不需要读取大字段，造成 IO 资源的浪费。  
 
规则 7：  
大表可以考虑水平拆分。  
大表影响查询效率，根据业务特性有很多拆分方式，像根据时间递增的数据，可以根据时间来分。以 id 划分的数据，可根据 id % 数据库个数的方式来拆分。  
 
规则 8：  
业务需要的相关索引是根据实际的设计所构造 sql 语句的 where 条件来确定的，业务不需要的不要建索引，不允许在联合索引（或主键）中存在多于的字段。特别是该字段根本不会在条件语句中出现。  
 
规则 9：  
唯一确定一条记录的一个字段或多个字段要建立主键或者唯一索引，不能唯一确定一条记录，为了提高查询效率建普通索引。  
 
 
规则 10：  
业务使用的表，有些记录数很少，甚至只有一条记录，为了约束的需要，也要建立索引或者设置主键。  
 
规则 11：  
对于取值不能重复，经常作为查询条件的字段，应该建唯一索引(主键默认唯一索引)，并且将查询条件中该字段的条件置于第一个位置。没有必要再建立与该字段有关的联合索引。  
 
规则 12：  
对于经常查询的字段，其值不唯一，也应该考虑建立普通索引，查询语句中该字段条件置于第一个位置，对联合索引处理的方法同样。  
 
规则 13：  
业务通过不唯一索引访问数据时，需要考虑通过该索引值返回的记录稠密度，原则上可能的稠密度最大不能高于 0.2，如果稠密度太大，则不合适建立索引了。  
当通过这个索引查找得到的数据量占到表内所有数据的 20% 以上时，则需要考虑建立该索引的代价，同时由于索引扫描产生的都是随机 I/O，生其效率比全表顺序扫描的顺序 I/O 低很多。数据库系统优化 query 的时候有可能不会用到这个索引。  
 
规则 14：  
需要联合索引(或联合主键)的数据库要注意索引的顺序。SQL语句中的匹配条件也要跟索引的顺序保持一致。  
注意：  
索引的顺势不正确也可能导致严重的后果。  
 
规则 15：  
表中的多个字段查询作为查询条件，不含有其他索引，并且字段联合值不重复，可以在这多个字段上建唯一的联合索引，假设索引字段为 (a1,a2,...an),则查询条件(a1 op val1,a2 op val2,...am op valm)m<=n,可以用到索引，查询条件中字段的位置与索引中的字段位置是一致的。  
 
规则 16：  
联合索引的建立原则(以下均假设在数据库表的字段a,b,c上建立联合索引(a,b,c))  
> 1. 联合索引中的字段应尽量满足过滤数据从多到少的顺序，也就是说差异最大的字段应该房子第一个字段  
> 2. 建立索引尽量与SQL语句的条件顺序一致，使SQL语句尽量以整个索引为条件，尽量避免以索引的一部分(特别是首个条件与索引的首个字段不一致时)作为查询的条件  
> 3. Where a=1,where a>=12 and a<15,where a=1 and b<5 ,where a=1 and b=7 and c>=40为条件可以用到此联合索引；而这些语句where b=10,where c=221,where b>=12 and c=2则无法用到这个联合索引。  
> 4. 当需要查询的数据库字段全部在索引中体现时，数据库可以直接查询索引得到查询信息无须对整个表进行扫描(这就是所谓的key-only)，能大大的提高查询效率。  
当 a，ab，abc 与其他表字段关联查询时可以用到索引  
> 5. 当 a，ab，abc 顺序而不是 b，c，bc，ac 为顺序执行 order by 或者 group 不要时可以用到索引  
> 6. 以下情况时，进行表扫描然后排序可能比使用联合索引更加有效  
> a.表已经按照索引组织好了  
> b.被查询的数据站所有数据的很多比例。  
 
规则 17：  
重要业务访问数据表时。但不能通过索引访问数据时，应该确保顺序访问的记录数目是有限的，原则上不得多于 10。  
 
规则 18：  
合理构造 Query 语句。  
> 1. Insert 语句中，根据测试，批量一次插入 1000 条时效率最高，多于 1000 条时，要拆分，多次进行同样的插入，应该合并批量进行。注意query 语句的长度要小于 mysqld 的参数 max_allowed_packet。  
> 2. 查询条件中各种逻辑操作符性能顺序是 and,or,in，因此在查询条件中应该尽量避免使用在大集合中使用 in  
> 3. 永远用小结果集驱动大记录集，因为在 mysql 中，只有 Nested Join 一种 Join 方式，就是说 mysql 的 join 是通过嵌套循环来实现的。通过小结果集驱动大记录集这个原则来减少嵌套循环的循环次数，以减少 IO 总量及 CPU 运算次数。  
> 4. 尽量优化 Nested Join 内层循环。  
> 5. 只取需要的 columns，尽量不要使用 select *  
> 6. 仅仅使用最有效的过滤字段，where 字句中的过滤条件少为好  
> 7. 尽量避免复杂的 Join 和子查询  
> Mysql 在并发这块做得并不是太好，当并发量太高的时候，整体性能会急剧下降，这主要与 Mysql 内部资源的争用锁定控制有关，MyIsam 用表锁，InnoDB 好一些用行锁。  
 
规则 19：  
应用系统的优化。  
> 1. 合理使用 cache，对于变化较少的部分活跃数据通过应用层的 cache 缓存到内存中，对性能的提升是成数量级的。  
> 2. 对重复执行相同的 query 进行合并，减少 IO 次数。  
> 3. 事务相关性最小原则。  

### 三十六条军规
#### 一、核心军规(5)
#### 1.1 尽量不在数据库做运算

* 别让脚趾头想事情，那是脑瓜子的职责
* 让数据库多做她擅长的事:
    * 尽量不在数据库做运算
    * 复杂运算秱到程序端CPU
    * 尽可能简单应用MySQL
* 举例: md5() / Order by Rand()

#### 1.2 控制单表数据量

* 一年内的单表数据量预估
    * 纯INT不超1000W
    * 含CHAR不超500W
* 合理分表不超载
    * USERID
    * DATE
    * AREA
    * ......
* 建议单库不超过300-400个表

#### 1.3 保持表身段苗条

* 表字段数少而精
    * IO高效
    * 全表遍历
    * 表修复快
    * 提高幵发
    * alter table快
* 单表多少字段合适?
* 单表1G体积 500W行评估
    * 顺序读1G文件需N秒
    * 单行不超过200Byte
    * 单表不超过50个纯INT字段
    * 单表不超过20个CHAR(10)字段
* 单表字段数上限控制在20~50个

#### 1.4 平衡范式不冗余

* 严格遵循三大范式?
* 效率优先、提升性能
* 没有绝对的对不错
* 适当时牺牲范式、加入冗余
* 但会增加代码复杂度

#### 1.5 拒绝3B

* 数据库幵发像城市交通
    * 非线性增长
* 拒绝3B
    * 大SQL (BIG SQL)
    * 大事务 (BIG Transaction)
    * 大批量 (BIG Batch)
* 详细解析见后

#### 1.6 核心军规小结

* 尽量不在数据库做运算
* 控制单表数据量
* 保持表身段苗条
* 平衡范式不冗余
* 拒绝3B

#### 二、字段类军规(6)

#### 2.1 用好数值字段类型

* 三类数值类型:
    * TINYINT(1Byte)
    * SMALLINT(2B)
    * MEDIUMINT(3B)
    * INT(4B)、BIGINT(8B)
    * FLOAT(4B)、DOUBLE(8B)
    * DECIMAL(M,D)
* BAD CASE:
    * INT(1) VS INT(11)
    * BIGINT AUTO_INCREMENT
    * DECIMAL(18,0)

#### 2.2 将字符转化为数字

* 数字型VS字符串型索引
    * 更高效
    * 查询更快
    * 占用空间更小
* 举例:用无符号INT存储IP，而非CHAR(15)
    * INT UNSIGNED
    * INET_ATON()
    * INET_NTOA()

#### 2.3 优先使用ENUM或SET

* 优先使用ENUM或SET
    * 字符串
    * 可能值已知且有限
* 存储
    * ENUM占用1字节，转为数值运算
    * SET视节点定，最多占用8字节
    * 比较时需要加' 单引号(即使是数值)
* 举例
    * `sex` enum('F','M') COMMENT '性别'
    * `c1` enum('0','1','2','3') COMMENT '职介审核'

#### 2.4 避免使用NULL字段

* 避免使用NULL字段
    * 很难进行查询优化
    * NULL列加索引，需要额外空间
    * 含NULL复合索引无效
* 举例
    * `a` char(32) DEFAULT NULL
    * `b` int(10) NOT NULL
    * `c` int(10) NOT NULL DEFAULT 0

#### 2.5 少用并拆分TEXT/BLOB

* TEXT类型处理性能远低亍VARCHAR
    * 强制生成硬盘临时表
    * 浪费更多空间
    * VARCHAR(65535)==>64K (注意UTF-8)
* 尽量不用TEXT/BLOB数据类型
* 若必须使用则拆分到单独的表
* 举例:

``` sql
CREATE TABLE t1 (
id INT NOT NULL AUTO_INCREMENT, data text NOT NULL,
‏PRIMARY KEY id
) ENGINE=InnoDB;
```

#### 2.6 不在数据库里存图片

#### 2.7 字段类军规小结

* 用好数值字段类型
* 将字符转化为数字
* 优先使用枚举ENUM/SET
* 避免使用NULL字段
* 少用幵拆分TEXT/BLOB
* 不在数据库里存图片

#### 三、索引类军规(5)

#### 3.1 谨慎合理添加索引

* 谨慎合理添加索引
    * 改善查询
    * 减慢更新
    * 索引不是赹多赹好
* 能不加的索引尽量不加
    * 综合评估数据密度和数据分布
    * 最好不赸过字段数20%
* 结合核心SQL优先考虑覆盖索引
* 举例
    * 不要给“性别”列创建索引

#### 3.2 字符字段必须建前缀索引

* 区分度
    * 单字母区分度:26
    * 4字母区分度:26*26*26*26=456,976
    * 5字母区分度:26*26*26*26*26=11,881,376
    * 6字母区分度:26*26*26*26*26*26=308,915,776
* 字符字段必须建前缀索引:

``` sql
(
`pinyin` varchar(100) DEFAULT NULL COMMENT '小区拼音', KEY `idx_pinyin` (`pinyin`(8)),
) ENGINE=InnoDB
```

#### 3.3 不在索引列做运算

* 不在索引列进行数学运算或凼数运算
    * 无法使用索引
    * 导致全表扫描
* 举例:

``` sql
BAD: SELECT * from table WHERE to_days(current_date) – to_days(date_col) <= 10
GOOD: SELECT * from table WHERE date_col >= DATE_SUB('2011-10- 22',INTERVAL 10 DAY);
```

#### 3.4 自增列或全局ID做INNODB主键

* 对主键建立聚簇索引
* 二级索引存储主键值
* 主键不应更新修改
* 按自增顺序揑入值
* 忌用字符串做主键
* 聚簇索引分裂
* 推荐用独立亍业务的AUTO_INCREMENT列或全局ID生成 器做代理主键
* 若不指定主键，InnoDB会用唯一且非空值索引代替

#### 3.5 尽量不用外键

* 线上OLTP系统(线下系统另论)
    * 外键可节省开发量
    * 有额外开销
    * 逐行操作
    * 可'到达'其它表，意味着锁
    * 高并发时容易死锁
* 由程序保证约束

#### 3.6 索引类军规小结

* 谨慎合理添加索引
* 字符字段必须建前缀索引
* 不在索引列做运算
* 自增列或全局ID做INNODB主键
* 尽量不用外键

#### 四、SQL类军规(15)

#### 4.1 SQL语句尽可能简单

* 大SQL VS 多个简单SQL
    * 传统设计思想
    * BUT MySQL NOT
    * 一条SQL叧能在一个CPU运算
    * 5000+ QPS的高幵发中，1秒大SQL意味着?
    * 可能一条大SQL就把整个数据库堵死
* 拒绝大SQL，拆解成多条简单SQL
    * 简单SQL缓存命中率更高
    * 减少锁表时间，特别是MyISAM
    * 用上多CPU

#### 4.2 保持事务(连接)短小

* 保持事务/DB连接短小精悍
    * 事务/连接使用原则:即开即用，用完即关
    * 与事务无关操作放到事务外面, 减少锁资源的占用
    * 不破坏一致性前提下，使用多个短事务代替长事务
* 举例
    * 发贴时的图片上传等待
    * 大量的sleep连接

#### 4.3 尽可能避免使用SP/TRIG/FUNC

* 线上OLTP系统(线下库另论)
    * 尽可能少用存储过程
    * 尽可能少用触发器
    * 减用使用MySQL凼数对结果进行处理
* 由客户端程序负责

#### 4.4 尽量不用 SELECT *

* 用SELECT * 时
* 更多消耗CPU、内存、IO、网络带宽
* 先向数据库请求所有列，然后丢掉不需要列?
* 尽量不用SELECT * ，叧取需要数据列 • 更安全的设计:减少表变化带来的影响
* 为使用covering index提供可能性
* SELECT/JOIN减少硬盘临时表生成，特别是有TEXT/BLOB时
* 举例:

``` sql
SELECT * FROM tag WHERE id = 999184;
SELECT keyword FROM tag WHERE id = 999184;
```

#### 4.5 改写OR为IN()

* 同一字段，将or改写为in()
* OR效率:O(n)
* IN 效率:O(Log n)
* 当n很大时，OR会慢很多
* 注意控制IN的个数，建议n小亍200
* 举例:

``` sql
SELECT * from opp WHERE phone='12347856' or phone='42242233' \G;
SELECT * from opp WHERE phone in ('12347856' , '42242233');
```

#### 4.6 改写OR为UNION

* 不同字段，将or改为union
* 减少对不同字段进行 "or" 查询
* Merge index往往很弱智
* 如果有足够信心:set global optimizer_switch='index_merge=off';
* 举例:

``` sql
SELECT * from opp WHERE phone='010-88886666' or cellPhone='13800138000';
SELECT * from opp WHERE phone='010-88886666' union SELECT * from opp WHERE cellPhone='13800138000';
```

#### 4.7 避免负向查询和% 前缀模糊查询

* 避免负向查询
    * NOT、!=、<>、!<、!>、NOT EXISTS、NOT IN、 NOT LIKE等
* 避免 % 前缀模糊查询
    * B+ Tree
    * 使用不了索引
    * 导致全表扫描
* 举例:

``` sql
SELECT * from post WHERE title like '北京%'; -- 298 rows in set (0.01 sec)
SELECT * from post WHERE title like '%北京%'; -- 572 rows in set (3.27 sec)
```

#### 4.8 COUNT(*)的几个例子

* 几个有趣的例子:
    * COUNT(COL) VS COUNT(*)
    * COUNT(*) VS COUNT(1)
    * COUNT(1) VS COUNT(0) VS COUNT(100)
* 示例:

``` sql
`id` int(10) NOT NULL AUTO_INCREMENT COMMENT '公司的id', `sale_id` int(10) unsigned DEFAULT NULL,
```

* 结论
    * COUNT(*)=count(1)
    *COUNT(0)=count(1)
    * COUNT(1)=count(100)
    * COUNT(*)!=count(col)
    * WHY?

#### 4.9 减少COUNT(*)

* MyISAM VS INNODB
    * 不带 WHERE COUNT()
    * 带 WHERE COUNT()
* COUNT(*)的资源开销大，尽量不用少用
* 计数统计
    * 实时统计:用memcache，双向更新，凌晨 跑基准
    * 非实时统计:尽量用单独统计表，定期重算

#### 4.10 LIMIT高效分页

* 传统分页:
    * SELECT * from table limit 10000,10;
* LIMIT原理:
    * Limit 10000,10  偏秱量赹大则赹慢
* 推荐分页:
    * SELECT * from table WHERE id>=23423 limit 11;
    *SELECT * from table WHERE id>=23434 limit 11;
* 分页方式二:
    * SELECT * from table WHERE id >= ( SELECT id from table limit 10000,1 ) limit 10;
* 分页方式三:
    * SELECT * FROM table INNER JOIN (SELECT id FROM table LIMIT 10000,10) USING (id);
* 分页方式四:
    * 程序取ID:SELECT id from table limit 10000,10;
    * SELECT * from table WHERE id in (123,456...);
* 可能需按场景分析幵重组索引
* 示例:

``` sql
MySQL> SELECT sql_no_cache * from post limit 10,10; 10 row in set (0.01 sec)
MySQL> SELECT sql_no_cache * from post limit 20000,10; 10 row in set (0.13 sec)
MySQL> SELECT sql_no_cache * from post limit 80000,10; 10 rows in set (0.58 sec)
MySQL> SELECT sql_no_cache id from post limit 80000,10; 10 rows in set (0.02 sec)
MySQL> SELECT sql_no_cache * from post WHERE id>=323423 limit 10; 10 rows in set (0.01 sec)
MySQL> SELECT * from post WHERE id >= ( SELECT sql_no_cache id from post limit 80000,1 ) limit 10; 10 rows in set (0.02 sec)
```

#### 4.11 用UNION ALL 而非 UNION

* 若无需对结果进行去重，则用UNION ALL
    * UNION有去重开销
* 举例:

``` sql
SELECT * FROM detail20091128 UNION ALL SELECT * FROM detail20110427 UNION ALL SELECT * FROM detail20110426 UNION ALL SELECT * FROM detail20110425 UNION ALL SELECT * FROM detail20110424 UNION ALL SELECT * FROM detail20110423;
```

#### 4.12 分解联接保证高并发

* 高幵发DB不建议进行两个表以上的JOIN
* 适当分解联接保证高幵发
    * 可缓存大量早期数据
    * 使用了多个MyISAM表
    * 对大表的小ID IN()
    * 联接引用同一个表多次
    * 举例:

``` sql
MySQL> SELECT * from tag JOIN post on tag_post.post_id=post.id WHERE tag.tag='二手玩具';

MySQL> SELECT * from tag WHERE tag='二手玩具';
MySQL> SELECT * from tag_post WHERE tag_id=1321;
MySQL> SELECT * from post WHERE post.id in (123,456,314,141);
```

#### 4.13 GROUP BY 去除排序

* GROUP BY 实现
    * 分组
    * 自劢排序
* 无需排序:Order by NULL
* 特定排序:Group by DESC/ASC
* 举例:

``` sql
MySQL> SELECT phone,count(*) from post group by phone limit 1 ; 1 row in set (2.19 sec)
MySQL> SELECT phone,count(*) from post group by phone order by null limit 1; 1 row in set (2.02 sec)
```

#### 4.14 同数据类型的列值比较

* 原则:数字对数字，字符对字符
* 数值列不字符类型比较
    * 同时转换为双精度
    * 进行比对
* 字符列不数值类型比较
    * 字符列整列转数值
    * 不会使用索引查询
* 举例:字符列不数值类型比较

``` sql
字段:`remark` varchar(50) NOT NULL COMMENT '备注, 默认为空',

MySQL>SELECT `id`, `gift_code` FROM gift WHERE `deal_id` = 640 AND remark=115127; 1 row in set (0.14 sec)
MySQL>SELECT `id`, `gift_code` FROM pool_gift WHERE `deal_id` = 640 AND remark='115127'; 1 row in set (0.005 sec)
```

#### 4.15 Load data 导数据

* 批量数据快导入:
    * 成批装载比单行装载更快，不需要每次刷新缓存
    * 无索引时装载比索引装载更快
    * Insert values ,values，values 减少索引刷新
    * Load data比insert快约20倍
* 尽量不用 INSERT ... SELECT
    * 延迟
    * 同步出错

#### 4.16 打散大批量更新

* 大批量更新凌晨操作，避开高峰
* 凌晨不限制
* 白天上限默认为100条/秒(特殊再议)
* 举例:

``` sql
update post set tag=1 WHERE id in (1,2,3); sleep 0.01;
update post set tag=1 WHERE id in (4,5,6); sleep 0.01;
......
```

#### 4.17 Know Every SQL

* SHOW PROFILE
* MySQLdumpslow
* EXPLAIN
* Show Slow Log
* SHOW QUERY_RESPONSE_TIME(Percona)
* MySQLsla
* Show Processlist

#### 4.18 SQL类军规小结

* SQL语句尽可能简单
* 保持事务(连接)短小
* 尽可能避免使用SP/TRIG/FUNC
* 尽量不用 SELECT *
* 改写OR语句
* 避免负向查询和% 前缀模糊查询
* 减少COUNT(*)
* LIMIT的高效分页
* 用UNION ALL 而非 UNION
* 分解联接保证高幵发
* GROUP BY 去除排序
* 同数据类型的列值比较
* Load data导数据
* 打散大批量更新
* Know Every SQL!

#### 五、约定类军规(5)

#### 5.1 隔离线上线下

* 构建数据库的生态环境
* 开发无线上库操作权限
* 原则:线上连线上，线下连线下
    * 实时数据用real库
    * 模拟环境用sim库
    * 测试用qa库
    * 开发用dev库

#### 5.2 禁止未经DBA确认的子查询

* MySQL子查询
    * 大部分情况优化较差
    * 特别WHERE中使用IN id的子查询  一般可用JOIN改写
* 举例:

``` sql
SELECT * from table1 where id id from table2) in (SELECT insert into table1 (SELECT * from table2); -- 可能导致复制异常
```

#### 5.3 永远不在程序端显式加锁

* 永远不在程序端对数据库显式加锁
    * 外部锁对数据库不可控
    * 高并发发时是灾难
    * 极难调试和排查
* 并发扣款等一致性问题
    * 采用事务
    * 相对值修改
    * Commit前二次较验冲突

#### 5.4 统一字符集为UTF8

* 字符集:
    * MySQL 4.1 以前叧有latin1
    * 为多语言支持增加多字符集
    * 也带来了N多问题
    * 保持简单
* 统一字符集:UTF8
* 校对规则:utf8_general_ci
* 乱码:SET NAMES UTF8

#### 5.5 统一命名规范

* 库表等名称统一用小写
    * Linux VS Windows
    * MySQL库表大小写敏感
    * 字段名的大小写不敏感
* 索引命名默认为“idx_字段名”
* 库名用缩写，尽量在2~7个字母
    * DataSharing ==> ds
* 注意避免用保留字命名
* ......

#### 5.6 注意避免用保留字命名

* 举例:

``` sql
SELECT * from return;
SELECT * from `return`;
```

<details>
<summary><b>MySQL系统关键字</b></summary>

* ADD
* ALL
* ALTER GOTO
* GRANT
* GROUP
* PURGE
* RAID0
* RANGE
* ANALYZE
* AND
* AS HAVING
* HIGH_PRIORIT Y
* HOUR_MICROSEC OND
* READ
* READS
* REAL
* ASC
* ASENSITIVE
* BEFORE HOUR_MINUTE
* HOUR_SECON D
* IF
* REFERENCES
* REGEXP
* RELEASE
* BETWEEN
* BIGINT
* BINARY IGNORE
* IN
* INDEX
* RENAME
* REPEAT
* REPLACE
* BLOB
* BOTH
* BY INFILE
* INNER
* INOUT
* REQUIRE
* RESTRICT
* RETURN
* CALL
* CASCADE
* CASE INSENSITIVE
* INSERT
* INT
* REVOKE
* RIGHT
* RLIKE
* CHANGE
* CHAR
* CHARACTER INT1
* INT2
* INT3
* SCHEMA
* SCHEMAS
* SECOND_MICROSEC OND
* CHECK
* COLLATE
* COLUMN INT4
* INT8
* INTEGER
* SELECT
* SENSITIVE
* SEPARATOR
* CONDITION
* CONNECTION
* CONSTRAINT INTERVAL
* INTO
* IS
* SET
* SHOW
* SMALLINT
* CONTINUE
* CONVERT
* CREATE ITERATE
* JOIN
* KEY
* SPATIAL
* SPECIFIC
* SQL
* CROSS
* CURRENT_DA TE
* CURRENT_TIM KEYS E
* KILL
* LABEL
* SQLEXCEPTION
* SQLSTATE
* SQLWARNING
* CURRENT_TIMESTA MP
* CURRENT_US ER
* CURSOR LEADING
* LEAVE
* LEFT
* SQL_BIG_RESUL T
* SQL_CALC_FOUND_R OWS
* SQL_SMALL_RESULT
* DATABASE
* DATABASES
* DAY_HOUR LIKE
* LIMIT
* LINEAR
* SSL
* STARTING
* STRAIGHT_JOIN
* DAY_MICROSECON D
* DAY_MINUTE
* DAY_SECOND LINES
* LOAD
* LOCALTIME
* TABLE
* TERMINATED
* THEN
* DEC
* DECIMAL
* DECLARE LOCALTIMESTAMP
* LOCK
* LONG
* TINYBLOB
* TINYINT
* TINYTEXT
* DEFAULT
* DELAYED
* DELETE LONGBLOB
* LONGTEXT
* LOOP
* TO
* TRAILING
* TRIGGER
* DESC
* DESCRIBE
* DETERMINISTI LOW_PRIORITY C
* MATCH
* MEDIUMBLOB
* TRUE
* UNDO
* UNION
* DISTINCT
* DISTINCTROW
* DIV MEDIUMINT
* MEDIUMTEXT
* MIDDLEINT
* UNIQUE
* UNLOCK
* UNSIGNED
* DOUBLE
* DROP
* DUAL
* MINUTE_MICROSECO ND
* MINUTE_SECO ND
* MOD
* UPDATE
* USAGE
* USE
* EACH
* ELSE
* ELSEIF MODIFIES
* NATURAL
* NOT
* USING
* UTC_DATE
* UTC_TIME
* ENCLOSED
* ESCAPED
* EXISTS
* NO_WRITE_TO_BINL OG
* NULL
* NUMERIC
* UTC_TIMESTAM P
* VALUES
* VARBINARY
* EXIT
* EXPLAIN
* FALSE ON
* OPTIMIZE
* OPTION
* VARCHAR
* VARCHARACTER
* VARYING
* FETCH
* FLOAT
* FLOAT4 OPTIONALLY
* OR
* ORDER
* WHEN
* WHERE
* WHILE
* FLOAT8
* FOR
* FORCE OUT
* OUTER
* OUTFILE
* WITH
* WRITE
* X509
* FOREIGN
* FROM
* FULLTEXT PRECISION
* PRIMARY
* PROCEDURE
* XOR
* YEAR_MONTH
* ZEROFILL
</details>

#### 5.7 约定类军规小结

* 隔离线上线下
* 禁止未经DBA确认的子查询上线
* 永远不在程序端显式加锁
* 统一字符集为UTF8
* 统一命名规范

### 规范
#### 1.1 数据库命名规范
1、所有数据库对象名称必须使用小写字母并用下划线分割。  
2、所有数据库对象名称禁止使用 MySQL 保留关键字（如果表名中包含关键字查询时，需要将其用单引号括起来）。  
3、数据库对象的命名要能做到见名识意，并且最后不要超过 32 个字符。  
4、临时库表必须以 tmp_ 为前缀并以日期为后缀，备份表必须以 bak_ 为前缀并以日期 ( 时间戳 ) 为后缀。  
5、所有存储相同数据的列名和列类型必须一致（一般作为关联列，如果查询时关联列类型不一致会自动进行数据类型隐式转换，会造成列上的索引失效，导致查询效率降低）。  

#### 1.2 数据库基本设计规范
1、所有表必须使用 InnoDB 存储引擎  
没有特殊要求（即 InnoDB 无法满足的功能如：列存储，存储空间数据等）的情况下，所有表必须使用 InnoDB 存储引擎（MySQL 5.5 之前默认使用 Myisam，5.6 以后默认的为 InnoDB）InnoDB 支持事务，支持行级锁，更好的恢复性，高并发下性能更好。  

2、数据库和表的字符集统一使用 UTF8MB4    
兼容性更好，统一字符集可以避免由于字符集转换产生的乱码，不同的字符集进行比较前需要进行转换会造成索引失效。  

3、所有表和字段都需要添加注释  
使用 comment 从句添加表和列的备注 从一开始就进行数据字典的维护。  

4、尽量控制单表数据量的大小，建议控制在 500 万以内  
500 万并不是 MySQL 数据库的限制，过大会造成修改表结构、备份、恢复都会有很大的问题，可以用历史数据归档（应用于日志数据），分库分表（应用于业务数据）等手段来控制数据量大小。  

5、谨慎使用 MySQL 分区表  
分区表在物理上表现为多个文件，在逻辑上表现为一个表 谨慎选择分区键，跨分区查询效率可能更低 建议采用物理分表的方式管理大数据。  

6、尽量做到冷热数据分离，减小表的宽度  
MySQL 限制每个表最多存储 4096 列，并且每一行数据的大小不能超过 65535 字节 减少磁盘 IO，保证热数据的内存缓存命中率（表越宽，把表装载进内存缓冲池时所占用的内存也就越大,也会消耗更多的 IO） 更有效的利用缓存，避免读入无用的冷数据 经常一起使用的列放到一个表中（避免更多的关联操作）  

7、禁止在表中建立预留字段  
预留字段的命名很难做到见名识义  
预留字段无法确认存储的数据类型，所以无法选择合适的类型 对预留字段类型的修改，会对表进行锁定  

8、禁止在数据库中存储图片，文件等大的二进制数据
通常文件很大，会短时间内造成数据量快速增长，数据库进行数据库读取时，通常会进行大量的随机 IO 操作，文件很大时，IO 操作很耗时   通常存储于文件服务器，数据库只存储文件地址信息。  

9、禁止在线上做数据库压力测试  

10、禁止从开发环境，测试环境直接连接生成环境数据库  

#### 1.3 数据库字段设计规范
1、优先选择符合存储需要的最小的数据类型  
原因：列的字段越大，建立索引时所需要的空间也就越大，这样一页中所能存储的索引节点的数量也就越少也越少，在遍历时所需要的 IO 次数也就越多， 索引的性能也就越差  
方法：
① 将字符串转换成数字类型存储，如：将IP地址转换成整形数据。
② 对于非负型的数据（如自增 ID、整型 IP）来说，要优先使用无符号整型来存储，因为无符号相对于有符号可以多出一倍的存储空间。  

2、避免使用 TEXT、BLOB 数据类型，最常见的 TEXT 类型就可以存储 64k 的数据  
建议把 BLOB 或是TEXT列分离到单独的扩展表中，TEXT 或 BLOB 类型只能使用前缀索引  

3、避免使用 ENUM 类型  
原因：修改 ENUM 值需要使用 ALTER 语句；ENUM 类型的 ORDER BY 操作效率低，需要额外操作；禁止使用数值作为 ENUM 的枚举值  

4、尽可能把所有列定义为 NOT NULL  
原因：索引 NULL 列需要额外的空间来保存，所以要占用更多的空间；进行比较和计算时要对 NULL 值做特别的处理。  

5、使用 TIMESTAMP（4 个字节）或 DATETIME 类型（8 个字节）存储时间  
TIMESTAMP 存储的时间范围 1970-01-01 00:00:01 ~ 2038-01-19-03:14:07。  
TIMESTAMP 占用 4 字节和 INT 相同，但比 INT 可读性高，超出 TIMESTAMP 取值范围的使用 DATETIME 类型存储。  

6、同财务相关的金额类数据必须使用 decimal 类型  
Decimal 类型为精准浮点数，在计算时不会丢失精度。占用空间由定义的宽度决定，每 4 个字节可以存储 9 位数字，并且小数点要占用一个字节。可用于存储比 bigint 更大的整型数据。  

#### 1.4 索引设计规范
1、限制每张表上的索引数量，建议单张表索引不超过 5 个  
索引并不是越多越好！索引可以提高效率同样也可以降低效率；索引可以增加查询效率，但同样也会降低插入和更新的效率，甚至有些情况下会降低查询效率。  
因为 MySQL 优化器在选择如何优化查询时，会根据统一信息，对每一个可以用到的索引来进行评估，以生成出一个最好的执行计划，如果同时有很多个索引都可以用于查询，就会增加 MySQL 优化器生成执行计划的时间，同样会降低查询性能。  

2、禁止给表中的每一列都建立单独的索引  
5.6 版本之前，一个 SQL 只能使用到一个表中的一个索引，5.6 以后，虽然有了合并索引的优化方式，但是还是远远没有使用一个联合索引的查询方式好。  

3、每个 InnoDB 表必须有个主键  
InnoDB 是一种索引组织表：数据的存储的逻辑顺序和索引的顺序是相同的。每个表都可以有多个索引，但是表的存储顺序只能有一种 InnoDB 是按照主键索引的顺序来组织表的。

不要使用更新频繁的列作为主键，不适用多列主键（相当于联合索引） 不要使用 UUID、MD5、HASH、字符串列作为主键（无法保证数据的顺序增长）。主键建议使用自增 ID 值。

4、常见的索引列建议  
出现在 SELECT、UPDATE、DELETE 语句的 WHERE 从句中的列。  
包含在 ORDER BY、GROUP BY、DISTINCT 中的字段。  
并不要将符合以上字段的列都建立一个索引，通常将字段建立联合索引效果更好。  
多表 JOIN 的关联列。  

5、选择索引列的顺序  
建立索引的目的是：希望通过索引进行数据查找，减少随机 IO，增加查询性能 ，索引能过滤出越少的数据，则从磁盘中读入的数据也就越少。  
➀ 区分度最高的放在联合索引的最左侧（区分度 = 列中不同值的数量 / 列的总行数）  
➁ 尽量把字段长度小的列放在联合索引的最左侧（因为字段长度越小，一页能存储的数据量越大，IO 性能也就越好）  
➂ 使用最频繁的列放到联合索引的左侧（这样可以比较少的建立一些索引）。  

6、避免建立冗余索引和重复索引  
因为这样会增加查询优化器生成执行计划的时间。  
重复索引示例：primary key(id)、index(id)、unique index(id)  
冗余索引示例：index(a,b,c)、index(a,b)、index(a）  

7、优先考虑覆盖索引  
覆盖索引就是包含了所有查询字段(where,select,ordery by,group by 包含的字段)的索引。  
对于频繁的查询优先考虑使用覆盖索引。  

8、避免使用外键约束  
不建议使用外键约束（foreign key），但一定要在表与表之间的关联键上建立索引。  
外键可用于保证数据的参照完整性，但建议在业务端实现。  
外键会影响父表和子表的写操作从而降低性能。  

#### 1.5 数据库 SQL 开发规范
1、建议使用预编译语句进行数据库操作  
预编译语句可以重复使用这些计划，减少 SQL 编译所需要的时间，还可以解决动态 SQL 所带来的 SQL 注入的问题；只传参数，比传递 SQL 语句更高效 相同语句可以一次解析，多次使用，提高处理效率。  

2、避免数据类型的隐式转换  
隐式转换会导致索引失效。如：select name,phone from customer where id = '111';  

3、充分利用表上已经存在的索引  
避免使用双 % 号的查询条件，如 a like '%123%'，（如果无前置 %，只有后置 %，是可以用到列上的索引的）；  
一个 SQL 只能利用到复合索引中的一列进行范围查询，如：有 a,b,c 列的联合索引，在查询条件中有 a 列的范围查询，则在 b,c 列上的索引将不会被用到，在定义联合索引时，如果a列要用到范围查找的话，就要把 a 列放到联合索引的右侧。  
使用 left join 或 not exists 来优化 not in 操作，因为 not in 也通常会使用索引失效。  

4、数据库设计时，应该要对以后扩展进行考虑  

5、程序连接不同的数据库使用不同的账号，禁止跨库查询  
为数据库迁移和分库分表留出余地；降低业务耦合度；避免权限过大而产生的安全风险。  

6、禁止使用 SELECT * 必须使用 SELECT <字段列表> 查询  
消耗更多的 CPU 和 IO 以网络带宽资源；可能无法使用覆盖索引；可减少表结构变更带来的影响  

7、禁止使用不含字段列表的 INSERT 语句  
如：insert into values ('a','b','c');   
应该：insert into t(c1,c2,c3) values ('a','b','c');  

8、避免使用子查询，可以把子查询优化为 JOIN 操作  
通常子查询在 in 子句中，且子查询中为简单 SQL ( 不包含 union、group by、order by、limit 从句 ) 时，才可以把子查询转化为关联查询进行优化。  
子查询性能差的原因：
子查询的结果集无法使用索引，通常子查询的结果集会被存储到临时表中，不论是内存临时表还是磁盘临时表都不会存在索引，所以查询性能会受到一定的影响。特别是对于返回结果集比较大的子查询，其对查询性能的影响也就越大。由于子查询会产生大量的临时表也没有索引，所以会消耗过多的 CPU 和 IO 资源，产生大量的慢查询。  

9、 避免使用 JOIN 关联太多的表  
对于 MySQL 来说，是存在关联缓存的，缓存的大小可以由 join_buffer_size 参数进行设置。  
在 MySQL 中，对于同一个 SQL 多关联（join）一个表，就会多分配一个关联缓存，如果在一个 SQL 中关联的表越多，所占用的内存也就越大。  

如果程序中大量的使用了多表关联的操作，同时 join_buffer_size 设置的也不合理的情况下，就容易造成服务器内存溢出的情况，就会影响到服务器数据库性能的稳定性。

同时对于关联操作来说，会产生临时表操作，影响查询效率 MySQL 最多允许关联 61 个表，建议不超过 5 个。

10、减少同数据库的交互次数  
数据库更适合处理批量操作，合并多个相同的操作到一起，可以提高处理效率  

11、对应同一列进行 or 判断时，使用 in 代替 or  
In 的值不要超过 500 个， in 操作可以更有效的利用索引，or 大多数情况下很少能利用到索引。  

12、禁止使用 order by rand() 进行随机排序  
会把表中所有符合条件的数据装载到内存中，然后在内存中对所有数据根据随机生成的值进行排序，并且可能会对每一行都生成一个随机值，如果满足条件的数据集非常大，就会消耗大量的 CPU 和 IO 及内存资源。  
推荐在程序中获取一个随机值，然后从数据库中获取数据的方式。  

13、WHERE 从句中禁止对列进行函数转换和计算  
对列进行函数转换或计算时会导致无法使用索引。  
不推荐：where date(create_time)='20190101'  
推荐：where create_time >= '20190101' and create_time < '20190102'  

14、在明显不会有重复值时使用 UNION ALL 而不是 UNION  
UNION 会把两个结果集的所有数据放到临时表中后再进行去重操作，UNION ALL 不会再对结果集进行去重操作。  

15、拆分复杂的大 SQL 为多个小 SQL  
大 SQL：逻辑上比较复杂，需要占用大量 CPU 进行计算的 SQL。  
小 SQL：一个 SQL 只能使用一个 CPU 进行计算。  
SQL 拆分后可以通过并行执行来提高处理效率。  

#### 1.6 数据库操作行为规范
1、超 100 万行的批量写（UPDATE、DELETE、INSERT）操作，要分批多次进行操作  
大批量操作可能会造成严重的主从延迟  
主从环境中，大批量操作可能会造成严重的主从延迟，大批量的写操作一般都需要执行一定长的时间，而只有当主库上执行完成后，才会在其他从库上执行，所以会造成主库与从库长时间的延迟情况。  

Binlog 日志为 row 格式时会产生大量的日志  
大批量写操作会产生大量日志，特别是对于 row 格式二进制数据而言，由于在 row 格式中会记录每一行数据的修改，我们一次修改的数据越多，产生的日志量也就会越多，日志的传输和恢复所需要的时间也就越长，这也是造成主从延迟的一个原因。  

避免产生大事务操作  
大批量修改数据，一定是在一个事务中进行的，这就会造成表中大批量数据进行锁定，从而导致大量的阻塞，阻塞会对 MySQL 的性能产生非常大的影响。
特别是长时间的阻塞会占满所有数据库的可用连接，这会使生产环境中的其他应用无法连接到数据库，因此一定要注意大批量写操作要进行分批。  

2、对于大表使用 pt-online-schema-change 修改表结构  
主要是避免大表修改产生的主从延迟，避免在对表字段进行修改时进行锁表。  

对大表数据结构的修改一定要谨慎，会造成严重的锁表操作，尤其是生产环境，是不能容忍的。  
pt-online-schema-change 会首先建立一个与原表结构相同的新表，并且在新表上进行表结构的修改，然后再把原表中的数据复制到新表中，并在原表中增加一些触发器。  

把原表中新增的数据也复制到新表中，在行所有数据复制完成之后，把新表命名成原表，并把原来的表删除掉,把原来一个 DDL 操作，分解成多个小的批次进行。

3、禁止为程序使用的账号赋予 super 权限  
当达到最大连接数限制时，还运行 1 个 有 super 权限的用户连接 super 权限只能留给 DBA 处理问题的账号使用。  

4、对于程序连接数据库账号，遵循权限最小原则  
程序使用数据库账号只能在一个 DB 下使用，不准跨库 程序使用的账号原则上不准有 drop 权限。  

