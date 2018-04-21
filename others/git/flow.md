
- git-flow 分支
> master 分支：用于上线的分支，保护性分支，只包含经过测试的稳定代码，开发人员不能直接工作在此分支上，也不能直接提交改动到 master 分支上。  
> develop分支：是开发人员进行任何新的开发的基础分支，当开始一个新的 feature 分支的时候，要从 develop 分出去；另外此分支也汇集所有的已完成的功能，等待合并到 master 分支上线。  
> 上面两个分支被称为「长期分支」，存在于项目的整个生命周期中；其他分支，是临时性的，根据需要来创建，当完成了自己的任务后，就会被删掉。  
> feature 分支：平常的开发工作使用最频繁的分支，功能分支。该分支默认从 develop 检出，在做功能性开发的时候，检出一个独立的分支，是版本控制中一个重要的原则。基于 develop，完成后 merge 回 develop。
> release 分支：准备要发布版本的分支, 用来修复 bug。 基于 develop, 完成后 merge 回 master 和 develop。  
> hotfix 分支：修复 master 上的问题, 等不及 release 版本就必须马上上线. 基于 master, 完成后 merge 回 master 和 develop 分支。  

```bash
# 目前流行的是 avh 版本的 git-flow
# 稳定版
brew install git-flow-avh
sudo apt-get install git-flow

# 初始化项目，不再需要 git init
# 分支命名的约定，使用默认值即可
git flow init

# 创建功能分支
git flow feature start [branch_name]

# 发布功能分支，方便其他人也可以用
git flow feature publish [branch_name]

# 跟踪功能分支
git flow feature track [branch_name]

# 获取功能分支，并拉取远程的变更
git flow feature pull origin [branch_name]

# 完成功能分支
git flow feature finish [branch_name]

# 创建 release 分支
git flow release start 1.1.9

# 发布 release 分支，方便其他人也可以用
git flow release publish 1.1.9

# 跟踪 release 分支
git flow release track 1.1.9

# 完成 release 分支
git flow release finish 1.1.9

# 创建 hotfix 分支
git flow hotfix start bug_fixed

# 完成 hotfix 分支
git flow hotfix finish bug_fixed
```
