
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

### PHP 函数式
PHP 支持一流的函数，意味着函数可以被赋值到一个变量。  
用户自定义以及内建函数都可以被变量引用以及动态调用。函数可以作为参数传递到其他函数中，函数也可以返回其他函数(这个特性被称为高阶函数)。  

递归，这是一个允许函数调用它本身的特性，它在语言中被支持。但是，大多数 PHP 代码都是使用迭代。

匿名函数(以及闭包支持)自从PHP5.3以后出现(2009)，PHP 5.4 增加了绑定闭包作用域到对象的特性，也提高回调的支持，如回调函数几乎在任何情况下都可以与匿名函数互换。  
闭包是一个可以访问从外部作用域引进来的非全局变量。理论上，闭包是一个带有参数的函数。那些参数在定义时被上下文封闭起来，对外部是不可见的。闭包可以用一个很干净利索的方式解决变量作用域的限制。
```php
$input = [1, 2, 3, 4, 5, 6];

// 创建一个新的匿名函数然后赋值给一个变量
$filter_even = function($item) {
    return ($item % 2) == 0;
};

// 内建函数 array_filter 接收数据和函数
$output = array_filter($input, $filter_even);
print_r($output);

// 函数不一定要被赋值到变量，下面这样也是可以的
$output = array_filter($input, function($item) {
    return ($item % 2) == 0;
});
print_r($output);

/**
 * 创建一个匿名过滤函数，只接收大于$min 的数
 *
 * 在“大于n”的过滤器之外返回单个过滤器
 */
function criteria_greater_than($min)
{
    return function($item) use ($min) {
        return $item > $min;
    };
}

// 在 input 数组中使用已选的过滤函数来调用 array_filter 函数
$output = array_filter($input, criteria_greater_than(3));
print_r($output); // 大于3的元素
```