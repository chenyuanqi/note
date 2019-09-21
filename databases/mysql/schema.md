
### 什么是 schema
在数据库中，schema ['skimə] 是数据库的组织和结构，称为模式。  

数据库模式 schema 是对一个数据库的结构描述。它包含 schema 对象，可以是表 (table)、列 (column)、数据类型 (data type)、视图 (view)、存储过程 (stored procedures)、关系 (relationships)、主键 (primary key)、外键 (foreign key) 等。  
数据库模式 schema 可以用一个可视化的图来表示，它显示了数据库对象及其相互之间的关系。  

如果把 database 看作是一个仓库，仓库很多房间（schema），一个 schema 代表一个房间，table 可以看作是每个房间中的储物柜，user 是每个 schema 的主人，有操作数据库中每个房间的权利，就是说每个数据库映射的 user 有每个 schema（房间）的钥匙。  
默认情况下，一个用户对应一个集合，用户的 schema 名等于用户名，并作为该用户缺省 schema。所以 schema 集合看上去像用户名。访问一个表时，如果没有指明该表属于哪个 schema，系统会自动加上缺省的 schema。一个对象的完整名称为 schema.object 而不属于 user.object。  

模式 schema 与数据库 database 是否等同，需要取决于数据库的供应商。  
Mysql 在物理上，模式 schema 与数据库 database 是同义的，所以模式和数据库其实是一回事。  
> 在 MySQL 中，CREATE SCHEMA 创建了一个数据库，这是因为 CREATE SCHEMA 是 CREATE DATABASE 的同义词。  

