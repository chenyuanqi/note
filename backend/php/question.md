
### 开发问题汇集
- BOM 头是什么，怎么去除
> BOM 头是放在 UTF-8 编码的文件的头部的三个字符（0xEF 0xBB 0xBF，即BOM）占用三个字节，用来标识该文件属于 UTF-8 编码。现在已经有很多软件识别 BOM 头，但是还有些不能识别 BOM 头，比如 PHP 就不能识别 BOM 头，所以 PHP 编码规范PSR-4：“无 BOM 的 UTF-8 格式”。同时这也是用 Windows 记事本编辑 UTF-8 编码后执行就会出错的原因了（Windows 记事本生成文件自带 BOM）。
```php
function remove_utf8_bom($text){
    $bom = pack('H*','EFBBBF');
    $text = preg_replace("/^$bom/", '', $text);

    return $text;
}
```

- & 引用
> $a =&$b; $a 和 $b 在这里是完全相同的，这并不是 $a 指向了 $b 或者相反，而是 $a 和 $b 指向了同一个地方；
> 引用做的第二件事是用引用传递变量，如 function foo(&$var)。  
> 
> 引用不是指针，$var=&$GLOBALS["baz"]; 这样的结构不会产生预期的效果；  
> 当使用 unset 掉一个引用，只是断开了变量名和变量内容之间的绑定，这并不意味着变量内容被销毁了。  
> 
> 值得注意的是，在 foreach 中使用 & 引用后，当 foreach 结束后，$key 和 $val 的变量也都不会被自动释放掉，此时$val 和 $arr[count($arr) - 1] 指向相同的内存地址。  

- include、include_once 和 require、require_once 的区别
> include 函数：会将指定的文件读入并且执行里面的程序（失败返回 FALSE 并且发出警告；成功返回 1，除非在包含文件中另外给出返回值）；  
> require 函数：会将目标文件的内容读入，并且把自己本身代换成这些读入的内容（失败发出致命错误并终止程序；成功返回 1，除非在包含文件中另外给出返回值）；  
> include_once 函数：在脚本执行期间包含并运行指定文件。此行为和 include 语句类似，唯一区别是如果该文件中已经被包含过，则不会再次包含。如同此语句名字暗示的那样，只会包含一次（失败没有返回且发出警告；成功返回 True）；  
> require_once 函数：和 require 语句完全相同，唯一区别是 PHP 会检查该文件是否已经被包含过，如果是则不会再次包含（失败没有返回且发出致命错误并终止程序；成功返回 True）。  
>
> include 与 require 除了在处理引入文件的方式不同外，最大的区别就是：include 在引入不存文件时产生一个警告且脚本还会继续执行，而 require 则会导致一个致命性错误且脚本停止执行。  
> include 与 require 的功能相同，但在用法上却有一些不同，include 是有条件包含函数，而 require 则是无条件包含函数（即不论条件是否成立都包含到文件里）。  
> 
> include_once（require_once）语句在脚本执行期间包含并运行指定文件；  
> include_once（require_once）行为和 include（require）语句类似，区别是如果该文件中的代码已经被包含了，则不会再次包含，只会包含一次；  
> include_once（require_once）需要查询一遍已加载的文件列表, 确认是否存在, 然后再加载。  
> 通常，include_once 和 require_once 会有性能上的下降，因为他们需要判断之前是否包含过（实际上，也不太需要去考虑，除非这已经影响到你程序的性能了）。  
> 
> 如果使用 APC 尽量也不要用 include_once，因为 apc.include_once_override 配置项没有很好的被实现。  
>（Alternative PHP Cache (APC)是一种对PHP有效的开放源高速缓冲储存器工具，它能够缓存 php 的中间码 opcode）  
> 
> require 通常放在 PHP 程序的最前面，PHP 程序在执行前，就会先读入 require 所指定引入的文件，使它变成 PHP 程序网页的一部份。同时，也可以这个方法将它引入网页中。  
> include 通常放在流程控制的处理部分中。PHP 程序网页在读到 include 的文件时，才将它读进来。这种方式，可以把程序执行时的流程简单化。
> 
> return 语句在返回引用的时候加上括号不但效率低而且会造成逻辑错误。  
> 理论上，include 和 require 后面加不加括号对执行结果没有区别，但是加上括号效率较低，所以后面能不加括号就不加括号。  
> 

额外阅读：[PHP 自动加载机制](https://blog.csdn.net/zhihua_w/article/details/52723402)

- isset 和 empty 的区别
> isset() 函数 一般用来检测变量是否设置  
> 若变量不存在则返回 FALSE   
> 若变量存在且其值为 NULL，也返回 FALSE   
> 若变量存在且值不为 NULL，则返回 TURE  
> 
> empty() 函数是检查变量是否为空  
> 若变量不存在则返回 TRUE 
> 若变量存在且其值为 ""、0、"0"、NULL、、FALSE、array()、var $var; 以及没有任何属性的对象，则返回 TURE   
> 若变量存在且值不为 ""、0、"0"、NULL、、FALSE、array()、var $var; 以及没有任何属性的对象，则返回 FALSE  

- 打开php.ini中的 safe_mode，会影响哪些函数
> safe_mode，php 安全模式，它提供一个基本安全的共享环境，在一个有多个用户账户存在的php开发的web服务器上  
> 当安全模式打开的时候，一些函数将被完全的禁止，而另一些函数的功能将会受到限制，  
> 如：chdir,move_uploaded_file,chgrp,parse_ini_file,chown,rmdir,copy,rename,fopen,require,mkdir,unlink 等。  
> 
> 注意，在 php5.3 以上版本，safe_mode 被弃用，在 php5.4 以上版本，则将此特性完全去除了。

- PHP 的垃圾收集机制 —— 引用计数  
> PHP可以自动进行内存管理，清除不再需要的对象。  
> PHP 使用了引用计数(reference counting)这种单纯的垃圾回收(garbage collection)机制，  
> 每个对象都内含一个引用计数器，每个 reference 连接到对象，计数器加 1。当 reference 离开生存空间或被设为NULL，计数器减 1。  
> 当某个对象的引用计数器为零时，PHP 知道你将不再需要使用这个对象，释放其所占的内存空间。

- 论坛中无限分类的实现原理
> 就像 windows 下新建一个文件夹，在新建的文件夹下又可以新建一个文件夹，这样无限循环下去，无限分类也是这样，父类可以分 出它子类，子类又可以分出它的子类，这样一直无限循环下去。  
> 
> 怎样实现？  
> 首先，创建类别表
```mysql
CREATE TABLE category(
cat_id smallint unsigned not null auto_increment primary key comment'类别ID',
cat_name VARCHAR(30)NOT NULL DEFAULT''COMMENT'类别名称',
parent_id SMALLINT UNSIGNED NOT NULL DEFAULT 0 COMMENT'类别父ID'
)engine=MyISAM charset=utf8;
```
> 编写一个函数，递归遍历，实现无限分类
```php
function tree($arr, $pid = 0, $level = 0)
{
    static $list = [];
    foreach($arr as $v){
        //如果父级分类一致，则将其存到 $list 中，并以此节点为根节点，遍历其子节点
        if ($v['parent_id'] == $pid){
            $v['level'] = $level;
            $list[] = $v;
            tree($arr, $v['cat_id'], $level + 1);
        }
    }

    return $list;
}
```

- 简述协程的优势
> 协程的特点在于是一个线程执行  
> 最大的优势就是协程极高的执行效率。因为子程序切换不是线程切换，而是由程序自身控制，因此，没有线程切换的开销，和多线程比，线程数量越多，协程的性能优势就越明显。  
> 第二大优势就是不需要多线程的锁机制，因为只有一个线程，也不存在同时写变量冲突，在协程中控制共享资源不加锁，只需要判断状态就好了，所以执行效率比多线程高很多。  
> 因为协程是一个线程执行，那怎么利用多核CPU呢？最简单的方法是多进程+协程，既充分利用多核，又充分发挥协程的高效率，可获得极高的性能。  

- PHP 进程如何 daemon 化
> daemon 音标: [‘di:mən]，中文含义为守护神或精灵的意思，即守护进程。  
> 守护进程简单地说就是可以脱离终端而在后台运行的进程，在 Linux 中是非常常见的一种进程，比如 apache 或者 mysql 等服务启动后，就会以守护进程的方式进驻在内存中。  
> 
> 在 Linux 中，大概有三种方式实现脚本后台化  
> 1、在命令后添加一个 & 符号，比如 php task.php & 。  
> 这个方法的缺点在于如果 terminal 终端关闭（无论是正常关闭还是非正常关闭），这个php进程都会随着终端关闭而关闭，其次是代码中如果有echo或者print_r之类的输出文本 , 会被输出到当前的终端窗口中。  
> 2、使用 nohup 命令，比如 nohup php task.php & 。  
> 默认情况下，代码中 echo 或者 print_r 之类输出的文本会被输出到 php 代码同级目录的 nohup.out 文件中；  
> 如果使用 exit 命令或者关闭按钮等正常手段关闭终端，该进程不会被关闭，依然会在后台持续运行；  
> 但是，如果终端遇到异常退出或者终止，该 php 进程也会随即退出。本质上，也不属于稳定可靠的 daemon 方案。  
> 3、使用 fork 和 setsid  
```php
// 一次fork  
$pid = pcntl_fork();
if ( $pid < 0 ) {
  exit( ' fork error. ' );
} else if( $pid > 0 ) {
  exit( ' parent process. ' );
}

// 将当前子进程提升会会话组组长 这是至关重要的一步 
if ( ! posix_setsid() ) {
  exit( ' setsid error. ' );
}

// 二次 fork
$pid = pcntl_fork();
if( $pid < 0 ){
  exit( ' fork error. ' );
} else if( $pid > 0 ) {
  exit( ' parent process. ' );
}

// 真正的逻辑代码，下面以循环写入文件为示例
for( $i = 1 ; $i <= 100 ; $i++ ){
  sleep( 1 );
  file_put_contents( 'daemon.log', $i, FILE_APPEND );
}
```

- 唯一 ID 的产生与方案
> 为什么要唯一 ID ?  
> 1、数据库的自增 ID 在分库的时候, 会是一场灾难。假设分两个库, 因为每个库都会开始从 1 开始自增, 这时候系统中将会出现两个 id 为 1 的用户  
> 2、自增 ID 会暴露用户量或者其他业务量  
> 3、自增 ID 会让有心者通过API得到任意用户的信息资料  
> 
> 有哪些解决方案呢?  
> 1、UUID，全称 Universally Unique Identifier，中文通用唯一标识符。这个是开放软件基金会组织提出的一个标准，为的就是解决分布式环境下生成唯一标识符的问题。  
> UUID 的长度是固定的 32 位，组织格式 8-4-4-4-12；当然，在用的时候，中间的分隔符是要去掉的。这货有几个问题不得不提，首先是字母数字混合，在一些传统数据库下，索引不太好做，不仅索引体积大， 查询效率也差，其次是它本身也非常大。  
> 2、 MongoDB ObjectId , 格式模样都很类似于 UUID，是 Mongodb 内置的一种数据类型，如果你在插入数据的时候不指定_id，那么 Mongodb 默认就会采用用这个货才填充_id， 在 Mongodb 这种类 key-value 性质的数据库中，有着不错的查询效率  
> 3、自建解决方案。需要保证全局空间唯一性、尽量采取数字类型而非数字字母混合方式、一定的时序行和含义、一定的可反解性 , 通过反解的结果可以知道该 ID 的相关信息。  
> 市面上有的几种解决方案为 Twitter 的 snowflake，Flikr 的数据库自增方案，Instagram 的数据库存储过程方案。  
> 比如 snowflask 使用了 64bit 来表示一个 id，组织格式 41(时间戳)-10(机器 ID)-12(自增序列)。推荐基于 snowflask 的 PHP ID 产生器：[DonkeyID](https://github.com/osgochina/donkeyid)  

- 自动加载与命名空间
> PSR4 利用了命名空间和 spl_autoload_register() 在 php-fig 的倡导下形成的一种开发者约定俗成的开发标准和规则。只要开发者开发的库满足 PSR 系列的标准，那么这个库就可以在任意一个支持 PSR 标准的任意框架或项目中运行，最终诞生了伟大的 php composer。  
> 
> 众所周知，php 中有个大名鼎鼎的魔法函数__autoload() 用来实现自动加载，但是__autoload 有个巨大的缺陷，就是无法同时加载多个 autoload 方法。  
> 于是，[spl_autoload_register()](http://php.net/manual/zh/function.spl-autoload-register.php) 来了，它可以注册多个 autoload 方法，即注册不同的自动加载机制。  
> 
> 命名空间，即 [namespace](http://www.php.net/manual/zh/language.namespaces.php)，是 php5.3 以后引入的新特性，用来解决包名冲突问题的。  
> 命名空间通常对应了文件夹目录的层次关系，当然这并不是官方的规定，而是人们利用命名空间做出的一种规范，为的是方便开发和协调。  
```php
// 声明三个不同加载函数
function autoload_1( $classname ){
  echo "autoload 1";
}
function autoload_2( $classname ){
  echo "autoload 2";
}
function autoload_3( $classname ){
  echo "autoload 3";
}

// 将三个不同的函数注册到 autoload 栈中
spl_autoload_register('autoload_1');
spl_autoload_register('autoload_2');
spl_autoload_register('autoload_3');
```

- 网站大规模并发处理方案
> https://www.awaimai.com/348.html

- PHP 读取大文件问题
> PHP 默认最大只给每个进程分配 128MB 内存，读取大文件的最原始想法就是改配置，但是这显然不是好办法。  
> 
> PHP 读取大文件，PHP 提供的标准类库 [SplFileObject](http://php.net/manual/en/class.splfileobject.php) 可以作为参考。  
> fgets 按行读取比 fgetc 按字符读取要快，而如果要读取指定行的内容可以考虑 ftell 和 fseek。  
> 总而言之，要读取大文件，需要依次按适当大小获取文件内容，显然这是一种稍微优秀一些的办法。  
> 
> 我们还可以使用生成器读取大文件（当然，使用流读取速度也很可观 stream_get_line）
```php
function readBigFile($filePath)
{
    # code...
    $fp = fopen($filePath, 'rb');
    while(false !== ($buffer = fgets( $fp, 4096 ))){
        # code...
    }

    fclose($handle);
}

foreach ($readBigFile('./test.txt') as $key => $value) {
    # code...
}
```

- API 安全性问题
> 安全是恒久的话题。API 主要的安全问题比如  
> 1、接口被大规模调用消耗系统资源，影响系统的正常访问，甚至系统瘫痪  
> 2、数据泄露  
> 3、伪造（篡改）数据，制造垃圾数据  
> 4、App被仿制…  
> 
> 那么，我们又该如何设计我们的 API 呢？  
> 需要保证对受限资源的登录授权、对请求做身份认证，并且防止篡改，重放攻击和对敏感的数据做加密。  
> 
> 对受限资源的登录授权的处理流程：客户端提交账号信息（用户名+密码）到服务端->服务端验证成功，返回AccessToken给客户端存储->访问受限资源时，客户端带入AccessToken就可访问  
> 对请求做身份认证的处理流程：初始时，服务端存有各 App 版本的 SIGN_KEY，客户端存有对应版本的 SIGN_KEY->当要发送请求之前，通过签名方法加密，得到一个 sign（如 sign=signature(path?query&imei&timetamp&SIGN_KEY)），发送请求的时候，连同sign一起发送给服务器端->服务器端首先验证时间戳 timestamp 是否有效，比如是服务器时间戳 5 分钟之前和之后的请求视为无效->服务端取对应版本的 SIGN_KEY 验证 sign 是否合法->为了防止重放攻击，需要检查 sign 是否在 redis中 存储，如不存在则存入 redis（缓存 5 分钟）  
> 对敏感数据加密：部署SSL基础设施（即HTTPS），敏感数据的传输全部基于SSL；仅对部分敏感数据做加密（例如账号+密码），并加入某种随机数作为加密盐，以防范数据被篡改  
>
> 签名机制是为了防止 API 被恶意调用，包括 API  
> 加密是为了保证敏感数据，敏感数据可以包括 token（token 和 uid 对应关系可以考虑 redis hash 类型的数据结构，key 就用 token，hash 中保存完整的用户信息）  
> token 本身与加密无关，只是 token 本身的含义总是跟加密似乎带点儿关系，但实际上 token 仅仅是个用户身份识别器  
> 只要客户端被反编译了，加密方式和签名机制都会暴露出来，所以安全是需要双方配合的  

- 消息队列
> 消息队列已经逐渐成为企业应用系统内部通信的核心手段。它具有 低耦合、可靠投递、广播、流量控制、最终一致性 等一系列功能。  
> 当前使用较多的 消息队列 有 RabbitMQ、RocketMQ、ActiveMQ、Kafka、ZeroMQ、MetaMQ 等，而部分 数据库 如 Redis、MySQL 以及 phxsql 也可实现消息队列的功能。  
> 
> 消息队列 是指利用 高效可靠 的 消息传递机制 进行与平台无关的 数据交流，并基于 数据通信 来进行分布式系统的集成。  
> 通过提供 消息传递 和 消息排队 模型，它可以在 分布式环境 下提供 应用解耦、弹性伸缩、冗余存储、流量削峰、异步通信、数据同步 等等功能，其作为 分布式系统架构 中的一个重要组件，有着举足轻重的地位。  
> 消息队列的特点主要是：采用异步处理模式，应用系统之间解耦合。  
> 消息队列的传输模式有：点对点模型，发布/订阅模型  
> 
> 理解 PHP Redis 实现消息队列 ... 


- web 负载均衡
> 负载均衡（Load Balance）是集群技术（Cluster）的一种应用，它可以将工作任务分摊到多个处理单元，从而提高并发处理能力。任何的负载均衡技术都要想办法建立某种一对多的映射机制：一个请求的入口映射到多个处理请求的节点，从而实现分而治之。  
> 常见的 web 负载均衡技术包括：DNS 轮询、IP 负载均衡和 CDN。其中 IP 负载均衡可以使用硬件设备或软件方式来实现。  
> 
> DNS 轮询是最简单的负载均衡方式，以域名作为访问入口，通过配置多条 DNS A 记录使得请求可以分配到不同的服务器。  
> DNS 轮询没有快速的健康检查机制，而且只支持 WRR 的调度策略导致负载很难“均衡”，通常用于要求不高的场景。并且 DNS 轮询方式直接将服务器的真实地址暴露给用户，不利于服务器安全。  
> 
> CDN（Content Delivery Network，内容分发网络），通过发布机制将内容同步到大量的缓存节点，并在 DNS 服务器上进行扩展，找到里边用户最近的缓存节点作为服务提供节点。  
> 因为很难自建大量的缓存节点，所以通常使用 CDN 运营商的服务。目前国内的服务商很少，而且按流量计费，价格也比较昂贵。  
> 
> IP 负载均衡是基于特定的 TCP/IP 技术实现的负载均衡。比如 NAT、DR、Turning 等是最常用的方式。  
> IP 负载均衡可以使用硬件设备，也可以使用软件实现。  
> 硬件设备的主要产品是 F5-BIG-IP-GTM（简称F5)，软件产品主要有 LVS、HAProxy、Nginx。其中 LVS、HAProxy 可以工作在 4-7 层，Nginx 工作在 7 层。硬件负载均衡设备可以将核心部分做成芯片，性能和稳定性更好，而且商用产品的可管理性、文档和服务都比较好。唯一的问题就是价格。  
> 软件负载均衡通常是开源软件。自由度较高，但学习成本和管理成本会比较大。  
> 
> 下面，继续说说 Nginx 的负载均衡实现  
> 1、轮询  
> 这种是默认的策略，把每个请求按顺序逐一分配到不同的 server，如果 server 挂掉，能自动剔除。
```
upstream  fengzp.com {   
    server   192.168.99.100:42000; 
    server   192.168.99.100:42001;  
}
```
> 2、最少连接  
> 把请求分配到连接数最少的 server  
```
upstream  fengzp.com {   
    least_conn;
    server   192.168.99.100:42000; 
    server   192.168.99.100:42001;  
}
```
> 3、权重  
> 使用 weight 来指定 server 访问比率，weight 默认是 1。以下配置会是 server2 访问的比例是 server1 的两倍。
```
upstream  fengzp.com {   
    server   192.168.99.100:42000 weight=1; 
    server   192.168.99.100:42001 weight=2;  
}
```
> 4、ip_hash  
> 每个请求会按照访问 ip 的 hash 值分配，这样同一客户端连续的 Web 请求都会被分发到同一 server 进行处理，可以解决 session 的问题。如果 server 挂掉，能自动剔除。
```
upstream  fengzp.com {   
    ip_hash;
    server   192.168.99.100:42000; 
    server   192.168.99.100:42001;  
}
```
> ip_hash 可以和 weight 结合使用。  


- 百万级数据导出的实现
> 导出大量的数据，在 PHP 设置层面主要是 set_time_limit(0) 和 ini_set('memory_limit', '1024M')，但是一个 PHP 程序占用那么大的内存的空间未免太奢侈；  
> excel 表的限制，PHPExcel_Settings::setCacheStorageMethod 方法可以更改缓冲方式来减小内存的使用，但是内存溢出还是不容易避免；  
```
Excel 2003 及以下的版本，一张表最大支持 65536 行数据，256 列
Excel 2007-2010 版本，一张表最大支持 1048576 行，16384 列
```
> csv 文件储存，既不限制数量，还能直接用 EXCEL 来查看，又能以后把文件导入数据库；但是，当我们用 putcsv() 输出缓存 buffer，如果几百万的数据一直用这个函数输出，会导致输出缓存太大而报错的；而且，使用 Excel 查看也查看不了全部  
> 
> 综上所述，数据的输出使用 csv 文件格式，并将数据分割保存在多个 csv 文件中，并且最后压缩成 zip 文件提供下载  
```php
function exportCsv(array $head, $data, $mark = 'attack_ip_info', $fileName = "test.csv")
{
    set_time_limit(0);
    $sqlCount = $data->count();
    // 输出 Excel 文件头
    header('Content-Type: application/vnd.ms-excel;charset=utf-8');
    header('Content-Disposition: attachment;filename="' . $fileName . '"');
    header('Cache-Control: max-age=0');

    //每次只从数据库取100000条以防变量缓存太大
    $sqlLimit = 100000;
    // 每隔$limit行，刷新一下输出buffer，不要太大，也不要太小
    $limit = 100000;
    // buffer计数器
    $cnt = 0;
    $fileNameArr = [];
    // 逐行取出数据，不浪费内存
    for($i = 0; $i < ceil($sqlCount / $sqlLimit); $i++){
    	// 生成临时文件
        $fp = fopen($mark . '_' . $i . '.csv', 'w'); 
        // 修改可执行权限
        // chmod('attack_ip_info_' . $i . '.csv',777);
        $fileNameArr[] = $mark . '_' . $i . '.csv';
        // 将数据通过fputcsv写到文件句柄
        fputcsv($fp, $head);
        $dataArr = $data->offset($i * $sqlLimit)->limit($sqlLimit)->get()->toArray();
        foreach($dataArr as $a){
            $cnt++;
            if ($limit == $cnt){
                // 刷新一下输出buffer，防止由于数据过多造成问题
                ob_flush();
                flush();
                $cnt = 0;
            }
            fputcsv($fp, $a);
        }
        fclose($fp);
    }
    // 进行多个文件压缩
    $zip = new ZipArchive();
    $zipFileName = $mark . ".zip";
    $zip->open($zipFileName, ZipArchive::CREATE); 
    foreach($fileNameArr as $file){
        $zip->addFile($file, basename($file));
    }
    $zip->close();
    foreach($fileNameArr as $file){
        unlink($file);
    }
    // 输出压缩文件提供下载
    header("Cache-Control: max-age=0");
    header("Content-Description: File Transfer");
    header('Content-disposition: attachment; filename=' . basename($zipFileName));
    header("Content-Type: application/zip");
    header("Content-Transfer-Encoding: binary"); 
    header('Content-Length: ' . filesize($zipFileName)); 
    // 输出文件
    @readfile($zipFileName);
    // 删除压缩包临时文件
    unlink($zipFileName);
}
```

- 简述 session 存储原理和实现 session 共享的方式以及防止 session_id 泄漏
> 当用户第一次访问站点时，PHP 会用 session_start () 函数为用户创建一个 session ID，这就是针对这个用户的唯一标识，每一个访问的用户都会得到一个自己独有的 session ID，这个 session ID 会存放在响应头里的 cookie 中，之后发送给客户端。这样客户端就会拥有一个该站点给他的 session ID。  
> 当用户第二次访问该站点时，浏览器会带着本地存放的 cookie (里面存有上次得到的 session ID) 随着请求一起发送到服务器，服务端接到请求后会检测是否有 session ID，如果有就会找到响应的 session 文件，把其中的信息读取出来；如果没有就跟第一次一样再创建个新的。  
> 通常站点的退出功能，实际上就是调用一下 session_destroy () 函数 (也有可能更复杂些)，把该用户的 session 文件删除，再把用户的 cookie 清除。这样客户端和服务端就算没有联系了。  
> 因为 HTTP 是无状态的，所以一次请求完成后客户端和服务端就不再有任何关系了，谁也不认识谁。但由于一些需要（如保持登录状态等），必须让服务端和客户端保持联系，session ID 就成了这种联系的媒介了。  
> 
> 客户端的工作：通过上面的分析我们可以知道 session 实际上是依赖与 cookie 的，当用户访问某一站点时，浏览器会根据用户访问的站点自动搜索可用的 cookie，如果有可用的就随着请求一起发送到了服务端。每次接收到服务端的响应时又会更新本地的 cookie 信息。当然也可以用 GET 方式来传递 session ID，但不推荐用 GET，这样不安全。  
> 服务端的工作：由上面的流程图可以看到，服务端实际上是把产生的一些数据存放在了 session 文件中，该文件的名字就是”sess“加上 session ID，这些文件的存放位置就是 phpinfo () 查到的 session.savepath 值。  
> 
> Session 共享: php 支持把 session 数据存储到 memcache/redis 内存服务器，手动把 seesion 的文件改为 nfs 网络系统文件，从而实现文件的跨机器共享。  
> 
> 防止 session_id 泄露：  
> 1、在登录验证成功后，通过重置 session，使之前的匿名 sessionId 失效，这样可以避免使用伪造的 sessionId 进行攻击。request.getSession ().invalidate (); 这样登录前与登录后的 sessionID 就不会相同。  
> 2、设置 httpOnly 属性。httponly 是微软对 cookie 做的扩展，该值指定 Cookie 是否可通过客户端脚本访问，解决用户的 cookie 可能被盗用的问题，减少跨站脚本攻击。主流的大多数浏览器已经支持此属性。httpOnly 是 cookie 的扩展属性，并不包含在 servlet2.x 的规范里。  
> 

- 如何果网站使用的 utf8 编码，为防止乱码出现，需要注意那些问题？
> 1、数据库中库和表字段中都用 utf8_general_ci 编码    
> 2、php 连接 mysql，指定数据库编码为 utf8 mysql_query (“set namesutf8”);  
> 3、php 文件指定头部编码为 utf-8 header (“content-type:text/html;charset=utf-8”);  
> 4、网站下所有文件的编码为 utf8  
> 5、html 文件指定编码为 utf-8 <content="text/html;charset=utf-8" />  

- Php 框架中单入口功能的实现原理以及必要性？
> 大部分 php 框架都是单一入口模式，进来的所有 http 请求都会指向一个文件，从这个入口文件进去，进去再去访问别的文件进行相应的操作，而这个原理就是 php 与服务器的配合，讲所有请求的路径转换给 index。Php 做字符串请求的判断即可达到路由解析功能，由于所有的 http 请求都由 index。Php 接收，可以集中的安全性检查。  

- 对命名空间的理解
> 一般地，在一个 php 程序源代码文件或同一个请求中是不允许有两个以上相同名称的类名、常量名或者函数名的，这样的话，程序在运行的时候就会报错。但是，在很多实际情况下，我们是无法避免在同一个文件下有两个以上相同名称的类名、常量名或者函数名的。这种情况下就需要引入命名空间。  
> 我们把这个 PHP 文件想象成是一个文件夹，而里面的两个函数想象成是两个文件，但是在同一个文件夹下是不允许有两个相同名称的文件的，所以我们就必须的将他们分到两个不同名称的文件夹中。命名空间的道理就是这样的。  
> 
> 命名空间主要解决下面两个问题：  
> 1、用户编写的代码与 PHP 内部的类 / 函数 / 常量或第三方类 / 函数 / 常量之间的名字冲突  
> 2、为很长的标识符创建一个别名，提高代码的可读性，减少代码的编写量。  

- $_SERVER [‘REMOTE_ADDR’] 和 $_SERVER [‘HTTP_X_FORWARED_FOR’] 的作用和区别
> php 中 $_SERVER 参数 HTTP_X_FORWARDED_FOR 和 REMOTE_ADDR 来获取客户端 IP。  
> 在 PHP 中使用 $_SERVER ["REMOTE_ADDR"] 来取得客户端的 IP 地址，但如果客户端是使用代理服务器来访问，那取到的就是代理服务器的 IP 地址，而不是真正的客户端 IP 地址。要想透过代理服务器取得客户端的真实 IP 地址，就要使用 $_SERVER ["HTTP_X_FORWARDED_FOR"] 来读取。  

### 开发进阶
- PHP 弱类型的实现
> PHP 是弱类型，动态的语言脚本。在申明一个变量的时候，并不需要指明它保存的数据类型。  
> 实际上，在 PHP 中声明的变量，在 ZE 中都是用结构体 zval 来保存的。zval 的实现上，_zvalue_value 是真正保存数据的关键部分。通过共用体实现的弱类型变量声明。  
> _zval_struct.type 中存储着一个变量的真正类型，根据 type 来选择如何获取 zvalue_value 的值。  
> 在 PHP 中，任何不属于 PHP 的内建的变量类型的变量，都会被看作资源来进行保存。 比如：数据库句柄、打开的文件句柄、打开的 socket 句柄。  
> 资源类型，会用 lval，此时它是一个整型指示器， 然后 PHP 会再根据这个指示器在 PHP 内建的一个资源列表中查询相对应的资源。  
> 正是因为 ZE 这样的处理方式，使 PHP 就实现了弱类型，而对于 ZE 的来说，它所面对的永远都是同一种类型 zval。  

- PHP 的生命周期
> ...

- PHP 框架 Laravel 和 Yii 路由原理对比
> ...

### 狗血的面试题
千万级别并发、负载均衡、索引原理

- 单点登录的实现
> 单点登录是在多个应用系统中，用户只需要登录一次就可以访问所有相互信任的应用系统的保护资源，若用户在某个应用系统中进行注销登录，所有的应用系统都不能再直接访问保护资源，像一些知名的大型网站，如：淘宝与天猫、新浪微博与新浪博客等都用到了这个技术。  
> 原理：有一个独立的认证中心，只有认证中心才能接受用户的用户名和密码等信息进行认证，其他系统不提供登录入口，只接受认证中心的间接授权。间接授权通过令牌实现，当用户提供的用户名和密码通过认证中心认证后，认证中心会创建授权令牌，在接下来的跳转过程中，授权令牌作为参数发送给各个子系统，子系统拿到令牌即得到了授权，然后创建局部会话。  

- 10g 的访问日志文件中？查看访问日志中访问次数前十的 IP？
> cat access_log | awk ‘{print $1}’ | uniq -c|sort -rn|head -10


