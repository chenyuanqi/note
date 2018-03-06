
### Pandas 是什么
Pandas 是 Python 的一个数据分析包。  
Pandas 以 NumPy 和 matplotlib 包为底层驱动，使导入和分析数据更容易。  

### Pandas 数据结构
> Series: 一维数组，类似于 python 中的基本数据结构 list，区别是 series 只允许存储相同的数据类型，这样可以更有效的使用内存，提高运算效率。  
> DataFrame: 二维的表格型数据结构，可以将 DataFrame 理解为 Series 的容器。(类似 R 中的 data.frame)  
> Panel：三维的数组，可以理解为 DataFrame 的容器。  
> Panel + dataframe + series => Pan(el)da(taframe)s(eries)  

### Pandas 安装
```bash
pip install pandas
# 指定版本
pip install -v pandas==0.22.0
```

### Pandas 数据分析基本原理
...

### Pandas 数据读写
| 文件格式 | 读取方法   |  写入方法  |
| ---- | :-----:  | :----: |
| Flat File | read_table/read_csv/read_fwf/read_msgpa | - |
| CSV | read_csv | to_csv |
| JSON | read_json | to_json |
| HTML | read_html | to_html |
| Local clipboard | read_clipboard | to_clipboard |
| MS Excel | read_excel | to_excel |
| HDF5 Format | read_hdf | to_hdf |
| Feather Format | read_feather | to_feather |
| Msgpack | read_msgpack | to_msgpack |
| Stata | read_stata | to_stata |
| SAS | read_sas | - | 
| Python Pickle Format | read_pickle | to_pickle |
| SQL | read_sql | to_sql |
| Google Big Query | read_gbq | to_gbq | 

注意，to_excel 存储 Excel 默认支持 .xlsx 格式，如报错，尝试执行如下命令
```bash
pip install openpyxl
```

### Pandas 基本方法
```python
# 导入 CSV 文件数据
import pandas # 通常，使用别名 pd: import pandas as pd
csv_content = pandas.read_csv("csv_file_of_path")

# 通过传递一个 list 对象来创建一个 Series
import numpy
pandas.Series([1, 3, 5, numpy.nan, 7, 9]) # 可指定索引(默认索引从 0 开始)，pandas.Series([1, 3, 5, numpy.nan, 7, 9], index=["a", "b", "c", "d", "e", "f"])
"""
0    1.0
1    3.0
2    5.0
3    NaN
4    7.0
5    9.0
dtype: float64
"""
# 访问 Series 元素
series_var = pandas.Series([1, 3, 5, numpy.nan, 7, 9])
series_var[2] # 5.0

# 通过传递时间索引创建 DataFrame
pandas.date_range('2018-03-01', periods=7)
"""
DatetimeIndex(['2018-03-01', '2018-03-02', '2018-03-03', '2018-03-04',
               '2018-03-05', '2018-03-06', '2018-03-07'],
              dtype='datetime64[ns]', freq='D')
"""

# 通过传递字典对象创建 DataFrame
df_var = pandas.DataFrame({ 'A' : 1.,
                            'B' : pandas.Timestamp('2018-03-01'),
                            'C' : pandas.Series(1,index=list(range(4)),dtype='float32'),
                            'D' : numpy.array([3] * 4,dtype='int32'),
                            'E' : pandas.Categorical(["test","train","test","train"]),
                            'F' : 'foo' })
"""
     A          B    C  D      E    F
0  1.0 2018-03-01  1.0  3   test  foo
1  1.0 2018-03-01  1.0  3  train  foo
2  1.0 2018-03-01  1.0  3   test  foo
3  1.0 2018-03-01  1.0  3  train  foo
"""
# 访问 DataFrame 元素
df_var["E"][2] # 'test'
# 插入 DataFrame 元素
df_var.insert(insert_index_number, insert_index_name, insert_index_values)
# 删除 DataFrame 元素
df_var.pop(index_name)
# 查看不同列的数据类型
df_var.dtypes
"""
A           float64
B    datetime64[ns]
C           float32
D             int32
E          category
F            object
dtype: object
"""
# 查看头 1 行，默认显示前 5 条
df_var.head(1)
"""
     A          B    C  D     E    F
0  1.0 2018-03-01  1.0  3  test  foo
"""
# 查看尾 1 行，默认显示后 5 条
df_var.tail(1)
"""
     A          B    C  D      E    F
3  1.0 2018-03-01  1.0  3  train  foo
"""
# 查看索引
df_var.index # Int64Index([0, 1, 2, 3], dtype='int64')
# 查看列
df_var.columns # Index(['A', 'B', 'C', 'D', 'E', 'F'], dtype='object')
# 查看底层 numpy 数据
df_var.values
"""
array([[1.0, Timestamp('2018-03-01 00:00:00'), 1.0, 3, 'test', 'foo'],
       [1.0, Timestamp('2018-03-01 00:00:00'), 1.0, 3, 'train', 'foo'],
       [1.0, Timestamp('2018-03-01 00:00:00'), 1.0, 3, 'test', 'foo'],
       [1.0, Timestamp('2018-03-01 00:00:00'), 1.0, 3, 'train', 'foo']],
      dtype=object)
"""
# 快速统计汇总，输出数据集的计数、最大值、最小值等
df_var.describe()
"""
         A    C    D
count  4.0  4.0  4.0
mean   1.0  1.0  3.0
std    0.0  0.0  0.0
min    1.0  1.0  3.0
25%    1.0  1.0  3.0
50%    1.0  1.0  3.0
75%    1.0  1.0  3.0
max    1.0  1.0  3.0
"""
# 计算最小值对应索引的标签
df_var.idxmin()
# 计算最大值对应索引的标签
df_var.idxmax()
# 统计非空数据的数量
df_var.count()
"""
A    4
B    4
C    4
D    4
E    4
F    4
dtype: int64
"""
# 针对 Series, 计算每个值对应的数量统计：series_var.value_counts()

# 数值数据的总和
df_var.sum()
# 数值数据的平均值
df_var.mean()
# 数值数据的算术中值
df_var.median()
# 其他常见的计算方法：
df_var.mad() # 平均绝对偏差
df_var.min() # 最小值
df_var.max() # 最大值
df_var.abs() # 绝对值
df_var.std() # 贝塞尔校准样本标准偏差
df_var.var() # 无偏差
df_var.sem() # 平均值的标准误差
df_var.skew() # 样品偏度（3 次）
df_var.kurt() # 样品偏度（4 次）
df_var.quantile() # 样本分位数（以 % 计算）
df_var.cumsum() # 累计总和
df_var.cummin() # 累计最小值
df_var.cummax() # 累计最大值

# 数据的转置
df_var.T
"""
                     0                    1                    2  \
A                    1                    1                    1   
B  2018-03-01 00:00:00  2018-03-01 00:00:00  2018-03-01 00:00:00   
C                    1                    1                    1   
D                    3                    3                    3   
E                 test                train                 test   
F                  foo                  foo                  foo   

                     3  
A                    1  
B  2018-03-01 00:00:00  
C                    1  
D                    3  
E                train  
F                  foo
"""

# 标签对齐
df_var.reindex(['one', 'two', 'three', 'four'])
"""
        A   B   C   D    E    F
one   NaN NaT NaN NaN  NaN  NaN
two   NaN NaT NaN NaN  NaN  NaN
three NaN NaT NaN NaN  NaN  NaN
four  NaN NaT NaN NaN  NaN  NaN
"""
# 由于数据的缺失等各种因素导致标签错位的现象，或者想匹配新的标签
# Pandas 提供 reindex() 的主要作用如下
# 1、重新排序现有数据以匹配新的一组标签。
# 2、在没有标签对应数据的位置插入缺失值（NaN）标记。
# 3、特殊情形下，使用逻辑填充缺少标签的数据（与时间序列数据高度相关）
# 重新排列的数据中，原有索引对应的数据能自动匹配，而新索引缺失的数据通过 NaN 补全。

# 按索引排序, axis 值为 0 或 1, ascending 是否倒序
df_var.sort_index(axis=1, ascending=False)
"""
     F      E  D    C          B    A
0  foo   test  3  1.0 2018-03-01  1.0
1  foo  train  3  1.0 2018-03-01  1.0
2  foo   test  3  1.0 2018-03-01  1.0
3  foo  train  3  1.0 2018-03-01  1.0
"""
# 按值排序 by 的值为 A, B, C, D, E, F，可同时多列
df_var.sort_values(by="E")
"""
     A          B    C  D      E    F
0  1.0 2018-03-01  1.0  3   test  foo
2  1.0 2018-03-01  1.0  3   test  foo
1  1.0 2018-03-01  1.0  3  train  foo
3  1.0 2018-03-01  1.0  3  train  foo
"""
```

### Pandas 数据筛选
```python
df_var = pandas.DataFrame({ 'A' : 1.,
                            'B' : pandas.Timestamp('2018-03-06'),
                            'C' : pandas.Series(1,index=list(range(6)),dtype='float32'),
                            'D' : numpy.array([3] * 6,dtype='int32')})
"""
     A          B    C  D
0  1.0 2018-03-06  1.0  3
1  1.0 2018-03-06  1.0  3
2  1.0 2018-03-06  1.0  3
3  1.0 2018-03-06  1.0  3
4  1.0 2018-03-06  1.0  3
5  1.0 2018-03-06  1.0  3
"""
# 根据数字索引选择
# 使用切片 data_frame.iloc[start:end:step]
df_var.iloc[:3] # 获取前 3 行，相当于 df_var.head(3)
# iloc 除了切片功能，还可以选择多行多列 data_frame.iloc[[row_name], [column_name]]
# 如取 1,3,5 行 df_var.iloc[[1, 3, 5]]
# 如取 2 - 4 列 df_var.iloc[:, 1:4]

# 根据标签名称选择
another_df_var = pandas.DataFrame(numpy.random.randn(6,5),index=list('abcdef'),columns=list('abcde'))
"""
          a         b         c         d         e
a -0.894370  0.441706  2.499390  0.288018 -1.333015
b  0.527419  0.000735 -0.629437 -1.458275  0.502275
c  1.184778 -1.007202 -0.690030  1.401235 -1.094956
d -0.173915  0.219659 -0.162941  0.601745  2.072009
e -0.278881  0.932112 -0.979195  1.920642  1.370296
f  0.686573 -0.553875  1.334577  1.180210 -0.048167
"""
# 获取前 3 行
another_df_var.loc['a':'c']
# loc 与 iloc 类似，数字与字符的区别

# 数据随机取样 
# sample(n=int) n 指返回的数量，默认 1
# sample(frac=float) frac 返回数量的比例，值在 0 ~ 1 区间，比如 .6 为返回 60% 的数量
# 对于 dataframe 的数据，sample 还可指定 axis 参数（0 或 1），其默认值为 0
another_df_var.sample()
"""
          a         b         c         d         e
b  0.527419  0.000735 -0.629437 -1.458275  0.502275
"""

# 条件语句
series_var = pandas.Series(range(-5,5))
series_var[(series_var < -2) | (series_var > 1)]
"""
0   -5
1   -4
2   -3
7    2
8    3
9    4
dtype: int64
"""
another_df_var[(another_df_var['a'] > 0) & (another_df_var['b'] > 0)]
"""
          a         b         c         d         e
b  0.527419  0.000735 -0.629437 -1.458275  0.502275
"""

# 使用 where 条件判断
another_df_var.where(another_df_var < 0)
# .where(df < 0) 会返回所有负值，而非负值默认被置为空值 NaN
# 也可以重新替换非负值如 another_df_var.where(another_df_var < 0, -another_df_var)

# 使用 query 选择
another_df_var.query('(a > 0) & (b > 0)')
# 等同于 another_df_var[(another_df_var['a'] > 0) & (another_df_var['b'] > 0)]
# 但是，query 更简洁且功能强大
```

### Pandas 缺失值处理
```python
# 缺失值主要是指数据丢失的现象，也就是数据集之中的某一块数据不存在
# 除此之外，存在但明显不正确的数据也被归为缺失值一类
# 例如，在一个时间序列数据集中，某一段数据突然发生了时间流错乱，那么，这一小块数据就是毫无意义的，可以被归为缺失值
# 除此之外，当我们使用索引对齐（reindex()）的方法时，也容易人为地导致缺失值的产生

# 检测、标记缺失值
# 不同类型数据的缺失均使用 NaN 标记，例外的是，时间序列里的时间戳丢失使用的是 NaT 标记
df_var = pandas.DataFrame(numpy.random.rand(5, 5), index=list('cafed'),columns=list('ABCDE'))
"""
          A         B         C         D         E
c  0.233613  0.236638  0.832488  0.701375  0.464164
a  0.728444  0.091974  0.236606  0.366514  0.029719
f  0.931431  0.225223  0.212345  0.866862  0.000830
e  0.851676  0.420660  0.643819  0.583977  0.395293
d  0.445159  0.471428  0.205275  0.379091  0.450958
"""
another_df_var = df_var.reindex(list('abcde'))
"""
          A         B         C         D         E
a  0.728444  0.091974  0.236606  0.366514  0.029719
b       NaN       NaN       NaN       NaN       NaN
c  0.233613  0.236638  0.832488  0.701375  0.464164
d  0.445159  0.471428  0.205275  0.379091  0.450958
e  0.851676  0.420660  0.643819  0.583977  0.395293
"""
# 重命名列名 df_var.columns = ['a','b','c', 'd', 'e']

# 判断每项是否为空
another_df_var.isnull()
"""
       A      B      C      D      E
a  False  False  False  False  False
b   True   True   True   True   True
c  False  False  False  False  False
d  False  False  False  False  False
e  False  False  False  False  False
"""
# 判断每项是否不为空another_df_var.notnull()

# 填充缺失值 (不直接改变原有数据)
another_df_var.fillna(0)
"""
          A         B         C         D         E
a  0.728444  0.091974  0.236606  0.366514  0.029719
b  0.000000  0.000000  0.000000  0.000000  0.000000
c  0.233613  0.236638  0.832488  0.701375  0.464164
d  0.445159  0.471428  0.205275  0.379091  0.450958
e  0.851676  0.420660  0.643819  0.583977  0.395293
"""
# 以缺失值前面的值填充，limit 参数设置连续填充的限制数量
another_df_var.fillna(method='pad', limit=1)
"""
          A         B         C         D         E
a  0.728444  0.091974  0.236606  0.366514  0.029719
b  0.728444  0.091974  0.236606  0.366514  0.029719
c  0.233613  0.236638  0.832488  0.701375  0.464164
d  0.445159  0.471428  0.205275  0.379091  0.450958
e  0.851676  0.420660  0.643819  0.583977  0.395293
"""
# 以缺失值后面的值填充 another_df_var.fillna(method='bfill') 若后面的值不存在，继续 NaN 标记
# 以 C 列、E 列的平均值填充 another_df_var.fillna(df.mean()['C':'E'])

# 清除缺失值（不直接改变原有数据）
# 缺失值比较少或者填充无意义时，可以直接清除缺失值
# 默认情况下省略参数 axis=0，即凡是存在缺失值的行均被直接移除
another_df_var.dropna()
"""
          A         B         C         D         E
a  0.728444  0.091974  0.236606  0.366514  0.029719
c  0.233613  0.236638  0.832488  0.701375  0.464164
d  0.445159  0.471428  0.205275  0.379091  0.450958
e  0.851676  0.420660  0.643819  0.583977  0.395293
"""
# 如果根据列移除，another_df_var.dropna(axis=1) 
# 移除所有包含少于 n 个非空值的列 another_df_var.dropna(axis=1,thresh=n) 

# 缺失值插值
# 即借助于一个函数（线性或非线性的函数），再根据已知数据去求解未知数据的值
# 插值的意义是尽量还原数据本来的样子
# 默认线性插值
another_df_var.interpolate()
"""
          A         B         C         D         E
a  0.728444  0.091974  0.236606  0.366514  0.029719
b  0.481029  0.164306  0.534547  0.533945  0.246941
c  0.233613  0.236638  0.832488  0.701375  0.464164
d  0.445159  0.471428  0.205275  0.379091  0.450958
e  0.851676  0.420660  0.643819  0.583977  0.395293
"""
# 如果数据的增长速率越来越快，可以选择 method='quadratic'二次插值
# 如果数据集呈现出累计分布的样子，推荐选择 method='pchip'
# 如果需要填补缺省值，以平滑绘图为目标，推荐选择 method='akima' （需要依赖 Scipy 库）
# 除此之外，插值的方法还有 method='barycentric' （需要依赖 Scipy 库）和 method='pchip' （需要依赖 Scipy 库）
```

### Pandas 时间序列分析
```python
# 时间序列（time series）是实证经济学的一种统计方法
# 它是采用时间排序的一组随机变量，国内生产毛额（GDP）、消费者物价指数（CPI）、股价指数、利率、汇率等等都是时间序列
# 时间序列的时间间隔可以是分秒（如高频金融数据），可以是日、周、月、季度、年、甚至更大的时间单位

# 时间戳
pandas.Timestamp('2018-03-01 13:30:59') # Timestamp('2018-03-01 13:30:59')

# 由时间戳构成的时间索引
pandas.date_range(start=None, end=None, periods=None, freq=’D’, tz=None, normalize=False,
name=None, closed=None, **kwargs)
pandas.date_range('1/3/2018', periods=24, freq='H')
"""
DatetimeIndex(['2018-01-03 00:00:00', '2018-01-03 01:00:00',
               '2018-01-03 02:00:00', '2018-01-03 03:00:00',
               '2018-01-03 04:00:00', '2018-01-03 05:00:00',
               '2018-01-03 06:00:00', '2018-01-03 07:00:00',
               '2018-01-03 08:00:00', '2018-01-03 09:00:00',
               '2018-01-03 10:00:00', '2018-01-03 11:00:00',
               '2018-01-03 12:00:00', '2018-01-03 13:00:00',
               '2018-01-03 14:00:00', '2018-01-03 15:00:00',
               '2018-01-03 16:00:00', '2018-01-03 17:00:00',
               '2018-01-03 18:00:00', '2018-01-03 19:00:00',
               '2018-01-03 20:00:00', '2018-01-03 21:00:00',
               '2018-01-03 22:00:00', '2018-01-03 23:00:00'],
              dtype='datetime64[ns]', freq='H')
"""

# 时间转换
pandas.to_datetime(arg, errors='raise', dayfirst=False, yearfirst=False, utc=None, box=True, 
format=None, exact=True, unit=None, infer_datetime_format=False, origin='unix')
pandas.to_datetime('1/4/2018 10:00', dayfirst=True) # Timestamp('2018-04-01 10:00:00')
pandas.to_datetime(['1/10/2018 10:00','2/10/2018 11:00','3/10/2018 12:00'])
"""
DatetimeIndex(['2018-01-10 10:00:00', '2018-02-10 11:00:00',
               '2018-03-10 12:00:00'],
              dtype='datetime64[ns]', freq=None)
"""
# to_date_time 可以转换 series、dataframe 的时间数据
# 遇到无法解析的数据，参数 errors 就派上用场了
# 参数 errors 包含 raise, ignore, coerce(强制)

# 时间序列检索
time_df_var = pandas.DataFrame(numpy.random.randn(100000,1), columns=['Value'], index=pandas.date_range('20180101', periods=100000, freq='T'))
time_df_var['2018-3-6'] # 检索日期不存在，抛 KeyError 异常
time_df_var['2018-3-6 14:00:00':'2018-3-6 17:00:00']
# 检索方法的参数类似 iloc、loc

# 时间序列计算
from pandas.tseries import offsets
dt = pandas.Timestamp('2018-3-6 10:59:59')
dt + offsets.DateOffset(months=1, days=2, hour=3) # Timestamp('2018-04-08 03:59:59')
dt - offsets.Week(3) # Timestamp('2018-02-13 10:59:59')
# offsets 还有很多方法，参考 https://pandas.pydata.org/pandas-docs/stable/timeseries.html

# 所有时间序列数据，沿着时间轴的方向向后移动
time_df_var.shift(3) # 向前移动则是 time_df_var.shift(-3)
# 向后移动 3 天
time_df_var.shift(3, freq='D')
# 移动索引
time_df_var.tshift(3)

# 重采样
ts = pandas.DataFrame(numpy.random.randn(50,1), columns=['Value' ], index=pandas.date_range('2017-01', periods=50, freq='D'))
# 升频，间隔变成小时，间隔变小，需对新增行填充
ts.resample('H').ffill()
# 降频
ts.resample('5D').sum()
```
