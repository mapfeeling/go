package _02312

import (
	"fmt"
	"testing"
)

func isWinner(player1 []int, player2 []int) int {
	s1, s2 := score(player1), score(player2)
	if s1 == s2 {
		return 0
	} else if s1 > s2 {
		return 1
	} else {
		return 2
	}
}

func score(player []int) int {
	res := 0
	for i, x := range player {
		if i > 0 && player[i-1] == 10 || i > 1 && player[i-2] == 10 {
			res += 2 * x
		} else {
			res += x
		}
	}
	return res
}

func TestIsWinner(t *testing.T) {
	player1 := []int{5, 6, 1, 10}
	player2 := []int{5, 1, 10, 5}
	fmt.Println(isWinner(player1, player2))
}
