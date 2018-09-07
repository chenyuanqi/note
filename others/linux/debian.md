
### 包管理器
apt-get 命令是 Debian Linux 发行版中的 APT 软件包管理工具，所有基于 Debian 的发行都使用这个包管理系统。
```bash
apt-get install xxx #安装软件

apt-get remove xxx ##删除软件

apt-get purge xxx #删除软件并删除配置文件

apt-get autoremove xxx #自动删除未使用的软件

apt-get update #检查软件并升级

apt-get upgrade #升级软件

apt-get clean #清理所有软件缓存

apt-get autoclean #清理旧版本的软件缓存

apt-cahe search xxx #搜索可安装的软件

apt-config dump #打印 apt-get 配置

# Ubuntu 16.04 新添加的命令
apt list #查看已安装的软件
apt edit-sources #快速编辑 apt 源

# 彻底卸载
sudo apt-get autoremove --purge <软件包>
```

