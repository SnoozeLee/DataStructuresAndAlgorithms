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
	if len(values) == 1 {
		return values[0]
	}
	if len(values) == 2 {
		return max(values[0], values[1])
	}

	// left[i] 以当前数value[i]作为末尾数字时的最大值
	left := make([]int, len(values))
	left[0] = values[0]
	left[1] = values[1]
	left[2] = values[2] + values[0]
	// 因为非负，所以加上values[0]肯定比单独选一个values[2]大

	for i := 3; i < len(values); i++ {
		if left[i-2] > left[i-3] {
			left[i] = values[i] + left[i-2]
		} else {
			left[i] = values[i] + left[i-3]
		}
	}

	length := len(values)
	return max(left[length-1], left[length-2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
