
### Panda 是什么
Pandas 以 NumPy 和 matplotlib 包为底层驱动，使得 Python 导入和分析数据更容易。  

### Panda 数据结构
> Series: 一维数组，类似于 python 中的基本数据结构 list，区别是 series 只允许存储相同的数据类型，这样可以更有效的使用内存，提高运算效率。  
> DataFrame: 二维的表格型数据结构，可以将 DataFrame 理解为 Series 的容器。(类似 R 中的 data.frame)  
> Panel：三维的数组，可以理解为 DataFrame 的容器。

### Panda 基本操作
```python
# 导入 CSV 文件数据
import pandas
csv_content = pandas.read_csv("csv_file_of_path")
```
