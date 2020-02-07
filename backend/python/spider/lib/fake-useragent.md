
### 1、概览
fake-useragent, 可以伪装生成 headers 请求头中的 User Agent 值。  

```bash
pip install fake-useragent
```

### 2、fake-useragent 的基本使用
各浏览器的 user-agent 值使用。
```python
from fake_useragent import UserAgent

ua = UserAgent()
# ie 浏览器的 user agent
print(ua.ie) # Mozilla/5.0 (Windows; U; MSIE 9.0; Windows NT 9.0; en-US)

# opera 浏览器
print(ua.opera) # Opera/9.80 (X11; Linux i686; U; ru) Presto/2.8.131 Version/11.11

# chrome浏览器
print(ua.chrome) # Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.2 (KHTML, like Gecko) Chrome/22.0.1216.0 Safari/537.2

# firefox浏览器
print(ua.firefox) # Mozilla/5.0 (Windows NT 6.2; Win64; x64; rv:16.0.1) Gecko/20121011 Firefox/16.0.1

# safri浏览器
print(ua.safari) # Mozilla/5.0 (iPad; CPU OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5355d Safari/8536.25
```

随意变换 headers。  
```python
from fake_useragent import UserAgent

ua = UserAgent()

print(ua.random) # Mozilla/5.0 (compatible; MSIE 10.0; Macintosh; Intel Mac OS X 10_7_3; Trident/6.0)
print(ua.random) # Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.2 (KHTML, like Gecko) Chrome/22.0.1216.0 Safari/537.2
print(ua.random) # Opera/9.80 (X11; Linux i686; U; ru) Presto/2.8.131 Version/11.11
```

具体的应用。  
```python
import requests
from fake_useragent import UserAgent

ua = UserAgent()
headers = {'User-Agent': ua.random}
url = 'url'
response = requests.get(url, headers=headers)
```
