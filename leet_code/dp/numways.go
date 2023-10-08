package main

import "fmt"

// Solution 青蛙跳台阶
func Solution(nums int) int {
	// f(n)=f(n-1)+f(n-2)
	if nums == 1 {
		return 1
	}
	if nums == 2 {
		return 2
	}
	return Solution(nums-1) + Solution(nums-2)

}

func main() {
	fmt.Println(Solution(5))
}
