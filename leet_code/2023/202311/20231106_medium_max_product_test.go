package _02311

import (
	"fmt"
	"testing"
)

func maxProduct(words []string) int {
	res := 0
	var marks = make([]int, len(words), len(words))
	for index, word := range words {
		for _, char := range word {
			marks[index] |= 1 << (char - 'a')
		}
	}
	for i, x := range marks {
		for j, y := range marks[:i] {
			if x&y != 0 {
				continue
			}

			if box := len(words[i]) * len(words[j]); box > res {
				res = box
			}

			//if x&y == 0 && len(words[i])*(len(words[j])) > res {
			//	res = len(words[i]) * (len(words[j]))
			//}
		}
	}
	return res
}

// 最大单词长度乘积
func TestMaxProduct(t *testing.T) {
	words := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	fmt.Println(maxProduct(words))
	fmt.Println(1 << 2)
}

// 给你一个字符串数组 words ,找出并返回 length(words[i]) * length(words[j]) 的最大值,并且这两个单词不含有公共字母
// 如果不存在这样的两个单词, 返回 0

// 与运算： 0 & 0 =0  0 & 1 =0  1 & 0 =0 1 & 1 =1
// 或运算： 1｜1 =1  1｜0 =1  0｜0 =0  0｜1=1
// 异或运算：0 ⊕ 0 = 0 1 ⊕ 0 = 1 0 ⊕ 1 = 1 1 ⊕ 1 = 0（同为0,异为1）
