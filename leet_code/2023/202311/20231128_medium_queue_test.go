package _02311

import (
	"fmt"
	"testing"
)

type FrontMiddleBackQueue struct {
	data []int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{
		data: make([]int, 0, 2),
	}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.data = append([]int{val}, this.data...)
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	middle := len(this.data) / 2
	box := append([]int{val}, this.data[middle:]...)
	this.data = append(this.data[:middle], box...)
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.data = append(this.data, val)
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if len(this.data) == 0 {
		return -1
	}
	ans := this.data[0]
	this.data = this.data[1:]
	return ans
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	n := len(this.data)
	if n == 0 {
		return -1
	}
	var middle int
	if n%2 == 1 {
		middle = n / 2
	} else {
		middle = n/2 - 1
	}

	ans := this.data[middle]
	this.data = append(this.data[:middle], this.data[middle+1:]...)
	return ans
}

func (this *FrontMiddleBackQueue) PopBack() int {
	n := len(this.data)
	if n == 0 {
		return -1
	}
	ans := this.data[n-1]
	this.data = this.data[:n-1]
	return ans
}

func TestFrontMiddleBackQueue(t *testing.T) {
	q := Constructor()
	fmt.Println("1", q.data)
	q.PushFront(1) // [1]
	fmt.Println("2", q.data)
	q.PushBack(2) // [1, 2]
	fmt.Println("3", q.data)
	q.PushMiddle(3) // [1, 3, 2]
	fmt.Println("4", q.data)
	q.PushMiddle(4) // [1, 4, 3, 2]
	fmt.Println("5", q.data)
	q.PopFront() // 返回 1 -> [4, 3, 2]
	fmt.Println("6", q.data)
	q.PopMiddle() // 返回 3 -> [4, 2]
	fmt.Println("7", q.data)
	q.PopMiddle() // 返回 4 -> [2]
	fmt.Println("8", q.data)
	q.PopBack() // 返回 2 -> []
	fmt.Println("9", q.data)
	q.PopBack() // 返回 -1 -> [] （队列为空）
	fmt.Println("10", q.data)
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
