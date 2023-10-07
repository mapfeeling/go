package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

/*
场景1、计数器
计数器是一种非常常见的场景
特别是在并发编程中,多个goroutine需要对同一个计数器进行操作时,就需要保证计数器的值是正确的
*/
func main() {
	var count int32 = 0
	for i := 0; i < 10; i++ {
		go func() {
			atomic.AddInt32(&count, 1)
		}()
	}
	// 等待goroutine执行完毕
	time.Sleep(time.Second)
	fmt.Printf("count: %d\n", count)
}
