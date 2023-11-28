package _02311

import (
	"fmt"
	"testing"
)

// 下面这段代码输出什么？
func change(s ...int) {
	s = append(s, 3)
}

func TestStudy20231128(t *testing.T) {
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	change(slice...)
	fmt.Println(slice)
	change(slice[0:2]...)
	fmt.Println(slice)
}

/*
[1,2,0,0,0]
[1,2,3,0,0]
*/
