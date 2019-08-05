
### Laravel 生命周期
整体流程：  
App 初始化容器（尚未启动）  
-> 绑定 Kernel 内核和异常处理  
-> Kernel 捕捉并处理请求  
-> 响应请求    

Application 初始化：  
初始化容器  
-> Events/Log/Router 注册基础服务  
-> 设置容器中 abstract 与 alias 的映射关系  
![Application 初始化流程](../image/app_start.png)  

Kernel 流程：  
初始化 Kernel（设置中间件）  
-> 捕获请求（路由调度，中间件栈，自定义异常处理）  
-> 处理请求（绑定 request），启动 Application（加载 .env 配置，加载 config 目录配置，设置错误和异常的 handler，设置 Facade 别名自动加载，注册服务提供者，启动服务提供者），使用管道和路由调度把请求通过中间件和路由  
-> 发送响应 -> Kernel 终止  

### Laravel 源码解读基础
