
### 什么是跨域
跨域是浏览器行为，是浏览器安全方面的考虑，是浏览器基于同源策略做出的限制。  

不允许跨域访问并非是浏览器限制了发起跨站请求，而是跨站请求可以正常发起，但是返回结果被浏览器拦截了。最好的例子是 CSRF 跨站攻击原理，请求是发送到了后端服务器，无论是否设置允许跨域。  
有些浏览器不允许从 HTTPS 跨域访问 HTTP，比如 Chrome 和 Firefox，这些浏览器在请求还未发出的时候就会拦截请求，这是特例。  

URL  
URL 全称是统一资源定位符，是对可以从互联网上得到的资源的位置和访问方法的一种简洁的表示，是互联网上标准资源的地址。互联网上的每个文件都有一个唯一的 URL，它包含的信息指出文件的位置以及浏览器应该怎么处理它。  
URL 的基本组成：模式（或称协议）、服务器名称（或 IP 地址，这其中包含端口）、路径和文件名。  

同源策略 (SOP)  
简单说来，要求协议（常见协议如 HTTP、HTTPS、FTP、TCP 等），域名，端口都相同才可以，否则均认为需要做跨域的处理。  

### 跨域的原因
造成跨域的两种策略：DOM 同源策略和 XHR 同源策略。  

DOM 同源策略  
浏览器会禁止对不同源页面 DOM 进行操作，主要针对场景是 iframe 跨域的情况，不同域名的 iframe 是限制互相访问的。  
如果没有这个限制，我们做一个假的网站（a.com），内嵌一个 iframe（指向银行网站（b.com）），把 iframe 做成银行网站的样子，用户有时候是看不出来的，这时如果用户输入账号密码，我们的主网站（假网站（a.com））可以跨域访问到指向银行网站（b.com）iframe 的 DOM 节点，就可以拿到用户的账户密码了。  

XHR 同源策略  
XMLHttpRequest 可以发送 ajax，浏览器禁止使用 XHR 对象向不同源的服务器地址发起 HTTP 请求。  
比如，用户登录了自己的银行页面（a.com），（a.com） 向用户的 cookie 中添加用户标识；同时，用户浏览了恶意页面（b.com），执行了页面中的恶意 AJAX 请求代码。恶意页面（b.com） 向（a.com）发起 AJAX HTTP 请求，请求会默认把（a.com）对应 cookie 也同时发送过去，银行页面从发送的 cookie 中提取用户标识，验证用户无误，response 中返回请求数据。此时数据就泄露了。  

### 跨域解决方案
支持跨域的标签  
html 中有许多标签是支持跨域的：img、script、css、video、audio、object、embed、applet、@font-face、frame、iframe。  
> <link src=""> 标签嵌入 CSS  
> 由于 CSS 的松散的语法规则，CSS 的跨域需要一个设置正确的 Content-Type 消息头  
> 不同浏览器有不同的限制： IE, Firefox, Chrome, Safari (跳至 CVE-2010-0051) 部分 和 Opera  
>
> @font-face 引入的字体  
> 一些浏览器允许跨域字体（ cross-origin fonts）  
> 一些需要同源字体（same-origin fonts）  
>
> <frame> 和 <iframe > 载入的任何资源  
> 站点可以使用 X-Frame-Options 消息头来阻止这种形式的跨域交互

跨域资源共享（CORS）  
CORS，W3C 提出了的一个标准。目前，所有浏览器都支持 CORS 功能。  
CORS 需要浏览器和服务器同时支持，如果设置了 CORS，那么浏览器就会允许我们发送的 ajax 突破同源策略，向不同的服务器发送请求。  
```js
// 简单请求
// 简单请求请求方法只有三种：HEAD、GET、POST    
// 简单请求的请求头不超出以下字段：Accept、Accept-Language、Content-Language、Last-Event-ID、Content-Type（只限于三个值：application/x-www-form-urlencoded，multipart/form-data，text/plain）


// 非简单请求
// 除了简单请求以外的 CORS 请求  
// 非简单请求是那种对服务器有特殊要求的请求，比如请求方法是 PUT 或 DELETE，或者 Content-Type 字段的类型是 application/json
// 非简单请求的 CORS 请求，会在正式通信之前，增加一次 HTTP 查询请求，称为 "预检" 请求（preflight）  
// 预检过程：
// 1、浏览器先询问服务器，当前网页所在的域名是否在服务器的许可名单之中，以及可以使用哪些 HTTP 动词和头信息字段
// 2、只有得到肯定答复，浏览器才会发出正式的 XMLHttpRequest 请求，否则就报错  
// 3、"预检" 请求用的请求方法是 OPTIONS，表示这个请求是用来询问的  
// 4、头信息里面，关键字段是 Origin，表示请求来自哪个源
// 预检请求字段包含：Origin（请求来源）、Access-Control-Request-Method（请求会用到哪些 HTTP 方法）、Access-Control-Request-Headers（该字段是一个逗号分隔的字符串，指定浏览器 CORS 请求会额外发送的头信息字段）、	Access-Control-Allow-Methods（该字段是一个逗号分隔的字符串，表明服务器支持的所有跨域请求的方法）、Access-Control-Allow-Headers（该字段是一个逗号分隔的字符串，表明服务器支持的所有头信息字段，不限于浏览器在 "预检" 中请求的字段）、Access-Control-Max-Age（指定本次预检请求的有效期，单位为秒）

// CORS 后台配置示例
let app = express()
app.all('*',(req,res,next)=>{
    // 消除中文乱码
    res.set('Content-Type','application/json;charset=utf-8');

    // 设置跨域访问CORS
    // 如果要发送 Cookie，Access-Control-Allow-Origin 就不能设为通配符 *
    res.header("Access-Control-Allow-Origin", "*");
    res.header("Access-Control-Allow-Headers", "X-Requested-With");
    res.header("Access-Control-Allow-Methods","PUT,POST,GET,DELETE,OPTIONS");
    res.header("X-Powered-By",' 3.2.1')
    res.header("Content-Type", "application/json;charset=utf-8");
    next();
})
```

JSONP（JSON with Padding）  
HTML 脚本元素是可以规避 SOP 检查的，我们可以利用这个，采用动态注入脚本的方式来解决跨域问题。  
基本思路：网页通过添加一个 `<script>` 元素，向服务器请求 JSON 数据，这种做法不受同源政策限制，服务器收到请求后，将数据放在一个指定名字的回调函数里传回来。  
实现步骤：首先前端先设置好回调函数，并将其作为 url 的参数；服务端接收到请求后，通过该参数获得回调函数名，并将数据放在参数中将其返回；收到结果后因为是 script 标签，所以浏览器会当做是脚本进行运行，从而达到跨域获取数据的目的。  
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>jsonp 模拟</title>
</head>
<body>
    <script>
	function jsonpCallback(data) {
	    alert('获得 name 是:' + data.name);
	}
    </script>
    <script src="http://127.0.0.1:3000?callback=jsonpCallback"></script>
</body>
</html>
```
```js
// server.js
const url = require('url');
	
require('http').createServer((req, res) => {

    const data = {
		name: 'vikey'
    };
    // url.parse：解析一个 URL 字符串并返回一个 URL 对象
    // query.callback：获得callback对应的方法名
    const callback = url.parse(req.url, true).query.callback;
    res.writeHead(200);
    // 返回一个函数的调用
    // 这里是：jsonpCallback({name:"vikey"})
    // jsonpCallback 是获取到的前端传递的方法名
    // 返回：jsonpCallback ({name:"vikey"}) 等同于 <script>jsonpCallback({name:"vikey"})</script>
    res.end(`${callback}(${JSON.stringify(data)})`);

}).listen(3000, '127.0.0.1');

console.log('启动服务，监听 127.0.0.1:3000');
```
JSONP 优点  
它不像 XMLHttpRequest 对象实现 Ajax 请求那样受到同源策略的限制；   
它的兼容性更好，在更加古老的浏览器中都可以运行，不需要 XMLHttpRequest 或 ActiveX 的支持；  
在请求完毕后可以通过调用 callback 的方式回传结果，将回调方法的权限给了调用方。  
JSONP 缺点（局限性）  
它支持 GET 请求而不支持 POST 等其它类行的 HTTP 请求，意味着 GET 请求的诸多限制 JSONP 也有；  
它只支持跨域 HTTP 请求这种情况；  
jsonp 在调用失败的时候不会返回各种 HTTP 状态码，如果脚本注入成功后，就会调用回调函数，但是注入失败后，没有任何提示；这就意味着，当 JSONP 遇到 404、505 或者其他服务器错误时，你是无法检测出错原因的;我们能够做的也只有超时，没有收到响应，便认为请求失败，执行对应的错误回调。  
借助 JSONP 有可能进行跨站请求伪造 (CSRF) 攻击。  

服务器代理（Server Proxy）  
跨域是因为浏览器的限制，在服务端是不存在跨域的。服务器代理，顾名思义，当你需要有跨域的请求操作时发送请求给后端，让后端帮你代为请求，然后最后将获取的结果发送给你。  
```js
const url = require('url');
const http = require('http');
const https = require('https');
// 我们请求的是 /api/newsList
const server = http.createServer((req, res) => {
    const path = url.parse(req.url).path.slice(1);
    if(path === 'newsList') { 
    // 服务器去另一个域名帮我们请求数据，进行转发
	https.get('/api/newsList', (resp) => {
	    let data = "";
	    resp.on('data', chunk => {
		data += chunk;
	    });
	    resp.on('end', () => {
		res.writeHead(200, {
		    'Content-Type': 'application/json; charset=utf-8'
		});
        // 将代为请求的数据返回给我们
		res.end(data);
	    });
	})		
    }
}).listen(3000, '127.0.0.1');

console.log('启动服务，监听 127.0.0.1:3000');
```
```js
// webpack 配置代理
// webpack 配置代理只是开发环境可以使用，线上环境一般是 nginx 配置代理，或者是服务器代码中添加 api 转发之类的功能
proxyTable: {  
    '/api': {  //代理地址  
        target: 'http://1.1.1.1:8000/',  //需要代理的地址  
        changeOrigin: true,  //是否跨域  
        secure: false,    
        pathRewrite: {  
            '^/api': '/' //本身的接口地址没有 '/api' 这种通用前缀，所以要rewrite，如果本身有则去掉  
        }
    }
}

// 发送请求
// 请求数据时URL前加上“/api”就可以跨域请求了
this.$axios.get('/api/queryRole', {params: params})  
    .then((res) => {  
        console.log(res);  
    }).catch((err) => {  
        console.log(err);  
    })
```
