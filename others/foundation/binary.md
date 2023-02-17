# 进制

## 进制的概念

进制即进位计数制，是利用固定的数字符号和统一的规则的带进位的计数方法。

任何一种进位计数制都有一个基数，基数为 **X** 的进位计数制称为 X 进制，表示每一个数位上的数运算时都是逢 **X**进一。

对于一个 **X** 进制的数，其具体数值由其中的每个数码和数码所在的数位决定。整数部分从右往左的第 m 个数位表示的权重是 ![img](https://cdn.nlark.com/yuque/__latex/cf28c2daee299bfe0220b54acd2921ae.svg) ，其中 ![img](https://cdn.nlark.com/yuque/__latex/6f8f57715090da2632453988d9a1501b.svg) 最小为 0；小数部分从左往右的第![img](https://cdn.nlark.com/yuque/__latex/7b8b965ad4bca0e41ab51de7b31363a1.svg)个数位表示的权重是![img](https://cdn.nlark.com/yuque/__latex/93fe45e540b6e5f5a34c3e7ec19487b8.svg)，其中 ![img](https://cdn.nlark.com/yuque/__latex/7b8b965ad4bca0e41ab51de7b31363a1.svg) 最小为 11。

十进制的 123.45，123.45 可以写成如下形式：  
![img](https://cdn.nlark.com/yuque/__latex/8f8a8cfa5f4f62b57934be14b8ec3917.svg)  

八进制的 720.5720.5 可以写成如下形式：

![img](https://cdn.nlark.com/yuque/__latex/c66d8804e39b415d192e5b1249883a6e.svg)

## 常见的进制

日常生活中，最常用的进制是十进制。十进制包括十个数码：0,1,2,3,4,5,6,7,8,9

计算机采用的进制是二进制。二进制包括两个数码：0,1

八进制和十六进制也是常见的进制。

八进制包含八个数码：0,1,2,3,4,5,6,7

十六进制包含十六个数码，除了 0 到 9 以外，还有 ![img](https://cdn.nlark.com/yuque/__latex/09b5d5101bf8face44ffdd9895afca95.svg)分别对应十进制下的 10,11,12,13,14,15

## 进制间的转换

### 非十进制转十进制

将非十进制数转成十进制数，只要将每个数位的加权和即可。

例如，将八进制数 ![img](https://cdn.nlark.com/yuque/__latex/bb091bf2402d5e98871c9f7f7489f7d2.svg)转成十进制：

![img](https://cdn.nlark.com/yuque/__latex/ce25bcad55505252ca44d2658e832e10.svg)

### 十进制转非十进制

将十进制数转成 X 进制数，需要对整数部分和小数部分分别转换。

对于整数部分，转换方式是将十进制数的整数部分每次除以 **X** 直到变成 0，并记录每次的余数，反向遍历每次的余数即可得到 **X** 进制表示。

例如，将十进制数 50 转成二进制：

![img](https://cdn.nlark.com/yuque/__latex/37b96a5c745377bdeb4adb99587e08d0.svg)

反向遍历每次的余数，依次是 1,1,0,0,1,0 因此十进制数 50 转成二进制数是 ![img](https://cdn.nlark.com/yuque/__latex/d1ec27561578f959b2351aa5ac3dd7ca.svg)

对于小数部分，转换方式是将十进制数的小数部分每次乘以 X 直到变成 0，并记录每次的整数部分，正序遍历每次的整数部分即可得到 X 进制表示。

例如，将十进制数 0.6875 转成二进制：

![img](https://cdn.nlark.com/yuque/__latex/2ec7792ecef782998d78f1c4996651d0.svg)

正序遍历每次的整数部分，依次是 1,0,1,1 因此十进制数 0.6875 转成二进制数是 ![img](https://cdn.nlark.com/yuque/__latex/dcf95520c6165bfc889f2c5422908726.svg)

需要注意的是，在一种进制下的有限小数，转成另一种进制之后可能变成无限循环小数。例如，十进制数 0.2 转成二进制数是 ![img](https://cdn.nlark.com/yuque/__latex/e8a6a574d01d00653cc6eb414b47de3a.svg)

### 二进制转十六进制

二进制数要转换为十六进制，就是以4位一段，分别转换为十六进制。

从右到左 4位一组

例如 100111110110101

左边不满4位的可以用0补满 0100,1111,1011,0101  
2进制0000对应16位进制0

```shell
0001>>>1
0010>>>2
0011>>>3
0100>>>4
0101>>>5
0110>>>6
0111>>>7
1000>>>8
1001>>>9
1010>>>A
1011>>>B
1100>>>C
1101>>>D
1110>>>E
1111>>>F
```

所以上面的2进制转为16进制为 4FB5

![img](https://cdn.nlark.com/yuque/0/2022/png/450565/1642752023705-26e5836e-c904-4428-99f8-a811188fb858.png)

### 其他进制间的转换

如果需要在两个不同的非十进制之间转换，常规的思路是先转成十进制数，再转成目标进制数。在一些特殊情况下，也可以不经过十进制，直接进行转换。

例如，将二进制数转成八进制数或十六进制数，以及将八进制数或十六进制数转成二进制数，都不需要经过十进制。一位八进制数可以表示成三位二进制数，一位十六进制数可以表示成四位二进制数。

例如，对于二进制数 ，按照三位一组进行分组，可以得到 ![img](https://cdn.nlark.com/yuque/__latex/537c75da88e302821f2b1f8f570147c6.svg) ，按照四位一组进行分组，可以得到 ![img](https://cdn.nlark.com/yuque/__latex/9d13e24f5d5d01c3cdff43d6c6a6fe0f.svg)，因此转成八进制数是 ![img](https://cdn.nlark.com/yuque/__latex/4c87f3656763820a67ed011c8504b346.svg)，转成十六进制数是 ![img](https://cdn.nlark.com/yuque/__latex/bc71eeac20d2c831848d3bd52c5b8475.svg)。
