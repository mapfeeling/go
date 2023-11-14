package _02311

import (
	"fmt"
	"testing"
)

func TestStudy20231114(t *testing.T) {
	var x *struct {
		s [][32]byte
	}
	// 注意这里不是定义一个结构体类型,而是定义一个结构体类型指针变量,
	// 即 x 是一个指针,指针类型是一个匿名结构体
	// 很显然,x 的值是 nil,因为没有初始化,可以打印证实这一点
	fmt.Printf("%T", x)

	println(len(x.s[99])) // 32
}

// len的详解
// func len(v Type) int
/*
内建函数 len 返回 v 的长度，这取决于具体类型：
数组: v 中元素的数量
数组指针: *v 中元素的数量（v 为 nil 时 panic）
切片、map: v 中元素的数量；若 v 为nil，len(v) 即为零
字符串: v 中字节的数量
通道: 通道缓存中队列（未读取）元素的数量；若 v 为 nil，len(v) 即为零
*/

/*
在规范中,有一节是关于 len 和 cap 的。有如下几个要点：
返回结果总是 int；
返回结果有可能是常量；
有时对函数参数不求值，即编译期确定返回值；
*/

/*
如果 v 的类型是数组或指向数组的指针,且表达式 v 没有包含 channel 接收或（非常量）函数调,，则返回值也是一个常量
这种情况下,不会对 v 进行求值（即编译期就能确定）
否则返回值不是常量,且会对 v 进行求值（即得运行时确定）
*/

// 其他类似情况
// 不会panic
func other() {
	var testdata *struct {
		a *[7]int
	}
	for i, _ := range testdata.a {
		fmt.Println(i)
	}
}

// 会panic
func otherCanPanic() {
	var testdata *struct {
		a *[7]int
	}
	for i, j := range testdata.a {
		fmt.Println(i, j)
	}
}
