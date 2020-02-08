# -*- coding:utf-8 -*-
import os
import time
import urllib
from urllib import request
from lxml import etree


def get_url(url):
    """
        1、构造请求代理 ip 网站链接

        :param url 国内高匿代理的链接
        :return 生成要爬取目标网址的链接
    """
    url_list = []
    for i in range(1,100):
        url_new = url + str(i)
        url_list.append(url_new)

    return url_list


def get_content(url):
    """
        2、获取网页内容

        :param url 目标网站链接
        :return 网页内容
    """
    user_agent = 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.22 Safari/537.36 SE 2.X MetaSr 1.0'
    headers = {'User-Agent': user_agent}
    req = urllib.request.Request(url=url, headers=headers)
    res = urllib.request.urlopen(req)
    content = res.read()

    return content.decode('utf-8')


def get_info(content):
    """
        3、提取网页中 ip 地址和端口号信息

        :param content 接收从 get_content 函数传来的网页内容
        :return 使用 etree 解析出 ip 和端口号，将端口号和 ip 写入 data
    """
    datas_ip = etree.HTML(content).xpath('//table[contains(@id,"ip_list")]/tr/td[2]/text()')
    datas_port = etree.HTML(content).xpath('//table[contains(@id,"ip_list")]/tr/td[3]/text()')
    with open("data.txt", "w") as fd:
        for i in range(0,len(datas_ip)):
            out = u""
            out += u"" + datas_ip[i]
            out += u":" + datas_port[i]
            fd.write(out + u"\n") # 所有ip和端口号写入data文件


def verify_ip(ip, port):
    """
        4、验证ip有效性

        :param ip
        :param port
        :return 使用 ProxyHandler 建立代理，使用代理 ip 访问某网址，查看是否得到响应；如数据有效，则保存到 data2.txt 文件
    """
    user_agent ='Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.22 Safari/537.36 SE 2.X MetaSr 1.0'
    headers = {'User-Agent':user_agent}
    proxy = {'http':'http://%s:%s'%(ip,port)}
    print(proxy)

    proxy_handler = urllib.request.ProxyHandler(proxy)
    opener = urllib.request.build_opener(proxy_handler)
    urllib.request.install_opener(opener)

    test_url = "https://www.baidu.com/"
    req = urllib.request.Request(url=test_url,headers=headers)
    time.sleep(6)
    try:
        res = urllib.request.urlopen(req)
        time.sleep(3)
        content = res.read()
        if content:
            print('that is ok')
            with open("data2.txt", "a") as fd:       # 有效ip保存到 data2 文件中
                fd.write(ip + u":" + port)
                fd.write("\n")
        else:
            print('its not ok')
    except urllib.request.URLError as e:
        print(e.reason)


if __name__ == '__main__':
    url = 'http://www.xicidaili.com/nn/'
    url_list = get_url(url)
    for i in url_list:
        print(i)
        content = get_content(i)
        time.sleep(3)
        get_info(content)

    with open("dali.txt", "r") as fd:
        datas = fd.readlines()
        for data in datas:
            print(data.split(u":")[0])
            # print('%d : %d'%(out[0],out[1]))
            verify_ip(data.split(u":")[0],data.split(u":")[1])