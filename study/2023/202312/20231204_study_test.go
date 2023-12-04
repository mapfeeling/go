package _02312

import (
	"fmt"
	"testing"
)

func TestStudy20231204(t *testing.T) {
	i := 1
	s := []string{"A", "B", "C"}
	i, s[i-1] = 2, "Z"
	fmt.Printf("s: %v \n", s)
}

// 问,上面代码会输出什么？
/*
A. s: [Z,B,C]
B. s: [A,Z,C]
*/

/*
知识点：多重赋值。
多重赋值分为两个步骤，有先后顺序:
计算等号左边的索引表达式和取址表达式，接着计算等号右边的表达式；
赋值；
所以本例，会先计算 s[i-1]，等号右边是两个表达式是常量，所以赋值运算等同于 i, s[0] = 2, "Z"
*/
