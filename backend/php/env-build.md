
### 开发环境部署 Link
[Laravel 开发环境部署](https://learnku.com/docs/laravel-development-environment/5.5)  

### 开发统一环境 Vagrant
```bash
# 安装
brew cask install virtualbox
brew cask install vagrant

# 下载镜像
vagrant box add --provider virtualbox centos/7

# 初始化
vagrant init centos/7 

# 挂载文件目录配置
# Vagrant.configure("2") do |config|
#  config.vm.box = "centos/7"
#  config.vm.synced_folder "/c/code", "/data/code"
#end
# 安装插件
vagrant plugin install vagrant-vbguest

# 查看 box 列表
vagrant box list
# 打包环境
vagrant package --base 导出的 box 的名称 --output 导出后的 box 名称

```
