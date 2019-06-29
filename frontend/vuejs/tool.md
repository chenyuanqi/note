
### Vue Devtools
[github](https://github.com/vuejs/vue-devtools)  
[chrome plugin](https://chrome.google.com/webstore/detail/vuejs-devtools/nhdogjmejiglipccpnnnanhbledajbpd?hl=zh-CN)  

### 安装
方式 1：chrome plugin 直接安装  
方式 2：手动安装  
```bash
git clone https://github.com/vuejs/vue-devtools.git
npm install
npm run build
```
浏览器输入：chrome://extensions/   
选择开发者模式，点击按钮【加载已解压的扩展程序...】，选择 build 生成的文件夹即可。  

### 日常使用
调试 vue 应用的时候，chrome 开发者工具中会看一个 vue 的一栏，点击之后就可以看见当前页面 vue 对象的一些信息。

在 Vue 面板，点击 Components 选项卡，我们会看到这个页面的组件和它们的所有状态，一览无余。当我们的状态改变的时候，Vue Devtools 也会及时更新，相当的方便。  

在 Vue 面板，Vuex 是 Vue 的核心插件之一，专门用来管理组件和应用的状态。

在 Vue 面板，Events 选项卡用来监视我们自定义的事件，（注意：不是原生的事件），它可以清晰地看到你触发的每一个事件的详细信息。