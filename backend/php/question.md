
### 开发问题汇集
- include、include_once 和 require、require_once 的区别
> include 函数：会将指定的文件读入并且执行里面的程序（失败返回 FALSE 并且发出警告；成功返回 1，除非在包含文件中另外给出返回值）；  
> require 函数：会将目标文件的内容读入，并且把自己本身代换成这些读入的内容（失败发出致命错误并终止程序；成功返回 1，除非在包含文件中另外给出返回值）；  
> include_once 函数：在脚本执行期间包含并运行指定文件。此行为和 include 语句类似，唯一区别是如果该文件中已经被包含过，则不会再次包含。如同此语句名字暗示的那样，只会包含一次（失败没有返回且发出警告；成功返回 True）；  
> require_once 函数：和 require 语句完全相同，唯一区别是 PHP 会检查该文件是否已经被包含过，如果是则不会再次包含（失败没有返回且发出致命错误并终止程序；成功返回 True）。  
>
> include 与 require 除了在处理引入文件的方式不同外，最大的区别就是：include 在引入不存文件时产生一个警告且脚本还会继续执行，而 require 则会导致一个致命性错误且脚本停止执行。  
> include 与 require 的功能相同，但在用法上却有一些不同，include 是有条件包含函数，而 require 则是无条件包含函数（即不论条件是否成立都包含到文件里）。  
> 
> include_once（require_once）语句在脚本执行期间包含并运行指定文件；  
> include_once（require_once）行为和 include（require）语句类似，区别是如果该文件中的代码已经被包含了，则不会再次包含，只会包含一次；  
> include_once（require_once）需要查询一遍已加载的文件列表, 确认是否存在, 然后再加载。  
> 通常，include_once 和 require_once 会有性能上的下降，因为他们需要判断之前是否包含过（实际上，也不太需要去考虑，除非这已经影响到你程序的性能了）。  
> 
> 如果使用 APC 尽量也不要用 include_once，因为 apc.include_once_override 配置项没有很好的被实现。  
>（Alternative PHP Cache (APC)是一种对PHP有效的开放源高速缓冲储存器工具，它能够缓存 php 的中间码 opcode）  
> 
> require 通常放在 PHP 程序的最前面，PHP 程序在执行前，就会先读入 require 所指定引入的文件，使它变成 PHP 程序网页的一部份。同时，也可以这个方法将它引入网页中。  
> include 通常放在流程控制的处理部分中。PHP 程序网页在读到 include 的文件时，才将它读进来。这种方式，可以把程序执行时的流程简单化。
> 
> return 语句在返回引用的时候加上括号不但效率低而且会造成逻辑错误。  
> 理论上，include 和 require 后面加不加括号对执行结果没有区别，但是加上括号效率较低，所以后面能不加括号就不加括号。  
> 

- 打开php.ini中的safe_mode，会影响哪些函数
> safe_mode，php安全模式，它提供一个基本安全的共享环境，在一个有多个用户账户存在的php开发的web服务器上  
> 当安全模式打开的时候，一些函数将被完全的禁止，而另一些函数的功能将会受到限制，  
> 如：chdir,move_uploaded_file,chgrp,parse_ini_file,chown,rmdir,copy,rename,fopen,require,mkdir,unlink 等。  
> 
> 注意，在 php5.3 以上版本，safe_mode 被弃用，在 php5.4 以上版本，则将此特性完全去除了。
