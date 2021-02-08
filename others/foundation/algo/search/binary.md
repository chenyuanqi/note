
### 二分查找
最简单的情况就是有序数组中不存在重复元素，我们在其中用二分查找值等于给定值的数据。
```php
// 非递归实现 - 注意：终止条件、区间上下界更新方法、返回值选择
function binarySearch($arr, $value)
{
	$low = 0;
	$high = count($arr) - 1;
	while ($low < $high) {
		$mid = $low + (($high - $low) >> 1);
		if ($arr[$mid] > $value) {
			$high = $mid - 1;
		} else if ($arr[$mid] < $value) {
			$low = $mid + 1;
		} else {
			return $mid;
		}
	}
	
	return -1;
}

// 递归实现
function binarySearch($arr, $value, $low = 0)
{
	$high = count($arr) - 1;
	return binarySearchInternal($arr, $value, $high, $low);
}

function binarySearchInternal($arr, $value, $high, $low = 0)
 {
	$mid = $low + (($high - $low) >> 1);
	if ($arr[$mid] > $value) {
		return binarySearchInternal($arr, $value, $low, $mid - 1);
	} else if($arr[$mid] < $value) {
		return binarySearchInternal($arr, $value, $mid + 1, $high);
	}

	return $mid;
}
```

### 二分查找变形 - 查找第一个值等于给定值的元素
```php
function binarySearch($arr, $value)
{
	$low = 0;
	$high = count($arr) - 1;
	while ($low <= $high) {
		$mid = $low + (($high - $low) >> 1);
		if ($arr[$mid] > $value) {
			$high = $mid - 1;
		} else if ($arr[$mid] < $value) {
			$low = $mid + 1;
		} else {
			if ($mid === 0 || $arr[$mid - 1] !== $value) {
			    return $mid;
			}
			
			$high = $mid - 1;
		}
	}
	
	return -1;
}
```

### 二分查找变形 - 查找最后一个值等于给定值的元素
```php
function binarySearch($arr, $value)
{
	$low = 0;
	$high = count($arr) - 1;
	while ($low <= $high) {
		$mid = $low + (($high - $low) >> 1);
		if ($arr[$mid] > $value) {
			$high = $mid - 1;
		} else if ($arr[$mid] < $value) {
			$low = $mid + 1;
		} else {
			if ($mid === 0 || $arr[$mid + 1] !== $value) {
			    return $mid;
			}
			
			$low = $mid + 1;
		}
	}
	
	return -1;
}
```

### 二分查找变形 - 查找第一个大于等于给定值的元素
```php
function binarySearch($arr, $value)
{
	$low = 0;
	$high = count($arr) - 1;
	while ($low <= $high) {
		$mid = $low + (($high - $low) >> 1);
		if ($arr[$mid] >= $value) {
			if ($mid === 0 || $arr[$mid - 1] < $value) {
			    return $mid;
			}

			$high = $mid - 1;
		} else {
			$low = $mid + 1;
		}
	}
	
	return -1;
}
```

### 二分查找变形 - 查找最后一个小于等于给定值的元素
```php
function binarySearch($arr, $value)
{
	$low = 0;
	$high = count($arr) - 1;
	while ($low <= $high) {
		$mid = $low + (($high - $low) >> 1);
		if ($arr[$mid] <= $value) {
			if ($mid === $high || $arr[$mid + 1] > $value) {
			    return $mid;
			}

            $low = $mid + 1;
		} else {
			$high = $mid - 1;
		}
	}
	
	return -1;
}
```
