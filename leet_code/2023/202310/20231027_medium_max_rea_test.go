package main

import (
	"fmt"
	"sort"
	"testing"
)

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	return calMax(horizontalCuts, h) * calMax(verticalCuts, w) % 1000000007
}

func calMax(arr []int, boardr int) int {
	res, pre := 0, 0
	for _, i := range arr {
		res = max(i-pre, res)
		pre = i
	}
	return max(res, boardr-pre)
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

// 切割后面积最大的蛋糕

func TestMaxRea(t *testing.T) {
	fmt.Println(maxArea(5, 4, []int{1, 2, 4}, []int{1, 3}))
}

/*
矩形蛋糕的高度为 h 且宽度为 w，给你两个整数数组 horizontalCuts 和 verticalCuts，其中：
 horizontalCuts[i] 是从矩形蛋糕顶部到第  i 个水平切口的距离
verticalCuts[j] 是从矩形蛋糕的左侧到第 j 个竖直切口的距离
请你按数组 horizontalCuts 和 verticalCuts 中提供的水平和竖直位置切割后，请你找出 面积最大 的那份蛋糕，并返回其 面积 。由于答案可能是一个很大的数字，因此需要将结果 对 109 + 7 取余 后返回
*/
