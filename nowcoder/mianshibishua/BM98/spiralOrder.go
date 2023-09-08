package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param matrix int整型二维数组
 * @return int整型一维数组
 */
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0]) // 题目保证每行的列数固定

	// 进行螺旋时的各个边界
	upEdge := -1
	downEdge := m
	leftEdge := -1
	rightEdge := n

	res := make([]int, 0)

	// 总共n个数
	i, j := 0, -1 // i-行坐标, j-列坐标
	count := 0
	for {
		// 先往右走
		j++
		for ; j < rightEdge; j++ {
			res = append(res, matrix[i][j])
			count++
			if count >= m*n {
				return res
			}
		}
		j--      // 越界才判断跳出循环，所以需要回退
		upEdge++ // 缩小上边界

		//	再往下走
		i++ // (i, j) 位置已经获取了，需要往下先走一步
		for ; i < downEdge; i++ {
			res = append(res, matrix[i][j])
			count++
			if count >= m*n {
				return res
			}
		}
		i--
		rightEdge--

		//	再往左走
		j--
		for ; j > leftEdge; j-- {
			res = append(res, matrix[i][j])
			count++
			if count >= m*n {
				return res
			}
		}
		j++
		downEdge--

		//	再往上走
		i--
		for ; i > upEdge; i-- {
			res = append(res, matrix[i][j])
			count++
			if count >= m*n {
				return res
			}
		}
		i++
		leftEdge++
	}
}
