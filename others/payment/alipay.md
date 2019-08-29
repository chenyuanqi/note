
### 支付宝支付
[支付宝支付文档](https://openhome.alipay.com/developmentDocument.htm)  


### 支付宝支付之 APP 支付
APP 支付步骤为：  
1、服务端获取支付宝的配置信息。  
> 支付时需要的配置信息有：  
> key：交易安全校验码  
> app_id：支付宝分配给开发者的应用 ID  

2、服务端生成订单信息。  
> 支付宝只需要知道的订单信息为：  
> subject：必填，商品的标题 / 交易标题 / 订单标题 / 订单关键字等  
> total_amount：必填，订单价格  
> out_trade_no：必填，商户唯一订单号  
> body：非必填，交易的具体描述信息  

3、服务端根据订单信息生成待校验数据。  
[APP 支付的详细请求参数](https://docs.open.alipay.com/204/105465)  

4、服务端生成请求给支付宝的加密字符串。  
```php
$sign = $alipaySubmit->buildRequestParaForApp($param_token);

// buildRequestParaForApp 的实现
// 对待签名参数数组排序
function argSort($para) {
    ksort($para);
    reset($para);

    return $para;
}
// 生成签名结果（推荐 RSA2，这里使用 RSA）
function rsaSign($data, $private_key_path) {
    $priKey = file_get_contents($private_key_path);
    $res = openssl_get_privatekey($priKey);
    openssl_sign($data, $sign, $res);
    openssl_free_key($res);
    //base64 编码
    $sign = base64_encode($sign);

    return $sign;
}
```

5、服务端将待校验数据和加密字符串拼接，返回给 APP。  
```php
$param_token['sign'] = urlencode($sign);

return http_build_query($param_token);
```

6、APP 将得到的数据请求支付宝客户端进行支付。

### 支付宝支付之 Web 支付
网页版支付步骤为：  
1、设置支付宝的配置信息。  
```php
    /* 调用授权接口 alipay.wap.trade.create.direct 获取授权码 token */
        
    //返回格式
    private  $format = "";
    //必填，不需要修改
    
    //版本
    private $v = "";
    //必填，不需要修改
    
    //请求号
    private $req_id = "";
    //必填，须保证每次请求都是唯一
    
    //**req_data详细信息**
    
    //服务器异步通知页面路径
    private $notify_url = "";
    //需http://格式的完整路径，不允许加?id=123这类自定义参数
    
    //页面跳转同步通知页面路径
    private $call_back_url = "";
    //需http://格式的完整路径，不允许加?id=123这类自定义参数
    
    //卖家支付宝账户
    private $seller_email = "";
    //必填
    
    //商户订单号
    private $out_trade_no = "";
    //商户网站订单系统中唯一订单号，必填
    
    //订单名称
    private $subject = "";
    //必填
    
    //付款金额
    private $total_fee = "";
    //必填
    
    //请求业务参数详细
    private $req_data = "";
    //必填
    
    //配置
    private $alipay_config = array();
```

2、向支付宝申请新订单，获取支付 token。  
```php
// 构造参数
$para_token = [
    "service" => "alipay.wap.trade.create.direct",
    //  合作者身份(partner ID)
    "partner" => trim($this->alipay_config['partner']),
    //  APP使用的是RSA，网页版使用的是MD5
    "sec_id" => trim($this->alipay_config['sign_type']),
    //  返回的数据格式
    "format"    => $this->format,
    //  版本号？
    "v" => $this->v,
    //  唯一的请求号
    "req_id"    => $this->req_id,
    //  请求参数
    "req_data"  => $req_data,
    //  字符集，一般为utf8即可。
    "_input_charset"    => trim(strtolower($this->alipay_config['input_charset']))
];

$para_filter = paraFilter($para_token);
$para_sort = argSort($para_filter);
$mysign = $this->buildRequestMysign($para_sort);
// 签名结果与签名方式加入请求提交参数组中
$para_sort['sign'] = $mysign;
// 过滤值为空的数据，过滤签名类型和签名
function paraFilter($para) {
    $para_filter = array();
    while (list ($key, $val) = each ($para)) {
        if($key == "sign" || $key == "sign_type" || $val == "")continue;
        else    $para_filter[$key] = $para[$key];
    }

    return $para_filter;
}
// 字典排序
function argSort($para) {
    ksort($para);
    reset($para);

    return $para;
}
// 签名
function buildRequestMysign($para_sort) {
    //把数组所有元素，按照“参数=参数值”的模式用“&”字符拼接成字符串
    $prestr = createLinkstring($para_sort);
    $mysign = "";
    switch (strtoupper(trim($this->alipay_config['sign_type']))) {
        case "MD5" :
            //  MD5直接将密钥拼接在字符串后面再进行MD5加密。
            $mysign = md5Sign($prestr, $this->alipay_config['key']);
            break;
        case "RSA" :
            //  RSA则是先读取商户的私钥，再用该密钥对字符串进行加密。
            $mysign = rsaSign($prestr, $this->alipay_config['private_key_path']);
            break;
        case "0001" :
            $mysign = rsaSign($prestr, $this->alipay_config['private_key_path']);
            break;
        default :
            $mysign = "";
    }
    
    return $mysign;
}

// 用构造好的参数请求支付宝后台申请新订单
$sResult = getHttpResponsePOST($this->alipay_gateway_new, $this->alipay_config['cacert'],$request_data,trim(strtolower($this->alipay_config['input_charset'])));
function getHttpResponsePOST($url, $cacert_url, $para, $input_charset = '') {

    if (trim($input_charset) != '') {
        $url = $url."_input_charset=".$input_charset;
    }
    $curl = curl_init($url);
    curl_setopt($curl, CURLOPT_SSL_VERIFYPEER, true);//SSL证书认证
    curl_setopt($curl, CURLOPT_SSL_VERIFYHOST, 2);//严格认证
    curl_setopt($curl, CURLOPT_CAINFO,$cacert_url);//证书地址
    curl_setopt($curl, CURLOPT_HEADER, 0 ); // 过滤HTTP头
    curl_setopt($curl,CURLOPT_RETURNTRANSFER, 1);// 显示输出结果
    curl_setopt($curl,CURLOPT_POST,true); // post传输数据
    curl_setopt($curl,CURLOPT_POSTFIELDS,$para);// post传输数据
    $responseText = curl_exec($curl);
    // 如果执行curl过程中出现异常，可打开此开关，以便查看异常内容
    curl_close($curl);
    
    return $responseText;
}
// 处理支付宝返回的数据，并获取 token
// URLDECODE 返回的信息
$html_text = urldecode($html_text);
// 解析远程模拟提交后返回的信息
$para_html_text = parseResponse($html_text);
// 获取 request_token
$request_token = $para_html_text['request_token'];
function parseResponse($str_text) {
    //以“&”字符切割字符串
    $para_split = explode('&',$str_text);
    //把切割后的字符串数组变成变量与数值组合的数组
    foreach ($para_split as $item) {
        //获得第一个=字符的位置
        $nPos = strpos($item,'=');
        //获得字符串长度
        $nLen = strlen($item);
        //获得变量名
        $key = substr($item,0,$nPos);
        //获得数值
        $value = substr($item,$nPos+1,$nLen-$nPos-1);
        //放入数组中
        $para_text[$key] = $value;
    }
    
    if( ! empty ($para_text['res_data'])) {
        //解析加密部分字符串
        if($this->alipay_config['sign_type'] == '0001') {
            $para_text['res_data'] = rsaDecrypt($para_text['res_data'], $this->alipay_config['private_key_path']);
        }
        
        //token从res_data中解析出来（也就是说res_data中已经包含token的内容）
        $doc = new DOMDocument();
        $doc->loadXML($para_text['res_data']);
        $para_text['request_token'] = $doc->getElementsByTagName( "request_token" )->item(0)->nodeValue;
    }
    
    return $para_text;
}
```

3、携带 token 进行订单支付。
```php
// 构造请求数据
$req_data = '<auth_and_execute_req><request_token>' . $request_token . '</request_token></auth_and_execute_req>';
//必填

//构造要请求的参数数组，无需改动
$parameter = array(
    "service" => "alipay.wap.auth.authAndExecute",
    //  合作者身份(partner ID)
    "partner" => trim($this->alipay_config['partner']),
    //  签名类型
    "sec_id" => trim($this->alipay_config['sign_type']),
    //  和步骤2一致
    "format"    => $this->format,
    "v" => $this->v,
    "req_id"    => $this->req_id,
    //  业务详细参数
    "req_data"  => $req_data,
    //  字符集，一般为utf8.
    "_input_charset"    => trim(strtolower($this->alipay_config['input_charset']))
);
```
将这些参数，在页面中传送给支付宝即可发起一次支付请求。  
在 PHP 中的实现就是将这些参数，渲染至 HTML 中，再将 HTML 中的表单提交即可。  


