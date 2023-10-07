package RWMutex

import (
	"fmt"
	"strconv"
	"sync"
)

// 声明全局变量，文件内容
var fileContext string

// 声明全局读写互斥锁
var rxMutex sync.RWMutex

// 声明全局等待组
var wg sync.WaitGroup

func main() {
	// 设置计数器
	wg.Add(5)
	for i := 1; i <= 5; i++ {
		name := "同学-" + strconv.Itoa(i)
		if i%2 == 0 {
			go readFile(name)
		} else {
			go writeFile(name, strconv.Itoa(i))
		}
	}
	// 等待所有计数器执行完成
	wg.Wait()
	fmt.Println("运行结束!")
}

// 读文件
func readFile(name string) {
	// 释放读锁
	defer rxMutex.RUnlock()
	// 获取读锁
	rxMutex.RLock()
	// 打印读取内容
	fmt.Printf("%s 获取读锁，读取内容为: %s \n", name, fileContext)
	// 计数器减一
	wg.Done()
}

// 写文件
func writeFile(name, s string) {
	// 释放写锁
	defer rxMutex.Unlock()
	// 获取写锁
	rxMutex.Lock()
	// 写入内容
	fileContext = fileContext + " " + s
	fmt.Printf("%s 获取写锁，写入内容: %s。 文件内容变成: %s \n", name, s, fileContext)
	// 计数器减一
	wg.Done()
}

/**输出
  同学-1 获取写锁，写入内容: 1。 文件内容变成:  1
  同学-4 获取读锁，读取内容为:  1
  同学-2 获取读锁，读取内容为:  1
  同学-5 获取写锁，写入内容: 5。 文件内容变成:  1 5
  同学-3 获取写锁，写入内容: 3。 文件内容变成:  1 5 3
  运行结束!
*/
