<?php

// TODO：
// 1、自动加载配置 config
// 2、http 请求响应
// 3、...
define('ROOT_PATH', realpath('./'));
define('CORE_PATH', realpath('./core'));
define('APP_PATH', realpath('./app'));

define('APP_ENV', 'dev');
define('APP_DEBUG', APP_ENV === 'dev');


require ROOT_PATH . DIRECTORY_SEPARATOR . 'vendor' . DIRECTORY_SEPARATOR . 'autoload.php';
require CORE_PATH . DIRECTORY_SEPARATOR . 'common' . DIRECTORY_SEPARATOR . 'function.php';
require CORE_PATH . DIRECTORY_SEPARATOR . 'bootstrap.php';

if (APP_DEBUG) {
    $whoops = new \Whoops\Run;
    $whoops->pushHandler(new \Whoops\Handler\PrettyPageHandler);
    $whoops->register();
} 

spl_autoload_register('\core\bootstrap::load');
\core\bootstrap::run();
