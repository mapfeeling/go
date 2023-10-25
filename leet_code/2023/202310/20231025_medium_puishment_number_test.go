package main

import (
	"fmt"
	"strconv"
	"testing"
)

func punishmentNumber(n int) int {
	var f func(string, int, int, int) bool
	f = func(s string, pos int, tot int, target int) bool {
		if pos == len(s) {
			return tot == target
		}
		sum := 0
		for i := pos; i < len(s); i++ {
			sum = sum*10 + int(s[i]-'0')
			if sum+tot > target {
				break
			}
			if f(s, i+1, sum+tot, target) {
				return true
			}
		}
		return false
	}
	res := 0
	for i := 1; i <= n; i++ {
		if f(strconv.Itoa(i*i), 0, 0, i) {
			res += i * i
		}
	}
	return res
}

func TestPunishmentNumber(t *testing.T) {
	fmt.Println(punishmentNumber(10))
}

//给你一个正整数 n ，请你返回 n 的 惩罚数
//n 的 惩罚数 定义为所有满足以下条件 i 的数的平方和
//1 <= i <= n
//i * i 的十进制表示的字符串可以分割成若干连续子字符串，且这些子字符串对应的整数值之和等于 i
