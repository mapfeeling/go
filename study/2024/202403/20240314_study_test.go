package _02403_test

import (
	"fmt"
	"testing"
)

var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	//use p
	fmt.Println(*p)
}

func Test20240314(t *testing.T) {
	p, err := foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Println(*p)
}

// 空指针,main()中的p会重新定义
