
### python -m 一句式命令
```bash
# 当前目录开启一个小的文件服务器, 默认端口8000   
python -m SimpleHTTPServer [port]  
# 另外，在 python 3.x 中，命令改成 python -m http.server   

# 开启邮件服务
python -m smtpd -n -c DebuggingServer localhost:20025

# 反重力
python -m antigravity

# 漂亮地格式化打印 json 数据  
echo '{"json":"obj"}' | python -mjson.tool  
# 高亮地打印json格式化  
echo '{"json":"obj"}' | python -mjson.tool | pygmentize -l json  

# 显示当前年份的日历  
python -m calendar  

# python 之禅  
python -m this  
```

### 命令行编写示例
...

### 其他

**生成可执行文件**  
[PyInstaller](http://www.pyinstaller.org/) 能将一个应用程序打包为独立可执行的文件，比如 Windows 下打包为 exe 文件。  
PyInstaller 相比于同类的优势：

- 支持 Python 2.7、Python 3.3-3.6；
- 生成的可执行文件字节数更小；
- 对第三方包的支持非常好，只需要将它们放到 Python 的解释器对应的文件夹中，PyInstaller 便可自动打包到最终生成的可执行文件中。  

```bash
pip install --upgrade pyinstaller

# 单文件模式：-F 打包成一个可执行文件，-w 去掉命令窗口（如果需要看到命令窗口中的错误，建议不要添加 -w）
pyinstaller -F script.py
pyinstaller -w -F script.py

# 将所有文件打包到一个可执行文件
pyinstaller --onefile script.py
```

PyInstaller 打包流程：  
1、梳理程序用到的第三方库有哪些，一定要确保程序用到的 Python 解释器所在的物理安装路径下  
2、将自己的程序代码放到 PyInstaller 的源文件根目录下  
3、执行命令 `pyinstaller script.py`  

