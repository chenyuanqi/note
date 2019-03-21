
### PHP 的错误
php 脚本自身原因，大部分是由于错误的语法，或者服务器环境造成的 php 脚本无法通过编译，或者无法正常执行的情况。  

注意：
1、并不是所有的错误都会造成脚本无法运行，部分错误发生后脚本继续运行  
2、正常情况下，错误并不能被 try catch 捕捉

错误的级别
1、Parse error 解析错误 (脚本不能继续执行)
```php
// 解析错误举例: 代码未以分号结尾
$num = 0 
//运行结果：Parse error: syntax error, unexpected 'echo' (T_ECHO) 
```
2、Fatal error 致命错误 (脚本不能继续执行)
```php
// 致命错误举例：调用未定义的函数
test();  
//运行结果：Fatal error: Uncaught Error: Call to undefined function test()
```
3、Warning 警告 (给出错误信息，脚本继续运行)
```php
// 警告举例：除以 0
1000 /0; 
// 运行结果：Warning: Division by zero
```
4、Notice 注意 (给出错误信息，脚本继续运行)  
```php
// 注意错误举例：输出未定义的变量
echo $a; 
// 运行结果：Notice: Undefined variable: a
```
5、deprecated：最低级别的错误 (脚本继续运行)，在 php7 以下版本会出现

error 相关函数
1、error_report 设置错误的报告级别  
在工作中我们经常会用到这个函数来设置 php 的错误报错，比如忽略 notice 报告：error_reporting (E_ALL ^ E_NOTICE);  
可用版本：(PHP 4, PHP 5, PHP 7)  
用法：error_reporting ([ int $level ] ) : int  

```php
// 关闭所有PHP错误报告
error_reporting(0);

// Report simple running errors
error_reporting(E_ERROR | E_WARNING | E_PARSE);

// 报告 E_NOTICE也挺好 (报告未初始化的变量
// 或者捕获变量名的错误拼写)
error_reporting(E_ERROR | E_WARNING | E_PARSE | E_NOTICE);

// 除了 E_NOTICE，报告其他所有错误
error_reporting(E_ALL ^ E_NOTICE);

// 报告所有 PHP 错误 (参见 changelog)
error_reporting(E_ALL);
```
2、set_error_handler 设置用户自定义的错误处理函数  
这个函数可以按照你自己定义的方式来处理错误，在后面的专题还会详细的讲解，这里先简单讲解一下  
可用版本：(PHP 4 >= 4.0.1, PHP 5, PHP 7)  
说明：set_error_handler ( callable $error_handler [, int $error_types = E_ALL | E_STRICT ] ) : mixed  
参数解释：  
$error_handler : 我们自己自定义的函数  
$error_types：错误的预定义常量， error_types 里指定的错误类型都会绕过 PHP 标准错误处理程序， 除非回调函数返回了 FALSE  
```php
set_error_handler('myErrorFun'); //把php原始的错误处理机制，变成我们的myErrorFun函数处理
function myErrorFun($errno, $message, $file, $line)
{
    echo '错误码是：'.$errno.'</br>';
    echo '错误的信息是'.$message.'</br>';
    echo '错误的文件是：'.$file.'</br>';
    echo '错误的行数是'.$line;
}

echo $a; //a是未定义的变量

// 打印结果：
// 错误码是：8
// 错误的信息是Undefined variable: a
// 错误的文件是：D:\project\mz_php_server\public\test.php
// 错误的行数是12
```

错误的相关配置（php.ini）：   
1、错误日志是否打开  
log_errors = On （Off）  

2、错误日志记录的位置  
error_log = php_errors.log  

3、是否打开错误显示  
display_errors = Off （Off）  

4、定义错误显示的级别  
error_reporting = E_ALL （错误的预定义常量）  

### PHP 的异常
程序的行为和预期不同，是属于逻辑上错误。什么意思呢，就是按照正常逻辑来说不会出现，但却发生了的。  

和错误的区别：  
1、错误是由于 php 脚本自身的问题导致的和逻辑无关，异常则是由于逻辑问题导致的  
2、在 php7 之前错误不能用 try catch 捕获，异常可以  

系统自定义异常类  
在 php 中，系统给我们内置好了一个 Exception 类，这个类是系统给我们定义好的，我们不需要去再定义异常处理类，可以直接使用。  
```php
// 异常的构造函数
// message：抛出的异常消息内容；code：异常代码；previous：异常链中的前一个异常
public Exception::__construct ([ string $message = "" [, int $code = 0 [, Throwable $previous = NULL ]]] )

$exception = new Exception; //直接new这个类，不需要去定义
var_dump($exception);
//打印结果：
object(Exception)#1 (7) {
  ["message":protected]=>
  string(0) ""
  ["string":"Exception":private]=>
  string(0) ""
  ["code":protected]=>
  int(0)
  ["file":protected]=>
  string(40) "D:\project\mz_php_server\public\test.php"
  ["line":protected]=>
  int(3)
  ["trace":"Exception":private]=>
  array(0) {
  }
  ["previous":"Exception":private]=>
  NULL
}
```

try catch 异常捕获  
```php
try {
    throw new Exception("大江东去，浪淘尽，千古风流人物"); //抛出异常
} catch (Exception $e) { //系统内置的异常处理类
    echo $e->getMessage(); //获取异常信息
}

//定义一个异常处理类，继承系统的异常处理类
class TestException extends Exception{
}
$type = 2;
try {
    if($type == 1){
        throw new Exception("type is 1");

    }elseif($type == 2){
        throw new TestException("type is 2");

    }
} catch (Exception $e) {
    echo $e->getMessage();
} catch (TestException $e) {
    echo $e->getMessage();
}

//try catch 还可以一层一层的向上抛出，和 js 的事件冒泡很像
//定义一个异常处理类，继承系统的异常处理类
class TestException extends Exception{
}

try {
    try {
        throw new TestException("我是测试嵌套的");
    } catch (Exception $e) { //捕获TestException
        throw $e; //把捕获的异常再次抛出
    }
} catch (Exception $e) { //捕获上一层抛出的TestException
    echo $e->getMessage();
}
```

异常相关函数  
```php
//系统提供了一个自定义的异常处理函数，用来处理没有用 try catch 捕获的异常
//exception_handler：当一个未捕获的异常发生时所调用函数的名称。 该处理函数需要接受一个参数，该参数是一个抛出的异常对象
set_error_handler ( callable $error_handler [, int $error_types = E_ALL | E_STRICT ] ) : mixed

set_exception_handler('myexception');
function myexception($exception){
    var_dump($exception->getMessage());
}
throw new Exception("测试一下自定义异常处理函");
```
