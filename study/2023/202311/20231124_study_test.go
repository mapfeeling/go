package _02311

import (
	"fmt"
	"testing"
)

// 下面这段代码能否正常结束？
func TestStudy20231124(t *testing.T) {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
	fmt.Println(v)
}

// 那这个呢？
func Study20231124() {
	v := []int{1, 2, 3}
	for i := 0; i < len(v); i++ {
		v = append(v, v[i])
		fmt.Println(len(v))
	}
}

// 无限循环

/*
会出现死循环，能正常结束
循环次数在循环开始前就已经确定，循环内改变切片的长度，不影响循环次数
*/
