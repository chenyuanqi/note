
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


### Yii2 源码解读
