
### 慢日志
```bash
mysql -uroot -p
# 查看是否开启慢日志
show variables like '%slow_query_log%';
# 当前查询，开启慢日志
set global slow_query_log=1;
```

开启慢日志配置
```
# /etc/my.cnf
slow_query_log =1
slow_query_log_file=/var/logs/mysql/slow-queries.log
```

使用 mysqldumpslow 分析慢日志
```bash
# 分析出使用频率（访问次数）最高的前 50 条慢 sql
mysqldumpslow -s c -t 50 /var/logs/mysql/slow-queries.log

# 分析处理速度最慢的 10 条 sql
mysqldumpslow -t 10 /var/logs/mysql/slow-queries.log
```

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
Percona XtraDB Cluster (简称 PXC) 集群是基于 Galera 2.x library，事务型应用下的通用的多主同步复制插件，主要用于解决强一致性问题，使得各个节点之间的数据保持实时同步以及实现多节点同时读写。提高了数据库的可靠性，也可以实现读写分离，是 MySQL 关系型数据库中大家公认的集群优选方案之一。  

安装前的准备
```bash
echo "192.168.1.1 node1" >> /etc/hosts
echo "192.168.1.2 node2" >> /etc/hosts
echo "192.168.1.3 node3" >> /etc/hosts

# 关闭防火墙
service iptables stop

# 安装基本依赖
yum -y install cmake gcc gcc-c++ libaio libaio-devel automake autoconf bzr  ncurses5-devel 
yum -y install perl-DBD-MySQL  perl-DBI  perl-Time-HiRes
```

开始安装
```bash
yum install http://www.percona.com/downloads/percona-release/redhat/0.1-3/percona-release-0.1-3.noarch.rpm
yum install Percona-XtraDB-Cluster-56
```

配置主节点 192.168.1.1  
```bash
# 初始化 
/usr/bin/mysql_install_db --basedir=/usr --user=mysql
# 启动 mysql
service mysql start
# 修改密码
/usr/bin/mysqladmin -u root -h localhost password 'pass'
# 创建用户
mysql -uroot -p
grant reload,lock tables,replication client on *.* to 'sstuser'@'%' identified by 'xxx';
# 关闭 mysql
service mysql stop
```
```
# /etc/my.cnf
[mysql]
user=root
password=pass

[mysqld]
datadir=/var/lib/mysql
user=mysql
server_id=1
wsrep_provider=/usr/lib64/libgalera_smm.so
wsrep_cluster_address="gcomm://192.168.1.1,192.168.1.2,192.168.1.3"
wsrep_sst_auth=wsrep_sst_auth=sstuser:xxx
wsrep_cluster_name=my_pxc_cluster
wsrep_sst_method=rsync
wsrep_node_address=192.168.1.1
wsrep_slave_threads=2
innodb_locks_unsafe_for_binlog=1
innodb_autoinc_lock_mode=2
binlog_format=ROW
```
启动 mysql 及 pxc 服务
```bash
/etc/init.d/mysql bootstrap-pxc
```

配置其他节点
```
[mysql]
user=root
password=pass

[mysqld]
datadir=/var/lib/mysql
user=mysql
server_id=2
wsrep_provider=/usr/lib64/libgalera_smm.so
wsrep_cluster_address="gcomm://192.168.1.1,192.168.1.2,192.168.1.3"
wsrep_sst_auth=wsrep_sst_auth=sstuser:xxx
wsrep_cluster_name=my_pxc_cluster
wsrep_sst_method=rsync
wsrep_node_address=192.168.1.2
wsrep_slave_threads=2
innodb_locks_unsafe_for_binlog=1
innodb_autoinc_lock_mode=2
binlog_format=ROW
```
