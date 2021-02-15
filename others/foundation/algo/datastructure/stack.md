
### 数据结构 - 栈
栈又叫堆栈，是限定只能在一端进行插入和删除操作的线性表，并且满足后进先出（LIFO）的特点。  
我们把允许插入和删除的一端叫做栈顶，另一个端叫做栈底，不含任何数据的栈叫做空栈。

栈支持通过数组 / 链表实现，通过数组实现的通常叫做顺序栈，通过链表实现的叫做链栈。  

堆栈在日常开发和软件使用中，应用非常广泛，比如我们的浏览器前进、倒退功能，编辑器 / IDE 中的撤销、取消撤销功能，程序代码中的函数调用、递归、四则运算等等，都是基于堆栈这种数据结构来实现的，就连著名的 stackoverflow 网站也是取「栈溢出」，需要求教之意。
```php
class Stack
{
    /**
     * 栈的数据
     *
     * @var array
     */
    private $data = [];

    /**
     * 栈的长度（元素个数）
     *
     * @var int
     */
    private $length = 0;

    /**
     * 栈的大小
     *
     * @var int
     */
    private $size = 0;

    /**
     * Stack constructor.
     *
     * @param int $size
     */
    public function __construct($size = 10)
    {
        $this->size = $size;
    }

    /**
     * 入栈
     *
     * @param $value
     *
     * @return bool
     */
    public function push($value)
    {
        if ($this->isFull()) {
            return false;
        }

        array_push($this->data, $value);
        ++$this->length;

        return true;
    }

    /**
     * 出栈
     *
     * @return bool|mixed
     */
    public function pop()
    {
        if ($this->isEmpty()) {
            return false;
        }

        --$this->length;

        return array_pop($this->data);
    }

    /**
     * 判断是否栈满
     *
     * @return bool
     */
    public function isFull()
    {
        return $this->length === $this->size;
    }

    /**
     * 判断是否为空栈
     *
     * @return bool
     */
    public function isEmpty()
    {
        return $this->length === 0;
    }
}
```
