package common

import (
	"fmt"
	"testing"
)

// 146 螺旋遍历二维数组
/*
给定一个二维数组 array，请返回「螺旋遍历」该数组的结果
螺旋遍历：从左上角开始，按照 向右、向下、向左、向上 的顺序 依次 提取元素，然后再进入内部一层重复相同的步骤，直到提取完所有元素
*/

func spiralArray(array [][]int) []int {
	if len(array) == 0 {
		return nil
	}
	var ans = make([]int, 0, len(array)*len(array[0]))
	top, bottom, left, right := 0, len(array)-1, 0, len(array[0])-1
	for bottom >= left && right >= left {
		for i := left; i <= right; i++ {
			ans = append(ans, array[top][i])
		}
		top++
		for i := top; i <= bottom; i++ {
			ans = append(ans, array[i][right])
		}
		right--
		if left > right || top > bottom {
			break
		}
		for i := right; i >= left; i-- {
			ans = append(ans, array[bottom][i])
		}
		bottom--
		for i := bottom; i > top; i-- {
			ans = append(ans, array[i][left])
		}
		left++
	}
	return ans
}

func TestSpiralArray(t *testing.T) {
	var arr = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralArray(arr))
}
