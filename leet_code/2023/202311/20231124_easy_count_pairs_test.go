package _02311

import (
	"fmt"
	"sort"
	"testing"
)

func countPairs(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] >= target {
				break
			}
			ans++
		}
	}
	return ans
}

func TestCountPairs(t *testing.T) {
	fmt.Println(countPairs([]int{-6, 2, 5, -2, -7, -1, 3}, 2))
}

/*
给你一个下标从 0 开始长度为 n 的整数数组 nums 和一个整数 target
请你返回满足 0 <= i < j < n 且 nums[i] + nums[j] < target 的下标对 (i, j) 的数目
*/
