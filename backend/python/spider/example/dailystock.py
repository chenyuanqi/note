# -*- coding:utf-8 -*-

from fake_useragent import UserAgent
import requests
from requests.exceptions import ReadTimeout, ConnectionError, RequestException
import random
import json


class DailyStock:
    def __init__(self):
        # 股票代码列表
        self.stock_code_list = ['sh600000', 'sz300253']
        # 新浪财经：https://finance.sina.com.cn/stock/
        self.sina_uri = 'http://hq.sinajs.cn/list='
        # 腾讯财经：http://stockhtm.finance.qq.com/sstock/ggcx/000001.shtml
        self.tencent_uri = 'http://web.ifzq.gtimg.cn/appstock/app/fqkline/get?_var=kline_dayqfq&param={code},{type},,,{number},qfq&r={rand}'
        # 其他参考
        # 东方财富：http://www.eastmoney.com/

    @staticmethod
    def get_data(url):
        try:
            ua = UserAgent()
            headers = {'User-Agent': ua.random}
            response = requests.get(url, headers=headers, timeout=5)

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
        except:
            print("Something error")

    def get_real_time_data(self):
        # 构建 url
        url = self.sina_uri + ",".join(self.stock_code_list)
        # 获取实时数据
        data = self.get_data(url)
        if data:
            # 格式化
            stock_list = data.split('\n')
            stock_list.pop()
            for stock in stock_list:
                stock_data = stock.split('=', maxsplit=1)[1].split(',')
                print("{} 实时价格：{}，今开：{}，昨收：{}，最高：{}，最低：{}"
                      .format(stock_data[0][1:], stock_data[3], stock_data[1], stock_data[2], stock_data[4],
                              stock_data[5]))

    def get_k_line_data(self):
        for stock_code in self.stock_code_list:
            # 构建 url
            url = self.tencent_uri.format(code=stock_code, type='day', number=50, rand=random.random())
            # 获取 k 线数据
            data = self.get_data(url)
            if data:
                # 格式化
                stock_data = json.loads(data.split('=', maxsplit=1)[1])
                if stock_data['code'] == 0:
                    stock_name = stock_data['data'][stock_code]['qt'][stock_code][1]
                    print('{}【{}】 日 k 线数据如下：'.format(stock_name, stock_code))
                    stock_days = stock_data['data'][stock_code]['qfqday']
                    for stock_day in reversed(stock_days):
                        print("{}  开：{}，收：{}，高：{}，低：{}，成交量：{}".format(*stock_day))
                    print("#" * 29)


def main():
    spider = DailyStock()
    spider.get_k_line_data()


if __name__ == '__main__':
    main()
