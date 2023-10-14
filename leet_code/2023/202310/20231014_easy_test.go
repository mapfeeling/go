package main

import (
	"fmt"
	"testing"
)

/*
你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素
你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间
*/
func singleNumber(nums []int) int {
	var box = make(map[int]struct{}, len(nums)/2)
	for _, val := range nums {
		if _, ok := box[val]; ok {
			delete(box, val)
		} else {
			box[val] = struct{}{}
		}
	}
	for k := range box {
		return k
	}
	return -1
}

// 0与其他任何数字异或都是其本身,非零数字与其本身异或为0
func singleNumberEasy(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}

func TestSingleNumber(t *testing.T) {
	var nums = []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
	fmt.Println(singleNumberEasy(nums))
}
