package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

/*
场景4、等待通知机制
在并发编程中,经常需要使用等待通知机制来协调不同的goroutine
使用atomic包提供的原子操作可以确保等待和通知的正确性
*/
var (
	count  int32 = 0
	notify int32 = 0
)

func worker(id int) {
	for {
		if atomic.LoadInt32(&notify) == 1 {
			break
		}
		// 进行工作
		atomic.AddInt32(&count, 1)
	}
}

func main() {
	// 启动多个goroutine进行工作
	for i := 0; i < 10; i++ {
		go worker(i)
	}
	// 等待一段时间后发出通知
	time.Sleep(time.Second)
	// 发送通知
	atomic.StoreInt32(&notify, 1)
	// 等待goroutine执行完毕
	time.Sleep(time.Second)
	fmt.Printf("count: %d\n", count)
}
