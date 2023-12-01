package _02312

import (
	"fmt"
	"testing"
)

func TestStudy20231201(t *testing.T) {
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)
}

// map遍历是无须的,所以可能的输出为2 or 3
