
### 一、Css 权重
CSS 权重指的是样式的优先级，有两条或多条样式作用于一个元素，权重高的那条样式对元素起作用，权重相同的，后写的样式会覆盖前面写的样式。

#### 1.1 权重的等级
Css 权重可以把样式的应用方式分为几个等级，按照等级来计算权重   
> 1.!important，加在样式属性值后，权重值为 10000  
> 2.内联样式，如：style=””，权重值为 1000  
> 3.ID 选择器，如：#content，权重值为 100  
> 4.类，伪类和属性选择器，如： content、:hover 权重值为 10  
> 5.标签选择器和伪元素选择器，如：div、p、:before 权重值为 1  
> 6.通用选择器（\*）、子选择器（\>）、相邻选择器（+）、同胞选择器（\~）、权重值为 0  

#### 1.2 权重的计算

### 二、Css 选择器
[Css 选择器演示](https://www.haorooms.com/tools/css_selecter/)  

#### 2.1 标签选择器
标签选择器，此种选择器影响范围大，建议尽量应用在层级选择器中。  
```css
*{margin:0;padding:0}
div{color:red}   
```

#### 2.2 id 选择器
通过 id 名来选择元素，元素的 id 名称不能重复，所以一个样式设置项只能对应于页面上一个元素，不能复用，id 名一般给程序使用，所以不推荐使用 id 作为选择器。  
```css
#box{color:red} 
```

#### 2.3 类选择器
通过类名来选择元素，一个类可应用于多个元素，一个元素上也可以使用多个类，应用灵活，可复用，是 css 中应用最多的一种选择器。  
```css
.red{color:red}
.big{font-size:20px}
.mt10{margin-top:10px} 
```

#### 2.4 层级选择器
主要应用在选择父元素下的子元素，或者子元素下面的子元素，可与标签元素结合使用，减少命名，同时也可以通过层级，防止命名冲突。
```html
<style>
.box span{color:red}
.box .red{color:pink}
.red{color:red}
</style>
<div class="box">
    <span>....</span>
    <a href="#" class="red">....</a>
</div>

<h3 class="red">....</h3>
```

#### 2.5 组选择器
多个选择器，如果有同样的样式设置，可以使用组选择器。  
```css
.box1,.box2,.box3{width:100px;height:100px}
.box1{background:red}
.box2{background:pink}
.box2{background:gold}
```

#### 2.6 伪类及伪元素选择器
常用的伪类选择器有 hover，表示鼠标悬浮在元素上时的状态，伪元素选择器有 before 和 after, 它们可以通过样式在元素中插入内容。
```css
.box1:hover{color:red}
.box2:before{content:'行首文字';}
.box3:after{content:'行尾文字';}
```

:before 可用于在某个元素之前插入某些内容，:after 可用于在某个元素之后插入某些内容。
```html
<style>
	p:before{
	    content: "H"
	}
	p:after{
	    content: "d"
	}
</style>
<p>ello Worl</p>
```
利用 :before 和 ：after 做个简单的对话框
```html
<style>
	.test-div{
	    position: relative;
	    width:150px;
	    height: 36px;
	    border:1px solid black;
	    border-radius:5px;
	    background: rgba(245,245,245,1)
	}
	.test-div:before,.test-div:after{
	    content: "";
	    display: block;
	    position: absolute;
	    top:8px;
	    width: 0;
	    height: 0;
	    border:6px solid transparent;
	}
	.test-div:before{
	    left:-11px;
	    border-right-color: rgba(245,245,245,1);
	    z-index:1
	}
	.test-div:after{
	    left:-12px;
	    border-right-color: rgba(0,0,0,1);
	    z-index: 0
	}
</style>
<div class="test-div"></div>
```

#### 2.7 属性选择器
属性选择器可以根据元素的属性及属性值来选择元素。  
```css
// 把包含标题（title）的所有元素变为红色
*[title] {
	color:red;
}

// 将同时有 href 和 title 属性的 HTML 超链接的文本设置为红色 
a[href][title] {
	color:red;
}

// 指定将连接为 www.xxx.com 的文字颜色变红
a[href="www.xxx.com"] {
	color: red;
}

// 根据属性值中的词列表的某个词进行选择
// <p class="important warning">This is a paragraph.</a>
// <p class="important">This is a paragraph.</a>
p[class~="important"] {
	color: red;
}

// 选择 attr 属性值以 "xxx" 开头的所有元素
[attr^="xxx"] {
	color: red;
}

// 选择 attr 属性值以 "xxx" 结尾的所有元素
[attr$="xxx"] {
	color: red;
}

// 选择 lang 属性等于 en 或以 en- 开头的所有元素
*[lang|="en"] {
	color: red;
}
```

#### 2.8 选择器特殊符号
1、>(大于号)  
大于号代表选择子元素  
```html
<!-- h1>strong ，只有第一个 h1 下面的 strong 被选中，第二个不起作用 -->
<h1>This is <strong class="haorooms">very</strong> <strong>very</strong> important.</h1>
<h1>This is <em>really <strong>very</strong></em> important.</h1>
```

2、+ 号  
选择相邻兄弟  
```html
<h1>This is a heading.</h1>
<p>This is paragraph.</p>
<p>This is paragraph.</p>

<style>
// h1 后面的第一个p元素会有 50px 的间距
h1 + p {
	margin-top:50px;
} 
</style>
```

3、~ 波浪号  
比如 p~ul 代表：所有相同的父元素中位于 p 元素之后的所有 ul 元素。  
```html
<div>一个 div 元素。</div>
<ul>
  <li>咖啡</li>
  <li>牛奶</li>
  <li>茶</li>
</ul>

<p>第一段。</p>
<ul>
  <li>咖啡</li>
  <li>牛奶</li>
  <li>茶</li>
</ul>
```