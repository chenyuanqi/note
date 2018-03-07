
### NumPy 是什么
NumPy 是一个运行速度非常快的数学库，主要用于数组计算。NumPy 可以让你在 Python 中使用向量和数学矩阵，以及许多用 C 语言实现的底层函数，你还可以体验到从未在原生 Python 上体验过的运行速度。  
NumPy 是 Python 在科学计算领域（数据科学、机器学习等）取得成功的关键之一。  

### Numpy 数据结构
| 类型 | 描述 |
| :-----:  | :----: |
| bool | 用一位存储的布尔类型（值为TRUE或FALSE） |
| inti | 由所在平台决定其精度的整数（一般为int32或int64） |
| int8 | 整数，范围为-128至127 |
| int16 | 整数，范围为-32768至32767 |
| int32 | 整数，范围为-231至231 -1 |
| int64 | 整数，范围为-263至263 - 1 |
| uint8 | 无符号整数，范围为0至255 |
| uint16 | 无符号整数，范围为0至65535 |
| uint32 | 无符号整数，范围为0至232 - 1 |
| uint64 | 无符号整数，范围为0至264 - 1 |
| float16 | 半精度浮点数（16位）：其中用1位表示正负号，5位表示指数，10位表示尾数 |
| float32 | 单精度浮点数（32位）：其中用1位表示正负号，8位表示指数，23位表示尾数 |
| float64或float | 双精度浮点数（64位）：其中用1位表示正负号，11位表示指数，52位表示尾数 |
| complex64 | 复数，分别用两个32位浮点数表示实部和虚部 |
| complex128或complex | 复数，分别用两个64位浮点数表示实部和虚部 |

### Numpy 数组操作
Numpy 的核心是多维数组（ndarrays）。
```python
# 导入 numpy
import numpy # 通常情况，开发者习惯用别名 import numpy as np

# 创建数组
numpy.array([0, 1, 2, 3, 4, 5]) # array([0, 1, 2, 3, 4, 5])
numpy.array((0, 1, 2, 3, 4, 5)) # array([0, 1, 2, 3, 4, 5])
numpy.arange(6) # array([0, 1, 2, 3, 4, 5])
numpy.arange(0, 100, 10) # 从 0 开始以 10 为步进一直到 100（不包含），返回 array([ 0, 10, 20, 30, 40, 50, 60, 70, 80, 90])
numpy.array([
    [1, 2],
    [3, 4]
]) # array([[1, 2], [3, 4]])

# 一维数组的读取
array_var = numpy.arange(6) # array([0, 1, 2, 3, 4, 5])
array_var[0] # 0
# 支持切片 array_var[start:end:step]
array_var[-1:] # array([5])
array_var[:2] + array_var[-2:] # array([4, 6])

# 数组的形状
array_var = numpy.arange(4)
array_var.shape # (4,)
# 修改数组的形状，即修改为几行几列
array_var.shape = 2,2 # array([[0, 1], [2, 3]])
# 数组元素的数据类型
array_var.dtype # dtype('int64')
# 数组元素的总个数
array_var.size
# 数组每一个条目所占的字节
array_var.itemsize # 8
# 数组有多少维
array_var.ndim # 2
# 数组所有元素占用的字节数
array_var.nbytes # 48

# 多维数组切片 array_var[column_start:column_end:column_step, row_start:row_end:row_step]
array_var = numpy.array([
    [1, 2],
    [3, 4]
])
array_var[0, 1:] # array([2])

# 数学操作
array_var - 1 # array([[0, 1], [2, 3]])
array_var + 1 # array([[2, 3], [4, 5]])
array_var * 2 # array([[2, 4], [6, 8]])
array_var / 10 # array([[0.1, 0.2], [0.3, 0.4]])
array_var % 2 # array([[1, 0], [1, 0]])
array_var ** 2 # array([[ 1,  4], [ 9, 16]])
array_var * array_var # array([[ 1,  4], [ 9, 16]])
array_var.min() # 1
array_var.max() # 4
array_var.sum() # 10
array_var.cumsum() # 累加器，返回 array([ 1,  3,  6, 10])

# 数组筛选
array_var[array_var >= 3] # array([3, 4])
numpy.where(array_var >= 3) # 返回元素的位置，(array([1, 1]), array([0, 1]))

```
