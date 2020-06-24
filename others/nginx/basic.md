
### 为什么是 Nginx
1、高并发、高性能（这是其他 web 服务器不具有的）  
2、可扩展性好（模块化设计，第三方插件生态圈丰富）  
3、高可靠性（可以在服务器行持续不间断的运行数年）  
4、热部署（这个功能对于 Nginx 来说特别重要，热部署指可以在不停止 Nginx 服务的情况下升级 Nginx）  
5、BSD 许可证（意味着我们可以将源代码下载下来进行修改然后使用自己的版本）  

[Nginx 官网](https://nginx.org/)  
[Ngins Github](https://github.com/nginx/nginx)  
[Nginx 文档](https://docshome.gitbooks.io/nginx-docs/content/)  
[Nginx Cookbook](https://huliuqing.gitbooks.io/complete-nginx-cookbook-zh/content/)  

### 主要组成
Nginx 二进制可执行文件：由各模块源码编译出一个文件  
Nginx.conf 配置文件：控制 Nginx 行为  
acess.log 访问日志： 记录每一条 HTTP 请求信息  
error.log 错误日志: 定位问题  

### 重要概念
- 正向代理 Vs 反向代理
> 正向代理：某些情况下，代理我们用户去访问服务器，需要用户手动的设置代理服务器的 ip 和端口号。正向代理比较常见的一个例子就是 VPN 了。  
> 反向代理：它是用来代理服务器的，代理我们要访问的目标服务器。代理服务器接受请求，然后将请求转发给内部网络的服务器，并将从服务器上得到的结果返回给客户端，此时代理服务器对外就表现为一个服务器。  
> 
> 所以，简单的理解，就是正向代理是为客户端做代理，代替客户端去访问服务器，而反向代理是为服务器做代理，代替服务器接受客户端请求。  

- 负载均衡
> 在高并发情况下需要使用，其原理就是将并发请求分摊到多个服务器执行，减轻每台服务器的压力，多台服务器 (集群) 共同完成工作任务，从而提高了数据的吞吐量。  
> Nginx 支持的 weight 轮询（默认）、ip_hash、fair、url_hash 这四种负载均衡调度算法。  
> 负载均衡相比于反向代理更侧重的时将请求分担到多台服务器上去，所以谈论负载均衡只有在提供某服务的服务器大于两台时才有意义。  

- 动静分离
> 动静分离是让动态网站里的动态网页根据一定规则把不变的资源和经常变的资源区分开来，动静资源做好了拆分以后，我们就可以根据静态资源的特点将其做缓存操作，这就是网站静态化处理的核心思路。

**nginx 常用命令**  
```bash
nginx # 启动 nginx
nginx -s reload # 向主进程发送信号，重新加载配置文件，热重启
nginx -s reopen # 重启 Nginx
nginx -s stop # 快速关闭
nginx -s quit # 等待工作进程处理完成后关闭
nginx -t # 查看当前 Nginx 配置是否有错误
nginx -t -c <配置路径> # 检查配置是否有问题，如果已经在配置目录，则不需要 - c

```

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
```
# nginx -t -c /etc/nginx/nginx.conf 配置语法检查
# nginx -s reload -c /etc/nginx/nginx.conf 不重启的方式加载配置
# rpm -ql nginx 查看安装到了哪些目录  

# Nginx使用用logrotate服务对日志进行切割的配置文件（eg：按天切割）
/etc/logrotate.d/nginx
# Nginx的核心目录
/etc/nginx
# 主要配置文件，Nginx启动的时候会读取
/etc/nginx/nginx.conf
/etc/nginx/conf.d
# nginx.conf没变更久读default.conf（默认Server加载的文件）
/etc/nginx/conf.d/default.conf
# Nginx对Python的wsgi配置
/etc/nginx/uwsgi_params
# fastcgi配置
/etc/nginx/fastcgi_params
# scgi配置
/etc/nginx/scgi_params
# Nginx缓存目录
/var/cache/nginx
# Nginx日志目录
/var/log/nginx
# Nginx默认网站存放的路径
/usr/share/nginx/html
/usr/share/nginx/html/50x.html
/usr/share/nginx/html/index.html
# 设置http的Content-Type与扩展名对应关系的配置文件
/etc/nginx/mime.types
# Nginx模块所在目录
/usr/lib64/nginx/modules
/etc/nginx/modules
# 二进制执行文件
/usr/sbin/nginx
/usr/sbin/nginx-debug
# 编码转换的映射文件
/etc/nginx/koi-utf
/etc/nginx/koi-win
/etc/nginx/win-utf
# 配置CentOS守护进程对Nginx的管理方式
/usr/lib/systemd/system/nginx-debug.service
/usr/lib/systemd/system/nginx.service
/etc/sysconfig/nginx
/etc/sysconfig/nginx-debug
# Nginx的文档
/usr/share/doc/nginx-1.16.0
/usr/share/doc/nginx-1.16.0/COPYRIGHT
/usr/share/man/man8/nginx.8.gz
# Nginx检测更新命令
/usr/libexec/initscripts/legacy-actions/nginx
/usr/libexec/initscripts/legacy-actions/nginx/check-reload
/usr/libexec/initscripts/legacy-actions/nginx/upgrade
```

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
比如 Nginx 默认是不支持 Lua 的，所以需要自己编译安装下
```bash
# 先安装 lua 库
yum install lua lua-devel-y

wget http://luajit.org/download/LuaJIT-2.0.5.tar.gz
tar -xzvf LuaJIT-2.0.5.tar.gz
cd LuaJIT-2.0.5
# 编译安装
make install PREFIX=/usr/local/LuaJIT
# 导入环境变量
export LUAJIT_LIB=/usr/local/LuaJIT/lib
export LUAJIT_INC=/usr/local/LuaJIT/include/luajit-2.0
# 安装 lua 模块到 nginx
./configure --prefix=/etc/nginx  --add-module=./ngx_devel_kit-0.3.1
./configure --prefix=/etc/nginx  --add-module=./lua-nginx-module-0.10.15

# 加载 lua 库到 ld.so.conf 文件
echo"/usr/local/LuaJIT/lib">>/etc/ld.so.conf
# 让动态函式库加载到缓存中
ldconfig

# 测试 lua
vi /etc/nginx/nginx.conf
# location /hello {
#        default_type 'text/plain';
#        content_by_lua 'ngx.say("welcome，lua comming~")';
# }

```

### Nginx 配置 waf 防火墙
waf 防火墙主要功能：  
> 拦截 Cookie 类型工具  
> 拦截异常 post 请求  
> 拦截 CC 洪水攻击  
> 拦截 URL  
> 拦截 arg（提交的参数）  

```bash
git clone https://github.com/loveshell/ngx_lua_waf
cd ngx_lua_waf
ls -alF 
# args 里面的规则 get 参数进行过滤的
# url 是只在 get 请求 url 过滤的规则
# post 是只在 post 请求过滤的规则
# whitelist 是白名单，里面的 url 匹配到不做过滤
# user-agent 是对 user-agent 的过滤规则
mkdir -p /etc/nginx/waf
mv * /etc/nginx/waf

# nginx http 模块添加如下配置
# server_tokens off; # 隐藏 nginx 版本信息
# lua_package_path "/etc/nginx/waf/?.lua";
# lua_shared_dict limit 10m;
# init_by_lua_file /etc/nginx/waf/init.lua;
# access_by_lua_file /etc/nginx/waf/waf.lua;

vim /etc/nginx/waf/config.lua
# RulePath 设置规则所在目录 /etc/nginx/waf/wafconf/
# attacklog 开启日志记录 on
# logdir 日志存放地址（不存在需要手动创建一下） /var/log/nginx/hack
# UrlDeny url防护
# CookieMatch 匹配 cookie
# postMatch 匹配 post
# whiteModule 开启白名单
# black_fileExt 不允许上传的后缀
# ipWhitelist ip白名单列表 {"xxx", "xxx"}
# ipBlocklist ip黑名单列表 {"xxx", "xxx"}
# CCDeny cc防护
# CCrate="100/60" 访问频率 60s 内访问 100 次
# html=[[...]] 拦截后显示的 html，可自定义
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
