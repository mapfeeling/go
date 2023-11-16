package _02311

import (
	"testing"
)

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

// 问,下面代码正常输出什么？
func TestStudy20231116(t *testing.T) {
	//var peo People = Student{}
	//think := "speak"
	//fmt.Println(peo.Speak(think))
}

// 值类型会默认生成其相对应的指针接受者方法,但该题中并没有实现值接受者的方法
