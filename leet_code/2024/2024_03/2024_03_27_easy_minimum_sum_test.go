package _024_03

import (
	"fmt"
	"testing"
)

/*
给你一个下标从 0 开始的整数数组 nums
如果下标三元组 (i, j, k) 满足下述全部条件，则认为它是一个 山形三元组:
i < j < k
nums[i] < nums[j] 且 nums[k] < nums[j]
请你找出 nums 中 元素和最小 的山形三元组，并返回其 元素和 。如果不存在满足条件的三元组，返回 -1
nums = [5,4,8,7,10,2]
*/

// 方法一：暴力解法
func minimumSum(nums []int) int {
	min := -1
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				for k := j + 1; k < len(nums); k++ {
					if nums[k] < nums[j] {
						sum := nums[i] + nums[j] + nums[k]
						if min == -1 {
							min = sum
						} else {
							if min > sum {
								min = sum
							}
						}
					}
				}
			}
		}
	}
	return min
}

// 方法二：双指针向中间逼近
func minimumSumBySlice(nums []int) int {
	// left表示从左到右i前的最小的值
	var left = []int{-1}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			continue
		}

		left = append(left, nums[i])
	}
}

func TestMinimumSum(t *testing.T) {
	fmt.Println(minimumSum([]int{8, 6, 1, 5, 3}))
	fmt.Println(minimumSum([]int{5, 4, 8, 7, 10, 2}))
}
