package main

import (
	"fmt"
	"math"
	"testing"
)

func categorizeBox(length int, width int, height int, mass int) string {
	var res []string
	nums104 := int(math.Pow10(4))
	nums109 := int(math.Pow10(9))
	if length >= nums104 || width >= nums104 || height >= nums104 || length*width*height >= nums109 {
		res = append(res, "Bulky")
	}
	if mass >= 100 {
		res = append(res, "Heavy")
	}
	if len(res) == 2 {
		return "Both"
	}
	if len(res) == 1 {
		if res[0] == "Heavy" {
			return "Heavy"
		}
		if res[0] == "Bulky" {
			return "Bulky"
		}
	}
	return "Neither"
}

func TestCategorizeBox(t *testing.T) {
	fmt.Println(categorizeBox(1000, 35, 700, 300))
}

/*
给你四个整数 length ，width ，height 和 mass ，分别表示一个箱子的三个维度和质量，请你返回一个表示箱子 类别 的字符串。
如果满足以下条件，那么箱子是 "Bulky" 的：
箱子 至少有一个 维度大于等于 10^4 。
或者箱子的 体积 大于等于 10^9 。
如果箱子的质量大于等于 100 ，那么箱子是 "Heavy" 的。
如果箱子同时是 "Bulky" 和 "Heavy" ，那么返回类别为 "Both" 。
如果箱子既不是 "Bulky" ，也不是 "Heavy" ，那么返回类别为 "Neither" 。
如果箱子是 "Bulky" 但不是 "Heavy" ，那么返回类别为 "Bulky" 。
如果箱子是 "Heavy" 但不是 "Bulky" ，那么返回类别为 "Heavy" 。
注意，箱子的体积等于箱子的长度、宽度和高度的乘积
*/
