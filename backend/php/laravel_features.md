
### Laravel 版本发布路线
- Laravel 5.6 - 2018 年 2 月份
> 一般发行版，提供 6 个月的 Bug 修复支持，一年的安全修复支持

- Laravel 5.5 – 2017 年 8 月份
> LTS 长久支持版本，会从这一刻开始停止 Laravel 5.1 的 Bug 修复，安全修复直到 2018 年 7 月份

- Laravel 5.4 – 2017 年 1 月份
> 一般发行版，提供 6 个月的 Bug 修复支持，一年的安全修复支持

- Laravel 5.3 – 2016 年 8 月份
> 一般发行版，提供 6 个月的 Bug 修复支持，一年的安全修复支持

- Laravel 5.2 – 2015 年 12 月份
> 一般发行版，提供 6 个月的 Bug 修复支持，一年的安全修复支持

- 5.1 LTS – 2015 年 6 月份
> LTS 长久支持版本，Bug 修复直到 2017 年 6 月份，安全修复直到 2018 年 6 月份。

- Laravel 的开始  
> 2011 年 6 月 9 日，Laravel 的创始人 Taylor Otwell 发布第一个测试版本

### Laravel 5.6 新特性
- 「[日志优化](http://laravelacademy.org/post/8805.html)」  
> 所有日志配置都存放在新的 config/logging.php 配置文件，轻松构建发送日志消息到多个处理器的日志”堆栈”。

- 「单机任务调度」
> 使用 onOneServer 方法可以指定任务只在一台机器上运行

- 「动态频率限制」
> 指定一个动态的最大请求次数 throttle:rate_limit，便于计算最大请求次数计数

- 「广播频道类」
> 要生成一个频道类，可以使用 Artisan 命令 make:channel（生成的频道类存放到 app/Broadcasting 目录下）  
> 然后，在 routes/channels.php 文件中注册这个频道类  
> 最后，可以将频道的授权逻辑放到频道类的 join 方法  

- 「API 控制器生成」
> 在使用 Artisan 命令执行 make:controller 时使用 --api 开关  

- 「模型序列化优化」
> 模型上已加载的关联关系在队列任务被处理时会自动进行重新加载  

- 「Eloquent 日期转化」
> 在转化声明中指定目标日期格式，该格式会在模型序列化为数组 / JSON 时使用  

- 「[Blade 组件别名](https://laravel-news.com/blade-component-include-aliases)」
> 假设一个 Blade 组件存放在 resources/views/components/alert.blade.php，你可以使用 component 方法将这个组件名从 components.alert 改为别名 alert  
> 然后，就可以使用别名来渲染了 @component('alert')

- 「[新的 Blade Directives](https://laravel-news.com/new-blade-directives-laravel-5-6)」

- 「Argon2 密码哈希」
> 要求 PHP 7.2.0+，支持通过 Argon2 算法进行密码哈希，默认的应用哈希驱动通过新增的 config/hashing.php 配置文件来控制  

- 「UUID 方法」 
> 新增 Str::uuid 和 Str::orderedUuid 方法  
> Str::orderedUuid 方法会生成一个时间戳最靠前的 UUID，通过诸如 MySQL 的数据库来索引，更简单，也更高效  

- 「[Collision](https://github.com/nunomaduro/collision)」  
> 包含一个 dev Composer 依赖，扩展包在通过命令行与 Laravel 应用交互时提供美观的错误报告  

- 「[Bootstrap 4](https://blog.getbootstrap.com/2018/01/18/bootstrap-4/)」  
> 所有前端脚手架、生成的分页链接升级到 Bootstrap 4

- 移除 「Artisan Optimize 命令」

### Laravel 5.5 新特性

### Laravel 5.4 新特性

### Laravel 5.3 新特性

### Laravel 5.2 新特性

### Laravel 5.1 新特性
