
### Memcached 是什么
Memcached 是一个自由开源的，高性能，分布式内存对象缓存系统。  
Memcached 是一种基于内存的 key-value 存储，Key 限制为 250 字节，单个 Value 最大 1MB，用来存储小块的任意数据（字符串、对象）。这些数据可以是数据库调用、API 调用或者是页面渲染的结果。  
Memcached 使用了多线程模式，开启 Memcached 服务器时使用 -t 参数可以指定要开启的线程数，但并不是线程数越多越好，一般设置为 CPU 核数，这样效率最高。此外，Memcached 使用了 NIO 模型以提升并发行能。  

Memcached 简洁而强大。  
它的简洁设计便于快速开发，减轻开发难度，解决了大数据量缓存的很多问题。它的 API 兼容大部分流行的开发语言，如 C、PHP、Java、Python、Ruby 等。  

**Memcached 特征**  

- 协议简单
- 基于 libevent 的事件处理
- 内置内存存储方式
- memcached 不互相通信的分布式

### 为什么要 Memcached
保存在 memcache 中的对象实际放置在内存中，所以 Memcached 非常高效。  
在实际使用中，通常把数据库查询的结果保存到 Memcached 中，下次访问时直接从 Memcached 中读取，而不再进行数据库查询操作，在很大程度上减少了数据库的负担。  

Memcached 主要用于解决如下问题：  
1、对数据库的高并发读写  
2、对海量数据的处理  

**Memcached 使用场景**  
一般的使用 Memcached 目的是，通过缓存数据库查询结果，减少数据库访问次数，以提高动态 Web 应用的速度、提高可扩展性。  

- 在动态系统中减少数据库负载，提升性能，做缓存，适合多读少写，大数据量的情况（如人人网大量查询用户信息、好友信息、文章信息等）。总原则是将经常需要从数据库读取的数据缓存在 memcached 中，这些数据包括经常被读取并且实时性要求不强可以等到自动过期的数据（如网站首页最新文章列表、排行榜等数据，使用典型的缓存策略，设置一过合理的过期时间，当数据过期以后再从数据库中读取）、经常被读取并且实时性要求强的数据（如用户的好友列表，用户文章列表，用户阅读记录等，数据先被载入到 memcached 中，当发生更改时就清除缓存）。
- 秒杀功能：一个人下单，要牵涉数据库读取，写入订单，更改库存，及事务要求， 对于传统型数据库来说，压力是巨大的。利用 memcached 的 incr/decr 功能，在内存存储 count 库存量， 秒杀 1000 台每人抢单主要在内存操作，速度非常快，抢到 count <= 1000 号的人会得到一个订单号，再去另一个页面慢慢支付。

### Memcached 淘汰策略及内存管理 
Memecache 在容量达到指定值后，将基于 LRU（Least Recently Used，最近最少被使用）算法自动删除不使用的缓存。在某些情况下 LRU 机制也会带来麻烦，如将不期待的数据从内存中清除，这种情况下启动 Memcache，可以通过 M 参数禁止 LRU 算法。此外，Memecache 只支持单一的淘汰策略，粒度较大，须谨慎使用。  

Memcached 没有直接采用 malloc/free 管理内存，而是采用 Slab Allocation 机制管理内存。  
首先从操作系统申请一大块内存，并将其分割成各种尺寸的块 Chunk，并把尺寸相同的块分成组 Slab Class。其中，Chunk 是用来存储 Key-Value 数据的最小单位。当 Memcached 接收到客户端发送过来的数据时，首先会根据数据大小选择一个最合适的 Slab Class，并通过查询 Memcached 保存的该 Slab Class 内空闲 Chunk 的列表，就可以找到一个可用于存储数据的 Chunk。当一条数据过期或者丢弃时，该记录所占用的 Chunk 就可以回收，重新添加到空闲列表中。  
从以上过程可以看出，Memcached 的内存管理制效率高，而且不会造成内存碎片，但它最大的缺点则是会造成空间浪费。每个 Chunk 都分配了特定长度的内存空间，所以变长数据无法充分利用这些空间。比如将 64 个字节的数据缓存到 88 个字节的 Chunk 中，剩余的 24 个字节就浪费掉了。

### Memcache 高可用方案
Memcached 不支持真正意义上的集群模式，也不支持主从副本以防止单点故障。为了保障 Memcached 服务的高可用，需要借助第三方软件或者自己设计编程实现。常用的第三方软件有 Repcached、Memagent、 memcached-ha 等。  
`注意：Memcached 在实现分布式群集部署时，Memcached 服务端之间是不能进行通讯的，也就是说服务端是伪分布式的，分布式将由客户端或者代理来实现。`

- 方案 1：一致性 Hash
Memcached 本身并不支持分布式，因此，可以在客户端通过一致性哈希这样的分布式算法来实现 Memcached 的分布式存储。  
当客户端向 Memcached 集群发送数据时，首先通过一致性哈希算法计算出该条数据的目标节点，然后将数据直接发送到该节点上存储。当客户端查询数据时，同样要计算出查询数据所在的节点，之后直接向该节点发送查询请求以获取数据。  
通过一致性哈希算法可以保证数据存放到不同的 Mamcached 上，分散了在单台机器上的风险，提高了可用性，但只能解决数据全部丢失的问题，部分数据仍可能丢失，比如当一台 Mamcached 所在节点宕机，它上面的数据还是会丢失。  

- 方案 2：Repcached  
Repcached，全称 Replication Cached 高可用技术，简称复制缓冲区技术。Repcached 可用来实现 Memcached 的复制功能。它所构建的主从方案是一个单主单从方案，不支持多主多从。但是，主从两个节点可以互相读写，从而可以达到互相同步的效果。  
假设主节点坏掉，从节点会很快侦测到连接断开，然后它会自动切换到监听状态（Listen）从而成为主节点，并等待新的从节点加入。但原来挂掉的主节点恢复之后，只能作为从节点通过人工手动的方式重新启动。它并不能抢占成为新的主节点，除非新的主节点挂掉。这就意味着，基于 Repcached 实现的 Memcached 主从文案中，主节点并不具备抢占功能。  

### Memcached 安装
```bash
# ubuntu
apt-get install -y libevent ibevent-dev
apt-get install -y memcached

# centos
yum install -y libevent libevent-devel 
yum install -y memcached

# 源码安装
wget http://memcached.org/latest
tar -zxvf memcached-x.x.x.tar.gz
cd memcached-x.x.x
./configure --prefix=/usr/local/memcached
make && make test
make install
```

### Memcached 命令
```bash
# 启动
# 参数说明：-d 表示后台启动，-m 表示分配内存，-u 表示运行用户，-l 表示监听 IP，-p 表示监听端口，-c 表示最大并发数，-P 表示 pid 文件位置
memcached -d -m 128 -u root  -l 127.0.0.1 -p 1121 -c 256 -P /tmp/memcached.pid
# 关闭 
# kill `cat /tmp/memcached.pid` 
ps -ef | grep memcached
kill -9 进程号
```

**memcached 指令**  
```bash
# 连接 
# telnet HOST PORT
telnet 127.0.0.1 11211

# 添加
# add key flags exptime bytes [noreply]
# value
add test 0 60 3
xxx
# 更新
# replace key flags exptime bytes [noreply]
# value
replace test 0 60 3
x
# 设置
# set key flags exptime bytes [noreply] 
# value
set test 0 60 3
xxx
# 后追加数据
# append key flags exptime bytes [noreply]
# value
append test 0 60 3
after
# 前追加数据
# prepend key flags exptime bytes [noreply]
# value
prepend test 0 60 3
before
# 获取
# get key1 key2 key3
get test
# gets 命令比普通的 get 命令多返回了一个数字。这个数字可以检查数据是否发生改变
# 当 key 对应的数据改变时，这个多返回的数字也会改变
gets test
# cas
# cas 即 checked and set 的意思，只有当最后一个参数和 gets 所获取的参数匹配时才能存储，否则返回 “EXISTS”
# cas key flags exptime bytes unique_cas_token [noreply]
# value
cas test 0 60 3
use-cas
# 删除
# delete key [noreply]
delete test
# 清空所有
# flush_all [after_time] [noreply]
flush_all

set visitors 0 900 2
# 自增
# incr key increment_value
incr visitors 5
# 自减
# decr key decrement_value
incr visitors 1

# 统计
stats
```

使用 java 实现 MemCache 的客户端。  
```java
public class MemCacheManager
{
    private static MemCacheManager instance = new MemCacheManager();
    
    /** XMemCache允许开发者通过设置节点权重来调节MemCache的负载，设置的权重越高，该MemCache节点存储的数据越多，负载越大 */
    private static MemcachedClientBuilder mcb = 
            new XMemcachedClientBuilder(AddrUtil.getAddresses("127.0.0.1:11211 127.0.0.2:11211 127.0.0.3:11211"), new int[]{1, 3, 5});
    private static MemcachedClient mc = null;
    
    /** 初始化加载客户端MemCache信息 */
    static
    {
        mcb.setCommandFactory(new BinaryCommandFactory()); // 使用二进制文件
        mcb.setConnectionPoolSize(10); // 连接池个数，即客户端个数
        try
        {
            mc = mcb.build();
        }
        catch (IOException e)
        {
            e.printStackTrace();
        }
        
    }
    
    private MemCacheManager()
    {
        
    }
    
    public MemCacheManager getInstance()
    {
        return instance;
    }
    
    /** 向MemCache服务器设置数据 */
    public void set(String key, int expiry, Object obj) throws Exception
    {
        mc.set(key, expiry, obj);
    }
    
    /** 从MemCache服务器获取数据 */
    public Object get(String key) throws Exception
    {
        return mc.get(key);
    }
    
    /**
     * MemCache通过compare and set即cas协议实现原子更新，类似乐观锁，每次请求存储某个数据都要附带一个cas值，MemCache
     * 比对这个cas值与当前存储数据的cas值是否相等，如果相等就覆盖老数据，如果不相等就认为更新失败，这在并发环境下特别有用
     */
    public boolean update(String key, Integer i) throws Exception
    {
        GetsResponse<Integer> result = mc.gets(key);
        long cas = result.getCas();
        // 尝试更新key对应的value
        if (!mc.cas(key, 0, i, cas))
        {
            return false;
        }
        return true;
    }
}
```

### Memcached 工作原理
memcached 是一套 C/S 模式架构的软件，在服务器端启动服务守护进程，可以为 memcached 服务器指定监听的 IP 地址、端口号、并发访问连接数、以及分配多少内存来处理客户端的请求的参数；memcached 软件是由 C 语言来实现的，全部代码仅有 2000 多行，采用的是异步 I/O，其实现方式是基于事件的单进程和单线程的。使用 libevent 作为事件通知机制，多个服务器端可以协同工作，但是这些服务器端之间是没有任何通信联系的，每个服务器只对自己的数据进行管理。应用程序端通过指定缓存服务器的 IP 地址和端口，就可以连接 memcached 服务进行相互通信。  

memcached 使用多路复用 I/O 模型。多路复用 I/O 是一种消息通知模式，用户连接做好 I/O 准备后，系统会通知我们这个连接可以进行 I/O 操作，这样就不会阻塞在某个用户连接。因此，memcache 才能支持高并发。  
memcached 使用了多线程机制，可以同时处理多个请求，线程数一般设置为 CPU 核数。

需要被缓存的数据以 Key/Value 键值对的形式保存在服务器端预分配的内存空间中，每个被缓存的数据都有唯一的标识 Key，操作 memcached 中的数据是通过这个唯一标识 Key 进行的。缓存到 Memcached 中的数据仅放置在 memcached 服务预分配的内存中，而非储存在磁盘中，因此存取速度非常快；由于 Memcached 服务自身没有对缓存的数据进行持久性存储的设计（memcached 软件开发的早期，仅为缓存而设计），因此，在服务器端的 memcached 服务进程重启之后，存储在内存中的这些数据就会丢失。而且，当内存中缓存的数据容量达到启动时设定的内存值时，就自动使用 LRU（最近最少使用算法）算法删除过期的缓存数据。  

为了满足数据可以持久性保留的需求，新浪网基于 memcached 服务开发了一款 NoSQL 软件，名字叫 MemcacheDB，可以实现在缓存的基础上增加了持久缓存的特性。  

**Memcached 内存管理机制**  
Memcached 利用 Slab Allocation 机制来分配和管理内存。  
传统的内存管理方式是使用完通过 malloc 分配的内存后通过 free 来回收内存。这种方式容易产生内存碎片并降低操作系统对内存的管理效率。  
但是，Slab Allocation 机制不存在这样的问题。它按照预先分配的大小，将分配的内存分割成特定长度的内存块，再把尺寸相同的内存块分成组，这些内存块不会释放，可以重复利用。  

Memcached 服务器端保存着一个空闲的内存块列表，当有数据存入时根据接收到的数据大小，分配一个能存下这个数据的最小内存块。这种方式有时会造成内存浪费，例如：将 200 字节的一个数据存入 300 字节的一个内存块中，就会有 100 字节的内存被浪费掉，不能被使用。  
避免浪费内存的办法：  

- 先计算出应用存入的数据大小，或把同一业务类型的数据存入一个 Memcached 服务器中，确保存入的数据大小相对均匀  
- 在 memcached 服务启动时，通过 -f 选项指定一个增长因子（或叫增长系数），它能控制内存组（slab）之间的大小差异。在应用中使用 Memcached 时，通常可以不重新设置这个参数，使用默认值 1.25 进行部署。如果想优化 memcached 对内存的使用，可以考虑重新计算数据的预期平均长度，调整这个参数来获得合适的设置值。  

**Memcached 的删除机制**  
Memcached 不会释放已分配的内存空间，在数据过期后，客户端不能通过 Key 取出它的值，其存储空间被重新利用。  

Memcached 使用的是一种 Lazy Expiration 策略，自己不会监控存入的 “Key/Value” 对是否过期，而是在获取 Key 值时查看记录的时间戳，检查 “key/value” 键值对的空间是否过期，这种策略不会在过期检测上浪费 CPU 资源。  

Memcached 在分配空间时，优先使用已经过期的 Key/Value 键值对空间，当分配的内存空间占满时，Memcached 就会使用 LRU（最近最少使用算法）算法来分配空间，删除最近最少使用的 Key/Value 键值对，将其空间分配给新 Key/Value 键值对。在某些情况下，如果不想使用 LRU 算法，那么可以通过 “-M” 参数来启动 Memcached，这样 Memcached 在内存耗尽时，会返回一个报错信息。  

### Memcached 疑难杂症
- memcache 与 memcached 的区别
> memcache 是基于 php 开发的，memcached 是基于 c 语言通过 libmemcached 与 memcached 服务器通信，因此性能更好（由于需要事先安装 libmemcached，因此 Windows 下不支持），并且支持的功能特性也更多，推荐使用后者。

- add、replace 和 set 的区别
> add 方法用于向 memcache 服务器添加一个要缓存的数据，如果 key 已存在则返回 false；  
> replace 方法用于替换一个指定 key 的缓存内容，如果 key 不存在则返回 false；  
> set 方法用于设置一个指定 key 的缓存内容，set 方法是 add 方法和 replace 方法的集合体。  
