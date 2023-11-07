package _02311

import (
	"fmt"
	"testing"
)

func vowelStrings(words []string, left int, right int) int {
	var (
		f = func(i uint8) bool {
			if i == 'a' || i == 'e' || i == 'i' || i == 'o' || i == 'u' {
				return true
			}
			return false
		}
		ans = 0
	)

	for i := left; i <= right && i < len(words); i++ {
		if f(words[i][0]) && f(words[i][len(words[i])-1]) {
			ans++
		}
	}
	return ans
}

func vowelStringsAnother(words []string, left int, right int) int {
	m := map[byte]struct{}{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}}
	ans := 0
	for i := left; i <= right; i++ {
		if _, ok := m[words[i][0]]; ok {
			if _, okk := m[words[i][len(words[i])-1]]; okk {
				ans++
			}
		}

	}
	return ans
}

// 统计范围内的元音字符串数
func TestVowelStrings(t *testing.T) {
	words := []string{"hey", "aeo", "mu", "ooo", "artro"}
	words2 := []string{"are", "amy", "u"}
	left, right := 1, 4
	fmt.Println(vowelStrings(words, left, right))
	fmt.Println(vowelStringsAnother(words2, 0, 2))
}

/*
给你一个下标从 0 开始的字符串数组 words 和两个整数：left 和 right
如果字符串以元音字母开头并以元音字母结尾，那么该字符串就是一个 元音字符串 ,其中元音字母是 'a'、'e'、'i'、'o'、'u'
返回 words[i] 是元音字符串的数目,其中 i 在闭区间 [left, right] 内
*/
