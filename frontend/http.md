
### http 状态码
100, 'Continue', 继续  
101, 'Switching Protocols', 交换协议  
102, 'Processing', 处理 RFC2518  
103, 'Early Hints', 提前暗示  

200, 'OK', 成功  
201, 'Created', 已创建（通常这是 PUT 方法得到的响应码）  
202, 'Accepted', 认可的（服务器已接受请求，但尚未处理）  
203, 'Non-Authoritative Information', 非授权信息  
204, 'No Content', 无内容  
205, 'Reset Content', 重置内容  
206, 'Partial Content', 部分内容  
207, 'Multi-Status', 多状态 RFC4918  
208, 'Already Reported', 已经播报 RFC5842  
226, 'IM Used', 异步使用 RFC3229  

300, 'Multiple Choices', 多选  
301, 'Moved Permanently', 永久移除（永久移动位置，被请求的资源已经被永久性的转移了位置）  
302, 'Found', 您请求的资源现在需要临时通过其他的 URI 来获取  
303, 'See Other', 查看其他位置（服务器发送该响应用来引导客户端使用 GET 方法访问另外一个 URI）  
304, 'Not Modified', 未修改（告诉客户端，所请求的内容距离上次访问并没有变化）   
305, 'Use Proxy', 使用代理，被请求的资源必须通过指定的代理才能访问到  
307, 'Temporary Redirect', 临时跳转（被请求的资源在临时从不同的 URL 响应请求）  
308, 'Permanent Redirect', 永久跳转 RFC7238  

400, 'Bad Request', 无效请求  
401, 'Unauthorized', 未授权请求  
402, 'Payment Required', 请求不允许  
403, 'Forbidden', 请求被禁止  
404, 'Not Found', 请求对象不存在  
405, 'Method Not Allowed', 方法不允许  
406, 'Not Acceptable', 无法接受（在进行服务器驱动内容协商后，没有发现合适的内容传回给客户端）  
407, 'Proxy Authentication Required', 需要代理身份验证  
408, 'Request Timeout', 请求超时  
409, 'Conflict', 冲突（由于和被请求的资源的当前状态之间存在冲突，请求无法完成）  
410, 'Gone', 请求丢失  
411, 'Length Required',   
412, 'Precondition Failed', 先决条件失败  
413, 'Payload Too Large', 负载太大  
414, 'URI Too Long', url 太长  
415, 'Unsupported Media Type', 媒体类型不支持  
416, 'Range Not Satisfiable', 范围不适合  
417, 'Expectation Failed', 预期失败  
418, 'I\'m a teapot', RFC2324  
421, 'Misdirected Request', 误导请求 RFC7540  
422, 'Unprocessable Entity', 无法处理的实体 RFC4918（请求格式正确，但是由于含有语义错误，无法响应）  
423, 'Locked', RFC4918  
424, 'Failed Dependency', RFC4918  
425, 'Reserved for WebDAV advanced collections expired proposal', 保留用于WebDAV高级集合过期提案 RFC2817  
426, 'Upgrade Required', RFC2817  
428, 'Precondition Required', 要求先决条件 RFC6585  
429, 'Too Many Requests', RFC6585  
431, 'Request Header Fields Too Large', 请求头字段太大 RFC6585  
451, 'Unavailable For Legal Reasons', 因法律原因无法获得 RFC7725  
499, 'Client has closed connection', 客户端主动关闭连接  

500, 'Internal Server Error', 内部服务器错误（服务器遇到了一个未曾预料的状况，无法完成对请求的处理，会在程序码出错时出现）  
501, 'Not Implemented', 未实现（服务器不支持当前请求所需要的某个功能）  
502, 'Bad Gateway', 无效网关（作为网关或者代理工作的服务器尝试执行请求时，从上游服务器接收到无效的响应）  
503, 'Service Unavailable', 服务不可用（由于临时的服务器维护或者过载，服务器当前无法处理请求）  
504, 'Gateway Timeout', 网关连接超时  
505, 'HTTP Version Not Supported', http 版本不支持  
506, 'Variant Also Negotiates', 服务器存在内部配置错误 RFC2295  
507, 'Insufficient Storage', 远程服务器返回错误 RFC4918  
508, 'Loop Detected', RFC5842  
510, 'Not Extended', RFC2774  
511, 'Network Authentication Required', 客户端需要进行身份验证以获得网络访问  


- 划重点
> 301、302 的区别  
> 301，302 对用户来说没有区别，他们看到效果只是一个跳转，浏览器中旧的 URL 变成了新的 URL。302 转向可能会有 URL 规范化及网址劫持的问题。可能被搜索引擎判为可疑转向，甚至认为是作弊。当网页 A 用 301 重定向转到网页 B 时，搜索引擎可以肯定网页 A 永久的改变位置，或者说实际上不存在了，搜索引擎就会把网页 B 当作唯一有效目标。  
> 
> 500、502、504 的处理
> 一般情况下，服务器发生 500 错误是比较好追查的，一般在日志里面都有错误栈打出，据此容易定位和修正代码错误。而 502、504 的发生，服务器基本都没有错误日志可供追查，因为通常并不是代码执行层面发生了错误。  
> 502： Bad Gateway；504: Gateway Timeout；  
> 从字面意思看来，都是因为 Gateway 发生了问题；那，什么是 Gateway 呢？对于 PHP 这种脚本语言，服务器的脚本执行程序（FastCGI 服务）发生错误或者出现超时就会报 502 或者 504。  
> 典型的 502 错误，如果 PHP-CGI 被卡住，所有进程都正在处理；那么，当新的请求进来时，PHP-CGI 将无法处理新的请求就会导致 502 出现；  
> 而 504 的典型错误，PHP 执行被阻塞住，导致 nginx 等服务迟迟收不到返回的数据，此时就会报 504 的错误。  
> 一般来说，502 与 php-fpm.conf 的设置有关，也与 php 的执行程序性能有关，网站的访问量大，而 php-cgi 的进程数偏少。针对这种情况的 502 错误，只需增加 php-cgi 的进程数。具体就是修改/usr/local/php/etc/php-fpm.conf文件，将其中的max_children值适当增加。  
> 504 请求没有到可以执行的 php-fpm，与 nginx.conf 的配置也有关系。  

### http 协议
HTTP 是一个客户端终端（用户）和服务器端（网站）请求和应答的标准（TCP）。通过使用 Web 浏览器、网络爬虫或者其它的工具，客户端发起一个 HTTP 请求到服务器上指定端口（默认端口为 80）。我们称这个客户端为用户代理程序（user agent）。应答的服务器上存储着一些资源，比如 HTML 文件和图像。我们称这个应答服务器为源服务器（origin server）。在用户代理和源服务器中间可能存在多个“中间层”，比如代理服务器、网关或者隧道（tunnel）。  

尽管 TCP/IP 协议是互联网上最流行的应用，但是在 HTTP 协议中，并没有规定必须使用它或它支持的层。事实上，HTTP 可以在任何互联网协议上，或其他网络上实现。HTTP 假定其下层协议提供可靠的传输。因此，任何能够提供这种保证的协议都可以被其使用。因此也就是其在 TCP/IP 协议族使用 TCP 作为其传输层。  

通常，由 HTTP 客户端发起一个请求，创建一个到服务器指定端口（默认是 80 端口）的 TCP 连接。HTTP 服务器则在那个端口监听客户端的请求。一旦收到请求，服务器会向客户端返回一个状态，比如"HTTP/1.1 200 OK"，以及返回的内容，如请求的文件、错误消息、或者其它信息。

### http header
Accept，指定客户端能够接收的内容类型，例如：Accept: text/plain, text/html  
Accept-Charset：浏览器申明自己接收的字符集  
Accept-Encoding：浏览器申明自己接收的编码方法，通常指定压缩方法，是否支持压缩，支持什么压缩方法 （gzip，deflate）  
Accept-Language：浏览器申明自己接收的语言语言跟字符集的区别：中文是语言，中文有多种字符集，比如 big5，gb2312，gbk 等等  
Accept-Ranges：WEB 服务器表明自己是否接受获取其某个实体的一部分（比如文件的一部分）的请求。bytes：表示接受，none：表示不接受  

Age：当代理服务器用自己缓存的实体去响应请求时，用该头部表明该实体从产生到现在经过多长时间了  
Authorization：当客户端接收到来自 WEB 服务器的 WWW-Authenticate 响应时，该头部来回应自己的身份验证信息给 WEB 服务器  
Content-Language：WEB 服务器告诉浏览器自己响应的对象的语言  
Content-Encoding：WEB 服务器表明自己使用了什么压缩方法（gzip，deflate）压缩响应中的对象   
Content-Length： WEB 服务器告诉浏览器自己响应的对象的长度  
Content-Range： WEB 服务器表明该响应包含的部分对象为整个对象的哪个部分  
Content-Type： WEB 服务器告诉浏览器自己响应的对象的类型  
Expired：WEB 服务器表明该实体将在什么时候过期，对于过期了的对象，只有在跟 WEB 服务器验证了其有效性后，才能用来响应客户请求。是 HTTP/1.0 的头部  
Host：客户端指定自己想访问的 WEB 服务器的域名 /IP 地址和端口号  
Last-Modified：WEB 服务器认为对象的最后修改时间，比如文件的最后修改时间，动态页面的最后产生时间等等  
Referer：浏览器向 WEB 服务器表明自己是从哪个网页 /URL 获得 / 点击当前请求中的网址 /URL  
Server: WEB 服务器表明自己是什么软件及版本等信息  
User-Agent: 浏览器表明自己的身份（是哪种浏览器）  

Cache-Control：  
请求：  
no-cache（不要缓存的实体，要求现在从 WEB 服务器去取）  
max-age：（只接受 Age 值小于 max-age 值，并且没有过期的对象）  
max-stale：（可以接受过去的对象，但是过期时间必须小于 max-stale 值）  
min-fresh：（接受其新鲜生命期大于其当前 Age 跟 min-fresh 值之和的缓存对象）  
响应：  
public(可以用 Cached 内容回应任何用户)  
private（只能用缓存内容回应先前请求该内容的那个用户）  
no-cache（可以缓存，但是只有在跟 WEB 服务器验证了其有效后，才能返回给客户端）  
max-age：（本响应包含的对象的过期时间）  
ALL:  no-store（不允许缓存）  

Connection：  
请求：  
close（告诉 WEB 服务器或者代理服务器，在完成本次请求的响应后，断开连接，不要等待本次连接的后续请求了）     
keepalive（告诉 WEB 服务器或者代理服务器，在完成本次请求的响应后，保持连接，等待本次连接的后续请求）  
响应：  
close（连接已经关闭）  
keepalive（连接保持着，在等待本次连接的后续请求）  
Keep-Alive：如果浏览器请求保持连接，则该头部表明希望 WEB 服务器保持连接多长时间（秒）。例如：Keep-Alive：300  


### HTTP 三层架构
```
                Request
浏览器（browser）——————> Internet ——————> Web Server
                <------          <------
                                 Response
```
### HTTP 两个核心
1、请求（Request）req  
2、响应（Response）res  

> 请求 URL（uniform resource locator）统一资源定位器，它可以用来标识一个资源，指明了如何定位这个资源（DNS、IP）  
> 请求信息包括请求行、空行、其他消息体  
> 
> 当浏览器接收并显示网页前，此网页所在的服务器会返回一个包含 HTTP 状态码的信息头（server header）用以响应浏览器的请求
> HTTP 状态码[参考](http://www.restapitutorial.com/httpstatuscodes.html)  
>
> HTTP 动词包括 HEAD、GET、POST、PUT、PATCH、DELETE  
>
> GET 和 POST 有各自的语意，通常 GET 用于获取资源，而 POST 用于修改资源  
> GET 和 POST 最直观的区别在于 GET 把参数包含在 URL 中，POST 通过 request body 传递参数（相对安全）；
> 实际上，GET 和 POST 本质上就是 TCP 链接，GET 产生一个 TCP 数据，而 POST 产生两个 TCP 数据包（Firefox 只发一次）  

### HTTP 一个记住
> HTTP 协议是无状态协议  
> 会话控制：服务器的实现 session 与客户端的支持 cookie  

### HTTPS 
在 HTTP 协议上增加了 SSL(secure socket layer)层（即 SSL 层 - 应用层 - 传输层 - 网络层 - 链路层 - 实体层）
```
# HTTPS 认证流程
                             发起请求
                      --------------------------->　　server 
                              下发证书
                      <---------------------------    server 
                      证书数字签名(用证书机构公钥加密)
                      --------------------------->　　证书机构 
                          证书数字签名验证通过
client(内置证书机构证书) <---------------------------  证书机构
                      公钥加密随机密码串(未来的共享秘钥)
                     --------------------------->　　 server私钥解密(非对称加密)
                        SSL协议结束　HTTP协议开始
                      <---------------------------    server(对称加密)
                            共享秘钥加密 HTTP
                     --------------------------->　 　server(对称加密)
```
> 核对证书： 证书机构的公开秘钥验证证书的数字签名  
> 使用非对称加密，公钥加密建立连接  
> 共享密钥加密  

### Websocket
Websocket 是一种持久化协议，基于 TCP，跟 HTTP 协议同级，属于全双工通信，由客户端建立连接（简而言之，客户端发送 GET 请求，告诉服务端要进行协议升级，服务器给客户端转换协议，进而建立 WebSocket 连接）。  
为了建立一个 WebSocket 连接，客户端浏览器首先要向服务器发起一个 HTTP 请求，这个请求和通常的 HTTP 请求不同，包含了一些附加头信息，其中附加头信息”Upgrade: WebSocket”表明这是一个申请协议升级的 HTTP 请求，服务器端解析这些附加的头信息然后产生应答信息返回给客户端，客户端和服务器端的 WebSocket 连接就建立起来了，双方就可以通过这个连接通道自由的传递信息，并且这个连接会持续存在直到客户端或者服务器端的某一方主动的关闭连接。  

> 长链接是指让客户端浏览器与服务器端保持长久的连接，并能持续通讯，它可以反向 ajax（服务器推技术）  
> 长链接技术有轮循（模拟型的实时），长轮循，SSE 等  

### HTTP2
由于 HTTP 1.1 包含了太多细节和可选内容，也为未来的扩展预留了很多选项，使得它过于庞大，而且它也很难完全使用出 TCP 协议所能提供的所有强大能力；另外，加载网站所需的资源，带来的恼人延迟 ... HTTP2 应运而生。  
> 克服延迟之道有很多，比如雪碧图（Spriting）、小图内联（Inlining）、JS 文件拼接（Concatenation）和分片（Sharding）  

改善 HTTP 协议显然是极好的事情  
> 降低协议对延迟的敏感  
> 修复 pipelining 和 head of line blocking 的问题  
> 防止主机需求更高的连接数量  
> 保留所有现有的接口，内容，URI 格式和结构  
> 由 IETF 的 HTTPbis 工作组来制定协议  

HTTP2 都带来哪些细节和概念呢？  
- 全双工通信  
> 服务端和客户端终端通信双方设备既是发送器，也是接收器  

- 二进制传输  
> HTTP2 是一个二进制协议，使成帧的使用变得更为便捷  
> HTTP2 会发送不同类型的二进制帧（公共字段：Type, Length, Flags, Stream Identifier 和 frame payload）  

- 多路复用  
> 多个请求共享一个 TCP 连接  

- 优先级和依赖性  
> 借助于 PRIORITY 帧，客户端同样可以告知服务器当前的流依赖于其他哪个流  
> 优先级和依赖关系可以在传输过程中被动态的改变  

- 头部压缩  
> 保证了 HTTP 的可重复性，提升请求速度  

- 重置 - 后悔药  
> HTTP2 可以通过发送 RST_STREAM 帧来实现重置当前传输的消息，从而避免浪费带宽和中断已有的连接  

- 服务器推送  
> 也叫“缓存推送”，客户端请求资源前，服务端主动将资源推送给客户端以备将来之需  
> 客户端可以通过发送一个 RST_STREAM 帧来中止服务器推送  

- 流量控制  
> 数据帧会受到流量控制，流的两端都必须告诉对方自己还有足够的空间来处理新的数据  


