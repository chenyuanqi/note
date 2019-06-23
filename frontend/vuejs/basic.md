
### 一、Vuejs 的安装
最简单的方法，直接引入一个 js 文件。
```html
<script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"></script>
```

在用 Vue 构建大型应用时推荐使用 NPM 安装 [1]。NPM 能很好地和诸如 webpack 或 Browserify 模块打包器配合使用。同时 Vue 也提供配套工具来开发 单文件组件。
```bash
npm install vue
```

`注意： Vue.js 的核心实现中使用了 ES5 的 Object.defineProperty 特性，IE8 及以下版本浏览器是不兼容的`

### 二、Vuejs 的核心
Vuejs 的核心：数据绑定 和 视图组件。  
Vue 的数据驱动视图：数据改变驱动了视图的自动更新，传统的做法你得手动改变 DOM 来改变视图，vuejs 只需要改变数据，就会自动改变视图。  
视图组件化：把整一个网页的拆分成一个个区块，每个区块我们可以看作成一个组件。网页由多个组件拼接或者嵌套组成。  

#### 2.1 数据驱动视图
数据驱动视图，解放 DOM 操作。  
```html
<div id="app">
{{ name }}
</div>
<script>
  let vm = new Vue({
      el:"#app",
      data:{
         name:"hello, vuejs.",
      }
  });
</script>
```
通过 new Vue( ) 创建一个实例 vm，参数是一个 json 对象，属性 el 提供一个在页面上存在的 DOM 元素 (id='app')，表明这个实例是关联指定的 DOM 节点。  
在 DOM 节点上，我们就可以使用双大括号 {{  }} 的语法来渲染 Vue 实例对象 data 中已经存在的属性值。  
一旦 name 的值被改变了，页面上立马跟着发生了变化，而不需要你再写 innerHTML 去更新视图了。这就是 Vuejs 的数据驱动视图。  

#### 2.2 双向绑定
Vue.js 提供了方便的语法指令，实现视图和数据的双向绑定，也就是说，不但数据变化了可以驱动视图，用户在页面上做了一些操作，也很方便地实现更新 model 层的数据。  
比如，监听用户在页面输入框输入的内容，然后将其实时更新在页面上。  

#### 2.3 组件化
组件化就是把页面中特定的区块独立抽出来，并封装成一个可方便复用的组件。  
组件的复用，Vuejs 提供了 props 属性来接受传递进来的参数。  

#### 2.4 Vue 的生命周期
创建一个 Vue 的实例，大致分几步走：  
new Vue() -> 设置数据 data -> 挂载元素 -> 数据渲染。  
```html
<div id="app">
  <!-- 数据渲染部分 -->
  我是{{ name }}，
  今年{{ age }}岁
</div>

<script>
let vm = new Vue({
	// 挂载元素
    el:'#app',

    // 实例 vm 的数据设定
    data:{
         name: "vikey",
         age :  18
    }
  });
</script>
```

实际上，每个 Vue 实例在被创建时都要经过一系列的初始化过程 —— 例如，需要设置数据监听、编译模板、将实例挂载到 DOM 并在数据变化时更新 DOM 等。同时在这个过程中也会运行一些叫做生命周期钩子的函数，这给了用户在不同阶段添加自己的代码的机会。  
参考 [Vue 生命周期图示](https://cn.vuejs.org/v2/guide/instance.html#%E7%94%9F%E5%91%BD%E5%91%A8%E6%9C%9F%E5%9B%BE%E7%A4%BA)

#### 2.5 Vue 实例的常用选项

