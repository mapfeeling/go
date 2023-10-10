package _go

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func TestDemo(t *testing.T) {
	// 生产数据 go
	// 消费数据 go
	// channel buf
	var (
		ch = make(chan int, 100)
		wg = sync.WaitGroup{}
	)

	wg.Add(2)
	defer wg.Wait()

	go func() {
		defer wg.Done()
		fmt.Println("我是生产者")
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		//close(ch)
	}()
	go func() {
		// 模拟
		defer wg.Done()
		fmt.Println("我是消费者")
		for {
			if v, ok := <-ch; ok {
				demoOther(v)
			} //} else {
			//	break
			//}
		}

	}()
}

func demo(i int) error {
	fmt.Println(i)
	var (
		url      = "http://www.baidu.com"
		urlParam = []string{url, "?=", strconv.Itoa(i)}
	)

	urlParamOther := strings.Join(urlParam, "")
	postRes, err := http.Get(urlParamOther)
	if err != nil {
		fmt.Println("err:", err)
		return err
	}
	fmt.Println(postRes)
	return nil
}

func demoOther(i int) {
	fmt.Println(i)

}
