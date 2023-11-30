package _02311

import (
	"fmt"
	"testing"
)

type Foo struct {
	bar string
}

func TestStudy20231130(t *testing.T) {
	s1 := []Foo{
		{"A"},
		{"B"},
		{"C"},
	}
	s2 := make([]*Foo, len(s1))
	for i, value := range s1 {
		s2[i] = &value
	}
	fmt.Println(s1[0], s1[1], s1[2])
	fmt.Println(s2[0], s2[1], s2[2])
}

/*
s2 的输出是 &{C} &{C} &{C}，在前面题目我们提到过，for range 使用短变量声明(:=)的形式迭代变量时
变量 i、value 在每次循环体中都会被重用，而不是重新声明
所以 s2 每次填充的都是临时变量 value 的地址，而在最后一次循环中，value 被赋值为{c}。因此，s2 输出的时候显示出了三个 &{c}
*/
