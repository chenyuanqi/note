
### 1、概览
urlib 库为 python3 的 HTTP 内置请求库。  

urilib 的四个模块：

- urllib.request: 用于获取网页的响应内容  
- urllib.error: 异常处理模块，用于处理异常的模块
- urllib.parse: 用于解析 url
- urllib.robotparse: 用于解析 robots.txt，主要用于看哪些网站不能进行爬取，不过少用

[官方文档](https://docs.python.org/3/library/urllib.html)  

### 2、Urllib 基本使用
**urllib.request**  
urllib.request.urlopen(url,data=None,[timeout,]\*,cafile=None,cadefault=False,context=None)

- url: 为请求网址  
- data: 请求时需要发送的参数
- timeout: 超时设置，在该时间范围内返回请求内容就不会报错

```python
from urllib import request

# 请求获取网页返回内容
response = request.urlopen('https://movie.douban.com/')
# 获取网页返回内容
print(response.read().decode('utf-8'))
# 获取状态码
print(response.status)
# 获取请求头
print(response.getheaders())

# 对请求头进行遍历
for k, v in response.getheaders():
    print(k, '=', v)
```

当需要添加请求 Header 时，需要用到 Request 对象。
```python
# 请求头
headers = {'User-Agent':'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.221 Safari/537.36 SE 2.X MetaSr 1.0'}
requests = request.Request('https://movie.douban.com/', headers=headers) # 加入自己的请求头更加接近浏览器
# 进行请求,把Request对象传入urlopen参数中
response = request.urlopen(requests)
print(response.read().decode('utf-8'))

# 如果网站需要进行登陆，这时需要用到 post 方法
from urllib import request, parse
# 使用post方法来进行模拟登陆豆瓣
data = {'source': 'None',
    'redir': 'https://www.douban.com/',
    'form_email': 'user',
    'form_password': 'passwd',
    'remember': 'on',
    'login': '登录'
    }
# 将 data 的字典类型转换为get请求方式
data = bytes(parse.urlencode(data), encoding='utf-8')
# 添加请求头还可以这样：Request.add_header("source", "None")
requests = request.Request('https://accounts.douban.com/login', headers=headers, data=data, method='POST')
response = request.urlopen(requests)
print(response.read().decode('utf-8'))

# 在登陆了网站之后，我们需要用到 cookie 来保存登陆信息，这时就需要获取 cookie 了
from http import cookiejar
# 获取cookie
cookie = cookiejar.CookieJar()
# 获取助手把cookie传进去
handler = request.HTTPCookieProcessor(cookie)
# 获取opener进行请求网站
opener = request.build_opener(handler)
# 请求网页
response = opener.open('https://movie.douban.com/')
# 打印 cookie
for c in cookie:
    print(c.name, '=', c.value)

# 保存 cookie 为文件
from http import cookiejar
# 将cookie保存在文件中
filename = 'cookie.txt'
cookie = cookiejar.MozillaCookieJar(filename) # 表示使用 Mozilla 的 cookie 方式存储和读取
# cookie = cookiejar.LWPCookieJar(filename) # 表示 Set-Cookie3 文件格式存储和读取
handler = request.HTTPCookieProcessor(cookie) # 创建 cookie 处理器
opener = request.build_opener(handler) # 利用 cookie 处理器构建 opener
opener.open('https://movie.douban.com/') # 进行请求网站，不能用 request.urlopen()
# 保存文件
cookie.save(ignore_discard=True, ignore_expires=True)

# 加载 cookie
from http import cookiejar
# 从 cookie 文件加载到网页上实现记住登陆
cookie = cookiejar.LWPCookieJar()
# 加载文件
cookie.load(filename, ignore_discard=True, ignore_expires=True)
handler = request.HTTPCookieProcessor(cookie) # 创建 cookie 处理器
opener = request.build_opener(handler)
opener.open('https://movie.douban.com/')

# 如果有时你在同一 ip 连续多次发送请求，会有被封 ip 的可能，这时我们还需要用到代理 ip 进行爬取
proxy = request.ProxyHandler({
    'https': 'https://106.60.34.111:80'
})
opener = request.build_opener(proxy)
opener.open('https://movie.douban.com/', timeout=1)
```

**urllib.error**  
因为有时这个 ip 或许也被封了，有可能会抛出异常。所以，将上面的使用代理 ip 的请求进行异常处理。  
```python
from urllib import request, error
try:
    proxy = request.ProxyHandler({
        'https': 'https://106.60.34.111:80'
    })
    opener = request.build_opener(proxy)
    opener.open('https://movie.douban.com/', timeout=1)
except error.HTTPError as e: # 是 error.URLError 的子类
    print(e.reason(), e.code(), e.headers())
except error.URLError as e: # 这个异常只有一个 reason 属性
    print(e.reason) 
```

**urllib.parse**  
解析 url:urllib.parse.urlparse (url, scheme='', allow_fragments=True)  

```python
from urllib import request, parse
# 解析 url 
print(parse.urlparse('https://movie.douban.com/'))
# 当 scheme 协议加了，而前面的 url 也包含协议，一般会忽略后面的 scheme 参数
print(parse.urlparse('https://movie.douban.com/', scheme='http')) # 加了 scheme 参数和没加的返回结果是有区别的
print(parse.urlparse('movie.douban.com/', scheme='http')) 
# 下面是结果
ParseResult(scheme='https', netloc='movie.douban.com', path='/', params='', query='', fragment='')
ParseResult(scheme='https', netloc='movie.douban.com', path='/', params='', query='', fragment='')
ParseResult(scheme='http', netloc='', path='movie.douban.com/', params='', query='', fragment='')

# 把元素串连成一个 url
url = ['http', 'www', 'baidu', 'com', 'dfdf', 'eddffa'] # 这里至少需要 6 个元素，要不然会抛出异常
print(parse.urlunparse(url)) # http://www/baidu.com?dfdf=eddffa

# urllib.parse.urljoin (): 这个是将第二个参数的 url 缺少的部分用第一个参数的 url 补齐
print(parse.urljoin('https://movie.douban.com/', 'index')) # https://movie.douban.com/index
print(parse.urljoin('https://movie.douban.com/', 'https://accounts.douban.com/login')) # https://accounts.douban.com/login

# urllib.parse.urlencode (): 这个方法是将字典类型的参数转为请求为 get 方式的字符串
data = {'name': 'sergiojuue', 'sex': 'boy'}
data = parse.urlencode(data)
print('https://accounts.douban.com/login'+data) # https://accounts.douban.com/loginname=sergiojuue&sex=boy
```
