<?php

namespace core\lib;

use Medoo\Medoo;

class model extends Medoo
{
    public function __construct()
    {
        parent::__construct([
            'type'     => 'mysql',
            'database' => 'api',
            'host'     => '127.0.0.1',
            'username' => 'root',
            'password' => 'root',
            'charset'  => 'utf8mb4',
            'port'     => 3306,
            'prefix'   => '',
        ]);
    }
}
