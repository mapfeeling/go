package RWMutex

import "sync"

/*
RWMutex是读写互斥锁，简称读写锁
该锁可以同时被多个读取者持有或被唯一个写入者持有
RWMutex类型锁跟Goroutine无关，可以由不同的Goroutine加锁、解锁
RWMutex也可以创建为其他结构体的字段；零值为解锁状态
*/

type RWMutex struct {
	w           sync.Mutex //用于控制多个写锁，获得写锁首先要获取该锁，如果有一个写锁在进行，那么再到来的写锁将会阻塞于此
	writerSem   uint32     //写阻塞等待的信号量，最后一个读者释放锁时会释放信号量
	readerSem   uint32     //读阻塞的协程等待的信号量，持有写锁的协程释放锁后会释放信号量
	readerCount int32      //记录读者个数
	readerWait  int32      //记录写阻塞时读者个数
}

// 读写锁堵塞场景

// 写锁需要阻塞写锁：一个协程拥有写锁时，其他协程写锁需要阻塞
// 写锁需要阻塞读锁：一个协程拥有写锁时，其他协程读锁需要阻塞
// 读锁需要阻塞写锁：一个协程拥有读锁时，其他协程写锁需要阻塞
// 读锁不能阻塞读锁：一个协程拥有读锁时，其他协程也可以拥有读锁

func (rw *RWMutex) RLock()   // 获取读锁,当一个协程拥有读锁时，其他协程写锁需要阻塞
func (rw *RWMutex) RUnlock() // 释放读锁
func (rw *RWMutex) Lock()    // 获取写锁，与Mutex完全一致
func (rw *RWMutex) Unlock()  // 释放写锁
