package closure

import (
	"fmt"
	"testing"
)

func doLater(msg string) func() {
	return func() {
		fmt.Println("Later:", msg)
	}
}

func TestDeferExecution(t *testing.T) {
	msg := "Hello World!"
	deferFunc := doLater(msg)
	defer deferFunc()
	fmt.Println("Doing something...")
}

/*
在上面的例子中,doLater 函数返回了一个闭包函数,该闭包函数在被调用时会输出 msg 变量的内容
我们将这个闭包函数通过 defer 延迟执行,这样在函数执行结束后,doLater 返回的闭包函数会在后续的代码之前被执行
*/
