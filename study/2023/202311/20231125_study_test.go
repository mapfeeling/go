package _02311

import (
	"fmt"
	"testing"
	"time"
)

// 下面这段代码输出什么？为什么？
func TestStudy20231125(t *testing.T) {
	var m = [...]int{1, 2, 3}

	for i, v := range m {
		go func() {
			fmt.Println(i, v)
		}()
	}

	time.Sleep(time.Second * 3)
}

/*输出:
2 3
2 3
2 3
*/

/*
for range 使用短变量声明 (:=) 的形式迭代变量
需要注意的是，变量 i、v 在每次循环体中都会被重用，而不是重新声明
各个 goroutine 中输出的 i、v 值都是 for range 循环结束后的 i、v 最终值
而不是各个 goroutine 启动时的 i, v值。可以理解为闭包引用，使用的是上下文环境的值
两种可行的 fix 方法
*/

/* a｜使用函数传递
for i, v := range m {
	go func(i,v int) {
		fmt.Println(i, v)
	}(i,v)
}
*/

/* b｜使用临时变量保留当前值
for i, v := range m {
	i := i           // 这里的 := 会重新声明变量，而不是重用
	v := v
	go func() {
		fmt.Println(i, v)
	}()
}
*/
