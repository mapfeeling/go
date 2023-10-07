package once

import (
	"sync"
	"sync/atomic"
)

/*
sync.Once 是 Go 语言中的一种同步原语,用于确保某个操作或函数在并发环境下只被执行一次
它只有一个导出的方法,即 Do,该方法接收一个函数参数
在 Do 方法被调用后,该函数将被执行,而且只会执行一次,即使在多个协程同时调用的情况下也是如此

sync.Once 主要用于以下场景:
单例模式:确保全局只有一个实例对象,避免重复创建资源
延迟初始化:在程序运行过程中需要用到某个资源时,通过 sync.Once 动态地初始化该资源
只执行一次的操作:例如只需要执行一次的配置加载、数据清理等操作
*/

// sync.once的实现原理

type Once struct {
	// 表示是否执行了操作
	done uint32
	// 互斥锁，确保多个协程访问时，只能一个协程执行操作
	m sync.Mutex
}

// Do Do方法不可以在Do中循环调用Do,否则会造成死锁
func (o *Once) Do(f func()) {
	// 判断 done 的值,如果是 0,说明 f 还没有被执行过
	if atomic.LoadUint32(&o.done) == 0 {
		// 构建慢路径(slow-path)，以允许对 Do 方法的快路径(fast-path)进行内联
		o.doSlow(f)
	}
}

func (o *Once) doSlow(f func()) {
	// 加锁
	o.m.Lock()
	defer o.m.Unlock()
	// 双重检查,避免 f 已被执行过
	if o.done == 0 {
		// 修改 done 的值
		defer atomic.StoreUint32(&o.done, 1)
		// 执行函数
		f()
	}
}

/*
1、sync.Once 结构体包含两个字段:done 和 mu
done 是一个 uint32 类型的变量，用于表示操作是否已经执行过；
m 是一个互斥锁，用于确保在多个协程访问时，只有一个协程能执行操作
2、sync.Once 结构体包含两个方法：Do 和 doSlow。Do 方法是其核心方法，它接收一个函数参数 f
首先它会通过原子操作atomic.LoadUint32（保证并发安全） 检查 done 的值，如果为 0，表示 f 函数没有被执行过，然后执行 doSlow 方法
3、在 doSlow 方法里，首先对互斥锁 m 进行加锁，确保在多个协程访问时，只有一个协程能执行 f 函数。
接着再次检查 done 变量的值，如果 done 的值仍为 0，说明 f 函数没有被执行过
此时执行 f 函数，最后通过原子操作 atomic.StoreUint32 将 done 变量的值设置为 1
*/

// 问:为什么会封装一个 doSlow 方法？
// 答：doSlow 方法的存在主要是为了性能优化
//将慢路径（slow-path）代码从 Do 方法中分离出来，使得 Do 方法的快路径（fast-path）能够被内联（inlined），从而提高性能

// 问:为什么会有双重检查（double check）的写法？
// 第一次检查：在获取锁之前，先使用原子加载操作 atomic.LoadUint32 检查 done 变量的值，如果 done 的值为 1，表示操作已执行，此时直接返回，不再执行 doSlow 方法。这一检查可以避免不必要的锁竞争。
// 第二次检查：获取锁之后，再次检查 done 变量的值，这一检查是为了确保在当前协程获取锁期间，其他协程没有执行过 f 函数。如果 done 的值仍为 0，表示 f 函数没有被执行过。
// 通过双重检查，可以在大多数情况下避免锁竞争，提高性能
