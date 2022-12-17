
### 环境安装
使用laradock。
```bash
# 安装docker
brew install --cask --appdir=/Applications docker

# 安装laradock
git clone https://github.com/Laradock/laradock.git
cd laradock
cp .env.example .env # 修改php版本
# 启动
docker-compose up -d nginx mysql redis workspace 

# 添加别名
# dockerbash='cd /Users/xxx/Code/laradock;docker-compose exec --user laradock workspace bash'
# dockerdown='cd /Users/xxx/Code/laradock;docker-compose stop'
# dockerenter='cd /Users/xxx/Code/laradock;docker-compose exec workspace bash'
# dockerup='cd /Users/xxx/Code/laradock;docker-compose up -d nginx mysql redis workspace'
```

laradock build错误：
- 0、node 报错  
- 关闭安装 node，本机上使用即可
- 1、0curl: (7) Failed to connect to raw.githubusercontent.com port 443: Connection refused  
- 通过的 www.ipaddress.com/ 查询的 raw.githubusercontent.com IP 地址，添加到 hosts  
- 2、Failed to fetch http://ppa.launchpad.net/ondrej/php/ubuntu/pool/main/p/php-igbinary/php8.0-igbinary_3.2.1+2.0.8-1+ubuntu18.04.1+deb.sury.org+1_amd64.deb  Connection timed out  
```
# 在 workspace 的 Dockfile 中 增加一段
find /etc/apt/sources.list.d/ -type f -name "ondrej-ubuntu-php-xenial.list" -exec  sed  -i.bak -r  's#deb(-src)?\s*http(s)?://ppa.launchpad.net#deb\1 https\2://launchpad.proxy.ustclug.org#ig' {} \; && \
//请注意，ubuntu 18.04 用下面这个
find /etc/apt/sources.list.d/ -type f -name "ondrej-ubuntu-php-bionic.list" -exec sed -i.bak -r 's#deb(-src)?\s*http(s)?://ppa.launchpad.net#deb\1 https\2://launchpad.proxy.ustclug.org#ig' {} \; && \
```

Mac 安装本地开发环境。
```bash
# 安装homebrew
/bin/zsh -c "$(curl -fsSL https://gitee.com/cunkai/HomebrewCN/raw/master/Homebrew.sh)"

# 安装mysql5.7
brew install mysql@latest
# 添加环境变量
echo 'export PATH="/opt/homebrew/opt/mysql@5.7/bin:$PATH"' >> ~/.zshrc
# 启动
mysql.server start 
mysql -u root -p 

# 安装redis，配置/opt/homebrew/etc/redis.conf
brew install redis
# 启动
brew services start redis

# 安装php7.4
brew install php@7.4
# 查看配置
php -ini

# 安装composer
curl -sS https://getcomposer.org/installer | php 
mkdir -p /usr/local/bin 
mv composer.phar /usr/local/bin/composer 
chmod +x /usr/local/bin/composer
# 更新&卸载
composer self-update 
rm /usr/local/bin/composer 

# 安装node
brew install node
# 更新
brew upgrade node 
# 或者：npm install -g n，n latest 
sudo npm install npm@latest -g

# 安装nginx
brew install nginx
# 使用nginx
brew services start nginx #启动
brew services stop nginx #停止
brew services restart nginx #重启
# 查看配置/usr/local/etc/nginx 或 /opt/homebrew/var/www，默认服务路径 /usr/local/var/www 或 /opt/homebrew/var/www
# 默认端口 8080

# 安装sshpass
wget http://sourceforge.net/projects/sshpass/files/sshpass/1.05/sshpass-1.05.tar.gz  
tar xvzf sshpass-1.05.tar.gz  
cd sshpass-1.05
./configure  
sudo make && make install  
```

PHP 升级问题。
```bash
brew update
# 普通升级
brew upgrade php
# 使用 shivammathur/homebrew-php 升级
brew tap shivammathur/php
brew install shivammathur/php/php@8.0

# 切换版本到8.0
brew link --overwrite --force php@8.0
# 切换回7.2
brew unlink php@8.0
brew link php@7.2

```

安装laravel。
```bash
composer config -g repo.packagist composer https://mirrors.aliyun.com/composer/
composer global require laravel/installer
```
确保 Composer 的全局 vendor/bin 目录包含在系统 $PATH 路径中，以便系统可以找到 laravel 可执行命令。在不同的操作系统中，这个目录的位置也有所不同，常见的几种操作系列存放位置罗列如下：
- macOS：$HOME/.composer/vendor/bin（$HOME 表示当前用户家目录，可以用 ~ 替代）  
- Windows：%USERPROFILE%\AppData\Roaming\Composer\vendor\bin（%USERPROFILE% 代表的也是当前用户家目录）  
- GNU/Linux：$HOME/.config/composer/vendor/bin 或者 $HOME/.composer/vendor/bin  

