
### Css 颜色体系
Css 颜色的表现方式主要有如下几种：  
色彩关键字、transparent（透明）、currentColor（原始 color 值）、rgb() 与 rgba()、hsl() 与 hsla()。  

### 色彩关键字
色彩关键字表示一个具体的颜色值，且它不区分大小写。譬如 color:red 的 red 即是一个色彩关键字。  

 CSS 标准 2，一共包含了 17 个基本颜色，分别是：  
 ![css2 基本色](./image/css2color.jpg)  

 CSS3，色彩关键字得到了极大的扩充，达到了 147 个。  
 [戳我查看 CSS 色彩关键字](https://developer.mozilla.org/zh-CN/docs/Web/CSS/color_value)   

 `主要：未知的关键字会让 CSS 属性无效。`  

### 可以设置颜色的属性
文本的颜色 color:red  
元素的背景色 background-color:red （包含各类渐变）  
元素的边框 border-color:red  
元素的盒阴影或文字阴影 box-shadow:0 0 0 1px red | text-shadow:5px 5px 5px red  
运用在一些滤镜当中 filter: drop-shadow(16px 16px 20px red)  
\<hr> 水平线的颜色  

一些无法直接设置，但是可以被得到或者继承当前元素 currentColor 的属性：  
\<img> 的 alt 文本。  
ul 列表项的小点  
```html
<div class="container">
  <ul>
    <li>list-type color</li>
  </ul>
  <img src="error/error.img" alt="错误的图片">
  <hr class="hr1"/>
  <hr class="hr2"/>
  <hr class="hr3"/>
</div>

<style>
.container{
  color:deeppink;
}

hr.hr1{
  color:red;
}
hr.hr2{
  background-color:red;
}
hr.hr3{
  border-top:1px solid red;
}
</style>
```

### transparent
transparent 的字面意思就是透明。它用来表示一个完全透明的颜色，即该颜色看上去将是背景色。也可以理解为它是 rgba(0,0,0,0) 的简写。  
`注意：在 CSS3 之前，transparent 关键字不是一个真实的颜色，只能用于 background-color 和 border-color 中，表示一个透明的颜色。而在支持 CSS3 的浏览器中，它被重新定义为一个真实的颜色，transparent 可以用于任何需要 color 值的地方，像 color 属性。`  

我们可以利用 transparent 绘制三角形。  
利用一个高宽为 0 的 div，设置它的 border，当任意三边的 border 颜色为 transparent 时，则可以得到任意朝向的一个三角形。  
```html
<div class='normal'></div>
<div class='triangle'></div>

<style>
.normal {
  width:0px;
  height:0px;
  margin:20px auto;
  border-top:50px solid yellowgreen;
  border-bottom:50px solid deeppink;
  border-left:50px solid bisque;
  border-right:50px solid chocolate;
}

.triangle {
  width:0px;
  height:0px;
  margin:20px auto;
  border-bottom:50px solid deeppink;
  border-left:50px solid transparent;
  border-right:50px solid transparent;
}
</style>
```

transparent 还可以用在 border，实现增大点击热区；用在 background，利用透明渐变，实现各种美妙的图形；配合 box-shadow ，在文本上运用 transparent，可以营造出一种文字发光的效果等等。  

### currentColor
currentColor 表示当前颜色，它来源自属性或者继承于它的父属性，即为当前 CSS 标签所继承或设定的文本颜色。  
```html
<div>currentColor</div>

<style>
div {
  width:200px;
  line-height:60px;
  margin:50px auto;
  font-size:24px;
  text-align:center;
  color:deeppink;
  border:1px solid currentColor;
  box-shadow:0px 0px 1px 1px currentColor;
}
</style>
```
如上，只在 color 里写了颜色，在 border 和 box-shadow 中使用了 currentColor 属性。  

currentColor 是 CSS3 新增的，在老版本浏览器下是无法识别的。但也有特例：
```html
<div>currentColor</div>


<style>
div{
  width:200px;
  line-height:60px;
  margin:50px auto;
  font-size:24px;
  text-align:center;
  color:#ff00ff;
  border:1px solid;
  box-shadow:0px 0px 1px 1px;
}
</style>
```
如上，只在 color 里写了颜色，border 的值为 1px solid，box-shadow 也是，并没有带上颜色值，但是依然表现为了 currentColor 的值。这是因为边框颜色和阴影颜色默认就是当前盒子的文本颜色，其中 border 兼容性很好，可以支持到 IE6 。  

border 和 box-shadow 是特例，但不是所有需要填写颜色值的属性不填写都会默认继承文本的值的。  

元素中将会得到或者继承元素 color 值有：  
元素的文本内容  
文本的轮廓  
元素的边框  
元素的盒阴影  
filter:drop-shadow()  
\<img> 的 alt 文本。也就是，当无法显示图像时，代替图像出现的文本，会继承这个颜色值。
列表项的小黑点和边框   
一些浏览器（比如 Chrome）水平线（ \<hr>）的边框颜色。（没有边框的话，颜色就不会受影响）。  

### rgb() 与 rgba()
rgb() 表示颜色的红 - 绿 - 蓝（red-green-blue (RGB)）模式，rgba() 多一个 a ，表示其透明度，取值为 0-1。

通常，十六进制符号 \#RRGGBB 中，RR 表示红色的深浅，GG 表示绿色的深浅，BB 表示蓝色的深浅。取值都是从 00 - FF，值越大表示该颜色越深。那么，如果采用 rgb(RR,GG,BB) 的写法，RR 取值 0~255 或者百分比，255 相当于 100%，和十六进制符号里的 F 或 FF。
`注意：颜色的叠加原理中，#FF00FF 红蓝叠加表示紫色， #FFFF00 红绿叠加表示黄色， #00FFFF 蓝绿叠加表示青色。`

### hsl() 与 hsla()
hsl() 被定义为色相 - 饱和度 - 明度（Hue-saturation-lightness），hsla() 多一个 a ，表示其透明度，取值为 0-1。  
hsl 相比 rgb 的优点是更加直观：你可以估算你想要的颜色，然后微调。它也更易于创建相称的颜色集合。  
```
色相（H）是色彩的基本属性，就是平常所说的颜色名称，如红色、黄色等。
饱和度（S）是指色彩的纯度，越高色彩越纯，低则逐渐变灰，取 0-100% 的数值。
明度（V），亮度（L），取 0-100%。
```

比如，使用 hsl 展示按钮 hover 和 active 的样式，只需改动 hsl 颜色值的第三个值达到了我们希望的效果。  
```html
<div>Btn</div>

<style>
div {
  width:140px;line-height:48px;
  text-align:center;
  margin:50px auto;
  color:#333;
  cursor:pointer;
  background:hsl(200, 60%, 60%);
  border:20px solid transparent;
  background-clip: padding-box;
}

div:hover {
  background:hsl(200, 60%, 50%);
  background-clip: padding-box;
}

div:active {
  background:hsl(200, 60%, 70%);
  background-clip: padding-box;
}
</style>
```

### rgb 到 hsl 的转换
在开发阶段我们只有一个 rgb 值，但是希望转换成 hsl 值，使用 chrome 开发者工具可以很便捷的做到，我们只需要选中我们想转换的颜色值，按住键盘左 shift，点击这个颜色表示框，即可进行转换。  
