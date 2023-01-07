package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type singleton struct{}

var instance *singleton

func GetInstance() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})

	return instance
}

func (s *singleton) SomeThing() {
	fmt.Println("单例对象的某方法")
}

func main() {
	s := GetInstance()
	s.SomeThing()
}

// Once.Do()方法的源代码：
// func (o *Once) Do(f func()) {　　　//判断是否执行过该方法，如果执行过则不执行
//     if atomic.LoadUint32(&o.done) == 1 {
//         return
//     }
//     // Slow-path.
//     o.m.Lock()
//     defer o.m.Unlock()
//     if o.done == 0 {
//         defer atomic.StoreUint32(&o.done, 1)
//         f()
//     }
// }
