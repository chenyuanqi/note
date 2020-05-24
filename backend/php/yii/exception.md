

### 什么是 Exception
Exception 即异常。在应用开发中，错误和异常处理机制是一块比较重要的模块。


### Yii2 的 Exception
yii2 的 yii\web\HttpException 默认对 web 请求友好，都是以 text/html 的方式返回错误描述，但是对 api 不友好（我们需要即时的以抛出异常的方式中断请求的处理，并以全局异常处理器格式化处理后统一返回给客户端）。

1、注册异常处理器  
yii2 也是以 controller/action 的方式定义一个异常处理器的，我们可以在配置项 components => errorHandler 中自定义。  
```php
# config/web.php
'components' => [
    'errorHandler' => [
        'errorAction' => 'exception/handler'
    ]
]
```

2、异常处理器  
定义相应的异常处理器，app\actions\ErrorApiAction 继承 yii\web\ErrorAction，可以拿到 yii2 为我们整理好的全局异常。
```php
# controllers/ExceptionController.php
namespace app\controllers;

use yii\web\Controller;

class ExceptionController extends Controller
{
    /**
     * 为 actionHandler 挂载独立的 action

     * @return array
     */
    public function actions()
    {
        return [
            'handler' => [
                'class' => 'app\actions\ErrorApiAction',
            ]
        ];
    }
}

#actions/ErrorApiAction.php
namespace app\actions;

use Yii;
use yii\web\ErrorAction;
use yii\web\Response;

class ErrorApiAction extends ErrorAction
{
    public function run()
    {
        // 根据异常类型设定相应的响应码
        Yii::$app->getResponse()->setStatusCodeByException($this->exception);
        // json 格式返回
        Yii::$app->getResponse()->format = Response::FORMAT_JSON;
        // 返回的内容数据
        return [
            'msg' => $this->exception->getMessage(),
            'err' => $this->exception->getCode()
        ];
    }
}
```

3、定义异常实体  
主要是简单的把状态码的传递封装一下，用更容易理解的类名来代理传递。  
```php
# exceptions/ApiBaseException.php
<?php
/**
 * api 异常基础类
 */
namespace app\exceptions;

class ApiBaseException extends \yii\web\HttpException
{
    public function __construct($message = null, $code = 0, \Exception $previous = null)
    {
        parent::__construct($this->statusCode, $message, $code, $previous);
    }
}

# exceptions/UnauthorizedException.php
<?php
/**
 * 401 unauthorized
 */

namespace app\exceptions;

class UnauthorizedException extends ApiBaseException
{
    public $statusCode = 401;
}
```

4、使用范例  
在一些 service logic model 中根据需要即时抛出异常即可，上层控制器拿到的永远都是正常的返回数据。  
```php
throw new UnauthorizedException("请认证后访问");
```
