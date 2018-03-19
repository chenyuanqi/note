
### Kylin 是什么
Apache Kylin 是一个开源的分布式分析引擎，最初由 eBay 开发贡献至开源社区。  
它提供 Hadoop 之上的 SQL 查询接口及多维分析（OLAP）能力以支持大规模数据，能够处理 TB 乃至 PB 级别的分析任务，能够在亚秒级查询巨大的 Hive 表，并支持高并发。  

Kylin 的使命是让大数据分析像使用数据库一样简单迅速，用户的查询请求可以在秒内返回。  

### Kylin 工作原理

### Kylin 主要特点
1、标准 SQL 接口
2、支持超大数据集
3、得益于预计算，亚秒级响应
4、可伸缩性和高吞吐率
5、BI 及可视化工具集成

### Kylin 工具 —— Kybot 
Kybot = Kylin + Robot  
[Kybot](https://kybot.io/home) 在线诊断，优化及服务 Kylin。
Kybot 平台含仪表盘、系统优化、故障排除及知识库。  

> 使用 Kybot 步骤：
> 1、注册 Kybot 官网账号
> 2、下载 Kybot client
> 3、执行 kybot/kybot.sh 生成诊断包
> 4、上传诊断包

### Kylin 与其他 OLAP 大数据分析框架
OLAP（On-Line Analytical Processing）线上分析处理，是一套以多维度方式分析数据，而能弹性地提供积存、下钻、和透视分析等操作，呈现集成性决策信息的方法，多用于决策支持系统、商务智能或数据仓库。其主要的功能，在于方便大规模数据分析及统计计算，对决策提供参考和支持。  

> kylin: 核心是 cube，cube 是一种预计算技术，基本思路是预先对数据作多维索引，查询时只扫描索引而不访问原始数据从而提速；支持标准 SQL  
> presto：facebook 开源的一个用 java 写的分布式数据查询框架，原生集成了 Hive、Hbase 和关系型数据库;
>         presto 背后所使用的执行模式与 Hive 有根本的不同，它没有使用 MapReduce，大部分场景下比 Hive 快一个数量级，其中的关键是所有的处理都在内存中完成；
>         presto 支持标准 SQL  
> druid：是一个实时处理时序数据的 Olap 数据库，因为它的索引首先按照时间分片，查询的时候也是按照时间线去路由索引；SQL 支持不好   
> spark SQL：基于 spark 平台上的一个 olap 框架，本质上也是基于 DAG 的 MPP， 基本思路是增加机器来并行计算，从而提高查询速度；没有可视化操作界面  

