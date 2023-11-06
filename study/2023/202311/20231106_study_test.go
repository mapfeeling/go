package _02311

import (
	"fmt"
	"testing"
)

// 下面 A、B 两处应该填入什么代码，才能确保顺利打印出结果？

/*
type S struct {
	m string
}

func f() *S {
	return __  //A
}

func main() {
	p := __    //B
	fmt.Println(p.m) //print "foo"
}
*/

type SThis struct {
	m string
}

func fThis() *SThis {
	return &SThis{m: "foo"} //A
}

func TestStudy20231106(t *testing.T) {
	p := fThis()     // B  *f() 或者 f()
	fmt.Println(p.m) // print "foo"
}

/*
f() 函数返回参数是指针类型，所以可以用 & 取结构体的指针
B 处,如果填 *f(),则 p 是 S 类型；如果填 f(),则 p 是 *S 类型,不过都可以使用 p.m 取得结构体的成员
*/
