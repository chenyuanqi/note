
### Yii2 生命周期
1、用户向入口脚本 web/index.php 发起请求  
2、入口脚本加载应用配置并创建一个应用实例去处理请求  
3、应用通过请求组件解析请求的路由  
4、应用创建一个控制器实例去处理请求  
5、控制器创建一个动作实例并针对操作执行过滤器  
6、如果任何一个过滤器返回失败，则动作取消  
7、如果所有过滤器都通过，动作将被执行  
8、动作会加载一个数据模型，或许是来自数据库  
9、动作会渲染一个视图，把数据模型提供给它  
10、渲染结果返回给响应组件  
11、响应组件发送渲染结果给用户浏览器  
![Yii2 生命周期](../image/yii-framework-flow.png)

**Yii2 一次请求的完整过程**  
从 yii\web\UrlManager 开始，最终回到 yii\web\Application。  

1、Yii2.0 框架使用了统一的入口脚本：index.php  
2、初始化配置项  
index.php 文件加载了各个文件夹下的配置项，再 new 一个 application（构造方法在 yii\web\Application 的父类 yii\base\Applicaton），主要根据配置项进行初始化。
3、创建实例 application 之后，调用 run 方法  
> before request 处理请求前的操作  
> handle request 真正的处理这次 HTTP 请求  
> after request 请求处理完成之后的操作  
> send response 将响应信息发送给客户端  

```php
$this->state = self::STATE_BEFORE_REQUEST;
$this->trigger(self::EVENT_BEFORE_REQUEST);

$this->state = self::STATE_HANDLING_REQUEST;
$response = $this->handleRequest($this->getRequest());

$this->state = self::STATE_AFTER_REQUEST;
$this->trigger(self::EVENT_AFTER_REQUEST);

$this->state = self::STATE_SENDING_RESPONSE;
$response->send();

$this->state = self::STATE_END;
```

4、请求的处理从 $this->handleRequest($this->getRequest()); 开始
```php
// getRequest 方法位于 yii\web\Application 中
public function getRequest()
{
    return $this->get('request');
}

// get 方法位于 yii\base\Application 的父类 yii\base\Module 的父类 yii\di\ServiceLocator
public function get($id, $throwException = true)
{
    if (isset($this->_components[$id])) {
        return $this->_components[$id];
    }

    if (isset($this->_definitions[$id])) {
        $definition = $this->_definitions[$id];
        if (is_object($definition) && !$definition instanceof Closure) {
            return $this->_components[$id] = $definition;
        } else {
            return $this->_components[$id] = Yii::createObject($definition);
        }
    } elseif ($throwException) {
        throw new InvalidConfigException("Unknown component ID: $id");
    } else {
        return null;
    }
}
```

5、run 中调用的 handleRequest() 位于 yii\web\Application 中  
> 从 Request 中获取用户请求路由  
> 调用这个路由对应的 action  

```php
// @param $request yii\web\Request
public function handleRequest($request)
{
    list ($route, $params) = $request->resolve();
    $this->requestedRoute = $route;
    //主要操作
    $result = $this->runAction($route, $params);
    if ($result instanceof Response) {
        return $result;
    } else {
        $response = $this->getResponse();
        if ($result !== null) {
            $response->data = $result;
        }
        return $response;
    }
}

public function resolve()
{
    $result = Yii::$app->getUrlManager()->parseRequest($this);
    if ($result !== false) {
    	// 匹配路由及参数
        list ($route, $params) = $result;
        if ($this->_queryParams === null) {
            $_GET = $params + $_GET; // preserve numeric keys
        } else {
            $this->_queryParams = $params + $this->_queryParams;
        }
        return [$route, $this->getQueryParams()];
    } else {
        throw new NotFoundHttpException(Yii::t('yii', 'Page not found.'));
    }
}

public function parseRequest($request)
{
    //在enable pretty url的前提下
    $pathInfo = $request->getPathInfo();
    //如果在rules中匹配到了 request 直接返回转换后的路由
    foreach ($this->rules as $rule) {
       if (($result = $rule->parseRequest($this, $request)) !== false) {
           return $result;
       }
    }
    //没有则使用默认方式，同时对路由进行检查，判断是否有多与一个的斜线
    //判断是否使用了.html后缀
    
    //代码省略，这里做的就是比较和截取的操作
    
    return [$pathInfo, []];
}
```

6、将 response 信息发送给客户端
```php
// yii\base\Application run()
$response = $this->handleRequest($this->getRequest());
$response->send();
return $response->exitStatus;
```

### Yii 源码解读基础
**__get 和 __set**  
魔术方法 \_\_get、\_\_set 是针对类而存在的。  
> \_\_get: 读取类不存在或者不可访问的属性时会被自动调用  
> \_\_set: 给类不存在或者不可访问的属性赋值时会被自动调用  

```php
class Configure
{
    public $name;

    private $_definitions = [];

    public function __get($name)
    {
        return isset($this->_definitions[$name]) ? $this->_definitions[$name] : null;
    }

    public function __set($name, $value)
    {
        $this->_definitions[$name] = $value;
    }
}

$config = [
    'class' => 'Test',
    'name' => 'vikey',
    'age' => 20,
];

$class = $config['class'];
unset($config['class']);
$object = new $class;

foreach ($config as $k => $v) {
    $object->$k = $v;
}

var_dump($object->name);
var_dump($object->age);
```

**控制反转和依赖注入、依赖倒置**  
“高内聚，低耦合”即类的内聚性是不是很高、耦合度是不是很低（简单来说，需要尽量让写出的程序易于维护，减少程序与程序之间的复杂性、耦合度），这一原则可以作为评判软件设计的好坏标准。  

控制反转（Inversion of Control，简称 IOC），字面上可以理解为对 xx 的控制进行了一个反转，即对 xx 的控制的另一种实现，是一种思路，一种逻辑思想。  
依赖注入（Dependency Injection，简称 DI），是 IOC 的一种典型实现。依赖注入就是把对象 A 所依赖的其他对象 B 或 C 或其他，以属性或者构造函数的方式传递到对象 A 的内部，而不是直接在对象 A 内实例化。其目的就是为了让对象 A 和其依赖的其他对象解耦，减少二者的依赖，即通过 “注入” 的方式去解决依赖问题。  

依赖倒置原则（Dependence Inversion Principle, DIP），是一种软件设计思想。  
传统软件设计中，上层代码依赖于下层代码，当下层出现变动时， 上层代码也要相应变化，维护成本较高。但是，DIP 的核心思想是上层定义接口，下层实现这个接口， 从而使得下层依赖于上层，降低耦合度，提高整个系统的弹性，这是一种经过实践证明的有效策略。  
```php
interface EmailSender
{
    public function send();
}

class EmailSenderByQq implements EmailSender
{
    public function send()
    {
    }
}

class EmailSenderBy163 implements EmailSender
{
    public function send()
    {
    }
}

class User
{
    public $emailSenderClass;

    public function __construct(EmailSender $emailSenderObject)
    {
        $this->emailSenderClass = $emailSenderObject;
    }

    public function register()
    {
        // other code
        $this->emailSenderClass->send();
    }
}

$user = new User(new EmailSenderBy163);
$user->register();
```

**反射**  
反射，是 php5 增加的功能。  
通过反射，可以提取出关于类、方法、属性、参数等的详细信息，即可以利用反射相关的一系列 API，来分析类所依赖的对象，并做自动实例化处理。
```php
// 简单版 Ioc 容器实现存储类及实例化类
Class Container
{
    private $_definitions;

    public function set($class, $definition)
    {
        $this->_definitions[$class] = $definition;
    }

    public function get($class, $params = [])
    {
        $definition = $this->_definitions[$class];
        return call_user_func($definition, $params);
    }
}

// 使用反射增强 Ioc 容器
Class Container
{
    public function get($class, $params = [])
    {
        return $this->build($class, $params);
    }

    public function build($class, $params)
    {
        $dependencies = [];

        $reflection = new ReflectionClass($class);
        $constructor = $reflection->getConstructor();
        if ($constructor !== null) {
            foreach ($constructor->getParameters() as $param) {
                $c = $param->getClass();
                if ($c !== null) {
                    $dependencies[] = $this->get($c->getName());
                }
            }
        }

        foreach ($params as $index => $param) {
            $dependencies[$index] = $param;
        }

        return $reflection->newInstanceArgs($dependencies);
    }
}
```


### Yii2 源码解读
yii2 的核心源码，位于 vendor/yiisoft/yii2 目录。

**入口文件**  
入口文件 web/index.php。按照 web server 的配置，入口文件是应用的开始，有且只有一个。  
入口文件加载 composer 的 autoload 文件，凡是通过 composer 安装的第三方库，都可以直接调用。
```php
// 定义当前应用是否处于 debug 模式，即调试模式
defined('YII_DEBUG') or define('YII_DEBUG', true);
// 定义当前应用的运行环境，比如开发环境，测试环境和生产环境等
defined('YII_ENV') or define('YII_ENV', 'dev');

// 加载 composer 的 autoload 文件
require(__DIR__ . '/../../vendor/autoload.php');

// 加载基础类文件 
require(__DIR__ . '/../../vendor/yiisoft/yii2/Yii.php');

// 在项目启动之前，加载一些全局的配置
require(__DIR__ . '/../../common/config/bootstrap.php');
require(__DIR__ . '/../config/bootstrap.php');

// 加载应用配置，并通过 yii\helpers\ArrayHelper::merge 方法合并处理
$config = yii\helpers\ArrayHelper::merge(
    require(__DIR__ . '/../../common/config/main.php'),
    require(__DIR__ . '/../../common/config/main-local.php'),
    require(__DIR__ . '/../config/main.php'),
    require(__DIR__ . '/../config/main-local.php')
);

// 实例化 yii\web\Application，并调用 run 方法运行
(new yii\web\Application($config))->run();
```

引入的 vendor/yiisoft/yii2/Yii.php，是一个继承 \yii\BaseYii 类的子类。  
Yii 类继承 \yii\BaseYii 之后，啥也没做，是个空类。尽管如此，在引用 BaseYii 的属性和方法时，最好还是用 Yii 来引用（比如用 Yii::getVersion() 获取当前 yii2 的版本号）。  
```php
require(__DIR__ . '/BaseYii.php');

class Yii extends \yii\BaseYii
{
}

spl_autoload_register(['Yii', 'autoload'], true, true);
// classes.php 保存 yii2 核心类和其类文件所属路径的映射关系（比如 'yii\base\Action' => YII2_PATH . '/base/Action.php'，其中 YII2_PATH 在 BaseYii.php 中的定义是 defined('YII2_PATH') or define('YII2_PATH', __DIR__);）
Yii::$classMap = require(__DIR__ . '/classes.php');
// 实例化一个全局的依赖注入容器
Yii::$container = new yii\di\Container();
```

**自动加载机制**  
上面 Yii.php 的 spl_autoload_register 函数定义了 Yii::autoload 为自动加载方法，即 yii\BaseYii::autoload 方法。这个方法是 yii2 内置的自动加载器，当 new 一个找不到的类的时候，这个方法就会被自动调用，这是 spl_autoload_register 函数的作用。  
```php
public static function autoload($className)
{
    if (isset(static::$classMap[$className])) {
    	// 加载核心类逻辑
        $classFile = static::$classMap[$className];
        if ($classFile[0] === '@') {
            $classFile = static::getAlias($classFile);
        }
    } elseif (strpos($className, '\\') !== false) {
    	// 加载非核心类逻辑
        $classFile = static::getAlias('@' . str_replace('\\', '/', $className) . '.php', false);
        if ($classFile === false || !is_file($classFile)) {
            return;
        }
    } else {
        return;
    }

    include($classFile);

    // ...
}
```

**别名 Alias**  
通过 Yii::setAlias 方法定义别名，定义的别名保存在 Yii::$aliases 属性上。  
通常，在入口文件中，require 的两个文件就引入了框架中一些的别名定义：  
```php
require(__DIR__ . '/../../common/config/bootstrap.php');
require(__DIR__ . '/../config/bootstrap.php');
```

**容器**  
关于 Container 创建对象的操作，我们可以使用 BaseYii::createObject 方法，该方法封装了 yii\di\Container 类的使用，所以通常直接用 Yii::createObject 方法创建对象或者调用可回调函数。  
```php
// Yii::createObject 方法的实现
// $type 是要创建的对象类型，$params 是创建的对象的构造参数
public static function createObject($type, array $params = [])
{
    if (is_string($type)) {
        return static::$container->get($type, $params);
    } elseif (is_array($type) && isset($type['class'])) {
        $class = $type['class'];
        unset($type['class']);
        return static::$container->get($class, $params, $type);
    } elseif (is_callable($type, true)) {
        return static::$container->invoke($type, $params);
    } elseif (is_array($type)) {
        throw new InvalidConfigException('Object configuration must be an array containing a "class" element.');
    }

    throw new InvalidConfigException('Unsupported configuration type: ' . gettype($type));
}
```
