
### 什么是 Solr
Solr 是一个构建在 Apache Lucene 之上的搜索服务器，它是一个开源、基于 Java 的信息检索库。Solr 查询是简单的 HTTP 请求，响应是结构化文档：主要是 JSON，也可以是 XML、CSV 或其他格式。这意味着任何使用 HTTP 的平台都可以和 Solr 集成。  

**Solr 目录结构说明**  

- contrib 一些功能插件 jar 包
- dist Solr 服务的 jar
- example Solr 提供的几个 Demo
- server\etc Jetty 的配置文件
- server\lib Jetty 的 libraries
- server\logs 默认日志输出目录，可在 resources 中配置
- server\solr-webapp Solr’s Admin UI 操作页面
- server\solr Solr 应用程序的核心目录
- server\solr\configsets 提供的配置集

**Solr 常用命令**  
```bash
# help
solr -help
solr start -help
solr stop -help

# 查看 Solr 运行信息
solr status

# Solr 默认访问端口 8983，相关配置定义在 \server\etc\jetty-http.xml 中
# root 用户启动时需加 -force
solr start -p 8983 -force
solr restart -p 8983 -force
solr stop -p 8983  # 关闭指定端口
solr stop -all # 关闭所有实例
```

**认证授权配置**  
在实际项目应用时，为了安全，我们可能需要对请求进行身份认证，而不是所有人都有权访问。Solr 支持多种方式的认证，像 Basic Authentication、JWT Authentication 等。  
Basic Authentication 认证方式的配置，首先要在 solr-8.2.0/server/solr 文件夹下创建文件 security.json，内容及其含义如下：  
```json
{
 "authentication":{ 
   "blockUnknown":  true, 
   "class": "solr.BasicAuthPlugin",
   "credentials": {"solr":"IV0EHq1OnNrj6gvR......"}, 
   "realm": "My Solr users", 
   "forwardCredentials": false 
 },
 "authorization":{
   "class": "solr.RuleBasedAuthorizationPlugin",
   "permissions": [{"name":"security-edit", "role":"admin"}], 
   "user-role": {"solr":"admin"}
 }
}
```
其中，BasicAuthPlugin & RuleBasedAuthorizationPlugin 启用基本身份认证和基于规则的授权插件。  
blockUnknown：true 表示不允许未经身份验证的请求通过。  
credentials：已定义了一个名为 solr 的用户，默认密码为 SolrRocks。  
permissions：角色 admin 已定义，并且具有编辑安全设置的权限。  
user-role：用户 solr 已被定义为 admin 角色。  
