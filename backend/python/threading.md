
### Python 多线程
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
