package main

import (
	"fmt"
	"testing"
)

/*
给你一个整数 n ，表示一张 无向图 中有 n 个节点，编号为 0 到 n - 1
同时给你一个二维整数数组 edges ，其中 edges[i] = [ai, bi] 表示节点 ai 和 bi 之间有一条 无向 边。
请你返回 无法互相到达 的不同 点对数目
*/

type UnionFind struct {
	Parents []int
	Sizes   []int
}

func NewUnionFind(n int) *UnionFind {
	parents, sizes := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		parents[i], sizes[i] = i, 1
	}
	return &UnionFind{
		Parents: parents,
		Sizes:   sizes,
	}
}

func (uf *UnionFind) Find(x int) int {
	if uf.Parents[x] == x {
		return x
	}
	uf.Parents[x] = uf.Find(uf.Parents[x])
	return uf.Parents[x]
}

func (uf *UnionFind) Union(x, y int) {
	rx, ry := uf.Find(x), uf.Find(y)
	if rx != ry {
		if uf.Sizes[x] > uf.Sizes[y] {
			uf.Parents[ry], uf.Parents[rx] = rx, uf.Sizes[rx]+uf.Sizes[ry]
		} else {
			uf.Parents[rx], uf.Sizes[ry] = ry, uf.Sizes[rx]+uf.Sizes[ry]
		}
	}
}

func (uf *UnionFind) GetSize(x int) int {
	return uf.Sizes[x]
}

func countPairs(n int, edges [][]int) int64 {
	uf := NewUnionFind(n)
	for _, edge := range edges {
		uf.Union(edge[0], edge[1])
	}
	var res int64
	for i := 0; i < n; i++ {
		res += int64(n - uf.GetSize(uf.Find(i)))
	}
	return res / 2

}

func TestCountPairs(t *testing.T) {
	fmt.Println(countPairs(7, [][]int{{0, 2}, {0, 5}, {2, 4}, {1, 6}, {5, 4}}))
}
