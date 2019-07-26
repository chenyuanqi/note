
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

.git 目录说明。  
.git/config 项目配置文件；  
.git/HEAD HEAD 文件；  
.git/refs/tags 标签。  

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

# 修改并提交到本地仓库
echo "append some content and excute commit." >> first.txt
git add first.txt
git commit -m "Try to append"
# 再次修改
echo "append more." >> first.txt
# 修正之前的提交
git add first.txt
git commit --amend -m "Try to append more."

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
git log --pretty=format:'%h %ad | %s%d [%an]' --graph --date=short

# 移动文件
mkdir src
git mv *.txt src
git commit -m "Moved .txt to src"
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

### 分支操作 
```bash
# 创建分支
git checkout -b new_branch

echo "The third file" > third.txt
git add third.txt
git commit -m "Add file"

# 切换分支
git checkout master

echo "This is a example file" > README.md
git add -a
git commit -m "Added README"

# 查看分叉的分支历史
git log --all

# 合并分支
git checkout new_branch
git merge master

# 冲突解决
git checkout master
echo "make conflict" >> README.md
git commit -am "make conflict"
git checkout new_branch
git merge master
# 修复冲突，编辑保留需要的代码
git add README.md
git commit -m "fixed conflict"

# 使用变基合并代码，回到合并前的代码
git log
git reset reset --hard <hash>
git checkout master
git log
git reset reset --hard <hash>
# 变基
git checkout new_branch
git rebase master
# 变基的最终结果与合并很相似，但是提交树却十分不同
# 变基后 new_branch 分支的提交树已被重写，以致 master 分支成为了其提交历史的一部分。这样使得提交链更加线性，且更易阅读
# 针对短期生命的本地分支使用变基，而对公开仓库的分支使用合并

# 合并回 master
git checkout master
git merge new_branch

# 
```
