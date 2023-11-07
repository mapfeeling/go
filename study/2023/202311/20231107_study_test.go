package _02311

import "testing"

// 问: 以下代码输出什么？
// A：不能编译；B：45；C：45.2；D：45.0
func TestStudy20231107(t *testing.T) {
	var ans float64 = 15 + 25 + 5.2
	println(ans)
}

/*
const a = 1 + 2 			// a == 3，是无类型常量
const b int8 = 1 + 2 	    // b == 3，是有类型常量，类型是 int8
// 而 1、2 这样的就是字面值常量
// a、b 这样的就是具名常量
*/
