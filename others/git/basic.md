
```bash
# 生成公钥 ~/.ssh/id_rsa.pub （用于仓库 ssh key 白名单）
ssh-keygen -t rsa -C "xxx@xxx.com"

# 创建线上项目后，拉取到本地
git clone git@xxx
cd xxx
# 变更文件后，查看当前 git 状态
git status [-sb]
# -sb 参数表示开启小白显示

# 添加到指定变更文件，增加跟踪、预备 commit
git add file_path
# 添加所有变更且未被忽略(gitignore)的文件 
git add .
# 对于同一个文件的多处变化，可以多次提交
git add -p [file_path]
# 此时，可以撤销 add 操作
git reset HEAD [file_path]

# 删除工作区文件，并且将这次删除放入暂存区
git rm [file1] [file2] ...

# 停止追踪指定文件，但该文件会保留在工作区
git rm --cached [file]

# 改名文件，并且将这个改名放入暂存区
git mv [file_original] [file_renamed]

# 提交到「暂存区」
git commit -m"commit detailed message."
# add + commit 所有文件已跟踪过的文件
git commit -am"commit detailed message."
# 使用一次新的commit，替代上一次提交
# 如果代码没有任何新变化，则用来改写上一次commit的提交信息
git commit --amend -m [message]

# 建立追踪关系，在现有分支与指定的远程分支之间
git branch --set-upstream [local_branch_name] [remote_branch_name]

# 推送到远程分支
git push remote_branch_name local_branch_name
# 如 git push origin master

# 删除本地分支
git branch -D local_branch_name
# 删除远程分支
git push origin :remote_branch_name

# 列出所有本地分支
$ git branch

# 列出所有远程分支
$ git branch -r

# 列出所有本地分支和远程分支
$ git branch -a

# 新建一个分支，但依然停留在当前分支
$ git branch [local_branch_name]

# 新建一个分支，并切换到该分支(+绑定远程分支)
$ git checkout -b [local_branch_name] [origin/remote_branch_name]

# 新建一个分支，指向指定commit
$ git branch [local_branch_name] [commit]

# 新建一个分支，与指定的远程分支建立追踪关系
$ git branch --track [local_branch_name] [remote_branch_name]

# 切换到指定分支，并更新工作区
git checkout [local_branch_name]

# 切换上一个分支
git checkout -

# 合并指定分支到当前分支
git merge [local_branch_name]

# 列出所有tag
git tag

# 新建一个tag在当前commit
git tag [tag]

# 新建一个tag在指定commit
git tag [tag] [commit]

# 删除本地tag
git tag -d [tag]

# 删除远程tag
git push origin :refs/tags/[tagName]

# 查看tag信息
git show [tag]

# 提交指定tag
git push [remote] [tag]

# 提交所有tag
git push [remote] --tags

# 新建一个分支，指向某个tag
git checkout -b [branch] [tag]

# 显示当前分支的版本历史
git log

# 显示commit历史，以及每次commit发生变更的文件
git log --stat

# 搜索提交历史，根据关键词
git log -S [keyword]

# 显示某个commit之后的所有变动，每个commit占据一行
git log [tag] HEAD --pretty=format:%s

# 显示某个commit之后的所有变动，其"提交说明"必须符合搜索条件
git log [tag] HEAD --grep feature

# 显示某个文件的版本历史，包括文件改名
git log --follow [file]
git whatchanged [file]

# 显示指定文件相关的每一次diff
git log -p [file]

# 显示过去5次提交
git log -5 --pretty --oneline

# 显示所有提交过的用户，按提交次数排序
git shortlog -sn

# 显示指定文件是什么人在什么时间修改过
git blame [file]

# 显示暂存区和工作区的差异
git diff

# 显示暂存区和上一个commit的差异
git diff --cached [file]

# 显示工作区与当前分支最新commit之间的差异
git diff HEAD

# 显示两次提交之间的差异
git diff [first-branch]...[second-branch]

# 显示今天你写了多少行代码
git diff --shortstat "@{0 day ago}"

# 显示某次提交的元数据和内容变化
git show [commit]

# 显示某次提交发生变化的文件
git show --name-only [commit]

# 显示某次提交时，某个文件的内容
git show [commit]:[filename]

# 显示当前分支的最近几次提交
git reflog

# 下载远程仓库的所有变动
git fetch [remote_branch_name]

# 显示所有远程仓库
git remote -v

# 显示某个远程仓库的信息
git remote show [remote_branch_name]

# 增加一个新的远程仓库，并命名
git remote add [shortname] [url]

# 取回远程仓库的变化，并与本地分支合并
git pull [remote_branch_name] [local_branch_name]

# 上传本地指定分支到远程仓库
git push [remote_branch_name] [local_branch_name]

# 强行推送当前分支到远程仓库，即使有冲突
git push [remote_branch_name] --force

# 推送所有分支到远程仓库
git push [remote_branch_name] --all

# 恢复暂存区的指定文件到工作区
git checkout [file]

# 恢复某个commit的指定文件到暂存区和工作区
git checkout [commit] [file]

# 恢复暂存区的所有文件到工作区
git checkout .

# 重置暂存区的指定文件，与上一次commit保持一致，但工作区不变
git reset [file]

# 重置暂存区与工作区，与上一次commit保持一致
git reset --hard

# 重置当前分支的指针为指定commit，同时重置暂存区，但工作区不变
git reset [commit]

# 重置当前分支的HEAD为指定commit，同时重置暂存区和工作区，与指定commit一致
git reset --hard [commit]

# 重置当前HEAD为指定commit，但保持暂存区和工作区不变
git reset --keep [commit]

# 新建一个commit，用来撤销指定commit
# 后者的所有变化都将被前者抵消，并且应用到当前分支
git revert [commit]

# 暂时将未提交的变化移除，稍后再移入
git stash
git stash pop

```
