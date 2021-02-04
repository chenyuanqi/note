
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
        $flag = false;
        for ($j = 0; $j < $arr_length - 1; ++$j) {
            $current = $arr[$j];
            $next    = $arr[$j + 1];
            if ($current > $next) {
                $arr[$j]     = $next;
                $arr[$j + 1] = $current;
                $flag        = true;
            }
        }

        if (!$flag) {
            break;
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

快速排序优化方案：  
1、单边递归优化（当本层完成了 partition 操作以后，让本层继续完成基准值左边的 partition 操作，而基准值右边的排序工作交给下一层递归函数去处理）  
```php
function quick_sort($arr, $l, $r) {
    while ($l < $r) {
        // 进行一轮 partition 操作
        // 获得基准值的位置
        $index = partition($arr, $l, $r);
        // 右侧正常调用递归函数 
        quick_sort($arr, $index + 1, $r);
        // 用本层处理左侧的排序
        $r = $index - 1;
    }
    return ;
}
```
2、三点取中选取基准值（每一轮取排序区间的头、尾和中间元素这三个值，然后把它们排序以后的中间值作为本轮的基准值）  
3、partition 操作优化（先从后向前找个小于基准值的数字放到前面，再从前向后找个大于基准值的数字放到后面，直到首尾指针相遇为止）  

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
选择排序算法的实现思路有点类似插入排序，也分已排序区间和未排序区间。但是选择排序每次会从未排序区间中找到最小的元素，将其放到已排序区间的末尾。  
选择排序的时间复杂度是 O(n^2)；由于不涉及额外的存储空间，所以是原地排序；由于涉及非相邻元素的位置交换，所以是不稳定的排序算法。  
```php
function selection_sort($arr)
{
    $arr_length = count($arr);
    if (!is_array($arr) || 0 === $arr_length) {
        return [];
    }

    for ($i = 0; $i < $arr_length; $i++) {
        $min= $i;
        for ($j = $i + 1; $j < $arr_length; $j++) {
            if ($arr[$j] < $arr[$min]) {
                $min = $j;
            }
        }
        if ($min != $i) {
            // 方式：交换位置
            $temp = $arr[$i];
            $arr[$i] = $arr[$min];
            $arr[$min] = $temp;

            // 方式：插入位置
            /*$temp = $arr[$min];
            for ($t = $min - 1; $t >= 0; --$t) {
                if ($arr[$t] <= $temp) {
                    break;
                }
                $arr[$t+1] = $arr[$t];
            }
            $arr[$i] = $temp;*/
        }
    }

    return $arr;
}
```

### 归并排序
所谓归并排序，指的是如果要排序一个数组，我们先把数组从中间分成前后两部分，然后对前后两部分分别排序，再将排好序的两部分合并在一起，这样整个数组就都有序了。  
归并排序使用了分治思想，分治，顾名思义，就是分而治之，将一个大问题分解成小的子问题来解决。说到这里，可能你就能联想起我们之前讲到的一个编程技巧 —— 递归，没错，归并排序就是通过递归来实现的。这个递归的公式是每次都将传入的待排序数组一分为二，直到不能分割，然后将排序后序列合并，最终返回排序后的数组。  
归并排序不涉及相等元素位置交换，是稳定的排序算法，时间复杂度是 O(nlogn)，要优于 O(n^2)，但是归并排序需要额外的空间存放排序数据，不是原地排序，最多需要和待排序数组同样大小的空间，所以空间复杂度是 O(n)。  
```php
function merge_sort($arr)
{
    if (count($arr) <= 1) {
        return  $arr;
    }

    merge_sort_c($arr, 0, count($arr) - 1);
    return $arr;
}

function merge_sort_c(&$arr, $p, $r)
{
    if ($p >= $r) {
        return;
    }

    $q = floor(($p + $r) / 2);
    merge_sort_c($arr, $p, $q);
    merge_sort_c($arr, $q + 1, $r);

    merge($arr, ['start' => $p, 'end' => $q], ['start' => $q + 1, 'end' => $r]);
}

function merge(&$arr, $arr_p, $arr_q){
    $temp = [];
    $i = $arr_p['start'];
    $j = $arr_q['start'];
    $k = 0;
    while ($i <= $arr_p['end'] && $j <= $arr_q['end']) {
        if ($arr[$i] <= $arr[$j]) {
            $temp[$k++] = $arr[$i++];
        } else {
            $temp[$k++] = $arr[$j++];
        }
    }

    if ($i <= $arr_p['end']) {
        for (; $i <= $arr_p['end']; $i++) {
            $temp[$k++] = $arr[$i];
        }
    }

    if ($j <= $arr_q['end']) {
        for (; $j <= $arr_q['end']; $j++) {
            $temp[$k++] = $arr[$j];
        }
    }

    for ($x = 0; $x < $k; $x++) {
        $arr[$arr_p['start'] + $x] = $temp[$x];
    }
}
```

