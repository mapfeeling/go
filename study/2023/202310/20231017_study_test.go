package _02310

import (
	"fmt"
	"time"
)

// 问执行一下代码会如何？

func main() {
	ch := make(chan struct{})
	go fmt.Println(<-ch)
	ch <- struct{}{}
	time.Sleep(time.Second)
}

// 死锁
// go 语句后面的函数调用,其参数会先求值,这和普通的函数调用求值一样
// 意思是说，函数调用之前，实参就被求值好了
// 因此这道题目 go fmt.Println(<-ch1) 语句中的 <-ch1 是在 main goroutine 中求值的
// 这相当于一个无缓冲的 chan，发送和接收操作都在一个 goroutine 中（main goroutine）进行，因此造成死锁

// 以defer举例
// 正确使用方式

/*
defer func() {
  recover()
}()
*/

// 而不是defer recover()
