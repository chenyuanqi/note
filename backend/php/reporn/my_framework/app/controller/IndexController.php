<?php

namespace app\controller;

use app\model\UserModel;
use core\lib\controller;

class IndexController extends controller
{
    public function actionIndex()
    {
        $model = new UserModel();
        $data  = $model->findAll();

        $this->response(10000, $data);
    }

    public function actionDemo()
    {
        $this->response(10001, []);
    }
}
