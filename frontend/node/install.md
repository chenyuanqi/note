
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