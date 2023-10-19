package _02310

import (
	"fmt"
	"testing"
)

// 下面代码下划线处可以填入哪个选项？

func Test20231019Study(t *testing.T) {
	var s1 []int
	//var s2 = []int{}
	// if __ == nil {
	if s1 == nil {
		fmt.Println("yes nil")
	} else {
		fmt.Println("no nil")
	}
}

/*
- A. s1
- B. s2
- C. s1、s2 都可以
*/

// 答案:A 因为nil切片等于nil,而空切片是一个空的切片集合
