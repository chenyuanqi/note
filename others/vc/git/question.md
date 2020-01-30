
### 常见问题

**代码已经 push 上去了才发现写错？**  
有的时候，代码 `push` 到了中央仓库，才发现有个 `commit` 写错了。这种问题的处理分两种情况：

- 1、出错的内容在你自己的 branch

假如是某个你自己独立开发的 `branch` 出错了，不会影响到其他人，那没关系，把写错的 `commit` 修改或者删除掉，然后再 `push` 上去就好了。不过……

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe2638ac5c1dd0?w=676&h=162&f=jpeg&s=95234)

由于你在本地对已有的 `commit` 做了修改，这时你再 `push` 就会失败，因为中央仓库包含本地没有的 `commit`s。但这个和前面讲过的情况不同，这次的冲突不是因为同事 `push` 了新的提交，而是因为你刻意修改了一些内容，这个冲突是你预料到的，你本来就希望用本地的内容覆盖掉中央仓库的内容。那么这时就不要乖乖听话，按照提示去先 `pull` 一下再 `push` 了，而是要选择「强行」`push`：

```
git push origin branch1 -f
```

`-f` 是 `--force` 的缩写，意为「忽略冲突，强制 `push`」。

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe2638ab7b7e6d?w=507&h=154&f=jpeg&s=74359)

这样，在本地修改了错误的 `commit`s，然后强制 `push` 上去，问题就解决了。

- 2、出错的内容已经合并到 master

这就不能用上面那招了。同事的工作都在 `master` 上，你永远不知道你的一次强制 `push` 会不会洗掉同事刚发上去的新提交。所以除非你是人员数量和行为都完全可控的超小团队，可以和同事做到无死角的完美沟通，不然一定别在 `master` 上强制 `push`。

在这种时候，你只能退一步，选用另一种策略：增加一个新的提交，把之前提交的内容抹掉。例如之前你增加了一行代码，你希望撤销它，那么你就做一个删掉这行代码的提交；如果你删掉了一行代码，你希望撤销它，那么你就做一个把这行代码还原回来的提交。这种事做起来也不算麻烦，因为 Git 有一个对应的指令：`revert`。

它的用法很简单，你希望撤销哪个 `commit`，就把它填在后面：

```
git revert HEAD^
```

上面这行代码就会增加一条新的 `commit`，它的内容和倒数第二个 `commit` 是相反的，从而和倒数第二个 `commit` 相互抵消，达到撤销的效果。

在 `revert` 完成之后，把新的 `commit` 再 `push` 上去，这个 `commit` 的内容就被撤销了。它和前面所介绍的撤销方式相比，最主要的区别是，这次改动只是被「反转」了，并没有在历史中消失掉，你的历史中会存在两条 `commit` ：一个原始 `commit` ，一个对它的反转 `commit`。

**reset 的本质是什么**  
在最新的 `commit` 写错时，可以用 `reset --hard` 来把 `commit` 撤销：

```
git reset --hard HEAD^

```

> 用这行代码可以撤销掉当前 `commit`

- reset 的本质：移动 HEAD 以及它所指向的 branch
实质上，`reset` 这个指令虽然可以用来撤销 `commit` ，但它的实质行为并不是撤销，而是移动 `HEAD` ，并且「捎带」上 `HEAD` 所指向的 `branch`（如果有的话）。也就是说，`reset` 这个指令的行为其实和它的字面意思 "reset"（重置）十分相符：它是用来重置 `HEAD` 以及它所指向的 `branch` 的位置的。

而 `reset --hard HEAD^` 之所以起到了撤销 `commit` 的效果，是因为它把 `HEAD` 和它所指向的 `branch` 一起移动到了当前 `commit` 的父 `commit` 上，从而起到了「撤销」的效果：

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe19c8a3235853?w=466&h=262&f=gif&s=120940)

> Git 的历史只能往回看，不能向未来看，所以把 `HEAD` 和 `branch` 往回移动，就能起到撤回 `commit` 的效果。

所以同理，`reset --hard` 不仅可以撤销提交，还可以用来把 `HEAD` 和 `branch` 移动到其他的任何地方。

```
git reset --hard branch2

```

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe333cb605b0de?w=434&h=642&f=gif&s=154560)

不过……`reset` 后面总是跟着的那个 `--hard` 是什么意思呢？

`reset` 指令可以重置 `HEAD` 和 `branch` 的位置，不过在重置它们的同时，对工作目录可以选择不同的操作，而对工作目录的操作的不同，就是通过 `reset` 后面跟的参数来确定的。

- reset --hard：重置工作目录
`reset --hard` 会在重置 `HEAD` 和 `branch` 的同时，重置工作目录里的内容。当你在 `reset` 后面加了 `--hard` 参数时，你的工作目录里的内容会被完全重置为和 `HEAD` 的新位置相同的内容。换句话说，就是你的未提交的修改会被全部擦掉。

例如你在上次 `commit` 之后又对文件做了一些改动：

```
git status

```

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe333cb5a0e894?w=621&h=236&f=jpeg&s=71681)

然后，你执行了 `reset` 并附上了 `--hard` 参数：

```
git reset --hard HEAD^

```

你的 `HEAD` 和当前 `branch` 切到上一条 `commit` 的同时，你工作目录里的新改动也一起全都消失了，不管它们是否被放进暂存区：

```
git status

```

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe333cb5dbef68?w=355&h=59&f=jpeg&s=20692)

可以看到，在 `reset --hard` 后，所有的改动都被擦掉了。

- reset --soft：保留工作目录
`reset --soft` 会在重置 `HEAD` 和 `branch` 时，保留工作目录和暂存区中的内容，并把重置 `HEAD` 所带来的新的差异放进暂存区。

什么是「重置 `HEAD` 所带来的新的差异」？就是这里：

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe333cb5c6a249?w=478&h=290&f=gif&s=202343)

由于 `HEAD` 从 `4` 移动到了 `3`，而且在 `reset` 的过程中工作目录的内容没有被清理掉，所以 `4` 中的改动在 `reset` 后就也成了工作目录新增的「工作目录和 `HEAD` 的差异」。这就是上面一段中所说的「重置 `HEAD` 所带来的差异」。

所以在同样的情况下：

```
git status

```

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe333cb5a0e894?w=621&h=236&f=jpeg&s=71681)

假设此时当前 `commit` 的改动内容是新增了 `laughters.txt` 文件：

```
git show --stat

```

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe333cb7cdd727?w=533&h=180&f=jpeg&s=47095)

如果这时你执行：

```
git reset --soft HEAD^

```

那么除了 `HEAD` 和它所指向的 `branch1` 被移动到 `HEAD^` 之外，原先 `HEAD` 处 `commit` 的改动（也就是那个 `laughters.txt` 文件）也会被放进暂存区：

```
git status

```

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe333cb7e6e40b?w=626&h=278&f=jpeg&s=90747)

这就是 `--soft` 和 `--hard` 的区别：`--hard` 会清空工作目录的改动，而 `--soft` 则会保留工作目录的内容，并把因为保留工作目录内容所带来的新的文件差异放进暂存区。

- reset 不加参数：保留工作目录，并清空暂存区
`reset` 如果不加参数，那么默认使用 `--mixed` 参数。它的行为是：保留工作目录，并且清空暂存区。也就是说，工作目录的修改、暂存区的内容以及由 `reset` 所导致的新的文件差异，都会被放进工作目录。简而言之，就是「把所有差异都混合（mixed）放在工作目录中」。

还以同样的情况为例：

```
git status

```

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe333cb5a0e894?w=621&h=236&f=jpeg&s=71681)

> 修改了 `games.txt` 和 `shopping list.txt`，并把 `games.txt` 放进了暂存区。

```
git show --stat

```

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe333cb7cdd727?w=533&h=180&f=jpeg&s=47095)

> 最新的 `commit` 中新增了 `laughters.txt` 文件。

这时如果你执行无参数的 `reset`：

```
git reset HEAD^

```

工作目录的内容和 `--soft` 一样会被保留，但和 `--soft` 的区别在于，它会把暂存区清空：

```
git status

```

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe333d086f9754?w=625&h=256&f=jpeg&s=70740)

总而言之，`reset` 指令的本质：重置 `HEAD` 以及它所指向的 `branch` 的位置。
同时，`reset` 的三种参数：

1.  `--hard`：重置位置的同时，清空工作目录的所有改动；
2.  `--soft`：重置位置的同时，保留工作目录和暂存区的内容，并把重置 `HEAD` 的位置所导致的新的文件差异放进暂存区。
3.  `--mixed`（默认）：重置位置的同时，保留工作目录的内容，并清空暂存区。

除了上面这三种参数，还有一些没有列出的较为不常用的参数；另外除了我讲的功能外，`reset` 其实也还有一些别的功能和用法。不过 `reset` 最关键的功能、用法和本质原理就是上面这些了，想了解更多的话，可以去官网了解一下。


**checkout 的本质**  
在前面的 `branch` 的部分，我说到 `checkout` 可以用来切换 `branch`：

```
git checkout branch2
```

![checkout](https://user-gold-cdn.xitu.io/2017/11/30/160089d53b4f65a5?w=458&h=572&f=gif&s=103354)

checkout

不过实质上，`checkout` 并不止可以切换 `branch`。`checkout` 本质上的功能其实是：签出（ checkout ）指定的 `commit`。

`git checkout branch名` 的本质，其实是把 `HEAD` 指向指定的 `branch`，然后签出这个 `branch` 所对应的 `commit` 的工作目录。所以同样的，`checkout` 的目标也可以不是 `branch`，而直接指定某个 `commit`：

```
git checkout HEAD^^
```

```
git checkout master~5
```

```
git checkout 78a4bc
```

```
git checkout 78a4bc^
```

这些都是可以的。

另外，如果你留心的话可能会发现，在 `git status` 的提示语中，Git 会告诉你可以用 `checkout -- 文件名` 的格式，通过「签出」的方式来撤销指定文件的修改：

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe34cc387ba541?w=616&h=160&f=jpeg&s=58078)

**checkout 和 reset 的不同**  
`checkout` 和 `reset` 都可以切换 `HEAD` 的位置，它们除了有许多细节的差异外，最大的区别在于：`reset` 在移动 `HEAD` 时会带着它所指向的 `branch` 一起移动，而 `checkout` 不会。当你用 `checkout` 指向其他地方的时候，`HEAD` 和 它所指向的 `branch` 就自动脱离了。

事实上，`checkout` 有一个专门用来只让 `HEAD` 和 `branch` 脱离而不移动 `HEAD` 的用法：

```
git checkout --detach
```

执行这行代码，Git 就会把 `HEAD` 和 `branch` 脱离，直接指向当前 `commit`：

![git checkout --detach](https://user-gold-cdn.xitu.io/2017/11/30/1600acce7b90b009?w=590&h=472&f=gif&s=94227)

git checkout --detach

**stash：临时存放工作目录的改动 【紧急任务下达】**  
"stash" 这个词，和它意思比较接近的中文翻译是「藏匿」，是「把东西放在一个秘密的地方以备未来使用」的意思。在 Git 中，`stash` 指令可以帮你把工作目录的内容全部放在你本地的一个独立的地方，它不会被提交，也不会被删除，你把东西放起来之后就可以去做你的临时工作了，做完以后再来取走，就可以继续之前手头的事了。

具体说来，`stash` 的用法很简单。当你手头有一件临时工作要做，需要把工作目录暂时清理干净，那么你可以：

```
git stash

```

就这么简单，你的工作目录的改动就被清空了，所有改动都被存了起来。

然后你就可以从你当前的工作分支切到 `master` 去给你的同事打包了……

打完包，切回你的分支，然后：

```
git stash pop

```

你之前存储的东西就都回来了。很方便吧？

> 注意：没有被 track 的文件（即从来没有被 add 过的文件不会被 stash 起来，因为 Git 会忽略它们。如果想把这些文件也一起 stash，可以加上 \`-u\` 参数，它是 \`--include-untracked\` 的简写。就像这样：  

```
git stash -u

```

**reflog ：引用的 log**  
`branch` 用完就删是好习惯，但有的时候，不小心手残删了一个还有用的 `branch` ，或者把一个 `branch` 删掉了才想起来它还有用，怎么办？

`reflog` 是 "reference log" 的缩写，使用它可以查看 Git 仓库中的引用的移动记录。如果不指定引用，它会显示 `HEAD` 的移动记录。假如你误删了 `branch1` 这个 `branch`，那么你可以查看一下 `HEAD` 的移动历史：

```
git reflog

```

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe3de05468c613?w=602&h=78&f=jpeg&s=51327)

从图中可以看出，`HEAD` 的最后一次移动行为是「从 `branch1` 移动到 `master`」。而在这之后，`branch1` 就被删除了。所以它之前的那个 `commit` 就是 `branch1` 被删除之前的位置了，也就是第二行的 `c08de9a`。

所以现在就可以切换回 `c08de9a`，然后重新创建 `branch1` ：

```
git checkout c08de9a
git checkout -b branch1

```

这样，你刚删除的 `branch1` 就找回来了。

> 注意：不再被引用直接或间接指向的 `commit`s 会在一定时间后被 Git 回收，所以使用 `reflog` 来找回删除的 `branch` 的操作一定要及时，不然有可能会由于 `commit` 被回收而再也找不回来。

`reflog` 默认查看 `HEAD` 的移动历史，除此之外，也可以手动加上名称来查看其他引用的移动历史，例如某个 `branch`：

```
git reflog master

```

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe3de0548714c7?w=629&h=98&f=jpeg&s=63093)
