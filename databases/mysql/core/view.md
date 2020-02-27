
### 视图
视图是在 mysql5 之后出现的，是一种虚拟表，行和列的数据来自于定义视图时使用的一些表中。   
`视图的数据是在使用视图的时候动态生成的，视图只保存了 sql 的逻辑，不保存查询的结果。`

视图是一种虚拟的表，具有和物理表相同的功能，可以对视图进行增、改、查操作。  
视图通常是一个表或者多个表的行或列的子集。  
视图与包含数据的表不一样，视图只包含使用时动态检索数据的查询；不包含任何列或数据。使用视图可以简化复杂的 sql 操作，隐藏具体的细节，保护数据；视图创建后，可以使用与表相同的方式利用它们。  

视图不能被索引，也不能有关联的触发器或默认值，如果视图本身内有 order by 则对视图再次 order by 将被覆盖。  

对于某些视图比如未使用联结子查询分组聚集函数 Distinct Union 等，是可以对其更新的，对视图的更新将对基表进行更新；但是视图主要用于简化检索，保护数据，并不用于更新，而且大部分视图都不可以更新。  

**视图的好处**  
获取数据更容易，相对于多表查询来说；简化复杂的 sql 操作，不用知道他的实现细节。  
视图能够对机密数据提供安全保护；隔离了原始表，可以不让使用视图的人接触原始的表，从而保护原始数据，提高了安全性。  
视图的修改不会影响基本表，提供了独立的操作单元，比较轻量。 

**使用场景**  
多个地方使用到同样的查询结果，并且该查询结果比较复杂的时候，我们可以使用视图来隐藏复杂的实现细节。  

**视图和表的区别**  


| | 语法 | 实际中是否占用物理空间 | 使用 |  
| :---: |  :---: |  :---: |  :---: |  
| 视图 | create view	 |只是保存了 sql 的逻辑 | 增删改查，实际上我们只使用查询 |  
|表	| create table | 保存了数据 |增删改查 |  

**创建视图**  
格式如下：  
```sql
create view 视图名
as
查询语句;
```

**修改视图**  
修改视图，有 2 种方式。  

方式 1：如果该视图存在，就修改，如果不存在，就创建新的视图。  
```sql
create or replace view 视图名
as
查询语句;
```

方式 2：直接修改视图。  
```sql
alter view 视图名
as 
查询语句;
```

**删除视图**  
可以同时删除多个视图，多个视图名称之间用逗号隔开。  
```sql
drop view 视图名1 [,视图名2] [,视图名n];
```

**查询视图结构**  
也有两种方式。  
```sql
/*方式1，显示视图查询的数据*/
desc 视图名称;
/*方式2，显示了视图的创建语句*/
show create view 视图名称;
```

**更新视图**  
视图的更新是更改视图中的数据，而不是更改视图中的 sql 逻辑。  
当对视图进行更新后，也会对原始表的数据进行更新。  
为了防止对原始表的数据产生更新，可以为视图添加只读权限，只允许读视图，不允许对视图进行更新。一般情况下，极少对视图进行更新操作。  
```sql
CREATE OR REPLACE VIEW myv4
  AS
  SELECT last_name,email
  from employees;

/*插入*/
insert into myv4 VALUES ('路人甲Java','javacode2018@163.com');
SELECT * from myv4 where email like 'javacode2018%';

/*修改*/
UPDATE myv4 SET last_name = '刘德华' WHERE last_name = '路人甲Java';
SELECT * from myv4 where email like 'javacode2018%';

/*删除*/
DELETE FROM myv4 where last_name = '刘德华';
SELECT * from myv4 where email like 'javacode2018%';
```

**视图的使用步骤**  
1、创建视图  
2、对视图执行查询操作  

```sql
/*案例1：查询姓名中包含a字符的员工名、部门、工种信息*/
/*①创建视图 myv1*/
CREATE VIEW myv1
AS
  SELECT
    t1.last_name,
    t2.department_name,
    t3.job_title
  FROM employees t1, departments t2, jobs t3
  WHERE t1.department_id = t2.department_id
        AND t1.job_id = t3.job_id;

/*②使用视图*/
SELECT * FROM myv1 a where a.last_name like 'a%';

/*案例2：查询各部门的平均工资级别*/
/*①创建视图 myv2*/
CREATE VIEW myv2
AS
  SELECT
    t1.department_id 部门id,
    t1.ag            平均工资,
    t2.grade_level   工资级别
  FROM (SELECT
          department_id,
          AVG(salary) ag
        FROM employees
        GROUP BY department_id)
       t1, job_grades t2
  WHERE t1.ag BETWEEN t2.lowest_sal AND t2.highest_sal;

/*②使用视图*/
SELECT * FROM myv2;
```


