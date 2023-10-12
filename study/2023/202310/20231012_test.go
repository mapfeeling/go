package _02310

import (
	"fmt"
	"testing"
)

func hello(i int) {
	fmt.Println(i)
}

func Test20231012(t *testing.T) {
	var i = 5
	defer hello(i)
	i = i + 10
}

// 以上代码输出什么？
// 5
// 太过于简单,不解释了
