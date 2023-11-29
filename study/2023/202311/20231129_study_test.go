package _02311

import (
	"fmt"
	"testing"
)

func TestStudy20231129(t *testing.T) {
	var a = []int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r) // [1,12,13,4,5]
	fmt.Println("a = ", a) // [1,12,13,4,5]
}

/*
这的 a 是一个切片，那切片是怎么实现的呢？
切片在 go 的内部结构有一个指向底层数组的指针，当 range 表达式发生复制时，副本的指针依旧指向原底层数组，所以对切片的修改都会反应到底层数组上，所以通过 v 可以获得修改后的数组元素
*/
