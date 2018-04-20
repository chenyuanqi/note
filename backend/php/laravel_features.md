
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
- 「[Laravel Horizon](http://laravelacademy.org/post/8492.html)」  
> Horizon 为基于 Redis 的 Laravel 队列提供了一个美观的后台和代码驱动的配置。Horizon 允许你轻松监控队列系统的关键指标，例如任务吞吐量、运行时和失败任务。
> 所有任务进程配置都存放在一个简单独立的配置文件中，从而允许你的配置保存在源码控制系统中以便整个团队的协作。

- 「[包自动发现](http://laravelacademy.org/post/8476.html)」  
> barryvdh/laravel-debugbar 扩展包可以自动发现并为你注册服务提供者和门面

- 「[API 资源](http://laravelacademy.org/post/8223.html)」  
> 资源类允许你简单、优雅地将模型和模型集合转化成 JSON  

- 「控制台命令自动注册」  
> 当创建新的控制台命令时，不再需要手动将其添加到 Console Kernel 的 $commands 属性列表中  
> 取而代之，在 kernel 的 commands 方法中会调用一个新的 load 方法，该方法会检索给定目录下的所有控制台命令并自动注册它们

- 「新的前端预置功能」  
> 可以使用 preset 命令将默认的 Vue 脚手架切换到 React 脚手架：php artisan preset react  
> 也可以使用 none 预置指令整个移除 JavaScript 和 CSS 框架脚手架：php artisan preset none  

- 「队列任务链」  
> 任务链允许你指定需要在一个序列中运行的队列任务列表，如果这个序列中的某个任务运行失败了，那么剩下的任务就不会再运行，要执行一个队列任务链，你可以在任意分发任务中使用 withChain 方法

- 「队列任务频率限制」  
> 应用队列使用的是 Redis，可以通过时间或并发量来控制队列任务的执行。  
```php
Redis::throttle('key')->allow(10)->every(60)->then(function () {
    // Job logic...
}, function () {
    // Could not obtain lock...

    return $this->release(10);
});
```

- 「基于时间的任务尝试」  
> 作为定义一个任务最终失败之前尝试次数的可选方案，你现在可以一个任务的超时时间，这样的话在给定时间范围内该任务就可以尝试很多次，要定义这样的一个时间，可以添加一个 retryUntil 方法到任务类  

- 「验证规则对象」  
> 通过 Artisan 命令 make:rule 将会在 app/Rules 目录下生成一个新的验证规则  
> 一个规则对象只包含两个方法：passes 和 message。passes 方法接收属性值和名称，然后基于属性值是否有效返回 true 或 false。message 方法会在验证失败时返回对应验证错误消息  

- 「[集成 Trusted Proxy](https://laravel-news.com/trusted-proxy)」  
> 直接集成「[Trusted Proxy](https://github.com/fideloper/TrustedProxy)」扩展包，并新增了一个 App\Http\Middleware\TrustProxies 中间件，这个中间件允许你快速定制化需要被应用信任的代理  

- 「按需通知」  
> 发送通知给应用中的非用户实体，使用新的 Notification::route 方法，可以在发送通知之前指定特别指定通知路由  

- 「可渲染的邮件对象」  
> 邮件对象可以直接从路由返回，从而快速在浏览器中预览邮件设计  

- 「可渲染 & 可报告异常」  
> 直接在异常中定义一个 render 方法，这样就可以直接在这个方法中设置自定义响应渲染逻辑，从而避免在异常处理器中堆积条件判断逻辑  
> 如果你还想要为异常自定义报告逻辑，可以在该类中定义一个 report 方法

- 「请求验证」  
> Illuminate\Http\Request 对象提供了一个 validate 方法，该方法允许你快速验证来自路由闭包或控制器的输入请求  
```php
use Illuminate\Http\Request;

Route::get('/comment', function (Request $request) {
    $request->validate([
        'title' => 'required|string',
        'body' => 'required|string',
    ]);

    // ...
});
```

- 「一致的异常处理」  
> 所有的 JSON 验证错误格式都可以通过在 App\Exceptions\Handler 类中定义单独的方法进行控制  
> 针对 JSON 验证响应的默认格式遵循以下约定  
```php
{
    "message": "The given data was invalid.",
    "errors": {
        "field-1": [
            "Error 1",
            "Error 2"
        ],
        "field-2": [
            "Error 1",
            "Error 2"
        ],
    }
}
```

- 「缓存锁」  
> 支持Redis 和 Memcached 缓存驱动获取和释放原子”锁”，该功能提供了一个获取任意锁的简单方法而不必担心任何竞争条件  

- 「Blade 优化」  
> Blade 提供一个 Blade::if 方法帮助你使用闭包快速定义自定义条件指令  

- 「[新的路由方法](https://laravel-news.com/laravel-5-5-router-view-and-redirect)」  
> Route::redirect 可以定义一个重定向到另一个 URI 的路由  
> 如果路由只需要返回一个视图，可以使用 Route::view  

- 「“粘性”数据库连接」  
> sticky 选项是可选的值，可用于允许在当前请求生命周期内立即读取刚刚写入数据库的记录  
> 如果 sticky 选项被开启并且在当前请求生命周期内在数据库上进行了一次”写”操作，任意后续的”读”操作将会使用”写”连接，这样就可以确保任何在当前请求周期内写入的数据可以立即在同一个请求生命周期内被正确地从数据库读取  
> 这可以看作是解决分布式数据库主从延迟的一种方案。

### Laravel 5.4 新特性

### Laravel 5.3 新特性

### Laravel 5.2 新特性

### Laravel 5.1 新特性
