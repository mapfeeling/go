package closure

import (
	"fmt"
	"testing"
)

func forEach(numbers []int, callback func(int)) {
	for _, num := range numbers {
		callback(num)
	}
}

// 回调函数（Callback Function）：闭包可以用作回调函数，在某些条件满足时执行

func TestCallback(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	forEach(numbers, func(num int) {
		fmt.Println("Number:", num)
	})
}
