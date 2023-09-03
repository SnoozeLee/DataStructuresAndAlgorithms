package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{1, 2, 3, 4, 5, 6, 7, 8}
	c := make([]int, 0)
	fmt.Println(solve(6, 2, a))
	fmt.Println(solve(6, 0, a))
	fmt.Println(solve(8, 4, b))
	fmt.Println(solve(8, 8, b))
	fmt.Println(solve(8, 12, b))
	fmt.Println(solve(0, 0, c))
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 旋转数组
 * @param n int整型 数组长度
 * @param m int整型 右移距离
 * @param a int整型一维数组 给定数组
 * @return int整型一维数组
 */
func solve(n int, m int, a []int) []int {
	if n != 0 && m >= n {
		m = m % n
	}
	count := 0
	startIdx := 0
	for count < n {
		// 可能存在循环组合，所以每次可能需要测试多组，然后通过统计操作个数的方式来判断跳出循环
		idx := startIdx
		tBefore := a[nextIdx(startIdx, n, -m)]
		var tNow int
		for {
			tNow = a[idx]
			a[idx] = tBefore
			tBefore = tNow
			idx = nextIdx(idx, n, m)
			// 完成一个数的位移，总体统计+1
			count++

			// 遇到当前组的第一个数了，说明已经操作完了
			if idx == startIdx {
				break
			}
		}
		startIdx++
	}

	return a
}

func nextIdx(idx, arrLen, offset int) int {
	idx += offset
	// 右移溢出
	if idx >= arrLen {
		return idx - arrLen
	}
	// 左移溢出
	if idx < 0 {
		return idx + arrLen
	}
	return idx
}
