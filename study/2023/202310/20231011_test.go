package _02310

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func Test20231011(t *testing.T) {
	d := struct {
		time.Time
		N int
	}{
		Time: time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
		N:    5,
	}

	m, _ := json.Marshal(d)
	fmt.Printf("%s", m)

}

// A：{"Time": "2020-12-20T00:00:00Z", "N": 5 }；B："2020-12-20T00:00:00Z"；C：{"N": 5}；D：<nil>
/*
time.Time 是内嵌的
通过内嵌来模拟继承的功能,t 的方法集包括了 time.Time 的方法集，所以，t 自动实现了 Marshaler 接口。因此答案是 B
*/
