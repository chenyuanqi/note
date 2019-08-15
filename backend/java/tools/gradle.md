
### 构建工具 Gradle
Ant 具有编译、测试和打包功能，其后出现的 Maven 在 Ant 的功能基础上又新增了依赖管理功能，而最新的 Gradle 又在 Maven 的功能基础上新增了对 Groovy 语言的支持。  

Gradle 和 Maven 的区别是，它使用 Groovy 这种特定领域语言（DSL）来管理构建脚本，而不再使用 XML 这种标记性语言。因为项目如果庞大的话，XML 很容易就变得臃肿。  

[Gradle 官网](https://gradle.org/)  

**常用命令**  
```bash
# 构建开发环境，下载定义的 jar 包
gradle eclipse

# 编译项目，生成 build 文件夹，并生成相应的 jar 或 war 包
gradle build

# 清理，删除 build 文件夹
gradle clean
```

### Gradle 安装
```bash
# ubuntu
sudo apt install -y gradle

# macos
brew install gradle

# 查看版本号
gradle -v
```

### Gradle 依赖管理配置
**项目添加 build.gradle**  
```
apply plugin: 'java'
apply plugin: 'eclipse'
sourceCompatibility = 1.5
version = '1.0'
jar {
    manifest {
        attributes 'Implementation-Title': 'Gradle Quickstart', 'Implementation-Version': version
    }
}
repositories {
    mavenCentral()
}
dependencies {
    compile group: 'commons-collections', name: 'commons-collections', version: '3.2'
    testCompile group: 'junit', name: 'junit', version: '4.+'
}
test {
    systemProperties 'property': 'value'
}
uploadArchives {
    repositories {
       flatDir {
           dirs 'repos'
       }
    }
}
```

