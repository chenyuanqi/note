
### 数组 Array
Java 语言中提供的数组是用来存储固定大小的同类型元素。  

```java
// 动态初始化：定义数组时未初始化元素时需要指定长度
int[] arr1 = new int[5];
System.out.println(arr1.length); // 数组的长度 5
System.out.println(arr1[0]); // 默认 0

// 静态初始化：定义数组时直接指定初始化的元素，不需要指定长度
int[] arr2 = new int[] { 68, 79, 91, 85, 62 };
// 简写为 int[] arr2 = { 68, 79, 91, 85, 62 };
// System.out.println(arr2.length); // 5
arr2 = new int[] { 1, 2, 3 };
System.out.println(arr2.length); // 3

// 字符串型的数组
String[] strArr = { "aaa", "bbb", "ccc" };
String str = strArr[0];
strArr[0] = "test"; // 数组项的改变不影响引用字符串型的 str
System.out.println(str); // aaa

// 遍历数组（for each 方式，不能指定排序）
// 遍历数组可以这样（能指定排序） for (int i=0; i < arr3.length; i++)
// for-each 的实现原理其实就是使用了普通的 for 循环和迭代器
int[] arr3 = { 1, 4, 19, 16, 25 };
for (int item : arr3) {
    System.out.println(item);
}

// 打印数组
import java.util.Arrays;
System.out.println(Arrays.toString(arr3));

// 数组排序
// 对数组排序会直接修改数组本身
// Arrays.sort(int[] array, int fromIndex, int toIndex) 可以对部分排序，默认全排
import java.util.Arrays;
Arrays.sort(arr3); // 默认从小到大升序排序
// 冒泡排序（从大到小排序）
for (int i = 0; i < arr3.length - 1; i++) {
    for (int j = 0; j < arr3.length - i - 1; j++) {
        if (arr3[j] < arr3[j+1]) {
            int tmp   = arr3[j];
            arr3[j]   = arr3[j+1];
            arr3[j+1] = tmp;
        }
    }
}

// 定义二维数组
int[][] arr = {
    { 1, 2, 3, 4 },
    { 5, 6, 7, 8 },
    { 9, 10, 11, 12 }
};
System.out.println(arr.length); // 3，即包含 3 个一维数组
System.out.println(ns[1][2]); // 7

// 遍历二维数组
for (int[] item : arr) {
    for (int i : item) {
        System.out.print(i + ", ");
    }
    System.out.println();
}
// 1, 2, 3, 4, 
// 5, 6, 7, 8, 
// 9, 10, 11, 12, 

// 打印二维数组
System.out.println(Arrays.deepToString(arr)); // [[1, 2, 3, 4], [5, 6, 7, 8], [9, 10, 11, 12]]
```
