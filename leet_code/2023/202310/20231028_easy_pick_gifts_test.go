package main

import (
	"container/heap"
	"fmt"
	"math"
	"testing"
)

type Box []int

func (b Box) Swap(i, j int) {
	//TODO implement me
	b[i], b[j] = b[j], b[i]
}

func (b *Box) Push(x any) {
	*b = append(*b, x.(int))
}

func (b *Box) Pop() any {
	n := len(*b)
	res := (*b)[n-1]
	*b = (*b)[0 : n-1]
	return res
}

func (b Box) Len() int {
	return len(b)
}

func (b Box) Less(i, j int) bool {
	return b[i] > b[j]
}

func pickGifts(gifts []int, k int) int64 {
	var box = &Box{}
	heap.Init(box)
	for _, gift := range gifts {
		box.Push(gift)
	}
	heap.Init(box)
	for k > 0 {
		cur := heap.Pop(box).(int)
		heap.Push(box, int(math.Floor(math.Sqrt(float64(cur)))))
		k--
	}
	var res int64
	for box.Len() > 0 {
		res += int64(box.Pop().(int))
	}
	return res
}

func TestPickGifts(t *testing.T) {
	fmt.Println(pickGifts([]int{25, 64, 9, 4, 100}, 4))
}

/*
给你一个整数数组 gifts ,表示各堆礼物的数量。每一秒，你需要执行以下操作：
选择礼物数量最多的那一堆
如果不止一堆都符合礼物数量最多，从中选择任一堆即可
选中的那一堆留下平方根数量的礼物（向下取整），取走其他的礼物
返回在 k 秒后剩下的礼物数量
*/
