package _02312

import (
	"fmt"
	"testing"
	"unsafe"
)

func minimumFuelCost(roads [][]int, seats int) int64 {
	gas := make([][]int, len(roads)+1)
	for _, road := range roads {
		gas[road[0]] = append(gas[road[0]], road[1])
		gas[road[1]] = append(gas[road[1]], road[0])
	}
	var (
		res int64
		dfs func(int, int) int
	)
	dfs = func(i int, j int) int {
		pSum := 1
		for _, val := range gas[i] {
			if val != j {
				pCount := dfs(val, i)
				pSum, res = pSum+pCount, res+int64((pCount+seats-1)/seats)

			}
		}
		return pSum
	}
	dfs(0, -1)
	return res
}

type Date struct {
	Home *Home
}

type Home struct {
	Kids *Kid
}

type Kid struct {
	Name string
}

func a(s []int) {

	s = append(s, 1)

}

func TestMinimumFuelCost(t *testing.T) {
	var ss = make([]int, 1, 1)
	fmt.Println(unsafe.Pointer(&ss))
	a(ss)
	fmt.Println(unsafe.Pointer(&ss))
	fmt.Println(len(ss))

	fmt.Println(minimumFuelCost([][]int{{0, 1}, {0, 2}, {0, 3}}, 5))
	var d = &Date{
		Home: &Home{
			Kids: &Kid{
				Name: "123",
			},
		},
	}
	d.Home.Kids.Name = "234"
	var q = &Date{}
	fmt.Println(q.Home)
	q.Home = &Home{
		Kids: &Kid{
			Name: "1",
		},
	}
	fmt.Println(q)
	fmt.Println(d.Home.Kids.Name)
	fmt.Println(".....")
	fmt.Println(1 & 7) //111 010

	var ret []int64
	// 选集起播 			 16777217
	// 全屏起播 			 16777218
	// 播起起播 			 16777220
	qq := int64(2 << (24 - 1))
	for _, option := range []int64{1, 2, 4} {
		if option&7 == option {
			ret = append(ret, qq+option)
		}
	}
	fmt.Println(ret)
	fmt.Println(2 << 3)
	var k string
	var a = &A{B: &BStruct{
		Name: k,
	}}
	fmt.Println(a)
}

type A struct {
	B *BStruct
}

type BStruct struct {
	Name string `json:"name"`
}
