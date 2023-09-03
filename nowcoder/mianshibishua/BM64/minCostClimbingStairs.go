package main

import "fmt"

func main() {
	cost := []int{2, 5, 20}
	fmt.Println(minCostClimbingStairs(cost))
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param cost int整型一维数组
 * @return int整型
 */
func minCostClimbingStairs(cost []int) int {
	step := len(cost)

	return minCost(cost, step)
}

func minCost(cost []int, step int) int {
	if step == 0 || step == 1 {
		return 0
	}
	costBefore2 := 0
	costBefore1 := 0
	var costToStep int
	for i := 2; i <= step; i++ {
		costToStep = min(costBefore1+cost[i-1], costBefore2+cost[i-2])
		costBefore2 = costBefore1
		costBefore1 = costToStep
	}
	return costToStep
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
