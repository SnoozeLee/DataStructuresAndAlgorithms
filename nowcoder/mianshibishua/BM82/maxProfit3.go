package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 两次交易所能获得的最大收益
 * @param prices int整型一维数组 股票每一天的价格
 * @return int整型
 */
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// 以当天分别作为 第1次购入、第1次售出、第2次购入、第2次售出 的各个结果
	// 与 之前已经判断出的各个结果 一对一比较，选出当前已经执行了
	// 第1次购入、第1次售出、第2次购入、第2次售出 的最优结果 （利润的数量）
	buy1Res, sell1Res, buy2Res, sell2Res := -prices[0], 0, -prices[0], 0
	//	注意求解顺序必须是从buy1到sell1到buy2到sell2，因为sell1必须引用当前最好的buy1状态，以此类推
	//	题目允许当天买当天卖再当天买当天卖，所以buy2Res初始值可以设置为 - prices
	for i := 1; i < len(prices); i++ {
		buy1Res = max(buy1Res, -prices[i])          // 购入时花钱，所以 prices[i] 记录正值，表示利润亏损
		sell1Res = max(sell1Res, buy1Res+prices[i]) // 售出时赚钱，所以 prices[i] 记录正值，表示利润增长
		// sell1Res 解释： 今天售出的钱 + 最优第一次购入的钱的状态（负值）
		// 即为假设当天售出时的利润，再用这个利润和昨天以及以前的最优解比较

		buy2Res = max(buy2Res, sell1Res-prices[i])
		sell2Res = max(sell2Res, buy2Res+prices[i])
	}
	return sell2Res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
