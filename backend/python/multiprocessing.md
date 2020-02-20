
### Python 多进程
进程就是操作系统中执行的一个程序，操作系统以进程为单位分配存储空间，每个进程都有自己的地址空间、数据栈以及其他用于跟踪进程执行的辅助数据，操作系统管理所有进程的执行，为它们合理的分配资源。进程可以通过 fork 或 spawn 的方式来创建新的进程来执行其他的任务，不过新的进程也有自己独立的内存空间，因此必须通过进程间通信机制（IPC，Inter-Process Communication）来实现数据共享，具体的方式包括管道、信号、套接字、共享内存区等。

多进程 Multiprocessing 和多线程 threading 类似，都是在 python 中用来并行运算的，Multiprocessing 用来弥补 threading 的一些劣势，比如 GIL。  
多核 / 多进程在同时间运行了多个任务使得程序运行更快，而多线程的运行时间甚至会比什么都不做的程序还要慢一点。  

多进程 Multiprocessing 和多线程 threading 的使用对比。  
```python
import multiprocessing as mp
import threading as td

def job(name):
    print(name)

if __name__ == '__main__':
	t1 = td.Thread(target=job,args=("thread",))
	p1 = mp.Process(target=job,args=("process",))
	t1.start()
	p1.start()
	t1.join()
	p1.join()
```

Queue 的功能是将每个核或线程的运算结果放在队里中，等到每个线程或核运行完毕后再从队列中取出结果，继续加载运算。  
多线程调用的函数不能有返回值，所以使用 Queue 存储多个线程运算的结果。  
```python
import multiprocessing as mp

def job(q):
    res=0
    for i in range(1000):
        res+=i+i**2+i**3
    q.put(res)

if __name__=='__main__':
    q = mp.Queue()
    p1 = mp.Process(target=job,args=(q,))
    p2 = mp.Process(target=job,args=(q,))
    p1.start()
    p2.start()
    p1.join()
    p2.join()
    res1 = q.get()
    res2 = q.get()
    print(res1+res2)
```

### Python 进程池
进程池就是将所要运行的东西放到池子里，Python 会自行解决多进程的问题。  
Pool 和之前的 Process 的不同点是丢向 Pool 的函数有返回值，而 Process 的没有返回值。  
```python
import multiprocessing as mp

def multicore():
	# 定义一个 Pool，默认调用是 CPU 的核数，传入 processes 参数可自定义 CPU 核数
    pool = mp.Pool(processes=3) # 定义 CPU 核数量为 3
    # map 放入迭代参数，返回多个结果
    res = pool.map(job, range(10))
    print(res)
    # apply_asnyc 只能放入一组参数，并返回一个结果
    res = pool.apply_async(job, (2,))
    # 用 get 获得结果
    print(res.get())
    # 通过迭代器实现 map 效果，i=0 时 apply 一次，i=1 时 apply 一次等等
    multi_res = [pool.apply_async(job, (i,)) for i in range(10)]
    # 从迭代器中取出
    print([res.get() for res in multi_res])
```

### 共享内存
只有用共享内存才能让 CPU 之间有交流。  
1、可以通过使用 Value 数据存储在一个共享的内存表中  
```python
import multiprocessing as mp

# d 和 i 参数用来设置数据类型的  
# d 表示一个双精浮点类型，i 表示一个带符号的整型
value1 = mp.Value('i', 0) 
value2 = mp.Value('d', 3.14)
```

2、可以通过 Array 和共享内存交互，来实现在进程之间共享数据。  
```python
# 和 numpy 不同，Array 只接受一维的，不能是多维的
array = mp.Array('i', [1, 2, 3, 4])
```

**参考数据形式**  

| Type code | C Type             | Python Type       | Minimum size in bytes |
| --------- | ------------------ | ----------------- | --------------------- |
| `'b'`     | signed char        | int               | 1                     |
| `'B'`     | unsigned char      | int               | 1                     |
| `'u'`     | Py_UNICODE         | Unicode character | 2                     |
| `'h'`     | signed short       | int               | 2                     |
| `'H'`     | unsigned short     | int               | 2                     |
| `'i'`     | signed int         | int               | 2                     |
| `'I'`     | unsigned int       | int               | 2                     |
| `'l'`     | signed long        | int               | 4                     |
| `'L'`     | unsigned long      | int               | 4                     |
| `'q'`     | signed long long   | int               | 8                     |
| `'Q'`     | unsigned long long | int               | 8                     |
| `'f'`     | float              | float             | 4                     |
| `'d'`     | double             | float             | 8                     |

### 进程锁
不同进程抢共享资源的问题，我们可以加进程锁来解决。  
```python
import multiprocessing as mp

def job(v, num, l):
    l.acquire() # 锁住
    for _ in range(5):
        time.sleep(0.1) 
        v.value += num # 获取共享内存
        print(v.value)
    l.release() # 释放

def multicore():
    l = mp.Lock() # 定义一个进程锁
    v = mp.Value('i', 0) # 定义共享内存
    p1 = mp.Process(target=job, args=(v,1,l)) # 需要将 lock 传入
    p2 = mp.Process(target=job, args=(v,3,l)) 
    p1.start()
    p2.start()
    p1.join()
    p2.join()

if __name__ == '__main__':
    multicore()
```
