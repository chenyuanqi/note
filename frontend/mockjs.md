
### mock 概览
如果将 mock 单独翻译过来，其意义为 “虚假、虚设”，因此在软件开发领域，我们也可以将其理解成 “虚假数据”，或者 “真实数据的替身”。  

[easy mock](https://easy-mock.com/docs)  


Mock 的好处：  

- 团队可以更好地并行工作  
当使用 mock 之后，各团队之间可以不需要再互相等待对方的进度，只需要约定好相互之间的数据规范（文档），即可使用 mock 构建一个可用的接口，然后尽快的进行开发和调试以及自测，提升开发进度的的同时，也将发现缺陷的时间点大大提前。  
- 开启 TDD（Test-Driven Development）模式，即测试驱动开发  
单元测试是 TDD 实现的基石，而 TDD 经常会碰到协同模块尚未开发完成的情况，但是有了 mock，这些一切都不是问题。当接口定义好后，测试人员就可以创建一个 Mock，把接口添加到自动化测试环境，提前创建测试。  
- 测试覆盖率  
比如一个接口在各种不同的状态下要返回不同的值，之前我们的做法是复现这种状态然后再去请求接口，这是非常不科学的做法，而且这种复现方法很大可能性因为操作的时机或者操作方式不当导致失败，甚至污染之前数据库中的数据。如果我们使用 mock，就完全不用担心这些问题。  
- 方便演示  
通过使用 Mock 模拟数据接口，我们即可在只开发了 UI 的情况下，无须服务端的开发就可以进行产品的演示。  
- 隔离系统  
在使用某些接口的时候，为了避免系统中数据库被污染，我们可以将这些接口调整为 Mock 的模式，以此保证数据库的干净。  

### mock 简单使用
先安装如下依赖
```bash
npm install --save-dev json-server
npm install --save-dev nodemon
```

路由 mock/routes.js 的编写
```javascript
module.exports = {
  "/comment/get": "/getComment"
}
```

数据库 mock/db.js 的编写  
```javascript
var Mock = require('mockjs');

// 通过使用mock.js，来避免手写数据
module.exports = {
  getComment: Mock.mock({
    "error": 0,
    "message": "success",
    "result|10": [{
      "author": "@name",
      "comment": "@cparagraph",
      "date": "@datetime"
    }]
  })
};
```

服务器 mock/server.js 的编写
```javascript
const jsonServer = require('json-server')
const db = require('./db.js')
const routes = require('./routes.js')
const port = 3000;

const server = jsonServer.create()
// 使用 mock 的数据生成对应的路由
const router = jsonServer.router(db)
const middlewares = jsonServer.defaults()
// 根据路由列表重写路由
const rewriter = jsonServer.rewriter(routes)

server.use(middlewares)
// 将 POST 请求转为 GET，满足可以接受 POST 和 GET 请求的需求
server.use((request, res, next) => {
  request.method = 'GET';
  next();
})

server.use(rewriter) // 注意：rewriter 的设置一定要在 router 设置之前
server.use(router)

server.listen(port, () => {
  console.log('open mock server at localhost:' + port)
})
```

接下来，运行如下命令即可
```bash
# 运行之后，浏览器访问 http://localhost:3000/comment/get
node_modules/.bin/nodemon --watch mock mock/server.js
```
