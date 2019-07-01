
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
  <!-- v-on 可缩写为 @ -->
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
vue 提供的 v-html 指令，可以输出 html 标签的效果。  
```html
<div id="app" v-html="name"></div>
<script>
let app = new Vue({
  el:"#app",
  data:{
   name:"<strong>在 html 中绑定数据</strong>"
  }
});
</script>
```

html 标签的属性也很重要，用 v-bind 指令可以绑定属性。  
对于 v-bind 指令，如果一个标签中需要绑定多个属性的话，就得连续写多个 v-bind。
```html
<div id="app">
 <!--在href前用v-bind：修饰-->
 <a v-bind:href="link">hello官网</a>
</div>
<script>
let app = new Vue({
  el:"#app",
  data:{
   link:"http://hello.com"
  }
});
</script>
```
渲染的属性值是布尔值的时候（true 和 false），效果就不一样了，并不是简单地将 true 或者 false 渲染出来，而是当值为 false 的时候，属性会被移除。  
```html
<div id="app">
 <!--用缩写的语法-->
 <button :disabled="dis_true">
    我是true
 </button>
 <button :disabled="dis_false">
    我是false
 </button>
</div>
<script>
let app = new Vue({
  el:"#app",
  data:{
    dis_true:true,
    dis_false:false
  }
});
</script>
```
我们看到，当属性值设置成 true 的时候，disabled 的值为解析成 disabled ，当属性值设置成 false 的时候，属性 disabled 直接被移除掉了。

vue 支持 javascript 表达式的写法  
```html
<div id="app">
    {{ num+3 }}<br>
    {{ ok ? 'yes':'no' }}<br>
    <a :href="'http://'+host">hello</a>
</div>
<script>
let app = new Vue({
  el:"#app",
  data:{
      num:2,
      ok:true,
      host:'hello.com'
  }
});
</script>
```
虽然 vue 支持 javascript 表达式运算，我们只会在运算比较简单的情况下才会这么玩，当运算比较繁琐复杂的时候，一定要用 vue 的 computed 属性来进行计算。  

#### 2.7 vue 必备指令
v-text 用于更新标签包含的文本，它的作用跟双大括号的效果一样  
```html
<div id="app">
  <p v-text="msg"></p>
  <!--效果相同-->
  <p>{{ msg }}</p>
</div>
<script>
let app = new Vue({
  el:"#app",
  data:{
      msg:'hello,vue'
  }
});
</script>
```

v-html 用于更新标签包含的 html，用于渲染 html 标签  
```html
<div id="app">
  <p v-html="msg"></p>
</div>
<script>
let app = new Vue({
  el:"#app",
  data:{
    msg:'<b>hello,vue</b>'
  }
});
</script>
```

v-show 用来控制元素的 display css 属性  
v-show 指令的取值为 true/false，分别对应着显示/隐藏。  
```html
<div id="app">
  <p v-show="show1">我是true</p>
  <!-- false 的时候，不会显示 -->
  <p v-show="show2">我是false</p>
</div>
<script>
let app = new Vue({
  el:"#app",
  data:{
      show1:true,
      show2:false
  }
});
</script>
```

v-if 控制元素是否需要被渲染出来  
v-if 指令的取值为 true/false，分别对应着渲染/不渲染元素。  
```html
<div id="app">
  <p v-if="if_1">我是true</p>
  <!-- 下面的 p 标签不会渲染 -->
  <p v-if="if_2">我是false</p>
</div>
<script>
let app = new Vue({
  el:"#app",
  data:{
      if_1:true,
      if_2:false
  }
});
</script>
```
如果需要频繁切换显示/隐藏的，就用 v-show；如果运行后不太可能切换显示/隐藏的，就用 v-if。  

v-else 与 v-if 结对，当 v-if 为 false 时执行  
```html
<div id="app">
  <p v-if="if_1">我是if</p>
  <p v-else>我是else</p>
</div>
<script>
let app = new Vue({
  el:"#app",
  data:{
      if_1:false
  }
});
</script>
```

v-for 迭代数组、对象和整数  
```html
<div id="app">
 <div v-for="(item,index) in list">
     {{index}}.{{item}}
 </div>
</div>
<script>
let app = new Vue({
  el:"#app",
  data:{
      list:['Tom','John','Lisa']
  }
});
</script>
```

v-bind 用于动态绑定 DOM 元素的属性  
v-bind 指令可以简写成一个冒号 :  
```html
<div id="app">
  <a :href="link">hello world</a>
</div>
<script>
let app = new Vue({
  el:"#app",
  data:{
      link:"http://helloworld.com"
  }
});
</script>
```

v-on 指令相当于绑定事件的监听器，绑定的事件触发了，可以指定事件的处理函数，配合 methods 使用  
```html
<div id="app">
 <button v-on:click="say('Tom')">
      点击
</button>
</div>
<script>
let app = new Vue({
  el:"#app",
  methods:{
    say(name){
      console.log('hello,'+name);
    }
  }
});
</script>
```

v-model 一般用在表单输入，它帮助我们轻易地实现表单控件和数据的双向绑定  
```html
 <div id="app">
    <input v-model="msg" type="text">
    <p>你输入的是：{{ msg }}</p>
 </div>
 <script>
 let app = new Vue({
    el:"#app",
    data:{
        msg:''
    }
 });
 </script>
```

v-once 只渲染一次，后面元素中的数据再更新变化，都不会重新渲染，可以用于优化更新性能  
```html
<div id="app">
 <input v-model="msg"  type="text">
 <p v-once>你输入：{{ msg }}</p>
</div>
<script>
let app = new Vue({
  el:"#app",
  data:{
      msg:'hello, everyone'
  }
});
</script>
```
由于 msg 有了初始值，第一次渲染的时候，input 控件和 p 标签都有了内容，由于 p 标签添加了 v-once 指令，所以后期再更新 msg 的值的时候，p 标签的内容不会发生改变。  

v-cloak 指令保持在元素上直到关联实例结束编译  
通常结合 CSS 规则，解决插值表达式"闪烁"问题。  
和 CSS 规则如 [v-cloak] { display: none } 一起用时，这个指令可以隐藏未编译的 Mustache 标签直到实例准备完毕。  
```html
<style>
// 带有v-cloak的标签隐藏
[v-cloak] {
  display: none;
}
</style>

// 标签
<div v-cloak>
  {{ message }}
</div>
```  

#### 2.8 动态绑定 class 和 style
对象语法：动态绑定的 class 的值是一个对象 {}
```html
 <div id="app">
  <!-- 当 isActive 为 true 时，class='active' -->
   <p :class="{'active':isActive}">这是文字</p>
 </div>
 <script>
    //创建一个实例 vm
    const vm = new Vue({
        el:"#app",
        data:{
            isActive:true
        }
    });
 </script>
```

数组语法：用数组语法来绑定 class  
```html
 <!-- 拼接数组内的所有样式 -->
 <p :class="[activeClass,errorClass]">
    这是文字
 </p>

  <script>
    //创建一个实例 vm
    const vm = new Vue({
        el:"#app",
        data:{
            activeClass:'active',
            errorClass:'error'
        }
    });
 </script>
```

绑定内联样式也有 2 种语法：对象语法【常用】和数组语法  
```html
 <p :style="{color:colorStyle}">
   这是文字
 </p> 
 <!-- 组合所有样式 -->
 <div :style="[baseStyles, overridingStyles]">我像风一样自由</div>
 <script>
    //创建一个实例vm
    const vm = new Vue({
        el:"#app",
        data:{
            colorStyle:'red',
            baseStyles:{color:'red'},
            overridingStyles:{'font-size':'10px'}
        }
    });
 </script>
```

### 三、Vuejs 组件
组件是我们人为地把页面合理地拆分成一个个区块，让这些区块更方便我们重复使用，有了组件，我们可以更高效合理地开发和维护我们的项目。  

#### 3.1 创建组件
vue 提供的全局 API: Vue.component() 可以用于创建组件。  
创建之后，可以用 `<my-article></my-article>` 的方式来使用组件。  
```html
<div id="app">
    <my-article></my-article>
</div>
<script>
  // 要确保实例 vm 在创建之前， <my-article> 组件已经被成功注册
  Vue.component('my-article',{
     template:`<div>
              <div>
              <h1>this is a article title</h1>
              <div>
              <span>2019/06/28</span>
              <span>Original</span>
              </div>
              </div>
              <img src="cover.jpg" alt="">
              </div>`
  });

  let vm = new Vue({
     el:"#app",
  });
</script>
```

#### 3.2 组件传参  
不但函数可以接受参数，vue 的组件也可以。  
```html
<div id="app">
    <my-article 
            v-for="item in articles" 
            :detail="item">
    </my-article>
</div>
<script>
  // 要确保实例 vm 在创建之前， <my-article> 组件已经被成功注册
  Vue.component('my-article',{
     // props 接受传参
     props:['detail'],
     template:`<div>
              <div>
              <h1>{{ detail.title }}</h1>
              <div>
              <span>{{ detail.date }}</span>
              <span v-show="detail.is_original">Original</span>
              </div>
              </div>
              <img src="detail.cover_url" alt="">
              </div>`
  });

  let vm = new Vue({
     el:"#app",
     data:{
       articles:[
           {
               title:"first",
               date:" 2019/06/28",
               is_original:true,
               cover_url:"cover.jpg"
           },
           {
               title:"second",
               date:" 2019/06/18",
               is_original:false,
               cover_url:"cover.jpg"
           }
       ]
   }
  });
</script>
```

#### 3.3 组件通信
##### 3.3.1 props 父组件→子组件传递数据  
【参考】 3.2 组件传参  

##### 3.3.2 自定义事件 子组件→父组件传递数据  
每一个 vue 实例都实现了事件接口，我们可以用它提供的 API $emit( eventName) 来触发一个事件。  
```html
<div id="app">
    <son @connect="say"></son>
 </div>
<script>
  Vue.component('son',{
    template:`<button @click="send">
                点击
               </button>`,
    data(){
        return{
            msg:'大家好，我是子组件的数据'
        }
    },
    methods:{
      send(){
          this.$emit('connect',this.msg);
      }
    }
 });

  const app = new Vue({
    el:"#app",
    methods:{
      say(msg){
        console.log(msg);
      }
    }
 });
 </script>
```

##### 3.3.3 非父子组件通信
非父子关系的组件，可以巧妙地利用一个空的 vue 实例来作为一个中央事件总线。但是在实际开发中，我们不会这么做，我们会使用专门用于状态管理的 vuex。  

### 四、Vue 进阶
#### 4.1 transition 组件
除了 CSS3 的 transition 属性，Vue 提供的组件也叫 transition，它们是两个不同的东西。  
Vue提供的 `<transition/>` 组件，会在下列四种情况下起作用：
> 1. 条件渲染（使用v-if）  
> 2. 条件展示（使用了v-show）  
> 3. 动态组件  
> 4. 组件根节点   

当一个被 `<transition/>` 组件包含的节点出现了以上的 4 种情况的任意一种的时候，Vue 自动嗅探目标元素是否应用了 CSS 过渡或动画，如果是，在恰当的时机添加/删除 CSS 类名。  
> 1.v-enter：进入过渡效果（enter）刚刚开始那一刻。在元素被插入或者 show 的时候生效，在下一个帧就会立刻移除，一瞬间的事。  
> 2.v-enter-active: 表示进入过渡（entering）的结束状态。在元素被插入时生效，在 transition/animation 完成之后移除。  
> 3.v-leave: 离开过渡（leave）的开始那一刻。在离开过渡被触发时生效，在下一个帧移除，也是一瞬间的事。  
> 4.v-leave-active：离开过渡（leaving）的结束状态。在离开过渡被触发时生效，在 transition/animation 完成之后移除。  
`注意：.v-enter 中的 v- 只是前缀，如果我们 <transition/> 组件的 name 值为 box，那么它实际的 class 为 .box-enter。`

整个动画的过程大概是这样的：  
> 当它进入过渡的时候（隐藏→显示），就会依次发生：
> 1. 添加 .box-enter 样式  
> 2. 删除 .box-enter 样式，添加 .box-enter-active 样式  
> 3. 删除 .box-enter-active 样式  
>  
> 当它离开过渡的时候（显示→隐藏），就会依次发生：
> 1. 添加 .box-leave 样式  
> 2. 删除 .box-leave 样式，添加 .box-leave-active 样式  
> 3. 删除 .box-leave-active 样式  

```html
 <style>
 /*box节点本身的样式*/
 .box{
    width:100%;
    height:100%;
    text-align: center;
    line-height: 200px;
    background: #ff8282;
    color: #Fff;
    /*以下两个默认值，可不写*/
    /*写上只是为了便于讲解，记住这两个*/
    opacity: 1;
    margin-left: 0;
 }

.box-enter-active,.box-leave-active{
    transition: all .8s;
 }

.box-enter{
    opacity: 0;
    margin-left: 100%;
 }
 
 .box-leave-active{
    opacity: 0;
    margin-left: 100%;
 }
 </style>
 <div id="app" class="app">
     <button 
            @click="showBox = !showBox" 
            class="btn">
        切换
     </button>
    <div class="container">
         <transition name="box">
            <div v-show="showBox" class="box">
                i am the box
            </div>
         </transition>
    </div>
 </div>

  <script>
    const app = new Vue({
        el:"#app",
        data:{
            showBox:false
        }
    });
 </script>
```

#### 4.2 vue-router
vue-router 是 vue 官方的路由插件，它和 vue.js 是深度集成的，适合用于构建 SPA 单页面应用。vue 的单页面应用是基于路由和组件的，相当于传统页面是基于 `<a/>` 标签链接和页面，路由用于设定访问路径，并将路径和组件映射起来，这样就可以实现通过路由 router 来切换组件。
```html
 <div id="app">
    <!--使用 router-link 组件来导航.-->
    <!-- 通过传入 `to` 属性指定链接. -->
    <div class="nav">

      <router-link to="/vue">
        简易vue
      </router-link>

      <router-link to="/es6">
        趣味ES6
      </router-link>

      <router-link to="/career">
        人在职场
   </router-link>

    </div>

    <div class="content">
       !--匹配到的组件将渲染在这里 -->
       <router-view></router-view>
    </div>

 </div>

 <script>
  // 定义路由对应的组件
  // 1.简易 vue 对应的视图组件
  const vueComponent = {
    // template 属性对应的内容，就是将会被替换渲染到 <router-view/> 组件的内容
    template:`<div>
                这里是《简易vue》教程
              </div>`
  };

  // 2.趣味 ES6 对应的视图组件
  const es6Component = {
    template:`<div>
                这里是《趣味ES6》教程
              </div>`
  };

  // 3.人在职场 对应的视图组件
  const careerComponent = {
    template:`<div>
                《混口饭吃》与《工资待遇》
              </div>`
  };

  // 创建 router 实例，并定义导航和组件的映射关系
  const router = new VueRouter({
  // 配置routes
  routes:[
      // 定义 3 个导航和组件的映射关系
      {
          path:"/vue",
          component:vueComponent
      },
      {
          path:"/es6",
          component:es6Component
      },
      {
          path:"/career",
          component:careerComponent
      },
  ]});

   //创建 vue 实例，注入路由 router
   const app = new Vue({
      el:"#app",
      router //此处是 ES6 语法，相当于 router:router
   });
 </script>
```

#### 4.3 vuex


#### 4.4 axios 实现数据请求


