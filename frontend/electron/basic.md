
### Electron 是什么
Electron 是用 HTML、CSS 和 JavaScript 来构建跨平台桌面应用程序的一个开源库。  
Electron 将 Chromium 和 Node.js 合并到同一个运行环境中，也是就说 Electron 其实就是谷歌浏览器加 Node.js，但你使用 Electron 来装载界面的时候，你不用考虑浏览器的兼容性，只要你的前端项目能够在谷歌浏览器中正常运行，那么你的项目就可以正常地在 Electron 中运行。  

用 Electron 开发的桌面应用是可以跨平台的，在开发模式下或者是生成模式下同样的一套代码可以运行在不同平台，打包之后的应用可以运行在 Windows、Mac 和 Linux 等不同的平台上。  
Electron 中引入一个很重的机制，主进程和渲染进程，关于主进程和渲染进程。  
用 Electron 写项目的时候，你也不用担心代码的调试问题，Electron 中引入来谷歌浏览器的开发者工具，只需要一段很少的代码就在项目中配置谷歌浏览器调试工具。  
Electron 给我们提供来很丰富的 API，这些 API 在你编码的过程中将会给你减去不少的代码量。  

### Electron 环境搭建
搭建 Electron 环境前，需要先安装 Nodejs。  
```bash
# 项目安装
npm install Electron --save-dev

# 全局安装
npm install Electron -g

# 自定义安装 - 指定 ia32 位数版本
npm install --arch=ia32 Electron
# 自定义安装 - 指定平台安装
npm install --platform=win32 Electron
```

### Electron 原理
Electron 是一个 Chromium + Node.js + NativeApis 的项目，这个项目集成的这些组件是 Electron 的核心 Chromium 是 Google 为发展 Chrome 浏览器而启动的开源项目，Chromium 相当于 Chrome 的工程版或称实验版（尽管 Chrome 自身也有 β 版阶段），新功能会率先在 Chromium 上实现，待验证后才会应用在 Chrome 上，故 Chrome 的功能会相对落后但较稳定。Electron API 就像 Node 一样，被设计成支持用户开发模块和应用程序。  

**主进程**  
ipcMain 是 Electron 的主进程的 EventEmitter 的实例，当在主进程中使用时，它处理从渲染器进程（网页）发送出来的异步和同步信息。 从渲染器进程发送的消息将被发送到该模块。  

- 发送消息时，事件名称为 channel
- 回复同步信息时，需要设置 event.returnValue
- 将异步消息发送回发件人，需要使用 event.sender.send (...)

**渲染进程**  
ipcRenderer 是渲染进程的 EventEmitter 的实例。 你可以使用它提供的一些方法从渲染进程（web 页面）发送同步或异步的消息到主进程。也可以接收主进程回复的消息。

**重要组件及事件**  
app 主要用于控制整个应用程序的生命周期。  
```javascript
const {app} = require('Electron')
// window-all-closed 事件是当所有窗口关闭时触发，这里处理 app 组件自动退出
app.on('window-all-closed', () => {
	app.quit()
})
```

autoUpdater 使应用程序能够自动更新，目前只支持 Windows 和 Mac。  
```javascript
import { autoUpdater } from 'Electron-updater'

// update-downloaded 是一个自动更新下载事件
autoUpdater.on('update-downloaded', () => {
  autoUpdater.quitAndInstall()
})

app.on('ready', () => {
  if (process.env.NODE_ENV === 'production') {
  	autoUpdater.checkForUpdates()
  }
})
```

BrowserWindow 创建和控制浏览器窗口。  
```javascript
const {BrowserWindow} = require('Electron')
let win = new BrowserWindow({width: 800, height: 600})
// closed 是关闭窗口事件
win.on('closed', () => {
	win = null
})
win.loadURL('https://github.com')
win.loadURL(`file://${__dirname}/app/index.html`)
```

static 方法是静态方法，是属于类本身的方法而不是类对象的方法。比如  
> BrowserWindow.getAllWindows () 返回 BrowserWindow []，返回所有窗口  
> BrowserWindow.getFocusedWindow () 返回当前锁定的窗口  
> BrowserWindow.fromWebContents (webContents) 从网页获得一个窗口  

webContents 渲染以及控制 web 页面。  
```javascript
const {BrowserWindow} = require('Electron')

let win = new BrowserWindow({width: 800, height: 1500})
win.loadURL('http://github.com')

let contents = win.webContents
console.log(contents)
```

<webview\> 标签是在一个独立的 frame 和进程里显示外部 web 内容。  

BrowserWindowProxy 操纵子浏览器窗口，使用 window.open 创建一个新窗口时会返回一个 BrowserWindowProxy 对象，并提供一个有限功能的子窗口。  

使用默认应用程序管理文件和 url，shell 模块提供与桌面集成相关的功能。  
```javascript
const {shell} = require('Electron')
// 在用户的默认浏览器中打开 URL 
shell.openExternal('https://github.com')
```
