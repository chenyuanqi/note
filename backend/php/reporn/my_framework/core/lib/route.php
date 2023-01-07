<?php

namespace core\lib;

class route
{
    public $controller;
    public $action;
    public function __construct()
    {
        $pathArr          = explode('/', trim($_SERVER['REQUEST_URI'], '/'));
        $this->controller = ucwords(($pathArr[0] ?? '') ?: 'index');
        $this->action     = 'action' . ucwords(($pathArr[1] ?? '') ?: 'index');
    }
}
