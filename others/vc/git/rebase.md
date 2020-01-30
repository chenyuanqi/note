
### rebase: 在新位置重新提交
有些人不喜欢 `merge`，因为在 `merge` 之后，`commit` 历史就会出现分叉，这种分叉再汇合的结构会让有些人觉得混乱而难以管理。如果你不希望 `commit` 历史出现分叉，可以用 `rebase` 来代替 `merge`。

`rebase` ，又是一个中国人看不懂的词。这个词的意思，你如果查一下的话：

![](https://user-gold-cdn.xitu.io/2017/11/21/15fdea7b6422ac3a?w=160&h=128&f=jpeg&s=8608)

> 哈？玩个 Git 就弯了？

其实这个翻译还是比较准确的。`rebase` 的意思是，给你的 `commit` 序列重新设置基础点（也就是父 `commit`）。展开来说就是，把你指定的 `commit` 以及它所在的 `commit` 串，以指定的目标 `commit` 为基础，依次重新提交一次。例如下面这个 `merge`：

```
git merge branch1

```

![](https://user-gold-cdn.xitu.io/2017/11/21/15fdea7b6646a1f3?w=640&h=454&f=gif&s=175263)

如果把 `merge` 换成 `rebase`，可以这样操作：

```
git checkout branch1
git rebase master

```

![](https://user-gold-cdn.xitu.io/2017/11/30/1600abd620a8e28c?w=698&h=518&f=gif&s=337134)

可以看出，通过 `rebase`，`5` 和 `6` 两条 `commit`s 把基础点从 `2` 换成了 `4` 。通过这样的方式，就让本来分叉了的提交历史重新回到了一条线。这种「重新设置基础点」的操作，就是 `rebase` 的含义。

另外，在 `rebase` 之后，记得切回 `master` 再 `merge` 一下，把 `master` 移到最新的 `commit`：

```
git checkout master
git merge branch1

```

![](https://user-gold-cdn.xitu.io/2017/12/2/160149e054fe485c?w=706&h=456&f=gif&s=207406)

> 为什么要从 `branch1` 来 `rebase`，然后再切回 `master` 再 `merge` 一下这么麻烦，而不是直接在 `master` 上执行 `rebase`？
> 
> 从图中可以看出，`rebase` 后的 `commit` 虽然内容和 `rebase` 之前相同，但它们已经是不同的 `commits` 了。如果直接从 `master` 执行 `rebase` 的话，就会是下面这样：
> 
> ![](https://user-gold-cdn.xitu.io/2017/12/2/16014b5a6919c0b7?w=650&h=428&f=gif&s=173918)
> 
> 这就导致 `master` 上之前的两个最新 `commit` 被剔除了。如果这两个 `commit` 之前已经在中央仓库存在，这就会导致没法 `push` 了：
> 
> ![](https://user-gold-cdn.xitu.io/2017/12/2/16014bc64d4337f8?w=643&h=640&f=jpeg&s=58468)
> 
> 所以，为了避免和远端仓库发生冲突，一般不要从 `master` 向其他 `branch` 执行 `rebase` 操作。而如果是 `master` 以外的 `branch` 之间的 `rebase`（比如 `branch1` 和 `branch2` 之间），就不必这么多费一步，直接 `rebase` 就好。

### rebase -i：交互式 rebase
`commit --amend` 可以修复最新 `commit` 的错误，但如果是倒数第二个 `commit` 写错了，怎么办？

如果不是最新的 `commit` 写错，就不能用 `commit --amend` 来修复了，而是要用 `rebase`。不过需要给 `rebase` 也加一个参数：`-i`。

`rebase -i` 是 `rebase --interactive` 的缩写形式，意为「交互式 rebase」。

所谓「交互式 rebase」，就是在 `rebase` 的操作执行之前，你可以指定要 `rebase` 的 `commit` 链中的每一个 `commit` 是否需要进一步修改。

那么你就可以利用这个特点，进行一次「原地 rebase」。

例如你是在写错了 `commit` 之后，又提交了一次才发现之前写错了：

```
git log
```

![](https://user-gold-cdn.xitu.io/2017/11/22/15fdf5fd00a27f45?w=614&h=290&f=jpeg&s=115583)

**开启交互式 rebase 过程**  
现在再用 `commit --amend` 已经晚了，但可以用 `rebase -i`：

```
git rebase -i HEAD^^
```

> 说明：在 Git 中，有两个「偏移符号」： `^` 和 `~`。
> 
> `^` 的用法：在 `commit` 的后面加一个或多个 `^` 号，可以把 `commit` 往回偏移，偏移的数量是 `^` 的数量。例如：`master^` 表示 `master` 指向的 `commit` 之前的那个 `commit`； `HEAD^^` 表示 `HEAD` 所指向的 `commit` 往前数两个 `commit`。
> 
> `~` 的用法：在 `commit` 的后面加上 `~` 号和一个数，可以把 `commit` 往回偏移，偏移的数量是 `~` 号后面的数。例如：`HEAD~5` 表示 `HEAD` 指向的 `commit`往前数 5 个 `commit`。

上面这行代码表示，把当前 `commit` （ `HEAD` 所指向的 `commit`） `rebase` 到 `HEAD` 之前 2 个的 `commit` 上：

![](https://user-gold-cdn.xitu.io/2017/11/22/15fdf5fd00522381?w=568&h=352&f=gif&s=182502)

如果没有 `-i` 参数的话，这种「原地 rebase」相当于空操作，会直接结束。而在加了 `-i` 后，就会跳到一个新的界面：

![](https://user-gold-cdn.xitu.io/2017/11/22/15fdf5fd04f46d6e?w=590&h=409&f=jpeg&s=137238)

**编辑界面：选择 commit 和对应的操作**  
这个编辑界面的最顶部，列出了将要「被 rebase」的所有 `commit`s，也就是倒数第二个 `commit` 「增加常见笑声集合」和最新的 `commit`「增加常见哭声集合」。需要注意，这个排列是正序的，旧的 `commit` 会排在上面，新的排在下面。

这两行指示了两个信息：

1.  需要处理哪些 `commit`s；
2.  怎么处理它们。

你需要修改这两行的内容来指定你需要的操作。每个 `commit` 默认的操作都是 `pick` （从图中也可以看出），表示「直接应用这个 `commit`」。所以如果你现在直接退出编辑界面，那么结果仍然是空操作。

但你的目标是修改倒数第二个 `commit`，也就是上面的那个「增加常见笑声集合」，所以你需要把它的操作指令从 `pick` 改成 `edit` 。 `edit` 的意思是「应用这个 commit，然后停下来等待继续修正」。其他的操作指令，在这个界面里都已经列举了出来（下面的 "Commands…" 部分文字），你可以自己研究一下。

![](https://user-gold-cdn.xitu.io/2017/11/22/15fdf5fd020c87f6?w=535&h=254&f=jpeg&s=86065)

把 `pick` 修改成 `edit` 后，就可以退出编辑界面了：

![](https://user-gold-cdn.xitu.io/2017/11/22/15fdf5fd007159fa?w=454&h=169&f=jpeg&s=44688)

上图的提示信息说明，`rebase` 过程已经停在了第二个 `commit` 的位置，那么现在你就可以去修改你想修改的内容了。

**修改写错的 commit**  
修改完成之后，和上节里的方法一样，用 `commit --amend` 来把修正应用到当前最新的 `commit`：

```
git add 笑声
git commit --amend
```

![](https://user-gold-cdn.xitu.io/2017/11/22/15fdf5fd04de0d40?w=407&h=96&f=jpeg&s=41778)

**继续 rebase 过程**  
在修复完成之后，就可以用 `rebase --continue` 来继续 `rebase` 过程，把后面的 `commit` 直接应用上去。

```
git rebase --continue
```

![](https://user-gold-cdn.xitu.io/2017/11/22/15fdf5fd54455c29?w=459&h=61&f=jpeg&s=29623)

然后，这次交互式 `rebase` 的过程就完美结束了，你的那个倒数第二个写错的 `commit` 就也被修正了：

![](https://user-gold-cdn.xitu.io/2017/11/22/15fdf5fd4e7d5257?w=548&h=348&f=gif&s=266072)

实质上，交互式 `rebase` 并不是必须应用在「原地 rebase」上来修改写错的 `commit` ，这只不过是它最常见的用法。你同样也可以把它用在分叉的 `commit` 上，不过这个你就可以自己去研究一下了。

### 用交互式 rebase 撤销提交
互式 `rebase` 可以用来修改某些旧的 `commit`s。其实除了修改提交，它还可以用于撤销提交。比如下面这种情况：

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe243fc7996318?w=536&h=354&f=jpeg&s=118972)

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe243fc74f48c7?w=447&h=273&f=jpeg&s=16849)

你想撤销倒数第二条 `commit`，那么可以使用 `rebase -i`：

```
git rebase -i HEAD^^

```

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe243fc7ac1154?w=554&h=261&f=jpeg&s=96953)

然后就会跳到 `commit` 序列的编辑界面，这个在之前的第 10 节里已经讲过了。和第 10 节一样，你需要修改这个序列来进行操作。不过不同的是，之前讲的修正 `commit` 的方法是把要修改的 `commit` 左边的 `pick` 改成 `edit`，而如果你要撤销某个 `commit` ，做法就更加简单粗暴一点：直接删掉这一行就好。

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe243fcf5f6607?w=542&h=240&f=jpeg&s=77581)

`pick` 的直接意思是「选取」，在这个界面的意思就是应用这个 `commit`。而如果把这一行删掉，那就相当于在 `rebase` 的过程中跳过了这个 `commit`，从而也就把这个 `commit` 撤销掉了。

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe243fce5804fd?w=548&h=330&f=gif&s=326602)

现在再看 `log`，就会发现之前的倒数第二条 `commit` 已经不在了。

```
git log

```

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe243fc7eb3b31?w=528&h=220&f=jpeg&s=71201)

**用 rebase --onto 撤销提交**  
除了用交互式 `rebase` ，你还可以用 `rebase --onto` 来更简便地撤销提交。

`rebase` 加上 `--onto` 选项之后，可以指定 `rebase` 的「起点」。一般的 `rebase`，告诉 Git 的是「我要把当前 `commit` 以及它之前的 `commit`s 重新提交到目标 `commit` 上去，这其中，`rebase` 的「起点」是自动判定的：选取当前 `commit` 和目标 `commit` 在历史上的交叉点作为起点。

例如下面这种情况：

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe24400508e3c8?w=515&h=360&f=jpeg&s=19522)

如果在这里执行：

```
git rebase 第3个commit

```

那么 Git 会自动选取 `3` 和 `5` 的历史交叉点 `2` 作为 `rebase` 的起点，依次将 `4` 和 `5` 重新提交到 `3` 的路径上去。

而 `--onto` 参数，就可以额外给 rebase 指定它的起点。例如同样以上图为例，如果我只想把 `5` 提交到 `3` 上，不想附带上 `4`，那么我可以执行：

```
git rebase --onto 第3个commit 第4个commit branch1

```

`--onto` 参数后面有三个附加参数：目标 `commit`、起点 `commit`（注意：rebase 的时候会把起点排除在外）、终点 `commit`。所以上面这行指令就会从 `4` 往下数，拿到 `branch1` 所指向的 `5`，然后把 `5` 重新提交到 `3` 上去。

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe24400d7d73d0?w=534&h=552&f=gif&s=199563)

同样的，你也可以用 `rebase --onto` 来撤销提交：

```
git rebase --onto HEAD^^ HEAD^ branch1

```

上面这行代码的意思是：以倒数第二个 `commit` 为起点（起点不包含在 `rebase` 序列里哟），`branch1` 为终点，`rebase` 到倒数第三个 `commit` 上。

也就是这样：

![](https://user-gold-cdn.xitu.io/2017/11/22/15fe243fce5804fd?w=548&h=330&f=gif&s=326602)

### 小结

本节介绍的是 `rebase` 指令，它可以改变 `commit` 序列的基础点。它的使用方式很简单：

```
git rebase 目标基础点

```

需要说明的是，`rebase` 是站在需要被 `rebase` 的 `commit` 上进行操作，这点和 `merge` 是不同的。

交互式 `rebase`，它可以在 `rebase` 开始之前指定一些额外操作。交互式 `rebase` 最常用的场景是修改写错的 `commit`，但也可以用作其他用途。它的大致用法：

1.  使用方式是 `git rebase -i 目标commit`；
2.  在编辑界面中指定需要操作的 `commit`s 以及操作类型；
3.  操作完成之后用 `git rebase --continue` 来继续 `rebase` 过程。

「撤销过往的提交」 方法有两种：

1.  用 `git rebase -i` 在编辑界面中删除想撤销的 `commit`s
2.  用 `git rebase --onto` 在 rebase 命令中直接剔除想撤销的 `commit`s

方法有两种，理念是一样的：在 `rebase` 的过程中去掉想撤销的 `commit`，让他它消失在历史中。
