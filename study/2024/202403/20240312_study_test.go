package _02403

import (
	"sync"
	"testing"
)

// 问以下代码输出什么？
// 即case中每个case都会先计算值
func Test_20240312_study(t *testing.T) {
	var wg sync.WaitGroup
	foo := make(chan int)
	bar := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case foo <- <-bar:
		default:
			println("default")
		}
	}()
	wg.Wait()
}

// 答案是panic,发生死锁
/*
Explain
package main

import (
	"fmt"
	"time"
)

func talk(msg string, sleep int) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(sleep) * time.Millisecond)
		}
	}()
	return ch
}

func fanIn(input1, input2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case ch <- <-input1:
			case ch <- <-input2:
			}
		}
	}()
	return ch
}

func main() {
	ch := fanIn(talk("A", 10), talk("B", 1000))
	for i := 0; i < 10; i++ {
		fmt.Printf("%q\n", <-ch)
	}
}
*/

/*
select {
case ch <- <-input1:
case ch <- <-input2:
}
*/

/*

// 改为这样就一切正常
select {
case t := <-input1:
  ch <- t
case t := <-input2:
  ch <- t
}
*/

// 又比如
/*
/ ch 是一个 chan int；
// getVal() 返回 int
// input 是 chan int
// getch() 返回 chan int
select {
  case ch <- getVal():
  case ch <- <-input:
  case getch() <- 1:
  case <- getch():
}
*/
