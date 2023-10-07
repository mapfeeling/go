package once

import (
	"fmt"
	"sync"
)

// func OnceFunc(f func()) func()
// OnceFunc返回一个可以并发调用的函数，它可以被调用多次
// 即使返回的函数被调用多次，f也只会被调用一次
func main() {
	var onceBody = func() {
		fmt.Println("Only once func")
	}
	var f = sync.OnceFunc(onceBody)
	for i := 0; i < 10; i++ {
		f()
	}
}
