// 执行下面代码会发生什么情况？
package main

import (
	"fmt"
	"testing"
	"time"
)

func TestStudy20230924(t *testing.T) {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}

/*
记住 channel 的一些关键特性：
给一个 nil channel 发送数据，造成永远阻塞
从一个 nil channel 接收数据，造成永远阻塞
给一个已经关闭的 channel 发送数据，引起 panic
从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值
无缓冲的channel是同步的，而有缓冲的channel是非同步的
15字口诀：“空读写阻塞，写关闭异常，读关闭空零”，往已经关闭的 channel 写入数据会 panic。
本题中，因为 main 在开辟完两个 goroutine 之后，立刻关闭了 ch， 结果就是 panic：
panic: send on closed channel
*/
