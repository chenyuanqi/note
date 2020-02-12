
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
```bash
pip install --upgrade pyinstaller

# 单文件模式
pyinstaller -F script.py

# 将所有文件打包到一个可执行文件
pyinstaller --onefile script.py
```
