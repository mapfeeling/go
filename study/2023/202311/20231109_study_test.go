package _02311

import (
	"fmt"
	"testing"
)

var aaa bool = true

// return 之后的 defer 语句会执行吗，下面这段代码输出什么？
func TestStudy20231109(t *testing.T) {
	defer func() {
		fmt.Println("1")
	}()
	if aaa == true {
		fmt.Println("2")
		return
	}
	defer func() {
		fmt.Println("3")
	}()
}

// defer 关键字后面的函数或者方法想要执行必须先注册
// return 之后的 defer 是不能注册的
// 也就不能执行后面的函数或方法
