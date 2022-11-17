
### pyenv
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
