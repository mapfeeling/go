package stack

import (
	"fmt"
	"testing"
)

// 使用切片实现一个栈,并实现pop和push功能

type Stack struct {
	data []int
}

// 入栈
func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}

// 出栈
func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top, true
}

// 获取栈顶元素
func (s *Stack) Top() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	return s.data[len(s.data)-1], true
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func TestSliceToStack(t *testing.T) {
	var stack = &Stack{
		data: make([]int, 0),
	}
	for i := 1; i <= 5; i++ {
		stack.Push(i)
	}
	for {
		if val, ok := stack.Pop(); ok {
			fmt.Println("出栈 ==>：", val)
		}
	}
}
