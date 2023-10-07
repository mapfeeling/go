package main

import (
	"fmt"
	"sync"
)

/*
golang经典题目:
一个字符串里面存在阿拉伯数字和大小写的英文字母,请实现阿拉伯数字和英文字母每两个交替打印
举例:字符串 0123456789ABCDefghi
预期输出结果:[A B 1 2 C D 3 4 e f 5 6]
*/
func main() {
	var s = "12345678ABCDefgh"
	var (
		resNumber []rune
		resWord   []rune
		ch        = make(chan struct{}, 1)
		res       []string
		wg        sync.WaitGroup
	)
	wg.Add(2)
	for _, v := range s {
		if v >= 48 && v <= 57 {
			resNumber = append(resNumber, v)
		}
		if (v >= 64 && v <= 90) || (v >= 97 && v <= 122) {
			resWord = append(resWord, v)
		}
	}

	go func() { //int
		defer wg.Done()
		for i := 0; i <= len(resNumber)-3; i = i + 2 {
			_ = <-ch
			res = append(res, fmt.Sprintf("%c", resNumber[i]), fmt.Sprintf("%c", resNumber[i+1]))
			ch <- struct{}{}
		}
	}()

	go func() { //string
		defer wg.Done()
		for i := 0; i <= len(resWord)-3; i = i + 2 {
			_ = <-ch
			res = append(res, fmt.Sprintf("%c", resWord[i]), fmt.Sprintf("%c", resWord[i+1]))
			ch <- struct{}{}
		}
	}()
	ch <- struct{}{}
	wg.Wait()
	fmt.Println(res)
}
