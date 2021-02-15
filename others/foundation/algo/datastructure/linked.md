
### 数据结构 - 链表
和数组不同，链表并不需要一块连续的内存空间，它通过 “指针” 将一组零散的内存块串联起来使用。

### 单链表
单链表是最原生的链表。  

单链表中有两个节点比较特殊，分别是第一个结点和最后一个结点。  
我们通常把第一个结点叫作头结点，把最后一个结点叫作尾结点。其中，头结点用来记录链表的基地址，有了它，我们就可以遍历得到整条链表。而尾结点特殊的地方是：指针不是指向下一个结点，而是指向一个空地址 NULL，表示这是链表上最后一个结点。  
对单链表而言，理论上来说，插入和删除节点的时间复杂度是 O (1)，查询节点的时间复杂度是 O (n)。
```php
class SingleLinkedList
{
    /**
     * 单链表头结点（哨兵节点）
     *
     * @var SingleLinkedListNode
     */
    public $head;

    /**
     * 单链表长度
     *
     * @var
     */
    private $length;

    /**
     * SingleLinkedList constructor.
     *
     * @param null $head
     */
    public function __construct($head = null)
    {
        $this->head = $head;
        if (!$head) {
            $this->head = new SingleLinkedListNode();
        }

        $this->length = 0;
    }

    /**
     * 获取链表长度
     *
     * @return int
     */
    public function size()
    {
        return $this->length;
    }

    /**
     * 插入数据（头插法）
     *
     * @param $data
     *
     * @return SingleLinkedListNode|bool
     */
    public function insert($data)
    {
        return $this->insertDataAfter($this->head, $data);
    }

    /**
     * 删除节点
     *
     * @param SingleLinkedListNode $node
     *
     * @return bool
     */
    public function delete(SingleLinkedListNode $node)
    {
        if (!$node) {
            return false;
        }

        $preNode = $this->getPreNode($node);
        if (empty($preNode)) {
            return false;
        }

        $preNode->next = $node->next;
        $this->length--;
        unset($node);

        return true;
    }

    /**
     * 通过索引获取节点
     *
     * @param int $index
     *
     * @return SingleLinkedListNode|null
     */
    public function getNodeByIndex($index)
    {
        if ($index >= $this->length) {
            return null;
        }

        $curNode = $this->head->next;
        for ($i = 0; $i < $index; ++$i) {
            $curNode = $curNode->next;
        }

        return $curNode;
    }

    /**
     * 获取某个节点的前置节点
     *
     * @param SingleLinkedListNode $node
     *
     * @return SingleLinkedListNode|bool|null
     */
    public function getPreNode(SingleLinkedListNode $node)
    {
        if (null == $node) {
            return false;
        }

        $curNode = $this->head;
        $preNode = $this->head;
        while ($curNode !== $node) {
            if ($curNode == null) {
                return null;
            }

            $preNode = $curNode;
            $curNode = $curNode->next;
        }

        return $preNode;
    }

    /**
     * 在某个节点后插入新的节点
     *
     * @param SingleLinkedListNode $originNode
     * @param                      $data
     *
     * @return SingleLinkedListNode|bool
     */
    public function insertDataAfter(SingleLinkedListNode $originNode, $data)
    {
        if (!$originNode) {
            return false;
        }

        $newNode = new SingleLinkedListNode();
        $newNode->data = $data;
        // 新节点的 next 指针指向插入节点的 next 指针
        // 插入节点的指针修改为新节点
        $newNode->next = $originNode->next;
        $originNode->next = $newNode;
        $this->length++;

        return $newNode;
    }

    /**
     * 在某个节点前插入新的节点（即该节点的前节点之后插入）
     *
     * @param SingleLinkedListNode $originNode
     * @param                      $data
     *
     * @return SingleLinkedListNode|bool
     */
    public function insertDataBefore(SingleLinkedListNode $originNode, $data)
    {
        if (!$originNode) {
            return false;
        }

        $preNode = $this->getPreNode($originNode);

        return $this->insertDataAfter($preNode, $data);
    }

    /**
     * 在某个节点后插入新的节点
     *
     * @param SingleLinkedListNode $originNode
     * @param SingleLinkedListNode $node
     *
     * @return SingleLinkedListNode|bool
     */
    public function insertNodeAfter(SingleLinkedListNode $originNode, SingleLinkedListNode $node)
    {
        if (!$originNode) {
            return false;
        }

        // 新节点的 next 指针指向插入节点的 next 指针
        // 插入节点的指针修改为新节点
        $node->next = $originNode->next;
        $originNode->next = $node;
        $this->length++;

        return $node;
    }

    /**
     * 输出单链表（注意：data 的数据为可输出类型）
     *
     * @return bool
     */
    public function print()
    {
        if (is_null($this->head->next)) {
            return false;
        }

        $curNode = $this->head;
        while (!is_null($curNode->next)) {
            echo $curNode->next->data . ' -> ';
            $curNode = $curNode->next;
        }
        echo 'NULL' . PHP_EOL;

        return true;
    }
}

class SingleLinkedListNode
{
    /**
     * 节点中的数据域
     *
     * @var null
     */
    public $data;

    /**
     * 节点中的指针域，指向下一个节点
     *
     * @var SingleLinkedListNode
     */
    public $next;

    /**
     * SingleLinkedListNode constructor.
     *
     * @param null $data
     */
    public function __construct($data = null)
    {
        $this->data = $data;
        $this->next = null;
    }
}
```

### 循环链表
循环链表和单链表的区别是尾节点指向了头结点，从而首尾相连，有点像贪吃蛇，可用于解决「约瑟夫环」问题。  
```php

```

### 双向链表
双向链表除了有一个指向下一个节点的指针外，还有一个用于指向上一个节点的指针，从而实现通过 O (1) 复杂度找到上一个节点。  

因为指向上一个节点的指针，使得双向链表在插入、删除节点时比单链表更高效。  
虽然单链表插入、删除时间复杂度已经是 O (1) ，但这只是针对插入、删除操作本身而言，还有后续的操作没计算在内。以删除为例，删除某个节点后，需要将其前驱节点的指针指向被删除节点的下一个节点，这样，我们还需要获取其前驱节点，在单链表中获取前驱节点的时间复杂度是 O (n)，所以综合来看单链表的删除、插入操作时间复杂度也是 O (n)。  
双向链表则不然，它有一个指针指向上一个节点，所以其插入和删除时间复杂度才是真正的 O (1)。  

对于有序链表而言，双向链表的查询效率显然也要高于单链表，不过更优的时间复杂度是靠更差的空间复杂度换取的，双向链表始终需要单链表的两倍空间。  
