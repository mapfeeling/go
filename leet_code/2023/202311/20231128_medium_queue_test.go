package _02311

import (
	"container/list"
	"fmt"
	"testing"
)

type FrontMiddleBackQueue struct {
	left  *list.List
	right *list.List
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{
		left:  list.New(),
		right: list.New(),
	}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.left.PushFront(val)
	if this.left.Len() > this.right.Len() {
		this.right.PushFront(this.left.Remove(this.left.Back()))
	}
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	this.left.PushBack(val)
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.right.Remove(this.right.Back())
	if this.left.Len() > this.right.Len() {
		this.right.PushFront(this.left.Remove(this.left.Back()))
	}
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if this.left.Len()+this.right.Len() == 0 {
		return -1
	}
	ans := this.left.Remove(this.left.Front())
	if this.left.Len()+2 == this.left.Len() {
		this.left.PushBack(this.right.Remove(this.right.Front()))
	}
	return ans.(int)
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.left.Len()+this.right.Len() == 0 {
		return -1
	}
	ans := this.left.Remove(this.left.Back())
	if this.left.Len()+2 == this.left.Len() {
		this.left.PushBack(this.right.Remove(this.right.Front()))
	}
	return ans.(int)
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if this.left.Len()+this.right.Len() == 0 {
		return -1
	}
	ans := this.right.Remove(this.right.Back())
	if this.left.Len() > this.right.Len() {
		this.right.PushBack(this.left.Remove(this.left.Back()))
	}
	return ans.(int)
}

func TestFrontMiddleBackQueue(t *testing.T) {
	q := Constructor()
	fmt.Println("1", q.left, q.right)
	q.PushFront(1) // [1]
	fmt.Println("2", q.left, q.right)
	q.PushBack(2) // [1, 2]
	fmt.Println("3", q.left, q.right)
	q.PushMiddle(3) // [1, 3, 2]
	fmt.Println("4", q.left, q.right)
	q.PushMiddle(4) // [1, 4, 3, 2]
	fmt.Println("5", q.left, q.right)
	q.PopFront() // 返回 1 -> [4, 3, 2]
	fmt.Println("6", q.left, q.right)
	q.PopMiddle() // 返回 3 -> [4, 2]
	fmt.Println("7", q.left, q.right)
	q.PopMiddle() // 返回 4 -> [2]
	fmt.Println("8", q.left, q.right)
	q.PopBack() // 返回 2 -> []
	fmt.Println("9", q.left, q.right)
	q.PopBack() // 返回 -1 -> [] （队列为空）
	fmt.Println("10", q.left, q.right)
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
