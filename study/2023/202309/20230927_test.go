package main

import (
	"fmt"
	"testing"
)

// 问,下面代码输出什么？

type person struct {
	name string
}

func Test20230927(t *testing.T) {
	var m map[person]int
	p := person{"mike"}
	fmt.Println(m[p])
}

// m 是一个 map,值是 nil,是一个nil map
// 从 nil map 中取值不会报错，而是返回相应的零值，这里值是 int 类型，因此返回 0
