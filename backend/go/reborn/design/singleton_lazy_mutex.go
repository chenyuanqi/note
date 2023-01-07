package main

import (
	"fmt"
	"sync"
)

//定义锁
var lock sync.Mutex

type singleton struct{}

var instance *singleton

func GetInstance() *singleton {
	//为了线程安全，增加互斥
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		return new(singleton)
	} else {
		return instance
	}
}

func (s *singleton) SomeThing() {
	fmt.Println("单例对象的某方法")
}

func main() {
	s := GetInstance()
	s.SomeThing()
}
