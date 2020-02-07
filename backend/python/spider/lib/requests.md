
### 1、概览
[requests 文档](https://requests.readthedocs.io/zh_CN/latest/)  
```bash
pip install requests
```

#### 1.1、实例引入
```python
# 引入Requests库
import requests

# 发起GET请求
response = requests.get('https://www.baidu.com/')
# 查看响应类型  requests.models.Response
print(type(response))
# 输出状态码
print(response.status_code)
# 输出响应内容类型  text
print(type(response.text))
# 输出响应内容
print(response.text)
# 输出cookies
print(response.cookies)
```

#### 1.2、各种请求方式
```python
import requests

# 发起POST请求
requests.post('http://httpbin.org/post')
# 发起PUT请求
requests.put('http://httpbin.org/put')
# 发起DELETE请求
requests.delete('http://httpbin.org/delete')
# 发送HEAD请求
requests.head('http://httpbin.org/get')
# 发送OPTION请求
requests.options('http://httpbin.org/get')
```

### 2、请求
#### 2.1 、基本 GET 请求
##### 2.1.1、基本写法
```python
import requests

response = requests.get('http://httpbin.org/get')
print(response.text)
```

##### 2.1.2、带参数的 GET 请求
```python
import requests

response = requests.get('http://httpbin.org/get?name=jyx&age=18')
print(response.text)
```

##### 2.1.3、带参数的 GET 请求 (2)
```python
import requests

# 分装GET请求参数
param = {'name':'jyx','age':19}
# 设置GET请求参数(Params)
response = requests.get('http://httpbin.org/get', params=param)
print(response.text)
```

##### 2.1.4、解析 json
```python
import requests

response = requests.get('http://httpbin.org/get')
# 获取响应内容
print(type(response.text))
# 如果响应内容是json,就将其转为json
print(response.json())
# 输出的是字典类型
print(type(response.json()))
```

##### 2.1.5、获取二进制数据
```python
import requests

response = requests.get('http://github.com/favicon.ico')

# str，bytes
print(type(response.text),type(response.content))
# 输出响应的文本内容
print(response.text)
# 输出响应的二进制内容
print(response.content)
# 下载二进制数据到本地
with open('favicon.ico','wb') as f:
  f.write(response.content)
  f.close()
```

##### 2.1.6、添加 headers
```python
import requests

# 设置User-Agent浏览器信息
headers = {
      "User-Agent": "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36"
}
# 设置请求头信息
response = requests.get('https://www.zhihu.com/explore',headers=headers)
print(response.text)
```

#### 2.2、基本 POST 请求
```python
import requests

# 设置传入post表单信息
data= { 'name':'jyx', 'age':18}
# 设置请求头信息
headers = {
      "User-Agent": "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36"
}
# 设置请求头信息和POST请求参数(data)
response = requests.post('http://httpbin.org/post', data=data, headers=headers)
print(response.text)
```

### 3、响应
#### 3.1 response 属性
```python
import requests

response = requests.get('http://www.jianshu.com/')
# 获取响应状态码
print(type(response.status_code),response.status_code)
# 获取响应头信息
print(type(response.headers),response.headers)
# 获取响应头中的cookies
print(type(response.cookies),response.cookies)
# 获取访问的url
print(type(response.url),response.url)
# 获取访问的历史记录
print(type(response.history),response.history)
```

#### 3.2、 状态码判断
```python
import requests

response = requests.get('http://www.jianshu.com/404.html')
# 使用request内置的字母判断状态码
if not response.status_code == requests.codes.ok:
    print('404-1')
response = requests.get('http://www.jianshu.com')
# 使用状态码数字判断
if not response.status_code == 200:
    print('404-2')
```

#### 3.3 requests 内置的状态字符
```
    100: ('continue',),
    101: ('switching_protocols',),
    102: ('processing',),
    103: ('checkpoint',),
    122: ('uri_too_long', 'request_uri_too_long'),

    200: ('ok', 'okay', 'all_ok', 'all_okay', 'all_good', '\\o/', '✓'),
    201: ('created',),
    202: ('accepted',),
    203: ('non_authoritative_info', 'non_authoritative_information'),
    204: ('no_content',),
    205: ('reset_content', 'reset'),
    206: ('partial_content', 'partial'),
    207: ('multi_status', 'multiple_status', 'multi_stati', 'multiple_stati'),
    208: ('already_reported',),
    226: ('im_used',),


    # Redirection.
    300: ('multiple_choices',),
    301: ('moved_permanently', 'moved', '\\'),
    302: ('found',),
    303: ('see_other', 'other'),
    304: ('not_modified',),
    305: ('use_proxy',),
    306: ('switch_proxy',),
    307: ('temporary_redirect', 'temporary_moved', 'temporary'),
    308: ('permanent_redirect', 'resume_incomplete', 'resume',), # These 2 to be removed in 3.0


    # Client Error.
    400: ('bad_request', 'bad'),
    401: ('unauthorized',),
    402: ('payment_required', 'payment'),
    403: ('forbidden',),
    404: ('not_found', '-'),
    405: ('method_not_allowed', 'not_allowed'),
    406: ('not_acceptable',),
    407: ('proxy_authentication_required', 'proxy_auth', 'proxy_authentication'),
    408: ('request_timeout', 'timeout'),
    409: ('conflict',),
    410: ('gone',),
    411: ('length_required',),
    412: ('precondition_failed', 'precondition'),
    413: ('request_entity_too_large',),
    414: ('request_uri_too_large',),
    415: ('unsupported_media_type', 'unsupported_media', 'media_type'),
    416: ('requested_range_not_satisfiable', 'requested_range', 'range_not_satisfiable'),
    417: ('expectation_failed',),
    418: ('im_a_teapot', 'teapot', 'i_am_a_teapot'),
    421: ('misdirected_request',),
    422: ('unprocessable_entity', 'unprocessable'),
    423: ('locked',),
    424: ('failed_dependency', 'dependency'),
    425: ('unordered_collection', 'unordered'),
    426: ('upgrade_required', 'upgrade'),
    428: ('precondition_required', 'precondition'),
    429: ('too_many_requests', 'too_many'),
    431: ('header_fields_too_large', 'fields_too_large'),
    444: ('no_response', 'none'),
    449: ('retry_with', 'retry'),
    450: ('blocked_by_windows_parental_controls', 'parental_controls'),
    451: ('unavailable_for_legal_reasons', 'legal_reasons'),
    499: ('client_closed_request',),


    # Server Error.
    500: ('internal_server_error', 'server_error', '/o\\', '✗'),
    501: ('not_implemented',),
    502: ('bad_gateway',),
    503: ('service_unavailable', 'unavailable'),
    504: ('gateway_timeout',),
    505: ('http_version_not_supported', 'http_version'),
    506: ('variant_also_negotiates',),
    507: ('insufficient_storage',),
    509: ('bandwidth_limit_exceeded', 'bandwidth'),
    510: ('not_extended',),
    511: ('network_authentication_required', 'network_auth', 'network_authentication'),
```

### 4、高级操作
#### 4.1、文件上传
```python
import requests

files = {'file':open('favicon.ico','rb')}
# 往POST请求头中设置文件(files)
response = requests.post('http://httpbin.org/post',files=files)
print(response.text)
```

#### 4.2、获取 cookies
```python
import requests

response = requests.get('https://www.baidu.com')
print(response.cookies)
for key,value in response.cookies.items():
    print(key,'=====',value)
```

#### 4.3、会话维持
##### 4.3.1、普通请求
```python
import requests

requests.get('http://httpbin.org/cookies/set/number/12456')
response = requests.get('http://httpbin.org/cookies')
# 本质上是两次不同的请求，session不一致
print(response.text)
```

##### 4.3.2、会话维持请求
```python
import requests

# 从Requests中获取session
session = requests.session()
# 使用seesion去请求保证了请求是同一个session
session.get('http://httpbin.org/cookies/set/number/12456')
response = session.get('http://httpbin.org/cookies')
print(response.text)
```

#### 4.4、证书验证
##### 4.4.1、无证书访问
```python
import requests

response = requests.get('https://www.12306.cn')
# 在请求https时，request会进行证书的验证，如果验证失败则会抛出异常
print(response.status_code)  
```

##### 4.4.2、关闭证书验证
```python
import requests

# 关闭验证，但是仍然会报出证书警告
response = requests.get('https://www.12306.cn',verify=False)
print(response.status_code)
```

##### 4.4.3、消除关闭证书验证的警告
```python
from requests.packages import urllib3
import requests

# 关闭警告
urllib3.disable_warnings()
response = requests.get('https://www.12306.cn',verify=False)
print(response.status_code)
```

##### 4.4.4、手动设置证书
```python
import requests

# 设置本地证书
response = requests.get('https://www.12306.cn', cert=('/path/server.crt', '/path/key'))
print(response.status_code)  
```

#### 4.5、代理设置
##### 4.5.1、设置普通代理
```python
import requests

proxies = {
    "http": "http://127.0.0.1:9743",
    "https": "https://127.0.0.1:9743",
}
# 往请求中设置代理(proxies
)
response = requests.get("https://www.taobao.com", proxies=proxies)
print(response.status_code)
```

##### 4.5.2、设置带有用户名和密码的代理
```python
import requests

proxies = {
    "http": "http://user:password@127.0.0.1:9743/",
}
response = requests.get("https://www.taobao.com", proxies=proxies)
print(response.status_code)
```

##### 4.5.3、设置 socks 代理
```bash
pip3 install requests[socks]
```

```python
import requests

proxies = {
    'http': 'socks5://127.0.0.1:9742',
    'https': 'socks5://127.0.0.1:9742'
}
response = requests.get("https://www.taobao.com", proxies=proxies)
print(response.status_code)
```

#### 4.6、超时设置
```python
import requests
from requests.exceptions import ReadTimeout

    
try:
   # 设置必须在500ms内收到响应，不然或抛出ReadTimeout异常
   response = requests.get("http://httpbin.org/get", timeout=0.5)
            print(response.status_code)
except ReadTimeout:
   print('Timeout')
```

#### 4.7、认证设置
```python
import requests
from requests.auth import HTTPBasicAuth

r = requests.get('http://120.27.34.24:9001', auth=HTTPBasicAuth('user', '123'))
# r = requests.get('http://120.27.34.24:9001', auth=('user', '123'))
print(r.status_code)
```

#### 4.8、异常处理
```python
import requests
from requests.exceptions import ReadTimeout, ConnectionError, RequestException

try:
   response = requests.get("http://httpbin.org/get", timeout = 0.5)
           print(response.status_code)
except ReadTimeout:
   # 超时异常
   print('Timeout')
except ConnectionError:
   # 连接异常
   print('Connection error')
except RequestException:
   # 请求异常
   print('Error')
```
