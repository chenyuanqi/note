
### 1、初始化
pyquery 允许你对 xml 文档进行 jquery 查询，它的 api 类似于 jquery。  
pyquery 使用 lxml 进行快速的 xml 和 html 操作。 
 
[官方文档参考](https://pyquery.readthedocs.io/en/latest/)  

#### 1.1、字符串初始化
```python
from pyquery import PyQuery as pq

html = '''
<div>
    <ul>
         <li class="item-0">first item</li>
         <li class="item-1"><a href="https://ask.hellobi.com/link2.html">second item</a></li>
         <li class="item-0 active"><a href="https://ask.hellobi.com/link3.html"><span class="bold">third item</span></a></li>
         <li class="item-1 active"><a href="https://ask.hellobi.com/link4.html">fourth item</a></li>
         <li class="item-0"><a href="https://ask.hellobi.com/link5.html">fifth item</a></li>
     </ul>
 </div>
'''

doc = pq(html)
print(doc('li'))
```

#### 1.2、URL 初始化
```python
from pyquery import PyQuery as pq

# 通过URL来获取
doc = pq(url='http://www.baidu.com')
# <class 'pyquery.pyquery.PyQuery'>
print(type(doc('title')))
# 输出选中的head标签
print(doc('head'))
```

#### 1.3、文件的初始化
```python
from pyquery import PyQuery as pq

# 通过文件来获取
doc = pq(filename='demo1.html')
# <class 'pyquery.pyquery.PyQuery'>
print(type(doc('li')))
# 输出所有的li标签
print(doc('li'))
```

### 2、基本的 CSS 选择器
```python
from pyquery import PyQuery as pq

html = '''
<div id="container">
    <ul class="list">
         <li class="item-0">first item</li>
         <li class="item-1"><a href="https://ask.hellobi.com/link2.html">second item</a></li>
         <li class="item-0 active"><a href="https://ask.hellobi.com/link3.html"><span class="bold">third item</span></a></li>
         <li class="item-1 active"><a href="https://ask.hellobi.com/link4.html">fourth item</a></li>
         <li class="item-0"><a href="https://ask.hellobi.com/link5.html">fifth item</a></li>
     </ul>
 </div>
'''

doc = pq(html)
# 选中id为container中的class为list中的li标签
print(doc('#container .list li'))
```

### 3、查找元素
#### 3.1、子元素
```python
from pyquery import PyQuery as pq

html = '''
<div id="container">
    <ul class="list">
         <li class="item-0">first item</li>
         <li class="item-1"><a href="https://ask.hellobi.com/link2.html">second item</a></li>
         <li class="item-0 active"><a href="https://ask.hellobi.com/link3.html"><span class="bold">third item</span></a></li>
         <li class="item-1 active"><a href="https://ask.hellobi.com/link4.html">fourth item</a></li>
         <li class="item-0"><a href="https://ask.hellobi.com/link5.html">fifth item</a></li>
     </ul>
 </div>
'''
doc = pq(html)
# 获取class为list的元素
items = doc('.list')
# <class 'pyquery.pyquery.PyQuery'>
print(type(items))
print(items)
# 在先前找到的元素中获取li标签
lis = items.find('li')
# <class 'pyquery.pyquery.PyQuery'>
print(type(lis))
print(lis)
# 获取先前找到的元素中的所有子元素
lis2 = items.children()
print(type(lis2))
print(lis2)
# 获取先前找到的元素中的class为active的元素
li3 = items.children('.active')
print(li3)
```

#### 3.2、父元素
```python
from pyquery import PyQuery as pq

html = '''
<div id="container">
    <ul class="list">
         <li class="item-0">first item</li>
         <li class="item-1"><a href="https://ask.hellobi.com/link2.html">second item</a></li>
         <li class="item-0 active"><a href="https://ask.hellobi.com/link3.html"><span class="bold">third item</span></a></li>
         <li class="item-1 active"><a href="https://ask.hellobi.com/link4.html">fourth item</a></li>
         <li class="item-0"><a href="https://ask.hellobi.com/link5.html">fifth item</a></li>
     </ul>
 </div>
'''
doc = pq(html)
# 获取class为list的元素
items = doc('.list')
# 获取所选元素的父元素
container = items.parent()
print(type(container))
print(container)
print("==========================")
# 获取所选元素的所有父元素
parents = items.parents()
print(type(parents))
print(parents)
print("==========================")
# 获取所选元素的所有父元素中class为container的元素
parent = items.parents('.container')
print(parent)
```

#### 3.3、兄弟元素
```python
from pyquery import PyQuery as pq

html = '''
<div id="container">
    <ul class="list">
         <li class="item-0">first item</li>
         <li class="item-1"><a href="https://ask.hellobi.com/link2.html">second item</a></li>
         <li class="item-0 active"><a href="https://ask.hellobi.com/link3.html"><span class="bold">third item</span></a></li>
         <li class="item-1 active"><a href="https://ask.hellobi.com/link4.html">fourth item</a></li>
         <li class="item-0"><a href="https://ask.hellobi.com/link5.html">fifth item</a></li>
     </ul>
 </div>
'''
doc = pq(html)
# 获取class为list的元素
items = doc('.list')
li = doc('.list .item-0.active')
# 查找选中元素的所有兄弟元素(不包含自己)
print(li.siblings())
# 查找选中元素的所有兄弟元素中class为active的元素(不包含自己)
print(li.siblings('.active'))
```

### 4、遍历
#### 4.1、单个元素
```python
from pyquery import PyQuery as pq

html = '''
<div id="container">
    <ul class="list">
         <li class="item-0">first item</li>
         <li class="item-1"><a href="https://ask.hellobi.com/link2.html">second item</a></li>
         <li class="item-0 active"><a href="https://ask.hellobi.com/link3.html"><span class="bold">third item</span></a></li>
         <li class="item-1 active"><a href="https://ask.hellobi.com/link4.html">fourth item</a></li>
         <li class="item-0"><a href="https://ask.hellobi.com/link5.html">fifth item</a></li>
     </ul>
 </div>
'''
doc = pq(html)

# 选中单个单个元素
li = doc('.item-0.active')
print(li)
```

#### 4.2、多个元素
```python
from pyquery import PyQuery as pq

html = '''
<div id="container">
    <ul class="list">
         <li class="item-0">first item</li>
         <li class="item-1"><a href="https://ask.hellobi.com/link2.html">second item</a></li>
         <li class="item-0 active"><a href="https://ask.hellobi.com/link3.html"><span class="bold">third item</span></a></li>
         <li class="item-1 active"><a href="https://ask.hellobi.com/link4.html">fourth item</a></li>
         <li class="item-0"><a href="https://ask.hellobi.com/link5.html">fifth item</a></li>
     </ul>
 </div>
'''
doc = pq(html)

# 查找所有li标签
lis = doc('li').items()
# <class 'generator'>
print(type(lis))
forli inlis:
    print(li)
```

### 5、获取信息
#### 5.1、获取属性
```python
from pyquery import PyQuery as pq

html = '''
<div id="container">
    <ul class="list">
         <li class="item-0">first item</li>
         <li class="item-1"><a href="https://ask.hellobi.com/link2.html">second item</a></li>
         <li class="item-0 active"><a href="https://ask.hellobi.com/link3.html"><span class="bold">third item</span></a></li>
         <li class="item-1 active"><a href="https://ask.hellobi.com/link4.html">fourth item</a></li>
         <li class="item-0"><a href="https://ask.hellobi.com/link5.html">fifth item</a></li>
     </ul>
 </div>
'''
doc = pq(html)
a = doc('.item-0.active a')
# <a href="https://ask.hellobi.com/link3.html"><span class="bold">third item</span></a>
print(a)
# link3.html    获取选中标签的href属性
print(a.attr('href'))
# link3.html
print(a.attr.href)
```

#### 5.2、获取文本
```python
from pyquery import PyQuery as pq

html = '''
<div id="container">
    <ul class="list">
         <li class="item-0">first item</li>
         <li class="item-1"><a href="https://ask.hellobi.com/link2.html">second item</a></li>
         <li class="item-0 active"><a href="https://ask.hellobi.com/link3.html"><span class="bold">third item</span></a></li>
         <li class="item-1 active"><a href="https://ask.hellobi.com/link4.html">fourth item</a></li>
         <li class="item-0"><a href="https://ask.hellobi.com/link5.html">fifth item</a></li>
     </ul>
 </div>
'''
doc = pq(html)
a = doc('.item-0.active a')
# <a href="https://ask.hellobi.com/link3.html"><span class="bold">third item</span></a>
print(a)
# 获取a标签的内容
print(a.text())
```

#### 5.3、获取 HTML
```python
from pyquery import PyQuery as pq

html = '''
<div id="container">
    <ul class="list">
         <li class="item-0">first item</li>
         <li class="item-1"><a href="https://ask.hellobi.com/link2.html">second item</a></li>
         <li class="item-0 active"><a href="https://ask.hellobi.com/link3.html"><span class="bold">third item</span></a></li>
         <li class="item-1 active"><a href="https://ask.hellobi.com/link4.html">fourth item</a></li>
         <li class="item-0"><a href="https://ask.hellobi.com/link5.html">fifth item</a></li>
     </ul>
 </div>
'''
doc = pq(html)
li = doc('.item-1.active')
# <li class="item-1 active"><a href="https://ask.hellobi.com/link4.html">fourth item</a></li>
print(li)
# 获取li标签的HTML
print(li.html())
```

### 6、DOM 操作
#### 6.1、addClass、removeClass
```python
from pyquery import PyQuery as pq

html = '''
<div id="container">
    <ul class="list">
         <li class="item-0">first item</li>
         <li class="item-1"><a href="https://ask.hellobi.com/link2.html">second item</a></li>
         <li class="item-0 active"><a href="https://ask.hellobi.com/link3.html"><span class="bold">third item</span></a></li>
         <li class="item-1 active"><a href="https://ask.hellobi.com/link4.html">fourth item</a></li>
         <li class="item-0"><a href="https://ask.hellobi.com/link5.html">fifth item</a></li>
     </ul>
 </div>
'''
doc = pq(html)
li = doc('.item-0.active')
print(li)
# 移除class
li.removeClass('active')
print(li)
# 添加class
li.addClass('active')
print(li)
```

#### 6.2、attr、css
```python
from pyquery import PyQuery as pq

html = '''
<div id="container">
    <ul class="list">
         <li class="item-0">first item</li>
         <li class="item-1"><a href="https://ask.hellobi.com/link2.html">second item</a></li>
         <li class="item-0 active"><a href="https://ask.hellobi.com/link3.html"><span class="bold">third item</span></a></li>
         <li class="item-1 active"><a href="https://ask.hellobi.com/link4.html">fourth item</a></li>
         <li class="item-0"><a href="https://ask.hellobi.com/link5.html">fifth item</a></li>
     </ul>
 </div>
'''
doc = pq(html)
li = doc('.item-0.active')
print(li)
# 添加name属性
li.attr('name', 'link')
print(li)
# 添加css样式
li.css('font-size', '14px')
print(li)
```

#### 6.3、remove
```python
from pyquery import PyQuery as pq

html = '''
<div class="wrap">
    Hello, World
    <p>This is a paragraph.</p>
 </div>
'''
doc = pq(html)
wrap = doc('.wrap')
print(wrap.text())
# 在选择的元素中找到p标签并移除
wrap.find('p').remove()
print(wrap.text())
```

#### 6.4、其他 DOM 方法
参考 [http://pyquery.readthedocs.io/en/latest/api.html](http://pyquery.readthedocs.io/en/latest/api.html)  


### 7、伪类选择器
```python
from pyquery import PyQuery as pq

html = '''
<div class="wrap">
    <div id="container">
        <ul class="list">
             <li class="item-0">first item</li>
             <li class="item-1"><a href="https://ask.hellobi.com/link2.html">second item</a></li>
             <li class="item-0 active"><a href="https://ask.hellobi.com/link3.html"><span class="bold">third item</span></a></li>
             <li class="item-1 active"><a href="https://ask.hellobi.com/link4.html">fourth item</a></li>
             <li class="item-0"><a href="https://ask.hellobi.com/link5.html">fifth item</a></li>
         </ul>
     </div>
 </div>
'''

doc = pq(html)
# 找到第一个li
li = doc('li:first-child')
print(li)
# 找到最后一个li
li = doc('li:last-child')
print(li)
# 找到第二个li
li = doc('li:nth-child(2)')
print(li)
# 找到第三个到最后的li
li = doc('li:gt(2)')
print(li)
# 找到第偶数个li
li = doc('li:nth-child(2n)')
print(li)
# 找到内容包含second的li
li = doc('li:contains(second)')
print(li)
```
