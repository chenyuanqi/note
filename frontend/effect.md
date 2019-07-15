
### 效果参考
[上海-豪情 前端作品](http://jikeytang.github.io/)  
[fgm](http://www.fgm.cc/learn/)  
[米空格 js 作品](http://www.laoshu133.com/Lab/)  
[上海 - 玄沐](http://k.swao.cn/js/)   
[移动在线做题组件](https://github.com/webjyh/Mexam)  
[Css 灵感](https://chokcoco.github.io/CSS-Inspiration/#/)  

### 文字内容溢出用省略号表示
```html
<div class="text_overflow">
    <div class="text_con"> 这是一段比较长的文字，用来测试是否文字溢出时会用省略号显示。</div>
    <div class="text_dotted">…</div>
</div>
<style>
.text_overflow {width:24em; height:1.3em; overflow:hidden; zoom:1;}
.text_overflow .text_con {float:left; height:1.3em; margin-right:3em; overflow:hidden;}
.text_overflow .text_dotted {width:3em; height:1.31em; float:right; margin-top:-1.3em;}
</style>
```

### 鼠标划过展示框的实现
```html
<div>
    <span data-content="测试隐藏的文本，鼠标移过展示框实现">测试</span>
</div>

<style>
div {
    margin-top: 30px;
}

span {
    position: relative;
    display: inline-block;
}

span:hover {
    cursor: pointer;
}

span:hover:before {
    content: attr(data-content);
    background-color: #2085c5;
    border-radius: 3px;
    color: #fff;
    padding: 10px;
    position: absolute;
    left: 100%;
    top: -70%;
    margin-left: 8px;
    white-space: pre;
}

span:hover:after {
    content: "";
    position: absolute;
    width: 0;
    height: 0;
    border-right: 8px solid #2085c5;
    border-top: 8px solid transparent;
    border-bottom: 8px solid transparent;
}
</style>
```

### 鼠标悬停提示
```html
<div id="div" class="demo">
    <ol class="list">
        <li><i></i><a href="#" data-title="铁路部门：火车票丢失需“先补买再退款”">铁路部门：火车票丢失需“先补买再退款”</a></li>
        <li><i></i><a href="#" data-title="新修订新疆宗教事务条例施行 禁止宣扬极端思想">新修订新疆宗教事务条例施行 禁止宣扬极端思想</a></li>
        <li><i></i><a href="#" data-title="长征火箭从天而降瞬间被拍下 民众围观(图)">长征火箭从天而降瞬间被拍下 民众围观(图)</a></li>
        <li><i></i><a href="#" data-title="呼格父母:今后不再上访 打算先给儿子买个好墓地">呼格父母:今后不再上访 打算先给儿子买个好墓地</a></li>
        <li><i></i><a href="#" data-title="中国驻印尼使馆:印尼未对持普通护照中方人员免签">中国驻印尼使馆:印尼未对持普通护照中方人员免签</a></li>
        <li><i></i><a href="#" data-title="陕西圈养大熊猫发生犬瘟热  一8岁大熊猫死亡">陕西圈养大熊猫发生犬瘟热  一8岁大熊猫死亡</a></li>
    </ol>
</div>

<style type="text/css">
.demo{ width:300px;text-align:left; padding:0; }
.list{ color:#2a6496;}
.list li{ position:relative; width:320px;list-style:decimal inside;line-height:35px;white-space:nowrap;overflow:hidden;text-overflow:ellipsis;}
.list a{font-size:18px;}
.tip{position:absolute;left: 0;display:none;color:#fff;background:rgba(0,0,0, 0.6); line-height:25px;font-size:12px;padding:2px 5px;border-radius:5px;z-index:1}
</style>

<script>
    (function(){
        var div = document.getElementById('div');
        var links = div.getElementsByTagName('a');
        var tip = createTip();
        for(var i = 0; i < links.length; i++){
            links[i].onmousemove = function(e){
                e = e || event;
                var title = this.getAttribute('data-title');
                tip.innerHTML = title;
                tip.style.display = 'block';
                tip.style.left = e.clientX + 10 + 'px';
                tip.style.top = e.clientY + 15 + 'px';
            }
            links[i].onmouseout = function(e){
                tip.style.display = 'none'
            }
        }
        function createTip(){
            var div = document.createElement('div');
            div.className = 'tip';
            return document.body.appendChild(div);
        }
    }());
</script>
```

### 轮播图的实现（常见于 banner）

### 瀑布流的实现（常见于图文加载）

### 让 Chrome 支持小于 12px 的文字
```html
<p><span>chrome 测试 10px</span></p>

<style>
p span{
	font-size:10px;
	-webkit-transform:scale(0.8); // transform:scale() 进行缩放
	display:block;
}
</style>
```

### textarea 禁止拖动
```css
resize: none;
// none：用户不能操纵机制调节元素的尺寸；
// both：用户可以调节元素的宽度和高度；
// horizontal：用户可以调节元素的宽度；
// vertical：让用户可以调节元素的高度；
// inherit：默认继承。
```

### footer 底部固定
```html
<div class='footer'>footer</div>

<style>
*{
  margin:0;
  padding:0;
}
body{
  height:2000px;
}
.footer {
  width:100%;
  height:30px;
  position: fixed;
  bottom:0;
  text-align:center;
  line-height:30px;
  background: #00CCCC;
}
</style>
```

### 透明度效果
```css
opacity: 0; 
cursor:pointer;  
-ms-filter:"progid:DXImageTransform.Microsoft.Alpha(Opacity=0)";
filter: alpha(opacity=0);
```

### 模糊滤镜效果
```css
-webkit-filter: blur(3px);
-moz-filter: blur(3px);
-o-filter: blur(3px);
-ms-filter: blur(3px);
filter: blur(3px);
```

### 阴影效果
```css
-webkit-box-shadow: 0 1px 1px rgba(0,0,0,.2);
-moz-box-shadow: 0 1px 1px rgba(0,0,0,.2);
box-shadow: 0 1px 1px rgba(0,0,0,.2);
```

### 给 input 的 placeholder 设置颜色
```css
::-webkit-input-placeholder { /* WebKit browsers */
    color:    #999;
}
:-moz-placeholder { /* Mozilla Firefox 4 to 18 */
    color:    #999;
}
::-moz-placeholder { /* Mozilla Firefox 19+ */
    color:    #999;
}
:-ms-input-placeholder { /* Internet Explorer 10+ */
    color:    #999;
}
```

### css 固定宽高 DIV 内部元素垂直居中
```html
<div class="outer">
    <div class="inner">haorooms内部内容</div><div class="v">cssHack</div>
</div>

<style>
* {
    margin: 0;
    padding: 0;
}
.outer {
    background-color: #ccc;
    font-size: 24px;
    height: 350px;
    text-align: center;
    overflow: hidden;
    width: 280px;
}
.outer  .inner,
.outer  .v {
    display: inline-block;
    zoom: 1;*display: inline; /* 用于触发支持IE67 inline-block */
}
.outer  .inner {            
    line-height: 1.8;
    padding: 0 4px 0 5px;
    vertical-align: middle;
    width: 262px;           
}
.outer  .v {
    line-height: 350px;
    text-indent:-9999px;
    width: 1px;         
}
</style>
```

### 一次性禁用所有表单元素
```html
<form>
    <fieldset disabled>
        <legend>完成您的购物订单</legend>
        <...>
    </fieldset>
</form>

<style>
fieldset[disabled] {
   -ms-pointer-events: none;
   pointer-events: none;
}
</style>
```

### form 表单两端对齐
```html
<div class="box">
    <div class="test">姓 名</div>
    <div class="test">所 在 地</div>
    <div class="test">工 作 单 位</div>
</div>

<style>
.test {
    text-align:justify;
    text-justify:distribute-all-lines;/*ie6-8*/
    text-align-last:justify;/* ie9*/
    -moz-text-align-last:justify;/*ff*/
    -webkit-text-align-last:justify;/*chrome 20+*/
}
@media screen and (-webkit-min-device-pixel-ratio:0){/* chrome*/
    .test:after{
        content:".";
        display: inline-block;
        width:100%;
        overflow:hidden;
        height:0;
    }
}
</style>
```

### 移除 HTML5 input 在 type="number" 时的上下小箭头
```css
// chrome
input::-webkit-outer-spin-button,input::-webkit-inner-spin-button{
  -webkit-appearance: none !important;
  margin: 0; 
}
// firefox
input[type="number"]{-moz-appearance:textfield;}
```

### 禁止用户选择复制
```html
<div unselectable="on" onselectstart="return false;">
    版权所有哦！
</div>

<style>
div {
	-webkit-user-select: none;
	-ms-user-select: none;
	-moz-user-select: none;
	-khtml-user-select: none;
	user-select: none;
}
</style>
```

### 禁止打开控制台
```js
document.onkeydown = document.onkeyup = document.onkeypress = function(event) {
    var e = event || window.event || arguments.callee.caller.arguments[0];

    if (e && e.keyCode == 123) {
            e.returnValue = false;
            return (false);
    }
}
```

### 让网站快速变灰
```html
<!-- 由于 FLASH 动画的颜色不能被 CSS 滤镜控制，可以在 FLASH 代码的和之间插入如下代码 -->
<param value="false" name="menu"/>
<param value="opaque" name="wmode"/>

<style>
html {
   filter: grayscale(100%);//IE浏览器
  -webkit-filter: grayscale(100%);//谷歌浏览器
  -moz-filter: grayscale(100%);//火狐
  -ms-filter: grayscale(100%);
  -o-filter: grayscale(100%);
  filter:progid:DXImageTransform.Microsoft.BasicImage(grayscale=1);
  -webkit-filter: grayscale(1);//谷歌浏览器
}
</style>
```

