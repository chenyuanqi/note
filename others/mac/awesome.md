
### Mac 常用软件
-- 浏览器  
chrome：http://www.google.cn/chrome/browser/   

-- 编辑器  
atom: https://atom.io/  
phpstorm: https://www.jetbrains.com/phpstorm/download/#section=mac  

-- 数据库管理  
mysql 客户端: https://sequelpro.com/   

-- 命令行  
Mac-cli：https://github.com/guarinogabriel/mac-cli/  
iterm：http://www.iterm2.com/index.html  
ohmyz: http://ohmyz.sh  

-- 包管理  
brew: http://brew.sh/  
```bash
$ cd /usr/local && git remote set-url origin https://git.coding.net/homebrew/homebrew.git
$ cd $home && brew update
```

-- 软件卸载  
appcleaner：https://freemacsoft.net/appcleaner/

-- IOS 文件管理  
Wondershare TunesGo（需付费）  

-- 其他工具  
印象笔记  
host 管理：https://oldj.github.io/SwitchHosts/  
快捷切换应用 Alfred：https://www.alfredapp.com/  
> 效率工作流参考 1: https://luolei.org/mac-alfred/  
> 效率工作流参考 2: http://www.alfredworkflow.com/  

滚动截屏：http://snip.qq.com  
快捷键查看 cheatsheet: https://www.cheatsheetapp.com/CheatSheet/  
mac 动态图制作：https://github.com/NickeManarin/ScreenToGif  
隐藏菜单栏：vanila  

### Mac 基本操作
-- 同时登陆两个 QQ  
在已经打开的 QQ 窗口中，按住 「command + N」 即可。  

-- MAC 软件（dmg，akp，app）出现程序已损坏的提示  
xxx.app 已损坏，你应该将它移到废纸篓」，并非你安装的软件已损坏，而是 Mac 系统的安全设置问题，因为这些应用都是破解或者汉化的，那么解决方法就是临时改变 Mac 系统安全设置。  
出现这个问题的解决方法：修改系统配置：系统偏好设置... -> 安全性与隐私，修改为任何来源。  
如果没有这个选项的话（macOS Sierra 10.12）,打开终端，执行：  
```bash
sudo spctl --master-disable
```

