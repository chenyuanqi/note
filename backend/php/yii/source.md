
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

// Container 中 set 方法
public function set($class, $definition = [], array $params = [])
{
    // 规范化类以及其定义，并以 $class 为下标，保存于 Container::_definitions 属性
    $this->_definitions[$class] = $this->normalizeDefinition($class, $definition);
    // 以 $class 为下标，将$class定义时的参数 $params 保存于 Container::_params 属性
    $this->_params[$class] = $params;
    // 重新定义的 $class，自然没必要保留单例的对象
    unset($this->_singletons[$class]);
    
    return $this;
}

// Container 中 get 方法
public function get($class, $params = [], $config = [])
{
    if (isset($this->_singletons[$class])) {
        // singleton
        return $this->_singletons[$class];
    } elseif (!isset($this->_definitions[$class])) {
        return $this->build($class, $params, $config);
    }

    $definition = $this->_definitions[$class];

    // 如果 $definition 是 callable，满足这个if
    if (is_callable($definition, true)) {
        // 调用 Container::resolveDependencies 方法解决依赖
        $params = $this->resolveDependencies($this->mergeParams($class, $params));
        // 调用 $definition
        $object = call_user_func($definition, $this, $params, $config);
    } 
    // 此时我们的$definition满足数组这一条件
    elseif (is_array($definition)) {
        $concrete = $definition['class'];
        unset($definition['class']);
        // 合并 $definition 中相同的 key，实质是以调用get方法时传递的 k-v 为准，k-v分别是对象的属性和属性值
        $config = array_merge($definition, $config);
        // $params是构造参数值，这里就是覆盖set时所定义的构造参数值，以当前get传递的构造参数值为准
        $params = $this->mergeParams($class, $params);
        // 以我们当前的例子，$class=testName,$concrete=frontend\components\Test,所以第一次走else
        // 如果我们当初直接定义的是类名 Yii::$container->set('frontend\components\Test'),此时就应该走下面这个if，然后直接build实例化了
        if ($concrete === $class) {
            $object = $this->build($class, $params, $config);
        } else {
            // 按照我们例子，走到这里之后，会继续调用 get 方法处理，即这是一个递归
            // 如果继续走下去，那么下一个get就会循环我们上文介绍的build，即从get方法的第4行代码走起
            // 不管怎么走，还是要build实例化 $class，如此，我们也就得到 $object 啦
            $object = $this->get($concrete, $params, $config);
        }
    } 
    // 如果 $definition 是对象，将其保存到 Container::_singletons 属性
    elseif (is_object($definition)) {
        return $this->_singletons[$class] = $definition;
    } 
    // 啥也不是，抛出异常
    else {
        throw new InvalidConfigException('Unexpected object definition type: ' . gettype($definition));
    }

    // 如果$class已经是单例了，则覆盖掉
    if (array_key_exists($class, $this->_singletons)) {
        // singleton
        $this->_singletons[$class] = $object;
    }
    // 返回最终的$object
    return $object;
}

// Container 中 build 方法
// 获取依赖，解析依赖，实例化依赖对象
protected function build($class, $params, $config)
{
    /* @var $reflection ReflectionClass */
    // 获取 $class 的依赖，返回 $class 的反射信息和依赖信息
    list ($reflection, $dependencies) = $this->getDependencies($class);

    // 依赖信息是通过构造方法获取的，所以如果我们通过 get 方法有给构造函数传递参数，理应以我们传递的为准，即覆盖掉 getDependencies 获取的
    foreach ($params as $index => $param) {
        $dependencies[$index] = $param;
    }

    // 解析依赖
    $dependencies = $this->resolveDependencies($dependencies, $reflection);
    if (!$reflection->isInstantiable()) {
        throw new NotInstantiableException($reflection->name);
    }
    // 如果$config为空，即不需要给 $class "注入" 属性和属性值，直接实例化
    if (empty($config)) {
        return $reflection->newInstanceArgs($dependencies);
    }

    // 检查 $class 是否实现了接口 yii\base\Configurable，我们的Test肯定是否了
    if (!empty($dependencies) && $reflection->implementsInterface('yii\base\Configurable')) {
        // set $config as the last parameter (existing one will be overwritten)
        $dependencies[count($dependencies) - 1] = $config;
        return $reflection->newInstanceArgs($dependencies);
    } else {
        // 实例化 $class，并"注入"属性，最后返回我们的对象 $object
        $object = $reflection->newInstanceArgs($dependencies);
        foreach ($config as $name => $value) {
            $object->$name = $value;
        }
        return $object;
    }
}

// 获取依赖
protected function getDependencies($class)
{
    // ①、判断 Container::_reflections 属性是否有保存 $class 的反射信息，如果有就直接返回，不用再解析，毕竟获取类的反射信息也是要消耗时间
    if (isset($this->_reflections[$class])) {
        return [$this->_reflections[$class], $this->_dependencies[$class]];
    }

    $dependencies = [];
    $reflection = new ReflectionClass($class);

    $constructor = $reflection->getConstructor();
    if ($constructor !== null) {
        foreach ($constructor->getParameters() as $param) {
            // ②、获取构造函数的参数时，通过 isDefaultValueAvailable 方法判断参数的默认值
            if ($param->isDefaultValueAvailable()) {
                $dependencies[] = $param->getDefaultValue();
            } else {
                $c = $param->getClass();
                // ③、调用 yii\di\Instance::of 方法处理依赖的类名，Instance::of 方法返回 Instace 类的实例
                $dependencies[] = Instance::of($c === null ? null : $c->getName());
            }
        }
    }
    // ④、将反射信息和依赖信息分别保存在 Container::_reflections和Container::_dependencies属性，防止重复实例化该类时重复解析反射信息和依赖信息
    $this->_reflections[$class] = $reflection;
    $this->_dependencies[$class] = $dependencies;

    return [$reflection, $dependencies];
}

// 解析依赖
protected function resolveDependencies($dependencies, $reflection = null)
{
    foreach ($dependencies as $index => $dependency) {
        if ($dependency instanceof Instance) {
            if ($dependency->id !== null) {
                $dependencies[$index] = $this->get($dependency->id);
            } elseif ($reflection !== null) {
                $name = $reflection->getConstructor()->getParameters()[$index]->getName();
                $class = $reflection->getName();
                throw new InvalidConfigException("Missing required parameter \"$name\" when instantiating \"$class\".");
            }
        }
    }
    return $dependencies;
}
```

**应用的生命周期之预初始化**  
应用程序类 yii\web\Application 的继承关系：  
yii\web\Application => yii\base\Application => yii\base\Module => yii\di\ServiceLocator（服务定位器） => yii\base\Component => yii\base\Object => yii\base\Configurable  
```php
// 执行 yii\base\Application::run 方法之前， yii\base\Application 的构造方法__construct 会先被执行
public function __construct($config = [])
{
    // yii\web\Application 对象被赋值给了 Yii::$app 静态属性
    Yii::$app = $this;
    // 调用 yii\base\Module::setInstance 方法，把当前 module 的实例保存到 yii\base\Application 的 loadedModules 属性
    static::setInstance($this);

    // 标记当前应用目前处于生命周期的哪个状态（共有 7 个状态）
    $this->state = self::STATE_BEGIN;

    // 预初始化配置
    $this->preInit($config);

    // 把 errorHandler 组件注册为应用的错误处理程序
    $this->registerErrorHandler($config);

    // 调用 yii\BaseYii::configure 方法，初始化 yii\web\Application 类的属性
    // 调用 yii\web\Application::init 方法完成初始化操作
    Component::__construct($config);
}

// 预初始化方法的目的：
// 1、检测配置是否正常，包括 id,basePath 必须配置项
// 2、设置 module 的 root 目录，并伴随一系列的别名设置，包括 @app、@vendor、@bower、@npm 、@runtime
// 3、时区的设置
// 4、合并核心 component 和自定义的 component
public function preInit(&$config)
{
    // 必须存在下标为id的单元，否则throw抛出异常
    if (!isset($config['id'])) {
        throw new InvalidConfigException('The "id" configuration for the Application is required.');
    }
    // 1、basePath 也是要必须设置的，其含义指的是当前module的root目录，不设置则抛出异常。我们看到basePath在 main.php中的值是 dirname(__DIR__)，即 frontend 目录
    // 2、调用 yii\base\Application::setBasePath 方法处理，会先调用 yii\base\Module::setBasePath 方法处理，yii\base\Module::setBasePath方法会获取 basePath的值的绝对路径，保存在 yii\base\Module::_basePath属性，表示当前module的root目录
    // 3、yii\base\Application::setBasePath 跟着会通过 Yii::setAlias('@app', $this->getBasePath()) 设置别名 @app 指向 yii\base\Module::_basePath 所指目录，不信你可以在其他位置打印下 Yii::getAlias('@app') 看看结果
    if (isset($config['basePath'])) {
        $this->setBasePath($config['basePath']);
        unset($config['basePath']);
    } else {
        throw new InvalidConfigException('The "basePath" configuration for the Application is required.');
    }

    // vendorPath，默认在common/config/main.php内配置的值是 dirname(dirname(__DIR__)) . '/vendor'
    // 调用 yii\base\Application::setVendorPath 方法处理，其目的是设置三个别名，@vendor、@bower、@npm 你可以在 yii\base\Application::setVendorPath 方法内看到，也就是说composer的安装包的位置，我们也是可以通过配置指定的
    if (isset($config['vendorPath'])) {
        $this->setVendorPath($config['vendorPath']);
        unset($config['vendorPath']);
    } else {
        // set "@vendor"
        // 如果没有指定 vendorPath,默认vendorPath是 yii\base\Module::_basePath 上一级目录下的vendor目录
        $this->getVendorPath();
    }

    // 如果配置了 runtimePath，即程序运行时的存储路径，比如日志路径，debug信息路径等都可以通过 runtimePath配置
    // 同时，yii\base\Application::setRuntimePath 方法会设置 @runtime 别名指向该路径
    if (isset($config['runtimePath'])) {
        $this->setRuntimePath($config['runtimePath']);
        unset($config['runtimePath']);
    } else {
        // set "@runtime"
        // 如果没有配置runtimePath，默认的runtimePath是 yii\base\Module::_basePath 目录下的runtime目录，即@runtime默认指向这里
        $this->getRuntimePath();
    }

    // 设置时区,我们尚未配置该选项，如果指定了，调用 yii\base\Application::setTimeZone会调用 date_default_timezone_set函数设置时区，否则将会以 php.ini内指定的时区为准，如果php.ini也未配置，默认时区格式是 UTC，也就是部分同学什么都没有配置时发现程序内的时间总是相差8小时的缘故
    if (isset($config['timeZone'])) {
        $this->setTimeZone($config['timeZone']);
        unset($config['timeZone']);
    } elseif (!ini_get('date.timezone')) {
        $this->setTimeZone('UTC');
    }

    // 设置自己的container，当然，大多数情况下是没有必要的，默认的container是yii\di\Container
    if (isset($config['container'])) {
        $this->setContainer($config['container']);

        unset($config['container']);
    }

    // merge core components with custom components
    // 合并核心component和自定义的component
    // 核心组件，在 yii\web\Application::coreComponents 和 yii\base\Application::coreComponents 方法内均有配置
    // 自定义的component指的是 $config['components'] 内的配置
    // 如果自定义的component跟核心component不冲突，则把核心component追加到 $config['components']；如果二者冲突且自定义的component结构规范，则以自定义的为准，规范指的是 是个数组且包含 class 下标
    foreach ($this->coreComponents() as $id => $component) {
        if (!isset($config['components'][$id])) {
            $config['components'][$id] = $component;
        } elseif (is_array($config['components'][$id]) && !isset($config['components'][$id]['class'])) {
            $config['components'][$id]['class'] = $component['class'];
        }
    }
}
```

**事件**  
事件分实例事件（针对某个实例进行的事件操作）、类级别事件（如 yii\web\User 类中 beforeLogin 和 afterLogin）、全局事件（Yii::$app 的实例事件）。  
事件的实现原理：先通过一个方法对某个事件注册相关的回调函数，触发事件的时候，再通过事件 ID 找到注册的事件进行回调。  
```php
// 绑定事件：存储到 yii\base\Component::_events[$name]
public function on($name, $handler, $data = null, $append = true)
{
    // 确保行为已被渲染到 $_behaviors 中
    $this->ensureBehaviors();
    if ($append || empty($this->_events[$name])) {
        $this->_events[$name][] = [$handler, $data];
    } else {
        array_unshift($this->_events[$name], [$handler, $data]);
    }
}

// 移除事件
public function off($name, $handler = null)
{
    $this->ensureBehaviors();
    if (empty($this->_events[$name]) && empty($this->_eventWildcards[$name])) {
        return false;
    }
    if ($handler === null) {
        unset($this->_events[$name], $this->_eventWildcards[$name]);
        return true;
    }

    $removed = false;
    // plain event names
    if (isset($this->_events[$name])) {
        foreach ($this->_events[$name] as $i => $event) {
            if ($event[0] === $handler) {
                unset($this->_events[$name][$i]);
                $removed = true;
            }
        }
        if ($removed) {
            $this->_events[$name] = array_values($this->_events[$name]);
            return $removed;
        }
    }

    // wildcard event names
    if (isset($this->_eventWildcards[$name])) {
        foreach ($this->_eventWildcards[$name] as $i => $event) {
            if ($event[0] === $handler) {
                unset($this->_eventWildcards[$name][$i]);
                $removed = true;
            }
        }
        if ($removed) {
            $this->_eventWildcards[$name] = array_values($this->_eventWildcards[$name]);
            // remove empty wildcards to save future redundant regex checks:
            if (empty($this->_eventWildcards[$name])) {
                unset($this->_eventWildcards[$name]);
            }
        }
    }

    return $removed;
}

// 触发事件
public function trigger($name, Event $event = null)
{
    $this->ensureBehaviors();
    if (!empty($this->_events[$name])) {
        if ($event === null) {
            $event = new Event;
        }
        if ($event->sender === null) {
            $event->sender = $this;
        }
        $event->handled = false;
        $event->name = $name;
        foreach ($this->_events[$name] as $handler) {
            // 存储数据
            $event->data = $handler[1];
            // 把 $event 传递给事件回调函数
            call_user_func($handler[0], $event);
            // 如果 trigger 时设置的 event 实例的 handled 属性为真，则该事件后续未调用的其他回调将不会被触发
            if ($event->handled) {
                return;
            }
        }
    }

    // 传递给类级别的事件触发
    Event::trigger($this, $name, $event);
}
```

**行为**  
行为是 yii\base\Behavior 或其子类的实例。  
行为，也称为 mixins， 可以无须改变类继承关系即可增强一个已有的 yii\base\Component 类功能。当行为附加到组件后，它将“注入”它的方法和属性到组件，然后可以像访问组件内定义的方法和属性一样访问它们。此外，行为通过组件能响应被触发的事件，从而自定义或调整组件正常执行的代码。  
实际上，行为就是一个类，通过某些特殊的方式（注入、绑定等），跟其他类进行了绑定，二者从而可以进行交互。行为，就是对当前类进行一个扩展，而且不用修改当前类。  
```php
// 行为绑定
public function attachBehavior($name, $behavior)
{
    // 确保行为已被渲染到 $_behaviors 中
    $this->ensureBehaviors();

    return $this->attachBehaviorInternal($name, $behavior);
}

// 绑定行为到组件内部
private function attachBehaviorInternal($name, $behavior)
{
    if (!($behavior instanceof Behavior)) {
        $behavior = Yii::createObject($behavior);
    }
    if (is_int($name)) {
        // 关联 component
        $behavior->attach($this);
        // 把行为添加到 yii\base\Component::_behaviors 属性
        $this->_behaviors[] = $behavior;
    } else {
        if (isset($this->_behaviors[$name])) {
            $this->_behaviors[$name]->detach();
        }
        $behavior->attach($this);
        $this->_behaviors[$name] = $behavior;
    }

    return $behavior;
}

public function __call($name, $params)
{
    $this->ensureBehaviors();
    foreach ($this->_behaviors as $object) {
        if ($object->hasMethod($name)) {
            return call_user_func_array([$object, $name], $params);
        }
    }
    throw new UnknownMethodException('Calling unknown method: ' . get_class($this) . "::$name()");
}
```
为了让 component 更强大，再看看行为事件。  
典型案例就是 yii\behaviors\TimestampBehavior 的使用，流程大致是这样的：  
yii\behaviors\TimestampBehavior::init =>  
yii\base\Behavior::attach =>  
yii\behaviors\AttributeBehavior::events =>  
yii\db\BaseActiveRecord::beforeSave =>  
yii\behaviors\AttributeBehavior::evaluateAttributes。    
```php
// AR 类中
public function behaviors()
{
    return [
        [
            'class' => TimestampBehavior::className(),
            'createdAtAttribute' => 'created_at',
            'updatedAtAttribute' => 'updated_at',
            'value' => date('Y-m-d H:i:s'),
        ],
    ];
}

// yii\behaviors\TimestampBehavior::init
if (empty($this->attributes)) {
    $this->attributes = [
        BaseActiveRecord::EVENT_BEFORE_INSERT => [$this->createdAtAttribute, $this->updatedAtAttribute],
        BaseActiveRecord::EVENT_BEFORE_UPDATE => $this->updatedAtAttribute,
    ];
}

// yii\behaviors\AttributeBehavior::events
return array_fill_keys(
    array_keys($this->attributes),
    'evaluateAttributes'
);

// yii\base\Behavior::attach
foreach ($this->events() as $event => $handler) {
    $owner->on($event, is_string($handler) ? [$this, $handler] : $handler);
}

// 计算属性
public function evaluateAttributes($event)
{
    if ($this->skipUpdateOnClean
        && $event->name == ActiveRecord::EVENT_BEFORE_UPDATE
        && empty($this->owner->dirtyAttributes)
    ) {
        return;
    }

    if (!empty($this->attributes[$event->name])) {
        $attributes = (array) $this->attributes[$event->name];
        $value = $this->getValue($event);
        foreach ($attributes as $attribute) {
            // ignore attribute names which are not string (e.g. when set by TimestampBehavior::updatedAtAttribute)
            if (is_string($attribute)) {
                if ($this->preserveNonEmptyValues && !empty($this->owner->$attribute)) {
                    continue;
                }
                // 为属性赋值
                $this->owner->$attribute = $value;
            }
        }
    }
}
```

**应用的生命周期之初始化**  
应用的初始化从 yii\base\Object::__construct 构造方法中调用 $this->init()，这里的 $this 就是 yii\web\Application 类的实例 Yii::$app。  
```php
public function init()
{
    // 标识状态为初始化
    $this->state = self::STATE_INIT;
    $this->bootstrap();
}

protected function bootstrap()
{
    // 获取请求组件
    $request = $this->getRequest();
    Yii::setAlias('@webroot', dirname($request->getScriptFile()));
    Yii::setAlias('@web', $request->getBaseUrl());

    parent::bootstrap();
}

// 父级 yii\base\Application::bootstrap
protected function bootstrap()
{
    // yii\web\Application 的 extensions 属性
    // 默认指的是 @vendor/yiisoft/extensions.php 内返回的数组
    if ($this->extensions === null) {
        $file = Yii::getAlias('@vendor/yiisoft/extensions.php');
        $this->extensions = is_file($file) ? include($file) : [];
    }

    // 循环处理 extensions,如果扩展存在 alias 项，则设置别名
    // 如果扩展存在 bootstrap 项，且经过 Container 容器实例化之后，bootstrap 项是 BootstrapInterface 的实例，调用该实例的 bootstrap 方法执行一些扩展的启动操作
    foreach ($this->extensions as $extension) {
        if (!empty($extension['alias'])) {
            foreach ($extension['alias'] as $name => $path) {
                Yii::setAlias($name, $path);
            }
        }

        if (isset($extension['bootstrap'])) {
            $component = Yii::createObject($extension['bootstrap']);
            if ($component instanceof BootstrapInterface) {
                Yii::trace('Bootstrap with ' . get_class($component) . '::bootstrap()', __METHOD__);
                $component->bootstrap($this);
            } else {
                Yii::trace('Bootstrap with ' . get_class($component), __METHOD__);
            }
        }
    }

    // bootstrap 属性，默认在 frontend/config/main.php 和 frontend/config/main-local.php 文件内都有配置，是个数组，包括 log、gii、debug 三个元素，这里循环处理
    foreach ($this->bootstrap as $class) {
        $component = null;
        // 这里判断 bootstrap 项是否是字符串，如果是且已经在 yii\di\ServiceLocator 中注册过，则直接获取
        // 如果是 yii\base\Module 中注册过的 module，则直接获取，否则抛出异常
        if (is_string($class)) {
            if ($this->has($class)) {
                $component = $this->get($class);
            } elseif ($this->hasModule($class)) {
                $component = $this->getModule($class);
            } elseif (strpos($class, '\\') === false) {
                throw new InvalidConfigException("Unknown bootstrapping component ID: $class");
            }
        }

        if (!isset($component)) {
            $component = Yii::createObject($class);
        }

        // 如果经过以上获取到的 $component 是 yii\base\BootstrapInterface 接口，则 trace 记录之后，调用 component 的 bootstrap 方法启动 component
        if ($component instanceof BootstrapInterface) {
            Yii::trace('Bootstrap with ' . get_class($component) . '::bootstrap()', __METHOD__);
            $component->bootstrap($this);
        } else {
            Yii::trace('Bootstrap with ' . get_class($component), __METHOD__);
        }
    }
}
```

**应用的生命周期之执行请求**  
应用的运行，从 yii\base\Application::run 开始。  
```php
public function run()
{
    try {
        $this->state = self::STATE_BEFORE_REQUEST;
        // 事件触发，事件的绑定在于预定义事件的时候
        $this->trigger(self::EVENT_BEFORE_REQUEST);

        $this->state = self::STATE_HANDLING_REQUEST;
        // 处理请求：捕获路由，处理路由以及根据路由规则调用对应的方法
        $response = $this->handleRequest($this->getRequest());

        $this->state = self::STATE_AFTER_REQUEST;
        $this->trigger(self::EVENT_AFTER_REQUEST);

        $this->state = self::STATE_SENDING_RESPONSE;
        $response->send();

        $this->state = self::STATE_END;

        return $response->exitStatus;
    } catch (ExitException $e) {
        $this->end($e->statusCode, isset($response) ? $response : null);
        return $e->statusCode;
    }
}
```

**请求**  
request 组件（yii\web\Request）是用来处理应用的请求。
```php
public function get($name = null, $defaultValue = null)
{
    if ($name === null) {
        return $this->getQueryParams();
    }
    return $this->getQueryParam($name, $defaultValue);
}

public function getQueryParams()
{
    if ($this->_queryParams === null) {
        return $_GET;
    }
    return $this->_queryParams;
}

public function getQueryParam($name, $defaultValue = null)
{
    $params = $this->getQueryParams();
    return isset($params[$name]) ? $params[$name] : $defaultValue;
}

public function getBodyParams()
{
    // 判断请求体参数是否是 null，可以通过 yii\web\Request::setBodyParams 方法修改并非实际意义的请求体参数集合
    if ($this->_bodyParams === null) {
        // 可以在 post 请求中添加 yii\web\Request::methodParam 对应的 key：_method，来修改当前的请求方法，但是请求体参数还是 $_POST
        if (isset($_POST[$this->methodParam])) {
            $this->_bodyParams = $_POST;
            unset($this->_bodyParams[$this->methodParam]);
            return $this->_bodyParams;
        }
        // 获取请求头的 content_type，不同的请求可以设置不同的 content_type，同时这也是解析的关键
        // 比如说客户端可以设置请求头 content_type 等于 'application/json; charset=UTF-8',也可以让其等于 'application/json',具体以分号";"前面的为准，紧接着下面的 if else 是这句话的实现
        $rawContentType = $this->getContentType();
        if (($pos = strpos($rawContentType, ';')) !== false) {
            // e.g. application/json; charset=UTF-8
            $contentType = substr($rawContentType, 0, $pos);
        } else {
            $contentType = $rawContentType;
        }
        // yii\web\Request::parsers 属性，我们可以在配置 request 组件的时候，配置该属性
        // 这个属性的值，是 content_type 和解析方法的映射，比如我们可以设定凡是 content_type=application/json 的请求，用 yii\web\JsonParser 解析。我们也可以自定义 content_type, 让解析的类继承 RequestParserInterface 接口即可
        if (isset($this->parsers[$contentType])) {
            $parser = Yii::createObject($this->parsers[$contentType]);
            if (!($parser instanceof RequestParserInterface)) {
                throw new InvalidConfigException("The '$contentType' request parser is invalid. It must implement the yii\\web\\RequestParserInterface.");
            }
            $this->_bodyParams = $parser->parse($this->getRawBody(), $rawContentType);
        } elseif (isset($this->parsers['*'])) {
            $parser = Yii::createObject($this->parsers['*']);
            if (!($parser instanceof RequestParserInterface)) {
                throw new InvalidConfigException("The fallback request parser is invalid. It must implement the yii\\web\\RequestParserInterface.");
            }
            $this->_bodyParams = $parser->parse($this->getRawBody(), $rawContentType);
        } 
        // 如果没有定义parsers, 这里会以POST为准
        elseif ($this->getMethod() === 'POST') {
            // PHP has already parsed the body so we have all params in $_POST
            $this->_bodyParams = $_POST;
        } 
        // 以上都不符合的情况下，则使用 mb_parse_str 函数解析
        else {
            $this->_bodyParams = [];
            mb_parse_str($this->getRawBody(), $this->_bodyParams);
        }
    }
    return $this->_bodyParams;
}
```

**响应**  
response 组件（yii\web\Response）是用来处理应用的响应。  
可以使用 Yii::$app->get ('response') 或者 Yii::$app->getResponse () 获取 Response 组件。  
响应从 yii\web\Application::handleRequest 方法的 $this->runAction 开始。  
```php
// yii\web\Application::runAction
public function runAction($route, $params = [])
{
    $parts = $this->createController($route);
    if (is_array($parts)) {
        /* @var $controller Controller */
        list($controller, $actionID) = $parts;
        $oldController = Yii::$app->controller;
        Yii::$app->controller = $controller;
        $result = $controller->runAction($actionID, $params);
        if ($oldController !== null) {
            Yii::$app->controller = $oldController;
        }

        return $result;
    }

    $id = $this->getUniqueId();
    throw new InvalidRouteException('Unable to resolve the request "' . ($id === '' ? $route : $id . '/' . $route) . '".');
}

// controller 的 runAction
public function runAction($id, $params = [])
{
    $action = $this->createAction($id);
    if ($action === null) {
        throw new InvalidRouteException('Unable to resolve the request: ' . $this->getUniqueId() . '/' . $id);
    }

    Yii::debug('Route to run: ' . $action->getUniqueId(), __METHOD__);

    if (Yii::$app->requestedAction === null) {
        Yii::$app->requestedAction = $action;
    }

    $oldAction = $this->action;
    $this->action = $action;

    $modules = [];
    $runAction = true;

    // call beforeAction on modules
    foreach ($this->getModules() as $module) {
        if ($module->beforeAction($action)) {
            array_unshift($modules, $module);
        } else {
            $runAction = false;
            break;
        }
    }

    $result = null;

    if ($runAction && $this->beforeAction($action)) {
        // run the action
        $result = $action->runWithParams($params);

        $result = $this->afterAction($action, $result);

        // call afterAction on modules
        foreach ($modules as $module) {
            /* @var $module Module */
            $result = $module->afterAction($action, $result);
        }
    }

    if ($oldAction !== null) {
        $this->action = $oldAction;
    }

    return $result;
}

// yii\web\Response::send 
public function send()
{   
    if ($this->isSent) {
        return;
    }
    $this->trigger(self::EVENT_BEFORE_SEND);
    $this->prepare();
    $this->trigger(self::EVENT_AFTER_PREPARE);
    $this->sendHeaders();
    $this->sendContent();
    $this->trigger(self::EVENT_AFTER_SEND);
    // 避免重复发送
    $this->isSent = true;
}
```

**cookie**  
cookie 主要涉及以下两部分：服务端设置 cookie，并向客户端发送 cookie；客户端获取 cookie。  

用 Response 组件描述服务端对 cookie 的设置，用 Request 组件描述客户端对 cookie 的获取等操作。  
实际上，Response 组件并没有直接对 cookie 进行操作，cookie 的操作主要由 yii\web\CookieCollection 负责。Response 组件的作用非常关键，虽然 yii\web\CookieCollection 负责管理 cookie，比如添加 / 删除等。但对服务端而言，通过 setcookie 方法把 cookie 发送给客户端才有意义，Response 的作用即在于此。  
通常，可以直接用 $response->cookies （$reponse 是 yii\web\Response 的实例）去描述 yii\web\CookieCollection 实例（间接的调用 $response->getCookies() 方法）。  
```php
class CookieCollection extends Object implements \IteratorAggregate, \ArrayAccess, \Countable
{
    public function __construct($cookies = [], $config = [])
    {
        $this->_cookies = $cookies;
        parent::__construct($config);
    }

    public function getIterator()
    {
        return new ArrayIterator($this->_cookies);
    }

    public function getValue($name, $defaultValue = null)
    {
        return isset($this->_cookies[$name]) ? $this->_cookies[$name]->value : $defaultValue;
    }

    public function has($name)
    {
        // 存在且值不为空，过期时间存在且大于当前时间
        return isset($this->_cookies[$name]) && $this->_cookies[$name]->value !== ''
            && ($this->_cookies[$name]->expire === null || $this->_cookies[$name]->expire === 0 || $this->_cookies[$name]->expire >= time());
    }

    public function add($cookie)
    {
        if ($this->readOnly) {
            throw new InvalidCallException('The cookie collection is read only.');
        }
        $this->_cookies[$cookie->name] = $cookie;
    }

    public function remove($cookie, $removeFromBrowser = true)
    {
        if ($this->readOnly) {
            throw new InvalidCallException('The cookie collection is read only.');
        }

        // 设置过期时间为 1s
        if ($cookie instanceof Cookie) {
            $cookie->expire = 1;
            $cookie->value = '';
        } else {
            $cookie = Yii::createObject([
                'class' => 'yii\web\Cookie',
                'name' => $cookie,
                'expire' => 1,
            ]);
        }

        // 浏览器设置cookie，其他直接移除
        if ($removeFromBrowser) {
            $this->_cookies[$cookie->name] = $cookie;
        } else {
            unset($this->_cookies[$cookie->name]);
        }
    }

    public function removeAll()
    {
        if ($this->readOnly) {
            throw new InvalidCallException('The cookie collection is read only.');
        }
        $this->_cookies = [];
    }

}

```

**session**  
session 组件 yii\web\Session 的定义跟 yii\web\CookieCollection 类似，但是父类不同。  
session 组件的实现主要是针对 $_SESSION 的操作，这一点跟原生 php 保持一致。  
```php
class Session extends Component implements \IteratorAggregate, \ArrayAccess, \Countable
{
    public function getIterator()
    {
        $this->open();
        return new SessionIterator;
    }

    public function open()
    {
        // 判断当前 session 的状态
        if ($this->getIsActive()) {
            return;
        }

        // 对 session_set_save_handler 方法的封装
        $this->registerSessionHandler();
        // 设置 cookie 参数
        $this->setCookieParamsInternal();

        YII_DEBUG ? session_start() : @session_start();

        // 再次判断当前 session 的状态
        if ($this->getIsActive()) {
            Yii::info('Session started', __METHOD__);
            // 更新 flash 数据
            $this->updateFlashCounters();
        } else {
            // 记录异常信息
            $error = error_get_last();
            $message = isset($error['message']) ? $error['message'] : 'Failed to start session.';
            Yii::error($message, __METHOD__);
        }
    }

    public function getIsActive()
    {
        return session_status() === PHP_SESSION_ACTIVE;
    }

    protected function registerSessionHandler()
    {
        // handler 属性是一个继承 \SessionHandlerInterface 的 object
        if ($this->handler !== null) {
            if (!is_object($this->handler)) {
                $this->handler = Yii::createObject($this->handler);
            }
            if (!$this->handler instanceof \SessionHandlerInterface) {
                throw new InvalidConfigException('"' . get_class($this) . '::handler" must implement the SessionHandlerInterface.');
            }
            YII_DEBUG ? session_set_save_handler($this->handler, false) : @session_set_save_handler($this->handler, false);
        } elseif ($this->getUseCustomStorage()) {
            // 如果让 yii\web\Session::getUseCustomStorage 方法 return true，会覆盖 yii\web\Session 中的方法（session 保存到 redis、mysql 等介质时）
            if (YII_DEBUG) {
                session_set_save_handler(
                    [$this, 'openSession'],
                    [$this, 'closeSession'],
                    [$this, 'readSession'],
                    [$this, 'writeSession'],
                    [$this, 'destroySession'],
                    [$this, 'gcSession']
                );
            } else {
                @session_set_save_handler(
                    [$this, 'openSession'],
                    [$this, 'closeSession'],
                    [$this, 'readSession'],
                    [$this, 'writeSession'],
                    [$this, 'destroySession'],
                    [$this, 'gcSession']
                );
            }
        }
    }

    private function setCookieParamsInternal()
    {
        $data = $this->getCookieParams();
        if (isset($data['lifetime'], $data['path'], $data['domain'], $data['secure'], $data['httponly'])) {
            if (PHP_VERSION_ID >= 70300) {
                session_set_cookie_params($data);
            } else {
                if (!empty($data['sameSite'])) {
                    throw new InvalidConfigException('sameSite cookie is not supported by PHP versions < 7.3.0 (set it to null in this environment)');
                }
                session_set_cookie_params($data['lifetime'], $data['path'], $data['domain'], $data['secure'], $data['httponly']);
            }

        } else {
            throw new InvalidArgumentException('Please make sure cookieParams contains these elements: lifetime, path, domain, secure and httponly.');
        }
    }

    public function get($key, $defaultValue = null)
    {
        $this->open();
        return isset($_SESSION[$key]) ? $_SESSION[$key] : $defaultValue;
    }

    public function set($key, $value)
    {
        $this->open();
        $_SESSION[$key] = $value;
    }

    public function destroy()
    {
        if ($this->getIsActive()) {
            $sessionId = session_id();
            $this->close();
            $this->setId($sessionId);
            $this->open();
            session_unset();
            session_destroy();
            $this->setId($sessionId);
        }
    }
}
```

**User 组件**  
User 组件指的是 yii\web\User。  
User 组件的作用是为应用提供用户认证状态的管理，即和 cookie、session 交互。而真正和用户信息相关的是 identityClass 指向的类，比如用户名、密码等。
```php
// user 组件的默认配置
[
    'components' => [
        'user' => [
            'identityClass' => 'common\models\User',
            'enableAutoLogin' => true,
            'identityCookie' => ['name' => '_identity-backend', 'httpOnly' => true],
        ],
    ]
]
```
User 组件实现登陆登出。
```php
protected function beforeLogin($identity, $cookieBased, $duration)
{
    $event = new UserEvent([
        'identity' => $identity,
        'cookieBased' => $cookieBased,
        'duration' => $duration,
    ]);
    $this->trigger(self::EVENT_BEFORE_LOGIN, $event);

    return $event->isValid;
}

public function login(IdentityInterface $identity, $duration = 0)
{
    if ($this->beforeLogin($identity, false, $duration)) {
        // 切换用户身份
        $this->switchIdentity($identity, $duration);
        $id = $identity->getId();
        $ip = Yii::$app->getRequest()->getUserIP();
        if ($this->enableSession) {
            $log = "User '$id' logged in from $ip with duration $duration.";
        } else {
            $log = "User '$id' logged in from $ip. Session not enabled.";
        }

        $this->regenerateCsrfToken();

        Yii::info($log, __METHOD__);
        $this->afterLogin($identity, false, $duration);
    }

    return !$this->getIsGuest();
}

public function switchIdentity($identity, $duration = 0)
{
    // 标识用户的状态信息
    $this->setIdentity($identity);

    if (!$this->enableSession) {
        return;
    }

    /* Ensure any existing identity cookies are removed. */
    if ($this->enableAutoLogin && ($this->autoRenewCookie || $identity === null)) {
        $this->removeIdentityCookie();
    }

    $session = Yii::$app->getSession();
    if (!YII_ENV_TEST) {
        $session->regenerateID(true);
    }
    $session->remove($this->idParam);
    $session->remove($this->authTimeoutParam);

    if ($identity) {
        // 存储 id  信息
        $session->set($this->idParam, $identity->getId());
        // authTimeout 如果用户不活动了，将会在该秒数之后自动退出
        if ($this->authTimeout !== null) {
            // 存储认证时间
            $session->set($this->authTimeoutParam, time() + $this->authTimeout);
        }
        // absoluteAuthTimeout 不管用户活不活动，时间到了就退出
        if ($this->absoluteAuthTimeout !== null) {
            // 存储绝对认证时间
            $session->set($this->absoluteAuthTimeoutParam, time() + $this->absoluteAuthTimeout);
        }
        // 自动登陆，注意 enableAutoLogin 是 User 组件的属性而不是表单值
        if ($this->enableAutoLogin && $duration > 0) {
            // 记录 cookie：用户标识、auth_key 和有效时间
            $this->sendIdentityCookie($identity, $duration);
        }
    }
}

public function logout($destroySession = true)
{
    $identity = $this->getIdentity();
    if ($identity !== null && $this->beforeLogout($identity)) {
        $this->switchIdentity(null);
        $id = $identity->getId();
        $ip = Yii::$app->getRequest()->getUserIP();
        Yii::info("User '$id' logged out from $ip.", __METHOD__);
        if ($destroySession && $this->enableSession) {
            Yii::$app->getSession()->destroy();
        }
        $this->afterLogout($identity);
    }

    return $this->getIsGuest();
}
```

**view 组件**  
view 组件用于渲染视图页面，他们的类之间的关系如下：  
```
frontend\controllers\SiteController  
=> yii\web\Controller  
=> yii\base\Controller  
=> yii\base\Component  
=> yii\web\View  
=> yii\base\View  
=> yii\base\Component  
```
view 组件获取视图文件内容的字符串过程：  
```php
// yii\base\Controller::render
public function render($view, $params = [])
{
    // getView 获取 view 组件的实例 
    // 跳转 yii\web\View::render，实际是 yii\base\View::render
    $content = $this->getView()->render($view, $params, $this);
    return $this->renderContent($content);
}

// yii\base\View::render
public function render($view, $params = [], $context = null)
{
    // 寻找指定的视图文件
    $viewFile = $this->findViewFile($view, $context);
    return $this->renderFile($viewFile, $params, $context);
}
// 支持多种格式查找视图文件
protected function findViewFile($view, $context = null)
{
    if (strncmp($view, '@', 1) === 0) {
        // e.g. "@app/views/main"
        $file = Yii::getAlias($view);
    } elseif (strncmp($view, '//', 2) === 0) {
        // e.g. "//layouts/main"
        $file = Yii::$app->getViewPath() . DIRECTORY_SEPARATOR . ltrim($view, '/');
    } elseif (strncmp($view, '/', 1) === 0) {
        // e.g. "/site/index"
        if (Yii::$app->controller !== null) {
            $file = Yii::$app->controller->module->getViewPath() . DIRECTORY_SEPARATOR . ltrim($view, '/');
        } else {
            throw new InvalidCallException("Unable to locate view file for view '$view': no active controller.");
        }
    } elseif ($context instanceof ViewContextInterface) {
        $file = $context->getViewPath() . DIRECTORY_SEPARATOR . $view;
    } elseif (($currentViewFile = $this->getRequestedViewFile()) !== false) {
        $file = dirname($currentViewFile) . DIRECTORY_SEPARATOR . $view;
    } else {
        throw new InvalidCallException("Unable to resolve view file for view '$view': no active view context.");
    }

    if (pathinfo($file, PATHINFO_EXTENSION) !== '') {
        return $file;
    }
    $path = $file . '.' . $this->defaultExtension;
    if ($this->defaultExtension !== 'php' && !is_file($path)) {
        $path = $file . '.php';
    }

    return $path;
}
// 渲染文件
public function renderFile($viewFile, $params = [], $context = null)
{
    $viewFile = $requestedFile = Yii::getAlias($viewFile);

    // 如果配置了主题
    if ($this->theme !== null) {
        $viewFile = $this->theme->applyTo($viewFile);
    }
    // 加载视图文件
    if (is_file($viewFile)) {
        $viewFile = FileHelper::localize($viewFile);
    } else {
        throw new ViewNotFoundException("The view file does not exist: $viewFile");
    }

    $oldContext = $this->context;
    if ($context !== null) {
        // 把 controller 的实例赋值给 yii\base\View::context 属性
        $this->context = $context;
    }
    $output = '';
    $this->_viewFiles[] = [
        'resolved' => $viewFile,
        'requested' => $requestedFile
    ];

    // 渲染之前和渲染之后，会分别调用 yii\base\View::beforeRender 和 yii\base\View::afterRender 方法触发 yii\base\View::EVENT_BEFORE_RENDER 和 yii\base\View::EVENT_AFTER_RENDER 两个事件
    if ($this->beforeRender($viewFile, $params)) {
        Yii::debug("Rendering view file: $viewFile", __METHOD__);
        $ext = pathinfo($viewFile, PATHINFO_EXTENSION);
        // renderers 可以配置下相关的模板引擎如 twig，根据后缀属性配置的键值找到具体的模板解析类
        if (isset($this->renderers[$ext])) {
            if (is_array($this->renderers[$ext]) || is_string($this->renderers[$ext])) {
                $this->renderers[$ext] = Yii::createObject($this->renderers[$ext]);
            }
            /* @var $renderer ViewRenderer */
            $renderer = $this->renderers[$ext];
            $output = $renderer->render($this, $viewFile, $params);
        } else {
            $output = $this->renderPhpFile($viewFile, $params);
        }
        $this->afterRender($viewFile, $params, $output);
    }

    array_pop($this->_viewFiles);
    $this->context = $oldContext;

    return $output;
}
// 渲染 php 文件
public function renderPhpFile($_file_, $_params_ = [])
{
    $_obInitialLevel_ = ob_get_level();
    ob_start();
    ob_implicit_flush(false);
    // 通过 extract 函数处理页面传参使用
    extract($_params_, EXTR_OVERWRITE);
    try {
        // 视图文件内的 $this 是 view 的实例
        require $_file_;
        return ob_get_clean();
    } catch (\Exception $e) {
        while (ob_get_level() > $_obInitialLevel_) {
            if (!@ob_end_clean()) {
                ob_clean();
            }
        }
        throw $e;
    } catch (\Throwable $e) {
        while (ob_get_level() > $_obInitialLevel_) {
            if (!@ob_end_clean()) {
                ob_clean();
            }
        }
        throw $e;
    }
}
```
之后，调用  yii\base\Controller::renderContent 方法处理（寻找布局文件，找不到就直接返回内容，找到就把内容渲染一下再做返回），最终交由 yii\web\Application::handleRequest 方法接管后续流程。  

此外，如果是 renderAjax，会回调 beginPage、head、beginBody、endBody、endPage 等方法处理环绕文件的渲染，以 ajax 方式返回页面内容（可配合 js/css 实现各种特效）。
```php
// yii/web/View::renderAjax
public function renderAjax($view, $params = [], $context = null)
{
    $viewFile = $this->findViewFile($view, $context);

    ob_start();
    ob_implicit_flush(false);

    $this->beginPage();
    $this->head();
    $this->beginBody();
    echo $this->renderFile($viewFile, $params, $context);
    $this->endBody();
    $this->endPage(true);

    return ob_get_clean();
}
```

**异常处理**  
try/catch 捕获的异常是有限的。  
set_exception_handler 函数（用户自定义异常处理的函数）专门处理未被 try/catch 捕获的异常；并且，类似程序中的 warning、notice 信息，需要另外一个函数捕获 set_error_handler。  
> 对于 php 初始化或者编译等产生的核心错误，我们并不能捕获并处理掉。  
> 以下级别的错误不能由用户定义的函数来处理： E_ERROR、 E_PARSE、 E_CORE_ERROR、 E_CORE_WARNING、 E_COMPILE_ERROR、 E_COMPILE_WARNING，和在调用 set_error_handler () 函数所在文件中产生的大多数 E_STRICT。  

在 yii\base\Application::\_\_construct 方法中，Yii 框架注册了  errorHandler 组件。  
那么，从 yii\base\Application::registerErrorHandler 方法的实现开始，看看 Yii 是如何实现异常处理的。  
```php
// yii\base\Application::registerErrorHandler
protected function registerErrorHandler(&$config)
{
	// YII_ENABLE_ERROR_HANDLER 常量在 yii\BaseYii 类中定义为 true
	// 当然 Yii 也允许用户自定义异常处理
    if (YII_ENABLE_ERROR_HANDLER) {
        if (!isset($config['components']['errorHandler']['class'])) {
            echo "Error: no errorHandler component is configured.\n";
            exit(1);
        }
        // 把 errorHandler 组件（yii\web\ErrorHandler）的定义保存起来
        $this->set('errorHandler', $config['components']['errorHandler']);
        unset($config['components']['errorHandler']);
        // 获取 errorHandler 组件，并调用 yii\web\ErrorHandler::register 方法开始执行
        $this->getErrorHandler()->register();
    }
}
//  yii\web\ErrorHandler::register，即 yii\base\ErrorHandler::register 
public function register()
{
	// 动态修改配置，避免错误信息直接显示到页面上
    ini_set('display_errors', false);
    // 自定义异常处理函数
    set_exception_handler([$this, 'handleException']);
    // 自定义错误处理函数
    if (defined('HHVM_VERSION')) {
        set_error_handler([$this, 'handleHhvmError']);
    } else {
        set_error_handler([$this, 'handleError']);
    }
    if ($this->memoryReserveSize > 0) {
        $this->_memoryReserve = str_repeat('x', $this->memoryReserveSize);
    }
    // 注册程序终止前执行的函数
    register_shutdown_function([$this, 'handleFatalError']);
}
// yii\base\ErrorHandler::handleException 
public function handleException($exception)
{
    if ($exception instanceof ExitException) {
        return;
    }

    $this->exception = $exception;

    // disable error capturing to avoid recursive errors while handling exceptions
    $this->unregister();

    // set preventive HTTP status code to 500 in case error handling somehow fails and headers are sent
    // HTTP exceptions will override this value in renderException()
    if (PHP_SAPI !== 'cli') {
        http_response_code(500);
    }

    try {
    	// 记录日志
    	// 日志记录部分取决于 Yii::error 方法的实现，如果在处理异常的过程中又发生了异常，则依赖 yii\base\ErrorHandler::handleFallbackExceptionMessage 方法了，这个方法就简单多了，直接输出错误信息，不过会有 YII_DEBUG 是否开启的区别
        $this->logException($exception);
        // 丢弃其他任何的输出
        if ($this->discardExistingOutput) {
            $this->clearOutput();
        }
        // 渲染 exception
        $this->renderException($exception);
        if (!YII_ENV_TEST) {
            \Yii::getLogger()->flush(true);
            if (defined('HHVM_VERSION')) {
                flush();
            }
            exit(1);
        }
    } catch (\Exception $e) {
        // an other exception could be thrown while displaying the exception
        $this->handleFallbackExceptionMessage($e, $exception);
    } catch (\Throwable $e) {
        // additional check for \Throwable introduced in PHP 7
        $this->handleFallbackExceptionMessage($e, $exception);
    }

    $this->exception = null;
}
// yii\web\ErrorHandler::renderException
protected function renderException($exception)
{
    if (Yii::$app->has('response')) {
        $response = Yii::$app->getResponse();
        // reset parameters of response to avoid interference with partially created response data
        // in case the error occurred while sending the response.
        $response->isSent = false;
        $response->stream = null;
        $response->data = null;
        $response->content = null;
    } else {
        $response = new Response();
    }

    $response->setStatusCodeByException($exception);

    // 区分 exception 是否是用户主动抛出的异常
    $useErrorView = $response->format === Response::FORMAT_HTML && (!YII_DEBUG || $exception instanceof UserException);

    if ($useErrorView && $this->errorAction !== null) {
        $result = Yii::$app->runAction($this->errorAction);
        if ($result instanceof Response) {
            $response = $result;
        } else {
            $response->data = $result;
        }
    } elseif ($response->format === Response::FORMAT_HTML) {
        if ($this->shouldRenderSimpleHtml()) {
            // AJAX request
            $response->data = '<pre>' . $this->htmlEncode(static::convertExceptionToString($exception)) . '</pre>';
        } else {
            // if there is an error during error rendering it's useful to
            // display PHP error in debug mode instead of a blank screen
            if (YII_DEBUG) {
                ini_set('display_errors', 1);
            }
            $file = $useErrorView ? $this->errorView : $this->exceptionView;
            $response->data = $this->renderFile($file, [
                'exception' => $exception,
            ]);
        }
    } elseif ($response->format === Response::FORMAT_RAW) {
        $response->data = static::convertExceptionToString($exception);
    } else {
        $response->data = $this->convertExceptionToArray($exception);
    }

    $response->send();
}
// yii\base\ErrorHandler::handleError 
public function handleError($code, $message, $file, $line)
{
    if (error_reporting() & $code) {
        // load ErrorException manually here because autoloading them will not work
        // when error occurs while autoloading a class
        if (!class_exists('yii\\base\\ErrorException', false)) {
            require_once __DIR__ . '/ErrorException.php';
        }
        // 抛出 yii\base\ErrorException 异常，经由 yii\base\ErrorHandler::handleException 方法捕获，最终渲染
        $exception = new ErrorException($message, $code, $code, $file, $line);

        // in case error appeared in __toString method we can't throw any exception
        $trace = debug_backtrace(DEBUG_BACKTRACE_IGNORE_ARGS);
        array_shift($trace);
        foreach ($trace as $frame) {
            if ($frame['function'] === '__toString') {
                $this->handleException($exception);
                if (defined('HHVM_VERSION')) {
                    flush();
                }
                exit(1);
            }
        }

        throw $exception;
    }

    return false;
}
// yii\base\ErrorHandler::handleFatalError
public function handleFatalError()
{
    unset($this->_memoryReserve);

    // load ErrorException manually here because autoloading them will not work
    // when error occurs while autoloading a class
    if (!class_exists('yii\\base\\ErrorException', false)) {
        require_once __DIR__ . '/ErrorException.php';
    }

    // 获取最终的错误信息
    $error = error_get_last();

    // 判断错误类型是否是核心错误（E_ERROR, E_PARSE, E_CORE_ERROR 等类型的错误）
    if (ErrorException::isFatalError($error)) {
        if (!empty($this->_hhvmException)) {
            $exception = $this->_hhvmException;
        } else {
            $exception = new ErrorException($error['message'], $error['type'], $error['type'], $error['file'], $error['line']);
        }
        $this->exception = $exception;

        $this->logException($exception);

        if ($this->discardExistingOutput) {
            $this->clearOutput();
        }
        $this->renderException($exception);

        // need to explicitly flush logs because exit() next will terminate the app immediately
        Yii::getLogger()->flush(true);
        if (defined('HHVM_VERSION')) {
            flush();
        }
        exit(1);
    }
}
```

**数据访问层 DAO**  
DAO（Data Access Objects）数据访问对象，DAO 十分友好，比如事务、读写分离、预处理等等，使用都非常简单。  
在 yii2 内，DAO 的封装基于 db 组件。  
```php
// 配置
[
	'components' => [
	    'db' => [
	        'class' => 'yii\db\Connection',
	        'dsn' => 'mysql:host=your host;dbname=your dbname',
	        'username' => 'your username',
	        'password' => 'your password',
	        'charset' => 'utf8',
	    ],
	],
];

// 使用
$sql = "SELECT id, username FROM user"; 
$users = Yii::$app->db->createCommand($sql)->queryAll();
```

db 组件指的是 yii\db\Connection，对 sql 处理的大致情况如下：  
```php
public function createCommand($sql = null, $params = [])
{
    /** @var Command $command */
    // 交给 yii\db\Command 类处理
    $command = new $this->commandClass([
        'db' => $this,
        'sql' => $sql,
    ]);
    return $command->bindValues($params);
}
// sql 处理，需要 yii\db\Command::setSql 方法的实现
// 除了重置的操作之外，调用 yii\db\Connection::quoteSql 方法对 sql 语句做处理
public function quoteSql($sql)
{
	// {{%tablename%}}
    return preg_replace_callback(
        '/(\\{\\{(%?[\w\-\. ]+%?)\\}\\}|\\[\\[([\w\-\. ]+)\\]\\])/',
        function ($matches) {
            if (isset($matches[3])) {
            	// 对列名处理
                return $this->quoteColumnName($matches[3]);
            }

            // 对表名处理
            return str_replace('%', $this->tablePrefix, $this->quoteTableName($matches[2]));
        },
        $sql
    );
}
```
