package main_test

import (
	"fmt"
	"sort"
	"testing"
)

func maxSatisfaction(satisfaction []int) int {
	sort.Slice(satisfaction, func(i, j int) bool {
		return satisfaction[i] > satisfaction[j]
	})
	fmt.Println(satisfaction)
	presum, ans := 0, 0
	for _, si := range satisfaction {
		if presum+si > 0 {
			presum += si
			fmt.Println(presum)
			ans += presum
		} else {
			break
		}
	}
	return ans
}

func Test20231023HardMaxSatisfaction(t *testing.T) {
	fmt.Println(maxSatisfaction([]int{-1, -8, 0, 5, -9}))
}

/*
一个厨师收集了他 n 道菜的满意程度 satisfaction ，这个厨师做出每道菜的时间都是 1 单位时间。
一道菜的 「 like-time 系数 」定义为烹饪这道菜结束的时间（包含之前每道菜所花费的时间）乘以这道菜的满意程度，也就是 time[i]*satisfaction[i] 。
返回厨师在准备了一定数量的菜肴后可以获得的最大 like-time 系数 总和。
你可以按 任意 顺序安排做菜的顺序，你也可以选择放弃做某些菜来获得更大的总和
*/
