
### 连接查询及原理
当我们查询的数据来源于多张表的时候，我们需要用到连接查询，连接查询使用率非常高。  

**笛卡尔积**  
笛卡尔积：有两个集合 A 和 B，笛卡尔积表示 A 集合中的元素和 B 集合中的元素任意相互关联产生的所有可能的结果。  
假如 A 中有 m 个元素，B 中有 n 个元素，A、B 笛卡尔积产生的结果有 m\*n 个结果，相当于循环遍历两个集合中的元素，任意组合。

sql 中笛卡尔积语法：
```sql
select 字段 from 表1,表2[,表N];
# 或者
select 字段 from 表1 join 表2 [join 表N];
```

**内连接**  
内连接相当于在笛卡尔积的基础上加上了连接的条件，当没有连接条件的时候，内连接上升为笛卡尔积。  

sql 中内连接语法：  
```sql
select 字段 from 表1 inner join 表2 on 连接条件;
# 或
select 字段 from 表1 join 表2 on 连接条件;
# 或
select 字段 from 表1, 表2 [where 关联条件]; # 简洁，建议使用
```

**外连接**  
`外连接查询结果 = 内连接的结果 + 主表中有的而内连接结果中没有的记录。`  
外连接涉及到 2 个表，分为：主表和从表，要查询的信息主要来自于哪个表，谁就是主表。  
外连接查询结果为主表中所有记录。如果从表中有和它匹配的，则显示匹配的值，这部分相当于内连接查询出来的结果；如果从表中没有和它匹配的，则显示 null。  

外连接分为 2 种：  
左外链接：使用 left join 关键字，left join 左边的是主表。  
右外连接：使用 right join 关键字，right join 右边的是主表。  

**左连接**  
sql 中左连接语法：  
```sql
select 列 from 主表 left join 从表 on 连接条件;
```

**右连接**  
sql 中右连接语法：  
```sql
select 列 from 从表 right join 主表 on 连接条件;
```

**表连接原理**  
在 MySQL 中，A left join B on condition 的执行过程如下：

1）以 table_A 为驱动表，检索 table_B  
2）根据 on 条件过滤 table_B 的数据，构建 table_A 结果集，并且添加外部行  
3）对结果集执行 where 条件过滤。如果 A 中有一行匹配 where 子句但是 B 中没有一行匹配 on 条件，则生成另一个 B 行，其中所有列设置为 NULL  
4）执行 group by 语句分组  
5）执行 having 语句对分组结果筛选  
6）执行 select 出结果集  
7）执行 distinct 对结果去重  
8）执行 order by 语句  
9）执行 limit 语句  

MySQL 会先进行连接查询，然后再使用 where 子句查询结果，再从结果执行 order by。所以如果被驱动表数据过大，会造成检索行过多。可以利用子查询先查询出一个较小的结果集，然后再用连接驱动。  
right join 的执行类似 left join ，只是表的角色相反。  

```sql
# 准备数据
drop table if exists test1;
create table test1(
  a int
);
drop table if exists test2;
create table test2(
  b int
);
insert into test1 values (1),(2),(3);
insert into test2 values (3),(4),(5);
mysql> select * from test1;
+------+
| a    |
+------+
|    1 |
|    2 |
|    3 |
+------+
3 rows in set (0.00 sec)

mysql> select * from test2;
+------+
| b    |
+------+
|    3 |
|    4 |
|    5 |
+------+
3 rows in set (0.00 sec)

# 内连接
mysql> select * from test1 t1,test2 t2;
+------+------+
| a    | b    |
+------+------+
|    1 |    3 |
|    2 |    3 |
|    3 |    3 |
|    1 |    4 |
|    2 |    4 |
|    3 |    4 |
|    1 |    5 |
|    2 |    5 |
|    3 |    5 |
+------+------+
9 rows in set (0.00 sec)

mysql> select * from test1 t1,test2 t2 where t1.a = t2.b;
+------+------+
| a    | b    |
+------+------+
|    3 |    3 |
+------+------+
1 row in set (0.00 sec)

# 左连接
mysql> select * from test1 t1 left join test2 t2 on t1.a = t2.b;
+------+------+
| a    | b    |
+------+------+
|    3 |    3 |
|    1 | NULL |
|    2 | NULL |
+------+------+
3 rows in set (0.00 sec)

# t1.a>10 只关联了 test1 表
mysql> select * from test1 t1 left join test2 t2 on t1.a>10;
+------+------+
| a    | b    |
+------+------+
|    1 | NULL |
|    2 | NULL |
|    3 | NULL |
+------+------+
3 rows in set (0.00 sec)

# 连接条件 1=1 值为 true，返回结果为笛卡尔积
mysql> select * from test1 t1 left join test2 t2 on 1=1;
+------+------+
| a    | b    |
+------+------+
|    1 |    3 |
|    2 |    3 |
|    3 |    3 |
|    1 |    4 |
|    2 |    4 |
|    3 |    4 |
|    1 |    5 |
|    2 |    5 |
|    3 |    5 |
+------+------+
9 rows in set (0.00 sec)
```
