
### 关系型数据库 Mysql

### Mysql 事务

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

# 删除数据表
DROP TABLE table_name;

# 写入数据
# INSERT INTO table_name ( field1, field2,...fieldN ) VALUES (value1, value2,...valueN );
INSERT INTO pet VALUES ('Puffball','Diane','hamster','f','1999-03-30',NULL);

# 更新数据
# UPDATE table_name SET field1=new-value1, field2=new-value2 [WHERE Clause]
UPDATE pet SET birth = '1989-08-31' WHERE name = 'Bowser';

# 删除数据
# DELETE FROM table_name [WHERE Clause]
DELETE pet WHERE id=100;

# 基本条件查询
SELECT name, email FROM pet WHERE birth >= '1998-1-1';

# 去重查询
SELECT DISTINCT owner FROM pet;

# is null 查询
SELECT name, birth FROM pet WHERE death IS NOT NULL;

# 模糊查询(_ 匹配一个，% 匹配零个或多个)
SELECT * FROM pet WHERE name LIKE '_b%';

# 正则查询(^ 开头定位符，$ 结尾定位符，. 任意字符，{number} 出现次数...)
SELECT * FROM pet WHERE name REGEXP '^b.{5}$';

# 统计
SELECT COUNT(*) FROM pet;

# 排序(默认升序)
SELECT name, birth FROM pet ORDER BY birth [DESC];

# 分组
SELECT sex, COUNT(*) FROM pet GROUP BY sex;

# 连表查询
SELECT pet.name,TIMESTAMPDIFF(YEAR,birth,date) AS age,remark FROM pet INNER JOIN event ON pet.name = event.name WHERE event.type = 'litter';

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

```

