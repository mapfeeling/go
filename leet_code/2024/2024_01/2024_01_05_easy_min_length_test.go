package _024_01

import (
	"fmt"
	"testing"
)

func minLength(s string) int {
	var box []byte
	for v := range s {
		box = append(box, s[v])
		for len(box) >= 2 && (string(box[len(box)-2:]) == "AB" || string(box[len(box)-2:]) == "CD") {
			box = box[:len(box)-2]
		}
	}
	return len(box)
}

func Test20241005MinLength(t *testing.T) {
	fmt.Println(minLength("ACBBD"))
}
