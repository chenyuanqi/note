
### PHP 的征程
PHP 走过了漫长的道路，成长为处理 web 的最卓越的语言。  
> PHP 发行版可以在「[PHP 博物馆](http://museum.php.net/)」找到。  

### PHP 7.4 新特性
有许多新的 PHP 特性会减少内存的使用并且大大提升 PHP 7.4 的性能。你将能够避免此编程语言之前的某些限制，编写更加简洁的代码，并更快地创建 web 解决方案。

1、短闭包：箭头函数的支持  
> 由于匿名函数或闭包主要应用于 JS 中，因此，他们在 PHP 中似乎很啰嗦，他们的实现和程序的维护也会更复杂一些。  
> 引入对箭头函数的支持使得 PHP 开发者大大简化他们的代码并且使语法更加简洁。这样，你代码的可读性和简洁性会大大提高。  

```php
$a = [1, 2, 3, 4, 5];

$b = array_map(fn($n) => $n * $n * $n, $a);

print_r($b);
```
关于短闭包的一些注意事项。它可以访问父作用域，不需要 use 关键字；$this 可以像普通的闭包一样使用；短闭包只能包含一行，也就是 return 语句。  

2、类型化属性的支持  
> 虽然之前不可能将声明方法用于类变量和属性（包括静态属性），但现在程序员能很轻松地进行编码，而无需创建特定的 getter 和 setter 方法。  
> 由于声明类型（不包括 void 和 callable），你可以使用可为空（Nullable）类型，即 int、float、array、string、object、iterable、self、bool 和 parent。

3、预加载  
> 预加载是在 OPcache 中加载文件、框架和库的过程，绝对是新版本的最佳补充。如果你使用框架，则必须为每个请求下载并重新编译其文件。  
> 在配置 OPcache 的时候，这些代码文件首次参与请求处理，然后每次都检查它们的更改。预加载使服务器可以将指定的代码文件加载到共享内存中。请务必注意，它们将始终可用于后续所有的请求，而无需检查其他文件的改变。  
> 在配置 OPcache 的时候，这些代码文件首次参与请求处理，然后每次都检查它们的更改。预加载使服务器可以将指定的代码文件加载到共享内存中。请务必注意，它们将始终可用于后续所有的请求，而无需检查其他文件的改变。  

4、协变量返回和协变量参数  
> 目前，PHP 中大多数是不变的参数类型和不变的返回类型，这带来了一些约束。  
> 随着协变量（类型从更具体到更通用）返回和协变量（类型从更通用到更具体）参数的引入，PHP 开发者们将能够将参数类型更改为超类型之一。

5、弱引用  
> 弱引用类（WeakReference class）允许 web 开发者们将链接保存到不阻止其销毁的对象中。  
> 请勿将弱引用类和弱引用扩展混淆。  
> 由于这些特性，它们更容易实现类似缓存的结构。  

```php
$obj = new stdClass;

$weakref = WeakReference::create($obj);

var_dump($weakref->get());

unset($obj);

var_dump($weakref->get());
```

6、合并分配运算符  
> 当你需要将三元运算符和 isset 方法一起使用时非常有用。  
> 如果它存在且不为空，那么就会返回第一个操作数，否则就会返回第二个操作数。  

```php
// 获取 $_GET['user'] 的值，如果它不存在则返回 nobody
$username = $_GET['user'] ?? 'nobody';
// 这等价于：
// $username = isset($_GET['user']) ? $_GET['user'] : 'nobody';

// 链式合并：将返回 $_GET['user']、$_POST['user'] 以及 noboody 中第一个不为 NULL 的值
$username = $_GET['user'] ?? $_POST['user'] ?? 'nobody';

$data['date'] ??= new DateTime();
// 这等价于：  
// $data['date'] = $data['date'] ?? new DateTime();
```

7、数组表达式中的展开运算符  
> 在数组中使用展开运算符。  
> 首先，展开运算符被认为是一种语言结构，而 array_merge 是一个函数；其次是针对常量数组 “编译时” 的优化。  
> 同样，展开运算符也有可能展开同一数组多次。  
> 此外，由于可以在扩展运算符的前后添加普通元素，因此 PHP 开发人员将能够在数组中使用其语法。  

```php
$parts = ['apple', 'pear'];
$fruits = ['banana', 'orange', ...$parts, 'watermelon'];
var_dump($fruits);
```

8、新的自定义对象序列化机制  
> 有两种新的可用魔术方法 \_\_serialize 和 \_\_unserialize。  
> 将 Serializable 接口的多功能性与实现 sleep 和 wakeup 方法结合起来，这种序列化机制使得 PHP 开发者可以避免与已存在的方法产生一些自定义的问题。  

9、为引用提供的反射  
> 类似于 symfony/var-dumper 之类的库，严重依赖 Reflection API 来准确罗列变量。  
> 原来，对于引用反射没有很好的支持，这迫使这些库只能依靠 hack 的方式来检测引用。在 PHP 7.4 中添加了 ReflectionReference 类来解决此问题。

10、支持从 \_\_toString () 方法抛出异常  
> 之前无法从 \_\_toString 方法中抛出异常。原因是标准库中的许多函数都执行从对象到字符串的转化，它们当中并非所有的都准备好正确的 “处理” 异常。  
> 作为该 RFC 的一部分，对代码库中的字符串转换进行了全面的审核，并取消了此限制。  

11、外部函数接口  
> 外部函数接口，简称 FFI，允许从用户区调用 C 代码。这意味着 PHP 扩展可以用纯 PHP 编写。  

12、连接优先级  
> 当在遇到没有圆括号包含的 '+' 或 ' - ' 表达式之前有 '.' 的时候，PHP 7.4 会提示弃用警告。  

```php
echo "sum: " . $a + $b;
// PHP 之前会像这样编译它： echo ("sum: " . $a) + $b;
// PHP 8 将使它如下编译： echo "sum :" . ($a + $b);
```

13、添加 mb_str_split 函数  
> 此函数提供与 str_split 多字节字符串相同的功能。  

14、永久支持 ext-hash 扩展；默认不启用 PEAR；弃用ext/wwdx；PHP 短标签被弃用；左关联三元运算符被弃用

### PHP 7.3 新特性
1、灵活的 Heredoc 和 Nowdoc 语法  
> 闭合标识符前支持缩进  
> 闭合标识符后不再强制换行 

```php
$query = <<<SQL

   SELECT *

   FROM `table`

   WHERE `column` = true;

   SQL;
```

2、函数调用时允许尾随逗号  
```php
use Foo\Bar\{
   Foo,
   Bar,
};

unset(
   $foo,
   $bar,
   $baz,
);
```

3、JSON_THROW_ON_ERROR  
> 解析 JSON 响应数据，有 json_encode() 以及 json_decode() 两个函数可供使用。不幸的是，它们都没有恰当的错误抛出表现。  
> json_encode 失败时仅会返回 false；json_decode 失败时则会返回 null，而 null 可作为合法的 JSON 数值。唯一获取错误的方法是，调用 json_last_error() 或 json_last_error_msg()，它们将分别返回机器可读和人类可读的全局错误状态。  
> 为 JSON 函数新增 JSON_THROW_ON_ERROR 常量用于忽略全局错误状态。当错误发生时，JSON 函数将会抛出 JsonException 异常，异常消息（message）为 json_last_error() 的返回值，异常代码（code）为 json_last_error_msg() 的返回值。  

```php
json_encode($data, JSON_THROW_ON_ERROR);
json_decode("invalid json", null, 512, JSON_THROW_ON_ERROR);
// 抛出 JsonException 异常
```

4、PCRE2 迁移  
> PCRE 作为正则表达式引擎，从 PHP 7.3 开始，PCRE2 将作为新的正则引擎大显身手。  
> PCRE2 严格要求，同时支持更多特性：  
> 相对后向引用 \g{+2}（等效于已存在的 \g{-2}）  
> PCRE2 版本检查 (?(VERSION>=x)...)  
> (\*NOTEMPTY) 和 (\*NOTEMPTY_ATSTART) 告知引擎勿返回空匹配  
> (\*NO_JIT) 禁用 JIT 优化  
> (\*LIMIT_HEAP=d) 限制堆大小为 d KB  
> (\*LIMIT_DEPTH=d) 设置回溯深度限制为 d  
> (\*LIMIT_MATCH=d) 设置匹配数量限制为 d  

5、list () 分配参考  
>  list() 可以赋值给引用

```php
$array = [1, 2];
list($a, &$b) = $array;
// 相当于
$array = [1, 2];
$a = $array[0];
$b =& $array[1];
```

6、is_countable 函数  
> 在 PHP 7.2 中，用 count () 获取对象和数组的数量。如果对象不可数，PHP 会抛出警告。  
> 新函数 is_countable() 对数组类型或者实现了 Countable 接口的实例的变量返回 true。

7、array_key_first(), array_key_last()  
> PHP 允许使用 reset() ，end() 和 key() 等方法，通过改变数组的内部指针来获取数组首尾的键和值。  
> 现在，为了避免这种内部干扰，PHP 7.3 推出了新的函数来解决这个问题：  
> array_key_first($array); 获取数组第一个元素的键名  
> array_key_last($array); 获取数组最后一个元素的键名  

8、Argon2 密码哈希增强功能  

9、弃用和删除 image2wbmp()  
> image2wbmp() 函数能够将图像输出为 WBMP 格式。  
> 另一个名为 imagewbmp() 的函数也同样具备单色转换的作用。出于重复原因，image2wbmp() 现已被废弃，你可使用 imagewbmp() 代替它。此函数被弃用后，再次调用它将会触发已弃用警告。待后续此函数被移除后，再次调用它将会触发致命错误。

10、弃用和删除不区分大小写的常量  

11、相同站点 Cookie  
> 建议在使用 cookies 时，增加同站点标志。

12、FPM 更新  
> FastCGI 进程管理器也进行了更新，现在提供了新的方式来记录 FPM 日志。  
> log_limit: 设置允许的日志长度，可以超过 1024 字符。  
> log_buffering: 允许不需要额外缓冲去操作日志。  
> decorate_workers_output: 当启用了 catch_workers_output 时，系统会去禁用渲染输出。

13、改进 Windows 下的文件删除
> 默认情况下，文件描述符以共享读、写、删除的方式去操作。这很有效的去映射 POSIX 并允许去删除正在使用中的文件。但这并不是 100% 都是一样的，不同的平台可能仍存在一些差异。删除操作之后，文件目录仍存在直到所有的文件操作被关闭。  

### PHP 7.2 新特性
1、新增「[object 对象类型](http://php.net/manual/zh/language.types.object.php)」，引进了可用于逆变参数输入和协变返回任何对象类型  
```php
function foo(object $obj) : object
{
    return new SplQueue();
}
```
2、支持通过「[dl()](http://php.net/manual/zh/function.dl.php)」函数实现用名称加载扩展  
3、支持继承抽象类后的抽象类重写抽象方法  
4、Argon2 算法加入生成密码散列 API  
5、新增 PDO 字符串扩展类型（国际化的字符集）  
6、优化「[PDOStatement::debugDumpParams()](http://php.net/manual/zh/pdostatement.debugdumpparams.php)」，增加额外的模拟调试信息  
7、LDAP 扩展支持 EXOP  
8、sockets 扩展增加地址信息函数「[socket_addrinfo_lookup()](http://php.net/manual/zh/function.socket-addrinfo-lookup.php)」、「[socket_addrinfo_connect()](http://php.net/manual/zh/function.socket-addrinfo-connect.php)」、「[socket_addrinfo_bind()](http://php.net/manual/zh/function.socket-addrinfo-bind.php)」、「[socket_addrinfo_explain()](http://php.net/manual/zh/function.socket-addrinfo-explain.php)」  
9、扩展参数类型，重写方法和接口实现的参数类型可以省略  
```php
interface A
{
    public function Test(array $input);
}

class B implements A
{
    public function Test($input){}
}
```
10、允许分组命名空间的尾部逗号  
```php
use Foo\Bar\{
    Foo,
    Bar,
    Baz,
};
```

### PHP 7.1 新特性
1、新增「可为空（Nullable）类型」，参数、返回值的类型允许为空  
```php
function language(): ?string
{
    return 'php';
}
```
2、新增「Void 函数」，返回值类型要么省去 return 语句，要么使用一个空的 return 语句   
```php
function swap(&$left, &$right) : void
{
    if ($left === $right) {
        return;
    }

    $tmp = $left;
    $left = $right;
    $right = $tmp;
}
```
3、支持短数组语法（[]）作为「[list()](http://php.net/manual/zh/function.list.php)」语法的一个备选项，可用于将数组的值赋给一些变量（包括在foreach中）  
```php
// 如下程式等价于 list($id, $name) = $data[0];
[$id, $name] = $data[0];

foreach ($data as list($id, $name)) {
    echo $data;
}
```
4、支持设置类常量的可见性  
```php
class Demo
{
    private const PRIVATE_CONST = 1;
}
```
5、新增「iterable 伪类」，可接受数组或者实现了 Traversable 接口的对象  
6、支持多异常捕获  
```php
try {
    // ...
} catch (FirstException | SecondException $e) {
    // ...
}
```
7、支持「[list()](http://php.net/manual/zh/function.list.php)」使用键名  
```php
list("id" => $id, "name" => $name) = $data[0];
// 如上程式等价于
["id" => $id, "name" => $name] = $data[0];
```
8、支持为负的字符串偏移量  
```php
// 使用 [] 或 {} 操作字符串下标  
echo "abcdef"[-2];
echo strpos("aabbcc", "b", -3);
```
9、ext/openssl 支持 AEAD，通过给「[openssl_encrypt()](http://php.net/manual/zh/function.openssl-encrypt.php)」和「[openssl_decrypt()](http://php.net/manual/zh/function.openssl-decrypt.php)」加解密函数添加额外参数实现  
10、新增「[pcntl_async_signals()](http://php.net/manual/zh/function.pcntl-async-signals.php)」，用于启用无需 ticks （这会带来很多额外的开销）的异步信号处理  
11、支持 curl 扩展使用 HTTP2，通过「[curl_multi_setopt()](http://php.net/manual/zh/function.curl-multi-setopt.php)」函数与新的常量 CURLMOPT_PUSHFUNCTION 来进行调节  

### PHP 7.0 新特性
1、支持标量「[类型声明](http://php.net/manual/zh/functions.arguments.php#functions.arguments.type-declaration)」  
```php
function sumOfInts(int ...$ints)
{
    return array_sum($ints);
}
```
2、支持「[返回值类型声明](http://php.net/manual/zh/functions.returning-values.php#functions.returning-values.type-declaration)」  
```php
function arraysSum(array ...$arrays): array
{
    return array_map(function(array $array): int {
        return array_sum($array);
    }, $arrays);
}
```
3、新增「[?? 语法糖](http://php.net/manual/zh/functions.arguments.php#functions.variable-arg-list)」，也叫 null 合并运算符  
```php
// 如果变量存在且值不为 NULL， 它就会返回自身的值，否则返回它的第二个操作数
// 如下程式等价于 $username = isset($_GET['user']) ? $_GET['user'] : 'nobody'; 
$username = $_GET['user'] ?? 'nobody'; 
```
4、新增太空船操作符（组合比较符）<=>，比较规则沿用「[常规比较规则](http://php.net/manual/zh/types.comparisons.php)」  
```php
// 当 $a 小于、等于或大于 $b 时它分别返回 -1、0 或 1
echo $a <=> $b;
```
5、支持通过 define() 定义常量数组  
6、支持通过new class 来实例化一个「[匿名类](http://php.net/manual/zh/language.oop5.anonymous.php)」  
7、支持 Unicode codepoint 转译语法  
8、优化 Closure::call()  
9、为「[unserialize()](http://php.net/manual/zh/function.unserialize.php)」提供过滤  
10、新增「[IntlChar 类](http://php.net/manual/zh/class.intlchar.php)」，暴露出更多的 ICU 功能  
11、优化「[assert 断言](http://php.net/manual/zh/class.intlchar.php)」，提供断言失败时抛出特定异常的预期能力  
12、支持「[use](http://php.net/manual/zh/language.namespaces.importing.php)」分组  
```php
use \namespace\{ClassA, ClassB, ClassC as C};
```
14、支持「[生成器](http://php.net/manual/zh/language.generators.php)」返回表达式  
15、支持「[生成器](http://php.net/manual/zh/language.generators.php)」委派其他生成器「[yield from](http://php.net/manual/zh/language.generators.syntax.php#control-structures.yield.from)」  
16、新增整数除法函数「[intdiv()](http://php.net/manual/zh/function.intdiv.php)」  
```php
// 如下程式返回 3
echo intdiv(10, 3); 
```
17、支持「[session_start()](http://php.net/manual/zh/language.generators.php)」接受 array 参数，覆盖配置文件中的会话选项  
18、新增「[preg_replace_callback_array()](http://php.net/manual/zh/function.preg-replace-callback-array.php)」多正则检索或替换回调  
19、新增高安全级别的随机字符串生成函数「[random_bytes()](http://php.net/manual/zh/function.random-bytes.php)」、高安全级别的随机整数生成函数「[random_int()](http://php.net/manual/zh/function.random-int.php)」  
20、支持「[list()](http://php.net/manual/zh/function.list.php)」函数来展开实现了 ArrayAccess 接口的对象  
21、支持克隆表达式上访问对象成员  
```php
(clone $foo)->bar();
```

### PHP 5.6 新特性
1、支持常量使用表达式，如 const THREE = TWO + 1;  
2、新增「[... 运算符](http://php.net/manual/zh/functions.arguments.php#functions.variable-arg-list)」，定义函数的可变长参数  
3、支持「[... 运算符](http://php.net/manual/zh/functions.arguments.php#functions.variable-arg-list)」展开参数，如 $operators = [2, 3]; add(1, ...$operators);  
4、新增「[** 幂运算](http://php.net/manual/zh/language.operators.arithmetic.php#115689)」，相当于 pow 函数  
5、支持「[使用命名空间](http://php.net/manual/zh/language.namespaces.importing.php)」use function 及 use const  
6、实现「[交互式调试器 PHPDBG](https://phpdbg.room11.org/)」  
7、修改「[默认编码 default_charset](http://php.net/manual/zh/ini.core.php#ini.default-charset)」默认值为 UTF-8  
8、支持「[php://input](http://php.net/manual/zh/wrappers.php.php#wrappers.php.input)」重用  
9、支持大于 2GB 的文件上传  
10、「[GMP](http://php.net/manual/zh/book.gmp.php)」支持运算符重载  
11、新增「[hash_equals](http://php.net/manual/zh/function.hash-equals.php)」函数，防止时序攻击的字符串比较  
12、新增「[魔术方法 __debugInfo](http://php.net/manual/zh/language.oop5.magic.php#language.oop5.magic.debuginfo)」，控制「[var_dump](http://php.net/manual/zh/function.var-dump.php)」 输出对象的属性和值  
13、支持 gost-crypto 散列算法  
14、提升 SSL/TLS 的支持  
15、「[pgsql](http://php.net/manual/zh/book.pgsql.php)」异步支持  

### PHP 5.5 新特性
1、新增「[生成器  Generator](http://php.net/manual/zh/language.generators.overview.php)」  
2、新增「[异常处理](http://php.net/manual/zh/language.exceptions.php)」关键字 finally  
3、「[foreach](http://php.net/manual/zh/control-structures.foreach.php)」支持「[list](http://php.net/manual/zh/function.list.php)」  
4、「[empty](http://php.net/manual/zh/function.empty.php)」方法支持任意表达式  
5、新增「[新的密码哈希 API](http://php.net/manual/zh/book.password.php)」  
6、支持数组、字符串解引用，如 [1, 2, 3][0] 或 '123'[0]  
7、改进图像处理「[GD 库](http://php.net/manual/zh/book.image.php)」,支持翻转、高级裁剪等图像处理  

### PHP 5.4 新特性
1、支持「[性状 Trait](http://php.net/manual/zh/language.oop5.traits.php)」  
2、支持「短数组语法」，如 $arr = [1, 2, 3];  
3、支持返回值为数组的函数进行成员访问解析，例如 foo()[0]  
4、新增「匿名函数」支持 $this  
5、<?= 将总是可用  
6、新增实例化时访问类成员，如 (new Foo)->foo();  
7、支持语法 Class::{expr}()  
8、内置「[Web Server](http://php.net/manual/zh/features.commandline.webserver.php)」  
9、SESSION 扩展可追踪[上传进度](http://php.net/manual/zh/session.upload-progress.php)  

### PHP 5.3 新特性
1、支持「[命名空间](http://php.net/manual/zh/language.namespaces.php)」  
2、支持「[后期静态绑定](http://php.net/manual/zh/language.oop5.late-static-bindings.php)」  
3、支持「[跳标签 goto](http://php.net/manual/zh/control-structures.goto.php)」  
4、支持「[匿名函数](http://php.net/manual/zh/functions.anonymous.php)」  
5、新增「[魔术方法 __callStatic](http://php.net/manual/zh/language.oop5.overloading.php#language.oop5.overloading.methods)」  
调用一个不存在的静态方法时被调用  
6、新增「[魔术方法 __invoke](http://php.net/manual/zh/language.oop5.magic.php#language.oop5.magic.invoke)」  
该魔术方法会在将一个对象作为函数调用时被调用  
7、新增「[Nowdoc 结构](http://php.net/manual/zh/language.types.string.php#language.types.string.syntax.nowdoc)」，类似「[Heredoc 结构](http://php.net/manual/zh/language.types.string.php#language.types.string.syntax.heredoc)」  
Heredoc 以三个左尖括号开始，后面跟一个标识符, 直到一个同样的顶格的标识符(不能缩进) 结束。就像双引号字符串一样，其中可以嵌入变量。  
Nowdoc 的行为像一个单引号字符串，不能在其中嵌入变量，和 Heredoc 唯一的区别就是，三个左尖括号后的标识符要以单引号括起来  
8、支持类外部使用 const 关键词声明「[常量](http://php.net/manual/zh/language.constants.syntax.php)」  
9、支持「[三元运算符](http://php.net/manual/zh/language.operators.comparison.php#language.operators.comparison.ternary)」简写“?:”  
10、支持「[异常](http://php.net/manual/zh/language.exceptions.php)」内嵌  
11、新增循环引用的垃圾回收器（默认开启）  
12、Phar 即 PHP Archive, 起初只是 Pear 中的一个库而已，后来在 PHP5.3 被重新编写成 C 扩展并内置到 PHP 中。  
Phar 用来将多个 .php 脚本打包 (也可以打包其他文件) 成一个 .phar 的压缩文件(通常是 ZIP 格式)。
目的在于模仿 Java 的 .jar, 不对，目的是为了让发布 PHP 应用程序更加方便。同时还提供了数字签名验证等功能。  

.phar 文件可以像 .php 文件一样，被 PHP 引擎解释执行，同时你还可以写出这样的代码来包含 (require) .phar 中的代码：
```php
require("xxoo.phar");
```

13、后期静态绑定  
PHP 的 OPP 机制，具有继承和类似虚函数的功能，例如如下的代码：
```php
class A
{
    public function callFuncTest()
    {
        print $this->funcTest();
    }

    public function funcTest()
    {
        return "A::funcTest";
    }
}

class B extends A
{
    public function funcTest()
    {
        return "B::funcTest";
    }
}

$b = new B;
$b->callFuncTest();
```
输出是：B::funcTest    
可以看到，当在 A 中使用 $this->funcTest() 时，体现了 “虚函数” 的机制，实际调用的是 B::funcTest()。
然而，如果将所有函数都改为静态函数：
```php
class A
{
    static public function callFuncTest()
    {
        echo self::funcTest();
    }

    static public function funcTest()
    {
        return "A::funcTest";
    }
}

class B extends A
{
    static public function funcTest()
    {
        return "B::funcTest";
    }
}

$b = new B;
$b->callFuncTest();
```
情况就没这么乐观了，输出是：A::funcTest  
这是因为 self 的语义本来就是 “当前类”，所以 PHP5.3 给 static 关键字赋予了一个新功能：后期静态绑定。

