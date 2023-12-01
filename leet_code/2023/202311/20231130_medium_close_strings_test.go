package _02311

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func closeStrings(word1 string, word2 string) bool {
	n1, n2 := len(word1), len(word2)
	if n1 != n2 {
		return false
	}
	m := make(map[int]int, n1)
	for _, v := range word1 {
		k := int(v - 'a')
		if _, ok := m[k]; ok {
			m[k]++
		} else {
			m[k] = 1
		}
	}
	m2 := make(map[int]int, n1)
	for _, v := range word2 {
		k := int(v - 'a')
		if _, ok := m2[k]; ok {
			m2[k]++
		} else {
			m2[k] = 1
		}
	}
	var (
		s1 = make([]int, 0, len(m))
		s2 = make([]int, 0, len(m2))
	)
	for _, v := range m {
		s1 = append(s1, v)
	}
	for _, v := range m2 {
		s2 = append(s2, v)
	}
	sort.Ints(s1)
	sort.Ints(s2)
	fmt.Println(s1)
	fmt.Println(s2)
	for k, v := range s1 {
		if v != s2[k] {
			return false
		}
	}
	return true
}

func TestCloseStrings(t *testing.T) {
	goCacheKey := strings.Join([]string{strconv.FormatInt(23, 10), ":", strconv.FormatInt(2232323, 10)}, "")
	fmt.Println(goCacheKey)
	fmt.Println(closeStrings("uau", "ssx"))
}
