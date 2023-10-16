package main

import (
	"fmt"
	"testing"
)

/*
给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次
请你找出并返回那个只出现了一次的元素
你必须设计并实现线性时间复杂度的算法且使用常数级空间来解决此问题
*/

func singleNumberMid(nums []int) int {
	var box = make(map[int]int, len(nums)/2)
	for _, v := range nums {
		box[v]++
	}
	for k, v := range box {
		if v == 1 {
			return k
		}
	}
	return -1
}

func TestMidSingleNumber(t *testing.T) {
	var nums = []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println(singleNumberMid(nums))
}
