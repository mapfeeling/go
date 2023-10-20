package main

import (
	"fmt"
	"testing"
)

/*
同积元组
给你一个由 不同 正整数组成的数组 nums
请你返回满足 a * b = c * d 的元组 (a, b, c, d) 的数量
其中 a、b、c 和 d 都是 nums 中的元素，且 a != b != c != d
*/

/*
示例 1：
输入：nums = [2,3,4,6]
输出：8
解释：存在 8 个满足题意的元组：
(2,6,3,4) , (2,6,4,3) , (6,2,3,4) , (6,2,4,3)
(3,4,2,6) , (4,3,2,6) , (3,4,6,2) , (4,3,6,2)

示例 2：
输入：nums = [1,2,4,5,10]
输出：16
解释：存在 16 个满足题意的元组：
(1,10,2,5) , (1,10,5,2) , (10,1,2,5) , (10,1,5,2)
(2,5,1,10) , (2,5,10,1) , (5,2,1,10) , (5,2,10,1)
(2,10,4,5) , (2,10,5,4) , (10,2,4,5) , (10,2,4,5)
(4,5,2,10) , (4,5,10,2) , (5,4,2,10) , (5,4,10,2)
*/

func tupleSameProduct(nums []int) int {
	var (
		ret int
		box = make(map[int]int)
	)

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			box[nums[i]*nums[j]]++
		}
	}

	for _, v := range box {
		ret += v * (v - 1) * 4
	}

	return ret
}

func TestTupleSameProduct(t *testing.T) {
	var nums = []int{1, 2, 4, 5, 10}
	var nums2 = []int{2, 3, 4, 6, 8, 12}
	var nums3 = []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192}
	fmt.Println(tupleSameProduct(nums))
	fmt.Println(tupleSameProduct(nums2))
	fmt.Println(tupleSameProduct(nums3))
}

/*
根据排列组合可以知道，如果存在 cnt(a×b)\textit{cnt}(a \times b)cnt(a×b) 个乘积相同的数对，则此时可以有 Ccnt(a×b)2C_{\textit{cnt}(a \times b)}^2C
cnt(a×b)2​种不同的数对组合
即此时可以构成的同积元组的数目为: cnt(a×b)×(cnt(a×b)−1)2×8=cnt(a×b)×(cnt(a×b)−1)×4\dfrac{\textit{cnt}(a \times b) \times (\textit{cnt}(a \times b) -1)}{2} \times 8 = {\textit{cnt}(a \times b) \times (\textit{cnt}(a \times b) -1)} \times 4
2
cnt(a×b)×(cnt(a×b)−1)
​
 ×8=cnt(a×b)×(cnt(a×b)−1)×4
*/
