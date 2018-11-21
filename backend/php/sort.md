
### 冒泡排序
冒泡排序实现思想：一组数据，依次比较相邻数据的大小，当左边的值大于右边的值时，交换位置，不断轮询，从而实现数据的排序从小到大冒泡。
```php
function bubble_sort($arr)
{
    $arr_length = count($arr);
    if (!is_array($arr) || 0 === $arr_length) {
        return false;
    }

    for ($i = 0; $i < $arr_length - 1; ++$i) {
        for ($j = 0; $j < $arr_length - 1; ++$j) {
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

快速排序实现思想：一组数据，以第一个数值为基准，依次比较后面的每个数据，如果后面的数值比基准数值大，则把该数值归置到 right_arr，否则归置到 left_arr；递归执行 left_arr 和 right_arr，最终与基准值合并成一个排序好的数组。

```php
function quick_sort($arr)
{
    $arr_length = count($arr);
    if (!is_array($arr) || 0 === $arr_length) {
        return [];
    }

    $core_val = $arr[0];
    $left_arr = [];
    $right_arr = [];

    for ($i = 1; $i < $arr_length; ++$i) {
        if ($arr[$i] > $core_val) {
            array_push($right_arr, $arr[$i]);
        } else {
            array_push($left_arr, $arr[$i]);
        }
    }

    $left_arr = quick_sort($left_arr);
    $right_arr = quick_sort($right_arr);

    return array_merge($left_arr, [$core_val], $right_arr);
}
```

### 插入排序
插入排序需要两个嵌套的循环，时间复杂度是 O(n^2)；没有额外的存储空间，是原地排序算法；不涉及相等元素位置交换，是稳定的排序算法。插入排序的时间复杂度和冒泡排序一样，也不是很理想，但是插入排序不涉及数据交换，从更细粒度来区分，性能要略优于冒泡排序。  

插入排序实现思想：将数组中的数据分为两个区间，已排序区间和未排序区间。初始已排序区间只有一个元素，就是数组的第一个元素。插入算法的核心思想是取未排序区间中的元素，在已排序区间中找到合适的插入位置将其插入，并保证已排序区间数据一直有序。重复这个过程，直到未排序区间中元素为空，算法结束。
```php
function insertion_sort($arr) 
{
    $arr_length = count($arr);
    if (!is_array($arr) || 0 === $arr_length) {
        return [];
    }

    for ($i = 0; $i < $arr_length; ++$i) {
        $temp = $arr[$i];
        for ($j = $i - 1; $j >= 0; --$j) {
            if ($arr[$j] <= $temp) {
                break;
            }
            $arr[$j+1] = $arr[$j];
        }
        $arr[$j+1] = $temp;
    }

    return $arr;
}
```

### 选择排序

