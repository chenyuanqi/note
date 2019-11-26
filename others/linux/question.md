
### 常见问题

- 文件误删如何恢复？ 
当我们执行 rm -f 删除文件时，其实只是删除了文件的目录索引节点，对于文件系统不可见，但是对于打开它的进程依然可见，即仍然可以使用先前发放的文件描述符读写文件，正是利用这样的原理，所以我们可以使用 I/O 重定向的方式来恢复文件。 
```bash
rm -f /root/selenium/Spider/MySql.Data.dll
# 查看当前是否有进程打开
lsof | grep /root/selenium/Spider/MySql.Data.dll
# 查看是否存在恢复数据
cat /proc/13067/fd/86 # /proc/进程号/fd：进程操作的文件描述符目录；86：文件描述符
# 使用 I/O 重定向恢复文件
cat /proc/23778/fd/86 > /root/selenium/Spider/MySql.Data.dll
```
