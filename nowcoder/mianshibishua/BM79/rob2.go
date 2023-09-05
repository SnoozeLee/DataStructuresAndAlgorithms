package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param nums int整型一维数组
 * @return int整型
 */
func rob(nums []int) int {
	res1 := robLtoR(nums, 0, len(nums)-2)
	res2 := robLtoR(nums, 1, len(nums)-1)
	if res1 > res2 {
		return res1
	}
	return res2
}

// 从L(包含)开始打劫,到R(包含)之前,不考虑环形
func robLtoR(nums []int, L, R int) int {
	if R-L < 0 {
		return 0
	}
	if R-L == 0 {
		return nums[L]
	}
	if R-L == 1 {
		if nums[L] > nums[L+1] {
			return nums[L]
		}
		return nums[L+1]
	}
	// 此时必有三家，且第三家的最佳方案一定是第一和第三家一起

	befores := [3]int{nums[L], nums[L+1], nums[L] + nums[L+2]}

	for i := L + 3; i <= R; i++ {
		var res int
		if befores[0] > befores[1] {
			res = befores[0] + nums[i]
		} else {
			res = befores[1] + nums[i]
		}

		befores[0] = befores[1]
		befores[1] = befores[2]
		befores[2] = res
	}

	if befores[1] > befores[2] {
		return befores[1]
	} else {
		return befores[2]
	}

}
