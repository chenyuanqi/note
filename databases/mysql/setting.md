
### 慢日志


### 主从复制
假设有两台服务器，一台做主，一台做从。  
`注意：主从数据库的版本请保持一致；数据库数据也请保持一致（把主库的数据库复制到从库并导入）；主数据库开启二进制日志，主从数据库的 server-id 必须唯一`  
```
MySQL 主信息：
IP：192.168.1.113
端口：3306
MySQL 从信息：
IP：192.168.1.115
端口：3306
```

- 主库操作步骤
> 1、创建一个目录：mkdir -p /usr/local/mysql/data/mysql-bin  
> 2、主 DB 开启二进制日志功能：vim /etc/my.cnf，添加一行：log-bin = /usr/local/mysql/data/mysql-bin  
> 指定同步的数据库，如果不指定则同步全部数据库，比如指定数据库名：binlog-do-db=demo  
> 3、主库关掉慢查询记录，用 SQL 语句查看当前是否开启：SHOW VARIABLES LIKE '%slow_query_log%'; 如果显示 OFF 则表示关闭，ON 表示开启  
> 4、重启主库 MySQL 服务  
> 5、进入 MySQL 命令行状态，执行 SQL 语句查询状态：SHOW MASTER STATUS; 在显示的结果中，我们需要记录下 File 和 Position 值，等下从库配置有用  
> 6、设置授权用户 slave01 使用 password 密码登录主库，这里 @ 后的 IP 为从库机子的 IP 地址，如果从库的机子有多个，我们需要多个这个 SQL 语句    
```mysql
grant replication slave on *.* to 'slave01'@'192.168.1.135' identified by 'password';
flush privileges;
```

- 从库操作步骤
> 1、从库开启慢查询记录，用 SQL 语句查看当前是否开启：SHOW VARIABLES LIKE '%slow_query_log%'; 如果显示 OFF 则表示关闭，ON 表示开启  
> 2、测试从库机子是否能连上主库机子：mysql -h 192.168.1.105 -u slave01 -p  
> 如果连不上，尝试临时关掉防火墙：service iptables stop  
> 或是添加防火墙规则：  
> 添加规则：iptables -I INPUT -p tcp -m tcp --dport 3306 -j ACCEPT  
> 保存规则：service iptables save  
> 重启 iptables：service iptables restart  
> 3、修改配置文件：vim /etc/my.cnf，把 server-id 改为跟主库不一样
> 4、在进入 MySQL 的命令行状态下，输入下面 SQL：  
```mysql
CHANGE MASTER TO
    master_host='192.168.1.113',
    master_user='slave01',
    master_password='123456',
    master_port=3306,
    master_log_file='主库值_File',
    master_log_pos=主库值_Position;
```
> 5、启动 slave 同步：START SLAVE; 查看从库机子同步状态：SHOW SLAVE STATUS;  
> 在查看结果中必须下面两个值都是 Yes 才表示配置成功：  
> Slave_IO_Running:Yes 如果值为 Connecting，则表示从机连不上主库，需要你进一步排查连接问题  
> Slave_SQL_Running:Yes 如果值为 No，可以检查从库下的错误日志：cat /usr/local/mysql/data/mysql-error.log 如果提示 uuid 错误，请尝试编辑从库的配置文件：/usr/local/mysql/data/auto.cnf的 server-uuid 值保证和主库的值不一样即可  

### 集群


