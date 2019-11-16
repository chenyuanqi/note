### NoSQL 简介            

NoSQL(NoSQL = Not Only SQL )，意即"不仅仅是 SQL“。NoSQL，指的是非关系型的数据库。NoSQL 有时也称作 Not Only SQL的缩写，是对不同于传统的关系型数据库的数据库管理系统的统称。

### MongoDB简介

MongoDB 是由 C++ 语言编写的，是一个基于分布式文件存储的开源数据库系统。在高负载的情况下，添加更多的节点，可以保证服务器性能。

MongoDB 旨在为 WEB 应用提供可扩展的高性能数据存储解决方案。

MongoDB 将数据存储为一个文档，数据结构由键值(key=>value)对组成。

MongoDB 采用 BSON 存储文档数据 (有大小限制，最大 16M)，文档类似于 JSON 对象。字段值可以包含其他文档，数组及文档数组。

**主要特点**

1. MongoDB的提供了一个面向文档存储，操作起来比较简单和容易。
2. 可以在MongoDB记录中设置任何属性的索引 (如：FirstName="Sameer",Address="8.Gandhi Road")来实现更快的排序。
3. 可以通过本地或者网络创建数据镜像，这使得MongoDB有更强的扩展性。
4. 如果负载的增加（需要更多的存储空间和更强的处理能力） ，它可以分布在计算机网络   中的其他节点上这就是所谓的分片。
5. Mongo支持丰富的查询表达式。查询指令使用JSON形式的标记，可轻易查询文档中内嵌的对象及数组。
6. MongoDb 使用update()命令可以实现替换完成的文档（数据）或者一些指定的数据字段 。Mongodb中的Map/reduce主要是用来对数据进行批量处理和聚合操作。
7. Map和Reduce。Map函数调用emit(key,value)遍历集合中所有的记录，将key与value传给Reduce函数进行处理。
8. Map函数和Reduce函数是使用Javascript编写的，并可以通过db.runCommand或mapreduce命令来执行MapReduce操作。
9. GridFS是MongoDB中的一个内置功能，可以用于存放大量小文件。
10. MongoDB允许在服务端执行脚本，可以用Javascript编写某个函数，直接在服务端执行，也可以把函数的定义存储在服务端，下次直接调用即可。
11. MongoDB支持各种编程语言:RUBY，PYTHON，JAVA，C++，PHP，C#等多种语言。1
12. MongoDB安装简单。

### MongoDB概念介绍

| SQL术语/概念 | MongoDB术语/概念 | 解释/说明                           |
| ------------ | ---------------- | ----------------------------------- |
| database     | database         | 数据库                              |
| table        | collection       | 表/集合                             |
| row          | document         | 行/文档                             |
| column       | filed            | 数据字段/域                         |
| index        | index            | 索引                                |
| primary key  | primary key      | 主键,MongoDB自动将_id字段设置为主键 |

一个集合中的文档实例

```json
{
    "username":"Tim",
    "age":18
}
```

- 数据库
数据库可以看成是一个电子化的文件柜，用户可以对文件中的数据运行新增、检索、更新、删除等操作。数据库是一个所有集合的容器，在文件系统中每一个数据库都有一个相关的物理文件。

- 集合
集合就是一组 MongoDB 文档。它相当于关系型数据库（RDBMS）中的表这种概念。集合位于单独的一个数据库中。一个集合内的多个文档可以有多个不同的字段。一般来说，集合中的文档都有着相同或相关的目的。

- 文档
文档由一组 key value 组成。文档是动态模式，这意味着同一集合里的文档不需要有相同的字段和结构。在关系型数据库中 table 中的每一条记录相当于 MongoDB 中的一个文档。

- mongod
mongod 是处理 MongoDB 系统的主要进程。它处理数据请求，管理数据存储，和执行后台管理操作。当我们运行 mongod 命令意味着正在启动 MongoDB 进程，并且在后台运行。  
mongod 参数中传递数据库存储路径，默认是 "/data/db"，端口号默认是 "27017"。  

- 命名空间
MongoDB 内部有预分配空间的机制，每个预分配的文件都用 0 进行填充。数据文件每新分配一次，它的大小都是上一个数据文件大小的 2 倍，每个数据文件最大 2G。  
MongoDB 每个集合和每个索引都对应一个命名空间，这些命名空间的元数据集中在 16M 的 \*.ns 文件中，平均每个命名占用约 628 字节，也即整个数据库的命名空间的上限约为 24000。  
如果每个集合有一个索引（比如默认的_id 索引），那么最多可以创建 12000 个集合。如果索引数更多，则可创建的集合数就更少了。同时，如果集合数太多，一些操作也会变慢。要建立更多的集合的话，MongoDB 也是支持的，只需要在启动时加上 “--nssize” 参数，这样对应数据库的命名空间文件就可以变得更大以便保存更多的命名。这个命名空间文件（.ns 文件）最大可以为 2G。  
每个命名空间对应的盘区不一定是连续的。与数据文件增长相同，每个命名空间对应的盘区大小都是随分配次数不断增长的。目的是为了平衡命名空间浪费的空间与保持一个命名空间数据的连续性。  
需要注意的一个命名空间 $freelist，这个命名空间用于记录不再使用的盘区（被删除的 Collection 或索引）。每当命名空间需要分配新盘区时，会先查看 $freelist 是否有大小合适的盘区可以使用，如果有就回收空闲的磁盘空间。  

- 分片
分片是将数据水平切分到不同的物理节点。当应用数据越来越大的时候，数据量也会越来越大。当数据量增长时，单台机器有可能无法存储数据或可接受的读取写入吞吐量。利用分片技术可以添加更多的机器来应对数据量增加以及读写操作的要求。



### MongoDB 使用场景
MongoDB 格式灵活，如果应用不需要事务及复杂 join 支持，则可以考虑使用它。主要应用场景如下：  
游戏场景，使用 MongoDB 存储游戏用户信息，用户的装备、积分等直接以内嵌文档的形式存储，方便查询、更新  
物流场景，使用 MongoDB 存储订单信息，订单状态在运送过程中会不断更新，以 MongoDB 内嵌数组的形式来存储，一次查询就能将订单所有的变更读取出来  
社交场景，使用 MongoDB 存储存储用户信息，以及用户发表的朋友圈信息，通过地理位置索引实现附近的人、地点等功能  
物联网场景，使用 MongoDB 存储所有接入的智能设备信息，以及设备汇报的日志信息，并对这些信息进行多维度的分析  
视频直播，使用 MongoDB 存储用户信息、礼物信息等  


### MongoDB 相关操作

```shell
show dbs #查看数据库列表
use test #创建并切换到test库
show collections #查看数据库的所有集合
# 集合相关操作
db.testCollection.insert({"name":"hello"}) #创建testCollection库并插入数据
db.testCollection.find() #查询testCollection所有数据
db.testCollection.drop()# 删除集合
db.testCollection.find().limit(5) #显示5条
#文档操作
db.users.find(
  {age: {$gt: 18}},                                  //查询条件
  {name: 1, address: 1}                        //查询显示字段
).limit(5)  

db.users.update(
   {age: {$gt: 18}},
   {$set: {status: "A"}},
   {multi: true}                        //multi指所有行修改
)

#统计文档的个数
db.testCollection.count()

# 查看 mongodb 连接
db.adminCommand("connPoolStats")

```

### MongoDB 数据类型

下表为MongoDB中常用的几种数据类型。

| 数据类型           | 描述                                                         |
| ------------------ | ------------------------------------------------------------ |
| String             | 字符串。存储数据常用的数据类型。在 MongoDB 中，UTF-8 编码的字符串才是合法的。 |
| Integer            | 整型数值。用于存储数值。根据你所采用的服务器，可分为 32 位或 64 位。 |
| Boolean            | 布尔值。用于存储布尔值（真/假）。                            |
| Double             | 双精度浮点值。用于存储浮点值。                               |
| Min/Max keys       | 将一个值与 BSON（二进制的 JSON）元素的最低值和最高值相对比。 |
| Arrays             | 用于将数组或列表或多个值存储为一个键。                       |
| Timestamp          | 时间戳。记录文档修改或添加的具体时间。                       |
| Object             | 用于内嵌文档。                                               |
| Null               | 用于创建空值。                                               |
| Symbol             | 符号。该数据类型基本上等同于字符串类型，但不同的是，它一般用于采用特殊符号类型的语言。 |
| Date               | 日期时间。用 UNIX 时间格式来存储当前日期或时间。你可以指定自己的日期时间：创建 Date 对象，传入年月日信息。 |
| Object ID          | 对象 ID。用于创建文档的 ID。                                 |
| Binary Data        | 二进制数据。用于存储二进制数据。                             |
| Code               | 代码类型。用于在文档中存储 JavaScript 代码。                 |
| Regular expression | 正则表达式类型。用于存储正则表达式。                         |

### 文档操作

#### 1. 插入

```shell
db.Collection.insertOne({"user":"zhang0"}) #插入单条
db.collection.insertMany([{},{}]) #插入多条
```

#### 2. 查询

```json
[
    {
        "item": "journal",
        "qty": 25,
        "size": {
            "h": 14,
            "w": 21,
            "uom": "cm"
        },
        "status": "A"
    },
    {
        "item": "notebook",
        "qty": 50,
        "size": {
            "h": 8.5,
            "w": 11,
            "uom": "in"
        },
        "status": "A"
    },
    {
        "item": "paper",
        "qty": 100,
        "size": {
            "h": 8.5,
            "w": 11,
            "uom": "in"
        },
        "status": "D"
    },
    {
        "item": "planner",
        "qty": 75,
        "size": {
            "h": 22.85,
            "w": 30,
            "uom": "cm"
        },
        "status": "D"
    },
    {
        "item": "postcard",
        "qty": 45,
        "size": {
            "h": 10,
            "w": 15.25,
            "uom": "cm"
        },
        "status": "A"
    }
]
```

```shell
db.inventory.find() #查找所有 =SELECT * FROM inventory
db.inventory.find( { "size.h": { $lt: 15 } } )
```

- 条件查询

```shell
db.inventory.find( { status: "D" } ) # SELECT * FROM inventory WHERE status = "D"
```

- in查询

```shell
db.inventory.find( { status: { $in: [ "A", "D" ] } } )
# SELECT * FROM inventory WHERE status in ("A", "D")
```

- AND

```shell
db.inventory.find( { status: "A", qty: { $lt: 30 } } )
#SELECT * FROM inventory WHERE status = "A" AND qty < 30
```

- OR

```shell
db.inventory.find( { $or: [ { status: "A" }, { qty: { $lt: 30 } } ] } )
# SELECT * FROM inventory WHERE status = "A" OR qty < 30
```

- 指定filed

等于1 标识查询。等于0 不显示。默认显示_id

```shell
db.inventory.find( { status: "A" }, { item: 1, status: 1 } )
# SELECT _id, item, status from inventory WHERE status = "A"
```

#### 3. 修改

- updateOne 更新单条记录

```shell
db.inventory.updateOne(
   { item: "paper" },
   {
     $set: { "size.uom": "cm", status: "P" },
     $currentDate: { lastModified: true }
   }
)

# UPDATE inventory SET item = paper WHERE `size.uom` = 'cm' AND status = 'p' LIMIT 1
```

- updateMany更新多条

```shell
db.inventory.updateMany(
   { "qty": { $lt: 50 } },
   {
     $set: { "size.uom": "in", status: "P" },
     $currentDate: { lastModified: true }
   }
)
```

- repalce

```shell
db.inventory.replaceOne(
   { item: "paper" },
   { item: "paper", instock: [ { warehouse: "A", qty: 60 }, { warehouse: "B", qty: 40 } ] }
)
```

#### 4. 删除

- deleteOne 删除单条

```shell
db.orders.deleteOne(
       { w : "majority", wtimeout : 100 }
   );
# DELETE FROM orders where w = "majority" AND  wtimeout = 100
```

- remove 删除整个文档

```shell
db.orders.remove() # DELETE FROM orders
```

#### 5. 创建索引

```shell
db.inventory.createIndex({"name":"username"})
```

`$text 	`查询运算符

```shell
db.stores.find( { $text: { $search: "java coffee shop" } } )
```

#### 5. 聚合操作

MongoDB中聚合(aggregate)主要用于处理数据(诸如统计平均值,求和等)，并返回计算后的数据结果。有点类似sql语句中的 count(*)。 

-  aggregate() 方法

```shell
db.mycol.aggregate([{$group : {_id : "$by_user", num_tutorial : {$sum : 1}}}])  
# select by_user, count(*) from mycol group by by_user
```

| $sum      | 计算总和。                                     | db.mycol.aggregate([{$group : {_id : "$by_user", num_tutorial : {$sum : "$likes"}}}]) |
| --------- | ---------------------------------------------- | ------------------------------------------------------------ |
| $avg      | 计算平均值                                     | db.mycol.aggregate([{$group : {_id : "$by_user", num_tutorial : {$avg : "$likes"}}}]) |
| $min      | 获取集合中所有文档对应值得最小值。             | db.mycol.aggregate([{$group : {_id : "$by_user", num_tutorial : {$min : "$likes"}}}]) |
| $max      | 获取集合中所有文档对应值得最大值。             | db.mycol.aggregate([{$group : {_id : "$by_user", num_tutorial : {$max : "$likes"}}}]) |
| $push     | 在结果文档中插入值到一个数组中。               | db.mycol.aggregate([{$group : {_id : "$by_user", url : {$push: "$url"}}}]) |
| $addToSet | 在结果文档中插入值到一个数组中，但不创建副本。 | db.mycol.aggregate([{$group : {_id : "$by_user", url : {$addToSet : "$url"}}}]) |
| $first    | 根据资源文档的排序获取第一个文档数据。         | db.mycol.aggregate([{$group : {_id : "$by_user", first_url : {$first : "$url"}}}]) |
| $last     | 根据资源文档的排序获取最后一个文档数据         | db.mycol.aggregate([{$group : {_id : "$by_user", last_url : {$last : "$url"}}}]) |

 聚合框架中常用的几个操作：

- $project：修改输入文档的结构。可以用来重命名、增加或删除域，也可以用于创建计算结果以及嵌套文档。
- $match：用于过滤数据，只输出符合条件的文档。$match使用MongoDB的标准查询操作。
- $limit：用来限制MongoDB聚合管道返回的文档数。
- $skip：在聚合管道中跳过指定数量的文档，并返回余下的文档。
- $unwind：将文档中的某一个数组类型字段拆分成多条，每条包含数组中的一个值。
- $group：将集合中的文档分组，可用于统计结果。
- $sort：将输入文档排序后输出。
- $geoNear：输出接近某一地理位置的有序文档

```shell
db.articles.aggregate( [   
	{ $match : { score : { $gt : 70, $lte : 90 } } },  
	{ $group: { _id: null, count: { $sum: 1 } } }                  
	] );  
#$match用于获取分数大于70小于或等于90记录，然后将符合条件的记录送到下一阶段$group管道操作符进行处理。
```

### SQL 到 Mongo 的对应表
这个列表是 PHP 版本的 [» SQL to Mongo](http://www.mongoing.com/docs/reference/sql-comparison.html) 对应表（在 MongoDB 官方手册中有更加通用的版本）。

| SQL 查询语句 | Mongo 查询语句 |
| --- | --- |
|CREATE TABLE USERS (a Number, b Number)       | 隐式的创建，或 [MongoDB::createCollection()](mongodb.createcollection.php). |
|INSERT INTO USERS VALUES(1,1)       |$db->users->insert(array("a" => 1, "b" => 1));       |
|SELECT a,b FROM users       |$db->users->find(array(), array("a" => 1, "b" => 1));       |
|SELECT * FROM users WHERE age=33       |$db->users->find(array("age" => 33));       |
|SELECT a,b FROM users WHERE age=33       |$db->users->find(array("age" => 33), array("a" => 1, "b" => 1));       |
|SELECT a,b FROM users WHERE age=33 ORDER BY name       |$db->users->find(array("age" => 33), array("a" => 1, "b" => 1))->sort(array("name" => 1));       |
|SELECT * FROM users WHERE age>33       |$db->users->find(array("age" => array('$gt' => 33)));       |
|SELECT * FROM users WHERE age<33       |$db->users->find(array("age" => array('$lt' => 33)));       |
|SELECT * FROM users WHERE name LIKE "%Joe%"       |$db->users->find(array("name" => new MongoRegex("/Joe/")));       |
|SELECT * FROM users WHERE name LIKE "Joe%"       |$db->users->find(array("name" => new MongoRegex("/^Joe/")));       |
|SELECT * FROM users WHERE age>33 AND age<=40       |$db->users->find(array("age" => array('$gt' => 33, '$lte' => 40)));       |
|SELECT * FROM users ORDER BY name DESC       |$db->users->find()->sort(array("name" => -1));       |
|CREATE INDEX myindexname ON users(name)       |$db->users->ensureIndex(array("name" => 1));       |
|CREATE INDEX myindexname ON users(name,ts DESC)       |$db->users->ensureIndex(array("name" => 1, "ts" => -1));       |
|SELECT * FROM users WHERE a=1 and b='q'       |$db->users->find(array("a" => 1, "b" => "q"));       |
|SELECT * FROM users LIMIT 20, 10       |$db->users->find()->limit(10)->skip(20);       |
|SELECT * FROM users WHERE a=1 or b=2       |$db->users->find(array('$or' => array(array("a" => 1), array("b" => 2))));       |
|SELECT * FROM users LIMIT 1       |$db->users->find()->limit(1);       |
|EXPLAIN SELECT * FROM users WHERE z=3       |$db->users->find(array("z" => 3))->explain()       |
|SELECT DISTINCT last_name FROM users       |$db->command(array("distinct" => "users", "key" => "last_name"));       |
|SELECT COUNT(*y) FROM users       |$db->users->count();       |
|SELECT COUNT(*y) FROM users where AGE > 30       |$db->users->find(array("age" => array('$gt' => 30)))->count();       |
|SELECT COUNT(AGE) from users       |$db->users->find(array("age" => array('$exists' => true)))->count();       |
|UPDATE users SET a=1 WHERE b='q'       |$db->users->update(array("b" => "q"), array('$set' => array("a" => 1)));       |
|UPDATE users SET a=a+2 WHERE b='q'       |$db->users->update(array("b" => "q"), array('$inc' => array("a" => 2)));       |
|DELETE FROM users WHERE z="abc"       |$db->users->remove(array("z" => "abc"));       |


### 原子性和事务处理

在 MongoDB 中，一个写操作的原子性是基于单个文档的，即使写操作是在单个文档内部更改多个嵌套文档。 

当一个写操作修改了多个文档，每个文档的更新是具有原子性的，但是整个操作作为一个整体是不具有原子性的，并且与其他操作可能会有所交替。但是，您可以使用:update:[`](http://www.mongoing.com/docs/core/write-operations-atomicity.html#id1)$isolated`操作将多个文档单的写操作*隔离*成单个的写操作， 

> update:`$isolated`操作将使写操作在集合上获得一个排他锁，甚至对于文档级别的锁存储引擎比如WiredTiger也是这样处理的。这也就是说在执行:update:`$isolated`操作运行期间会导致WiredTiger单线程运行。 