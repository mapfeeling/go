package main

import "fmt"

func SliceRise(s []int) {
	s = append(s, 0)
	for i := range s {
		s[i]++
	}
}

func main() {
	s1 := []int{1, 2}
	s2 := s1
	s2 = append(s2, 3)
	SliceRise(s1)
	SliceRise(s2)
	fmt.Println(s1, s2)
}

// 打印结果
// 这是一个及其经典的切片相关的题目
/*
首先 s1  s2 传到函数发生了值拷贝，所以 s1  s2 的 length 属性肯定不会变，所以 s1  的打印只能看到2个值；s2 的打印能看到 3 个值。
又因为虽然是值拷贝，但底层数组拷贝的是指针，但s1的副本进入函数之后因为初始容量不够扩容，所以引用的数组和s1 不同，所以s1打印[1,2],不会受副本++影响
s2 进入函数没有扩容，所以s2副本改变影响到s2 所以打印[2,3,4].
*/
