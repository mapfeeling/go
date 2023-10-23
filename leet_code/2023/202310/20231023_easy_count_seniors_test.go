package main

import (
	"fmt"
	"strconv"
	"testing"
)

func countSeniors(details []string) int {
	var count int
	for _, value := range details {
		if len(value) == 15 {
			atoI, _ := strconv.Atoi(value[11:13])
			if atoI > 60 {
				count += 1
			}
		}
	}
	return count
}

func countSeniorsUpdate(details []string) int {
	var count int
	for _, value := range details {
		age := 10*int(value[11]-'0') + int(value[12]-'0')
		if age > 60 {
			count += 1
		}
	}
	return count
}

func TestCountSeniors(t *testing.T) {
	var s = "7868190130M7522"
	fmt.Println(s[11:13])
}

/*
给你一个下标从 0 开始的字符串 details
details 中每个元素都是一位乘客的信息，信息用长度为 15 的字符串表示，表示方式如下:
前十个字符是乘客的手机号码
接下来的一个字符是乘客的性别
接下来两个字符是乘客的年龄
最后两个字符是乘客的座位号
请你返回乘客中年龄 严格大于 60 岁 的人数
*/
