package _02310

import "fmt"

// 下面代码输出什么？

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func main() {
	fmt.Println(increaseA())
	fmt.Println(increaseB())
}

/*
A. 1 1
B. 0 1
C. 1 0
D. 0 0
*/

// 答案为B
