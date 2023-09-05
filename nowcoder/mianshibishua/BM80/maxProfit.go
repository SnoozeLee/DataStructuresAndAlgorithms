package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param prices int整型一维数组
 * @return int整型
 */
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	// 动态规划解法：
	// 从第0天到今天为止的最高收益
	// 子问题拆解成：
	// 从第0天到昨天时（不一定是昨天卖出）的最高收 vs 以今天作为卖出日的最高收益
	min := prices[0]
	dp := make([]int, len(prices))
	dp[0] = 0
	for i := 1; i < len(prices); i++ {
		dp[i] = max(dp[i-1], prices[i]-min)

		// 维护min
		if prices[i] < min {
			min = prices[i]
		}
	}
	return dp[len(prices)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
