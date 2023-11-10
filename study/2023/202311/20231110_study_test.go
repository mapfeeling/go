package _02311

import (
	"fmt"
	"testing"
)

func TestStudy20231110(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	s2[1] = 4
	fmt.Println(s1) // 1,2,4
	s2 = append(s2, 5, 6, 7)
	fmt.Println(s1) // 1,2,4
}
