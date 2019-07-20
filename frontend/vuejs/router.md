
### 前端路由实现
前端的路由实现一般有两种，一种是 Hash 路由，另外一种是 History 路由。  

History 路由，是指 History 接口允许操作浏览器的曾经在标签页或者框架里访问的会话历史记录。
```
# History.length 是一个只读属性，返回当前 session 中的 history 个数，包含当前页面在内
# 举个例子，对于新开一个 tab 加载的页面当前属性返回值 1 

# History.state 返回一个表示历史堆栈顶部的状态的值，这是一种可以不必等待 popstate 事件而查看状态的方式

# History.back() 前往上一页，用户可点击浏览器左上角的返回按钮模拟此方法，等价于 history.go (-1)  
# 注意：当浏览器会话历史记录处于第一页时调用此方法没有效果，而且也不会报错  

# History.forward() 在浏览器历史记录里前往下一页，用户可点击浏览器左上角的前进按钮模拟此方法，等价于 history.go (1) 
# 注意：当浏览器历史栈处于最顶端时 (当前页面处于最后一页时) 调用此方法没有效果也不报错  

# History.go(n) 通过当前页面的相对位置从浏览器历史记录 (会话记录) 加载页面  

# history.pushState () 和 history.replaceState ()，这两个 API 都接收三个参数，分别是状态对象、标题和地址。两个 API 都会操作浏览器的历史记录，而不会引起页面的刷新；pushState 会增加一条新的历史记录，而 replaceState 则会替换当前的历史记录
# 状态对象（state object）是一个 JavaScript 对象，与用 pushState () 方法创建的新历史记录条目关联。无论何时用户导航到新创建的状态，popstate 事件都会被触发，并且事件对象的 state 属性都包含历史记录条目的状态对象的拷贝  
# 标题（title），FireFox 浏览器目前会忽略该参数，虽然以后可能会用上。考虑到未来可能会对该方法进行修改，传一个空字符串会比较安全。或者，你也可以传入一个简短的标题，标明将要进入的状态  
# 地址（URL），新的历史记录条目的地址。浏览器不会在调用 pushState () 方法后加载该地址，但之后，可能会试图加载，例如用户重启浏览器。新的 URL 不一定是绝对路径；如果是相对路径，它将以当前 URL 为基准；传入的 URL 与当前 URL 应该是同源的，否则，pushState () 会抛出异常。该参数是可选的；不指定的话则为文档当前 URL
```

Hash 路由，我们经常在 url 中看到 #，这个 # 有两种情况，一个是我们所谓的锚点，比如典型的回到顶部按钮原理、Github 上各个标题之间的跳转等，但是路由里的 # 不叫锚点，我们称之为 hash。  
现在的前端主流框架的路由实现方式都会采用 Hash 路由，当 hash 值发生改变的时候，我们可以通过 hashchange 事件监听到，从而在回调函数里面触发某些方法。  
```html
<!-- 简单模拟 vue router 切换路由 -->
<div class="router_box">
    <a href="/home" class="router">主页</a>
    <a href="/news" class="router">新闻</a>
    <a href="/about" class="router">关于</a>
</div>
<div id="router-view"></div>
<script>
function Vue(parameters) {
	let vue = {};
	vue.routes = parameters.routes || [];
	vue.init = function() {
	    document.querySelectorAll(".router").forEach((item, index) => {
	        item.addEventListener("click", function(e) {
	            let event = e || window.event;
	            event.preventDefault();
	            window.location.hash = this.getAttribute("href");
	        }, false);
    });

    window.addEventListener("hashchange", () => {
        vue.routerChange();
    });

    vue.routerChange();
};

vue.routerChange = () => {
    let nowHash = window.location.hash;
    let index = vue.routes.findIndex((item, index) => {
        return nowHash == ('#' + item.path);
    });
    if (index >= 0) {
        document.querySelector("#router-view").innerHTML = vue.routes[index].component;
    } else {
        let defaultIndex = vue.routes.findIndex((item, index) => {
            return item.path == '*';
        });
        if (defaultIndex >= 0) {
            window.location.hash = vue.routes[defaultIndex].redirect;
        }
    }
};

vue.init();
}

new Vue({
	routes: [{
	    path: '/home',
	    component: "<h1>主页</h1><a href='#'>这里是主页内容</a>"
	}, {
	    path: '/news',
	    component: "<h1>新闻</h1><a href='#'>这里是新闻内容</a>"
	}, {
	    path: '/about',
	    component: '<h1>关于</h1><h4>关于...</p>'
	}, {
	    path: '*',
	    redirect: '/home'
	}]
});
</script>
```
