package main

import (
	"fmt"
	"math"
	"sort"
)

// 最长递增子序列
// nums = [10,9,2,5,1,2,3,7,101,18]
// 输出: 4
// nums = [0,1,0,3,2,3]
// 输出: 4

func lengthOfLIS(nums []int) int {
	// minNum[k] 表示 LIS 的长度为 k 时，最后一个数字的最小值。
	// 初始 LIS 的长度为 0 ，最后一个数的最小值为 MIN ，方便后续处理。
	minNum := []int{math.MinInt32}
	fmt.Println(minNum)
	// 遍历每个数
	for _, num := range nums {
		// 寻找 minNum 中第一个大于等于 num 的下标 k ，
		// 则说明以 num 为 LIS 的最后一个数时，对应的 LIS 的长度最长为 k
		k := sort.Search(len(minNum), func(i int) bool { return minNum[i] >= num })
		fmt.Println(k)
		if len(minNum) == k {
			// num 是第一个使 LIS 长度达到 k 的数，所以直接放入 minNum
			minNum = append(minNum, num)
		} else {
			// 此时存在长度为 k 的 LIS ，
			// 因为前面二分寻找的是第一个大于等于 num 的下标，
			// 所以 minNum[k] >= num ，可以直接更新为 num
			minNum[k] = num
		}
		fmt.Println(minNum)
	}

	// LIS 的长度最长为 len(minNum) - 1
	return len(minNum) - 1
}

func main() {
	var SolutionLengthOfLisSlice = []int{10, 9, 2, 5, 1, 2, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(SolutionLengthOfLisSlice))
}
