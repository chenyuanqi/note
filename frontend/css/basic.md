
### 一 CSS 简介
CSS 即层叠样式表（Cascading  Style  Sheets），也称为级联样式表。  
CSS 书写位置：  
- 内嵌式：在 head 标签中直接书写  
```Html
<style type="text/css"></style>
```
- 外连式：
```Html
<link rel="stylesheet" href="demo.css">
```
- 行内式：通过给标签直接设置style属性来设置样式。

### 二 CSS 选择器
#### 2.1 常见基础选择器
```Html
<style type="text/css">
    * {}                        /* 通配符选择器，控制所有 */
    p { color: red; }           /* 标签选择器 */
    .test { color: green;}      /* 类选择器 */
    #oDiv { color: blue;}       /* id 选择器 */
</style>
<div class="test"></div>
<div id="oDiv"></div>
```
id 选择器：多个标签不要同时使用同一 id 选择器，一个标签只能调用一个 id。  

类选择器：多个标签可以同时调用一个类样式， 一个标签可以调用多个类样式。  
类选择器命名规范：  
- 不能以纯数字或者以数字开头定义类名
- 不推荐使用汉字定义类名
- 不能以特殊符号或者以特殊符号开头（“_”除外）定义类名
- 不建议使用标签名或者属性名定义类名

#### 2.2 标签指定选择器
特点关系：既.....又.....   如下案例使用 id p#id 也可以实现：
```Html
<style type="text/css">
   p.one{ color: red;}
</style>
........
<p class="one">one</p>
<p class="two">two</p>
```

#### 2.3 后代选择器
特点关系：标签之间必须属于嵌套关系，选择器之间用空格隔开，如下案例，选择 p 标签下的类名为 p1 的标签。
```Html
<style type="text/css">
   p #p1{ color: red;}
</style>
```

#### 2.4 并集选择器
```Html
<style type="text/css">
   div,p,span{ color: red;}
</style>
```

#### 2.5 属性选择器
```Html
<style>
    input[type=text][class="password"] {
        background-color: red;
    }
</style>

<input type="text" id="username">
<input type="password">
```

#### 2.6 伪类选择器
```
:link      /* 未访问的链接 */
:visited   /* 已访问的链接 */
:hover     /* 鼠标移动到链接上 */
:active    /* 选定的链接 */
```

### 三 CSS 与文字
#### 3.1 常见文字设置
```
设置颜色    color: red
            支持 rgb，颜色名，16 进制数值等
设置居中    text-align: center;  
            设置盒子内文本居中，不能让标签居中，当然我们可以把行内元素、行内块元素看成文本
设置大小    font-size: 16px;
            经常使用相对长度单位，如像素
设置字体    font-family:"宋体","微软雅黑";
            可以设置单个值，也可以设置多个，即浏览器不支持第一个字体，则会尝试下一个，如果都不支持，那么使用系统默认的字体。
设置加粗    font-weight
            直接设置数字        100-900（必须是100的整数倍，推荐使用700）
            直接设置设置英文：   bold   (700)   
                              normal (400，即文字字体正常显示)
                              其他：bolder、lighter
设置样式    font-style
            italic            文字斜体显示；
            normal            文字正常显示；
            oblique           文字倾斜显示；
整体设置    font:normal 12px/36px ' 微软雅黑 ';
            font 同时设置文字的几个属性，写的顺序有兼容问题，建议按照如下顺序写： font：是否加粗 字号 / 行高 字体
设置行高    line-height:24px;
设置下划线  text-decoration:none; 
设置首行缩进 text-indent:24px;
设置水平对齐 text-align:center; 
```
注意：
- <font color=red> 现在网页中普遍使用 14px+，尽量使用偶数字字号，ie6 等老式浏览器支持奇数会有 bug</font>
- 当需要设置英文字体时，英文字体名必须位于中文字体名之前。
- 如果字体名中包含空格、#、$等符号，则该字体必须加英文状态下的单引号或双引号，例如 font-family: "Times New Roman";
- 在 CSS 中设置字体名称，直接写中文是可以的。但是在文件编码（GB2312、UTF-8 等）不匹配时会产生乱码的错误。xp 系统不支持 类似微软雅黑的中文。可以使用英文代替： font-family:"Microsoft Yahei"，或者使用Unicode编码：font-family: "\5FAE\8F6F\96C5\9ED1"。
- 平时我们很少给文字加斜体，反而喜欢给斜体标签（em，i）改为普通模式

#### 3.2 属性连写
```
语法：
    选择器{font: font-style  font-weight  font-size/line-height  font-family;}
示范：
    p {  font: italic 700 30px 宋体; }
```
注意： 联写必须有 font-size 和 font-family(其他属性可以不写)，且顺序不能更换。

#### 3.3 行高 line-height
浏览器默认文字大小：16px，  
浏览器默认行高大小：18px。  
行高的定义：两行文本基线的距离，就是行高（这里才是正确量取行高的方法）  
行高 = 文字大小 + 上间距 + 下间距   
三者加起来如果等于盒子高度，正好文字居中！  

当行高值为父容器的高度时，可以让父容器中的文字垂直显示。
如果单独给一个元素设置行高：

| 行高单位 | 赋值 | 文字大小 | 行高值 |
| :------| ------: | :------: |:------: |
| px | 20px | 20px |20px|
| em | 2em | 20px |40px|
| % | 120%	| 20px | 24px |
| 不带单位 | 2 | 20px | 40px |
总结：当给一个独立的元素设置行高值的时候，除了以px为单位的行高值与文字大小无关，其他都与文字大小有关（与文字大小的积）。

盒子嵌套，给父元素设置行高值，子元素的行高问题：
|行高单位|设置行高|父文字|子文字|行高|
| :------| ------: | :------: |:------: |:------: |
||px|	20px|	20px|	30px|	20px
em|	2em|	20px|	30px|	40px
|%|	120%|	20px|	30px|	24px
|不带单位|	2	|20px	|30px|	60px
总结：行高可以继承。当父元素设置了行高值（不带单位除外），子元素的行高值都是父元素行高值乘以父元素文字大小。

总结图表：
| 给父元素设置行高 | 子元素行高结果 | 
| :------| ------: | 
| px | 行高=父元素行高|
| em | 行高=父元素文字大小*行高|
| % | 行高=父元素文字大小*行高 |
| 不带单位 | 行高=子元素文字大小*行高|


### 四 标签分类
#### 4.1 块级元素
代表：div  p  li   h1
特点：
- 元素独占一行显示，与宽度无关，即不给宽度直接占满一行
- 支持所有 CSS 命令（可以设置宽度和高度）
- 当嵌套一个块级元素，子元素如果不设置宽度，子元素的宽度为父元素的宽度。

#### 4.2 行内元素
代表:  span，a，font，strong
特点：
- 元素可以在同一行上显示
- 宽高由内容撑开
- 不支持上下 margin
- 代码换行会被解析

#### 4.3 行内块元素
代表： Image,  input(表单控件)
特点： 元素在一行上显示，可以设置宽高，没有宽度时由内容撑开。
案例：给下列的 div 设置 text-align:center 一样能居中，所以行内块元素可以当做文字来处理。
```html
<div style="text-align:center">
    <span>111</span>
</div>
```

#### 4.4 元素转换
display: inline 将元素转化为行内元素
display：inline-block   将元素转化行内块元素
display: block  将元素转化为块元素
注意：
inline-block 也可以实现类似 float 的效果，但是同一行的元素间有间隔。

### 五 CSS 特性
- 层叠性：某个元素同时出现了多个同级别的同名样式，那么书写在后面的样式将会覆盖之前的样式。
- 继承性：子元素会继承父元素的：颜色、大小、字体、行高。注意：a 标签不能继承父元素文字颜色（层叠掉了），h 标签不能继承父元素文字大小。
- 优先级：
默认样式 0 < 标签选择器 1 < 类选择器 10 < id选择器 100 < 行内样式 1000 < !important 1000以上  
注意：继承的权重为0；权重可以叠加

left 比 right 权重高。有 left 又有 right 的时候，执行 left 的值。
top 比 bottom 权重高。有 top 又有 bottom 的时候，执行 top 的值。

### 六 伪类
a:link{属性:值;}     超链接默认状态，与 a{} 一样  
a:visited{属性:值;}  超链接访问过后的样式  
a:hover{属性:值;}	鼠标放到超链接上的样式  
a:active{}      	超链接激活状态下的样式  
:focus{}            获取焦点的时候的样式  

### 七 CSS 背景
#### 7.1 背景色 背景图
background 属性是 css 中应用比较多，且比较重要的一个属性，它是负责给盒子设置背景图片和背景颜色的，background 是一个复合属性，它可以分解成如下几个设置项：  
> background-color 设置背景颜色  
> background-image 设置背景图片地址  
> background-repeat 设置背景图片如何重复平铺  
> background-position 设置背景图片的位置  
> background-attachment 设置背景图片是固定还是随着页面滚动条滚动  

实际应用中，我们可以用 background 属性将上面所有的设置项放在一起，而且也建议这么做，这样做性能更高，而且兼容性更好，比如：
```css
background: #00FF00 url (bgimage.gif) no-repeat left center fixed”;  
// #00ff00 是设置 background-color；  
// url (bgimage.gif) 是设置 background-image
// no-repeat 是设置 background-repeat
// left center 是设置 background-position
// fixed 是设置 background-attachment，各个设置项用空格隔开，有的设置项不写也是可以的，它会使用默认值  
```
注意：插入背景图一定要设置高度和宽度
```Html
div{
    background-color: red;
    background-image: url("1.jpg");
    width: 200px;
    height: 200px;
}
```

#### 7.2 背景平铺
repeat 		默认值   
no-repeat 	不平铺  
repeat-x  	横向平铺  
repeat-y 	纵向平铺  

#### 7.3 背景位置
设置具体值： left| right| top| bottom| cneter  
background-position: top right;	  
设置具体值时不区分顺序，分别为水平、垂直方向

#### 7.4 背景固定于滚动
background-attachment:fixed;  
fixed：图片固定   scroll：滚动（默认）  

#### 7.5 属性连写
没有数量限制和先后顺序限制：
```html
background: url("1.png") red no-repeat 30px 40px;
```

### 八 盒模型
#### 8.1 盒模型简介
元素在页面中显示成一个方块，类似一个盒子，CSS 盒子模型就是使用现实中盒子来做比喻。  
盒模型主要用来网页布局，盒子包括：border(边框)、内边距（padding）、外边距（margin）。
![盒模型](./image/box.jpg)  

#### 8.2 边框border
border-width: 边框宽度  
border-style: 边框样式，solid（实线）、dotted（点线）、dashed（虚线）  
border-color：边框颜色  
border书写方式：  
- border-left: 1px solid green  同理设置其他边框；
- border-top-color:green 		同理设置其他属性；
- border: solid 1px red;    联写
注意：border属性联写：没有先后顺序限制，边框颜色、宽度可以不写。  

#### 8.3 内边距 padding
内边距设置内容距离盒子边框之间的距离。  
padding-left: 左边距  
padding-right: 右边距  
padding-top: 上边距  
padding-bottom: 下边距  
属性联写：  
padding: 10px;上，右，下，左的距离为 10px  
padding: 10px 20px; 上下 10px，左右 20px  
padding: 10px 20px 30px;上 10px  左右 20px   下 30px  
padding: 10px 20px 30px 40px; 上， 右 ， 下， 左  

#### 8.4 盒子大小计算
盒子宽度 = 内容宽度+左右边框+左右内边距；  
注意：进行页面布局时，如果给盒子设置了内边距，对应的要将内容宽度或者高度减去相应的值。  

子盒子在父盒子宽度范围内，父盒子的 padding 不会影响子盒子大小。

#### 8.5 外边距 margin
外边距设置盒子与盒子之间的距离。属性同 padding。  
注意：当两个盒子垂直显示的时候，外边距以设置的最大值为准。  
行内元素只有左右 margin，没有上下 margin，同理 padding 也是。  
注意：当两个盒子垂直显示的时候，外边距以设置的最大值为准。行内元素只有左右 margin，没有上下 margin，同理 padding 也是。

### 九 浮动
#### 9.1 文档流与浮动
元素默认的显示方式（如块级元素独占一行）就是标准流（文档流）。    
文档流，是指盒子按照 html 标签编写的顺序依次从上到下，从左到右排列，块元素占一行，行内元素在一行之内从左到右排列，先写的先排列，后写的排在后面，每个盒子都占据自己的位置。  

浮动的目的是为了让块级元素能够在一行显示。  
浮动常用来解决文字环绕图片问题，也可用来制作导航栏、网页布局。  
浮动用法：float:left| right    元素默认是没有浮动的，即 float:none  
特点：  
- 设置了浮动的元素不占原来的位置（脱标）    
- 浮动可以行内元素转化为行内块元素

浮动的特性
> 1 浮动元素有左浮动 (float:left) 和右浮动 (float:right) 两种  
> 2 浮动的元素会向左或向右浮动，碰到父元素边界、浮动元素、未浮动的元素才停下来  
> 3 相邻浮动的块元素可以并在一行，超出父级宽度就换行  
> 4 浮动让行内元素或块元素自动转化为行内块元素  
> 5 浮动元素后面没有浮动的元素会占据浮动元素的位置，没有浮动的元素内的文字会避开浮动的元素，形成文字饶图的效果  
> 6 父元素内整体浮动的元素无法撑开父元素，需要清除浮动  
> 7 浮动元素之间没有垂直 margin 的合并  

#### 9.2 清除浮动
浮动后，后续的盒子会浮上来，经常采取的做法是：浮动的元素都被包裹在一个透明的父盒子中，父盒子只要拥有自己的宽高，那么就不会对整体布局造成影响。  
当然也可以选择清除浮动，清除浮动不是删除浮动，而是清除浮动对布局的影响。  

当子元素设置了浮动，父元素没有高度的时候，会造成页面布局混乱，这种情况下进行清除浮动。即解决父盒子高度为 0 的问题。  
如果父盒子没有设置高度，那么高度由子盒子撑开；  
如果子盒子这时使用了浮动，那么父盒子就无法撑开，不显示。  

清除浮动方式：  
方式一：谁出问题给谁加一个 clearfix 类名，使用clear:left|  right  | both  
```html
    <style>
        #div1 {
            width: 500px;
            background-color: red;
        }
        #div2 {
            width: 200px;
            height: 200px;
            background-color: green;
            float:right;
        }
        .clearfix {
            clear: both;
        }
    </style>
</head>
<body>
<div id="div1">
    <div id="div2"></div>
    <div class="clearfix"></div>
</div>
```

方式二：给父盒子设置 overflow:hidden 使用此属性用来触发 bfc。  
如果父盒子中有定位的元素，一般不推荐使用该种方式清除浮动，因为子盒子中的元素如果超出了父盒子的高度，超出部分会被切掉。  

方式三：使用伪元素清除浮动
```html
    <style>
        .div1 {
            width: 500px;
            background-color: red;
        }
        .div2 {
            width: 200px;
            height: 200px;
            background-color: green;
            float:right;
        }
        .clearfix:after {
            content: "";
            display: block;
            clear: both;
            height: 0;
            line-height: 0;
            visibility: hidden;
        }
        .clearfix {
            zoom: 1;    /*兼容IE*/
        }
    </style>
</head>
<body>
<div class="div1 clearfix">
    <div class="div2"></div>
</div>
```

方式四：使用双伪元素清除浮动（淘宝、小米采用）
```html
.clearfix:before,.clearfix:after{
    display: table;
    content: "";
}
.clearfix:after {
    clear: both;
}
.clearfix {
    zoom: 1;
}  

```

### 十 定位 position
#### 10.1 简介
定位就是元素位于定在某个位置，比如轮播图中左右控制按钮，比如某个区域内的不断变换的小图标，都可以使用定位实现。    
定位写法：
```css
postion: static;         /* 定位模式 */
left: 100px;            /* 边偏移 */
top: 100px;
```
边偏移：left, right,top,bottom（分别代表相对于父元素的左右上下距离）  

定位模式：
- static：静态定位，也称为自动定位，是元素的默认定位方式，即按照元素标准流的显示方式，静态定位约等于标准流，所以静态定位中偏移量的设置是无效的，一般用于清除定位。没有定位，元素出现在正常的文档流中，相当于取消定位属性或者不设置定位属性  
- relative：相对定位，
- absolute：绝对定位，
- fixed：固定定位，

#### 10.2 相对定位 relative
相对定位：设置了相对定位后，新的展示位置根据自己原来的位置 +或者- 定位位置,但是该元素在 CSS 中仍然是占据着原来的位置。元素所占据的文档流的位置不变，元素本身相对文档流的位置进行偏移。    
`注意：相对定位不能进行元素的模式转换`  

子元素设置绝对定位，父元素设置相对定位（子绝父相，也要看情况）
```html
 div {
    width: 200px;
    height: 200px;
    background: aqua;
    margin: 100px;
    position: relative;
    top: 10px;
    left: 10px;
}
```

#### 10.3 绝对定位 absolute
绝对定位 position:absolute; 类似浮动不再占据位置。元素脱离文档流，不占据文档流的位置，可以理解为漂浮在文档流的上方，相对于上一个设置了相对或者绝对或者固定定位的父级元素来进行定位，如果找不到，则相对于 body 元素进行定位。    
特点：
- 给单独的元素设置绝对定位，以浏览器左上角（body）为基准
- 给行内元素设置绝对定位后，该元素转化为了行内块元素
- 给盒子设置了绝对定位，该盒子不占位置（脱标）
- 当盒子发生嵌套关系的时候：
    - 父盒子没有定位，子盒子定位，以浏览器左上角为基准; 
    - 父盒子设置定位，子盒子设置定位，以父盒子左上角为基准。  

#### 10.4 固定定位 fixed
固定定位不占位置（脱标），将行内元素转化为行内块元素，且不会随着浏览器的滚动条滚动而变化。  
元素脱离文档流，不占据文档流的位置，可以理解为漂浮在文档流的上方，相对于浏览器窗口进行定位。  

#### 10.5 定位盒子层级关系 z-index
定位元素是浮动的正常的文档流之上的，可以用 z-index 属性来设置元素的层级。  

后定位的盒子的层级要高于前面定位的盒子的层级，使用 z-index 设置定位盒子之间的层级关系，z-index 可以取负数。  
只有定位的盒子（除了static）才有 z-index；  
如果都是绝对定位，他们默认的 z-index 都是0；

### 十一、其他
#### 11.1 元素溢出
当子元素的尺寸超过父元素的尺寸时，需要设置父元素显示溢出的子元素的方式，设置的方法是通过 overflow 属性来设置。  
```
overflow 的设置项： 
1、visible 默认值。内容不会被修剪，会呈现在元素框之外。
2、hidden 内容会被修剪，并且其余内容是不可见的，此属性还有清除浮动、清除 margin-top 塌陷的功能。
3、scroll 内容会被修剪，但是浏览器会显示滚动条以便查看其余的内容。
4、auto 如果内容被修剪，则浏览器会显示滚动条以便查看其余的内容。
5、inherit 规定应该从父元素继承 overflow 属性的值。
```

#### 11.2 元素特性
元素就是标签，布局中常用的有三种标签，块元素、内联元素、内联块元素。  

块元素  
块元素，也可以称为行元素，布局中常用的标签如：div、p、ul、li、h1~h6、dl、dt、dd 等等都是块元素，它在布局中的行为：  
> 支持全部的样式  
> 如果没有设置宽度，默认的宽度为父级宽度 100%  
> 盒子占据一行、即使设置了宽度  

内联元素  
内联元素，也可以称为行内元素，布局中常用的标签如：a、span、em、b、strong、i 等等都是内联元素，它们在布局中的行为：  
> 支持部分样式（不支持宽、高、margin 上下、padding 上下）  
> 宽高由内容决定  
> 盒子并在一行  
> 代码换行，盒子之间会产生间距  
> 子元素是内联元素，父元素可以用 text-align 属性设置子元素水平对齐方式，用 line-height 属性值设置垂直对齐方式

解决内联元素间隙的方法   
> 1、去掉内联元素之间的换行  
> 2、将内联元素的父级设置 font-size 为 0，内联元素自身再设置 font-size  

内联块元素  
内联块元素，也叫行内块元素，是新增的元素类型，现有元素没有归于此类别的，img 和 input 元素的行为类似这种元素，但是也归类于内联元素，我们可以用 display 属性将块元素或者内联元素转化成这种元素。它们在布局中表现的行为：  
> 支持全部样式  
> 如果没有设置宽高，宽高由内容决定  
> 盒子并在一行  
> 代码换行，盒子会产生间距  
> 子元素是内联块元素，父元素可以用 text-align 属性设置子元素水平对齐方式，用 line-height 属性值设置子元素垂直对齐方式  

这三种元素，可以通过 display 属性来相互转化，不过实际开发中，块元素用得比较多，所以我们经常把内联元素转化为块元素，少量转化为内联块，而要使用内联元素时，直接使用内联元素，而不用块元素转化了。  

display 属性  
display 属性是用来设置元素的类型及隐藏的，常用的属性有：  
> 1、none 元素隐藏且不占位置  
> 2、block 元素以块元素显示  
> 3、inline 元素以内联元素显示  
> 4、inline-block 元素以内联块元素显示  

#### 11.3 一些基本常识
在 css 中，后面的样式代码会覆盖之前的相同的样式代码。  
a 标签的四个伪类执行顺序时 :link，:visited，:hover，:active。  

移动端点击事件（或页面设置 viewpoint）会有延迟，原因时等待 300 ms 看用户是否点击还是双击缩放，解决方法是禁止缩放、设置默认视口宽度为设备宽度、设置 css touch-action:none、fastclick.js。  
