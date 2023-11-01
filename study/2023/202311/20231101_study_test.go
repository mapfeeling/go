package _02311

import (
	"fmt"
	"testing"
)

type Person struct {
	age int
}

func TestStudy20231101(t *testing.T) {
	person := &Person{28}

	// 1.
	defer fmt.Println(person.age)

	// 2. 匿名函数 - person的指针
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	// 3. 闭包引用
	defer func() {
		fmt.Println(person.age)
	}()

	person.age = 29
}

/*
输出:
29 29 28
*/

/*
变量 person 是一个指针变量
1.person.age 此时是将 28 当做 defer 函数的参数，会把 28 缓存在栈中，等到最后执行该 defer 语句的时候取出，即输出 28；
2.defer 缓存的是结构体 Person{28} 的地址，最终 Person{28} 的 age 被重新赋值为 29，所以 defer 语句最后执行的时候，依靠缓存的地址取出的 age 便是 29，即输出 29；
3.很简单，闭包引用，输出 29
又由于 defer 的执行顺序为先进后出，即 3 2 1，所以输出 29 29 28
*/
