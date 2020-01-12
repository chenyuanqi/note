
```bash
# 全局配置保存在：~/.gitconfig，本地仓库配置保存在：.git/config
# 设置用户信息（多个不同源的仓库，不建议使用 --global）
git config user.name "xxx"
git config user.email xxx@xxx.com
# 设置 UI 及大小写敏感
git config color.ui auto
git config core.ignorecase false 

# 设置别名（多人使用的环境，不建议使用 --global）
git config alias.s "status -sb"
git config alias.v "checkout vikey"
git config alias.d "checkout develop"
git config alias.m "checkout master"
git config alias.b "branch -vv"
git config alias.pr "remote prune origin"
git config alias.lg "log --color --graph --pretty=format:'% Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit"

# 编辑 git 配置文件
git config -e 
```
