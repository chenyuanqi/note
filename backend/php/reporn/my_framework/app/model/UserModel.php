<?php

namespace app\model;

use core\lib\model;

class UserModel extends model
{
    public $table = 'users';
    
    public function findAll()
    {
        return $this->select($this->table, '*');
    }

    public function findOneById($id)
    {
        return $this->get($this->table, '*', array(
            'user_id' => $id,
        ));
    }

    public function updateById($id, $data)
    {
        return $this->update($this->table, $data, [
            'id' => $id,
        ]);
    }

    public function save($data)
    {
        return $this->insert($this->table, $data);
    }
}
