package main

// 此处使用 暴力枚举， 从时间复杂度角度上看应该不可行的
// 不过判题通过了

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
	max := 0
	// 暴力解法 i 为购入日， j 为出售日
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}
