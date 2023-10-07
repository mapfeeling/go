package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

/*
场景2、状态标志
在并发编程中,经常需要使用状态标志来控制不同的goroutine的执行
如果多个goroutine需要共享同一个状态标志,使用atomic包可以保证状态标志的值在不同的goroutine之间正确同步
*/
func main() {
	var flag int32 = 0
	go func() {
		atomic.StoreInt32(&flag, 1)
	}()
	// 等待goroutine执行完毕
	time.Sleep(time.Second)
	if atomic.LoadInt32(&flag) == 1 {
		fmt.Println("flag is true")
	} else {
		fmt.Println("flag is false")
	}
}
