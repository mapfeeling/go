package once

import (
	"fmt"
	"math/rand"
	"sync"
)

// func OnceValues[T1, T2 any](f func() (T1, T2)) func() (T1, T2)
// OnceValues和OnceValue的功能类似，只不过返回两个参数，仅此而已

func main() {
	randValue := func() (int, int) {
		return rand.Int(), rand.Int() + 1
	}

	val := sync.OnceValues(randValue)
	for i := 0; i < 10; i++ {
		fmt.Println(val())
	}
}
