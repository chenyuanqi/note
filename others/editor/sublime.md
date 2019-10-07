
### sublime 简介
Sublime Text 是一款具有代码高亮、语法提示、自动完成且反应快速的编辑器软件，不仅具有华丽的界面，还支持插件扩展机制，用她来写代码，绝对是一种享受。相比于难于上手的 Vim，浮肿沉重的 Eclipse，VS，即便体积轻巧迅速启动的 Editplus、Notepad++，在 SublimeText 面前大略显失色，无疑这款性感无比的编辑器是 Coding 和 Writing 最佳的选择，没有之一。

### sublime 链接
[sublime 官网](https://www.sublimetext.com/)  

### sublime package control
[package control](https://packagecontrol.io/)  

优雅使用 Sublime Text，插件则是不可缺少的存在。  
Package Control 本身是一个为了方便管理插件的插件。  

```bash
# install
# ctrl + ` 或  View > show console
# 输入一下代码，回车（可能需要等几分钟）  
import urllib.request,os,hashlib; h = '6f4c264a24d933ce70df5dedcf1dcaee' + 'ebe013ee18cced0ef93d5f746d80ef60'; pf = 'Package Control.sublime-package'; ipp = sublime.installed_packages_path(); urllib.request.install_opener( urllib.request.build_opener( urllib.request.ProxyHandler()) ); by = urllib.request.urlopen( 'http://packagecontrol.io/' + pf.replace(' ', '%20')).read(); dh = hashlib.sha256(by).hexdigest(); print('Error validating download (got %s instead of %s), please try manual install' % (dh, h)) if dh != h else open(os.path.join( ipp, pf), 'wb' ).write(by)
```

安装插件：ctrl + shift + p 选择 Package Control: Install Packge   
```
# 主题 #
# 安装 ayu
# 使用 ayu：ctrl + shift + p > ayu: Activate theme > ayu mirage  

# markdown 预览 #
# 安装 markdownpreview
# 安装 LiveReload
# 
# 配置
# Preferences – Package Settings – Markdown Preview – Setting 的 user 中
{
    "enable_autoreload": true,
}
# ctrl + shift + p > LiveReload: Enable/disable plug-ins > Enable: Simple Reload
# Preferences > Key Bindings-User 设置快捷键
{ "keys": ["ctrl+shift+alt+m"], "command": "markdown_preview", "args": {"target": "browser", "parser":"markdown"} }

#
```

### sublime 技巧

**sublime 定制快捷键**  
[Sublime Text3 快捷键汇总](https://blog.csdn.net/moyan_min/article/details/11530751)  
打开 Preferences -> Key Bindings - User，设置快捷键如下：  
```json
[
    { "keys": ["ctrl+f9"], "command": "build" },
    { "keys": ["f10"], "command": "build", "args": {"variant": "Run"} },
    { "keys": ["ctrl+shift+x"], "command": "toggle_comment", "args": { "block": true } },
]
```

**Sublime 拼写检查**  
如果你经常使用 SublimeText 从事英文创作，那么启用拼写检查就非常有用处了。选择 Preferences > Settings – User 菜单，添加以下代码：
```
"spell_check": true,
```
