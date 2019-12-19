
### 什么是 SOAP
SOAP（Simple Object Access Protocol）即简单对象访问协议，定义了数据对象传输的格式，以便在网络的节点之间交换信息。SOAP 是基于 XML 的简易协议，可使应用程序在 HTTP 之上进行信息交换。  
经常和 SOAP 放在一起谈论的 WSDL（Web Service Description Language）即 Web 服务描述语言。WSDL 用于描述一个 Web Service，即 WSDL 用来说明某个 Web 服务该怎样使用，有怎样的接口方法，支持怎样的参数，会有怎样的返回。由于支持 SOAP 的服务端接口是经常使用 WSDL 来描述，因此我们会看到它们总被放在一起讨论，于是在这种情况下，WSDL 常常被形容成 SOAP 服务的使用说明书，但是请注意，本质上它们之间不存在依赖关系。  


**SOAP 与 REST**  
REST，Representational State Transfer，即表现层状态转换，指的是一种为了信息能在互联网上顺利传递而设计的软件架构风格。  
SOAP 是协议，但 REST 是风格，而非协议或标准，至于 HTTP，它是 REST 风格的重要载体。  
SOAP 和 REST，由于概念层次上的不同，其实原本是无法放到一起比较的。  
对于互联网来说，SOAP 已经是一项“古老”的技术了，晚辈 REST 似乎更切合互联网的潮流。在大多数情况下，REST 要易用和流行得多，于是很多人都不喜欢繁琐的 SOAP 协议。SOAP 一般强大而复杂，REST 一般返璞归真。

### 为什么 SOAP
SOAP 提供了一种标准的方法，使得运行在不同的操作系统并使用不同的技术和编程语言的应用程序可以互相进行通信。

### SOAP 语法
一条 SOAP 消息就是一个普通的 XML 文档，包含下列元素：

- 必需的 Envelope 元素，可把此 XML 文档标识为一条 SOAP 消息
- 可选的 Header 元素，包含头部信息
- 必需的 Body 元素，包含所有的调用和响应信息
- 可选的 Fault 元素，提供有关在处理此消息所发生错误的信息

一些重要的语法规则：

- SOAP 消息必须用 XML 来编码
- SOAP 消息必须使用 SOAP Envelope 命名空间
- SOAP 消息必须使用 SOAP Encoding 命名空间
- SOAP 消息不能包含 DTD 引用
- SOAP 消息不能包含 XML 处理指令

**SOAP 消息的基本结构**  
```
<?xml version="1.0"?>
<soap:Envelope
xmlns:soap="http://www.w3.org/2001/12/soap-envelope"
soap:encodingStyle="http://www.w3.org/2001/12/soap-encoding">

<soap:Header>
...
</soap:Header>

<soap:Body>
...
  <soap:Fault>
  ...
  </soap:Fault>
</soap:Body>

</soap:Envelope>
```

### SOAP 实例
SOAP 请求如下
```
POST /InStock HTTP/1.1
Host: www.example.org
Content-Type: application/soap+xml; charset=utf-8
Content-Length: nnn

<?xml version="1.0"?>
<soap:Envelope
xmlns:soap="http://www.w3.org/2001/12/soap-envelope"
soap:encodingStyle="http://www.w3.org/2001/12/soap-encoding">

<soap:Body xmlns:m="http://www.example.org/stock">
  <m:GetStockPrice>
    <m:StockName>IBM</m:StockName>
  </m:GetStockPrice>
</soap:Body>

</soap:Envelope>
```

SOAP 响应如下
```
HTTP/1.1 200 OK
Content-Type: application/soap+xml; charset=utf-8
Content-Length: nnn

<?xml version="1.0"?>
<soap:Envelope
xmlns:soap="http://www.w3.org/2001/12/soap-envelope"
soap:encodingStyle="http://www.w3.org/2001/12/soap-encoding">

<soap:Body xmlns:m="http://www.example.org/stock">
  <m:GetStockPriceResponse>
    <m:Price>34.5</m:Price>
  </m:GetStockPriceResponse>
</soap:Body>

</soap:Envelope>
```
