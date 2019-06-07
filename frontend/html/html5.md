### 一 HTML5 初识
#### 1.1 HTML5 语法变化
HTML5 不仅仅是一个标记语言，而是制定了 Web 应用的标准：更多的语义化标签，新的表单，多媒体，canvas，数据存储，地理应用等等。
```
标签不再区分大小写
单标签可统一写为 <img />，常见的有：area，base，br，col，hr，img，input，link，mata
省略结束标签元素：dt，dd，li，option，p，thead，tbody，tr，td，t
可省略全部标签元素：html，head，body，tbdy
允许省略属性名
允许省略属性值的引号
删除了少量的元素和属性，如 font，HTML5 推荐使用 CSS 来控制
```

H5同样扩展了很多标签的功能：
```html
<figure>
    <figcaption>这是图片</figcaption>		//这里会换行
    <img src="./1.jpg">
</figure>
```

#### 1.2 HTML5 网页结构
常见的网页结构有：html4，xhtl，html5
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
</body>
</html>

```

可以设置多个 meta 标签：
```html
// 过期时间
<meta http-equiv="Expires" content="Sat Sep 27 16:12:33 CST 2009">
// 禁止从缓存中读取页面
<meta http-equiv="Pragma " content="no-cache ">
// 自动刷新
<meta http-equiv="Refresh" content="2" URL="http://www.a.com"> 
// 网页过期删Cookie
<meta http-equiv="Set-cookie" content="name=value expires=Sat Sep 27 16:12:33 CST 2009,path=/">
```

#### 1.3 HTML5 新增元素
```
文档结构元素：
<article>   代表一片文章，内部可包含 <article><header><footer><section>
<section>   页面内容进行分块，内部可包含： <h1><article><section>
<nav>       导航条
<aside>     侧边栏
<header>    头部
<footer>    脚注
<hgroup>    组织多个 <h1> 这样的标题
<figure>    一个独立的图片区域，包含多个图片标签，<figcaption>表示图片区域标题

语义元素：
<mark>      显示页面中重点关注的内容，浏览器通常使用黄色显示
<time>      标注时间，向用户展示，
                datatime 属性：标注时间，向计算机展示
                <time datetime="2012-04-01"> 2012 年 4 月</time>
                <time>2012-02-03</time>

功能元素：
<meter>     表示一个已知最大值和最小值的技术仪表，如电池剩余电量。
<progress>  表示进度条
```

#### 1.4 HTML5 新增属性
```
contentEditable
    如果将该属性设为 true，那么浏览器会允许开发者直接编辑该 HTML 元素里的内容，且如果一个 HTML 元素的父元素是可编辑的，那么他自己默认也是可编辑的。

isContentEditable
    用返回值 true、false 来判断

designMode
    相当于一个全局的 contentEditable 属性，如果把整个页面的 designMode 设置为 on，那么所有支持编辑的元素都会变为可编辑状态，除非设置为 off，如下所示：
	<body ondbclick="document.designMode='on'">

hidden
    hidden属性可以代替 CSS 样式中的 display 属性

spellcheck
    HTML5 为 <input><textarea>等元素增加了 spellcheck 属性，支持 true、false 属性值，以判断是否需要浏览器对文本进行校验，如：对拼错的单词进行提示。

disabled
    新属性 disabled 直接就可以让 input 无法选择，而老版的 html 中要使用:
    disabled="disabled"
```

#### 1.5 语义化标签
本质上新语义标签与 <div>、<span> 没有区别，使用时除了在 HTML 结构上需要注意外，其它和普通标签的使用无任何差别，可以理解成 <div class="nav"> 相当于 <nav>。尽量避免全局使用 header、footer、aside 等语义标签。
常见语义化标签有：
```
header 			定义 section 或 page 的页眉。页面的头部
nav 			定义导航链接。一般定义导航
main			定义主要区域 
section 		定义文档中的节 
aside			定义内容之外的内容 侧边栏
footer 			定义 section 或 page 的页脚。
article 		定义文章。
mark 			定义记号文本，自带可改变的背景色，带UI，不常用
figure 			定义媒介内容的分组，以及它们的标题。
figcaption		定义 figure 元素的标题
details			定义元素的细节
summary		    定义可见的 <details> 元素的标题。
progress		定义任何类型的任务的进度，带 UI，不常用
```
在不支持 HTML5 新标签的浏览器里，会将这些新的标签解析成行内元素(inline)对待，所以我们只需要将其转换成块元素(block)即可使用。
但是在 IE9 版本以下，并不能正常解析这些新标签，但是可以通过
 document.createElement('footer') 创建的自定义标签（创建出来的元素是行内元素，因此一般情况下要给：display:block）。
在实际开发中，一般使用第三方兼容文件html5shiv来解决上述问题：
```html
<script src="html5shiv.js"></script>
<!--[if lte IE 9]>
<script type="text/javascript" src="htmlshiv.js"></script>
<![endif]--
```

### 二 表单
#### 2.1 表单自动联想
```html
    <input type="text" list="data">
    <datalist id="data">
        <option>男</option>
        <option>女</option>
    </datalist>
```

#### 2.2 表单 type
```
email 		输入email格式
tel 		手机号码，移动设备上得到焦点后会弹出键盘
url 		只能输入 url 格式
number 	    只能输入数字，此时有属性 min  max  step 等
search 		搜索框，输入内容时候会出现清除按钮
range 		范围 滑动条
color 		拾色器
time		时间，可控制时间，且有增减按钮
date 		日期，可控制日期，且有日期控件！
month 	    月份，同上
week 		星期，同上
datetime    时间日期
```

部分类型是针对移动设备生效的，且具有一定的兼容性，在实际应用当中可选择性的使用，比如：
```html
<form>
    <label>
        邮箱：<input type="email" name="email" class="email">
    </label>
    <input type="submit" value="提交">
</form>
```
上述表单类型设置为 email 后，可以验证邮箱书写的合法性。

#### 2.3 表单属性
```
placeholder 	占位符
autofocus 		获取焦点
multiple 		文件上传多选或多个邮箱地址  
autocomplete    自动完成，用于表单元素，也可用于表单自身(on/off)
form 			指定表单项属于哪个 form，处理复杂表单时会需要
novalidate 		关闭验证，可用于 <form> 标签
required 		必填项
pattern 		正则表达式 验证表单

autocapitalize	iOS独有属性，设置为 off 关闭首字母大写
autocorrect     iOS独有属性，设置为 off 关闭输入自动修正
```

#### 2.4 表单事件
```
oninput         用户输入内容时触发，可用于移动端输入字数统计
oninvalid       验证不通过时触发
```

### 三 多媒体
在 HTML5 之前，在网页上播放音频/视频的通用方法是利用 Flash 来播放，但是大多情况下，并非所有用户的浏览器都安装了 Flash 插件，由此使得处理音频/视频播放变的非常复杂，并且移动设备的浏览器并不支持 Flash 插件。

#### 3.1 音频
```html
<audio src="./test.mp3"></audio>
//由于版权问题，不同浏览器可以播放的音频格式不同，适配方案如下：
<audio controls>
    <source src="./test.mp3">
    <source src="./test.wav">
    <source src="./test.ogg">
    您的浏览器不支持播放该格式多媒体
</audio>
```

音频支持的属性：
```
muted			静音
controls		显示控制条
autoplay 		自动播放
controls 		是否显不默认播放控件
loop 			循环播放
preload 		预加载 同时设置 autoplay 时些属性失效
auto：          加载；
none：          不加载；
metadata：      加载元数据（封面、歌词、作者）
```

#### 3.2 视频
```html
<video src="./test.mp4"></video>
//由于版权问题，不同浏览器可以播放的视频格式不同，适配方案如下：
<video controls>
    <source src="./test.mp4">
    <source src="./test.ogg">
    <source src="./test.ogg">
    您的浏览器不支持播放该格式多媒体
</video>
```

视频支持的属性：
```
autoplay        自动播放
controls        是否显示默认播放控件
loop            循环播放
preload         预加载，同时设置了 autoplay，此属性将失效
width           设置播放窗口宽度
height          设置播放窗口的高度

```

#### 3.3 多媒体控制
```
方法：
    load()			加载
    play()			播放
    pause()			暂停
属性：
    currentTime 	视频播放的进度
    paused			视频播放的状态
    duration		视频总时长
事件：
    oncanplay		是否能播放
    ontimeupdate	播放时触发，报告当前的播放进度
    onended		    播放完时触发

```

#### 3.4 全屏
全屏:           video.webkitRequestFullScreen();
webkit: 	    webkit RequestFullscreen()
moz: 	        moz RequestFullScreen()
ms: 		    ms RequestFullscreen()
其他:	        RequestFullscreen()
出现浏览器前缀的原因：为了兼容低版本的自己的浏览器，比如说 webkit 为了兼容低版本的 chrome。

### 四 新增 DOM 操作
#### 4.1 选择器
```
//获取复合匹配条件的第1个元素，参数可以是标签、类、id
document.querySelector("div"); 
//获取复合匹配条件的所有元素，以伪数组形式     
document.querySelectorAll("div");  
```

#### 4.2 类名操作
```
Node.classList.add('class') 		添加class
Node.classList.remove('class') 	    移除class
Node.classList.toggle('class') 		切换class，有则移除，无则添加
Node.classList.contains('class') 	检测是否存在class
```
#### 4.3 自定义属性data
```html
<ul id="items">
    <li data-flag="1" class="first"></li>
    <li data-rest="Monday" class="second"></li>
    <li data-week-li="test" class="third"></li>
</ul>
<script>
    //注意 json属性必须带""
    let json = {
        "name": "zs",
        "age": 30
    }

    let jsonStr = JSON.stri
    
</script>
```

#### 4.4 JSON 对象新增方法
json 与字符串的转换我们之前一直使用的是 eval 方法，在 H5 中增加了以下两个方法：
```js
//注意 json属性必须带""
let json = {
    "name": "zs",
    "age": 30
}

let jsonStr = JSON.stringify(json);
let jsonObj = JSON.parse(jsonStr);

console.log(jsonStr);   //{"name":"zs","age":30}
console.log(jsonObj);   //{name: "zs", age: 30}
```

#### 4.5 延迟加载 defer 与 async
defer：延迟加载，会按顺序执行，最好只在加载外部 JS 时使用，其他场合会有兼容问题。
async：异步加载，加载完就触发，有顺序问题。
```html
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <script defer="defer"></script>
</head>
```
在 head 内的标签由于使用了 defer，那么就不会先加载，而是 onload 之后再加载。

#### 4.6 拖拽
为元素增加属性 draggable="true" 可以设置此元素是否可以进行拖拽。其中图片、链接默认是开启的，但是这时候的拖动没有带走数据，可以设置 ondragstart 事件就可以携带数据。
```
拖拽元素事件（页面中设置了draggable="true"属性的元素）
ondragstart	    应用于拖拽元素，当拖拽开始时调用，只执行一次
ondrag 		    应用于拖拽元素，整个拖拽过程都会调用，一直执行
ondragleave	    应用于拖拽元素，当鼠标离开拖拽元素时触发一次
ondragend		应用于拖拽元素，当拖拽结束时调用

目标元素事件（页面中任何一个元素都可以成为目标元素）		
ondragenter	    应用于目标元素，当拖拽元素进入时触发一次
ondragover		应用于目标元素，当停留在目标元素上时调用
ondrop		    应用于目标元素，当在目标元素上松开鼠标时调用
ondragleave	    应用于目标元素，当鼠标离开目标元素时调用

火狐下拖拽对象（dataTransfer）兼容：
ev.dataTransfer.setData(key,value)
ev.dataTransfer.getData(key) 
```

### 五 Web 存储
web 存储即把数据存储在浏览器中。传统方式我们以 document.cookie 来进行存储的，但是由于其存储大小只有 4k 左右，并且解析也相当的复杂，HTML5 规范则提出解决方案：Storage 存储（WebSQL、IndexDB已经被 w3c 放弃）。分别有两种存储方式：
```
会话存储：
    设置：window.sessionStorage.setItem("name","lisi");
    获取：window.sessionStorage.getItem("name");
    删除：window.sessionStorage.removeItem("name");
    清除：window.sessionStorage.clear();
    说明：生命周期为关闭浏览器窗口,在同一个窗口下数据可以共享，可存储大约5M
    
本地存储：
    获取：window.localStorage
    获取：window.localStorage.getItem("name");
    删除：window.localStorage.removeItem("name");
    清除：window.localStorage.clear();
    说明：永久生效，除非手动删除,可以多窗口共享,可存储大约20M
```

### 六 地理位置
HTML5 提供了 Geolocation 地理位置支持。目前大多数浏览器都可以支持（IE9+）。  
出于安全考虑，部分最新的浏览器只允许通过 HTTPS 协议使用 GeolocationAPI，在 http 协议下回跑出异常，当然本地开发不会有问题。  
访问地理位置方式：
```js
navigator.geolocation           //会提示用户是否允许获取该API的权限
```
示例代码：
```js
//判断浏览器是否兼容代码
if (navigator.geolocation) {
    navigator.geolocation.getCurrentPosition(
        function success(position) {
            console.log("获取位置成功，当前位置为：", position);
        },
        function errror(error) {
            console.log("获取位置错误，错误为：", error);
        }, 
        {
            timeout: 10000,
        }
    )
} else {
    //没有获取到地理位置执行的操作
}
```

获取位置成功后，返回的position参数对象，包含下列信息：
```
timestamp： 获取位置时的时间戳
ciirds：    坐标信息对象，内部包含：latitude（纬度），longitude（经度），accuracy（坐标精度，单位为米）
```

获取位置失败后，返回的错误对象，常见为：
```
code:   错误标识，1为用户拒绝分享位置 2为获取用户位置失败 3为获取超时 0为其他错误
```

传入的位置参数对象：
```
timeout:允许获取位置的超时时间，单位为毫秒
enableHighAccuracy： 布尔值是否获取高精度信息 
maximumAge：用户位置信息缓存的最大时间，默认为0，单位是毫秒  
```

当用户位置发生变化时，可以通过watchPostion方法监听用户的位置信息，该方法和getCurrentPosition使用方式一致：
```javascript
let watchID = navigator.geolocation.watchPosition(success, error, option);
//取消监听
navigator.geolocation.clearWatch(watchID);
```

### 七 检测用户网络状态
我们可以通过 window.onLine 来检测，用户当前的网络状况，返回一个布尔值
```js
window.online       //用户网络连接时被调用
window.offline      //用户网络断开时被调用
window.addEventListener("online",function(){
    alert("已经建立了网络连接")
});
window.addEventListener("offline",function(){
    alert("已经失去了网络连接")
});
```

### 八 访问设备
#### 8.1 访问摄像头
```
新标准：部分浏览器不支持，支持的浏览器有：Fifox36+（需要moz前缀），Chrome47+（需要webkit前缀）
navigator.getUserMedia  
旧标准：目前已废除，支持的浏览器有：Edg12+，Firefox17+，Chrome21+，AndroidBrowser56+，Chrome for Android57+，UC11.4+
navigator.mediaDevices.getUserMedia  
```

案例代码：
```html
<!DOCTYPE html> 
<html>  
  <head>  
    <meta charset="UTF-8">  
    <title></title>  
    <style>
      #canvas,#video {
        float: left;  
        margin-right: 10px;  
        background: #fff;  
      }      
      .box {  
        overflow: hidden;  
        margin-bottom: 10px;  
      }
    </style>
  </head>  
  <body>  
    <div class="box">
      <video id="video" width="400" height="300"></video>
      <canvas id="canvas"></canvas>
    </div>
    <button id="live">直播</button>
    <button id="snap">截图</button>
    <script>  
      var video = document.getElementById('video');
      var canvas = document.getElementById('canvas');  
      var ctx = canvas.getContext('2d');  
      var width = video.width;  
      var height = video.height;  
      canvas.width = width;  
      canvas.height = height;   
      function liveVideo(){  
        var URL = window.URL || window.webkitURL;   // 获取到window.URL对象
        navigator.getUserMedia({  
          video: true  
        }, function(stream){  
          video.src = URL.createObjectURL(stream);   // 将获取到的视频流对象转换为地址
          video.play();   // 播放
          //点击截图     
          document.getElementById("snap").addEventListener('click', function() {  
            ctx.drawImage(video, 0, 0, width, height);  
            var url = canvas.toDataURL('image/png');  
            document.getElementById('download').href = url;  
          });
        }, function(error){  
          console.log(error.name || error);  
        });  
      }  
      document.getElementById("live").addEventListener('click',function(){  
        liveVideo();  
      });    
    </script>  
  </body> 
</html>
```

#### 8.2 传感器与摇一摇实战
现代手机都内置了传感器，通过传感器，可以感知手机的方向和位置变化。  
手机方向变化和位置变化的解释：
```
X轴：左右横贯手机方向，当手机绕X轴旋转时移动方向称为Beta
Y轴：上下纵贯手机方向，时手机绕Y轴旋转时移动方向称为Gamma
X轴：垂直手机平面方向，当手机绕Z轴旋转时移动方向称为Alpha
```

方向事件 deviceorientation：设备方向发生变化时触发  
```javascript
window.addEventListener('deviceorientation', orientationHandler , true);
```

回调函数 orientationHandler   在注册后，会被定时调用，并会受到一个DeviceOrientationEvent类型的参数，通过该茶树获取设备的方向信息，如下所示：
```
absolute：如果方向数据跟地球坐标系和设备坐标系有差异，则为true，如果方向设备由设备本身的坐标系提供，则为false
alpha：设备在该方向上的旋转角度，范围是0-360
beta：设备在该方向上的旋转角度，范围是0-360
gamma：设备在该方向上的旋转角度，范围是0-360
```

移动事件 devicemotion：设备位置发生变化时触发：
```javascript
window.addEventListener('devicemotion', motionHandler , true);
```

同样的，回调函数motionHandler被注册后，也会被定时调用，并收到一个DevicemotionEvnent的参数，该参数可以访问设备的方向和位置信息，参数如下：
```
acceleration：设备在XYZ三个轴方向上移动的距离，已抵消重力加速
accelerationIncludingGravity：设备在XYZ三个轴的方向上移动的距离，包含重力加速
rotationRate：设备在Alpha，Beta，Gamma三个方向旋转的角度
interval：从设备获取数据的频率（单位毫秒）
```

摇一摇案例：摇一摇可以理解为手机不低于已定的速度在移动。  
实现的方法是监听 devicemotion 事件后，判断设备在XYZ三个方向上的移动距离与前一次移动的距离差，并除以两次事件触发的时间差，即为移动设备移动的速度
```html
<!DOCTYPE html>
<html lang='en'>

<head>
    <meta charset='UTF-8' />
    <meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0,user-scalable=no">
    <title>陀螺仪实现摇一摇</title>
</head>

<body>
    <div>用力摇一摇你的手机</div>
    <script type="text/javascript">
        var SHAKE_SPEED_THRESHOLD = 300; // 摇动速度阈值
        var lastTime = 0; // 上次变化的时间
        var x = y = z = lastX = lastY = lastZ = 0; // 位置变量初始化

        function motionHandler(evt) {
            var acceleration = evt.accelerationIncludingGravity; // 取得包含重力加速的位置信息
            var curTime = Date.now(); // 取得当前时间
            if ((curTime - lastTime) > 120) { // 判断
                var diffTime = curTime - lastTime; // 两次变化时间差
                lastTime = curTime; // 保存此次变化的时间
                x = acceleration.x;
                y = acceleration.y;
                z = acceleration.z;
                var speed = Math.abs(x + y + z - lastX - lastY - lastZ) / diffTime * 1000; // 计算速度
                if (speed > SHAKE_SPEED_THRESHOLD) { // 速度是否大于预设速度
                    alert("你摇动了手机");
                }
                lastX = x; // 保存此次变化的位置x
                lastY = y; // 保存此次变化的位置y
                lastZ = z; // 保存此次变化的位置z
            }
        }
        if (window.DeviceMotionEvent) {
            window.addEventListener('devicemotion', motionHandler, false);
        } else {
            alert('您的设备不支持位置感应');
        }
    </script>
</body>

</html>
```

## 九 离线应用
离线应用可以帮助用户在没有网络时使用 web 程序，H5 的离线功能包含：离线资源缓存、在线状态监测、本地数据存储等。  
离线 web 应用比普通的 web 应用多了一个描述文件，该文件用来列出需要缓存和永不缓存的资源，描述文件的扩展名为：.manifest 或者 .appcache(推荐使用)。  
首先需要在项目目录下创建 offline.appcache 文件：
```
CACHE MANIFEST      # 说明这是离线应用描述文件
CACHE:              # 会被缓存的资源列表
index.html
index.js
NETWORK:            # 总是从web获取的资源列表
test.js
```

html 文件需要添加如下配置：
```html
<html lmanifest="./offline.appcache">
```

### 十 跨文档通信
在过去，跨源或者跨窗口之间的通信往往是与服务端进行数据交互来实现的，并且需要借助轮训或者 Connect 技术来监听消息。  
H5 提供了 PostMessages 实现安全的跨源通信：
```js
otherWindow.postMessage(message, targetOrigin, [transfer]);

// otherWindow:其他窗口的引用，比如iframe中的contentWindow属性，执行widnow.open返回的窗口对象。
//三个参数：
//message:将要发给其他窗口的数据
// targetOrigin:通过创窗口的origin属性来指定哪些窗口能接收到消息事件，其值可以是字符“*”（表示无限制）或者一个url
// transfer:可选参数，是一串和message同时传递的Transferable对象，这些对象的所有权将被转译给消息的接收方，而发送乙方将不再保有所有权
```

iframe 应用实例：
```html
<button id="btn">点击发送消息给iframe</button>
<iframe src="http:127.0.0.1/iframe.html"></iframe>
<script>
    let btn = document.querySelector("#btn");
    let data = ["周一","周二","周五"];
    btn.onclick = function(){
        alert("执行发送数据给iframe？");
        window.parent.postMessage(data, "http:127.0.0.1/iframe.html");
    }
</script>
```

iframe 接受数据：
```html
<script>
    window.addEventListener("data", e =>{
        console.log("origin=", e.origin);
        console.log("data=", e.data);
    });
</script>
```

### 十一 Ajax2.0 新特性
设置 HTTP 请求超时时间
```js
xhr.timeout = 3000;
xhr.ontimeout = function(e){}
```

使用 FormData 表单对象管理表单数据：
```js
let formData = new FormData();
formData.append("username", "zs");
formData.append("id", 2331);
xhr.send(formData);
```

ajax2.0 上传文件
```js
for (let i = 0; i < files.length; i++) {            //files是一个选择文件的表单元素
    formData.append("files[]", files[i]);
}
```

全新的跨域请求，需要服务端配合。  
获取服务端的二进制数据：
```js
xhr.open("GET", "...");
xhr.responseType = 'blob';                      //表示服务器传回的数据是二进制对象
```

## 十二 Server Sent Event 服务器主动通信
像邮件这样的应用，服务端发送了邮件后，需要主动通知客户端。传统的做法是，客户端不断的去轮训服务端，这种做法十分不优雅！  
H5 提供了 Server Sent Event 技术解决上述问题：
- 简单轻量
- 单向数据传送（服务端向客户端传送）
- 基于 HTTP 协议
- 默认支持断线重连接
- 自定义发送数据类型

```js
if (typeof (EventSource) !== "undefined") {
    let source = new EventSource("/test");

    source.onmessage = function (e) {
        div.innerHTML = e.data;
    }
} else {
    alert("浏览器不支持服务器发送事件");
}
```