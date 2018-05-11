
### 什么是设计模式
在面向对象中，类是用于生成对象的代码模版，设计模式是用于解决共性问题的代码模版；遵循这样的模板，我们可以快速设计出优秀的代码。
（注意：设计模式只是模板，不是具体的代码；它是为了代码复用，增加可维护性。）

设计模式的宗旨：**重用**。  
面向对象的设计模式都是类之间关系的组合，譬如依赖注入
```php
class Human{}

class Woman
{
    public function __construct(Human $human){}
}
```

### UML 类图
UML 类图（类图中的类，与面向对象语言中的类的概念是对应的）是一种结构图，用于描述一个系统的静态结构。类图以反映类结构和类之间关系为目的，用以描述软件系统的结构，是一种静态建模方法。  
UML 类图是面向对象设计的辅助工具，但不是必须工具。  

UML 类图推荐使用免费的 [UMLet 工具](http://www.umlet.com/umletino/umletino.html)。

类与类之间的关系主要有六种：继承→实现→组合→聚合→关联→依赖。  
继承关系也称泛化关系（Generalization），用于描述父类与子类之间的关系。父类又称作基类，子类则称作派生类。在继承关系中，子类继承父类的所有功能，父类所具有的属性、方法，子类应该都有。子类中除了与父类一致的信息以外，还可以包括额外的信息。  
实现关系（Implementation），主要用来规定接口和实现类的关系。接口（包括抽象类）是方法的集合，在实现关系中，类实现了接口，类中的方法实现了接口声明的所有方法。  
组合关系（Composition），整体与部分的关系，但是整体与部分不可以分开。组合关系表示类之间整体与部分的关系，整体和部分有一致的生存期。一旦整体对象不存在，部分对象也将不存在，是同生共死的关系。  
聚合关系（Aggregation），整体和部分的关系，整体与部分可以分开。聚合关系也表示类之间整体与部分的关系，成员对象是整体对象的一部分，但是成员对象可以脱离整体对象独立存在。  
关联关系（Association），表示一个类的属性保存了对另一个类的一个实例（或多个实例）的引用。关联关系是类与类之间最常用的一种关系，表示一类对象与另一类对象之间有联系。组合、聚合也属于关联关系，只是关联关系的类间关系比其他两种要弱。关联关系有四种：双向关联、单向关联、自关联、多重数关联。  
依赖关系（Dependence），假设 A 类的变化引起了 B 类的变化，则说明 B 类依赖于 A 类。大多数情况下，依赖关系体现在某个类的方法使用另一个类的对象作为参数。依赖关系是一种“使用”关系，特定事物的改变有可能会影响到使用该事物的其他事物，在需要表示一个事物使用另一个事物时使用依赖关系。  

### 设计原则
设计模式有六大原则（SOLID）。  

单一职责原则（ SRP —— Single responsibility principle ），是指一个类只负责一个功能领域中的相应职责，或者可以定义为：就一个类而言，应该只有一个引起它变化的原因。它的目的是类的复杂性降低，可读性提高，可维护性提高。

开闭原则（OCP —— Open Close Principle），即软件模块应该对扩展开放，对修改关闭。比如，在程序需要进行新增功能的时候，不能去修改原有的代码，而是新增代码，实现一个热插拔的效果（热插拔：灵活的去除或添加功能，不影响到原有的功能）。开闭原则的目的是为了使程序的扩展性好，易于维护和升级。  

里氏代换原则（LSP —— Liskov Substitution Principle）是继承复用的基石，只有当衍生类可以替换掉基类，软件单位的功能不受到影响时，基类才能真正被复用，而衍生类也能够在基类的基础上增加新的行为。比如球类，原本是一种体育用品，它的衍生类有篮球、足球、排球、羽毛球等等，如果衍生类替换了基类的原本方法，如把体育用品改成了食用品（那么软件单位的功能受到影响），就不符合里氏代换原则。那么，里氏替换的目的是对实现抽象化的具体步骤的规范。  

接口隔离原则（ISP —— Interface Segregation Principle），即使用多个隔离的接口，比使用单个接口要好。比如登录和注册是属于用户模块的两个接口，比写成一个接口要好的多。接口隔离的目的是提高程序设计灵活性。  

依赖反转原则（DIP —— Dependence Inversion Principle）即针对接口编程，而不是针对实现编程。以计算机系统为例，无论主板、CPU、内存、硬件都是在针对接口设计的，如果针对实现来设计，内存就要对应到针对某个品牌的主板，那么会出现换内存需要把主板也换掉的尴尬。它的目的是降低模块之间的耦合。  

还有一个原则是迪米特法则（DP —— Demeter Principle），它也称最少知道原则，一个实体应当尽量少的与其他实体之间发生相互作用，使得系统功能模块相对独立。比如，一个类公开的 public 属性或方法越多，修改时涉及的面也就越大，变更引起的风险扩散也就越大。它的目的是降低类之间的耦合，减少对其他类的依赖。  

> 据说好的代码要符合「高内聚，低耦合」，那什么是高内聚低耦合呢？  
> 内聚是从功能角度来度量模块内的联系，一个好的内聚模块应当恰好做一件事。它描述的是模块内的功能联系；  
> 耦合是软件结构中各模块之间相互连接的一种度量，耦合强弱取决于模块间接口的复杂程度、进入或访问一个模块的点以及通过接口的数据。  

### 创建型设计模式
创建型模式对类的实例化过程进行了抽象，将软件模块中对象的创建和对象的使用分离。  
为了使软件的结构更加清晰，外界对于这些对象只需要知道它们共同的接口，而不需要清楚其具体的实现细节。这样，就使得整个系统的设计更加符合单一职责原则。  
创建型模式在创建什么(What)，由谁创建(Who)，何时创建(When)等方面都为软件设计者提供了尽可能大的灵活性。  
创建型模式隐藏了类的实例的创建细节，通过隐藏对象如何被创建，和组合在一起达到使整个系统独立的目的。  

主要的创建型设计模式包括简单工厂模式、工厂模式、抽象工厂模式，单例模式。  

工厂模式，就是负责生成其他对象的类或方法。比如交通工具的实现  
```php
interface Vehicle
{
    public function drive();
}

class Car implements Vehicle
{
    public function drive()
    {
        echo '汽车靠四个轮子滚动行走';
    }
}

class Ship implements Vehicle
{
    public function drive()
    {
        echo '轮船靠螺旋桨划水前进';
    }
}

class Aircraft implements Vehicle
{
    public function drive()
    {
        echo '飞机靠螺旋桨和机翼的升力飞行';
    }
}

// 工厂类，根据类名实现类的创建
class VehicleFactory
{
    public static function build($className = null)
    {
        $className = ucfirst($className);
        if ($className && class_exists($className)) {
            return new $className();
        }

        return null;
    }
}
// 不用 new，直接就可以调用了
VehicleFactory::build('Car')->drive();
VehicleFactory::build('Ship')->drive();
VehicleFactory::build('Aircraft')->drive();
```

------  
单例模式，只创建一个对象的类。常见的有 Database 类、Cache 类、配置类、Session 类、 Cache 类等等。  
以 Database 类为例，如果不使用单例模式的话，需要每次都创建一个 Database 对象，多增加一个数据库的连接，用户量大的情况下会给数据库和服务器性能带来巨大的影响。  
```php
class Database
{
    // 声明$instance为私有静态类型，用于保存当前类实例化后的对象
    private static $instance = null;
    // 数据库连接句柄
    private $db = null;

    // 构造方法声明为私有方法，禁止外部程序使用new实例化，只能在内部new
    private function __construct($config = [])
    {
        $dsn = sprintf('mysql:host=%s;dbname=%s', $config['db_host'], $config['db_name']);
        $this->db = new PDO($dsn, $config['db_user'], $config['db_pass']);
    }

    // 这是获取当前类对象的唯一方式
    public static function getInstance($config = [])
    {
        // 检查对象是否已经存在，不存在则实例化后保存到$instance属性
        if(self::$instance == null) {
            self::$instance = new self($config);
        }
        return self::$instance;
    }

    // 获取数据库句柄方法
    public function db()
    {
        return $this->db;
    }

    // 声明成私有方法，禁止克隆对象
    private function __clone(){}
    // 声明成私有方法，禁止重建对象
    private function __wakeup(){}
}

// 使用单例模式
$config = [
    'db_name' => 'user',
    'db_host' => 'localhost',
    'db_user' => 'root',
    'db_pass' => 'root'
];
$dbObj = Database::getInstance($config);
```
从如上代码可以知道，单例模式的特点是 4 私 1 公，一个私有静态属性，构造方法私有，克隆方法私有，重建方法私有，一个公共静态方法。  
单例模式在应用请求的整个生命周期中都有效，这有点类似全局变量，会降低程序的可测试性。大部分情况下，也可以用依赖注入来代替单例模式，避免在应用中引入不必要的耦合。所以，对于仅需生成一个对象的类，首先考虑用依赖注入方式，其次考虑用单例模式来实现。  

### 结构型设计模式
主要的结构型设计模式包括适配器模式、组合模式、外观模式、代理模式。  

适配器模式，就是根据客户端需要，将某个类的接口转换成特定样式的接口，以解决类之间的兼容问题。当我们的代码依赖一些外部的 API，或者依赖一些可能会经常更改的类，那么应该考虑用适配器模式。  
```php
/**
 * 适配器接口，所有的支付适配器都需实现这个接口
 * 不管第三方支付实现方式如何，对于客户端来说，都用 pay()方法完成支付
 */
interface PayAdapter
{
    public function pay();
}

/**
 * 支付宝适配器
 */
class AlipayAdapter implements PayAdapter
{
    public function pay()
    {
        // 实例化 Alipay 类
        // 用 Alipay 的方法实现支付，后续如果改变直接修改这里即可
        $alipay = new Alipay();
        $alipay->sendPayment();
    }
}

/**
 * 微信支付适配器
 */
class WechatPayAdapter implements PayAdapter
{
    public function pay()
    {
        $wechatPay = new WechatPay();
        $wechatPay->scan();
        $wechatPay->doPay();
    }
}

// 客户端代码
$alipay = new AlipayAdapter();
// 用 pay() 方法实现支付
$alipay->pay();
```
大的应用都会不断地加入新库和新 API。为避免它们的变更引发问题，应该用适配器模式包装起来，提供应用统一的引用方式，这样就会让我们的代码更具结构化，便于管理和扩展。  

------  
组合模式是将对象组合成树形结构，以表示‘部分-整体’的层次结构。  
在组合模式中，客户端访问独立对象和组合对象（或称对象集合）一样；独立对象是一个有特定功能的对象，它不引用其他任何其他对象，组合对象则是一个提供相似功能对象的集合，主要用来管理独立对象，并为客户端提供和独立对象一样的访问方式。  
```php
/**
 * 规范独立对象和组合对象必须实现的方法，保证它们提供给客户端统一的
 * 访问方式
 */
abstract class Filesystem
{
    protected $name;

    public function __construct($name)
    {
        $this->name = $name;
    }

    public abstract function getName();
    public abstract function getSize();
}

/**
 * 目录类
 */
class Dir extends Filesystem
{
    private $filesystems = [];

    // 组合对象必须实现添加方法
    // 因为传入参数规定为 Filesystem 类型，所以目录和文件都能添加
    public function add(Filesystem $filesystem)
    {
        $key = array_search($filesystem, $this->filesystems);
        if ($key === false) {
            $this->filesystems[] = $filesystem;
        }
    }

    // 组合对象必须实现移除方法
    public function remove(Filesystem $filesystem)
    {
        $key = array_search($filesystem, $this->filesystems);
        if ($key !== false) {
            unset($this->filesystems[$key]);
        }
    }

    public function getName()
    {
        return '目录：' . $this->name;
    }

    public function getSize()
    {
        $size = 0;
        foreach ($this->filesystems as $filesystem) {
            $size += $filesystem->getSize();
        }

        return $size;
    }
}

/**
 * 独立对象：文本文件类
 */
class TextFile extends Filesystem
{
    public function getName()
    {
        return '文本文件：' . $this->name;
    }

    public function getSize()
    {
        return 10;
    }
}

/**
 * 独立对象：图片文件类
 */
class ImageFile extends Filesystem
{
    public function getName()
    {
        return '图片：' . $this->name;
    }

    public function getSize()
    {
        return 100;
    }
}

/**
 * 独立对象：视频文件类
 */
class VideoFile extends Filesystem
{
    public function getName()
    {
        return '视频：'. $this->name;
    }

    public function getSize()
    {
        return 200;
    }
}
// 创建 home 目录，并加入三个文件
$dir = new Dir('home');
$dir->add(new TextFile('text.txt'));
$dir->add(new ImageFile('bg.png'));
$dir->add(new VideoFile('film.mp4'));

// 在 home 下创建子目录 source
$subDir = new Dir('source');
$dir->add($subDir);

// 创建一个 text2.txt，并放到子目录 source 中
$text2 = new TextFile('text2.txt');
$subDir->add($text2);

// 打印信息
// 文本文件：text2.txt-->10
echo $text2->getName(), '-->', $text2->getSize();
// 目录：source --> 10
echo $subDir->getName(), ' --> ',$subDir->getSize();
// 目录：home --> 320
echo $dir->getName(), ' --> ', $dir->getSize();
```
组合模式分为安全模式和透明模式，这是根据接口中是否包含管理对象的方法来区分的。  
在组合模式中，组合对象和独立对象必须实现一个接口。其中，组合对象必须包含添加和删除节点对象。组合模式通过和装饰模式有着类似的结构图，但是组合模式旨在构造类，而装饰模式重在不生成子类即可给对象添加职责。并且，装饰模式重在修饰，而组合模式重在表示。

### 行为型设计模式
主要的行为型设计模式包括命令模式、迭代器模式、策略模式、观察者模式。  

策略模式定义了一组相同类型的算法，算法之间独立封装，并且可以互换代替。每一个算法（处理方式）称为一个策略，这些算法是同一类型问题的多种处理方式，但是具体行为有差别。  
在应用中，就可以根据环境的不同，选择不同的策略来处理问题。  
比如数组的输出有序列化输出、JSON 字符串输出和数组格式输出等方式，每种输出方式都可以独立封装起来，作为一个策略；如果要把数组保存到数据库则可以选择使用序列化方式转化输出，如果要给 APP 作接口则可以用 JSON 字符串输出。  
```php
/**
 * 策略接口
 */
interface OutputStrategy
{
    public function render($array);
}

/**
 * 策略类 1：返回序列化字符串
 */
class SerializeStrategy implements OutputStrategy
{
    public function render($array)
    {
        return serialize($array);
    }
}

/**
 * 策略类 2：返回 JSON 编码后的字符串
 */
class JsonStrategy implements OutputStrategy
{
    public function render($array)
    {
        return json_encode($array);
    }
}

/**
 * 策略类 3：直接返回数组
 */
class ArrayStrategy implements OutputStrategy
{
    public function render($array)
    {
        return $array;
    }
}

/**
 * 环境角色类
 */
class Output
{
    private $outputStrategy;

    public function __construct(OutputStrategy $outputStrategy)
    {
        $this->outputStrategy = $outputStrategy;
    }

    public function renderOutput($array)
    {
        return $this->outputStrategy->render($array);
    }
}

// 需要返回序列化字符串
$output = new Output(new SerializeStrategy());
$data = $output->renderOutput($arr);

// 需要返回 JSON
$output = new Output(new JsonStrategy());
$data = $output->renderOutput($arr);
```
策略模式主要是用来分离算法，根据相同的行为抽象来做不同的具体策略实现。  
策略模式结构清晰明了、使用简单直观。并且耦合度相对而言较低，扩展方便。同时操作封装也更为彻底，数据更为安全。  
策略模式的缺点是随着策略的增加，子类也会变得繁多；但是，这样并不会影响系统的运行，所以在复杂业务中应该考虑使用。  

------  
观察者模式，也称发布-订阅模式，定义了一个被观察者和多个观察者的、一对多的对象关系，当被观察者状态发生变化的时候，它的所有观察者都会收到通知，并自动更新。观察者模式通常用在实时事件处理系统、组件间解耦、数据库驱动的消息队列系统，同时也是 MVC 设计模式中的重要组成部分。  
比如订单的处理，当订单创建后，系统会发送邮件和短信，并保存日志记录。  
```php
/**
 * 被观察者接口
 */
interface Observable
{
    // 添加 / 注册观察者
    public function attach(Observer $observer);
    // 删除观察者
    public function detach(Observer $observer);
    // 触发通知
    public function notify();
}

/**
 * 被观察者
 * 职责：添加观察者到 $observers 属性中，有变动时通过 notify() 方法执行通知
 */
class Order implements Observable
{
    // 保存所有观察者
    private $observers = [];
    // 订单状态
    private $state = 0;

    // 添加（注册）观察者
    public function attach(Observer $observer)
    {
        $key = array_search($observer, $this->observers);
        if ($key === false) {
            $this->observers[] = $observer;
        }
    }

    // 移除观察者
    public function detach(Observer $observer)
    {
        $key = array_search($observer, $this->observers);
        if ($key !== false) {
            unset($this->observers[$key]);
        }
    }

    // 遍历调用观察者的 update() 方法进行通知，不关心其具体实现方式
    public function notify()
    {
        foreach ($this->observers as $observer) {
            // 把本类对象传给观察者，以便观察者获取当前类对象的信息
            $observer->update($this);
        }
    }

    // 订单状态有变化时发送通知
    public function addOrder()
    {
        $this->state = 1;
        $this->notify();
    }

    // 获取提供给观察者的状态
    public function getState()
    {
        return $this->state;
    }
}

/**
 * 观察者接口
 */
interface Observer
{
    // 接收到通知的处理方法
    public function update(Observable $observable);
}

/**
 * 观察者1：发送邮件
 */
class Email implements Observer
{
    public function update(Observable $observable)
    {
        $state = $observable->getState();
        if ($state) {
            echo '发送邮件：您已经成功下单';
        } else {
            echo '发送邮件：下单失败，请重试';
        }
    }
}

/**
 * 观察者2：短信通知
 */
class Message implements Observer
{
    public function update(Observable $observable)
    {
        $state = $observable->getState();
        if ($state) {
            echo '短信通知：您已下单成功';
        } else {
            echo '短信通知：下单失败，请重试';
        }
    }
}

/**
 * 观察者3：记录日志
 */
class Log implements Observer
{
    public function update(Observable $observable)
    {
        echo '记录日志：生成了一个订单记录';
    }
}

// 创建观察者对象
$email = new Email();
$message = new Message();
$log = new Log();
// 创建被观察者对象 —— 订单
$order = new Order();

// 向订单对象中注册 3 个观察者：发送邮件、短信通知、记录日志
$order->attach($email);
$order->attach($message);
$order->attach($log);
// 添加订单，添加时会自动发送通知给观察者
$order->addOrder();
```
如上代码，扩展观察者是非常方便的。  
在观察者模式中，被观察者完全不需要关心观察者，当自身状态有变化时，遍历执行所有观察者的 update() 方法即可完成通知；被观察者通过添加 attach() 方法，提供给观察者注册，使自己变得可见；当被观察者改变时，给注册的观察者发送通知，至于观察者如何处理通知，被观察者不需要关心。  

这是一种良好的设计，对象之间不必相互理解，同样能够相互通信。  
在面向对象的编程中，任何对象的状态都非常重要，它们是对象间交互的桥梁；当一个对象的改变需要被其他对象关注时，观察者模式就派上用场了。  
