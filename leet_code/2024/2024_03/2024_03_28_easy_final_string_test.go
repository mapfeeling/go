package _024_03

import (
	"fmt"
	"testing"
)

/*
你的笔记本键盘存在故障，每当你在上面输入字符 'i' 时，它会反转你所写的字符串。而输入其他字符则可以正常工作
给你一个下标从 0 开始的字符串 s ，请你用故障键盘依次输入每个字符
返回最终笔记本屏幕上输出的字符串
*/

// 2810故障键盘
func TestFinalString(t *testing.T) {
	fmt.Println(finalString("string"))
}

func finalString(s string) string {

	var (
		queue       []rune
		willReverse bool // 转置
	)
	for _, c := range s {
		if c != 'i' {
			queue = append(queue, c)
		} else {
			willReverse = !willReverse
			reverse(queue)
		}
	}

	return string(queue)
}

func reverse(arr []rune) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
