
### PHP 的征程
PHP 走过了漫长的道路，成长为处理 web 的最卓越的语言。  
> PHP 发行版可以在[「PHP 博物馆」](http://museum.php.net/) 找到。  

### PHP 7.2 新特性

### PHP 7.1 新特性

### PHP 7.0 新特性

### PHP 5.6 新特性
1、支持常量使用表达式，如 const THREE = TWO + 1;  
2、新增「[... 运算符](http://php.net/manual/zh/functions.arguments.php#functions.variable-arg-list)」，定义函数的可变长参数  
3、支持「[... 运算符](http://php.net/manual/zh/functions.arguments.php#functions.variable-arg-list)」展开参数，如 $operators = [2, 3]; add(1, ...$operators);  
4、新增「[** 幂运算](http://php.net/manual/zh/language.operators.arithmetic.php#115689)」，相当于 pow 函数  
5、支持「[使用命名空间](http://php.net/manual/zh/language.namespaces.importing.php)」use function 及 use const  
6、实现「[交互式调试器 PHPDBG](https://phpdbg.room11.org/)」  
7、修改「[默认编码 default_charset](http://php.net/manual/zh/ini.core.php#ini.default-charset)」默认值为 UTF-8  
8、支持「[php://input](http://php.net/manual/zh/wrappers.php.php#wrappers.php.input)」重用  
9、支持大于 2GB 的文件上传  
10、「[GMP](http://php.net/manual/zh/book.gmp.php)」支持运算符重载  
11、新增「[hash_equals](http://php.net/manual/zh/function.hash-equals.php)」函数，防止时序攻击的字符串比较  
12、新增「[魔术方法 __debugInfo](http://php.net/manual/zh/language.oop5.magic.php#language.oop5.magic.debuginfo)」，控制「[var_dump](http://php.net/manual/zh/function.var-dump.php)」 输出对象的属性和值  
13、支持 gost-crypto 散列算法  
14、提升 SSL/TLS 的支持  
15、「[pgsql](http://php.net/manual/zh/book.pgsql.php)」异步支持  

### PHP 5.5 新特性
1、新增「[生成器  Generator](http://php.net/manual/zh/language.generators.overview.php)」  
2、新增「[异常处理](http://php.net/manual/zh/language.exceptions.php)」关键字 finally  
3、「[foreach](http://php.net/manual/zh/control-structures.foreach.php)」支持「[list](http://php.net/manual/zh/function.list.php)」  
4、「[empty](http://php.net/manual/zh/function.empty.php)」方法支持任意表达式  
5、新增「[新的密码哈希 API](http://php.net/manual/zh/book.password.php)」  
6、支持数组、字符串解引用，如 [1, 2, 3][0] 或 '123'[0]  
7、改进图像处理「[GD 库](http://php.net/manual/zh/book.image.php)」,支持翻转、高级裁剪等图像处理  

### PHP 5.4 新特性
1、支持「[性状 Trait](http://php.net/manual/zh/language.oop5.traits.php)」  
2、支持「短数组语法」，如 $arr = [1, 2, 3];  
3、支持返回值为数组的函数进行成员访问解析，例如 foo()[0]  
4、新增「匿名函数」支持 $this  
5、<?= 将总是可用  
6、新增实例化时访问类成员，如 (new Foo)->foo();  
7、支持语法 Class::{expr}()  
8、内置「[Web Server](http://php.net/manual/zh/features.commandline.webserver.php)」  
9、SESSION 扩展可追踪[上传进度](http://php.net/manual/zh/session.upload-progress.php)  

### PHP 5.3 新特性
1、支持「[命名空间](http://php.net/manual/zh/language.namespaces.php)」  
2、支持「[后期静态绑定](http://php.net/manual/zh/language.oop5.late-static-bindings.php)」  
3、支持「[跳标签 goto](http://php.net/manual/zh/control-structures.goto.php)」  
4、支持「[匿名函数](http://php.net/manual/zh/functions.anonymous.php)」  
5、新增「[魔术方法 __callStatic](http://php.net/manual/zh/language.oop5.overloading.php#language.oop5.overloading.methods)」  
6、新增「[魔术方法 __invoke](http://php.net/manual/zh/language.oop5.magic.php#language.oop5.magic.invoke)」  
7、新增「[Nowdoc 结构](http://php.net/manual/zh/language.types.string.php#language.types.string.syntax.nowdoc)」，类似「[Heredoc 结构](http://php.net/manual/zh/language.types.string.php#language.types.string.syntax.heredoc)」  
8、支持类外部使用 const 关键词声明「[常量](http://php.net/manual/zh/language.constants.syntax.php)」  
9、支持「[三元运算符](http://php.net/manual/zh/language.operators.comparison.php#language.operators.comparison.ternary)」简写“?:”  
10、支持「[异常](http://php.net/manual/zh/language.exceptions.php)」内嵌  
11、新增循环引用的垃圾回收器（默认开启）  
