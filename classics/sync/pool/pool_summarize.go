package pool

import (
	"sync"
	"unsafe"
)

/*
频繁的创建对象，GC每次需要检查这些对象是否可以回收，然后释放对象
sync.Pool采用对象池的形式，把不用的对象回收起来，而不是直接GC掉，下次再使用的时候不用重新创建对象，使用对象池已有的对象。这样做有两方面收益：
减轻GC负担
减少新对象的申请 复用对象，避免每次使用对象都要创建对象
*/
type noCopy struct{}
type Pool struct {
	noCopy noCopy

	// local是一个指向[p]poolLocal的指针，它指向的是一个数组，此数组中的元素是poolLocal类型
	// 数组的大小是固定的，数量与p的数量是相同的，也是每个p对应数组中的1个元素
	local unsafe.Pointer // local fixed-size per-P pool, actual type is [P]poolLocal
	// local数组的大小
	localSize uintptr // size of the local array

	// victim类型与local是一样的，可以理解为local的中转站，过渡存储local用的
	victim     unsafe.Pointer // local from previous cycle
	victimSize uintptr        // size of victims array

	// New optionally specifies a function to generate
	// a value when Get would otherwise return nil.
	// It may not be changed concurrently with calls to Get.
	// New方法，返回的空接口类型，该字段也是Pool唯一暴露给外界的字段
	// New方法可以赋值一个能够产生值的函数，在调用Get方法的时候可以用
	// New方法来产生一个value，如果没有给New赋值，调用Get时将返回nil.
	New func() any
}

// Pool是一个可以分别存取的临时对象的集合
// 可以被看作是一个存放可重用对象的值的容器、过减少GC来提升性能，是Goroutine并发安全的
// 有两个方法 Get()、Set()

// Get Get方法没有取得item：如p.New非nil，Get返回调用p.New的结果；否则返回nil
func (p *Pool) Get() interface{}

// Put 将对象放入对象池中
func (p *Pool) Put(x interface{})

// New初始化Pool实例
var bufferPool = &sync.Pool{
	New: func() interface{} {
		println("Create pool instance")
		return struct{}{}
	}}

// 为什么 sync.Pool 不适合用于像 socket 长连接或数据库连接池?
// 因为，我们不能对 sync.Pool 中保存的元素做任何假设，以下事情是都可以发生的：
// 1、Pool 池里的元素随时可能释放掉，释放策略完全由 runtime 内部管理；
// 2、Get 获取到的元素对象可能是刚创建的，也可能是之前创建好 cache 住的。使用者无法区分；
// 3、Pool 池里面的元素个数你无法知道；
// 所以，只有的你的场景满足以上的假定，才能正确的使用 Pool 。sync.Pool 本质用途是增加临时对象的重用率，减少 GC 负担。
// 划重点：临时对象。所以说，像 socket 这种带状态的，长期有效的资源是不适合 Pool 的

// 总结：
// 1、sync.Pool 本质用途是增加临时对象的重用率，减少 GC 负担；
// 2、不能对 Pool.Get 出来的对象做预判，有可能是新的（新分配的），有可能是旧的（之前人用过，然后 Put 进去的）；
// 3、不能对 Pool 池里的元素个数做假定，你不能够；
// 4、sync.Pool 本身的 Get, Put 调用是并发安全的，sync.New 指向的初始化函数会并发调用，里面安不安全只有自己知道；
// 5、当用完一个从 Pool 取出的实例时候，一定要记得调用 Put，否则 Pool 无法复用这个实例，通常这个用 defer 完成
