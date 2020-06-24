
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

### 屏蔽 IP
在 nginx 的配置文件 nginx.conf 中加入如下配置，也可以放到 http, server, location, limit_except 语句块，需要注意相对路径
```
include ip-black-list.conf;

# ip-black-list.conf 内容如下
deny 165.91.122.67;

deny IP;   # 屏蔽单个 ip 访问
allow IP;  # 允许单个 ip 访问
deny all;  # 屏蔽所有 ip 访问
allow all; # 允许所有 ip 访问
deny 123.0.0.0/8   # 屏蔽整个段即从 123.0.0.1 到 123.255.255.254 访问的命令
deny 124.45.0.0/16 # 屏蔽 IP 段即从 123.45.0.1 到 123.45.255.254 访问的命令
deny 123.45.6.0/24 # 屏蔽 IP 段即从 123.45.6.1 到 123.45.6.254 访问的命令

# 除了几个 IP 外，其他全部拒绝
allow 1.1.1.1; 
allow 1.1.1.2;
deny all; 
```

### 屏蔽 .git 等文件
```
location ~ (.git|.gitattributes|.gitignore|.svn) {
    deny all;
}
```

### 防盗链
盗链是指服务提供商自己不提供服务的内容，通过技术手段绕过其它有利益的最终用户界面 (如广告)，直接在自己的网站上向最终用户提供其它服务提供商的服务内容，骗取最终用户的浏览和点击率。受益者不提供资源或提供很少的资源，而真正的服务提供商却得不到任何的收益。  
网站盗链会大量消耗被盗链网站的带宽，而真正的点击率也许会很小，严重损害了被盗链网站的利益。  

防止盗链的方法主要有如下几种：  
1、不定期更名文件或者目录  
2、限制引用页  
原理是，服务器获取用户提交信息的网站地址，然后和真正的服务端的地址相比较， 如果一致则表明是站内提交，或者为自己信任的站点提交，否则视为盗链。实现时可以使用 HTTP_REFERER 和 htaccess 文件 (需要启用 mod_Rewrite)，结合正则表达式去匹配用户的每一个访问请求。  
```
location ~* \.(gif|jpg|png|swf|flv)$ {
   root html
   valid_referers none blocked *.nginxcn.com;
   if ($invalid_referer) {
     rewrite ^/ www.nginx.cn
     #return 404;
   }
}
```

3、文件伪装  
文件伪装是目前用得最多的一种反盗链技术，一般会结合服务器端动态脚本 (PHP/JSP/ASP)。实际上用户请求的文件地址，只是一个经过伪装的脚本文件，这个脚本文件会对用户的请求作认证，一般会检查 Session，Cookie 或 HTTP_REFERER 作为判断是否为盗链的依据。而真实的文件实际隐藏在用户不能够访问的地方，只有用户通过验证以后才会返回给用户。  
4、加密认证  
这种反盗链方式，先从客户端获取用户信息，然后根据这个信息和用户请求的文件名字一起加密成字符串 (Session ID) 作为身份验证。只有当认证成功以后，服务端才会把用户需要的文件传送给客户。一般我们会把加密的 Session ID 作为 URL 参数的一部分传递给服务器，由于这个 Session ID 和用户的信息挂钩，所以别人就算是盗取了链接，该 Session ID 也无法通过身份认证，从而达到反盗链的目的。这种方式对于分布式盗链非常有效。  
5、随机附加码  
每次在页面里生成一个附加码，并存在数据库里，和对应的图片相关，访问图片时和此附加码对比，相同则输出图片，否则输出 404 图片。  
6、加入水印  

### 防盗图
```
location ~ \/public\/(css|js|img)\/.*\.(js|css|gif|jpg|jpeg|png|bmp|swf) {
    valid_referers none blocked *.jslite.io;
    if ($invalid_referer) {
        rewrite ^/  http://wangchujiang.com/piratesp.png;
    }
}
```

### 反爬虫
根据 User-Agent 过滤请求，通过一个简单的正则表达式，就可以过滤不符合要求的爬虫请求 (初级爬虫)。
```
location / {
    # ~* 表示不区分大小写的正则匹配
    if ($http_user_agent ~* "python|curl|java|wget|httpclient|okhttp") {
        return 503;
    }
    # 正常处理
    # ...
}
```

### 反向代理
反向代理是一个 Web 服务器，它接受客户端的连接请求，然后将请求转发给上游服务器，并将从服务器得到的结果返回给连接的客户端。  
```
# 简单的配置
server {  
  listen       80;                                                        
  server_name  localhost;                                              
  client_max_body_size 1024M;  # 允许客户端请求的最大单文件字节数

  location / {
    proxy_pass                         http://localhost:8080;
    proxy_set_header Host              $host:$server_port;
    proxy_set_header X-Forwarded-For   $remote_addr; # HTTP的请求端真实的IP
    proxy_set_header X-Forwarded-Proto $scheme;      # 为了正确地识别实际用户发出的协议是 http 还是 https
  }

  location /node/ {
    proxy_pass http://127.0.0.1:9502;
  }
}

# 复杂一些的配置
server {
    #侦听的80端口
    listen       80;
    server_name  git.example.cn;
    location / {
        proxy_pass   http://localhost:3000;
        #以下是一些反向代理的配置可删除
        proxy_redirect             off;
        #后端的Web服务器可以通过X-Forwarded-For获取用户真实IP
        proxy_set_header           Host $host;
        client_max_body_size       10m; #允许客户端请求的最大单文件字节数
        client_body_buffer_size    128k; #缓冲区代理缓冲用户端请求的最大字节数
        proxy_connect_timeout      300; #nginx跟后端服务器连接超时时间(代理连接超时)
        proxy_send_timeout         300; #后端服务器数据回传时间(代理发送超时)
        proxy_read_timeout         300; #连接成功后，后端服务器响应时间(代理接收超时)
        proxy_buffer_size          4k; #设置代理服务器（nginx）保存用户头信息的缓冲区大小
        proxy_buffers              4 32k; #proxy_buffers缓冲区，网页平均在32k以下的话，这样设置
        proxy_busy_buffers_size    64k; #高负荷下缓冲大小（proxy_buffers*2）
    }
}
```

代理到上游服务器的配置中，最重要的是 proxy_pass 指令。  
代理模块中的一些常用指令：  
| 指令 | 说明 |
| ---- | ---- |
| proxy_connect_timeout  | Nginx从接受请求至连接到上游服务器的最长等待时间 |
| proxy_send_timeout  | 后端服务器数据回传时间(代理发送超时) |
| proxy_read_timeout  | 连接成功后，后端服务器响应时间(代理接收超时) |
| proxy_cookie_domain | 替代从上游服务器来的Set-Cookie头的domain属性 |
| proxy_cookie_path   | 替代从上游服务器来的Set-Cookie头的path属性 |
| proxy_buffer_size   | 设置代理服务器（nginx）保存用户头信息的缓冲区大小 |
| proxy_buffers       | proxy_buffers缓冲区，网页平均在多少k以下 |
| proxy_set_header    | 重写发送到上游服务器头的内容，也可以通过将某个头部的值设置为空字符串，而不发送某个头部的方法实现 |
| proxy_ignore_headers | 这个指令禁止处理来自代理服务器的应答。 | 
| proxy_intercept_errors | 使nginx阻止HTTP应答代码为400或者更高的应答。 | 

### 负载均衡
upstream 指令启用一个新的配置区段，在该区段定义一组上游服务器。这些服务器可能被设置不同的权重，也可能出于对服务器进行维护，标记为 down。  
```
upstream gitlab {
    ip_hash;
    # upstream的负载均衡，weight 是权重，可以根据机器配置定义权重
    # weigth 参数表示权值，权值越高被分配到的几率越大
    server 192.168.122.11:8081 ;
    server 127.0.0.1:82 weight=3;
    server 127.0.0.1:83 weight=3 down;
    server 127.0.0.1:84 weight=3; max_fails=3  fail_timeout=20s;
    server 127.0.0.1:85 weight=4;;
    keepalive 32;
}

server {
    #侦听的80端口
    listen       80;
    server_name  git.example.cn;
    location / {
        proxy_pass   http://gitlab;    #在这里设置一个代理，和upstream的名字一样
        #以下是一些反向代理的配置可删除
        proxy_redirect             off;
        #后端的Web服务器可以通过X-Forwarded-For获取用户真实IP
        proxy_set_header           Host $host;
        proxy_set_header           X-Real-IP $remote_addr;
        proxy_set_header           X-Forwarded-For $proxy_add_x_forwarded_for;
        client_max_body_size       10m;  #允许客户端请求的最大单文件字节数
        client_body_buffer_size    128k; #缓冲区代理缓冲用户端请求的最大字节数
        proxy_connect_timeout      300;  #nginx跟后端服务器连接超时时间(代理连接超时)
        proxy_send_timeout         300;  #后端服务器数据回传时间(代理发送超时)
        proxy_read_timeout         300;  #连接成功后，后端服务器响应时间(代理接收超时)
        proxy_buffer_size          4k; #设置代理服务器（nginx）保存用户头信息的缓冲区大小
        proxy_buffers              4 32k;# 缓冲区，网页平均在32k以下的话，这样设置
        proxy_busy_buffers_size    64k; #高负荷下缓冲大小（proxy_buffers*2）
        proxy_temp_file_write_size 64k; #设定缓存文件夹大小，大于这个值，将从upstream服务器传
    }
}
```
每个请求按时间顺序逐一分配到不同的后端服务器，如果后端服务器 down 掉，能自动剔除。  

upstream 模块能够使用 3 种负载均衡算法：轮询、IP 哈希、最少连接数。  
RR 轮询： 默认情况下使用轮询算法，不需要配置指令来激活它，它是基于在队列中谁是下一个的原理确保访问均匀地分布到每个上游服务器；
IP 哈希： 通过 ip_hash 指令来激活，Nginx 通过 IPv4 地址的前 3 个字节或者整个 IPv6 地址作为哈希键来实现，同一个 IP 地址总是能被映射到同一个上游服务器；
最少连接数： 通过 least_conn 指令来激活，该算法通过选择一个活跃数最少的上游服务器进行连接。如果上游服务器处理能力不同，可以通过给 server 配置 weight 权重来说明，该算法将考虑到不同服务器的加权最少连接数。  

Nginx 默认是 RR 策略，所以不需要其他更多的设置。  
```
# RR 核心配置
upstream test {
    server localhost:8080;
    server localhost:8081;
}

server {
    listen       81;
    server_name  localhost;
    client_max_body_size 1024M;
 
    location / {
        proxy_pass http://test;
        proxy_set_header Host $host:$server_port;
    }
}
```

Nginx 权重策略，指定轮询几率，weight 和访问比率成正比，用于后端服务器性能不均的情况。
```
# 设置 10 次访问中一般只会有 1 次会访问到 8081，而有 9 次会访问到 8080
upstream test {
    server localhost:8080 weight=9;
    server localhost:8081 weight=1;
}
```

RR 和权重都存在一个问题：下一个请求来的时候请求可能分发到另外一个服务器，导致比如登陆信息等不一致。  
iphash 的每个请求按访问 ip 的 hash 结果分配，这样每个访客固定访问一个后端服务器，可以解决如上问题。  
```
upstream test {
    ip_hash;
    server localhost:8080;
    server localhost:8081;
}
```

fair 是第三方模块，按后端服务器的响应时间来分配请求，响应时间短的优先分配。  
```
upstream backend {
    fair;
    server localhost:8080;
    server localhost:8081;
}
```

url_hash 也是第三方模块，按访问 url 的 hash 结果来分配请求，使每个 url 定向到同一个后端服务器，后端服务器为缓存时比较有效。 在 upstream 中加入 hash 语句，server 语句中不能写入 weight 等其他的参数，hash_method 是使用的 hash 算法
```
upstream backend {
    hash $request_uri;
    hash_method crc32;
    server localhost:8080;
    server localhost:8081;
}
```

**server指令可选参数：**

1. weight：设置一个服务器的访问权重，数值越高，收到的请求也越多；
2. fail_timeout：在这个指定的时间内服务器必须提供响应，如果在这个时间内没有收到响应，那么服务器将会被标记为 down 状态；
3. max_fails：设置在 fail_timeout 时间之内尝试对一个服务器连接的最大次数，如果超过这个次数，那么服务器将会被标记为 down;
4. down：标记一个服务器不再接受任何请求；
5. backup：一旦其他服务器宕机，那么有该标记的机器将会接收请求。

keepalive 指令：Nginx 服务器将会为每一个 worker 进行保持同上游服务器的连接
```
upstream backend {
    server 127.0.0.1:8080;
    keepalive 32;
}

server {
    ...
    location /api/ {
        proxy_pass http://backend;
        proxy_http_version 1.1;
        proxy_set_header Connection "";
    }
}
```
