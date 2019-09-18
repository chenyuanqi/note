
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

**应用的生命周期**  
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
