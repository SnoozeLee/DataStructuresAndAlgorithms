package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param m int整型
 * @param n int整型
 * @return int整型
 */
func uniquePaths(m int, n int) int {
	// 为了是rectRow最小，所以n应该要小。由于m，n对称，所以选min(m,n)作为n
	if m < n {
		t := m
		m = n
		n = t
	}

	//	m行n列
	//  i行j列

	// 由于下一行只依赖于上一行，所以只要一行就能求解该动态规划问题
	rectRow := make([]int, n) // 一行中的n个数。
	for j := 0; j < n; j++ {
		rectRow[j] = 1
	}
	// 遍历每一行，第0行已经求解了，所以从第1行起进行运算
	for i := 1; i < m; i++ {
		// rectRow[0] 永远是1
		for j := 1; j < n; j++ {
			rectRow[j] = rectRow[j] + rectRow[j-1]
		}
	}

	return rectRow[n-1]
}
