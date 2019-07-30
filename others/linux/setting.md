
### ssh 指定登陆用户及 ip
```
# /etc/ssh/sshd_config
PermitRootLogin yes # no 时，限制远程登陆
AllowUsers root@ip

# 重启 ssh
/etc/init.d/sshd restart
```
