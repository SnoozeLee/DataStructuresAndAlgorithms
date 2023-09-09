package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 最少货币数
 * @param arr int整型一维数组 the array
 * @param aim int整型 the target
 * @return int整型
 */
func minMoney(arr []int, aim int) int {
	if aim == 0 {
		return 0
	}
	// aims 到达某个目标所需要的最少钞票数量
	aims := make([]int, aim+1) // aims长度至少为aim+1时才有下标aim
	for i := 1; i < len(aims); i++ {
		aims[i] = -1 // 用-1表示不可达
	}
	for i := 0; i < len(aims); i++ {
		// 首先要只操作可到达的点
		if aims[i] >= 0 {
			for j := 0; j < len(arr); j++ {
				// 下一个点在边界内且(下一个点的票数大于当前数+1或下一个数是不可达)
				if i+arr[j] < len(aims) && (aims[i]+1 < aims[i+arr[j]] || aims[i+arr[j]] == -1) {
					aims[i+arr[j]] = aims[i] + 1
				}
			}
		}
	}
	return aims[aim]
}
