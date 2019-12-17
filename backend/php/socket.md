
### Socket 是什么
TCP (Transmission Control Protocol) 即传输控制协议 / 网间协议，是一个工业标准的协议集，它是为广域网（WANs）设计的。  
UDP（User Data Protocol，用户数据报协议）是与 TCP 相对应的协议。它是属于 TCP/IP 协议族中的一种。TCP 和 UDP 都处于运输层。  
Socket，是封装好的通信协议的接口，提供网络通讯的能力，更加方便使用协议栈。简而言之，Socket 封装好了一套 API，比如，create、listen、connect、accept、send、read 和 write 等等，更方便使用 TCP（Socket 是应用层与 TCP/IP 协议族通信的中间软件抽象层）。  

Socket 的操作是 I/O 的集合。  
我们知道，网络 I/O 延迟给应用带来极大的负面影响。而 I/O 的过程设计有两个对象，一个是 IO 调用者（进程 process 或者线程 thread），另一个是系统内核 kernel。  
以 read 操作为例子，read 的过程经历了两个步骤：  
> 1、等待数据就绪  
> 2、将数据从内核拷贝到进程中  

I/O 模型有几种类型：阻塞 I/O（bloking I/O）、非阻塞 I/O（non-bloking I/O）、多路复用 I/O（multiplexing I/O）、信号驱动式 I/O（signal-driven I/O）、异步 I/O （asynchronous I/O）。  
> 阻塞 I/O，对于网络 I/O，用户进程通过网络传输等待数据的到达。等待过程中，进程被阻塞。等到数据从网络接收完毕，内核开始复制数据到内存。直到内核返回结果，用户进程才解除阻塞。  
> 非阻塞 I/O, 不断询问系统数据是否准备好。例如 socket accept 操作，调用 accept，立刻返回结果。通过返回的结果来判断数据是否准备好，如果还没准备好，继续再问系统，直到数据准备好。进程并没有阻塞，但是一直占用 CPU，所以这个不断询问的操作做了很多无用功，浪费资源。  
> 多路复用 I/O, 通过某种机制让系统通知进程其所等待的数据已经准备好。多路复用解决了非阻塞浪费 CPU 资源的问题。多路复用有三个著名的库 select、poll 和 epoll。比如 select 是系统级别的函数，PHP 里 socket_select 就是调用系统的 select()。select 不断轮询文件的描述符的读写就绪状态，如果发现就绪，就通知进程处理数据。另外，select 不支持超过 1024 个的描述符，所以超过 1024 个连接，select 会处理不来。  

Socket 主要应用于异步编程，比如网页实时聊天、消息推送、代理、数据转发、游戏服务器等。  
workerman 的 socket 就是用 PHP socket 扩展，只是对其进行工程化开发，成了一个框架。另外，swoole 是用 C 实现，它的 socket 是 C 库的 socket，更加底层可控。  

**Socket 基本原理**  
一个生活中的场景。  
你要打电话给一个朋友，先拨号，朋友听到电话铃声后提起电话，这时你和你的朋友就建立起了连接，就可以讲话了。等交流结束，挂断电话结束此次交谈。  
![socket 基本原理](./image/socket-basic.jpg)  

服务器端先初始化 Socket，然后与端口绑定 (bind)，对端口进行监听 (listen)，调用 accept 阻塞，等待客户端连接。在这时如果有个客户端初始化一个 Socket，然后连接服务器 (connect)，如果连接成功，这时客户端与服务器端的连接就建立了。客户端发送数据请求，服务器端接收请求并处理请求，然后把回应数据发送给客户端，客户端读取数据，最后关闭连接，一次交互结束。  

**Socket 相关函数**  
Socket 通信依次会进行 Socket 创建、Socket 监听、Socket 收发、Socket 关闭几个阶段。  
`注意：测试之前需要把 php.ini 中 extension=php_sockets.dll 打开`

```php
socket_accept(); // 接收一个socket连接
socket_bind(); // 把socket绑定在一个IP地址和端口上
socket_clear_error(); // 清除socket的错误或者最后的错误代码
socket_close(); // 关闭一个socket资源
socket_connect(); // 开始一个socket连接
socket_create_listen(); // 在指定端口打开一个socket监听
socket_create_pair(); // 产生一对没有区别的socket到一个数组里
socket_create(); // 产生一个socket，相当于产生一个socket的数据结构
socket_get_option(); // 获取socket选项
socket_getpeername(); // 获取远程类似主机的ip地址
socket_getsockname(); // 获取本地socket的ip地址
socket_iovec_add(); // 添加一个新的向量到一个分散/聚合的数组
socket_iovec_alloc(); // 这个函数创建一个能够发送接收读写的iovec数据结构
socket_iovec_delete(); // 删除一个已经分配的iovec
socket_iovec_fetch(); // 返回指定的iovec资源的数据
socket_iovec_free(); // 释放一个iovec资源
socket_iovec_set(); // 设置iovec的数据新值
socket_last_error(); // 获取当前socket的最后错误代码
socket_listen(); // 监听由指定socket的所有连接
socket_read(); // 读取指定长度的数据
socket_readv(); // 读取从分散/聚合数组过来的数据
socket_recv(); // 从socket里结束数据到缓存
socket_recvfrom(); // 接受数据从指定的socket，如果没有指定则默认当前socket
socket_recvmsg(); // 从iovec里接受消息
socket_select(); // 多路选择
socket_send(); // 这个函数发送数据到已连接的socket
socket_sendmsg(); // 发送消息到socket
socket_sendto(); // 发送消息到指定地址的socket
socket_set_block(); // 在socket里设置为块模式
socket_set_nonblock(); // socket里设置为非块模式
socket_set_option(); // 设置socket选项
socket_shutdown(); // 这个函数允许你关闭读、写、或者指定的socket
socket_strerror(); // 返回指定错误号的详细错误
socket_write(); // 写数据到socket缓存
socket_writev(); // 写数据到分散/聚合数组
```

**Socket 案例**  
server.php (服务器端) 代码如下：  
```php
// 设置不限请求超时时间
set_time_limit(0);

$ip = '127.0.0.1';
$port = 8099;

// 创建 socket
if(($sock = socket_create(AF_INET,SOCK_STREAM,SOL_TCP)) < 0) {
    echo "socket_create() 失败的原因是:".socket_strerror($sock)."\n";
    exit();
}
// 把 socket 绑定在一个IP地址和端口上
if(($ret = socket_bind($sock,$ip,$port)) < 0) {
    echo "socket_bind() 失败的原因是:".socket_strerror($ret)."\n";
    exit();
}
// 监听由指定 socket 的所有连接
if(($ret = socket_listen($sock,4)) < 0) {
    echo "socket_listen() 失败的原因是:".socket_strerror($ret)."\n";
    exit();
}

// 次数
$count = 0;

do{
    // 接收一个 Socket 连接
    if (($msgsock = socket_accept($sock)) < 0) {
        echo "socket_accept() failed: reason: " . socket_strerror($msgsock) . "\n";
        break;
    } else {
        // 发送到客户端
        $msg = "测试成功! \n";
        socket_write($msgsock, $msg, strlen($msg));

        echo "测试成功了啊\n";
        // 获得客户端的输入
        $buf = socket_read($msgsock, 2048);

        $talkback = "收到的信息:$buf\n";
        echo $talkback;

        // 第 5 次结束
        if(++$count >= 5){
            break;
        }
    }
    // 关闭 socket
    socket_close($msgsock);
}while(true);
```

client.php (客户端) 代码如下：  
```php
error_reporting(E_ALL);
// 设置不限请求超时时间
set_time_limit(0);

echo "<h2>TCP/IP Connection</h2>\n";

$ip = '127.0.0.1';
$port = 8099;

// 创建 socket
if(($socket = socket_create(AF_INET,SOCK_STREAM,SOL_TCP)) < 0) {
    echo "socket_create() 失败的原因是:".socket_strerror($socket)."\n";
    exit();
}
echo "OK. \n";

echo "试图连接 '$ip' 端口 '$port'...\n";

// 连接 socket
if(($result = socket_connect($socket, $ip, $port)) < 0){
    echo "socket_connect() 失败的原因是:".socket_strerror($sock)."\n";
    exit();
}
echo "连接OK\n";

$in .= "hello flycorn\r\n";
$out = '';

// 写数据到 socket 缓存
if(!socket_write($socket, $in, strlen($in))) {
    echo "socket_write() 失败的原因是:".socket_strerror($sock)."\n";
    exit();
}
echo "发送到服务器信息成功！\n";
echo "发送的内容为:$in \n";

// 读取指定长度的数据
while($out = socket_read($socket, 2048)) {
    echo "接收服务器回传信息成功！\n";
    echo "接收的内容为:",$out;
}

echo "关闭 SOCKET...\n";
socket_close($socket);
echo "关闭 OK\n";
```

下面，我们用命令行之行如下命令
```bash
# 打开终端
php server.php

# 新开一个终端
php client.php
```
