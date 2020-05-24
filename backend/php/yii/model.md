
### Model 常用方法
```php
// 多语言的情况下，可翻译属性标签
public function attributeLabels()
{
    return [
        'name' => \Yii::t('app', 'Your name'),
        'email' => \Yii::t('app', 'Your email address'),
        'subject' => \Yii::t('app', 'Subject'),
        'body' => \Yii::t('app', 'Content'),
    ];
}

// 验证规则
public function rules()
{
    return [
        // 在"register" 场景下 username, email 和 password 必须有值
        [['username', 'email', 'password'], 'required', 'on' => 'register'],

        // 在 "login" 场景下 email 和 password 必须有值
        [['email', 'password'], 'required', 'on' => 'login'],

        // email 属性必须是一个有效的电子邮箱地址
        ['email', 'email'],
    ];
}
// 场景设置安全属性（只覆盖配置中的字段）
public function scenarios()
{
    return [
        'login' => ['username', 'password'],
        'register' => ['username', 'email', 'password'],
    ];
}

// 字段映射
public function fields()
{
    return [
        // 字段名和属性名相同
        'id',

        // 字段名为 "email"，对应属性名为 "email_address"
        'email' => 'email_address',

        // 字段名为 "name", 值通过 PHP 代码返回
        'name' => function () {
            return $this->first_name . ' ' . $this->last_name;
        },
    ];
}
```


### Model 关联模型
表与表的关系，可以是一对一、一对多或是多对一。  

举个例子，我们又两张表：  
customer（用户表）（id, name）  
order（订单表）（id, customer_id, price）  
这里，customer 表和 order 表之间是一对多的关系，通过 customer_id 字段关联。相应的 model 设计如下：  
```php
class Customer extends \yii\db\ActiveRecord
{
    /**
     * 获取订单信息
     */
    public function getOrders()
    {
        // 一个用户对应多个订单，一对多的关系使用 hasMany()关联
        // 关联的条件：【关联表字段】 => 【当前 model 字段】
        return $this->hasMany(Order::className(), ['customer_id' => 'id'])->asArray();
    }
}

class Order extends \yii\db\ActiveRecord
{
    /**
     * 获取用户信息
     */
    public function getCustomer()
    {
        // 一个订单对应一个用户，一对一的关系使用 hasOne() 关联
        // 同样的，关联的条件：【关联表字段】 => 【当前 model 字段】
        return $this->hasOne(Customer::className(), ['id' => 'customer_id'])->asArray();
    }
}
```

那么，我们可以这样使用：
```php
// 查询客户名为张三的订单信息
$customer = Customer::find()->where(['name' => '张三'])->one();
// 可以通过 $customer->getOrder() 方法调用 customer.php 模型中 getOrders() 方法
$orders = $customer->getOrders()->all();
var_dump($orders);exit;

// 也可以使用 $customer->orders 属性调用，
// 当程序没有找到 orders 属性时，PHP 会调用__get()函数，程序会自动寻找 getOrdesr() 方法调用
// 这里不用加 all()，程序会自动识别，如果使用的是hasMany则加 all()，hasOne 则加上 one()
$orders = $customer->orders;
var_dump($orders);exit;
 
// 查询订单 id 为 1 的客户信息
$order = Order::find()->where(['id' => 1])->one();
// 调用 order.php 模型中 getCustomer() 方法
$customer = $order->customer;
var_dump($customer);exit;

// 当 customer 订单量过大时，可以使用 with 语句
// 相当于执行了两条 SQL 语句 select * from customer 和 select * from order where customer_id in(...)
$customers = Customer::find()->with('order')->asArray()->all();
foreach ($customers as $customer) {
    $orders = $customer['order'];  // 取得 order 表的关联数据
}
// joinWith () 的用法和 with () 差不多，joinWith () 可以指定连接类型，默认 LEFT JOIN 连接
```
