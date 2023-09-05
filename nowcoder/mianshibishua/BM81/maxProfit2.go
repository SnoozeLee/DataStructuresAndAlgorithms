package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 计算最大收益
 * @param prices int整型一维数组 股票每一天的价格
 * @return int整型
 */
func maxProfit(prices []int) int {
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		// 只要第二天是涨的，今天就持有(可能已经持有了)。只要第二天跌了，今天就售出(可能已经售出了)
		if prices[i+1] > prices[i] {
			profit += prices[i+1] - prices[i]
		}
	}
	return profit
}
