package prime

import (
	"fmt"
	"math/big"
	"testing"
)

func multi(str1, str2 string) (result string) {
	if len(str1) == 0 && len(str2) == 0 {
		result = "0"
		return
	}
	var (
		index1 = len(str1) - 1
		index2 = len(str2) - 2
		left   int
	)
	for index1 >= 0 && index2 >= 0 {
		c1 := str1[index1] - '0'
		c2 := str2[index2] - '0'
		sum := int(c1) + int(c2) + left
		if sum > 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
		index2--
	}

	for index1 >= 0 {
		c1 := str1[index1] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
	}

	for index2 >= 0 {
		c1 := str2[index2] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index2--
	}
	if left == 1 {
		result = fmt.Sprintf("1%s", result)
	}
	return

}

func TestMulti(t *testing.T) {
	println(multi("123", "123"))
	println(add("123", "123"))
}

func add(a, b string) string {
	bigA, _ := new(big.Int).SetString(a, 10)
	bigB, _ := new(big.Int).SetString(b, 10)
	sum := new(big.Int).Add(bigA, bigB)
	return sum.String()
}
