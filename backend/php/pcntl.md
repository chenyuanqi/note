
### 多进程扩展
php 要使用多进程需要安装 pcntl 扩展，在 PHP 中进程控制支持默认是关闭的。您需要使用 --enable-pcntl 配置选项重新编译 PHP 的 CGI 或 CLI 版本以打开进程控制支持。  
pcntl 支持实现了 Unix 方式的进程创建，程序执行，信号处理以及进程的中断。可惜的是，进程控制不能被应用在 Web 服务器环境，当其被用于 Web 服务环境时可能会带来意外的结果。PCNTL 现在使用了 ticks 作为信号处理的回调机制，可以使用 declare () 语句在程序中指定允许发生回调的位置。

**孤儿进程**  
当一个父进程退出，而它的一个或多个子进程还在运行，那么这些子进程将会成为孤儿进程。
换言之，孤儿进程是没有父进程的进程，不会产生什么危害，因为孤儿进程会被 init 进程 (进程号为 1) 所管理，并由 init 进程对它们完成状态收集工作。  

**僵尸进程**  
当子进程退出，而父进程仍在运行，且没有调用 wait 或 waitpid 获取子进程的状态信息，那么子进程的进程描述符仍然保存在系统中，这类进程称之为僵死进程。僵尸进程，会一直占用有限的进程号，当系统的进程号用尽的时候，就不能创建新的进程了。  

### pcntl_fork
使用 pcntl_fork 来 fork 出多个进程来并行执行代码。

pcntl_fork — 在当前进程当前位置产生分支（子进程）。  
注：fork 是创建了一个子进程，父进程和子进程 都从 fork 的位置开始向下继续执行，不同的是父进程执行过程中，得到的 fork 返回值为子进程 号，而子进程得到的是 0。  
```php
if (!function_exists('pcntl_fork')) {
    die("pcntl extention is must !");
}

// fork 后父进程会走自己的逻辑，子进程从处开始走自己的逻辑，堆栈信息会完全复制给子进程内存空间，父子进程相互独立
// fork 首先会执行父进程逻辑再执行子进程逻辑
$pid = pcntl_fork();
if ($pid == -1) {
    die("创建子进程失败!");
} elseif ($pid) { 
    // > 0
    // 父进程逻辑
    pcntl_wait($status); // 等待子进程结束，为了防止子进程变成僵尸进程
} else {
    // 0
    // 子进程逻辑
}
```

多进程可以用来进行大文件的处理，如一个文件有几亿行数据，则可以将文件拆分成多个小文件进程处理。  
假如文件有 10 万行，则可以拆分成 4 个文件，每个文件 2.5 万行，可以使用 split。  
```php
shell_exec('split -l 25000 -d access.log prefix_name')

// 3个子进程处理
for ($i = 0; $i < 3; $i++) {
    $pid = pcntl_fork();
    if ($pid == -1) {
        die("创建子进程失败!");
    } elseif ($pid) {
        // 父进程
    } else {
        $content = file_get_contents("prefix_name0" . $i);
        // 处理逻辑
        exit;
    }
}

while (pcntl_waitpid(0, $status) != -1) {
    $status = pcntl_wexitstatus($status); // 回收子进程
    echo "回收进程" . $status;
}
```
