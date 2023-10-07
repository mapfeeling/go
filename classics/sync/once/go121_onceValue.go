package once

import (
	"fmt"
	"math/rand"
	"sync"
)

// func OnceValue[T any](f func() T) func() T
// OnceValue返回一个函数， 这个函数会返回f的返回值
// 多次调用都会返回同一个值

func main() {
	randValue := func() int {
		return rand.Int()
	}
	val := sync.OnceValue(randValue)
	for i := 0; i < 10; i++ {
		fmt.Println(val())
	}
}
