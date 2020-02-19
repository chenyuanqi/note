
### 安装 Node.js 的最佳方案 —— NVM
在我们的日常开发中经常会遇到这种情况：手上有好几个项目，每个项目的需求不同，进而不同项目必须依赖不同版的 NodeJS 运行环境。如果没有一个合适的工具，这个问题将非常棘手。  

NVM 应运而生，NVM 是 Nodejs 版本管理工具，有点类似于 Python 的 virtualenv 或者 Ruby 的 rvm，每个 Node 版本的模块都会被安装在各自版本的沙箱里面（因此切换版本后模块需重新安装）。开发环境中，我们需要经常对 Node 版本进行切换测试兼容性和一些模块对 Node 版本的限制。NVM 也很好了解决了同台机器上多个项目 Node 版本不兼容的窘境。  

### NVM 安装
Mac 下，
```bash
# 安装命令行工具
xcode-select --install
# 安装 nvm
curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.33.11/install.sh | bash
```

Linux 下，
```bash
# 尽量避免在 Ubuntu 上使用  apt-get 来安装 Node.js，如果你已经这么做了，请手动移除
sudo apt-get purge nodejs && sudo apt-get autoremove && sudo apt-get autoclean

# 安装 nvm
curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.33.11/install.sh | bash

# 对 inotify 做以下配置
echo fs.inotify.max_user_watches=524288 | sudo tee -a /etc/sysctl.conf && sudo sysctl -p
```

Windows 下，  
NVM 官方版本并不支持 Windows，目前来看，使用 [Node.js 官方安装包](https://nodejs.org/en/download/ )来安装是最稳妥的方式。


### NVM 基本使用
```bash
# 查看本地所有可以用的 Node.js 版本
nvm list

# 查看服务器端可用的 Node.js 版本：
nvm ls-remote

# 推荐使用 8.* LTS 版本 (长久维护版本) ，使用以下命令安装
nvm install 8.11.2

# 设置默认版本
nvm use 8.11.2
nvm alias default 8.11.2

# 使用淘宝进行加速 NPM
npm config set registry=https://registry.npm.taobao.org

# 将 NPM 更新到最新
npm install -g npm
```

### npm —— Node 生态利器
Node 世界里，一切皆模块，而安装模块，皆是 npm i （也就是 npm install 的缩写）。  
如果说 Node 是打开了前后端同一种语言的能力大门，那么 npm 就是让千军万马通过大门的高速公路，完全赋能了 Node，让它的背后长出了一个无比繁荣的军工城市群，成千上万的手艺人在那里无时无刻的制造趁手的工具，你能想到的几乎所有功能，只要想偷懒，就可以前往免费取回来。  
大家可以输入 `npm -h npm -l` 来查看 npm 提供的命令集。  

**包（模块）服务 - npm registry**  
npm 全称 Node Package Manager，只是它早就不是只为 Node 服务的包（模块）管理工具，海洋般的前端模块也一并被它纳入怀中。  
有了 npm，让包的下载变成一件特别 easy 的事情，比如命令行里丢进去 `npm i lodash@4.17.11` 运行后， 就安装好了版本是 4.17.11 的 lodash 包。  

模块汇聚成了包，包是从哪里下载的呢？registry。  
registry 直译就是注册表，再直白一点，就是所有可被下载的模块，需要有一个地方存储和记录他们，并且对外提供一个可查和下载的服务，对于 Node 模块来说，npm registry 就是这个服务，那么整个 npm 包括三个部分：  
> npm 官网，网址 www.npmjs.com，可以通过网站直接查询一个模块的相关信息  
> npm registry，https://registry.npmjs.org/ 则是模块查询下载的服务 API  
> npm cli(https://github.com/npm/cli)，则是一个命令行下载工具，通过它来从 registry 下载模块

需要我们注意的是，这个 npm registry 不过是 Isaaz 公司提供给 Node 社区的服务，也是 npm 默认使用的服务，但它并不是唯一的服务，只要你有意愿，你都可以搭建自己的 registry，甚至是使用第三方的 registry，比如 [淘宝 NPM 镜像](https://npm.taobao.org/)，它也有自己的 cli - [cnpm cli](https://github.com/cnpm/cnpm)，每三十分钟同步一次，大家在国内如果网络不通畅，建议使用淘宝镜像代替 npm 来作为模块下载的服务，使用的办法非常简单：  
```bash
npm i lodash@4.17.11 --registry=https://registry.npm.taobao.org
# 或者通过 cnpm 来代替
npm install -g cnpm --registry=https://registry.npm.taobao.org
cnpm install lodash@4.17.11
```
如果想给自己团队搭建私有 registry，可以参考 [cnpmjs.org](https://github.com/cnpm/cnpmjs.org)。除了 npm 官方的 registry 付费私有服务，还可以先注册一个阿里云账号，然后打开 [Aliyun Registry](https://node.console.aliyun.com/registry/home)。  

**包的地图 - npm init**  
npm init 可以直接创建一个 package.json，在它里面主要记录了：  

- 当前模块/项目名称、版本、描述、作者  
- 配置了当前项目依赖的模块及版本  
- 配置了当前项目是在哪个 git repo 下  
- 当前模块的主入口文件 main  
- ...  

我们也可以 `npm init --yes` 可以直接创建默认值的 package.json。  

**包的安装 - npm install**  
我们最常用的 npm install 通常会搭配 --save 和 --save-dev 来使用，分别把模块安装到 dependencies 和 devDependencies，一个是运行时依赖的模块，一个是本地开发时依赖的模块，当然也可以简写，比如 `npm i xx -S` 和 `npm i xx -D`，我个人会建议大家在本地安装模块，一定要指定 --save 或者 -S，来确保本地模块正确的添加到 package.json，那么具体 install 都支持哪些类型的模块呢，其实我们可以通过 `npm help install` 找到答案。  
```bash
# 项目中已有 package.json，可以直接 npm intall 安装所有依赖项
npm install (with no args, in package dir)
# scope 通常用于管理私有模块，以 @ 开头，没有 @ 则反之
# npm install some-pkg
# npm install @scott/some-pkg
npm install [<@scope>/]<name>
# 可以安装特定 tag 的模块，默认是 latest，如：
# npm install lodash@latest
npm install [<@scope>/]<name>@<tag>
# npm install lodash@4.17.11
# npm install @scott/some-pkg@1.0.0
npm install [<@scope>/]<name>@<version>
# 安装一个某个范围内的版本
# npm install lodash@">=2.0.0 <3.0.0"
# npm install @scott/som-pkg@">=2.0.0 <3.0.0"
npm install [<@scope>/]<name>@<version range>
# npm install git+ssh://git@github.com:tj/commander.js.git
npm install <git-host>:<git-user>/<repo-name>
# 以 git 仓库地址来安装
# npm install https://github.com/petkaantonov/bluebird.git
npm install <git repo url>
# 安装本地的 tar 包
# npm install /Users/black/Downloads/request-2.88.1.tar.gz
npm install <tarball file>
# 以 tar 包地址来安装
# npm install https://github.com/caolan/async/tarball/v2.3.0
# npm install https://github.com/koajs/koa/archive/2.5.3.tar.gz
npm install <tarball url>
# 从本地文件夹安装
# npm install ../scott/some-module
npm install <folder>
# 卸载也很简单
npm uninstall some-pkg -S
# 或者简写，加上 -S 是把卸载也同步到 package.json 中
npm un some-pkg -S
```
npm 有如此多样的安装方式，给了我们很多想象力，比如在强运维大量服务器存储保障的前提下，可以把所有的模块全部 tarball 形式从本地上传到服务器，可以保证所有模块代码的绝对一致性，甚至 npm registry 不稳定的时候也不影响，更不用提私有模块 scope 的组合使用。  
抛开 npm 带来的便捷，反过来的一个问题就是模块的版本管理，怎样保证我本地安装的版本，跟服务器上运行的版本，两份代码是一模一样的，关于这一点我们先往下看，了解 `npm versions` 后再来讨论版本的管理问题。  

**包版本 - npm semver version**  
只要一个文件夹里面有 package.json，里面的基本信息完备，无论里面有多少个 js 模块，整个文件夹便可以看做是包，我们所谓的 npm - 包管理工具，其实本质上就是管理这个文件夹的版本，将几十上百个甚至上千个项目所依赖的包，全部下载到本地，全部整理到 node_modules 下面管理，每个包都是一个独立的文件夹，比如安装了 bluebird 和 lodash 的项目，package.json 的依赖是这样的：  
```
"dependencies": {
  "bluebird": "^3.5.2",
  "lodash": "^4.17.11"
}

# 在 node_modules 里面是这样的
.
├── node_modules
│   ├── bluebird
│   └── lodash
├── package-lock.json
└── package.json
```

每个文件夹都是一个 package（包），每个包都有自己依赖的其他包，每个包也都有自己的名称和版本，每个包作者对于版本的管理都不尽相同，有依赖大版本的，有依赖小版本的，有奇数偶数策略的等等，我们说下业界最常见的一种版本管理方式，就是 Semantic Versioning 2.0.0，大家可以前往 [semver.org](https://semver.org/) 查看详情，这里简单解释下，版本号比如 v4.5.1，v 是 version 的缩写，4.5.1 被 . 分开成三端，这三端分别是：major minor patch，也就是 主版本号.次版本号.修订号：  

- major: breaking changes (做了不兼容的 API 修改)  
- minor: feature add（向下兼容的功能性新增）  
- patch: bug fix, docs（向下兼容的问题修正）  

拿 lodash@4.17.11 举例（不一定准确，仅示例），4 代表主版本，17 就是次版本，11 就是修订号，如果每一个变动都严格遵守，且每次都是 +1 的话，可以这样理解：lodash 经历了 3 次大的断代更新，也即从 1 到 4，同时在 4 的大版本上，经历了 17 次的功能更新，并且向下兼容，至于 bug 修复之类也有 11 次，每个包实际执行并不一定严格遵守这种语义化版本规范，所以也会带来一些管理困扰，但真正的困扰我们的反而不是版本号本身，而是包与包之间的依赖关系，以及包自身的版本稳定性（背后的代码稳定性）。  

**包目录层级 - npm node_modules**  
在 npm 的升级历史中，有这样的一个重大的变化，那就是 node_module 是包依赖安装层级，也就是 npm2 时代和 npm3+ 时代，在 npm2 时代，一个项目的 node_modules 目录是递归安装的，它是按照依赖关系进行文件夹的嵌套，比如：  
```
.
├── connect-mongo
│   ├── node_modules
│   │   └── mongodb
│   │       ├── node_modules
├── mongoose
│   ├── node_modules
│   │   ├── mongodb
│   │   │   └── node_modules
│   │   └── sliced
├── async
├── grunt
│   ├── node_modules
│   │   ├── async
│   │   └── which
└── underscore
```
通常一个项目依赖三四十个包就算比较多了，在 node_modules 里面，也就三四十个目录，进去找一个包的源代码，或者去它的 node_modules 里继续向下找，会非常省事，尤其是当我去 review 源码去查找关键字的时候，但它的缺点也有很多，比如嵌套可能会出现很深的情况，会遇到 windows 的文件路径长度限制，当然最敏感的是，会导致大量的代码冗余，比如我们上面，connect-mongo 和 mongoose 里面都用到 mongodb，grunt 里面也用到了 async 等等，这会导致整个项目体积特别的臃肿。  
所幸是 npm3 时代里面策略改成了平铺结构，全部一股脑平铺到 node_modules 下面。  
但是要注意，新的 npm 并不会无脑的平铺，而是会有一套算法来做同名且同版本的包去重，合理规划目录的嵌套层级，这样可以保证即便是有同名但是版本不同的模块，不会在 node_modules 里面冲突，同时只要不在同级冲突，npm 会尽可能把能复用的模块往高层级安装，这样可以达到最大程度的模块重用，代码冗余就大幅降低。  

**锁包 - npm shrinkwrap**  
除了安装策略外，npm 另外一个重大的升级，就是我们熟悉的 package-lock 文件，这是 npm5 以后带来的新特性，package-lock，顾名思义，就是把包的版本锁住，保证它的代码每一行每一个字节都恒久不变，为什么需要这样一种看上去奇葩的策略，我们还得结合上面的 Semantic Versioning 也就是包的语义化版本来说事。  
在一个 package.json 里的 dependencies 里面，包的依赖版本可以这样写：  
```
"lodash": "~3.9.0",
"lodash": "^3.9.0",
"lodash": ">3.9.0",
"lodash": ">=1.0.0-rc.2",
"lodash": "*"
...
```
最常见的就是 ~ 和 ^ 这两种写法，它俩有什么区别呢，~ 意思是，选择一个最近的小版本依赖包，比如 ~3.9.0 可以匹配到所有的 3.9.x 版本，但是不会匹配到 3.10.0，而 ^ 则是匹配最新的大版本，比如 ^3.9.0 可以匹配到所有的 3.x.x，但是不会匹配到 4.0.0。他们的好处很明显，就是当一个包有一些 bug， 作者修复之后，不需要我们开发者主动到 package.json 里，一个个的修改过去，事实上我们开发者也无从知晓作者什么时候升级了包，甚至我们都不知道里面有没有 bug，所以依靠 ~ 和 ^，它就能自动晋升到较新版本的包，里面包含了最新的代码，只不过 ^ 比 ~ 更加激进，可能会导致新包与项目的不兼容，而 ~ 会友好很多，但也不能保证 100% 的兼容，因为所有的包版本都是包作者自行管理的，作者的技术实力和版本意识也是有限的，它这次升级会不会导致你的项目出现问题，我们心里是没底的。  
于是千古难题出现了，我们既想享受静默升级的好处，又要避免静默升级背后包代码的不兼容性，这两个实际上是冲突的，静默升级一定会带来代码变动，代码变动一定会带来兼容风险，而且，就算是我们把版本写死为 3.9.0，也是不能保证 3.9.0 的这个包，它自身又向下依赖的很多别的包，这些别的包又依赖了别的包，他们的包策略如果是语义化的，照样会带来包依赖树的不稳定（任何一个底层包代码有语义化升级）。  
所以，路被堵死了，意味着除非我们把整个 node_modules 保存到本地，上传到 git 仓库，全量上传到服务器，我们根本无法保证代码的不变性，据淘宝的工程师讲，他们某段时间也确实是这么干的，全包上传，全包回滚，粗暴但实用。  
那么到底应该怎么办呢，大家可能猜到了，答案就是 package-lock.json，也就是 npm 的锁包。  

大家可以在本地的一个空目录下，执行 `npm init --yes && npm i lodash async -S`，然后我们来看下 package-lock.json 里面的内容：
```
{
  "name": "npm",
  "version": "1.0.0",
  "lockfileVersion": 1,
  "requires": true,
  "dependencies": {
    "async": {
      "version": "2.6.1",
      "resolved": "http://registry.npm.taobao.org/async/download/async-2.6.1.tgz",
      "integrity": "sha1-skWiPKcZMAROxT+kaqAKPofGphA=",
      "requires": {
        "lodash": "^4.17.10"
      }
    },
    "lodash": {
      "version": "4.17.11",
      "resolved": "http://registry.npm.taobao.org/lodash/download/lodash-4.17.11.tgz",
      "integrity": "sha1-s56mIp72B+zYniyN8SU2iRysm40="
    }
  }
}
```
version 就是包的准确版本号（无语义化的跃迁）， resolved 则是一个明确 tar 包地址，它是唯一不变的，并且还有 integrity 这个内容 hash 的值，他们三个就决定了这个包准确身份信息，这样第一个问题就解决了，那就是特定版本的包代码不变性，然后第二个问题，这些包向下依赖的包如何不变？  
这个是通过每个包的 requires 字段实现，它实际上跟每个包的内部 package.json 的 dependencies 里的包是一一对应的，所以包的依赖关系也有了，无论嵌套多少层级，在 lock 文件里面，它都有 version、resolved、integrity 来保证单包不变性，那么整包就保证了代码不变，可以把 package-lock.json 理解为一个详细描述代码版本的快照文件，它储存了 node_modules 当前的包代码状态，无论被哪个团队成员拿走项目，无论是本地还是服务器上 npm install，都能依据 package-lock.json 里面的包状态，原封不动的复原 node_modules 里面的代码版本。  

这个就是锁包功能，其实在 npm5 之前就提供了，也就是 `npm shrinkwrap`，它需要手动执行，而现在则是自动生成。如果你完全不依赖锁包功能，则可以将它关闭： `npm config set package-lock false`。  

**包脚本 - npm scripts**  
npm 最强大的能力，除了 install 安装能力，就是脚本能力，在 package.json 里的 scripts 里配置的各种任务，都可以这样直接调用：  
```bash
npm start
npm run dev
```
结合 npm 社区海量的包资源，跨平台执行也完全没有问题，比如 rm -rf 在 windows 下不支持，或者考虑支持 windows/linux 都可以设置环境变量，都可以换一个模块来执行，比如：  
```
"scripts": {
  "build": "npm run build:prod",
  "clean:dist": "rimraf ./dist",
  "build:prod": "cross-env NODE_ENV=production webpack"
}

npm run clean:dist
npm run build:prod
npm run build
```
npm scripts 如此之强大，甚至直接替换历史产物 grunt/gulp，尤其是处理一些构建预准备工作或构建后任务，比如先检查代码规范，再跑单元测试，最后跑构建，构建成功了就发一个钉钉通知到团队等等，这些任务可能是级联关系也可能是并行关系，在 npm scripts 里面也轻松搞定，比如：  
```
"scripts": {
  // 通过 && 分隔，如果 clean:dist 任务失败，则不会执行后面的构建任务
  "build:task1": "npm run clean:dist && npm run build:prod"
  // 通过 ; 分隔，无论 clean:dist 是否成功，运行后都继续执行后面的构建任务
  "build:task2": "npm run clean:dist;npm run build:prod"
  // 通过 || 分隔，只有当 clean:dist 失败，才会继续执行后面的构建任务
  "build:task3": "npm run clean:dist;npm run build:prod"
  "clean:dist": "rimraf ./dist",
  "build:prod": "cross-env NODE_ENV=production webpack",
  // 对一个命令传配置参数，可以通过 -- --prod
  // 比如 npm run compile:prod 相当于执行 node ./r.js --prod
  "compile:prod": "npm run compile -- --prod",
  "compile": "node ./r.js",
}
```
npm scripts 可以构建非常复杂的任务，但是 npm scripts 也会带来一些问题，比如非常复杂的 scripts 会带来非常复杂的依赖队列，不好维护，针对这一点，建议把每个独立的任务都分拆开进行组合，可以把复杂的任务独立写入到一个本地的脚本中，比如 task.js。  

如果需要底层系统命令支撑，又实在找不到跨平台的包，也可以在它里面，使用 shelljs 来调用系统命令，甚至不仅仅局限于 Node 的包，在 script 里面调用 python 脚本和 bash 脚本也一样溜，相信我，npm scripts 会给你打开一片新天地，大家有时间也可以研究下 [npmasbuildtool](https://github.com/marcusoftnet/npmasbuildtool/blob/master/package.json) 的 scripts 清单

**包执行工具 - npx**  
npx 是 npm 自带的非常酷炫的功能，直接执行依赖包里的二进制文件，比如：  
```
# 先安装一个 cowsay
npm install cowsay -D
# 包里的二进制文件会被放到 node_modules/.bin 目录下
ll node_modules/.bin/
# total 0
# lrwxr-xr-x 1 16:34 cowsay -> ../cowsay/cli.js
# lrwxr-xr-x 1 16:34 cowthink -> ../cowsay/cli.js

# 直接通过 npx 来调用 cowsay 里的二进制文件
npx cowthink Node 好玩么
 _____________
( Node 好玩么 )
 -------------
        o   ^__^
         o  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```
甚至我们 `npm i webpack -D` 以后，可以直接 npx http-server 把静态服务开起来。  

**包发布 - npm publish**  
如何发布一个包呢？  
首先你要有一个 npm 的账号和 Github 账号，可以分别到 npmjs.com 和 github.com 注册，各自都注册且验证邮箱后（Github 还需要配置 ssh key），这些都搞定后，就可以准备开发和发布 NPM 包了，整个流程很简单，总共都不超过 10 步：  
1、本地（或者从 Github 上）创建创建一个空项目，拉到本地  
2、增加 .gitignore 忽略文件和 README  
3、npm init 生成 package.json  
4、编写功能代码，增加到目录 /lib（也可以增加一个 bin 文件夹存放脚本）  
5、npm install 本地包进行测试  
6、npm publish 发布包  
7、npm install 线上包进行验证  
8、修改代码发布一个新版本  

