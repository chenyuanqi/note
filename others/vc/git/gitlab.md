
### Gitlab 安装与配置
```bash
# 安装依赖
yum install -y curl policycoreutils-python openssh-server
# 启动 ssh 服务，并且设置为开机启动
systemctl enable sshd
systemctl start sshd
# 安装 firewalld
yum install -y firewalld
systemctl unmask firewalld
systemctl enable firewalld
# 开放 ssh、http 服务（对应配置文件位于 /etc/firewalld/zones）
firewall-cmd --add-service=ssh --permanent
firewall-cmd --add-service=http --permanent
firewall-cmd --zone=public --add-port=80/tcp --permanent
firewall-cmd --zone=public --add-port=22/tcp --permanent
firewall-cmd --reload
# 查看端口状态
firewall-cmd --zone=public --query-port=80/tcp
# 查看端口访问权限
firewall-cmd --query-port=80/tcp
# 添加端口外部访问
firewall-cmd --add-port=80/tcp
# 查看配置
firewall-cmd --zone=public --list-all
# 启动防火墙
systemctl start firewalld
systemctl status firewalld
# 启用 postfix 邮件服务
yum install -y postfix
systemctl enable postfix
systemctl start postfix

# 添加 gitlab 社区版 package
curl https://packages.gitlab.com/install/repositories/gitlab/gitlab-ce/script.rpm.sh | sudo bash
yum install -y gitlab-ce
```

配置 Gitlab：/etc/gitlab/gitlab.rb
```
# external_url 'http://gitlab.chenyuanqi.com'
```

```bash
# gitlab 重新配置
gitlab-ctl reconfigure
# 重启 gitlab
gitlab-ctl restart

# 重置管理员密码
gitlab-rails console production
# ruby 脚本（用户 1 即 root），其他用户 user = User.find_by(email: 'user@host')
user = User.where(id: 1).first
user.password = '新密码'
user.password_confirmation = '新密码'　
user.save!

# 查看 gitlab 日志
gitlab-ctl tail

# 检查 gitlab
gitlab-rake gitlab:check SANITIZE=true

# 数据库关系升级
gitlab-rake db:migrate
# 清理缓存
gitlab-rake cache:clear

# 更新gitlab包
yum update gitlab-ce

# 升级gitlab
yum install gitlab-ce

# 升级数据命令
gitlab-ctl pg-upgrade

# gitlab 卸载
gitlab-ctl stop
rpm -e gitlab-ce
```

访问 ip/url:port 即可配置第密码（用户 root）；
登陆后，创建组，添加用户，创建项目即可使用。

常用配置 /etc/gitlab/gitlab.rb
```
# 外部访问 GitLab 使用的 URL
external_url 'https://gitlab.chenyuanqi.com'
 
# GitLab 发送 Webhook HTTP 请求后，等待响应的超时
gitlab_rails['webhook_timeout'] = 60
 
# 设置 SSH 主机名
gitlab_rails['gitlab_ssh_host'] = 'gitlab.chenyuanqi.com'
 
# 设置时区
gitlab_rails['time_zone'] = 'Asia/Shanghai'
 
# 运行 GitLab 组件使用的操作系统用户
user['username'] = "gitlab"
user['group'] = "gitlab"
 
# 必须匹配 Unicorn 的监听端口
gitlab_workhorse['auth_backend'] = "http://localhost:2099"
 
# Unicorn 是一个 Ruby Web 服务器，它提高 GitLab 的 Web Interface
unicorn['port'] = 2099
 
 
# 我们通常会在 Unicorn 前面加上 Nginx 作为反向代理
nginx['enable'] = true
nginx['client_max_body_size'] = '250m'
nginx['redirect_http_to_https'] = true
# Nginx 的 HTTP 端口
nginx['redirect_http_to_https_port'] = 2099
# Nginx 的 HTTPS 端口
nginx['listen_port'] = 2443
# TLS Termination 配置
nginx['ssl_certificate'] = "/etc/letsencrypt/live/gmem.cc/cert.pem"
nginx['ssl_certificate_key'] = "/etc/letsencrypt/live/gmem.cc/privkey.pem
 
 
# 禁止普通用户创建组
gitlab_rails['gitlab_default_can_create_group'] = false
# 禁止改变用户名
gitlab_rails['gitlab_username_changing_enabled'] = false

## Default theme
gitlab_rails['gitlab_default_theme'] = 3
##   BASIC  = 1
##   MARS   = 2
##   MODERN = 3
##   GRAY   = 4
##   COLOR  = 5
```


### Gitlab 备份与恢复
修改 /etc/gitlab/gitlab.rb 默认存放备份文件的目录，自动备份就配置 crond
```
gitlab_rails['backup_path'] = '/mnt/backups' 
# 备份只保留 7 天
gitlab_rails['backup_keep_time'] = 604800   
```
```bash
# 重载配置
gitlab-ctl reconfigure 
# 执行备份
gitlab-rake gitlab:backup:create
```

恢复，先进入备份的目录
```bash
# 停止相关数据连接服务
gitlab-ctl stop unicorn
gitlab-ctl stop sidekiq

# 恢复数据
gitlab-rake gitlab:backup:restore BACKUP=1483533591_2017_01_04_gitlab_backup.tar
# 启动Gitlab
sudo gitlab-ctl start  


# 恢复过程中没有权限
mkdir /var/opt/gitlab/backups
chown git /var/opt/gitlab/backups
chmod 700 /var/opt/gitlab/backups

# 恢复成功页面报没有权限的错误
sudo chown -R git:git /var/opt/gitlab/git-data/repositories
sudo chmod -R ug+rwX,o-rwx /var/opt/gitlab/git-data/repositories
sudo chmod -R ug-s /var/opt/gitlab/git-data/repositories
sudo find /var/opt/gitlab/git-data/repositories -type d -print0 | sudo xargs -0 chmod g+s
# 备份文件报没有权限
sudo chown -R git:git 1483533591_2017_01_04_gitlab_backup.tar
```

### Gitlab 使用 Https
下载证书放到 /etc/gitlab/ssh 中，并配置 /etc/gitlab/gitlab.rb
```
nginx['enable'] = true
nginx['redirect_http_to_https'] = true
nginx['redirect_http_to_https_port'] = 80
```
重新配置 gitlab 并重启即可
```bash
# gitlab 重新配置
gitlab-ctl reconfigure
# 重启 gitlab
gitlab-ctl restart
```

### Gitlab 问题
1、Error: Cannot allocate memory  
```bash
# 如果不打算开启休眠功能，物理内存在 8G 以下，则交换内存 swap 设置为与物理内存一样大
# 如果物理内存在 8G 以上，swap 空间设置为 8G 即可
/bin/dd if=/dev/zero of=/var/swap.1 bs=1M count=1024
/sbin/mkswap /var/swap.1
/sbin/swapon /var/swap.1
```

2、Error: 502  
gitlab 重启的几分钟内会报 502，这个可以忽略；  
排查是否端口被占用，若占用修改配置文件 /etc/gitlab/gitlab.rb，再重载配置
```
unicorn['port'] = 9090
```

3、Error: git clone git@gitlab.chenyuanqi.com port 22: Connection refused  
修改配置文件 /etc/gitlab/gitlab.rb  
```
gitlab_rails['gitlab_shell_ssh_port'] = 2099
```
`注意：会暴露服务器 ssh 端口，通常使用 https 也可以满足开发要求`
