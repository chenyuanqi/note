
### sqlite 是什么
SQLite，是一款轻型的数据库，是遵守 ACID 的关系数据库管理系统，SQLite 是 D.RichardHipp 用 C 语言编写的开源嵌入式数据库引擎。它是完全独立的，不具有外部依赖性。  
SQLite 可以在所有主要的操作系统上运行，并且支持大多数计算机语言。  

区别于其他数据库引擎，SQLite 引擎不是一个独立的进程，而是根据应用程序要求进行动态或静态连接到程序中成为它的一个主要部分。SQLite 直接访问其存储文件。整个数据库 (定义、表、索引和数据本身) 都在宿主主机上存储在一个单一的文件中。

SQLite 第一个 Alpha 版本诞生于 2000 年 5 月，至今已经有 19 个年头。  

**sqlite 特性**  

- ACID 事务
- 零配置 -- 无需安装和管理配置
- 支持大多数 SQL92 标准，比如：ATTACH DATABASE，BEGIN TRANSACTION，COMMENT 等等，同时它还支持事务处理功能
- 单个数据库的所有信息储存在单一磁盘文件中，适用于作为应用程序文件的存储格式
- 支持数据库大小扩展至 2TB
- 源码代码量小，大致 13 万行 C 代码，4.43M
- 简单易用的 API
- 用 ANSI-C 编写，同时加入 TCL 绑定。可通过 Wrapper 支持其他语言的绑定
- 良好注释的源代码，并且有着 90% 以上的测试覆盖率
- 可作为单独头文件包，编译简单，易于添加到大型项目中
- 独立：没有额外依赖
- 支持多平台多开发语言：UNIX (Linux, Mac OS-X, Android, IOS), Windows (Win32, WinCE, WinRT), C, PHP, Perl, Java, C#,Python, Ruby
- 源码完全的开源，你可以用于任何用途，包括出售它
- 命令行接口，可通过命令行来管理数据库

### 为什么要 sqlite
sqlite 试图为单独的应用程序和设备提供本地的数据存储，sqlite 强调经济性、效率性、可靠性、独立性，和简单性。  

**sqlite 使用场景**  
1、作为应用程序文件格式  
相对于使用 fopen() 以 XML，JSON，CSV 或者其他序列化方法将应用程序数据写入到磁盘，使用 SQLite 数据库来保存数据不仅可以避免写入过程中的解析器故障，而且数据将具有良好的跨平台性。另一方面，对数据库的事务特点可以保证数据更新的一致性。  

2、作为小程序的数据库  
SQLite 具有较小的代码量，拥有较好的内存和磁盘管理机制，而且不需要维护，可靠。因此适合作为手机，PDA,MP3 播放器等设备的数据库引擎  

3、作为网页站点的数据库  
由于网站通常需要将一些无配置信息存放在磁盘上，SQLite 适合于作为中小型站点的数据库  

4、作为企业数据库管理系统  

### sqlite 安装
Windows 下，访问 [sqlite 下载页面](http://www.sqlite.org/download.html) 下载相应版本的压缩包，解压后添加文件夹路径到 PATH 环境变量即可使用。  
Linux 下，几乎所有操作系统自带 sqlite，如果未安装则访问 [sqlite 下载页面](http://www.sqlite.org/download.html) 下载 sqlite-autoconf-*.tar.gz，然后执行如下命令。  
```bash
tar xvzf sqlite-autoconf-*.tar.gz
cd sqlite-autoconf-*
./configure --prefix=/usr/local
make && make test
make install
```

### sqlite 命令
```bash
# 进入 sqlite 命令行
sqlite3
# 查看帮助文档
sqlite>.help
# 退出
sqlite>.quit
```

### sqlite 语法
sqlite 不区分大小写的，但也有一些命令是大小写敏感的，比如 GLOB 和 glob 在 SQLite 的语句中有不同的含义。  

SQLite 注释是附加的注释，可以在 SQLite 代码中添加注释以增加其可读性。  
```bash
.help -- 这是一个简单的注释
```

所有的 sqlite 语句可以以任何关键字开始，如 SELECT、INSERT、UPDATE、DELETE、ALTER、DROP 等，所有的语句以分号（;）结束。  
语句跟 Mysql 大同小异，如下简单示例。
```bash
SELECT column1, column2....columnN
FROM   table_name
WHERE  CONDITION-1 {AND|OR} CONDITION-2;

ALTER TABLE table_name ADD COLUMN column_def...;

ALTER TABLE table_name RENAME TO new_table_name;

ATTACH DATABASE 'DatabaseName' As 'Alias-Name';

BEGIN;
or
BEGIN EXCLUSIVE TRANSACTION;
...
COMMIT;
```

### sqlite 数据类型
实际上，sqlite 是无类型的。这意味着你可以保存任何类型的数据到你所想要保存的任何表的任何列中，无论这列声明的数据类型是什么 (只有自动递增 Integer Primary Key 才有用)。

- NULL：空值
- INTEGER：带符号的 32 位整型，具体取决有存入数字的范围大小
- SMALLINT：16 位的整型
- REAL：浮点数字，存储为 8-byte IEEE 浮点数
- DECIMAL(p,s)：精确值 p 是指全部有几个数 (digits) 大小值，s 是指小数点后有几位数，默认 p 为 5，s 为 0
- FLOAT:32 位实数
- DOUBLE:64 位实数
- GRAPHIC(n)：单位是双字节的 n 长度字串，n 不超过 127
- CHAR(n)：n 长度的字串，n 不超过 254
- VARCHAR(n)：长度不固定且最大长度为 n 的字串，n 不超过 4000
- VARGRAPHIC(n)：长度不固定且最大长度为 n 的双字节字串，n 不超过 2000
- TEXT：字符串文本
- BLOB：二进制对象
- DATE：包含年月日
- TIME：包含时分秒
- DATETIME：日期时间格式，如 2019-09-09
- TIMESTAMP：包含年、月、日、时、分、秒、千分之一秒
