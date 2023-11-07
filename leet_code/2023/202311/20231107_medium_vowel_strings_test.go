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

func TestVowelStrings(t *testing.T) {
	words := []string{"hey", "aeo", "mu", "ooo", "artro"}
	words2 := []string{"are", "amy", "u"}
	left, right := 1, 4
	fmt.Println(vowelStrings(words, left, right))
	fmt.Println(vowelStringsAnother(words2, 0, 2))
}
