package main

// 这里的解法是我以为只能选两头或以下的牛来巡逻，但实际应该是随意选
// 所以以下的解法只适用与选择两个不相邻数的最大加和
// 不能用来解此题

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	arr2 := []int{1, 2, 4, 1}
	fmt.Println(maxPatrolValue(arr))
	fmt.Println(maxPatrolValue(arr2))
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param values int整型一维数组
 * @return int整型
 */
func maxPatrolValue(values []int) int {
	maxValue := values[0]
	if len(values) == 1 {
		return 1
	} else if len(values) == 2 {
		return max(values[0], values[1])
	}

	// 记录当前坐标下，左边（含本身）能获取到的最大值，
	leftMaxValue := make([]int, len(values))
	leftMaxValue[0] = maxValue
	for i := 1; i < len(values); i++ {
		if values[i] > maxValue {
			maxValue = values[i]
		}
		leftMaxValue[i] = maxValue
	}

	// 生成以某个数作为结尾时，两个数的最大值
	maxList := make([]int, len(values))
	maxList[0] = leftMaxValue[0]
	maxList[1] = leftMaxValue[1]
	for i := 2; i < len(values); i++ {
		maxList[i] = maxList[i-2] + values[i]
	}

	maxRes := maxList[0]
	for _, v := range maxList {
		if maxRes < v {
			maxRes = v
		}
	}
	return maxRes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
