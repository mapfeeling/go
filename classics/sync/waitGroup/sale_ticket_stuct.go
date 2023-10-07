package waitGroup

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// 声明一个票池
type ticketPool struct {
	over int
	lock sync.Mutex
	wg   sync.WaitGroup
}

// 定义售票方法
func (t *ticketPool) sellTicket(windowName string) {
	// 等待组减一
	defer t.wg.Done()
	for {
		// 加锁
		t.lock.Lock()
		if t.over > 0 {
			time.Sleep(10 * time.Millisecond)
			t.over--
			fmt.Printf("%s 卖出一张票，余票: %d \n", windowName, t.over)
		} else {
			// 无票，跳无限循环并解锁
			t.lock.Unlock()
			fmt.Printf("%s 票已卖完! \n", windowName)
			break
		}
		// 正常售票流程解锁
		t.lock.Unlock()
	}
}
func main() {
	// 创建一个票池
	ticketP := ticketPool{over: 10}
	fmt.Printf("T:%T v: %v \n", ticketP, ticketP)
	// 设置窗口数量
	windowNum := 3
	// 设置等待组计数器
	ticketP.wg.Add(windowNum)
	// 定义3个窗口售票
	for i := 1; i <= windowNum; i++ {
		go ticketP.sellTicket("窗口" + strconv.Itoa(i))
	}
	ticketP.wg.Wait()
	fmt.Println("运行结束!")
}

/**输出
  窗口3 卖出一张票，余票: 9
  窗口3 卖出一张票，余票: 8
  窗口1 卖出一张票，余票: 7
  窗口2 卖出一张票，余票: 6
  窗口3 卖出一张票，余票: 5
  窗口1 卖出一张票，余票: 4
  窗口2 卖出一张票，余票: 3
  窗口3 卖出一张票，余票: 2
  窗口1 卖出一张票，余票: 1
  窗口2 卖出一张票，余票: 0
  窗口1 票已卖完!
  窗口2 票已卖完!
  窗口3 票已卖完!
  运行结束!
*/
