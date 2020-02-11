
### 1、概述
Selenium [sə'liniəm]。  
Selenium Python bindings 使用非常简洁方便的 API 让你去使用像 Firefox, IE, Chrome, Remote 等等 这样的 Selenium WebDrivers（Selenium web 驱动器）。  

Selenium Python bindings 提供了一个简单的 API，让你使用 Selenium WebDriver 来编写功能 / 校验测试。 通过 Selenium Python 的 API，你可以非常直观的使用 Selenium WebDriver 的所有功能。  

[selenium 中文文档](https://selenium-python-zh.readthedocs.io/en/latest/index.html)  

```bash
pip install selenium
```

如果你想使用一个远程的 WebDriver，Selenium 服务是唯一的依赖。  
Selenium server 是一个 JAVA 工程，Java Runtime Environment (JRE) 1.6 或者更高的版本是推荐的运行环境。  
```bash
# 下载页：http://seleniumhq.org/download/
java -jar selenium-server-standalone-2.x.x.jar
```

### 2、基本使用
```python
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.support.wait import WebDriverWait

browser = webdriver.Chrome()
browser.get('https://www.taobao.com/')
# 显示等待 10s
wait = WebDriverWait(browser, 10)
# 等待直到元素加载出
input = wait.until(EC.presence_of_element_located((By.ID, 'q')))
# 等待直到元素可点击
button = wait.until(EC.element_to_be_clickable((By.CSS_SELECTOR, '.btn-search')))
print(input, button)


# 创建一个浏览器对象
browser = webdriver.Chrome()
try:
    # 开启一个浏览器并访问https://www.baidu.com
    browser.get('https://www.baidu.com')
    # 在打开的网页响应中根据id查找元素   获取到查询框
    input = browser.find_element_by_id('kw')
    # 向查询框中输入Python
    input.send_keys('Python')
    # 模拟回车
    input.send_keys(Keys.ENTER)
    # 显示等待， 等待10秒
    wait = WebDriverWait(browser, 10)
    # 显式等待指定某个条件，然后设置最长等待时间。如果在这个时间还没有找到元素，那么便会抛出异常
    wait.until(EC.presence_of_element_located((By.ID,'content_left')))
    # 输出当前的url
    print(browser.current_url)
    # 输出Cookies
    print(browser.get_cookies())
    # 输出页面响应内容
    print(browser.page_source)
finally:
    pass
    # 关闭浏览器
    browser.close()
```

**Selenium 声明浏览器对象**  
```python
from selenium import webdriver

browser = webdriver.Chrome()
browser = webdriver.Firefox()
browser = webdriver.Edge()
browser = webdriver.PhantomJS()
browser = webdriver.Safari()
```

**Selenium 查找元素**   
查找单个元素的方法：  

- find_element_by_name   通过 name 查找
- find_element_by_xpath  通过 xpath 查找
- find_element_by_link_text   通过链接查找
- find_element_by_partial_link_text    通过部分链接查找
- find_element_by_tag_name   通过标签名称查找
- find_element_by_class_name   通过类名查找
- find_element_by_css_selector  通过 css 选择武器查找

```python
from selenium import webdriver
from selenium.webdriver.common.by import By

# 申明一个浏览器对象
browser = webdriver.Chrome()
# 使用浏览器访问淘宝
browser.get('https://www.taobao.com')
# 在响应结果中通过id查找元素
input_first = browser.find_element_by_id('q')
# 在响应结果中通过css选择器查找元素
input_second = browser.find_element_by_css_selector('#q')
# 第二种方式，通过使用 By.xxx 指定查找方式
input = browser.find_element(By.ID,'q')
print(input)
# 在响应结果中通过xpath查找元素
input_third = browser.find_element_by_xpath('//*[@id="q"]')
print(input_first)
print(input_second)
print(input_third)
browser.close()
```

查多个元素的方法：  

- find_elements_by_name
- find_elements_by_xpath
- find_elements_by_link_text
- find_elements_by_partial_link_text
- find_elements_by_tag_name
- find_elements_by_class_name
- find_elements_by_css_selector

```python
from selenium import webdriver
from selenium.webdriver.common.by import By

# 申明一个浏览器对象
browser = webdriver.Chrome()
# 使用浏览器访问淘宝
browser.get('https://www.taobao.com')
# 根据选择查找多个元素
input_first = browser.find_elements_by_css_selector('.service-bd li')
input_second = browser.find_elements(By.CSS_SELECTOR,'.service-bd li')
print(input_first)
print(input_second)
browser.close()
```

**Selenium 元素交互操作**   
对获取的元素调用[交互方法](http://selenium-python.readthedocs.io/api.html#module-selenium.webdriver.remote.webelement)。  
```python
import time
from selenium import webdriver
from selenium.webdriver.common.by import By

# 申明一个浏览器对象
browser = webdriver.Chrome()
# 使用浏览器访问淘宝
browser.get('https://www.taobao.com')
# 根据 ID 查找元素
input_search = browser.find_element(By.ID,'q')
# 模拟输入 PSV 到输入框
input_search.send_keys('PSV')
time.sleep(2)
# 清空输入
input_search.clear()
input_search.send_keys('3DS')
button = browser.find_element(By.CLASS_NAME,'btn-search')
# 模拟点击
button.click()
```

**Selenium 交互动作**   
[更多操作](http://selenium-python.readthedocs.io/api.html#module-selenium.webdriver.common.action_chains)  

```python
from selenium import webdriver
from selenium.webdriver import ActionChains

browser = webdriver.Chrome()
url = 'http://www.runoob.com/try/try.php?filename=jqueryui-api-droppable'
browser.get(url)
# 切换id为iframeResult的frame
browser.switch_to.frame('iframeResult')
source = browser.find_element_by_css_selector('#draggable')
target = browser.find_element_by_css_selector('#droppable')
actions = ActionChains(browser)
actions.drag_and_drop(source, target)
actions.perform()
```

**Selenium 执行 JavaScript**   
```python
from selenium import webdriver

# 申明一个浏览器对象
browser = webdriver.Chrome()
browser.get('https://www.zhihu.com/explore')
# 执行JavaScript脚本
browser.execute_script('window.scrollTo(0, document.body.scrollHeight)')
browser.execute_script('alert("To Bottom")')
```

**Selenium 获取元素信息**  
```python
from selenium import webdriver
from selenium.webdriver.common.by import By

# 申明一个浏览器对象
browser = webdriver.Chrome()
browser.get('https://www.zhihu.com/explore')
logo = browser.find_element(By.ID,'zh-top-link-logo')
# 获取属性
print(logo.get_attribute('class'))
print(logo)
browser.close()

# 申明一个浏览器对象
browser = webdriver.Chrome()
browser.get('https://www.zhihu.com/explore')
submit = browser.find_element(By.ID,'zu-top-add-question')
# 获取文本值
print(submit.text)
print(submit)
browser.close()

# 申明一个浏览器对象
browser = webdriver.Chrome()
browser.get('https://www.zhihu.com/explore')
submit = browser.find_element(By.ID,'zu-top-add-question')
# 获取 id   0.04584255991652042-1
print(submit.id)
# 获取位置  {'y': 7, 'x': 675}
print(submit.location)
# 获取标签名称    button
print(submit.tag_name)
# 获取大小  {'width': 66, 'height': 32}
print(submit.size)
browser.close()
```

**Selenium Frame**  
```python
from selenium import webdriver
from selenium.common.exceptions import NoSuchElementException

browser = webdriver.Chrome()
url = 'http://www.runoob.com/try/try.php?filename=jqueryui-api-droppable'
browser.get(url)
# 将操作的响应数据换成 iframeResult
browser.switch_to.frame('iframeResult')
source = browser.find_element_by_css_selector('#draggable')
print(source)

try:
    logo = browser.find_element_by_class_name('logo')
exceptNoSuchElementException:
    print('NO LOGO')
# 切换成父元素
browser.switch_to.parent_frame()
logo = browser.find_element_by_class_name('logo')
print(logo)
print(logo.text)
```

**Selenium 等待**  
等待分隐式等待和显示等待。  
> 当使用了隐式等待执行测试的时候，如果 WebDriver 没有在 DOM 中找到元素，将继续等待，超出设定时间后则抛出找不到元素的异常，换句话说，当查找元素或元素并没有立即出现的时候，隐式等待将等待一段时间再查找 DOM，默认的时间是 0  
> 
> 显式等待指定某个条件，然后设置最长等待时间。如果在这个时间还没有找到元素，那么便会抛出异常了。  

- title_is 标题是某内容
- title_contains 标题包含某内容
- presence_of_element_located 元素加载出，传入定位元组，如 (By.ID, 'p')
- visibility_of_element_located 元素可见，传入定位元组
- visibility_of 可见，传入元素对象
- presence_of_all_elements_located 所有元素加载出
- text_to_be_present_in_element 某个元素文本包含某文字
- text_to_be_present_in_element_value 某个元素值包含某文字
- frame_to_be_available_and_switch_to_it frame 加载并切换
- invisibility_of_element_located 元素不可见
- element_to_be_clickable 元素可点击
- staleness_of 判断一个元素是否仍在 DOM，可判断页面是否已经刷新
- element_to_be_selected 元素可选择，传元素对象
- element_located_to_be_selected 元素可选择，传入定位元组
- element_selection_state_to_be 传入元素对象以及状态，相等返回 True，否则返回 False
- element_located_selection_state_to_be 传入定位元组以及状态，相等返回 True，否则返回 False
- alert_is_present 是否出现 Alert

[更多操作](http://selenium-python.readthedocs.io/api.html#module-selenium.webdriver.support.expected_conditions)  
```python
from selenium import webdriver

# 隐式等待
browser = webdriver.Chrome()
# 等待 10 秒
browser.implicitly_wait(10)
browser.get('https://www.zhihu.com/explore')
input = browser.find_element_by_class_name('zu-top-add-question')
print(input)

# 显示等待
from selenium.webdriver.common.by import By
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.support.wait import WebDriverWait

browser = webdriver.Chrome()
browser.get('https://www.taobao.com/')
# 显示等待 10s
wait = WebDriverWait(browser, 10)
# 等待直到元素加载出
input = wait.until(EC.presence_of_element_located((By.ID, 'q')))
# 等待直到元素可点击
button = wait.until(EC.element_to_be_clickable((By.CSS_SELECTOR, '.btn-search')))
print(input, button)
```

**Selenium 前进后退**  
```python
import time
from selenium import webdriver

browser = webdriver.Chrome()
browser.get('https://www.baidu.com/')
browser.get('https://www.taobao.com/')
browser.get('https://www.python.org/')
# 后退
browser.back()
time.sleep(1)
# 前进
browser.forward()
browser.close()
```

**Selenium Cookies**  
```python
from selenium import webdriver

browser = webdriver.Chrome()
browser.get('https://www.zhihu.com/explore')
# 获得 cookies
print(browser.get_cookies())
# 添加 cookie
browser.add_cookie({'name': 'name', 'domain': 'www.zhihu.com', 'value': 'germey'})
print(browser.get_cookies())
# 删除所有 cookies
browser.delete_all_cookies()
print(browser.get_cookies())
```

**Selenium 选项卡管理**  
```python
import time
from selenium import webdriver

browser = webdriver.Chrome()
browser.get('https://www.baidu.com')
# 打开一个选项卡
browser.execute_script('window.open()')
print(browser.window_handles)
# 选择第二个选项卡
browser.switch_to_window(browser.window_handles[1])
browser.get('https://www.taobao.com')
time.sleep(1)
browser.switch_to_window(browser.window_handles[0])
browser.get('https://python.org')
```

**Selenium 异常处理**  
```python
from selenium import webdriver
from selenium.common.exceptions import TimeoutException, NoSuchElementException

browser = webdriver.Chrome()
try:
    browser.get('https://www.baidu.com')
exceptTimeoutException:
    print('Time Out')
try:
    browser.find_element_by_id('hello')
exceptNoSuchElementException:
    print('No Element')
finally:
    browser.close()
```

