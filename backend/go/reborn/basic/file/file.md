
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

### 常用文件操作
UNIX 的一个基础设计就是"万物皆文件"(everything is a file)。我们不必知道操作系统的设备驱动把什么映射给了一个文件描述符，操作系统为设备提供了文件格式的接口。  
Go 语言中的 reader 和 writer 接口也类似。我们只需简单的读写字节，不必知道 reader 的数据来自哪里，也不必知道 writer 将数据发送到哪里。你可以在 /dev 下查看可用的设备，有些可能需要较高的权限才能访问。

**创建空文件**  
```go
package main

import (
    "log"
    "os"
)

var (
    newFile *os.File
    err     error
)

func main() {
    newFile, err = os.Create("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    log.Println(newFile)
    newFile.Close()
}
```

**Truncate裁剪文件**  
```go
package main

import (
    "log"
    "os"
)

func main() {
    // 裁剪一个文件到100个字节。
    // 如果文件本来就少于100个字节，则文件中原始内容得以保留，剩余的字节以null字节填充。
    // 如果文件本来超过100个字节，则超过的字节会被抛弃。
    // 这样我们总是得到精确的100个字节的文件。
    // 传入0则会清空文件。

    err := os.Truncate("test.txt", 100)
    if err != nil {
        log.Fatal(err)
    }
}
```

**获取文件信息**  
```go
package main

import (
    "fmt"
    "log"
    "os"
)

var (
    fileInfo os.FileInfo
    err      error
)

func main() {
    // 如果文件不存在，则返回错误
    fileInfo, err = os.Stat("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("File name:", fileInfo.Name())
    fmt.Println("Size in bytes:", fileInfo.Size())
    fmt.Println("Permissions:", fileInfo.Mode())
    fmt.Println("Last modified:", fileInfo.ModTime())
    fmt.Println("Is Directory: ", fileInfo.IsDir())
    fmt.Printf("System interface type: %T\n", fileInfo.Sys())
    fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
}
```

**重命名和移动**  
```go
package main

import (
    "log"
    "os"
)

func main() {
    originalPath := "test.txt"
    newPath := "test2.txt"
    err := os.Rename(originalPath, newPath)
    if err != nil {
        log.Fatal(err)
    }
}
```

**删除文件**  
```go
package main

import (
    "log"
    "os"
)

func main() {
    err := os.Remove("test.txt")
    if err != nil {
        log.Fatal(err)
    }
}
```

**打开和关闭文件**  
```go
package main

import (
    "log"
    "os"
)

func main() {
    // 简单地以只读的方式打开。下面的例子会介绍读写的例子。
    file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    file.Close()

    // OpenFile提供更多的选项。
    // 最后一个参数是权限模式permission mode
    // 第二个是打开时的属性    
    file, err = os.OpenFile("test.txt", os.O_APPEND, 0666)
    if err != nil {
        log.Fatal(err)
    }
    file.Close()

    // 下面的属性可以单独使用，也可以组合使用。
    // 组合使用时可以使用 OR 操作设置 OpenFile的第二个参数，例如：
    // os.O_CREATE|os.O_APPEND
    // 或者 os.O_CREATE|os.O_TRUNC|os.O_WRONLY

    // os.O_RDONLY // 只读
    // os.O_WRONLY // 只写
    // os.O_RDWR // 读写
    // os.O_APPEND // 往文件中添建（Append）
    // os.O_CREATE // 如果文件不存在则先创建
    // os.O_TRUNC // 文件打开时裁剪文件
    // os.O_EXCL // 和O_CREATE一起使用，文件不能存在
    // os.O_SYNC // 以同步I/O的方式打开
}
```

**检查文件是否存在**  
```go
package main

import (
    "log"
    "os"
)

var (
    fileInfo *os.FileInfo
    err      error
)

func main() {
    // 文件不存在则返回error
    fileInfo, err := os.Stat("test.txt")
    if err != nil {
        if os.IsNotExist(err) {
            log.Fatal("File does not exist.")
        }
    }
    log.Println("File does exist. File information:")
    log.Println(fileInfo)
}
```

**检查读写权限**  
```go

package main

import (
    "log"
    "os"
)

func main() {
    // 这个例子测试写权限，如果没有写权限则返回error。
    // 注意文件不存在也会返回error，需要检查error的信息来获取到底是哪个错误导致。
    file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
    if err != nil {
        if os.IsPermission(err) {
            log.Println("Error: Write permission denied.")
        }
    }
    file.Close()

    // 测试读权限
    file, err = os.OpenFile("test.txt", os.O_RDONLY, 0666)
    if err != nil {
        if os.IsPermission(err) {
            log.Println("Error: Read permission denied.")
        }
    }
    file.Close()
}
```

**改变权限、拥有者、时间戳**  
```go
package main

import (
    "log"
    "os"
    "time"
)

func main() {
    // 使用Linux风格改变文件权限
    err := os.Chmod("test.txt", 0777)
    if err != nil {
        log.Println(err)
    }

    // 改变文件所有者
    err = os.Chown("test.txt", os.Getuid(), os.Getgid())
    if err != nil {
        log.Println(err)
    }

    // 改变时间戳
    twoDaysFromNow := time.Now().Add(48 * time.Hour)
    lastAccessTime := twoDaysFromNow
    lastModifyTime := twoDaysFromNow
    err = os.Chtimes("test.txt", lastAccessTime, lastModifyTime)
    if err != nil {
        log.Println(err)
    }
}
```

**创建硬链接和软链接**  
一个普通的文件是一个指向硬盘的inode的地方。硬链接创建一个新的指针指向同一个地方。只有所有的链接被删除后文件才会被删除。硬链接只在相同的文件系统中才工作。你可以认为一个硬链接是一个正常的链接。

symbolic link，又叫软连接，和硬链接有点不一样，它不直接指向硬盘中的相同的地方，而是通过名字引用其它文件。他们可以指向不同的文件系统中的不同文件。并不是所有的操作系统都支持软链接。
```go
package main

import (
    "os"
    "log"
    "fmt"
)

func main() {
    // 创建一个硬链接。
    // 创建后同一个文件内容会有两个文件名，改变一个文件的内容会影响另一个。
    // 删除和重命名不会影响另一个。
    err := os.Link("original.txt", "original_also.txt")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("creating sym")
    // Create a symlink
    err = os.Symlink("original.txt", "original_sym.txt")
    if err != nil {
        log.Fatal(err)
    }

    // Lstat返回一个文件的信息，但是当文件是一个软链接时，它返回软链接的信息，而不是引用的文件的信息。
    // Symlink在Windows中不工作。
    fileInfo, err := os.Lstat("original_sym.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Link info: %+v", fileInfo)

    //改变软链接的拥有者不会影响原始文件。
    err = os.Lchown("original_sym.txt", os.Getuid(), os.Getgid())
    if err != nil {
        log.Fatal(err)
    }
}
```

**复制文件**  
```go
package main

import (
    "os"
    "log"
    "io"
)

func main() {
    // 打开原始文件
    originalFile, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer originalFile.Close()

    // 创建新的文件作为目标文件
    newFile, err := os.Create("test_copy.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer newFile.Close()

    // 从源中复制字节到目标文件
    bytesWritten, err := io.Copy(newFile, originalFile)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Copied %d bytes.", bytesWritten)

    // 将文件内容flush到硬盘中
    err = newFile.Sync()
    if err != nil {
        log.Fatal(err)
    }
}
```

**跳转到文件指定位置(Seek)**  
```go
package main

import (
    "os"
    "fmt"
    "log"
)

func main() {
    file, _ := os.Open("test.txt")
    defer file.Close()

    // 偏离位置，可以是正数也可以是负数
    var offset int64 = 5

    // 用来计算offset的初始位置
    // 0 = 文件开始位置
    // 1 = 当前位置
    // 2 = 文件结尾处
    var whence int = 0
    newPosition, err := file.Seek(offset, whence)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Just moved to 5:", newPosition)

    // 从当前位置回退两个字节
    newPosition, err = file.Seek(-2, 1)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Just moved back two:", newPosition)

    // 使用下面的技巧得到当前的位置
    currentPosition, err := file.Seek(0, 1)
    fmt.Println("Current position:", currentPosition)

    // 转到文件开始处
    newPosition, err = file.Seek(0, 0)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Position after seeking 0,0:", newPosition)
}
```

**写文件**  
可以使用os包写入一个打开的文件。因为Go可执行包是静态链接的可执行文件，你import的每一个包都会增加你的可执行文件的大小。其它的包如io、｀ioutil｀、｀bufio｀提供了一些方法，但是它们不是必须的。
```go
package main

import (
    "os"
    "log"
)

func main() {
    // 可写方式打开文件
    file, err := os.OpenFile(
        "test.txt",
        os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
        0666,
    )
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // 写字节到文件中
    byteSlice := []byte("Bytes!\n")
    bytesWritten, err := file.Write(byteSlice)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Wrote %d bytes.\n", bytesWritten)
}
```

**快写文件**  
ioutil包有一个非常有用的方法WriteFile()可以处理创建或者打开文件、写入字节切片和关闭文件一系列的操作。如果你需要简洁快速地写字节切片到文件中，你可以使用它。
```go
package main

import (
    "io/ioutil"
    "log"
)

func main() {
    err := ioutil.WriteFile("test.txt", []byte("Hi\n"), 0666)
    if err != nil {
        log.Fatal(err)
    }
}
```

**使用缓存写**  
bufio包提供了带缓存功能的writer，所以你可以在写字节到硬盘前使用内存缓存。当你处理很多的数据很有用，因为它可以节省操作硬盘I/O的时间。在其它一些情况下它也很有用，比如你每次写一个字节，把它们攒在内存缓存中，然后一次写入到硬盘中，减少硬盘的磨损以及提升性能。
```go
package main

import (
    "log"
    "os"
    "bufio"
)

func main() {
    // 打开文件，只写
    file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // 为这个文件创建buffered writer
    bufferedWriter := bufio.NewWriter(file)

    // 写字节到buffer
    bytesWritten, err := bufferedWriter.Write(
        []byte{65, 66, 67},
    )
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Bytes written: %d\n", bytesWritten)

    // 写字符串到buffer
    // 也可以使用 WriteRune() 和 WriteByte()   
    bytesWritten, err = bufferedWriter.WriteString(
        "Buffered string\n",
    )
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Bytes written: %d\n", bytesWritten)

    // 检查缓存中的字节数
    unflushedBufferSize := bufferedWriter.Buffered()
    log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

    // 还有多少字节可用（未使用的缓存大小）
    bytesAvailable := bufferedWriter.Available()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Available buffer: %d\n", bytesAvailable)

    // 写内存buffer到硬盘
    bufferedWriter.Flush()

    // 丢弃还没有flush的缓存的内容，清除错误并把它的输出传给参数中的writer
    // 当你想将缓存传给另外一个writer时有用
    bufferedWriter.Reset(bufferedWriter)

    bytesAvailable = bufferedWriter.Available()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Available buffer: %d\n", bytesAvailable)

    // 重新设置缓存的大小。
    // 第一个参数是缓存应该输出到哪里，这个例子中我们使用相同的writer。
    // 如果我们设置的新的大小小于第一个参数writer的缓存大小， 比如10，我们不会得到一个10字节大小的缓存，
    // 而是writer的原始大小的缓存，默认是4096。
    // 它的功能主要还是为了扩容。
    bufferedWriter = bufio.NewWriterSize(
        bufferedWriter,
        8000,
    )

    // resize后检查缓存的大小
    bytesAvailable = bufferedWriter.Available()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Available buffer: %d\n", bytesAvailable)
}
```

**读取最多N个字节**  
os.File提供了文件操作的基本功能， 而io、ioutil、bufio提供了额外的辅助函数。
```go
package main

import (
    "os"
    "log"
)

func main() {
    // 打开文件，只读
    file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // 从文件中读取len(b)字节的文件。
    // 返回0字节意味着读取到文件尾了
    // 读取到文件会返回io.EOF的error
    byteSlice := make([]byte, 16)
    bytesRead, err := file.Read(byteSlice)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Number of bytes read: %d\n", bytesRead)
    log.Printf("Data read: %s\n", byteSlice)
}
```

**读取正好N个字节**  
```go
package main

import (
    "os"
    "log"
    "io"
)

func main() {
    // Open file for reading
    file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }

    // file.Read()可以读取一个小文件到大的byte slice中，
    // 但是io.ReadFull()在文件的字节数小于byte slice字节数的时候会返回错误
    byteSlice := make([]byte, 2)
    numBytesRead, err := io.ReadFull(file, byteSlice)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Number of bytes read: %d\n", numBytesRead)
    log.Printf("Data read: %s\n", byteSlice)
}
```

**读取至少N个字节**  
```go
package main

import (
    "os"
    "log"
    "io"
)

func main() {
    // 打开文件，只读
    file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }

    byteSlice := make([]byte, 512)
    minBytes := 8
    // io.ReadAtLeast()在不能得到最小的字节的时候会返回错误，但会把已读的文件保留
    numBytesRead, err := io.ReadAtLeast(file, byteSlice, minBytes)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Number of bytes read: %d\n", numBytesRead)
    log.Printf("Data read: %s\n", byteSlice)
}
```

**读取全部字节**  
```go
package main

import (
    "os"
    "log"
    "fmt"
    "io/ioutil"
)

func main() {
    file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }

    // os.File.Read(), io.ReadFull() 和
    // io.ReadAtLeast() 在读取之前都需要一个固定大小的byte slice。
    // 但ioutil.ReadAll()会读取reader(这个例子中是file)的每一个字节，然后把字节slice返回。
    data, err := ioutil.ReadAll(file)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Data as hex: %x\n", data)
    fmt.Printf("Data as string: %s\n", data)
    fmt.Println("Number of bytes read:", len(data))
}
```

**快读到内存**  
```go
package main

import (
    "log"
    "io/ioutil"
)

func main() {
    // 读取文件到byte slice中
    data, err := ioutil.ReadFile("test.txt")
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Data read: %s\n", data)
}
```

**使用缓存读**  
有缓存写也有缓存读。缓存reader会把一些内容缓存在内存中。它会提供比os.File和io.Reader更多的函数,缺省的缓存大小是4096，最小缓存是16。
```go
package main

import (
    "os"
    "log"
    "bufio"
    "fmt"
)

func main() {
    // 打开文件，创建buffered reader
    file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    bufferedReader := bufio.NewReader(file)

    // 得到字节，当前指针不变
    byteSlice := make([]byte, 5)
    byteSlice, err = bufferedReader.Peek(5)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Peeked at 5 bytes: %s\n", byteSlice)

    // 读取，指针同时移动
    numBytesRead, err := bufferedReader.Read(byteSlice)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Read %d bytes: %s\n", numBytesRead, byteSlice)

    // 读取一个字节, 如果读取不成功会返回Error
    myByte, err := bufferedReader.ReadByte()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Read 1 byte: %c\n", myByte)     

    // 读取到分隔符，包含分隔符，返回byte slice
    dataBytes, err := bufferedReader.ReadBytes('\n')
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Read bytes: %s\n", dataBytes)           

    // 读取到分隔符，包含分隔符，返回字符串
    dataString, err := bufferedReader.ReadString('\n')
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Read string: %s\n", dataString)     

    //这个例子读取了很多行，所以test.txt应该包含多行文本才不至于出错
}
```

**使用 scanner**  
Scanner是bufio包下的类型,在处理文件中以分隔符分隔的文本时很有用。通常我们使用换行符作为分隔符将文件内容分成多行。在CSV文件中，逗号一般作为分隔符。os.File文件可以被包装成bufio.Scanner，它就像一个缓存reader。我们会调用Scan()方法去读取下一个分隔符，使用Text()或者Bytes()获取读取的数据。

分隔符可以不是一个简单的字节或者字符，有一个特殊的方法可以实现分隔符的功能，以及将指针移动多少，返回什么数据。如果没有定制的SplitFunc提供，缺省的ScanLines会使用newline字符作为分隔符，其它的分隔函数还包括ScanRunes和ScanWords,皆在bufio包中。
```go
// To define your own split function, match this fingerprint
type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)

// Returning (0, nil, nil) will tell the scanner
// to scan again, but with a bigger buffer because
// it wasn't enough data to reach the delimiter

// 为一个文件创建了bufio.Scanner，并按照单词逐个读取
package main

import (
    "os"
    "log"
    "fmt"
    "bufio"
)

func main() {
    file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    scanner := bufio.NewScanner(file)

    // 缺省的分隔函数是bufio.ScanLines,我们这里使用ScanWords。
    // 也可以定制一个SplitFunc类型的分隔函数
    scanner.Split(bufio.ScanWords)

    // scan下一个token.
    success := scanner.Scan()
    if success == false {
        // 出现错误或者EOF是返回Error
        err = scanner.Err()
        if err == nil {
            log.Println("Scan completed and reached EOF")
        } else {
            log.Fatal(err)
        }
    }

    // 得到数据，Bytes() 或者 Text()
    fmt.Println("First word found:", scanner.Text())

    // 再次调用scanner.Scan()发现下一个token
}
```

**打包(zip) 文件**  
```go
// This example uses zip but standard library
// also supports tar archives
package main

import (
    "archive/zip"
    "log"
    "os"
)

func main() {
    // 创建一个打包文件
    outFile, err := os.Create("test.zip")
    if err != nil {
        log.Fatal(err)
    }
    defer outFile.Close()

    // 创建zip writer
    zipWriter := zip.NewWriter(outFile)


    // 往打包文件中写文件。
    // 这里我们使用硬编码的内容，你可以遍历一个文件夹，把文件夹下的文件以及它们的内容写入到这个打包文件中。
    var filesToArchive = []struct {
        Name, Body string
    } {
        {"test.txt", "String contents of file"},
        {"test2.txt", "\x61\x62\x63\n"},
    }

    // 下面将要打包的内容写入到打包文件中，依次写入。
    for _, file := range filesToArchive {
            fileWriter, err := zipWriter.Create(file.Name)
            if err != nil {
                    log.Fatal(err)
            }
            _, err = fileWriter.Write([]byte(file.Body))
            if err != nil {
                    log.Fatal(err)
            }
    }

    // 清理
    err = zipWriter.Close()
    if err != nil {
            log.Fatal(err)
    }
}
```

**抽取(unzip) 文件**  
```go
// This example uses zip but standard library
// also supports tar archives
package main

import (
    "archive/zip"
    "log"
    "io"
    "os"
    "path/filepath"
)

func main() {
    zipReader, err := zip.OpenReader("test.zip")
    if err != nil {
        log.Fatal(err)
    }
    defer zipReader.Close()

    // 遍历打包文件中的每一文件/文件夹
    for _, file := range zipReader.Reader.File {
        // 打包文件中的文件就像普通的一个文件对象一样
        zippedFile, err := file.Open()
        if err != nil {
            log.Fatal(err)
        }
        defer zippedFile.Close()

        // 指定抽取的文件名。
        // 你可以指定全路径名或者一个前缀，这样可以把它们放在不同的文件夹中。
        // 我们这个例子使用打包文件中相同的文件名。
        targetDir := "./"
        extractedFilePath := filepath.Join(
            targetDir,
            file.Name,
        )

        // 抽取项目或者创建文件夹
        if file.FileInfo().IsDir() {
            // 创建文件夹并设置同样的权限
            log.Println("Creating directory:", extractedFilePath)
            os.MkdirAll(extractedFilePath, file.Mode())
        } else {
            //抽取正常的文件
            log.Println("Extracting file:", file.Name)

            outputFile, err := os.OpenFile(
                extractedFilePath,
                os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
                file.Mode(),
            )
            if err != nil {
                log.Fatal(err)
            }
            defer outputFile.Close()

            // 通过io.Copy简洁地复制文件内容
            _, err = io.Copy(outputFile, zippedFile)
            if err != nil {
                log.Fatal(err)
            }
        }
    }
}
```

**压缩文件**  
```go
// 这个例子中使用gzip压缩格式，标准库还支持zlib, bz2, flate, lzw
package main

import (
    "os"
    "compress/gzip"
    "log"
)

func main() {
    outputFile, err := os.Create("test.txt.gz")
    if err != nil {
        log.Fatal(err)
    }

    gzipWriter := gzip.NewWriter(outputFile)
    defer gzipWriter.Close()

    // 当我们写如到gizp writer数据时，它会依次压缩数据并写入到底层的文件中。
    // 我们不必关心它是如何压缩的，还是像普通的writer一样操作即可。
    _, err = gzipWriter.Write([]byte("Gophers rule!\n"))
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Compressed data written to file.")
}
```

**解压缩文件**  
```go
// 这个例子中使用gzip压缩格式，标准库还支持zlib, bz2, flate, lzw
package main

import (
    "compress/gzip"
    "log"
    "io"
    "os"
)

func main() {
    // 打开一个gzip文件。
    // 文件是一个reader,但是我们可以使用各种数据源，比如web服务器返回的gzipped内容，
    // 它的内容不是一个文件，而是一个内存流
    gzipFile, err := os.Open("test.txt.gz")
    if err != nil {
        log.Fatal(err)
    }

    gzipReader, err := gzip.NewReader(gzipFile)
    if err != nil {
        log.Fatal(err)
    }
    defer gzipReader.Close()

    // 解压缩到一个writer,它是一个file writer
    outfileWriter, err := os.Create("unzipped.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer outfileWriter.Close()

    // 复制内容
    _, err = io.Copy(outfileWriter, gzipReader)
    if err != nil {
        log.Fatal(err)
    }
}
```

**临时文件和目录**  
ioutil提供了两个函数: TempDir() 和 TempFile()。使用完毕后，调用者负责删除这些临时文件和文件夹。有一点好处就是当你传递一个空字符串作为文件夹名的时候，它会在操作系统的临时文件夹中创建这些项目（/tmp on Linux）。os.TempDir()返回当前操作系统的临时文件夹。
```go
package main

import (
     "os"
     "io/ioutil"
     "log"
     "fmt"
)

func main() {
     // 在系统临时文件夹中创建一个临时文件夹
     tempDirPath, err := ioutil.TempDir("", "myTempDir")
     if err != nil {
          log.Fatal(err)
     }
     fmt.Println("Temp dir created:", tempDirPath)

     // 在临时文件夹中创建临时文件
     tempFile, err := ioutil.TempFile(tempDirPath, "myTempFile.txt")
     if err != nil {
          log.Fatal(err)
     }
     fmt.Println("Temp file created:", tempFile.Name())

     // ... 做一些操作 ...

     // 关闭文件
     err = tempFile.Close()
     if err != nil {
        log.Fatal(err)
    }

    // 删除我们创建的资源
     err = os.Remove(tempFile.Name())
     if err != nil {
        log.Fatal(err)
    }
     err = os.Remove(tempDirPath)
     if err != nil {
        log.Fatal(err)
    }
}
```

**通过HTTP下载文件**  
```go
package main

import (
     "os"
     "io"
     "log"
     "net/http"
)

func main() {
     newFile, err := os.Create("devdungeon.html")
     if err != nil {
          log.Fatal(err)
     }
     defer newFile.Close()

     url := "http://www.devdungeon.com/archive"
     response, err := http.Get(url)
     defer response.Body.Close()

     // 将HTTP response Body中的内容写入到文件
     // Body满足reader接口，因此我们可以使用ioutil.Copy
     numBytesWritten, err := io.Copy(newFile, response.Body)
     if err != nil {
          log.Fatal(err)
     }
     log.Printf("Downloaded %d byte file.\n", numBytesWritten)
}
```

**哈希和摘要**
```go
package main

import (
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
    "log"
    "fmt"
    "io/ioutil"
)

func main() {
    // 得到文件内容
    data, err := ioutil.ReadFile("test.txt")
    if err != nil {
        log.Fatal(err)
    }

    // 计算Hash
    fmt.Printf("Md5: %x\n\n", md5.Sum(data))
    fmt.Printf("Sha1: %x\n\n", sha1.Sum(data))
    fmt.Printf("Sha256: %x\n\n", sha256.Sum256(data))
    fmt.Printf("Sha512: %x\n\n", sha512.Sum512(data))
}
```
上面的例子复制整个文件内容到内存中，传递给hash函数。另一个方式是创建一个hash writer, 使用Write、WriteString、Copy将数据传给它。下面的例子使用 md5 hash,但你可以使用其它的Writer。
```go
package main

import (
    "crypto/md5"
    "log"
    "fmt"
    "io"
    "os"
)

func main() {
    file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    //创建一个新的hasher,满足writer接口
    hasher := md5.New()
    _, err = io.Copy(hasher, file)
    if err != nil {
        log.Fatal(err)
    }

    // 计算hash并打印结果。
    // 传递 nil 作为参数，因为我们不通参数传递数据，而是通过writer接口。
    sum := hasher.Sum(nil)
    fmt.Printf("Md5 checksum: %x\n", sum)
}
```

