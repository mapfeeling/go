package _02311

import (
	"container/list"
	"fmt"
	"testing"
)

type FrontMiddleBackQueueInLink struct {
	left  *list.List
	right *list.List
}

func ConstructorInLink() FrontMiddleBackQueueInLink {
	return FrontMiddleBackQueueInLink{
		left:  list.New(),
		right: list.New(),
	}
}

func (this *FrontMiddleBackQueueInLink) PushFront(val int) {
	this.left.PushFront(val)
	if this.left.Len() > this.right.Len() {
		this.right.PushFront(this.left.Remove(this.left.Back()))
	}
}

func (this *FrontMiddleBackQueueInLink) PushMiddle(val int) {
	this.left.PushBack(val)
}

func (this *FrontMiddleBackQueueInLink) PushBack(val int) {
	this.right.Remove(this.right.Back())
	if this.left.Len() > this.right.Len() {
		this.right.PushFront(this.left.Remove(this.left.Back()))
	}
}

func (this *FrontMiddleBackQueueInLink) PopFront() int {
	ans := this.left.Remove(this.left.Front())
	if this.left.Len()+2 == this.left.Len() {
		this.left.PushBack(this.right.Remove(this.right.Front()))
	}
	return ans.(int)
}

func (this *FrontMiddleBackQueueInLink) PopMiddle() int {
	ans := this.left.Remove(this.left.Back())
	if this.left.Len()+2 == this.left.Len() {
		this.left.PushBack(this.right.Remove(this.right.Front()))
	}
	return ans.(int)
}

func (this *FrontMiddleBackQueueInLink) PopBack() int {
	ans := this.right.Remove(this.right.Back())
	if this.left.Len() > this.right.Len() {
		this.right.PushBack(this.left.Remove(this.left.Back()))
	}
	return ans.(int)
}

func TestFrontMiddleBackQueueInLink(t *testing.T) {
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
