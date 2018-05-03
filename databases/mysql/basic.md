
### 关系型数据库 Mysql

### Mysql 事务

### MyISAM 和 InnoDB 的区别
> 1、MyISAM 查询效率更高，但是不支持事物  
> 2、InnoDB 插入、更新较高，支持事物处理  
> 3、MyISAM 支持表锁， InnoDb 支持行锁  
> 4、MyISAM 是默认引擎，InnoDB 需要指定  
> 5、InnoDB 不支持 FULLTEXT 类型的索引  

### Mysql 基本增删改查
```bash

# 包含查询
# 6 在  '[1,2,3,6]' 内，示例
SELECT * FROM `articles_onlines` WHERE FIND_IN_SET('6', TRIM(TRAILING ']' FROM TRIM(LEADING '[' FROM `article_tags`)));
# 6 在 '1,2,3,6' 内，示例
SELECT * FROM `articles_onlines` WHERE FIND_IN_SET(6, `article_tags`);
# 某字段是否存在字符串中，譬如 tags_id in ('1,2,3')
SELECT * FROM `article_tabs` WHERE `tags_id` IN '1,2,3';

```

