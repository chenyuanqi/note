

### Git 异常处理
1、本地工作区文件恢复  
本地工作区文件的撤回，撤回到当前分支最新的远程代码。  
```bash
# 语法：git checkout <filename/dirname>
git checkout test.php
```

2、远程分支删除后，删除本地分支及关联  
一般情况下会用本地命令 git branch --set-upstream-to=origin/master master 建立本地分支与远程分支的关联。  
```bash
# 去除本地分支与远程分支的关联
# 语法：git branch --unset-upstream <branchname>
git branch --unset-upstream feature/test

# 本地删除分支及删除远程分支
# 删除本地分支命令: git branch -d/-D <BranchName>
# 删除远程分支命令: git push origin --delete <BranchName>
```

3、修改提交时的备注内容  
命令会修改提交时的 commit-id，即会覆盖原本的提交，需要谨慎操作。  
```bash
# 语法：git commit --amend
```

4、修改分支名  
开发都是拥有极快手速的人，建了个分支一不小心打错了某个字母或者两个字母打反了，可能就与本意存在较大误差了，Git 提供一种已经拉取了分支，在上面开发了不少的内容，但后来发现原本拉的分支名字就有问题的修复方法。  
```bash
# 语法：git branch -m <oldbranch> <newbranch>
git branch -m feature/stor-200830 feature/story-200830
```

5、撤回提交  
①已将更改提交到本地，需要撤回提交。  
文件变更记录与未提交之前的文件变更记录是一致的，只是撤销了 commit 的操作。  
```bash
# 语法： git reset --soft [<commit-id>/HEAD~n>]
git reset --soft HEAD~1
```
②用新的更改替换撤回的更改  
提交之中可能有些地方需要优化，我们可以撤销本次的 commit 以及文件暂存状态，修改之后再重新添加到暂存区进行提交。已变更的文件都未添加到暂存区，撤销了 commit 和 add 的操作。  
```bash
# 语法： git reset --mixed [<commit-id>/HEAD~n>]
git reset --mixed HEAD~1
```
③本地提交了错误的文件  
本地将完全错误的，本不应提交的内容提交到了仓库，需要进行撤销，可以使用 --hard 参数。已追踪文件的变更内容都消失了，撤销了 commit 和 add 的操作，同时撤销了本地已追踪内容的修改；未追踪的内容不会被改变。从上面的效果可以看到，文件的修改都会被撤销。--hard 参数需要谨慎使用。  
```bash
# 语法： git reset --hard [<commit-id>/HEAD~n>]
git reset --hard HEAD~1
```  

6、撤销本地分支合并  
实际操作中，总会有很多的干扰，导致我们合并了并不该合并的分支到目标分支上。解决这种问题的方式有两种，git reset 和 git revert。从需要提交到远程分支的角度来讲，reset 能够 “毁尸灭迹”，不让别人发现我们曾经错误的合并过分支（注：多人协作中，需要谨慎使用）；revert 则会将合并分支和撤回记录一并显示在远程提交记录上。  
```bash
# 语法：git revert <commit-id>
git revert 700920
```

7、恢复误删的本地分支  
本地分支拉取之后，由于疏忽被删除，而且本地的分支并没有被同步到远程分支上，此时想要恢复本地分支。  
```bash
# 查看该仓库下的所有历史操作
git reflog
# 语法：git checkout -b <branch-name> <commit-id>
git checkout -b feature/delete HEAD@{2}
```
命令执行完成后，分支恢复到 HEAD@{2} 的快照，即从 master 分支拉取 feature/delete 分支的内容，仍然缺少 “新增 xxx 文件” 的提交，直接将文件内容恢复到最新的提交内容，使用命令 git reset --hard HEAD@{1} 即可实现硬性覆盖本地工作区内容的目的。git reflog 命令获取到的内容为本地仓库所有发生过的变更，可谓恢复利器。  

8、不确定哪个分支有自己提交的 commit  
工作中会经常碰到一种场景，某个提交先后合并到了各个分支上，但后来发现提交的这个修改是有问题的，需要排查到底哪个分支包含这个提交，然后将其修复掉。  
```bash
# 需要先确定有问题的提交的 commit-id
# 然后查看本地所有的分支
# Git 提供了一种能够直接通过 commit-id 查找出包含该内容分支的命令
# 语法：git branch --contains <commit-id>
git branch --contains 700920
```



