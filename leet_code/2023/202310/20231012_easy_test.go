package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func findTheArrayConcVal(nums []int) int64 {
	var res int
	if len(nums) < 2 {
		for _, val := range nums {
			res += val
		}
	} else {
		var (
			box = make([]int, 0, len(nums)/2+1)
			f   = func(i, j int) (int, error) {
				return strconv.Atoi(strings.Join([]string{strconv.Itoa(i), strconv.Itoa(j)}, ""))
			}
		)
		for {
			if len(nums) == 0 {
				break
			}
			if len(nums) == 1 {
				box = append(box, nums[0])
				break
			}
			if len(nums) == 2 {
				res, _ = f(nums[0], nums[1])
				break
			}
			a, b := nums[:1], nums[len(nums)-1:]
			nums = nums[1 : len(nums)-1]
			atoi, _ := f(a[0], b[0])
			box = append(box, atoi)
		}

		for _, v := range box {
			res += v
		}
	}

	return int64(res)
}

func Test20231012(t *testing.T) {
	var nums = []int{5, 14, 13, 8, 12}
	var nums2 = []int{92, 14}
	fmt.Println(findTheArrayConcVal(nums))
	fmt.Println(findTheArrayConcVal(nums2))
}

/*
给你一个下标从 0 开始的整数数组 nums
现定义两个数字的 串联 是由这两个数值串联起来形成的新数字
例如，15 和 49 的串联是 1549
nums 的 串联值 最初等于 0 执行下述操作直到 nums 变为空：
如果 nums 中存在不止一个数字，分别选中 nums 中的第一个元素和最后一个元素，将二者串联得到的值加到 nums 的 串联值 上，然后从 nums 中删除第一个和最后一个元素
如果仅存在一个元素，则将该元素的值加到 nums 的串联值上，然后删除这个元素
返回执行完所有操作后 nums 的串联值
*/
