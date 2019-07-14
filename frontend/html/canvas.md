
### 什么是 canvas
canvas 中文翻译就是” 画布”，是 HTML5 中重要的元素，和 audio、video 元素类似完全不需要任何外部插件就能够运行。  
canvas 提供了强大的图形的处理功能 (绘制，变换，像素处理…)，但是 canvas 元素本身并不绘制图形，它只是相当于一张空画布。如果开发者需要向 canvas 上绘制图形，则必须使用 JavaScript 脚本进行绘制。  

### canvas 能做什么
基础图形的绘制  
文字的绘制  
图形的变形和图片的合成  
图片和视频的处理  
动画的实现  
小游戏的制作  

### 支持的浏览器
大多数现代浏览器都是支持 Canvas 的，比如 Firefox, safari, chrome, opera 的最近版本以及 IE9 都支持（IE8 及以下不支持 HTML5, 但是我们可以进行提示用户更新到最新的版本）。  

### canvas 基本概念
定义 canvas 元素与定义其他普通元素并无任何不同，它除了可以指定 id, style ,class ,hidden 等通用属性之外，还可以设置 width 和 height 两个属性。  
在网页中定义 canvas 元素之后，它只是一张空白的画布，想要在画布上绘画，一定要经过下面几步：  
1、获取 canvas 元素对应的 DOM 对象，这必须是一个 canvas 对象  
2、调用 canvas 对象的 getContext() 方法，该方法返回一个 canvasRenderingContext2D 对象，该对象可以绘制图形  
3、调用 canvasRenderingContext2D 对象的方法进行绘图  
```html
<canvas id="canvas_demo" style="background-color: #FF0;"></canvas>

<script>
var canvas_demo = document.getElementById("canvas_demo");
var ctx = canvas_demo.getContext("2d");
// canvas 设置宽高，实际是把 canvas 默认的 300 * 150 的图片强行拉伸压缩，会导致问题
// canvas_demo.width = "500";
// canvas_demo.height = "500";

// 开始绘制
ctx.beginPath();

// 左上角是圆心，横轴向右是正，纵轴向下是正
// 设置绘制起点
ctx.moveTo(0,0);

// 设置绘制下一个点
ctx.lineTo(200,100);

// 设置绘制下一个点
ctx.lineTo(150,150);
    
// 结束绘制
ctx.closePath();

// 设置线的宽度
ctx.lineWidth = 10;

//设置绘制的样式
ctx.strokeStyle = "red";

//绘制点之间的线路
ctx.stroke();
    
// 设置填充样式
ctx.fillStyle = "green";

// 填充当前视图
ctx.fill();
</script>
```

