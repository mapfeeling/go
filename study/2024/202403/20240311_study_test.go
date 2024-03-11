package _02403

import "fmt"

type Math struct {
	x, y int
}

var m = map[string]Math{
	"foo": Math{2, 3},
}

func main() {
	// m["foo"].x = 4
	// 为什么注释放开后,编译报错
	fmt.Println(m["foo"].x)
}

// 答案:
// 错误原因：对于类似 X = Y的赋值操作，必须知道 X 的地址，才能够将 Y 的值赋给 X，但 go 中的 map 的 value 本身是不可寻址的
// 两种解决办法:
// 1、使用临时变量
// tmp:= m["foo"]
// tmp.x = 4
// m["foo"] = tmp
// 2、修改结构体
// type Math struct {
//     x, y int
// }
// var m = map[string]*Math{
//	"foo": &Math{2, 3},
// }
