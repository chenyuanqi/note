
### 什么是 Redis
「[Redis](https://redis.io/)」是一个开源的使用C语言编写、支持网络、可基于内存亦可持久化的日志型、Key-Value数据库，并提供多种语言的 API。

### Redis 的特点
- 支持数据持久化
> Redis 可以将内存中的数据保存在磁盘中，重启的时候可以再次加载进行使用。
>   
> 方式一（RDB）：根据指定的规则，定时周期性的把内存中更新的数据写入到磁盘里。  
> RDB的方式是通过快照（snapshot）完成，当符合规则时 Redis 会把内存的数据生成一个副本并存储在硬盘中，这个过程称之为“快照”。  
> 
> 方式二（AOF）：把修改的操作记录追加到文件里，默认情况 Redis 没有开启 AOF 方式，可以通过 appendonly 命令来启用，如：appendonly yes。  
> 
> 两种方式的区别：RDB 方式性能较高，但是可能会引起一定程度的数据丢失，AOF 方式正好相反。  

- 支持主从复制
> 主机会自动将数据同步到从机，可以进行读写分离

- 丰富的数据类型
> Redis 不仅仅支持简单的key-value类型的数据，同时还提供 list，set，zset，hash 等数据结构的存储。  

- 单进程单线程高性能服务器
> 启动一个实例只能用一个 CPU，所以用 Redis 可以用多个实例，一个实例用一个 CPU 以便于提高效率。  

- crash safe 和 recovery slow
> Redis 崩溃后，数据相对安全，但是恢复起来比较缓慢，所以生产环境不建议一个 Redis 实例数据太多{（20-30）G 数据内存对应（96-128）G 实际内存）}，这种 20%-23% 的比例比较合适，因为磁盘读到内存的恢复时间也很慢，可以使用 ssd 磁盘来提高磁盘读取速度。  

- 性能极高
> 读写性能优异，Redis 单机 qps（每秒的并发）可以达到的速度是 110000 次/s，写的速度是 81000 次/s，适合小数据量高速读写访问。  

- Redis 的原子性
> Redis 的所有操作都是原子性的，意思是要么成功执行要么失败完全不执行。  
> 单个操作是原子性的。  
> 多个操作也支持事务，即原子性，通过 MULTI 和 EXEC 指令包起来。  
> 
> Redis 支持数据的备份即 master-slave 模式的数据备份。  
> Redis 还支持 publish/subscribe, key 过期。

### Redis 的缺陷与陷阱
> 1、Redis 不具备自动容错和恢复功能，主机从机的宕机都会导致前端部分读写请求失败，需要等待机器重启或者手动切换前端的 IP 才能恢复  
> 2、主机宕机，宕机前有部分数据未能及时同步到从机，切换 IP 后还会引入数据不一致的问题，降低了系统的可用性  
> 3、Redis 较难支持在线扩容，在集群容量达到上限时在线扩容会变得很复杂。为避免这一问题，运维人员在系统上线时必须确保有足够的空间，这对资源造成了很大的浪费  
> 4、内存管理开销大（不要超过物理内存的 3/5）  

### Redis 的安装
```bash
wget http://download.redis.io/releases/redis-4.0.9.tar.gz
tar -zxvf redis-4.0.9.tar.gz -C /usr/local/
cd redis-4.0.9
# 若编译失败，重新执行 make MALLOC=libc
make

# 启动 redis 服务
./redis-server
# 启动 redis 客户端，一般 ./redis-cli 即可
./redis-cli -h 127.0.0.1 -p 6379 -a "password" 
```

### Redis 的配置
Redis 的配置文件位于 Redis 的安装目录下，名为 redis.conf（Ubuntu 下，一般在 /etc/redis/6379.conf）。  
```bash
# 在 Redis 客户端下，使用命令获取所有配置项
config get *
# 获取某一项配置 
config get loglevel

# 修改配置
config set loglevel debug
```

> 配置参数大致如下  
![redis_conf](./static/redis_conf.png)

### Redis 的五种数据类型及使用
- key-value  
> Redis key 值是二进制安全的
> 
> 太长的键值不好，比如 1024 字节的键值就不好，不仅因为消耗内存，在数据中查找这类键值的成本也很高  
> 太短的键值通常也不好，键值的设置应该像变量命名一样，能标识出它的含义  
> 最好坚持一种模式，比如 “user:666:password”  
> key 建议 10 到 20 个字符  
> value 建议 string 不要超过 2K，set 元素不要超过 5000 字节，如果内容长度太多，可以根据内容长度规划不同的实例端口  
> ![redis_keys_command](./static/redis_keys_command.png)  
```bash
# 客户端和服务器连接正常
redis> ping
PONG
# 设置 key-value
redis> set name redis
OK
# 获取 key
redis> get name
"redis"
# 判断 key 是否存在
redis> exists name
(integer) 1
# 设置过期时间，过期后 key 对应的 value 变为 nil
redis> expire name 3600
(integer) 1
# 获取所有 key
redis> keys *
1) "name"
# 重命名 key
redis> rename name NAME
OK
# 查看 key 的类型
redis> type NAME
string
# 删除 key
redis> del NAME
(integer) 1
```

- 字符串（string）  
> Redis 最简单而且最常用的数据类型之一，字符串（string）类型也可以用来存储数字，并支持对数字的加减操作  
> ![redis_string_command](./static/redis_string_command.png)   
```bash
# 查看帮助（查看某个命令 help COMMAND）
redis> help @string

# 设置
redis> set name redis
OK
# 获取长度
redis> strlen name
(integer) 5
# 追加字符串
redis> append name 2018
(integer) 9
redis> get name
"redis2018"

# 使用数字
redis> set age 18
OK
# 自增 1
redis> incr age
(integer) 19
# 增加 2
redis> incrby age 2
(integer) 21
# 自减 1
redis> decr age
(integer) 20
# 减去 2
redis> decrby age 2
(integer) 18
```

- 哈希（hash）  
> Redis hash 是一个 string 类型的 field 和 value 的映射表，hash 特别适合用于存储对象，能够存储 key 对多个属性的数据  
> ![redis_hash_command](./static/redis_hash_command.png)   
```bash
# 查看帮助（查看某个命令 help COMMAND）
redis> help @hash

# 设置 hash key 的 field 和 value
redis> hset user name redis
(integer) 1
redis> hset user age 18
(integer) 1
# 获取 hash key 的 field 对应的 value
redis> hget user name
"redis"
# 获取 hash key 所有的 field
redis> hkeys user
1) "name"
2) "age"
# 获取 hash key 所有的 value
redis> hvals user
1) "redis"
2) "18"
# 获取 hash key 的 field 数量
redis> hlen user
(integer) 2
# 获取 hash key 所有的 field 和 value
redis> hgetall user
1) "name"
2) "redis"
3) "age"
4) "18"
# 扫描 hash key 的键值对
redis> hscan user 0
1) "0"
2) 1) "name"
   2) "redis"
   3) "age"
   4) "18"
```

- 列表（list）  
> Redis 列表就是有序元素的序列，比如 1,2,3,4,5,6 就是一个列表；  
> Redis 列表是简单的字符串列表，按照插入顺序排序，可以在头部或尾部添加元素  
>   
> Redis 列表是基于 Linked List 实现（用数组实现的 List 和用 Linked List 实现的 List，在属性方面大不相同）  
> 这意味着即使在一个列表中有数百万个元素，在头部或尾部添加一个元素的操作，其时间复杂度也是常数级别的  
> 也就是说，用 LPUSH 命令在十个元素的列表头部添加新元素，和在千万元素列表头部添加新元素的速度相同  
> ![redis_list_command](./static/redis_list_command.png)   
```bash
# 查看帮助（查看某个命令 help COMMAND）
redis> help @list

# 设置列表元素（可以 1 个或多个元素）
redis> lpush demo 1 2 3
(integer) 3
# 查看列表长度
redis> llen demo
(integer) 3
# 获取指定范围的元素
redis> lrange demo 0 1
1) "3"
2) "2"
# 移除一个元素
redis> lpop demo
"3"
# 移除指定个数的元素
redis> lrem demo 2 2
(integer) 1
redis> lrange demo 0 3
1) "1"
```

- 集合（set）
> Redis 集合是 string 类型的无序集合  
> 集合成员是唯一的，这就意味着集合中不能出现重复的数据  
> ![redis_set_command](./static/redis_set_command.png)   
```bash
# 查看帮助（查看某个命令 help COMMAND）
redis> help @set

# 添加集合元素（可以 1 个或多个元素）
redis> sadd lang english japanese
(integer) 2
# 获取集合元素个数
redis> scard lang
(integer) 2
# 判断集合是否存在某元素
redis> sismember lang chinese
(integer) 0
# 获取集合所有元素
redis> smembers lang
1) "japanese"
2) "english"
```

- 有序集合（sorted set）
> Redis 有序集合和集合一样也是 string 类型元素的集合，且不允许重复的成员  
> 不同的是，Redis 有序集合每个元素都会关联一个 double 类型的分数，redis 正是通过分数来为集合中的成员进行从小到大的排序  
> ![redis_sorted_set_command](./static/redis_sorted_set_command.png)   
```bash
# 查看帮助（查看某个命令 help COMMAND）
redis> help @sorted_set

# 添加有序集合元素
redis> zadd country 1 china
(integer) 1
redis> zadd country 2 US
(integer) 1
redis> zadd country 3 japan
(integer) 1
# 查看集合元素的分数
redis> zscore country china
"1"
# 查看集合
redis> zrange country 0 3
1) "china"
2) "US"
3) "japan"
# 逆向查看集合
redis> zrevrange country 0 3
1) "japan"
2) "US"
3) "china"
# 根据分数范围获取集合元素
redis> zrangebyscore country 2 3
1) "US"
2) "japan"
```

### Redis 的集群
Redis 有 3 种集群策略

- 主从
> 1 台机器可写，作为主；另外 2 台可读，作为从，类似于 MySQL 的主从复制，不过 Redis 没有 BINLOG 机制

- 哨兵
> 增加一台机器作为哨兵，监控 3 台主从机器  
> 当主节点挂机的时候，机器内部进行选举，从集群中从节点里指定一台机器升级为主节点，从而实现高可用  
> 当主节点恢复的时候，加入到从节点中继续提供服务

- 集群
> Redis3.0 以后增加了集群的概念，可以实现多主多从结构，实现正真的高可用

### Redis 与 PHP
「[phpredis](https://github.com/phpredis/phpredis)」是使用 c 写的 php 扩展，「[predis](https://github.com/nrk/predis)」是使用纯 php 写的第三方包。  
在性能上，当然是 phpredis 扩展更好一些。  

phpredis 扩展和 predis 在连接的保持上是有区别的
> phpredis 在扩展中使用 c 可以保持 php-fpm 到 redis 的长连接，所以一个 php-fpm 进程上的多个请求是复用同一个连接的  
> phpredis 的 pconnect 就是长连接方式  
> predis 是使用 php 的 socket 来连接 redis，所以需要每次请求连接 redis  

phpredis 和 predis 的性能差距没有跨数量级。  
根据业务场景，如果业务非常依赖 redis，并且单机 qps 需要支持的比较大，建议使用 phpredis；如果是小规模业务，建议使用 predis（laravel 官方推荐使用，便捷）。  