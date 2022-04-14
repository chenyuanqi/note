
### Go 环境配置
Go 团队已经将版本发布节奏稳定在每年发布两次大版本上，一般是在二月份和八月份发布。Go 团队承诺对最新的两个 Go 稳定大版本提供支持，比如目前最新的大版本是 Go 1.17，那么 Go 团队就会为 Go 1.17 和 Go 1.16 版本提供支持。  

**Go 安装**  
```bash
# linux
wget -c https://golang.google.cn/dl/go1.16.5.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.16.5.linux-amd64.tar.gz
# vim $HOME/.profile 
export PATH=$PATH:/usr/local/go/bin
source ~/.profile
go version

# mac
# wget -c https://golang.google.cn/dl/go1.16.5.darwin-amd64.pkg
# 可视化安装
# vim $HOME/.profile 
# export PATH=$PATH:/usr/local/go/bin
# source ~/.profile

# mac 多版本
brew install go
brew install go@1.15
# 切换版本
brew unlink go
brew link go@1.15
```

**Go 升级**
```bash
# 设置代理镜像
go env -w GOPROXY=https://goproxy.cn,direct
# Mac
export GOPROXY="https://goproxy.cn,direct"

# Linux
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf /home/nikhita/Downloads/go1.8.1.linux-amd64.tar.gz
echo $PATH | grep "/usr/local/go/bin"

# Mac
brew update
brew upgrade
brew upgrade go
```

### Go 环境变量
环境变量（environment variables）一般是指在操作系统中用来指定操作系统运行环境的一些参数，如：临时文件夹位置和系统文件夹位置等。环境变量是在操作系统中一个具有特定名字的对象，它包含了一个或者多个应用程序所将使用到的信息。  
```bash
sudo vi ~/.zshrc
# GOPATH用于指定我们的开发工作区，是存放源代码、测试文件、库静态文件、可执行文件的目录
export GOPATH=$HOME/golang
# GOROOT表示 Go 语言的安装目录，当系统中存在多个版本的Go SDK时，通过设置这个环境变量，可方便我们在不同的Go SDK版本之间切换
export GOROOT=$HOME/go1_17
# 追加 path
export PATH=$PATH:$GOROOT/bin
# GOBIN 表示程序编译后二进制命令的安装目录，一般设置为 GOPATH/bin
export GOBIN=$GOPATH/bin
source  ~/.zshrc
# 查看环境变量
go env
# 改写环境变量 go env -w <NAME>=<VALUE>
# 设置代理镜像
go env -w GOPROXY=https://goproxy.cn,direct
```

