
### Composer 是什么
Composer 是PHP依赖管理的工具，它允许你声明你的项目所依赖的库并且它将为你管理(安装/更新) 它们。  
Composer 不是包管理器，它用"per project" 这种方式处理包，虽然它提供了一种全局安装选项，但默认情况下它不会在全局安装任何东西。  
基本上, Composer 允许你声明和管理PHP项目中的每个依赖项。  

### Composer 安装与更新
```bash
# 本地安装
php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');"
php -r "if (hash_file('SHA384', 'composer-setup.php') === '544e09ee996cdf60ece3804abc52599c22b1f40f4323403c44d44fdfdd586475ca9813a858088ffbc1f233e9b180f061') { echo 'Installer verified'; } else { echo 'Installer corrupt'; unlink('composer-setup.php'); } echo PHP_EOL;"
php composer-setup.php --install-dir=bin --filename=composer
php -r "unlink('composer-setup.php');"

# 全局安装（建议）
# download page: https://getcomposer.org/download/ （选择要下载的版本）
wget https://getcomposer.org/download/1.6.3/composer.phar
mv composer.phar/usr/local/bin/composer
# 或者这样
curl -s https://getcomposer.org/installer | php

# 更新到最新可用版本
sudo composer self-update # 或 composer self update
# 回滚版本
sudo composer self-update --rollback
```

### Composer 加速
1、去除 xdebug 的警告
```bash
# 找到 php 配置文件的位置
php -i | grep "php.ini"

# 创建 composer 使用的 php 配置
sudo cp php.ini php-composer.ini
sudo vi php-composer.ini
# 在 php-composer.ini 中注释掉这一行
;zend_extension="/your/path/to/xdebug.so"

# 在 ~/.bashrc 中，配置 composer 的别名
vi ~/.bashrc
alias composer="php -c php-composer.ini /usr/local/bin/composer"
# 重新加载 ~/.bashrc
source ~/.bashrc
```

2、使用 prestissimo
```bash
composer global require "hirak/prestissimo"
```

3、使用中国全量镜像
```bash
composer config -g repo.packagist composer https://packagist.laravel-china.org
```

### Composer 的使用
Composer 将所有依赖项下载到项目的 vendor 文件夹中，  
创建 vendor 后，生成 vendor/autoload.php 文件。  
```php
// include 或者 require 到代码文件中, 即可使用依赖包中所提供的所有功能
require __DIR__ . '/vendor/autoload.php';
```

Composer 提供的自动加载对于访问项目的依赖是非常有用的，也可以访问我们自己应用程序的代码。
```php
// 在 composer.json 文件中，使用 composer 的自动加载
{
    "autoload": {
        "psr-4": {"Foo\\": "src/"}
    }
}
```

Composer 默认会阻止不稳定的包被安装，如需允许，执行如下命令配置其最低稳定性。
```bash
composer config minimum-stability dev
```

```bash
# 查看版本
composer --version # 或 composer -V

# 显示给定命令的帮助页面
composer help <command>

# 搜索依赖包
composer search <keyword>

# 显示一组本地被更新的依赖
composer status

# 诊断系统的常见错误
composer diagnose

# 查看配置
composer config [--global] --list
# 修改配置
composer config [--global] <config_name> <config_value>

# 启动向导，创建 composer.json 文件（不建议手动修改）
composer init
# 没有 composer.lock 时，下载 composer.json 相关依赖包，并生成 composer.lock；
# 有 composer.lock 时，下载 composer.lock 中已经注册的依赖及相应版本
composer install
# 更新依赖
composer update <package>:<version>
# 添加依赖包
composer require [--dev] <package>:<version>
# 全局添加依赖包
composer global require <package>:<version>

# 列出所有可用的软件包
composer show
# 列出某包的详细信息
composer show <package>
# 依赖性检测（检测已安装在项目中的某个包，是否正在被其它的包所依赖，并列出他们）
composer depends --link-type=require <package>
# 检测 composer.json 文件是否是有效的
composer validate
# 打印自动加载索引
composer dump-autoload

# 自动克隆仓库，并检出指定的版本 
composer create-project <package> <path> <version>
# 比如创建 laravel 项目
composer create-project --prefer-dist laravel/laravel project_name
```

### 向 Packagist 提交一个 composer package
> 1、注册并登陆 [Packagist](https://packagist.org/)（或使用 GitHub 账号）  
> 2、创建 package 项目，并提交到 GitHub，可参考 [elasticsearch-php](https://github.com/chenyuanqi/elasticsearch)  
> 3、在 [Packagist Submit](https://packagist.org/packages/submit) 填写 package 项目的 GitHub 地址，点击 check  
> 4、要更新 package 项目，除了提交代码 GitHub 之外，到自己的 Packagist 仓库点击 Update 按钮即可  

### Composer 问题及解决
1、[Composer\Downloader\TransportException]  
Your configuration does not allow connections to http://packagist.org/packages.json. See https://getcomposer.org/do  
c/06-config.md#secure-http for details.  

```bash
# 解决方法：换 htt p源，更改配置不要使用 https 加密连接
composer config -g repo.packagist composer http://packagist.phpcomposer.com
composer config -g secure-http false
```

2、Failed to decode response: zlib_decode(): data error
```bash
# 解决方法
composer clear:cache 
composer self-update --update-keys
```
