package _02311

import (
	"fmt"
	"testing"
)

func maxSliceIntSum(nums []int) (ans int) {
	ans = nums[0]
	for index, num := range nums {
		var (
			currVal = num
			//minVal  = math.MinInt32
		)
		for j := index + 1; j < len(nums); j++ {
			currVal += nums[j]
			if currVal > ans {
				ans = currVal
			}
		}
	}
	return ans
}

func TestFindTheCity(t *testing.T) {
	// []int 连续数组的最大和 +- 1 or n
	var nums = []int{1, -2, 3, 10, -4, 7, 2, -5}
	fmt.Println(maxSliceIntSum(nums))
	s := "string"
	fmt.Println(len(s))
	ss := []rune(s)
	fmt.Println(ss, []byte(s))
	ss[0] = 116
	fmt.Println(string(ss))
	fmt.Println(ss)
	//fmt.Println(ss[])
}
