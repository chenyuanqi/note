
### 选择排序
选择排序算法的实现思路有点类似插入排序，也分已排序区间和未排序区间。但是选择排序每次会从未排序区间中找到最小的元素，将其放到已排序区间的末尾（最小元素与未排序第一个元素交换位置）。

```php
function selectionSort($arr, $sort = 'ASC')
{
	$length = count($arr);
	if ($length <= 1) {
		return $arr;
	}

    // 要替换的已选择区后第一个元素，从数组第 1 个元素（即下标 0）开始
	for ($i=0; $i < $length - 1; $i++) { 
	    // 未选择区要替换的目标键
		$replaceKey = $i;
		// 未选择区，查找要替换的目标值的位置
		for ($j=$i; $j < $length - 1; ++$j) { 
		    // 未选择目标值与已选择区每个元素逐个比较
		    // 不符合顺序的元素，将 j+1 的值移动到 j 的值
			if (($sort === 'ASC' && $arr[$replaceKey] > $arr[$j+1]) 
			|| ($sort === 'DESC' && $arr[$replaceKey] < $arr[$j+1])) {
			    $replaceKey = $j+1;
			}
		}

		if ($replaceKey !== $i) {
		    [$arr[$i], $arr[$replaceKey]] = [$arr[$replaceKey], $arr[$i]];
		}
	}
	return $arr;
}
```
