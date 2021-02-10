
### 归并排序
归并排序的核心思想：如果要排序一个数组，我们先把数组从中间分成前后两部分，然后对前后两部分分别排序，再将排好序的两部分合并在一起，这样整个数组就都有序了。

```php
function mergeSort($arr, $sort = 'ASC') 
{
	$length = count($arr);
	if ($length <= 1) {
		return $arr;
	}

	return mergeSortRecursive($arr, 0, $length - 1, $sort);
}


function mergeSortRecursive($arr, $l, $r, $sort = 'ASC')
{
    if ($l >= $r) {
        return [$arr[$r]];
    }

    // 获取中间的位置作为分区点
    $pivot = (int)(($l + $r) / 2);

    // 分治：二分处理
    $left = mergeSortRecursive($arr, $l, $pivot, $sort);
    $right = mergeSortRecursive($arr, $pivot + 1, $r, $sort);
    // 合并处理
    return mergeArray($left, $right, $sort);
}

function mergeArray($left, $right, $sort = 'ASC')
{
    $leftLength = count($left);
    $rightLength = count($right);

    $leftIndex = $rightIndex = 0;
    $tempArr = [];
    // 循环遍历左右数组，写入【临时数组】
    // 循环条件：左右数组下标之和 !== 左右数组长度之和
    do {
    	// 当右数组下标未超出范围，且左数组已遍历完成或：
    	// 1、升序排序，【左数组项】>=【右数组项】时
    	// 2、降序排序，【左数组项】<=【右数组项】时
    	// 将右数组项写入临时数组
    	if ($rightIndex < $rightLength 
    		&& ($leftIndex === $leftLength 
    			|| ($sort === 'ASC' && $left[$leftIndex] >= $right[$rightIndex]) || ($sort === 'DESC' && $left[$leftIndex] <= $right[$rightIndex]))) {
    		$tempArr[] = $right[$rightIndex++];
    	} else {
            $tempArr[] = $left[$leftIndex++];
    	}
    } while ($leftIndex + $rightIndex !== $leftLength + $rightLength);

    return $tempArr;
}
```

