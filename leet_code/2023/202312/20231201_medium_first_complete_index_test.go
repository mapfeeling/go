package _02312

import (
	"fmt"
	"testing"
)

func firstCompleteIndex(arr []int, mat [][]int) int {
	var (
		n, m     = len(mat), len(mat[0])
		mp       = make(map[int][2]int, n*m)
		row, col = make([]int, n, n), make([]int, m, m)
	)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			mp[mat[i][j]] = [2]int{i, j}
		}
	}

	for k, v := range arr {
		val := mp[v]
		row[val[0]]++
		if row[val[0]] == m {
			return k
		}
		col[val[1]]++
		if col[val[1]] == n {
			return k
		}
	}

	return -1
}

func TestFirstCompleteIndex(t *testing.T) {
	var (
		//arr = []int{2, 8, 7, 4, 1, 3, 5, 6, 9}
		arr = []int{6, 2, 3, 1, 4, 5}
		//mat = [][]int{{3, 2, 5}, {1, 4, 6}, {8, 7, 9}}
		mat = [][]int{{5, 1}, {2, 4}, {6, 3}}
	)
	fmt.Println(firstCompleteIndex(arr, mat))
}
