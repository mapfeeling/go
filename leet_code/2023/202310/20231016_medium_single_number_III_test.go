package main

import (
	"fmt"
	"testing"
)

// 给你一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次
// 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案

func mediumSingleNumber(nums []int) []int {
	var (
		box = make(map[int]int, len(nums)/2)
		res = make([]int, 0)
	)

	for _, v := range nums {
		box[v]++
	}

	for k, v := range box {
		if v != 1 {
			continue
		}
		res = append(res, k)
	}

	return res
}

/*
func singleNumber(nums []int) []int {
    sum:=0;
    for _,num := range nums{
        sum^=num;
    }

    a:=0;
    b:=0;
    mostRightOne:=sum&(^sum+1)
    //在mostRightOne位置上的数可以分为两类 0 1
    for _,num := range nums{
        if num&mostRightOne==0 {
            a^=num
        }else{
            b^=num
        }
    }
    return []int{a,b}
}
*/

func TestMediumSingleNumber(t *testing.T) {
	var nums = []int{1, 2, 1, 3, 2, 5}
	fmt.Println(mediumSingleNumber(nums))
}
