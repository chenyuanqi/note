<?php

/**
 * raw 调试输出
 */
if (!function_exists('dump')) {
    function dump(...$params)
    {
        foreach ($params as $param) {
            var_dump($param);
        }
        exit(1);
    }
}

