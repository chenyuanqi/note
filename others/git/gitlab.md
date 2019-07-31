
### Gitlab 安装与配置
```bash
# 安装依赖
yum install -y curl policycoreutils-python openssh-server
# 启动 ssh 服务，并且设置为开机启动
systemctl enable sshd
systemctl start sshd
# 安装 firewalld
yum install firewalld
systemctl unmask firewalld
systemctl enable firewalld
systemctl start firewalld
# 开放 ssh、http服务
firewall-cmd --add-service=ssh --permanent
firewall-cmd --add-service=http --permanent
firewall-cmd --reload
# 启用 postfix 邮件服务
yum install postfix
systemctl enable postfix
systemctl start postfix

# 添加 gitlab 社区版 package
curl https://packages.gitlab.com/install/repositories/gitlab/gitlab-ce/script.rpm.sh | sudo bash
yum install -y gitlab-ce
```

配置 Gitlab：/etc/gitlab/gitlab.rb
```
# external_url 'http://gitlab.xxx.com'
```

```bash
# gitlab 重新配置
gitlab-ctl reconfigure
# 重启 gitlab
gitlab-ctl restart

# 查看 gitlab 日志
gitlab-ctl tail

# gitlab 卸载
gitlab-ctl stop
rpm -e gitlab-ce
```

访问 ip/url:port 即可配置第一个用户及密码；
登陆后，创建组-> 创建项目 -> 配置 ssh。

### 
