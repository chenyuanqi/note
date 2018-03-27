
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
# 此时，可以撤销 add 操作
git reset HEAD [file_path]

# 提交到「暂存区」
git commit -m"commit detailed message."
# add + commit 所有文件已跟踪过的文件
git commit -am"commit detailed message."

# 推送到远程分支
git push remote_branch_name local_branch_name
```
