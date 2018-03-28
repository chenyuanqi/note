
### Linux 文件系统
```html
根文件系统(rootfs)：
    root filesystem

LSB, FHS: (FileSystem Heirache Standard)
    /etc, /usr, /var, /root, /home, /dev

    /boot：引导文件存放目录，内核文件(vmlinuz)、引导加载器(bootloader, grub)都存放于此目录；
    /bin：供所有用户使用的基本命令；不能关联至独立分区，OS启动即会用到的程序；
    /sbin：管理类的基本命令；不能关联至独立分区，OS启动即会用到的程序；
    /lib：基本共享库文件，以及内核模块文件(/lib/modules)；
    /lib64：专用于x86_64系统上的辅助共享库文件存放位置；
    /etc：配置文件目录(纯文本文件)；
    /home/USERNAME：普通用户家目录；
    /root：管理员的家目录；
    /media：便携式移动设备挂载点；
        cdrom
        usb
    /mnt：临时文件系统挂载点；
    /dev：设备文件及特殊文件存储位置；
        b: block device，随机访问
        c: character device，线性访问
    /opt：第三方应用程序的安装位置；
    /srv：系统上运行的服务用到的数据；
    /tmp：临时文件存储位置；
    /usr: universal shared, read-only data；
        bin: 保证系统拥有完整功能而提供的应用程序；
        sbin:
        lib：
        lib64：
        include: C程序的头文件(header files)；
        share：结构化独立的数据，例如doc, man等；
        local：第三方应用程序的安装位置；
            bin, sbin, lib, lib64, etc, share

    /var: variable data files
        cache: 应用程序缓存数据目录；
        lib: 应用程序状态信息数据；
        local：专用于为/usr/local下的应用程序存储可变数据；
        lock: 锁文件
        log: 日志目录及文件；
        opt: 专用于为/opt下的应用程序存储可变数据；
        run: 运行中的进程相关的数据；通常用于存储进程的pid文件；
        spool: 应用程序数据池；
        tmp: 保存系统两次重启之间产生的临时数据；

    /proc: 用于输出内核与进程信息相关的虚拟文件系统；
    /sys：用于输出当前系统上硬件设备相关信息的虚拟文件系统；
    /selinux: security enhanced Linux，selinux相关的安全策略等信息的存储位置；
```

### Linux 应用程序的组成部分
二进制程序：/bin, /sbin, /usr/bin, /usr/sbin, /usr/local/bin, /usr/local/sbin  
库文件：/lib, /lib64, /usr/lib, /usr/lib64, /usr/local/lib, /usr/local/lib64  
配置文件：/etc, /etc/DIRECTORY, /usr/local/etc  
帮助文件：/usr/share/man, /usr/share/doc, /usr/local/share/man, /usr/local/share/doc  

### Linux 文件类型
- (f)：普通文件；  
d: 目录文件；  
b: 块设备；  
c: 字符设备；  
l: 符号链接文件；  
p: 管道文件；  
s: 套接字文件；socket;   
