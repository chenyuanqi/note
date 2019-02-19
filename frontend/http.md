
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

### http header
Accept，指定客户端能够接收的内容类型，例如：Accept: text/plain, text/html  

