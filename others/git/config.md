
```bash
# 设置用户信息（多个不同源的仓库，不建议使用 --global）
git config [--global] user.name "xxx"
git config [--global] user.email xxx@xxx.com
git config [--global] color.ui auto

# 设置别名（多人使用的环境，不建议使用 --global）
git config [--global] alias.s "status -sb"
git config [--global] alias.v "checkout vikey"
git config [--global] alias.d "checkout develop"
git config [--global] alias.m "checkout master"
git config [--global] alias.b "branch -vv"
git config [--global] alias.pr "remote prune origin"
git config [--global] alias.lg "log --color --graph --pretty=format:'% Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit"
```
