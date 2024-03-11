package _024_01

import (
	"fmt"
	"testing"
)

/*
给定平面上 n 对 互不相同 的点 points ，其中 points[i] = [xi, yi]
回旋镖 是由点 (i, j, k) 表示的元组 ，其中 i 和 j 之间的距离和 i 和 k 之间的欧式距离相等（需要考虑元组的顺序）
返回平面上所有回旋镖的数量
*/

func numberOfBoomerangs(points [][]int) (ans int) {
	for _, p := range points {
		cnt := make(map[int]int, len(points))
		for _, q := range points {
			dis := (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])
			cnt[dis]++
		}
		for _, m := range cnt {
			ans += m * (m - 1)
		}
	}
	return
}

func Test20240108(t *testing.T) {
	var points = [][]int{{0, 0}, {1, 0}, {2, 0}}
	fmt.Println(numberOfBoomerangs(points))
}
