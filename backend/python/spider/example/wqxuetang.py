# -*- coding:utf-8 -*-

from selenium import webdriver
from webdriver_manager.chrome import ChromeDriverManager
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.action_chains import ActionChains
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.support.ui import Select
from bs4 import BeautifulSoup
import pdfkit
import time
import os

class Wqxuetang:
    def __init__(self):
        self.account = 'xxxx'
        self.password = 'xxxx'
        self.output_path = os.getcwd() + os.path.sep + 'pdf'
        if not os.path.exists(self.output_path):
            os.makedirs(output_path)

    def get_content(self):
        content_list = []
        try:
            # option = webdriver.ChromeOptions()
            # option.add_argument('headless')
            # driver = webdriver.Chrome(ChromeDriverManager().install(), options=option)
            driver = webdriver.Chrome(ChromeDriverManager().install())
            driver.get(url)

            # 显示等待
            wait = WebDriverWait(driver, 10)
            # 显式等待指定某个条件，然后设置最长等待时间。如果在这个时间还没有找到元素，那么便会抛出异常
            wait.until(EC.presence_of_element_located((By.ID, 'pagebox')))

            # 获取当前窗口
            # current_window = driver.current_window_handle

            # 根据 class 查找按钮
            button = driver.find_element_by_class_name('goBranch')
            button.click()
            time.sleep(3)

            # 切换窗口
            # all_windows = driver.window_handles
            # print(all_windows)
            # for window in all_windows:
            #     if window != current_window:
            #         driver.switch_to.window(window)
            # print(driver.page_source)

            # 登录页
            login = driver.find_element_by_id('mobilelogin-form-link')
            login.click()
            time.sleep(6)
            # 填写账号
            phone = driver.find_element_by_name('account')
            phone.send_keys(self.account)
            time.sleep(6)
            # 填写密码
            password = driver.find_element_by_name('password')
            password.send_keys(self.password)
            time.sleep(6)
            # 点击按钮
            buttons = driver.find_elements_by_class_name('btn-block')
            buttons[0].click()
            time.sleep(10)

            # 滚动页面
            for i in range(page_count):
                time.sleep(6)
                ActionChains(driver).key_down(Keys.DOWN).perform()

            # 页面响应内容
            html = driver.page_source
            bs4 = BeautifulSoup(html, 'lxml')
            content_list = bs4.select('#pagebox')
            time.sleep(3)
        finally:
            # 关闭浏览器
            driver.close()

            return content_list[0] if content_list else ''

    def export_pdf(self, content):
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
            {content}
            </body>
            </html>
            '''
            # windows 下需要配置此项
            config = pdfkit.configuration(wkhtmltopdf='C:/Program Files/wkhtmltopdf/bin/wkhtmltopdf.exe')
            pdfkit.from_string(html, self.output_path + os.path.sep + f'{title}.pdf', configuration=config)
        except Exception as e:
            # 导出存在异常
            print(e)

    def download(self):
        content = self.get_content()
        if content:
            self.export_pdf(content)


def main():
    wq = Wqxuetang()
    wq.download()


if __name__ == '__main__':
    url = 'https://lib-nuanxin.wqxuetang.com/read/pdf/3187145'
    title = '数据结构教程'
    page_count = 471
    # 执行爬取
    main()
