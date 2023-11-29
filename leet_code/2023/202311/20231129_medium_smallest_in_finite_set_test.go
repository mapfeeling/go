package _02311

import (
	"container/heap"
)

type SmallestInfiniteSet struct {
	smallest int
	addbacks IntHeap
	dups     map[int]struct{}
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Constructor20231129() SmallestInfiniteSet {
	return SmallestInfiniteSet{1, make([]int, 0), make(map[int]struct{})}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if len(this.addbacks) > 0 && this.addbacks[0] < this.smallest {
		num := heap.Pop(&this.addbacks).(int)
		delete(this.dups, num)
		return num
	} else {
		num := this.smallest
		this.smallest++
		return num
	}
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if this.smallest <= num {
		return
	}
	if _, ok := this.dups[num]; !ok {
		heap.Push(&this.addbacks, num)
		this.dups[num] = struct{}{}
	}
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */

// 无限集合中的最小数字
/*
现有一个包含所有正整数的集合 [1, 2, 3, 4, 5, ...] 。
实现 SmallestInfiniteSet 类：
SmallestInfiniteSet() 初始化 SmallestInfiniteSet 对象以包含 所有 正整数
int popSmallest() 移除 并返回该无限集中的最小整数
void addBack(int num) 如果正整数 num 不 存在于无限集中，则将一个 num 添加 到该无限集最后
*/
