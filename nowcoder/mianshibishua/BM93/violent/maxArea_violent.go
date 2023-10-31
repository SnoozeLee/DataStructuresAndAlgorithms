package violent

// 暴力解法 - 运行超时

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param height int整型一维数组
 * @return int整型
 */
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	max := 0

	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			area := min(height[i], height[j]) * (j - i)
			if area > max {
				max = area
			}
		}
	}
	return max
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
