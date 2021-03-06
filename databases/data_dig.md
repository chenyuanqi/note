
### 什么是数据挖掘
数据挖掘是指有组织有目的地收集数据、分析数据，并从这些大量数据提取出需要的有用信息，从而寻找出数据中存在的规律、规则、知识以及模式、关联、变化、异常和有意义的结构。  
数据挖掘是一种从大量数据中寻找存在的规律、规则、知识以及模式、关联、变化、异常和有意义的结构的技术，是统计学、数据库技术和人工智能技术等技术的综合。

数据挖掘的本质就是寻找出数据中存在的规律、规则、知识以及模式、关联、变化、异常和有意义的结构。

数据挖掘是一门涉及面很广的交叉学科，包括数理统计、人工智能、计算机等。涉及机器学习、数理统计、神经网络、数据库、模式识别、粗糙集、模糊数学等相关技术。

### 数据挖掘的价值、目的、作用
- 数据挖掘的价值
> 数据挖掘大部分的价值在于利用数据挖掘技术改善预测模型，产生学术价值、促进生产、产生并促进商业利益，一切都是为了商业价值（数据——>信息——>知识——>商业）。

- 数据挖掘的目的
> 数据挖掘的最终目的是要实现数据的价值，所以，单纯的数据挖掘是没有多大意义的。

- 数据挖掘的作用
> 从大量数据中寻找存在的规律、规则、知识以及模式、关联、变化、异常和有意义的结构。

### 数据挖掘的背景
数据正在以空前的速度增长，现在的数据是海量的大数据。  
现在，不缺乏数据，但是却面临一个尴尬的境地 —— 数据极其丰富，信息知识匮乏。  

海量的大数据已经远远超出了人类的理解能力，如果不借助强大的工具和技术，很难弄清楚大数据中所蕴含的信息和知识。重要决策如果只是基于决策制定者的个人经验，而不是基于信息、知识丰富的数据，那么，这就极大地浪费了数据，也极大地给我们的商业、学习、工作、生产带来不便和巨大的阻碍。所以，能够方便、高效、快速地从大数据里提取出巨大的信息和知识是必须面对的，因此，数据挖掘技术应运而生。数据挖掘填补了数据和信息、知识之间的鸿沟。  

数据挖掘技术有助于实现从 DT（数据时代）向 KT（知识时代）转变。

### 数据挖掘软件及其发展
第一代，代表软件：Salford Systems 公司早期的 CART 系统。  
第二代，代表软件：SAS Enterprise Miner；DBMiner，DBMiner 是加拿大 SimonFraser 大学开发的一个多任务数据挖掘系统，它的前身是 DBLearn。  
第三代，代表软件：SPSS Clementine，SPSS Clementine 是 SPSS 公司的一个数据挖掘平台；RapidMiner，RapidMiner 是世界领先的数据挖掘解决方案。  
第四代，正在开发。  

### 数据挖掘技术及其分类
- 数据挖掘技术（方法）分为两大类
> （1）预言（Predication）：用历史预测未来  
> （2）描述（Description）：了解数据中潜在的规律  

- 数据挖掘技术
> 数据挖掘常用的方法有：分类、聚类、回归分析、关联规则、神经网络、特征分析、偏差分析等。这些方法从不同的角度对数据进行挖掘。  
> 
> 分类的含义：找出数据库中的一组数据对象的共同特点并按照分类模式将其划分为不同的类。分类是依靠给定的类别对对象进行划分的。  
> 分类的目的（作用）：通过分类模型，将数据库中的数据项映射到某个给定的类别中。  
> 分类的应用：客户的分类、客户的属性和特征分析、客户满意度分析、客户的购买趋势预测、应用分类、趋势预测等。  
> 主要的分类方法：决策树、KNN 法 (K-Nearest Neighbor)、SVM 法、VSM 法、Bayes 法、神经网络等。  
> 分类算法的局限：分类作为一种监督学习方法，要求必须事先明确知道各个类别的信息，并且断言所有待分类项都有一个类别与之对应。  
> 但是，很多时候上述条件得不到满足，尤其是在处理海量数据的时候，如果通过预处理使得数据满足分类算法的要求，则代价非常大，这时候可以考虑使用聚类算法。  
> 
> 聚类的含义：聚类指事先并不知道任何样本的类别标号，按照对象的相似性和差异性，把一组对象划分成若干类，并且每个类里面对象之间的相似度较高，不同类里面对象之间相似度较低或差异明显。  
> 我们并不关心某一类是什么，我们需要实现的目标只是把相似的东西聚到一起，聚类是一种无监督学习。  
> 聚类与分类的区别：聚类类似于分类，但是，与分类不同的是，聚类不依靠给定的类别对对象进行划分，而是根据数据的相似性和差异性将一组数据分为几个类别。  
> 聚类的目的：聚类与分类的目的不同。聚类是要按照对象的相似性和差异性将对象进行分类，属于同一类别的数据间的相似性很大，但不同类别之间数据的相似性很小，跨类的数据关联性很低。  
> 组内的相似性越大，组间差别越大，聚类就越好。  
> 聚类的方法（算法）：主要的聚类算法可以划分为如下几类，划分方法、层次方法、基于密度的方法、基于网格的方法、基于模型的方法。  
> 每一类中都存在着得到广泛应用的算法， 划分方法中有 k-means 聚类算法、层次方法中有凝聚型层次聚类算法、基于模型方法中有神经网络聚类算法。  
> 聚类的应用：它可以应用到客户群体的分类、客户背景分析、客户购买趋势预测、市场的细分等。  
> 
> 回归分析的含义：回归分析是一个统计预测模型，用以描述和评估因变量与一个或多个自变量之间的关系；  
> 反映的是事务数据库中属性值在时间上的特征，产生一个将数据项映射到一个实值预测变量的函数，发现变量或属性间的依赖关系。  
> 回归分析的目的（作用）：回归分析反映了数据库中数据的属性值在时间上的特征，通过函数表达数据映射的关系来发现属性值之间的依赖关系。  
> 回归分析的应用：回归分析方法被广泛地用于解释市场占有率、销售额、品牌偏好及市场营销效果。  
> 它可以应用到市场营销的各个方面，如客户寻求、保持和预防客户流失活动、产品生命周期分析、销售趋势预测及有针对性的促销活动等。  
> 回归分析的主要研究问题：数据序列的趋势特征、数据序列的预测、数据间的相关关系等。  
> 
> 关联规则的含义：关联规则是隐藏在数据项之间的关联或相互关系，即可以根据一个数据项的出现推导出其他数据项的出现。  
> 关联规则是描述数据库中数据项之间所存在的关系的规则。  
> 关联规则的目的（作用）：发现隐藏在数据间的关联或相互关系，从一件事情的发生，来推测另外一件事情的发生，从而更好地了解和掌握事物的发展规律等等。  
> 关联规则的挖掘过程主要包括两个阶段：第一阶段为从海量原始数据中找出所有的高频项目组；第二阶段为从这些高频项目组产生关联规则。  
> 关联规则的应用：关联规则挖掘技术已经被广泛应用于金融行业企业中用以预测客户的需求，各银行在自己的 ATM 机上通过捆绑客户可能感兴趣的信息供用户了解并获取相应信息来改善自身的营销。  
> 
> 神经网络作为一种先进的人工智能技术，因其自身自行处理、分布存储和高度容错等特性非常适合处理非线性的问题，以及那些以模糊、不完整、不严密的知识或数据为特征的问题，它的这一特点十分适合解决数据挖掘的问题。  
> 典型的神经网络模型主要分为三大类：第一类是以用于分类预测和模式识别的前馈式神经网络模型，其主要代表为函数型网络、感知机。  
> 第二类是用于联想记忆和优化算法的反馈式神经网络模型，以 Hopfield 的离散模型和连续模型为代表。  
> 第三类是用于聚类的自组织映射方法，以 ART 模型为代表。虽然神经网络有多种模型及算法，但在特定领域的数据挖掘中使用何种模型及算法并没有统一的规则，而且人们很难理解网络的学习及决策过程。  
> 
> Web 数据挖掘的含义：Web 数据挖掘是一项综合性技术，指 Web 从文档结构和使用的集合C中发现隐含的模式 P，如果将 C 看做是输入，P 看做是输出，那么 Web 挖掘过程就可以看做是从输入到输出的一个映射过程。  
> Web 数据挖掘的研究对象：是以半结构化和无结构文档为中心的 Web，这些数据没有统一的模式，数据的内容和表示互相交织，数据内容基本上没有语义信息进行描述，仅仅依靠 HTML 语法对数据进行结构上的描述。当前越来越多的 Web 数据都是以数据流的形式出现的，因此对 Web 数据流挖掘就具有很重要的意义。  
> 目前常用的 Web 数据挖掘算法：PageRank 算法、HITS 算法、LOGSOM 算法。这三种算法提到的用户都是笼统的用户，并没有区分用户的个体。  
> Web 数据挖掘的应用：可以利用 Web 的海量数据进行分析，收集政治、经济、政策、科技、金融、各种市场、竞争对手、供求信息、客户等有关的信息，集中精力分析和处理那些对企业有重大或潜在重大影响的外部环境信息和内部经营信息，并根据分析结果找出企业管理过程中出现的各种问题和可能引起危机的先兆，对这些信息进行分析和处理，以便识别、分析、评价和管理危机。  
> 目前 Web 数据挖掘面临着一些问题：用户的分类问题、网站内容时效性问题，用户在页面停留时间问题，页面的链入与链出数问题等。  
> 
> 特征分析的含义：特征分析是从数据库中的一组数据中提取出关于这些数据的特征式，这些特征式表达了该数据集的总体特征。  
> 特征分析的目的（作用）：在于从海量数据中提取出有用信息，从而提高数据的使用效率。  
> 特征分析的应用：如营销人员通过对客户流失因素的特征提取，可以得到导致客户流失的一系列原因和主要特征，利用这些特征可以有效地预防客户的流失。  
> 
> 偏差分析的含义：偏差是数据集中的小比例对象。通常，偏差对象被称为离群点、例外、野点等。偏差分析就是发现与大部分其他对象不同的对象。  
> 偏差分析的应用：在企业危机管理及其预警中，管理者更感兴趣的是那些意外规则。  
> 意外规则的挖掘可以应用到各种异常信息的发现、分析、识别、评价和预警等方面。而其成因源于不同的类、自然变异、数据测量或收集误差等。  

- “异常”的定义
> Hawkins 给出了异常的本质性的定义：异常是数据集中与众不同的数据，使人怀疑这些数据并非随机偏差，而是产生于完全不同的机制。  
> 聚类算法对异常的定义：异常是聚类嵌于其中的背景噪声。  
> 异常检测算法对异常的定义：异常是既不属于聚类也不属于背景噪声的点。他们的行为与正常的行为有很大不同。  

### 数据挖掘的流程
- 常用的数据挖掘的流程模式
> 数据挖掘有很多不同的实施方法，如果只是把数据拉到 Excel 表格中计算一下算是数据分析， 而不是数据挖掘。常用的数据挖掘流程 CRISP-DM 和 SEMMA。

- 数据挖掘的一般流程
> 通常情况下，数据挖掘需要有 7 个步骤：数据集选取或构造、数据预处理、数据转换、数据建模、进行数据挖掘、结果分析（模型评估）与改进、知识表示。  
> （1）数据集选取或构造 根据任务的目的，选择数据集；或者从实际中构造自己需要的数据。  
> （2）数据预处理  
> 确定数据集后，就开始对数据进行预处理，使得数据经过预处理后能够为我们所用。数据预处理提高了数据质量：准确性、完整性和一致性。  
> 
> 数据预处理包括：数据清理、数据集成、数据规约、数据变换和离散化。  
> 
> 数据清理的含义：数据清洗指的是删除原始数据集中的无关数据、重复数据，平滑噪声数据，筛选掉与挖掘主题无关的数据，处理缺失值、异常值等等。  
> 数据清理的目的和作用：主要是删除原始数据集中的无关数据、重复数据，平滑噪声数据，筛选掉与挖掘主题无关的数据，处理缺失值、异常值等等。  
> 在数据库中的数据有一些是不完整的（有些感兴趣的属性缺少属性值）、含噪声的（包含错误的属性值）、不一致的（同样的信息不同的表示方式），因此需要进行数据清理，将完整、正确、一致的数据信息存入数据仓库中。不然，挖掘的结果会差强人意。  
> 数据清理的内容：忽略元祖、人工填写缺失值、使用属性的中心度量填充、给定同一类所有样本的属性均值或中位数填充、最可能的值填充。  
> 
> 数据集成的含义：把不同来源、格式、特点性质的数据在逻辑上或物理上有机地集中，从而为企业提供全面的数据共享。  
> 数据集成的目的和作用：把不同来源、格式、特点性质的数据在逻辑上或物理上有机地集中，从而为企业提供全面的数据共享。  
> 数据集成的内容：实体识别、冗余和相关分析(卡方检验，相关系数，协方差等，用 spss 比较方便)。  
> 
> 数据规约的含义：数据归约是指在尽可能保持数据原貌的前提下，最大限度地精简数据量。  
> 数据规约的目的和作用：在海量数据上进行复杂的数据分析扣数据挖掘将需要很长时间，而在执行多数的数据挖掘算法时，即使在少量数据上也需要很长的时间，使得在原有的海量数据上进行这种分析不现实或不可行。所以必须利用数据规约技术将数据量减少，缩短数据分析和数据挖掘的时间。  
> 数据规约技术可以用来得到数据集的规约表示，数据集经过规约表示后，数据集就小得多，但仍然接近于保持原数据的完整性，并且规约后执行数据挖掘结果与规约前执行结果相同或几乎相同。  
> 数据规约的方法：维规约(小波变换和主成分分析，最常用)、数量规约(较小的数据替代原始数据)、数据压缩(有损无损两种，尤其对于图像视频等多媒体常用)。  
> 数据归约主要有两个途径：属性选择和数据采样，分别针对原始数据集中的属性和记录。  
> 特征归约：是从原有的特征中删除不重要或不相关的特征，或者通过对特征进行重组来减少特征的个数。其原则是在保留、甚至提高原有判别能力的同时减少特征向量的维度。  
> 特征归约算法的输入是一组特征，输出是它的一个子集。  
> 在领域知识缺乏的情况下进行特征归约时一般包括3个步骤：  
> 搜索过程：在特征空间中搜索特征子集，每个子集称为一个状态由选中的特征构成。  
> 评估过程：输入一个状态，通过评估函数或预先设定的阈值输出一个评估值搜索算法的目的是使评估值达到最优。  
> 分类过程：使用最终的特征集完成最后的算法。 特征归约处理的效果： ///更少的数据，提高挖掘效率 ///更高的数据挖掘处理精度 ///简单的数据挖掘处理结果 ///更少的特征。  
> 
> 数据变换的含义：数据变换指的是将数据转换或者统一为适合进行数据挖掘的形式。  
> 数据变换的内容: 包括光滑、聚集、数据泛化、规范化和属性构造。  
> 光滑：去掉数据中的噪声，主要有分箱、回归和聚类等方法。  
> 聚集：通过对数据仓库中的数据进行简单的汇总和聚集来获得统计信息，以便对数据进行更高层次的分析。  
> 数据泛化：使用概念分层的方式，利用高层的概念来替换低层或原始数据。  
> 规范化：对属性数据进行缩放，使之可以落入到一个特定区域之间，主要有最小-最大规范化、Z-Score 规范化（利用均值和标准差）以及小数定标（除以10的n次方，使之落到[-1,1]）规范化等方法。  
> 属性构造：构造新的属性并添加到属性集合中以便帮助挖掘。  
> 
> 数据离散化的含义：数据离散化是指将连续的数据进行分段，使其变为一段段离散化的区间。分段的原则有基于等距离、等频率、优化的方法。  
> 数据离散化的主要原因：算法需要比如决策树、朴素贝叶斯等算法，都是基于离散型的数据展开的。如果要使用该类算法，必须将离散型的数据进行。有效的离散化能减小算法的时间和空间开销，提高系统对样本的分类聚类能力和抗噪声能力。  
> 数据转换就是将进过处理后的数据转换为特征，这些特征要尽可能的准确地描述数据，并且使得算法达到最优。  
> 
> （3）数据转换  
> 数据转换就是将进过处理后的数据转换为特征，这些特征要尽可能的准确地描述数据，并且使得算法达到最优。  
> 
> （4）数据建模  
> 数据建模的含义：根据模型优缺点和任务情况，选择适宜本任务的最佳模型。  
> 选择最佳模型的方式：其中一种方式是对每个模型都进行训练，再统计测试数据的误差，选择误差最小的模型即可。  
> 调整模型参数：选好模型后，还需要调整模型的参数，使得模型表现尽可能最优。主要方法有手动调优、网格搜索、随机搜索以及基于贝叶斯的参数调优方法。  
> 
> （5）进行数据挖掘  
> 根据数据仓库中的数据信息，选择合适的分析工具，应用统计方法、事例推理、决策树、规则推理、模糊集、甚至神经网络、遗传算法等方法处理信息， 得出有用的分析信息。  
> 
> （6）结果分析（模型评估）和改进模型  
> 结果分析（模型评估）：从商业角度来看，结果分析（模型评估）就是由行业专家来验证数据挖掘结果的正确性。分析的对象主要是模型的优缺点，客观公正的评判自己的作品(能有高手帮忙最好啦)能清醒自己的认知。  
> 改进模型：就是从分析当中来，对模型进行改进。  
> 
> （7）知识表示  
> 知识表示就是将数据挖掘所得到的分析信息以可视化的方式呈现给用户，或作为新的知识存放在知识库中，供其他应用程序使用。  

数据挖掘过程是一个反复循环的过程，每一个步骤如果没有达到预期目标，都需要回到前面的步骤，重新调整并执行。  
不是每件数据挖掘的工作都需要这里列出的每一步，例如在某个工作中不存在多个数据源的时候，数据集成的步骤便可以省略。  
在数据挖掘中，至少 60% 的费用可能要花在步骤（1）信息收集阶段，而至少 60% 以上的精力和时间是花在数据预处理过程上。  

### 数据挖掘的九大规律
1. 每个数据挖掘解决方案的根源都是有商业目的的。  
2. 数据挖掘过程的每一步都需要以商业信息为中心。  
3. 数据挖掘过程前期的数据准备工作要超过整个过程的一半。  
4. 对数据挖掘者来说，没有免费的午餐，数据挖掘的任何一个过程都是来之不易的。  
5. 在数据的世界里，总是有模式、规律可循的，您找不到模式和规律不是因为这些模式和规律不存在，而是因为您还没有发现它。  
6. 数据挖掘可以把商业领域的信息放大。  
7. 预测可以为我们增加信息。  
8. 数据挖掘模式的精准和稳定并不决定数据挖掘过程的价值，换句话说，技术手段再精妙，没有商业意义和合适的商业应用，这个数据挖掘也是没有价值的。  
9. 所有的模式都会变化。  

### 数据挖掘的应用
现在的数据挖掘技术在商业上的应用已经相当广泛了，因为对数据挖掘技术进行支撑的三种基础技术已经发展成熟。  
这三种基础技术是：   
（1）海量数据收集和存储技术。   
（2）强大的计算集群和分布式计算技术。   
（3）数据挖掘算法。  

数据挖掘的最终目的是要实现数据的价值，所以，单纯的数据挖掘是没有多大意义的，只有以商业利益、生产价值为目的的数据挖掘才是有意义的，才能更好地实现数据的价值。实现数据的价值的最佳方式有很多，而商业智能（Business Intelligence , BI）是在企业中实现数据价值的最佳方式之一。  

### 数据挖掘的常用算法
- 数据挖掘算法的含义
> 数据挖掘算法是根据数据创建数据挖掘模型的一组试探方法和计算模式。  

- 数据挖掘算法的作用
> 实现数据挖掘模型的功能，找出数据中的潜在规律。  

- 数据挖掘算法的三要素
> （1）模式记述语言：反映了算法可以发现什么样的知识。  
> （2）模式评价：反映了什么样的模式可以称为知识。  
> （3）模式探索：包括针对某一特定模式对参数空间的探索和对模式空间的探索。  

- 常用算法
> （1）决策树  
> 决策树是一个预测模型，它代表的是对象属性值与对象值之间的一种映射关系。从数据产生决策树的技术叫做决策树学习。  
> 决策树中的每个节点表示某个对象，每个分叉路径则代表的某个可能的属性值，而每个叶结点则对应具有上述属性值的子对象。决策树仅有单一输出，若需要多个输出，可以建立独立的决策树以处理不同输出。  
> 每个决策树都表述了一种树型结构，它由它的分支来对该类型的对象依靠属性进行分类。每个决策树可以依靠对源数据库的分割进行数据测试，这个过程可以递归式的对树进行修剪，当不能再进行分割或一个单独的类可以被应用于某一分支时，递归过程就完成了。  
> 决策树的工作原理   
> a、决策树一般都是自上而下的来生成的。  
> b、选择分割的方法有多种，但是目的都是一致的，即对目标类尝试进行最佳的分割。  
> c、从根节点到叶子节点都有一条路径，这条路径就是一条“规则”。  
> d、决策树可以是二叉的，也可以是多叉的。  
> 决策树的代表算法 —— C4.5 算法。C4.5 算法是 ID3 算法的一个改进算法，C4.5 算法是机器学习算法中的一种分类决策树算法，分类决策树算法是从大量事例中进行提取分类规则的自上而下的决策树。  
> 
> （2）SVM（支持向量机）  
> 支持向量机英文为 Support Vector Machine，简称 SV 机（论文中一般简称 SVM）。它是一种监督式学习的方法，它广泛的应用于统计分类以及回归分析中。支持向量机属于一般化线性分类器。支持向量机也被称为最大边缘区分类器。  
> 
> （3）SVM 的主要思想  
> 它是针对线性可分情况进行分析，对于线性不可分的情况，通过使用非线性映射算法将低维输入空间线性不可分的样本转化为高维特征空间使其线性可分，从而使得高维特征空间采用线性算法对样本的非线性特征进行线性分析成为可能。  
> 它基于结构风险最小化理论之上在特征空间中建构最优分割超平面，使得学习器得到全局最优化，并且在整个样本空间的期望风险以某个概率满足一定上界。  
> SVM的优势 ///可以解决小样本情况下的机器学习问题。 ///可以提高泛化性能。 ///可以解决高维问题。 ///可以解决非线性问题。 ///可以避免神经网络结构选择和局部极小点问题。  
> 
> （4）贝叶斯（Bayes）分类器  
> 贝叶斯分类器是用于分类的贝叶斯网络。贝叶斯网络是一个带有概率注释的有向无环图，图中的每一个结点均表示一个随机变量，图中两结点间若存在着一条弧，则表示这两结点相对应的随机变量是概率相依的，反之则说明这两个随机变量是条件独立的。目前研究较多的贝叶斯分类器主要有四种，分别是：Naive Bayes、TAN、BAN 和 GBN。  
> 贝叶斯分类器的分类原理  
> 通过某对象的先验概率，利用贝叶斯公式计算出其后验概率，即该对象属于某一类的概率，选择具有最大后验概率的类作为该对象所属的类。  
> 应用贝叶斯网络分类器进行分类  
> 应用贝叶斯网络分类器进行分类主要分成两阶段。第一阶段是贝叶斯网络分类器的学习，即从样本数据中构造分类器。第二阶段是贝叶斯网络分类器的推理，即计算类结点的条件概率，对分类数据进行分类。  
> 这两个阶段的时间复杂性均取决于特征值间的依赖程度，甚至可以是NP完全问题（世界七大数学难题之一），因而在实际应用中，往往需要对贝叶斯网络分类器进行简化。  
> 
> （5）朴素贝叶斯（Naive Bayes）分类器  
> 朴素贝叶斯分类发源于古典数学理论，有着坚实的数学基础，以及稳定的分类效率。同时，所需估计的参数很少，对缺失数据不太敏感，算法也比较简单。理论上，与其他分类方法相比具有最小的误差率。  
> 
> （6）k 最近邻算法（k-Nearest Neighbor algorithm）
> K 最近邻(k-Nearest Neighbor，KNN)分类算法，是一个理论上比较成熟的方法，也是最简单的机器学习算法之一。  
> k最近邻算法的思路  
> 如果一个样本在特征空间中的 k 个最相似(即特征空间中最邻近)的样本中的大多数属于某一个类别，则该样本也属于这个类别。 KNN 算法中，所选择的邻居都是已经正确分类的对象。该方法在定类决策上只依据最邻近的一个或者几个样本的类别来决定待分样本所属的类别。  
> KNN 方法虽然从原理上也依赖于极限定理，但在类别决策时，只与极少量的相邻样本有关。由于 KNN 方法主要靠周围有限的邻近的样本，而不是靠判别类域的方法来确定所属类别的，因此对于类域的交叉或重叠较多的待分样本集来说，KNN 方法较其他方法更为适合。  
> KNN 算法的应用  
> 不仅可以用于分类，还可以用于回归。该算法比较适用于样本容量比较大的类域的自动分类，而那些样本容量较小的类域采用这种算法比较容易产生误分。  
> KNN 算法的缺点  
> 该算法在分类时的不足是，当样本不平衡时，如一个类的样本容量很大，而其他类样本容量很小时，有可能导致当输入一个新样本时，该样本的 K 个邻居中大容量类的样本占多数。因此可以采用权值的方法（和该样本距离小的邻居权值大）来改进。  
> 
> （7）回归树分类器  
> 如果要选择在很大范围的情形下性能都好的、同时不需要应用开发者付出很多的努力并且易于被终端用户理解的分类技术的话，那么 Brieman,Friedman, Olshen 和 Stone（1984）提出的分类树方法是一个强有力的竞争者。  
> 
> （8）人工神经网络（ANN, artificial neural network）  
> 人工神经网络的含义 是由具有适应性的简单单元组成的广泛并行互连的网络，它的组织能够模拟生物神经系统对真实世界物体所作出的交互反应。  
> 人工神经网络研究的局限性  
> a、研究受到脑科学研究成果的限制。  
> b、缺少一个完整、成熟的理论体系。  
> c、研究带有浓厚的策略和经验色彩。  
> 人工神经网络的特点  
> 一般而言, ANN 与经典计算方法相比并非优越,只有当常规方法解决不了或效果不佳时 ANN 方法才能显示出其优越性。尤其对问题的机理不甚了解或不能用数学模型表示的系统,如故障诊断、特征提取和预测等问题, ANN 往往是最有利的工具。  
> 另一方面，ANN 对处理大量原始数据而不能用规则或公式描述的问题, 表现出极大的灵活性和自适应性。  

人工神经网络以其具有自学习、自组织、较好的容错性和优良的非线性逼近能力，受到众多领域学者的关注。在实际应用中，80%～90% 的人工神经网络模型是采用误差反传算法或其变化形式的网络模型（简称 BP 网络），目前主要应用于函数逼近、模式识别、分类和数据压缩或数据挖掘。  
