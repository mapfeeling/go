package container

import (
	"container/ring"
	"fmt"
	"testing"
)

// 循环链表

func TestRing(t *testing.T) {
	// 初始化环
	r := ring.New(3)
	fmt.Println("环的长度:", r.Len())
	// 循环赋值
	for i := 1; i <= 3; i++ {
		r.Value = i
		r = r.Next()
	}
	// 对环进行循环操作
	r.Do(func(a any) {
		i := a.(int)
		a = i + 1
		fmt.Println(a)
	})

	{
		fmt.Println("r的打印", r.Value)
		r = r.Move(1)
		fmt.Println("r的打印", r.Value)
	}

	{
		fmt.Println("r的当前节点值", r.Value)
		r = r.Prev()
		fmt.Println("执行r.Prev()的当前节点值", r.Value)
	}

	{
		fmt.Println("r的当前节点值", r.Value)
		r = r.Next()
		fmt.Println("执行r.Next()的当前节点值", r.Value)
	}

	{
		// 连接两个环
		rAnother := ring.New(10)
		for i := 10; i <= 19; i++ {
			rAnother.Value = i
			rAnother = rAnother.Next()
		}
		r = r.Link(rAnother)
		r.Do(func(a any) {
			fmt.Println(a)
		})
	}

	{
		// 从当前元素开始,删除n个元素
		println("从当前元素开始,删除n个元素")
		r.Unlink(3) // 不能使用r = r.Unlink(3) 不然就是将删除的元素赋值给r
		r.Do(func(a any) {
			fmt.Println(a)
		})
	}
	{
		RingRangeMove()
	}

}

func RingRangeMove() {
	// 1-100的100数,每相隔2,删除该元素,问最后剩下的元素是什么？
	println(" 1-100的100数,每相隔2,删除该元素,问最后剩下的元素是什么？")
	r := ring.New(100)

	for i := 1; i <= 100; i++ {
		r.Value = i
		r = r.Next()
	}

	for r.Len() > 1 {
		r.Move(2)
		r.Unlink(1)
	}

	r.Do(func(a any) {
		fmt.Println(a)
	})
}
