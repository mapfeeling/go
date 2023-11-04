package _02311

import (
	"fmt"
	"testing"
)

type Person20231103 struct {
	age int
}

func Reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Test20231103Study(t *testing.T) {
	person := &Person20231103{28}

	// 1.
	defer fmt.Println(person.age) // 28

	// 2.
	defer func(p *Person20231103) {
		fmt.Println(p.age)
	}(person)

	// 3.
	defer func() {
		fmt.Println(person.age)
	}()

	person = &Person20231103{29}

}

// 28
// 29
// 28
