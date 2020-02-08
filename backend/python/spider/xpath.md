
### 1、概览
XPath，全称 XML Path Language，即 XML 路径语言，它是一门在 XML 文档中查找信息的语言。XPath 最初设计是用来搜寻 XML 文档的，但是它同样适用于 HTML 文档的搜索。  

XPath 的选择功能十分强大，它提供了非常简洁明了的路径选择表达式，另外它还提供了超过 100 个内建函数用于字符串、数值、时间的匹配以及节点、序列的处理等等，几乎所有我们想要定位的节点都可以用 XPath 来选择。  
XPath 于 1999 年 11 月 16 日 成为 W3C 标准，它被设计为供 XSLT、XPointer 以及其他 XML 解析软件使用，更多的文档可以访问其官方网站：https://www.w3.org/TR/xpath/。

[XPath 文档](https://developer.mozilla.org/zh-CN/docs/Web/XPath)  

```bash
pip install lxml 
```

### 2、XPath 基本使用

**XPath 常用规则**  

| 表达式 | 描述 |  
| :---: | :---: |  
| nodename | 选取此节点的所有子节点 |  
| / | 从当前节点选取直接子节点 |  
| // | 从当前节点选取子孙节点 |  
| . | 选取当前节点 |  
| .. | 选取当前节点的父节点 |  
| @ | 选取属性 |  

例如：//title[@lang='eng'] 表示一个 XPath 规则，它就代表选择所有名称为 title，同时属性 lang 的值为 eng 的节点。  

**XPath 初步使用**  
```python
from lxml import etree

# HTML 文本中的最后一个 li 节点是没有闭合的，但是 etree 模块可以对 HTML 文本进行自动修正
text = '''
<div>
    <ul>
         <li><a href="https://ask.hellobi.com/link1.html">first item</a></li>
         <li><a href="https://ask.hellobi.com/link2.html">second item</a></li>
         <li><a href="https://ask.hellobi.com/link3.html">third item</a></li>
         <li><a href="https://ask.hellobi.com/link4.html">fourth item</a></li>
         <li><a href="https://ask.hellobi.com/link5.html">fifth item</a>
     </ul>
 </div>
'''
html = etree.HTML(text) # 调用 HTML 类进行初始化，构造了一个 XPath 解析对象
result = etree.tostring(html) # 输出修正后的 HTML 代码
print(result.decode('utf-8')) # 转成 str 类型（经过处理之后 li 节点标签被补全，并且还自动添加了 body、html 节点）

# 直接读取文本文件进行解析（test.html 内容即 text）
html = etree.parse('test.html', etree.HTMLParser())
result = etree.tostring(html)
print(result.decode('utf-8')) # 输出结果略有不同，多了一个 DOCTYPE 的声明，不过对解析无任何影响
```

**XPath 所有节点**  
我们一般会用 // 开头的 XPath 规则来选取所有符合要求的节点。  
```python
from lxml import etree

html = etree.parse('test.html', etree.HTMLParser())
result = html.xpath('//*') # * 代表匹配所有节点，也就是整个 HTML 文本中的所有节点都会被获取
# 也可以指定节点名称，比如想获取所有 li 节点
# result = html.xpath('//li')
# print(result[0]) # 取第一个 li 节点
print(result) # 返回形式是一个列表，每个元素是 Element 类型，其后跟了节点的名称，所有的节点都包含在列表中
# [<Element html at 0x10510d9c8>, <Element body at 0x10510da08>, <Element div at 0x10510da48>, <Element ul at 0x10510da88>, <Element li at 0x10510dac8>, <Element a at 0x10510db48>, <Element li at 0x10510db88>, <Element a at 0x10510dbc8>, <Element li at 0x10510dc08>, <Element a at 0x10510db08>, <Element li at 0x10510dc48>, <Element a at 0x10510dc88>, <Element li at 0x10510dcc8>, <Element a at 0x10510dd08>]
```

**XPath 子节点**  
通过 / 或 // 即可查找元素的子节点或子孙节点。  
```python
from lxml import etree

html = etree.parse('test.html', etree.HTMLParser())
# 获取了所有 li 节点的所有直接 a 子节点
result = html.xpath('//li/a')
# 如果我们要获取所有子孙节点就该使用 // 
# result = html.xpath('//ul//a') # 结果是一样的
# 需要注意的是：//ul/a 就无法获取任何结果，因为在 ul 节点下没有直接的 a 子节点
print(result)
```

**XPath 父节点**  
可以用 .. 来获取父节点，也可以通过 parent:: 来获取父节点。  
```python
from lxml import etree

html = etree.parse('test.html', etree.HTMLParser())
# 先选中 href 是 link4.html 的 a 节点，然后再获取其父节点，然后再获取其 class 属性
result = html.xpath('//a[@href="https://ask.hellobi.com/link4.html"]/../@class')
# result = html.xpath('//a[@href="https://ask.hellobi.com/link4.html"]/parent::*/@class')
print(result) # ['item-1']
```

**XPath 属性匹配**  
在选取的时候我们可以用 @ 符号进行属性过滤。  
```python
from lxml import etree

html = etree.parse('test.html', etree.HTMLParser())
# 选取 class 为 item-1 的 li 节点
result = html.xpath('//li[@class="item-0"]')
print(result) # 符合条件的 li 节点有两个：[<Element li at 0x10a399288>, <Element li at 0x10a3992c8>]
```

**XPath 文本获取**  
用 XPath 中的 text () 方法可以获取节点中的文本。  
```python
from lxml import etree

html = etree.parse('test.html', etree.HTMLParser())
result = html.xpath('//li[@class="item-0"]/text()')
print(result) # ['\n     ']
# li 的直接子节点都是 a 节点，而且自动修正的 li 节点的尾标签换行，所以获取到的是换行符 \n
# 应该这样获取
# result = html.xpath('//li[@class="item-0"]/a/text()') # ['first item', 'fifth item']
# 或者
# result = html.xpath('//li[@class="item-0"]//text()') # ['first item', 'fifth item', '\n     ']
```

**XPath 属性获取**  
节点属性的获取还是用 @ 符号就可以。
```python
from lxml import etree

html = etree.parse('test.html', etree.HTMLParser())
# 获取所有 li 节点下所有 a 节点的 href 属性
result = html.xpath('//li/a/@href')
print(result) # ['link1.html', 'link2.html', 'link3.html', 'link4.html', 'link5.html']
# 注意：属性匹配是中括号加属性名和值来限定某个属性，而此处的 @href 指的是获取节点的某个属性
```

**XPath 属性多值匹配**  
有时候某些节点的某个属性可能有多个值。  
```python
from lxml import etree

text = '''
<li class="li li-first"><a href="https://ask.hellobi.com/link.html">first item</a></li>
'''
html = etree.HTML(text)
# contains () 方法：第一个参数传入属性名称，第二个参数传入属性值
# 注意：使用属性匹配获取是无法匹配的
result = html.xpath('//li[contains(@class, "li")]/a/text()')
print(result) # ['first item']
```

**XPath 多属性匹配**  
有时候根据多个属性才能确定一个节点，需要同时匹配多个属性才可以，则可以使用运算符 and 来连接。  
```python
from lxml import etree

text = '''
<li class="li li-first" name="item"><a href="https://ask.hellobi.com/link.html">first item</a></li>
'''
html = etree.HTML(text)
result = html.xpath('//li[contains(@class, "li") and @name="item"]/a/text()')
print(result)
```

and 其实是 XPath 中的运算符，另外还有很多运算符，如 or、mod 等等。  

| 运算符 | 描述 | 实例 | 返回值 |  
| :---: | :---: | :---: | :---: |  
| or | 或 | price=9.80 or price=9.70 | 如果 price 是 9.80，则返回 true。如果 price 是 9.50，则返回 false。 |  
| and | 与 | price>9.00 and price<9.90 | 如果 price 是 9.80，则返回 true。如果 price 是 8.50，则返回 false。 |  
| mod | 计算除法的余数 | 5 mod 21 |  
| \ | 计算两个节点集 | //book \//cd | 返回所有拥有 book 和 cd 元素的节点集 |  
| + | 加法 | 6 + 410 |  
| - | 减法 | 6 - 42 |  
| * | 乘法 | 6 * 424 |  
| div | 除法 | 8 div 42 |  
| = | 等于 | price=9.80 | 如果 price 是 9.80，则返回 true。如果 price 是 9.90，则返回 false。 |  
| != | 不等于 | price!=9.80 | 如果 price 是 9.90，则返回 true。如果 price 是 9.80，则返回 false。 |  
| < | 小于 | price<9.80 | 如果 price 是 9.00，则返回 true。如果 price 是 9.90，则返回 false。 |  
| <= | 小于或等于 | price<=9.80 | 如果 price 是 9.00，则返回 true。如果 price 是 9.90，则返回 false。 |  
| > | 大于 | price>9.80 | 如果 price 是 9.90，则返回 true。如果 price 是 9.80，则返回 false。 |  
| >= | 大于或等于 | price>=9.80 | 如果 price 是 9.90，则返回 true。如果 price 是 9.70，则返回 false。|  

**XPath 按序选择**  
在选择的时候可能某些属性同时匹配了多个节点，但是我们只想要其中的某个节点，如第二个节点，或者最后一个节点 ... 这时候可以利用中括号传入索引的方法获取特定次序的节点。  
```python
from lxml import etree

text = '''
<div>
    <ul>
         <li><a href="https://ask.hellobi.com/link1.html">first item</a></li>
         <li><a href="https://ask.hellobi.com/link2.html">second item</a></li>
         <li><a href="https://ask.hellobi.com/link3.html">third item</a></li>
         <li><a href="https://ask.hellobi.com/link4.html">fourth item</a></li>
         <li><a href="https://ask.hellobi.com/link5.html">fifth item</a>
     </ul>
 </div>
'''
html = etree.HTML(text)
result = html.xpath('//li[1]/a/text()') # 选取第一个 li 节点。注意这里和代码中不同，序号是以 1 开头的，不是 0 开头的。
print(result) # ['first item']
result = html.xpath('//li[last()]/a/text()') # 选取最后一个 li 节点
print(result) # ['first item']
result = html.xpath('//li[position()<3]/a/text()') # 选取位置小于 3 的 li 节点
print(result) # ['first item', 'second item']
result = html.xpath('//li[last()-2]/a/text()') # 选取倒数第三个 li 节点
print(result) # ['first item']
```

XPath 中提供了 100 多个函数，包括存取、数值、字符串、逻辑、节点、序列等处理功能，具体所有的函数作用可以参考[这里](http://www.w3school.com.cn/xpath/xpath_functions.asp)。  

**XPath 节点轴选择**  
XPath 提供了很多节点轴选择方法，英文叫做 XPath Axes，包括获取子元素、兄弟元素、父元素、祖先元素等等，在一定情况下使用它可以方便地完成节点的选择。  
```python
from lxml import etree

text = '''
<div>
    <ul>
         <li><a href="https://ask.hellobi.com/link1.html"><span>first item</span></a></li>
         <li><a href="https://ask.hellobi.com/link2.html">second item</a></li>
         <li><a href="https://ask.hellobi.com/link3.html">third item</a></li>
         <li><a href="https://ask.hellobi.com/link4.html">fourth item</a></li>
         <li><a href="https://ask.hellobi.com/link5.html">fifth item</a>
     </ul>
 </div>
'''
html = etree.HTML(text)
# 调用 ancestor 轴，可以获取所有祖先节点，其后需要跟两个冒号，然后是节点的选择器，这里我们直接使用了 *，表示匹配所有节点
result = html.xpath('//li[1]/ancestor::*') 
print(result) # [<Element html at 0x107941808>, <Element body at 0x1079418c8>, <Element div at 0x107941908>, <Element ul at 0x107941948>]
# 加了限定条件，这次在冒号后面加了 div，这样得到的结果就只有 div 这个祖先节点
result = html.xpath('//li[1]/ancestor::div')
print(result) # [<Element div at 0x107941908>]
# 调用 attribute 轴，可以获取所有属性值，其后跟的选择器还是 *，这代表获取节点的所有属性，返回值就是 li 节点的所有属性值
result = html.xpath('//li[1]/attribute::*')
print(result) # ['item-0']
# 调用 child 轴，可以获取所有直接子节点，在这里我们又加了限定条件选取 href 属性为 link1.html 的 a 节点
result = html.xpath('//li[1]/child::a[@href="https://ask.hellobi.com/link1.html"]')
print(result) # [<Element a at 0x1079418c8>]
# 调用 descendant 轴，可以获取所有子孙节点，这里我们又加了限定条件获取 span 节点，所以返回的就是只包含 span 节点而没有 a 节点
result = html.xpath('//li[1]/descendant::span')
print(result) # [<Element span at 0x107941948>]
# 调用 following 轴，可以获取当前节点之后的所有节点，这里我们虽然使用的是 * 匹配，但又加了索引选择，所以只获取了第二个后续节点
result = html.xpath('//li[1]/following::*[2]')
print(result) # [<Element a at 0x1079418c8>]
# 调用 following-sibling 轴，可以获取当前节点之后的所有同级节点，这里我们使用的是 * 匹配，所以获取了所有后续同级节点
result = html.xpath('//li[1]/following-sibling::*')
print(result) # [<Element li at 0x107941948>, <Element li at 0x107941988>, <Element li at 0x1079419c8>, <Element li at 0x107941a08>]
```

以上是 XPath 轴的简单用法，更多的轴的使用可以参考[这里](http://www.w3school.com.cn/xpath/xpath_axes.asp)。
