
### 冒泡排序
冒泡排序实现思想：一组数据，依次比较相邻数据的大小，当左边的值大于右边的值时，交换位置，不断轮询，从而实现数据的排序从小到大冒泡。
```php
function bubble_sort($arr)
{
    $arr_length = count($arr);
    if (!is_array($arr) || 0 === $arr_length) {
        return false;
    }

    for ($i = 0; $i < $arr_length; ++$i) {
        $j_length = $arr_length - $i - 1;
        for ($j = 0; $j < $j_length; ++$j) {
            $current = $arr[$j];
            $next    = $arr[$j + 1];
            if ($current > $next) {
                $arr[$j]     = $next;
                $arr[$j + 1] = $current;
            }
        }
    }

    return $arr;
}
```

### 快速排序
快速排序是对冒泡排序的一种改进。

快速排序实现思想：通过一趟排序将待排记录分割成独立的两部分，其中一部分的关键字均比另一部分记录的关键字小，则可分别对这两部分记录继续进行快速排序，整个排序过程可以递归进行，以达到整个序列有序的目的。即找到当前数组中的任意一个元素（一般选择第一个元素），作为标的，新建两个空数组，遍历这个数组元素，如果数组的值比标的小，那么就放到左边的数组，否则放到右面的数组，然后再对这两个数组进行同样的操作。
```php
function quick_sort($arr)
{

}
```

