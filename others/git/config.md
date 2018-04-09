
```bash
# 全局配置保存在：~/.gitconfig，本地仓库配置保存在：.git/config
# 设置用户信息（多个不同源的仓库，不建议使用 --global）
git config [--global] user.name "xxx"
git config [--global] user.email xxx@xxx.com
# 设置 UI 及大小写敏感
git config [--global] color.ui auto
git config [--global] core.ignorecase false 

# 设置别名（多人使用的环境，不建议使用 --global）
git config [--global] alias.s "status -sb"
git config [--global] alias.v "checkout vikey"
git config [--global] alias.d "checkout develop"
git config [--global] alias.m "checkout master"
git config [--global] alias.b "branch -vv"
git config [--global] alias.pr "remote prune origin"
git config [--global] alias.lg "log --color --graph --pretty=format:'% Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit"
```
