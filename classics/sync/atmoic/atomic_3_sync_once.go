package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
场景3、单例模式
在多线程环境下,如果需要确保只有一个实例被创建,可以使用atomic包来实现单例模式
此处
*/
type singletonAnother struct{}

var (
	instanceAnother *singletonAnother
	once            sync.Once
)

func getInstanceAnother() *singletonAnother {
	// 初始化
	initializeAnother()
	return instanceAnother
}

func initializeAnother() {
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
			//fmt.Printf("%p\n", getInstance())
			once.Do(func() {
				fmt.Println("初始化一次")
				getInstanceAnother()
			})
		}()
	}
	// 等待goroutine执行完毕
	time.Sleep(time.Second)
}
