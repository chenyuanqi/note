
### 迭代法  
迭代法，就是不断地用旧的变量值，递推计算新的变量值。  
人类并不擅长重复性的劳动，而计算机却很适合做这种事。很多问题的现象，思考其本质，看看不断更新变量值或者缩小搜索的区间范围，是否可以获得最终的解（或近似解、局部最优解），如果是就可以尝试迭代法。

```java
/**
 * 求和
 * 
 * @param  number 统计最终数值
 * @return
 */
public static long sum(int number) 
{
    long result = 0;
    for (int i = 0; i <= number; ++i){
        result += i;
    }

    return result;
}
```

一般的，迭代法可以运用在以下几个方面：  
> 1、求数值的精确或者近似解，比如二分法。  
> 2、在一定范围内查找目标值，比如二分查找。  
> 3、机器学习算法中的迭代，很多时候机器学习的过程，就是根据已知的数据和一定的假设，求一个局部最优解。  

### 数学归纳法
“归纳”，是一种从经验事实中找出普遍特征的认知方法。  
数学归纳法，就是指直接从理论上证明某个结论，节约大量的计算资源和时间（形成公式）。  

数学归纳法的一般步骤是这样的：
> 1、证明基本情况是否成立；  
> 2、假设基本条件成立，再证明其他共性条件也是成立的。  

和使用迭代法的计算相比，数学归纳法最大的特点就在于“归纳”二字。它已经总结出了规律，只要我们能够证明这个规律是正确的（从第一步正确，第二步正 确，第三步正确，一直推导到最后一步是正确的），就没有必要进行逐步的推算，可以节省很多时间和资源。

递归调用的代码和数学归纳法的逻辑是一致的。递归把计算交给计算机，归纳把计算交给人，前者是拿计算机的计算成本换人的时间，后者是拿人的时间换计算机的计算成本。  

### 递归
递归就是将复杂的问题，每次都解决一点点，并将剩下的任务转化成为更简单的问题等待下次求解，如此反复，直到最简单的形式。  

递归和循环其实都是迭代法的实现，而且在某些场合下，它们的实现是可以相互转化的。但是，对于某些应用场景，递归确很难被循环取代。  
> 1、递归的核心思想和数学归纳法类似，并更具有广泛性：将当前的问题化解为两部分：一个当前所采取的步骤（可能是进行一次运算、做一个选择等）和另一个更简单的问题（运算的结果离目标值更近）。  
> 2、递归会使用计算机的函数嵌套调用。而函数的调用本身，就可以保存很多中间状态和变量值，因此极大的方便了编程的处理。

```java
/**
 * 阶乘
 * 
 * @param  number 阶乘的最终值
 * @return
 */
public static long factorial(int number) 
{
    if (number <= 1) {
        return 1;
    }

    return number * factorial(number - 1);
}
```

**归并排序中的分治思想**  
对一堆杂乱无序的数字，按照从小到大或者从大到小的规则进行排序，这是计算机领域非常经典的问题。归并排序算法的核心就是“归并”，也就是把两个有序的数列合并起来，形成一个更大的有序数列。  

分而治之（Divide and Conquer），我们通常简称为分治。它的思想就是，将一个复杂的问题，分解成两个甚至多个规模 相同或类似的子问题，然后对这些子问题再进一步细分，直到最后的子问题变得很简单， 很容易就能被求解出来，这样这个复杂的问题就求解出来了。  

归并排序通过分治的思想，把长度为 n 的数列，每次简化为两个长度为 n/2 的数列，这更 有利于计算机的并行处理，只需要 log n 次归并。  
归并排序使用了分治的思想，而这个过程需要使用递归来实现。  
```java
/**
 * 归并排序
 * 
 * @param arr 待排序数组
 */
public static int[] mergingSort(int[] arr) {
    if (arr.length <= 1) {
        return arr;
    }

    int num = arr.length >> 1;
    int[] leftArr = Arrays.copyOfRange(arr, 0, num);
    int[] rightArr = Arrays.copyOfRange(arr, num, arr.length);

    return mergeTwoArray(mergingSort(leftArr), mergingSort(rightArr));
}

private static int[] mergeTwoArray(int[] arr1, int[] arr2) {
    int i = 0, j = 0, k = 0;
    int[] result = new int[arr1.length + arr2.length];
    while(i < arr1.length && j < arr2.length){
        if (arr1[i] <= arr2[j]) {
            result[k++] = arr1[i++];
        } else {
            result[k++] = arr2[j++];
        }
    }

    // 序列 1 中多余的元素移入新数组
    while(i < arr1.length){
        result[k++] = arr1[i++];
    }
    // 序列 2 中多余的元素移入新数组
    while(j < arr2.length){
        result[k++] = arr2[j++];
    }

    return result;
}
```

当需要排序的数组很大（比如达到 1024GB 的时候），我们没法把这些数据都塞入一台普通机器的内存里。这时候，我们可以把这个超级大的数据集，分解为多个更小的数据集（比如 16GB 或者更小），然后分配到多台机器，让它们并行地处理。等所有机器处理完后，中央服务器再进行结果的合并。由于多个小任务间不会相互干扰，可以同时处理，这样会大大增加处理的速度，减少等待时间。  
在单台机器上实现归并排序的时候，我们只需要在递归函数内，实现数据分组以及合并就行了。而在多个机器之间分配数据的时候，递归函数内除了分组及合并，还要负责把数据分发到某台机器上。  

分布式架构中 MapReduce 的主要步骤：  
> 1、数据分割和映射  
> 分割是指将数据源进行切分，并将分片发送到 Mapper 上。映射是指 Mapper 根据应用的需求，将内容按照键 - 值的匹配，存储到哈希结构中。  
> 2、规约  
> 归约是指接受到的一组键值配对，如果是键内容相同的配对，就将它们的值归并。由于哈希映射的关系，MapReduce 还需要 洗牌的步骤，也就是将键 - 值的配对不断地发给对应的 Reducer 进行归约。普通的分治或递归不一定需要洗牌的步骤。  
> 3、合并  
> 为了提升洗牌阶段的效率，可以选择减少发送到归约阶段的键 - 值配对。具体做法是在数 据映射和洗牌之间，加入合并的过程，在每个 Mapper 节点上先进行一次本地的归约。然后只将合并的结果发送到洗牌和归约阶段。

### 排列   
从 n 个不同的元素中取出 m（1<=m<=n）个不同的元素，按照一定的顺序排成一列，这个过程就叫排列。当 m = n 的时候，这就是全排列。  

排列的公式 n(n-1)(n-2)...(n-m+1) = n!/(n-m)!  
```java
import java.util.Arrays;

public class Range
{
    public static void main(String[] args)
    {
        arrangementSelect(new String[]{"1","2","3"}, 3);
    }

    /**
     * 排列选择（从列表中选择 n 个排列）
     *
     * @param dataList 待选列表
     * @param n 选择个数
     */
    public static void arrangementSelect(String[] dataList, int n)
    {
        System.out.println(String.format("A(%d, %d) = %d", dataList.length, n, arrangement(dataList.length, n)));

        arrangementSelect(dataList, new String[n], 0);
    }

    /**
     * 排列选择
     *
     * @param dataList 待选列表
     * @param resultList 前面（resultIndex-1）个的排列结果
     * @param resultIndex 选择索引，从 0 开始
     */
    private static void arrangementSelect(String[] dataList, String[] resultList, int resultIndex)
    {
        int resultLen = resultList.length;

        // 全部选择完时，输出排列结果
        if (resultIndex >= resultLen) {
            System.out.println(Arrays.asList(resultList));
            return ;
        }

        // 递归选择下一个
        for (int i = 0; i < dataList.length; i++) {
            // 判断待选项是否存在于排列结果中
            boolean isExists = false;
            for (int j = 0; j < resultIndex; j++) {
                if (dataList[i].equals(resultList[j])) {
                    isExists = true;
                    break;
                }
            }

            // 排列结果不存在该项，才可选择
            if (!isExists) {
                resultList[resultIndex] = dataList[i];
                arrangementSelect(dataList, resultList, resultIndex + 1);
            }
        }
    }

    /**
     * 计算阶乘数，即 n! = n * (n-1) * ... * 2 * 1
     *
     * @param n
     * @return
     */
    public static long factorial(int n)
    {
        return (n > 1) ? n * factorial(n - 1) : 1;
    }

    /**
     * 计算排列数，即 A(n, m) = n! / (n-m)!
     *
     * @param n 所有排列数
     * @param m 抽取排列数
     * @return
     */
    public static long arrangement(int n, int m)
    {
        return (n >= m) ? factorial(n) / factorial(n - m) : 0;
    }
}
```

排列可以帮助我们生成很多可能性，鉴于此，排列最多的用途就是穷举法，也就是，列出所有可能的情况，一个一个验证，然后看每种情况是否符合条件的解。  
比如田忌赛马所使用的出战顺序，黑客们的暴力破解密码（鉴于存在暴力破解，通常程序也会加上每天尝试密码的次数限制，所以足够长度的多类型字符密码相对安全）。

### 组合
组合是指，从 n 个不同元素中取出 m（1≤m≤n）个不同的元素。对于所有 m 取值的组合之全集合，我们可以叫作全组合（All Combination）。例如对于集合 {1, 2, 3} 而言，全组合就是 {空集, {1}, {2}, {3}, {1, 2}, {1,3} {2, 3}, {1, 2, 3}}。  

组合的公式 总排列/m! = n!/(m! * (n-m)!)  
```java
import java.util.Arrays;

public class Combination
{
    public static void main(String[] args)
    {
        combinationSelect(new String[]{"1","2","3"}, 2);
    }

    /**
     * 组合选择（从列表中选择 n 个组合）
     *
     * @param dataList 待选列表
     * @param n 选择个数
     */
    public static void combinationSelect(String[] dataList, int n)
    {
        System.out.println(String.format("C(%d, %d) = %d", dataList.length, n, combination(dataList.length, n)));

        combinationSelect(dataList, 0, new String[n], 0);
    }

    /**
     * 组合选择
     *
     * @param dataList 待选列表
     * @param dataIndex 待选开始索引
     * @param resultList 前面（resultIndex-1）个的组合结果
     * @param resultIndex 选择索引，从 0 开始
     */
    private static void combinationSelect(String[] dataList, int dataIndex, String[] resultList, int resultIndex)
    {
        int resultLen = resultList.length;
        int resultCount = resultIndex + 1;
        // 全部选择完时，输出组合结果
        if (resultCount > resultLen) {
            System.out.println(Arrays.asList(resultList));
            return ;
        }

        // 递归选择下一个
        for (int i = dataIndex; i < dataList.length + resultCount - resultLen; i++) {
            resultList[resultIndex] = dataList[i];
            combinationSelect(dataList, i + 1, resultList, resultIndex + 1);
        }
    }

    /**
     * 计算阶乘数，即 n! = n * (n-1) * ... * 2 * 1
     *
     * @param n
     * @return
     */
    public static long factorial(int n)
    {
        return (n > 1) ? n * factorial(n - 1) : 1;
    }

    /**
     * 计算组合数，即 C (n, m) = n!/((n-m)! * m!)
     *
     * @param n
     * @param m
     * @return
     */
    public static long combination(int n, int m) {
        return (n >= m) ? factorial(n) / factorial(n - m) / factorial(m) : 0;
    }
}
```

组合可以说是排列的兄弟，两者类似但又有所不同，这两者的区别：排列考虑了取出的元 素它们之间的顺序，而组合无需考虑这种顺序（在递归的时候，传入下一个嵌套调用函数的剩余元素，只需 要包含当前被选元素之后的那些，以避免重复的组合）。  

在自然语言处理中，我们需要用多元文法把临近的几个单词合并起来，组合成一个新的词组。  
普通的多元文法定死了每个元组内单词出现的顺序，但是事实上，多个单词出现时，我们可以不用关心它们的顺序，而只关心它们的组合。这样，我们就可以对多元组内的单词进行某种形式的标准化，即使原来的单词出现顺序有所不同，经过这个标准化过程之后，都会变成唯一的顺序。
