package main

import (
	"fmt"
	"strconv"
	"testing"
)

// 统计能整除数字的位数
func countDigits(num int) int {
	count := 0
	for _, v := range strconv.Itoa(num) {
		if num%int(v-'0') == 0 {
			count++
		}
	}
	return count
}

func TestCountDigits(t *testing.T) {
	num := 1248
	fmt.Println(countDigits(num))
}

/*
给你一个整数 num ，返回 num 中能整除 num 的数位的数目。
如果满足 nums % val == 0 ，则认为整数 val 可以整除 nums
*/
