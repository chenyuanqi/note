
### Nginx 安装
```bash
# 重定向依赖
yum install -y pcre-devel 
yum -y install gcc make gcc-c++ wget
# https 需要的依赖
yum -y install openssl openssl-devel 

# 选择合适的 nginx 版本
wget http://nginx.org/download/nginx-x.x.x.tar.gz
tar zxf nginx-x.x.x.tar.gz
cd nginx-x.x.x
./configure
make && make install
# 一般会安装在 /usr/local/nginx
/usr/local/nginx/sbin/nginx -t
# 设置全局 nginx（~/.bash_profile）PATH=$PATH:$HOME/bin:/usr/local/nginx/sbin/

# 启动
/usr/local/nginx/sbin/nginx

# 重启
/usr/local/nginx/sbin/nginx -s reload

# 关闭进程
/usr/local/nginx/sbin/nginx -s stop

# 平滑关闭 nginx
/usr/local/nginx/sbin/nginx -s quit

# 查看 nginx 版本
/usr/local/nginx/sbin/nginx -V 

# nginx 卸载
yum remove nginx
# 编译安装的删除 /usr/local/nginx 目录（如果配置了自启动脚本，删除脚本相关内容）
```

Nginx 安装参数说明  
| 参数 | 说明 |
| ---- | ---- |
| --prefix=`<path>` | Nginx安装路径。如果没有指定，默认为 /usr/local/nginx。 |
| --sbin-path=`<path>` | Nginx可执行文件安装路径。只能安装时指定，如果没有指定，默认为`<prefix>`/sbin/nginx。 |
| --conf-path=`<path>` | 在没有给定-c选项下默认的nginx.conf的路径。如果没有指定，默认为`<prefix>`/conf/nginx.conf。 |
| --pid-path=`<path>` | 在nginx.conf中没有指定pid指令的情况下，默认的nginx.pid的路径。如果没有指定，默认为 `<prefix>`/logs/nginx.pid。 |
| --lock-path=`<path>` | nginx.lock文件的路径。 |
| --error-log-path=`<path>` | 在nginx.conf中没有指定error_log指令的情况下，默认的错误日志的路径。如果没有指定，默认为 `<prefix>`/- logs/error.log。 |
| --http-log-path=`<path>` | 在nginx.conf中没有指定access_log指令的情况下，默认的访问日志的路径。如果没有指定，默认为 `<prefix>`/- logs/access.log。 |
| --user=`<user>` | 在nginx.conf中没有指定user指令的情况下，默认的nginx使用的用户。如果没有指定，默认为 nobody。 |
| --group=`<group>` | 在nginx.conf中没有指定user指令的情况下，默认的nginx使用的组。如果没有指定，默认为 nobody。 |
| --builddir=DIR | 指定编译的目录 |
| --with-rtsig_module | 启用 rtsig 模块 |
| --with-select_module --without-select_module | 允许或不允许开启SELECT模式，如果 configure 没有找到更合适的模式，比如：kqueue(sun os),epoll (linux kenel 2.6+), rtsig(- 实时信号)或者/dev/poll(一种类似select的模式，底层实现与SELECT基本相 同，都是采用轮训方法) SELECT模式将是默认安装模式|
| --with-poll_module --without-poll_module | Whether or not to enable the poll module. This module is enabled by, default if a more suitable method such as kqueue, epoll, rtsig or /dev/poll is not discovered by configure. |
| --with-http_ssl_module | Enable ngx_http_ssl_module. Enables SSL support and the ability to handle HTTPS requests. Requires OpenSSL. On Debian, this is libssl-dev. 开启HTTP SSL模块，使NGINX可以支持HTTPS请求。这个模块需要已经安装了OPENSSL，在DEBIAN上是libssl  |
| --with-http_realip_module | 启用 ngx_http_realip_module |
| --with-http_addition_module | 启用 ngx_http_addition_module |
| --with-http_sub_module | 启用 ngx_http_sub_module |
| --with-http_dav_module | 启用 ngx_http_dav_module |
| --with-http_flv_module | 启用 ngx_http_flv_module |
| --with-http_stub_status_module | 启用 "server status" 页 |
| --without-http_charset_module | 禁用 ngx_http_charset_module |
| --without-http_gzip_module | 禁用 ngx_http_gzip_module. 如果启用，需要 zlib 。 |
| --without-http_ssi_module | 禁用 ngx_http_ssi_module |
| --without-http_userid_module | 禁用 ngx_http_userid_module |
| --without-http_access_module | 禁用 ngx_http_access_module |
| --without-http_auth_basic_module | 禁用 ngx_http_auth_basic_module |
| --without-http_autoindex_module | 禁用 ngx_http_autoindex_module |
| --without-http_geo_module | 禁用 ngx_http_geo_module |
| --without-http_map_module | 禁用 ngx_http_map_module |
| --without-http_referer_module | 禁用 ngx_http_referer_module |
| --without-http_rewrite_module | 禁用 ngx_http_rewrite_module. 如果启用需要 PCRE 。 |
| --without-http_proxy_module | 禁用 ngx_http_proxy_module |
| --without-http_fastcgi_module | 禁用 ngx_http_fastcgi_module |
| --without-http_memcached_module | 禁用 ngx_http_memcached_module |
| --without-http_limit_zone_module | 禁用 ngx_http_limit_zone_module |
| --without-http_empty_gif_module | 禁用 ngx_http_empty_gif_module |
| --without-http_browser_module | 禁用 ngx_http_browser_module |
| --without-http_upstream_ip_hash_module | 禁用 ngx_http_upstream_ip_hash_module |
| --with-http_perl_module | 启用 ngx_http_perl_module |
| --with-perl_modules_path=PATH | 指定 perl 模块的路径 |
| --with-perl=PATH | 指定 perl 执行文件的路径 |
| --http-log-path=PATH | Set path to the http access log |
| --http-client-body-temp-path=PATH | Set path to the http client request body temporary files |
| --http-proxy-temp-path=PATH | Set path to the http proxy temporary files |
| --http-fastcgi-temp-path=PATH | Set path to the http fastcgi temporary files |
| --without-http | 禁用 HTTP server |
| --with-mail | 启用 IMAP4/POP3/SMTP 代理模块 |
| --with-mail_ssl_module | 启用 ngx_mail_ssl_module |
| --with-cc=PATH | 指定 C 编译器的路径 |
| --with-cpp=PATH | 指定 C 预处理器的路径 |
| --with-cc-opt=OPTIONS | Additional parameters which will be added to the variable CFLAGS. With the use of the system library PCRE in FreeBSD, it is necessary to indicate --with-cc-opt="-I /usr/local/include". If we are using select() and it is necessary to increase the number of file descriptors, then this also can be assigned here: --with-cc-opt="-D FD_SETSIZE=2048". |
| --with-ld-opt=OPTIONS | Additional parameters passed to the linker. With the use of the system library PCRE in - FreeBSD, it is necessary to indicate --with-ld-opt="-L /usr/local/lib". |
| --with-cpu-opt=CPU | 为特定的 CPU 编译，有效的值包括：pentium, pentiumpro, pentium3, pentium4, athlon, opteron, amd64, sparc32, sparc64, ppc64 |
| --without-pcre | 禁止 PCRE 库的使用。同时也会禁止 HTTP rewrite 模块。在 "location" 配置指令中的正则表达式也需要 PCRE 。 |
| --with-pcre=DIR | 指定 PCRE 库的源代码的路径。 |
| --with-pcre-opt=OPTIONS | Set additional options for PCRE building. |
| --with-md5=DIR | Set path to md5 library sources. |
| --with-md5-opt=OPTIONS | Set additional options for md5 building. |
| --with-md5-asm | Use md5 assembler sources. |
| --with-sha1=DIR | Set path to sha1 library sources. |
| --with-sha1-opt=OPTIONS | Set additional options for sha1 building. |
| --with-sha1-asm | Use sha1 assembler sources. |
| --with-zlib=DIR | Set path to zlib library sources. |
| --with-zlib-opt=OPTIONS | Set additional options for zlib building. |
| --with-zlib-asm=CPU | Use zlib assembler sources optimized for specified CPU, valid values are: pentium, pentiumpro |
| --with-openssl=DIR | Set path to OpenSSL library sources |
| --with-openssl-opt=OPTIONS | Set additional options for OpenSSL building |
| --with-debug | 启用调试日志 |
| --add-module=PATH | Add in a third-party module found in directory PATH |  

### Nginx 配置基础
nginx.conf 是主配置文件，由若干个部分组成，每个大括号`{}`表示一个部分。每一行指令都由分号结束 `;`，标志着一行的结束。  

#### 常用正则

| 正则 | 说明 | 正则 | 说明 |
| ---- | ---- | ---- | ---- | 
| `. ` | 匹配除换行符以外的任意字符 | `$ ` | 匹配字符串的结束 |
| `? ` | 重复 0 次或 1 次 | `{n} ` | 重复 n 次 |
| `+ ` | 重复 1 次或更多次 | `{n,} ` | 重复 n 次或更多次 |
| `*` | 重复 0 次或更多次 | `[c] ` | 匹配单个字符 c |
| `\d ` |匹配数字 | `[a-z]` | 匹配 a-z 小写字母的任意一个 |
| `^ ` | 匹配字符串的开始 | - | - |

#### 全局变量

| 变量 | 说明 | 变量 | 说明 |
| ---- | ---- | ---- | ---- | 
| $args | 这个变量等于请求行中的参数，同$query_string | $remote_port | 客户端的端口。 |
| $content_length | 请求头中的Content-length字段。 | $remote_user | 已经经过Auth Basic Module验证的用户名。 |
| $content_type | 请求头中的Content-Type字段。 | $request_filename | 当前请求的文件路径，由root或alias指令与URI请求生成。 |
| $document_root | 当前请求在root指令中指定的值。 | $scheme | HTTP方法（如http，https）。 |
| $host | 请求主机头字段，否则为服务器名称。 | $server_protocol | 请求使用的协议，通常是HTTP/1.0或HTTP/1.1。 |
| $http_user_agent | 客户端agent信息 | $server_addr | 服务器地址，在完成一次系统调用后可以确定这个值。 |
| $http_cookie | 客户端cookie信息 | $server_name | 服务器名称。 |
| $limit_rate | 这个变量可以限制连接速率。 | $server_port | 请求到达服务器的端口号。 |
| $request_method | 客户端请求的动作，通常为GET或POST。 | $request_uri | 包含请求参数的原始URI，不包含主机名，如：/foo/bar.php?arg=baz。 |
| $remote_addr | 客户端的IP地址。 | $uri | 不带请求参数的当前URI，$uri不包含主机名，如/foo/bar.html。 |
| $document_uri | 与$uri相同。 | - | - |

例如请求：`http://localhost:3000/test1/test2/test.php`

$host：localhost  
$server_port：3000  
$request_uri：/test1/test2/test.php  
$document_uri：/test1/test2/test.php  
$document_root：/var/www/html  
$request_filename：/var/www/html/test1/test2/test.php  

### Nginx 第三方模块安装方法
```bash
./configure --prefix=/nginx 安装目录  --add-module=/第三方模块目录
```

### Nginx 监控
使用 ngxtop 实时解析 nginx 访问日志，并且将处理结果输出到终端，功能类似于系统命令 top。  
如果要指定访问日志文件和 / 或日志格式，请使用 -f 指定格式 和 -a 指定正则选项。  
```
# 安装 ngxtop
pip install ngxtop

# 实时状态
ngxtop
# 状态为 404 的前 10 个请求的路径：
ngxtop top request_path --filter 'status == 404'

# 发送总字节数最多的前 10 个请求
ngxtop --order-by 'avg(bytes_sent) * count'

# 排名前十位的 IP，例如，谁攻击你最多
ngxtop --group-by remote_addr

# 打印具有 4xx 或 5xx 状态的请求，以及 status 和 http referer
ngxtop -i 'status >= 400' print request status http_referer

# 响应状态为 200，发送的平均正文字节以 'foo' 开始：
ngxtop avg bytes_sent --filter 'status == 200 and request_path.startswith("foo")'

# 使用 common 日志格式从远程机器分析 apache 访问日志
ssh remote tail -f /var/log/apache2/access.log | ngxtop -f common
```

### 重定向
permanent 永久性重定向，请求日志中的状态码为 301。  
redirect 临时重定向，请求日志中的状态码为 302。  
```
# 重定向整个网站
server {
    server_name old-site.com
    return 301 $scheme://new-site.com$request_uri;
}

# 重定向单页
server {
    location = /oldpage.html {
        return 301 http://example.org/newpage.html;
    }
}

# 重定向整个子路径
location /old-site {
    rewrite ^/old-site/(.*) http://example.org/new-site/$1 permanent;
}
```

### 内容缓存
允许浏览器基本上永久地缓存静态内容，Nginx 将为您设置 Expires 和 Cache-Control 头信息。  
```
location /static {
    root /data;
    expires max; # 不缓存设置为 -1
}
```

### 文件缓存
```
open_file_cache max=1000 inactive=20s;
open_file_cache_valid 30s;
open_file_cache_min_uses 2;
open_file_cache_errors on;
```

### ssl 缓存
```
ssl_session_cache shared:SSL:10m;
ssl_session_timeout 10m;
```

### Gzip 压缩
```
gzip  on;
gzip_buffers 16 8k;
gzip_comp_level 6;
gzip_http_version 1.1;
gzip_min_length 256;
gzip_proxied any;
gzip_vary on;
gzip_types
    text/xml application/xml application/atom+xml application/rss+xml application/xhtml+xml image/svg+xml
    text/javascript application/javascript application/x-javascript
    text/x-json application/json application/x-web-app-manifest+json
    text/css text/plain text/x-component
    font/opentype application/x-font-ttf application/vnd.ms-fontobject
    image/x-icon;
gzip_disable  "msie6";
```
