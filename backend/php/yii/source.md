
### Yii 生命周期


**Yii 一次请求的完整过程**
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

### Yii 源码解读基础 —— 组件


