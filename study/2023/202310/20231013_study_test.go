package _02310

import (
	"fmt"
	"testing"
)

// 下面这段代码输出什么？
type PeopleThis struct{}

func (p *PeopleThis) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *PeopleThis) ShowB() {
	fmt.Println("showB")
}

type TeacherTHis struct {
	PeopleThis
}

func (t *TeacherTHis) ShowB() {
	fmt.Println("teacher showB")
}

func Test20231013Study(t *testing.T) {
	teacher := TeacherTHis{}
	teacher.ShowA()
}

//showA
//showB
//知识点：结构体嵌套
//这道题可以结合第 12 天的第三题一起看
//Teacher 没有自己 ShowA()，所以调用内部类型 People 的同名方法，需要注意的是第 5 行代码调用的是 People 自己的 ShowB 方法
