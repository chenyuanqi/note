
### 数据结构 - 跳表
链表加多级索引的结构，就是跳表。比如每2个节点提取一个节点到上一级，我们把抽出来的那一级叫做索引或索引层。  
跳表使用空间换时间的设计思路，通过构建多级索引来提高查询的效率，实现了基于链表的“二分查找”。跳表是一种动态数据结构，支持快速地插入、删除、查找操作，时间复杂度都是 O(logn)。  
跳表的空间复杂度是 O(n)。跳表的实现非常灵活，可以通过改变索引构建策略，有效平衡执行效率和内存消耗。  

```php
class SkipList
{
    /**
     * 索引最大层数
     * 
     * @var mixed 
     */
    public $indexMaxLevel;

    /**
     * 头节点
     * 
     * @var SNode 
     */
    protected $headNode;

    public function __construct($indexLevel = 0)
    {
        $this->indexMaxLevel = $indexLevel;
        $this->headNode = new SNode();
    }

    /**
     * 添加
     * 
     * @param $data
     *
     * @return SNode
     */
    public function add($data)
    {
        $newNode = new SNode($data);
        for ($level = $this->getRandomLevel(), $node = $this->headNode; $level >= 0; $level--) {
            while (isset($node->next[$level]) && $data < $node->next[$level]->data) {
                $node = $node->next[$level];
            }
            
            if (isset($node->next[$level])) {
                $newNode->next[$level] = $node->next[$level];
            }
            $node->next[$level] = $newNode;
        }
        
        return $newNode;
    }

    /**
     * 删除
     * 
     * @param $data
     *
     * @return bool
     */
    public function delete($data)
    {
        $deleted = false;
        for ($level = $this->headNode->getMaxLevel(), $node = $this->headNode; $level >= 0; $level--) {
            while (isset($node->next[$level]) && $data < $node->next[$level]->data) {
                $node = $node->next[$level];
            }
            
            if (isset($node->next[$level]) && $data == $node->next[$level]->data) {
                $node->next[$level] = isset($node->next[$level]->next[$level]) 
                    ? $node->next[$level]->next[$level] : null;
                $deleted = true;
            }
        }
        
        return $deleted;
    }

    /**
     * 查找
     * 
     * @param $data
     *
     * @return bool|mixed
     */
    public function find($data)
    {
        for ($level = $this->headNode->getMaxLevel(), $node = $this->headNode; $level >= 0; $level--) {
            while (isset($node->next[$level]) && $data < $node->next[$level]->data) {
                $node = $node->next[$level];
            }
            
            if (isset($node->next[$level]) && $data == $node->next[$level]->data) {
                return $node->next[$level];
            }
        }
        
        return false;
    }

    /**
     * 获取随机层数
     * 
     * @return int
     */
    private function getRandomLevel()
    {
        return mt_rand(0, $this->indexMaxLevel);
    }
}

class SNode
{
    /**
     * 数据域
     *
     * @var null
     */
    public $data;

    /**
     * 指针域，引用 SNode 对象
     *
     * @var array
     */
    public $next = [];

    public function __construct($data = null)
    {
        $this->data = $data;
    }

    /**
     * 获取当前节点索引层数
     *
     * @return int
     */
    public function getMaxLevel()
    {
        return count($this->next) - 1;
    }
}
```
