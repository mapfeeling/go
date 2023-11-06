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
	p := fThis()     //B
	fmt.Println(p.m) //print "foo"
}
