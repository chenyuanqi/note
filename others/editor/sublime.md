
### sublime link
[sublime 官网](https://www.sublimetext.com/)  

### sublime package control
[package control](https://packagecontrol.io/)  

Package Control 本身是一个为了方便管理插件的插件

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