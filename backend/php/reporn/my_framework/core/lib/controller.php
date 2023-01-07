<?php

namespace core\lib;

class controller 
{
    public function corsHeader() 
    {
        if ($referer = $_SERVER['HTTP_REFERER'] ?? '') {
            $parsed = parse_url($referer);
            $origin = sprintf("%s://%s", $parsed['scheme'] ?? "http", $parsed['host']);
        }

        $xorigin = $_SERVER['HTTP_X_ORIGIN'] ?? null;
        $origin = $_SERVER['HTTP_ORIGIN'] ?? null;
        header("Access-Control-Allow-Origin: " . $xorigin ?: $origin);
        header("Access-Control-Allow-Headers: x-requested-with, X-Origin, Origin, Accept, Content-Type, Referer, Referrer-Policy, User-Agent, Cookie, access-token, crossdomain, withCredentials, authorization");
        header("Access-Control-Allow-Credentials: true");
        header('Access-Control-Allow-Methods: OPTIONS, GET, POST, DELETE, PUT');
        header('Access-Control-Max-Age: 60');
        header('Content-Type: application/json');
    }

    public function response($code = 10000, $data = [])
    {
        $this->corsHeader();

        exit(json_encode(['code' => $code, 'data' => $data], JSON_UNESCAPED_UNICODE));
    }
}
