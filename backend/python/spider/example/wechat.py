# -*- coding:utf-8 -*-

import os
import time
import re
from fake_useragent import UserAgent
import requests
from requests.exceptions import ReadTimeout, ConnectionError, RequestException
from bs4 import BeautifulSoup
import wechatsogou
import pdfkit
from PyPDF2 import PdfFileMerger, PdfFileReader


class WechatSpider:
    def __init__(self):
        self.ws_api = wechatsogou.WechatSogouAPI(captcha_break_time=3)
        self.output_path = os.getcwd() + os.path.sep + 'pdf'
        if not os.path.exists(self.output_path):
            os.makedirs(self.output_path)

    @staticmethod
    def get_items():
        try:
            ua = UserAgent()
            headers = {'User-Agent': ua.random}
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
        except:
            print("Something error")

    def export_pdf(self, article_url, title):
        # https://github.com/JazzCore/python-pdfkit/wiki/Installing-wkhtmltopdf
        try:
            title = self.format_title(title)
            content = self.ws_api.get_article_content(article_url)
        except:
            print("export error")

        try:
            html = f'''
            <!DOCTYPE html>
            <html lang="en">
            <head>
                <meta charset="UTF-8">
                <title>{title}</title>
            </head>
            <body>
            <h2 style="text-align: center;font-weight: 400;">{title}</h2>
            {content['content_html']}
            </body>
            </html>
            '''

            # windows 下需要配置此项
            config = pdfkit.configuration(wkhtmltopdf='C:/Program Files/wkhtmltopdf/bin/wkhtmltopdf.exe')
            pdfkit.from_string(html, self.output_path + os.path.sep + f'{title}.pdf', configuration=config)
        except Exception as e:
            # 导出存在异常
            print(e)

    def merge_pdf(self):
        # 获取所有 pdf 文件
        file_list = [f for f in os.listdir(self.output_path) if f.endswith('.pdf')]
        file_list.sort(key=lambda k: int(k.split('#')[0]))
        merger = PdfFileMerger(strict=False)
        for file_name in file_list:
            file_path = os.path.join(self.output_path, file_name)
            bookmark_name = os.path.basename(os.path.splitext(file_name)[0]).split('###')[1]
            # PdfReadError: Unexpected destination '/__WKANCHOR_2'
            # 屏蔽 pdf.py 这行：raise utils.PdfReadError("Unexpected destination %r" % dest)
            merger.append(PdfFileReader(open(file_path, 'rb')), bookmark=bookmark_name, import_bookmarks=True)

        with open(os.path.join(self.output_path, merge_name + '.pdf'), "wb") as file:
            merger.write(file)

    @staticmethod
    def format_title(title):
        # 去除标题中间的空格
        title = ''.join(title.split())
        # 去除标题特殊符号
        title = title.replace('“', "'").replace('”', "'").replace('，', ',')
        # 文件名不能包含 \/:*?<>
        title = title.replace('\\', '').replace('/', '').replace(':', '-').replace('*', '')\
            .replace('?', '').replace('<', '').replace('>', '')

        return title

    def from_url(self):
        number = 0
        content = self.get_items()
        bs4 = BeautifulSoup(content, 'lxml')
        url_list = bs4.select('#js_content p')

        for url_item in url_list:
            link_list = url_item.find_all('a')
            link_len = len(link_list)
            if link_len:
                # 链接多个，取最后一个
                if link_len > 1:
                    link_str = link_list[-1]
                else:
                    link_str = link_list[0]

                title = link_str.get_text()
                if not title:
                    continue

                prefix = url_item.select('p strong')
                if prefix:
                    title = prefix[0].get_text() + title
                number = number + 1
                title = str(number) + " ### " + title
                print(title, link_str['href'])
                # 执行导出
                self.export_pdf(link_str['href'], title)
                time.sleep(3)

    def from_html(self):
        # 关注公众号后，使用网页版查看历史消息并获取源码保存至 source.html
        with open("source.html", 'r', encoding='utf-8') as f:
            number = 0
            bs4 = BeautifulSoup(f, 'lxml')
            history_list = bs4.select('.weui_media_bd')

            for history in history_list:
                info = history.find('h4')
                if info:
                    title = info.get_text()
                    # 标题处理
                    if title.find('原创'):
                        title = title.replace('原创', '')
                    title = title.strip()
                    date = history.select('.weui_media_extra_info')
                    # 拼接日期
                    if date and title:
                        title = "【" + date[0].get_text().strip() + "】" + title
                    number = number + 1
                    title = str(number) + " ### " + title
                    # 执行导出
                    print(title, info['hrefs'])
                    self.export_pdf(info['hrefs'], title)
                    time.sleep(3)
                else:
                    print('非图文文章，继续~')

    def start(self):
        if mode == 'url' and url:
            self.from_url()
        elif mode == 'html':
            self.from_html()
        elif mode == 'merge':
            self.merge_pdf()
        elif mode == 'direct' and url and title:
            self.export_pdf(url, title)


def main():
    spider = WechatSpider()
    spider.start()


if __name__ == '__main__':
    mode = input("请选择模式（direct、url、html、merge）：")
    if mode == 'url':
        url = input('请输入要爬取的 url（请确保包含微信公众号汇总历史连接）：').strip()
    elif mode == 'html':
        if not os.path.exists("source.html"):
            print("请关注公众号后，使用网页版查看历史消息并获取源码保存至 source.html")
            exit()
    elif mode == 'merge':
        merge_name = input('请输入合并 pdf 的名字：')
    elif mode == 'direct':
        title = input('请输入要爬取的标题：').strip()
        url = input('请输入要爬取的 url：').strip()
    else:
        print('请重新输入 direct、url、html、merge')
        exit()

    main()
