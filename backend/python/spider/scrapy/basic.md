
### Scrapy 概述
Scrapy['skræpi:] 是 Python 开发的一个非常流行的网络爬虫框架，可以用来抓取 Web 站点并从页面中提取结构化的数据，被广泛的用于数据挖掘、数据监测和自动化测试等领域。下图展示了 Scrapy 的基本架构，其中包含了主要组件和系统的数据处理流程（图中带数字的红色箭头）。  
![scrapy 架构](../../../../others/static/images/scrapy-architecture.png)

**组件**  
1、Scrapy 引擎（Engine）：Scrapy 引擎是用来控制整个系统的数据处理流程。   
2、调度器（Scheduler）：调度器从 Scrapy 引擎接受请求并排序列入队列，并在 Scrapy 引擎发出请求后返还给它们。  
3、下载器（Downloader）：下载器的主要职责是抓取网页并将网页内容返还给蜘蛛（Spiders）。  
4、蜘蛛（Spiders）：蜘蛛是有 Scrapy 用户自定义的用来解析网页并抓取特定 URL 返回的内容的类，每个蜘蛛都能处理一个域名或一组域名，简单的说就是用来定义特定网站的抓取和解析规则。  
5、条目管道（Item Pipeline）：条目管道的主要责任是负责处理有蜘蛛从网页中抽取的数据条目，它的主要任务是清理、验证和存储数据。当页面被蜘蛛解析后，将被发送到条目管道，并经过几个特定的次序处理数据。每个条目管道组件都是一个 Python 类，它们获取了数据条目并执行对数据条目进行处理的方法，同时还需要确定是否需要在条目管道中继续执行下一步或是直接丢弃掉不处理。条目管道通常执行的任务有：清理 HTML 数据、验证解析到的数据（检查条目是否包含必要的字段）、检查是不是重复数据（如果重复就丢弃）、将解析到的数据存储到数据库（关系型数据库或 NoSQL 数据库）中。  
6、中间件（Middlewares）：中间件是介于 Scrapy 引擎和其他组件之间的一个钩子框架，主要是为了提供自定义的代码来拓展 Scrapy 的功能，包括下载器中间件和蜘蛛中间件。  

**数据处理流程**  
Scrapy 的整个数据处理流程由 Scrapy 引擎进行控制，通常的运转流程包括以下的步骤：  
1、引擎询问蜘蛛需要处理哪个网站，并让蜘蛛将第一个需要处理的 URL 交给它。  
2、引擎让调度器将需要处理的 URL 放在队列中。  
3、引擎从调度那获取接下来进行爬取的页面。  
4、调度将下一个爬取的 URL 返回给引擎，引擎将它通过下载中间件发送到下载器。  
5、当网页被下载器下载完成以后，响应内容通过下载中间件被发送到引擎；如果下载失败了，引擎会通知调度器记录这个 URL，待会再重新下载。  
6、引擎收到下载器的响应并将它通过蜘蛛中间件发送到蜘蛛进行处理。  
7、蜘蛛处理响应并返回爬取到的数据条目，此外还要将需要跟进的新的 URL 发送给引擎。  
8、引擎将抓取到的数据条目送入条目管道，把新的 URL 发送给调度器放入队列中。  

上述操作中的 2-8 步会一直重复直到调度器中没有需要请求的 URL，爬虫停止工作。  

### Scrapy 基本使用
先创建虚拟环境并在虚拟环境下使用 pip 安装 scrapy。  
```bash
# 安装
pip install scrapy
# 创建一个 Scrapy 项目
scrapy startproject douban
cd douban
# alias tree="find . -print | sed -e 's;[^/]*/;|____;g;s;____|; |;g'"
tree
# .
# |____ scrapy.cfg # Scrapy 项目的配置文件，其内定义了项目的配置文件路径、部署相关信息等内容
# |____ douban
# | |____ spiders # 包含一个个 Spider 的实现，每个 Spider 都有一个文件
# | | |____ __init__.py
# | | |____ __pycache__
# | |____ __init__.py
# | |____ __pycache__
# | |____ middlewares.py # 定义 Spider Middlewares 和 Downloader Middlewares 的实现
# | |____ settings.py # 定义项目的全局配置
# | |____ items.py # 定义 Item 数据结构，所有的 Item 的定义都可以放这里
# | |____ pipelines.py # 定义 Item Pipeline 的实现，所有的 Item Pipeline 的实现都可以放这里
```
根据数据处理流程，基本上需要我们做的有以下几件事情：  
1、在 items.py 文件中定义字段，这些字段用来保存数据，方便后续的操作。  
```python
# -*- coding: utf-8 -*-

# Define here the models for your scraped items
#
# See documentation in:
# https://doc.scrapy.org/en/latest/topics/items.html

import scrapy


class DoubanItem(scrapy.Item):

    name = scrapy.Field()
    year = scrapy.Field()
    score = scrapy.Field()
    director = scrapy.Field()
    classification = scrapy.Field()
    actor = scrapy.Field()
```

2、在 spiders 文件夹中编写自己的爬虫。  
通过 Scrapy 提供的爬虫模板创建了 Spider，执行如下命令  
```bash
scrapy genspider movie movie.douban.com --template=crawl
```

生成的爬虫 spiders/movie.py 的代码如下  
```python
# -*- coding: utf-8 -*-
import scrapy
from scrapy.selector import Selector
from scrapy.linkextractors import LinkExtractor
from scrapy.spiders import CrawlSpider, Rule

from douban.items import DoubanItem


# 在 Scrapy 框架中，我们自定义的蜘蛛都继承自 scrapy.spiders.Spider
class MovieSpider(CrawlSpider):
    name = 'movie' # 爬虫的名字
    allowed_domains = ['movie.douban.com'] # 允许爬取的域名，不在此范围的链接不会被跟进爬取
    start_urls = ['https://movie.douban.com/top250'] # 起始 URL 列表，当我们没有重写 start_requests() 方法时，就会从这个列表开始爬取
    #  LinkExtractor 对象会自动完成对新的链接的解析，该对象中有一个名为 extract_link 的回调方法
    rules = (
        Rule(LinkExtractor(allow=(r'https://movie.douban.com/top250\?start=\d+.*'))),
        Rule(LinkExtractor(allow=(r'https://movie.douban.com/subject/\d+')), callback='parse_item'),
    )
    # 其他属性配置
    # custom_settings：用来存放蜘蛛专属配置的字典，这里的设置会覆盖全局的设置
    # cawler：由 from_crawler () 方法设置的和蜘蛛对应的 Crawler 对象，Crawler 对象包含了很多项目组件，利用它我们可以获取项目的配置信息，如调用 crawler.settings.get () 方法
    # settings：用来获取爬虫全局设置的变量
    # 

    def parse_item(self, response):
        sel = Selector(response)
        item = DoubanItem()
        # Scrapy 支持用 XPath 语法和 CSS 选择器进行数据解析，对应的方法分别是 xpath 和 css
        item['name']=sel.xpath('//*[@id="content"]/h1/span[1]/text()').extract()
        item['year']=sel.xpath('//*[@id="content"]/h1/span[2]/text()').re(r'\((\d+)\)')
        item['score']=sel.xpath('//*[@id="interest_sectl"]/div/p[1]/strong/text()').extract()
        item['director']=sel.xpath('//*[@id="info"]/span[1]/a/text()').extract()
        item['classification']= sel.xpath('//span[@property="v:genre"]/text()').extract()
        item['actor']= sel.xpath('//*[@id="info"]/span[3]/a[1]/text()').extract()
        return item

    # 其他方法配置
    # start_requests ()：此方法用于生成初始请求，它返回一个可迭代对象。该方法默认是使用 GET 请求访问起始 URL，如果起始 URL 需要使用 POST 请求来访问就必须重写这个方法
    # parse()：当 Response 没有指定回调函数时，该方法就会被调用，它负责处理 Response 对象并返回结果，从中提取出需要的数据和后续的请求，该方法需要返回类型为 Request 或 Item 的可迭代对象（生成器当前也包含在其中，因此根据实际需要可以用 return 或 yield 来产生返回值）
    # closed ()：当蜘蛛关闭时，该方法会被调用，通常用来做一些释放资源的善后操作
```

到这里，我们已经可以通过下面的命令让爬虫运转起来。  
```bash
scrapy crawl movie

# 可以在控制台看到爬取到的数据，如果想将这些数据保存到文件中，可以通过 -o 参数来指定文件名
# Scrapy 支持我们将爬取到的数据导出成 JSON、CSV、XML、pickle、marshal 等格式
scrapy crawl moive -o result.json
```

3、在 pipelines.py 中完成对数据进行持久化的操作。  
利用 Pipeline 我们可以完成以下操作：清理 HTML 数据，验证爬取的数据；丢弃重复的不必要的内容；将爬取的结果进行持久化操作。  
```python
# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://doc.scrapy.org/en/latest/topics/item-pipeline.html
import pymongo

from scrapy.exceptions import DropItem
from scrapy.conf import settings
from scrapy import log


class DoubanPipeline(object):

    def __init__(self):
        connection = pymongo.MongoClient(settings['MONGODB_SERVER'], settings['MONGODB_PORT'])
        db = connection[settings['MONGODB_DB']]
        self.collection = db[settings['MONGODB_COLLECTION']]

    def process_item(self, item, spider):
        #Remove invalid data
        valid = True
        for data in item:
          if not data:
            valid = False
            raise DropItem("Missing %s of blogpost from %s" %(data, item['url']))
        if valid:
        #Insert data into database
            new_moive=[{
                "name":item['name'][0],
                "year":item['year'][0],
                "score":item['score'],
                "director":item['director'],
                "classification":item['classification'],
                "actor":item['actor']
            }]
            self.collection.insert(new_moive)
            log.msg("Item wrote to MongoDB database %s/%s" %
            (settings['MONGODB_DB'], settings['MONGODB_COLLECTION']),
            level=log.DEBUG, spider=spider) 
        return item
```

4、修改 settings.py 文件对项目进行配置。  
```python
# -*- coding: utf-8 -*-

# Scrapy settings for douban project
#
# For simplicity, this file contains only settings considered important or
# commonly used. You can find more settings consulting the documentation:
#
#     https://doc.scrapy.org/en/latest/topics/settings.html
#     https://doc.scrapy.org/en/latest/topics/downloader-middleware.html
#     https://doc.scrapy.org/en/latest/topics/spider-middleware.html

BOT_NAME = 'douban'

SPIDER_MODULES = ['douban.spiders']
NEWSPIDER_MODULE = 'douban.spiders'


# Crawl responsibly by identifying yourself (and your website) on the user-agent
USER_AGENT = 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_3) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.54 Safari/536.5'

# Obey robots.txt rules
ROBOTSTXT_OBEY = True

# Configure maximum concurrent requests performed by Scrapy (default: 16)
# CONCURRENT_REQUESTS = 32

# Configure a delay for requests for the same website (default: 0)
# See https://doc.scrapy.org/en/latest/topics/settings.html#download-delay
# See also autothrottle settings and docs
DOWNLOAD_DELAY = 3
RANDOMIZE_DOWNLOAD_DELAY = True
# The download delay setting will honor only one of:
# CONCURRENT_REQUESTS_PER_DOMAIN = 16
# CONCURRENT_REQUESTS_PER_IP = 16

# Disable cookies (enabled by default)
COOKIES_ENABLED = True

MONGODB_SERVER = '120.77.222.217'
MONGODB_PORT = 27017
MONGODB_DB = 'douban'
MONGODB_COLLECTION = 'movie'

# Disable Telnet Console (enabled by default)
# TELNETCONSOLE_ENABLED = False

# Override the default request headers:
# DEFAULT_REQUEST_HEADERS = {
#   'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8',
#   'Accept-Language': 'en',
# }

# Enable or disable spider middlewares
# See https://doc.scrapy.org/en/latest/topics/spider-middleware.html
# SPIDER_MIDDLEWARES = {
#    'douban.middlewares.DoubanSpiderMiddleware': 543,
# }

# Enable or disable downloader middlewares
# See https://doc.scrapy.org/en/latest/topics/downloader-middleware.html
# DOWNLOADER_MIDDLEWARES = {
#    'douban.middlewares.DoubanDownloaderMiddleware': 543,
# }

# Enable or disable extensions
# See https://doc.scrapy.org/en/latest/topics/extensions.html
# EXTENSIONS = {
#    'scrapy.extensions.telnet.TelnetConsole': None,
# }

# Configure item pipelines
# See https://doc.scrapy.org/en/latest/topics/item-pipeline.html
ITEM_PIPELINES = {
    'douban.pipelines.DoubanPipeline': 400,
}

LOG_LEVEL = 'DEBUG'

# Enable and configure the AutoThrottle extension (disabled by default)
# See https://doc.scrapy.org/en/latest/topics/autothrottle.html
#AUTOTHROTTLE_ENABLED = True
# The initial download delay
#AUTOTHROTTLE_START_DELAY = 5
# The maximum download delay to be set in case of high latencies
#AUTOTHROTTLE_MAX_DELAY = 60
# The average number of requests Scrapy should be sending in parallel to
# each remote server
#AUTOTHROTTLE_TARGET_CONCURRENCY = 1.0
# Enable showing throttling stats for every response received:
#AUTOTHROTTLE_DEBUG = False

# Enable and configure HTTP caching (disabled by default)
# See https://doc.scrapy.org/en/latest/topics/downloader-middleware.html#httpcache-middleware-settings
HTTPCACHE_ENABLED = True
HTTPCACHE_EXPIRATION_SECS = 0
HTTPCACHE_DIR = 'httpcache'
HTTPCACHE_IGNORE_HTTP_CODES = []
HTTPCACHE_STORAGE = 'scrapy.extensions.httpcache.FilesystemCacheStorage'
```

### Scrapy 进阶应用
**Selector 的用法**  

**Spider 的用法**

**DownloaderMiddleware 的用法**  

**对接 Selenium**

### Scrapy 分布式
Scrapy 分布式实现，大致需要如下步骤：  
1、安装 Scrapy-Redis。  
2、配置 Redis 服务器。  
3、修改配置文件。  

- SCHEDULER = 'scrapy_redis.scheduler.Scheduler'  
- DUPEFILTER_CLASS = 'scrapy_redis.dupefilter.RFPDupeFilter'  
- REDIS_HOST = '1.2.3.4'  
- REDIS_PORT = 6379  
- REDIS_PASSWORD = '1qaz2wsx'  
- SCHEDULER_QUEUE_CLASS = 'scrapy_redis.queue.FifoQueue'  
- SCHEDULER_PERSIST = True（通过持久化支持接续爬取）  
- SCHEDULER_FLUSH_ON_START = True（每次启动时重新爬取）  

4、Scrapyd 分布式部署  
安装 Scrapyd，然后修改配置文件 /etc/scrapyd/scrapyd.conf。  
安装 Scrapyd-Client，将项目打包成 Egg 文件，将打包的 Egg 文件通过 addversion.json 接口部署到 Scrapyd 上。  
