package _02311

import (
	"fmt"
	"testing"
)

// 以下代码输出什么？
func TestStudy20231112(t *testing.T) {
	m := map[int]string{0: "zero", 1: "one"}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

/*
1 one
0 zero
*/
// 或者
/*
0 zero
1 one
*/
