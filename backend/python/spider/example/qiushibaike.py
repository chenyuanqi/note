# -*- coding:utf-8 -*-

import re
from fake_useragent import UserAgent
import requests
from requests.exceptions import ReadTimeout, ConnectionError, RequestException


class QuishibaikeSpider:
    def __init__(self):
        self.page = 1
        self.uri = "http://www.qiushibaike.com/hot/page/"

        # 是否继续运行下一页
        self.enable = False

    def get_page(self, page):
        try:
            ua = UserAgent()
            headers = {'User-Agent': ua.random}
            url = self.uri + str(page)
            response = requests.get(url, headers=headers, timeout=5)

            return response.content.decode("utf-8")
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

    def get_page_items(self):
        content = self.get_page(self.page)
        self.page += 1
        if not content:
            print("页面加载失败....")

            return None
        pattern = re.compile('<div.*?content">.*?<span.*?>(.*?)</span>.*?<div.*?', re.S)
        items = re.findall(pattern, content)

        for item in items:
            # 是否含有图片
            have_img = re.search("img", item)
            # 如果不含有图片，把它加入 list 中
            if not have_img:
                replace_br = re.compile('<br/>')
                text = re.sub(replace_br, "\n", item)
                print("{}\n".format(text.strip()))

    def next_page(self):
            self.get_page_items()
            # 每当输入回车一次，判断一下是否要加载新页面
            is_quit = input("是否继续读取内容（输入 Q 退出）")
            # 如果输入 Q 则程序结束
            if is_quit.upper() == "Q":
                self.enable = False

                return False

    def start(self):
        print("正在读取糗事百科...")
        self.enable = True
        # 先加载一页内容
        self.get_page_items()
        while self.enable:
            self.next_page()


def main():
    spider = QuishibaikeSpider()
    spider.start()


if __name__ == '__main__':
    main()

