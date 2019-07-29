
### Samba 安装
查看系统是否已安装：
```bash
# CentOS：
rpm -qa | grep samba
# Ubuntu：
dpkg -l | grep samba
```
如果未安装：
```bash
# CentOS 6：
yum install samba samba-client samba-common
# Ubuntu：
sudo apt-get install -y samba samba-client
```

### Samba 服务器配置
配置文件位置：/etc/samba/smb.conf  
```bash
# 修改前先备份
cp /etc/samba/smb.conf /etc/samba/smb.conf.backup
sudo vim /etc/samba/smb.conf
```

```
[global]
       # WORKGROUP 表示 Windows 默认的工作组名称，一般共享给 windows 是设置为 WORKGROUP，此字段不重要，无需与 Windows 域保持一致
       workgroup = WORKGROUP  
       # 指定 samba 的安全等级，安全等级分别有四种：share（其他人不需要账号密码即可访问共享目录）、user（检查账号密码）、server（表示检查密码由另外一台服务器负责）、domain（指定 Windows 域控制服务器来验证用户的账号和密码） 
       # 注: samba 4 不再支持 security = share (查看版本 smbd --version)
       # ubuntu 下配置文件默认没有这项，可以自己配置上去
       security = user 
       passdb backend = tdbsam
       printing = cups
       printcap name = cups
       printcap cache time = 750
       cups options = raw
       map to guest = Bad User
       include = /etc/samba/dhcp.conf
       logon path = \\%L\profiles\.msprofile
       logon home = \\%L\%U\.9xprofile
       logon drive = P:
       max connections = 0
       deadtime = 0
       max log size = 500
[share_demo]
       path = /home/<your path> # 分享的目录
       browsable =yes
       writable = yes
       read only = no
       guest ok=no     
       create mask = 0646
       force create mode = 0646
       directory mask = 0747
       force directory mode = 0747
```

### Samba 命令
```bash
# 启动服务（CentOS 6）：
sudo service samba restart
service smb restart 
# 启动服务（CentOS 7）：
systemctl start smb.service # 启动 samba
systemctl enable smb.service # 激活
systemctl status smb.service # 查询 samba 状态（启动 samba 前后可以用查询验证）
 # 启动服务（Ubuntu 16.04.3 -- ljoaquin 提供）：
sudo service smbd restart
```

### Samba 登录失败
可能原因:  
> linux 防火墙  
> Windows 用户密码都正确，错误提示"未知的用户名和密码"时，regedit 打开注册表，删除键值 HKEY_LOCAL_MACHINE\System\CurrentControlSet\Control\Lsa 中的 LMCompatibilityLevel，无需重启计算机
