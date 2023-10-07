package mutex

import (
	"fmt"
	"sync"
)

/*
在Go语言中，如果结构体（struct）中包含了sync.Mutex类型的字段，那么该结构体无法使用值接收器（value receiver）
这是因为sync.Mutex类型本身是一个引用类型，而不是值类型
在Go语言中，方法可以通过值接收器（value receiver）或指针接收器（pointer receiver）定义
值接收器是将方法与结构体的值关联，而指针接收器是将方法与结构体的指针关联
当一个结构体包含sync.Mutex类型的字段时，如果该结构体的方法使用了值接收器，会导致复制该结构体的值
由于sync.Mutex是一个引用类型，复制后的值将只包含sync.Mutex的副本，而不是原始的sync.Mutex
这将导致在方法中对sync.Mutex的操作不会影响到原始结构体的sync.Mutex
为了解决这个问题，可以使用指针接收器来定义结构体的方法。指针接收器将方法与结构体的指针关联，从而避免了复制结构体的值
*/

/*
core: 值接收器是将方法与结构体的值关联，而指针接收器是将方法与结构体的指针关联
*/

type MutexPointReceiver struct {
	mutex sync.Mutex
	data  int
}

func (m *MutexPointReceiver) Increment() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.data++
}

func (m *MutexPointReceiver) GetValue() int {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.data
}

func main() {
	s := &MutexPointReceiver{data: 42}
	s.Increment()
	fmt.Println(s.GetValue()) // 输出: 43
}

/*
go语言接受者的选取
何时使用值类型
1.如果接受者是一个 map，func 或者 chan，使用值类型(因为它们本身就是引用类型)。
2.如果接受者是一个 slice，并且方法不执行 reslice 操作，也不重新分配内存，使用值类型。
3.如果接受者是一个小的数组或者原生的值类型结构体类型(比如 time.Time 类型)，而且没有可修改的字段和指针，又或者接受者是一个简单地基本类型像是 int 和 string，使用值类型就好了。

一个值类型的接受者可以减少一定数量的垃圾生成，如果一个值被传入一个值类型接受者的方法，一个栈上的拷贝会替代在堆上分配内存(但不是保证一定成功)，所以在没搞明白代码想干什么之前，别因为这个原因而选择值类型接受者。

何时使用指针类型
1.如果方法需要修改接受者，接受者必须是指针类型。
2.如果接受者是一个包含了 sync.Mutex 或者类似同步字段的结构体，接受者必须是指针，这样可以避免拷贝。
3.如果接受者是一个大的结构体或者数组，那么指针类型接受者更有效率。(多大算大呢？假设把接受者的所有元素作为参数传给方法，如果你觉得参数有点多，那么它就是大)。
4.从此方法中并发的调用函数和方法时，接受者可以被修改吗？一个值类型的接受者当方法调用时会创建一份拷贝，所以外部的修改不能作用到这个接受者上。如果修改必须被原始的接受者可见，那么接受者必须是指针类型。
5.如果接受者是一个结构体，数组或者 slice，它们中任意一个元素是指针类型而且可能被修改，建议使用指针类型接受者，这样会增加程序的可读性
当你看完这个还是有疑虑，还是不知道该使用哪种接受者，那么记住使用指针接受者
*/
