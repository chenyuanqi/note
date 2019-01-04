
### 常用代码
```php
// 获取 post 传参 (结果是 array)
// Yii::$app->request->post()
Yii::$app->request->bodyParams 
// 获取 get 传参 (结果是 array)
// Yii::$app->request->get()
Yii::$app->request->queryParams;

// 获取接口传过来的 JSON 数据
Yii::$app->request->rawBody;

//  防止 SQL 和 Script 注入
use yii\helpers\Html;
use yii\helpers\HtmlPurifier;

echo Html::encode($view_hello_str) // 可以原样显示<script></script>代码  
echo HtmlPurifier::process($view_hello_str)  // 可以过滤掉<script></script>代码
```

### 常用 AR
```php
//  is not null条件查询
$query->andWhere(['not', ['state' => null]])

// 多条件组合
$logQuery->andWhere([
'OR',
['user' => array_values(array_unique($userIdArr))],
['like', 'aname', $params->word],
]);
User::find()->where(['and', ['xxx' => 0, 'yyy' => 2], ['>', 'zzz', $time]]);

// 使用 find_in_set
User::find() ->where(new Expression('FIND_IN_SET(:status, status)'))->addParams([':status' => 1])->all();

// 模糊查询，第 4 个参数是否自动给关键词前后加"%"
$query = User::find() ->where(['LIKE', 'name', $keyword.'%', false]);

//避免 select 里面的子查询被识别成字段
$query = User::find() ->select([ new Expression('count(*) as count , count(distinct mobile) as mnumber') ]);

// 关于使用事务
Yii::$app->db->transaction(function() {
    $order = new Order($customer);
    $order->save();
    $order->addItems($items);
});
// 这相当于下列冗长的代码：
$transaction = Yii::$app->db->beginTransaction();
try {
    $order = new Order($customer);
    $order->save();
    $order->addItems($items);
    $transaction->commit();
} catch (\Exception $e) {
    $transaction->rollBack();
    throw $e;
}
```

### 常用 Model
```php
// 验证规则
public function rules()
{
	[['email'], 'required'],
	[['email'], 'email'],
	[['email'], 'unique'],

	[['last_logined_ip'], 'ip'],

	[['cover_thumb'], 'url'],

	[['status'], 'default', 'value' => static::STATUS_OFF],
	[['status'], 'in', 'range' => array_keys(static::$statusLabel)],

    [['target_id'], 'integer'],

    [['score'], 'number', 'min' => 1, 'max' => 10, 'integerOnly' => true],

	[['salt'], 'string', 'length' => 32],

    [['content'], 'string', 'max' => 100],
	[['content'], 'trim'],

	[['password'], 'string', 'min' => 6],
    [
        ['password'],
        'match',
        'pattern' => '/^\d*$/',
        'not' => true,
        'message' => '{attribute}不能是纯数字.'
    ],

    [
        ['confirm_password'],
        'compare',
        'compareAttribute' => 'new_password',
        'operator' => '===',
        'message' => '两次输入的密码不一致.',
    ],

    [
        ['owner'],
        'required',
        'when' => function ($model){
            return static::OWNER_TYPE_COMPANY == $model->owner_type;
        }
    ],

    [
        ['title'],
        'filter',
        'filter' => function ($value){
            return StringHelper::msubstr($value, 0, 50);
        }
    ],

    [
        ['audit_status'],
        'filter',
        'filter' => function (){
            return Teachers::AUDIT_STATUS_REVIEW_SUCCESS;
        },
        'when' => function ($model){
            return static::STATUS_NORMAL == $model->status;
        }
    ],

    [
        ['joined_count'],
        'integer',
        'when' => function (self $model){
            return !$model->joined_count instanceof Expression;
        }
    ],
}

// 状态常量定义
const IS_DELAY_YES = 1;
const IS_DELAY_NO  = 0;
public static $isDelayLabel = [
    self::IS_DELAY_YES => ['title' => '是', 'value' => self::IS_DELAY_YES],
    self::IS_SYN_NO => ['title' => '否', 'value' => self::IS_SYN_NO],
];


// 自动记录写入、更新时间
use yii\behaviors\TimestampBehavior;
/**
 * @return array
 */
public function behaviors()
{
    return [
        [
            'class' => TimestampBehavior::className(),
            'createdAtAttribute' => 'created_at',
            'updatedAtAttribute' => 'updated_at',
        ],
    ];
}

// 写入数据
$xxx = new XXX();
$data = [
    'company' => $data['company'],
    'name' => $data['name']
];
$xxx->setAttributes($data);
if (false === $xxx->save()){
    // 取第一条报错，current($xxx->getErrors())[0]
    throw new ServiceException('插入失败，error：' . Json::encode($xxx->getErrors()), 0);
}

// ---
$xxx = new XXX();
// $xxx->loadDefaultValues();
$attributes = Yii::$app->request->post();
$queryForm->load($attributes, '');
if ($queryForm->validate()){
	return $this->renderJson([
        'code' => 400,
        'message' => current($queryForm->getErrors())[0]
    ]);
}
$xxx->user_id = $attributes->user_id;
if (false === $xxx->save()){
    throw new ServiceException('插入失败，error：' . Json::encode($xxx->getErrors()), 0);
}

// 写入或更新操作
$table = XXX::tableName();
$sql = <<<sql
INSERT INTO {$table}(`company`,`year`,`month`,`day`,`date`,`total_points`,`gain_points`,`consume_points`,`gain_points_man`,`gain_points_time`,`created_at`,`updated_at`) VALUES(
{$data['company']},{$year},{$month},{$day},'{$data['date']}',{$balancePoint},{$usePoint},{$getPoint},1,{$userNum},{$timestamp},{$timestamp}
) ON DUPLICATE KEY UPDATE `total_points`={$balancePoint},`consume_points`=`consume_points`+{$getPoint},`gain_points`=`gain_points`+{$usePoint},`gain_points_man`=`gain_points_man`+1,`gain_points_time`={$userNum},`updated_at`={$timestamp}
sql;
$query = \Yii::$app->db->createCommand($sql);
$result = $query->execute();
if (!$result){
    throw new ServiceException('保存xxx失败', 2);
}

/**
 * 追加一个用户的学习时间
 *
 * @param $time
 */
public function appendLastLearnTimes($time)
{
    $this->last_learn_times = ($time . ',' . $this->last_learn_times);
    $times = $this->getLastLearnTimesArray();

    $this->last_learn_times = implode(',', array_slice($times, 0, 10));
}

/**
 * 获取用户最后的学习时间
 *
 * @return array
 */
public function getLastLearnTimesArray()
{
    $times = array_unique(explode(',', $this->last_learn_times));
    rsort($times);

    return array_values($times);
}
```

### 常用 Function
```php
/**
 * 时间戳转日期
 *
 * @param $time
 * @param string $format
 * @return string
 */
public static function timeToStr($time, $format = 'Y-m-d H:i:s')
{
    return $time <= 0 ? '' : date($format, $time);
}

/**
 * 中文截取字符串
 *
 * @param   str string [必选] 需要截取的字符串;
 * @param   length int [必须] 截取字符的长度,按照一个汉字的长度算作一个字符;
 * @param   start string [可选] 从那里开始截取;
 * @param   suffix string [可选] 截取字符后加上的后缀,默认为@...;
 * @param   charset enum('gbk','utf-8') [可选] 字符的编码,默认为@utf-8;
 * @return string;
 */
public static function msubstr($str, $start = 0, $length = null, $suffix = '...', $charset = 'utf-8')
{
    $length = null === $length ? strlen($length) : $length;

    switch($charset){
        case 'utf-8':
            $charLen = 3;
            break;
        case 'UTF8':
            $charLen = 3;
            break;
        default:
            $charLen = 2;
    }
    // 小于指定长度，直接返回
    if (strlen($str) <= ($length * $charLen)){
        return $str;
    }elseif (function_exists('mb_substr')){
        $slice = mb_substr($str, $start, $length, $charset);
    }elseif (function_exists('iconv_substr')){
        $slice = iconv_substr($str, $start, $length, $charset);
    }else{
        $re['utf-8'] = "/[\x01-\x7f]|[\xc2-\xdf][\x80-\xbf]|[\xe0-\xef][\x80-\xbf]{2}|[\xf0-\xff][\x80-\xbf]{3}/";
        $re['gb2312'] = "/[\x01-\x7f]|[\xb0-\xf7][\xa0-\xfe]/";
        $re['gbk'] = "/[\x01-\x7f]|[\x81-\xfe][\x40-\xfe]/";
        $re['big5'] = "/[\x01-\x7f]|[\x81-\xfe]([\x40-\x7e]|\xa1-\xfe])/";
        preg_match_all($re[$charset], $str, $match);
        $slice = join("", array_slice($match[0], $start, $length));
    }

    return $slice . $suffix;
}

/**
 * 去除全角空白符
 *
 * @param string $str
 * @return string
 */
public static function trimBlankCharacter($str)
{
    return preg_replace('/(^[\s\x{3000}]*)|([\s\x{3000}]*$)/u', '', strval($str));
}

```
