
### Matplotlib 是什么
Matplotlib 是一个类似 Matlab 的工具包，是 Python 的一个绘图库。  
Matplotlib 可以创建简单的散点图，正弦曲线，甚至是三维图形。  
Matplotlib 经常用作 Python 的数据可视化。  

### Matplotlib 画正弦曲线
```python
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
plt.figure()
plt.plot(array_var, sin_array_var)
# 也可以画多条曲线 plt.plot(array_var, numpy.sin(array_var), x, numpy.cos(array_var))
plt.show()  
plt.savefig("picture_path/picture_name")
```
