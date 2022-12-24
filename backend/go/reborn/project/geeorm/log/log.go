package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	// log.Lshortfile 支持显示文件名和代码行号
	// 红色
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	// 蓝色
	infoLog = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers = []*log.Logger{errorLog, infoLog}
	mu      sync.Mutex
)

// log methods
var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

// log levels
const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

// SetLevel controls log level
// 根据 level 设置输出日志是否打印
// 如果设置为 ErrorLevel，infoLog 的输出会被定向到 ioutil.Discard，即不打印该日志
func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level {
		// Discard 放弃输出
		errorLog.SetOutput(ioutil.Discard)
	}
	if InfoLevel < level {
		// Discard 放弃输出
		infoLog.SetOutput(ioutil.Discard)
	}
}
