
### 什么是 SQL 注入
通过把 SQL 语句插入到 Web 表单提交或输入域名或页面请求的查询字符串，最终达到欺骗服务器执行恶意的 SQL 语句。具体来说，它是利用现有应用程序，将（恶意的）SQL 命令注入到后台数据库引擎执行的能力，它可以通过在 Web 表单中输入（恶意）SQL 语句得到一个存在安全漏洞的网站上的数据库，而不是按照设计者意图去执行 SQL 语句。
```
# SQL 查询代码
strSQL = "SELECT * FROM users WHERE (name = '" + userName + "') and (pw = '"+ passWord +"');"

# 攻击字符
userName = "1' OR '1'='1";
passWord = "1' OR '1'='1";

# 最终结果，相当于 "SELECT * FROM users;"
strSQL = "SELECT * FROM users WHERE (name = '1' OR '1'='1') and (pw = '1' OR '1'='1');"
```

### SQL 注入分类
按照参数类型，SQL 注入分为数字型和字符型。  
> 页面正常返回结果： select 字段名 from 表名 where id = 1 and 1=1;
> 页面返回错误： select 字段名 from 表名 where id = 1 and 1=1;  
> 页面返回错误： select 字段名 from 表名 where id = 1';

根据数据库返回的结果，SQL 注入通常分回显注入（页面中获取返回结果）、报错注入（错误信息直接显示在页面）、盲注（没有直接显示结果也没有报错信息）。  

按照注入位置及方式不同，SQL 注入分为 post 注入，get 注入，cookie 注入，盲注，延时注入，搜索注入，base64 注入。  

### 预防措施
> 1、对用户输入的内容要时刻保持警惕，可以通过正则表达式，或限制长度；对单引号和双 "-" 进行转换等  
> 2、只有客户端的验证等于没有验证  
> 3、永远不要把服务器错误信息暴露给用户（应用的异常信息应该给出尽可能少的提示，最好使用自定义的错误信息对原始错误信息进行包装）  
> 4、不要把机密信息直接存放，加密或者 hash 掉密码和敏感的信息  
> 5、使用 pdo 过滤传入的参数，使用参数化查询  
> 6、永远不要使用动态拼装 sql, 可以使用参数化的 sql 或者直接使用存储过程进行数据查询存取  

把一些 sql 语句进行过滤，比如 delete update insert select * 或者使用 PDO 占位符进行转义

**PDO为什么可以防范 SQL 注入**  
虽然 PHP 可以用 mysql_real_escape_string() 函数过滤用户提交的值，但是也有缺陷。  
PDO 当调用 prepare () 时，查询语句已经发送给了数据库服务器，此时只有占位符？发送过去，没有用户提交的数据；当调用到 execute () 时，用户提交过来的值才会传送给数据库，他们是分开传送的，两者独立的。
```php
$dbHandle = new PDO('mysql:dbname=dbtest;host=127.0.0.1;charset=utf8', 'user', 'password');
$dbHandle->setAttribute(PDO::ATTR_EMULATE_PREPARES, false);
$dbHandle->exec("set names 'utf8'"); 

$sql = "select * from test where name = ? and password = ?";
$statement = $dbHandle->prepare($sql); 
$result = $statement->execute([$name, $password]); 
```
