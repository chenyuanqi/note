
### Spring boot
从本质上来说，Spring Boot 就是 Spring, 它做了那些没有它你也会去做的 Spring Bean 配置。  
Spring Boot 是 Spring 官方发展十几年后推出的重量级产品，具有快速开发、快速部署、方便配置、便于监控等特性，这些特性将会重构整个研发流程、提升研发效率，达到快速开发、交付的目的，同时也让 Spring Boot 成为 Java 领域最佳微服务架构落地技术。  

Spring 将很多魔法带入了 Spring 应用程序的开发之中:  
1、自动配置：针对很多 Spring 应用程序常见的应用功能，Spring Boot 能自动提供相关配置  
2、起步依赖：告诉 Spring Boot 需要什么功能，它就能引入需要的库  
3、命令行界面：这是 Spring Boot 的可选特性，借此你只需写代码就能完成完整的应用程序，无需传统项目构建  
4、Actuator：让你能够深入运行中的 Spring Boot 应用程序，一套究竟  

**初始化 Spring Boot 项目**  
由 IDE 快速创建（在 File 菜单里面选择 New > Project, 然后选择 Spring Initializr），并新建一个 HelloController.java
```java
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;

// @RequestMapping 也可以用作整个类 url 的前缀
@RestController
public class HelloController.java
{
    // 如果是多个 url
    // @RequestMapping(value = {"/hello", "/hi"}, method = RequestMethod.GET)
    // 不使用 method 参数将默认支持所有（不建议）
    @RequestMapping(value = "/hello", method = RequestMethod.GET)
    public String sayHello()
    {
        return "Hello, spring boot";
    }
}
```

Spring boot 启动方式：  
1、使用 ide 启动，运行 Application.java 文件即可  
2、命令行执行 mvn spring-boot:start  
3、先编译生成 jar 包，再运行包  
```bash
mvn install # 生成 jar_path
mvn -jar jar_path
```

**spring loaded 实现热部署**  
在 pom.xml 添加依赖
```xml
<dependency>
    <groupId>org.springframework</groupId>
    <artifactId>springloaded</artifactId>
    <version>1.2.5.RELEASE</version>
</dependency>
```
之后通过 mvn spring-boot:run 启动就支持热部署了。  

### Spring boot 项目配置
项目配置在 src/main/resources/application.properties，可以添加属性及自定义属性
```
spring.datasource.url = jdbc:mysql://127.0.0.1:3306/
spring.datasource.username = root
spring.datasource.password = 123456
spring.datasource.driver-class-name = com.mysql.jdbc

server.port = 80
server.servlet.context-path = /app

custom.project.name = "xxx"
custom.project.timezone = "xxx"
custom.project.reference = ${custom.project.name} 所在时区 ${custom.project.timezone}
```
在使用的地方通过注解 \@Value(value="${config.name}") 就可以绑定到你想要的属性上面
```java
@Value("${custom.project.name}")
private String name;
```

如果属性太多，我们还可以定义一个 ConfigBean.java 类  
```java
@Component
@ConfigurationProperties(prefix = "custom.project")
public class ConfigBean 
{
    private String name;
    private String timezone;

    // 不能省略 getter 和 setter
    public String getName() 
    {
        return name;
    }

    public void setName(String name) 
    {
        this.name = name;
    }

    public String getTimezone() 
    {
        return timezone;
    }

    public void setTimezone(String timezone) 
    {
        this.timezone = timezone;
    }
}
```
同时，在 Spring boot 入口类加上 \@EnableConfigurationProperties 并指明要加载哪个 bean。  
```java
@SpringBootApplication
@EnableConfigurationProperties({ConfigBean.class})
public class Application 
{
    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }
}
```
在使用它的 Controller 引入 ConfigBean 即可  
```java
@RestController
public class SiteController 
{
    @Autowired
    ConfigBean configBean;

    @RequestMapping("/")
    public String home()
    {
        return configBean.getName() + configBean.getTimezone();
    }
}
```

**多环境配置**  
通常，我们会存在多个环境的配置文件：  
1、application-dev.properties：开发环境  
2、application-test.properties：测试环境  
3、application-prod.properties：生产环境  
在 application.properties 通过 spring.profiles.active 切换环境
```
spring.profiles.active = dev
```

**配置文件优先级**  
1、src/main/resources/config > src/main/resources > Classpath 根目录  
2、相同优先级位置上，application.yml > application.properties  

**命令行使用项目配置**  
```bash
java -jar xx.jar --server.port=9090 --spring.profiles.active=dev
```

**推荐使用 yaml 配置文件**
```yaml
spring:
    datasource:
        url: jdbc:mysql://127.0.0.1:3306/dbName
        username: root
        password: 123456
        driver-class-name: com.mysql.jdbc

server:
    port: 80
    servlet:
        context-path: /app

custom:
    project:
        name: xxx
        timezone: xxx
        reference: "${custom.project.name} 所在时区 ${custom.project.timezone}"
```

### Spring boot 控制器

| 注解 | 说明 |  
| :--- | :--- |  
| \@Controller | 处理 http 请求 |  
| \@RestController | Spring 4 之后新加的注解，原返回 json 需要 \@ResponseBody 配合 \@Controller | 
| \@RequestMapping | 配置 url 映射 |  
| \@GetMapping | 组合注解（即缺省 method 是 get） |  
| \@PathVariable | 获取 url 中的数据 |  
| \@RequestParam | 获取请求参数的值 |  
| \@RequestBody | 获取 body 参数的值 |  
| \@RequestHeader | 获取请求头参数 |  
| \@CookieValue | 获取 cookie 的值 |

```java
import org.springframework.web.bind.annotation.*;

@RestController
public class HelloController {
    @RequestMapping(value = "/say/{id}", method = RequestMethod.GET)
    public String say(@PathVariable("id") Integer id)
    {
        return "id: " + id;
    }

    @RequestMapping(value = "/getParam", method = RequestMethod.GET)
    public String getParam(@RequestParam(value="id", required=false, defaultValue = "0") Integer id)
    {
        return "id: " + id;
    }

    @RequestMapping(value = "/postParam1", method = RequestMethod.POST)
    public String getParam(@RequestBody Person person)
    {
        return "body: " + person.toString();
    }

    @RequestMapping(value = "/postParam2", method = RequestMethod.POST)
    public String getParam(@RequestBody Map<String, String> person)
    {
        return "body-name: " + person.get("name");
    }
}
```
控制器获取参数还可以使用 HttpServletRequest 来获取，get 和 post 都可以。

### Spring boot 视图
如果是静态页面，src/main/resources 下的目录 static、public、resources、META-INF/resources 都会将静态访问 (html / 图片等) 映射到其自动配置的静态目录。  
```java
// 静态文件在 src/main/resources/static/index.html
@Controller
public class HtmlController 
{
    @GetMapping("/templates")
    public String html() 
    {
        return "index.html";
    }
}
```

如果是动态页面，使用 Thymeleaf 组件。在 pom.xml 中配置
```xml
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-thymeleaf</artifactId>
</dependency>
```
接下来使用它
```java
@Controller
public class TemplatesController 
{  
    @GetMapping("/templates")
    public String test() 
    {
        return "index";
    }  
}
```

### Spring boot 数据库操作
Mysql 的处理使用 Spring-Data-Jpa。  
JPA（Java Persistence APl）定义了一系列对象持久化的标准，目前实现这一规范的产品有 Hibernate、TopLink 等。  

**使用数据库前的准备**  
1、在 pom.xml 添加如下依赖  
```xml
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-data-jpa</artifactId>
</dependency>

<dependency>
    <groupId>mysql</groupId>
    <artifactId>mysql-connector-java</artifactId>
</dependency>
```

2、配置文件 application.yml 添加如下配置  
```yaml
spring:
    datasource:
        driver-class-name: com.mysql.cj.jdbc.Driver
        url: jdbc:mysql://192.168.10.10:3306/dbName
        username: homestead
        password: secret
    jpa:
        hibernate:
            ddl-auto: create # 默认 none 什么都不做，create 新建新表，create-drop 应用开始时创建表停下时删除表，update 不存在就新建存在就保留，validate 验证已存在表对应的类中属性是否与表结构一致（不一致会报错）
            show-sql: true
```

3、创建 Model 实体类  
```java
import javax.persistence.*;

@Entity
@Table(name = "t_user")
public class User
{
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Integer id;

    @Column(name = "user_name")
    private String userName;

    private String password;

    public Integer getId()
    {
        return id;
    }

    public void setId(Integer id)
    {
        this.id = id;
    }

    public String getUserName()
    {
        return userName;
    }

    public void setUserName(String userName)
    {
        this.userName = userName;
    }

    public String getPassword()
    {
        return password;
    }

    public void setPassword(String password)
    {
        this.password = password;
    }
}
```

**Restful api**  
1、创建 UserRepository  
```java
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface UserRepository extends JpaRepository<User, Integer>
{
    public List<User> findByAge(Integer age);
}
```

2、创建 UserController，实现 crud  
```java
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;
import java.util.List;
import java.util.Optional;

@RestController
public class UserController
{
    @Autowired
    private UserRepository userRepository;

    @GetMapping(value = "/users")
    public List<User> lists()
    {
        return userRepository.findAll();
    }

    @GetMapping(value = "/users/{id}")
    public Optional<User> findUserById(@PathVariable("id") Integer id)
    {
        return userRepository.findById(id);
    }

    @PostMapping(value = "/users")
    public User insert(
            @RequestParam("userName") String userName,
            @RequestParam("password") String password
    ) {
        User user = new User();

        user.setUserName(userName);
        user.setPassword(password);

        return userRepository.save(user);
    }

    @PutMapping(value = "/users/{id}")
    public User insert(
            @PathVariable("id") Integer id,
            @RequestParam("userName") String userName,
            @RequestParam("password") String password
    ) {
        User user = new User();

        user.setId(id);
        user.setUserName(userName);
        user.setPassword(password);

        return userRepository.save(user);
    }

    @DeleteMapping(value = "/users/{id}")
    public void delete(@PathVariable("id") Integer id)
    {
        userRepository.deleteById(id);
    }
}
```
3、如果需要使用事务，使用注解 \@Transactional 即可  

### Spring boot 配合 Nginx
Spring 应用作为服务运行，NGINX 代理允许将应用部署到非特权端口。  
```
server {        
   listen 80;        
   listen [::]:80;  
           
   server_name example.com;    
         
   location / {             
           proxy_pass http://localhost:8080/;              
           proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;              
           proxy_set_header X-Forwarded-Proto $scheme;              
           proxy_set_header X-Forwarded-Port $server_port;         
   } 
}
```
