
### 什么是 Jquery
Jquery 是 js 的一个库，封装了我们开发过程中常用的一些功能，方便我们调用，提高开发效率。

### Jquery 版本
Jquery2.0.3 并不完全兼容 IE 6、7、8，解决方案如下
```html
// 不是 IE 就用 2.03，IE8 以上也用 2.0.3，IE8 及以下用 1.11 版本 
<!--[if !IE]> -->
     <script src="/Scripts/Jquery-2.0.3.min.js"></script>
 <!-- <![endif]-->

 <!--[if lte IE 8]>
     <script src="/Scripts/Jquery-1.11.1.min.js"></script>
 <![endif]-->

 <!--[if gt IE 8]>
      <script src="/Scripts/Jquery-2.0.3.min.js"></script>
 <![endif]-->
```

### Jquery 入口函数
```js
$(document).ready(function () {
  alert("入口函数");
})

// 简写形式
$(function () {
  alert("入口函数");
});

// 文档准备就绪才执行某些操作的时候
window.onload() = function(){ /* 要执行的代码 */ }   
```

### Jquery 多库共存
在页面中使用了包括 Jquery 在内的多个 JavaScript 库，就有可能会发生冲突。  
在 Jquery 中提供了 noConflict() 工具函数来消除这些冲突。会导致冲突的库被加载之后的任何时候，我们都可以调用下面的方法来消除冲突。
```js
// 将 $ 还原到非 Jquery 库所定义的含义
Jquery.noConflict();
```

### Jquery 常用函数
```js
// 获取 url 参数
$.getUrlParam = function (name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
    var r = window.location.search.substr(1).match(reg);
    
    return r !== null ? unescape(r[2]) : null;
}
```

### Jquery 常用选择器
```
基本选择器  
$('*')	匹配页面所有元素  
$('#id')	id 选择器  
$('.class')	类选择器  
$('element')	标签选择器  

组合 / 层次选择器  
$('E,F')	多元素选择器，用”, 分隔，同时匹配元素 E 或元素 F  
$('E F')	后代选择器，用空格分隔，匹配 E 元素所有的后代（不只是子元素、子元素向下递归）元素 F  
$(E>F)	子元素选择器，用”>” 分隔，匹配 E 元素的所有直接子元素  
$('E+F')	直接相邻选择器，匹配 E 元素之后的相邻的同级元素 F  
$('E~F')	普通相邻选择器（弟弟选择器），匹配 E 元素之后的同级元素 F（无论直接相邻与否）  
$('.class1.class2')	匹配类名中既包含 class1 又包含 class2 的元素  

基本过滤选择器  
$("E:first")	所有 E 中的第一个  
$("E:last")	所有 E 中的最后一个  
$("E:not(selector)")	按照 selector 过滤 E  
$("E:even")	所有 E 中 index 是偶数  
$("E:odd")	所有 E 中 index 是奇数  
$("E:eq(n)")	所有 E 中 index 为 n 的元素  
$("E:gt(n)")	所有 E 中 index 大于 n 的元素  
$("E:ll(n)")	所有 E 中 index 小于 n 的元素  
$(":header")	选择 h1~h7 元素  
$("div:animated")	选择正在执行动画效果的元素  

内容过滤器  
$('E:contains(value)')	内容中包含 value 值的元素  
$('E:empty')	内容为空的元素  
$('E:has(F)')	子元素中有 F 的元素，$('div:has (a)'): 包含 a 标签的 div  
$('E: parent')	父元素是 E 的元素，$('td: parent'): 父元素是 td 的元素  

可视化选择器  
$('E:hidden')	所有被隐藏的 E  
$('E:visible')	所有可见的 E  

属性过滤选择器  
$('E[attr]')	含有属性 attr 的 E  
$('E[attr=value]')	属性 attr=value 的 E  
$('E[attr !=value]')	属性 attr！=value 的 E  
$('E[attr ^=value]')	属性 attr 以 value 开头的 E  
$('E[attr $=value]')	属性 attr 以 value 结尾的 E  
$('E[attr =value]')	属性 attr 包含 value 的 E  
$('E[attr][attr =value]')	可以连用  

子元素过滤器  
$('E:nth-child(n)')	E 的第 n 个子节点  
$('E:nth-child(3n+1)')	E 的 index 符合 3n+1 表达式的子节点  
$('E:nth-child(even)')	E 的 index 为偶数的子节点  
$('E:nth-child(odd)')	E 的 index 为奇数的子节点  
$('E:first-clild')	所有 E 的第一个子节点  
$('E:last-clild')	所有 E 的最后一个子节点  
$('E:only-clild')	只有唯一子节点的 E 的子节点  

表单元素选择器  
$('E:type')	特定类型的 input  
$(':checked')	被选中的 checkbox 或 radio  
$('option: selected')	被选中的 option  

除了选择器，筛选节点的方法还可以：  
.find (selector) 查找集合每个元素的子节点  
.filter (selector) 过滤当前集合内元素  
```

### Jquery 实用方法
```js
// 获取一个表格中的 tr 元素的个数
$("#table tr").size(); 
// 获取一个表格下标为 2 的 tr 元素
$($("#table tr").get(2));
// 获取某个元素在包装集中的位置
$(".element").index(); 

// each() 方法可以循环遍历包装集中的每一个元素
$(".contentDiv").each( function(index, element) {

});  

// first() 函数会返回包装集中的第一个元素
$(".contentDiv").first().css("background-color", "#fff"); 

// last() 函数会返回包装集中的最后一个元素
$(".contentDiv").last().css("background-color", "#fff");

// next() 方法会取得一个包含匹配的元素集合中每一个元素紧邻的后面同辈元素
$(".contentDiv").first().next().css("background-color", "#fff");        

// prev() 方法取得一个包含匹配的元素集合中每一个元素紧邻的前一个同辈元素
$(".contentDiv").first().prev().css("background-color", "#fff");  

// parent() 方法取得一个包含着所有匹配元素的唯一父元素
$("#sonDiv").parent().css("background-color","#fff"); 

// children() 方法取得一个包含匹配的元素集合中每一个元素的所有子元素的元素集合
$("#parentDiv").children().each(function() {

});

// siblings() 方法获取的是相同父元素中的所有同辈元素
$("#sonDiv").siblings().each(function() {

});

// 修改元素属性
$("#img").attr("alt", "这是一张图片");
// 获取元素属性
$("#img").attr("alt");

// 添加 class 和移除 class
$("#img").addClass("active").removeClass("normal");
// 切换使用或丢弃 class
$("#img").toggleClass("active");

// 设置和获取元素的样式
$(".element").css("color","#fff");  
$(".element").css("color");
```

### Jquery 事件
DOM2 级事件处理程序定义了 2 个方法，用于处理指定和删除事件处理程序的操作，这 2 个方法是：addEventListener 和 removeEventListener。
```js
// DOM2 级事件模型支持事件捕获,在 DOM2 级事件处理程序中，提供了第 3 个参数来控制使用事件冒泡还是事件捕获
btn.addEventListener("click",function(){

}, true);  
```

因为各个浏览器在处理事件的方法上存在如此大的差异，导致开发人员在编写 js 事件处理程序时非常吃力。Jquery 事件模型屏蔽了浏览器之间的差异，它提供了如下的功能：
> 提供简历事件处理程序的统一方法  
> 允许在每一个元素上为每一个事件类型建立多个处理程序  
> 采用标准的事件类型名称  
> 使 Event 实例可用作处理程序的参数  
> 对 Event 实例的最常用的属性进行规范化  
> 为取消事件和阻塞默认操作提供了统一的方法  


常用事件  
```js
// hover([over,]out) 当鼠标移动到一个匹配的元素上面时，会触发指定的第一个函数；当鼠标移出这个元素时，会触发指定的第二个函数  
$("td").hover(
  function () {
    $(this).addClass("hover");
  },
  function () {
    $(this).removeClass("hover");
  }
); 

// toggle([speed],[easing],[fn]) 用于绑定两个或多个事件处理器函数，以响应被选元素的轮流的 click 事件
$("li").toggle(
  function () {
    $(this).addClass("visibled");
  },
  function () {
    $(this).removeClass("visibled");
  }
);

// blur([[data],fn]) 在元素失去焦点的时候触发，既可以是鼠标行为，也可以是按 tab 键离开的
$(".btn").blur({name:"tom"},function(event){
  console.log(event.data.name);
});

// change([[data],fn]) 当元素的值发生改变时，会发生 change 事件
$(selector).change();   

// click([[data],fn]) 会调用执行绑定到 click 事件的所有函数
// dblclick([[data],fn]) 在很短的时间内发生两次 click，即是一次 double click 事件
$("li").click();      

// keydown([[data],fn]) 当键盘或按钮被按下时，发生 keydown 事件
// keypress([[data],fn]) 当键盘或按钮被按下时，会发生 keypress 事件
// keyup([[data],fn]) 当键盘或按钮被松开时，发生 keyup 事件
$("input").keyup(function(){
  $("input").css("background-color","#fff");
});   

// mousedown([[data],fn]) 当鼠标指针移动到元素上方，并按下鼠标按键时，会发生 mousedown 事件
// mouseenter([[data],fn]) 当鼠标指针穿过元素时，会发生 mouseenter 事件  
// mouseleave([[data],fn]) 当鼠标指针离开元素时，会发生 mouseleave 事件
// mousemove([[data],fn] 当鼠标指针在指定的元素中移动时，就会发生 mousemove 事件
// mouseout([[data],fn]) 当鼠标指针从元素上移开时，发生 mouseout 事件
// mouseover([[data],fn]) 当鼠标指针位于元素上方时，会发生 mouseover 事件
// mouseup([[data],fn]) 当在元素上放松鼠标按钮时，会发生 mouseup 事件

// resize([[data],fn]) 当调整浏览器窗口的大小时，发生 resize 事件
// select([[data],fn]) 当 textarea 或文本类型的 input 元素中的文本被选择时，会发生 select 事件
// submit([[data],fn]) 当提交表单时，会发生 submit 事件，该事件只适用于表单元素
```

事件处理  
bind(type,[data],fn) 方法为每个匹配元素的特定事件绑定事件处理函数，该方法可以同时绑定多个事件类型  
```js
$("#btn").bind({
  click:function(){$("p").slideToggle();},
  mouseover:function(){$("body").css("background-color","red");},  
  mouseout:function(){$("body").css("background-color","#FFF");}  
});  
```
unbind(type,[data|fn]]) 方法是 bind() 的反向操作，从每一个匹配的元素中删除绑定的事件。如果没有参数，则删除所有绑定的事件。  
one(type,[data],fn) 方法为每一个匹配元素的特定事件绑定一个一次性的事件处理函数；也就是说，通过 one() 执行的事件只会被执行一次。  
trigger(type,[data]) 方法在每一个匹配的元素上触发某类事件。  
triggerHandler(type, [data]) 方法会触发指定的事件类型上所有绑定的处理函数。但不会执行浏览器默认动作，也不会产生事件冒泡。  
live(type, [data], fn) 方法，由于 bind() 方法在新添加元素时无法为新元素绑定事件，所以 live() 方法来绑定新添加元素的事件。  
die(type, [fn]) 方法是从元素中删除先前用 live() 方法绑定的所有事件，如果 die() 方法不带参数，则所有绑定的 live 事件都会被移除。如果提供了 type 参数，那么会移除对应的 live 事件。如果也指定了第二个参数 function, 则只移出指定的事件处理函数。  
delegate(selector,[type],[data],fn) 另一个事件委派的方法，代替 live()。  
undelegate([selector,[type],fn]) 方法删除由 delegate() 方法添加的一个或多个事件处理程序，代替 die()。  
on(events,[selector],[data],fn) 统一代替 bind()，live() 和 delegate()。  
off(events,[selector],[fn]) 方法移除用 on() 绑定的事件处理程序。  

事件委派  
live() 方法的原理是事件委托，而事件委托是通过冒泡机制来实现的。  
在 Jquery 的文档筛选中有一个重要的方法：closest()。这是在 Jquery1.3 中新增的方法，它用于从元素本身开始，逐级向上级元素匹配，并返回最先匹配的元素。closest 会首先检查当前元素是否匹配，如果匹配则直接返回元素本身。如果不匹配则向上查找父元素，一层一层往上，直到找到匹配选择器的元素。如果什么都没找到则返回一个空的 Jquery 对象。   
由于 closest 可以从元素自己本身开始查找，所以 closest 非常适合用于做事件委派。
```js
/* 模拟事件委派 */
$(function(){
  $(document).bind("click",function(event){
    var obj = $(event.target).closest(".row");
    if((obj[0]) == event.target){
    	alert("find!");
    }
  })
  $("#container").append("<div class='row'>div2</div>");
}); 
```
事件委派是指绑定在祖先元素上的事件处理函数可以对在后代上触发的事件作出回应。因为所有的元素事件都会通过冒泡到根元素中，所以我们只需要在文档根元素上绑定事件即可，然后在处理时，判断一下该事件是否是我们需要处理的对象就可以了。  
那么，事件委派的每一个事件都会冒泡到 document 上面去，影响性能，为了提高性能，live() 方法提供了第二个参数来指定绑定的上下文，说明事件委派的根对象是什么。  
```js
// live() 绑定的上下文是#container 容器，也就是说事件只会冒泡到这个容器上，不会再冒泡到文档的根元素
$(".row","#container").live("click",function(event){
  alert($(this).html())
})   
```

### Jquery 动画
Jquery 提供了一组方法用于为 HTML 元素制作简单的动画效果。  

show([speed,[easing],[fn]]) 方法用于显示隐藏的元素，如果选择的元素是可见的，这个方法将不会改变任何东西。无论这个元素是通过 hide() 方法隐藏的还是在 CSS 里设置了 display:none;，这个方法都将有效。    
hide([speed,[easing],[fn]]) 方法用于隐藏元素，如果选择的元素是隐藏的，这个方法将不会改变任何东西。  
toggle([speed],[easing],[fn]) 方法用于切换一个元素的可见性。如果元素是可见的，切换为隐藏的；如果元素是隐藏的，切换为可见的。  

slideDown([speed],[easing],[fn]) 方法通过高度变化（向下增大）来动态地显示所有匹配的元素，在显示完成后可选地触发一个回调函数。这个动画效果只调整元素的高度，可以使匹配的元素以 “滑动” 的方式显示出来。  
slideUp([speed],[easing],[fn]) 方法通过高度变化（向上减小）来动态地隐藏所有匹配的元素，在隐藏完成后可选地触发一个回调函数。这个动画效果只调整元素的高度，可以使匹配的元素以 “滑动” 的方式显示出来。  
slideToggle([speed],[easing],[fn]) 方法通过高度变化来切换所有匹配元素的可见性，并在切换完成后可选地触发一个回调函数。这个动画效果只调整元素的高度，可以使匹配的元素以 “滑动” 的方式显示出来。  

fadeIn([speed],[easing],[fn]) 方法通过不透明度的变化来实现所有匹配元素的淡入效果，并在动画完成后可选地触发一个回调函数。这个动画只调整元素的不透明度，也就是说所有匹配的元素的高度和宽度不会发生变化。  
fadeOut([speed],[easing],[fn]) 方法通过不透明度的变化来实现所有匹配元素的淡出效果，并在动画完成后可选地触发一个回调函数。这个动画只调整元素的不透明度，也就是说所有匹配的元素的高度和宽度不会发生变化。  
fadeToggle([speed],[easing],[fn]) 方法通过不透明度的变化来开关所有匹配元素的淡入和淡出效果，并在动画完成后可选地触发一个回调函数。这个动画只调整元素的不透明度，也就是说所有匹配的元素的高度和宽度不会发生变化。  
fadeTo([[speed],opacity,[easing],[fn]]) 方法把所有匹配元素的不透明度以渐进方式调整到指定的不透明度，并在动画完成后可选地触发一个回调函数。这个动画只调整元素的不透明度，也就是说所有匹配的元素的高度和宽度不会发生变化。  

animate(params,[speed],[easing],[fn]) 方法是用于创建自定义动画的函数，这个函数的关键在于指定动画形式及结果样式属性对象。这个对象中每个属性都表示一个可以变化的样式属性（如 “height”、“top” 或 “opacity”）。  
`注意：所有指定的属性必须用骆驼形式，比如用 marginLeft 代替 margin-left。`  
```
# 可以被动画的属性
# borderWidth、width、height、fontSize、opacity、margin、padding、bottom、left、right、top、wordSpacing
```

### Jquery 插件化
一个 Jquery 插件可以实现某项功能，完成某种炫酷的效果。  
使用 Jquery 来编写 Jquery 插件有 2 种形式：全局 Jquery 插件和基于包装集的 Jquery 插件。
```js
// 通过 $.方法名称就可以定义一个全局的 jQuery 插件
$.say = function(name){
  alert("hello," + name);
} 
// 调用也非常简单
$.say("Vikey");

// 基于包装集的 Jquery 插件，形式：$.fn.myplugins（$.fn 是 jQuery.prototype 的另外一种写法）
// prototype 是 JavaScript 的，Jquery 不建议直接在原型上扩展方法，所以提供了一个 $.fn 给开发者用于插件的编写
(function($){
    $.fn.mySlider = function(options){
    	// 设置插件的参数
        var settings = $.extend({
            autoplay: true, // 是否自动播放
            interval: 3000, // 自动播放的时间间隔
            speed: 400, //动画的速度
            direction:'ltr' //从左到右或从右到左
        },options || {});
 
        // 当前幻灯片的索引
		var current = 0;
		// 幻灯片jQuery对象
		var content = $(this);
		// ul元素
		var ulElement = content.find("ul");
		// li元素
		var item = ulElement.find("li");
		// 为实现无限循环效果进行图片分组
		var group = item.length;
		// next导航按钮
		var next = $(".next");
		// prev导航按钮
		var prev = $(".prev"); 
 
        var slider = {
        	// 幻灯片的初始化方法
            init:function(){
            	// 自动播放的时长
                timer = ''; 
                if(item.length > 1){
                    ulElement.css({
                        'width' : (item.length+2) * 100 + "%",
                        'margin-left': -100 + '%',
                        'left': 0
                    });
                    item.css({
                        'float': 'left',
                        'listStyle': 'none',
                         'width': 100 / (item.length + 2) + '%'
                    });
                    ulElement.html(item.slice(0).clone());
                    ulElement.append(item.slice(0, 1).clone());
                    ulElement.prepend(item.slice(-1).clone());
                }
            },
            // 自动播放
            autoPlay: function() {
                var _this = this;
                if (settings.autoplay) {
                    timer = setInterval(function() {
                        if (settings.direction == 'ltr') {
                            _this.toPrev();
                        } else {
                            _this.toNext();
                        }
                    }, settings.interval);
                }
            },
            // 停止自动播放
            stopAuto: function() {
                if (settings.autoplay) {
                    clearInterval(timer);
                }
            },
            // 幻灯片的动画函数
            animate: function(num) {
                var _this = this;
                ulElement.animate({
                    left: -num * 100 + '%'
                }, settings.speed,function(){
                    if (num < 0) {
                        current = group - 1;
                        ulElement.css('left', -(group - 1) * 100 + '%');
                    } else if (num >= group) {
                        current = 0;
                        ulElement.css('left', 0);
                    } else {
                        current = num;
                    }
                })
            },
            // 跳转到下一帧
            toNext: function() {
                current = current + 1;
                this.animate(current);
            },
            // 跳转到上一帧
            toPrev: function() {
                current = current - 1;
                this.animate(current);
            }
        }
 
        slider.init();
        slider.autoPlay();
         
        next.on("click",function(event){
            slider.toNext();
        })
 
        prev.on("click",function(event){
            slider.toPrev();
        });

        // 在匿名函数的内部，this 对象引用的是整个包装集，不需要在使用 $(this) 来重新包装
        // 基于包装集的函数一定要能够支持链式编程结构，所以一定要返回 this 对象
        return this;
    }
})(jQuery);    
// 调用
$('.slider-warp').mySlider();
```

在编写一个 Jquery 插件的时候，我们需要考虑以下一些问题：
1、一般 Jquery 插件都需要通过一个独立的 js 文件来存储，我们应该如何确定插件命名的规则？  
> 在为 Jquery 插件进行命名的时候，通常都遵循 jquery.命名空间.js 的命名规则，比如幻灯片插件可以命名为：jquery.mySlider.js  
2、如果将来 $ 这个符号被 jQuery.noConflict 之后，如何保证插件还能继续使用？  
> 可以将 $ 改为 Jquery，但是这种方式会增加我们的插件开发的工作量  
> 最好的办法是将整个插件都写在一个闭包中
```js
(function($){
   //插件在此定义
})(jQuery);
```
3、如何为插件设置参数？  
> 如果插件参数较少，可以直接定义；如果较多，设置可选参数
```js
// 参数较少时
$.myPlugins = function(param1, param2){
  // ......
}

// 参数较多时
$.myPlugins = function(param1, options){
  //在代码中通过 extend 方法来完成覆盖
  var settings = $.extend({
    a2:"value1",
    a3:"value2",
    a4:"value3",
    a5:"value4",
    a6:"value5",
    a7:"value6",
  } ,options || {});
}
```

### Jquery 与 Ajax
Ajax 是 Asynchronous JavaScript And XML 的简写，它是一种能够向服务器请求额外数据而不需要卸载页面的技术。通俗来说，Ajax 是一种可以使客户端和服务器进行异步交互，在不刷新整个页面的情况下，对页面进行局部更新的技术。    

jQuery 对原生 JavaScript 中的 Ajax 请求代码做了封装，屏蔽了各个浏览器之间的差异，使用起来非常简单。  

load(url, [data], [callback]) 方法可以加载远程的 HTML 文件，并将内容插入到指定的 DOM 元素中。
```js
// 取 01.html 中 id 是 title 的内容
$("#content").load("01.html #title");  
// 加载数据
$("#content").load("01.html",function(data){
  console.info(data);
});  
```

$.get(url, [data], [callback], [type]) 方法可以通过远程 HTTP GET 请求载入信息。这是一个简单的 GET 请求功能以取代复杂 $.ajax 方法。请求成功时可调用回调函数。如果需要在出错时执行函数，请使用 $.ajax 方法。  
> data 是待发送 Key/value 参数  
> type 是返回内容格式，xml, html, script, json, text, _default  
```js
// 也可以使用 $.getJSON()
$.get("person.json",function(data){
  console.info(data.name);
}, "json");
```
$.post(url, [data], [callback], [type]) 方法通过远程 HTTP POST 请求载入信息。这是一个简单的 POST 请求功能以取代复杂 $.ajax 方法。请求成功时可调用回调函数。如果需要在出错时执行函数，请使用 $.ajax 方法。
