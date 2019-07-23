
### 设置
设置 Git 以便准备开始工作
```bash
# 设置用户名和邮箱
git config --global user.name "Your Name"
git config --global user.email "your_email@whatever.com"

# 设置行尾首选项
git config --global core.autocrlf input
git config --global core.safecrlf true
```

### 创建项目与基本操作
```bash
mkdir demo
cd demo
echo "this is a demo" > first.txt

# 初始化 git
git init
# 添加文件到暂存区
git add first.txt
# 提交文件到本地仓库
git commit -m "First Commit"
# 查看当前 git 状态
git status
# 再进行修改
echo "update." >> first.txt
# 撤销工作区的修改【撤销】
git checkout first.txt

# 修改一下并添加到暂存区
echo "update one more." >> first.txt
git add first.txt
# 撤销暂存区的内容【撤销】
git reset HEAD first.txt

# 修改并提交到本地仓库
echo "append some content and excute commit." >> first.txt
git add first.txt
git commit -m "Try to revert our commit"
# 撤销已经提交的更改【撤销】
git revert HEAD --no-edit

# 加个文件
echo "the second file" > second.txt
git add second.txt
# 对同一个文件二次修改
echo "append the comment" >> second.txt
# 此时的状态是一次更改已被暂存，且准备提交；二次更改还未暂存（如果现在提交，那么第二次的修改将不会保存到本地仓库中）
git status
git commit -m "Second commit"
# 此时 second.txt 还有修改未提交到暂存区，可以再次提交
git add second.txt
git commit -m "Append comment"
# 查看提交的历史
git log --pretty=format:'%h %ad | %s%d [%an]' --graph --date=short'
```

### 别名
设置别名，更方便的操作（$HOME/.gitconfig）
```
[alias]
  co = checkout
  ci = commit
  st = status
```
 Shell 支持别名或简写 ($HOME/.profile)
 ```
alias gs='git status '
alias gc='git commit'
 ```

### 给版本打标签
```bash
# 添加 v1 版本的测试版标签
git tag v1-beta
# 查看已有的标签
git tag
# 从已有的标签中检出代码
git checkout v1-beta

# 删除无效标签
git tag -d v1-beta
```
