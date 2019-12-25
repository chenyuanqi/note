
### 什么是 Cocos2d
Cocos2d-x 是 MIT 许可证下发布的一款功能强大的开源游戏引擎，允许开发人员使用 C++、Javascript 及 Lua 三种语言来进行游戏开发，它支持所有常见平台，包括 iOS、Android、Windows、macOS、Linux。  

Cocos2d-x 用户不仅包括个人开发者和游戏开发爱好者，还包括许多知名大公司如 Zynga、Wooga、Gamevil、Glu、GREE、Konami、TinyCo、HandyGames、IGG 及 Disney Mobile 等。  
使用 Cocos2d-x 开发的许多游戏占据苹果应用商店和谷歌应用商店排行榜，同时许多公司如触控、谷歌、微软、ARM，英特尔及黑莓的工程师在 Cocos2d-x 领域也非常活跃。  
在中国，每一年的手游榜单大作，Cocos2d-x 从未缺席，市场份额占 50% 以上，游戏品类覆盖从轻度休闲，热火棋牌，到横版，SLG，重度 MMO 等市面全品类。  

[Cocos2d-x 官网](https://www.cocos.com/)  
[Cocos2d-x api](https://docs.cocos2d-x.org/api-ref/index.html)  
[Cocos2d-x Github](https://github.com/cocos2d/cocos2d-x)  
[Cocos2d-x 测试项目](https://github.com/cocos2d/cocos2d-x/tree/v3/tests)  
[Cocos 引擎社区](https://forum.cocos.org/)  
[Cocos 引擎微博](https://weibo.com/cocos2dx?is_hot=1)  

[Cocos ide 旧版](http://www.cocos2d-x.org/filedown/cocos-code-ide-win64-1.2.0.exe)  
[Cocos ide 新版](https://www.cocos.com/creator)  

游戏引擎是一种特殊的软件，它提供游戏开发时需要的常见功能；引擎会提供许多组件，使用这些组件能缩短开发时间，让游戏开发变得更简单；专业引擎通常会能比自制引擎表现出更好的性能。游戏引擎通常会包含渲染器，2D/3D 图形元素，碰撞检测，物理引擎，声音，控制器支持，动画等部分。  
Cocos2d-x 就是这样的一个游戏引擎，它提供了许多易于使用的组件，有着更好的性能，还同时支持移动端和桌面端。Cocos2d-x 通过封装底层图形接口提供了易用的 API，降低了游戏开发的门槛，让使用者可以专注于开发游戏，而不用关注底层的技术细节。更重要的是 Cocos2d-x 是一个完全开源的游戏引擎，这就允许您在游戏开发过程中根据实际需要，定制化引擎的功能，如果您想要一个功能但又不知如何修改，提出这个需求，全世界的开发者可以一起为您完成。  
只使用 Cocos2d-x 引擎，你就能完成一款游戏的开发，因为 Cocos2d-x 提供了游戏开发所需的一切。  

**引擎特性**  

- 现代化的 C++ API
- 立足于 C++ 同时支持 JavaScript/Lua 作为开发语言
- 可以跨平台部署，支持 iOS、Android、Windows、macOS 和 Linux
- 可以在 PC 端完成游戏的测试，最终发布到移动端
- 完善的游戏功能支持，包含精灵、动作、动画、粒子特效、场景转换、事件、文件 IO、数据持久化、骨骼动画、3D

**基本概念**  

- 导演 (Director)  
一个常见的 Director 任务是控制场景替换和转换。  
Director 是一个共享的单例对象，可以在代码中的任何地方调用。当游戏设计好时，Director 就负责场景的转换。  

- 场景 (Scene)  
场景是被渲染器 (renderer) 画出来的，场景是由很多小的对象拼接而成，所有的对象组合在一起，形成了最终的结果。  
渲染器负责渲染精灵和其它的对象进入屏幕。

- 场景图 (Scene Graph)  
场景图是一种安排场景内对象的数据结构，它把场景内所有的节点 (Node) 都包含在一个树 (tree) 上。  

- 精灵 (Sprite)  
所有的游戏都有精灵这个对象，精灵是在屏幕上移动的对象，它能被控制。  
精灵很容易被创建，它有一些可以被配置的属性，比如：位置，旋转角度，缩放比例，透明度，颜色等等。  
精灵 (Sprite) 和节点 (Node) 的区别：如果你能控制它，它才是一个精灵，如果无法控制，那就只是一个节点。  

- 锚点 (anchor point)  
所有的节点 (Node) 对象都有锚点值，精灵 (Sprite) 是节点 (Node) 的子类，自然也具有锚点。  
锚点是节点对象在计算坐标位置时的一个基准点。  

- 动作 (Action)  
动作可以让精灵在场景中移动，比如从一个点移动到另外一个点；比如还可以创建一个动作序列 (Sequence) ，让精灵按照这个序列做连续的动作，在动作过程中你可以改变精灵的位置，旋转角度，缩放比例等等。  

- 序列 (Sequence)   
序列就是多个动作按照特定顺序的一个排列，当然反向执行这个序列也是可以的。  

- 节点关系  
Cocos2d-x 的 节点关系，是被附属和附属的关系，就像数据结构中的父子关系，如果两个节点被添加到一个父子关系中，那么父节点的属性变化会被自动应用到子节点中。  
需要注意的是，不是所有的父节点属性都会被自动应用到子节点。  

- 日志输出  
游戏正在运行的时候，为了了解程序的运行过程或者为了查找一个 BUG，这时候可以使用 log() 在控制台看到一些运行时信息。  
