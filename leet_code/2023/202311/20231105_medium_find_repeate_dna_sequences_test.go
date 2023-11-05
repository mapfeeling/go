package _02311

import (
	"fmt"
	"testing"
)

const L = 10

func findRepeatedDnaSequences(s string) []string {
	var ans []string
	cnt := map[string]int{}
	for i := 0; i <= len(s)-L; i++ {
		sub := s[i : i+L]
		cnt[sub]++
		if cnt[sub] == 2 {
			ans = append(ans, sub)
		}
	}
	return ans
}

// 重复的dna序列
func TestFindRepeatedDnaSequences(t *testing.T) {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	fmt.Println(findRepeatedDnaSequences(s))
}

/*
DNA序列 由一系列核苷酸组成,缩写为 'A', 'C', 'G' 和 'T'
例如，"ACGAATTCCG" 是一个 DNA序列
在研究 DNA 时，识别 DNA 中的重复序列非常有用
给定一个表示 DNA序列 的字符串 s ,返回所有在 DNA 分子中出现不止一次的 长度为 10 的序列(子字符串)
你可以按 任意顺序 返回答案
*/
