package _02311

import (
	"fmt"
	"testing"
)

type Myint int

func (i Myint) PrintInt() {
	fmt.Println(i)
}

//func (i int) PrintInt ()  {
//	fmt.Println(i)
//}

// 基于类型创建的方法必须定义在同一个包内,上面的代码基于 int 类型创建了 PrintInt() 方法,由于 int 类型和方法 PrintInt() 定义在不同的包内,所以编译出错
// 解决的办法可以定义一种新的类型
// compilation error 基础类型不能绑定方法
func TestStudy20231115(t *testing.T) {
	// var i int = 1
	//i.PrintInt()
	var i Myint = 1
	i.PrintInt()
}
