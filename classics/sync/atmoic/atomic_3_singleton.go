package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

/*
场景3、单例模式
在多线程环境下,如果需要确保只有一个实例被创建,可以使用atomic包来实现单例模式
*/
type singleton struct{}

var (
	instance          *singleton
	initialized       uint32
	nonInitalizeCount int
)

func getInstance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		nonInitalizeCount++
		return instance
	}
	// 初始化
	initialize()
	return instance
}

func initialize() {
	if atomic.LoadUint32(&initialized) == 0 {
		// 进行初始化操作
		instance = &singleton{}
		atomic.StoreUint32(&initialized, 1)
	}
}

func main() {
	// 多个goroutine并发访问单例对象
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("%p\n", getInstance())
		}()
	}
	// 等待goroutine执行完毕
	time.Sleep(time.Second)
	fmt.Println("未初始化的数量:", nonInitalizeCount)
}
