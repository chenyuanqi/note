
### 顶级配置
```
#定义 Nginx 运行的用户和用户组
user nginx;

#进程文件
pid /var/run/nginx.pid;

#错误日志位置和级别，debug、info、notice、warn、error、crit
error_log  /var/log/nginx/error.log warn;

#Nginx worker 的进程数，一般可设置为可用的CPU内核数。
worker_processes 8;

#每个 worker 打开文件描述符的最大数量限制。理论值应该是最多打开文件数（系统的值ulimit -n）与 nginx 进程数相除，但是 nginx 分配请求并不均匀，所以建议与ulimit -n的值保持一致。
worker_rlimit_nofile 65535;
```

修改系统文件打开数量限制：
```bash
sudo sh -c ulimit -HSn 65535 //临时修改
```

在文件尾部添加：
```
* soft nofile 200000
* hard nofile 200000
```

### Events 模块
```
events {
    #设置一个worker进程同时打开的最大连接数
    worker_connections 2048;

    #告诉nginx收到一个新连接通知后接受尽可能多的连接
    multi_accept on;

    #设置用于复用客户端线程的轮询方法。如果你使用Linux 2.6+，你应该使用epoll。如果你使用*BSD，你应该使用kqueue。
    use epoll;
}
```

### HTTP 模块
```
http {
    #隐藏 Nginx 的版本号，提高安全性。
    server_tokens off;

    #开启高效文件传输模式，sendfile 指令指定 Nginx 是否调用sendfile 函数来输出文件，对于普通应用设为 on，如果用来进行下载等应用磁盘 IO 重负载应用，可设置为 off，以平衡磁盘与网络 I/O 处理速度，降低系统的负载。
    sendfile on;

    #是否开启目录列表访问，默认关闭。
    autoindex off;

    #告诉 Nginx 在一个数据包里发送所有头文件，而不一个接一个的发送
    tcp_nopush on;

    #告诉 Nginx 不要缓存数据，而是一段一段的发送--当需要及时发送数据时，就应该给应用设置这个属性，这样发送一小块数据信息时就不能立即得到返回值。Nginx 默认会始终工作在 tcp nopush 状态下。但是当开启前面的 sendfile on; 时，它的工作特点是 nopush 的最后一个包会自动转转换到 nopush off。为了减小那200ms的延迟，开启 nodelay on; 将其很快传送出去。结论就是 sendfile on; 开启时，tcp_nopush 和 tcp_nodelay 都是on 是可以的。
    tcp_nodelay on;

    #日志格式设定
    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
    '$status $body_bytes_sent "$http_referer" '
    '"$http_user_agent" "$http_x_forwarded_for"';
    #定义访问日志，设置为 off 可以关闭日志，提高性能
    access_log /var/log/nginx/access.log main;


    #连接超时时间，单位是秒
    keepalive_timeout 120;

    #读取HTTP头部的超时时间，默认值 60。客户端与服务器建立连接后将开始接收HTTP头部，在这个过程中，如果在一个时间间隔（超时时间）内没有读取到客户端发来的字节，则认为超时，并向客户端返回408 ("Request timed out")响应。
    client_header_timeout 60;

    #默认值 60。与client_header_timeout相似，只是这个超时时间只在读取HTTP包体时才有效。
    client_body_timeout 10;

    #发送响应的超时时间，默认值 60。即Nginx服务器向客户端发送了数据包，但客户端一直没有去接收这个数据包。如果某个连接超过send_timeout定义的超时时间，那么Nginx将会关闭这个连接。
    send_timeout 60;

    #连接超时后将通过向客户端发送RST包来直接重置连接。这个选项打开后，Nginx会在某个连接超时后，不是使用正常情形下的四次握手关闭TCP连接，而是直接向用户发送RST重置包，不再等待用户的应答，直接释放Nginx服务器上关于这个套接字使用的所有缓存（如TCP滑动窗口）。相比正常的关闭方式，它使得服务器避免产生许多处于FIN_WAIT_1、FIN_WAIT_2、TIME_WAIT状态的TCP连接。注意，使用RST重置包关闭连接会带来一些问题，默认情况下不会开启。
    reset_timedout_connection off;

    #要限制连接，必须先有一个容器对连接进行计数，"zone=" 是给它一个名字，可以随便叫，这个名字要跟下面的 limit_conn 一致。$binary_remote_addr 用二进制来储存客户端的地址，1m 可以储存 32000 个并发会话。
    limit_conn_zone $binary_remote_addr zone=addr:5m;

    #给定的key设置最大连接数。这里key是addr，我们设置的值是100，也就是说我们允许每一个IP地址最多同时打开有100个连接。
    limit_conn addr 100;

    #对每个连接限速100k。这如果一个IP允许两个并发连接，那么这个IP就是限速200K。
    limit_rate 100k; 

    #include 是一个在当前文件中包含另一个文件内容的指令。这里我们使用它来加载文件扩展名与文件类型映射表。nginx根据映射关系，设置http请求响应头的Content-Type值。当在映射表找不到时，使用nginx.conf中default-type指定的默认值。
    include /etc/nginx/mime.types;

    #设置文件使用的默认的MIME-type
    default_type text/html;

    #默认编码
    charset UTF-8;

    #该模块可以读取预先压缩的gz文件，这样可以减少每次请求进行gzip压缩的CPU资源消耗。该模块启用后，nginx首先检查是否存在请求静态文件的gz结尾的文件，如果有则直接返回该gz文件内容。
    gzip_static off;  

    #开启 gzip 压缩。
    gzip on;

    # 禁用客户端为 IE6 时的 gzip功能。
    gzip_disable "msie6";

    #Nginx做为反向代理的时候启用。可选值：off|expired|no-cache|no-sotre|private|no_last_modified|no_etag|auth|any
    gzip_proxied any;

    #设置允许压缩的页面最小字节数，页面字节数从header头中的Content-Length中进行获取。建议设置成大于1k的字节数，小于1k可能会越压越大。
    gzip_min_length 1024;

    #设置数据的压缩等级。这个等级可以是1-9之间的任意数值，9是最慢但是压缩比最大的。
    gzip_comp_level 5;

    #设置系统获取几个单位的缓存用于存储gzip的压缩结果数据流。 例如 4 4k 代表以4k为单位，按照原始数据大小以4k为单位的4倍申请内存。如果没有设置，默认值是申请跟原始数据相同大小的内存空间去存储gzip压缩结果。
    gzip_buffers 4 16k;

    #设置需要压缩的数据格式。Nginx默认只对text/html进行压缩。
    gzip_types text/plain text/css application/json application/x-javascript text/xml application/xml application/xml+rss text/javascript;

    #为打开文件指定缓存，默认是没有启用的，max 指定缓存数量，建议和打开文件数一致，inactive 是指经过多长时间文件没被请求后删除缓存。
    open_file_cache max=65535 inactive=30s;

    #多长时间检查一次缓存的有效信息
    open_file_cache_valid 30s;

    #open_file_cache指令中的inactive参数时间内文件的最少使用次数，如果超过这个数字，文件描述符一直是在缓存中打开的。出现 Last-Modified 不变的情况，就是因为当nginx对一个静态文件缓存后，如果30s内还在访问它，那么它的缓存就一直存在，直到30s内你不访问了为止。
    open_file_cache_min_uses 2;
    #是否记录cache错误
    open_file_cache_errors on;

    include /etc/nginx/conf.d/*.conf;
    include /etc/nginx/sites-enabled/*;
}
```

### SERVER 模块
```
server {
    #监听端口，nginx 会根据请求的 HOST 来决定使用哪个 SERVER 段的配置。如果没有匹配的 server_name，则默认使用配置文件中第一个。加上 default_server 则可以以指定没有匹配时的默认规则。
    #listen 80;
    listen 80 default_server;

    #域名可以有多个，用空格隔开
    server_name www.test.com test.com;
    root /user/share/nginx/html/test;

    #404页面配置
    error_page   404   /404.html;

    #配置 ssl，有需要时开启。
    ssl on;
    ssl_certificate /etc/nginx/ssl/server.crt;
    ssl_certificate_key /etc/nginx/ssl/server.key;

    location / {
        index   index.html index.php;
    }

    #图片缓存时间设置
    location ~ .*.(gif|jpg|jpeg|png|bmp|swf)$ {
        expires 10d;
    }

    #JS和CSS缓存时间设置
    location ~ .*.(js|css)?$ {
        expires 1h;
    }

    location ~ [^/]\.php(/|$) {
        fastcgi_index   index.php;
        #开启 PATH_INFO 支持，作用就是把参数按照给定的正则表达式分割成 $fastcgi_script_name 和 $fastcgi_path_info。
        #例如：请求 index.php/id/1 不加此行配置时，fastcgi_script_name 是 /index.php/id/1，fastcgi_path_info 是空。
        #加上之后，fastcgi_script_name 是 index.php，fastcgi_path_info 是 /id/1
        fastcgi_split_path_info ^(.+\.php)(.*)$;

        #此值即是 PHP 中 $_SERVER['SCRIPT_FILENAME'] 的值
        fastcgi_param   SCRIPT_FILENAME $document_root$fastcgi_script_name;
        fastcgi_param   PATH_INFO               $fastcgi_path_info;
        fastcgi_param   PATH_TRANSLATED $document_root$fastcgi_path_info;

        #指定FastCGI服务器监听端口与地址。须和 PHP-FPM 的设置相同。
        #fastcgi_pass   127.0.0.1:9000;
        fastcgi_pass    unix:/var/run/php5-fpm.sock;
        include fastcgi_params;
    }
}
```