
### 冒泡排序
冒泡排序只会操作相邻的两个数据。
每次冒泡操作都会对相邻的两个元素进行比较，看是否满足大小关系要求；如果不满足就让它俩互换。
```php
function bubbleSort($arr, $sort = 'ASC') 
{
	$length = count($arr);
	if ($length <= 1) {
		return $arr;
	}

	for ($i=0; $i < $length - 1; ++$i) { 
		// 无交换位置，则标志已完成排序
		$completedFlag = true;
		for ($j=$i+1; $j < $length; ++$j) { 
			if (($sort === 'ASC' && $arr[$j] < $arr[$i]) 
				|| ($sort === 'DESC' && $arr[$j] > $arr[$i])) {
				$completedFlag = false;
			    [$arr[$i], $arr[$j]] = [$arr[$j], $arr[$i]];
			}
		}

		if ($completedFlag) {
			break;
		}
	}

	return $arr;
}
```

