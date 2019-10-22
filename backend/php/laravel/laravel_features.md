
### Laravel 版本发布路线
- Laravel 6.0 LTS - 2019 年 9 月份
> LTS 长久支持版本，Bug 修复直到 2021 年 9 月份，安全修复直到 2022 年 9 月份。

- Laravel 5.8 - 2019 年 2 月份
> 一般发行版，提供 6 个月的 Bug 修复支持，一年的安全修复支持

- Laravel 5.7 - 2018 年 9 月份
> 一般发行版，提供 6 个月的 Bug 修复支持，一年的安全修复支持

- Laravel 5.6 - 2018 年 2 月份
> 一般发行版，提供 6 个月的 Bug 修复支持，一年的安全修复支持

- Laravel 5.5 LTS – 2017 年 8 月份
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

### Laravel 6.0 新特性
Laravel 6.0 的发布标志着 laravel 框架开始使用语义化版本，此外，该版本还包含了对 Laravel Vapor 的支持、优化了授权响应、任务中间件、懒集合、子查询优化以及很多其它细节优化。要求 PHP 7.2+。  

- 语义版本号
> Laravel 框架 (Laravel /framework) 包遵循语义版本控制标准。这使得框架与已经遵循此版本控制标准的其他第一方 Laravel 包保持一致，Laravel 的发布周期将保持不变。

- 优化授权响应
> 在此之前，围绕授权策略提供自定义错误信息给终端用户非常困难，Laravel6 提供 Gate::inspect 方法来授权策略响应。  
```php
$response = Gate::inspect('view', $flight);

if ($response->allowed()) {
    // 用户已授权...
}

if ($response->denied()) {
    // 用户未授权，返回响应信息
    echo $response->message();
}
```

- 任务中间件
> 任务中间件允许中间件到队列任务中对其进行过滤。  
> 使用中间件可以避免在任务类的 handle() 方法中编写与主体业务逻辑无关的代码。  
```php
// 在任务类中定义中间件方法
public function middleware()
{
     return [new SomeMiddleware];
}

// 分发任务时可通过through指定中间件
SomeJob::dispatch()->through([new SomeMiddleware]);
```

- 惰性集合
> 对于处理大量数据的集合 (包括 Eloquent 模型集合)，惰性集合是一个改变（既定的）游戏规则者。一个新的 `lighting \Support\LazyCollection` 类利用 PHP 的生成器在处理大型数据集时保持低内存。

- Eloquent 子查询增强
> 详见 https://learnku.com/laravel/t/33324

- Laravel UI
> Laravel 5.x 自带的前端脚手架，现在被分离成一个独立的 laravel/ui Composer 包。这允许在主框架之外，方便迭代 UI 脚手架。  
```bash
# 如果用传统的 Bootstrap/Vue/ 搭建
composer require laravel/ui
php artisan ui vue --auth
```

### Laravel 5.8 新特性
Laravel 5.8 在 Laravel 5.7 的基础上继续进行优化，包括引入新的 Eloquent 关联关系（远层一对一）、优化邮箱验证、基于约定的授权策略类自动注册、DynamoDB 缓存及 Session 驱动、优化任务调度器的时区配置、支持分配多个认证 guard 到广播频道、PSR-16 缓存驱动规范、优化 artisan serve 命令、支持 PHPUnit 8.0、支持 Carbon 2.0、支持 Pheanstalk 4.0，以及多个 bug 修复和可用性的提升。  
要求 PHP 7.1.3+。

- Eloquent HasOneThrough 关联关系
> Eloquent 现在提供了对 HasOneThrough 关联类型的支持。  
> 例如，假设 Supplier 模型类与 Account 模型类之间是一对一关联，并且 Account 模型类与 AccountHistory 模型类之间也是一对一关联，那么我们说 Supplier 模型类与 AccountHistory 模型类之间可以通过 hasOneThrough 方法基于 Account 模型类建立远层的一对一关联
```php
/**
 * Get the account history for the supplier.
 */
public function accountHistory()
{
    return $this->hasOneThrough(AccountHistory::class, Account::class);
}
```

- 自动发现模型对应策略类
> 在之前版本中，每个模型类对应的授权策略类需要在应用的 AuthServiceProvider 中显式注册。  
```php
/**
 * The policy mappings for the application.
 *
 * @var array
 */
protected $policies = [
    'App\User' => 'App\Policies\UserPolicy',
];
```
> Laravel 5.8 引入了模型类对应策略类的自动发现机制，只要模型类和策略类遵循标准的 Laravel 命名约定（在约定的命名空间及对应目录下）。  
> 即策略类都必须位于模型类所在目录的子目录 Policies 中，例如，如果模型类都位于 app 目录下，则策略类必须位于 app/Policies 目录下。此外，策略类的名称必须和模型类相匹配，并且有一个 Policy 后缀，因此，User 模型类对应的策略类就是 UserPolicy。  
> 如果你想要提供自己的策略类发现逻辑，可以使用 Gate::guessPolicyNamesUsing 方法注册自定义的回调。通常，该方法会在应用自带的服务提供者类 AuthServiceProvider 中调用。  
```php
use Illuminate\Support\Facades\Gate;

Gate::guessPolicyNamesUsing(function ($modelClass) {
    // 返回相应的策略类名称...
});
```
`注：任何显式注册在 AuthServiceProvider 中的模型映射策略类优先级高于这种自动发现获取的策略类。`

- PSR-16 缓存规范
> 为了在存储缓存项时允许更细粒度的过期时间并遵守 PSR-16 缓存标准，我们将缓存项的有效期单位从分钟调整到秒。

- 多个广播认证 Guard
> 在 Laravel 之前的版本中，私有和存在广播频道通过应用的默认认证 guard 对用户进行认证。从 Laravel 5.8 开始，你可以分配多个不同的 guard 来对请求进行认证。  
```php
Broadcast::channel('channel', function() {
    // ...
}, ['guards' => ['web', 'admin']])
```

- Token Guard 令牌哈希算法
> Laravel 的 token guard 用于提供最基本的 API 认证，现在支持以 SHA-256 哈希算法对 API 令牌进行存储，这样比存储纯文本令牌更加安全。

- 优化邮箱验证
> Laravel 5.8 通过 SwiftMailer 提供的 egulias/email-validator 扩展包对验证器的底层邮箱验证逻辑进行了优化。之前版本的邮箱验证逻辑偶尔会将有效的邮箱地址，如 example@bär.se 判定为无效。

- 默认调度器时区
> Laravel 允许你使用 timezone 方法自定义调度任务的时区
```php
$schedule->command('inspire')
     ->hourly()
     ->timezone('America/Chicago');

// 不过，当每个调度任务的时区都一样时这样编写代码显得笨重和累赘
// 现在，可以在 app/Console/Kernel.php 文件中定义一个 scheduleTimezone 方法，用来返回在所有调度任务中使用的默认时区
/**
 * Get the timezone that should be used by default for scheduled events.
 *
 * @return \DateTimeZone|string|null
 */
protected function scheduleTimezone()
{
    return 'America/Chicago';
}
```

- Artisan 调用优化
> Laravel 允许你通过 Artisan::call 方法调用 Artisan 命令。  
```php
use Illuminate\Support\Facades\Artisan;
// 在之前发布的版本中，命令的选项通过数组以第二个参数的方式传递到该方法
Artisan::call('migrate:install', ['database' => 'foo']);
// 在 Laravel 5.8 中，可以传递整个命令，包括命令的选项，只需传递一个参数即可
Artisan::call('migrate:install --database=foo');
```

- 测试辅助方法 mock/spy
> 为了让模拟对象更方便，Laravel 测试用例基类中新增了 mock 和 spy 方法，这两个方法会自动绑定模拟类到容器中。
```php
// Laravel 5.7
$this->instance(Service::class, Mockery::mock(Service::class, function ($mock) {
    $mock->shouldReceive('process')->once();
}));

// Laravel 5.8
$this->mock(Service::class, function ($mock) {
    $mock->shouldReceive('process')->once();
});
```

- Eloquent 资源键保留
> 当从路由中返回一个 Eloquent 资源集合时，Laravel 会重置集合的键以便它们以简单的数字顺序呈现。  
> 当使用 Laravel 5.8 时，你可以添加一个 preserveKeys 属性到资源类表明资源类的键是否保留。默认情况下，为了和之前版本的 Laravel 保持一致，这些键会被重置；而如果 preserveKeys 属性值为 true 时，集合键会被保留。  

- 更高阶的 Eloquent 方法 orWhere
> Laravel 5.8 引入了更高阶的 orWhere 方法，允许你以方法链的方式链接作用域，而不必使用闭包。  
```php
// 在之前版本的 Laravel 中，通过 or 查询操作符连接不同的 Eloquent 模型作用域时需要使用闭包回调
// scopePopular and scopeActive methods defined on the User model...
$users = App\User::popular()->orWhere(function (Builder $query) {
    $query->active();
})->get();

// laravel 5.8
$users = App\User::popular()->orWhere->active()->get();
```

- Artisan Serve 命令优化
> 在之前版本的 Laravel 中，Artisan 的 serve 命令会在 8000 端口上提供服务，如果有一个 serve 命令已经在监听这个端口，则再次运行 artisan serve 命令会失败。从 Laravel 5.8 开始，serve 会扫描从 8000 到 8009 之间的所有有效端口，以便你可以一次运行多个 serve 命令。

- Blade 文件映射
> 编译 Blade 模板时，Laravel 现在会添加一行注释到编译后文件的顶部，其中包含了编译前 Blade 模板文件的路径。 

- DynamoDB 缓存 / Session 驱动
> Laravel 5.8 引入了 [DynamoDB](https://aws.amazon.com/dynamodb/) 缓存和 Session 驱动，DynamoDB 是一个由 AWS 提供的无服务器 NoSQL 数据库，dynamodb 缓存的默认配置可以在 Laravel 5.8 的缓存配置文件中找到。

- 支持 Carbon 2.0

- 支持 Pheanstalk 4.0
> Laravel 5.8 提供了对队列库 Pheanstalk ~4.0 版本的支持。  


### Laravel 5.7 新特性
Laravel 5.7 引入了一些新特性并修复了很多 5.6 版本中的 bug，要求 PHP 7.1.3+。  

- 新的目录结构
> resources 目录移除了 assets 子目录并将之前在 assets 目录下的子目录移到 resources 目录下。

- 新的自定义分页
> 提供了一个新的分页方法 linksOnEachSide($number) 来自定义分页器上显示的链接数目。  

- 优化错误消息
> 跟踪动态调用 Eloquent 模型引起的错误消息将变得更加简单

- Laravel Nova
> [Laravel Nova](https://nova.laravel.com/) 是一个专门为 Laravel 应用打造的、美观的、一流的后台管理面板，当然，Nova 的核心功能还是通过 Eloquent 管理底层数据库记录，不过在这个核心功能之上，Nova 还支持过滤器、透镜、动作、动作队列、授权、自定义工具、自定义卡片、自定义字段等额外功能。

- 邮箱验证
> Laravel 框架自带的认证脚手架代码引入了邮箱验证功能，为了实现这个功能，框架自带的 users 表迁移还新增了一个时间戳字段 email_verified_at。  
> 如果想提示新注册用户验证他们的邮箱，User 模型类需要实现 MustVerifyEmail 接口。一旦 User 模型类标记实现 MustVerifyEmail 接口，新注册的用户将会收到一封邮件，其中包含已签名的验证链接，点击这个链接，Laravel 将会自动在数据库中记录验证时间并将用户重定向到你设置的页面。此外，默认 HTTP kernel 中还新增了一个 verified 中间件，用于过滤那些未验证邮箱的用户。  

- 未登录用户授权  
> 在之前版本的 Laravel 中，对于未登录访客的权限判断，无论是 Gate 还是 Policy，都会自动返回 false。不过，在 Laravel 5.7 中，你可以通过声明一个可选的类型提示或者将用户参数定义默认设置为 null，以便访客可以通过授权检查。  
```php
Gate::define('update-post', function (?User $user, Post $post) {
    // ...
});
```

- Symfony Dump Server
> 通过引入扩展包 Laravel Dump Server 在 Laravel 应用中集成了 Symfony 的 dump-server 命令（php artisan dump-server）。  
> 一旦服务器启动，所有对 dump 函数的调用结果将会输出到 dump-server 控制台窗口，而不是显示在浏览器中，从而允许开发者在不打乱 HTTP 响应输出的情况下检查某些值。

- 通知本地化
> 允许开发者通过本地化语言发送通知，甚至能够在通知队列中记住本地化设置。  
```php
// 为了实现这个功能，Illuminate\Notifications\Notification 类新增了一个 locale 方法来设置期望语言，应用会在通知被格式化时切换到 locale 设置语言，并在格式化完成后切换回之前的语言
$user->notify((new InvoicePaid($invoice))->locale('es'));

// 还可以通过 Notification 门面来本地化多个通知实体
Notification::locale('es')->send($users, new InvoicePaid($invoice));
```

- URL 生成器 & 可调用语法
> 除了只接收字符串外，Laravel 的 URL 生成器现在还可以在生成指向控制器动作的 URL 时接收「可调用」语法。  
```php
action([UserController::class, 'index']);
```

- 文件系统读写流
> Laravel 的文件系统现在集成了 readStream 和 writeStream 方法。  
```php
Storage::disk('s3')->writeStream(
    'remote-file.zip',
    Storage::disk('local')->readStream('local-file.zip')
);
```

### Laravel 5.6 新特性
要求 PHP 7.1.3+

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
要求 PHP 7.0+

- 「[Laravel Horizon](http://laravelacademy.org/post/8492.html)」  
> Horizon 为基于 Redis 的 Laravel 队列提供了一个美观的后台和代码驱动的配置。Horizon 允许你轻松监控队列系统的关键指标，例如任务吞吐量、运行时和失败任务。
> 所有任务进程配置都存放在一个简单独立的配置文件中，从而允许你的配置保存在源码控制系统中以便整个团队的协作。

- 「[包自动发现](http://laravelacademy.org/post/8476.html)」  
> barryvdh/laravel-debugbar 扩展包可以自动发现并为你注册服务提供者和门面

- 「[Whoops 优雅报错](https://laravel-news.com/whoops-laravel-5-5)」  
> whoops 扩展包重新回归  

- 「[API 资源](http://laravelacademy.org/post/8223.html)」  
> 资源类允许你简单、优雅地将模型和模型集合转化成 JSON  

- 「控制台命令自动注册」  
> 当创建新的控制台命令时，不再需要手动将其添加到 Console Kernel 的 $commands 属性列表中  
> 取而代之，在 kernel 的 commands 方法中会调用一个新的 load 方法，该方法会检索给定目录下的所有控制台命令并自动注册它们

- 「[新的前端预置功能](https://laravel-news.com/frontend-presets)」  
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

- 「[验证规则对象](https://laravel-news.com/laravel-5-5-custom-validator-rules)」  
> 通过 Artisan 命令 make:rule 将会在 app/Rules 目录下生成一个新的验证规则  
> 一个规则对象只包含两个方法：passes 和 message。passes 方法接收属性值和名称，然后基于属性值是否有效返回 true 或 false。message 方法会在验证失败时返回对应验证错误消息  

- 「[集成 Trusted Proxy](https://laravel-news.com/trusted-proxy)」  
> 直接集成「[Trusted Proxy](https://github.com/fideloper/TrustedProxy)」扩展包，并新增了一个 App\Http\Middleware\TrustProxies 中间件，这个中间件允许你快速定制化需要被应用信任的代理  

- 「按需通知」  
> 发送通知给应用中的非用户实体，使用新的 Notification::route 方法，可以在发送通知之前指定特别指定通知路由  

- 「[可渲染的邮件对象](https://laravel-news.com/render-mailables)」  
> 邮件对象可以直接从路由返回，从而快速在浏览器中预览邮件设计  

- 「[可渲染 & 可报告异常](https://laravel-news.com/custom-exception-reporting)」  
> 直接在异常中定义一个 render 方法，这样就可以直接在这个方法中设置自定义响应渲染逻辑，从而避免在异常处理器中堆积条件判断逻辑  
> 如果你还想要为异常自定义报告逻辑，可以在该类中定义一个 report 方法

- 「[自定义默认报错视图](https://laravel-news.com/custom-exception-reporting)」  
> 自定义 4xx、5xx 页面  

- 「[请求验证](https://laravel-news.com/request-data-validator-improvements)」  
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

- [新增异常抛出方法](https://laravel-news.com/throw_if-throw_unless)  
> 新增 throw_if 和 throw_unless 方法  

- 「缓存锁」  
> 支持 Redis 和 Memcached 缓存驱动获取和释放原子”锁”，该功能提供了一个获取任意锁的简单方法而不必担心任何竞争条件  

- 「[Blade 优化](https://laravel-news.com/bladeif)」  
> Blade 提供一个 Blade::if 方法帮助你使用闭包快速定义自定义条件指令  

- 「[新的路由方法](https://laravel-news.com/laravel-5-5-router-view-and-redirect)」  
> Route::redirect 可以定义一个重定向到另一个 URI 的路由  
> 如果路由只需要返回一个视图，可以使用 Route::view  

- 「“粘性”数据库连接」  
> sticky 选项是可选的值，可用于允许在当前请求生命周期内立即读取刚刚写入数据库的记录  
> 如果 sticky 选项被开启并且在当前请求生命周期内在数据库上进行了一次”写”操作，任意后续的”读”操作将会使用”写”连接，这样就可以确保任何在当前请求周期内写入的数据可以立即在同一个请求生命周期内被正确地从数据库读取  
> 这可以看作是解决分布式数据库主从延迟的一种方案。

- 「[使用 casts 定义数据表数据的类型](https://laravel-news.com/laravel-5-5-pivot-casting)」

- 「[dd 和 dump 加入 colletions](https://laravel-news.com/dd-and-dump-collections)」  

- 「[动态模板](https://laravel-news.com/dd-and-dump-collections)」  
> 使用 View::first

### Laravel 5.4 新特性
- 「[when 加入 colletions](https://laravel-news.com/laravel-collections-when-method)」  
> 允许使用 when 对执行有条件 where 的操作进行处理，而不需要中断链式操作  

- 「[支持 Markdown 编写邮件和通知](https://laravel-news.com/laravel-markdown-emails)」  

- 「[Laravel Dusk](http://laravelacademy.org/post/7047.html)」  
> 提供优雅的、易于使用的浏览器自动化测试 API  

- 「[Laravel Mix](http://laravelacademy.org/post/7047.html)」  
> Laravel Mix 是 Laravel Elixir 的精神继承者，完全基于 Webpack 而不是 Gulp  
> 使用通用 CSS 和 JavaScript 预处理器定义 Laravel 应用的 Webpack 构建步骤提供了流式 API  
> 通过简单的方法链，可以定义流式资源管道  
```php
mix.js('resources/assets/js/app.js', 'public/js')
   .sass('resources/assets/sass/app.scss', 'public/css');
```

- 「[Blade 组件&插槽](http://laravelacademy.org/post/6780.html)」  
> 命名插槽允许我们在单个组件中定义多个插槽  
> 命名插槽可以通过@slot指令进行注入，任意在@slot指令中的内容都会被传递给$slot变量  
```php
@component('alert')
    @slot('title')
        Forbidden
    @endslot

    You are not allowed to access this resource!
@endcomponent
```

- 「[广播中的模型绑定](http://laravelacademy.org/post/6851.html)」  
> 频道路由也可以显式或隐式进行模型绑定  
```php
use App\Order;

Broadcast::channel('order.{order}', function ($user, Order $order) {
    return $user->id === $order->user_id;
});
```

- 「集合高阶消息传递」  
> 集合现在支持“高阶消息传递”，从而精简对集合的操作  
> 目前支持高阶消息传递的集合方法有：contains、each、every、filter、first、map、partition、reject、sortBy、sortByDesc 和 sum  

- 「基于对象的 Eloquent 事件」  
> Eloquent 事件处理器可以被映射到事件对象上，这为我们处理 Eloquent 事件并让其变得易于测试提供了一种直观的方式  

- 「[任务级的重试&超时](http://laravelacademy.org/post/6922.html)」  
> 队列任务的“重试”和“超时”设置可以在任务类中为每一个任务配置独立的“重试”次数和“超时”时间  
```php
namespace App\Jobs;

class ProcessPodcast implements ShouldQueue
{
    /**
     * The number of times the job may be attempted.
     *
     * @var int
     */
    public $tries = 5;

    /**
     * The number of seconds the job can run before timing out.
     *
     * @var int
     */
    public $timeout = 120;
}
```

- 「请求清理中间件」  
> 在默认中间件堆栈中引入了两个新的中间件：TrimStrings 和 ConvertEmptyStringsToNull  

- 「[“实时”门面](https://laravel.com/docs/5.4/mocking)」  
> 可以轻松将任意类实时转化为一个门面，只需要将导入类包裹在 Facades 命名空间中即可  

- 「自定义透视表模型」  
> 所有隶属 belongsToMany 关联关系的透视表模型都使用同一个内置的 Pivot 模型实例，可以为这些数据透视表自定义模型类  

- 「[优化 Redis 集群支持](http://laravelacademy.org/post/6974.html)」  
> 可以在同一个应用中定义Redis连接指向多个主机和多个集群  

- 「[迁移默认字符换长度](https://laravel-news.com/laravel-5-4-key-too-long-error)」  
> 默认使用 utf8mb4 字符编码，该编码支持对 “emojis” 进行排序  
> 在 AppServiceProvider 中调用 Schema::defaultStringLength 方法来实现手动配置迁移命令生成的默认字符串长度  

### Laravel 5.3 新特性
- 「[通知（Notifications）](http://laravelacademy.org/post/6116.html)」  
> Laravel Notifications 为我们提供了简单、优雅的 API 用于在不同的发行渠道中发送通知，例如邮件、SMS、Slack 等等  

- 「[WebSockets／事件广播](http://laravelacademy.org/post/6026.html)」  
> 通过为已私有和已存在的 WebSocket 频道添加频道级认证对事件广播进行了极大的优化和提升  
> Laravel Echo，通过NPM安装的全新的JavaScript包，将和Laravel 5.3一起发布  
> Laravel Echo 用于为订阅频道以及在客户端 JavaScript 应用中监听服务器端事件提供了简单、优美的 API  
> Laravel Echo 包含对 Pusher 和 Socket.io 的支持  

- 「[Laravel Passport（OAuth2 服务器）](http://laravelacademy.org/post/5993.html)」  
> 使用 Laravel Passport让API 认证变得简单  
> Laravel Passport 可以在几分钟内为应用提供一个完整的 Oauth2 服务器实现，Passport 基于 Alex Bilbie 维护的 League OAuth2 server 实现    

- 「[搜索（Laravel Scout）](http://laravelacademy.org/post/6277.html)」  
> Laravel Scout 提供了一个简单的、基于驱动的针对 Eloquent 模型的全文搜索解决方案    
> Laravel Scout 会自动同步更新 Eloquent 记录的搜索索引  

- 「[支持邮寄对象](http://laravelacademy.org/post/6095.html)」  
> 对象可以以一个简单对象的形式表示邮件信息，而不再需要在闭包中自定义邮件信息  

- 「[存储上传文件](http://laravelacademy.org/post/5881.html#ipt_kb_toc_5881_8)」  
> 通过在上传文件实例上使用新的 store 方法使得用户上传文件变得简单，只需要简单调用 store 方法并传入文件保存路径即可  

- 「[Webpack & Laravel Elixir](http://laravelacademy.org/post/5962.html)」  
> gulpfile.js 使用 webpack 编译 javascript  

- 「[前端架构](http://laravelacademy.org/post/5956.html)」  
> 不再从 cdn 中加载前端资源，所有依赖默认定义在 package.json 中  
> 单文件 vue 组件开箱支持  

- 「路由文件」  
> 新的顶层目录 routes 包含两个 http 路由文件 web 和 api  

- 「[闭包控制台命令](http://laravelacademy.org/post/6220.html)」  
> 除了通过命令类定义之外，现在 Artisan 命令还可以在 app/Console/Kernel.php 文件的 commands 方法中以简单闭包的方式定义  
> commands 方法会加载 routes/console.php 文件，从而允许你基于闭包、以路由风格定义控制台命令  

- 「[$loop 变量](http://laravelacademy.org/post/5919.html)」  
> 在 Blade 模板中循环遍历的时候，$loop 变量将会在循环中生效  
```php
@foreach ($users as $user)
    @if ($loop->first)
        This is the first iteration.
    @endif

    @if ($loop->last)
        This is the last iteration.
    @endif

    <p>This is user {{ $user->id }}</p>
@endforeach
```

### Laravel 5.2 新特性
- 「多认证驱动」
> 可以定义多个认证驱动（不只是默认的、基于 session 的认证驱动），还有多个认证模型以及用户表，并且可以独立控制其认证处理（登录、注册、密码重置）

- 「认证脚手架」
> 提供了便捷的方式来创建前台认证视图，只需在终端执行如下 Artisan 命令即可
```bash
# 生成纯文本的、兼容 Bootstrap 样式的视图用于登录、注册和密码重置
# 使用相应路由更新路由文件
php artisan make:auth
```

- 「[隐式模型绑定](http://laravelacademy.org/post/2784.html#ipt_kb_toc_2784_18)」
> 隐式模型绑定使得在路由或控制器中直接注入相应模型实例更加便捷
```php
use App\User;

Route::get('/user/{user}', function (User $user) {
    return $user;
});
```

- 「[中间件组](http://laravelacademy.org/post/2803.html)」
> 中间件组允许你通过单个方便的键来对相关路由中间件进行分组，从而为某个路由一次指定多个中间件

- 「[访问频率限制]」
> 一个新的访问频率限制中间件已经被内置到框架中，从而允许你轻松限制给定 IP 地址在指定时间内对某个路由发起请求的数目
```php
//  限制某个 IP 地址每分钟只能访问某个路由 60 次
Route::get('/api/users', ['middleware' => 'throttle:60,1', function () {
    // something
}]);
```

- 「[数组输入验证](http://laravelacademy.org/post/3279.html)」
> 表单字段的数组输入验证
```php
// 验证给定数组输入字段中的每一个 email 是唯一的
// 使用 * 来指定验证数组字段
$validator = Validator::make($request->all(), [
    'person.*.email' => 'email|unique:users'
]);
```

- 「[Eloquent 全局作用域优化](http://laravelacademy.org/post/2995.html#global-scopes)」
> 全局查询作用域只需实现一个简单的方法 apply 即可

### Laravel 5.1 新特性
- 「LTS」
> Laravel 5.1 是 Laravel 第一个长期支持版本，将会提供两年的 bug 修复和安全修复  
> 这是迄今为止，Laravel提供的最大跨度的支持，并且将会持续为更多的企业用户及普通用户提供稳定平滑的支持

- 「[PSR-2](https://github.com/PizzaLiu/PHP-FIG/blob/master/PSR-2-coding-style-guide-cn.md)」
> PSR-2 被采取为默认风格指南，此外，所有代码生成器已经被更新到生成兼容 PSR-2 语法的代码  

- 「文档」
> Laravel 文档的每一个页面都进行了一丝不苟的审查和引人注目的优化，所有代码示例都被审查并且扩展到更好的支持上下文相关性  

- 「目录结构」
> 为了更好地表达意图，app/Commands 目录被重命名为 app/Jobs，此外，app/Handlers 被合并到 app/Listeners 目录  
> 并不是破坏式的改变，所以使用 Laravel 5.1 并不强制要求更新到新的目录结构  

- 「[事件广播](http://laravelacademy.org/post/198.html)」
> 在很多现代的 web 应用中，web 套接字被用于实现实时的，即时更新的用户接口  
> 当服务器上的某些数据更新后，通常一条消息将会通过 websocket 连接发送到客户端并进行处理  
> 广播 Laravel 事件允许你在服务端代码和客户端 JavaScript 框架之间共享相同的事件名称  

- 「[中间件参数](http://laravelacademy.org/post/57.html)」
> 中间件可以接受额外的自定义参数  

- 「[测试革新](http://laravelacademy.org/post/238.html)」
> 内置的测试功能获得了引入注目的提升，多个新方法提供了平滑的，富有变现力的接口和应用进行交互并测试响应

- 「[模型工厂](http://laravelacademy.org/post/238.html#model-factories)」
> 通过使用模型工厂附带一种简单的方式类创建Eloquent模型存根  
> 模型工厂允许你为 Eloquent 模型定义一系列默认属性，然后为测试或数据库填充生成模型实例  
> 模型工厂还可以利用强大的PHP扩展库 Faker 类生成随机的属性数据  

- 「[Artisan优化](http://laravelacademy.org/post/170.html)」
> Artisan 命令可以通过使用一个简单的，类似路由风格的“签名”（提供了一个非常简单的接口来定义命令行参数和选项）来定义

- 「加密」
> 在之前的 Laravel 版本中，加密是通过 PHP 扩展 mcrypt 来进行处理的  
> 从 5.1 开始，加密改由通过 PHP 的另一个扩展 openssl 进行处理，因为该扩展较前者而言维护的更加活跃  
