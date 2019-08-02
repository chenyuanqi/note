
### svn 安装
```bash
yum install -y httpd subversion mod_dav_svn
# 查看 subversion 安装位置
rpm -ql subvserion 
# 查看版本
svnserve --version

# 避免启动告警
echo "export LC_ALL=C" >> /etc/profile

# 卸载
yum remove subversion
```

### svn 代码库创建
```bash
mkdir -p /opt/svn/repo
svnadmin create /opt/svn/repo
```
自动建立 repo 库，查看 /opt/svn/repo 文件夹发现包含了 conf, db,format,hooks, locks, README.txt 等文件，说明一个 SVN 库已经建立。  

### svn 配置
在 /source/svn/repo/conf 目录下有三个文件：  
> passwd ： 里面保存用户信息，基本格式为：user = passwd
> authz ： 里面保存用户的分组信息，以及每个组的访问权限
> svnserve.conf ： 里面保存服务器的基本配置（下面介绍）

在 passwd 文件中添加用户
```
[users]
admin = 123456
vikey = 123456
```

在 authz 文件中配置用户权限
```
[groups]
admin = admin

# 该项目的权限配置
[/]
@admin=rw      # admin分组配置读写权限（实际中可以增加只读分组）
* = r          # 其他用户的权限为只读
```

svnserve.conf 配置
```
[general]
# 匿名用户权限（none：拒绝， write：读写， read：只读权限）
anon-access = none # 控制非鉴权用户访问版本库的权限
auth-access = write # 控制鉴权用户访问版本库的权限
password-db = passwd  # 指定用户名口令文件名
authz-db = authz # 指定权限配置文件名
realm = repo # 指定版本库的认证域，即在登录时提示的认证域名称
```

### svn 启动与停止
```bash
# 防火墙需要开启 3690 端口
firewall-cmd --zone=public --add-port=3690/tcp --permanent
firewall-cmd --zone=public --add-port=3690/udp --permanent
firewall-cmd --reload

# 创建的工程 repo 启动指定目录为 /opt/svn
svnserve -d -r /opt/svn/repo

# 查看 svn
ps aux | grep svnserve

# 停止 svn
killall svnserve 
```

启动后，测试下载项目
```bash
svn checkout svn://server_ip/svn/repo
```

错误解决：  
1、svn: E204900: Can't open file '/opt/svn/repo/format': Permission denied   
解决：执行临时关闭 SELinux 命令 sudo setenforce 0  

2、客户端检出：No repository found in 'svn://xx.xx.xx.xx'  
解决：启动目录是配置的父级目录 svnserve -d -r /opt/svn  

3、Error: Authorization failed
解决：authz 需要配置具体项目权限，否则务必配置 [/]  

### svn 的常用命令
```bash
# 检出仓库，需要输入账号密码
svn co [svn_url] / svn checkout [svn_url]

# 查看状态
svn st / svn status
# ?：不在 svn 的控制中
# M：内容被修改
# C：发生冲突
# A：预定加入到版本库
# K：被锁定

# 查看信息
svn info
# 查看指定文件的详细信息
svn info [file_path]

# 查看变动
svn diff [file_path]

# 添加未追踪文件
svn add [file_path / dir_path]

# 提交改动
svn commit -m "[COMMENT]"

# 删除文件
svn delete [file_path]
svn commit -m "delete file name"

# 还原改动
svn revert [file_path]

# 更新所有提交
svn up / svn update

# 切换分支
svn sw [svn_url] / svn switch [svn_url]

# 建立分支
svn copy [your_trunk_url] [your_feature_branch_url] -m "[COMMENT]"

# 查看指定文件的日志
svn log [file_path]

# 锁定文件
svn lock -m 'commit message' [file_path]
# 解锁文件
svn unlock [file_path]

# 比较文件差异
svn diff [file_path] 
# 比较文件版本 1 和版本 2 的差异
svn diff -r version1:version2 [file_path]  
```

### svn 常用链接
[svn 使用文档](https://tortoisesvn.net/docs/release/TortoiseSVN_zh_CN/index.html)  
[svn 版本控制](http://svnbook.red-bean.com/nightly/zh/svn-book.html)  

[svn 客户端](https://tortoisesvn.net/downloads.zh.html)  


### svn 冲突解决
更新代码，显示冲突
```bash
$ svn update
$ Conflict discovered in [your file]
  Select: (p) postpone, (df) diff-full, (e) edit,
  (mc) mine-conflict, (tf) theirs-conflict,
  (s) show all options:
```
冲突符号说明：

| 符号 | 说明 |
| ---: | :---- |
|p|标记冲突，暂不处理|
|df|显示所有冲突|
|e|编辑冲突|
|mc|冲突以本地文件为准|
|tf|冲突以远程仓库为准|
|s|显示所有选项|

一般先输入df命令看冲突是否严重，如果不严重则通过e直接编辑，编辑页面通常为
```
<<<<<<< .mine
[your_version]
=======
[their_version]
>>>>>>> [version]
```
在<<<<<<< .mine 和 >>>>>>> [version]之间解决冲突，.mine 是自己修改后的代码，version 是别人提交的最新代码，然后保存。

回到 Select 界面，此时会多出一个(r) resolve的命令。输入 r 通知 SVN 已解决冲突。

**使用 postpone 解决冲突**  
如果冲突很严重，需要和提交者讨论解决，可以输入p标记，此时输入svn status显示：
```
C [your_file]
? [your_file].working
? [your_file].merge-left.[version]
? [your_file].merge-right.[version]
```

冲突文件说明：

| 文件 | 说明 |
| ----: | :---- |
|[your_file] | 所有冲突标记在该文件 |
|[your_file].working | 当前工作副本 |
|[your_file].merge-left.[version] | 产生冲突前基础版本 |
|[your_file].merge-right.[version] | 仓库里的最新版本 |

用以下命令解决冲突
```bash
$ svn resolve --accept [base | working | mine-conflict | theirs-conflict | mine-full | theirs-full] [conflicting file] 
```
示例：svn resolve --accept=working readme.txt

| 参数 | 说明 | 
| ----: | :---- |
| base | 将 [your_file].merge-left.[version] 做为最终结果 |
| working | 把 [your_file] 解决冲突后的结果做为最终结果 |
| mine-conflict | 将 [your_file].working 做为最终结果 |
| theirs-conflict | 将 [your_file].merge-right.[version] 做为最终结果 |
| mine-full | 将所有 [your_file].working 做为最终结果 |
| theirs-full | 将所有 [your_file].merge-right.[version] 做为最终结果 |

或者手动解决冲突并执行如下命令
```bash
svn resolved
```

解决冲突后，文件状态变为 M，这时再向仓库提交代码即可。

### svn 合并分支
feature 分支通过测试后就可以合并到 trunk 分支。

首先切换到 trunk 分支，然后执行以下命令
```bash
$ svn merge [your_feature_branch_url]
```

merge 还具有回滚的功能：
```bash
svn merge -r old:new .
```
注意不要少最后一个点，这表示把 new 版本会滚到 old 版本

顺利的话，feature 分支就合并到 trunk 分支了，但是如果有别人和你修改了同一段代码并且提交到 trunk 分支就可能再次出现冲突。同样先解决冲突再提交。

### svn 删除分支
完成功能开发，合并到 trunk 后，删除 feature 分支
```bash
$ svn delete [your_feature_branch_url] -m [your_log]
```

