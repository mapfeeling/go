package _02311

import (
	"fmt"
	"testing"
)

func TestStudy2023(t *testing.T) {
	var f = func() {
		defer fmt.Println("D")
		fmt.Println("F")
	}
	f()
	fmt.Println("M")
}

// F
// D
// M
