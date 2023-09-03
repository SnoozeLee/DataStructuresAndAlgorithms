package main

import "math"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param pond int整型一维数组
 * @return int整型
 */
func minJump(pond []int) int {
	n := len(pond)
	if len(pond) == 1 {
		return 0
	}

	steps := make([]int, n) // 到达某一节点所需的跳跃数
	steps[0] = 0
	for i := 1; i < n; i++ {
		steps[i] = math.MaxInt // 无穷远
	}

	// 要求到达pond[n-1]且题目保证终点一定可到达，意味着计算到pond[n-2]就已经得出答案了
	for i := 0; i < n-1; i++ {
		for j := 1; j <= pond[i] && i+j < n; j++ {
			if steps[i]+1 < steps[i+j] {
				steps[i+j] = steps[i] + 1
			}
		}
	}
	return steps[n-1]
}
