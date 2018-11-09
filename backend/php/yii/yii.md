
### Yii 的常用链接
[官网](https://www.yiiframework.com/)  
[官方中文网](https://www.yiichina.com/)  
[yii 第一社区](https://getyii.com/)  
[yii 学习社区-阿北](https://nai8.me/)  
[yiigist](https://yiigist.com/)  
[yii 速查表](https://nai8.me/tool-sc.html)  

### Yii 的生命周期
- 请求到响应的生命周期
```
-------------------------      ----------------------       -------------
| 入口脚本（web/index.php）|     | 应用                |      | 请求处理组件 |
|                        | ——〉|                     | 〈—— --------------
| 1、加载配置文件 -〉启动应用|     | 解析路由 -〉创建控制器 |
-------------------------      ----------------------
                                           |
                                           ∨
                               ----------------------
                               | 控制器              |
                               |                    |      --------------
                               | 创建动作 -〉执行过滤  | 〈—— | 响应处理组件  |
                               |                    |      --------------
                               | 加载数据            | 〈—— | 数据库 -〉模型|
                               |                    |      --------------
                               | 渲染视图            | 〈—— | 视图         |
                               ---------------------       --------------
```

### Yii 的一些概念
- web 应用入口脚本
> 包含 Yii 类文件  
> 读取配置文件，实例化应用主体  
> 解析路由，明确需要创建的路由  

- 应用主体 
> 它是 yii\web\Application 类的实例  
> 它是管理 Yii 应用系统整体结构和生命周期的对象  
> 要求每个应用入口脚本只能创建一个应用主体  
> 可以使用 \Yii::$app 访问  
> 
> 应用主体的配置  
> $config 变量给应用主体这个对象的属性进行初始化赋值  
> $config 变量是从配置文件 web.php 加载而来  
> 
> 应用主体的属性  
> id 属性用来区分其他应用的唯一标识 ID  
> basePath 指定该应用的根目录  
> components 注册多个在其他地方使用的应用组件  
> defaultRoute 如何加载控制器  

- 别名
> 别名是用来表示文件路径和 url ，目的是为了避免在代码中硬编码一些绝对路径和 url  
> 一个别名必须以 @ 字符开头  
> 别名的设置可以这样 Yii::setAlias('@foo', '/path/to')  
> 别名的获取相应的 Yii::getAlias('@foo'), 或在路径中直接使用 new FileCache(['cachePath' => '@foo/cache'])  


### Yii 的跨域问题
```php
// 1、入口处理浏览器的 option 验证请求
if ('OPTIONS' == $_SERVER['REQUEST_METHOD']) {
    header('Access-Control-Allow-Origin: *');
    header('Access-Control-Allow-Methods: GET, POST, OPTIONS');
    header('Access-Control-Allow-Credentials: true');
    header('Access-Control-Allow-Headers: Content-Type, Content-Length, User-Agent, Accept-Language, Accept-Encoding, Authorization, Auth-Sign, Client-date, Client-Timestamp, Accept, X-Requested-With');

    echo json_encode(['code' => 200, 'message' => 'ok']);exit();
}

// 2、相应 Response 基类设置头部
public function sendHeaders()
{
  $this->headers->add("Access-Control-Allow-Origin", "*");
  $this->headers->add("Access-Control-Allow-Methods", "GET, POST, OPTIONS");
  $this->headers->add('Access-Control-Allow-Credentials', 'true');
  $this->headers->add("Access-Control-Allow-Headers", "Content-Type, Content-Length, User-Agent, Accept-Language, Accept-Encoding, Authorization, Auth-Sign, Client-date, Client-Timestamp, Accept, X-Requested-With");
  $this->headers->add('Hxxx', gethostname());

  return parent::sendHeaders();
}
// 或在控制器限制行为
[
    'class' => Cors::className(),
    'cors' => [
        'Origin' => ['*'],
        'Access-Control-Request-Method' => [],
        'Access-Control-Request-Headers'=>['*']
    ],
]
```
