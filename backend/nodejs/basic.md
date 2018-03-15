
### Nodejs 安装及更新
```bash
# Ubuntu
# Install nodejs
curl -sL https://deb.nodesource.com/setup_8.x| sudo -E bash -
sudo apt-get install -y nodejs
# Upgrade nodejs
npm install -g n
# 更新到文档版本
n stable 
# 或者更新到最新版本
n latest
```

### Nodejs 执行流程及小试牛刀
JavaScript code -> c++ v8 引擎 -> 机器码
```javascript
// 显示当前所在目录绝对路径
console.log(__dirname);
// 显示当前文件绝对路径
console.log(__filename);

// 定时执行一段代码
time = 0
var timer = setInterval(time => {
    time += 2;
    console.log(time + " Seconds Passed.");
    if (time > 10) {
        clearInterval(timer);
    }
}, 2000)
```

### Nodejs 函数
```javascript
// 基本函数
function hello() {
    console.log("hello nodejs");
}

// 匿名函数
function hello = function() {
    console.log("hello, nodejs");
}

// 回调函数
function callFunc(fun) {
     fun();
}
callFunc(hello);

// 常见回调写法
sayHello = callFunc(function(name) {
    console.log("hello " + name)
});
sayHello("nodejs");
```

### Nodejs 模块
```javascript
callModule = function() {
    console.log("Hi, node module");
}

var pi = 3.14;

function plus(a, b) {
     return a + b;
}

module.exports = {
    callModule: callModule, 
    plus: plus,
    pi: pi
}
// 拆开写也是可以的
// 比如 module.exports.callModule = callModule;

// 模块的使用示例
require "./demo"
// require 返回 module 的 exports 对象，所以可以赋值；模块的引入不需要加后缀 .js
```

exports 与 module.exports 的区别?  
[参考链接](https://cnodejs.org/topic/52308842101e574521c16e06)  

### Nodejs 事件
```javascript
var events = require("events") ;
// 创建事件监听
var myEmitter = new events.EventEmitter();
// 绑定监听事件
myEmitter.on("print", function(message) {
    console.log(message);
});
// 触发事件
myEmitter.emit("print", "hello, node events");

// 使用事件
var events = require("events");
// 工具库 util
var util = require("util");

var Person = function(name) {
    this.name = name;
}

util.inherits(Person, events.EventEmitter);

var lili = new Person("lili");
var lucy = new Person("lucy");

var persons = [lili, lucy];

persons.forEach(function(person) {
    person.on("say", function(message) {
        console.log(person.name + " said: " + message);
    });
})

lili.emit("say", "I want an apple");
lucy.emit("say", "sorry, I dont");
```

### Nodejs 读写文件
```javascript
// 使用文件库 fs
var fs = require("fs");

// 同步读取
fs.readFileSync("file_path", "utf8");
// 异步读取
fs.readFile("file_path", "utf8", function(err, data) {
    console.log(data, err);
});
// 异步写入
fs.writeFile("file_path", data, function() {
    console.log("write it");
});

// 创建和删除目录
fs.unlink("file_path", function() {
    console.log("delete file.");
});
fs.mkdirSync("dir_path", function() {
    fs.readFile("file_path", "utf8", function(err, data) {
        fs.writeFile("dir_path/file_path", data, function() {
            console.log("copy success");
        });
    });
});
fs.rmdirSync("dir_path");
// 更多方法，参考官方文档 api: https://nodejs.org/dist/latest-v8.x/docs/api/
```

### Nodejs 流和管道
在 Linux 中，命令 ls | grep app 是以 ls 的输出流向 grep 处理的输入, | 为管道；  
在 http 中，请求就像是输入流，响应则是输出流。  

那么，流的作用是什么呢？  
流能处理数据，能提高性能。  

```javascript
var fs = require("fs");

var myReadStream = fs.createReadStream(__dirname + "file_path", "utf8");
// 以 buffer 分段处理文件
// 设置编码还可以这样， myReadStream.setEncoding("utf8");
var myWriteStream = fs.createWriteStream(__dirname + "file_path", "utf8");

var data = "";
myReadStream.on("data", function(chunk) {
    data += chunk;
    myWriteStream.write(chunk, "utf8");
    // myWriteStream.end();
    // myWriteStream.on("finish", function() {});
});

myReadStream.on("end", function() {
    console.log(data);
})

// 使用管道读写文件
var myReadStream = fs.createReadStream(__dirname + "file_path", "utf8");
var myWriteStream = fs.createWriteStream(__dirname + "file_path", "utf8");

myReadSteam.pipe(myWriteStream);
```

### Web 服务器
```javascript
var http = require("http")

var onRequest = function(request, response) {
    console.log("Request received");
    response.writeHead(200, { "Content-Type": "text/plain" });
    // 如果是传送 json, Content-Type 为 application/json
    // 如果是 HTML 页面，则 Content-Type 为 text/html，此时，可以使用 fs 模块获取相应文件的内容，并通过管道输出；但是，管道输出时，不需要 reponse.end(); 声明

    // response.write("Hi");
    response.end("Hi");
    // myObj = { name: "vikey"}; response.end(JSON.stringify(myObj));
    // JSON.stringify(obj) 和 JSON.parse(json_str) 是序列化与反序列化的过程
}

var server = http.createServer(onRequest);

server.listen(8080, "127.0.0.1");
```

### 包管理器 Npm
npm 的更新  
```bash
npm install npm@latest -g
```

npm 常用命令
```bash
# 初始化项目
npm init 
# 安装项目依赖
npm install --save package_name 
# 安装项目的开发依赖
npm install --save-dev package_name 

# 搜索指定包
npmsearch package_name
# 安装指定包
npm install (-g) package_name
# 更新指定包
npm update (-g) package_name
# 卸载指定包
npm uninstall (-g) package_name

# 使用 package.json 中的 script 命令
npm run "command_name"
```

nodemon，一个无刷新监控文件更改工具  
```bash
# 全局安装
npm install -g nodemon
# 简单使用
nodemon js_file
```
