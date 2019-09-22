
### 什么是 SPL
SPL 即 PHP 标准库(standard PHP library)，从 PHP5.0 起内置的组件和接口，并且从 PHP5.3 已逐渐的成熟。  
SPL 是用于解决典型问题(standard problems)的一组接口与类的集合。  
> spl 数据结构，解决数据怎么存储的问题  
> spl 元素遍历，解决数据怎么查看问题  
> spl 常用方法的统一调用，如数字、集合的大小，自定义遍历  
> spl_autoload_register 实现类自动加载  

似乎众多 PHP 开发者基本没有用 SPL 甚至闻所未闻，一个重要的原因是它的文档极少。  
SPL 对 PHP 引擎进行了扩展，例如 ArrayAccess、Countable 和 SeekableIterator 等接口，它们用于以数组形式操作对象。同时，还可以使用 RecursiveIterator、ArrayObejcts 等其他迭代器进行数据的迭代操作。SPL 还内置几个的对象例如 Exceptions、SplObserver、Spltorage 以及 spl_autoload_register、spl_classes、iterator_apply 等辅助函数（helper functions），用于重载对应的功能。  
```php
class MyLoader 
{
    public static function doAutoload($class) 
    {
        // 本模块对应的 autoload 操作
    }
}

spl_autoload_register( array('MyLoader', 'doAutoload') );
```

SPL 包含如下内容：迭代器、基础接口、数据结构、基础函数、异常、其他。  

### SPL 接口
**SPL 核心概念 Iterator 迭代器接口**  
Iterator 迭代器是常见设计模式之一，普遍应用于一组数据中的统一的遍历操作。可以毫不夸张的说，SPL 提供了所有需要的对应数据类型的迭代器。  

Iterator 迭代器能够使许多不同的数据结构，都能有统一的操作界面，比如一个数据库的结果集、同一个目录中的文件集、或者一个文本中每一行构成的集合。SPL 规定，所有部署了 Iterator 界面的 class，都可以用在 foreach Loop 中。Iterator 界面中包含 5 个必须部署的方法：current()、key()、next()、rewind()、valid()。  
```php
interface Iterator extends Traversable{
    //返回当前元素
    public mixed current ( void );
    
    //返回当前元素的键
    public scalar key ( void );
    
    //向前移动到下一个元素
    public void next ( void );
    
    //返回到迭代器的第一个元素
    public void rewind ( void );
    
    //检查当前位置是否有效
    public boolean valid ( void );
}
```

SPL 提供如下迭代器：  

- RecursiveIterator
- RecursiveIteratorIterator
- OuterIterator
- IteratorIterator
- FilterIterator
- RecursiveFilterIterator
- ParentIterator
- SeekableIterator
- LimitIterator
- GlobIterator
- CachingIterator
- RecursiveCachingIterator
- NoRewindIterator
- AppendIterator
- RecursiveIteratorIterator
- InfiniteIterator
- RegexIterator
- RecursiveRegexIterator
- EmptyIterator
- RecursiveTreeIterator
- ArrayIterator

**ArrayAccess 数组式访问接口**  
实现 ArrayAccess 接口，可以使得 object 像 array 那样操作。ArrayAccess 接口包含四个必须实现的方法：offsetExists()、offsetGet()、offsetSet()、offsetUnset()。  
```php
interface ArrayAccess {
    //检查一个偏移位置是否存在 
    public mixed offsetExists ( mixed $offset  );
    
    //获取一个偏移位置的值 
    public mixed offsetGet( mixed $offset  );
    
    //设置一个偏移位置的值 
    public mixed offsetSet ( mixed $offset  );
    
    //复位一个偏移位置的值 
    public mixed offsetUnset  ( mixed $offset  );
}
```

**IteratorAggregate 聚合式迭代器接口**  
假设对象 A 实现了上面的 ArrayAccess 接口，这时候虽然可以像数组那样操作，却无法使用 foreach 遍历，除非实现了前面提到的 Iterator 接口。另一个解决方法是，有时会需要将数据和遍历部分分开，这时就可以实现 IteratorAggregate 接口。它规定了一个 getIterator() 方法，返回一个使用 Iterator 接口的 object。  
```php
IteratorAggregate extends Traversable {
    /* 获取一个外部迭代器 */
    abstract public Traversable getIterator ( void )
}
```

`注意：虽然都继承自 Traversable，但这是一个无法在 PHP 脚本中实现的内部引擎接口。我们直接使用 IteratorAggregate 或 Iterator 接口来代替它。`  

**RecursiveIterator**  
RecursiveIterator 用于遍历多层数据，它继承了 Iterator 接口，因而也具有标准的 current()、key()、next()、 rewind() 和 valid() 方法。同时，它自己还规定了 getChildren() 和 hasChildren() 方法。getChildren() 必须返回一个实现 RecursiveIterator 的对象。  

**SeekableIterator**  
SeekableIterator 接口也是 Iterator 接口的延伸，除了 Iterator 的 5 个方法以外，还规定了 seek() 方法，参数是元素的位置，返回该元素。如果该位置不存在，则抛出 OutOfBoundsException。  

**Countable**  
Countable 规定了一个 count() 方法，返回结果集的数量。  

### SPL 数据结构
数据结构是计算机存储、组织数据的方式。

SPL 提供了一些数据结构基本类型的实现。  
虽然我们可以使用传统的变量类型来描述数据结构，例如用数组来描述堆栈（Strack），但是毕竟它们不是专门用于描述数据结构的，一次误操作就有可能破坏该堆栈。  

SPL 提供了双向链表、堆栈、队列、堆、降序堆、升序堆、优先级队列、定长数组、对象容器。

**SplDoublyLinkedList 双向链表**  
双链表是一种重要的线性存储结构，对于双链表中的每个节点，不仅仅存储自己的信息，还要保存前驱和后继节点的地址。
```php
SplDoublyLinkedList implements Iterator , ArrayAccess , Countable {    
	/* 方法 */
    public __construct ( void )    
    public void add ( mixed $index , mixed $newval )    
    public mixed bottom ( void )//双链表的尾部节点
    public int count ( void )//双联表元素的个数
    public mixed current ( void )//当前记录
    public int getIteratorMode ( void ) //获取迭代模式
    public bool isEmpty ( void )//检测双链表是否为空
    public mixed key ( void )//当前节点索引
    public void next ( void )//移到下条记录
    public bool offsetExists ( mixed $index )//指定index处节点是否存在
    public mixed offsetGet ( mixed $index )//获取指定index处节点值
    public void offsetSet ( mixed $index , mixed $newval )//设置指定index处值
    public void offsetUnset ( mixed $index )//删除指定index处节点
    public mixed pop ( void )//从双链表的尾部弹出元素
    public void prev ( void )//移到上条记录
    public void push ( mixed $value )//添加元素到双链表的尾部
    public void rewind ( void )//将指针指向迭代开始处
    public string serialize ( void )//序列化存储
    public void setIteratorMode ( int $mode )//设置迭代模式
    public mixed shift ( void )//双链表的头部移除元素
    public mixed top ( void )//双链表的头部节点
    public void unserialize ( string $serialized )//反序列化
    public void unshift ( mixed $value )//双链表的头部添加元素
    public bool valid ( void )//检查双链表是否还有节点
}
```

**SplStack 栈**  
栈 (Stack) 是一种特殊的线性表，因为它只能在线性表的一端进行插入或删除元素 (即进栈和出栈)。  
栈是一种后进先出 (LIFO) 的数据结构。  
```php
SplStack extends SplDoublyLinkedList implements Iterator , ArrayAccess , Countable {
	/* 方法 */
	__construct ( void )
	setIteratorMode ( int $mode ) : void

	/* 继承的方法 */
	public SplDoublyLinkedList::add ( mixed $index , mixed $newval ) : void
	public SplDoublyLinkedList::bottom ( void ) : mixed
	public SplDoublyLinkedList::count ( void ) : int
	public SplDoublyLinkedList::current ( void ) : mixed
	public SplDoublyLinkedList::getIteratorMode ( void ) : int
	public SplDoublyLinkedList::isEmpty ( void ) : bool
	public SplDoublyLinkedList::key ( void ) : mixed
	public SplDoublyLinkedList::next ( void ) : void
	public SplDoublyLinkedList::offsetExists ( mixed $index ) : bool
	public SplDoublyLinkedList::offsetGet ( mixed $index ) : mixed
	public SplDoublyLinkedList::offsetSet ( mixed $index , mixed $newval ) : void
	public SplDoublyLinkedList::offsetUnset ( mixed $index ) : void
	public SplDoublyLinkedList::pop ( void ) : mixed
	public SplDoublyLinkedList::prev ( void ) : void
	public SplDoublyLinkedList::push ( mixed $value ) : void
	public SplDoublyLinkedList::rewind ( void ) : void
	public SplDoublyLinkedList::serialize ( void ) : string
	public SplDoublyLinkedList::setIteratorMode ( int $mode ) : void
	public SplDoublyLinkedList::shift ( void ) : mixed
	public SplDoublyLinkedList::top ( void ) : mixed
	public SplDoublyLinkedList::unserialize ( string $serialized ) : void
	public SplDoublyLinkedList::unshift ( mixed $value ) : void
	public SplDoublyLinkedList::valid ( void ) : bool
}
```

**SplQueue 队列**  
队列也是非常实用的一种数据结构，可以通过加权对值进行排序，由于排序在 php 内部实现，业务代码中将精简不少而且更高效。通过 SplPriorityQueue::setExtractFlags(int  $flag) 设置提取方式可以提取数据（等同最大堆）、优先级、和两者都提取的方式。  
队列是一种先进先出 (FIFO) 的数据结构。使用队列时插入在一端进行而删除在另一端进行。  
```php
SplFixedArray implements Iterator , ArrayAccess , Countable {
　　/* 方法 */　　
　　public __construct ([ int $size = 0 ] )
　　public int count ( void )
　　public mixed current ( void )
　　public static SplFixedArray fromArray ( array $array [, bool $save_indexes = true ] )
　　public int getSize ( void )
　　public int key ( void )
　　public void next ( void )
　　public bool offsetExists ( int $index )
　　public mixed offsetGet ( int $index )
　　public void offsetSet ( int $index , mixed $newval )
　　public void offsetUnset ( int $index )
　　public void rewind ( void )
　　public int setSize ( int $size )
　　public array toArray ( void )
　　public bool valid ( void )
　　public void __wakeup ( void )
}
```

**SplHeap 堆**  
堆 (Heap) 就是为了实现优先队列而设计的一种数据结构，它是通过构造二叉堆 (二叉树的一种) 实现。根节点最大的堆叫做最大堆或大根堆（SplMaxHeap），根节点最小的堆叫做最小堆或小根堆（SplMinHeap）。二叉堆还常用于排序 (堆排序)。  
```php
abstract SplHeap implements Iterator , Countable {    
    /* 方法 用法同双向链表一致 */
    public __construct ( void )    
    abstract protected int compare ( mixed $value1 , mixed $value2 )    
    public int count ( void )    
    public mixed current ( void )    
    public mixed extract ( void )    
    public void insert ( mixed $value )    
    public bool isEmpty ( void )    
    public mixed key ( void )    
    public void next ( void )    
    public void recoverFromCorruption ( void )    
    public void rewind ( void )    
    public mixed top ( void )    
    public bool valid ( void )
}
```

### SPL 类
SPL 除了定义一系列 Interfaces 以外，还提供一系列的内置类，它们对应不同的任务，大大简化了编程。  
```php
// 查看所有内置类
foreach(spl_classes() as $key=>$value){
	echo $key.' -&gt; '.$value.'<br />';
}
```

**SplFileInfo 和 SplFileObject**  
SplFileInfo 和 SplFileObject 两个类用来处理文件操作。  
```php
$file = new SplFileInfo('foo-bar.txt');
 
print_r([
    'getATime' => $file->getATime(), //最后访问时间
    'getBasename' => $file->getBasename(), //获取无路径的basename
    'getCTime' => $file->getCTime(), //获取inode修改时间
    'getExtension' => $file->getExtension(), //文件扩展名
    'getFilename' => $file->getFilename(), //获取文件名
    'getGroup' => $file->getGroup(), //获取文件组
    'getInode' => $file->getInode(), //获取文件inode
    'getLinkTarget' => $file->getLinkTarget(), //获取文件链接目标文件
    'getMTime' => $file->getMTime(), //获取最后修改时间
    'getOwner' => $file->getOwner(), //文件拥有者
    'getPath' => $file->getPath(), //不带文件名的文件路径
    'getPathInfo' => $file->getPathInfo(), //上级路径的SplFileInfo对象
    'getPathname' => $file->getPathname(), //全路径
    'getPerms' => $file->getPerms(), //文件权限
    'getRealPath' => $file->getRealPath(), //文件绝对路径
    'getSize' => $file->getSize(),//文件大小，单位字节
    'getType' => $file->getType(),//文件类型 file  dir  link
    'isDir' => $file->isDir(), //是否是目录
    'isFile' => $file->isFile(), //是否是文件
    'isLink' => $file->isLink(), //是否是快捷链接
    'isExecutable' => $file->isExecutable(), //是否可执行
    'isReadable' => $file->isReadable(), //是否可读
    'isWritable' => $file->isWritable(), //是否可写
]);

// 文件遍历
try {
    foreach(new SplFileObject('foo-bar.txt') as $line) {
        echo $line;
    }
} catch (Exception $e) {
    echo $e->getMessage();
}

// 查找指定行
try {
    $file = new SplFileObject('foo-bar.txt');
    $file->seek(2);
    echo $file->current();
} catch (Exception $e) {
    echo $e->getMessage();
}
```

**DirectoryIterator**  
DirectoryIterator 用来查看一个目录中的所有文件和子目录。  
```php
try{
    foreach ( new DirectoryIterator('./') as $Item ){
        echo $Item.'<br />';
    }
}catch(Exception $e){
    echo 'No files Found!<br />';
}
```

**ArrayObject 和 ArrayIterator**  
ArrayObject 可以将 Array 转化为 object。ArrayIterator 实际是 ArrayObject 类的补充，并提供遍历功能。  
```php
$array = ['koala', 'kangaroo', 'wombat', 'wallaby', 'emu', 'kiwi', 'kookaburra', 'platypus'];

$arrayObj = new ArrayObject($array);

for($iterator = $arrayObj->getIterator();
   $iterator->valid();
   $iterator->next()
){
    echo $iterator->key() . ' => ' . $iterator->current() . '<br />';
}

try {
    $object = new ArrayIterator($array);
    foreach($object as $key=>$value){
        echo $key.' => '.$value.'<br />';
    }
}catch (Exception $e){
    echo $e->getMessage();
}
```

### SPL 异常
SPL 提供一系列标准异常。  

- BadFunctionCallException
- BadMethodCallException
- DomainException
- InvalidArgumentException
- LengthException
- LogicException
- OutOfBoundsException 
- OutOfRangeException 
- OverflowException 
- RangeException 
- RuntimeException 
- UnderflowException 
- UnexpectedValueException 

### SPL 函数
SPL 提供如下函数：  

- class_implements — 返回指定的类实现的所有接口。
- class_parents — 返回指定类的父类。
- class_uses — Return the traits used by the given class
- iterator_apply — 为迭代器中每个元素调用一个用户自定义函数
- iterator_count — 计算迭代器中元素的个数
- iterator_to_array — 将迭代器中的元素拷贝到数组
- spl_autoload_call — 尝试调用所有已注册的__autoload () 函数来装载请求类
- spl_autoload_extensions — 注册并返回 spl_autoload 函数使用的默认文件扩展名。
- spl_autoload_functions — 返回所有已注册的__autoload () 函数。
- spl_autoload_register — 注册给定的函数作为 \_\_autoload 的实现
- spl_autoload_unregister — 注销已注册的 \_\_autoload () 函数
- spl_autoload — \_\_autoload () 函数的默认实现
- spl_classes — 返回所有可用的 SPL 类
- spl_object_hash — 返回指定对象的 hash id
- spl_object_id — Return the integer object handle for given object
