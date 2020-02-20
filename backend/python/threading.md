
### Python 多线程
一个进程还可以拥有多个并发的执行线索，简单的说就是拥有多个可以获得 CPU 调度的执行单元，这就是所谓的线程。由于线程在同一个进程下，它们可以共享相同的上下文，因此相对于进程而言，线程间的信息共享和通信更加容易。当然在单核 CPU 系统中，真正的并发是不可能的，因为在某个时刻能够获得 CPU 的只有唯一的一个线程，多个线程共享了 CPU 的执行时间。使用多线程实现并发编程为程序带来的好处是不言而喻的，最主要的体现在提升程序的性能和改善用户体验。  

多线程是加速程序计算的有效方式。

python 提供了两个模块来实现多线程 thread 和 threading ，thread 有一些缺点，在 threading 得到了弥补。
```python
# -*- coding: UTF-8 -*-
import threading
from time import ctime,sleep

def music(name):
    for i in range(2):
        print "I was listening to %s. %s" %(name, ctime())
        sleep(1)

def movie(name):
    for i in range(2):
        print "I was at the %s! %s" %(name, ctime())
        sleep(4)

threads = []
# 定义线程
t1 = threading.Thread(target=music,args=(u'爱情买卖',))
threads.append(t1)
t2 = threading.Thread(target=movie,args=(u'阿凡达',))
threads.append(t2)

# 获取已激活的线程数：threading.active_count()
# 查看所有线程信息：threading.enumerate()
# 查看正在运行的线程：threading.current_thread()

if __name__ == '__main__':
    for t in threads:
        # 将线程声明为守护线程，必须在 start 前
        t.setDaemon(True)
        # 让线程开始工作
        t.start() 

    # join 对控制多个线程的执行顺序非常关键
    # 在子线程完成运行之前，这个子线程的父线程将一直被阻塞
    t.join()

    print "all over %s" %ctime()

    # I was listening to 爱情买卖. Tue Aug 13 06:33:06 2019
    # I was at the 阿凡达! Tue Aug 13 06:33:06 2019
    # I was listening to 爱情买卖. Tue Aug 13 06:33:07 2019
    # I was at the 阿凡达! Tue Aug 13 06:33:10 2019
    # all over Tue Aug 13 06:33:14 2019
```

### Python GIL
Python 提供了多线程包，但并不支持真正意义上的多线程。  

Python 的设计上，有一个必要的环节，就是全局解释器锁 Global Interpreter Lock (GIL)，它让 Python 还是一次性只能处理一件事情。  
> 尽管 Python 完全支持多线程编程， 但是解释器的 C 语言实现部分在完全并行执行时并不是线程安全的。   
> 实际上，解释器被一个全局解释器锁保护着，它确保任何时候都只有一个 Python 线程执行。 GIL 最大的问题就是 Python 的多线程程序并不能利用多核 CPU 的优势 （比如一个使用了多个线程的计算密集型程序只会在一个单 CPU 上面运行）。  
> 在讨论普通的 GIL 之前，有一点要强调的是 GIL 只会影响到那些严重依赖 CPU 的程序（比如计算型的）而不是 I/O。  
> 如果你的程序大部分只会涉及到 I/O，比如网络交互，那么使用多线程就很合适， 因为它们大部分时间都在等待。实际上，你完全可以放心的创建几千个 Python 线程， 现代操作系统运行这么多线程没有任何压力，没啥可担心的。  
> 有一点要注意的是，线程不是专门用来优化性能的。  

解决 GIL 缺点的两种策略：  
1、使用 multiprocessing 模块来创建一个进程池， 并像协同处理器一样的使用它  
2、使用 C 扩展编程技术  

### Python 线程锁
lock 在不同线程使用同一共享内存时，能够确保线程之间互不影响。  

使用 lock 的方法是， 在每个线程执行运算修改共享内存之前，执行 lock.acquire() 将共享内存上锁， 确保当前线程执行时，内存不会被其他线程访问，执行运算完毕后，使用 lock.release() 将锁打开， 保证其他的线程可以使用该共享内存。  
```python
import threading

def job1():
    global A,lock
    lock.acquire()
    for i in range(10):
        A+=1
        print('job1',A)
    lock.release()

def job2():
    global A,lock
    lock.acquire()
    for i in range(10):
        A+=10
        print('job2',A)
    lock.release()

if __name__== '__main__':
    lock=threading.Lock()
    A=0
    t1=threading.Thread(target=job1)
    t2=threading.Thread(target=job2)
    t1.start()
    t2.start()
    t1.join()
    t2.join()
```

### 多线程使用示例
#### 1 默认启动主线程

一般的，程序默认执行只在一个线程，这个线程称为主线程，例子演示如下：

导入线程相关的模块 `threading`:

```python
import threading
```

threading的类方法 `current_thread()`返回当前线程：

```python
t = threading.current_thread()
print(t) # <_MainThread(MainThread, started 139908235814720)>
```

所以，验证了程序默认是在`MainThead`中执行。

`t.getName()`获得这个线程的名字，其他常用方法，`getName()`获得线程`id`,`isAlive()`判断线程是否存活等。

```python
print(t.getName()) # MainThread
print(t.ident) # 139908235814720
print(t.isAlive()) # True
```

以上这些仅是介绍多线程的`背景知识`，因为到目前为止，我们有且仅有一个"干活"的主线程

#### 2 创建线程

创建一个线程：

```python
my_thread = threading.Thread()
```

创建一个名称为`my_thread`的线程：

```python
my_thread = threading.Thread(name='my_thread')
```

创建线程的目的是告诉它帮助我们做些什么，做些什么通过参数`target`传入，参数类型为`callable`，函数就是可调用的：

```python
def print_i(i):
    print('打印i:%d'%(i,))
my_thread = threading.Thread(target=print_i,args=(1,))
```

`my_thread`线程已经全副武装，但是我们得按下发射按钮，启动start()，它才开始真正起飞。

```python
my_thread().start()
```

打印结果如下，其中`args`指定函数`print_i`需要的参数i，类型为元祖。

```python
打印i:1
```

至此，多线程相关的核心知识点，已经总结完毕。但是，仅仅知道这些，还不够！光纸上谈兵，当然远远不够。

接下来，聊聊应用多线程编程，最本质的一些东西。

**3 交替获得 CPU 时间片**

为了更好解释，假定计算机是单核的，尽管对于`cpython`，这个假定有些多余。

开辟3个线程，装到`threads`中:

```python
import time
from datetime import datetime
import threading


def print_time():
    for _ in range(5): # 在每个线程中打印5次
        time.sleep(0.1) # 模拟打印前的相关处理逻辑
        print('当前线程%s,打印结束时间为:%s'%(threading.current_thread().getName(),datetime.today()))


threads = [threading.Thread(name='t%d'%(i,),target=print_time) for i in range(3)]
```

启动3个线程：

```python
[t.start() for t in threads]
```

打印结果如下，`t0`,`t1`,`t2`三个线程，根据操作系统的调度算法，轮询获得CPU时间片，注意观察，`t2`线程可能被连续调度，从而获得时间片。

```python
当前线程t0,打印结束时间为:2020-01-12 02:27:15.705235
当前线程t1,打印结束时间为:2020-01-12 02:27:15.705402
当前线程t2,打印结束时间为:2020-01-12 02:27:15.705687
当前线程t0,打印结束时间为:2020-01-12 02:27:15.805767
当前线程t1,打印结束时间为:2020-01-12 02:27:15.805886
当前线程t2,打印结束时间为:2020-01-12 02:27:15.806044
当前线程t0,打印结束时间为:2020-01-12 02:27:15.906200
当前线程t2,打印结束时间为:2020-01-12 02:27:15.906320
当前线程t1,打印结束时间为:2020-01-12 02:27:15.906433
当前线程t0,打印结束时间为:2020-01-12 02:27:16.006581
当前线程t1,打印结束时间为:2020-01-12 02:27:16.006766
当前线程t2,打印结束时间为:2020-01-12 02:27:16.007006
当前线程t2,打印结束时间为:2020-01-12 02:27:16.107564
当前线程t0,打印结束时间为:2020-01-12 02:27:16.107290
当前线程t1,打印结束时间为:2020-01-12 02:27:16.107741
```

#### 3 多线程抢夺同一个变量

多线程编程，存在抢夺同一个变量的问题。

比如下面例子，创建的10个线程同时竞争全局变量`a`:
​

```python
import threading


a = 0
def add1():
    global a    
    a += 1
    print('%s  adds a to 1: %d'%(threading.current_thread().getName(),a))
    
threads = [threading.Thread(name='t%d'%(i,),target=add1) for i in range(10)]
[t.start() for t in threads]
```

执行结果：

```python
t0  adds a to 1: 1
t1  adds a to 1: 2
t2  adds a to 1: 3
t3  adds a to 1: 4
t4  adds a to 1: 5
t5  adds a to 1: 6
t6  adds a to 1: 7
t7  adds a to 1: 8
t8  adds a to 1: 9
t9  adds a to 1: 10
```

结果一切正常，每个线程执行一次，把`a`的值加1，最后`a` 变为10，一切正常。

运行上面代码十几遍，一切也都正常。

所以，我们能下结论：这段代码是线程安全的吗？

NO！

多线程中，只要存在同时读取和修改一个全局变量的情况，如果不采取其他措施，就一定不是线程安全的。

尽管，有时，某些情况的资源竞争，暴露出问题的概率`极低极低`：

本例中，如果线程0 在修改a后，其他某些线程还是get到的是没有修改前的值，就会暴露问题。



但是在本例中，`a = a + 1`这种修改操作，花费的时间太短了，短到我们无法想象。所以，线程间轮询执行时，都能get到最新的a值。所以，暴露问题的概率就变得微乎其微。

#### 4 代码稍作改动，叫问题暴露出来

只要弄明白问题暴露的原因，叫问题出现还是不困难的。

想象数据库的写入操作，一般需要耗费我们可以感知的时间。

为了模拟这个写入动作，简化期间，我们只需要延长修改变量`a`的时间，问题很容易就会还原出来。

```python
import threading
import time


a = 0
def add1():
    global a    
    tmp = a + 1
    time.sleep(0.2) # 延时0.2秒，模拟写入所需时间
    a = tmp
    print('%s  adds a to 1: %d'%(threading.current_thread().getName(),a))
    
threads = [threading.Thread(name='t%d'%(i,),target=add1) for i in range(10)]
[t.start() for t in threads]
```

重新运行代码，只需一次，问题立马完全暴露，结果如下：

```python
t0  adds a to 1: 1
t1  adds a to 1: 1
t2  adds a to 1: 1
t3  adds a to 1: 1
t4  adds a to 1: 1
t5  adds a to 1: 1
t7  adds a to 1: 1
t6  adds a to 1: 1
t8  adds a to 1: 1
t9  adds a to 1: 1
```

看到，10个线程全部运行后，`a`的值只相当于一个线程执行的结果。

下面分析，为什么会出现上面的结果：

这是一个很有说服力的例子，因为在修改a前，有0.2秒的休眠时间，某个线程延时后，CPU立即分配计算资源给其他线程。直到分配给所有线程后，根据结果反映出，0.2秒的休眠时长还没耗尽，这样每个线程get到的a值都是0，所以才出现上面的结果。



以上最核心的三行代码：

```python
tmp = a + 1
time.sleep(0.2) # 延时0.2秒，模拟写入所需时间
a = tmp
```





#### 5 加上一把锁，避免以上情况出现

知道问题出现的原因后，要想修复问题，也没那么复杂。

通过python中提供的锁机制，某段代码只能单线程执行时，上锁，其他线程等待，直到释放锁后，其他线程再争锁，执行代码，释放锁，重复以上。

创建一把锁`locka`:

```python
import threading
import time


locka = threading.Lock()
```

通过 `locka.acquire()` 获得锁，通过`locka.release()`释放锁，它们之间的这些代码，只能单线程执行。

```python
a = 0
def add1():
    global a    
    try:
        locka.acquire() # 获得锁
        tmp = a + 1
        time.sleep(0.2) # 延时0.2秒，模拟写入所需时间
        a = tmp
    finally:
        locka.release() # 释放锁
    print('%s  adds a to 1: %d'%(threading.current_thread().getName(),a))
    
threads = [threading.Thread(name='t%d'%(i,),target=add1) for i in range(10)]
[t.start() for t in threads]
```

执行结果如下：

```python
t0  adds a to 1: 1
t1  adds a to 1: 2
t2  adds a to 1: 3
t3  adds a to 1: 4
t4  adds a to 1: 5
t5  adds a to 1: 6
t6  adds a to 1: 7
t7  adds a to 1: 8
t8  adds a to 1: 9
t9  adds a to 1: 10
```

一起正常，其实这已经是单线程顺序执行了，就本例子而言，已经失去多线程的价值，并且还带来了因为线程创建开销，浪费时间的副作用。

程序中只有一把锁，通过 `try...finally`还能确保不发生死锁。但是，当程序中启用多把锁，还是很容易发生死锁。

注意使用场合，避免死锁，是我们在使用多线程开发时需要注意的一些问题。

