
### 视图
视图是一种虚拟的表，具有和物理表相同的功能，可以对视图进行增、改、查操作。视图通常是一个表或者多个表的行或列的子集。  
视图创建脚本如下：  
```mysql
create view vname as
select column_names
from table_name
where condition
```

- 视图的优点
> 获取数据更容易，相对于多表查询来说；  
> 视图能够对机密数据提供安全保护；  
> 视图的修改不会影响基本表，提供了独立的操作单元，比较轻量。  

