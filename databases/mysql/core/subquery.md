
### 子查询
出现在 select 语句中的 select 语句，称为`子查询`或`内查询`；外部的 select 查询语句，称为`主查询`或`外查询`。  

子查询按结果集的行列数不同分为 4 种：  

- 标量子查询（结果集只有一行一列）  
- 列子查询（结果集只有一列多行）  
- 行子查询（结果集有一行多列）  
- 表子查询（结果集一般为多行多列）  

按子查询出现在主查询中的不同位置分：  

- select 后面：仅仅支持标量子查询  
- from 后面：支持表子查询  
- where 或 having 后面：支持标量子查询（单列单行）、列子查询（单列多行）、行子查询（多列多行）  
- exists 后面（即相关子查询）：表子查询（多行、多列）  

**select 后面的子查询**  
子查询位于 select 后面的，仅仅支持标量子查询。  
```sql
# 查询每个部门员工个数
SELECT
  a.*,
  (SELECT count(*)
   FROM employees b
   WHERE b.department_id = a.department_id) AS 员工个数
FROM departments a;

# 查询员工号 = 102 的部门名称
SELECT (SELECT a.department_name
        FROM departments a, employees b
        WHERE a.department_id = b.department_id
              AND b.employee_id = 102) AS 部门名;
```

**from 后面的子查询**  
将子查询的结果集充当一张表，要求必须起别名，否则这个表找不到。然后，将真实的表和子查询结果表进行连接查询。  
```sql
# 查询每个部门平均工资
SELECT
  department_id,
  avg(a.salary)
FROM employees a
GROUP BY a.department_id;

# 薪资等级表
SELECT *
FROM job_grades;

# 将上面 2 个结果连接查询，筛选条件:平均工资 between lowest_sal and highest_sal;
SELECT
  t1.department_id,
  sa AS '平均工资',
  t2.grade_level
FROM (SELECT
        department_id,
        avg(a.salary) sa
      FROM employees a
      GROUP BY a.department_id) t1, job_grades t2
WHERE
  t1.sa BETWEEN t2.lowest_sal AND t2.highest_sal;
```

**where 和 having 后面的子查询**  
where 或 having 后面，可以使用：标量子查询（单行单列行子查询）、列子查询（单列多行子查询）、行子查询（多行多列）。

特点：  
1、子查询放在小括号内  
2、子查询一般放在条件的右侧  
3、标量子查询，一般搭配着单行操作符使用，多行操作符   >、<、>=、<=、=、<>、!=  
4、列子查询，一般搭配着多行操作符使用  
> in (not in)：列表中的 “任意一个”  
> any 或者 some：和子查询返回的 “某一个值” 比较，比如 a>som (10,20,30)，a 大于子查询中任意一个即可，a 大于子查询中最小值即可，等同于 a>min (10,20,30)。  
> all：和子查询返回的 “所有值” 比较，比如 a>all (10,20,30)，a 大于子查询中所有值，换句话说，a 大于子查询中最大值即可满足查询条件，等同于 a>max (10,20,30);

5、子查询的执行优先于主查询执行，因为主查询的条件用到了子查询的结果  

mysql 中的 in、any、some、all：  
in，any，some，all 分别是子查询关键词之一。  
in：in 常用于 where 表达式中，其作用是查询某个范围内的数据。  
any 和 some 一样： 可以与 =、>、>=、<、<=、<> 结合起来使用，分别表示等于、大于、大于等于、小于、小于等于、不等于其中的任何一个数据。  
all：可以与 =、>、>=、<、<=、<> 结合是来使用，分别表示等于、大于、大于等于、小于、小于等于、不等于其中的其中的所有数据。  

标量子查询 sql 示例：  
```sql
/*①查询 abel 的工资【改查询是标量子查询】*/
SELECT salary
FROM employees
WHERE last_name = 'Abel';

/*②查询员工信息，满足 salary>① 的结果*/
SELECT *
FROM employees a
WHERE a.salary > (SELECT salary
                  FROM employees
                  WHERE last_name = 'Abel');
```

子查询 + 分组函数 sql 示例：  
```sql
/* 查询最低工资大于50号部门最低工资的部门id和其最低工资【having】 */
/* ①查询 50 号部门的最低工资 */
SELECT min(salary)
FROM employees
WHERE department_id = 50;
/* ②查询每个部门的最低工资 */
SELECT
  min(salary),
  department_id
FROM employees
GROUP BY department_id;
/* ③在②的基础上筛选，满足 min(salary)>① */
SELECT
  min(a.salary) minsalary,
  department_id
FROM employees a
GROUP BY a.department_id
HAVING min(a.salary) > (SELECT min(salary)
                        FROM employees
                        WHERE department_id = 50);
```

列子查询需要搭配多行操作符使用：in (not in)、any/some、all。为了提升效率，最好去重一下 distinct 关键字。  
列子查询 sql 示例：  
```sql
/* 返回 location_id 是 1400 或 1700 的部门中的所有员工姓名 */
/* 方式1 */
/* ①查询 location_id 是 1400 或 1700 的部门编号 */
SELECT DISTINCT department_id
FROM departments
WHERE location_id IN (1400, 1700);

/* ②查询员工姓名，要求部门是①列表中的某一个 */
SELECT a.last_name
FROM employees a
WHERE a.department_id IN (SELECT DISTINCT department_id
                          FROM departments
                          WHERE location_id IN (1400, 1700));

/* 方式2：使用 any 实现 */
SELECT a.last_name
FROM employees a
WHERE a.department_id = ANY (SELECT DISTINCT department_id
                             FROM departments
                             WHERE location_id IN (1400, 1700));

/* 拓展，下面与 not in 等价 */
SELECT a.last_name
FROM employees a
WHERE a.department_id <> ALL (SELECT DISTINCT department_id
                             FROM departments
                             WHERE location_id IN (1400, 1700));
```

行子查询（结果集一行多列）sql 示例：  
```sql
/*查询员工编号最小并且工资最高的员工信息*/
/* ①查询最小的员工编号 */
SELECT min(employee_id)
FROM employees;
/* ②查询最高工资 */
SELECT max(salary)
FROM employees;
/* ③方式1：查询员工信息 */
SELECT *
FROM employees a
WHERE a.employee_id = (SELECT min(employee_id)
                       FROM employees)
      AND salary = (SELECT max(salary)
                    FROM employees);

/* 方式 2 */
SELECT *
FROM employees a
WHERE (a.employee_id, a.salary) = (SELECT
                                     min(employee_id),
                                     max(salary)
                                   FROM employees);
/* 方式 3 */
SELECT *
FROM employees a
WHERE (a.employee_id, a.salary) in (SELECT
                                     min(employee_id),
                                     max(salary)
                                   FROM employees);
```

exists 后面（也叫做相关子查询）：  
1、语法：exists (玩转的查询语句)。  
2、exists 查询结果：1 或 0，exists 查询的结果用来判断子查询的结果集中是否有值。  
3、一般来说，能用 exists 的子查询，绝对都能用 in 代替，所以 exists 用的少。  
4、和前面的查询不同，这先执行主查询，然后主查询查询的结果，在根据子查询进行过滤，子查询中涉及到主查询中用到的字段，所以叫相关子查询。  
```sql
/* exists入门案例 */
SELECT exists(SELECT employee_id
              FROM employees
              WHERE salary = 300000) AS 'exists返回1或者0';

/* 查询所有员工部门名 */
SELECT department_name
FROM departments a
WHERE exists(SELECT 1
             FROM employees b
             WHERE a.department_id = b.department_id);

/* 使用 in 实现 */
SELECT department_name
FROM departments a
WHERE a.department_id IN (SELECT department_id
                          FROM employees);
```

NULL 的大坑：in 的情况下，子查询中列的值为 NULL 的时候，外查询的结果为空。  
建议：建表时，列不允许为空。  
```sql
SELECT *
FROM departments a
WHERE a.department_id NOT IN (SELECT department_id
                              FROM employees b);
```
