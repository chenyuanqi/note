
### Go 文件处理
Go 的文件处理，重点在于文件而非目录或者通用的文件系统，特别是如何读写标准格式（如 XML 和 JSON 格式）的文件以及自定义的纯文本和二进制格式文件。

**JSON 文件的读写操作**  
Go 语言内建对 JSON 的支持，使用内置的 encoding/json 标准库，开发人员可以轻松使用Go程序生成和解析 JSON 格式的数据。
```go
package main
import (
    "encoding/json"
    "fmt"
    "os"
)
type Website struct {
    Name   string `xml:"name,attr"`
    Url    string
    Course []string
}
func main() {
    info := []Website{{"Golang", "http://c.biancheng.net/golang/", []string{"http://c.biancheng.net/cplus/", "http://c.biancheng.net/linux_tutorial/"}}, {"Java", "http://c.biancheng.net/java/", []string{"http://c.biancheng.net/socket/", "http://c.biancheng.net/python/"}}}
    // 创建文件
    filePtr, err := os.Create("info.json")
    if err != nil {
        fmt.Println("文件创建失败", err.Error())
        return
    }
    defer filePtr.Close()
    // 创建Json编码器
    encoder := json.NewEncoder(filePtr)
    // 写Json文件
    err = encoder.Encode(info)
    if err != nil {
        fmt.Println("编码错误", err.Error())
    } else {
        fmt.Println("编码成功")
    }

    // .................................................
    filePtr, err := os.Open("./info.json")
    if err != nil {
        fmt.Println("文件打开失败 [Err:%s]", err.Error())
        return
    }
    defer filePtr.Close()
    var info []Website
    // 创建json解码器
    decoder := json.NewDecoder(filePtr)
    // 读Json文件
    err = decoder.Decode(&info)
    if err != nil {
        fmt.Println("解码失败", err.Error())
    } else {
        fmt.Println("解码成功")
        fmt.Println(info)
    }
}
```

**纯文本文件的读写操作**  
Go 语言的 fmt 包中打印函数强大而灵活，写纯文本数据非常简单直接。
```go
package main
import (
    "bufio"
    "fmt"
    "os"
)
func main() {
    //创建一个新文件，写入内容
    filePath := "./output.txt"
    file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        fmt.Printf("打开文件错误= %v \n", err)
        return
    }
    //及时关闭
    defer file.Close()
    //写入内容
    str := "http://c.biancheng.net/golang/\n" // \n\r表示换行  txt文件要看到换行效果要用 \r\n
    //写入时，使用带缓存的 *Writer
    writer := bufio.NewWriter(file)
    for i := 0; i < 3; i++ {
        writer.WriteString(str)
    }
    //因为 writer 是带缓存的，因此在调用 WriterString 方法时，内容是先写入缓存的
    //所以要调用 flush方法，将缓存的数据真正写入到文件中。
    writer.Flush()

    // ..............................................
    //打开文件
    file, err := os.Open("./output.txt")
    if err != nil {
        fmt.Println("文件打开失败 = ", err)
    }
    //及时关闭 file 句柄，否则会有内存泄漏
    defer file.Close()
    //创建一个 *Reader ， 是带缓冲的
    reader := bufio.NewReader(file)
    for {
        str, err := reader.ReadString('\n') //读到一个换行就结束
        if err == io.EOF {                  //io.EOF 表示文件的末尾
            break
        }
        fmt.Print(str)
    }
    fmt.Println("文件读取结束...")
}
```

**buffer 读取文件**  
buffer 是缓冲器的意思，Go 语言要实现缓冲读取需要使用到 bufio 包。bufio 包本身包装了 io.Reader 和 io.Writer 对象，同时创建了另外的 Reader 和 Writer 对象，因此对于文本 I/O 来说，bufio 包提供了一定的便利性。

buffer 缓冲器的实现原理就是，将文件读取进缓冲（内存）之中，再次读取的时候就可以避免文件系统的 I/O 从而提高速度。同理在进行写操作时，先把文件写入缓冲（内存），然后由缓冲写入文件系统。  
bufio 和 io 包中有很多操作都是相似的，唯一不同的地方是 bufio 提供了一些缓冲的操作，如果对文件 I/O 操作比较频繁的，使用 bufio 包能够提高一定的性能。  
```go
// 使用 bufio 包写入文件
// 在 bufio 包中，有一个 Writer 结构体，而其相关的方法支持一些写入操作
//Writer 是一个空的结构体，一般需要使用 NewWriter 或者 NewWriterSize 来初始化一个结构体对象
type Writer struct {
        // contains filtered or unexported fields
}
//NewWriterSize 和 NewWriter 函数
//返回默认缓冲大小的 Writer 对象(默认是4096)
func NewWriter(w io.Writer) *Writer
//指定缓冲大小创建一个 Writer 对象
func NewWriterSize(w io.Writer, size int) *Writer
//Writer 对象相关的写入数据的方法
//把 p 中的内容写入 buffer，返回写入的字节数和错误信息。如果 nn < len(p)，返回错误信息中会包含为什么写入的数据比较短
func (b *Writer) Write(p []byte) (nn int, err error)
//将 buffer 中的数据写入 io.Writer
func (b *Writer) Flush() error
//以下三个方法可以直接写入到文件中
//写入单个字节
func (b *Writer) WriteByte(c byte) error
//写入单个 Unicode 指针返回写入字节数错误信息
func (b *Writer) WriteRune(r rune) (size int, err error)
//写入字符串并返回写入字节数和错误信息
func (b *Writer) WriteString(s string) (int, error)

// ..........................
package main
import (
    "bufio"
    "fmt"
    "os"
)
func main() {
    name := "demo.txt"
    content := "http://c.biancheng.net/golang/"
    fileObj, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        fmt.Println("文件打开失败", err)
    }
    defer fileObj.Close()
    writeObj := bufio.NewWriterSize(fileObj, 4096)
    //使用 Write 方法,需要使用 Writer 对象的 Flush 方法将 buffer 中的数据刷到磁盘
    buf := []byte(content)
    if _, err := writeObj.Write(buf); err == nil {
        if err := writeObj.Flush(); err != nil {
            panic(err)
        }
        fmt.Println("数据写入成功")
    }
}


// 使用 bufio 包读取文件
// bufio 包的相关的 Reader 函数方法
//首先定义了一个用来缓冲 io.Reader 对象的结构体，同时该结构体拥有以下相关的方法
type Reader struct {
}
//NewReader 函数用来返回一个默认大小 buffer 的 Reader 对象（默认大小是 4096） 等同于 NewReaderSize(rd,4096)
func NewReader(rd io.Reader) *Reader
//该函数返回一个指定大小 buffer（size 最小为 16）的 Reader 对象，如果 io.Reader 参数已经是一个足够大的 Reader，它将返回该 Reader
func NewReaderSize(rd io.Reader, size int) *Reader
//该方法返回从当前 buffer 中能被读到的字节数
func (b *Reader) Buffered() int
//Discard 方法跳过后续的 n 个字节的数据，返回跳过的字节数。如果 0 <= n <= b.Buffered()，该方法将不会从 io.Reader 中成功读取数据
func (b *Reader) Discard(n int) (discarded int, err error)
//Peekf 方法返回缓存的一个切片，该切片只包含缓存中的前 n 个字节的数据
func (b *Reader) Peek(n int) ([]byte, error)
//把 Reader 缓存对象中的数据读入到 []byte 类型的 p 中，并返回读取的字节数。读取成功，err 将返回空值
func (b *Reader) Read(p []byte) (n int, err error)
//返回单个字节，如果没有数据返回 err
func (b *Reader) ReadByte() (byte, error)
//该方法在 b 中读取 delimz 之前的所有数据，返回的切片是已读出的数据的引用，切片中的数据在下一次的读取操作之前是有效的。如果未找到 delim，将返回查找结果并返回 nil 空值。因为缓存的数据可能被下一次的读写操作修改，因此一般使用 ReadBytes 或者 ReadString，他们返回的都是数据拷贝
func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
//功能同 ReadSlice，返回数据的拷贝
func (b *Reader) ReadBytes(delim byte) ([]byte, error)
//功能同 ReadBytes，返回字符串
func (b *Reader) ReadString(delim byte) (string, error)
//该方法是一个低水平的读取方式，一般建议使用 ReadBytes('\n') 或 ReadString('\n')，或者使用一个 Scanner 来代替。ReadLine 通过调用 ReadSlice 方法实现，返回的也是缓存的切片，用于读取一行数据，不包括行尾标记（\n 或 \r\n）
func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
//读取单个 UTF-8 字符并返回一个 rune 和字节大小
func (b *Reader) ReadRune() (r rune, size int, err error)

// ............................
package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)
func main() {
    fileObj, err := os.Open("demo.txt")
    if err != nil {
        fmt.Println("文件打开失败：", err)
        return
    }
    defer fileObj.Close()
    //一个文件对象本身是实现了io.Reader的 使用bufio.NewReader去初始化一个Reader对象，存在buffer中的，读取一次就会被清空
    reader := bufio.NewReader(fileObj)
    buf := make([]byte, 1024)
    //读取 Reader 对象中的内容到 []byte 类型的 buf 中
    info, err := reader.Read(buf)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("读取的字节数:" + strconv.Itoa(info))
    //这里的buf是一个[]byte，因此如果需要只输出内容，仍然需要将文件内容的换行符替换掉
    fmt.Println("读取的文件内容:", string(buf))
}
```

**文件的写入、追加、读取、复制操作**  
Go 语言的 os 包下有一个 OpenFile 函数，其原型：func OpenFile(name string, flag int, perm FileMode) (file *File, err error)  
其中 name 是文件的文件名，如果不是在当前路径下运行需要加上具体路径；flag 是文件的处理参数，为 int 类型，根据系统的不同具体值可能有所不同，但是作用是相同的。一些常用的 flag 文件处理参数：
- O_RDONLY：只读模式打开文件；
- O_WRONLY：只写模式打开文件；
- O_RDWR：读写模式打开文件；
- O_APPEND：写操作时将数据附加到文件尾部（追加）；
- O_CREATE：如果不存在将创建一个新文件；
- O_EXCL：和 O_CREATE 配合使用，文件必须不存在，否则返回一个错误；
- O_SYNC：当进行一系列写操作时，每次都要等待上次的 I/O 操作完成再进行；
- O_TRUNC：如果可能，在打开时清空文件。

```go
package main
import (
    "bufio"
    "fmt"
    "os"
)
func main() {
    //创建一个新文件，写入内容 5 句 “http://c.biancheng.net/golang/”
    filePath := "e:/code/golang.txt"
    file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        fmt.Println("文件打开失败", err)
    }
    //及时关闭file句柄
    defer file.Close()
    //写入文件时，使用带缓存的 *Writer
    write := bufio.NewWriter(file)
    for i := 0; i < 5; i++ {
        write.WriteString("http://c.biancheng.net/golang/ \n")
    }
    //Flush将缓存的文件真正写入到文件中
    write.Flush()
}
```

**文件锁操作**  
使用 Go 语言开发一些程序的时候，往往出现多个进程同时操作同一份文件的情况，这很容易导致文件中的数据混乱。这时我们就需要采用一些手段来平衡这些冲突，文件锁（flock）应运而生。  
对于 flock，最常见的例子就是 Nginx，进程运行起来后就会把当前的 PID 写入这个文件，当然如果这个文件已经存在了，也就是前一个进程还没有退出，那么 Nginx 就不会重新启动，所以 flock 还可以用来检测进程是否存在。  
flock 是对于整个文件的建议性锁。也就是说，如果一个进程在一个文件（inode）上放了锁，其它进程是可以知道的（建议性锁不强求进程遵守）。最棒的一点是，它的第一个参数是文件描述符，在此文件描述符关闭时，锁会自动释放。而当进程终止时，所有的文件描述符均会被关闭。所以很多时候就不用考虑类似原子锁解锁的事情。
```go
package main
import (
    "fmt"
    "os"
    "sync"
    "syscall"
    "time"
)
//文件锁
type FileLock struct {
    dir string
    f   *os.File
}
func New(dir string) *FileLock {
    return &FileLock{
        dir: dir,
    }
}
//加锁
func (l *FileLock) Lock() error {
    f, err := os.Open(l.dir)
    if err != nil {
        return err
    }
    l.f = f
    err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
    if err != nil {
        return fmt.Errorf("cannot flock directory %s - %s", l.dir, err)
    }
    return nil
}
//释放锁
func (l *FileLock) Unlock() error {
    defer l.f.Close()
    return syscall.Flock(int(l.f.Fd()), syscall.LOCK_UN)
}
func main() {
    // Windows 系统不支持 pid 锁，需要在 Linux 或 Mac 系统下才能正常运行
    test_file_path, _ := os.Getwd()
    locked_file := test_file_path
    wg := sync.WaitGroup{}
    // 同时启动 10 个 goroutinue，但在程序运行过程中，只有一个 goroutine 能获得文件锁（flock）。其它的 goroutinue 在获取不到 flock 后，会抛出异常的信息。这样即可达到同一文件在指定的周期内只允许一个进程访问的效果
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(num int) {
            flock := New(locked_file)
            err := flock.Lock()
            if err != nil {
                wg.Done()
                fmt.Println(err.Error())
                return
            }
            fmt.Printf("output : %d\n", num)
            wg.Done()
        }(i)
    }
    wg.Wait()
    time.Sleep(2 * time.Second)
}
```
flock 属于建议性锁，不具备强制性。一个进程使用 flock 将文件锁住，另一个进程可以直接操作正在被锁的文件，修改文件中的数据，原因在于 flock 只是用于检测文件是否被加锁，针对文件已经被加锁，另一个进程写入数据的情况，内核不会阻止这个进程的写入操作，也就是建议性锁的内核处理策略。  
flock 主要三种操作类型：
- LOCK_SH：共享锁，多个进程可以使用同一把锁，常被用作读共享锁；
- LOCK_EX：排他锁，同时只允许一个进程使用，常被用作写锁；
- LOCK_UN：释放锁。

进程使用 flock 尝试锁文件时，如果文件已经被其他进程锁住，进程会被阻塞直到锁被释放掉，或者在调用 flock 的时候，采用 LOCK_NB 参数。在尝试锁住该文件的时候，发现已经被其他服务锁住，会返回错误，错误码为 EWOULDBLOCK。

flock 锁的释放非常具有特色，即可调用 LOCK_UN 参数来释放文件锁，也可以通过关闭 fd 的方式来释放文件锁（flock 的第一个参数是 fd），意味着 flock 会随着进程的关闭而被自动释放掉。

