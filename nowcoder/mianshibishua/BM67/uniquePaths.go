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
	var rect = make([][]int, m)
	// 假定m为行数，n为列数
	for i := 0; i < m; i++ {
		rect[i] = make([]int, n)
		rect[i][0] = 1 // 第0列置一
	}
	// 第0行置一
	for j := 1; j < n; j++ {
		rect[0][j] = 1
	}
	// 从第1行第1列起动态规划
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			rect[i][j] = rect[i][j-1] + rect[i-1][j]
		}
	}
	return rect[m-1][n-1]
}
