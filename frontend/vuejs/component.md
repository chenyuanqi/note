
### Vue 组件
组件是可复用的 Vue 实例，且带有一个名字。它们与 new Vue 接收相同的选项，例如 data、computed、watch、methods 以及生命周期钩子等。  
**Vue 组件分类**  
一般来说，Vue.js 组件主要分成三类：

- 由 vue-router 产生的每个页面，它本质上也是一个组件（.vue），主要承载当前页面的 HTML 结构，会包含数据获取、数据整理、数据可视化等常规业务。整个文件相对较大，但一般不会有 props 选项和 自定义事件，因为它作为路由的渲染，不会被复用，因此也不会对外提供接口。

在项目开发中，我们写的大部分代码都是这类的组件（页面），协同开发时，每人维护自己的页面，很少有交集。这类组件相对是最好写的，因为主要是还原设计稿，完成需求，不需要太多模块和架构设计上的考虑。

- 不包含业务，独立、具体功能的基础组件，比如日期选择器、模态框等。这类组件作为项目的基础控件，会被大量使用，因此组件的 API 进行过高强度的抽象，可以通过不同配置实现不同的功能。比如开源的 iView，就是包含了 50 多个这样基础组件的 UI 组件库。

每个公司都有自己的组件使用规范或组件库，但要开发和维护一套像 iView 这样的组件库，投入的人力和精力还是很重的，所以出于成本考虑，很多项目都会使用已有的开源组件库。  
独立组件的开发难度要高于第一类组件，因为它的侧重点是 API 的设计、兼容性、性能、以及复杂的功能。

- 业务组件。它不像第二类独立组件只包含某个功能，而是在业务中被多个页面复用的，它与独立组件的区别是，业务组件只在当前项目中会用到，不具有通用性，而且会包含一些业务，比如数据请求；而独立组件不含业务，在任何项目中都可以使用，功能单一，比如一个具有数据校验功能的输入框。

业务组件更像是介于第一类和第二类之间，在开发上也与独立组件类似，但寄托于项目，你可以使用项目中的技术栈，比如 Vuex、axios、echarts 等，所以它的开发难度相对独立组件要容易点，但也有必要考虑组件的可维护性和复用性。


### Vue 组件 API
Vue.js 组件的三个 API：prop、event、slot。  
一个再复杂的组件，都是由三部分组成的：prop、event、slot，它们构成了 Vue.js 组件的 API。如果你开发的是一个通用组件，那一定要事先设计好这三部分，因为组件一旦发布，后面再修改 API 就很困难了，使用者都是希望不断新增功能，修复 bug，而不是经常变更接口。如果你阅读别人写的组件，也可以从这三个部分展开，它们可以帮助你快速了解一个组件的所有功能。

**属性 prop**  
`prop` 定义了这个组件有哪些可配置的属性，组件的核心功能也都是它来确定的。写通用组件时，props 最好用**对象**的写法，这样可以针对每个属性设置类型、默认值或自定义校验属性的值，这点在组件开发中很重要，然而很多人却忽视，直接使用 props 的数组用法，这样的组件往往是不严谨的。比如我们封装一个按钮组件 `<i-button>`：  

```
<template>
  <button :class="'i-button-size' + size" :disabled="disabled"></button>
</template>
<script>
  // 判断参数是否是其中之一
  function oneOf (value, validList) {
    for (let i = 0; i < validList.length; i++) {
      if (value === validList[i]) {
        return true;
      }
    }
    return false;
  }

  export default {
    props: {
      size: {
        validator (value) {
          return oneOf(value, ['small', 'large', 'default']);
        },
        default: 'default'
      },
      disabled: {
        type: Boolean,
        default: false
      }
    }
  }
</script>

```

使用组件：

```
<i-button size="large"></i-button>
<i-button disabled></i-button>

```

组件中定义了两个属性：尺寸 size 和 是否禁用 disabled。其中 size 使用 `validator` 进行了值的自定义验证，也就是说，从父级传入的 size，它的值必须是指定的 **small、large、default** 中的一个，默认值是 default，如果传入这三个以外的值，都会抛出一条警告。

要注意的是，组件里定义的 props，都是**单向数据流**，也就是只能通过父级修改，组件自己不能修改 props 的值，只能修改定义在 data 里的数据，非要修改，也是通过自定义事件通知父级，由父级来修改。

在使用组件时，也可以传入一些标准的 html 特性，比如 **id**、**class**：

```
<i-button id="btn1" class="btn-submit"></i-button>

```

这样的 html 特性，在组件内的 `<button>` 元素上会继承，并不需要在 props 里再定义一遍。这个特性是默认支持的，如果不期望开启，在组件选项里配置 `inheritAttrs: false` 就可以禁用了。


**插槽 slot**  
如果要给上面的按钮组件 `<i-button>` 添加一些文字内容，就要用到组件的第二个 API：插槽 slot，它可以分发组件的内容，比如在上面的按钮组件中定义一个插槽：

```
<template>
  <button :class="'i-button-size' + size" :disabled="disabled">
    <slot></slot>
  </button>
</template>

```

这里的 `<slot>` 节点就是指定的一个插槽的位置，这样在组件内部就可以扩展内容了：

```
<i-button>按钮 1</i-button>
<i-button>
  <strong>按钮 2</strong>
</i-button>

```

当需要多个插槽时，会用到具名 slot，比如上面的组件我们再增加一个 slot，用于设置另一个图标组件：

```
<template>
  <button :class="'i-button-size' + size" :disabled="disabled">
    <slot name="icon"></slot>
    <slot></slot>
  </button>
</template>

```

```
<i-button>
  <i-icon slot="icon" type="checkmark"></i-icon>
  按钮 1
</i-button>

```

这样，父级内定义的内容，就会出现在组件对应的 slot 里，没有写名字的，就是默认的 slot。

在组件的 `<slot>` 里也可以写一些默认的内容，这样在父级没有写任何 slot 时，它们就会出现，比如：

```
<slot>提交</slot>

```


**自定义事件 event**  
现在我们给组件 `<i-button>` 加一个点击事件，目前有两种写法，我们先看自定义事件 event（部分代码省略）：

```
<template>
  <button @click="handleClick">
    <slot></slot>
  </button>
</template>
<script>
  export default {
    methods: {
      handleClick (event) {
        this.$emit('on-click', event);
      }
    }
  }
</script>

```

通过 `$emit`，就可以触发自定义的事件 `on-click` ，在父级通过 `@on-click` 来监听：

```
<i-button @on-click="handleClick"></i-button>

```

上面的 click 事件，是在组件内部的 `<button>` 元素上声明的，这里还有另一种方法，直接在父级声明，但为了区分原生事件和自定义事件，要用到事件修饰符 `.native`，所以上面的示例也可以这样写：

```
<i-button @click.native="handleClick"></i-button>

```

如果不写 `.native` 修饰符，那上面的 `@click` 就是**自定义事件** click，而非**原生事件** click，但我们在组件内只触发了 `on-click` 事件，而不是 `click`，所以直接写 `@click` 会监听不到。  


### Vue 组件通信
一般来说，组件可以有以下几种关系：

![组件关系](https://user-gold-cdn.xitu.io/2018/10/18/166864d066bbcf69?w=790&h=632&f=png&s=36436)

A 和 B、B 和 C、B 和 D 都是父子关系，C 和 D 是兄弟关系，A 和 C 是隔代关系（可能隔多代）。组件间经常会通信，Vue.js 内置的通信手段一般有两种：

*   `ref`：给元素或组件注册引用信息；
*   `$parent` / `$children`：访问父 / 子实例。

这两种都是直接得到组件实例，使用后可以直接调用组件的方法或访问数据，比如下面的示例中，用 ref 来访问组件（部分代码省略）：

```
// component-a
export default {
  data () {
    return {
      title: 'Vue.js'
    }
  },
  methods: {
    sayHello () {
      window.alert('Hello');
    }
  }
}

```

```
<template>
  <component-a ref="comA"></component-a>
</template>
<script>
  export default {
    mounted () {
      const comA = this.$refs.comA;
      console.log(comA.title);  // Vue.js
      comA.sayHello();  // 弹窗
    }
  }
</script>

```

`$parent` 和 `$children` 类似，也是基于当前上下文访问父组件或全部子组件的。

这两种方法的弊端是，无法在**跨级**或**兄弟**间通信，比如下面的结构：

```
// parent.vue
<component-a></component-a>
<component-b></component-b>
<component-b></component-b>

```

我们想在 component-a 中，访问到引用它的页面中（这里就是 parent.vue）的两个 component-b 组件，那这种情况下，就得配置额外的插件或工具了，比如 Vuex 和 Bus 的解决方案；不过，它们都是依赖第三方插件的存在，这在开发独立组件时是不可取的。


### Vue 的构造器 API：extend 与手动挂载 $mount
这两个 Vue.js 内置但却不常用的 API —— extend 和 $mount，它们经常一起使用。不常用，是因为在业务开发中，基本没有它们的用武之地，但在独立组件开发时，在一些特定的场景它们是至关重要的。  

**应用场景**  
我们在写 Vue.js 时，不论是用 CDN 的方式还是在 Webpack 里用 npm 引入的 Vue.js，都会有一个根节点，并且创建一个根实例，比如：

```
<body>
  <div id="app"></div>
</body>
<script>
  const app = new Vue({
    el: '#app'
  });
</script>

```

Webpack 也类似，一般在入口文件 main.js 里，最后会创建一个实例：

```
import Vue from 'vue';
import App from './app.vue';

new Vue({
  el: '#app',
  render: h => h(App)
});

```

因为用 Webpack 基本都是前端路由的，它的 html 里一般都只有一个根节点 `<div id="app"></div>`，其余都是通过 JavaScript 完成，也就是许多的 Vue.js 组件（每个页面也是一个组件）。

有了初始化的实例，之后所有的页面，都由 vue-router 帮我们管理，组件也都是用 `import` 导入后局部注册（也有在 main.js 全局注册的），不管哪种方式，组件（或页面）的创建过程我们是无需关心的，只是写好 `.vue` 文件并导入即可。这样的组件使用方式，有几个特点：

1.  所有的内容，都是在 `#app` 节点内渲染的；
2.  组件的模板，是事先定义好的；
3.  由于组件的特性，注册的组件只能在当前位置渲染。

比如你要使用一个组件 `<i-date-picker>`，渲染时，这个自定义标签就会被替换为组件的内容，而且在哪写的自定义标签，就在哪里被替换。换句话说，常规的组件使用方式，只能在规定的地方渲染组件，这在一些特殊场景下就比较局限了，例如：

1.  组件的模板是通过调用接口从服务端获取的，需要动态渲染组件；
2.  实现类似原生 `window.alert()` 的提示框组件，它的位置是在 `<body>` 下，而非 `<div id="app">`，并且不会通过常规的组件自定义标签的形式使用，而是像 JS 调用函数一样使用。

一般来说，在我们访问页面时，组件就已经渲染就位了，对于场景 1，组件的渲染是异步的，甚至预先不知道模板是什么。对于场景 2，其实并不陌生，在 jQuery 时代，通过操作 DOM，很容易就能实现，你可以沿用这种思路，只是这种做法不那么 Vue，既然使用 Vue.js 了，就应该用 Vue 的思路来解决问题。对于这两种场景，Vue.extend 和 vm.$mount 语法就派上用场了。

**用法**  

上文我们说到，创建一个 Vue 实例时，都会有一个选项 `el`，来指定实例的根节点，如果不写 `el` 选项，那组件就处于未挂载状态。`Vue.extend` 的作用，就是基于 Vue 构造器，创建一个“子类”，它的参数跟 `new Vue` 的基本一样，但 `data` 要跟组件一样，是个函数，再配合 `$mount` ，就可以让组件渲染，并且挂载到任意指定的节点上，比如 body。

比如上文的场景，就可以这样写：

```
import Vue from 'vue';

const AlertComponent = Vue.extend({
  template: '<div>{{ message }}</div>',
  data () {
    return {
      message: 'Hello, Aresn'
    };
  },
});

```

这一步，我们创建了一个构造器，这个过程就可以解决异步获取 template 模板的问题，下面要手动渲染组件，并把它挂载到 body 下：

```
const component = new AlertComponent().$mount();

```

这一步，我们调用了 `$mount` 方法对组件进行了手动渲染，但它仅仅是被渲染好了，并没有挂载到节点上，也就显示不了组件。此时的 `component` 已经是一个标准的 Vue 组件实例，因此它的 `$el` 属性也可以被访问：

```
document.body.appendChild(component.$el);

```

当然，除了 body，你还可以挂载到其它节点上。

`$mount` 也有一些快捷的挂载方式，以下两种都是可以的：

```
// 在 $mount 里写参数来指定挂载的节点
new AlertComponent().$mount('#app');
// 不用 $mount，直接在创建实例时指定 el 选项
new AlertComponent({ el: '#app' });

```

实现同样的效果，除了用 extend 外，也可以直接创建 Vue 实例，并且用一个 Render 函数来渲染一个 .vue 文件：

```
import Vue from 'vue';
import Notification from './notification.vue';

const props = {};  // 这里可以传入一些组件的 props 选项

const Instance = new Vue({
  render (h) {
    return h(Notification, {
      props: props
    });
  }
});

const component = Instance.$mount();
document.body.appendChild(component.$el);

```

这样既可以使用 .vue 来写复杂的组件（毕竟在 template 里堆字符串很痛苦），还可以根据需要传入适当的 props。渲染后，如果想操作 Render 的 `Notification` 实例，也是很简单的：

```
const notification = Instance.$children[0];

```

因为 Instance 下只 Render 了 Notification 一个子组件，所以可以用 `$children[0]` 访问到。

> 如果你还不理解这样做的目的，没有关系，后面小节的两个实战你会感受到它的用武之地。

需要注意的是，我们是用 `$mount` 手动渲染的组件，如果要销毁，也要用 `$destroy` 来手动销毁实例，必要时，也可以用 `removeChild` 把节点从 DOM 中移除。


### Vue.js 容易忽略的 API

**nextTick**  
nextTick 是 Vue.js 提供的一个函数，并非浏览器内置。nextTick 函数接收一个回调函数 `cb`，在下一个 DOM 更新循环之后执行。比如下面的示例：

```
<template>
  <div>
    <p v-if="show" ref="node">内容</p>
    <button @click="handleShow">显示</button>
  </div>
</template>
<script>
  export default {
    data () {
      return {
        show: false
      }
    },
    methods: {
      handleShow () {
        this.show = true;
        console.log(this.$refs.node);  // undefined
        this.$nextTick(() => {
          console.log(this.$refs.node);  // <p>内容</p>
        });
      }
    }
  }
</script>

```

当 `show` 被置为 true 时，这时 p 节点还未被渲染，因此打印出的是 undefined，而在 nextTick 的回调里，p 已经渲染好了，这时能正确打印出节点。

nextTick 的源码在 [https://github.com/vuejs/vue/blob/dev/src/core/util/next-tick.js](https://github.com/vuejs/vue/blob/dev/src/core/util/next-tick.js)，可以看到，Vue.js 使用了 `Promise`、`setTimeout` 和 `setImmediate` 三种方法来实现 nextTick，在不同环境会使用不同的方法。


**v-model 语法糖**  
`v-model` 常用于表单元素上进行数据的双向绑定，比如 `<input>`。除了原生的元素，它还能在自定义组件中使用。

v-model 是一个语法糖，可以拆解为 props: value 和 events: input。就是说组件必须提供一个名为 value 的 prop，以及名为 input 的自定义事件，满足这两个条件，使用者就能在自定义组件上使用 `v-model`。比如下面的示例，实现了一个数字选择器：

```
<template>
  <div>
    <button @click="increase(-1)">减 1</button>
    <span style="color: red;padding: 6px">{{ currentValue }}</span>
    <button @click="increase(1)">加 1</button>
  </div>
</template>
<script>
  export default {
    name: 'InputNumber',
    props: {
      value: {
        type: Number
      }
    },
    data () {
      return {
        currentValue: this.value
      }
    },
    watch: {
      value (val) {
        this.currentValue = val;
      }
    },
    methods: {
      increase (val) {
        this.currentValue += val;
        this.$emit('input', this.currentValue);
      }
    }
  }
</script>

```

props 一般不能在组件内修改，它是通过父级修改的，因此实现 v-model 一般都会有一个 `currentValue` 的内部 data，初始时从 value 获取一次值，当 value 修改时，也通过 watch 监听到及时更新；组件不会修改 value 的值，而是修改 currentValue，同时将修改的值通过自定义事件 `input` 派发给父组件，父组件接收到后，由父组件修改 value。所以，上面的数字选择器组件可以有下面两种使用方式：

```
<template>
  <InputNumber v-model="value" />
</template>
<script>
  import InputNumber from '../components/input-number/input-number.vue';

  export default {
    components: { InputNumber },
    data () {
      return {
        value: 1
      }
    }
  }
</script>

```

或：

```
<template>
  <InputNumber :value="value" @input="handleChange" />
</template>
<script>
  import InputNumber from '../components/input-number/input-number.vue';

  export default {
    components: { InputNumber },
    data () {
      return {
        value: 1
      }
    },
    methods: {
      handleChange (val) {
        this.value = val;
      }
    }
  }
</script>

```

如果你不想用 `value` 和 `input` 这两个名字，从 Vue.js 2.2.0 版本开始，提供了一个 `model` 的选项，可以指定它们的名字，所以数字选择器组件也可以这样写：

```
<template>
  <div>
    <button @click="increase(-1)">减 1</button>
    <span style="color: red;padding: 6px">{{ currentValue }}</span>
    <button @click="increase(1)">加 1</button>
  </div>
</template>
<script>
  export default {
    name: 'InputNumber',
    props: {
      number: {
        type: Number
      }
    },
    model: {
      prop: 'number',
      event: 'change'
    },
    data () {
      return {
        currentValue: this.number
      }
    },
    watch: {
      value (val) {
        this.currentValue = val;
      }
    },
    methods: {
      increase (val) {
        this.currentValue += val;
        this.$emit('number', this.currentValue);
      }
    }
  }
</script>

```

在 model 选项里，就可以指定 prop 和 event 的名字了，而不一定非要用 value 和 input，因为这两个名字在一些原生表单元素里，有其它用处。  


