
### Go 文件读写
在 Go 语言中，文件使用指向 os.File 类型的指针来表示的，也叫做文件句柄。  
```golang
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

func main() {
	// 变量 inputFile 是 *os.File 类型的
    inputFile, inputError := os.Open("input.dat")
    if inputError != nil {
        fmt.Printf("An error occurred on opening the inputfile\n" +
            "Does the file exist?\n" +
            "Have you got acces to it?\n")
        return // exit the function on error
    }
    defer inputFile.Close()

    // 使用 bufio.NewReader 来获得一个读取器变量
    inputReader := bufio.NewReader(inputFile)
    for {
        inputString, readerError := inputReader.ReadString('\n')
        fmt.Printf("The input was: %s", inputString)
        // 一旦读取到文件末尾，变量 readerError 的值将变成非空（事实上，其值为常量 io.EOF）
        if readerError == io.EOF {
            return
        }      
    }

    // 将整个文件的内容读到一个字符串里
    inputFile := "products.txt"
    outputFile := "products_copy.txt"
    buf, err := ioutil.ReadFile(inputFile) // 第一个返回值的类型是 []byte，里面存放读取到的内容
    if err != nil {
        fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
        // panic(err.Error())
    }
    fmt.Printf("%s\n", string(buf))
    err = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex
    if err != nil {
        panic(err.Error())
    }

    // 带缓冲的读取
    buf := make([]byte, 1024)
	// ...
	n, err := inputReader.Read(buf) // 变量 n 的值表示读取到的字节数
	if (n == 0) { 
		break
	}

    // 按列读取文件中的数据
    file, err := os.Open("products2.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    var col1, col2, col3 []string
    for {
        var v1, v2, v3 string
        _, err := fmt.Fscanln(file, &v1, &v2, &v3)
        // scans until newline
        if err != nil {
            break
        }
        col1 = append(col1, v1)
        col2 = append(col2, v2)
        col3 = append(col3, v3)
    }
    fmt.Println(col1)
    fmt.Println(col2)
    fmt.Println(col3)

    // path 包里包含一个子包叫 filepath，这个子包提供了跨平台的函数，用于处理文件名和路径
    import "path/filepath"
    filename := filepath.Base(path)
}
```

**写文件**  
```golang
package main

import (
	"os"
	"bufio"
	"fmt"
)

func main () {
	// var outputWriter *bufio.Writer
	// var outputFile *os.File
	// var outputError os.Error
	// var outputString string
	// 以只写模式打开文件 output.dat，如果文件不存在则自动创建
	outputFile, outputError := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return  
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"

	for i:=0; i<10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}
```

### Go 文件拷贝
如何拷贝一个文件到另一个文件？最简单的方式就是使用 io 包。  
```golang
// filecopy.go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	CopyFile("target.txt", "source.txt")
	fmt.Println("Copy done!")
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
```
`注意 defer 的使用：当打开 dst 文件时发生了错误，那么 defer 仍然能够确保 src.Close() 执行。如果不这么做，src 文件会一直保持打开状态并占用资源。`

### Go 文件压缩
compress 包提供了读取压缩文件的功能，支持的压缩文件格式为：bzip2、flate、gzip、lzw 和 zlib。
```golang
package main

import (
	"fmt"
	"bufio"
	"os"
	"compress/gzip"
)

func main() {
	fName := "MyFile.gz"
	var r *bufio.Reader
	fi, err := os.Open(fName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], fName,
			err)
		os.Exit(1)
	}
	defer fi.Close()
	fz, err := gzip.NewReader(fi)
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}

```
