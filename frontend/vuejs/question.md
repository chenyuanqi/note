
### Vue 常见问题
1、什么是 MVVM，与 MVC 有什么区别  
MVVM 模式是由经典的软件架构 MVC 衍生来的。当 View（视图层）变化时，会自动更新到 ViewModel（视图模型），反之亦然。View 和 ViewModel 之间通过双向绑定（data-binding）建立联系。与 MVC 不同的是，它没有 Controller 层，而是演变为 ViewModel。

ViewModel 通过双向数据绑定把 View 层和 Model 层连接了起来，而 View 和 Model 之间的同步工作是由 Vue.js 完成的，我们不需要手动操作 DOM，只需要维护好数据状态。


2、Vue 的生命周期  
[Vue.js 生命周期](https://cn.vuejs.org/v2/api/#%E9%80%89%E9%A1%B9-%E7%94%9F%E5%91%BD%E5%91%A8%E6%9C%9F%E9%92%A9%E5%AD%90) 主要有 8 个阶段：

*   创建前 / 后（beforeCreate / created）：在 beforeCreate 阶段，Vue 实例的挂载元素 el 和数据对象 data 都为 undefined，还未初始化。在 created 阶段，Vue 实例的数据对象 data 有了，el 还没有。
    
*   载入前 / 后（beforeMount / mounted）：在 beforeMount 阶段，Vue 实例的 $el 和 data 都初始化了，但还是挂载之前为虚拟的 DOM 节点，data 尚未替换。在 mounted 阶段，Vue 实例挂载完成，data 成功渲染。
    
*   更新前 / 后（beforeUpdate / updated）：当 data 变化时，会触发 beforeUpdate 和 updated 方法。这两个不常用，且不推荐使用。
    
*   销毁前 / 后（beforeDestroy / destroyed）：beforeDestroy 是在 Vue 实例销毁前触发，一般在这里要通过 removeEventListener 解除手动绑定的事件。实例销毁后，触发 destroyed。


3、Vue.js 2.x 双向绑定原理

这个问题几乎是面试必问的，回答也是有深有浅。基本上要知道核心的 API 是通过 `Object.defineProperty()` 来劫持各个属性的 setter / getter，在数据变动时发布消息给订阅者，触发相应的监听回调，这也是为什么 Vue.js 2.x 不支持 IE8 的原因（IE 8 不支持此 API，且无法通过 polyfill 实现）。

Vue.js 文档已经对 [深入响应式原理](https://cn.vuejs.org/v2/guide/reactivity.html) 解释的很透彻了。  


4、为什么组件中的 data 必须是一个函数，然后 return 一个对象，而 new Vue 实例里，data 可以直接是一个对象？  
因为组件是用来复用的，JS 里对象是引用关系，这样作用域没有隔离，而 new Vue 的实例，是不会被复用的，因此不存在引用对象的问题。  


5、v-show 与 v-if 区别  
`v-show` 只是 CSS 级别的 `display: none;` 和 `display: block;` 之间的切换，而 `v-if` 决定是否会选择代码块的内容（或组件）。  

频繁操作时，使用 `v-show`；一次性渲染完的，使用 `v-if`。

因为当 `v-if="false"` 时，内部组件是不会渲染的，所以在特定条件才渲染部分组件（或内容）时，可以先将条件设置为 `false`，需要时（或异步，比如 $nextTick）再设置为 `true`，这样可以优先渲染重要的其它内容，合理利用，可以进行性能优化。  


6、计算属性和 watch 的区别  
计算属性是自动监听依赖值的变化，从而动态返回内容；监听是一个过程，在监听的值变化时，可以触发一个回调，并做一些事情。所以，区别来源于用法，只是需要动态值，那就用计算属性；需要知道值的改变后执行业务逻辑，才用 watch，用反或混用虽然可行，但都是不正确的用法。

computed 是一个对象时，它有 get 和 set 两个选项。  
methods 是一个方法，它可以接受参数，而 computed 不能；computed 是可以缓存的，methods 不会；一般在 `v-for` 里，需要根据当前项动态绑定值时，只能用 methods 而不能用 computed，因为 computed 不能传参。  
computed 可以依赖其它 computed，甚至是其它组件的 data。  

watch 是一个对象时，它有如下配置：  

*   handler 执行的函数
*   deep 是否深度
*   immediate 是否立即执行


7、

