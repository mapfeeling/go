package main

import (
	"fmt"
	"sync/atomic"
)

/*
AddInt32，AddInt64：用于对int32、int64类型的变量进行原子性的加法操作
CompareAndSwapInt32，CompareAndSwapInt64：用于对int32、int64类型的变量进行比较并交换操作
SwapInt32，SwapInt64：用于对int32、int64类型的变量进行原子性的交换操作
LoadInt32，LoadInt64：用于对int32、int64类型的变量进行原子性的读操作
StoreInt32，StoreInt64：用于对int32、int64类型的变量进行原子性的写操
*/

func main() {
	var number int32
	if ok := atomic.CompareAndSwapInt32(&number, 0, 1); ok {
		fmt.Println("atomic.CompareAndSwapInt32", number)
	}
	atomic.StoreInt32(&number, 999)
	fmt.Println("atomic.StoreInt32", number)
}
