
### 数据结构 —— 栈
栈是一个线性结构，在计算机中是一个相当常见的数据结构。它的特点是只能在某一端添加或删除数据，遵循先进后出的原则。  
每种数据结构都可以有很多种方式来实现，栈其实可以看成是数组的一个子集。  
```php
class Stack {
    
    private static $instance;
    
    private static $stack = [];

    private function __construct() {}

    private function __clone() {}

    private function __wakeup() {}
    
    public static function getInstance()
    {
        if (!static::$instance instanceof static){
            static::$instance = new static();
        }

        return static::$instance;
    }

    public function push($item) 
    {
        array_push(static::$stack, $item);
        
        return static::$stack;
    }

    public function pop() 
    {
        return array_pop(static::$stack);
    }

    public function clear() 
    {
        return static::$stack = [];
    }

    public function getCount() 
    {
        return count(static::$stack);
    }

    public function isEmpty() 
    {
        return 0 === static::getCount();
    }
}
```
栈的应用，比如实现面包屑导航栏。

### 数据结构 —— 队列
队列一个线性结构，特点是在某一端添加数据，在另一端删除数据，遵循先进先出的原则。队列主要有单链队列和循环队列。  
```php

```

### 数据结构 —— 链表
链表是一个线性结构，同时也是一个天然的递归结构。链表结构可以充分利用计算机内存空间，实现灵活的内存动态管理。但是链表失去了数组随机读取的优点，同时链表由于增加了结点的指针域，空间开销比较大。
```php

```

### 数据结构 —— 二叉树


###
