
### 算法
- 算法的概念
> 解决特定问题求解步骤的描述，在计算机中表现为指令的有限序列，并且每条指令表示一个或多个操作  
> 一个问题可以有多种算法，每种算法都有不同的效率  
> 一个算法具有五个特征：有穷性、确切性、输入项、输出项、可行性  
> 比如 1+2+3+...+n => n/2 * (n+1)
> 
> 算法分析的目的在于选择合适的算法和改进算法  
> 
> 时间复杂度  
> 执行算法所需要的计算工作量。一般来说，计算机算法问题规模 n 的函数 f(n)，算法的时间复杂度也因此记做 T(n) = O(f(n))  
> 问题的规模 n 越大，算法执行的时间增长率与 f(n) 的增长率正相关，称作渐进时间复杂度（Asymptotic Time Complexity）  
>
> 时间复杂度的计算方式  
> 1、得出算法的计算次数公式  

```php
// 时间复杂度就是 O(n)
$sum = 0;
for($i = 1;$i <= $n; ++$i) {
    $sum += $i;
}
```

> 2、用常数 1 来取代所有时间中的所有加法常数  
> 如上的算法中，比如 $n 固定为 3，那我们不认为时间复杂度是 O(3)，而是 O(1)  
> 3、在修改后的运行次数函数中，只保留最高阶项  
> 如上的算法中，比如 $n = $n + $n，那么时间复杂度就是 O(n^2)  
> 4、如果最高阶存在且不是 1，则去除与这个项相乘的常数  
> 最高阶是个粗略的值，而不是精确的值  
> 5、其他  
> 常数阶：O(1)  
> 线性阶：O(n)  
> 平方/立方阶：O(n^2) \ O(n^3)  
> 特殊平方阶：O(n^2+n+1) => O(n^2)  
> 对数阶：O(log2n)  
> nlog2n 阶、指数阶  
> 一般地，时间复杂度的效率：O(1)>O(log2n)>O(n)>O(nlog2n)>O(n^2)>O(n^3)>O(2^n)>O(n!)>O(n^n)  
> 
> 时间复杂度的其他概念  
> 最坏情况：最坏情况时的运行时间，一种保证，如果没有特别说明，说的时间复杂度就是最坏情况的时间复杂度  
> 平均情况：期望的运行时间  
> 
> 空间复杂度  
> 算法需要消耗的内存空间，记做 S(n) = O(f(n))  
> 包括程序代码所占用的空间，输入数据所占用的空间和辅助变量所占用的空间这三个方面  
> 计算和表示方法和时间复杂度类似，一般用复杂度的渐近性来表示  
> 
> 空间复杂度的计算方式  
> 有时用空间换取时间  
> 冒泡排序的元素交换，空间复杂度是 O(1)；即冒泡排序使用一个临时变量交换数值  

- 排序算法
> 主要有冒泡排序，直接插入排序，希尔排序，选择排序，快速排序，堆排序，归并排序  
> 比如，冒泡排序的原理是两两相邻的数进行比较，如果反序就交换，否则不交换；它的时间复杂度最坏 O(n^2)，平均 O(n^2)；它的空间复杂度是 O(1)  
> 
> 快速排序、归并排序的理想时间复杂度都是 O(nlog2n)，但是快速排序的时间复杂度并不稳定，最坏情况下复杂度为 O(n^2)，所以最理想的还是归并排序  
> 
> 附录：常用的排序算法的时间复杂度和空间复杂度  
| :---: | :---: | :---: | :---: |
| 排序法 | 最差时间分析 | 平均时间复杂度 | 稳定度 | 空间复杂度 |  
| 冒泡排序 | O(n2) | O(n2) | 稳定 | O(1) |  
| 快速排序 | O(n2) | O(n*log2n) | 不稳定 | O(log2n)~O(n) |  
| 选择排序 | O(n2) | O(n2) | 稳定 | O(1) |  
| 二叉树排序 | O(n2) | O(n*log2n) | 不一定 | O(n) |  
| 插入排序 | O(n2) | O(n2) | 稳定 | O(1) |  
| 堆排序 | O(n*log2n) | O(n*log2n) | 不稳定 | O(1) |  
| 希尔排序 | O | O | 不稳定 | O(1) |  

- 查找算法
> 二分查找原理：从数组的中间元素开始，如果中间元素正好是要查找的元素，搜索结束；如果某一个特定元素大于或者小于中间元素，则在数组大于或小于中间元素的那一半中查找，而且跟开始一样从中间开始比较，如果某一步数组为空，代表找不到  
> 二分查找的时间复杂度最差 O(log2n)，平均 O(log2n)，空间复杂度中迭代是 O(1)，递归是 O(log2n)  
> 
> 顺序查找原理：按一定的顺序检查数组中的每一个元素，直到找到所要找到的特定值为止  
> 顺序查找的时间复杂度最差 O(n)，平均 O(n)，空间复杂度是 O(1)  
> 
> 二分查找算法的时间复杂度最差是 O(log2n)，顺序查找的时间复杂度最差 O(n)，所以二分查找法最快，但是递归情况下，二分查找更消耗内存  

- 其他算法
> 斐契那波递归与非递归的实现
```php
// 递归实现
function Fib($number = 0) { 
    if($number <= 0) {
        return 0;
    }

    if ($number== 1 || $number == 2) {
        return 1;
    }

    return Fib($number - 1) + Fib($number - 2);
}

// 非递归实现
function Fib($number = 0) {
    if($number <= 0) {
        return 0;
    }

    if ($number== 1 || $number == 2) {
        return 1;
    }

    $arr = [1, 1];
    for ($i = 2;$i < $number; ++$i) {
        $arr[$i] = $arr[$i - 1] + $arr[$i - 2];
    } 
    
    return $arr[$number - 1];
}
```

> PHP 内置函数的实现  
> 不使用 PHP 函数实现字符串反转、不使用 array_merge 合并数组  
```php
function str_reverse($str) {
    $reverse_str = '';
    $str_len = 0;
    do {
        if (!isset($str[$str_len])) {
            break;
        }
    } while(++$str_len);

    for ($i = $str_len - 1;$i >= 0;--$i) {
        $reverse_str .= $str[$i];
    }
    
    return $reverse_str;
}
```
### 数据结构
> Array - 数组，最简单、应用最广泛的数据结构之一；它的特征是使用连续的内存来存储，数组中的所有元素必须是相同的类型或类型的衍生（同质数据结构），元素可以通过下标直接访问  
> LinkedList - 链表，线性表的一种，最基本最简单也是最常用的数据结构；它的特征是元素之间的关系是一对一的关系（除了第一个和最后一个元素，其他元素都是首尾相接），顺序存储结构和链式存储结构两种存储方式  
> Stack - 栈，和对列相似，一个带有数据存储特性的数据结构；它的特性是存储数据时先进后出的，栈只有一个出口，只能从栈顶部增加和移除元素  
> Heap - 堆，又叫二叉堆，近似完全二叉树的数据结构；它的特性是子节点的键值或者索引总是小于它的父节点、每个节点的左右子树又是一个二叉堆，根节点最大的堆或者大根堆，最小的叫最小堆或者小根堆  
> List - 线性表，有零个或多个元素组成的有限序列；它的特性是线性表是一个序列，0 个元素构成的线性表是空表，第一个元素无先驱、最后一个元素无后继、其他元素都只有一个先驱和后继、有长度，长度是元素个数，长度有限  
> doubly-linked-list - 双向链表，每个元素都是一个对象，每个对象有一个关键字 key 和两个指针（next 和 prev）  
> queue - 队列，先进先出，并发中使用可以安全将对象从一个任务传给另一个任务   
> set - 集合，保存不重复的元素  
> map - 字典，关联数组，也叫字典或者键值对  
> graph - 图，通常使用邻接矩阵和邻接表表示，前者易实现但是对于稀疏矩阵会浪费较多空间、后者使用链表的方式存储信息但是对于图搜索时间复杂度较高  

### 常见算法题
1、字符串反转的实现
```php
function strReverse($str, $encoding = 'utf-8') {
    $result = '';
    $len = mb_strlen($str);
    for($i=$len-1; $i>=0; $i--){
        $result .= mb_substr($str, $i, 1, $encoding);
    }
    
    return $result;
}

echo strReverse('abccccdd');
```

2、获取最大子串长度及子串  
```php
function getLongestSubStr($str) {
    $length = 0;
    $result = '';
    if (empty($str)) {
        goto RESULT;
    }

    $arr = str_split($str);
    $maxLen = strlen($str) - 1;
    $tempArr = [];
    $tempStr = $prevStr = '';
    $tempLen = 0;
    $sameFlag = false;
    foreach($arr as $k => $v){
        $tempFlag = $v === $prevStr;
        $prevStr = $v;
        $tempStr .= $v;

        if ($sameFlag && !$tempFlag){
            $tempStr = substr($tempStr, -1);
        } elseif (!$sameFlag && $tempFlag) {
            $tempStrResult = substr($tempStr, 0, -2);
            $tempStrResult && array_push($tempArr, [$tempStrResult => strlen($tempStrResult)]);
            $tempStr = substr($tempStr, -2);
        } elseif ($k === $maxLen) {
            $tempStr && array_push($tempArr, [$tempStr => strlen($tempStr)]);
        }
        
        $sameFlag = $tempFlag;
    }

    asort($tempArr);
    
    $longestArr = end($tempArr);
    $result     = array_keys($longestArr)[0];
    $length     = array_values($longestArr)[0];

    RESULT:
    return [$length, $result];
}

getLongestSubStr('aaabcdddffsfsda'); // 5, sfsda
```

3、斐波那契数列的求解  
```php
function fibonacci($num) {
    if ($num === 0) {
        return 0;
    }

    if ($num === 1) {
        return 1;
    }

    return fibonacci($num - 2) + fibonacci($num - 1);
}
```

4、遍历指定目录下所有子目录和子文件  
```php
function scanDir($path){
    $dirHandle = opendir($path);

    $files = [];
    while(($temp = readdir($dirHandle)) !== false){
        if($temp === '.' || $temp === '..'){
            continue;
        }

        $newPath = $path. DIRECTORY_SEPARATOR . $temp;
        if(is_dir(!$newPath)){
            $files[$temp] = scanDir($newPath);
        } else {
            $files[] = $temp;
        }
    }

    closedir($handle);

    return $files;
} 
```

