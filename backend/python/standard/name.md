
```python
# 命名是种艺术
# 要直指要义，不要通用名称
# 用复数形式命名集合变量，不要单数省字符，不要使用已经存在的名称
# 避免命名空间中存在冗余信息，言简意赅更好

# 常量集中放在使用模块的顶部
SQL_USER = 'user_name'  
SQL_PASSWORD = 'secret'  
SQL_URI = 'postgres://%s:%s@localhost/db' % (  
    SQL_USER, SQL_PASSWORD  
)  
MAX_THREADS = 4  

# 公有变量
mysql_connection  

# 私有变量
_base_secret  
__apk_secret  

# 模块或包名，简短、使用小写字母、并且不带下划线
# 比如要实现一个协议
telnetlib

# 类名，使用大驼峰
class SQLEngine(object):
    pass

# 常规属性
tables 
persons_addresses 

# 布尔属性
is_connected = False
has_cache = False

# 公有方法名
get_user_info(user_id=0)

# 私有方法名
_generate_token(token_length=6)

```
