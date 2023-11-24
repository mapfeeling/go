package _02311

import (
	"fmt"
	"sort"
	"testing"
)

// 2216、美化数组的最少删除数

func minDeletion(nums []int) int {
	n, ans, check := len(nums), 0, true
	for i := 0; i+1 < n; i++ {
		if nums[i] == nums[i+1] && check {
			ans++
		} else {
			check = !check
		}
	}
	if (n-ans)%2 != 0 {
		ans++
	}
	return ans
}

func TestMinDeletion(t *testing.T) {
	fmt.Println(minDeletion([]int{}))
	var s = []string{"adce", "abc", "ab", "abvsd", "z", "w"}
	for i := 0; i < len(s); i++ {
		sort.Slice(s, func(a, b int) bool {
			n := min(len(s[a]), len(s[b]))
			for i := 0; i < n; i++ {
				if s[a][i]-s[b][i] > 0 {
					return false
				}
			}
			return true
		})
	}
	fmt.Println(s)
}

/*
给你一个下标从 0 开始的整数数组 nums ，如果满足下述条件，则认为数组 nums 是一个 美丽数组 :
nums.length 为偶数
对所有满足 i % 2 == 0 的下标 i ，nums[i] != nums[i + 1] 均成立
注意，空数组同样认为是美丽数组
你可以从 nums 中删除任意数量的元素。当你删除一个元素时，被删除元素右侧的所有元素将会向左移动一个单位以填补空缺，而左侧的元素将会保持 不变
返回使 nums 变为美丽数组所需删除的 最少 元素数目
*/
