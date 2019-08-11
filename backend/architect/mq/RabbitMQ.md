
### RabbitMQ
RabbitMQ 是高可靠、高可用的消息中间件。  
RabbitMQ 是基于 Erlang 语言，实现了 AMQP 协议实现的消息队列中间件。  

![RabbitMQ 整体架构](./images/rabbitmq-01.png)  

生产者将消息发送给交换器是需要一个 RoutingKey ，当 BindingKey 和 RoutingKey 匹配时，消息会被路由到对应的队列中。 这里要说明的是 BindingKey 并不是在所有的情况都生效，还依赖交换器的类型。Fanout 类型的交换机会无视 BindingKey，会将消息发送给所有的交换机绑定的所有的队列。  
> 队列（Queue）：队列是 RabbitMQ 的内部对象，用于存储消息的  
> 交换机 （Exchange）: 交换机，生产者将消息发送到 Exchange (交换机），由交换机将消息路由到一个或者多个队列上，或者路由不到，将返回给生产者或者直接丢弃  
> RoutingKey ：路由键，生产者将消息发送给交换机时，一般会指定一个 RoutingKey （路由键），用来指定这个消息的路由规则，这个 RoutingKey 需要与交换机类型和绑定键（ BindingKey）联合使用才能生效  
> Binding ：绑定， RabbitMQ 通过绑定将交换机和队列关联起来，在绑定的时候，一般会指定一个绑定键 （BindingKey），这样 RabbitMQ 就知道如何正确的将消息路由到队列  

### RabbitMQ 安装
```bash
#centeros7 安装 erlang
yum install erlang
#启动扩展源
yum install epel-release
#下载rabbitmq源文件
wget http://www.rabbitmq.com/releases/rabbitmq-server/v3.6.6/rabbitmq-server-3.6.6-1.el7.noarch.rpm
#安装
yum install rabbitmq-server-3.6.6-1.el7.noarch.rpm 
#尝试执行
yum install socat
# 添加开机启动RabbitMQ服务
chkconfig rabbitmq-server on
# 启动服务
/sbin/service rabbitmq-server start
# 查看服务状态
/sbin/service rabbitmq-server status 
# 停止服务
 /sbin/service rabbitmq-server stop
# 查看当前所有用户
 rabbitmqctl list_users
# 查看默认guest用户的权限
 rabbitmqctl list_user_permissions guest
由于RabbitMQ默认的账号用户名和密码都是guest。为了安全起见, 先删掉默认用户
rabbitmqctl delete_user guest
# 添加新用户
rabbitmqctl add_user username password
# 设置用户
rabbitmqctl set_user_tags username administrator
# 赋予用户默认vhost的全部操作权限
rabbitmqctl set_permissions -p / username ".*" ".*" ".*"
# 查看用户的权限
rabbitmqctl list_user_permissions username
#这是打开管理插件的命令.
rabbitmq-plugins enable rabbitmq_management
```

