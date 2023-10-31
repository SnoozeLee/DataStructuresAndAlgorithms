package main

// 双指针解法
// 时间复杂度O(n)
// 空间复杂度O(1)

// 两边往内靠拢，每次靠拢舍弃短边，记录最大值

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

	i := 0
	j := len(height) - 1
	for i < j {
		area := min(height[i], height[j]) * (j - i)
		if area > max {
			max = area
		}
		if height[i] <= height[j] {
			i++
		} else {
			j--
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
