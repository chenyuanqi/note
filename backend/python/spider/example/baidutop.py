# -*- coding:utf-8 -*-

from fake_useragent import UserAgent
import requests
from requests.exceptions import ReadTimeout, ConnectionError, RequestException
from bs4 import BeautifulSoup


class BaiduTopSpider:
    def __init__(self):
        self.url = "http://top.baidu.com/"

    def get_items(self):
        try:
            ua = UserAgent()
            headers = {'User-Agent': ua.random}
            response = requests.get(self.url, headers=headers, timeout=5)

            return response.content.decode("gbk")
        except ReadTimeout:
            # 超时异常
            print('Timeout')
        except ConnectionError:
            # 连接异常
            print('Connection error')
        except RequestException:
            # 请求异常
            print('Request Error')
        except Exception:
            print("Something error")

    def search_hot(self, content):
        bs4 = BeautifulSoup(content, 'lxml')
        search_list = bs4.select('ul.item-list')[7].select('li .item-hd a.list-title')
        for index, text in enumerate(search_list):
            print("{}.{}".format(index + 1, text.get_text()))

    def livelihood_hot(self, content):
        bs4 = BeautifulSoup(content, 'lxml')
        live_list = bs4.select('ul.item-list')[3].select('li .item-hd a.list-title')
        for index, text in enumerate(live_list):
            print("{}.{}".format(index + 1, text.get_text()))

    def realtime_hot(self, content):
        bs4 = BeautifulSoup(content, 'lxml')
        hot_list = bs4.select('#hot-list li>a.list-title')
        for index, text in enumerate(hot_list):
            print("{}.{}".format(index + 1, text.get_text()))

    def week_hot(self, content):
        bs4 = BeautifulSoup(content, 'lxml')
        week_list = bs4.select('ul.list')[1].select('li>a.list-title')
        for index, text in enumerate(week_list):
            print("{}.{}".format(index + 1, text.get_text()))

    def start(self):
        content = self.get_items()
        if content:
            print("实时热点：")
            self.realtime_hot(content)
            print("############################")
            print("7 日关注：")
            self.week_hot(content)
            print("############################")
            print("民生热点：")
            self.livelihood_hot(content)
            print("############################")
            print("热门搜索：")
            self.search_hot(content)


def main():
    spider = BaiduTopSpider()
    spider.start()


if __name__ == '__main__':
    main()

