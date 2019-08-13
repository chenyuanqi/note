
### Maven 是什么
Maven 是基于项目对象模型（Project Object Model），可以通过一小段描述信息来管理项目的构建、报告和文档的项目管理工具，提供了一个仓库的概念，统一管理项目所依赖的第三方 jar 包，最大可能避免了由于环境变量的不同在不同电脑之间无法运行的问题。  

Maven 可以在不同的项目之间有效进行项目的管理。  

[Maven 官网](http://maven.apache.org/)  

### Maven 安装使用
**Maven 安装**  
```bash
sudo apt -y install maven
```

**pom.xml 基本元素**  
1、groupId 项目组，和包名一致  
2、artifactId 标识项目中的模块，格式建议使用【项目名称-模块名称】  
3、version 版本号，快照版本是一种特殊的版本  
4、modelVersion 固定值 4.0.0  
5、maven 路径约定：java 文件放在 src/main/java 目录，test 文件放在 src/test/java 目录  
```
<!---pom.xml 示例--->
<modelVersion>4.0.0</modelVersion>  
<groupId>com.vikey.maven.hello</groupId>  
<artifactId>hello-test</artifactId>  
<version>0.0.1-SNAPSHOT</version> 
```

**Maven 常用命令**  
```bash
mvn --version # 或 mvn -v 查看 maven 版本
mvn compile # 编译项目
mvn test # 执行单元测试
mvn clean # 清除产生的项目，默认清理 target 目录中的数据
mvn package # 打包
mvn install # 打包，并将包放入本地仓库中
```

### Maven 生命周期
maven 有 3 个独立的阶段：clean、default、site。  
其中，clean 清理包括 pre-clean、clean、post-clean；default 核心的构建包括 compile、test、package、install；site 生成项目站点包括 pre-site、site（生成项目的站点文档）、post-site、site-deploy（发布站点到服务器）。


### Maven 常见问题

