# -*- coding:utf-8 -*-

from fake_useragent import UserAgent
import requests
from requests.exceptions import ReadTimeout, ConnectionError, RequestException


class BoxOfficeSpider:
    def __init__(self):
        self.url = "http://www.endata.com.cn/API/GetData.ashx"

    def get_items(self):
        try:
            ua = UserAgent()
            headers = {'User-Agent': ua.random}
            data = {'MethodName': 'BoxOffice_GetPcHomeList'}
            response = requests.post(self.url, data=data, headers=headers, timeout=5)

            return response.json()
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

    def start(self):
        result = self.get_items()
        if result and result['Status'] == 1:
            data = result['Data']['Table1']
            for item in data:
                print("No {}. {} 票房：{} 万，排片占比：{}%，上映天数：{}".format(item['tid'], item['MovieName'], item['boxoffice'], item['paipian'], item['releasedate']))


def main():
    spider = BoxOfficeSpider()
    spider.start()


if __name__ == '__main__':
    main()

