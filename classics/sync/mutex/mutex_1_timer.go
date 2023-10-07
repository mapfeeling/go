package mutex

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// SafeCounter 定义计数器
type SafeCounter struct {
	mu sync.Mutex
	m  map[string]int64
}

// Inc 计数器生成
// 读者你知道下面为什么必须要使用指针接受者吗？
// “Load”通过值传递锁：类型“SafeCounter”包含“sync.Mutex”，即“sync.Locker”
func (s SafeCounter) Inc(key string) {
	defer wg.Done()
	// 加锁
	s.mu.Lock()
	// 对计数器的数值 增值+1
	s.m[key]++
	// 解锁
	s.mu.Unlock()
}

func (s SafeCounter) Load(key string) int64 {
	//defer wg.Done()
	// 加锁
	s.mu.Lock()
	val := s.m[key]
	// 解锁
	s.mu.Unlock()
	return val
}
func main() {
	// 初始化计数器
	var sc = &SafeCounter{m: make(map[string]int64)}
	// 开启1000个数生成计数器
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go sc.Inc("AAA")
	}
	wg.Wait()
	// 打印计数器
	fmt.Println(sc.Load("AAA"))
}
