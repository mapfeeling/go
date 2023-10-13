package leet_code

import (
	"fmt"
	"sort"
	"testing"
)

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
示例 1：
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]

示例 2：
输入：nums = []
输出：[]

示例 3：
输入：nums = [0]
输出：[]

示例4：
输入：nums=[-2,0,0,2,2]
输出：[[-2,0,2]]

示例5：
输入：[-1,0,1,2,-1,-4,-2,-3,3,0,4]
输出：[[-4,0,4],[-4,1,3],[-3,-1,4],[-3,0,3],[-3,1,2],[-2,-1,3],[-2,0,2],[-1,-1,2],[-1,0,1]]

提示：
0 <= nums.length <= 3000
-10^5 <= nums[i] <= 10^5
*/

func threeSum(nums []int) (res [][]int) {
	n := len(nums)
	if n < 3 {
		return
	}
	sort.Ints(nums)

	for left := 0; left < n; left++ {
		if left > 0 && nums[left] == nums[left-1] {
			continue
		}
		target := 0 - nums[left]
		right := n - 1
		for mid := left + 1; mid < right; mid++ {
			// 去重
			if mid > left+1 && nums[mid] == nums[mid-1] {
				continue
			}
			for mid < right && nums[mid]+nums[right] > target {
				right--
			}
			if mid != right && nums[mid]+nums[right] == target {
				res = append(res, []int{nums[left], nums[mid], nums[right]})
			}
		}
	}

	return
}
func TestNew(t *testing.T) {
	var nums = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

//输入：nums = [-1,0,1,2,-1,-4]
// -4,-1,-1,0,1,2
//输出：[[-1,-1,2],[-1,0,1]]
