
### rem：移动 web 适配利器
rem（font size of the root element），根据网页的根元素来设置字体大小。和 em（font size of the element）的区别是，em 是根据其父元素的字体大小来设置，而 rem 是根据网页的跟元素（html）来设置字体大小的。  

现在大部分浏览器 IE9+，Firefox、Chrome、Safari、Opera ，如果我们不修改相关的字体配置，都是默认 html 标签显示 font-size 是 16px。  

对于 rem 来说，它可以用来做移动端的响应式适配。比如，给一个 p 标签设置 12px 的字体如下：
```css
p {
    font-size: 0.75rem; // 12÷16=0.75（rem）
}
```  
基本上使用 rem 这个单位来设置字体大小基本上是这个套路，好处是假如用户自己修改了浏览器的默认字体大小，那么使用 rem 时就可以根据用户的调整的大小来显示了。 但是 rem 不仅可以适用于字体，同样可以用于 width height margin 这些样式的单位。  

兼容性支持：
ios：6.1 系统以上都支持  
android：2.1 系统以上都支持  
大部分主流浏览器都支持  

rem 也不是万能的，比如当用作图片或者一些不能缩放的展示时，必须要使用固定的 px 值，因为缩放可能会导致图片压缩变形等。

### rem 屏幕适配
一般地，移动端适配的方法可以分为如下三种情况：  
> 1.简单一点的页面，一般高度直接设置成固定值，宽度一般撑满整个屏幕。  
> 2.稍复杂一些的是利用百分比设置元素的大小来进行适配，或者利用 flex 等 css 去设置一些需要定制的宽度。  
> 3.再复杂一些的响应式页面，需要利用 css3 的 media query 属性来进行适配，大致思路是根据屏幕不同大小，来设置对应的 css 样式。  

rem 可以兼顾到以上这些方面。  
```html
<div class="con"></div>
<style>
html {
    font-size: 16px; // 修改值即会影响到 .con 的宽高
}

.con {
    width: 10rem;
    height: 10rem;
    background-color: red;
}
</style>
 ```

rem 数值计算  
如果利用 rem 来设置 css 的值，一般要通过一层计算才行，比如如果要设置一个长宽为 100px 的 div，那么就需要计算出 100px 对应的 rem 值是 100 / 16 =6.25rem，这在我们写 css 中，其实算比较繁琐的一步操作了。  
所以，对于没有使用 sass 的工程：  
为了方便起见，可以将 html 的 font-size 设置成 100px，这样在写单位时，直接将数值除以 100 在加上 rem 的单位就可以了。  
对于使用 sass 的工程：  
前端构建中，完全可以利用 scss 来解决这个问题。  
```css
@function px2rem($px){
    $rem : 37.5px; // 基准值
    @return ($px/$rem) + rem;
}

.xxx{
	height: px2rem(90px);
    width: px2rem(90px);
}
```
rem 的基准值是根据我们所拿到的视觉稿来决定的，因为我们所写出的页面是要在不同的屏幕大小设备上运行的，所以我们在写样式的时候必须要先以一个确定的屏幕来作为参考，这个就由我们拿到的视觉稿来定。假如我们拿到的视觉稿是以 iphone6 的屏幕为基准设计的，Phone6 的屏幕大小是 375px，那么 rem = window.innerWidth/10（避免 html 的 font-size 太大，随意定义） 就得到 37.5px 了。  

动态设置 html 的 font-size  
通过不同的屏幕去动态设置 html 的 font-size，通常有两种办法：  
1、利用 css 的 media query 来设置  
```css
@media (min-device-width : 375px) and (max-device-width : 667px) and (-webkit-min-device-pixel-ratio : 2){
      html{font-size: 37.5px;}
}
```
2、利用 javascript 来动态设置 根据我们之前算出的基准值，我们可以利用 js 动态算出当前屏幕所适配的 font-size  
```js
document.getElementsByTagName('html')[0].style.fontSize = window.innerWidth / 10 + 'px';
```

rem 适配进阶  
一般我们获取到的视觉稿大部分是 iphone6 的，所以我们看到的尺寸一般是双倍大小的，在使用 rem 之前，我们一般会自觉的将标注 / 2。  
1.设计给的稿子双倍的原因是 iphone6 这种屏幕属于高清屏，也即是设备像素比 (device pixel ratio) dpr 比较大，所以显示的像素较为清晰。  
2.一般手机的 dpr 是 1，iphone4，iphone5 这种高清屏是 2，iphone6s plus 这种高清屏是 3，可以通过 js 的 window.devicePixelRatio 获取到当前设备的 dpr，所以 iphone6 给的视觉稿大小是（\*2）750×1334  
3.拿到了 dpr 之后，我们就可以在 viewport meta 头里，取消让浏览器自动缩放页面，而自己去设置 viewport 的 content 例如（这里之所以要设置 viewport 是因为我们要实现 border1px 的效果，加入我给 border 设置了 1px，在 scale 的影响下，高清屏中就会显示成 0.5px 的效果）  
```js
meta.setAttribute('content', 'initial-scale=' + 1/dpr + ', maximum-scale=' + 1/dpr + ', minimum-scale=' + 1/dpr + ', user-scalable=no');
```
4.设置完之后配合 rem，修改
```css
@function px2rem($px){
    $rem : 75px;
    @return ($px/$rem) + rem;
}
```
双倍 75，这样就可以完全按照视觉稿上的尺寸来了。不用在 / 2 了，这样做的好处是：  
1.解决了图片高清问题  
2.解决了 border 1px 问题（我们设置的 1px，在 iphone 上，由于 viewport 的 scale 是 0.5，所以就自然缩放成 0.5px）  

### rem 案例
[网易新闻](http://3g.163.com/touch/news/subchannel/all?version=v_standard)  
[聚划算](https://jhs.m.taobao.com/m/index.htm#!all)  
