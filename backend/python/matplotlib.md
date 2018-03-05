
### Matplotlib 是什么
Matplotlib 是一个类似 Matlab 的工具包，是 Python 的一个绘图库。  
Matplotlib 可以创建简单的散点图，正弦曲线，甚至是三维图形。  
Matplotlib 经常用作 Python 的数据可视化。  

### Matplotlib 安装
```bash
sudo apt-get install python-matplotlib
# 或者使用 pip 安装
pip install matplotlib
```

### Matplotlib 画正弦曲线
```python
# 简易版
import numpy as np
from matplotlib import pylot as plt

x = np.linspace(-2 * np.pi, 2 * np.pi, 1000)

plt.plot(x, np.sin(x))
plt.show()
```

```python
# 啰嗦版
# 生成等间隔的数据 linspace(start, step, length)
import numpy
array_var = numpy.linspace(0, 2 * numpy.pi, 21)
"""
array([0.        , 0.31415927, 0.62831853, 0.9424778 , 1.25663706,
       1.57079633, 1.88495559, 2.19911486, 2.51327412, 2.82743339,
       3.14159265, 3.45575192, 3.76991118, 4.08407045, 4.39822972,
       4.71238898, 5.02654825, 5.34070751, 5.65486678, 5.96902604,
       6.28318531])
"""
# 绘制正弦曲线
import matplotlib
matplotlib.use("pdf")
import matplotlib.pyplot as plt

# 解决中文乱码问题
ttc_font = matplotlib.font_manager.FontProperties(fname="ttc_file_of_path", size=14)
matplotlib.rcParams["axes.unicode_minus"] = False

sin_array_var = numpy.sin(array_var)
"""
array([ 0.00000000e+00,  3.09016994e-01,  5.87785252e-01,  8.09016994e-01,
        9.51056516e-01,  1.00000000e+00,  9.51056516e-01,  8.09016994e-01,
        5.87785252e-01,  3.09016994e-01,  1.22464680e-16, -3.09016994e-01,
       -5.87785252e-01, -8.09016994e-01, -9.51056516e-01, -1.00000000e+00,
       -9.51056516e-01, -8.09016994e-01, -5.87785252e-01, -3.09016994e-01,
       -2.44929360e-16])
"""
# 需安装 tk 库：sudo apt-get install python3-tk  
# 生成画布
plt.figure(figsize=(8, 6), dpi=80)

# 设置标题
plt.title("正弦曲线", fontproperties=ttc_font)
# 显示网格
plt.grid(True)
# 声明图例，使显示正常
plt.legend()

# 设置X轴
plt.xlabel("X轴", fontproperties=ttc_font)
plt.xlim(-4.0, 4.0)
plt.xticks(numpy.linspace(-4, 4, 9, endpoint=True))

# 设置Y轴
plt.ylabel("Y轴", fontproperties=ttc_font)
plt.ylim(-1.0, 1.0)
plt.yticks(numpy.linspace(-1, 1, 9, endpoint=True))

# 绘制曲线，plt.plot(x轴数据, y轴数据, 曲线类型,图例说明,曲线线宽)
plt.plot(array_var, sin_array_var)
# 也可以画多条曲线 plt.plot(array_var, numpy.sin(array_var), x, numpy.cos(array_var))
# 显示图形
plt.show()  
# 保存图片，默认 png 格式
plt.savefig("picture_path/picture_name")
```

### Matplotlib 绘制 3D 曲面图
```python
# 载入模块
import numpy as np
import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D

# 生成画布
axes_figure = plt.figure()
# 创建 3D 图形对象
axes_instance = Axes3D(axes_figure)

# 生成 x, y 轴数据
x = np.arange(-2, 2, 0.1)
y = np.arange(-2, 2, 0.1)
# 编织 x, y 轴数据
x, y = np.meshgrid(x, y)
# 计算 z 轴数据
z = np.sqrt(x ** 2 + y ** 2)

# 绘制曲面图，并使用 cmap 着色
axes_instanceax.plot_surface(x, y, z, cmap=plt.cm.winter)
# 显示图形
plt.show()
```

### Matplotlib 高级 API 绘图库 Seaborn
Matplotlib 拥有 3000 多页的官方文档，上千个方法以及数万个参数，其复杂程度使得开发者常常伤透脑筋。  
Seaborn 基于 Matplotlib 核心库进行了更高级的 API 封装，可以让你轻松地画出更漂亮的图形。  
Seaborn 的漂亮主要体现在配色更加舒服、以及图形元素的样式更加细腻。  

```bash
# 安装 Seaborn
pip install seaborn
```

尝试绘制热力图
```python
# 载入绘图模块
import numpy as np
import matplotlib.pyplot as plt
import seaborn

# 生成 10x10 的随机矩阵
matrix_data = np.random.rand(10, 10)

# 绘制 heatmap
seaborn.heatmap(data=matrix_data)
# 显示图片
plt.show()
```

更多，请查阅[官网](http://seaborn.pydata.org/index.html)  