
### 什么是 swoole
swoole 是基于 C 开发的一个 php 扩展，类似你熟悉的 Mysqli、cURL 等等。但是 swoole 更强大，它几乎重新定义了 php 的高度，让 php 变得更加无可挑剔。  
Swoole 是 PHP 语言的高性能网络通信框架，提供了 PHP 语言的异步多线程服务器，异步 TCP/UDP 网络客户端，异步 MySQL，数据库连接池，AsyncTask，消息队列，毫秒定时器，异步文件读写，异步 DNS 查询。 Swoole 虽然是标准的 PHP 扩展，实际上与普通的扩展不同。普通的扩展只是提供一个库函数。而 swoole 扩展在运行后会接管 PHP 的控制权，进入事件循环。当 IO 事件发生后，swoole 会自动回调指定的 PHP 函数。

### swoole 解决了哪些问题
swoole 其实更多的是解决 php 在某些方面的缺陷，比如即时通讯、异步任务、消息队列等等。

### php cli 模式
cli 模式下，主要使用如下命令：php -m(module),php -c(config),php -i (info),php -f(file),php -v(version)，php -r(run)
```bash
# 运行某 php 文件
php a.php
php -f a.php

# 查看命令行下该 php 的版本信息
php -v

# 查看 cli 模式下 phpinfo 的信息
php -i

# 查看 php 所加载的配置文件所在路径
php -i | grep php.ini

# 指定命令行模式下 php 所加载的配置文件
php -c /custom/directory/custom-file.ini

# 校验某 php 文件是否有语法错误
php -l a.php

# 查看 php 都加载了哪些模块
php -m

# 命令行下直接运行 php 代码
php -r "var_dump(get_cfg_var('display_errors'));"
```

### 进程和线程
对于操作系统而言，进程就是一个任务，比方说你打开了一个记事本，那就启动了一个进程，打开了两个浏览器，就是另外开启了两个进程，再或者说我现在在 word 内写文章，打开 word 也会占用一个进程。也就是说，一个进程至少要干一件事情。  
```bash
# 查看 nginx 进程
ps aux | grep nginx
```
有些情况下，一个进程会同时做一些事情，比如说 word。它可以同时进行打字、拼写检查等操作。注意这里我们说的同时进行。像这样，在一个进程内部，同时运行着多个 “子任务”，我们就可以把这些子任务称之为 “线程”。即进程是由多个线程组成的，一个进程至少要有一个线程。实际上，线程是操作系统最小的执行单元。  
apache 其实就是一种多进程实现的案例。当父进程监听到有新的请求时，就会 fork 出新的子进程来对之进行处理。Linux 的 fork() 函数通过系统调用即可实现创建一个与原进程几乎相同的进程。对于多任务，通常我们会设计 Master-Worker 模式，即一个 Master 进程负责分配任务，多个 Worker 进程负责执行任务。同理，如果是多线程，Master 就是主线程，Worker 就是子线程。  

多进程的优点就是稳定性很高，如果一个进程挂了，不会影响其他子进程，当然，如果主进程挂了那就都玩完（主进程挂点的可能性微乎其微）。而对于多线程，这个恐怕就是致命的缺点了，因为所有线程共享内存，如果某一个线程挂了，那这个进程几乎就崩溃了。  
性能方面，不论是进程还是线程，如果启动太多，无疑都会带来 CPU 的调度问题，因为进程或者线程的切换，本身就非常耗费资源。数量达到一定程度的时候，CPU 和内存就消耗殆尽，电脑就死机了。  
举个例子：使用过 windows 的用户都知道，如果我们打开的软件越多（开启的进程也就越多），电脑就会越卡，甚至装死机没反应。  
线程与进程相比，自然是要比进程更轻量一些，而且线程之间是共享内存的，所以不同线程之间的交互就显得容易实现。而对于多进程之间的通信，需要借助消息队列，共享内存等复杂的方式才可以实现。  

### IO 模型
IO 即 Input/Output, 输入和输出的意思。在计算机的世界里，涉及到数据交换的地方，比如磁盘、网络等，就需要 IO 接口。  

通常，IO 是相对的。  
php 本身是单线程的，当 php 进程被挂起的时候，像上面的读取磁盘数据，往磁盘写数据，在 IO 操作之前 php 代码就没办法继续执行了。用户可以同时访问你的网站实际上是 web 服务器的功劳。    

同步 IO 模型下，主线程只能被挂起等待，但是在异步 IO 模型中，主线程发起 IO 指令后，可以继续执行其他指令，没有被挂起，也没有切换线程的操作。因此，使用异步 IO 明显可以提高了系统性能。  

### TCP/IP 和 UDP
- 打开一个浏览器，然后输入网址后回车，展现了一个网页的内容的过程是怎样的呢？
> 浏览器通过 TCP/IP 协议建立到服务器的 TCP 连接  
> 客户端向服务器发送 HTTP 协议请求包，请求服务器里的资源文档  
> 服务器向客户端发送 HTTP 协议应答包，如果请求的资源包含有动态语言的内容，那么服务器会调用动态语言的解释引擎负责处理 “动态内容”，并将处理得到的数据返回给客户端  
> 客户端与服务器断开，由客户端解释 HTML 文档，在客户端屏幕上渲染图形结果

计算机为了联网，就必须遵循通信协议。早期的互联网有很多协议，但是最重要的就非 TCP 协议和 IP 协议莫属了。所以，我们把互联网的协议简称为 TCP/IP 协议。  
TCP 协议是一种面向连接、可靠的、基于 IP 之上的传出层协议。TCP 协议是建立在 IP 协议之上，专门负责建立可靠连接，并保证数据包顺序到达。TCP 协议会通过握手建立连接，然后，对每个 IP 包编号，确保对方顺序收到，如果出现丢包，则重新发送。  
TCP 协议还有一个更重要的特点，它是基于数据流的。这就好比客户端和服务端要进行数据交互，中间有一个管子连接着，这个时候交互数据就好比管子中的水，当数据在传输（水在流动）的过程中，服务端是无法知道哪段数据是我们想要的完整数据。  
相对于 TCP, 使用 UDP 协议进行通信的最大区别就是，UDP 不需要建立连接，给他一个 ip 和端口，就可以直接发送数据包了。但是，能不能成功到达就不知道了。虽然 UDP 传输不可靠，但是速度快。对于一些对数据要求不高的场景，使用 UDP 通信无疑是不错的选择。  

### swoole 的安装与升级
```bash
apt-get update
apt-get install make autoconf gcc

# 编译安装 swoole
wget http://pecl.php.net/get/swoole-1.9.6.tgz
tar zxvf swoole-1.9.6.tgz
cd swoole-1.9.6
/usr/local/opt/php70/bin/phpize
./configure
make 
sudo make install
# pecl 安装 swoole
pecl install swoole
# 在 php 配置文件 php.ini 添加 extension=swoole.so

# 查看 swoole 版本
php --ri swoole | grep Version

# 编译升级，重新安装一边相应的版本即可
# pecl 升级 swoole
pecl upgrade swoole
```

### 初识 swoole
swoole 是既支持异步，也支持同步模式。  

- socket 是什么？
> socket 即套接字，是用来与另一个进程进行跨网络通信的文件（linux 中一切都可以理解为 “文件”）  
> 比如客户端可以借助 socket 与服务器之间建立连接  
> 也可以把 socket 理解为一组函数库，它确实也就是一堆函数  

基于套接字接口的网络应用的描述，大致是这样的：服务器创建一个 socket，绑定 ip 和端口，在该端口处进行监听，然后通过 accept 函数阻塞。当有新的客户端连接进来时，server 接收客户端数据并处理数据，然后返回给客户端，客户端关闭连接，server 关闭该客户端，一次连接交互完成。  

server，顾名思义，就是服务器。下面简单使用下 swoole server
```php
// 创建一个 server 对象
$serv = new swoole_server('127.0.0.1', 9501);
// 开启两个 worker 进程，官方建议设置为 CPU 核数的 1-4 倍
$serv->set([
    'worker_num' => 2,
]);

// 有新的客户端连接时，worker进程内会触发该回调
$serv->on('Connect', function ($serv, $fd) {
    echo "new client connected." . PHP_EOL;
});
// server接收到客户端的数据后，worker进程内触发该回调
$serv->on('Receive', function ($serv, $fd, $fromId, $data) {
    // 收到数据后发送给客户端
    $serv->send($fd, 'Server '. $data);
});
// 客户端断开连接或者server主动关闭连接时 worker进程内调用
$serv->on('Close', function ($serv, $fd) {
    echo "Client close." . PHP_EOL;
});

// 启动server
$serv->start();
```
执行 `php server.php` 即可启动如上 swoole server 了。

### swoole task
AsyncTask，即异步任务。我们可以利用 AsyncTask 将一个耗时的任务投递到队列中，由进程池异步去执行。  
```php
$serv = new swoole_server("127.0.0.1", 9501);
// 配置 task 进程的数量
$serv->set([
    'task_worker_num' => 1,
]);

$serv->on('Connect', function ($serv, $fd) {
    echo "new client connected." . PHP_EOL;
});
$serv->on('Receive', function ($serv, $fd, $fromId, $data) {
    echo "worker received data: {$data}" . PHP_EOL;
    
    // 投递一个任务到task进程中
    $serv->task($data);

    // 通知客户端server收到数据了
    $serv->send($fd, 'This is a message from server.');
    
    // 为了校验task是否是异步的，这里和task进程内都输出内容，看看谁先输出
    echo "worker continue run."  . PHP_EOL;
});
/**
 * 注册 onTask 回调
 *
 * $serv swoole_server
 * $taskId 投递的任务id,因为task进程是由worker进程发起，所以多worker多task下，该值可能会相同
 * $fromId 来自那个worker进程的id
 * $data 要投递的任务数据
 */
$serv->on('Task', function ($serv, $taskId, $fromId, $data) {
    echo "task start. --- from worker id: {$fromId}." . PHP_EOL;
    for ($i=0; $i < 5; $i++) { 
        sleep(1);
        echo "task runing. --- {$i}" . PHP_EOL;
    }
    echo "task end." . PHP_EOL;
});
/**
 * 注册 onFinish 回调
 * 只有在task进程中调用了finish方法或者return了结果，才会触发finish
 */
$serv->on('Finish', function ($serv, $taskId, $data) {
    echo "finish received data '{$data}'" . PHP_EOL;
});

$serv->start();
```
在 worker 进程收到数据后，直接调用 swoole_server->task 函数把数据投递给 task 进程，随后在 swoole_server->task 调用后和 task 进程内都输出内容。  

### swoole 进程模型
swoole 的进程模型可以用 Master-Manager-Worker 来形容，即在 Master-Worker 的基础上又增加了一层 Manager 进程。  
Master 进程就是我们所说的主进程，掌管生杀大权，它挂了，那底下的都得玩完。Master 进程，包括主线程，多个 Reactor 线程等。Master 进程是一个多线程程序。  
每一个线程都有自己的用途，比如主线程用于 Accept、信号处理等操作，而 Reactor 线程是处理 tcp 连接，处理网络 IO，收发数据的线程。主线程的 Accept 操作，socket 服务端经常用 accept 阻塞；信号处理，信号就相当于一条消息，比如我们经常操作的 Ctrl+C 其实就是给 Master 进程的主线程发送一个 SIGINT 的信号，意思就是你可以终止啦。  
通常，主线程处理完新的连接后，会将这个连接分配给固定的 Reactor 线程，并且这个 Reactor 线程会一直负责监听此 socket（socket 即套接字，是用来与另一个进程进行跨网络通信的文件，文件可读可写），换句话就是当此 socket 可读时，会读取数据，并将该请求分配给 worker 进程，这也就解释了 worker 进程内的回调函数 onReceive 的第三个参数 $fromId 的含义；当此 socket 可写时，会把数据发送给 tcp 客户端。  

在 linux 中，父进程可以通过调用 fork 函数创建一个新的子进程，子进程是父进程的一个副本，几乎但不完全相同，二者的最大区别就是都拥有自己独立的进程 ID，即 PID。  
对于多线程的 Master 进程而言，想要多个 Worker 进程就必须 fork 操作，但是 fork 操作是不安全的，所以，在 swoole 中，有一个专职的 Manager 进程，Manager 进程就专门负责 worker/task 进程的 fork 操作和管理。换句话说，对于 worker 进程的创建、回收等操作全权有 “保姆” Manager 进程来进行管理。  
通常，worker 进程被误杀或者由于程序的原因会异常退出，Manager 进程为了保证服务的稳定性，会重新拉起新的 worker 进程，意思就是 Worker 进程你发生意外 “死” 了，没关系，我自身不“死”，就可以 fork 千千万万个你。  
```
# 6 个主要的回调函数
Master进程：
    启动：onStart
    关闭：onShutdown
Manager进程：
    启动：onManagerStart
    关闭：onManagerStop
Worker进程：
    启动：onWorkerStart
    关闭：onWorkerStop
```
```php
$serv = new swoole_server('127.0.0.1', 9501);
$serv->set([
    'worker_num' => 2,
    'task_worker_num' => 1,
]);
$serv->on('Connect', function ($serv, $fd) {
});
$serv->on('Receive', function ($serv, $fd, $fromId, $data) {
});
$serv->on('Close', function ($serv, $fd) {
});
$serv->on('Task', function ($serv, $taskId, $fromId, $data) {
});
$serv->on('Finish', function ($serv, $taskId, $data) {
});
// swoole_set_process_name 修改进程名，mac 下不支持
$serv->on("start", function ($serv){
    swoole_set_process_name('server-process: master');
});
// 以下回调发生在Manager进程
$serv->on('ManagerStart', function ($serv){
    swoole_set_process_name('server-process: manager');
});
$serv->on('WorkerStart', function ($serv, $workerId){
    // $serv->setting 可以获取配置的 server 信息，在 swoole 中预留了一些 swoole_server 的属性，我们可以在回调函数中访问
    // 如$serv->connections 属性获取当前 server 的所有的连接，$serv->master_pid 属性获取当前 server 的主进程 id
    if($workerId >= $serv->setting['worker_num']) {
        swoole_set_process_name("server-process: task");
    } else {
        swoole_set_process_name("server-process: worker");
    }
});
$serv->start();
```
在 onWorkerStart 回调中，$workerId 表示的是一个值，这个值的范围是 0~worker_num，worker_num 是我们的对 worker 进程的配置，其中 0~worker_num 表示 worker 进程的标识，包括 0 但不包括 worker_num；worker_num~worker_num+task_worker_num 是 task 进程的标识, 包括 worker_num 不包括 worker_num+task_worker_num。  
