package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param matrix int整型二维数组 the matrix
 * @return int整型
 */
func minPathSum(matrix [][]int) int {
	// m行n列
	m := len(matrix)
	n := len(matrix[0])

	// 如果m<n，则用一列来保存状态
	if m < n {
		arr := make([]int, m)
		// 初始化首列
		arr[0] = matrix[0][0]
		for i := 1; i < m; i++ {
			arr[i] = matrix[i][0] + arr[i-1]
		}
		// 滚动各列
		for j := 1; j < n; j++ {
			arr[0] = matrix[0][j] + arr[0]
			for i := 1; i < m; i++ {
				if arr[i-1] < arr[i] {
					arr[i] = arr[i-1] + matrix[i][j]
				} else {
					arr[i] = arr[i] + matrix[i][j]
				}
			}
		}

		return arr[m-1]
	}

	// 如果是m>=n，则用一行来保存状态
	arr := make([]int, n)
	// 初始化首行
	arr[0] = matrix[0][0]
	for j := 1; j < n; j++ {
		arr[j] = arr[j-1] + matrix[0][j]
	}

	for i := 1; i < m; i++ {
		arr[0] = arr[0] + matrix[i][0]
		for j := 1; j < n; j++ {
			if arr[j-1] < arr[j] {
				arr[j] = arr[j-1] + matrix[i][j]
			} else {
				arr[j] = arr[j] + matrix[i][j]
			}
		}
	}

	return arr[n-1]
}
