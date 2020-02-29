
### 游标概述
游标（Cursor）是处理数据的一种方法，为了查看或者处理结果集中的数据，游标提供了在结果集中一次一行遍历数据的能力。  
游标只能在存储过程和函数中使用。  

### 游标的基本使用
游标相当于一个指针，这个指针指向 select 语句的第一行数据，可以通过移动指针来遍历后面的数据。  

**游标的使用步骤**  
声明游标：这个过程只是创建了一个游标，需要指定这个游标需要遍历的 select 查询，声明游标时并不会去执行这个 sql。  
```sql
# 注意：一个 begin end 中只能声明一个游标
DECLARE 游标名称 CURSOR FOR 查询语句;
```
打开游标：打开游标的时候，会执行游标对应的 select 语句。  
```sql
open 游标名称;
```
遍历数据：使用游标循环遍历 select 结果中每一行数据，然后进行处理。   
```sql
# 取出当前行的结果，将结果放在对应的变量中，并将游标指针指向下一行的数据
# 当调用 fetch 的时候，会获取当前行的数据，如果当前行无数据，会引发 mysql 内部的 NOT FOUND 错误
fetch 游标名称 into 变量列表;
```
关闭游标：游标使用完之后一定要关闭。  
```sql
close 游标名称;
```

**游标使用示例**  
mysql 的设置默认是不允许创建函数，临时开启则在命令行执行如下命令：
```sql
# 开启需要在 my.cnf 里面设置：log-bin-trust-function-creators=1; 并重启 mysql 服务
SET GLOBAL log_bin_trust_function_creators = 1;
```
`要注意的是，有主从复制的时候，从机必须要设置，不然会导致主从同步失败。`  

写一个函数，计算 test 表中 a、b 字段所有的和。  
```sql
/*删除函数*/
DROP FUNCTION IF EXISTS func_sum;
/*声明结束符为$*/
DELIMITER $
/*创建函数*/
CREATE FUNCTION func_sum(v_max_a int)
  RETURNS int
  BEGIN
    /*用于保存结果*/
    DECLARE v_total int DEFAULT 0;
    /*创建一个变量，用来保存当前行中 a 的值*/
    DECLARE v_a int DEFAULT 0;
    /*创建一个变量，用来保存当前行中 b 的值*/
    DECLARE v_b int DEFAULT 0;
    /*创建游标结束标志变量*/
    DECLARE v_done int DEFAULT FALSE;
    /*创建游标*/
    DECLARE cur_test CURSOR FOR SELECT a,b from test where a<=v_max_a;
    /*设置游标结束时v_done的值为true，可以v_done来判断游标是否结束了*/
    DECLARE CONTINUE HANDLER FOR NOT FOUND SET v_done=TRUE;
    /*设置v_total初始值*/
    SET v_total = 0;
    /*打开游标*/
    OPEN cur_test;
    /*使用Loop循环遍历游标*/
    a:LOOP
      /*先获取当前行的数据，然后将当前行的数据放入v_a,v_b中，如果当前行无数据，v_done会被置为true*/
      FETCH cur_test INTO v_a, v_b;
      /*通过v_done来判断游标是否结束了，退出循环*/
      if v_done THEN
        LEAVE a;
      END IF;
      /*对v_total值累加处理*/
      SET v_total = v_total + v_a + v_b;
    END LOOP;
    /*关闭游标*/
    CLOSE cur_test;
    /*返回结果*/
    RETURN v_total;
  END $
/*结束符置为;*/
DELIMITER ;

# 效果查询
SELECT a,b FROM test;
SELECT func_sum(1);
SELECT func_sum(2);
```

游标的详细执行过程：  
游标中有个指针，当打开游标的时候，才会执行游标对应的 select 语句，这个指针会指向 select 结果中第一行记录。  
当调用 fetch 游标名称时，会获取当前行的数据，如果当前行无数据，会触发 NOT FOUND 异常。  
当触发 NOT FOUND 异常的时候，我们可以使用一个变量来标记一下，如下代码：  
```sql
DECLARE CONTINUE HANDLER FOR NOT FOUND SET v_done=TRUE;
```
当游标无数据触发 NOT FOUND 异常的时候，将变量 v_down 的值置为 TURE，循环中就可以通过 v_down 的值控制循环的退出。  
如果当前行有数据，则将当前行数据存到对应的变量中，并将游标指针指向下一行数据，如下语句：  
```sql
fetch 游标名称 into 变量列表;
```

### 嵌套游标
写个存储过程，遍历 test2、test3，将 test2 中的 a 字段和 test3 中的 b 字段任意组合，插入到 test1 表中。  
```sql
/*删除存储过程*/
DROP PROCEDURE IF EXISTS proc1;
/*声明结束符为$*/
DELIMITER $
/*创建存储过程*/
CREATE PROCEDURE proc1()
  BEGIN
    /*创建一个变量，用来保存当前行中a的值*/
    DECLARE v_a int DEFAULT 0;
    /*创建游标结束标志变量*/
    DECLARE v_done1 int DEFAULT FALSE;
    /*创建游标*/
    DECLARE cur_test1 CURSOR FOR SELECT a FROM test2;
    /*设置游标结束时v_done1的值为true，可以v_done1来判断游标cur_test1是否结束了*/
    DECLARE CONTINUE HANDLER FOR NOT FOUND SET v_done1=TRUE;
    /*打开游标*/
    OPEN cur_test1;
    /*使用Loop循环遍历游标*/
    a:LOOP
      FETCH cur_test1 INTO v_a;
      /*通过v_done1来判断游标是否结束了，退出循环*/
      if v_done1 THEN
        LEAVE a;
      END IF;

      BEGIN
        /*创建一个变量，用来保存当前行中b的值*/
        DECLARE v_b int DEFAULT 0;
        /*创建游标结束标志变量*/
        DECLARE v_done2 int DEFAULT FALSE;
        /*创建游标*/
        DECLARE cur_test2 CURSOR FOR SELECT b FROM test3;
        /*设置游标结束时v_done1的值为true，可以v_done1来判断游标cur_test2是否结束了*/
        DECLARE CONTINUE HANDLER FOR NOT FOUND SET v_done2=TRUE;

        /*打开游标*/
        OPEN cur_test2;
        /*使用Loop循环遍历游标*/
        b:LOOP
          FETCH cur_test2 INTO v_b;
          /*通过v_done1来判断游标是否结束了，退出循环*/
          if v_done2 THEN
            LEAVE b;
          END IF;

          /*将v_a、v_b插入test1表中*/
          INSERT INTO test1 VALUES (v_a,v_b);
        END LOOP b;
        /*关闭cur_test2游标*/
        CLOSE cur_test2;
      END;

    END LOOP;
    /*关闭游标cur_test1*/
    CLOSE cur_test1;
  END $
/*结束符置为;*/
DELIMITER ;

# 效果查询
DELETE FROM test1
CALL proc1();
SELECT * from test1;
```
