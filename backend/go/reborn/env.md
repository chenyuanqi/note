
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
wget -c https://golang.google.cn/dl/go1.16.5.darwin-amd64.pkg
# 可视化安装
# vim $HOME/.profile 
export PATH=$PATH:/usr/local/go/bin
source ~/.profile
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
```

