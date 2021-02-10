

### PHP 数组
在 PHP 这种动态语言中，因为数组底层是通过散列表实现的，所以功能异常强大。PHP 的数组可以存储任何类型数据，如果与 Java 对比的话，PHP 数组集成了 Java 的数组、List、Set、Map 于一身，所以写代码的效率比 Java 高了几个量级。

抛开 PHP 或 JavaScript 这种动态语言，对于传统的数组，比如 C 语言和 Java 中的数组，在使用之前都需要声明数组存储数据的类型和数组的大小，数组的优点是可以通过下标值随机访问数组内的任何元素，算法复杂度是 O(1)，非常高效，但是缺点是删除/插入元素比较费劲，以删除为例，需要在删除某个元素后，将后续元素都往前移一位，如果是插入，则需要将插入位置之后的元素都往后移，所以对数组的插入/删除而言，算法复杂度是 O(n)，当然了，这个是针对 C / Java 这种语言而言，PHP 不受此约束，因为它不是传统一样上的数组。


### 数据结构 - 数组
数组（Array）是一种线性表数据结构。它用一组连续的内存空间，来存储一组具有相同类型的数据。
```php
class OriginalArray 
{
	private $data;
	// 容量
	private $capacity;
	private $length = 0;

	public function __construct(int $capacity)
	{
        if($capacity <= 0) {
            throw new Exception('容量值必须为大于 0 的正整数');
        }

        $this->data = array_pad([], $capacity, null);
        $this->capacity = $capacity;
	}

	public function find($index)
	{
		if ($this->checkOutOfRange($index)) {
			throw new Exception("索引值超出数组范围");
		}

        // data[index]_address = base_address + i * data_type_size
		return $this->data[$index];
	}

	public function insert(int $index, $value)
	{
		if ($this->checkOutOfRange($index)) {
			throw new Exception("索引值超出数组范围");
		} else if ($this->checkIfFull()) {
			throw new Exception("数组容量已达上限");
		} else if (!is_null($this->find($index))) {
			throw new Exception("索引下已存在数据");
		}

		$this->data[$index] = $value;
		++$this->length;
	}

	public function udpate(int $index, $value)
	{
        if ($this->checkOutOfRange($index)) {
			throw new Exception("索引值超出数组范围");
		} else if (is_null($this->find($index))) {
			throw new Exception("索引下不存在数据");
		}

		$this->data[$index] = $value;
	}

	public function delete(int $index)
	{
		if (!is_null($this->find($index))) {
			$this->data[$index] = null;
		    --$this->length;
		}
	}

	public function print()
	{
		if ($this->length === 0) {
			echo 'Empty Array' . PHP_EOL;
			return ;
		}

		for ($i=0; $i < $this->length; $i++) { 
			var_dump($this->data[$i]);
		}
	}

	public function checkIfFull()
	{
		return $this->length === $this->capacity;
	}

	private function checkOutOfRange($index)
	{
		return $index < 0 || $index >= $this->capacity;
	}
}

$arr = new OriginalArray(6);
$arr->print();
$arr->insert(0, 'zero');
$arr->insert(1, 'one');
$arr->insert(2, 'two');
$arr->insert(3, 'three');
$arr->insert(4, 'four');
$arr->insert(5, 'five');
// $arr->insert(6, 'six');
$arr->print();
$arr->delete(1);
$arr->print();
```
