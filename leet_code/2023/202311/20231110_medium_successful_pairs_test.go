package _02311

import (
	"fmt"
	"sort"
	"testing"
)

// 咒语和药水的成功对数
func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	res := make([]int, len(spells))
	for i, spell := range spells {
		res[i] = len(potions) - sort.SearchInts(potions, (int(success)-1)/spell+1)
	}
	return res
}

func TestSuccessfulPairs(t *testing.T) {
	var (
		spells  = []int{5, 1, 3}
		potions = []int{1, 2, 3, 4, 5}
		success = int64(7)
	)
	fmt.Println(successfulPairs(spells, potions, success))
	fmt.Println(sort.Find())
}

/*
给你两个正整数数组 spells 和 potions ，长度分别为 n 和 m ，其中 spells[i] 表示第 i 个咒语的能量强度，potions[j] 表示第 j 瓶药水的能量强度
同时给你一个整数 success 。一个咒语和药水的能量强度 相乘 如果 大于等于 success ，那么它们视为一对 成功 的组合
请你返回一个长度为 n 的整数数组 pairs，其中 pairs[i] 是能跟第 i 个咒语成功组合的 药水 数目
*/
