
### 什么是 Elasticsearch
Elaticsearch 简写是 ES， Elasticsearch 是一个开源的高扩展的分布式全文检索引擎，它可以近乎实时的存储、检索数据；本身扩展性很好，可以扩展到上百台服务器，处理 PB 级别的数据。  
Elasticsearch 使用 Java 开发，并且使用 Lucene 作为其核心来实现所有索引和搜索的功能，Elasticsearch 的目的是通过简单的 RESTful API 来隐藏 Lucene 的复杂性，从而让全文搜索变得简单。  

> Lucene 只是一个库。  
> 如果想要直接使用它，你必须使用 Java 来作为开发语言并将其直接集成到你的应用中；糟糕的是，Lucene 非常复杂，你需要深入了解检索的相关知识来理解它是如何工作的。  
> Elasticsearch 通过简单的 RESTful API 来隐藏了 Lucene 的复杂性，从而让全文搜索变得简单。  

Elasticsearch 主要解决如下问题：  
> 1、快速检索相关数据  
> 2、能够统计检索结果  

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
> 同一个分片（Shard）的备份数据，一个分片可能会有 0 个或多个副本，这些副本中的数据保证强一致或最终一致

- 全文检索
> 全文检索是对一篇文章进行索引，可以根据关键字搜索，类似于 mysql 的模糊查询（like "%keyword%"）   
> 全文索引是把内容根据词的意义进行分词，然后分别创建索引  

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
> Mapping 类似于关系型数据库的表、字段、表和字段的关系

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

### Elasticsearch 的一些链接
[国外社区](https://discuss.elastic.co/)  
[国内社区](https://elasticsearch.cn/)  

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
```