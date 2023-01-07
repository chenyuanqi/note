<?php

namespace core;

use core\lib\log;
use core\lib\route;

class bootstrap
{
    public static $classMap = [];

    public static function run()
    {
        $route           = new route();
        $controllerClass = $route->controller;
        $action          = $route->action;
        $controllerFile  = APP_PATH . DIRECTORY_SEPARATOR . 'controller' . DIRECTORY_SEPARATOR . $controllerClass . 'Controller.php';
        $controllerClass = '\\app\controller\\' . $controllerClass . 'Controller';
        if (is_file($controllerFile)) {
            include $controllerFile;
            $controller = new $controllerClass();
            log::write($controllerClass . '--' . $action);
            
            return $controller->$action();
        } 
        
        throw new \Exception("Controller Not Found.");
    }

    public static function load($class)
    {
        if (isset($classmap[$class])) {
            return true;
        } 
        
        $class = str_replace('\\', '/', $class);
        $filePath = ROOT_PATH . DIRECTORY_SEPARATOR . $class . '.php';
        if (is_file($filePath)) {
            include $filePath;
            self::$classMap[$class] = $filePath;
        } 

        return false;
    }
}
