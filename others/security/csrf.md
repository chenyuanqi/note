
### 什么是 CSRF 攻击
CSRF 全名是 Cross-site request forgery，即「跨站请求伪造攻击」，是一种对网站的恶意利用，CSRF 比 XSS 更具危险性。

攻击者一般会使用吸引人的图片去引导用户点击进去他设定好的全套，然后你刚登录的A网站没有关闭，这时候攻击者会利用JS事件去模拟用户请求A网站信息，从而就得到了目的。预防措施：为表单提交都加上自己定义好的token然后加密好，后台也一样的规则然后进行对比。

CSRF 攻击最早在 2001 年被发现，由于它的请求是从用户的 IP 地址发起的，因此在服务器上的 web 日志中可能无法检测到是否受到了 CSRF 攻击，正是由于它的这种隐蔽性，很长时间以来都没有被公开的报告出来，直到 2007 年才真正的被人们所重视。  

CSRF 可以盗用受害者的身份，完成受害者在 web 浏览器有权限进行的任何操作，想想吧，能做的事情太多了。比如，以你的名义发送诈骗邮件，用你的账号购买商品，泄露个人隐私...

### CSRF 产生机理
要完成一个 CSRF 攻击，必须具备以下几个条件：
1、受害者已经登录到了目标网站（你的网站）并且没有退出  
2、受害者有意或者无意的访问了攻击者发布的页面或者链接地址  

利用 CSRF 攻击，主要包含两种方式，一种是基于 GET 请求方式的利用，另一种是基于 POST 请求方式的利用。
- GET 请求利用
> 使用 GET 请求方式的利用是最简单的一种利用方式，其隐患的来源主要是由于在开发系统的时候没有按照 HTTP 动词的正确使用方式来使用造成的。对于 GET 请求来说，它所发起的请求应该是只读的，不允许对网站的任何内容进行修改。  
> 但是事实上并不是如此，很多网站在开发的时候，研发人员错误的认为 GET/POST 的使用区别仅仅是在于发送请求的数据是在 Body 中还是在请求地址中，以及请求内容的大小不同。对于一些危险的操作比如删除文章，用户授权等允许使用 GET 方式发送请求，在请求参数中加上文章或者用户的 ID，这样就造成了只要请求地址被调用，数据就会产生修改。  
> 

- POST 请求利用
> 相对于 GET 方式的利用，POST 方式的利用更加复杂一些，难度也大了一些。攻击者需要伪造一个能够自动提交的表单来发送 POST 请求。  
> 只要想办法实现用户访问的时候自动提交表单就可以了。
```
<script>
$(function() {
    $('#CSRF_forCSRFm').trigger('submit');
});
</script>
<form action="http://a.com/xxx" id="CSRF_form" method="post">
    <input name="uid" value="111" type="hidden">
</form>
```

### 预防措施
防范 CSRF 攻击，其实本质就是要求网站能够识别出哪些请求是非正常用户主动发起的。这就要求我们在请求中嵌入一些额外的授权数据，让网站服务器能够区分出这些未授权的请求，比如说在请求参数中添加一个字段，这个字段的值从登录用户的 Cookie 或者页面中获取的（这个字段的值必须对每个用户来说是随机的，不能有规律可循）。攻击者伪造请求的时候是无法获取页面中与登录用户有关的一个随机值或者用户当前 cookie 中的内容的，因此就可以避免这种攻击。  

- Synchronizer token pattern
> 令牌同步模式（Synchronizer token pattern，简称 STP）是在用户请求的页面中的所有表单中嵌入一个 token，在服务端验证这个 token 的技术。token 可以是任意的内容，但是一定要保证无法被攻击者猜测到或者查询到。攻击者在请求中无法使用正确的 token，因此可以判断出未授权的请求。

- Cookie-to-Header Token
> 对于使用 Js 作为主要交互技术的网站，将 CSRF 的 token 写入到 cookie 中，  
> 然后使用 javascript 读取 token 的值，在发送 http 请求的时候将其作为请求的 header，  
> 最后服务器验证请求头中的 token 是否合法。

- 验证码
> 使用验证码可以杜绝 CSRF 攻击，但是这种方式要求每个请求都输入一个验证码，显然没有哪个网站愿意使用这种粗暴的方式，用户体验太差，用户会疯掉的。

### 简单实现 STP
首先在 index.php 中，创建一个表单，在表单中，我们将 session 中存储的 token 放入到隐藏域，这样，表单提交的时候 token 会随表单一起提交  
```
<?php
$token = sha1(uniqid(rand(), true));
$_SESSION['token'] = $token;
?>
<form action="buy.php" method="post">
    <input type="hidden" name="token" value="<?=$token; ?>" />
    ... 表单内容
</form>
```
在服务端校验请求参数的buy.php中，对表单提交过来的 token 与 session 中存储的 token 进行比对，如果一致说明 token 是有效的
```
<?php
if ($_POST['token'] != $_SESSION['token']) {
    // TOKEN无效
    throw new \Exception('Token无效，请求为伪造请求');
}
// TOKEN有效，表单内容处理
```
对于攻击者来说，在伪造请求的时候是无法获取到用户页面中的这个token值的，因此就可以识别出其创建的伪造请求。

### Laravel 框架中的 VerifyCSRFToken 中间件
在 Laravel 框架中，使用了 VerifyCSRFToken 这个中间件来防范 CSRF 攻击。  

在页面的表单中使用 {{ CSRF_field() }} 来生成 token，该函数会在表单中添加一个名为 _token 的隐藏域，该隐藏域的值为 Laravel 生成的 token，Laravel 使用随机生成的 40 个字符作为防范 CSRF 攻击的 token。  
```php
$this->put('_token', Str::random(40));
```
如果请求是 ajax 异步请求，可以在 meta 标签中添加 token
```php
<meta name="CSRF-token" content="{{ CSRF_token() }}">
```
使用jquery作为前端的框架时候，可以通过以下配置将该值添加到所有的异步请求头中
```javascript
$.ajaxSetup({
    headers: {
        'X-CSRF-TOKEN': $('meta[name="CSRF-token"]').attr('content')
    }
});
```
在启用 session 的时候，Laravel 会生成一个名为_token的值存储到 session 中。而使用前面两种方式在页面中加入的 token 就是使用的这一个值。在用户请求到来时，VerifyCSRFToken 中间件会对符合条件的请求进行 CSRF 检查
```php
if (
  $this->isReading($request) ||
  $this->runningUnitTests() ||
  $this->shouldPassThrough($request) ||
  $this->tokensMatch($request)
) {
  return $this->addCookieToResponse($request, $next($request));
}

throw new TokenMismatchException;
```
在 if 语句中有四个条件，只要任何一个条件结果为 true 则任何该请求是合法的，否则就会抛出 TokenMismatchException 异常，告诉用户请求不合法，存在 CSRF 攻击。  

第一个条件 $this->isReading($request) 用来检查请求是否会对数据产生修改
```php
protected function isReading($request)
{
    return in_array($request->method(), ['HEAD', 'GET', 'OPTIONS']);
}
```
这里判断了请求方式，如果是 HEAD，GET，OPTIONS 这三种请求方式则直接放行。你可能会感到疑惑，为什么 GET 请求也要放行呢？这是因为 Laravel 认为这三个请求都是请求查询数据的，如果一个请求是使用 GET 方式，那无论请求多少次，无论请求参数如何，都不应该最数据做任何修改。

第二个条件顾名思义是对单元测试进行放行，第三个是为开发者提供了一个可以对某些请求添加例外的功能，最后一个 $this->tokensMatch($request) 则是真正起作用的一个，它是 Laravel 防范 CSRF 攻击的关键
```php
$sessionToken = $request->session()->token();
$token = $request->input('_token') ?: $request->header('X-CSRF-TOKEN');

if (! $token && $header = $request->header('X-XSRF-TOKEN')) {
  $token = $this->encrypter->decrypt($header);
}

if (! is_string($sessionToken) || ! is_string($token)) {
  return false;
}

return hash_equals($sessionToken, $token);
```
Laravel 会从请求中读取 _token 参数的的值，这个值就是在前面表单中添加的 CSRF_field() 函数生成的。如果请求是异步的，那么会读取 X-CSRF-TOKEN 请求头，从请求头中读取 token 的值。  

最后使用 hash_equals 函数验证请求参数中提供的 token 值和 session 中存储的 token 值是否一致，如果一致则说明请求是合法的。  

你可能注意到，这个检查过程中也会读取一个名为 X-XSRF-TOKEN 的请求头，这个值是为了提供对一些 javascript 框架的支持（比如 Angular），它们会自动的对异步请求中添加该请求头，而该值是从 Cookie 中的 XSRF-TOKEN 中读取的，因此在每个请求结束的时候，Laravel 会发送给客户端一个名为 XSRF-TOKEN 的 Cookie 值
```php
$response->headers->setCookie(
    new Cookie(
        'XSRF-TOKEN', $request->session()->token(), time() + 60 * $config['lifetime'],
        $config['path'], $config['domain'], $config['secure'], false
    )
);
```

### 没有绝对安全的系统
有一个事实是我们无法回避的：没有绝对安全的系统，你有一千种防御对策，攻击者就有一千零一种攻击方式，但不管如何，我们都要尽最大的努力去将攻击者拦截在门外。  

作为一名 web 方向的研发人员，无论你是从事业务逻辑开发还是做单纯的技术研究，了解一些安全方面的知识都是很有必要的，多关注一些安全方向的动态，了解常见的攻击方式以及应对策略，必将在你成长为一名大牛的路上为你 “推波助澜”。



