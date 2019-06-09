### 一 CSS3简介
CSS3 在 CSS2 基础上，增强或新增了许多特性，目前浏览器支持程度差，需要添加私有前缀，但是移动端支持率较高。
CSS3 手册使用贴士：
```
{}      表示范围					[]		表示全部可选项  
||		表示或者        			|		表示多选一
？      表示0个或者1个			*		表示0个或者多个
```

### 二 CSS3 选择器
#### 2.1 常见选择器
CSS3 选择器与 jQuery 中提供的大部分选择器类似，以前常用的选择器：
```
div         标签选择器
.box        类名选择器
div p       后代选择器
div.box     交集选择器
div,p,span  并集选择器
div>p       子代选择器
*           通配符选择器
div+p       选中div后面的第一个p
div~p       选中div后面所有p
```

#### 2.2 属性选择器
其特点是通过属性来选择元素，具体有以下 5 种形式：
```
div[attr]				  	属性名 = attr 的 div
div[attr="value"]		  	属性 attr 值为 value 的 div
div[attr~="value"]			属性值包含 value 字符
div[attr*="value"]			属性值包含 value 字符并且在“任意”位置
div[attr^="value"]			属性值包含 value 字符并且在“开始”位置
div[attr$="value"]			属性值包含 value 字符并且在“结束”位置
div[attr|="value"]			属性值是 value 或以 “value-” 开头的值（比如说zh-cn）
```

#### 2.3 伪类选择器
##### 2.3.1 以前的伪类选择器
```
:link、:active、:visited、:hover
```

##### 2.3.2 结构伪类
以某元素相对于其父元素或兄弟元素的位置来获取元素：
```
E:first-child			第一个子元素
E:last-child			最后一个子元素
E:nth-child(n) 		    第 n 个子元素，计算方法是 E 元素的全部兄弟元素；
E:nth-last-child(n) 	同 E:nth-child(n) 相似，只是倒着计算；
E:nth-of-type(n)  	    表示E父元素中的第 n 个字节点，且类型为 E
E:nth-last-of-type(n)   表示E父元素中的第 n 个字节点，且类型为 E,从后向前计算
E:empty 			    没有任何子节点（包括空格）的E元素,(子节点包含文本节点)
```
注意：
- n 遵循线性变化，其取值 0、1、2、3、4、... 但是当 n<=0 时，选取无效。
- n 可是多种形式：nth-child(2n)、nth-child(2n+1)、nth-child(-1n+5) 等；

代码示例：
```css
li:nth-child(2n-1){           // 选中所有的奇数的 li
    color: red;
}
li:nth-child(7n){             // 选中所有的 7 的倍数的 li 
    color: red;
}  
li:nth-child(-1n+5){         // 选中前面 5 个
    color: red;
}
li:nth-last-child(-1n+5){    // 选中后面 5 个
    color: red;
}
li:nth-child(even){          // 所有的偶数
    color:red
}
li:nth-child(odd){           // 所有的奇数
    color:blue;
}
```

##### 2.3.3 input 相关伪类
```
结合锚点进行使用，处于当前锚点的元素会被选中
E:target 
E:enabled  
E:disabled
E:checked
```

##### 2.3.4 :not
```html
<style>
p:not(.p0) {
    	background-color: #2aabd2;
}
</style>
<p>1111</p>
<p>1111</p>
<p class="p0">1111</p>
```

#### 2.4 伪元素选择器
CSS2 中 E:before 或者 E:after，是属于伪类的，并且没有伪元素的概念，CSS3 中提出伪元素的概念，如 E::before 和 E::after，并且归属到了伪元素当中，伪类里就不再存在 E:before 或者 E:after 伪类，这样做的目的是用来做兼容处理。

E::before、E::after 是一个行内元素，需要转换成块元素，当然这里只写一个冒号浏览器仍然能识别。  
E::first-letter 文本的第一个字母或字（如中文、日文、韩文等）；  
E::first-line 文本第一行；  文本第一行高亮；  
E::selection 可改变选中文本的样式；  

":" 与 "::" 区别在于区分伪类和伪元素  

### 三 盒模型
#### 3.1 CSS3 盒模型简介
CSS3 中可以通过 box-sizing 属性来指定盒模型，根据不同的盒模型，采用不同的盒子大小计算方式，优点是兼容性好。
```
box-sizing		指定盒模型类型
                content-box     标准模型(默认)：盒子宽度 = width+border+padding
                border-box      怪异模型：盒子宽度 = width  
```
同时CSS3引入了私有化前缀机制，即针对不同内核采用不同的属性前缀。
```
-webkit-		webkit内核 Safari and Chrome 
-moz-			火狐
-ms-			IE
-o-				欧朋
案例：
-webkit-box-reflect:below
```

#### 3.2 弹性盒模型
```html
<style>
    .box {
        height: 100px;
        border: solid 1px #2aabd2;
        padding: 10px;
        display: -webkit-box;
    }
    .box div {
        width: 100px;
        height: 100px;
        background-color: greenyellow;
        border: 1px solid #2aabd2;
    }
    .box div:nth-of-type(1) {
        -webkit-box-ordinal-group: 2;       /* 设置具体顺序 */
        -webkit-box-flex: 1;                /* 设置伸缩比例 */
    }
    .box div:nth-of-type(2) {
        -webkit-box-ordinal-group: 3;      
        -webkit-box-flex: 2;
    }
    .box div:nth-of-type(3) {
        -webkit-box-ordinal-group: 1;      
        -webkit-box-flex: 3;
    }
</style>

<div class="box">
    <div style="background-color: red ">第一个</div>
    <div style="background-color: orange ">第二个</div>
    <div style="background-color: green ">第三个</div>
</div>
```
注意：
- 查看上述元素的位置顺序。
- 使用弹性盒模型的时候，盒子元素必须要加 display:box 或 display:inline-box，且属性名前要加上浏览器内核前缀：display: -webkit-box;
常用属性列表：
```
box-orient		定义盒模型的布局方向，默认
horizontal 	    水平显示（默认），
vertical		垂直显示
 
box-direction	元素排列顺序，
                normal      正序（默认），
                reverse     反序

box-ordinal-group 设置元素的具体位置

box-flex 		定义盒子的弹性空间
                子元素的尺寸 =
                盒子尺寸*子元素box-flex值 / 所有子元素box-flex属性值之和 

box-pack 		对盒子富裕的空间进行管理
                start   所有子元素在盒子左侧显示，富裕空间在右侧
                end 	所有子元素在盒子右侧显示，富裕空间在左侧
                center 	所有子元素居中
                justify 富余空间在子元素之间平均分布

box-align 		在垂直方向上对元素的位置进行管理
                star    所有子元素在据顶
                end     所有子元素在据底
                center  所有子元素居中
```

#### 3.3 阴影与倒影
```
阴影：
box-shadow:[inset] x y blur [spread] color
    参数一 	新增投影方式：inset（内投影）不设置（外投影）
            x、y：		阴影偏移
            blur：		模糊半径
            spread：	扩展阴影半径，先扩展原有形状，再开始画阴影
            color：		阴影颜色
倒影
box-reflect:below 10px（只支持-webkit-）
    参数一  倒影方式
            above|below|left|right;
    参数二	倒影距离
```

### 四 新增属性
#### 4.1 颜色相关属性
新增了 RGBA、HSLA 模式，其中的 A 表示透明度通道，即可以设置颜色值的透明度，相较 opacity，它们不具有继承性，即不会影响子元素的透明度。
RGBA 是代表 Red（红色） Green（绿色） Blue（蓝色）和 Alpha 的色彩空间。
```
rgba(255,0,0,0.1)  R、G、B 取值范围 0~255
H 	色调 		    取值范围 0~360，0/360 表示红色、120 表示绿色、240 表示蓝色
S 	饱和度 	        取值范围 0%~100%
L 	亮度 		    取值范围	0%~100%
A 	透明度 	        取值范围	0~1

关于透明度：
background-color:hsla(0, 23%, 56%, 1)
opacity 只能针对整个盒子设置透明度，子盒子及内容会继承父盒子的透明度；
transparent 不可调节透明度，始终完全透明
background-color: transparent;
使用 rgba 来控制颜色，相对 opacity ，不具有继承性

```
#### 4.2 文本相关属性
文本阴影： text-shadow
```
案例一：
    text-shadow: 3px 4px 5px #ccc
        3px     水平偏移量. 正值向右 负值向左
        4px     垂直偏移量. 正值向下 负值向上
        5px     模糊度，模糊度不能为负值，值越大越模糊
        #ccc    设置对象阴影的颜色.
案例二：
    text-shadow:2px 2px 0px red, 2px 2px 4px green;
        阴影叠加：先渲染后面的，再渲染前面的
好玩的案例：
层叠：
    color:red; font-size:100px; font-weight:bold; text-shadow:2px 2px 0px white, 4px 4px 0px red;
光晕：
    color:white; font-size:100px; text-shadow:0 0 10px #fff, 0 0 20px #fff, 0 0 30px #fff, 0 0 40px #ff00de, 0 0 70px #ff00de, 0 0 80px #ff00de, 0 0 100px #ff00de, 0 0 150px #ff00de;
火焰文字：
    text-shadow: 0 0 20px #fefcc9, 10px -10px 30px #feec85, -20px -20px 40px #ffae34, 20px -40px 50px #ec760c, -20px -60px 60px #cd4606, 0 -80px 70px #973716, 10px -90px 80px #451b0e; font-family:Verdana, Geneva, sans-serif; font-size:100px; font-weight:bold; color:white;
```
文字描边：
```
webkit-text-stroke:宽度 颜色    （只有webkit内核浏览器才支持文字描边）
```
文字排版：
```
direction:rtl;unicode-bidi:bidi-override;
定义文字排列方式(全兼容),rtl 从右向左排列,ltr 从右向左排列
注意要配合unicode-bidi 一块使用，才能实现排序，没有这个属性，那么文字只会走到右边或者左边，而不会改变顺序。
```
文字溢出：
```
text-overflow 定义省略文本的处理方式
clip  无省略号
ellipsis 省略号 (注意配合overflow:hidden和white-space:nowrap一块使用) 
```
文字字体：
```
@font-face {
    font-family: ‘myziti’;			/*字体名*/
    src: url('111-webfont.eot');
    src: url('111-webfont.eot?#iefix') format('embedded-opentype'),
         url('111-webfont.woff') format('woff'),
         url('111-webfont.ttf') format('truetype'),
         url('111-webfont.svg#untitledregular') format('svg');
    font-weight: normal;
    font-style: normal;

}
使用：
p { font-family:‘myziti’}

转换字体格式生成兼容代码
http://www.fontsquirrel.com/fontface/generator

```

#### 4.3 边框新增属性
边框圆角 border-radius: 
```
border-radius: 1-4个数字 / 1-4个数字
	前面是水平轴半径，后面是垂直轴半径，不给“/”则水平和垂直一样（即正圆圆角）。
    值的单位可以是px，%等。

1个：都一样
border-radius: 一样
2个：对角
border-radius: 左上&右下    右上&左下
3个：斜对角
border-radius: 左上    右上&左下    右下
4个：全部，顺时针
border-radius: 左上    右上    右下    左下

```
边框背景：
```
border-image 					设置边框的背景图片
border-image-source:url(“”) 	设置边框图片的地址   
border-image-slice:27,27,27,27  浏览器会自动去裁剪图片.
border-left-width:20px;      	指定边框的宽度
border-image-repeat: stretch;   边框平铺的样式
stretch                         拉升  
round                           会自动调整缩放比例
repeat                          重复
```
边框阴影：
box-shadow  与 text-shadow 用法差不多
1、水平偏移量 正值向右 负值向左；
2、垂直偏移量 正值向下 负值向上；
box-shadow: 5px 5px 27px red, -5px -5px 27px green;
3、模糊度是不能为负值；
4、inset 可以设置内阴影；
 
设置边框阴影不会改变盒子的大小，即不会影响其兄弟元素的布局。
可以设置多重边框阴影，实现更好的效果，增强立体感。

#### 4.4 背景新增属性
背景控制：
```
background-size: 500px 500px;		控制背景大小，也可以是%
background-size:cover;				完全覆盖盒子，但不能保证是否完整显示
background-size:contain				背景图最大化在盒子中等比例显示，不保证铺满
```
cover会自动调整缩放比例，保证图片始终填充满背景区域，如有溢出部分则会被隐藏。
整个背景图片完整显示在背景区域.
contain会自动调整缩放比例，保证图片始终完整显示在背景区域。
 
注意：背景图默认从padding就开始平铺了，为了让背景从内容盒子才开始平铺显示，可以设置背景原点，backgound-origin:padding-box，将默认值改为：content-box
```
background-origin ： border | padding | content 
border-box： 	从border区域开始显示背景。 
padding-box： 	从padding区域开始显示背景。 
content-box： 	从content区域开始显示背景。
```
背景剪裁：
```
background-clip: padding-box  超出padding-box的裁剪，同理有border-box，content-box
```
多背景：
```
background: url() no-repeat left top,
            url() no-repeat left top;
```
背景渐变：
```
线性渐变:
background:linear-gradient(
to right 表示方向 (left,top,right,left ,也可以使用度数)
         yellow,  渐变起始颜色
green   渐变终止颜色
)
不写渐变方向，默认从上往下渐变

径向渐变:
radial-gradient径向渐变指从一个中心点开始沿着四周产生渐变效果
background: radial-gradient(
           150px  at  center,
           yellow,
           green
);  

```

#### 4.5 缩放
缩放属性：resize:both;overflow:auto;		
both 			水平垂直都可以缩放
horizontal 		只有水平方向可以缩放
vertical 		只有垂直方向可以缩放
注意：一定要配合overflow:auto 一块使用

#### 4.6 遮罩
```html
<style>
    .box {
        width: 800px;
        height: 600px;
        background: url(***.jpg) no-repeat;
        background-size: 100% 100%;
        border: 10px solid #000;
        -webkit-mask: url(mask.png) no-repeat;
        transition: 1s;
    }

    .box:hover {
        -webkit-mask-position: 100% 100%;
    }
</style>
```

### 五 过渡 transition
过渡是 CSS3 中具有颠覆性的特征之一，可以实现元素不同状态间的平滑过渡（补间动画），经常用来制作动画效果。  
补间动画：自动完成从起始状态到终止状态的的过渡，不用管中间的状态  
帧动画：扑克牌切换，通过一帧帧的画面按照顺序和速度播放，如电影  
```html
<style>
    div {
        height: 100px;
        width: 100px;
        background-color: #5cb85c;
    }
    div:hover {
        height: 500px;
        width: 500px;
        background-color: #2aabd2;
        /*动画完成时间：可以不指定过渡对象，或指定其中某项/多项*/
        transition: height 2s, width 3s,background-color 5s;     
    }
</style>
```

常用属性（注意属性可以连写）：
```
transition-property         要运动的样式  （all || [attr] || none）
transition-duration         运动时间
transition-delay            延迟时间
transition-timing-function  运动形式 
                                ease：（逐渐变慢）默认值
                                linear：（匀速）
                                ease-in：(加速)
                                ease-out：（减速）
                                ease-in-out：（先加速后减速）
cubic-bezier                    贝塞尔曲线（ x1, y1, x2, y2 ）          http://matthewlein.com/ceaser/ 
transition:hover            鼠标移入时动画

transition的end事件：
Webkit内核：   obj.addEventListener('webkitTransitionEnd',function(){},false);
firefox内核：  obj.addEventListener('transitionend',function(){},false);

每改变一次样式，都会触发一次end事件（如果绑定了该事件）。
```

### 六 转换 transform
转换是 CSS3 中具有颠覆性的特征之一，可以实现元素的位移、旋转、变形、缩放，甚至支持矩阵方式，配合过渡和动画，可以取代大量之前只能靠 Flash 才可以实现的效果。在 css3 当中，通过 transform(变形) 来实现 2d 或者 3d 转换,其中 2d 有：缩放、移动、旋转。  
transform-origin 的意思是变换的中心在哪里。

旋转：rotate(deg)   
可以对元素进行旋转，正值为顺时针，负值为逆时针。
```css
body:hover div {
    -webkit-transform: rotate(30deg); /* 顺时针旋转 */
}
```

斜切：
```css
/* XY 控制斜切的方向 */
-webkit-transform: skewX(30deg);
-webkit-transform: skewY(30deg);
-webkit-transform: skew(30deg，30deg);
```

缩放：scale(x, y) 
可以对元素进行水平和垂直方向的缩放，x、y 的取值可为小数，不可为负值；scale(x, y) transform: scale(1.5,1.5)

移动：
translate(x, y) 可以改变元素的位置，x、y 可为负值；  
x 相对于自身在水平方向移动。  
y 相对于自身在垂直方向移动。  
位置还原：transform: translate(0,0) rotate(0)  

矩阵：
2D 转换的效果其实都是通过矩阵来实现的。很多效果 IE 不支持，但是 IE 下也可以使用矩阵，我们可以通过矩阵函数来兼容 IE678。  
在标准浏览器中的矩阵函数  
matrix(1,0,0,1,0,0);  // 参数为初始化的值。  
在 IE 下的矩阵函数（没有标准下的后面 2 个参数，且与标准顺序不一致：  
且第二个参数、第三个参数位置与标准浏览器位置相反）：  
filter(“progid:DXImageTransform.Microsoft.Matrix( M11= 1, M12= 0, M21= 0 , M22=1,SizingMethod='auto expand'”);
```
参数 1：X 轴缩放值 
参数 2：Y 轴的倾斜度，计算方式为：Math.tan(xDeg/180*Math.PI)
（IE6789下为参数 3）
参数 3：X 轴的倾斜度，计算方式为：Math.tan(xDeg/180*Math.PI)
（IE6789 下为参数 2）
参数 4：Y 轴缩放值  
参数 5：X 位移
参数 6：Y 位移
``` 

实例代码：
```js
var box = document.getElementById('box');
box.onclick = function () {
    //X轴缩放0.5,Y轴缩放0.2,X轴斜切0.3,Y轴斜切0.4
   box.style.transform = 'matrix(0.5,0.4,0.3,0.2,0,0)';  //IE10及以上以及大部分高版本浏览器支持
   box.style.filter = "progid:DXImageTransform.Microsoft.Matrix( M11= 0.5, M12= 0.3, M21= 0.4 , M22= 0.2, SizingMethod='auto expand')";
    // box.style.WebkitTransform = 'matrix(0.5,0,0,0.2,0,0)';       //低版本chrome，以下同理
    // box.style.MozTransform = 'matrix(0.5,0*0.2,0,1*0.2,0,0)';
}
```

矩阵旋转：
通过矩阵实现旋转--注意 IE 下顺序
参数 1 = Math.cos(deg/180*Math.PI);   
参数 2 = Math.sin(deg/180*Math.PI);  
参数 3 = -Math.sin(deg/180*Math.PI);  
参数 4 = Math.cos(deg/180*Math.PI);  
但是在 IE 下旋转并不是围绕中心点来展开，需要做处理，案例如下：  
```html
    <style>
        .box{width:100px;height:100px;margin:30px auto; position:relative;border:1px solid #000;}
        #box{width:100px;height:100px;background:red; position:absolute;left:0;top:0;}
    </style>

    <div class="box">
        <div id="box"></div>
    </div>

    <script>
        var oBox=document.getElementById("box");
        var iDeg=0;
        setInterval(function(){
            iDeg++;
            toRotate(oBox,iDeg);
        },30);
        function toRotate(obj,iDeg)
        {
            var a=Math.round(Math.cos(iDeg/180*Math.PI)*100)/100;
            var b=Math.round(Math.sin(iDeg/180*Math.PI)*100)/100;
            var c=-b;
            var d=a;
            obj.style.WebkitTransform="matrix("+a+","+b+","+c+","+d+",0,0)";
            obj.style.MozTransform="matrix("+a+","+b+","+c+","+d+",0,0)";
            obj.style.transform="matrix("+a+","+b+","+c+","+d+",0,0)";
            obj.style.filter="progid:DXImageTransform.Microsoft.Matrix( M11="+a+", M12= "+c+", M21= "+b+" , M22="+d+",SizingMethod='auto expand')";
            obj.style.left=(obj.parentNode.offsetWidth-obj.offsetWidth)/2+"px";
            obj.style.top=(obj.parentNode.offsetHeight-obj.offsetHeight)/2+"px";
        }
    </script>
```

3D转换：
```
<style>
        #box {
            width: 400px;
            height: 400px;
            margin: 100px auto;
            border: 1px solid #2aabd2;
            transform-style: preserve-3d;    /* 建立 3D 空间 */
            perspective: 150px      /* 景深基点--站在多远观看*/
        }
        #mini {
            margin: 100px auto;
            width: 200px;
            height: 200px;
            background-color: #5cb85c;
            transition: 1s;
        }
        #box:hover #mini{
            transform: rotateX(180deg);
        }
    </style>

<div id="box">
    <div id="mini">666</div>
</div>

常见属性：
transform-style: preserve-3d;    建立 3D 空间             
perspective: 150px              景深基点--站在多远观看
backface-visibility:hidden          设置盒子背面是否隐藏


新增的一些函数：
rotateX()
rotateY()
rotateZ()
translateZ()        必须配合透视（即景深基点）使用
scaleZ()
```

### 七 动画 animation
```
动画是 CSS3 中具有颠覆性的特征之一，可通过设置多个节点来精确控制一个或一组动画，常用来实现复杂的动画效果。
代码示例：
div {
    width: 200px;
    height: 200px;
    background-color: #5cb85c;
    animation: 2S move;         /* 指定动画名称和时间 */
}
@keyframes move {
    50%{
        width: 1000px;
    }
}

动画问题：动画执行完毕后，盒子会回到原始状态，如果我们不想其回到原始状态，一般可以给其设定一个额外的 class。
    <style>
        .div {
        width: 100px;
            height: 100px;
            background-color: #5cb85c;
        }
        .move {
            animation: 2S move;         /* 指定动画名称和时间 */
            width: 500px;
        }
        @keyframes move {
            0%{
                width: 100px;
            }
            100%{
                width: 500px;
            }
        }
    </style>
<div class="div move"></div>

注意：参数 forward 可以让动画保持在最后一帧

keyFrames：即关键帧。只需指明两个状态，之间的过程由计算机自动计算

格式 1
@keyframes 动画名称
{
    动画状态
}

格式 2
@keyframes  动画名称
{
    from { background:red; }
    to { background:green; }
}
可以只有 to

animation：定义动画的执行时间和执行行为是谁（名字）。当然也有一些可选参数如：animation: 2S move ease-in;

常见属性与对应值：

animation-name          必要属性-动画名称（关键帧名称）
animation-duration      必要属性-动画持续时间
例如：
-webkit-animation-name: ‘miaov';
-webkit-animation-duration: 4s;

animation-timing-function   动画运动形式
linear          匀速
ease            缓冲
ease-in         由慢到快
ease-out        由快到慢
ease-in-out     由慢到快再到慢

animation-play-state            播放状态（running 播放 和 paused 暂停 ）

animation-delay             动画延迟（只是第一次）

animation-iteration-count       重复次数（infinite 为无限次）

animation-direction         动画是否重置后再开始播放
alternate           动画直接从上一次停止的位置开始执行
normal              动画第二次直接跳到0%的状态开始执行

cubic-bezier(number, number, number, number)：   
特定贝塞尔曲线类型，4 个数值需在 [0, 1] 区间内


end 函数：
obj.addEventListener('webkitAnimationEnd', function (){}, false);
```