
### Yii2 微服务概念
把一个大型的单个应用程序和服务拆分为数个甚至数十个的支持微服务，它可扩展单个组件而不是整个的应用程序堆栈，从而满足服务等级协议。

传统的开发模式就是把所有功能都放在一个包里，基本不存在依赖，这样的优势在于开发简单，集中式管理，功能都在本地，不存在分布式的管理和调度消耗。但缺点也很明显：效率低，开发都在同一个项目改代码，相互等待，冲突不断。稳定性差，一个微小的问题，都可能导致整个应用挂掉。另外在资源利用上表现出明显的劣势，比如电商双11大促场景，下单压力非常大，评价的压力相对较少，那么我们希望临时增配应对双11的大流程，只能全部增配，而不能定点只对订单服务增配。所以微服务的架构开始慢慢流行并应用于大型的网站平台。

那么引入今天的主题，Yii 如何做微服务？Yii 可以轻松使用，而不需要基本和高级模板中包含的功能。换句话说，Yii 已经是一个微框架。不需要由模板提供的目录结构与 Yii 一起工作。

### 安装 Yii
为您的项目创建一个目录并将工作目录更改为该路径。示例中使用的命令是基于 Unix 的，但在 Windows 中也存在类似的命令。
```
mkdir micro-app
cd micro-app
```
Note：需要一些 Composer 的知识才能继续。如果您还不知道如何使用 composer，请花些时间阅读 Composer 指南。

使用您最喜爱的编辑器在 micro-app 目录下创建 composer.json 文件并添加以下内容：
```
{
    "require": {
        "yiisoft/yii2": "~2.0.0"
    },
    "repositories": [
        {
            "type": "composer",
            "url": "https://asset-packagist.org"
        }
    ]
}
```
保存文件并运行 composer install 命令。这将安装框架及其所有依赖项。

### 创建项目结构
安装框架之后，需要为此应用程序创建一个入口点。入口点是您尝试打开应用程序时将执行的第一个文件。出于安全原因，建议将入口点文件放在一个单独的目录中，并将其设置为Web根目录。

创建一个 web 目录并将 index.php 放入其中，内容如下：
```
<?php

// comment out the following two lines when deployed to production
defined('YII_DEBUG') or define('YII_DEBUG', true);
defined('YII_ENV') or define('YII_ENV', 'dev');

require(__DIR__ . '/../vendor/autoload.php');
require(__DIR__ . '/../vendor/yiisoft/yii2/Yii.php');

$config = require __DIR__ . '/../config.php';
(new yii\web\Application($config))->run();
```
还要创建一个名为 config.php 的文件，它将包含所有的应用程序配置：
```
<?php
return [
    'id' => 'micro-app',

    //设置`micro-app`的根目录
    'basePath' => __DIR__,

    // 控制器所在目录。
    'controllerNamespace' => 'micro\controllers',

    // 设置命名空间为 micro
    'aliases' => [
        '@micro' => __DIR__,
    ],

    //默认访问地址
    'defaultRoute' => 'home/index',

    'components' => [
        //请求配置
        'request' => [
            'cookieValidationKey' => 'test&123456',
            'parsers' => [
                'application/json' => 'yii\web\JsonParser',
            ]
        ],

        //Url 美化
        'urlManager' => [
            'enablePrettyUrl' => true,
            'showScriptName' => false,
            'enableStrictParsing' => false,
            'rules' => [
                '<controller:\w+>/<action:\w+>/<id:\w+>'   => '<controller>/<action>',
            ],
        ],

        //数据库配置
        'db' => [
            'class' => 'yii\db\Connection',
            'dsn' => 'mysql:host=localhost;dbname=micro',
            'username' => 'root',
            'password' => '数据库密码',
            'charset' => 'utf8',
        ],
    ],

];
```
Info：尽管配置可以保存在 index.php 文件中，建议单独使用它。 这样它也可以用于控制台应用程序，如下所示。

您的项目现在已经准备进行编码了。尽管由您决定项目目录结构，只要您遵守命名空间即可。

### 创建第一个控制器
在创建控制器之前，创建一个 controllers/base 目录并创建一个基础控制器 BaseController。
```
<?php
namespace micro\controllers\base;

use yii\web\Controller;

class BaseController extends Controller
{
    //关闭 csrf 验证
    public $enableCsrfValidation = false;
}
```
然后在 controller 文件夹下面 新建一个 SiteController.php，这是默认的 控制器将处理没有路径信息的请求。
```
<?php

namespace micro\controllers;

use yii\web\Controller;

class HomeController extends BaseController
{
    public function actionIndex()
    {
        return '欢迎来到 Yii2.0 微服务！';
    }
}
```
如果您想为此控制器使用不同的名称，则可以配置 yii\base\Application::$defaultRoute 进行更改。 例如，对于 HomeController 将会是 'defaultRoute' => 'home/index'。

在这一点上，项目结构应该如下所示：
```
micro-app/
├── composer.json
├── config.php
├── web/
    └── index.php
└── controllers/
    └── base
        └── BaseController.php
    └── HomeController.php
└── vendor
```
如果您尚未设置 Web 服务器，则可能需要查看Web服务器配置文件示例。 另一种选择是使用 yii serve 命令，它将使用 PHP 内置 web 服务器。 您可以通过以下方式从 micro-app / 目录运行它：
```
vendor/bin/yii serve --docroot=./web
```
在浏览器中打开应用程序URL现在应该打印出“欢迎来到 Yii2.0 微服务！”，它已经在 HomeController::actionIndex() 中返回。

Info：在我们的示例中，我们已将默认应用程序名称空间 app 更改为 micro， 以表明您不受此名称的限制（如果您是这样认为）， 然后调整 controllers namespace 并设置正确的别名。