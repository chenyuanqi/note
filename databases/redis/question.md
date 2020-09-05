

### Redis 常见问题
1、Redis 属于单线程还是多线程？不同的版本有什么区别？  
在 Redis 4.0 之前，Redis 是单线程运行的，但单线程并不意味着性能低，类似单线程的程序还有 Nginx 和 NodeJs 他们都是单线程程序，但是效率并不低。Redis 的 FAQ（Frequently Asked Questions，常见问题）也回到过[这个问题](https://redis.io/topics/faq)。  Redis 是基于内存操作的，因此他的瓶颈可能是机器的内存或者网络带宽而并非 CPU，既然 CPU 不是瓶颈，那么自然就采用单线程的解决方案了，况且使用多线程比较麻烦。Redis 之所以在 4.0 之前一直采用单线程的模式是因为以下三个原因：  

- 使用单线程模型是 Redis 的开发和维护更简单，因为单线程模型方便开发和调试；
- 即使使用单线程模型也并发的处理多客户端的请求，主要使用的是多路复用；
- 对于 Redis 系统来说，主要的性能瓶颈是内存或者网络带宽而并非 CPU。

在 Redis 4.0 中开始支持多线程了，例如后台删除等功能。Redis 在 4.0 中引入了惰性删除（也可以叫异步删除），意思就是说我们可以使用异步的方式对 Redis 中的数据进行删除操作了，例如 unlink key / flushdb async / flushall async 等命令。这样处理的好处是不会导致 Redis 主线程卡顿，会把这些删除操作交给后台线程来执行。  
`tips：通常情况下使用 del 指令可以很快的删除数据，而当被删除的 key 是一个非常大的对象时，例如时包含了成千上万个元素的 hash 集合时，那么 del 指令就会造成 Redis 主线程卡顿，因此使用惰性删除可以有效的避免 Redis 卡顿的问题。`


扩展问题：Redis 4.0 之前是单线程的，为什么还能这么快？  
官方使用基准测试的结果是，单线程的 Redis 吞吐量可以达到 10W / 每秒。Redis 速度比较快的原因有以下几点：

- 基于内存操作：Redis 的所有数据都存在内存中，因此所有的运算都是内存级别的，所以他的性能比较高；  
- 数据结构简单：Redis 的数据结构比较简单，是为 Redis 专门设计的，而这些简单的数据结构的查找和操作的时间复杂度都是 O (1)，因此性能比较高；  
- 多路复用和非阻塞 I/O：Redis 使用 I/O 多路复用功能来监听多个 socket 连接客户端，这样就可以使用一个线程连接来处理多个请求，减少线程切换带来的开销，同时也避免了 I/O 阻塞操作，从而大大提高了 Redis 的性能；  
- 避免上下文切换：因为是单线程模型，因此就避免了不必要的上下文切换和多线程竞争，这就省去了多线程切换带来的时间和性能上的消耗，而且单线程不会导致死锁问题的发生。  

扩展问题：什么是 I/O 多路复用？  
套接字的读写方法默认情况下是阻塞的，例如当调用读取操作 read 方法时，缓冲区没有任何数据，那么这个线程就会阻塞卡在这里，直到缓冲区有数据或者是连接被关闭时，read 方法才可以返回，线程才可以继续处理其他业务。  
但这样显然降低了程序的整体执行效率，而 Redis 使用的就是非阻塞的 I/O，这就意味着 I/O 的读写流程不再是阻塞的，读写方法都是瞬间完成并返回的，也就是他会采用能读多少读多少能写多少写多少的策略来执行 I/O 操作，这显然更符合我们对性能的追求。但这种非阻塞的 I/O 依然存在一个问题，那就是当我们执行读取数据操作时，有可能只读取了一部分数据，同样写入数据也是这种情况，当缓存区满了之后我们的数据还没写完，剩余的数据何时写何时读就成了一个问题。  
而 I/O 多路复用就是解决上面的这个问题的，使用 I/O 多路复用最简单的实现方式就是使用 select 函数，此函数为操作系统提供给用户程序的 API 接口，是用于监控多个文件描述符的可读和可写情况的，这样就可以监控到文件描述符的读写事件了，当监控到相应的事件之后就可以通知线程处理相应的业务了，这样就保证了 Redis 读写功能的正常执行了。  
`tips：现在的操作系统已经很少使用 select 函数了，改为调用 epoll（linux）和 kqueue（MacOS）等函数了，因为 select 函数在文件描述符特别多时性能非常的差。`  

扩展问题：了解 Redis 6.0 多线程吗？  
Redis 4.0 版本中虽然引入了多线程，但此版本中的多线程只能用于大数据量的异步删除，然而对于非删除操作的意义并不是很大。  
如果我们使用多线程就可以分摊 Redis 同步读写 I/O 的压力，以及充分的利用多核 CPU 的资源，并且可以有效的提升 Redis 的 QPS。在 Redis 中虽然使用了 I/O 多路复用，并且是基于非阻塞 I/O 进行操作的，但 I/O 的读和写本身是堵塞的，比如当 socket 中有数据时，Redis 会通过调用先将数据从内核态空间拷贝到用户态空间，再交给 Redis 调用，而这个拷贝的过程就是阻塞的，当数据量越大时拷贝所需要的时间就越多，而这些操作都是基于单线程完成的。  
因此在 Redis 6.0 中新增了多线程的功能来提高 I/O 的读写性能，他的主要实现思路是将主线程的 IO 读写任务拆分给一组独立的线程去执行，这样就可以使多个 socket 的读写可以并行化了，但 Redis 的命令依旧是由主线程串行执行的。需要注意的是 Redis 6.0 默认是禁用多线程的，可以通过修改 Redis 的配置文件 redis.conf 中的 io-threads-do-reads 等于 true 来开启多线程，完整配置为 io-threads-do-reads true，除此之外我们还需要设置线程的数量才能正确的开启多线程的功能，同样是修改 Redis 的配置，例如设置 io-threads 4 表示开启 4 个线程。  
Redis 6 引入的多线程 I/O 特性对性能提升至少是一倍以上。  
`tips：关于线程数的设置，官方的建议是如果为 4 核的 CPU，建议线程数设置为 2 或 3，如果为 8 核 CPU 建议线程数设置为 6，线程数一定要小于机器核数，线程数并不是越大越好。`  


2、Redis 有哪些数据类型？  
Redis 最常用的数据类型有 5 种：String（字符串类型）、Hash（字典类型）、List（列表类型）、Set（集合类型）、ZSet（有序集合类型）。  

- 字符串类型（Simple Dynamic Strings 简称 SDS），译为：简单动态字符串，它是以键值对 key-value 的形式进行存储的，根据 key 来存储和获取 value 值。字符串的常见使用场景：存放用户（登录）信息；存放文章详情和列表信息；存放和累计网页的统计信息（存储 int 值）...  
- 字典类型 (Hash) 又被成为散列类型或者是哈希表类型，它是将一个键值 (key) 和一个特殊的 “哈希表” 关联起来，这个 “哈希表” 表包含两列数据：字段和值。可以使用字典类型来存储用户信息，并且使用字典类型来存储此类信息就无需手动序列化和反序列化数据了，所以使用起来更加的方便和高效。通常情况下字典类型会使用数组的方式来存储相关的数据，但发生哈希冲突时才会使用链表的结构来存储数据。  
- 列表类型 (List) 是一个使用链表结构存储的有序结构，它的元素插入会按照先后顺序存储到链表结构中，因此它的元素操作 (插入和删除) 时间复杂度为 O (1)，所以相对来说速度还是比较快的，但它的查询时间复杂度为 O (n)，因此查询可能会比较慢。列表的典型使用场景有两个：消息队列 —— 列表类型可以使用 rpush 实现先进先出的功能，同时又可以使用 lpop 轻松的弹出（查询并删除）第一个元素，所以列表类型可以用来实现消息队列；文章列表 —— 对于博客站点来说，当用户和文章都越来越多时，为了加快程序的响应速度，我们可以把用户自己的文章存入到 List 中，因为 List 是有序的结构，所以这样又可以完美的实现分页功能，从而加速了程序的响应速度。  
- 集合类型 (Set) 是一个无序并唯一的键值集合。集合类型的经典使用场景：微博关注我的人和我关注的人都适合用集合存储，可以保证人员不会重复；中奖人信息也适合用集合类型存储，这样可以保证一个人不会重复中奖。集合类型（Set）和列表类型（List）的区别如下：列表可以存储重复元素，集合只能存储非重复元素；列表是按照元素的先后顺序存储元素的，而集合则是无序方式存储元素的。  
- 有序集合类型 (Sorted Set) 相比于集合类型多了一个排序属性 score（分值），对于有序集合 ZSet 来说，每个存储元素相当于有两个值组成的，一个是有序结合的元素值，一个是排序值。有序集合的存储元素值也是不能重复的，但分值是可以重复的。有序集合的经典使用场景：学生成绩排名；粉丝列表，根据关注的先后时间排序。  

我们经常会把用户的登录信息存储在 Redis 中，但通常的做法是先将用户登录实体类转为 JSON 字符串存储在 Redis 中，然后读取时先查询数据再反序列化为 User 对象，这个过程看似没什么问题，但我们可以有更优的解决方案来处理此问题，比如我们可以使用 Hash 存储用户的信息，这样就无需序列化的过程了，并且读取之后无需反序列化，直接使用 Map 来接收就可以了，这样既提高了程序的运行速度有省去了序列化和反序列化的业务代码。


3、如何实现查询附近的人？  
查询附近的人或者是附近的商家是一个实用且常用的功能，比如微信中 “附近的人” 或是美团外卖中 “附近商家” 等。  
如何实现的呢？我们所处的任何位置都可以用经度和纬度来标识，经度的范围 -180 到 180，纬度的范围为：-90 到 90。纬度以赤道为界，赤道以南为负数，赤道以北为正数；经度以本初子午线 (英国格林尼治天文台) 为界，东边为正数，西边为负数。这样我们所处的位置才能在地球上被标注出来，这也成为了我们能够查询出两点之间距离的基础。  
从而让查询附近的人变得简单了，我们只需要查询出附近几个点和自己的距离，再进行排序就可以实现查询附近人的功能了，然而使用 Redis 让这一切更简单了，Redis 为我们提供了专门用来存储地理位置的类型 GEO，我们使用它以及它所内置的方法就可以轻松的实现查询附近的人了。我们可以使用 Redis 3.2 版本中新增了 GEO 类型，以及它的 georadius 命令来实现查询附近的人。  
`添加位置信息我们需要使用 geoadd 命令，它的语法为：geoadd key longitude latitude member [longitude latitude member ...]。`  
`georadius 命令的相关语法为： georadius key longitude latitude radius m|km|ft|mi [WITHCOORD] [WITHDIST] [WITHHASH] [COUNT count] [ASC|DESC]`  

查询附近的人看似是一个复杂的问题，但深入研究之后发现借助 Redis 还是很好实现的，但别高兴的太早这只是入门题。  
扩展问题：如何查询位置的经纬度信息？  
我们可以借助在线的坐标查询系统来获取经纬度的值，例如百度的坐标系统。  
扩展问题：如何在代码中实现查询附近的人？  
```java
import redis.clients.jedis.GeoCoordinate;
import redis.clients.jedis.GeoRadiusResponse;
import redis.clients.jedis.GeoUnit;
import redis.clients.jedis.Jedis;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class GeoHashExample {
    public static void main(String[] args) {
        Jedis jedis = new Jedis("127.0.0.1", 6379);
        Map<String, GeoCoordinate> map = new HashMap<>();
        // 添加小明的位置
        map.put("xiaoming", new GeoCoordinate(116.404269, 39.913164));
        // 添加小红的位置
        map.put("xiaohong", new GeoCoordinate(116.36, 39.922461));
        // 添加小美的位置
        map.put("xiaomei", new GeoCoordinate(116.499705, 39.874635));
        // 添加小二
        map.put("xiaoer", new GeoCoordinate(116.193275, 39.996348));
        jedis.geoadd("person", map);
        // 查询小明和小红的直线距离
        System.out.println("小明和小红相距：" + jedis.geodist("person", "xiaoming",
                "xiaohong", GeoUnit.KM) + " KM");
        // 查询小明附近 5 公里的人
        List<GeoRadiusResponse> res = jedis.georadiusByMemberReadonly("person", "xiaoming",
                5, GeoUnit.KM);
        for (int i = 1; i < res.size(); i++) {
            System.out.println("小明附近的人：" + res.get(i).getMemberByString());
        }
    }
}
```
扩展问题：GEO 类型的底层是如何实现的？  
GEO 类型的底层其实是借助 ZSet（有序集合）实现的，因此我们可以使用 zrem 命令来删除地理位置信息。Redis 内部使用 ZSet 来保存位置对象的，它使用 ZSet 的 Score 来存储经纬度对应的 52 位的 GEOHASH 值的。  
```java
void geoaddCommand(client *c) {
    // 参数校验
    if ((c->argc - 2) % 3 != 0) {
        /* Need an odd number of arguments if we got this far... */
        addReplyError(c, "syntax error. Try GEOADD key [x1] [y1] [name1] "
                         "[x2] [y2] [name2] ... ");
        return;
    }
    // 参数提取 Redis
    int elements = (c->argc - 2) / 3;
    int argc = 2+elements*2; /* ZADD key score ele ... */
    robj **argv = zcalloc(argc*sizeof(robj*));
    argv[0] = createRawStringObject("zadd",4);
    argv[1] = c->argv[1]; /* key */
    incrRefCount(argv[1]);
    // 参数遍历+转换
    int i;
    for (i = 0; i < elements; i++) {
        double xy[2];
        // 提取经纬度
        if (extractLongLatOrReply(c, (c->argv+2)+(i*3),xy) == C_ERR) {
            for (i = 0; i < argc; i++)
                if (argv[i]) decrRefCount(argv[i]);
            zfree(argv);
            return;
        }
        // 将经纬度转换为 52 位的 geohash 作为分值 & 提取对象名称
        GeoHashBits hash;
        geohashEncodeWGS84(xy[0], xy[1], GEO_STEP_MAX, &hash);
        GeoHashFix52Bits bits = geohashAlign52Bits(hash);
        robj *score = createObject(OBJ_STRING, sdsfromlonglong(bits));
        robj *val = c->argv[2 + i * 3 + 2];
        // 设置有序集合的对象元素名称和分值
        argv[2+i*2] = score;
        argv[3+i*2] = val;
        incrRefCount(val);
    }
    replaceClientCommandVector(c,argc,argv);
    // 调用 zadd 命令，存储转化好的对象
    zaddCommand(c);
}
```


4、Redis 如何实现限流功能？  
限流的主要目的就是为了保证整个系统的正常运行，比如以车辆限流为了，它的作用主要有两个，一个是为了保证我们生存空间的资源少受污染，尤其是近几年雾霾已经越来越严重了，如果不采取相应的手段会导致生态系统更加恶化，第二，目前车辆的增长速度已经远远的超过了市政道路的新建速度，尤其是上班的时候大家都在赶时间，如果车流量太大的话就会造成严重的交通拥堵，那么导致的直接后果就是大家上班都会迟到，为了解决这个问题所有需要限行。回到程序的这个层面也是一样，假设我们的系统只能为 10 万人同时提供购物服务，但是某一天因为老罗带货突然就涌进了 100 万用户，那么导致的直接后果就是服务器瘫痪，谁也甭想买东西了，所以这个时候我们需要 “限流” 的功能保证先让一部分用户享受购物的服务，而其他用户进行排队等待购物，这样就可以让整个系统正常的运转了。  
使用 Redis 如何实现限流功能？我们可以使用 Redis 中的 ZSet（有序集合）加上滑动时间算法来实现简单的限流。所谓的滑动时间算法指的是以当前时间为截止时间，往前取一定的时间，比如往前取 60s 的时间，在这 60s 之内运行最大的访问数为 100，此时算法的执行逻辑为，先清除 60s 之前的所有请求记录，再计算当前集合内请求数量是否大于设定的最大请求数 100，如果大于则执行限流拒绝策略，否则插入本次请求记录并返回可以正常执行的标识给客户端。（如果要精确的保证每分钟最多访问 100 次，需要记录下每次访问的时间。因此对每个用户，我们使用一个列表类型的键来记录他最近 100 次访问时间，一旦键中的元素超过 100 个，就判断时间最早的元素距离现在的时间是否小于 1 分钟。如果是则表示最近 1 分钟内的访问次数超过了 100 次；如果不是就讲现在的时间加入到列表中，同时把最早的元素删除）  
使用 ZSet 的方式加上滑动时间的算法固然可以实现简单的限流，但是这个解决方案存在一定的问题，比如当我们允许 60s 的最大访问次数为 1000w 的时候，此时如果使用 ZSet 的方式就会**占用大量的空间用来存储**请求的记录信息，并且它的判断和添加属于两条 Redis 执行指令是**非原子**单元的，所以使用它可以出现问题。那有没有更好的实现 Redis 限流的方法呢？限流的算法还有哪些？  
其他的常见限流算法还有以下两个：  

- 漏桶算法
漏桶算法的灵感源于漏斗。漏桶算法类似于生活中的漏斗，无论上面的水流倒入漏斗有多大，也就是无论请求有多少，它都是以均匀的速度慢慢流出的。当上面的水流速度大于下面的流出速度时，漏斗会慢慢变满，当漏斗满了之后就会丢弃新来的请求；当上面的水流速度小于下面流出的速度的话，漏斗永远不会被装满，并且可以一直流出。  
在一定范围内，比如 60s 内只能有 10 个请求，当第一秒时就到达了 10 个请求，那么剩下的 59s 只能把所有的请求都给拒绝掉，而漏桶算法可以解决这个问题。  
漏桶算法的实现步骤是，先声明一个队列用来保存请求，这个队列相当于漏斗，当队列容量满了之后就放弃新来的请求，然后重新声明一个线程定期从任务队列中获取一个或多个任务进行执行，这样就实现了漏桶算法。  
- 令牌算法
在令牌桶算法中有一个程序以某种恒定的速度生成令牌，并存入令牌桶中，而每个请求需要先获取令牌才能执行，如果没有获取到令牌的请求可以选择等待或者放弃执行。令牌算法最简单的实现方式是 Google 开源的 guava 包，但它是单机版的限流方案。  

有没有更好的分布式限流解决方案呢？  
答案是有的，它就是 Redis 4.0 提供的 Redis-Cell 模块，该模块使用的是漏斗算法，并且提供了原子的限流指令，而且依靠 Redis 这个天生的分布式程序就可以实现完美的限流了。Redis-Cell 实现限流的方法也很简单，只需要使用一条指令 cl.throttle 即可。  
```
> cl.throttle mylimit 15 30 60
1）（integer）0 # 0 表示获取成功，1 表示拒绝
2）（integer）15 # 漏斗容量
3）（integer）14 # 漏斗剩余容量
4）（integer）-1 # 被拒绝之后，多长时间之后再试（单位：秒）-1 表示无需重试
5）（integer）2 # 多久之后漏斗完全空出来
```


5、Redis 如何处理已经过期的数据？  
Redis 内存用完之后的内存淘汰策略，它主要是用来出来异常情况下的数据清理；而这里说的是 Redis 的键值过期之后的数据处理，是正常情况下的数据清理。  
典型的回答是：在 Redis 中维护了一个过期字典，会将所有已经设置了过期时间的键值全部存储到此字典中。获取键值的执行流程是，当有键值的访问请求时 Redis 会先判断此键值是否在过期字典中，如果没有表示键值没有设置过期时间（永不过期），然后就可以正常返回键值数据了；如果此键值在过期字典中则会判断当前时间是否小于过期时间，如果小于则说明此键值没有过期可以正常返回数据，反之则表示数据已过期，会删除此键值并且返回给客户端 nil。这是键值数据的方法流程，同时也是过期键值的判断和删除的流程。  


扩展问题：常用的删除策略有哪些？Redis 使用了什么删除策略？  
常见的过期策略，有三种：
定时删除（在设置键值过期时间时，创建一个定时事件，当过期时间到达时，由事件处理器自动执行键的删除操作；优点是保证内存可以被尽快的释放；缺点是在 Redis 高负载的情况下或有大量过期键需要同时处理时，会造成 Redis 服务器卡顿，影响主业务执行）。  
惰性删除（不主动删除过期键，每次从数据库获取键值时判断是否过期，如果过期则删除键值，并返回 null；优点因为每次访问时，才会判断过期键，所以此策略只会使用很少的系统资源；缺点是系统占用空间删除不及时，导致空间利用率降低，造成了一定的空间浪费）。  
定期删除（每隔一段时间检查一次数据库，随机删除一些过期键。 Redis 默认每秒进行 10 次过期扫描，此配置可通过 Redis 的配置文件 redis.conf 进行配置，配置键为 hz 它的默认值是 hz 10 。优点是通过限制删除操作的时长和频率，来减少删除操作对 Redis 主业务的影响，同时也能删除一部分过期的数据减少了过期键对空间的无效占用；缺点是内存清理方面没有定时删除效果好，同时没有惰性删除使用的系统资源少）。  
为了兼顾存储空间和性能，Redis 采用了惰性删除加定期删除的组合删除策略。  


6、Redis 内存用完会怎样？  
在某些极端情况下，软件为了能正常运行会做一些保护性的措施，比如运行内存超过最大值之后的处理，以及键值过期之后的处理等，都属于此类问题。  
Redis 的内存用完指的是 Redis 的运行内存超过了 Redis 设置的最大内存，此值可以通过 Redis 的配置文件 redis.conf 进行设置，设置项为 maxmemory，我们可以使用 config get maxmemory 来查看设置的最大运行内存。  
`当此值为 0 时，表示没有内存大小限制，直到耗尽机器中所有的内存为止，这是 Redis 服务器端在 64 位操作系统下的默认值。`  
当 Redis 的内存用完之后就会触发 Redis 的内存淘汰策略，Redis 内存淘汰策略可以使用 config get maxmemory-policy 命令来查看。Redis 服务器默认采用的是 noeviction 策略，此策略表示当运行内存超过最大设置内存时，不淘汰任何数据，但新增操作会报错。此值可通过修改 redis.conf 文件进行修改。Redis 的内存最大值和内存淘汰策略都可以通过配置文件进行修改，或者是使用命令行工具进行修改。使用命令行工具进行修改的优点是操作简单，成功执行完命令之后设置的策略就会生效，我们可以使用 confg set xxx 的方式进行设置，但它的缺点是不能进行持久化，也就是当 Redis 服务器重启之后设置的策略就会丢失。另一种方式就是为配置文件修改的方式，此方式虽然较为麻烦，修改完之后要重启 Redis 服务器才能生效，但优点是可持久化，重启 Redis 服务器设置不会丢失。  

扩展问题：Redis 内存淘汰策略有哪些？分别代表什么含义？  
Redis 内存淘汰在 4.0 版本之后一共有 8 种：

- noeviction：不淘汰任何数据，当内存不足时，新增操作会报错，Redis 默认内存淘汰策略；
- allkeys-lru：淘汰整个键值中最久未使用的键值；
- allkeys-random：随机淘汰任意键值；
- volatile-lru：淘汰所有设置了过期时间的键值中最久未使用的键值；
- volatile-random：随机淘汰设置了过期时间的任意键值；
- volatile-ttl：优先淘汰更早过期的键值；
- volatile-lfu：淘汰所有设置了过期时间的键值中，最少使用的键值；
- allkeys-lfu：淘汰整个键值中最少使用的键值。

扩展问题：内存淘汰策略采用了什么算法？  
内存淘汰策略决定了内存淘汰算法，从以上八种内存淘汰策略可以看出，它们中虽然具体的实现细节不同，但主要的淘汰算法有两种：LRU 算法和 LFU 算法。  
LRU 全称是 Least Recently Used 译为最近最少使用，是一种常用的页面置换算法，选择最近最久未使用的页面予以淘汰。LRU 算法需要基于链表结构，链表中的元素按照操作顺序从前往后排列，最新操作的键会被移动到表头，当需要内存淘汰时，只需要删除链表尾部的元素即可。LRU 算法有一个缺点，比如说很久没有使用的一个键值，如果最近被访问了一次，那么它就不会被淘汰，即使它是使用次数最少的缓存，那它也不会被淘汰，因此在 Redis 4.0 之后引入了 LFU 算法。  
LFU 全称是 Least Frequently Used 翻译为最不常用的，最不常用的算法是根据总访问次数来淘汰数据的，它的核心思想是 “如果数据过去被访问多次，那么将来被访问的频率也更高”。 LFU 解决了偶尔被访问一次之后，数据就不会被淘汰的问题，相比于 LRU 算法也更合理一些。 在 Redis 中每个对象头中记录着 LFU 的信息。  


7、Redis 如何实现分布式锁？  
锁是多线程编程中的一个重要概念，它是保证多线程并发时顺利执行的关键。我们通常所说的 “锁” 是指程序中的锁，也就是单机锁，例如 Java 中的 Lock 和 ReadWriteLock 等，而所谓的分布式锁是指可以使用在多机集群环境中的锁。  
首先来说 Redis 作为一个独立的三方系统（通常被作为缓存中间件使用），其天生的优势就是可以作为一个分布式系统来使用，因此使用 Redis 实现的锁都是分布式锁。使用 Redis 实现分布式锁可以通过以下两种手段来实现：

- 使用 incr 方式实现；
原理很简单，我们每次的加锁（上锁）都使用 incr 命令，如果执行的结果为 1 的话表示加锁成功，释放锁则使用 del 命令来实现。而当我们某个程序正在是使用锁时，我们继续使用 incr 会导致返回的结果不为 1，在结果不为 1 的情况下，我们就可以判断此锁正在被使用中，这样就可以实现分布式的功能了。  
```bash
127.0.0.1:6379> incr lock # 加锁
(integer) 1
127.0.0.1:6379> del lock # 释放锁
(integer) 0
127.0.0.1:6379> incr key # 第一次加锁
(integer) 1
127.0.0.1:6379> incr key # 第二次加锁
(integer) 2
```
- 使用 setnx 方式实现。
当我们使用 setnx 创建键值成功时，则表明加锁成功，否则代码加锁失败。我们可以使用执行的结果是否为 1 来判断加锁是否成功。  
```bash
127.0.0.1:6379> setnx lock true
(integer) 1 #创建锁成功
#逻辑业务处理...
127.0.0.1:6379> del lock
(integer) 1 #释放锁
127.0.0.1:6379> setnx lock true # 第一次加锁
(integer) 1
127.0.0.1:6379> setnx lock true # 第二次加锁
(integer) 0
```

扩展问题：分布式锁的死锁问题？  
死锁是并发编程中一个常见的问题，以单机锁的死锁来说，当两个线程都持有了自己锁资源并试图获取对方锁资源时就会造成死锁的诞生。那么，分布式锁的死锁是如何发生的呢？在系统中当一个程序在创建了分布式锁之后，因为某些特殊的原因导致程序意外退出了，那么这个锁将永远不会被释放，就造成了死锁的问题。    
因此为了解决死锁问题，我们最简单的方式就是设置锁的过期时间，这样即使出现了程序意外退出的情况，那么等待此锁超过了设置的过期时间之后就会释放此锁，这样其他程序就可以继续使用了。但这样依然会有问题，因为命令 setnx 和 expire 处理是一前一后非原子性的，因此如果在它们执行之间，出现断电和 Redis 异常退出的情况，因为超时时间未设置，依然会造成死锁。然而指的庆幸的是在 Redis 2.6.12 版本之后，新增了一个强大的功能，我们可以使用一个原子操作也就是一条命令来执行 setnx 和 expire 操作了。  
```bash
127.0.0.1:6379> set lock true ex 30 nx
OK #创建锁成功
127.0.0.1:6379> set lock true ex 30 nx
(nil) #在锁被占用的时候，企图获取锁失败
```
使用 set 命令只解决创建锁的问题，那执行中的极端问题，和释放锁极端问题，我们依旧要考虑。例如，我们设置锁的最大超时时间是 30s，但业务处理使用了 35s，这就会导致原有的业务还未执行完成，锁就被释放了，新的程序和旧程序一起操作就会带来线程安全的问题。执行超时的问题处理带来线程安全问题之外，还引发了另一个问题：锁被误删。 


8、如何保证 Redis 消息队列中的数据不丢失？  
消息队列（Message Queue）是一种进程间通信或同一进程的不同线程间的通信方式，它的实现流程是一方会将消息存储在队列中，而另一方则从队列中读取相应的消息，消息队列提供了异步的通信协议，也就是说消息的发送者和接收者无需同时与消息队列进行交互。  
使用消息队列有如下好处：

- 削峰填谷：将某一个时刻急速上升的请求压力转嫁给消息队列来处理，保证应用系统的正常运转；
- 系统解耦：使用消息队列我们可以将生产者和消费者的耦合代码进行分离，变成两个完全独立的模块，从而更加方便维护与改造；
- 更加可靠：使用消息队列可以永久性的保留业务数据，保证了数据在传输过程中不会因为意外情况，例如掉电而造成的数据丢失；
- 扩展性：使用消息队列具备更好的可扩展性，例如用户在完善了个人信息之后，刚开始的需求是增加用户的经验值，有一天产品部门的同学要求完善了个人资料之后不但要增加用户的经验值还要加一定的虚拟金币，那么我们无需改动太多的业务代码，只需要完善了个人信息之后，给增加金币的 channel 中发送一条增加金币的消息即可，即使过两天产品的同学告诉你需要把功能改回去，你也无需改动太多的代码，只需要注释掉发送消息的代码即可，即使要扩展更多的功能也就是非常方便的。

Redis 实现消息队列的方式有四种：

- 使用 List 方式来实现消息队列，主要使用的是 lpush/rpop 来实现消息的先进先出；
- ZSet 实现方式，此方式与 List 方式类似，它是使用 zadd 和 zrangebyscore 来实现存入和读取；
- Redis 自身提供的发布订阅模式，也就是使用 Publisher (发布者) 和 Subscriber (订阅者) 来实现消息队列；
- 使用 Redis 5.0 版本中提供的 Stream 来实现消息队列，它主要使用的是 xadd/xread 来实现消息的读取和存储。

Redis 想要保住消息队列中的数据不丢失，必须要做到以下两点：

- 必须开启消息的持久化功能，负责在 Redis 重启之后消息就会全部丢失；
- 需要使用 Stream 中提供的消息确认功能，保证消息能够被正常消费。

扩展问题：如何实现消息队列？  
Stream 实现消息队列的代码如下：
```java
import com.google.gson.Gson;
import redis.clients.jedis.Jedis;
import redis.clients.jedis.StreamEntry;
import redis.clients.jedis.StreamEntryID;
import utils.JedisUtils;

import java.util.AbstractMap;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class StreamGroupExample {
    private static final String _STREAM_KEY = "mq"; // 流 key
    private static final String _GROUP_NAME = "g1"; // 分组名称
    private static final String _CONSUMER_NAME = "c1"; // 消费者 1 的名称
    private static final String _CONSUMER2_NAME = "c2"; // 消费者 2 的名称
    public static void main(String[] args) {
        // 生产者
        producer();
        // 创建消费组
        createGroup(_STREAM_KEY, _GROUP_NAME);
        // 消费者 1
        new Thread(() -> consumer()).start();
        // 消费者 2
        new Thread(() -> consumer2()).start();
    }
    /**
     * 创建消费分组
     * @param stream    流 key
     * @param groupName 分组名称
     */
    public static void createGroup(String stream, String groupName) {
        Jedis jedis = JedisUtils.getJedis();
        jedis.xgroupCreate(stream, groupName, new StreamEntryID(), true);
    }
    /**
     * 生产者
     */
    public static void producer() {
        Jedis jedis = JedisUtils.getJedis();
        // 添加消息 1
        Map<String, String> map = new HashMap<>();
        map.put("data", "redis");
        StreamEntryID id = jedis.xadd(_STREAM_KEY, null, map);
        System.out.println("消息添加成功 ID：" + id);
        // 添加消息 2
        Map<String, String> map2 = new HashMap<>();
        map2.put("data", "java");
        StreamEntryID id2 = jedis.xadd(_STREAM_KEY, null, map2);
        System.out.println("消息添加成功 ID：" + id2);
    }
    /**
     * 消费者 1
     */
    public static void consumer() {
        Jedis jedis = JedisUtils.getJedis();
        // 消费消息
        while (true) {
            // 读取消息
            Map.Entry<String, StreamEntryID> entry = new AbstractMap.SimpleImmutableEntry<>(_STREAM_KEY,
                    new StreamEntryID().UNRECEIVED_ENTRY);
            // 阻塞读取一条消息（最大阻塞时间 120s）
            List<Map.Entry<String, List<StreamEntry>>> list = jedis.xreadGroup(_GROUP_NAME, _CONSUMER_NAME, 1,
                    120 * 1000, true, entry);
            if (list != null && list.size() == 1) {
                // 读取到消息
                Map<String, String> content = list.get(0).getValue().get(0).getFields(); // 消息内容
                System.out.println("Consumer 1 读取到消息 ID：" + list.get(0).getValue().get(0).getID() +
                        " 内容：" + new Gson().toJson(content));
            }
        }
    }
    /**
     * 消费者 2
     */
    public static void consumer2() {
        Jedis jedis = JedisUtils.getJedis();
        // 消费消息
        while (true) {
            // 读取消息
            Map.Entry<String, StreamEntryID> entry = new AbstractMap.SimpleImmutableEntry<>(_STREAM_KEY,
                    new StreamEntryID().UNRECEIVED_ENTRY);
            // 阻塞读取一条消息（最大阻塞时间 120s）
            List<Map.Entry<String, List<StreamEntry>>> list = jedis.xreadGroup(_GROUP_NAME, _CONSUMER2_NAME, 1,
                    120 * 1000, true, entry);
            if (list != null && list.size() == 1) {
                // 读取到消息
                Map<String, String> content = list.get(0).getValue().get(0).getFields(); // 消息内容
                System.out.println("Consumer 2 读取到消息 ID：" + list.get(0).getValue().get(0).getID() +
                        " 内容：" + new Gson().toJson(content));
            }
        }
    }
}
```
其中，jedis.xreadGroup () 方法的第五个参数 noAck 表示是否自动确认消息，如果设置 true 收到消息会自动确认 (ack) 消息，否则则需要手动确认。可以看出，同一个分组内的多个 consumer 会读取到不同消息，不同的 consumer 不会读取到分组内的同一条消息。  


9、使用 Redis 如何实现延迟队列？  
延迟消息队列在我们的日常工作中经常会被用到，比如支付系统中超过 30 分钟未支付的订单，将会被取消，这样就可以保证此商品库存可以释放给其他人购买，还有外卖系统如果商家超过 5 分钟未接单的订单，将会被自动取消，以此来保证用户可以更及时的吃到自己点的外卖，等等。  
延迟消息队列的常见实现方式是通过 ZSet 的存储于查询来实现，它的核心思想是在程序中开启一个一直循环的延迟任务的检测器，用于检测和调用延迟任务的执行。ZSet 实现延迟任务的方式有两种，第一种是利用 zrangebyscore 查询符合条件的所有待处理任务，循环执行队列任务；第二种实现方式是每次查询最早的一条消息，判断这条信息的执行时间是否小于等于此刻的时间，如果是则执行此任务，否则继续循环检测。  

扩展：如何实现一个延迟消息队列？  
可以使用 Java 语言中自带的 DelayQueue 数据类型来实现一个延迟消息队列。  
```java
public class DelayTest {
    public static void main(String[] args) throws InterruptedException {
        DelayQueue delayQueue = new DelayQueue();
        delayQueue.put(new DelayElement(1000));
        delayQueue.put(new DelayElement(3000));
        delayQueue.put(new DelayElement(5000));
        System.out.println("开始时间：" +  DateFormat.getDateTimeInstance().format(new Date()));
        while (!delayQueue.isEmpty()){
            System.out.println(delayQueue.take());
        }
        System.out.println("结束时间：" +  DateFormat.getDateTimeInstance().format(new Date()));
    }

    static class DelayElement implements Delayed {
        // 延迟截止时间（单面：毫秒）
        long delayTime = System.currentTimeMillis();
        public DelayElement(long delayTime) {
            this.delayTime = (this.delayTime + delayTime);
        }
        @Override
        // 获取剩余时间
        public long getDelay(TimeUnit unit) {
            return unit.convert(delayTime - System.currentTimeMillis(), TimeUnit.MILLISECONDS);
        }
        @Override
        // 队列里元素的排序依据
        public int compareTo(Delayed o) {
            if (this.getDelay(TimeUnit.MILLISECONDS) > o.getDelay(TimeUnit.MILLISECONDS)) {
                return 1;
            } else if (this.getDelay(TimeUnit.MILLISECONDS) < o.getDelay(TimeUnit.MILLISECONDS)) {
                return -1;
            } else {
                return 0;
            }
        }
        @Override
        public String toString() {
            return DateFormat.getDateTimeInstance().format(new Date(delayTime));
        }
    }
}
```


10、如何在海量数据中查询一个值是否存在？  
因为是海量数据，所以我们就无法将每个键值都存起来，然后再从结果中检索数据了，我们只能依靠专门处理此问题的 “特殊功能” 和相关方法来实现数据的查找。  
统计一个值是否在海量数据中可以使用布隆过滤器，布隆过滤器（Bloom Filter）是 1970 年由布隆提出的。它实际上是一个很长的二进制向量和一系列随机映射函数。布隆过滤器可以用于检索一个元素是否在一个集合中。它的优点是空间效率和查询时间都比一般的算法要好的多，缺点是有一定的误识别率和删除困难。也就是说布隆过滤器的优点就是计算和查询速度很快，但是缺点也很明显就是存在一定的误差。  
在 Redis 中布隆过滤器的用法如下：

- bf.add 添加元素；
- bf.exists 判断某个元素是否存在；
- bf.madd 添加多个元素；
- bf.mexists 判断多个元素是否存在；
- bf.reserve 设置布隆过滤器的准确率。

布隆过滤器在数据库领域的使用也比较广泛，例如：HBase、Cassandra、LevelDB、RocksDB 内部都有使用布隆过滤器。  
在 Redis 中使用布隆过滤器需要先安装布隆过滤器的相关模块，在 Redis 4.0 版本之后就可以使用 modules (扩展模块) 的方式引入相应的功能。  
Redis 布隆过滤器的实现，依靠的是它数据结构中的一个位数组，每次存储键值的时候，不是直接把数据存储在数据结构中，因为这样太占空间了，它是利用几个不同的无偏哈希函数，把此元素的 hash 值均匀的存储在位数组中，也就是说，每次添加时会通过几个无偏哈希函数算出它的位置，把这些位置设置成 1 就完成了添加操作。当进行元素判断时，查询此元素的几个哈希位置上的值是否为 1，如果全部为 1，则表示此值存在，如果有一个值为 0，则表示不存在。因为此位置是通过 hash 计算得来的，所以即使这个位置是 1，并不能确定是那个元素把它标识为 1 的，因此布隆过滤器查询此值存在时，此值不一定存在，但查询此值不存在时，此值一定不存在。并且当位数组存储值比较稀疏的时候，查询的准确率越高，而当位数组存储的值越来越多时，误差也会增大。  
```java
import redis.clients.jedis.Jedis;
import utils.JedisUtils;

import java.util.Arrays;

public class BloomExample {
    private static final String _KEY = "userlist";

    public static void main(String[] args) {
        Jedis jedis = JedisUtils.getJedis();
        for (int i = 1; i < 100001; i++) {
            bfAdd(jedis, _KEY, "user_" + i);
            boolean exists = bfExists(jedis, _KEY, "user_" + (i + 1));
            if (exists) {
                System.out.println("找到了" + i);
                break;
            }
        }
        System.out.println("执行完成");
    }

    /**
     * 添加元素
     * @param jedis Redis 客户端
     * @param key   key
     * @param value value
     * @return boolean
     */
    public static boolean bfAdd(Jedis jedis, String key, String value) {
        String luaStr = "return redis.call('bf.add', KEYS[1], KEYS[2])";
        Object result = jedis.eval(luaStr, Arrays.asList(key, value),
                Arrays.asList());
        if (result.equals(1L)) {
            return true;
        }
        return false;
    }

    /**
     * 查询元素是否存在
     * @param jedis Redis 客户端
     * @param key   key
     * @param value value
     * @return boolean
     */
    public static boolean bfExists(Jedis jedis, String key, String value) {
        String luaStr = "return redis.call('bf.exists', KEYS[1], KEYS[2])";
        Object result = jedis.eval(luaStr, Arrays.asList(key, value),
                Arrays.asList());
        if (result.equals(1L)) {
            return true;
        }
        return false;
    }
}
```


11、常用的 Redis 优化手段有哪些？  
最有效的提高 Redis 性能的方案就是在没有必要开启持久化的情况下，关闭 Redis 的持久化功能，这样每次对 Redis 的操作就无需进行 IO 磁盘写入了，因此性能会提升很多。  
其他优化 Redis 的常见手段有，缩短键值对的存储长度（键值对的长度是和性能成反比的）和不使用耗时长的 Redis 命令。

- 决定禁止使用 keys 命令；
- 避免一次查询所有的成员，要使用 scan 命令进行分批的，游标式的遍历；
- 通过机制严格控制 Hash、Set、Sorted Set 等结构的数据大小；
= 将排序、并集、交集等操作放在客户端执行，以减少 Redis 服务器运行压力；
- 删除 (del) 一个大数据的时候，可能会需要很长时间，所以建议用异步删除的方式 unlink，它会启动一个新的线程来删除目标数据，而不阻塞 Redis 的主线程。

管道技术 (Pipeline) 是客户端提供的一种批处理技术，用于一次处理多个 Redis 命令，从而提高整个交互的性能。也就是说 Pipeline 并不是 Redis 服务器的功能，而是客户端提供的批量处理 Redis 命令的功能。  
通常情况下 Redis 是单行执行的，客户端先向服务器发送请求，服务端接收并处理请求后再把结果返回给客户端，这种处理模式在非频繁请求时不会有任何问题，但如果出现集中大批量的请求时，因为每个请求都要经历先请求再响应的过程，这就会造成网络资源浪费，此时就需要管道技术来把所有的命令整合一次发给服务端，再一次响应给客户端，这样就能大大的提升了 Redis 的响应速度。  

Redis 还有一些其他的优化手段，例如以下这些：

- 使用 lazy free（延迟删除）特性；
- 设置键值的过期时间；
- 使用 slowlog 优化耗时命令；
- 避免大量数据同时失效；
- 客户端使用优化；
- 限制 Redis 内存大小；
- 使用物理机而非虚拟机安装 Redis 服务；
- 检查数据持久化策略；
- 使用分布式架构来增加读写速度。


12、如何设计不宕机的 Redis 高可用服务？  
想要设计一个高可用的 Redis 服务，那么一定要从 Redis 的多机功能来考虑，比如 Redis 的主从、哨兵以及 Redis 集群服务。  
主从同步 (主从复制) 是 Redis 高可用服务的基石，也是多机运行中最基础的一个，它是将从前的一台 Redis 服务器，变为一主多从的多台 Redis 服务器，这样我们就可以将 Redis 的读写分离，而这个 Redis 服务器也能承载更多的并发操作。  
Redis Sentinel（哨兵模式）使用监控 Redis 主从服务器的，当 Redis 的主从服务器出现问题时，可以利用哨兵模式自动的实现容灾恢复。  
Redis Cluster（集群）是 Redis 3.0 版本中推出的 Redis 集群方案，它是将数据分布在不同的服务器上，以此来降低系统对单主节点的依赖，并且可以大大的提高 Redis 服务的读写性能。Redis Cluster 拥有所有主从同步和哨兵的所有优点，并且可以实现多主多从的集群服务，相当于将单台 Redis 服务器的性能平行扩展到了集群中，并且它还有自动容灾恢复的功能。  

扩展问题：Redis 主从同步如何开启？它的数据同步方式有几种？  
我们可以将多台单独启动的 Redis 服务器设置为一个主从同步服务器，它的设置方式有两种，一种是在运行时将自己设置为目标服务器的从服务器；另一种是启动时将自己设置为目标服务器的从服务器。
在 Redis 运行过程中，我们可以使用 replicaof host port 命令，把自己设置为目标 IP 的从服务器。如果主服务设置了密码，需要在从服务器输入主服务器的密码，使用 config set masterauth 主服务密码 命令的方式。  
```bash
127.0.0.1:6379> replicaof 127.0.0.1 6380
OK
127.0.0.1:6377> config set masterauth pwd654321
OK
```

主从同步的数据同步方式有以下三种。  
① 完整数据同步  
当有新的从服务器连接时，为了保障多个数据库的一致性，主服务器会执行一次 bgsave 命令生成一个 RDB 文件，然后再以 Socket 的方式发送给从服务器，从服务器收到 RDB 文件之后再把所有的数据加载到自己的程序中，就完成了一次全量的数据同步。  
② 部分数据同步  
在 Redis 2.8 之前每次从服务器离线再重新上线之前，主服务器会进行一次完整的数据同步，然后这种情况如果发生在离线时间比较短的情况下，只有少量的数据不同步却要同步所有的数据是非常笨拙和不划算的，在 Redis 2.8 这个功能得到了优化。 Redis 2.8 的优化方法是当从服务离线之后，主服务器会把离线之后的写入命令存储在一个特定大小的队列中，队列是可以保证先进先出的执行顺序的，当从服务器重写恢复上线之后，主服务会判断离线这段时间内的命令是否还在队列中，如果在就直接把队列中的数据发送给从服务器，这样就避免了完整同步的资源浪费。  
③ 无盘数据同步  
从前面的内容我们可以得知，在第一次主从连接的时候，会先产生一个 RDB 文件，再把 RDB 文件发送给从服务器，如果主服务器是非固态硬盘的时候，系统的 I/O 操作是非常高的，为了缓解这个问题，Redis 2.8.18 新增了无盘复制功能，无盘复制功能不会在本地创建 RDB 文件，而是会派生出一个子进程，然后由子进程通过 Socket 的方式，直接将 RDB 文件写入到从服务器，这样主服务器就可以在不创建 RDB 文件的情况下，完成与从服务器的数据同步。  
要使用无须复制功能，只需把配置项 repl-diskless-sync 的值设置为 yes 即可，它默认配置值为 no。  

Redis Cluster 是无代理模式去中心化的运行模式，客户端发送的绝大数命令会直接交给相关节点执行，这样大部分情况请求命令无需转发，或仅转发一次的情况下就能完成请求与响应，所以集群单个节点的性能与单机 Redis 服务器的性能是非常接近的，因此在理论情况下，当水平扩展一倍的主节点就相当于请求处理的性能也提高了一倍，所以 Redis Cluster 的性能是非常高的。  
Redis Cluster 的搭建方式有两种，一种是使用 Redis 源码中提供的 create-cluster 工具快速的搭建 Redis 集群环境，另一种是配置文件的方式手动创建 Redis 集群环境。create-cluster 搭建的方式虽然速度很快，但是该方式搭建的集群主从节点数量固定以及槽位分配模式固定，并且安装在同一台服务器上，所以只能用于测试环境。  
故障发现里面有两个重要的概念：疑似下线 (PFAIL-Possibly Fail) 和确定下线 (Fail)。集群中的健康监测是通过定期向集群中的其他节点发送 PING 信息来确认的，如果发送 PING 消息的节点在规定时间内，没有收到返回的 PONG 消息，那么对方节点就会被标记为疑似下线。一个节点发现某个节点疑似下线，它会将这条信息向整个集群广播，其它节点就会收到这个消息，并且通过 PING 的方式监测某节点是否真的下线了。如果一个节点收到某个节点疑似下线的数量超过集群数量的一半以上，就可以标记该节点为确定下线状态，然后向整个集群广播，强迫其它节点也接收该节点已经下线的事实，并立即对该失联节点进行主从切换。这就是疑似下线和确认下线的概念，这个概念和哨兵模式里面的主观下线和客观下线的概念比较类似。  

当一个节点被集群标识为确认下线之后就可以执行故障转移了，故障转移的执行流程如下：

- 从下线的主节点的所有从节点中，选择一个从节点 (选择的方法详见下面 “新主节点选举原则” 部分)；
- 从节点会执行 SLAVEOF NO ONE 命令，关闭这个从节点的复制功能，并从从节点转变回主节点，原来同步所得的数据集不会被丢弃；
- 新的主节点会撤销所有对已下线主节点的槽指派，并将这些槽全部指派给自己；
- 新的主节点向集群广播一条 PONG 消息，这条 PONG 消息是让集群中的其他节点知道此节点已经由从节点变成了主节点，并且这个主节点已经接管了原本由已下线节点负责处理的槽位信息；
- 新的主节点开始处理相关的命令请求，此故障转移过程完成。  

新主节点选举的方法是这样的：

- 集群的纪元 (epoch) 是一个自增计数器，初始值为 0；
- 而每个主节点都有一次投票的机会，主节点会把这一票投给第一个要求投票的从节点；
- 当从节点发现自己正在复制的主节点确认下线之后，就会向集群广播一条消息，要求所有有投票权的主节点给此从节点投票；
- 如果有投票权的主节点还没有给其他人投票的情况下，它会向第一个要求投票的从节点发送一条消息，表示把这一票投给这个从节点；
- 当从节点收到投票数量大于集群数量的半数以上时，这个从节点就会当选为新的主节点。


13、其他问题  
**Redis 持久化**  
Redis 持久化总共有以下三种方式：

- 快照方式（RDB, Redis DataBase）将某一个时刻的内存数据，以二进制的方式写入磁盘；
- 文件追加方式（AOF, Append Only File），记录所有的操作命令，并以文本的形式追加到文件中；
- 混合持久化方式，Redis 4.0 之后新增的方式，混合持久化是结合了 RDB 和 AOF 的优点，在写入的时候，先把当前的数据以 RDB 的形式写入文件的开头，再将后续的操作命令以 AOF 的格式存入文件，这样既能保证 Redis 重启时的速度，又能减低数据丢失的风险。

RDB 优点如下

- RDB 的内容为二进制的数据，占用内存更小，更紧凑，更适合做为备份文件；
- RDB 对灾难恢复非常有用，它是一个紧凑的文件，可以更快的传输到远程服务器进行 Redis 服务恢复；
- RDB 可以更大程度的提高 Redis 的运行速度，因为每次持久化时 Redis 主进程都会 fork () 一个子进程，进行数据持久化到磁盘，Redis 主进程并不会执行磁盘 I/O 等操作；
- 与 AOF 格式的文件相比，RDB 文件可以更快的重启。

RDB 缺点如下

- 因为 RDB 只能保存某个时间间隔的数据，如果中途 Redis 服务被意外终止了，则会丢失一段时间内的 Redis 数据；
- RDB 需要经常 fork () 才能使用子进程将其持久化在磁盘上。如果数据集很大，fork () 可能很耗时，并且如果数据集很大且 CPU 性能不佳，则可能导致 Redis 停止为客户端服务几毫秒甚至一秒钟。

Redis 默认是关闭 AOF 持久化的，想要开启 AOF 持久化，有以下两种方式：通过命令行的方式；通过修改配置文件的方式（redis.conf）。  

AOF 优点如下

- AOF 持久化保存的数据更加完整，AOF 提供了三种保存策略：每次操作保存、每秒钟保存一次、跟随系统的持久化策略保存，其中每秒保存一次，从数据的安全性和性能两方面考虑是一个不错的选择，也是 AOF 默认的策略，即使发生了意外情况，最多只会丢失 1s 钟的数据；
- AOF 采用的是命令追加的写入方式，所以不会出现文件损坏的问题，即使由于某些意外原因，导致了最后操作的持久化数据写入了一半，也可以通过 redis-check-aof 工具轻松的修复；
- AOF 持久化文件，非常容易理解和解析，它是把所有 Redis 键值操作命令，以文件的方式存入了磁盘。即使不小心使用 flushall 命令删除了所有键值信息，只要使用 AOF 文件，删除最后的 flushall 命令，重启 Redis 即可恢复之前误删的数据。

AOF 缺点如下

- 对于相同的数据集来说，AOF 文件要大于 RDB 文件；
- 在 Redis 负载比较高的情况下，RDB 比 AOF 性能更好；
- RDB 使用快照的形式来持久化整个 Redis 数据，而 AOF 只是将每次执行的命令追加到 AOF 文件中，因此从理论上说，RDB 比 AOF 更健壮。

RDB 和 AOF 持久化各有利弊，RDB 可能会导致一定时间内的数据丢失，而 AOF 由于文件较大则会影响 Redis 的启动速度，为了能同时使用 RDB 和 AOF 各种的优点，Redis 4.0 之后新增了混合持久化的方式。

在开启混合持久化的情况下，AOF 重写时会把 Redis 的持久化数据，以 RDB 的格式写入到 AOF 文件的开头，之后的数据再以 AOF 的格式化追加的文件的末尾。  

混合持久化优点如下：

- 混合持久化结合了 RDB 和 AOF 持久化的优点，开头为 RDB 的格式，使得 Redis 可以更快的启动，同时结合 AOF 的优点，有减低了大量数据丢失的风险。

混合持久化缺点如下：

- AOF 文件中添加了 RDB 格式的内容，使得 AOF 文件的可读性变得很差；
- 兼容性差，如果开启混合持久化，那么此混合持久化 AOF 文件，就不能用在 Redis 4.0 之前版本了。 

**缓存雪崩**  
缓存雪崩是指在短时间内，有大量缓存同时过期，导致大量的请求直接查询数据库，从而对数据库造成了巨大的压力，严重情况下可能会导致数据库宕机的情况叫做缓存雪崩。  
崩对系统造成的影响，那如何解决缓存雪崩的问题？ 缓存雪崩的常用解决方案有以下几个。  
1. 加锁排队  
加锁排队可以起到缓冲的作用，防止大量的请求同时操作数据库，但它的缺点是增加了系统的响应时间，降低了系统的吞吐量，牺牲了一部分用户体验。  
```java
// 缓存 key
String cacheKey = "userlist";
// 查询缓存
String data = jedis.get(cacheKey);
if (StringUtils.isNotBlank(data)) {
    // 查询到数据，直接返回结果
    return data;
} else {
    // 先排队查询数据库，在放入缓存
    synchronized (cacheKey) {
        data = jedis.get(cacheKey);
        if (!StringUtils.isNotBlank(data)) { // 双重判断
            // 查询数据库
            data = findUserInfo();
            // 放入缓存
            jedis.set(cacheKey, data);
        }
        return data;
    }
}
```
2. 随机化过期时间  
为了避免缓存同时过期，可在设置缓存时添加随机时间，这样就可以极大的避免大量的缓存同时失效。   
```java
// 缓存原本的失效时间
int exTime = 10 * 60;
// 随机数生成类
Random random = new Random();
// 缓存设置
jedis.setex(cacheKey, exTime+random.nextInt(1000) , value);
```
3. 设置二级缓存  
二级缓存指的是除了 Redis 本身的缓存，再设置一层缓存，当 Redis 失效之后，先去查询二级缓存。 例如可以设置一个本地缓存，在 Redis 缓存失效的时候先去查询本地缓存而非查询数据库。  

**缓存穿透**  
缓存穿透是指查询数据库和缓存都无数据，因为数据库查询无数据，出于容错考虑，不会将结果保存到缓存中，因此每次请求都会去查询数据库，这种情况就叫做缓存穿透。  
缓存穿透会给数据库造成很大的压力。 缓存穿透的解决方案有以下几个。  
1. 使用过滤器  
我们可以使用过滤器来减少对数据库的请求，例如使用布隆过滤器，它的原理是将数据库的数据哈希到 bitmap 中，每次查询之前，先使用布隆过滤器过滤掉一定不存在的无效请求，从而避免了无效请求给数据库带来的查询压力。  
2. 缓存空结果  
另一种方式是我们可以把每次从数据库查询的数据都保存到缓存中，为了提高前台用户的使用体验 (解决长时间内查询不到任何信息的情况)，我们可以将空结果的缓存时间设置的短一些，例如 3-5 分钟。  

**缓存击穿**  
缓存击穿指的是某个热点缓存，在某一时刻恰好失效了，然后此时刚好有大量的并发请求，此时这些请求将会给数据库造成巨大的压力，这种情况就叫做缓存击穿。  
它的解决方案有以下 2 个。  
1. 加锁排队  
此处理方式和缓存雪崩加锁排队的方法类似，都是在查询数据库时加锁排队，缓冲操作请求以此来减少服务器的运行压力。  
2. 设置永不过期  
对于某些热点缓存，我们可以设置永不过期，这样就能保证缓存的稳定性，但需要注意在数据更改之后，要及时更新此热点缓存，不然就会造成查询结果的误差。  

**缓存预热**  
首先来说，缓存预热并不是一个问题，而是使用缓存时的一个优化方案，它可以提高前台用户的使用体验。 缓存预热指的是在系统启动的时候，先把查询结果预存到缓存中，以便用户后面查询时可以直接从缓存中读取，以节约用户的等待时间。   
缓存预热的实现思路有以下三种：

- 把需要缓存的方法写在系统初始化的方法中，这样系统在启动的时候就会自动的加载数据并缓存数据；
- 把需要缓存的方法挂载到某个页面或后端接口上，手动触发缓存预热；
- 设置定时任务，定时自动进行缓存预热。
