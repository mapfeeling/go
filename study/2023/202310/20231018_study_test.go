package _02310

import (
	"fmt"
	"testing"
)

func Test2023Study(t *testing.T) {
	var f1 = func() {
		a := []int{2: 1}
		fmt.Println(a)
	}
	var f2 = func() {
		var x = []int{4: 44, 55, 66, 1: 77, 88}
		println(len(x), x[2])
	}
	f1()
	f2()
}

// 问：f1和f2分别输出什么？
// f1-----> A：编译错误；B：[2 1]；C：[0 0 1]；D：[0 1]  ----> C
// f2-----> A：5 66；B：5 88；C：7 88；D：以上都不对     ----> C
