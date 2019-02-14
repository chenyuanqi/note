
### 常用函数
- call_user_func 和 call_user_func_array
> mixed call_user_func ( callable $callback [, mixed $parameter [, mixed $... ]] )  
> 返回回调函数的返回值，如果错误则返回 FALSE。  
> 值得注意的是：传入 call_user_func 函数的参数不能为引用传递。  
> 
> mixed call_user_func_array ( callable $callback , array $param_arr )
> 要被传入回调函数的数组，这个数组得是索引数组。  
> 同样，返回回调函数的返回值，如果错误则返回 FALSE。  
> 注意，在函数中注册有多个回调内容时(如使用 call_user_func() 与 call_user_func_array())，如在前一个回调中有未捕获的异常，其后的将不再被调用。  
> 它们的区别：call_user_func 不支持引用参数，call_user_func_array 支持引用参数；传入参数的方式不同。  
```php
// call_user_func
function sayName($name){
    echo "My name is $name";
}

call_user_func('sayName', 'Tom');

class Test {
    static function sayHello($name){
        echo "Hello, my name is $name";
    }
}

$class_name = 'Test';
call_user_func(array($class_name, 'sayHello'), 'Tom');
call_user_func($class_name. '::sayHello', 'Tom');

$obj = new Test();
call_user_func(array($obj, 'sayHello'), 'Tom');

// call_user_func_array
function foobar($arg1, $arg2){
    echo __FUNCTION__ . " got $arg1 and $arg2\n";
}
class Foo {
    function bar($arg1, $arg2){
        echo __METHOD__ . " got $arg1 and $arg2\n";
    }
}

// 调用foobar函数并传入两个参数
call_user_func_array("foobar", array("one", "two"));

// 调用$foo->bar()方法并传入两个参数
$foo = new foo;
call_user_func_array(array($foo, "bar"), array("three", "four"));

function increment(&$var){
    $var++;
}

$a = 0;
call_user_func('increment', $a);
echo $a."\n";//0

call_user_func_array('increment', array(&$a));
echo $a."\n";//1
```


