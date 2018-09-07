
### 包管理器
yum 命令是在 Fedora 和 Red Hat 以及 SUSE 中基于 rpm 的软件包管理器，它可以使系统管理人员交互和自动化地更细与管理RPM软件包，能够从指定的服务器自动下载 RPM 包并且安装，可以自动处理依赖性关系，并且一次安装所有依赖的软体包，无须繁琐地一次次下载、安装。
```bash
yum install xxx #安装软件

yum erase xxx #删除软件
yum -y remove xxx ##删除软件

yum info xxx #查看软件信息

yum list #列出软件列表

yum reinstall xxx #重新安装软件

yum search xxx #搜索软件

yum update #升级所有软件
```
