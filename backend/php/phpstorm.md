
### PHPStorm 常用链接
* [插件](http://plugins.jetbrains.com/phpstorm) - 官网插件。
* [注册](http://idea.lanyus.com/) - IntelliJ IDEA 注册码。
* [小贴士](https://phpstorm.tips/) - 操作秀起来。  

### PHPStorm 常用插件
* **Laravel Plugin** - 支持 Laravel
* **.env files support** - 支持.env 文件
* **Ideolog** - 有好的插件 .log 文件
* **Material Theme UI** - Material Theme 主题
* **.ignore** - 友好的查看 .ignore 文件
* **LiveEdit** - 可以实时编辑 HTML/CSS/JavaScript
* **Markdown Navigator** - 支持 Markdown
* **PHP composer.json support** - 支持 composer.json 文件
* **Php Inspections (EA Extended)** - PHP 的静态代码分析工具
* **Styled Components** - 利用标记的模板文字和 CSS 的强大功能，样式化组件允许您编写实际的 CSS 代码来设置组件样式
* **Translation** - 最好用的翻译插件
* **ApiDebugger** - 一个开源的接口调试插件
* **[PHP Annotations](https://github.com/Haehnchen/idea-php-annotation-plugin)** - PhpStorm 注释
* **[BrowseWordAtCaret](https://plugins.jetbrains.com/plugin/201-browsewordatcaret)** - 高亮选中词语

### PHPStorm 优化
- Java VM options
> PHPStorm 依赖 java 运行环境，说白了也就是 java 虚拟机，找到`help > Edit Custom VM Options`，然后在这个文件里可以根据需要增加或减少 PHPstorm 使用的内存
```
-Xms512m
-Xmx2048m

-Dawt.useSystemAAFontSettings=lcd
-Dawt.java2d.opengl=true

# 这一条只适合于Mac, 可以使java调用优化过的图形引擎
-Dapple.awt.graphics.UseQuartz=true

```

- 自定义 properties
> 进入`help > Edit Custom Properties`来设置 PHPStorm 的自定义属性.
```
editor.zero.latency.typing=true
```
上面这条，改变的是 PHPstorm 如何渲染字体：立即渲染文字，而不是先进行内容分析。可能会因此导致偶尔有那么一瞬间文字都是不带样式的，但是整体上会顺畅很多。

- Inspections and plugins（检查和插件）
> PHPstorm 的一大问题就是太强大了，默认加了很多功能，而我们可能平时根本用不到。  
> 找到`preferences -> plugins`，把我们根本用不到的很多插件`plugin`，禁用掉！  
> 不要担心禁的太多，如果你勾掉一个插件的时候，它又被另外一个插件依赖，它会提示你的；而且，在特定的情境下，当 PHPstorm 觉得你应该启用一个插件的时候，它也会提示你的。  
> 
> 禁用不必要的插件是第一步，但是禁用代码检查（inspections），往往可能影响更大。找到`Settings > Editor > Inspections`，根据自己的情况看看哪些时候其实不需要实时的代码检查

- Language injection（其它语言的插入）
> 有一个插件其实特别影响性能，就是`IntelliLang`. 这个插件支持一种语言在其他的文件格式中也照样能被识别，比如说当你在一个 PHP 文件中插入 HTML，或者用到 HTML 的代码自动补齐或高亮显示功能时。  
> 当然，并不建议完全禁用掉这个插件，但是呢，可能有些特定的语言插入支持，你并不会用到，这个时候你可以到`Settings > Editor > Language Injections`下，把当前项目里不可能用到的第三方语言插入，都勾掉。

- 排除对特定项目目录的索引
> 在 `Settings > Directories` 下可以将特定的目录标记排除，然后 PHPstorm 就不会索引其中的文件了。建议排除的目录一般是类似`cache`、`public`、`storage`等包含资源编译文件的，当然还有两个大头，就是`vendor` 和`node_modules`目录。

- vendor 目录的问题
> 排除掉`vendor`目录，意味着就不能基于那里面的组件进行自动补全（auto-complete）了，所以这可能不是个好主意。但是呢，有个小技巧就是，你可以整体上排除掉`vendor`目录，然后在`Settings > Languages & Frameworks > PHP`下，将你真正用到的组件目录给额外添加上。

- Node modules 目录
> `Node modules`目录实际上默认已经被排除掉了，但是呢，在 `Settings > Languages & Frameworks > JavaScript > Libraries`下，你会看到，它们又被额外引入进来了，假设说你写 js 不是那么多，你也可以在这里将其完全排除掉。当然这些呢，都是基于项目的，你可以在不同的项目作不同的选择。

- 删除之前版本的 phpstorm 缓存文件夹
> 经常，每次你更新了 PHPstorm，它就会创建一个新的 cache 文件，而不会自动删除你上一个版本的 cache 文件夹，这往往会占用大量的系统盘空间，如果你用了某一个版本的 PHPstorm 很长时间，这个文件夹一般都是好几 GB。  
> 
> 在 Mac 上，你可以查找类似`‘PhpStorm2016.x’`或`~/Library/Caches`的文件夹，然后删除它；  
> 在 windows 上，在你的当前用户目录查找类似`.WebIde`的文件夹，将多出来的删掉。

- phpstorm 不断重新索引（re-index）的问题
> 在新近版本的 PS 上，这里一般指 2018 以后的版本，你可能会发现 PS 老是自动地不断重新索引，在右下角会出现 indexing 的状态条，导致你编辑文件的操作很容易被打断，或者说 PS 的各种方便的提示、补全功能会因此而暂停，遇到这个问题，可以尝试如下解决：

* 重建缓存： `File > Invalidate Caches/Restart`, 也即重新生成缓存，然后重启，一般此问题会解决
* 如果第一步没解决问题，那么到 plugin 中禁用掉`.ignore`插件，该插件存在的 bug 也会导致此问题，重启后观察解决情况