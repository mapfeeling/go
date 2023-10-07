package mutex

import (
	"internal/race"
	"sync/atomic"
	"unsafe"
)

/*
互斥锁的公平性

互斥锁有两种运行模式：正常模式和饥饿模式
在正常模式下，请求锁的goroutine会按 FIFO(先入先出)的顺序排队，依次被唤醒，但被唤醒的goroutine并不能直接获取锁，而是要与新请求锁的goroutines去争夺锁的所有权
但是这其实是不公平的，因为新请求锁的goroutines有一个优势：他们正在cpu上运行且数量可能比较多，所以新唤醒的goroutine在这种情况下很难获取锁。在这种情况下，如果这个goroutine获取失败，会直接插入到队列的头部
如果一个等待的goroutine超过1ms时仍未获取到锁，会将这把锁转换为饥饿模式。

在饥饿模式下，互斥锁的所有权会直接从解锁的goroutine转移到队首的goroutine
并且新到达的goroutines不会尝试获取锁，会直接插入到队列的尾部

如果一个等待的goroutine获得了锁的所有权，并且满足以下两个条件之一：
(1) 它是队列中的最后一个goroutine
(2) 它等待的时间少于 1 毫秒(hard code在代码里)
它会将互斥锁切回正常模式

普通模式具有更好的性能，因为即使有很多阻塞的等待锁的goroutine，一个goroutine也可以尝试请求多次锁
而饥饿模式则可以避免尾部延迟这种bad case
*/

type Mutex struct {
	// state字段就是这把mutex的状态，二进制低3位对应锁的状态，将state右移3位代表mutex的数量
	state int32
	// sema（信号量）用来唤醒goroutine
	sema uint32
}

// Mutex是一种互斥锁
// Mutex的零值是一个未锁定的互斥锁,即sync.mutex的0值就是一把开启状态的互斥锁
//
// Mutex在首次使用后不得复制,mutex 数据结构禁止拷贝，拷贝mutex的结果是未定义的 即要用指针接受者而非值接受者
//
// 在Go存储器模型的术语中,第n次解锁调用“同步于”第m次锁定调用
// 对于任何n<m,对TryLock的成功调用相当于对Lock的调用
// 对TryLock的失败调用未建立任何“之前同步”关系

// 关键常量
const (
	mutexLocked           = 1 << iota // mutex is locked mutex是一个加锁状态，对应0b001
	mutexWoken                        // 是否有goroutine被唤醒or新来的goroutine，在尝试获取锁， 对应0b010
	mutexStarving                     // 当前锁处于饥饿状态， 对应0b10
	mutexWaiterShift      = iota      // waiter的数量移位，通过mutex.state>> mutexWaiterShift可以得到waiter的数量
	starvationThresholdNs = 1e6       // 这里就是上文中提到的转换成饥饿模式的时间限制，在源码里写死为1e6 ns,也就是1ms)
)

// 接口

type Locker interface {
	Lock()   // 锁住m，如果 m 已经加锁，则阻塞直到 m 解锁
	UnLock() // 解锁 m，如果 m 未加锁会导致运行时错误
}

// 方法

// Lock 方法锁住 m，如果 m 已经加锁，则阻塞直到 m 解锁
func (m *Mutex) Lock() {
	// Fast path: grab unlocked mutex.
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		if race.Enabled {
			race.Acquire(unsafe.Pointer(m))
		}
		return
	}
	// Slow path (outlined so that the fast path can be inlined) 慢速路径（概述以便可以内联快速路径）
	m.lockSlow()
}

// Unlock 方法解锁 m，如果 m 未加锁会导致运行时错误
func (m *Mutex) Unlock() {
	if race.Enabled {
		_ = m.state
		race.Release(unsafe.Pointer(m))
	}

	// Fast path: drop lock bit.
	new := atomic.AddInt32(&m.state, -mutexLocked)
	if new != 0 {
		// Outlined slow path to allow inlining the fast path.
		// To hide unlockSlow during tracing we skip one extra frame when tracing GoUnblock.
		m.unlockSlow(new)
	}
}

// 注意事项
/*
1、在一个 goroutine 获得 Mutex 后，其他 goroutine 只能等到这个 goroutine 释放该 Mutex
2、使用 Lock() 加锁后，不能再继续对其加锁，直到利用 Unlock() 解锁后才能再加锁
3、在 Lock() 之前使用 Unlock() 会导致 panic 异常
4、已经锁定的 Mutex 并不与特定的 goroutine 相关联，这样可以利用一个 goroutine 对其加锁，再利用其他 goroutine 对其解锁
5、在同一个 goroutine 中的 Mutex 解锁之前再次进行加锁，会导致死锁
6、适用于读写不确定，并且只有一个读或者写的场景
*/

// mutex涉及的runtime调用
// 确定当前goroutine是否可以自旋（因为自旋是非常消耗cpu的，所以对自旋操作也有一定的次数限制）
func runtime_canSpin(i int) bool

// 执行自旋操作
func runtime_doSpin()

// 获取当前毫秒时间戳
func runtime_nanotime() int64

// 信号量的实现，对应信号量的P原语操作，s的值代表信号量，包含一个fifo等待队列，如果lifo为true，则放入队首。skipframes只有在开启tracing时有效。
func runtime_SemacquireMutex(s *uint32, lifo bool, skipframes int)

// 信号量的实现，对应信号量的V原语操作，如果handoff为true，则会将锁直接转移给队首goroutine
func runtime_Semrelease(s *uint32, handoff bool, skipframes int)
