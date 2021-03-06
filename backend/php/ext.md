
### PHP 扩展开发
[PHP Internals Book](http://www.phpinternalsbook.com/index.html) 和 [PHP Internals](https://phpinternals.net/) 两个非官方站点提供稍微全面些的 PHP7 扩展开发的指导。  

更多参考：  
[深入理解 PHP 内核](http://www.php-internals.com/)  

### 为什么需要 PHP 扩展开发

- 当 PHP 的语法特性无法满足我们的要求时。比如在 PHP5.5 之前，没有 Generator，所以如果我们想要在 PHP 中使用协程，就必须在底层实现一个上下文切换的库。
- 当我们需要使用一个 C/C++ 的库时。只有当你充分阅读并理解它的源码以后你才有可能用 PHP 重写这个库，而直接封装成 PHP 扩展你只往往需要理解它暴露的接口就可以了。简单高效。
- 当 PHP 的执行速度真的成为我们项目的性能瓶颈时（比如 yaf 和 swoole 等扩展的存在）。

### PHP 扩展快速入门
1、下载 PHP 源码到本地目录  
解压后进入 PHP 源码的 ext 目录，在此目录下有一个名为 ext_skel 的 shell 脚本文件。接下来我们将使用它来生成我们的扩展的基本骨架。当然，如果你够牛也可以不用它，直接自己编写必要的文件。

2、生成扩展的基本骨架  
在 ext 目录执行如下命令:  
```bash
./ext_skel --extname=foobar
cd ..
vim ext/foobar/config.m4
./buildconf
./configure --[with|enable]-foobar
make
./sapi/cli/php -f ext/foobar/foobar.php
vim ext/foobar/foobar.c
make
```

3、编辑 config.m4 文件  
config.m4 文件是用来配置扩展的行为，比如说明扩展编译选项，是否使用第三方库，扩展的源码组成等等。  
```
// 去掉注释
dnl PHP_ARG_ENABLE(foobar, whether to enable foobar support,
// 去掉注释
nl [  --enable-foobar           Enable foobar support])
```

4、编辑 php_foobar.h 文件，声明一个函数  
php_foobar.h 是一个 C 的头文件，需要在这个头文件里声明一个方法。  
```
PHP_FUNCTION(confirm_foobar_compiled); /* For testing, remove later. */

// 这里声明一个名为 hello 的 PHP 空间的函数，在 PHP 的代码里就可以像普通函数一样调用它
// 我们将在 foobar.c 文件中编写它的函数体。
PHP_FUNCTION(hello);
```

5、编辑 foobar.c 文件，定义函数体  
foobar.c 是扩展主要实现的地方。  
```
const zend_function_entry foobar_functions[] = {
    PHP_FE(confirm_foobar_compiled, NULL)       /* For testing, remove later. */
    // 向 PHP 空间注册一个函数 hello
    PHP_FE(hello, NULL)；
    PHP_FE_END  /* Must be the last line in foobar_functions[] */
};

PHP_FUNCTION(halo){
    php_printf("hello");
}
```

6、编译安装扩展  
扩展编译分动态编译和静态编译两种方法。动态编译需要在扩展目录中执行 phpize 命令（确保安装 php-dev 工具集），一定要在扩展的目录执行才有效，否则将得到一个错误提示。  
```bash
phpize
./configure --enable-foobar --with-config-path=/usr/local/php/bin/php-config
make
# 在 php.ini 中添加 extension=foobar.so
```

7、尝试使用扩展   
```php
hello();
```

