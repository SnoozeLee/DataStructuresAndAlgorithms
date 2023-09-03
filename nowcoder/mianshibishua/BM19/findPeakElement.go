package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param nums int整型一维数组
 * @return int整型
 */
func findPeakElement(nums []int) int {
	// write code here
	L := 0
	R := len(nums) - 1
	for {
		if R == L {
			return L
		}
		if R == L+1 {
			if nums[R] > nums[L] {
				return R
			}
			return L
		}
		mid := L + ((R - L) >> 1)
		if nums[mid] > nums[mid-1] {
			//	峰值在mid（含本身）到R之间
			L = mid
		} else {
			//	峰值在L到mid-1之间
			R = mid - 1
		}
	}
}

func main() {
	fmt.Println(findPeakElement([]int{2, 4, 1, 2, 7, 8, 4}))
}
