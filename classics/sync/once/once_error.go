package once

/*
sync.Once 提供的 Do 方法并没有返回值
意味着如果我们传入的函数如果发生 error 导致初始化失败,后续调用 Do 方法也不会再初始化
为了避免这个问题,我们可以实现一个 类似 sync.Once 的并发原语
*/

import (
	"sync"
	"sync/atomic"
)

type OnceStrengthen struct {
	done uint32
	m    sync.Mutex
}

func (o *OnceStrengthen) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 0 {
		return o.doSlow(f)
	}
	return nil
}

func (o *OnceStrengthen) doSlow(f func() error) error {
	o.m.Lock()
	defer o.m.Unlock()
	var err error
	if o.done == 0 {
		err = f()
		// 只有没有 error 的时候，才修改 done 的值
		if err == nil {
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}

/*
上述代码实现了一个加强的 Once 结构体。与标准的 sync.Once 不同
这个实现允许 Do 方法的函数参数返回一个 error。如果执行函数没有返回 error
则修改 done 的值以表示函数已执行
这样，在后续的调用中，只有在没有发生 error 的情况下，才会跳过函数执行，避免初始化失败
*/
