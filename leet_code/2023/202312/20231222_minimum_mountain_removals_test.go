package _02312

import (
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
)

type AA struct {
	Data int `json:"data"`
	Box  int `json:"box,omitempty"`
}

func TestAb(t *testing.T) {
	//var sm = make([]map[interface{}]interface{}, 0, 8)
	//for i := 0; i < 8; i++ {
	//	sm = append(sm, map[interface{}]interface{}{
	//		i:              i,
	//		"movie":        i,
	//		"content_type": i,
	//	})
	//}
	//fmt.Println(sm)
	//for _, m := range sm {
	//	if movieId, _ := m["movie"].(int); movieId > 0 {
	//		m["content_type"] = 13
	//	}
	//}
	//fmt.Println(sm)

	var a = AA{
		//Data: 0,
		Box: 0,
	}
	marshal, err := sonic.Marshal(a)
	if err != nil {
		return
	}
	fmt.Println("--->", string(marshal))
	fmt.Println(a)
}
