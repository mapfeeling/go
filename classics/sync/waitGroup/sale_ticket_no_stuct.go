package main

// sync.waitGroup不作为结构体属性使用
import (
	"fmt"
	"sync"
	"time"
)

// 声明全局等待组
var wg sync.WaitGroup

// 声明全局锁
var mutex sync.Mutex

// 因为sync.WaitGroup不是引用类型，当它在函数之间传递的时候，需要使用指针
// 声明全局余票
var ticket int = 10

func main() {
	// 设置等待组计数器
	wg.Add(3)
	// 窗口卖票
	go saleTicket("窗口A", &wg)
	go saleTicket("窗口B", &wg)
	go saleTicket("窗口C", &wg)
	wg.Wait()
	fmt.Println("运行结束!")
}

// 卖票流程
func saleTicket(windowName string, wg *sync.WaitGroup) {
	// 卖票流程结束后关闭
	defer wg.Done()
	for {
		// 加锁
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(10 * time.Millisecond)
			ticket--
			fmt.Printf("%s 卖出一张票，余票: %d \n", windowName, ticket)
		} else {
			fmt.Printf("%s 票已卖完! \n", windowName)
			// 解锁
			mutex.Unlock()
			break
		}
		// 解锁
		mutex.Unlock()
	}
}

/**输出
  窗口C 卖出一张票，余票: 9
  窗口C 卖出一张票，余票: 8
  窗口B 卖出一张票，余票: 7
  窗口A 卖出一张票，余票: 6
  窗口C 卖出一张票，余票: 5
  窗口B 卖出一张票，余票: 4
  窗口A 卖出一张票，余票: 3
  窗口C 卖出一张票，余票: 2
  窗口B 卖出一张票，余票: 1
  窗口A 卖出一张票，余票: 0
  窗口C 票已卖完!
  窗口B 票已卖完!
  窗口A 票已卖完!
  运行结束!
*/
