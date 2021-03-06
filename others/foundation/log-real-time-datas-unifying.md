
### 日志：每个软件工程师都应该知道的有关实时数据的统一概念
原文链接： [The Log: What every software engineer should know about real-time data's unifying abstraction - Jay Kreps](https://engineering.linkedin.com/distributed-systems/log-what-every-software-engineer-should-know-about-real-time-datas-unifying)  

学习笔记：[The Log（我所读过的最好的一篇分布式技术文章）](https://www.cnblogs.com/foreach-break/p/notes_about_distributed_system_and_The_log.html)  

### 开篇
我在六年前加入到 LinkedIn 公司，那是一个令人兴奋的时刻：我们刚开始面临单一庞大的集中式数据库的限制问题，需要过渡到一套专门的分布式系统。 这是一个令人兴奋的经历：我们构建、部署和运行分布式图数据库（distributed graph database）、分布式搜索后端（distributed search backend）、 Hadoop 以及第一代和第二代键值数据存储（key-value store），而且这套系统一直运行至今。  

这个过程中，我学到的最有益的事情是我们所构建这套系统的许多组件其核心都包含了一个很简单的概念：日志。 日志有时会叫成 预先写入日志（write-ahead logs）、提交日志（commit logs）或者事务日志（transaction logs），几乎和计算机本身形影不离， 是许多分布式数据系统（distributed data system）和实时应用架构（real-time application architecture）的核心。  

不懂得日志，你就不可能真正理解数据库、NoSQL 存储、键值存储（key value store）、数据复制（replication）、paxos、Hadoop、版本控制（version control），甚至几乎任何一个软件系统；然而大多数软件工程师对日志并不熟悉。我有意于改变这个现状。 本文我将带你浏览有关日志需要了解的一切，包括日志是什么，如何在数据集成（data integration）、实时处理（real time processing）和系统构建中使用日志。  

### 第一部分：日志是什么？
日志可能是一种最简单的不能再简单的存储抽象，只能追加、按照时间完全有序（totally-ordered）的记录序列。日志看起来的样子：  
![基本日志](https://content.linkedin.com/content/dam/engineering/en-us/blog/migrated/log.png)  

在日志的末尾添加记录，读取日志记录则从左到右。每一条记录都指定了一个唯一的顺序的日志记录编号。  
日志记录的次序（ordering）定义了『时间』概念，因为位于左边的日志记录表示比右边的要早。 日志记录编号可以看作是这条日志记录的『时间戳』。 把次序直接看成是时间概念，刚开始你会觉得有点怪异，但是这样的做法有个便利的性质：解耦了 时间 和 任一特定的物理时钟（physical clock）。引入分布式系统后，这会成为一个必不可少的性质。  
`【译注】 分布式系统的 时间、次序、时钟是个最基础根本的问题`  

日志记录的内容和格式是什么对于本文讨论并不重要。另外，不可能一直给日志添加记录，因为总会耗尽存储空间。稍后我们会再回来讨论这个问题。  
所以，日志 和 文件或数据表（table）并没有什么大的不同。文件是一系列字节，表是由一系列记录组成，而日志实际上只是一种按照时间顺序存储记录的数据表或文件。  
讨论到现在，你可能奇怪为什么要讨论这么简单的概念？只能追加的有序的日志记录究竟又是怎样与数据系统生产关系的？ 答案是日志有其特定的目标：它记录了什么时间发生了什么事情。而对分布式数据系统，在许多方面，这是要解决的问题的真正核心。  
不过，在我们进行更加深入的讨论之前，让我先澄清有些让人混淆的概念。每个程序员都熟悉另一种日志记录的定义 —— 应用使用 syslog 或者 log4j 写入到本地文件里的无结构的错误信息或者追踪信息。为了区分，这种情形的称为『应用日志记录』。 应用日志记录是我说的日志概念的一种退化。两者最大的区别是：文本日志意味着主要用来方便人去阅读，而构建我所说的『日志（journal）』或者『数据日志（data logs）』是用于程序的访问。  
（实际上，如果你深入地思考一下，会觉得人去阅读某个机器上的日志这样的想法有些落伍过时了。 当涉及很多服务和服务器时，这样的做法很快就变得难于管理， 我们的目的很快就变成 输入查询 和 输出用于理解多台机器的行为的图表， 因此，文件中的字句文本 几乎肯定不如 本文所描述的结构化日志 更合适。）  

**数据库中的日志**  
我不知道日志概念的起源 —— 可能就像二分查找（binary search）一样，发明者觉得太简单了而不是一项发明。早在 IBM 的系统 R 出现时候日志就出现了。 在数据库里的用法是在崩溃的时候用它来保持各种数据结构和索引的同步。为了保证操作的原子性（atomic）和持久性（durable)， 在对数据库维护的所有各种数据结构做更改之前，数据库会把要做的更改操作的信息写入日志。 日志记录了发生了什么，而每个表或者索引都是更改历史中的一个投影。由于日志是立即持久化的，发生崩溃时，可以作为恢复其他所有持久化结构的可靠来源。  
随着时间的推移，日志的用途从 ACID 的实现细节成长为数据库间复制数据的一种方法。结果证明，发生在数据库上的更改序列 即是与远程副本数据库（replica database）保持同步 所需的操作。 Oracle、MySQL 和 PostgreSQL 都包括一个日志传送协议（log shipping protocol），传输日志给作为备库（Slave）的复本（replica）数据库。 Oracle 还把日志产品化为一个通用的数据订阅机制，为非 Oracle 数据订阅用户提供了 XStreams和 GoldenGate，在 MySQL 和 PostgreSQL 中类似设施是许多数据架构的关键组件。  
正是由于这样的起源，机器可识别的日志的概念主要都被局限在数据库的内部。日志作为做数据订阅机制的用法似乎是偶然出现的。 但这正是支持各种的消息传输、数据流和实时数据处理的理想抽象。  

**分布式系统中的日志**  
日志解决了两个问题：更改动作的排序和数据的分发，这两个问题在分布式数据系统中更是尤为重要。 协商达成一致的更改动作的顺序（或是协商达成不一致做法并去做有副作用的数据拷贝）是分布式系统设计的核心问题之一。  
分布式系统以日志为中心的方案是来自于一个简单的观察，我称之为状态机复制原理（State Machine Replication Principle）：如果两个相同的、确定性的进程从同一状态开始，并且以相同的顺序获得相同的输入，那么这两个进程将会生成相同的输出，并且结束在相同的状态。  
听起来有点难以晦涩，让我们更加深入的探讨，弄懂它的真正含义。  
确定性（deterministic）意味着处理过程是与时间无关的，而且不让任何其他『带外数据（out of band）』的输入影响处理结果。例如，如果一个程序的输出会受到线程执行的具体顺序影响，或者受到 getTimeOfDay 调用、或者其他一些非重复性事件的影响，那么这样的程序一般被认为是非确定性的。  
进程 状态 是进程保存在机器上的任何数据，在进程处理结束的时候，这些数据要么保存在内存里，要么保存在磁盘上。  
当碰到以相同的顺序输入相同的内容的情况时，应该触发你的条件反射：这个地方要引入日志。下面是个很直觉的意识：如果给两段确定性代码相同的日志输入，那么它们就会生产相同的输出。  
应用到分布式计算中就相当明显了。你可以把用多台机器都执行相同事情的问题化简为实现用分布式一致性日志作为这些处理的输入的问题。 这里日志的目的是把所有非确定性的东西排除在输入流之外，以确保处理这些输入的各个复本（replica）保持同步。  
当你理解了这个以后，状态机复制原理就不再复杂或深奥了：这个原理差不多就等于说的是『确定性的处理过程就是确定性的』。不管怎样，我认为它是分布式系统设计中一个更通用的工具。  
这样方案的一个美妙之处就在于：用于索引日志的时间戳 就像 用于保持副本状态的时钟 —— 你可以只用一个数字来描述每一个副本，即这个副本已处理的最大日志记录的时间戳。 日志中的时间戳一一对应了 副本的完整状态。  
根据写进日志的内容，这个原理可以有不同的应用方式。举个例子，我们可以记录一个服务的输入请求日志，或者从请求到响应服务的状态变化日志，或者服务所执行的状态转换命令的日志。理论上来说，我们甚至可以记录各个副本执行的机器指令序列的日志 或是 所调用的方法名和参数序列的日志。只要两个进程用相同的方式处理这些输入，这些副本进程就会保持一致的状态。  
对日志用法不同群体有不同的说法。数据库工作者通常说成 物理日志（physical logging）和 逻辑日志（logical logging）。物理日志是指记录每一行被改变的内容。逻辑日志记录的不是改变的行而是那些引起行的内容改变的 SQL 语句（insert、update 和 delete 语句）。  
分布式系统文献通常把处理和复制（processing and replication）方案宽泛地分成两种。『状态机器模型』常常被称为主 - 主模型（active-active model）， 记录输入请求的日志，各个复本处理每个请求。 对这个模型做了细微的调整称为『主备模型』（primary-backup model），即选出一个副本做为 leader，让 leader 按请求到达的顺序处理请求，并输出它请求处理的状态变化日志。 其他的副本按照顺序应用 leader 的状态变化日志，保持和 leader 同步，并能够在 leader 失败的时候接替它成为 leader。  
![分布式日志请求](https://content.linkedin.com/content/dam/engineering/en-us/blog/migrated/active_and_passive_arch.png)  

为了理解两种方式的差异，我们来看一个不太严谨的例子。假定有一个要复制的『算法服务』，维护一个独立的数字作为它的状态（初始值为 0），可以对这个值进行加法和乘法运算。 主 - 主方式所做的可能是输出所进行的变换的日志，比如『+1』、『*2』等。各个副本都会应用这些变换，从而经过一系列相同的值。 而主备方式会有一个独立的 Master 执行这些变换并输出结果日志，比如『1』、『3』、『6』等。 这个例子也清楚的展示了为什么说顺序是保证各副本间一致性的关键：加法和乘法的顺序的改变将会导致不同的结果。  
分布式日志可以看作是建模一致性（consensus）问题的数据结构。因为日志代表了『下一个』追加值的一系列决策。 你需要眯起眼睛才能从 Paxos 算法簇中找到日志的身影，尽管构建日志是它们最常见的实际应用。 Paxos 通过称为 multi-paxos 的一个扩展协议来构建日志，把日志建模为一系列一致性值的问题，日志的每个记录对应一个一致性值。 日志的身影在 ZAB、RAFT 和 Viewstamped Replication 等其它的协议中明显得多，这些协议建模的问题直接就是维护分布式一致的日志。  
个人有一点感觉，在这个问题上，我们的思路被历史发展有些带偏了，可能是由于过去的几十年中，分布式计算的理论远超过了其实际应用。 在现实中，一致性问题是有点被过于简单化了。计算机系统几乎不需要决定单个的值，要的是处理一序列的请求。所以，日志而不是一个简单的单值寄存器，是更自然的抽象。  
此外，对算法的专注掩盖了系统底层所需的日志抽象。个人觉得，我们最终会更关注把日志作为一个商品化的基石而不是考虑它的实现，就像我们经常讨论哈希表而不纠结于它的细节，比如使用线性探测的杂音哈希（the murmur hash with linear probing）还是某个变种。日志将成为一种大众化的接口，可以有多种竞争的算法和实现，以提供最好的保证和最佳的性能。  

**表与事件的二象性（duality）**  
让我们继续聊一下数据库。变更日志 和 表之间有着迷人的二象性。 日志类似借贷清单和银行处理流水，而数据库表则是当前账户的余额。如果有变更日志，你就可以应用这些变更生成数据表并得到当前状态。 表记录的是每条数据的最后状态（日志的一个特定时间点）。可以认识到日志是更基本的数据结构：日志除了可用来创建原表，也可以用来创建各类衍生表。（是的，表可以是非关系型用户用的键值数据存储（keyed data store））  
这个过程也是可逆的：如果你对一张表进行更新，你可以记录这些变更，并把所有更新的『变更日志（changelog）』发布到表的状态信息中。 这些变更日志正是你所需要的支持准实时的复制。 基于此，你就可以清楚的理解表与事件的二象性： 表支持了静态数据，而日志记录了变更。日志的魅力就在于它是变更的完整记录，它不仅仅包含了表的最终版本的内容， 而且可以用于重建任何存在过其它版本。事实上，日志可以看作是表每个历史状态的一系列备份。  

这可能会让你想到源代码的版本控制（source code version control）。源码控制和数据库之间有着密切的关系。 版本管理解决了一个和分布式数据系统要解决的很类似的问题 —— 管理分布式的并发的状态变更。 版本管理系统建模的是补丁序列（the sequence of patches），实际上这就是日志。 你可以检出当前的代码的一个『快照』并直接操作，这个代码快照可以类比成表。 你会注意到，正如有状态的分布式系统一样，版本控制系统通过日志来完成复制：更新代码即是拉下补丁并应用到你的当前快照中。

从销售日志数据库的公司 Datomic那里，大家可以看到一些这样的想法。 这个视频比较好地介绍了他们如何在系统中应用这些想法。 当然这些想法不是只专属于这个系统，这十多年他们贡献了很多分布式系统和数据库方面的文献。

### 第二部分：数据集成  
我先解释一下我说的是『数据集成』（data integration）是什么，还有为什么我觉得它很重要，然后我们再来看看它是如何和日志建立关系的。  
数据集成是指使一个组织的所有数据对这个组织的所有的服务和系统可用。  
数据集成』还不是一个常见的用语，但是我找不到一个更好的。大家更熟知的术语 ETL （译注：ETL 是 Extraction-Transformation-Loading 的缩写，即数据提取、转换和加载） 通常只是覆盖了数据集成的一个有限子集 —— 主要在关系型数据仓库的场景。但我描述的东西很大程度上可以理解为，将 ETL 推广至覆盖实时系统和处理流程。  
你一定不会听到数据集成就兴趣盎然地屏住呼吸，并且天花乱坠的想到大数据的概念，但尽管如此，我相信这个陈词滥调的『让数据可用』的问题是组织可以关注的更有价值的事情之一。  
对数据的高效使用遵循一种马斯洛的需要层次理论。 金字塔的基础部分包含捕获所有相关数据，能够将它们全部放到适当的处理环境中（可以是一个华丽的实时查询系统，或仅仅是文本文件和 python 脚本构成的环境）。 这些数据需要以统一的方式建模，以方便读取和处理。 一旦这些以统一的方式捕获数据的基本需求得到满足，那么在基础设施上以不同方法处理这些数据就变得理所当然 —— MapReduce、实时查询系统等等。  
显而易见但值得注意的一点：如果没有可靠的、完整的数据流，Hadoop 集群只不过是个非常昂贵而且安装麻烦的供暖器。 一旦有了数据和处理（data and processing），人们的关注点就会转移到良好的数据模型和一致且易于理解的语义这些更精致的问题上来。 最后，关注点会转移到更高级处理上 —— 更好的可视化、生成报表以及处理和预测算法。  
以我的经验，大多数组织在这个数据金字塔的底部存在巨大的漏洞 —— 缺乏可靠的完整的数据流 —— 却想直接跳到高级数据模型技术上。这样做完全是本未倒置。  
所以问题是，我们如何在组织中构建贯穿所有数据系统的可靠数据流？  

**数据集成：两个难题之事件数据管道**
第一个趋势是增长的事件数据（event data）。事件数据记录的是发生的事情，而不是已存在的事情。在 Web 系统中，这就意味着用户活动日志，还有为了可靠地操作和监控数据中心机器的价值而记录的机器级别的事件和统计数字。人们倾向于称它们为『日志数据（log data）』，因为它们经常被写到应用日志中，但这样的说法混淆了形式与功能。这些数据是现代 Web 的核心：归根结底，Google 的财富来自于建立在点击和展示（clicks and impressions）上的相关性管道（relevance pipeline），而这些点击和展示正是事件。  
这样的事情并不是仅限于 Web 公司，只是 Web 公司已经完全数字化，所以更容易完成。财务数据长久以来一直是以事件为中心的。 RFID（无线射频识别）使得能对物理设备做这样的跟踪。 随着传统的业务和活动的数字化（digitization）， 我认为这个趋势仍将继续。  
这种类型的事件数据记录了发生的事情，往往比传统数据库应用要大好几个数量级。这对于处理提出了重大的挑战。  

**数据集成：两个难题之专用的数据系统（specialized data systems）的爆发**  
第二个趋势来自于专用的数据系统的爆发，这些数据系统在最近五年开始流行并且可以免费获得。 专门用于 OLAP、搜索、简单在线存储、 批处理、图分析（graph analysis）等等 的数据系统已经出现。  
更加多样化的数据同时变成更加大量，而且这些数据期望放到更多的系统中，这些需求同时要解决，导致了一个巨大的数据集成问题。  

**日志结构化的（log-structured）数据流**  
处理系统之间的数据流，日志是最自然的数据结构。解决方法很简单：提取所有组织的数据，并放到一个用于实时订阅的中心日志中。  
每个逻辑数据源都可以建模为它自己的日志。 一个数据源可以看作 一个输出事件日志的应用（如点击或页面的浏览），或是 一个接受修改的数据库表。 每个订阅消息的系统都尽可能快的从日志读取信息，将每条新的记录应用到自己的存储中，同时向前滚动日志文件中的自己的位置。 订阅方可以是任意一种数据系统 —— 缓存、Hadoop、另一个网站中的另一个数据库、一个搜索系统，等等。  
举个例子，日志概念为每个变更提供了逻辑时钟，所有的订阅方都可以比较这个逻辑时钟。 这极大简化了如何去推断不同的订阅系统的状态彼此是否一致的，因为每个系统都持有了一个读到哪儿的『时间点』。  
为了让讨论更具体些，我们考虑一个简单的案例，有一个数据库和一组缓存服务器集群。 日志提供了一个方法可以同步更新到所有这些系统，并推断出每个系统的所处在的时间点。 我们假设做了一个写操作，对应日志记录 X，然后要从缓存做一次读操作。 如果我们想保证看到的不是过时的数据，我们只需保证，不要去读取那些复制操作还没有跟上 X 的缓存即可。  
日志也起到缓冲的作用，使数据的生产异步于数据的消费。有许多原因使得这一点很重要，特别是在多个订阅方消费数据的速度各不相同的时候。 这意味着一个数据订阅系统可以宕机或是下线维护，在重新上线后再赶上来：订阅方可以按照自己的节奏来消费数据。 批处理系统如 Hadoop 或者是一个数据仓库，或许只能每小时或者每天消费一次数据，而实时查询系统可能需要及时到秒。 无论是起始的数据源还是日志都感知感知各种各样的目标数据系统，所以消费方系统的添加和删除无需去改变传输管道。  
特别重要的是：目标系统只知道日志，而不知道来源系统的任何细节。 无论是数据来自于一个 RDBMS、一种新型的键值存储，还是由一个不包含任何类型实时查询的系统所生成的，消费方系统都无需关心。 这似乎是一个小问题，但实际上却是至关重要的。  

这里我使用术语『日志』取代了『消息系统』或者『发布 - 订阅』，因为在语义上明确得多，并且准确得多地描述了在实际实现中支持数据复制时你所要做的事。 我发现『发布订阅』只是表达出了消息的间接寻址（indirect addressing of messages） —— 如果你去比较两个发布 - 订阅的消息系统的话，会发现他们承诺的是完全不同的东西，而且大多数模型在这一领域没什么用。 你可以认为日志是一种有持久性保证和强有序（strong ordering）语义的消息系统。 在分布式系统中，这个通信模型有时有个（有些可怕的）名字叫做原子广播（atomic broadcast）。  
值得强调的是，日志仍然只是基础设施，并不是精通数据流这个故事的结束： 故事的剩余部分围绕着元数据（metadata)、schemas、兼容性以及处理数据结构及其演化的所有细节来展开。 但是，除非有一种可靠的通用的方式来处理数据流的机制，否则语义细节总是次要的。  

**在 LinkedIn**  
随着 LinkedIn 从集中式关系数据库过渡到一套分布式系统，我注意到数据集成的问题在迅速地演变。  
目前我们主要的数据系统包括：  

- 搜索
- Social Graph
- Voldemort （键值存储）
- Espresso （文档存储）
- 推荐引擎
- OLAP 查询引擎
- Hadoop
- Terradata
- Ingraphs （监控图表和指标服务）

每一个都是专用的分布式系统，在各自的专门领域提供高级的功能。  
使用日志作为数据流的这个想法，甚至在我到这里之前，就已经在 LinkedIn 的各个地方开始浮现了。 我们开发的最早的一个基础设施是一个称为 databus的服务， 它在我们早期的 Oracle 表上提供了一种日志缓存的抽象，用于可伸缩地订阅数据库修改，给我们的 social graph 和搜索索引输入数据。  
我先简单介绍一些历史以提供讨论的上下文。在发布我们自己键值存储之后，大约是 2008 年我开始参与这个项目。 我接着的一个项目是把一个运行的 Hadoop 部署用起来，迁移我们的一些推荐处理上来。 由于缺乏这方面的经验，我们只计划了几周时间完成数据的导入导出，剩下的时间则用来实现复杂的预测算法。 就这样我们开始了长途跋涉。  
我们本来计划是仅仅将数据从现存的 Oracle 数据仓库中剖离。 但是我们首先发现将数据从 Oracle 中迅速取出简直是一个黑魔法（dark art）。 更糟的是，数据仓库的处理过程并不适合于 我们为 Hadoop 设计的生产批处理过程 —— 大部分处理都是不可逆的，并且与要生成的具体报表相关。 最终我们采取的办法是，避免使用数据仓库，直接访问源数据库和日志文件。 最后，我们实现了一个管道，用于完成加载数据到我们的键值存储并生成结果。  
这种普通常见的数据拷贝最终成为原来开发项目的主要内容之一。 糟糕的是，只要在任何时间任意管道有一个问题，Hadoop 系统基本上就是废的 —— 在错误的数据基础上运行复杂的算法只会产生更多的错误数据。  
虽然我们已经使用了一种很通用的构建方式，但是每个数据源都需要自定义的安装配置。这也被证明是大量错误与失败的根源。 我们用 Hadoop 实现的网站功能开始流行起来，而我们发现自己有一大把需要协作的工程师。 每个用户都有他们想要集成的一大把的系统，并且想要导入的一大把新数据源。  
有些东西在我面前开始渐渐清晰起来。  
首先，我们已建成的通道虽然有一些杂乱，但实际上是极有价值的。 仅在一个新的处理系统（Hadoop）中让数据可用于处理 就开启了大量的可能性。 基于这些数据过去很难实现的计算如今已变为可能。 许多新的产品和分析技术都来源于把多个数据片块放在一起，这些数据过去被锁定在特定的系统中。  
第二，可靠的数据加载需要数据通道的深度支持，这点已经变得很清晰了。 如果我们可以捕获所有我们需要的结构，就可以使得 Hadoop 数据全自动地加载， 这样不需要额外的手动操作就可以添加新的数据源或者处理 schema 变更 —— 数据就会自动的出现在 HDFS，并且 Hive 表就会自动的为新数据源生成恰当的列。  
第三，我们的数据覆盖率仍然很低。 如果看一下 LinkedIn 所有数据在 Hadoop 中可用的比率，仍然很不完整。 相比接入并运转一个新数据源所要做的努力，完整接入一个数据源更不容易。  
我们曾经推行的方式是为每个数据源和目标构建自定义的数据加载，很显然这是不可行的。 我们有几十个数据系统和数据仓库。把这些系统和仓库联系起来，就会导致任意两两系统间构建自定义的管道，如下所示：  
![datapipeline-complex](https://content.linkedin.com/content/dam/engineering/en-us/blog/migrated/datapipeline_complex.png)  

我们需要尽可能的将每个消费者与数据源隔离。理想情形下，它们应该只与一个单独的数据源集成，就能访问到所有数据。  
这个思想是增加一个新的数据系统 —— 它可以作为数据来源或者数据目的地 —— 集成工作只需连接这个新系统到一个单独的管道，而无需连接到每个数据消费方。  
这样的经历使得我专注于创建 Kafka，把 我们所知的消息系统的特点 与 在数据库和分布式系统内核常用的日志概念 结合起来。 我们首先需要一个实体作为所有的活动数据的中心管道，并逐步的扩展到其他很多的使用方式，包括 Hadoop 之外的数据、数据监控等等。  
在相当长的时间内，Kafka 是独一无二的（有人会说是怪异） —— 作为一个底层设施，它既不是数据库，也不是日志文件收集系统，更不是传统的消息系统。 但是最近 Amazon 提供了非常非常类似 Kafka 的服务，称之为 Kinesis。 相似度包括了分片（partition）处理的方式，数据的持有方式，甚至包括有点特别的 Kafka API 分类（分成高端和低端消费者）。 我很开心看到这些，这表明了你已经创建了很好的底层设施抽象，AWS 已经把它作为服务提供！ 他们对此的想法看起来与我所描述的完全吻合： 管道联通了所有的分布式系统，诸如 DynamoDB,RedShift,S3 等，同时作为使用 EC2 进行分布式流处理的基础。  

**ETL 与数据仓库的关系**  
我们再来聊聊数据仓库。数据仓库旨在包含支撑数据分析的规整的集成的结构化数据。 这是一个非常好的理念。对不了解数据仓库概念的人来说，数据仓库的用法是： 周期性的从源数据库抽取数据，把它们转化为可理解的形式，然后把它导入中心数据仓库。 对于数据集中分析和处理，拥有高度集中的位置存放全部数据的规整副本对于数据密集的分析和处理是非常宝贵的资产。 在更高层面上，无论你使用传统的数据仓库 Oracle 还是 Teradata 或 Hadoop， 这个做法不会有太多变化，可能调整一下抽取和加载数据的顺序。  
数据仓库是极其重要的资产，它包含了的和规整的数据，但是实现此目标的机制有点过时了。  
对于以数据为中心的组织，关键问题是把规整的集成的数据联结到数据仓库。 数据仓库是个批处理查询基础设施：它们适用于各类报表和临时性分析，特别是当查询包含了简单的计数、聚合和过滤。 但是如果批处理系统是唯一一个包含规整的完整的数据的仓库， 这就意味着，如果一个系统需要 实时数据输入的实时系统（如实时处理、实时搜索索引、实时监控等系统），这些数据是不可用的。

依我之见，ETL 包括两件事。 首先，它是数据抽取和清理的处理 —— 本质上就是释放被锁在组织的各类系统中的数据，去除特定于系统的约束。 第二，依照数据仓库的查询重构数据，例如使其符合关系数据库类型系统， 强制使用星型 schema（star schema）、雪花型 schema（snowflake schema），可能会打散数据成高性能的列格式（column format），等等。同时做好这两件事是有困难的。 这些集成仓库的规整的数据除了要索引到实时存储系统中，也应当可用于实时或是低时延处理中。

在我看来，正是因为这个原因有了额外好处：使得数据仓库 ETL 大大提升了 *** 组织级 *** 的可伸缩性（scalable）。 典型的问题是数据仓库团队要负责收集和整理组织中各个团队所生成的全部数据。 两边的收益是不对称的：数据的生产者常常并不知晓在数据仓库中数据的使用情况， 结果产生的数据，为了转成为可用的形式，抽取过程很难或是很繁重，转换过程很难统一规模化。 当然，中心团队的规模不可能跟上组织中其它团队增长， 结果数据的覆盖率总是参差不齐的，数据流是脆弱的，跟进变更是缓慢的。

较好的做法是有一个中央管道即日志，用定义良好的 API 来添加数据。 集成这个管道和提供良好的结构化的输入数据所需的职责由提供数据的生产者承担。 这意味着作为系统设计和实现一部分的生产者，在交付到中心通道时， 必须考虑其输出和输入的数据要有良好结构形式的问题。 新的存储系统的加入对于数据仓库团队是无关紧要的，因为他们现在只有一个中心结点去集成。 （译注：原来要集成的是其它各个相关的系统，工作是被简化了的） 数据仓库团队需只处理更简单的问题，从中心日志中加载结构化的输入数据、完成特定于他们系统的数据转换。
![pipeline-ownership](https://content.linkedin.com/content/dam/engineering/en-us/blog/migrated/pipeline_ownership.png)  

从上面讨论可以看出，当考虑采纳传统的数据仓库之外额外的数据系统时，组织级的伸缩性（organizational scalability）显得尤为重要。 例如，想为组织的所有的数据集提供搜索能力。 或者想为数据流的监控的次级监控（sub-second monitoring）添加实时数据趋势和告警。 无论是哪个情况，传统的数据仓库的基础设施，甚至是 Hadoop 集群都将不再适合。 更糟的是，用于支持数据加载的 ETL 处理管道可能输入不了数据到其它系统， 和带动不了要动用数据仓库这样的大企业下的那些基础设备。 这样的做法应该是不可行的，可能可以解释为什么多数组织对他们的所有数据很难轻松具备这样的能力。 反之，如果组织能导出标准的结构良好的数据， 那么任何新的系统要使用所有数据仅仅需要提供一个用于集成的管道接到中央管道上即可。

关于数据规整化和转换在哪里进行，这种架构也引出了的不同观点：  
> 在添加数据到公司全局日志之前，由数据的生产者完成。  
> 由在日志上的一个实时转换器完成，转换器生成一个新的转换过的日志。  
> 作为加载过程的一部分，由目标系统完成。  

最好的模型是数据发布到日志之前由数据生产者完成数据规整化。 这样可以确保数据是处于规范形式（canonical form）的， 并且不需要保留数据 从原来生产系统的特定代码或是原来存储系统的维护方式所带来的任何遗留属性。 这些细节最好由产成数据的团队来处理，因为他们最了解他们自己的数据。 这个阶段所使用的任何逻辑都应该是无损的和可逆的。

可以实时完成的任何类型有附加值的转换操作都应该作为原始日志数据的后处理环节完成。 这类操作包括了事件数据的会话管理，或者附加上大家感兴趣的派生字段。 原始日志仍然是可用的，但这样的实时处理生产了包含增强数据（augmented data）的派生日志。

最后，只有针对目标系统的聚合操作才应该加到加载过程中。 比如可能包括在数据仓库中为分析和报表而做的把数据转化成特定的星型或者雪花状 schema。 因为在这个阶段（一般比较自然地对应到传统的 ETL 处理阶段），现在处理的是一组规整得多和统一得多的流， 处理过程已经大大简化了。  

**日志文件与事件**  
我们再来聊聊这种架构的附带的优势：支持解耦的事件驱动的系统。

在 Web 行业取得活动数据的典型方法是把打日志到文本文件中， 然后这些文本文件分解进入数据仓库或者 Hadoop 用于聚合和查询。 这做的问题和所有批处理的 ETL 做法一样：数据流耦合了数据仓库系统的能力和处理计划（processing schedule）。

在 LinkedIn，是以在中心日志完成处理的方式构建事件数据。 Kafka 做为中心的有多个订阅方的事件日志，定义数百种事件类型， 每种类型都会捕获一个特定动作类型的独特属性。 这样的方式覆盖从页面浏览、广告展示、搜索到服务调用、应用异常的方方面面。

为了进一步理解这一优势，设想一个简单的场景 —— 显示在工作职位页面提交的职位信息。 职位页面应当只包括显示职位所需的逻辑。 然而，在足够动态站点中，这很容易就变成与职位显示无关的额外逻辑的点缀。 例如，我们将对如下的系统进行集成：  
> 发送数据到 Hadoop 和数据仓库中，以做离线数据处理    
> 浏览计数，确保查看者不是一个内容爬虫  
> 聚合浏览信息，在职位提交者的分析页面显示  
> 记录浏览信息，确保合适地设置了用户的推荐职位的展示上限（不想重复地展示同样内容给用户）  
> 推荐系统可能需要记录浏览，以正确的跟踪职位的流行程度  

用不了多久，简单的职位显示变得相当的复杂。 与此同时，还要增加职位显示的其它终端 —— 移动终端应用等等 —— 这样的逻辑必须继续实现，复杂程度被不断地提升。 更糟的是，我们需要交互的系统是多方需求交织缠绕在一起的 —— 负责显示职位的工程师需要知道多个其它系统和功能，才可以确保集成的正确。 这里仅是简单描述了问题，实际应用系统只会更加复杂。

『事件驱动』风格提供了简化这类问题的方案。 职位显示页面现在只负责显示职位并记录显示职位的信息，如职位相关属性、页面浏览者及其它有价值的信息。 其它所有关心这个信息的系统诸如推荐系统、安全系统、职位提交分析系统和数据仓库，只需订阅上面的输出数据进行各自的处理。 显示代码并不需要关注其它的系统，也不需要因为增加了数据的消费者而做改变。  

**构建可伸缩的日志**  
当然，把发布者与订阅者分离不再是什么新鲜事了。 但是如果要给一个需要按用户扩展的（consumer-scale）网站提供多个订阅者的实时提交日志， 那么可伸缩性就会成为你所面临的首要挑战。 如果我们不能创建快速、低成本和可伸缩的日志以满足实际大规模的使用，把日志用作统一集成机制只不过是个美好的幻想。  
人们普遍认为分布式日志是缓慢的、重量级的抽象（并且通常只把它与『元数据』类的使用方式联系在一起，可能用 Zookeeper 才合适）。但有了一个专注于大数据流的深思熟虑的实现可以打破上面的想法。 在 LinkedIn，目前每天通过 Kafka 写入超过 600 亿条不同的消息。 （如果算上数据中心之间镜像的消息，那么这个数字会是数千亿。）

为了支持这样的规模，我们在 Kafka 中使用了一些技巧：
> 日志分片  
> 通过批量读出和写入来优化吞吐量  
> 规避无用的数据拷贝

为了确保水平可扩展性，我们把日志进行切片：  
![partitioned-log](https://content.linkedin.com/content/dam/engineering/en-us/blog/migrated/partitioned_log.png)  

每个分片的日志是有序的，但是分片之间没有全局的次序（这个有别于在你的消息中可能包含的挂钟时间）。 由写入者决定消息发送到特定的日志分片，大部分使用者以某种键值（如用户 id）来进行分片。 追加日志时，分片方式在片段之间可以不需要协调，并且可以使系统的吞吐量与 Kafka 集群大小线性增长。

每个分片通过可配置数字指定数据复制的复本个数，每个复本都有一个分片日志完全一致的一份拷贝。 任何时候都有一个复本作为 leader，如果 leader 出错了，会有一个复本接替成为 leader。

缺少跨分片的全局顺序是个局限，但是我们没有发现它成为大问题。 事实上，与日志的交互一般来源于成百上千个不同的处理流程，所以为所有处理提供全局顺序没有意义。 转而需要的是，我们提供的每个分片有序的保证，和 Kafka 提供的 同一发送者发送给同一分区的消息以相同的顺序交付到接收者 的保证。

日志，和文件系统一样，对于顺序读写可以方便地优化。日志可以把小的读写合成大的高吞吐量的操作。 Kafka 非常积极做这方面的优化。客户端向服务器端的数据发送、磁盘写入、服务器之间复制、到消费者数据传递和数据提交确认 都会做批处理。

最后，Kafka 使用简单的二进制格式维护内存日志、磁盘日志和传送网络数据。这使得我们可以使用包括『0 拷贝的数据传输』在内的大量的优化机制。

这些优化的积累效应是往往以磁盘和网络的速度在读写数据，即使维护的数据集大大超出内存大小。

这些自卖自夸的介绍不意味着是关于 Kafka 的主要内容，我就不再深入细节了。 LinkedIn 方案的更细节说明在这儿，Kafka 设计的详细说明在这儿，你可以读一下。

### 第三部分：日志与实时流处理
到目前为止，我只讲述了系统之间拷贝数据的理想机制。但是在存储系统之间搬运字节不是所要讲述内容的全部。 最终会发现，『日志』是流的另一种说法， 并且日志是流处理的核心。

但是，等会儿，流处理到底是什么呢？

如果你是上世纪 90 年代晚期或者 21 世纪初数据库文化或者成功了一半的数据基础设施产品的爱好者，那么你就可能会把流处理与建创 SQL 引擎或者『箱子和箭头』（boxes and arrows）接口用于事件驱动的处理联系起来。

如果你关注大量出现的开源数据库系统，你就可能把流处理和一些这领域的系统关联起来， 比如 Storm、Akka、S4和 Samza。 但是大部分人会把这些系统看为异步消息处理系统，与支持群集的远程过程调用（RPC）层没什么差别 （而事实上这一领域一些系统确实是如此）。

这些观点都有一些局限性。流处理即与 SQL 无关，也不局限于实时流处理。 还没有根本的原因，限制你不能使用多种不同的语言来表达计算，处理昨天的或者一个月之前的流数据。

我把流处理视为更广泛的概念：持续数据流处理的基础设施。 我认为计算模型可以像 MapReduce 或者分布式处理框架一样通用，但是有能力生成低时延的结果。

处理模型的真正驱动力是数据收集方法。成批收集数据自然是批处理。当数据是持续收集的，自然也应该持续不断地处理。

美国的统计调查是一个成批收集数据的经典例子。 统计调查周期性的开展，用的是蛮力调查，通过挨门挨户的走访统计美国公民的信息。 在 1790 年统计调查刚刚开始，这样做是很合理的。 那时的数据收集本质就是面向批处理的，包括了骑马到周边人家，用纸笔记录，然后把成批的记录运输到人们统计数据的中心站点。 现在，在描述这个统计过程时，人们立即会想到为什么我们不保留出生和死亡的记录，这样就可以算出人口统计信息，这些信息或是持续即时计算出来或者按需要时间隔计算。

这是一个极端的例子，但是现在大量的数据传输处理仍然依赖于周期性的转录和批量的传输和集成。 处理批量转录数据的唯一方法就是批量的处理。 但是随着这些批处理被持续的数据输入所取代，人们自然而然的开始向持续处理转变，以平滑地使用所需的处理资源并且减少延迟。

例如在 LinkedIn 几乎完全没有批量数据收集。我们大部分的数据要么是活动数据或者要么是数据库变更，两者都是不间断地发生的。 事实上，你想到的任何商业业务，底层的机制几乎都是不间断的处理，正如 Jack Bauer 所说的，事件的发生是实时的。 当数据以成批的方式收集，几乎总是由这些原因所致：有一些人为的步骤；缺少数字化；或是非数字化流程的历史古董不能自动化。 当使用邮件或者人工方式，传输和处理数据是非常缓慢的。刚开始转成自动化时，总是保持着原来流程的形式，所以这样的情况会持续相当长的时间。

每天运行的『批量』处理作业常常在模拟一种窗口大小是一天的持续计算。 当然，底层的数据其实总是在变化着的。 在 LinkedIn，这样的做法如此之常见（并且在 Hadoop 做到这些的实现机制如此之复杂）， 以至于我们实现了一整套框架来管理增量的 Hadoop 工作流。

由此看来，对于流处理我很容易得出不同观点： 它处理的是包含时间概念的底层数据并且不需要静态的数据快照， 所以可以以用户可控频率生产输出而不是等待数据集的『都』到达后再生产输出（译注：数据是会持续的，所以实际上不会有『都』达到的时间点）。 从这个角度上讲，流处理是广义上的批处理，随着实时数据的流行，流处理会是很重要处理方式。

那么，为什么流处理的传统观点大家之前会认为更合适呢？ 我个人认为最大的原因是缺少实时数据收集，使得持续处理之前是学术性的概念。

我觉得，是否缺少实时数据的收集决定了商用流处理系统的命运。 当他们的客户还是用面向文件的每日批量处理完成 ETL 和数据集成时， 建设流处理系统的公司专注于提供处理引擎来连接实时数据流，而结果是当时几乎没有人真地有实时数据流。 其实我在 LinkedIn 工作的初期，有一家公司想把一个非常棒的流处理系统卖给我们， 但是因为当时我们的所有数据都按小时收集在的文件里， 所以用上这个系统我们能做到的最好效果就是在每小时的最后把这些文件输入到流处理系统中。 他们意识到这是个普遍问题。 下面的这个异常案例实际上是证明上面规律： 流处理获得一些成功的一个领域 —— 金融领域，这个领域在过去，实时数据流就已经标准化，并且流处理已经成为了瓶颈。

甚至于在一个健康的批处理的生态中，我认为作为一种基础设施风格，流处理的实际应用能力是相当广阔的。 我认为它填补了实时数据请求 / 响应服务和离线批量处理之间的缺口。现代的互联网公司，我觉得大约 25% 的代码可以划分到这个情况。

事实证明，日志解决了流处理中最关键的一些技术问题，后面我会进一步讲述， 但解决的最大的问题是日志使得多个订阅者可以获得实时的数据输入。 对技术细节感兴趣的朋友，我们已经开源了 Samza， 它正是基于这些理念建设的一个流处理系统。 很多这方面的应用的更多技术细节我们在此文档中有详细的描述。

**数据流图（data flow graphs）**  
![dag](https://content.linkedin.com/content/dam/engineering/en-us/blog/migrated/dag.png)  

流处理最有趣的特点是它与流处理系统的内部组织无关， 但是与之密切相关的是，流处理是怎么扩展了之前在数据集成讨论中提到的认识：输入数据（data feed）是什么。 我们主要讨论了原始数据（primary data）的 feeds 或说 日志 —— 各种系统执行所产生的事件和数据行。 但是流处理允许我们包括了由其它 feeds 计算出的 feeds。 在消费者看来，这些派生的 feeds 和 用于生成他们的原始数据的 feeds 看下来没什么差别。 这些派生的 feeds 可以按任意的复杂方式封装组合。

让我们再深入一点这个问题。 对于我们的目标，流处理作业是指从日志读取数据和将输出写入到日志或其它系统的任何系统。 用于输入和输出的日志把这些处理系统连接成一个处理阶段的图。 事实上，以这样的风格使用中心化的日志，你可以把组织全部的数据抓取、转化和工作流仅仅看成是一系列的写入它们的日志和处理过程。

流处理器根本不需要高大上的框架： 可以是读写日志的一个处理或者一组处理过程，但是为了便于管理处理所用的代码，可以提供一些额外的基础设施和支持。

在集成中日志的目标是双重的：

首先，日志让各个数据集可以有多个订阅者并使之有序。 让我们回顾一下『状态复制』原理来记住顺序的重要性。 为了更具体地说明，设想一下从数据库中更新数据流 —— 如果在处理过程中把对同一记录的两次更新重新排序，可能会产生错误的输出。 这里的有序的持久性要强于 TCP 之类所提供的有序，因为不局限于单一的点对点链接，并且在流程处理失败和重连时仍然要保持有序。

其次，日志提供了处理流程的缓冲。 这是非常基础重要的。如果多个处理之间是非同步的，那么生成上行流数据的作业生成数据可能比另一个下行流数据作业所能消费的更快。 这种情况下，要么使处理进程阻塞，要么引入缓冲区，要么丢弃数据。 丢弃数据似乎不是个好的选择，而阻塞处理进程，会使得整个的数据流的图被迫中止处理。 日志是一个非常非常大的缓冲，允许处理进程的重启或是失败，而不影响流处理图中的其它部分的处理速度。 要扩展数据流到一个更庞大的组织，这种隔离性极其重要，整个处理是由组织中不同的团队提供的处理作业完成的。 不能因为某个作业发生错误导致影响前面作业，结果整个处理流程都被卡住。

Storm和 Sama都是按这种风格构建，能用 kafka 或其它类似的系统作为它们的日志。  

**有状态的实时流处理**  
一些实时流处理做的只是无状态的单次记录的转换，但有很多使用方式需要在流处理的某个大小的时间窗口内进行更复杂的计数、聚合和关联（join）操作。 比如，给一个的事件流（如用户点击的流）附加上做点击操作用户的信息， —— 实际上即是关联点击流到用户的账户数据库。 这类流程最终总是要处理者维护一些状态信息： 比如在计算一个计数时，需要维护到目前为止的计数器。 在处理者可能挂掉的情况下，如何维护正确的状态？

最简单的方案是把状态保存在内存中。但是如果处理流程崩溃，会丢失中间状态。 如果状态是按窗口维护的，处理流程只能会回退到日志中窗口开始的时间点上。 但是，如果计数的时间窗口是 1 个小时这么长，那么这种方式可能不可行。

另一个方案是简单地存储所有的状态到远程的存储系统，通过网络与这些存储关联起来。 但问题是没了数据的局部性并产生很多的网络间数据往返（network round-trip）。

如何才能即支持像处理流程一样分片又支持像『表』一样的存储呢？

回顾一下关于表和日志二象性的讨论。它正好提供了把流转换成与这里我们处理中所需的表的工具，同时也是一个解决表的容错的处理机制。

流处理器可以把它的状态保存在本地的『表』或『索引』中 —— bdb、leveldb 甚至是些更不常见的组件，如 Lucene或 fastbit索引。 这样一些存储的内容可以从它的输入流生成（可能做过了各种转换后的输入流）。 通过记录关于本地索引的变更日志，在发生崩溃、重启时也可以恢复它的状态。 这是个通用的机制，用于保持 任意索引类型的分片之间相互协作（co-partitioned）的本地状态 与 输入流数据 一致。

当处理流程失败时，可以从变更日志中恢复它的索引。 每次备份时，即是日志把本地状态转化成一种增量记录。

这种状态管理方案的优雅之处在于处理器的状态也是做为日志来维护。 我们可以把这个日志看成是数据库表变更的日志。 事实上，这些处理器本身就很像是自维护的分片之间相互协作的表。 因为这些状态本身就是日志，所以其它处理器可以订阅它。 如果处理流程的目标是更新结点的最后状态并且这个状态又是流程的一个自然的输出，那么这种方式就显得尤为重要。

再组合使用上用于解决数据集成的数据库输出日志，日志和表的二象性的威力就更加明显了。 从数据库中抽取出来的变更日志可以按不同的形式索引到各种流处理器中，以关联到事件流上。

在 Samza 和这些大量实际例子中， 我们说明了这种风格的有状态流处理管理的更多细节。

**日志合并（log compaction）**  
![log-compaction](https://content.linkedin.com/content/dam/engineering/en-us/blog/migrated/log_compaction_0.png)  

当然，我们不能奢望一直保存着全部变更的完整日志。 除非想要使用无限空间，日志总是要清理。 为了让讨论更具体些，我会介绍一些 Kafka 这方面的实现。 在 Kafka 中，清理有两种方式，取决于数据包括的是键值存储的更新还是事件数据。 对于事件数据，Kafka 支持仅维护一个窗口的数据。通常，窗口配置成几天，但窗口也可以按空间大小来定。 对于键值存储的更新，尽管完整日志的一个优点是可以回放以重建源系统的状态（一般是另一个系统中重建）。

但是，随着时间的推移，保持完整的日志会使用越来越多的空间，并且回放的耗时也会越来越长。 因此在 Kafka 中，我们支持不同类型的保留方式。 我们删除过时的记录（如这些记录的主键最近更新过）而不是简单的丢弃旧日志。 这样做我们仍然保证日志包含了源系统的完整备份，但是现在我们不再重现原系统曾经的所有状态，仅是最近的哪些状态。 这一功能我们称之为日志合并。

### 第四部分：系统构建
最后我要讨论的是在线数据系统设计中日志的角色。

日志服务在分布式数据库中服务于数据流 可以类比 日志服务在大型组织机构中服务于数据集成。 在这两个应用场景中，日志要解决的问题 都是 数据流、一致性和可恢复性。   如果组织不是一个很复杂的分布式数据系统呢，它究竟是什么？  

**分解单品方式而不是打包套餐方式（Unbundling）？**  
如果你换个角度，可以把组织的系统和数据流整体看做整个一个分布式数据： 把所有独立的面向查询的系统（如 Redis、SOLR、Hive 表，等等）看做只是你的数据的特定的索引； 把流处理系统（如 Storm、Samza）看做只是一种很成熟的触发器和视图的具体机制。 我注意到，传统的数据库人员非常喜欢这样的观点，因为他们终于能解释通，这些不同的数据系统到底是做什么用的 —— 它们只是不同的索引类型而已！

不可否认这段时间涌现了大量类型的数据系统，但实际上，这方面的复杂性早就存在。 即使是在关系数据库的鼎盛时期，组织里就有种类繁多的关系数据库！ 因为大型机，所有的数据都存储在相同的位置，所以可能并没有真正的数据集成。 有很多推动力要把数据分离到多个系统：数据伸缩性、地理地域、安全性和性能隔离是最常见的。 这些问题可以通过一个好的系统来解决： 比如组织使用单个 Hadoop 保存所有数据来服务大量各式各样的客户，这样做是可能的。

所以处理的数据向分布式系统变迁的过程中，已经有了个可能的简单方法： 把大量的不同系统的小实例合到少数的大集群中。

许多的系统还不足好到支持这个方法：它们没有安全，或者不能保证性能隔离性，或者伸缩性不够好。 不过这些问题都是可以解决的。

依我之见，不同系统大量涌现的原因是构建分布式数据系统的难度。 把关注点消减到单个查询类型或是用例，各个系统可以把关注范围控制到一组能构建出来的东西上。 但是把全部这些系统运行起来，这件事有非常多的复杂性。

我觉得解决这类问题将来有三个可能的方向：

第一种可能性是延续现状：各个分离的系统在往后很长的一段时间里基本保持不变。 发生这种可能要么是因为建设分布式系统的困难很难克服， 要么系统的专用化（specialization）能让各个系统的便得性（convenience）和能力（power）达到一个新的高度。 只要现状不变，为了能够使用数据，数据集成问题将仍会最核心事情之一。 如果是这样，用于集成数据的外部日志将会非常的重要。

第二种可能性是一个统一合并的系统（re-consolidation），这个系统具备足够的通用性，逐步把所有不同的功能合并成单个超极系统。 这个超级系统表面看起来类似关系数据库，但在组织中使用方式会非常不一样，因为只能用一个大系统而不是无数个小系统。 在这样的世界里，除了系统自身的内部，不存在真正的数据集成问题。 我觉得，因为建设这样的系统的实际困难，使这个情况不太可能发生。

还有另一种可能的结果，呃，其实我觉得这个结果对工程师很有吸引力。 新一代数据系统的一个让人感兴趣的特征是，这个系统几乎是完全开源的。 开源提供了另一个可能性：数据基础架构不用是打包套餐式的而是分解单品成一组服务及面向应用的 API。 在 Java 栈中，你可以看到这种状况在一定程度上已经发生了：  

- Zookeeper处理系统之间的协调的很多问题。 （或许诸如 Helix 或 Curator等高级别抽象可以有些帮助）。
- Mesos和 YARN 处理虚拟化（virtualization）和资源管理。
- Lucene和 LevelDB等嵌入式类库做为索引。
- Netty、Jetty 和 更高层封装如 Finagle、rest.li处理远程通信。
- Avro、Protocol Buffers、Thrift和 umpteen zillion等其它类库处理序列化。
- Kafka 和 Bookeeper提供后端支持的日志。

如果你把上面这些叠成一堆，换个角度去看，它会有点像是乐高版（lego version）的分布式数据系统工程。 你可以把这些零件拼装在一起，创建大量的可能的系统。 显然，上面说的不是面向 主要关心 API 及 API 实现的最终用户， 但在一个更多样化和模块化且持续演变的世界中，这可能一条途径可以通往简洁的单个系统。 因为随着可靠的、灵活的构建模块的出现，实现分布式系统的时间由年缩减为周，聚合形成大型整体系统的压力将会消失。

**日志在系统架构中的地位**
提供外部日志的系统允许各个系统抛弃很多各自的复杂性，依靠共享的日志。在我看来，日志可以做到以下事情：

- 通过对节点的并发更新的排序处理，处理了数据一致性（无论即时的还是最终的一致）
- 提供节点之间的数据复制
- 为写入者提供『提交（commit）』语义（仅当写入数据确保不会丢失时才会收到完成确认（acknowledge））
- 为系统提供外部的数据订阅
- 对于丢失数据的失败了的复本，提供恢复或是启动一个新复本的能力
- 调整节点间的数据平衡

这就是一个数据分布式系统所要做的主要部分，实际上，剩下的大部分内容是与 最终用户要面对的查询 API 和索引策略相关的。 这正是不同系统间的应该变化的部分，例如：一个全文搜索查询语句可能需要查询所有分区， 而一个主键查询只需要查询负责这个主键数据的单个节点就可以了。  
![system](https://content.linkedin.com/content/dam/engineering/en-us/blog/migrated/system.png)  

下面我们来看下系统是如何工作的。 系统被分为两个逻辑部分：日志和服务层。日志按顺序捕获状态变化。 服务节点存储索引提供查询服务需要的所有信息（比如键值存储的索引可能会类似 BTree 或 SSTable，搜索系统可能用的是倒排索引（inverted index））。 写入操作可以直接进入日志，尽管可能经过服务层的代理。 在写入日志的时候会生成逻辑时间戳（称为日志中的索引），如果系统是分区的，我也假定是分区的， 那么日志和服务节点会包含相同分区个数，尽管两者的机器台数可能相差很多。

服务节点订阅日志，并按照日志存储的顺序尽快把日志写到它的本地索引中。

客户端只要在查询语句中提供某次写入操作的时间戳，就可以有从任何节点『读到该次写入』的语义（read-your-write semantics） —— 服务节点收到该查询语句后，会将其中的时间戳与自身索引的位置比较， 如果必要，服务节点会延迟请求直到它的索引至少已经跟上那个时间戳，以避免提供的是旧数据。

服务节点可能会或可能不会需要感知 master 身份或是当选 leader。 对很多简单的使用场景，服务节点集群可以完全无需 leader，因为日志是正确真实的信息源。

分布式系统所需要处理的一件比较复杂的事是 恢复失败节点 和 在结点之间移动分片（partition）。 典型的做法是仅保留一个固定窗口的数据，并把这个数据和分片中存储数据的一个快照关联。 另一个相同效果的做法是，让日志保留数据的完整拷贝，并对日志做垃圾回收。 这把大量的复杂性从特定于系统的系统服务层移到了通用的日志中。

有了这个日志系统，你得到一个成熟完整的订阅 API，这个 API 可以订阅数据存储的内容，驱动到其它系统的 ETL 操作。 事实上，许多系统都可以共享相同的日志以提供不同的索引，如下所示：  
![full-stack](https://content.linkedin.com/content/dam/engineering/en-us/blog/migrated/full-stack.png)  

注意，这样的以日志为中心的系统是如何做到本身即是 在其它系统中要处理和加载的数据流的提供者的呢？ 同样，流处理器既可以消费多个输入流，然后通过这个流处理器输出把这些输入流的数据索引到其它系统中。

我觉得把系统分解成日志和查询 API 的观点很有启迪性，因为使得查询相关的因素与系统的可用性和一致性方面解耦。 我其实觉得这更是个好用的思路，可以对于没按这种方式构建的系统做概念上的分解。

值得一提的是，尽管 Kafka 和 Bookeeper 都是一致性日志，但并不是必须的。 你可以轻松把Dynamo之类的数据库作为你的系统的 最终一致的AP日志和键值对服务层。 这样的日志使用起来很灵活，因为它会重传了旧消息并依赖订阅者的信息处理（很像 Dynamo 所做的）。

很多人认为在日志中维护数据的单独拷贝（特别是做全量数据拷贝）太浪费。 然而事实上，有几个因素可以让这个不成为问题。 首先，日志可以是一种特别高效的存储机制。在我们 Kafka 生产环境的服务器上，每个数据中心都存储了超过75TB的数据。 同时其它的许多服务系统需要的是多得多的内存来提供高效的服务（例如文本搜索，它通常是全在内存里）。 其次，服务系统会用优化过的硬件。例如，我们的在线数据系统或者基于内存提供服务或者使用固态硬盘。 相反，日志系统只需要线性读写，因此很合适用TB级的大硬盘。 最后，如上图所示，多个系统使用日志数据提供服务，日志的成本是分摊到多个索引上。 上面几点合起来使得外部日志的开销相当小。

LinkedIn 正是使用这个模式构建了它很多的实时查询系统。 这些系统的数据来自数据库（使用作为日志概念的数据总线，或是来自 Kafka 的真正日志），提供了在这个数据流上特定的分片、索引和查询能力。 这也是我们实现搜索、social graph 和 OLAP 查询系统的方式。 事实上，把单个数据源（无论来自 Hadoop 的在线数据源还是派生数据源）复制到多个在线服务系统中，这个做法很常见。 这种方式经过了验证可以大大简化系统的设计。 系统根本不需要给外部提供写入 API，Kafka 和数据库通过日志给查询系统提供记录和变更流。 各个分片的结点在本地完成写操作。 这些结点只要机械地把日志中的数据转录到自己的存储中。失败的结点通过回放上游的日志就可以恢复。

系统的强度取决于日志的使用方式。一个完全可靠的系统把日志用作数据分片、结点的存储、负载均衡，以及所有和数据一致性和数据传播（propagation）有关的方面。 在这样的架构中，服务层实际上只不过是一种『缓存』，可以通过直接写日志就能完成某种处理。

### 结束语
如果你从头一直做读到了这，那么我对日志的理解你大部分都知道了。

这里再给一些有意思参考资料，你可以再去看看。

人们会用不同的术语描述同一事物，当你从数据库系统到分布式系统、从各类企业级应用软件到广阔的开源世界查看资料时，这会让人有一些困惑。无论如何，在大方向上还是有一些共同之处。










