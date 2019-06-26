
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
1. beforeCreate  
此阶段为实例初始化之后，此时的数据观察和事件配置都没好准备好。此时的实例中的 data 和 el 还是 undefined，不可用的。  
2. created  
beforeCreate 之后紧接着的钩子就是创建完毕 created，我们同样打印一下数据 data 和挂载元素 el。此时，我们能读取到数据 data 的值，但是 DOM 还没生成，所以属性 el 还不存在，输出 $data 为一个 Object 对象，而 $el 的值为 undefined。  
3. beforeMount  
此阶段为即将挂载。$el 不再是 undefined，而是成功关联到我们指定的 dom 节点 <div id=”app”></div>，但此时 {{ name }} 还没有被成功地渲染成我们 data中 的数据。  
4. mounted  
mounted 也就是挂载完毕阶段，到了这个阶段，数据就会被成功渲染出来。此时打印属性 el，我们看到 {{ name }} 已经成功渲染成我们 data.name 的值。  
5. beforeUpdate  
当修改 vue 实例的 data 时，vue 就会自动帮我们更新渲染视图，在这个过程中，vue 提供了 beforeUpdate 的钩子给我们，在检测到我们要修改数据的时候，更新渲染视图之前就会触发钩子 beforeUpdate。  
在控制台把 app.name 的值修改，在更新视图之前 beforeUpdate 打印视图上的值，结果依然为原来的值，表明在此阶段，视图并未重新渲染更新。  
6. updated  
此阶段为更新渲染视图之后，此时再读取视图上的内容，已经是最新的内容。在控制台把 app.name 的值修改，在此阶段，视图将重新渲染更新。   
7. beforeDestroy  
调用实例的 destroy() 方法可以销毁当前的组件，在销毁前，会触发 beforeDestroy 钩子。   
8. destroyed  
成功销毁之后，会触发 destroyed 钩子，此时该实例与其他实例的关联已经被清除，它与视图之间也被解绑。  
销毁之前，修改 name 的值，可以成功修改视图显示的内容，一旦效用实例的 $destroy() 方法销毁之后，实例与视图的关系解绑，再修改 name 的值，已于事无补，视图再也不会更新了，说明实例成功被销毁了。  
9. actived  
keep-alive 组件被激活的时候调用。  
10. deactivated  
keep-alive 组件停用时调用。  
```html
<!-- ref 属性，用于获取 DOM 元素（beforeUpdate 用到） -->
<div ref="app" id="app">{{name}}</div>
<script>
  let app = new Vue({
    el:"#app",
      data:{
       name:"vikey"
    },
    beforeCreate(){
       console.log('即将创建');
       console.log(this.$data);
       console.log(this.$el);
    },
    created(){
      console.log('创建完毕');
      console.log(this.$data);
      console.log(this.$el);
   },
   beforeMount(){
      console.log('即将挂载');
      console.log(this.$el);
   },
   mounted(){
      console.log('挂载完毕');
      console.log(this.$el);
   },
    beforeUpdate(){
      console.log('=即将更新渲染=');
      let name = this.$refs.app.innerHTML;
      console.log('name:'+name);
   },
    updated(){
      console.log('==更新成功==');
      let name = this.$refs.app.innerHTML;
      console.log('name:'+name);
     },
      beforeDestroy(){
       console.log('销毁之前');
     },
     destroyed(){
       console.log('销毁成功');
     }
  });
</script>
```


#### 2.5 Vue 实例的常用选项
定义一个 Vue 实例，4 个常用参数选项：filters 过滤器、computed  计算属性、methods  方法、watch 观察。  

filters 过滤器  
顾名思义，就是定义过滤数据处理的方法。  
```html
<div id="app">
  数字1：{{ num1 | toInt }}<br>
  数字2：{{ num2 | toInt }}<br>
  数字3：{{ num3 | toInt }}
</div>

<script>
let vm = new Vue({
    //挂载元素
  el:'#app',
    //实例vm的数据
  data:{
         num1:33.141,
         num2:46.212,
         num3:78.541
    },
    //过滤器
  filters:{
      toInt(value){
         return parseInt(value);
      }
    }
  });
</script>
```
如上面代码所示，通过管道符 | 把函数 toInt 放在变量后面即可，num1，num2，num3 会分别作为参数 value 传入 toInt( value )方法进行运算，并返回一个整数。  

computed  计算属性  
有时候，我们拿到一些数据，需要经过处理计算后得到的结果，才是我们要展示的内容。计算属性 computed 的定义语法和过滤器 filters 类似，但是用法不一。  
```html
<div id="app">
 总和：{{ sum }}
</div>
<script>
let vm = new Vue({
//挂载元素
el:'#app',
  //实例vm的数据
  data:{
         num1:1,
         num2:3,
         num3:6
  },
  //计算属性
  computed:{
      sum(){
        return this.num1 + this.num2 + this.num3
      }
  }
});
</script>
```
需要注意的是，sum 的值是依赖 data 中的数据 num1，num2，num3 的，一旦它们发生了变化，sum 的值也会自动更新。  

methods  方法  
在 methods 中，我们可以定义一些方法，供组件使用。  
```html
<div id="app">
  {{ a }}
  <button v-on:click="plus">加1</button>
</div>
<script>
let vm = new Vue({
  //挂载元素
  el:'#app',
  //实例vm的数据
  data:{
         a:0
    },
  //方法
  methods:{
        plus(){
            return this.a++;
        }
    }
 });
</script>
```

watch 观察  
watch 选项是 Vue 提供的用于检测指定的数据发生改变的 api。  
```html
<div id="app">
  {{ a }}
  <button v-on:click="plus">加1</button>
</div>
<script>
let vm = new Vue({
  //挂载元素
  el:'#app',
  //实例vm的数据
  data:{
         a:0
    },
  //方法
  methods:{
        plus(){
            return this.a++;
        }
    },
  //观察
  watch:{
        // a() 表示我们要观察监听的就是数据 a
        a(){
          console.log(`有变化了，最新值：`);
          console.log(this.a);
        }
    }
 });
</script>
```

#### 2.6 在 html 中绑定数据
