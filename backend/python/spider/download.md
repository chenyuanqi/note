
### Python 下载资源
小文件下载。  
比如下载一张图片：
```python
import requests

image_url = "https://www.python.org/static/community_logos/python-logo-master-v3-TM.png"
r = requests.get(image_url) 
with open("python_logo.png",'wb') as f:
    f.write(r.content)
```

大文件下载。  
如果文件比较大的话，那么下载下来的文件先放在内存中，内存还是比较有压力的。所以为了防止内存不够用的现象出现，我们要想办法把下载的文件分块写到磁盘中。
```python
import requests

file_url = "http://codex.cs.yale.edu/avi/db-book/db4/slide-dir/ch1-2.pdf"
r = requests.get(file_url, stream=True)
with open("python.pdf", "wb") as pdf:
    for chunk in r.iter_content(chunk_size=1024):
        if chunk:
            pdf.write(chunk)
```

批量文件下载。  
思路很简单，首先读取网页的内容，再从网页中抽取链接信息，比如通过 a 标签，然后再从抽取出的链接中过滤出我们想要的链接，比如我们只想下载 MP4 文件，那么我们可以通过文件名过滤所有链接：  
```python
import requests
from bs4 import BeautifulSoup


archive_url = "http://www-personal.umich.edu/~csev/books/py4inf/media/"


def get_video_links():
    r = requests.get(archive_url)
    soup = BeautifulSoup(r.content, 'html5lib')
    links = soup.findAll('a')
    video_links = [archive_url + link['href'] for link in links if link['href'].endswith('mp4')]

    return video_links


def download_video_series(video_links):
    for link in video_links:
        file_name = link.split('/')[-1]

        print("Downloading file:%s" % file_name)
        r = requests.get(link, stream=True)

        # download started
        with open(file_name, 'wb') as f:
            for chunk in r.iter_content(chunk_size=1024 * 1024):
                if chunk:
                    f.write(chunk)

        print("%s downloaded!\n" % file_name)


    print("All videos downloaded!")

    return


if __name__ == "__main__":
    video_links = get_video_links()
    download_video_series(video_links)
```

显示下载进度。  
```python
# 请求关键参数：stream=True
# 默认情况下，当你进行网络请求后，响应体会立即被下载。你可以通过 stream 参数覆盖这个行为，推迟下载响应体直到访问 Response.content 属性。
tarball_url = 'https://github.com/kennethreitz/requests/tarball/master'
r = requests.get(tarball_url, stream=True)

# 此时仅有响应头被下载下来了，连接保持打开状态，因此允许我们根据条件获取内容
if int(r.headers['content-length']) < TOO_LONG:
  content = r.content

# 进一步使用 Response.iter_content 和 Response.iter_lines 方法来控制工作流，或者以 Response.raw 从底层 urllib3 的 urllib3.HTTPResponse
# 保持活动状态（持久连接）归功于 urllib3，同一会话内的持久连接是完全自动处理的，同一会话内发出的任何请求都会自动复用恰当的连接
from contextlib import closing
with closing(requests.get('http://httpbin.org/get', stream=True)) as r:
    # Do things with the response here.
```
`注意：只有当响应体的所有数据被读取完毕时，连接才会被释放到连接池；所以，确保将 stream 设置为 False 或读取 Response 对象的 content 属性。`


下载文件并显示进度条。  
```python
with closing(requests.get(self.url(), stream=True)) as response:
    chunk_size = 1024 # 单次请求最大值
    content_size = int(response.headers['content-length']) # 内容体总大小
    progress = ProgressBar(self.file_name(), total=content_size,
                                     unit="KB", chunk_size=chunk_size, run_status="正在下载", fin_status="下载完成")
    with open(file_name, "wb") as file:
       for data in response.iter_content(chunk_size=chunk_size):
           file.write(data)
           progress.refresh(count=len(data))

# 进度条的实现
class ProgressBar(object):
    def __init__(self, title,
                 count=0.0,
                 run_status=None,
                 fin_status=None,
                 total=100.0,
                 unit='', sep='/',
                 chunk_size=1.0):
        super(ProgressBar, self).__init__()
        self.info = "【%s】%s %.2f %s %s %.2f %s"
        self.title = title
        self.total = total
        self.count = count
        self.chunk_size = chunk_size
        self.status = run_status or ""
        self.fin_status = fin_status or " " * len(self.status)
        self.unit = unit
        self.seq = sep

    def __get_info(self):
        # 【名称】状态 进度 单位 分割线 总数 单位
        _info = self.info % (self.title, self.status,
                             self.count/self.chunk_size, self.unit, self.seq, self.total/self.chunk_size, self.unit)
        return _info

    def refresh(self, count=1, status=None):
        self.count += count
        # if status is not None:
        self.status = status or self.status
        # 将结束符改为 “\r”，输出完成之后，光标会回到行首，并不换行
        end_str = "\r"
        if self.count >= self.total:
        	# 结束符也可以使用 “\d”，为退格符，光标回退一格，可以使用多个，按需求回退
        	# 在结束这一行输出时，将结束符改回 “\n” 或者不指定使用默认
            end_str = '\n'
            self.status = status or self.fin_status
        print(self.__get_info(), end=end_str)
```

