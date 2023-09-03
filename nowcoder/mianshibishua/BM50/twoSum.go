package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param numbers int整型一维数组
 * @param target int整型
 * @return int整型一维数组
 */
func twoSum(numbers []int, target int) []int {
	m := make(map[int]int, len(numbers))
	for i := 0; i < len(numbers); i++ {
		if v, ok := m[target-numbers[i]]; ok {
			return []int{v + 1, i + 1} // 题目要求的下标是从1计数的
		} else {
			m[numbers[i]] = i
		}
	}
	// 这里不会被触发
	return nil
}
