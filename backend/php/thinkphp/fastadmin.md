
### FastAdmin 概览
FastAdmin 是一款基于 ThinkPHP5+Bootstrap 的极速后台开发框架。  

[FastAdmin 官网](https://www.fastadmin.net/)  
[FastAdmin Github](https://github.com/karsonzhang/fastadmin)   
[FastAdmin 文档](https://doc.fastadmin.net/docs)  

FastAdmin 主要特性：  

- 基于 Auth 验证的权限管理系统  
- 强大的一键生成功能  
- 完善的前端功能组件开发  
- 强大的插件扩展功能，在线安装卸载升级插件  
- ...  

FastAdmin 目录结构：  
```
├── addons                  //插件存放目录
├── application           //应用目录
│   ├── admin             //后台管理应用模块
│   ├── api               //API应用模块
│   ├── common             //通用应用模块
│   ├── extra             //扩展配置目录
│   ├── index             //前台应用模块
│   ├── build.php
│   ├── command.php        //命令行配置
│   ├── common.php         //通用辅助函数
│   ├── config.php         //基础配置
│   ├── database.php       //数据库配置
│   ├── route.php          //路由配置
│   ├── tags.php           //行为配置
├── extend
│   └── fast               //FastAdmin扩展辅助类目录
├── public
│   ├── assets
│   │   ├── build            //打包JS、CSS的资源目录
│   │   ├── css                //CSS样式目录
│   │   ├── fonts            //字体目录
│   │   ├── img
│   │   ├── js
│   │   │   ├── backend
│   │   │   └── frontend     //后台功能模块JS文件存放目录
│   │   ├── libs            //Bower资源包位置
│   │   └── less            //Less资源目录
│   └── uploads                //上传文件目录
│   ├── index.php            //应用入口主文件
│   ├── install.php          //FastAdmin安装引导
│   ├── admin.php            //后台入口文件,强烈建议修改
│   ├── robots.txt
│   └── router.php
├── runtime                    //缓存目录    
├── thinkphp                //ThinkPHP5框架核心目录
├── vendor                    //Compposer资源包位置
├── .bowerrc                //Bower目录配置文件
├── LICENSE
├── README.md
├── bower.json                //Bower前端包配置
├── build.php                    
├── composer.json            //Composer包配置
└── think
```

### FastAdmin 安装与配置
```bash
# 克隆 FastAdmin 到本地
git clone https://gitee.com/karson/fastadmin.git ProjectName
# 目录权限设置
chown www:www ProjectName -R
chmod 655 ProjectName -R
chmod u+w ProjectName/runtime -R
chmod u+w ProjectName/public/uploads -R
cd ProjectName
# 下载前端插件依赖包
bower install
# 下载 PHP 依赖包
composer install
# 修改配置文件
cp .env.sample .env
vim .env # 配置数据库信息（开发环境开启调试追踪模式）
# 一键创建数据库并导入数据
php think install
# Admin url:http://www.yoursite.com/oJBjCPnRxw.php
# Admin username:admin
# Admin password:xfVW96nyJR
```

添加虚拟主机并绑定到项目的 public 目录（为了安全，安装完成后会在 public 目录生成随机后台入口，请通过随机后台入口登录管理后台）  
```
server {
    listen 80;
    listen 443 ssl http2;
    server_name .fastadmin.test;
    root "/home/vagrant/code/main/fastadmin/public";
    index  index.html index.htm index.php;

    charset utf-8;

    location / {
        if (!-e $request_filename) {
           rewrite  ^(.*)$ /index.php?s=/$1 last;
           break;
        }
    }

    location ~ ^/(uploads|assets)/.*\.(php|php5|jsp)$ {
        deny all;
    }

    location = /favicon.ico { access_log off; log_not_found off; }
    location = /robots.txt  { access_log off; log_not_found off; }

    access_log off;
    error_log  /var/log/nginx/fastadmin.test-error.log error;

    sendfile off;

    client_max_body_size 100m;

    location ~ \.php(.*)$ {
        fastcgi_pass unix:/var/run/php/php7.2-fpm.sock;
        fastcgi_index index.php;
        fastcgi_split_path_info ^((?U).+\.php)(/?.+)$;
        include fastcgi_params;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        fastcgi_param  PATH_INFO  $fastcgi_path_info;

        fastcgi_intercept_errors off;
        fastcgi_buffer_size 16k;
        fastcgi_buffers 4 16k;
        fastcgi_connect_timeout 300;
        fastcgi_send_timeout 300;
        fastcgi_read_timeout 300;
    }

    location ~ /\.ht {
        deny all;
    }

    ssl_certificate     /etc/nginx/ssl/fastadmin.test.crt;
    ssl_certificate_key /etc/nginx/ssl/fastadmin.test.key;
}
```

### FastAdmin 命令行
FastAdmin 基于 ThinkPHP5 强大的命令行功能扩展了一系列命令行功能，可以很方便的一键生成 CRUD、生成权限菜单、压缩打包 CSS 和 JS、安装配置插件等功能。  

```bash
# 创建一个 fa_test 数据表，编辑好表字段结构，并且一定写上字段注释和表注释
# 生成 fa_test 表的 CRUD
php think crud -t test
# 生成 fa_test 表的 CRUD 且一键生成菜单
php think crud -t test -u 1
# 删除 fa_test 表生成的 CRUD
php think crud -t test -d 1
# 生成 fa_test 表的 CRUD 且控制器生成在二级目录下
php think crud -t test -c ProjectName/test
# 生成 fa_test_log 表的 CRUD 且生成对应的控制器为 testlog
php think crud -t test_log -c testlog
# 生成 fa_test 表的 CRUD 且对应的模型名为 testmodel
php think crud -t test -m testmodel
# 生成 fa_test 表的 CRUD 且生成关联模型 category，外链为 category_id，关联表主键为 id
php think crud -t test -r category -k category_id -p id
# 生成 fa_test 表的 CRUD 且所有以 list 或 data 结尾的字段都生成复选框
php think crud -t test --setcheckboxsuffix=list --setcheckboxsuffix=data
# 生成 fa_test 表的 CRUD 且所有以 image 和 img 结尾的字段都生成图片上传组件
php think crud -t test --imagefield=image --imagefield=img
# 关联多个表,参数传递时请按顺序依次传递，支持以下几个参数 relation/relationmodel/relationforeignkey/relationprimarykey/relationfields/relationmode
php think crud -t test --relation=category --relation=admin --relationforeignkey=category_id --relationforeignkey=admin_id
# 生成 v_phealth_db2 数据库下的 fa_test 表的CRUD
php think crud -t test --db=v_phealth_db2

# 一键生成 test 控制器的权限菜单
php think menu -c test
# 一键生成 ProjectName/test 控制器的权限菜单
php think menu -c ProjectName/test
# 删除test控制器生成的菜单
php think menu -c test -d 1
# 一键全部重新所有控制器的权限菜单
php think menu -c all-controller

# 一键压缩打包前后台的 JS 和 CSS
php think min -m all -r all
# 一键压缩打包后台的 JS 和 CSS
php think min -m backend -r all
# 一键压缩打包前后台的 JS
php think min -m all -r js
# 一键压缩打包后台的 CSS
php think min -m backend -r css

# 一键生成 API 文档
php think api --force=true
# 指定 https://www.example.com 为 API 接口请求域名,默认为空
php think api -u https://www.example.com --force=true
# 输出自定义文件为 myapi.html,默认为 api.html
php think api -o myapi.html --force=true
# 修改 API 模板为 mytemplate.html，默认为 index.html
php think api -e mytemplate.html --force=true
# 修改标题为 FastAdmin,作者为作者
php think api -t FastAdmin -a Karson --force=true
# 查看 API 接口命令行帮助
php think api -h
```