package _02310

import (
	"fmt"
	"testing"
)

//问下面输出什么？

func Test20231025(t *testing.T) {
	v := new(int)
	fmt.Println(*v)
	*v = 2
	println(5 / +-*v)
}

// +x  是 0+x
// -x  是 0-x
