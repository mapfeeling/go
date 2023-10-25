package main

import (
	"fmt"
	"testing"
)

func numRollsToTarget(n int, k int, target int) int {
	mod := int(1e9 + 7)

}

func TestNumRollsToTarget(t *testing.T) {
	var (
		n      = 2
		k      = 30
		target = 7
	)
	fmt.Println(numRollsToTarget(n, k, target))
}

/*
投骰子等于目标和的方法数
这里有 n 个一样的骰子，每个骰子上都有 k 个面，分别标号为 1 到 k
给定三个整数 n ,  k 和 target ，返回可能的方式(从总共 kn 种方式中)滚动骰子的数量，使正面朝上的数字之和等于 target
答案可能很大，你需要对 109 + 7 取模
*/
