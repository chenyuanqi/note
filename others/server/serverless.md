
### 什么是 Serverless
Serverless 是由事件（event）驱动（e.g. http,pub/sub）的全托管计算服务。  
Serverless（无服务器架构）是指服务端逻辑由开发者实现，运行在无状态的计算容器中，由事件触发，完全被第三方管理，其业务层面的状态则存储在数据库或其他介质中。  
就像无线互联网实际有的地方也需要用到有线连接一样，无服务器架构仍然在某处有服务器。开发者无需关注服务器，只需关注代码即可。  

Serverless 相对于 serverful，对业务用户强调 noserver（serverless 并不是说没有服务器，只是业务人员无需关注服务器了，代码仍然是运行在真实存在的服务器上）的运维理念，业务人员只需要聚焦业务逻辑代码。

- 弱化了存储和计算之间的联系。服务的储存和计算被分开部署和收费，存储不再是服务本身的一部分，而是演变成了独立的云服务，这使得计算变得无状态化，更容易调度和扩缩容，同时也降低了数据丢失的风险。  
- 代码的执行不再需要手动分配资源。不需要为服务的运行指定需要的资源（比如使用几台机器、多大的带宽、多大的磁盘等），只需要提供一份代码，剩下的交由 serverless 平台去处理就行了。当前阶段的实现平台分配资源时还需要用户方提供一些策略，例如单个实例的规格和最大并发数，单实例的最大 cpu 使用率。理想的情况是通过某些学习算法来进行完全自动的自适应分配。  
- 按使用量计费。Serverless 按照服务的使用量（调用次数、时长等）计费，而不是像传统的 serverful 服务那样，按照使用的资源（ECS 实例、VM 的规格等）计费。

Serverless 是云原生技术发展的高级阶段，可以使开发者更聚焦在业务逻辑，而减少对基础设施的关注。

**Serverless 的分类**  
Serverless 通常包含了两个领域 BaaS（Backend as a Service）和 FaaS（Function as a Service）。  

BaaS（Backend as a Service）后端即服务，一般是一个个的 API 调用后端或别人已经实现好的程序逻辑，比如身份验证服务 Auth0，这些 BaaS 通常会用来管理数据，还有很多公有云上提供的我们常用的开源软件的商用服务，比如亚马逊的 RDS 可以替代我们自己部署的 MySQL，还有各种其它数据库和存储服务。  

FaaS（Functions as a Service）函数即服务，FaaS 是无服务器计算的一种形式，当前使用最广泛的是 AWS 的 Lambada。FaaS 本质上是一种事件驱动的由消息触发的服务，FaaS 供应商一般会集成各种同步和异步的事件源，通过订阅这些事件源，可以突发或者定期的触发函数运行。  
两者都为我们的计算资源提供了弹性的保障，BaaS 其实依然是服务外包，而 FaaS 使我们更加关注应用程序的逻辑，两者使我们不需要关注应用程序所在的服务器，但实际上服务器依然是客观存在的。  

**Serverless 的使用场景**  
虽然 Serverless 的应用很广泛，但是其也有局限性，Serverless 比较适合以下场景。  

- 异步的并发，组件可独立部署和扩展
- 应对突发或服务使用量不可预测（主要是为了节约成本，因为 Serverless 应用在不运行时不收费）
- 短暂、无状态的应用，对冷启动时间不敏感
- 需要快速开发迭代的业务（因为无需提前申请资源，因此可以加快业务上线速度）  

使用场景示例如机器学习及 AI 模型处理、图片处理、流处理、聊天机器人、发送通知、WebHook、轻量级 API、物联网等。  

**Serverless 的优缺点**  
Serverless 的优点主要有如下这些。  

- 降低运营成本
- 降低开发成本
- 扩展能力强，增加缩放的灵活性
- 更简单的管理
- “绿色” 的计算
- 降低风险
- 减少资源开销
- 缩短创新周期

Serverless 的缺点主要有如下这些。 

- 状态管理，对于无状态的处理，有状态的 Serverless 服务丧失了灵活性
- 延迟高
- 本地测试很棘手，缺乏调试和开发工具
- 不适合长时间运行应用
- 完全依赖于第三方服务

### Hello，Serverless
先按官网的 demo，进行实验。  
开始之前，除了拥有一台电脑，你还需要有一个 AWS 账号。AWS 提供一年的免费试用，你所需要做的就是办一张支持 visa 的信用卡。  
```bash
npm install -g serverless # yarn global add serverless
```

设置 AWS 凭证。
> 登录 AWS 账号，然后点击进入 IAM (即，Identity & Access Management)。  
> 点击用户，然后添加用户，如 serveless-admin，并在『选择 AWS 访问类型』里，勾上编程访问。  
> 点击下一步权限，选择『直接附加现有策略』，输入 AdministratorAccess，然后创建用户。  
> 创建用户。随后，会生成访问密钥 ID 和 私有访问密钥，请妥善保存好。  
> 然后导出证书，并使用 serverless depoy 保存到本地。  

```bash
# 创建 hello-world 服务，生成两个文件：handler.js 和 serverless.yml
serverless create --template aws-nodejs --path hello-world

# 部署及测试
serverless deploy -v
# 触发函数
serverless invoke -f hello -l
# 获取相应的日志
serverless logs -f hello -t
```
