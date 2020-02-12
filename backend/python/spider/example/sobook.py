# -*- coding:utf-8 -*-

from fake_useragent import UserAgent
import requests
from requests.exceptions import ReadTimeout, ConnectionError, RequestException
from bs4 import BeautifulSoup
from selenium import webdriver
from webdriver_manager.chrome import ChromeDriverManager
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.by import By
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.support.wait import WebDriverWait


class SobookSpider:
    def __init(self):
        pass

    def search_jiumo(self):
        url = 'https://www.jiumodiary.com'
        try:
            option = webdriver.ChromeOptions()
            option.add_argument('headless')
            driver = webdriver.Chrome(ChromeDriverManager().install(), options=option)
            driver.get(url)
            # 根据id查找元素，获取到查询框
            input = driver.find_element_by_id('SearchWord')
            # 向查询框中输入关键词
            input.send_keys(keyword)
            # 模拟回车
            input.send_keys(Keys.ENTER)
            # 显示等待
            wait = WebDriverWait(driver, 3)
            # 显式等待指定某个条件，然后设置最长等待时间。如果在这个时间还没有找到元素，那么便会抛出异常
            wait.until(EC.presence_of_element_located((By.XPATH, '//ul/li')))
            # 页面响应内容
            html = driver.page_source
            bs4 = BeautifulSoup(html, 'lxml')
            content_list = bs4.select('#result-ul li')
            print('鸠摩搜索结果如下：')
            for i, content in enumerate(content_list):
                link = content.select('a')
                print(i + 1, link[0].get_text(), link[0]['href'])
        finally:
            # 关闭浏览器
            driver.close()

    def search_zqbook(self):
        url = 'https://www.zqbook.top'
        try:
            option = webdriver.ChromeOptions()
            option.add_argument('headless')
            driver = webdriver.Chrome(ChromeDriverManager().install(), options=option)
            driver.get(url)
            # 根据 css 查找元素，获取到查询框
            input = driver.find_element_by_css_selector('.input input')
            # 向查询框中输入关键词
            input.send_keys(keyword)
            # 模拟回车
            input.send_keys(Keys.ENTER)
            # 显示等待
            wait = WebDriverWait(driver, 3)
            # 显式等待指定某个条件，然后设置最长等待时间
            wait.until(EC.presence_of_element_located((By.ID, 'books')))
            # 页面响应内容
            html = driver.page_source
            bs4 = BeautifulSoup(html, 'lxml')
            content_list = bs4.select('#books .item')
            print('智奇搜书结果如下：')
            for i, content in enumerate(content_list):
                link = content.select('a')
                print(i + 1, link[0].h3.get_text(), link[0]['href'])
        finally:
            # 关闭浏览器
            driver.close()

    def search_epubw(self):
        url = 'https://epubw.com/'
        try:
            option = webdriver.ChromeOptions()
            option.add_argument('headless')
            driver = webdriver.Chrome(ChromeDriverManager().install(), options=option)
            driver.get(url)
            # 根据 css 查找元素，获取到查询按钮
            button = driver.find_element_by_css_selector('.site-nav .navto-search')
            # 点击按钮
            button.click()
            # 输入关键字
            input = driver.find_element_by_name('s')
            input.send_keys(keyword)
            input.send_keys(Keys.ENTER)
            # 显示等待
            wait = WebDriverWait(driver, 3)
            # 显式等待指定某个条件，然后设置最长等待时间
            wait.until(EC.presence_of_element_located((By.CLASS_NAME, 'container')))
            # 页面响应内容
            html = driver.page_source
            bs4 = BeautifulSoup(html, 'lxml')
            not_found = bs4.select('.f404')
            if not_found:
                print('epubw 搜索无结果...')
                return False

            content_list = bs4.select('.content .row article')
            print('epubw 搜索结果如下：')
            for i, content in enumerate(content_list):
                link = content.select('a')
                print(i + 1, content.get_text().strip().replace('\n', '').replace('\r', ''), link[0]['href'])
        finally:
            # 关闭浏览器
            driver.close()

    def search_d4j(self):
        url = 'https://www.d4j.cn/'
        try:
            option = webdriver.ChromeOptions()
            option.add_argument('headless')
            driver = webdriver.Chrome(ChromeDriverManager().install(), options=option)
            # driver = webdriver.Chrome(ChromeDriverManager().install())
            driver.get(url)
            # 输入关键字
            input = driver.find_element_by_id('search')
            input.send_keys(keyword)
            input.send_keys(Keys.ENTER)
            # 显示等待
            wait = WebDriverWait(driver, 6)
            # 显式等待指定某个条件，然后设置最长等待时间
            wait.until(EC.presence_of_element_located((By.CLASS_NAME, 'kratos-hentry')))
            # 页面响应内容
            html = driver.page_source
            bs4 = BeautifulSoup(html, 'lxml')
            content_list = bs4.select('#main article')
            print('三秋书屋搜索结果如下：')
            for i, content in enumerate(content_list):
                link = content.select('h2 a')
                if link:
                    print(i + 1, link[0].get_text().strip(), link[0]['href'])
                else:
                    title = content.select('h1')
                    if title:
                        print(i + 1, title[0].get_text().strip(), driver.current_url)
        finally:
            # 关闭浏览器
            driver.close()

    def search_lorefree(self):
        url = 'https://ebook.lorefree.com/'
        try:
            option = webdriver.ChromeOptions()
            option.add_argument('headless')
            driver = webdriver.Chrome(ChromeDriverManager().install(), options=option)
            driver.get(url)
            # 输入关键字
            input = driver.find_element_by_name('s')
            input.send_keys(keyword)
            input.send_keys(Keys.ENTER)
            # 显示等待
            wait = WebDriverWait(driver, 6)
            # 显式等待指定某个条件，然后设置最长等待时间
            wait.until(EC.presence_of_element_located((By.CLASS_NAME, 'body-content')))
            # 页面响应内容
            html = driver.page_source
            bs4 = BeautifulSoup(html, 'lxml')
            content_list = bs4.select('.body-content .row')
            print('LoreFree 搜索结果如下：')
            for i, content in enumerate(content_list):
                link = content.select('.book-card .book-content a')
                if link:
                    print(i + 1, link[0].get_text().strip(), url + link[0]['href'])
        finally:
            # 关闭浏览器
            driver.close()

    def start(self):
        print("鸠摩搜索ing...")
        self.search_jiumo()

        print("智奇搜书ing...")
        self.search_zqbook()

        print("epubw 搜索ing...")
        self.search_epubw()

        print("三秋书屋搜索ing...")
        self.search_d4j()

        print("LoreFree 搜索ing...")
        self.search_lorefree()


def main():
    spider = SobookSpider()
    spider.start()


if __name__ == '__main__':
    keyword = input('请输入关键词检索相关书籍：')
    main()
