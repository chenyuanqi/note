<?php

namespace core\lib;

class log
{
    public function __construct()
    {
    }

    public static function write($msg, $file = 'log')
    {
        $path = ROOT_PATH . '/runtime/logs';
        if (!is_dir($path)) {
            @mkdir($path, '0777', true);
        }

        $msg = "{$msg} " . date('Y-m-d H:i:s');
        return file_put_contents($path . DIRECTORY_SEPARATOR . $file . '-' . date('Ymd') . '.php', is_array($msg) ? json_encode($msg) : $msg . PHP_EOL, FILE_APPEND);
    }

    public static function __callStatic($method, $args) 
    {
        return (new static)->$method(...$args);
    }
}
