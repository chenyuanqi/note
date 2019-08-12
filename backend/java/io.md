
### Java IO
IO 一直是软件开发中的核心部分之一，伴随着海量数据增长和分布式系统的发展，IO 扩展能力愈发重要。  

Java IO 流的 40 多个类都是从如下 4 个抽象类基类中派生出来的：  
> InputStream/Reader: 所有的输入流的基类，前者是字节输入流，后者是字符输入流  
> OutputStream/Writer: 所有输出流的基类，前者是字节输出流，后者是字符输出流  

**IO 按操作方式分类**  
![java-io-01](./image/java-io-01.png)  

**IO 按操作对象分类**  
![java-io-02](./image/java-io-02.png)  

一些基本概念与认识  
> 区分同步或异步（synchronous/asynchronous）。简单来说，同步是一种可靠的有序运行机制，当我们进行同步操作时，后续的任务是等待当前调用返回，才会进行下一步； 而异步则相反，其他任务不需要等待当前调用返回，通常依靠事件、回调等机制来实现任务间次序关系。  
> 区分阻塞与非阻塞（blocking/non-blocking）。在进行阻塞操作时，当前线程会处于阻塞状态，无法从事其他任务，只有当条件就绪才能继续，比如 ServerSocket 新连接建立 完毕，或数据读取、写入操作完成；而非阻塞则是不管 IO 操作是否结束，直接返回，相应操作在后台继续处理。  
> 不能一概而论认为同步或阻塞就是低效，具体还要看应用和系统特征。  
> 
> IO 不仅仅是对文件的操作，网络编程中比如 Socket 通信是典型的 IO 操作目标。  
> 
> 输入流、输出流（InputStream/OutputStream）是用于读取或写入字节的，例如操作图片文件。  
> Reader/Writer 则是用于操作字符，增加了字符编解码等功能，适用于类似从文件中读取或者写入文本信息。本质上计算机操作的都是字节，不管是网络通信还是文件读取，Reader/Writer 相当于构建了应用逻辑和原始数据之间的桥梁。  
> BuferedOutputStream 等带缓冲区的实现，可以避免频繁的磁盘读写，进而提高 IO 处理效率。这种设计利用了缓冲区，将批量数据进行一次操作，但在使用中千万别忘了 fush。

```java
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;

// 创建字节输入流对象
FileInputStream fis = new FileInputStream("test.txt");
// 创建字节输出流对象
FileOutputStream fos = new FileOutputStream("copy.txt");

// 读写文件 1
// byte[] bys = new byte[1024 * 1024]; // 这里的数据一般是 1024 或者其整数倍
// int len = 0;
// while ((len = fis.read(bys)) != -1) {
//     fos.write(bys, 0, len);
// }
// 
// 读写文件 2
int by = 0;
while ((by = fis.read()) != -1) {
    fos.write(by);
}

// 释放资源
fos.close();
fis.close();
```

[Java IO 体系学习](https://blog.csdn.net/nightcurtis/article/details/51324105)  

### Java NIO
Java NIO 是 java 1.4 之后新出的一套 IO 接口，NIO 中的 N 可以理解为 Non-blocking，不单纯是 New。

NIO 与 IO的区别：  
IO 是面向流的，NIO 是面向缓冲区的；IO 流是阻塞的，NIO 流是不阻塞的；NIO 有选择器，而 IO 没有。  

NIO 的读写方式：  
从通道进行数据读取 ：创建一个缓冲区，然后请求通道读取数据。  
从通道进行数据写入 ：创建一个缓冲区，填充数据，并要求通道写入数据。  

NIO 的主要组成部分：  
> Bufer，高效的数据容器，除了布尔类型，所有原始数据类型都有相应的 Bufer 实现，Buffer 本质上就是一块内存区，我们从 Channel 中读取数据到 buffers 里，从 Buffer 把数据写入到 Channels。一个 Buffer 有三个属性是必须掌握的，分别是：capacity 容量、position 位置、limit 限制。    
> Channel，类似在 Linux 之类的操作系统上看到的文件描述符，是 NIO 中被用来支持批量式 IO 操作的一种抽象。File 或者 Socket，通常被认为是比较高层次的抽象，而 Channel 则是更加操作系统底层的一种抽象，这也使得 NIO 得以充分利用现代操作系统底层机制，获得特定场景的性能优化，例如，DMA（Direct Memory Access）等。不同层次的抽象是相互关联的，我们可以通过 Socket 获取 Channel，反之亦然。
  
> Selector，是 NIO 实现多路复用的基础，它提供了一种高效的机制，可以检测到注册在 Selector 上的多个 Channel 中，是否有 Channel 处于就绪状态，进而实现了单线程对 多 Channel 的高效管理。Selector同样是基于底层操作系统机制，不同模式、不同版本都存在区别。  
> Chartset，提供 Unicode 字符串定义，NIO 也提供了相应的编解码器等。  

[Java NIO 体系学习](https://blog.csdn.net/a953713428/article/details/64907250)  

### Java AIO
AIO 是异步 IO 的缩写。  
虽然 NIO 在网络操作中，提供了非阻塞的方法，但是 NIO 的 IO 行为还是同步的。对于 NIO 来说，我们的业务线程是在 IO 操作准备好时，得到通知，接着就由这个线程自行进行 IO 操作，IO 操作本身是同步的。  
但是，对于 AIO 来说，则更加进了一步，它不是在 IO 准备好时再通知线程，而是在 IO 操作已经完成后，再给线程发出通知。因此 AIO 是不会阻塞的，此时我们的业务逻辑将变成一个回调函数，等待 IO 操作完成后，由系统自动触发。

[Java AIO 体系学习](https://blog.csdn.net/x_i_y_u_e/article/details/52223406)
