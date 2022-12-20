
### Python 重生
[自学是门手艺](https://github.com/selfteaching/the-craft-of-selfteaching)  
[100 天 python](https://github.com/jackfrued/Python-100-Days)   
[python 问答](https://github.com/wklken/stackoverflow-py-top-qa)  

**环境配置 pyenv**  
```bash
# 安装
brew install pyenv
# 版本
pyenv -v

# 查看可安装版本
pyenv install --list
# 安装指定版本
pyenv install 3.10.2
pyenv rehash # 刷新
# 卸载指定版本
pyenv uninstall x.x.x
# 查看所有安装版本
pyenv versions

# 切换版本
pyenv global 3.7.3 #全局切换
pyenv local 3.7.3 # 当前所在文件夹切换
```

用 pyenv versions 查看，明明已经切换成功，但是用 python -V 却还是系统版本。原因是 pyenv 没有加到 $PATH 环境变量里去，解决办法如下：
```
export PYENV_ROOT=~/.pyenv
export PATH=$PYENV_ROOT/shims:$PATH
```

**安装 jupyter**  
```bash
pip install jupyterlab
jupyter lab
```

**文件操作**  
```python
file = open('test.txt', 'w')
file.write('1')
print("文件当前指针：". file.tell())
print('回到开头')
file.seek(0)
print("文件当前指针：". file.tell())
file.close()

with open('test.txt') as f:
	f.readline()
``` 

**异常处理**  
```python
try:
	print(i)
	raise ValueError('xxx')
except NameError:
    print('NameError')
except Exception:
	print('Exception')
finally:
	i = 2022
```

**函数**  
```python
def sum(num1, num2):
	return num1 + num2
print(sum(1, 1))
```


