
### 数据结构 —— 栈
栈是一个线性结构，在计算机中是一个相当常见的数据结构。它的特点是只能在某一端添加或删除数据，遵循先进后出的原则。  
每种数据结构都可以有很多种方式来实现，栈其实可以看成是数组的一个子集  
```php
class Stack {
    private $stack;

    public function __construct() {
        self::$stack = []
    }

    public function push($item) {
        array_push(static::$stack, $item);
    }

    public function pop() {
        array_pop(static::$stack);
    }

    public function peek() {
        return static::$stack[static::getCount() - 1];
    }

    public function getCount() {
        return count(self::$stack);
    }

    public function isEmpty() {
        return 0 === static::getCount();
    }
}
```
栈的应用，比如...

### 数据结构 —— 队列
队列一个线性结构，特点是在某一端添加数据，在另一端删除数据，遵循先进先出的原则。队列主要有单链队列和循环队列。  
```php

```

