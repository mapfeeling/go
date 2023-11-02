package closure

import (
	"fmt"
	"testing"
)

func addGenerator() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func TestFuncFactory(t *testing.T) {
	addFunc := addGenerator()
	fmt.Println(addFunc(1))
	fmt.Println(addFunc(2))
	fmt.Println(addFunc(3))
}

/*
addGenerator 函数返回了一个闭包函数,该闭包函数可以对一个变量 sum 进行累加操作
每次调用返回的闭包函数时,sum 的值都会保留,并在累加后返回结果
这样,我们可以通过 addGenerator 来创建多个独立的累加器
*/
