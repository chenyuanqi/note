
### Hive 权限控制
Hive 的权限需要在 hive-site.xml 文件中设置才会起作用，配置默认的是 false。  
需要把 hive.security.authorization.enabled 设置为 true，并对不同的用户设置不同的权限，例如 select ,drop 等的操作。  

### Hive 优化数据倾斜
1、万能方法
```
hive.groupby.skewindata=true
```

2、大小表关联
small_table join big_table

3、数据中有大量0或NULL
```sql
on case when (x.uid = '-' or x.uid = '0' or x.uid is null)  
then concat('dp_hive_search',rand()) else x.uid  
end = f.user_id; 
```

4、大大表关联
```sql
Select *  
from dw_log t11  
join (  
    select t1.*  
    from (  
        select user_id from dw_log group by user_id  
        ) t  
        join dw_user t1  
        on t.user_id=t1.user_id  
) t12  
on t11.user_id=t12.user_id 
```

5、count distinct 时存在大量特殊值
```sql
select cast(count(distinct user_id)+1 as bigint) as user_cnt  
from tab_a  
where user_id is not null and user_id <> ''
```

6、空间换时间
```sql
select day,  
count(case when type='session' then 1 else null end) as session_cnt,  
count(case when type='user' then 1 else null end) as user_cnt  
from (  
    select day,session_id,type  
        from (  
            select day,session_id,'session' as type  
            from log  
            union all  
            select day user_id,'user' as type  
            from log  
        )  
        group by day,session_id,type  
    ) t1  
group by day  
```
