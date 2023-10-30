package _02310

import "testing"

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func Test20231030(t *testing.T) {
	println(f1()) // 1
	println(f2()) // 5
	println(f3()) // 1
}

/*
go的return不是原子操作,分为两步执行: 1、对返回值赋值  2、真正的return返回
而defer操作在1和2操作之间执行
*/
