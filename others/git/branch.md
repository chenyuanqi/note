
### Git 分支
分支代表了一条独立的开发流水线。  
在 Git 中，分支是你日常开发流程中的一部分。当你想要添加一个新的功能或是修复一个 bug 时 —— 不管 bug 是大是小 —— 你都应该新建一个分支来封装你的修改。这确保了不稳定的代码永远不会被提交到主代码库中，它同时给了你机会，在并入主分支前清理你 feature 分支的历史。  
分支代表了一系列提交的顶端 —— 而不是提交的容器，分支历史通过提交之间的关系来推断。这使得 Git 的合并模型变成了动态的。

```bash
# 列出仓库中所有分支
git branch

# 创建一个名为 <branch> 的分支
git branch <branch>

# 删除指定分支，Git 会阻止你删除包含未合并更改的分支
git branch -d <branch>
# 强制删除指定分支
git branch -D <branch>

# 将当前分支命名为 <branch>
git branch -m <branch>

# 查看特定分支
git checkout <existing-branch>

# 创建并查看 <new-branch>，-b 选项告诉 Git 在运行 git checkout <new-branch> 之前运行 git branch <new-branch>
git checkout -b <new-branch> <existing-branch>
```

**分离的 HEAD**  
HEAD 是 Git 指向当前快照的引用。  
git checkout 命令内部只是更新 HEAD，指向特定分支或提交。当它指向分支时，Git 不会报错，但当你 `checkout <commit>` 或 `checkout tag_name` 时，它会进入「分离 HEAD」状态。  
```bash
git checkout <commit>
```
有个警告会告诉你所做的更改和项目的其余历史处于「分离」的状态。  

当 HEAD 处于分离状态（不依附于任一分支）时，提交操作可以正常进行，但是不会更新任何已命名的分支。一旦此后你切换到别的分支，比如说 master，那么这个提交节点（可能）再也不会被引用到，然后就会被丢弃掉了。但是，如果你想保存这个状态，可以用命令 git checkout -b branch_name 来创建一个新的分支。  

你应该永远在分支上开发 —— 而绝不在分离的 HEAD 上。这样确保你一直可以引用到你的新提交。不过，如果你只是想查看旧的提交，那么是否处于分离 HEAD 状态并不重要。  

