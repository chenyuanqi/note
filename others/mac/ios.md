
### IOS 实用技巧
```
# 去除烦人的系统更新提示  
1. 设置 > 通用 > 存储空间 > 找到更新包，删除  
2. 设置 > iTunes&App Stores > 找到更新，关闭  
3. Safari 打开 https://oldcat.me/web/NOOTA11.mobileconfig - 安装该 provision 文件 - 重启  

# mac、ipad/iphone 文件传输
使用 AirDrop (隔空传送) 功能，ipad 该功能位于设置 > 通用 > 隔空传送；  
mac 和 ipad 都设置为允许所有人接收，mac 右键选择文件或文件夹 > 共享 > AirDrop > 点击要传送的机器名开始传送；  
ipad 接受后会弹出要存储的位置，选择即可。  

# 扫描二维码
使用相机扫描即可

```

### IOS 推荐应用
```
Thor: 抓包神器  
AdBye Pro: 广告杀手 (Notice: 某些应用会被误伤)    
Shadowrocket: 翻墙的小火箭  

VVebo: 微博客户端  
讯飞输入法：有态度的语音输入  
Picsew: 网页截图+截图合并（收费 1 软妹币）   
to:day: 计划+提醒    

```

### IOS 命令行
```
# 朗读指定内容
say anything

# 查看当前电源的使用方案
pmset -g
# 设置不休眠
pmset noidle
# 设置电池供电时，显示器5分钟内进入睡眠
sudo pmset -b displaysleep 5

# 

```

### IOS 解压缩
```
.tar 
解包：tar xvf FileName.tar
打包：tar cvf FileName.tar DirName
（注：tar是打包，不是压缩！）
———————————————
.gz
解压1：gunzip FileName.gz
解压2：gzip -d FileName.gz
压缩：gzip FileName

.tar.gz 和 .tgz
解压：tar zxvf FileName.tar.gz
压缩：tar zcvf FileName.tar.gz DirName
———————————————
.bz2
解压1：bzip2 -d FileName.bz2
解压2：bunzip2 FileName.bz2
压缩： bzip2 -z FileName

.tar.bz2
解压：tar jxvf FileName.tar.bz2
压缩：tar jcvf FileName.tar.bz2 DirName
———————————————
.bz
解压1：bzip2 -d FileName.bz
解压2：bunzip2 FileName.bz
压缩：未知

.tar.bz
解压：tar jxvf FileName.tar.bz
压缩：未知
———————————————
.Z
解压：uncompress FileName.Z
压缩：compress FileName
.tar.Z

解压：tar Zxvf FileName.tar.Z
压缩：tar Zcvf FileName.tar.Z DirName
———————————————
.zip
解压：unzip FileName.zip
压缩：zip FileName.zip DirName
———————————————
.rar
解压：rar x FileName.rar
压缩：rar a FileName.rar DirName
———————————————
.lha
解压：lha -e FileName.lha
压缩：lha -a FileName.lha FileName
———————————————
.rpm
解包：rpm2cpio FileName.rpm | cpio -div
———————————————
.deb
解包：ar p FileName.deb data.tar.gz | tar zxf -
———————————————
.tar .tgz .tar.gz .tar.Z .tar.bz .tar.bz2 .zip .cpio .rpm .deb .slp .arj .rar .ace .lha .lzh .lzx .lzs .arc .sda .sfx .lnx .zoo .cab .kar .cpt .pit .sit .sea
解压：sEx x FileName.*
压缩：sEx a FileName.* FileName
```
