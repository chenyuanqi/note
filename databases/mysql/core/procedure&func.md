
### 存储过程
存储过程是一组预编译好的 sql 语句集合，可以理解成批处理语句。  

存储过程的优点如下：  

- 预先编译，而不需要每次运行时编译，提高了数据库执行效率；
- 封装了一系列操作，对于一些数据交互比较多的操作，相比于单独执行 SQL 语句，可以减少网络通信量；
- 具有可复用性，减少了数据库开发的工作量；
- 安全性高，可以让没有权限的用户通过存储过程间接操作数据库；
- 更易于维护

存储过程的缺点是：  

- 可移植性差，存储过程将应用程序绑定到了数据库上；
- 开发调试复杂：没有好的 IDE；
- 修改复杂，需要重新编译，有时还需要更新程序中的代码以更新调用

**创建存储过程**  
格式如下：  
```sql
create procedure 存储过程名([参数模式] 参数名 参数类型)
begin
    存储过程体
end
```
参数模式有 3 种：  
in：该参数可以作为输入，也就是该参数需要调用方传入值。  
out：该参数可以作为输出，也就是说该参数可以作为返回值。  
inout：该参数既可以作为输入也可以作为输出，也就是说该参数需要在调用的时候传入值，又可以作为返回值。  
参数模式默认为 IN。  

一个存储过程可以有多个输入、多个输出、多个输入输出参数。  

**调用存储过程**  
注意：调用存储过程关键字是 call。
```sql
call 存储过程名称(参数列表);
```

**删除存储过程**  
存储过程只能一个个删除，不能批量删除。
```sql
# if exists：表示存储过程存在的情况下删除
drop procedure [if exists] 存储过程名称;
```

**修改存储过程**  
存储过程不能修改，若涉及到修改的，可以先删除，然后重建。  

**查看存储过程**  
可以查看存储过程详细创建语句。  
```sql
show create procedure 存储过程名称;
```

**空参列表的存储过程示例**  
mysql 默认结束符是分号。  
delimiter 用来设置结束符，当 mysql 执行脚本的时候，遇到结束符的时候，会把结束符前面的所有语句作为一个整体运行。  
```sql
/*设置结束符为$*/
DELIMITER $
/*如果存储过程存在则删除*/
DROP PROCEDURE IF EXISTS proc1;
/*创建存储过程proc1*/
CREATE PROCEDURE proc1()
  BEGIN
    INSERT INTO t_user VALUES (1,30,'路人甲Java');
    INSERT INTO t_user VALUES (2,50,'刘德华');
  END $

/*将结束符置为;*/
DELIMITER ;

# 调用存储过程
CALL proc1();

# 查看效果
select * from t_user;
+----+-----+---------------+
| id | age | name          |
+----+-----+---------------+
|  1 |  30 | 路人甲Java    |
|  2 |  50 | 刘德华        |
+----+-----+---------------+
```

**带 in 参数的存储过程示例**  
```sql
/*设置结束符为$*/
DELIMITER $
/*如果存储过程存在则删除*/
DROP PROCEDURE IF EXISTS proc2;
/*创建存储过程proc2，前两个参数没标明默认 in*/
CREATE PROCEDURE proc2(id int,age int,in name varchar(16))
  BEGIN
    INSERT INTO t_user VALUES (id,age,name);
  END $

/*将结束符置为;*/
DELIMITER ;

# 调用存储过程
/*创建了3个自定义变量*/
SELECT @id:=3,@age:=56,@name:='张学友';
/*调用存储过程*/
CALL proc2(@id,@age,@name);

# 查看效果
select * from t_user;
+----+-----+---------------+
| id | age | name          |
+----+-----+---------------+
|  1 |  30 | 路人甲Java    |
|  2 |  50 | 刘德华        |
|  3 |  56 | 张学友        |
+----+-----+---------------+
```

**带 out 参数的存储过程示例**
```sql
delete a from t_user a where a.id = 4;
/*如果存储过程存在则删除*/
DROP PROCEDURE IF EXISTS proc3;
/*设置结束符为$*/
DELIMITER $
/*创建存储过程proc3，前两个参数没标明默认 in*/
CREATE PROCEDURE proc3(id int,age int,in name varchar(16),out user_count int,out max_id INT)
  BEGIN
    INSERT INTO t_user VALUES (id,age,name);
    /*查询出t_user表的记录，放入user_count中,max_id用来存储t_user中最小的id*/
    SELECT COUNT(*),max(id) into user_count,max_id from t_user;
  END $

/*将结束符置为;*/
DELIMITER ;

# 调用存储过程
/*创建了3个自定义变量*/
SELECT @id:=4,@age:=55,@name:='郭富城';
/*调用存储过程*/
CALL proc3(@id,@age,@name,@user_count,@max_id);

# 查看效果
select @user_count,@max_id;
+-------------+---------+
| @user_count | @max_id |
+-------------+---------+
|           4 |       4 |
+-------------+---------+
```

**带 inout 参数的存储过程示例**
```sql
/*如果存储过程存在则删除*/
DROP PROCEDURE IF EXISTS proc4;
/*设置结束符为$*/
DELIMITER $
/*创建存储过程proc4*/
CREATE PROCEDURE proc4(INOUT a int,INOUT b int)
  BEGIN
    SET a = a*2;
    select b*2 into b;
  END $

/*将结束符置为;*/
DELIMITER ;

# 调用存储过程
/*创建了2个自定义变量*/
set @a=10,@b:=20;
/*调用存储过程*/
CALL proc4(@a,@b);

# 查看效果
SELECT @a,@b;
+------+------+
| @a   | @b   |
+------+------+
|   20 |   40 |
+------+------+
```

### 自定义函数
mysql 5.0 开始支持自定义函数，自定义函数是一组预编译好的 sql 语句集合，可以理解成批处理语句。类似于 java 中的方法，但是必须有返回值。  

**创建自定义函数**  
注意：参数是可选的，返回值是必须的。  
```sql
create function 函数名(参数名称 参数类型)
returns 返回值类型
begin
    函数体
end
```

**调用自定义函数**  
```sql
select 函数名(实参列表);
```

**删除自定义函数**   
```sql
drop function [if exists] 函数名;
```

**查看自定义函数**  
```sql
show create function 函数名;
```

**无参自定义函数示例**  
```sql
/*删除fun1*/
DROP FUNCTION IF EXISTS fun1;
/*设置结束符为$*/
DELIMITER $
/*创建函数*/
CREATE FUNCTION fun1()
  returns INT
  BEGIN
    DECLARE max_id int DEFAULT 0;
    SELECT max(id) INTO max_id FROM t_user;
    return max_id;
  END $
/*设置结束符为;*/
DELIMITER ;

# 调用并查看结果
SELECT fun1();
+--------+
| fun1() |
+--------+
|      4 |
+--------+
```

**有参自定义函数示例**  
```sql
/*删除函数*/
DROP FUNCTION IF EXISTS get_user_id;
/*设置结束符为$*/
DELIMITER $
/*创建函数*/
CREATE FUNCTION get_user_id(v_name VARCHAR(16))
  returns INT
  BEGIN
    DECLARE r_id int;
    # SELECT INTO 语句从一个表复制数据，然后把数据插入到另一个新表中
    SELECT id INTO r_id FROM t_user WHERE name = v_name;
    return r_id;
  END $
/*设置结束符为;*/
DELIMITER ;

# 调用并查看结果
SELECT get_user_id(name) from t_user;
+-------------------+
| get_user_id(name) |
+-------------------+
|                 1 |
|                 2 |
|                 3 |
|                 4 |
+-------------------+
```

### 存储过程 Vs 自定义函数
存储过程的关键字为 procedure，返回值可以有多个，调用时用 call， 一般用于执行比较复杂的的过程体、更新、创建等语句。  
自定义函数的关键字为 function， 返回值必须有一个，调用时用 select，一般用于查询单个值并返回。  
存储过程一般是作为一个独立的部分来执行，而自定义函数可以作为查询语句的一个部分来调用。  


| | 存储过程	| 函数 |  
| :---: | :---: |  :---: |  
| 返回值 | 可以有 0 个或者多个 | 必须有一个 |  
| 关键字| procedure | function |  
| 调用方式 | call | select |  