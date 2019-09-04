
### 常用增删改查
```php
use DB;
// 查
$users = DB::select ('select * from users where active = ?', [1]);
// 查询所有信息
$users = DB::table('users')->get();
// 查询 name 为 John 的第一行
$user = DB::table('users')->where('name', 'John')->first();
// 聚合
$users = DB::table('users')->count();
$price = DB::table('orders')->max('price');
// 取所有不同
$users = DB::table('users')->distinct()->count();
// 排序
$users = DB::table('users')->orderBy('field', 'value')->get();
// 插入
DB::insert ('insert into users (id, name) values (?, ?)', [1, 'Dayle']);
DB::table('users')->insert([
    ['email' => 'taylor@example.com', 'votes' => 0],
    ['email' => 'dayle@example.com', 'votes' => 0]
]);
$id = DB::table('users')->insertGetId(
    ['email' => 'john@example.com', 'votes' => 0]
);
// 更新
DB::insert ('insert into users (id, name) values (?, ?)', [1, 'Dayle']);
DB::table('users')->where('id', 1)->update(['votes' => 1]);
DB::table('users')->increment('votes', 5);
// 更新数字加字段组合
DB::table('users')->increment('votes', 1, ['name' => 'John']);
// 删除
$deleted = DB::delete ('delete from users');
DB::table('users')->delete();
DB::table('users')->where('votes', '>', 100)->delete();
// 无返回值删除
DB::statement ('drop table users');
// 开启事务
DB::beginTransaction();
// 回滚事务
DB::rollBack();
// 提交事务
DB::commit();


// 设置关联模型的表
protected $table = '';
// 设置是否表中有 create_at 和 update_at
public $timestamps = fasle;
// 设置日期存储格式
protected $dateFormat = 'U';

// 根据 id 查询，返回 model 或 null
Model::find($id);
// 根据 id 查询，返回 model 或抛出 error 异常
Model::findOrFail($id)
// 根据 id 查询，返回在数据库中找到的第一条记录或 null
Model::where(['id' => $id])->first();
// 根据 id 查询，返回在数据库中找到的第一条记录或抛出 error 异常
Model::where(['id' => $id])->firstOrFail();
// 原生条件
$builder->whereRaw("job_id in (select id from jobs where status=1)");
// 获取与查询匹配的模型集合
Model::get();
// 返回给定列
Model::get()->pluck([$field]);
// 集合转数组
Model::get()->toArray()

// 分页查询
Model::query()->orderBy('created_at', 'desc')->paginate(config('website.admin.page_size'));

// 写入
Model::insert([
    'title' => $data['title'],
    'description' => $data['description'],
]);

// 更新
$data = Model::find($id);
$data->title = 'xxx';
$data->save();
// 或者
Model::where('id','=';$id)->update(['title' => 'xxx'])

// 软删除
// model 增加 use SoftDeletes;
// 数据表增加 deleted_at
// 迁移增加 $table->softDeletes();
// 删除具体使用
Model::find($id)->delete();
Model::destroy($id);
// 强制删除
Model::find($id)->forceDelete();
```

### 代码片段
```php
// 获取执行 sql（返回结果带参数及执行时间）
DB::enableQueryLog();
// $builder ...
dd(DB::getQueryLog());

// 获取执行 sql（返回结果不带参数）
$builder->toSql();
```

### 常用命令


### 常见问题

