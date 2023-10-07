package RWMutex

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// 定义一个文件结构体
type fileResource struct {
	content string
	wg      sync.WaitGroup
	rwLock  sync.RWMutex
}

// 读文件
func (f *fileResource) readFile(name string) {
	// 释放读锁
	defer f.rwLock.RUnlock()
	// 获取读锁
	f.rwLock.RLock()
	// 打印读取内容
	time.Sleep(time.Second)
	fmt.Printf("%s 获取读锁，读取内容为: %s \n", name, f.content)
	// 计数器减一
	f.wg.Done()
}

// 写文件
func (f *fileResource) writeFile(name, s string) {
	// 释放写锁
	defer f.rwLock.Unlock()
	// 获取写锁
	f.rwLock.Lock()
	// 写入内容
	time.Sleep(time.Second)
	f.content = f.content + " " + s
	fmt.Printf("%s 获取写锁，写入内容: %s。 文件内容变成: %s \n", name, s, f.content)
	// 计数器减一
	f.wg.Done()
}
func main() {
	// 声明结构体
	var file fileResource
	// 设置计数器
	file.wg.Add(5)
	for i := 1; i <= 5; i++ {
		name := "同学-" + strconv.Itoa(i)
		if i%2 == 0 {
			go file.readFile(name)
		} else {
			go file.writeFile(name, strconv.Itoa(i))
		}
	}
	// 等待所有计数器执行完成
	file.wg.Wait()
	fmt.Println("运行结束!")
}

/**输出
  同学-5 获取写锁，写入内容: 5。 文件内容变成:  5
  同学-1 获取写锁，写入内容: 1。 文件内容变成:  5 1
  同学-2 获取读锁，读取内容为:  5 1
  同学-3 获取写锁，写入内容: 3。 文件内容变成:  5 1 3
  同学-4 获取读锁，读取内容为:  5 1 3
  运行结束!
*/
