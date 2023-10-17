package minor

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// f(n)
// 1 nil
// 12-> 1 2 12 nil
// 123->1 2 3 13 23 123 3 nil
// f(n+1)=
func subsetsAnother(nums []int) (res []int) {
	res = append(res, nums[:1]...)
	var f func(box *[]int, nextTarget int) (res *[]int)
	f = func(box *[]int, nextTarget int) (res *[]int) {
		for _, v := range *box {
			*res = append(*res, 10*v+nextTarget)
		}
		fmt.Println(res)
		return
	}

	for i := 1; i <= len(nums)-1; i++ {
		res = *f(&res, nums[i])
	}

	fmt.Println("res-----", res)
	return
}
func subsets(nums []int) (res [][]int) {
	//n := len(nums)
	var f = func(nums []int, target int) bool {
		targetStr := strings.Split(strconv.Itoa(target), "")
		//fmt.Println(target)
		//fmt.Println(targetStr)
		var box = make([]int, 0, len(targetStr))
		for _, v := range targetStr {
			if val, err := strconv.Atoi(v); err == nil {
				box = append(box, val)
			}
		}
		for _, valBox := range box {
			for _, valNums := range nums {
				if valBox != valNums {
					return false
				}
			}
		}
		return true
	}
	for i := 0; i <= 4321; i++ {
		if f(nums, i) {
			fmt.Println(i)
		}
	}
	return nil
}

func TestSubsets(t *testing.T) {
	//fmt.Println(subsets([]int{1, 2, 3, 4}))
	fmt.Println(subsetsAnother([]int{1, 2, 3, 4}))
}
