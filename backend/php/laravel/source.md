
### Laravel 生命周期
整体流程：  
入口（public/index.php）  
-> App 初始化容器（尚未启动）  
-> 绑定 Kernel 内核和异常处理  
-> Kernel 捕捉并处理请求  
-> 响应请求    

Application 初始化：  
初始化容器  
-> Events/Log/Router 注册基础服务  
-> 设置容器中 abstract 与 alias 的映射关系  
![Application 初始化流程](../image/app_start.png)  

Kernel 流程：  
初始化 Kernel（设置中间件）  
-> 捕获请求（路由调度，中间件栈，自定义异常处理）  
-> 处理请求（绑定 request），启动 Application（加载 .env 配置，加载 config 目录配置，设置错误和异常的 handler，设置 Facade 别名自动加载，注册服务提供者，启动服务提供者），使用管道和路由调度把请求通过中间件和路由  
-> 发送响应 -> Kernel 终止  

### Laravel 源码解读基础

**依赖注入和控制反转**  
依赖注入的目的在于解耦，通过把控制权交给外部达到控制的反转。  
如下，日志的实现可以有文件、数据库等等，记录用户登陆的日志信息的实现方式就可以通过构造函数的参数传递；当需要增加或切换日志实现方式时，不影响用户操作类，完美符合设计模式原则。  
```php
interface Log
{
    public function write($message);   
}

class FileLog implements Log
{
    public function write($message)
    {
        echo "Filelog::write:{$message}";
    }   
}

class DatabaseLog implements Log
{
    public function write($message)
    {
        echo "DatabaseLog::write:{$message}";
    }   
}

class User 
{
    protected $log;

    public function __construct(Log $log)
    {
        $this->log = $log;   
    }

    public function login()
    {
        // ... 
        $this->log->write("xxx Logined at " . date('Y-m-d H:i:s'));
    }
}

$user = new User(new DatabaseLog());
$user->login();
```

**反射**  
反射可以理解成根据类名返回该类的任何信息，比如该类有什么方法，参数，变量等等。  
```php
// 继续上面依赖注入的代码
//
// 获取 User 的反射对象
$reflector = new reflectionClass(User::class);
// 获取 User 的构造函数
$constructor = $reflector->getConstructor();
/*
ReflectionMethod {#234
  +name: "__construct"
  +class: "User"
  parameters: {
    $log: ReflectionParameter {#243
      +name: "log"
      position: 0
      typeHint: "Log"
    }
  }
  extra: {
    file: "当前文件路径"
    line: "28 to 31"
    isUserDefined: true
  }
  modifiers: "public"
}
 */
// 获取 User 构造函数的所有依赖参数
$dependencies = $constructor->getParameters();
/*
array:1 [
  0 => ReflectionParameter {#235
    +name: "log"
    position: 0
    typeHint: "Log"
  }
]
 */
// 创建 user 对象
// $reflector->newInstance($dependencies = []);
$user = $reflector->newInstance(new DatabaseLog());

// make 通过反射机制完成依赖注入
function make($concrete){
    $reflector = new ReflectionClass($concrete);
    $constructor = $reflector->getConstructor();
    if(null === $constructor) {
        return $reflector->newInstance();
    }

    $dependencies = $constructor->getParameters();
    $instances = getAllDependencies($dependencies);

    return $reflector->newInstanceArgs($instances);
}

function getAllDependencies($paramters) {
    $dependencies = [];
    foreach ($paramters as $paramter) {
        $dependencies[] = make($paramter->getClass()->name);
    }

    return $dependencies;
}

$user = make('User');
$user->login()
```

### Laravel Ios 容器


### Laravel 服务提供者


### Laravel 契约


### Laravel Facade


