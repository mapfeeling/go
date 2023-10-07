package cond

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 声明互斥锁
	var mutex sync.Mutex
	// 声明条件变量
	cond := sync.NewCond(&mutex)
	for i := 1; i <= 10; i++ {
		go func(i int) {
			// 获取锁
			cond.L.Lock()
			// 释放锁
			defer cond.L.Unlock()
			// 等待通知,阻塞当前协程
			cond.Wait()
			// 等待通知后打印输出
			fmt.Printf("输出:%d ! \n", i)
		}(i)
	}
	// 单个通知
	time.Sleep(time.Second)
	fmt.Println("单个通知A！")
	cond.Signal()
	time.Sleep(time.Second)
	fmt.Println("单个通知B！")
	cond.Signal()

	// 广播通知
	time.Sleep(time.Second)
	fmt.Println("广播通知！并睡眠1秒，等待其他协程输出!")
	cond.Broadcast()
	// 等待其他协程处理完
	time.Sleep(time.Second)
	fmt.Println("运行结束！")
}

/**输出
  单个通知A！
  输出:1 !
  单个通知B！
  输出:4 !
  广播通知！并睡眠1秒，等待其他协程输出!
  输出:10 !
  输出:2 !
  输出:3 !
  输出:8 !
  输出:9 !
  输出:6 !
  输出:5 !
  输出:7 !
  运行结束！
*/
