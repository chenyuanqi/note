
### 快速排序
快排的核心思想：  
如果要排序数组中下标从 l 到 r 之间的一组数据，我们选择 l 到 r 之间的任意一个数据作为 pivot（分区点位置）；  
遍历 l 到 r 之间的数据，将小于 pivot 的放到左边，将大于 pivot 的放到右边，将 pivot 放到中间。  
经过这一步骤之后，数组 l 到 r 之间的数据就被分成了三个部分，前面 l 到 pivot-1 之间都是小于 pivot 的，中间是 pivot，后面的 pivot+1 到 r 之间是大于 pivot 的。  

如果分区点选得不够合理，快排的时间复杂度可能会退化为 O(n^2)。那么，我们可以怎样优化分区点的选择呢？  
1、三数取中法
从区间的首、尾、中间，分别取出一个数，然后对比大小，取这 3 个数的中间值作为分区点。  
这样每间隔某个固定的长度，取数据出来比较，将中间值作为分区点的分区算法，肯定要比单纯取某一个数据更好。  
但是，如果要排序的数组比较大，那“三数取中”可能就不够了，可能要“五数取中”或者“十数取中”。  
2、随机法  
随机法就是每次从要排序的区间中，随机选择一个元素作为分区点。  
这种方法并不能保证每次分区点都选的比较好，但是从概率的角度来看，也不大可能会出现每次分区点都选得很差的情况，所以平均情况下，这样选的分区点是比较好的。时间复杂度退化为最糟糕的 O(n2) 的情况，出现的可能性不大。

```php
function quickSort($arr, $sort = 'ASC')
{
	$length = count($arr);
	if ($length <= 1) {
		return $arr;
	}

	$pivot = calculatePivot($arr, 0, $length - 1);

	$leftArr = [];
    $rightArr = [];
    for ($i = 0; $i <= $length - 1; ++$i) {
    	// 注意：分区点不参与对比
        if ($i === $pivot) {
            continue;
        }
        
        // 与【分区点的值】做对比，
        // 1、升序排序，>【分区点的值】放入【右数组】
        // 2、降序排序，<【分区点的值】放入【左数组】
        if (($sort === 'ASC' && $arr[$i] > $arr[$pivot]) || ($sort === 'DESC' && $arr[$i] < $arr[$pivot])) {
            array_push($rightArr, $arr[$i]);
        } else {
            array_push($leftArr, $arr[$i]);
        }
    }

    $leftArr = quickSort($leftArr, $sort);
    $rightArr = quickSort($rightArr, $sort);

	return array_merge($leftArr, [$arr[$pivot]], $rightArr);
}

function calculatePivot($arr, $l, $r)
{
	// 三数取中法获取分区点位置
	$pivot = (int)(($l + $r) / 2);
	// 中间数的条件：大于等于【其他数 1】并且小于等于【其他数 2】
	if (($arr[$l] >= $arr[$pivot] && $arr[$l] <= $arr[$r]) || ($arr[$l] <= $arr[$pivot] && $arr[$l] >= $arr[$r])) {
		return $l;
	} else if (($arr[$r] >= $arr[$pivot] && $arr[$r] <= $arr[$l]) || ($arr[$r] <= $arr[$pivot] && $arr[$r] >= $arr[$l])) {
		return $r;
	}

	return $pivot;
}
```

