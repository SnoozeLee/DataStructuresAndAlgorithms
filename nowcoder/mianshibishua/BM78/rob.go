package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param nums int整型一维数组
 * @return int整型
 */
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}
	// 此时至少三家
	bestSums := make([]int, len(nums))
	bestSums[0] = nums[0]
	bestSums[1] = nums[1]
	bestSums[2] = nums[2] + nums[0] // 要求nums[0] 非负，不然还要多加判断

	for i := 3; i < len(nums); i++ {
		// 这里的代码要求题目保证nums[i] 非负。如果nums[i]可能是负数的话，则子问题不只两个
		if bestSums[i-2] > bestSums[i-3] {
			bestSums[i] = bestSums[i-2] + nums[i]
		} else {
			bestSums[i] = bestSums[i-3] + nums[i]
		}
	}

	if bestSums[len(nums)-1] > bestSums[len(nums)-2] {
		return bestSums[len(nums)-1]
	} else {
		return bestSums[len(nums)-2]
	}
}
