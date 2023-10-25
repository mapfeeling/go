package prime

import (
	"testing"
	"time"
)

func worker() chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 42
	}()
	return ch
}

func TestChan(t *testing.T) {
	timeStart := time.Now()

	_, _ = <-worker(), <-worker()
	// 等价于
	// _ = <-worker()
	// _ = <-worker()

	println(int(time.Since(timeStart).Seconds()))
}
