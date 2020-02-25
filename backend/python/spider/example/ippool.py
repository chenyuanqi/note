# -*- coding:utf-8 -*-

# IP 地址取自国内髙匿代理 IP 网站：http://www.xicidaili.com/nn/
# 仅仅爬取首页 IP 地址就足够一般使用
from bs4 import BeautifulSoup
import requests
import random


class IpPool:
    def __init__(self):
        self.url = 'http://www.xicidaili.com/nn/'
        self.ip_list = []

    def get_ip_list(self, headers):
        web_data = requests.get(self.url, headers=headers)
        soup = BeautifulSoup(web_data.text, 'lxml')
        ips = soup.find_all('tr')
        for i in range(1, len(ips)):
            ip_info = ips[i]
            tds = ip_info.find_all('td')
            ip = tds[1].text + ':' + tds[2].text
            if self.validate_ip(ip):
                self.ip_list.append(ip)

        return self.ip_list

    @staticmethod
    def validate_ip(ip, protocol='http'):
        url = "https://www.baidu.com/"
        try:
            proxy_host = protocol+"://"+ip
            result = requests.get(url, proxies={protocol: proxy_host}, timeout=3)
        finally:
            if result.status_code == 200:
                return True
            else:
                return False

    def get_random_ip(self):
        proxy_list = []
        for ip in self.ip_list:
            proxy_list.append('http://' + ip)
        proxy_ip = random.choice(proxy_list)
        proxies = {'http': proxy_ip}

        return proxies

    def get_proxies(self):
        headers = {
            'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36'
        }
        self.get_ip_list(headers=headers)

        return self.get_random_ip()


if __name__ == '__main__':
    pool = IpPool()
    proxy = pool.get_proxies()
    print(proxy)

