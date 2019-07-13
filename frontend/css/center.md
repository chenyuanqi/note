
居中，很常见的页面处理。  
主要会用到如下方式:
```css
// 行内元素
text-align: center;
// 块元素
margin: 0 auto;
// 定位方式
position: relative | absolute; left: 50%;
// 垂直方向居中
vertical-align: middle;
// 元素转换
transform: translate(-50%);
// 浮动元素居中
margin-left: 50%; transform: translate(-50%);
// 绝对定位的 div 居中
top:0; left:0; bottom:0; right:0; margin: auto;
```

### css 水平垂直居中实现方式
水平垂直居中包括行内元素居中，以及块级元素居中。  
```html
<!-- 基础样式 -->
<style>
.outer {
    width: 400px;
    height: 400px;
    border: 1px solid red;
}
.outer .inner {
    width: 50px;
    height: 50px;
    border: 1px solid blue;
}
</style>

<!-- 行内元素 html 结构 -->
<div class="outer">
    <span class="inner"></span>
</div>

<!-- 块级元素结构 -->
<div class="outer">
	<div class="inner"></div>
</div>
```

### 水平居中
1、行内元素  
```css
.outer {
    text-align: center;
}
```

2、块级元素
```css
// a.内部元素必须定义宽度，不然 margin 属性会无效
.outer .inner {
    margin: auto;
}

// b.如果不想定义宽度，可以设置内部元素的 display 为 table，它的宽度会由内部元素来撑开
.outer .inner {
	display: table;
    margin: auto;
}

// c.为内部元素设置 display 为 inline，将它转换为行内元素，再对父元素使用 text-align: center 可以实现水平居中
// 缺点就是内部元素无法设置宽度
.outer {
    text-align: center;
}
.outer .inner {
    display: inline;
}

// d.采用 inline-block，则可以为内部元素设置宽度
.outer {
    text-align: center;
}
.outer .inner {
    display: inline-block;
}

// e.不需要知道内部元素宽度
.outer .inner {
    position: relative;
    left: 50%;
    transform: translateX(-50%);
}

// f.负边距 + 绝对定位
.outer {
    position: relative;
}
.outer .inner {
    position: absolute;
    left: 50%;
    margin-left: -25px;
}

// g.用的最多的方式，但低版本浏览器会有兼容问题
.outer {
    display: flex;
    justify-content: center;  // 主轴上居中
}
```

### 垂直居中
1、行内元素  
```css
// 外部元素设置 line-height
.outer {
    line-height: 400px;
}
```

2、块级元素
```css
// a.使用绝对定位将内部元素的顶部定位在中间，再通过 margin-top 负值拉回高度，需要提前知道内部元素的高度
.outer {
    position: relative;
}
.outer .inner {
    position: absolute;
    top: 50%;
    margin-top: -25px;
}

// b.不需要知道内部元素的高度，兼容性也很好
.outer {
    position: relative;
}
.outer .inner {
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    margin: auto;
}

// c.relative + transform
// 可以使用 position: absolute 方式，但要对应地将外部元素设置成 position: relative
.outer .inner {
    position: relative;
    top: 50%;
    transform: translateY(-50%);
}

// d.vertical-align + table-cell
.outer {
    display: table-cell;
    vertical-align: middle;
}

// e.新建一个 inner 的兄弟元素，高度撑开整个容器，再对 inner 使用 vertical-align: middle 使元素居中，不需要知道内部元素的高度
<div class="outer">
	<div class="inner"></div>
	<div class="sibling"></div>
</div>

.outer .inner {
    vertical-align: middle;
    display: inline-block;
}
.outer .slibing {
    height: 400px;
    display: inline-block;
    vertical-align: middle;
}

// f.通过伪元素去撑开高度
.inner {
    display: inline-block;
    vertical-align: middle;
}
.outer::before {
    content: '';
    height: 100%;
    display: inline-block;
    vertical-align: middle;
}

// g.使用 flexbox
.outer {
    display: flex;
    align-items: center;
}
```