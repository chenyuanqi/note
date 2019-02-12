
# Clean Code PHP

## 目录

  1. [介绍](#介绍)
  2. [变量](#变量)
     * [使用见字知意的变量名](#使用见字知意的变量名)
     * [同一个实体要用相同的变量名](#同一个实体要用相同的变量名)
     * [使用便于搜索的名称 (part 1)](#使用便于搜索的名称-part-1)
     * [使用便于搜索的名称 (part 2)](#使用便于搜索的名称-part-2)
     * [使用自解释型变量](#使用自解释型变量)
     * [避免深层嵌套，尽早返回 (part 1)](#避免深层嵌套尽早返回-part-1)
     * [避免深层嵌套，尽早返回 (part 2)](#避免深层嵌套尽早返回-part-2)
     * [少用无意义的变量名](#少用无意义的变量名)
     * [不要添加不必要上下文](#不要添加不必要上下文)
     * [合理使用参数默认值，没必要在方法里再做默认值检测](#合理使用参数默认值没必要在方法里再做默认值检测)
  3. [表达式](#表达式)
     * [使用恒等式](#使用恒等式)
  4. [函数](#函数)
     * [函数参数（最好少于2个）](#函数参数-最好少于2个)
     * [函数应该只做一件事](#函数应该只做一件事)
     * [函数名应体现他做了什么事](#函数名应体现他做了什么事)
     * [函数里应当只有一层抽象abstraction](#函数里应当只有一层抽象abstraction)
     * [不要用flag作为函数的参数](#不要用flag作为函数的参数)
     * [避免副作用](#避免副作用)
     * [不要写全局函数](#不要写全局函数)
     * [不要使用单例模式](#不要使用单例模式)
     * [封装条件语句](#封装条件语句)
     * [避免用反义条件判断](#避免用反义条件判断)
     * [避免条件判断](#避免条件判断)
     * [避免类型检查 (part 1)](#避免类型检查-part-1)
     * [避免类型检查 (part 2)](#避免类型检查-part-2)
     * [移除僵尸代码](#移除僵尸代码)
  5. [对象和数据结构 Objects and Data Structures](#对象和数据结构)
     * [使用 getters 和 setters Use object encapsulation](#使用-getters-和-setters)
     * [给对象使用私有或受保护的成员变量](#给对象使用私有或受保护的成员变量)
  6. [类](#类)
     * [少用继承多用组合](#少用继承多用组合)
     * [避免连贯接口](#避免连贯接口)
     * [推荐使用 final 类](#推荐使用-final-类)
  7. [类的SOLID原则 SOLID](#solid)
     * [S: 单一职责原则 Single Responsibility Principle (SRP)](#单一职责原则)
     * [O: 开闭原则 Open/Closed Principle (OCP)](#开闭原则)
     * [L: 里氏替换原则 Liskov Substitution Principle (LSP)](#里氏替换原则)
     * [I: 接口隔离原则 Interface Segregation Principle (ISP)](#接口隔离原则)
     * [D: 依赖倒置原则 Dependency Inversion Principle (DIP)](#依赖倒置原则)
  8. [别写重复代码 (DRY)](#别写重复代码-dry)
  9. [翻译](#翻译)

## 介绍


本文参考自 Robert C. Martin的[*Clean Code*](https://www.amazon.com/Clean-Code-Handbook-Software-Craftsmanship/dp/0132350882)  书中的软件工程师的原则
,适用于PHP。 这不是风格指南。 这是一个关于开发可读、可复用并且可重构的PHP软件指南。

并不是这里所有的原则都得遵循，甚至很少的能被普遍接受。 这些虽然只是指导，但是都是*Clean Code*作者多年总结出来的。

本文受到 [clean-code-javascript](https://github.com/ryanmcdermott/clean-code-javascript) 的启发

虽然很多开发者还在使用PHP5，但是本文中的大部分示例的运行环境需要PHP 7.1+。

## 翻译说明

翻译完成度100%，最后更新时间2017-12-25。本文由 php-cpm 基于 [yangweijie版本](https://github.com/yangweijie/clean-code-php) 的[clean-code-php](https://github.com/jupeter/clean-code-php)翻译并同步大量原文内容。

原文更新频率较高，我的翻译方法是直接用文本比较工具逐行对比。优先保证文字内容是最新的，再逐步提升翻译质量。

阅读过程中如果遇到各种链接失效、内容老旧、术语使用错误和其他翻译错误等问题，欢迎大家积极提交PR。

## **变量**

### 使用见字知意的变量名

**坏:**

```php
$ymdstr = $moment->format('y-m-d');
```

**好:**

```php
$currentDate = $moment->format('y-m-d');
```

**[⬆ 返回顶部](#目录)**

### 同一个实体要用相同的变量名

**坏:**

```php
getUserInfo();
getUserData();
getUserRecord();
getUserProfile();
```

**好:**

```php
getUser();
```

**[⬆ 返回顶部](#目录)**

### 使用便于搜索的名称 (part 1)

写代码是用来读的。所以写出可读性高、便于搜索的代码至关重要。
命名变量时如果没有有意义、不好理解，那就是在伤害读者。
请让你的代码便于搜索。

**坏:**
```php
// What the heck is 448 for?
$result = $serializer->serialize($data, 448);
```

**好:**

```php
$json = $serializer->serialize($data, JSON_UNESCAPED_SLASHES | JSON_PRETTY_PRINT | JSON_UNESCAPED_UNICODE);
```

### 使用便于搜索的名称 (part 2)

**坏:**

```php
// What the heck is 4 for?
if ($user->access & 4) {
    // ...
}
```

**好:**

```php
class User
{
    const ACCESS_READ = 1;
    const ACCESS_CREATE = 2;
    const ACCESS_UPDATE = 4;
    const ACCESS_DELETE = 8;
}

if ($user->access & User::ACCESS_UPDATE) {
    // do edit ...
}
```

**[⬆ 返回顶部](#目录)**

### 使用自解释型变量

**坏:**

```php
$address = 'One Infinite Loop, Cupertino 95014';
$cityZipCodeRegex = '/^[^,]+,\s*(.+?)\s*(\d{5})$/';
preg_match($cityZipCodeRegex, $address, $matches);

saveCityZipCode($matches[1], $matches[2]);
```

**不错:**

好一些，但强依赖于正则表达式的熟悉程度

```php
$address = 'One Infinite Loop, Cupertino 95014';
$cityZipCodeRegex = '/^[^,]+,\s*(.+?)\s*(\d{5})$/';
preg_match($cityZipCodeRegex, $address, $matches);

[, $city, $zipCode] = $matches;
saveCityZipCode($city, $zipCode);
```

**好:**

使用带名字的子规则，不用懂正则也能看的懂

```php
$address = 'One Infinite Loop, Cupertino 95014';
$cityZipCodeRegex = '/^[^,]+,\s*(?<city>.+?)\s*(?<zipCode>\d{5})$/';
preg_match($cityZipCodeRegex, $address, $matches);

saveCityZipCode($matches['city'], $matches['zipCode']);
```

**[⬆ 返回顶部](#目录)**

### 避免深层嵌套，尽早返回 (part 1)

太多的if else语句通常会导致你的代码难以阅读，直白优于隐晦

**糟糕:**

```php
function isShopOpen($day): bool
{
    if ($day) {
        if (is_string($day)) {
            $day = strtolower($day);
            if ($day === 'friday') {
                return true;
            } elseif ($day === 'saturday') {
                return true;
            } elseif ($day === 'sunday') {
                return true;
            } else {
                return false;
            }
        } else {
            return false;
        }
    } else {
        return false;
    }
}
```

**好:**

```php
function isShopOpen(string $day): bool
{
    if (empty($day)) {
        return false;
    }

    $openingDays = [
        'friday', 'saturday', 'sunday'
    ];

    return in_array(strtolower($day), $openingDays, true);
}
```

**[⬆ 返回顶部](#目录)**

### 避免深层嵌套，尽早返回 (part 2)

**糟糕的:**

```php
function fibonacci(int $n)
{
    if ($n < 50) {
        if ($n !== 0) {
            if ($n !== 1) {
                return fibonacci($n - 1) + fibonacci($n - 2);
            } else {
                return 1;
            }
        } else {
            return 0;
        }
    } else {
        return 'Not supported';
    }
}
```

**好:**

```php
function fibonacci(int $n): int
{
    if ($n === 0 || $n === 1) {
        return $n;
    }

    if ($n > 50) {
        throw new \Exception('Not supported');
    }

    return fibonacci($n - 1) + fibonacci($n - 2);
}
```

**[⬆ 返回顶部](#目录)**

### 少用无意义的变量名

别让读你的代码的人猜你写的变量是什么意思。
写清楚好过模糊不清。

**坏:**

```php
$l = ['Austin', 'New York', 'San Francisco'];

for ($i = 0; $i < count($l); $i++) {
    $li = $l[$i];
    doStuff();
    doSomeOtherStuff();
    // ...
    // ...
    // ...
  // 等等, `$li` 又代表什么?
    dispatch($li);
}
```

**好:**

```php
$locations = ['Austin', 'New York', 'San Francisco'];

foreach ($locations as $location) {
    doStuff();
    doSomeOtherStuff();
    // ...
    // ...
    // ...
    dispatch($location);
}
```

**[⬆ 返回顶部](#目录)**

### 不要添加不必要上下文

如果从你的类名、对象名已经可以得知一些信息，就别再在变量名里重复。

**坏:**

```php
class Car
{
    public $carMake;
    public $carModel;
    public $carColor;

    //...
}
```

**好:**

```php
class Car
{
    public $make;
    public $model;
    public $color;

    //...
}
```

**[⬆ 返回顶部](#目录)**

### 合理使用参数默认值，没必要在方法里再做默认值检测

**不好:**

不好，`$breweryName` 可能为 `NULL`.

```php
function createMicrobrewery($breweryName = 'Hipster Brew Co.'): void
{
    // ...
}
```

**还行:**

比上一个好理解一些，但最好能控制变量的值

```php
function createMicrobrewery($name = null): void
{
    $breweryName = $name ?: 'Hipster Brew Co.';
    // ...
}
```

**好:**

如果你的程序只支持 PHP 7+, 那你可以用 [type hinting](http://php.net/manual/en/functions.arguments.php#functions.arguments.type-declaration) 保证变量 `$breweryName` 不是 `NULL`.

```php
function createMicrobrewery(string $breweryName = 'Hipster Brew Co.'): void
{
    // ...
}
```

**[⬆ 返回顶部](#目录)**

## 表达式

### [使用恒等式](http://php.net/manual/en/language.operators.comparison.php)

**不好:**

简易对比会将字符串转为整形

```php
$a = '42';
$b = 42;

if( $a != $b ) {
   //这里始终执行不到
}
```

对比 $a != $b 返回了 `FALSE` 但应该返回 `TRUE` !
字符串 '42' 跟整数 42 不相等

**好:**

使用恒等判断检查类型和数据

```php
$a = '42';
$b = 42;

if ($a !== $b) {
    // The expression is verified
}
```

The comparison `$a !== $b` returns `TRUE`.

**[⬆ 返回顶部](#目录)**

## 函数

### 函数参数（最好少于2个）

限制函数参数个数极其重要，这样测试你的函数容易点。有超过3个可选参数参数导致一个爆炸式组合增长，你会有成吨独立参数情形要测试。

无参数是理想情况。1个或2个都可以，最好避免3个。再多就需要加固了。通常如果你的函数有超过两个参数，说明他要处理的事太多了。 如果必须要传入很多数据，建议封装一个高级别对象作为参数。

**坏:**

```php
function createMenu(string $title, string $body, string $buttonText, bool $cancellable): void
{
    // ...
}
```

**好:**

```php
class MenuConfig
{
    public $title;
    public $body;
    public $buttonText;
    public $cancellable = false;
}

$config = new MenuConfig();
$config->title = 'Foo';
$config->body = 'Bar';
$config->buttonText = 'Baz';
$config->cancellable = true;

function createMenu(MenuConfig $config): void
{
    // ...
}
```

**[⬆ 返回顶部](#目录)**

### 函数应该只做一件事

这是迄今为止软件工程里最重要的一个规则。当一个函数做超过一件事的时候，他们就难于实现、测试和理解。当你把一个函数拆分到只剩一个功能时，他们就容易被重构，然后你的代码读起来就更清晰。如果你光遵循这条规则，你就领先于大多数开发者了。

**坏:**

```php
function emailClients(array $clients): void
{
    foreach ($clients as $client) {
        $clientRecord = $db->find($client);
        if ($clientRecord->isActive()) {
            email($client);
        }
    }
}
```

**好:**

```php
function emailClients(array $clients): void
{
    $activeClients = activeClients($clients);
    array_walk($activeClients, 'email');
}

function activeClients(array $clients): array
{
    return array_filter($clients, 'isClientActive');
}

function isClientActive(int $client): bool
{
    $clientRecord = $db->find($client);

    return $clientRecord->isActive();
}
```

**[⬆ 返回顶部](#目录)**

### 函数名应体现他做了什么事

**坏:**

```php
class Email
{
    //...

    public function handle(): void
    {
        mail($this->to, $this->subject, $this->body);
    }
}

$message = new Email(...);
// 啥？handle处理一个消息干嘛了？是往一个文件里写吗？
$message->handle();
```

**好:**

```php
class Email 
{
    //...

    public function send(): void
    {
        mail($this->to, $this->subject, $this->body);
    }
}

$message = new Email(...);
// 简单明了
$message->send();
```

**[⬆ 返回顶部](#目录)**

### 函数里应当只有一层抽象abstraction

当你抽象层次过多时时，函数处理的事情太多了。需要拆分功能来提高可重用性和易用性，以便简化测试。
（译者注：这里从示例代码看应该是指嵌套过多）

**坏:**

```php
function parseBetterJSAlternative(string $code): void
{
    $regexes = [
        // ...
    ];

    $statements = explode(' ', $code);
    $tokens = [];
    foreach ($regexes as $regex) {
        foreach ($statements as $statement) {
            // ...
        }
    }

    $ast = [];
    foreach ($tokens as $token) {
        // lex...
    }

    foreach ($ast as $node) {
        // parse...
    }
}
```

**坏:**

我们把一些方法从循环中提取出来，但是`parseBetterJSAlternative()`方法还是很复杂，而且不利于测试。

```php
function tokenize(string $code): array
{
    $regexes = [
        // ...
    ];

    $statements = explode(' ', $code);
    $tokens = [];
    foreach ($regexes as $regex) {
        foreach ($statements as $statement) {
            $tokens[] = /* ... */;
        }
    }

    return $tokens;
}

function lexer(array $tokens): array
{
    $ast = [];
    foreach ($tokens as $token) {
        $ast[] = /* ... */;
    }

    return $ast;
}

function parseBetterJSAlternative(string $code): void
{
    $tokens = tokenize($code);
    $ast = lexer($tokens);
    foreach ($ast as $node) {
        // 解析逻辑...
    }
}
```

**好:**

最好的解决方案是把 `parseBetterJSAlternative()`方法的依赖移除。

```php
class Tokenizer
{
    public function tokenize(string $code): array
    {
        $regexes = [
            // ...
        ];

        $statements = explode(' ', $code);
        $tokens = [];
        foreach ($regexes as $regex) {
            foreach ($statements as $statement) {
                $tokens[] = /* ... */;
            }
        }

        return $tokens;
    }
}

class Lexer
{
    public function lexify(array $tokens): array
    {
        $ast = [];
        foreach ($tokens as $token) {
            $ast[] = /* ... */;
        }

        return $ast;
    }
}

class BetterJSAlternative
{
    private $tokenizer;
    private $lexer;

    public function __construct(Tokenizer $tokenizer, Lexer $lexer)
    {
        $this->tokenizer = $tokenizer;
        $this->lexer = $lexer;
    }

    public function parse(string $code): void
    {
        $tokens = $this->tokenizer->tokenize($code);
        $ast = $this->lexer->lexify($tokens);
        foreach ($ast as $node) {
            // 解析逻辑...
        }
    }
}
```

这样我们可以对依赖做mock，并测试`BetterJSAlternative::parse()`运行是否符合预期。

**[⬆ 返回顶部](#目录)**

### 不要用flag作为函数的参数

flag就是在告诉大家，这个方法里处理很多事。前面刚说过，一个函数应当只做一件事。 把不同flag的代码拆分到多个函数里。

**坏:**
```php
function createFile(string $name, bool $temp = false): void
{
    if ($temp) {
        touch('./temp/'.$name);
    } else {
        touch($name);
    }
}
```

**好:**

```php
function createFile(string $name): void
{
    touch($name);
}

function createTempFile(string $name): void
{
    touch('./temp/'.$name);
}
```
**[⬆ 返回顶部](#目录)**

### 避免副作用

一个函数做了比获取一个值然后返回另外一个值或值们会产生副作用如果。副作用可能是写入一个文件，修改某些全局变量或者偶然的把你全部的钱给了陌生人。

现在，你的确需要在一个程序或者场合里要有副作用，像之前的例子，你也许需要写一个文件。你想要做的是把你做这些的地方集中起来。不要用几个函数和类来写入一个特定的文件。用一个服务来做它，一个只有一个。

重点是避免常见陷阱比如对象间共享无结构的数据，使用可以写入任何的可变数据类型，不集中处理副作用发生的地方。如果你做了这些你就会比大多数程序员快乐。

**坏:**

```php
// Global variable referenced by following function.
// If we had another function that used this name, now it'd be an array and it could break it.
$name = 'Ryan McDermott';

function splitIntoFirstAndLastName(): void
{
    global $name;

    $name = explode(' ', $name);
}

splitIntoFirstAndLastName();

var_dump($name); // ['Ryan', 'McDermott'];
```

**好:**

```php
function splitIntoFirstAndLastName(string $name): array
{
    return explode(' ', $name);
}

$name = 'Ryan McDermott';
$newName = splitIntoFirstAndLastName($name);

var_dump($name); // 'Ryan McDermott';
var_dump($newName); // ['Ryan', 'McDermott'];
```

**[⬆ 返回顶部](#目录)**

### 不要写全局函数
在大多数语言中污染全局变量是一个坏的实践，因为你可能和其他类库冲突
并且调用你api的人直到他们捕获异常才知道踩坑了。让我们思考一种场景：
如果你想配置一个数组，你可能会写一个全局函数`config()`，但是他可能
和试着做同样事的其他类库冲突。

**坏:**

```php
function config(): array
{
    return  [
        'foo' => 'bar',
    ]
}
```

**好:**

```php
class Configuration
{
    private $configuration = [];

    public function __construct(array $configuration)
    {
        $this->configuration = $configuration;
    }

    public function get(string $key): ?string
    {
        return isset($this->configuration[$key]) ? $this->configuration[$key] : null;
    }
}
```

加载配置并创建 `Configuration` 类的实例

```php
$configuration = new Configuration([
    'foo' => 'bar',
]);
```

现在你必须在程序中用 `Configuration` 的实例了

**[⬆ 返回顶部](#目录)**

### 不要使用单例模式

单例是一种 [反模式](https://en.wikipedia.org/wiki/Singleton_pattern).  以下是解释：Paraphrased from Brian Button:
 1. 总是被用成全局实例。They are generally used as a **global instance**, why is that so bad? Because **you hide the dependencies** of your application in your code, instead of exposing them through the interfaces. Making something global to avoid passing it around is a [code smell](https://en.wikipedia.org/wiki/Code_smell).
 2. 违反了[单一响应原则]()They violate the [single responsibility principle](#single-responsibility-principle-srp): by virtue of the fact that **they control their own creation and lifecycle**.
 3. 导致代码强耦合They inherently cause code to be tightly [coupled](https://en.wikipedia.org/wiki/Coupling_%28computer_programming%29). This makes faking them out under **test rather difficult** in many cases.
 4. 在整个程序的生命周期中始终携带状态。They carry state around for the lifetime of the application. Another hit to testing since **you can end up with a situation where tests need to be ordered** which is a big no for unit tests. Why? Because each unit test should be independent from the other.

这里有一篇非常好的讨论单例模式的[根本问题((http://misko.hevery.com/2008/08/25/root-cause-of-singletons/)的文章，是[Misko Hevery](http://misko.hevery.com/about/) 写的。

**坏:**

```php
class DBConnection
{
    private static $instance;

    private function __construct(string $dsn)
    {
        // ...
    }

    public static function getInstance(): DBConnection
    {
        if (self::$instance === null) {
            self::$instance = new self();
        }

        return self::$instance;
    }

    // ...
}

$singleton = DBConnection::getInstance();
```

**好:**

```php
class DBConnection
{
    public function __construct(string $dsn)
    {
        // ...
    }

     // ...
}
```

创建 `DBConnection` 类的实例并通过 [DSN](http://php.net/manual/en/pdo.construct.php#refsect1-pdo.construct-parameters) 配置.

```php
$connection = new DBConnection($dsn);
```

现在你必须在程序中 使用 `DBConnection` 的实例了

**[⬆ 返回顶部](#目录)**

### 封装条件语句

**坏:**

```php
if ($article->state === 'published') {
    // ...
}
```

**好:**

```php
if ($article->isPublished()) {
    // ...
}
```

**[⬆ 返回顶部](#目录)**

### 避免用反义条件判断

**坏:**

```php
function isDOMNodeNotPresent(\DOMNode $node): bool
{
    // ...
}

if (!isDOMNodeNotPresent($node))
{
    // ...
}
```

**好:**

```php
function isDOMNodePresent(\DOMNode $node): bool
{
    // ...
}

if (isDOMNodePresent($node)) {
    // ...
}
```

**[⬆ 返回顶部](#目录)**

### 避免条件判断

这看起来像一个不可能任务。当人们第一次听到这句话是都会这么说。
"没有`if语句`我还能做啥？" 答案是你可以使用多态来实现多种场景
的相同任务。第二个问题很常见， “这么做可以，但为什么我要这么做？”
 答案是前面我们学过的一个Clean Code原则：一个函数应当只做一件事。
 当你有很多含有`if`语句的类和函数时,你的函数做了不止一件事。
 记住，只做一件事。

**坏:**

```php
class Airplane
{
    // ...

    public function getCruisingAltitude(): int
    {
        switch ($this->type) {
            case '777':
                return $this->getMaxAltitude() - $this->getPassengerCount();
            case 'Air Force One':
                return $this->getMaxAltitude();
            case 'Cessna':
                return $this->getMaxAltitude() - $this->getFuelExpenditure();
        }
    }
}
```

**好:**

```php
interface Airplane
{
    // ...

    public function getCruisingAltitude(): int;
}

class Boeing777 implements Airplane
{
    // ...

    public function getCruisingAltitude(): int
    {
        return $this->getMaxAltitude() - $this->getPassengerCount();
    }
}

class AirForceOne implements Airplane
{
    // ...

    public function getCruisingAltitude(): int
    {
        return $this->getMaxAltitude();
    }
}

class Cessna implements Airplane
{
    // ...

    public function getCruisingAltitude(): int
    {
        return $this->getMaxAltitude() - $this->getFuelExpenditure();
    }
}
```

**[⬆ 返回顶部](#目录)**

### 避免类型检查 (part 1)

PHP是弱类型的,这意味着你的函数可以接收任何类型的参数。
有时候你为这自由所痛苦并且在你的函数渐渐尝试类型检查。
有很多方法去避免这么做。第一种是统一API。

**坏:**

```php
function travelToTexas($vehicle): void
{
    if ($vehicle instanceof Bicycle) {
        $vehicle->pedalTo(new Location('texas'));
    } elseif ($vehicle instanceof Car) {
        $vehicle->driveTo(new Location('texas'));
    }
}
```

**好:**

```php
function travelToTexas(Traveler $vehicle): void
{
    $vehicle->travelTo(new Location('texas'));
}
```

**[⬆ 返回顶部](#目录)**

### 避免类型检查 (part 2)

如果你正使用基本原始值比如字符串、整形和数组，要求版本是PHP 7+，不用多态，需要类型检测，
那你应当考虑[类型声明](http://php.net/manual/en/functions.arguments.php#functions.arguments.type-declaration)或者严格模式。
提供了基于标准PHP语法的静态类型。 手动检查类型的问题是做好了需要好多的废话，好像为了安全就可以不顾损失可读性。
保持你的PHP 整洁，写好测试，做好代码回顾。做不到就用PHP严格类型声明和严格模式来确保安全。

**坏:**

```php
function combine($val1, $val2): int
{
    if (!is_numeric($val1) || !is_numeric($val2)) {
        throw new \Exception('Must be of type Number');
    }

    return $val1 + $val2;
}
```

**好:**

```php
function combine(int $val1, int $val2): int
{
    return $val1 + $val2;
}
```

**[⬆ 返回顶部](#目录)**

### 移除僵尸代码

僵尸代码和重复代码一样坏。没有理由保留在你的代码库中。如果从来没被调用过，就删掉！
因为还在代码版本库里，因此很安全。

**坏:**
```php
function oldRequestModule(string $url): void
{
    // ...
}

function newRequestModule(string $url): void
{
    // ...
}

$request = newRequestModule($requestUrl);
inventoryTracker('apples', $request, 'www.inventory-awesome.io');
```

**好:**

```php
function requestModule(string $url): void
{
    // ...
}

$request = requestModule($requestUrl);
inventoryTracker('apples', $request, 'www.inventory-awesome.io');
```

**[⬆ 返回顶部](#目录)**


## 对象和数据结构

### 使用 getters 和 setters

在PHP中你可以对方法使用`public`, `protected`, `private` 来控制对象属性的变更。

* 当你想对对象属性做获取之外的操作时，你不需要在代码中去寻找并修改每一个该属性访问方法
* 当有`set`对应的属性方法时，易于增加参数的验证
* 封装内部的表示
* 使用set*和get*时，易于增加日志和错误控制
* 继承当前类时，可以复写默认的方法功能
* 当对象属性是从远端服务器获取时，get*，set*易于使用延迟加载

此外，这样的方式也符合OOP开发中的[开闭原则](#开闭原则)

**坏:**

```php
class BankAccount
{
    public $balance = 1000;
}

$bankAccount = new BankAccount();

// Buy shoes...
$bankAccount->balance -= 100;
```

**好:**

```php
class BankAccount
{
    private $balance;

    public function __construct(int $balance = 1000)
    {
      $this->balance = $balance;
    }

    public function withdraw(int $amount): void
    {
        if ($amount > $this->balance) {
            throw new \Exception('Amount greater than available balance.');
        }

        $this->balance -= $amount;
    }

    public function deposit(int $amount): void
    {
        $this->balance += $amount;
    }

    public function getBalance(): int
    {
        return $this->balance;
    }
}

$bankAccount = new BankAccount();

// Buy shoes...
$bankAccount->withdraw($shoesPrice);

// Get balance
$balance = $bankAccount->getBalance();
```

**[⬆ 返回顶部](#目录)**

### 给对象使用私有或受保护的成员变量

* 对`public`方法和属性进行修改非常危险，因为外部代码容易依赖他，而你没办法控制。**对之修改影响所有这个类的使用者。** `public` methods and properties are most dangerous for changes, because some outside code may easily rely on them and you can't control what code relies on them. **Modifications in class are dangerous for all users of class.**
* 对`protected`的修改跟对`public`修改差不多危险，因为他们对子类可用，他俩的唯一区别就是可调用的位置不一样，**对之修改影响所有集成这个类的地方。**  `protected` modifier are as dangerous as public, because they are available in scope of any child class. This effectively means that difference between public and protected is only in access mechanism, but encapsulation guarantee remains the same. **Modifications in class are dangerous for all descendant classes.**
* 对`private`的修改保证了这部分代码**只会影响当前类**`private` modifier guarantees that code is **dangerous to modify only in boundaries of single class** (you are safe for modifications and you won't have [Jenga effect](http://www.urbandictionary.com/define.php?term=Jengaphobia&defid=2494196)).

所以，当你需要控制类里的代码可以被访问时才用`public/protected`，其他时候都用`private`。

可以读一读这篇 [博客文章](http://fabien.potencier.org/pragmatism-over-theory-protected-vs-private.html) ，[Fabien Potencier](https://github.com/fabpot)写的.

**坏:**

```php
class Employee
{
    public $name;

    public function __construct(string $name)
    {
        $this->name = $name;
    }
}

$employee = new Employee('John Doe');
echo 'Employee name: '.$employee->name; // Employee name: John Doe
```

**好:**

```php
class Employee
{
    private $name;

    public function __construct(string $name)
    {
        $this->name = $name;
    }

    public function getName(): string
    {
        return $this->name;
    }
}

$employee = new Employee('John Doe');
echo 'Employee name: '.$employee->getName(); // Employee name: John Doe
```

**[⬆ 返回顶部](#目录)**

## 类

### 少用继承多用组合

正如  the Gang of Four 所著的[*设计模式*](https://en.wikipedia.org/wiki/Design_Patterns)之前所说，
我们应该尽量优先选择组合而不是继承的方式。使用继承和组合都有很多好处。
这个准则的主要意义在于当你本能的使用继承时，试着思考一下`组合`是否能更好对你的需求建模。
在一些情况下，是这样的。

接下来你或许会想，“那我应该在什么时候使用继承？” 
答案依赖于你的问题，当然下面有一些何时继承比组合更好的说明：

1. 你的继承表达了“是一个”而不是“有一个”的关系（人类-》动物，用户-》用户详情）
2. 你可以复用基类的代码（人类可以像动物一样移动）
3. 你想通过修改基类对所有派生类做全局的修改（当动物移动时，修改她们的能量消耗）

**糟糕的:**

```php
class Employee 
{
    private $name;
    private $email;

    public function __construct(string $name, string $email)
    {
        $this->name = $name;
        $this->email = $email;
    }

    // ...
}


// 不好，因为 Employees "有" taxdata
// 而 EmployeeTaxData 不是 Employee 类型的


class EmployeeTaxData extends Employee 
{
    private $ssn;
    private $salary;
    
    public function __construct(string $name, string $email, string $ssn, string $salary)
    {
        parent::__construct($name, $email);

        $this->ssn = $ssn;
        $this->salary = $salary;
    }

    // ...
}
```

**好:**

```php
class EmployeeTaxData 
{
    private $ssn;
    private $salary;

    public function __construct(string $ssn, string $salary)
    {
        $this->ssn = $ssn;
        $this->salary = $salary;
    }

    // ...
}

class Employee 
{
    private $name;
    private $email;
    private $taxData;

    public function __construct(string $name, string $email)
    {
        $this->name = $name;
        $this->email = $email;
    }

    public function setTaxData(string $ssn, string $salary)
    {
        $this->taxData = new EmployeeTaxData($ssn, $salary);
    }

    // ...
}
```

**[⬆ 返回顶部](#目录)**

### 避免连贯接口

[连贯接口Fluent interface](https://en.wikipedia.org/wiki/Fluent_interface)是一种
旨在提高面向对象编程时代码可读性的API设计模式，他基于[方法链Method chaining](https://en.wikipedia.org/wiki/Method_chaining)

有上下文的地方可以降低代码复杂度，例如[PHPUnit Mock Builder](https://phpunit.de/manual/current/en/test-doubles.html)
和[Doctrine Query Builder](http://docs.doctrine-project.org/projects/doctrine-dbal/en/latest/reference/query-builder.html)
，更多的情况会带来较大代价：

While there can be some contexts, frequently builder objects, where this
pattern reduces the verbosity of the code (for example the [PHPUnit Mock Builder](https://phpunit.de/manual/current/en/test-doubles.html)
or the [Doctrine Query Builder](http://docs.doctrine-project.org/projects/doctrine-dbal/en/latest/reference/query-builder.html)),
more often it comes at some costs:

1. 破坏了 [对象封装](https://en.wikipedia.org/wiki/Encapsulation_%28object-oriented_programming%29)
2. 破坏了 [装饰器模式](https://en.wikipedia.org/wiki/Decorator_pattern)
3. 在测试组件中不好做[mock](https://en.wikipedia.org/wiki/Mock_object)
4. 导致提交的diff不好阅读

了解更多请阅读 [连贯接口为什么不好](https://ocramius.github.io/blog/fluent-interfaces-are-evil/)
，作者 [Marco Pivetta](https://github.com/Ocramius).

**坏:**

```php
class Car
{
    private $make = 'Honda';
    private $model = 'Accord';
    private $color = 'white';

    public function setMake(string $make): self
    {
        $this->make = $make;

        // NOTE: Returning this for chaining
        return $this;
    }

    public function setModel(string $model): self
    {
        $this->model = $model;

        // NOTE: Returning this for chaining
        return $this;
    }

    public function setColor(string $color): self
    {
        $this->color = $color;

        // NOTE: Returning this for chaining
        return $this;
    }

    public function dump(): void
    {
        var_dump($this->make, $this->model, $this->color);
    }
}

$car = (new Car())
  ->setColor('pink')
  ->setMake('Ford')
  ->setModel('F-150')
  ->dump();
```

**好:**

```php
class Car
{
    private $make = 'Honda';
    private $model = 'Accord';
    private $color = 'white';

    public function setMake(string $make): void
    {
        $this->make = $make;
    }

    public function setModel(string $model): void
    {
        $this->model = $model;
    }

    public function setColor(string $color): void
    {
        $this->color = $color;
    }

    public function dump(): void
    {
        var_dump($this->make, $this->model, $this->color);
    }
}

$car = new Car();
$car->setColor('pink');
$car->setMake('Ford');
$car->setModel('F-150');
$car->dump();
```

**[⬆ 返回顶部](#目录)**

### 推荐使用 final 类

能用时尽量使用 `final` 关键字:

1. 阻止不受控的继承链
2. 鼓励 [组合](#少用继承多用组合).
3. 鼓励 [单一职责模式](#单一职责模式).
4. 鼓励开发者用你的公开方法而非通过继承类获取受保护方法的访问权限.
5. 使得在不破坏使用你的类的应用的情况下修改代码成为可能.

The only condition is that your class should implement an interface and no other public methods are defined.

For more informations you can read [the blog post](https://ocramius.github.io/blog/when-to-declare-classes-final/) on this topic written by [Marco Pivetta (Ocramius)](https://ocramius.github.io/).

**Bad:**

```php
final class Car
{
    private $color;
    
    public function __construct($color)
    {
        $this->color = $color;
    }
    
    /**
     * @return string The color of the vehicle
     */
    public function getColor() 
    {
        return $this->color;
    }
}
```

**Good:**

```php
interface Vehicle
{
    /**
     * @return string The color of the vehicle
     */
    public function getColor();
}

final class Car implements Vehicle
{
    private $color;
    
    public function __construct($color)
    {
        $this->color = $color;
    }
    
    /**
     * {@inheritdoc}
     */
    public function getColor() 
    {
        return $this->color;
    }
}
```

**[⬆ 返回顶部](#目录)**

## SOLID

**SOLID** 是Michael Feathers推荐的便于记忆的首字母简写，它代表了Robert Martin命名的最重要的五个面对对象编码设计原则

 * [S: 单一职责原则 (SRP)](#职责原则)
 * [O: 开闭原则 (OCP)](#开闭原则)
 * [L: 里氏替换原则 (LSP)](#里氏替换原则)
 * [I: 接口隔离原则 (ISP)](#接口隔离原则)
 * [D: 依赖倒置原则 (DIP)](#依赖倒置原则)


### 单一职责原则

Single Responsibility Principle (SRP)

正如在Clean Code所述，"修改一个类应该只为一个理由"。
人们总是易于用一堆方法塞满一个类，如同我们只能在飞机上
只能携带一个行李箱（把所有的东西都塞到箱子里）。这样做
的问题是：从概念上这样的类不是高内聚的，并且留下了很多
理由去修改它。将你需要修改类的次数降低到最小很重要。
这是因为，当有很多方法在类中时，修改其中一处，你很难知
晓在代码库中哪些依赖的模块会被影响到。

**坏:**

```php
class UserSettings
{
    private $user;

    public function __construct(User $user)
    {
        $this->user = $user;
    }

    public function changeSettings(array $settings): void
    {
        if ($this->verifyCredentials()) {
            // ...
        }
    }

    private function verifyCredentials(): bool
    {
        // ...
    }
}
```

**好:**

```php
class UserAuth 
{
    private $user;

    public function __construct(User $user)
    {
        $this->user = $user;
    }
    
    public function verifyCredentials(): bool
    {
        // ...
    }
}

class UserSettings 
{
    private $user;
    private $auth;

    public function __construct(User $user) 
    {
        $this->user = $user;
        $this->auth = new UserAuth($user);
    }

    public function changeSettings(array $settings): void
    {
        if ($this->auth->verifyCredentials()) {
            // ...
        }
    }
}
```

**[⬆ 返回顶部](#目录)**

### 开闭原则

Open/Closed Principle (OCP)

正如Bertrand Meyer所述，"软件的工件（ classes, modules, functions 等）
应该对扩展开放，对修改关闭。" 然而这句话意味着什么呢？这个原则大体上表示你
应该允许在不改变已有代码的情况下增加新的功能

**坏:**

```php
abstract class Adapter
{
    protected $name;

    public function getName(): string
    {
        return $this->name;
    }
}

class AjaxAdapter extends Adapter
{
    public function __construct()
    {
        parent::__construct();

        $this->name = 'ajaxAdapter';
    }
}

class NodeAdapter extends Adapter
{
    public function __construct()
    {
        parent::__construct();

        $this->name = 'nodeAdapter';
    }
}

class HttpRequester
{
    private $adapter;

    public function __construct(Adapter $adapter)
    {
        $this->adapter = $adapter;
    }

    public function fetch(string $url): Promise
    {
        $adapterName = $this->adapter->getName();

        if ($adapterName === 'ajaxAdapter') {
            return $this->makeAjaxCall($url);
        } elseif ($adapterName === 'httpNodeAdapter') {
            return $this->makeHttpCall($url);
        }
    }

    private function makeAjaxCall(string $url): Promise
    {
        // request and return promise
    }

    private function makeHttpCall(string $url): Promise
    {
        // request and return promise
    }
}
```

**好:**

```php
interface Adapter
{
    public function request(string $url): Promise;
}

class AjaxAdapter implements Adapter
{
    public function request(string $url): Promise
    {
        // request and return promise
    }
}

class NodeAdapter implements Adapter
{
    public function request(string $url): Promise
    {
        // request and return promise
    }
}

class HttpRequester
{
    private $adapter;

    public function __construct(Adapter $adapter)
    {
        $this->adapter = $adapter;
    }

    public function fetch(string $url): Promise
    {
        return $this->adapter->request($url);
    }
}
```

**[⬆ 返回顶部](#目录)**

### 里氏替换原则

Liskov Substitution Principle (LSP)

这是一个简单的原则，却用了一个不好理解的术语。它的正式定义是
"如果S是T的子类型，那么在不改变程序原有既定属性（检查、执行
任务等）的前提下，任何T类型的对象都可以使用S类型的对象替代
（例如，使用S的对象可以替代T的对象）" 这个定义更难理解:-)。

对这个概念最好的解释是：如果你有一个父类和一个子类，在不改变
原有结果正确性的前提下父类和子类可以互换。这个听起来依旧让人
有些迷惑，所以让我们来看一个经典的正方形-长方形的例子。从数学
上讲，正方形是一种长方形，但是当你的模型通过继承使用了"is-a"
的关系时，就不对了。

**坏:**

```php
class Rectangle
{
    protected $width = 0;
    protected $height = 0;

    public function setWidth(int $width): void
    {
        $this->width = $width;
    }

    public function setHeight(int $height): void
    {
        $this->height = $height;
    }

    public function getArea(): int
    {
        return $this->width * $this->height;
    }
}

class Square extends Rectangle
{
    public function setWidth(int $width): void
    {
        $this->width = $this->height = $width;
    }

    public function setHeight(int $height): void
    {
        $this->width = $this->height = $height;
    }
}

function printArea(Rectangle $rectangle): void
{
    $rectangle->setWidth(4);
    $rectangle->setHeight(5);
 
    // BAD: Will return 25 for Square. Should be 20.
    echo sprintf('%s has area %d.', get_class($rectangle), $rectangle->getArea()).PHP_EOL;
}

$rectangles = [new Rectangle(), new Square()];

foreach ($rectangles as $rectangle) {
    printArea($rectangle);
}
```

**好:**

最好是将这两种四边形分别对待，用一个适合两种类型的更通用子类型来代替。

尽管正方形和长方形看起来很相似，但他们是不同的。
正方形更接近菱形，而长方形更接近平行四边形。但他们不是子类型。
尽管相似，正方形、长方形、菱形、平行四边形都是有自己属性的不同形状。

```php
interface Shape
{
    public function getArea(): int;
}

class Rectangle implements Shape
{
    private $width = 0;
    private $height = 0;

    public function __construct(int $width, int $height)
    {
        $this->width = $width;
        $this->height = $height;
    }

    public function getArea(): int
    {
        return $this->width * $this->height;
    }
}

class Square implements Shape
{
    private $length = 0;

    public function __construct(int $length)
    {
        $this->length = $length;
    }

    public function getArea(): int
    {
        return $this->length ** 2;
    }
}

function printArea(Shape $shape): void
{
    echo sprintf('%s has area %d.', get_class($shape), $shape->getArea()).PHP_EOL;
}

$shapes = [new Rectangle(4, 5), new Square(5)];

foreach ($shapes as $shape) {
    printArea($shape);
}
```

**[⬆ 返回顶部](#目录)**

### 接口隔离原则

Interface Segregation Principle (ISP)

接口隔离原则表示："调用方不应该被强制依赖于他不需要的接口"

有一个清晰的例子来说明示范这条原则。当一个类需要一个大量的设置项，
为了方便不会要求调用方去设置大量的选项，因为在通常他们不需要所有的
设置项。使设置项可选有助于我们避免产生"胖接口"

**坏:**

```php
interface Employee
{
    public function work(): void;

    public function eat(): void;
}

class HumanEmployee implements Employee
{
    public function work(): void
    {
        // ....working
    }

    public function eat(): void
    {
        // ...... eating in lunch break
    }
}

class RobotEmployee implements Employee
{
    public function work(): void
    {
        //.... working much more
    }

    public function eat(): void
    {
        //.... robot can't eat, but it must implement this method
    }
}
```

**好:**

不是每一个工人都是雇员，但是每一个雇员都是一个工人

```php
interface Workable
{
    public function work(): void;
}

interface Feedable
{
    public function eat(): void;
}

interface Employee extends Feedable, Workable
{
}

class HumanEmployee implements Employee
{
    public function work(): void
    {
        // ....working
    }

    public function eat(): void
    {
        //.... eating in lunch break
    }
}

// robot can only work
class RobotEmployee implements Workable
{
    public function work(): void
    {
        // ....working
    }
}
```

**[⬆ 返回顶部](#目录)**

### 依赖倒置原则

Dependency Inversion Principle (DIP)

这条原则说明两个基本的要点：
1. 高阶的模块不应该依赖低阶的模块，它们都应该依赖于抽象
2. 抽象不应该依赖于实现，实现应该依赖于抽象

这条起初看起来有点晦涩难懂，但是如果你使用过 PHP 框架（例如 Symfony），你应该见过
依赖注入（DI），它是对这个概念的实现。虽然它们不是完全相等的概念，依赖倒置原则使高阶模块
与低阶模块的实现细节和创建分离。可以使用依赖注入（DI）这种方式来实现它。最大的好处
是它使模块之间解耦。耦合会导致你难于重构，它是一种非常糟糕的的开发模式。

**坏:**

```php
class Employee
{
    public function work(): void
    {
        // ....working
    }
}

class Robot extends Employee
{
    public function work(): void
    {
        //.... working much more
    }
}

class Manager
{
    private $employee;

    public function __construct(Employee $employee)
    {
        $this->employee = $employee;
    }

    public function manage(): void
    {
        $this->employee->work();
    }
}
```

**好:**

```php
interface Employee
{
    public function work(): void;
}

class Human implements Employee
{
    public function work(): void
    {
        // ....working
    }
}

class Robot implements Employee
{
    public function work(): void
    {
        //.... working much more
    }
}

class Manager
{
    private $employee;

    public function __construct(Employee $employee)
    {
        $this->employee = $employee;
    }

    public function manage(): void
    {
        $this->employee->work();
    }
}
```

**[⬆ 返回顶部](#目录)**

## 别写重复代码 (DRY)

试着去遵循[DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself) 原则.

尽你最大的努力去避免复制代码，它是一种非常糟糕的行为，复制代码
通常意味着当你需要变更一些逻辑时，你需要修改不止一处。

试想一下，如果你在经营一家餐厅并且你在记录你仓库的进销记录：所有
的土豆，洋葱，大蒜，辣椒等。如果你有多个列表来管理进销记录，当你
用其中一些土豆做菜时你需要更新所有的列表。如果你只有一个列表的话
只有一个地方需要更新。

通常情况下你复制代码是应该有两个或者多个略微不同的逻辑，它们大多数
都是一样的，但是由于它们的区别致使你必须有两个或者多个隔离的但大部
分相同的方法，移除重复的代码意味着用一个function/module/class创
建一个能处理差异的抽象。

用对抽象非常关键，这正是为什么你必须学习遵守在[类](#类)章节写
的SOLID原则，不合理的抽象比复制代码更糟糕，所以务必谨慎！说了这么多，
如果你能设计一个合理的抽象，那就这么干！别写重复代码，否则你会发现
任何时候当你想修改一个逻辑时你必须修改多个地方。

**坏:**

```php
function showDeveloperList(array $developers): void
{
    foreach ($developers as $developer) {
        $expectedSalary = $developer->calculateExpectedSalary();
        $experience = $developer->getExperience();
        $githubLink = $developer->getGithubLink();
        $data = [
            $expectedSalary,
            $experience,
            $githubLink
        ];

        render($data);
    }
}

function showManagerList(array $managers): void
{
    foreach ($managers as $manager) {
        $expectedSalary = $manager->calculateExpectedSalary();
        $experience = $manager->getExperience();
        $githubLink = $manager->getGithubLink();
        $data = [
            $expectedSalary,
            $experience,
            $githubLink
        ];

        render($data);
    }
}
```

**好:**

```php
function showList(array $employees): void
{
    foreach ($employees as $employee) {
        $expectedSalary = $employee->calculateExpectedSalary();
        $experience = $employee->getExperience();
        $githubLink = $employee->getGithubLink();
        $data = [
            $expectedSalary,
            $experience,
            $githubLink
        ];

        render($data);
    }
}
```

**极好:**

最好让代码紧凑一点

```php
function showList(array $employees): void
{
    foreach ($employees as $employee) {
        render([
            $employee->calculateExpectedSalary(),
            $employee->getExperience(),
            $employee->getGithubLink()
        ]);
    }
}
```

**[⬆ 返回顶部](#目录)**

## 翻译

其他语言的翻译:

* :cn: **Chinese:**
   * [php-cpm/clean-code-php](https://github.com/php-cpm/clean-code-php)
* :ru: **Russian:**
   * [peter-gribanov/clean-code-php](https://github.com/peter-gribanov/clean-code-php)
* :es: **Spanish:**
   * [fikoborquez/clean-code-php](https://github.com/fikoborquez/clean-code-php)
* :brazil: **Portuguese:**
   * [fabioars/clean-code-php](https://github.com/fabioars/clean-code-php)
   * [jeanjar/clean-code-php](https://github.com/jeanjar/clean-code-php/tree/pt-br)
* :thailand: **Thai:**
   * [panuwizzle/clean-code-php](https://github.com/panuwizzle/clean-code-php)
* :fr: **French:**
   * [errorname/clean-code-php](https://github.com/errorname/clean-code-php)
* :vietnam: **Vietnamese**
   * [viethuongdev/clean-code-php](https://github.com/viethuongdev/clean-code-php)
* :kr: **Korean:**
   * [yujineeee/clean-code-php](https://github.com/yujineeee/clean-code-php)
* :tr: **Turkish:**
   * [anilozmen/clean-code-php](https://github.com/anilozmen/clean-code-php)
   
**[⬆ 返回顶部](#目录)**
