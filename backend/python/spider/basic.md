
### 快速理解 HTTP 协议
HTTP 协议是互联网应用中，客户端（浏览器）与服务器之间进行数据通信的一种协议。协议中规定了客户端应该按照什么格式给服务器发送请求，同时也约定了服务端返回的响应结果应该是什么格式。  
HTTP 协议本身是非常简单的。它规定，只能由客户端主动发起请求，服务器接收请求处理后返回响应结果，同时 HTTP 是一种无状态的协议，协议本身不记录客户端的历史请求记录。  

HTTP 请求由 3 部分组成，分别是请求行、请求首部、请求体，首部和请求体是可选的，并不是每个请求都需要的。  
> 请求行是每个请求必不可少的部分，它由 3 部分组成，分别是请求方法（method)、请求 URL（URI）、HTTP 协议版本，以空格隔开。  
> HTTP 协议中最常用的请求方法有：GET、POST、PUT、DELETE。GET 方法用于从服务器获取资源，90% 的爬虫都是基于 GET 请求抓取数据。  
> 请求 URL 是指资源所在服务器的路径地址，比如上图的例子表示客户端想获取 index.html 这个资源，它的路径在服务器 foofish.net 的根目录（/）下面。  
> 
> 因为请求行所携带的信息量非常有限，以至于客户端还有很多想向服务器要说的事情不得不放在请求首部（Header），请求首部用于给服务器提供一些额外的信息，比如 User-Agent 用来表明客户端的身份，让服务器知道你是来自浏览器的请求还是爬虫，是来自 Chrome 浏览器还是 FireFox。  
> HTTP/1.1 规定了 47 种首部字段类型。HTTP 首部字段的格式很像 Python 中的字典类型，由键值对组成，中间用冒号隔开。比如：User-Agent: Mozilla/5.0  
> 因为客户端发送请求时，发送的数据（报文）是由字符串构成的，为了区分请求首部的结尾和请求体的开始，用一个空行来表示，遇到空行时，就表示这是首部的结尾，请求体的开始。  
> 
> 请求体是客户端提交给服务器的真正内容，比如用户登录时的需要用的用户名和密码，比如文件上传的数据，比如注册用户信息时提交的表单信息。

服务端接收请求并处理后，返回响应内容给客户端，同样地，响应内容也必须遵循固定的格式浏览器才能正确解析。HTTP 响应也由 3 部分组成，分别是：响应行、响应首部、响应体，与 HTTP 的请求格式是相对应的。  
> 响应行同样也是 3 部分组成，由服务端支持的 HTTP 协议版本号、状态码、以及对状态码的简短原因描述组成。  
> 状态码是响应行中很重要的一个字段。通过状态码，客户端可以知道服务器是否正常处理的请求。如果状态码是 200，说明客户端的请求处理成功，如果是 500，说明服务器处理请求的时候出现了异常。404 表示请求的资源在服务器找不到。除此之外，HTTP 协议还很定义了很多其他的状态码。  
> 
> 响应首部和请求首部类似，用于对响应内容的补充，在首部里面可以告知客户端响应体的数据类型是什么？响应内容返回的时间是什么时候，响应体是否压缩了，响应体最后一次修改的时间。  
> 
> 响应体（body）是服务器返回的真正内容，它可以是一个 HTML 页面，或者是一张图片、一段视频等等。	

[更多参考《图解 HTTP》](https://gitbook.cn/gitchat/geekbook/5acad23fdbd50e7493d38305)  

### 优雅的 HTTP 库 requests
Python 提供了很多模块来支持 HTTP 协议的网络编程，urllib、urllib2、urllib3、httplib、httplib2，都是和 HTTP 相关的模块，看名字觉得很反人类，更糟糕的是这些模块在 Python2 与 Python3 中有很大的差异，如果业务代码要同时兼容 2 和 3，写起来会让人崩溃。  
幸运的是，繁荣的 Python 社区给开发者带来了一个非常惊艳的 HTTP 库 requests，一个真正给人用的 HTTP 库。  

requests 实现了 HTTP 协议中绝大部分功能，它提供的功能包括 Keep-Alive、连接池、Cookie 持久化、内容自动解压、HTTP 代理、SSL 认证、连接超时、Session 等很多特性，最重要的是它同时兼容 python2 和 python3。  

requests 的安装可以直接使用 pip 方法。  
```bash
pip install requests
```

requests 的使用。  
```python
import requests
import json

# requests 除了支持 GET 请求外，还支持 HTTP 规范中的其它所有方法，包括 POST、PUT、DELTET、HEADT、OPTIONS 方法
response = requests.get("https://foofish.net") # 请求返回 Response 对象
# Response 对象是 对 HTTP 协议中服务端返回给浏览器的响应数据的封装，响应的中的主要元素包括：状态码、原因短语、响应首部、响应体等等，这些属性都封装在 Response 对象中
response.status_code # 响应状态码
response.reason # 原因短语
for name,value in response.headers.items(): # 响应首部
	print("%s:%s" % (name, value))
response.content # 响应内容

# 指定编码类型
response.encoding = 'utf-8'

# 构建请求查询参数
args = {"p": 4, "s": 20}
response = requests.get("http://fav.foofish.net", params = args)
response.url # http://fav.foofish.net/?p=4&s=20

# 构建请求首部 Headers
r = requests.get(url, headers={'user-agent': 'Mozilla/5.0'})

# 代理设置
# 当爬虫频繁地对服务器进行抓取内容时，很容易被服务器屏蔽掉，代理是明智的选择
proxies = {
  'http': 'http://10.10.1.10:3128',
  'https': 'http://10.10.1.10:1080',
}
requests.get('http://example.org', proxies=proxies)

# 超时设置
# 请求线程一直阻塞，直到有响应返回才处理后面的逻辑，超时设置是必须的
r = requests.get("http://www.google.coma", timeout=5)

# 当一个网站不安全，需要你用证书验证时
response = requests.get('https://www.12306.cn', verify=False) 
# 不加这个关键字参数的话会出现验证错误问题，因为这个网站的协议不被信任

# 构建 POST 请求数据 - 作为表单数据传输给服务器
payload = {'key1': 'value1', 'key2': 'value2'}
r = requests.post("http://httpbin.org/post", data=payload)
# 构建 POST 请求数据 - 作为 json 格式的字符串格式传输给服务器
payload = {'some': 'data'}
r = requests.post('http://httpbin.org/post', json=payload)

# 上传文件
files = {'picture': open('baidu.png', 'rb')}
response = requests.post('http://httpbin.org/post', files=files)
print(response.text)

# Response 中的响应体
# content 是 byte 类型，适合直接将内容保存到文件系统或者传输到网络中
r = requests.get("https://pic1.zhimg.com/v2-2e92ebadb4a967829dcd7d05908ccab0_b.jpg")
type(r.content) # <class 'bytes'>
with open("test.jpg", "wb") as f: # 另存为 test.jpg
	f.write(r.content)
# 一个普通的 HTML 页面，需要对文本进一步分析时，使用 text
r = requests.get("https://foofish.net/understand-http.html") 
type(r.text) # <class 'str'>
re.compile('xxx').findall(r.text)
#  API 接口返回的内容是 json 格式的数据时，那么可以直接使用 json () 方法返回一个经过 json.loads () 处理后的对象
r = requests.get('https://www.v2ex.com/api/topics/hot.json')
r.json() # [{'id': 352833, 'title': '...'}]

# 获取头部
print(r.headers)
# 获取cookie
print(r.cookies)

# Session
session  = requests.Session() # 构建会话
session.post(login_url, data={username, password}) #　登录 url
# 登录后才能访问的 url
r = session.get(home_url)
session.close()

# 请求异常处理
import requests
from requests.exceptions import ReadTimeout, ConnectTimeout, HTTPError, ConnectionError, RequestException
# 捕捉异常
try:
    response = requests.get('http://httpbin.org/get', timeout=0.1) # 规定时间内未响应就抛出异常
    print(response.text)
except ReadTimeout as e:
    print('请求超时')
except ConnectionError as e:
    print('连接失败')
except RequestException as e:
    print('请求失败')

```

**HTML 文本解析库 BeautifulSoup**  
BeautifulSoup 是一个用于解析 HTML 文档的 Python 库，通过 BeautifulSoup，你只需要用很少的代码就可以提取出 HTML 中任何感兴趣的内容，此外，它还有一定的 HTML 容错能力，对于一个格式不完整的 HTML 文档，它也可以正确处理。
```bash
# BeautifulSoup3 被官方放弃维护
pip install beautifulsoup4
```

学习 BeautifulSoup4 前有必要先对 HTML 文档有一个基本认识，如下代码，HTML 是一个树形组织结构。  
```html
<html>  
    <head>
     <title>hello, world</title>
    </head>
    <body>
        <h1>BeautifulSoup</h1>
        <p>如何使用BeautifulSoup</p>
    <body>
</html>
```
它由很多标签（Tag）组成，比如 html、head、title 等等都是标签；  
一个标签对构成一个节点，比如 <html\>...</html\> 是一个根节点；  
节点之间存在某种关系，比如 h1 和 p 互为邻居，他们是相邻的兄弟（sibling）节点；  
h1 是 body 的直接子（children）节点，还是 html 的子孙（descendants）节点；  
body 是 p 的父（parent）节点，html 是 p 的祖辈（parents）节点；  
嵌套在标签之间的字符串是该节点下的一个特殊子节点，比如 “hello, world” 也是一个节点，只不过没名字。

```python
from bs4 import BeautifulSoup  

text = """
<html>  
    <head>
     <title >hello, world</title>
    </head>
    <body>
        <h1>BeautifulSoup</h1>
        <p class="bold">如何使用BeautifulSoup</p>
        <p class="big" id="key1"> 第二个p标签</p>
        <a href="http://foofish.net">python</a>
    </body>
</html>  
"""
# 构建一个 BeautifulSoup 对象需要两个参数：将要解析的 HTML 文本字符串，告诉 BeautifulSoup  使用哪个解析器来解析 HTML（解析器负责把 HTML 解析成相关的对象，而 BeautifulSoup 负责操作数据（增删改查））
# “html.parser” 是 Python 内置的解析器，“lxml” 则是一个基于 c 语言开发的解析器，它的执行速度更快，不过它需要额外安装
soup = BeautifulSoup(text, "html.parser")
# 获取 title 标签
soup.title # <title>hello, world</title>
# 获取 p 标签
soup.p # <p class="bold">\u5982\u4f55\u4f7f\u7528BeautifulSoup</p>
# 获取 p 标签的内容
soup.p.string # u'\u5982\u4f55\u4f7f\u7528BeautifulSoup'

# BeatifulSoup 将 HTML 抽象成为 4 类主要的数据类型，分别是 Tag , NavigableString , BeautifulSoup，Comment 
# BeautifulSoup 对象代表整个 HTML 文档
type(soup)
# 每个标签节点就是一个 Tag 对象
type(soup.h1) # <class 'bs4.element.Tag'>
# NavigableString 对象一般是包裹在 Tag 对象中的字符串
type(soup.p.string) # <class 'bs4.element.NavigableString'>

# 每个 Tag 都有一个名字，它对应 HTML 的标签名称
soup.h1.name # u'h1'
# 标签还可以有属性，属性的访问方式和字典是类似的，它返回一个列表对象
soup.p['class'] # [u'bold']
# 获取标签中的内容，直接使用 .stirng 即可获取，它是一个 NavigableString 对象，你可以显式地将它转换为 unicode 字符串
soup.p.string # u'\u5982\u4f55\u4f7f\u7528BeautifulSoup'
type(soup.p.string) # <class 'bs4.element.NavigableString'>
unicode_str = unicode(soup.p.string) # unicode_str：u'\u5982\u4f55\u4f7f\u7528BeautifulSoup'
```

如何从 HTML 中找到我们关心的数据？  
BeautifulSoup 提供了两种方式，一种是遍历，另一种是搜索，通常两者结合来完成查找任务。
> 遍历文档树，顾名思义，就是是从根节点 html 标签开始遍历，直到找到目标元素为止，遍历的一个缺陷是，如果你要找的内容在文档的末尾，那么它要遍历整个文档才能找到它，速度上就慢了。  
> 因此还需要配合第二种方法。  
> 遍历文档树的另一个缺点是只能获取到与之匹配的第一个子节点，例如，如果有两个相邻的 p 标签时，第二个标签就没法通过 .p 的方式获取，这是需要借用 next_sibling 属性获取相邻的节点。  
> 此外，还有很多不怎么常用的属性，比如：.contents 获取所有子节点，.parent 获取父节点，更多的参考请查看官方文档。  
> 
> 搜索文档树是通过指定标签名来搜索元素，还可以通过指定标签的属性值来精确定位某个节点元素，最常用的两个方法就是 find 和 find_all。这两个方法在 BeatifulSoup 和 Tag 对象上都可以被调用。  
> find_all 的语法：find_all(name , attrs , recursive , text , \*\*kwargs)  
> find_all 的返回值是一个 Tag 组成的列表，方法调用非常灵活，所有的参数都是可选的  
> find 方法跟 find_all  类似，唯一不同的地方是，它返回的单个 Tag 对象而非列表，如果没找到匹配的节点则返回 None。如果匹配多个 Tag，只返回第 0 个  

```python
# 通过遍历文档树的方式获取标签节点可以直接通过 .标签名的方式获取，如获取 body 标签
soup.body # <body>\n<h1>BeautifulSoup<...

# 找到所有标签名为 title 的节点
soup.find_all("title") # [<title>hello, world</title>]
soup.find_all("p") 
# [<p class="bold">\xc8\xe7\xba\xce\xca....</p>, 
# <p class="big"> \xb5\xda\xb6\xfe\xb8\xf6p...</p>]

# 找到所有 class 属性为 big 的 p 标签
soup.find_all("p", "big") # 等价于 soup.find_all("p", class_="big")，[<p class="big"> \xb5\xda\xb6\xfe\xb8\xf6p\xb1\xea\xc7\xa9</p>]
# 查找有 href 属性值为 “http://foofish.net” 的标签
soup.find_all(href="http://foofish.net") # [<a href="http://foofish.net">python</a>]

# 支持正则
import re
soup.find_all(href=re.compile("^http")) # [<a href="http://foofish.net">python</a>]

# 具体的属性值
soup.find_all(id="key1") # [<p class="big" id="key1"> \xb5\xda\xb6\xfe\xb8\xf6p\xb1\xea\xc7\xa9</p>]
# 布尔值（True/Flase）表示有属性或者没有该属性
soup.find_all(id=True)

# 遍历和搜索相结合查找，先定位到 body 标签，缩小搜索范围，再从 body 中找 a 标签
body_tag = soup.body
body_tag.find_all("a") # [<a href="http://foofish.net">python</a>]

body_tag.find("a") # <a href="http://foofish.net">python</a>
body_tag.find("p") # <p class="bold">\xc8\xe7\xba\xce\xca\xb9\xd3\xc3BeautifulSoup</p>

# 获取标签里面内容，除了可以使用 .string 之外，还可以使用 get_text 方法
# 不同的地方在于前者返回的一个 NavigableString 对象，后者返回的是 unicode 类型的字符串
p1 = body_tag.find('p').get_text() # u'\xc8\xe7\xba\xce\xca\xb9\xd3\xc3BeautifulSoup'
type(p1) # <type 'unicode'>
```

**正则表达式**  
正则表达式处理文本有如疾风扫秋叶，绝大部分编程语言都内置支持正则表达式，它应用在诸如表单验证、文本提取、替换等场景。爬虫系统更是离不开正则表达式，用好正则表达式往往能收到事半功倍的效果。  

在 Python 中，正则表达式一般用原始字符串的形式来定义。比如对于字符 "\b" 来说，它在 ASCII 中是有特殊意义的，表示退格键，而在正则表达式中，它是一个特殊的元字符，用于匹配一个单词的边界，为了能让正则编译器正确地表达它的意义就需要用原始字符串，当然也可以使用反斜杠 “\” 对常规定义的字符串进行转义。  

正则表达式由普通文本字符和特殊字符（元字符）两种字符组成。  
元字符在正则表达式中具有特殊意义，它让正则表达式具有更丰富的表达能力。例如，正则表达式 r"a.d"中 ，字符 ‘a’ 和 ‘d’ 是普通字符，’.’ 是元字符，. 可以指代任意字符，它能匹配 ‘a1d’、’a2d’、’acd’ 。  
```python
rex = r"a.d"   # 正则表达式文本
original_str = "and"  # 原始文本

pattern = re.compile(rex)  # 正则表达式对象
m = pattern.match(original_str)  # 匹配对象，m：<_sre.SRE_Match object at 0x101c85b28>
# 等价于 re.match(r"a.d", "and")
# 如果原文本字符串与正则表达式匹配，那么就会返回一个 Match 对象
# 当不匹配时，match 方法返回的 None，通过判断 m 是否为 None 可进行表单验证
```

基本元字符有如下这些：  

- .：匹配除换行符以外的任意一个字符，例如：”a.c” 可以完全匹配 “abc”，也可以匹配 “abcef” 中的 “abc”
- \： 转义字符，使特殊字符具有本来的意义，例如： 1\.2 可以匹配 1.2
- [...]：匹配方括号中的任意一个字符，例如：a [bcd] e 可以匹配 abe、ace、ade，它还支持范围操作，比如：a 到 z 可表示为 “a-z”，0 到 9 可表示为 “0-9”，注意，在 “[]” 中的特殊字符不再有特殊意义，就是它字面的意义，例如：[.*]就是匹配。或者 *
- [^...]，字符集取反，表示只要不是括号中出现的字符都可以匹配，例如：a [^bcd] e 可匹配 aee、afe 等

```python
# group 方法返回原字符串 (abcef) 中与正则表达式相匹配的那部分子字符串 (abc)
# 提前是要匹配成功 match 方法才会返回 Match 对象，进而才有 group 方法
re.match(r"a.c", "abc").group() # abc
re.match(r"a.c", "abcef").group() # abc
re.match(r"1\.2", "1.2").group() # 1.2
re.match(r"a[0-9]b", "a2b").group() # a2b
re.match(r"a[0-9]b", "a5b11").group() # a5b
re.match(r"a[.*?]b", "a.b").group() # a.b
re.match(r"abc[^\w]", "abc!123").group() # abc!
```

预设元字符有如下这些：  

- \w 匹配任意一个单词字符，包括数字和下划线，它等价于 [A-Za-z0-9_]，例如 a\wc 可以匹配 abc、acc
- \W 匹配任意一个非单词字符，与 \w 操作相反，它等价于 [^A-Za-z0-9_]，例如： a\Wc 可匹配 a!c
- \s 匹配任意一个空白字符，空格、回车等都是空白字符，例如：a\sc 可以配 a\nc，这里的 \n 表示回车
- \S 匹配任意一个非空白字符
- \d 匹配任意一个数字，它等价于 [0-9]，例如：a\dc 可匹配 a1c、a2c …
- \D 匹配任意一个非数字

边界匹配（相关的符号专门用于修饰字符）：  

- ^ 匹配字符的开头，在字符串的前面，例如：^abc 表示匹配 a 开头，后面紧随 bc 的字符串，它可以匹配 abc
- $ 匹配字符的结尾，在字符串的末尾位置，例如： hello$

```python
re.match(r"^abc","abc").group() # abc
re.match(r"^abc$","abc").group() # abc
```

如果希望匹配的字符重复出现，重复匹配：  

- * 重复匹配零次或者更多次
- ? 重复匹配零次或者一次
- + 重复匹配 1 次或者多次
- {n} 重复匹配 n 次
- {n,} 重复匹配至少 n 次
- {n, m} 重复匹配 n 到 m 次

```python
# 简单匹配身份证号码，前面17位是数字，最后一位可以是数字或者字母X
re.match(r"\d{17}[\dX]", "42350119900101153X").group() # 42350119900101153X
# 匹配 5 到 12 的 QQ 号码
re.match(r"\d{5,12}$", "4235011990").group() # 4235011990
```

逻辑分支：  

- | 把表达式分为左右两部分，先尝试匹配左边部分，如果匹配成功就不再匹配后面部分了，这是逻辑 “或” 的关系


```python
# abc|cde 可以匹配abc 或者 cde，但优先匹配abc
re.match(r"aa(abc|cde)","aaabccde").group() # aaabc
# 既可以匹配 3 位区号 8 位号码，也可以匹配 4 位区号 7 位号码
re.match(r"0\d{2}-\d{8}|0\d{3}-\d{7}", "0755-4348767").group() # 0755-4348767
```

分组：  
如果想要重复匹配多个字符，需要用子表达式（也叫分组）来表示，分组用小括号”()” 表示。  
关于分组，group 方法可用于提取匹配的字符串分组，默认它会把整个表达式的匹配结果当做第 0 个分组，就是不带参数的 group() 或者是 group(0)，第一组括号中的分组用 group(1) 获取，以此类推。  
```python
# 从字符串中提取出想要的信息
m = re.match(r"(\d+)(\w+)", "123abc")
# 分组０，匹配整个正则表达式
m.group() # 等价于 m.group(0)，123abc
m.group(1) # 分组 1，匹配第一对括号，123
m.group(2) # 分组2，匹配第二对括号，abc

# 通过指定名字的方式获取
# 第一个分组的名字是 number
# 第二个分组的名字是 char
m = re.match(r"(?P<number>\d+)(?P<char>\w+)", "123abc")
m.group("number") # 123
```

贪婪与非贪婪：  
默认情况下，正则表达式重复匹配时，在使整个表达式能得到匹配的前提下尽可能匹配多的字符，我们称之为贪婪模式，是一种贪得无厌的模式。例如： r"a.\*b" 表示匹配 a 开头 b 结尾，中间可以是任意多个字符的字符串，如果用它来匹配 aaabcb，那么它会匹配整个字符串。
有时，我们希望尽可能少的匹配，怎么办？只需要在量词后面加一个问号” ？”，在保证匹配的情况下尽可能少的匹配，比如刚才的例子，我们只希望匹配 aaab，那么只需要修改正则表达式为 r"a.\*?b"。  
```python
re.match(r"a.*b", "aaabcb").group() # aaabcb
re.match(r"a.*?b", "aaabcb").group() # aaab
```

### 爬虫小技巧
1、写脚本的时候记得带 request headers，特别是 host 和 user-agent 这俩个字段（通常带这两个就足够了，比如某些网站如果没有带上 ua，会一直拿不到 html），这个就是体现在对 http 协议理解不够。

2、注意请求参数被放到 cookies 的情况。  

3、ip 的访问频率，基本的职业素养是不要影响别人正常运营。  
处理的方式主要是：每一次请求时变更 ua，在 request headers 里带上 referer，设置时延，通过改变代理 ip 的方式等。  

4、验证码的问题，可以通过机器学习训练出可以识别 xx 网的验证码，也可以购买打码平台接口等。

5、


