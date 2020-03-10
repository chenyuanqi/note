
### Egret 白鹭引擎
Egret 是一套 HTML5 游戏开发解决方案，使用 Egret 开发的游戏可以轻松发布到 HTML5、iOS、Android、微信小游戏、Facebook、QQ 玩一玩、百度小游戏、Blockchain Game 等各个平台运行。  

Egret 包含以下生态：

- Egret Launcher：相当于一个启动器，负责 Egret 所有构件的管理和使用。  
- Egret Engine：遵循 HTML5 标准的 2D 引擎及全新打造的 3D 引擎。  
- Egret Wing：支持主流开发语言与技术的编辑器。  
- Egret Native：将基于 Egret 引擎开发的 HTML5 游戏轻松转换为 Android 或 iOS 的原生游戏。  
- ResDepot：可视化资源管理工具。  
- Texture Merger：纹理集打包和动画转换工具。  
- Egret Inspector：供 Chrome 开发者使用的插件。  
- Egret Feather：粒子编辑器。  
- DragonBones：面向设计师的 2D 游戏动画和富媒体内容创作平台。  

egret 游戏默认是 30 帧，代码需要用 typeScript 进行开发。大部分 api 是以 egret 开头的，读取资源的是用 RES；每个 egret 应有且只有一个舞台（也就是 stage 对象），舞台是 egret 显示架构中最根本的显示容器，舞台的坐标原点位于左上角。

### Egret 常用 Api

```javascript
// 关于文本
let label:egret.TextField = new egret.TextField(); 
label.text = "hello world!"; 

// 关于图片
let img:egret.Bitmap = new egret.Bitmap();
img.texture = RES.getRes("imgName");

// 关于形状
// 画个红色矩形框
let shp:egret.Shape = new egret.Shape();
shp.graphics.beginFill( 0xff0000, 1);
shp.graphics.drawRect( 0, 0, 100, 200 );
shp.graphics.endFill();

// 关于声音
let sound:egret.Sound = RES.getRes("mp3Name");
sound.play();
sound.stop();

// 关于事件
// 触摸事件（相当于点击）
this.addEventListener(egret.TouchEvent.TOUCH_TAP, this.onTouchTap, this);
this.removeEventListener(egret.TouchEvent.TOUCH_TAP, this.onTouchTap, this);

// 关于计时器
// 参数为时间间隔（ms）和执行次数
let timer:egret.Timer = new egret.Timer(500, 5); 
// 边计时边触发
timer.addEventListener(egret.TimerEvent.TIMER, this.timerFunc, this);
// 计时结束触发
timer.addEventListener(egret.TimerEvent.TIMER_COMPLETE, this.timerComFunc, this);
// 开始计时
timer.start();
// 暂停计时
timer.stop();
// 重新计时
timer.reset();

// 关于数据存储
// 存储数据
let key:string = "score";
let value:string = "100";
egret.localStorage.setItem(key, value);
// 读取数据
let score:string = egret.localStorage.getItem(key);
// 移除数据
egret.localStorage.removeItem(key);
// 清除所有数据
egret.localStorage.clear();
```

