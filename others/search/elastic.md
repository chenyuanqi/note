
### 什么是 Elasticsearch
Elaticsearch 简写是 ES， Elasticsearch 是一个开源的高扩展的分布式全文检索引擎，它可以近乎实时的存储、检索数据；本身扩展性很好，可以扩展到上百台服务器，处理 PB 级别的数据。  
Elasticsearch 使用 Java 开发，并且使用 Lucene 作为其核心来实现所有索引和搜索的功能，Elasticsearch 的目的是通过简单的 RESTful API 来隐藏 Lucene 的复杂性，从而让全文搜索变得简单。  

> Lucene 只是一个库  
> 如果想要直接使用它，你必须使用 Java 来作为开发语言并将其直接集成到你的应用中；糟糕的是，Lucene 非常复杂，你需要深入了解检索的相关知识来理解它是如何工作的  
> Elasticsearch 通过简单的 RESTful API 来隐藏了 Lucene 的复杂性，从而让全文搜索变得简单  

Elasticsearch 主要解决如下问题：  
> 1、快速检索相关数据  
> 2、能够统计检索结果  

### Elasticsearch 的安装配置
> java 版本要求：最低 1.7 

- Elasticsearch 安装
> [ELK 部署](https://blog.csdn.net/laoyang360/article/details/73368740)  
> [Windows 部署](https://blog.csdn.net/laoyang360/article/details/51900235)  

- Elasticsearch 配置
> [Mapping 模板](https://blog.csdn.net/laoyang360/article/details/78396928)  
> [高性能配置清单](https://blog.csdn.net/laoyang360/article/details/77985822)  

- Elasticsearch 集群
> [集群部署](https://blog.csdn.net/laoyang360/article/details/72850834)  

- Elasticsearch 插件
> [Head](https://blog.csdn.net/laoyang360/article/details/51472821)  
> [Kibana](https://blog.csdn.net/laoyang360/article/details/51472888)  
> [Logstash](https://blog.csdn.net/laoyang360/article/details/51472914)  
> [Marvel](https://blog.csdn.net/laoyang360/article/details/51472902)  
> [Graph](https://blog.csdn.net/laoyang360/article/details/51472931)  
> [IK](https://blog.csdn.net/laoyang360/article/details/51472953)  
> 
> Github  
> [head](https://github.com/mobz/elasticsearch-head)  
> [bigdesk](https://github.com/hlstudio/bigdesk)  
> [kopt](https://github.com/lmenezes/elasticsearch-kopf)  
> [sql](https://github.com/NLPchina/elasticsearch-sql)  
> [ik](https://github.com/medcl/elasticsearch-analysis-ik)  
> [pinyin](https://github.com/gitchennan/elasticsearch-analysis-lc-pinyin)  
> [同义词](https://github.com/bells/elasticsearch-analysis-dynamic-synonym)  
> [简繁转换](https://github.com/medcl/elasticsearch-analysis-stconvert)  

### Elasticsearch 的工作原理
ElasticSearch 的节点启动后，它会默认使用多播的方式（multicast，使用单播需要修改配置）寻找集群中的其它节点，并与之建立连接；利用倒排索引检索后并把相应结果返回。  
> 常规索引：文档 -> 关键词的映射过程（正向索引），耗时，需遍历整个文档  
> 倒排索引：关键词 -> 文档的映射，把正向索引的结果重新构造成倒排索引[矩阵]  
> 倒排索引（inverted index）也称为反向索引、置入档案或反向档案，是一种索引方法，被用来存储在全文搜索下某个单词在一个文档或一组文档中的存储位置的映射
> 倒排索引是文档检索系统中常用的、最快的数据结构

### Elasticsearch 的核心概念
- 集群（Cluster）
> Elasticearch 可以作为一个独立的搜索服务器；但是，为了处理大型数据集，实现容错和高可用性，Elasticearch 可以运行在许多台互相合作的服务器上  
> 这些服务器的集合称为 Elasticearch 集群

- 节点（Node）
> 物理概念，一个运行的 Elasticearch 实例，一般是一台机器上的一个进程

- 索引（Index）
> 逻辑概念，包括配置信息 mapping 和倒排正排数据文件，一个索引的数据文件可能会分布于一台机器，也有可能分布于多台机器  
> 索引的另外一层意思是倒排索引文件  

- 分片（Shard）
> 为了支持更大量的数据，索引一般会按某个维度分成多个部分，每个部分就是一个分片，分片被节点（Node）管理  
> 一个节点（Node）一般会管理多个分片，这些分片可能是属于同一份索引，也有可能属于不同索引，但是为了可靠性和可用性，同一个索引的分片尽量会分布在不同节点（Node）上  
> 分片有两种，主分片和副本分片  

- 副本（Replica）
> 同一个分片（Shard）的备份数据，一个分片可能会有 0 个或多个副本，这些副本中的数据保证强一致或最终一致。副本可以在分片故障时提供备用服务，保证数据不丢失，多个副本还可以提升搜索操作的吞吐量和性能。  

- 全文检索
> 全文检索是对一篇文章进行索引，可以根据关键字搜索，类似于 mysql 的模糊查询（like "%keyword%"）   
> 全文索引是把内容根据词的意义进行分词，然后分别创建索引  

- elasticsearch document 路由原理
> 路由算法：shard = hash(routing) % number_of_primary_shards  
> 
> 一个 index 数据会被分为多片，每片在一个 shard 中，当客户端创建 document 时，elasticsearch 就需要决定 document 放在这个 index 的哪个 shard，这个过程就叫 document routing，即数据路由。  
> primary shard 一旦由 index 建立是不可修改的，但是 replica shard 却是可以随时修改。  
> 默认的 routing 就是 _id，也可以在发送请求的时候，手动指定一个routing value，比如说 put /index/type/id?routing=user_id。  
> 手动指定 routing value 是很有用的，可以保证某一类 document 一定被路由到一个 shard 上去，那么在后续进行应用级别的负载均衡，以及提升批量读取的性能的时候，是很有帮助的。

- elasticsearch document 增删改内部原理
> 1、客户端选择一个 node 发送请求过去，这个 node 就是 coordinating node（协调节点）  
> 2、coordinating node，对 document 进行路由，将请求转发给对应的 node（有 primary shard）  
> 3、实际的 node 上的 primary shard 处理请求，然后将数据同步到 replica node  
> 4、coordinating node，如果发现 primary node 和所有 replica node 都搞定之后，就返回响应结果给客户端  

- elasticsearch document 写一致原理
> 我们在发送任何一个增删改操作的时候，比如说 put /index/type/id，都可以带上一个 consistency 参数，指明我们想要的写一致性是什么？  
> 如 put /index/type/id?consistency=quorum  
> one：要求这个写操作，只要有一个 primary shard 是 active 活跃可用的，就可以执行  
> all：要求这个写操作，必须所有的 primary shard 和 replica shard 都是活跃的，才可以执行这个写操作  
> quorum：默认的值，要求所有的 shard 中，必须是大部分的 shard 都是活跃可用的，才可以执行这个写操作  
> 
> quroum = int( (primary + number_of_replicas) / 2 ) + 1，当 number_of_replicas > 1 时才生效  
> 如果节点数少于 quorum 数量，可能导致 quorum 不齐全，进而导致无法执行任何写操作  
> 如果 quorum 不齐全时，默认等待 1 分钟；等待期间，期望活跃的 shard 数量可以增加，最后实在不行，就会 timeout  
> 其实，写操作时加一个 timeout 参数，如 put /index/type/id?timeout=30，即设定 quorum 不齐全时，timeout 时长可以缩短也可以增长  
- elasticsearch document 查询原理  
> 对于读请求，不一定就将请求转发到 primary shard 上，也可以转发到 replica shard 上，因为 replica shard 是可以服务所有读请求的。  
> 1、客户端发送请求到任意一个 node，成为 coordinate node（协调节点）  
> 2、coordinate node 对 document 进行路由（之后就知道在哪个 primary shard 上），将请求转发到对应的 node，此时会使用 round-robin 随机轮询算法，在 primary shard 以及其所有 replica 中随机选择一个，让读请求负载均衡  
> 3、接收请求的 node 返回 document 给 coordinate node  
> 4、coordinate node 返回 document 给客户端  
> 
> 特殊情况：document 如果还在建立索引过程中，可能只有 primary shard 有，任何一个 replica shard 都没有，此时可能会导致无法读取到 document，但是 document 完成索引建立之后，primary shard 和 replica shard 就都有了  

- timeout 机制  
> 默认情况下，没有所谓的 timeout，即每个 shard 搜索特别慢，需要花费几分钟的时间，那么搜索请求也会等待几分钟才会返回。  
> timeout 机制，指定每个 shard，就只能在 timeout 时间范围内，将搜索到的部分数据（也有可能全部搜索到了），直接返回给 client 程序，而不是等到所有的数据都搜索出来才返回。确切的说，一次搜索请求可以在用户指定的 timeout 时长内完成，为一些时间敏感的应用提供了良好的支持。

- elasticsearch 搜索原理
> multi-index 和 multi-type 搜索模式  
> multi-index 和 multi-type 搜索模式可以一次性搜索多个 index 和多个 type 下的数据。  
> 
> /_search：所有索引，所有type下的所有数据都搜索出来  
> /index1/_search：指定一个index，搜索其下所有type的数据  
> /index1,index2/_search：同时搜索两个index下的数据
> /*1,*2/_search：按照通配符去匹配多个索引  
> /index1/type1/_search：搜索一个index下指定的type的数据  
> /index1/type1,type2/_search：可以搜索一个index下多个type的数据  
> /index1,index2/type1,type2/_search：搜索多个index下的多个type的数据  
> /_all/type1,type2/_search：_all，可以代表搜索所有index下的指定type的数据  
> 
> client 发送一个搜索请求，会把请求打到所有的 primary shard 上执行，因为每个 shard 都包含部分数据，所以每个 shard 都可能会包含搜索请求的结果。但是，如果 primary shard 有 replica shard，那么请求也可以打到 replica shard 上。  

- elasticsearch document _all metadata 原理和作用
> elasticsearch 中的 _all 元数据，在建立索引的时候，我们插入一条 document，它里面包含了多个 field，此时，elasticsearch 会自动将多个 field 的值，全部用字符串的方式串联起来，变成一个长的字符串，作为 _all field 的值，同时建立索引。  
> 之后，如果在搜索的时候，没有对某个 field 指定搜索，就默认搜索 _all field，其中是包含了所有 field 的值。  
> 如下例子，"vikey 27 vikey@xxx.com shenzhen"将作为这一条 document 的 _all field 的值，同时进行分词后建立对应的倒排索引。  
```json
{
  "name": "vikey",
  "age": 27,
  "email": "vikey@xxx.com",
  "address": "shenzhen"
}
```

- filter 与 query 对比
> filter，仅仅只是按照搜索条件过滤出需要的数据而已，不计算任何相关度分数，对相关度没有任何影响；还有内置的自动 cache 最常使用 filter 的数据  
> query，会去计算每个document相对于搜索条件的相关度，并按照相关度进行排序；而且无法 cache 结果  
> 
> 一般来说，如果你是在进行搜索，需要将最匹配搜索条件的数据先返回，那么用 query；如果你只是要根据一些条件筛选出一部分数据，不关注其排序，那么用 filter  
> 除非是你的这些搜索条件，你希望越符合这些搜索条件的 document 越排在前面返回，那么这些搜索条件要放在 query 中；如果你不希望一些搜索条件来影响你的 document 排序，那么就放在 filter 中即可

- 字符串如何排序？
> 如果对一个 string field 进行排序，结果往往不准确，因为分词后是多个单词，再排序就不是我们想要的结果了；通常的解决方案是，将一个 string field 建立两次索引，一个使用分词，用来进行搜索；一个不分词（"index": "not_analyzed"），用来进行排序

- elasticsearch document 打分规则（关联匹配程度）  
> elasticsearch 打分使用的是 TF/IDF（term frequency/inverse document frequency） 算法  
> 1、Term frequency：搜索文本中的各个词条在field文本中出现了多少次，出现次数越多，就越相关  
> 2、Inverse document frequency：搜索文本中的各个词条在整个索引的所有文档中出现了多少次，出现的次数越多，就越不相关  
> 3、Field-length norm：field 长度越长，相关度越弱  
> 
> 分析 _score 是如何被计算出来的：GET /index/type/_search?explain -d "{}"

- elasticsearch doc value
> 搜索的时候，要依靠倒排索引；排序的时候，需要依靠正排索引，看到每个 document 的每个 field，然后进行排序，所谓的正排索引，其实就是 doc values  
> 在建立索引的时候，一方面会建立倒排索引，以供搜索用；一方面会建立正排索引，也就是doc values，以供排序，聚合，过滤等操作使用
doc values 是被保存在磁盘上的，此时如果内存足够，os 会自动将其缓存在内存中，性能还是会很高；如果内存不足够，os 会将其写入磁盘上

- query phase
> 1、搜索请求发送到某一个 coordinate node，构构建一个 priority queue，长度以 paging 操作 from 和 size 为准，默认为 10  
> 2、coordinate node 将请求转发到所有 shard，每个 shard 本地搜索，并构建一个本地的 priority queue  
> 3、各个 shard 将自己的 priority queue 返回给 coordinate node，并构建一个全局的 priority queue  

### Elasticsearch 数据架构的主要概念
- 索引（Index）
> 类似于关系型数据库中的数据库（DataBase）

- 类型（Type）
> 类似于一个关系型数据库中的数据表（Table），可以有多个（但是 Elasticsearch 6.x 将废除）

- 文档（Document）
> 类似于一个关系型数据库中数据表的每一行记录（ROW）

- 列（Field）
> 类似于关系型数据库的字段属性（Column）

- 映射（Mapping）
> Mapping 定义索引（Index）下的类型（Type）的字段处理规则，即索引如何建立、索引类型、是否保存原始索引 JSON 文档、是否压缩原始 JSON 文档、是否需要分词处理、如何进行分词处理等  
> Mapping，就是 index 的 type 的元数据，每个 type 都有一个自己的 Mapping，决定了数据类型，建立倒排索引的行为，还有进行搜索的行为  
> Mapping 类似于关系型数据库的表、字段、表和字段的关系  
> Mapping 核心数据类型：string，byte，short，integer，long，float，double，boolean，date  
> 查看 Mapping：GET /index/_mapping/type  
> 
> dynamic mapping，自动为我们建立 index，创建 type，以及 type 对应的 mapping，mapping 中包含了每个 field 对应的数据类型，以及如何分词等设置  
> 可以手动在创建数据之前，先创建 index 和 type，以及 type 对应的 mapping  
> 
> 只能创建 index 时手动建立 mapping，或者新增 field mapping，但是不能 update field mapping
```
# 新增 field mapping
PUT /index
{
  "mappings": {
    "article": {
      "properties": {
        "author_id": {
          "type": "integer"
        }
      }
    }
  }
}
```

- 复制（Replicas）  
> 拷贝索引一份或多份  

- 增删改查（crud）
> Elasticsearch 有增 PUT/POST、删 Delete、改 _update、查 GET 等 Api 操作  
> 这就类似于关系型数据库的增 Insert、删 Delete、改 Update、查 Search

- 版本控制  
> 内部版本控制，_version 自增长，修改数据后 _version 自增 1  
> 外部版本控制（version_type=external），为了保持内外版本的一致性，需要检查请求中的 _version 值是否大于当前的 _version 值  

- 锁机制
> 乐观锁，假设不会发生并发冲突，只在提交操作时检查是否违反数据。Elasticsearch 的锁机制使用的就是乐观锁  
> 悲观锁，假设会发生并发冲突，屏蔽一切可能违反数据完整性的操作  

- 端口  
> ElasticSearch 一般会有两个端口，一个是集群维护端口（默认 9300）， 一个是对外提供的端口（默认 9200）  

- 分词器
> 切分词语，提升召回率（recall，搜索的时候，增加能够搜索到的结果的数量）  
> character filter：在一段文本进行分词之前，先进行预处理  
> tokenizer：分词  
> 
> 一个分词器，将一段文本进行各种处理，最后处理好的结果才会拿去建立倒排索引  

### Elasticsearch 的特点和优势
> 1、分布式实时文件存储，可将每一个字段存入索引，使其可以被检索到  
> 2、实时分析的分布式搜索引擎  
> 分布式：将索引拆分成多个分片，每个分片可以有零个或者多个副本  
> 集群中的每个数据节点都可以承载一个或多个分片，并且协调处理各种操作；负载再平衡和路由在大多数情况下自动完成  
> 3、可以运行在单台服务器上；也可以扩展到上百台的服务器，处理 PB 级别的结构化或非结构化数据  
> 4、支持插件机制，分词插件、同步插件、Hadoop插件、可视化插件等

### Elasticsearch 的集群部署
随着业务的拓展、数据量的增多，部署分布式 Elasticsearch 集群刻不容缓。  

Elasticsearch 集群中节点一般有三种角色：主节点、客户端节点和数据节点。  
> 主节点（master node）主要用于元数据（metadata）的处理，比如索引的新增、删除、分片分配等  
> 客户端节点（client node）起到路由请求的作用，实际上可以看做是一个负载均衡器  
> 数据节点（data node）保存了数据分片，它负责数据的相关操作，比如分片的 CRUD，以及搜索和整合操作，这些操作都比较消耗 CPU、内存和 I/O 资源  
> master node, data node, client node 都各自有一份自己的数据，只是 master 具备所有读写权限，而 data node 和 client node 只读
> 需要修改数据时，需要通过 master，然后同步到各个节点  
> master 在扩展中，可能遭遇瓶颈（可能性低）
> master 通过配置可做为 master、data、client 三者

需要注意的是，**各节点间 hostname 需要和节点名称设置为一致，并且需要关闭防火墙（防火墙会导致无法正常通信，head 插件不能看到节点数据信息）**。

### Elasticsearch 的索引存储
Elasticsearch 1.X 版本的存储模块可以控制索引数据的存储方式（内存或磁盘上）。使用内存方式可以得到更好的性能，但是受限于实际可用的物理内存大小。虽然索引数据可以存储到内存中，但是相比基于 mmap 的存储方式，并没有太多的性能改善。  
Elasticsearch 2.X 版本删除直接存储内存的方式，并在内存中映射索引使用 mmap 系统调用（缺省默认值default_fs，以达到最优性能）。  

在创建索引的时候，可以指定 setting 的配置项 index.store.type。  
> 在 Windows x64 系统上使用 mmapfs   
> 在 Windows x86 系统上使用 simplefs   
> 在 Linux 系统上默认使用 default_fs (hybrid niofs 和 mmapfs)

### Elasticsearch 的疑难杂症
- 如何选择版本
> 新手建议从最新版本学起  
> 已经接触过 2.x 以前的版本，为了提升性能及系统的稳定性，也建议了解新版本的新特性并升级到最新的版本

- 分片怎么配置  
> 每个节点的分片数量保持在低于每 1GB 堆内存对应集群的分片在 20-25 之间  
> 分片大小为 50GB 通常被界定为适用于各种用例的限制

- Elasticsearch 有哪些推荐配置
> 禁用自动创建索引 action.auto_create_index: false  
> 关闭自动发现节点 discovery.zen.ping.multicast.enabled: false  

- Elasticsearch 健康状态怎样
> 请求 GET /_cat/health?v  
> status=green 高可用，必须存在副本  
> status=yellow 副本不正常，副本数量为 0  
> status=red 数据丢失，副本过多，会消耗资源，影响写入、更新速度、查询吞吐量（主副分片的数据同步有时间差，可指定查询副本，但实时性不太好）  

- bulk 批处理的最佳处理记录是多少
> 最佳值跟硬件配置、文档复杂度以及当前集群的负载有关，可根据仪表性能的消耗来判断  

- 怎样实现统计去重（distinct）
> [参考链接](https://github.com/laoyang360/deep_elasticsearch/blob/master/04_distinct)  

- 相关性打分机制是怎样的
> ElasticSearch 采用的默认相关性打分采用的是 Lucene 的 [TF-IDF 技术](https://lucene.apache.org/core/5_5_0/core/org/apache/lucene/search/similarities/TFIDFSimilarity.html)  
> 更多参考[权威指南 > 相关性](https://es.xiaoleilu.com/056_Sorting/90_What_is_relevance.html)、[相关性打分机制](https://zhuanlan.zhihu.com/p/27951938)  

- 更新和删除文档的过程是怎样的
> 删除和更新文档都是写操作，但是 Elasticsearch 的文档是不可变的，因此不能被删除或者改动以展示其变更  
> 磁盘上的每个段都有一个相应的 .del 文件，当删除请求发送后，文档并没有真的被删除，而是在 .del 文件中被标记为删除  
> 该文档依然能匹配查询，但是会在结果中被过滤  
> 当段合并时，在 .del 文件中被标记为删除的文档将不会被写入新段  
> 在新的文档被创建时，Elasticsearch 会为该文档指定一个版本号  
> 当文档执行更新时，旧版本的文档在 .del 文件中被标记为删除，新版本的文档被索引到一个新段，旧版本的文档依然能匹配查询，但是会在结果中被过滤掉

- 搜索的过程是怎样的
> 搜索分成两个过程：初始查询阶段和取回阶段  
> 
> 初始查询阶段，查询会广播到索引中每一个分片拷贝（主分片或者副本分片），每个分片在本地执行搜索并构建一个匹配文档的大小为 from + size 的优先队列
> 每个分片返回各自优先队列中所有文档的 ID 和排序值给协调节点，协调节点合并这些值到自己的优先队列中，从而产生一个全局排序后的结果列表
> 
> 取回阶段，协调节点辨别出哪些文档需要被取回并向相关的分片提交多个 GET 请求，每个分片加载并丰富文档
> 如果有需要的话，接着返回文档给协调节点。一旦所有的文档都被取回了，协调节点返回结果给客户端

- http 协议中 get 是否可以带上 request body  
> HTTP 协议，一般不允许 get 请求带上 request body，但是因为 get 更加适合描述查询数据的操作，因此还是这么用了  
> 碰巧，很多浏览器，或者是服务器，也都支持 GET + request body 模式；如果遇到不支持的场景，也可以用 POST /_search  

### Elasticsearch 的一些链接
[国外社区](https://discuss.elastic.co/)  
[国内社区](https://elasticsearch.cn/)  

[Elasticsearch 官方文档](https://www.elastic.co/guide/index.html)  
[Elasticsearch 5.4 中文文档](http://cwiki.apachecn.org/pages/viewpage.action?pageId=4260364)  
[Elasticsearch 权威指南](https://es.xiaoleilu.com/)  
[ELK 教程](http://www.cnblogs.com/xing901022/p/4704319.html)  

### Elasticsearch 的增删改查
Elasticsearch 的增删改查都是通过发起 HTTP 请求；请求方式繁多，尤其是查询，建议根据需要查阅[官方文档](https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html)。

```
# 新增记录（不带 ID，随机分配）
POST /index/type?pretty
{
  "field_name": "field_value"
}

# 删除文档
DELETE /index/type/[id]?pretty

# 更新数据
POST /index/type/[id]/_update?pretty
{
  "doc":{"field_name":"field_value"}
}

# 查询全部
GET /index/_search
{
  "query": {"match_all": {}}
}

# query dsl 基本语法
{
    QUERY_NAME: {
        ARGUMENT: VALUE,
        ARGUMENT: VALUE,...
    }
}
{
    QUERY_NAME: {
        FIELD_NAME: {
            ARGUMENT: VALUE,
            ARGUMENT: VALUE,...
        }
    }
}
# 组合多个搜索条件：title 必须包含 elasticsearch，content 可以包含 elasticsearch 也可以不包含，author_id 必须不为 1  
GET /index/type/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "title": "elasticsearch"
          }
        }
      ],
      "should": [
        {
          "match": {
            "content": "elasticsearch"
          }
        }
      ],
      "must_not": [
        {
          "match": {
            "author_id": 111
          }
        }
      ]
    }
  }
}

```