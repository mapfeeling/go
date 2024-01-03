package _02401

import (
	"fmt"
	"testing"
)

// 以下代码有什么问题，说明原因？

type student struct {
	Name string
	Age  int
}

func Test20240103Study(t *testing.T) {
	//定义map
	m := make(map[string]*student)

	//定义student数组
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	//将数组依次添加到map中
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	//打印map
	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}
}

/*
map 中的 3 个 key 均指向数组中最后一个结构体
分析
foreach 中，stu 是结构体的一个拷贝副本，所以m[stu.Name]=&stu实际上一致指向同一个指针， 最终该指针的值为遍历的最后一个struct的值拷贝
*/

// 正确写法:
/*
1、
for _, stu := range stus {
		s := stu
		m[stu.Name] = &s
	}
2、
// 遍历结构体数组，依次赋值给map
    for i := 0; i < len(stus); i++  {
        m[stus[i].Name] = &stus[i]
    }
*/
