package prime

import (
	"fmt"
	"testing"
)

// go的return和defer的执行顺序
func TestDeferAndReturn(t *testing.T) {
	fmt.Println("return", test1())
	fmt.Println("return", test2())
}

func test1() (i int) {
	i = 0
	defer func() {
		fmt.Println("defer1")
	}()
	defer func() {
		i += 1
		fmt.Println("defer2")
	}()
	return
}

func test2() int {
	i := 0
	defer func() {
		fmt.Println("defer1")
	}()
	defer func() {
		i += 1
		fmt.Println("defer2")
	}()
	return i
}

/*
return的正确执行操作:
1、return前将返回值赋值
2、检查是否有defer并执行
3、return携带返回值并退出函数
*/
