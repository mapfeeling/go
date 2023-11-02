package closure

import (
	"fmt"
	"testing"
)

func newCounter() func() int {
	count := 0
	return func() int {
		count++
		fmt.Println(count)
		return count
	}
}

// 常见的闭包创建方式就是在一个函数内部创建另一个函数,通过另一个函数访问这个函数的局部变量

func TestClosureIntroduction(t *testing.T) {
	f := newCounter()
	f()
	f()
	f1 := newCounter()
	f1()
	f1()
	f1()
}

/*
当调用newCounter()时,newCounter函数执行完毕退出,但由于匿名函数f存储了newCounter中的变量count
所以count并没有销毁，而是被封装在了函数f中。因此,当你通过f()调用函数f时,它还可以访问和修改count

这就是闭包的特性,通过内部的函数的方式获取其所在函数的引用环境的变量和参数访问和修改权限
在本质上,闭包是将函数内部和函数外部连接起来的桥梁
*/

/*
闭包产生的条件:
1、在函数 A 内部直接或者间接返回一个函数 B
2、B 函数内部使用着 A 函数的私有变量(私有数据)
3、A 函数外部有一个变量接受着函数 B
*/

// https://zhuanlan.zhihu.com/p/645853924
