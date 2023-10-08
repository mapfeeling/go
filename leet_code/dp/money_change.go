package main

import "fmt"

/*
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回-1
示例1:
输入: coins = [1, 2, 5], amount = 11
输出: 3
解释: 11 = 5 + 5 + 1
示例 2:
输入: coins = [2], amount = 3
输出: -1
*/

func coinChange(amount int, coins []int) int {
	var dp = make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
		fmt.Println(dp)
		for j := 0; j < len(coins); j++ {
			// 面额不大于当前金额且发现了更少的兑换情况
			if coins[j] <= i && dp[i-coins[j]]+1 < dp[i] {
				dp[i] = dp[i-coins[j]] + 1
			}
		}
	}
	//无法兑换
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func main() {
	var (
		coins  = []int{1, 2, 5}
		amount = 11
	)
	fmt.Println(coinChange(amount, coins))
}
