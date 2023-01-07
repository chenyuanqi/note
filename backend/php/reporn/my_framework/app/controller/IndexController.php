<?php

namespace app\controller;

use app\model\UserModel;
use core\lib\controller;

class IndexController extends controller
{
    public function actionIndex()
    {
        $model = new UserModel();
        $list  = $model->findAll();

        $this->response(10000, compact('list'));
    }

    public function actionDemo()
    {
        $this->response(10001, []);
    }
}
