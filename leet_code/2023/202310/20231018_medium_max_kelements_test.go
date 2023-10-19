package main

import (
	"fmt"
	"sort"
	"testing"
)

// 执行k次操作后的最大分数
// 降序排序数组，维护一个大根堆
func heapify(arr []int, k, i int) {
	dummy := i
	left, right := i*2+1, i*2+2
	if left < k && arr[left] > arr[dummy] {
		dummy = left
	}
	if right < k && arr[right] > arr[dummy] {
		dummy = right
	}
	if i != dummy {
		arr[i], arr[dummy] = arr[dummy], arr[i]
		heapify(arr, k, dummy)
	}
}

func maxKelements(nums []int, k int) int64 {
	var (
		ret int
		f   = func(nums []int) {
			sort.Slice(nums, func(i, j int) bool {
				if nums[i] > nums[j] {
					return true
				}
				return false
			})
		}
	)
	f(nums)

	for i := 0; i < k; i++ {
		ret += nums[0]
		var dummy int
		if nums[0]%3 == 0 {
			dummy = nums[0] / 3
		} else {
			dummy = nums[0]/3 + 1
		}
		nums[0] = dummy
		heapify(nums, len(nums), 0)
	}

	return int64(ret)
}

func TestMaxKelements(t *testing.T) {
	fmt.Println(maxKelements([]int{1, 10, 3, 3, 3}, 3))
}

/*
给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。你的 起始分数 为 0 。
在一步 操作 中：
选出一个满足 0 <= i < nums.length 的下标 i ，
将你的 分数 增加 nums[i] ，并且
将 nums[i] 替换为 ceil(nums[i] / 3) 。
返回在 恰好 执行 k 次操作后，你可能获得的最大分数。
向上取整函数 ceil(val) 的结果是大于或等于 val 的最小整数
*/
