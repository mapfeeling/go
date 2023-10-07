package pool

import (
	"io"
	"os"
	"sync"
)

/*
newPrinter就是调用的sync.Pool.Get(),拿到pp指针
首先是做了一些format操作。然后调用free()方法，将使用过得pp放回到ppFree中。归还之前将p的部分字段重置，以保证下次调用的是原始pp
*/

//fmt.Printf

func Printf(format string, a ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	p := newPrinter()
	//...
	//使用完后释放，将pp归还给了sync.pool
	p.free()
	return
}

// 返回一个pp结构体指针
func newPrinter() *pp {
	p := ppFree.Get().(*pp)
	p.panicking = false
	p.erroring = false
	p.wrapErrs = false
	p.fmt.init(&p.buf)
	return p
}

var ppFree = sync.Pool{
	New: func() interface{} { return new(pp) },
}

// 结构体pp
type pp struct {
	//...
}
