
### easyswoole 概览
EasySwoole 是一款基于 Swoole Server 开发的常驻内存型的分布式 PHP 框架，专为 API 而生，摆脱传统 PHP 运行模式在进程唤起和文件加载上带来的性能损失。 EasySwoole 高度封装了 Swoole Server 而依旧维持 Swoole Server 原有特性，支持同时混合监听 HTTP、自定义 TCP、UDP 协议，让开发者以最低的学习成本和精力编写出多进程，可异步，高可用的应用服务。

[easyswoole 官网](https://www.easyswoole.com/)  
[easyswoole Github](https://github.com/easy-swoole)  
[easyswoole Demo](https://github.com/easy-swoole/demo/)   

特性：  
- 强大的 TCP/UDP Server 框架，多线程，EventLoop，事件驱动，异步，Worker 进程组，Task 异步任务，毫秒定时器，SSL/TLS 隧道加密  
- EventLoop API，让用户可以直接操作底层的事件循环，将 socket，stream，管道等 Linux 文件加入到事件循环中  
- 定时器、协程对象池、HTTP\SOCK 控制器、分布式微服务、RPC 支持  

优势：  
- 简单易用开发效率高  
- 并发百万 TCP 连接  
- TCP/UDP/UnixSock  
- 支持异步 / 同步 / 协程  
- 支持多进程 / 多线程  
- CPU 亲和性 / 守护进程  

目录结构：  
EasySwoole 的目录结构是非常灵活的，基本上可以任意定制，没有太多的约束，但是仍然建议遵循下面的目录结构，方便开发。  
```
project                   项目部署目录
├─App                     应用目录(可以有多个)
│  ├─HttpController       控制器目录
│  │  └─Index.php         默认控制器
│  └─Model                模型文件目录
├─Log                     日志文件目录
├─Temp                    临时文件目录
├─vendor                  第三方类库目录
├─composer.json           Composer架构
├─composer.lock           Composer锁定
├─EasySwooleEvent.php     框架全局事件
├─easyswoole              框架管理脚本
├─dev.php                 开发配置文件
├─produce.php             生产配置文件
```

服务器配置：  
```
# Nginx
server {
    root /data/wwwroot/;
    server_name local.swoole.com;
    location / {
        proxy_http_version 1.1;
        proxy_set_header Connection "keep-alive";
        proxy_set_header X-Real-IP $remote_addr;
        if (!-f $request_filename) {
             proxy_pass http://127.0.0.1:9501;
        }
    }
}

# Apache
<IfModule mod_rewrite.c>
  Options +FollowSymlinks
  RewriteEngine On
  RewriteCond %{REQUEST_FILENAME} !-d
  RewriteCond %{REQUEST_FILENAME} !-f
  # RewriteRule ^(.*)$ index.php/$1 [QSA,PT,L]  fcgi下无效
  RewriteRule ^(.*)$  http://127.0.0.1:9501/$1 [QSA,P,L]
  # 请开启 proxy_mod proxy_http_mod request_mod
</IfModule>
```

### easyswoole orm
关于模型的定义。
```php
namespace App\Model;

use EasySwoole\ORM\AbstractModel;
use EasySwoole\ORM\Utility\Schema\Table;

class UserModel extends AbstractModel
{
	// 指定连接名
	protected $connectionName = 'default';
	// 指定表名
	protected $tableName = 'user';
	// 是否自动填充时间戳
	protected $autoTimeStamp = true;
	protected $createTime = 'created_at';
    protected $updateTime = 'updated_at';

    /**
     * 表的获取（获取建表语句 $model->schemaInfo()->__toString()）
     * 
     * @return Table
     */
    public function schemaInfo(bool $isCache = true): Table
    {
        $table = new Table();
        $table->colInt('id')->setIsPrimaryKey(true);
        $table->colChar('name', 255)->setIsNotNull()->setColumnComment('姓名');
        $table->colInt('age')->setIsNotNull()->setDefaultValue(0)->setColumnComment('年龄');
        $table->colTinyInt('is_vip', 1)->setIsNotNull()->setDefaultValue(0)->setColumnComment('是否 vip');

        return $table;
    }
}
```

Orm 增删改查。  
```php
$user = UserModel::create();
// 插入一条新用户记录，save 返回值是 true 或 false
$user->data([
    'attr' => 'value'
]);
$user->save();
// 或者这样新增
$model = UserModel::create([
    'name' => 'vikey',
    'age'  => 18,
]);
$res = $model->save();
// 批量添加：saveAll($data, $replace = true, $transaction = true): [key => boolean]
UserModel::create()->saveAll([
    ['name' => 'test1'],
    ['name' => 'test2'],
]);

// 根据 id 删除用户，删除成功返回影响行号/失败返回 false
$user = UserModel::create()->get($param['id']); // 根据 id 查询
if (!$user){
    return '操作数据不存在，请检查再试';
}
$res = $user->destroy();
// 删除多条记录
$res = UserModel::create()->destroy('2,4,5');// 指定多个参数每个参数为不同主键
$res = UserModel::create()->destroy([3, 7]);// 数组指定多个主键
// 清空表数据
$res = UserModel::create()->destroy(null,true);

// 更新，返回值为 true/false
// 通过已有 model 更新
$user = UserModel::create()->get(1);
$user->update([
  'is_vip' => 1
]); // 或者 $user->is_vip = 1; $user->update();
// 通过 where 条件更新
$res = UserModel::create()->update(['name' => 'new vikey'], ['id' => 1]);
$res = UserModel::create()->update(['is_vip' => 0], ['vip_time' => 0]);
UserModel::create()->update([
    'age' => QueryBuilder::inc(3), // 自增 3
    'number' => QueryBuilder::dec(4), // 自降 4
], [
    'name' => 'vikey'
]);

// 多条件查询
// get 返回值：
// 数据不存在返回 false 或 null
// 如设置第二参数为 true 则返回数组
// 其他返回 model，model 属性获取如 $res->name
$res = UserModel::create()->get([
  'uid'    => 1,
  'state'  => 0,
  'is_vip' => 1
]);
// 多个主键
$user = UserModel::create()->where([1,2,3])->all();
// 复杂条件数组
$user = UserModel::create()->where([
    'age'  => [[18,23], 'between'],
    'name' => ['siam', 'like', 'or'],
]);
// 走 builder 原生 where
$res = UserModel::create()->where('id', 1, '!=')->get();
$res = UserModel::create()->where('id', 1, 'like')->get();
// alias 用于设置当前数据表的别名
$res = UserModel::create()->alias('siam')->where(['siam.name' => 'test'])->all();
// group 方法可以将结果分组
$res = UserModel::create()->field('sum(age) as age, `name`')->group('name')->all();
// order 方法可用于将原生字符串设置为 order by 子句
$res = UserModel::create()->order('id', 'DESC')->get();
// 连表操作
$res = UserModel::create()->join('table','table.col1 = user.col2')->get();
$res = UserModel::create()->alias('u')->join('table as t','t.col1 = u.col2')->get();

// 聚合查询
$max = UserModel::create()->max('age');
$min = UserModel::create()->min('age');
$count = UserModel::create()->count(); // 获取总记录数
$avg = UserModel::create()->avg('age');
$sum = UserModel::create()->sum('age');

// 获取器：对 model 的属性获取时自动处理
// 定义数据表中存在的字段
protected function getStatusAttr($value, $data)
{
    $status = [-1=>'删除',0=>'禁用',1=>'正常',2=>'待审核'];
    return $status[$value];
}
// 定义数据表中不存在的字段
protected function getTestAttr($value, $data)
{
  return 'test-'.$data['id'];
}
$res = UserModel::create()->get(3);
var_dump($res->status);
var_dump($res->test);

// 修改器：修改字段赋值时自动进行处理
protected function setNameAttr($value, $data)
{
    return $value."_统一后缀";
}
$model = new UserModel([
    'name' => 'siam',
    'age'  => 21,
]);
$model->save(); // name 存入后值为: siam_加一个统一后缀
```

Orm 事务操作。
```php
$user = UserModel::create()->get(4);
$user->age = 4;
// 开启事务
DbManager::getInstance()->startTransaction($connectionNames = 'default');
// 更新操作
$res = $user->update();
// 直接回滚
$rollback = DbManager::getInstance()->rollback();
// 返回 false 因为连接已经回滚，事务关闭
$commit = DbManager::getInstance()->commit();
var_dump($commit);
```


