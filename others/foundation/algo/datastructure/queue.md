
### 数据结构 - 队列
和栈一样，队列也是一种特殊的线性表结构。  
队列是在一端插入，另一端删除，就跟我们平常排队一样的道理，从队尾入队，在队头出去，所以队列的特性是先入先出（FIFO），允许插入的一端叫队尾，允许删除的一端叫队头。  

和栈一样，队列也可以通过数组和链表实现，通过数组实现的叫顺序队列，通过链表实现的叫做链式队列。  

链式栈只需要一个栈顶指针就可以了，因为只允许在栈顶插入删除，但是链式队列需要两个指针，一个指向队头，一个指向队尾。  
通过数组实现的顺序队列有一个问题，就是随着队列元素的插入和删除，队尾指针和队头指针不断后移，而导致队尾指针指向末尾无法插入数据，这时候有可能队列头部还是有剩余空间的。我们当然可以通过数据搬移的方式把所有队列数据往前移，但这会增加额外的时间复杂度，如果频繁操作数据量很大的队列，显然对性能有严重损耗，对此问题的解决方案是使用循环队列，即把队列头尾连起来。此时判断队列是否为空的条件还是 tail === head，但是判断队列是否满的条件就变成了 (tail+1) % maxsize === head，maxsize 是数组的长度，浪费一个空间是为了避免混淆判断空队列的条件。  
队列的应用非常广泛，比如我们常见的消息队列就是队列的典型应用场景。  
```php
class Queue
{
    /**
     * 队列的数据
     *
     * @var array
     */
    private $data = [];

    /**
     * 队列的长度（元素个数）
     * @var int
     */
    private $maxSize = 0;

    /**
     * 队列头部位置
     *
     * @var int
     */
    private $head = 0;

    /**
     * 队列尾部位置
     *
     * @var int
     */
    private $tail = 0;

    public function __construct($size = 10)
    {
        $this->maxSize = ++$size;
    }

    /**
     * 入队
     *
     * @param $value
     *
     * @return bool
     */
    public function enQueue($value)
    {
        if ($this->isFull()) {
            return false;
        }

        array_push($this->data, $value);
        $this->tail = (++$this->tail) % $this->maxSize;

        return true;
    }

    /**
     * 出队
     *
     * @return bool|mixed
     */
    public function deQueue()
    {
        if ($this->isEmpty()) {
            return false;
        }

        $this->head = (++$this->head) % $this->maxSize;

        return array_shift($this->data);
    }

    /**
     * 判断队列是否为满
     *
     * @return bool
     */
    public function isFull()
    {
        return ($this->tail + 1) % $this->maxSize === $this->head;
    }

    /**
     * 判断是否队列为空
     *
     * @return bool
     */
    public function isEmpty()
    {
        return $this->head === $this->tail;
    }

    /**
     * 获取队列长度
     *
     * @return int
     */
    public function size()
    {
        return ($this->tail - $this->head + $this->maxSize) % $this->maxSize;
    }
}
```