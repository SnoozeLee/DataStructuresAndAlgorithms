package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param numbers int整型一维数组
 * @return int整型
 */
func MoreThanHalfNum_Solution(numbers []int) int {
	// write code here
	m := make(map[int]int, len(numbers)/2+1)
	var res int
	for _, n := range numbers {
		if _, ok := m[n]; ok {
			m[n] = m[n] + 1
		} else {
			m[n] = 1
		}

		if m[n] > len(numbers)/2 {
			res = n
			break
		}
	}
	return res
}
