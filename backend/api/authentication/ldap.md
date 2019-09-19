
### 什么是 LDAP 认证
如果你遇到好多的系统（常见是 GitLab、Bug 平台、Confluence 这些成熟在内部使用的第三方平台），每个系统都要输入用户名密码，输的还都是一样的，基本这些都是基于 LDAP 做的。  

LDAP 英文全称是 Lightweight Directory Access Protocol，意思是轻量目录访问协议，这个协议是基于 TCP/IP 的，意思是它的数据存储是基于目录文件的。  

LDAP 是一个登录认证的协议，在 Web 页面体现的是，N 多系统都要输入相同的用户名密码，多用户内部的第三方平台 GitLab、Confluence、Jira、Nexus。  

LDAP 和 SSO 单点登录最大的差别就是，LDAP 是每个系统都要输入用户名密码，SSO 是跳转到一个专门的页面的登录。  

LDAP 的特点：  

- LDAP 的结构用树来表示，而不是用表格。正因为这样，就不能用 SQL 语句了
- LDAP 可以很快地得到查询结果，不过在写方面，就慢得多
- LDAP 提供了静态数据的快速查询方式
- Client/server 模型，Server 用于存储数据，Client 提供操作目录信息树的工具
- 这些工具可以将数据库的内容以文本格式（LDAP 数据交换格式，LDIF）呈现在您的面前
- LDAP 是一种开放 Internet 标准，LDAP 协议是跨平台的 Interent 协议

### LDAP 基本概念
**Entry**  
Entry 即条目，也叫记录项，是 LDAP 中最基本的颗粒，就像字典中的词条，或者是数据库中的记录。通常对 LDAP 的添加、删除、更改、检索都是以条目为基本对象的。
> dn：每一个条目都有一个唯一的标识名（distinguished Name ，DN），如 "cn=baby,ou=marketing,ou=people,dc=mydomain,dc=org" 。通过 DN 的层次型语法结构，可以方便地表示出条目在 LDAP 树中的位置，通常用于检索。  
> rdn：一般指 dn 逗号最左边的部分，如 cn=baby。它与 RootDN 不同，RootDN 通常与 RootPW 同时出现，特指管理 LDAP 中信息的最高权限用户。  
> Base DN：LDAP 目录树的最顶部就是根，也就是所谓的 “Base DN"，如"dc=mydomain,dc=org"。  

**Attribute**  
每个条目都可以有很多属性（Attribute），比如常见的人都有姓名、地址、电话等属性。每个属性都有名称及对应的值，属性值可以有单个、多个，比如你有多个邮箱。  

属性不是随便定义的，需要符合一定的规则，而这个规则可以通过 schema 制定。比如，如果一个 entry 没有包含在 inetorgperson 这个 schema 中的 objectClass: inetOrgPerson，那么就不能为它指定 employeeNumber 属性，因为 employeeNumber 是在 inetOrgPerson 中定义的。  
LDAP 为人员组织机构中常见的对象都设计了属性 (比如 commonName，surname)。下面有一些常用的别名：  

|属性 | 别名 | 语法 | 描述 | 值 (举例) |
| :--- | :---: | :--- | :--- | :--- |
|commonName | cn | Directory String  |  姓名 | sean |
|surname | sn | Directory String  |  姓 |  Chow |
|organizationalUnitName | ou | Directory String  |  单位（部门）名称    IT_SECTION |
|organization   | 　o | Directory String  |  组织（公司）名称  |  example|
|telephoneNumber 　|   Telephone Number  |  电话号码   | 110|
|objectClass 　 |  　|   内置属性   | organizationalPerson|

**ObjectClass**  
对象类是属性的集合，LDAP 预想了很多人员组织机构中常见的对象，并将其封装成对象类。比如人员（person）含有姓（sn）、名（cn）、电话 (telephoneNumber)、密码 (userPassword) 等属性，单位职工 (organizationalPerson) 是人员 (person) 的继承类，除了上述属性之外还含有职务（title）、邮政编码（postalCode）、通信地址 (postalAddress) 等属性。  

通过对象类可以方便的定义条目类型。每个条目可以直接继承多个对象类，这样就继承了各种属性。如果 2 个对象类中有相同的属性，则条目继承后只会保留 1 个属性。对象类同时也规定了哪些属性是基本信息，必须含有 (Must 活 Required，必要属性)：哪些属性是扩展信息，可以含有（May 或 Optional，可选属性）。  

对象类有三种类型：结构类型（Structural）、抽象类型 (Abstract) 和辅助类型（Auxiliary）。结构类型是最基本的类型，它规定了对象实体的基本属性，每个条目属于且仅属于一个结构型对象类。抽象类型可以是结构类型或其他抽象类型父类，它将对象属性中共性的部分组织在一起，称为其他类的模板，条目不能直接集成抽象型对象类。辅助类型规定了对象实体的扩展属性。每个条目至少有一个结构性对象类。  
对象类本身是可以相互继承的，所以对象类的根类是 top 抽象型对象类。  

**Schema**  
对象类（ObjectClass）、属性类型（AttributeType）、语法（Syntax）分别约定了条目、属性、值。这些构成了模式 (Schema)—— 对象类的集合。条目数据在导入时通常需要接受模式检查，它确保了目录中所有的条目数据结构都是一致的。  
schema（一般在 /etc/ldap/schema/ 目录）在导入时要注意前后顺序。  

**backend & database**  
ldap 的后台进程 slapd 接收、响应请求，但实际存储数据、获取数据的操作是由 Backends 做的，而数据是存放在 database 中，所以你可以看到往往你可以看到 backend 和 database 指令是一样的值如 bdb 。一个 backend 可以有多个 database instance，但每个 database 的 suffix 和 rootdn 不一样。openldap 2.4 版本的模块是动态加载的，所以在使用 backend 时需要 moduleload back_bdb 指令。  

bdb 是一个高性能的支持事务和故障恢复的数据库后端，可以满足绝大部分需求。许多旧文档里（包括官方）说建议将 bdb 作为首选后端服务（primary backend），但 2.4 版文档明确说 hdb 才是被首先推荐使用的，这从 2.4.40 版默认安装后的配置文件里也可以看出。hdb 是基于 bdb 的，但是它通过扩展的索引和缓存技术可以加快数据访问，修改 entries 会更有效率。  

另外 config 是特殊的 backend，用来在运行时管理 slapd 的配置，它只能有一个实例，甚至无需显式在 slapd.conf 中配置。  

**TLS & SASL**  
分布式 LDAP 是以明文的格式通过网络来发送信息的，包括 client 访问 ldap 的密码（当然一般密码已然是二进制的），SSL/TLS 的加密协议就是来保证数据传送的保密性和完整性。  

SASL （Simple Authenticaion and Security Layer）简单身份验证安全框架，它能够实现 openldap 客户端到服务端的用户验证，也是 ldapsearch、ldapmodify 这些标准客户端工具默认尝试与 LDAP 服务端认证用户的方式（前提是已经安装好 Cyrus SASL）。SASL 有几大工业实现标准：Kerveros V5、DIGEST-MD5、EXTERNAL、PLAIN、LOGIN。  

Kerveros V5 是里面最复杂的一种，使用 GSSAPI 机制，必须配置完整的 Kerberos V5 安全系统，密码不再存放在目录服务器中，每一个 dn 与 Kerberos 数据库的主体对应。DIGEST-MD5 稍微简单一点，密码通过 saslpasswd2 生成放在 sasldb 数据库中，或者将明文 hash 存到 LDAP dn 的 userPassword 中，每一个 authid 映射成目录服务器的 dn，常和 SSL 配合使用。  
参考[将 LDAP 客户端配置为使用安全性](http://docs.oracle.com/cd/E19957-01/820-0293/6nc1tbp0h/index.html)。  

EXTERNAL 一般用于初始化添加 schema 时使用，如 ldapadd -Y EXTERNAL -H ldapi:/// -f /etc/openldap/schema/core.ldif。  

**LDIF**  
LDIF（LDAP Data Interchange Format，数据交换格式）是 LDAP 数据库信息的一种文本格式，用于数据的导入导出，每行都是 “属性：值” 对，见 [openldap ldif 格式示例](http://seanlook.com/2015/01/22/openldap_ldif_example/)。  
